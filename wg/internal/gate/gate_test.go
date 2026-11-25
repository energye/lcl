package gate

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestMissingProbeCoverage(t *testing.T) {
	cfg := probeConfig{
		Version:         1,
		RequireProbeFor: []string{"wg/widgets/"},
		Probes: []Probe{{
			ID: "feedback", Covers: []string{"wg/widgets/feedback.go"}, Command: []string{"go", "test"},
		}},
	}
	missing := MissingProbeCoverage(cfg, []string{"wg/widgets/feedback.go", "wg/widgets/timepicker.go",
		"wg/widgets/timepicker_test.go", "docs/readme.md"})
	if len(missing) != 1 || missing[0] != "wg/widgets/timepicker.go" {
		t.Fatalf("missing=%v, want [wg/widgets/timepicker.go]", missing)
	}
}

func TestAllPass(t *testing.T) {
	if allPass(nil) {
		t.Fatal("空步骤不能通过")
	}
	if allPass([]Step{{ID: "a", Status: StatusFail}}) {
		t.Fatal("FAIL 不能通过")
	}
	if !allPass([]Step{{ID: "a", Status: StatusPass}, {ID: "b", Status: StatusPass}}) {
		t.Fatal("全 PASS 应通过")
	}
}

func TestCheckRejectsMissingReport(t *testing.T) {
	root := t.TempDir()
	err := Check(root, "tree", Options{ReportDir: "reports", ProbeFile: "probes.json"})
	if err == nil || !strings.Contains(err.Error(), "报告缺失") {
		t.Fatalf("err=%v, want 报告缺失", err)
	}
}

func TestCheckRejectsStaleAndFailedReports(t *testing.T) {
	root := newGitRepo(t)
	tree := gitOutputForTest(t, root, "write-tree")
	opts := Options{ReportDir: "reports", ProbeFile: "probes.json"}

	stale := &Report{Version: ReportVersion, TreeHash: "stale", Status: StatusPass}
	if err := WriteReport(root, opts.ReportDir, stale); err != nil {
		t.Fatal(err)
	}
	staleData, err := os.ReadFile(ReportPath(root, opts.ReportDir, stale.TreeHash))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(ReportPath(root, opts.ReportDir, tree), staleData, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := Check(root, tree, opts); err == nil || !strings.Contains(err.Error(), "staged tree") {
		t.Fatalf("err=%v, want stale tree", err)
	}

	failed := &Report{Version: ReportVersion, TreeHash: tree, Status: StatusFail}
	if err := WriteReport(root, opts.ReportDir, failed); err != nil {
		t.Fatal(err)
	}
	if err := Check(root, tree, opts); err == nil || !strings.Contains(err.Error(), "状态=FAIL") {
		t.Fatalf("err=%v, want FAIL status", err)
	}

	badStep := &Report{
		Version: ReportVersion, TreeHash: tree, Status: StatusPass,
		Steps: []Step{{ID: "build", Status: StatusTimeout, Detail: "timeout"}},
	}
	if err := WriteReport(root, opts.ReportDir, badStep); err != nil {
		t.Fatal(err)
	}
	if err := Check(root, tree, opts); err == nil || !strings.Contains(err.Error(), "build=TIMEOUT") {
		t.Fatalf("err=%v, want TIMEOUT step", err)
	}
}

func newGitRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	gitOutputForTest(t, root, "init")
	if err := os.WriteFile(filepath.Join(root, "file.txt"), []byte("init\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "file.txt")
	gitOutputForTest(t, root, "-c", "user.name=wg-gate", "-c", "user.email=wg-gate@example.invalid", "commit", "-m", "init")
	return root
}

func gitOutputForTest(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func TestEnsureStagedRejectsUnstagedAndUntrackedGateSource(t *testing.T) {
	root := newGitRepo(t)
	if err := os.WriteFile(filepath.Join(root, "file.txt"), []byte("changed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := EnsureStaged(root); err == nil || !strings.Contains(err.Error(), "未暂存改动") {
		t.Fatalf("err=%v, want 未暂存改动", err)
	}
	gitOutputForTest(t, root, "add", "file.txt")
	if err := EnsureStaged(root); err != nil {
		t.Fatalf("已暂存改动应通过: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "wg", "new"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "wg", "new", "main.go"), []byte("package main\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := EnsureStaged(root); err == nil || !strings.Contains(err.Error(), "未暂存门禁源码") {
		t.Fatalf("err=%v, want 未暂存门禁源码", err)
	}
}

func TestTaskOnlyTreeDrift(t *testing.T) {
	root := newGitRepo(t)
	if err := os.MkdirAll(filepath.Join(root, "wg", "tasks", "task"), 0o755); err != nil {
		t.Fatal(err)
	}
	taskPath := filepath.Join(root, "wg", "tasks", "task", "task.json")
	if err := os.WriteFile(taskPath, []byte("{\"state\":\"VERIFIED\"}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "wg")
	before := gitOutputForTest(t, root, "write-tree")
	if err := os.WriteFile(taskPath, []byte("{\"state\":\"CLOSED\"}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "wg")
	after := gitOutputForTest(t, root, "write-tree")
	taskOnly, err := OnlyTaskChanges(root, before, after)
	if err != nil || !taskOnly {
		t.Fatalf("taskOnly=%v err=%v", taskOnly, err)
	}
	if err := os.WriteFile(filepath.Join(root, "file.txt"), []byte("code changed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "file.txt")
	codeChanged := gitOutputForTest(t, root, "write-tree")
	taskOnly, err = OnlyTaskChanges(root, before, codeChanged)
	if err != nil || taskOnly {
		t.Fatalf("taskOnly=%v err=%v, want false", taskOnly, err)
	}
}

func TestCheckAllowsTaskOnlyDrift(t *testing.T) {
	root := newGitRepo(t)
	for i := 0; i < 60; i++ {
		dir := filepath.Join(root, "wg", "examples", fmt.Sprintf("f%02d", i))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := writeJSON(filepath.Join(root, "wg", "gate.probes.json"), probeConfig{Version: 1, RequireProbeFor: []string{"wg/widgets/"}}); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "wg")
	tree := gitOutputForTest(t, root, "write-tree")
	files, err := ChangedFiles(root, tree)
	if err != nil {
		t.Fatal(err)
	}
	report := &Report{
		Version: ReportVersion, TreeHash: tree, Status: StatusPass, ChangedFiles: files,
		Audits: AuditSummary{Expected: 60, Passed: 60},
	}
	for _, spec := range baseSteps() {
		report.Steps = append(report.Steps, Step{ID: spec.ID, Status: StatusPass})
	}
	for _, spec := range diagSteps("wg-diag") {
		report.Steps = append(report.Steps, Step{ID: spec.ID, Status: StatusPass})
	}
	for i := 0; i < 60; i++ {
		report.Steps = append(report.Steps, Step{ID: "audit:f" + fmt.Sprintf("%02d", i), Status: StatusPass})
	}
	report.Steps = append(report.Steps, Step{ID: "probes", Status: StatusPass})
	if err := WriteReport(root, ".shots/gates", report); err != nil {
		t.Fatal(err)
	}
	taskPath := filepath.Join(root, "wg", "tasks", "task", "task.json")
	if err := os.MkdirAll(filepath.Dir(taskPath), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(taskPath, []byte("{\"state\":\"CLOSED\"}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "wg/tasks")
	after := gitOutputForTest(t, root, "write-tree")
	if err := Check(root, after, Options{ReportDir: ".shots/gates", ProbeFile: "wg/gate.probes.json"}); err != nil {
		t.Fatalf("任务证据变化后 check 应通过: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "file.txt"), []byte("code changed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "file.txt")
	codeChanged := gitOutputForTest(t, root, "write-tree")
	if err := Check(root, codeChanged, Options{ReportDir: ".shots/gates", ProbeFile: "wg/gate.probes.json"}); err == nil {
		t.Fatal("代码变化后 check 不应通过")
	}
}

func TestReportRoundTrip(t *testing.T) {
	root := t.TempDir()
	report := &Report{
		Version: ReportVersion, TreeHash: strings.Repeat("a", 40), Status: StatusPass,
		StartedAt: time.Now(), FinishedAt: time.Now(), ChangedFiles: []string{"wg/a.go"},
	}
	if err := WriteReport(root, ".shots/gates", report); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(ReportPath(root, ".shots/gates", report.TreeHash))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), report.TreeHash) || !strings.Contains(string(data), `"status": "PASS"`) {
		t.Fatalf("report 内容异常: %s", data)
	}
}
