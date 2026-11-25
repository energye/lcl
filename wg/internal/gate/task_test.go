package gate

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestTaskAndFeedbackImport(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "wg", "examples", "drawer"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "wg", "examples", "drawer", "spec.json"), []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	task, err := NewTask(root, "drawer-task", "Drawer", "Drawer 任务", []string{"默认宽度 378"})
	if err != nil {
		t.Fatal(err)
	}
	if task.State != TaskNew || len(task.Specs) != 1 || task.Specs[0] != "wg/examples/drawer/spec.json" {
		t.Fatalf("task=%+v", task)
	}

	feedback := Feedback{
		Version: 1, ID: "drawer-feedback", Component: "Drawer", Severity: "P2",
		Steps: []string{"打开抽屉"}, Expected: "宽度 378", Actual: "宽度 320", Reviewer: "reviewer",
	}
	feedbackPath := filepath.Join(root, "feedback.json")
	if err := writeJSON(feedbackPath, feedback); err != nil {
		t.Fatal(err)
	}
	imported, err := ImportFeedback(root, "drawer-feedback-task", feedbackPath)
	if err != nil {
		t.Fatal(err)
	}
	if imported.State != TaskSpecRed || imported.Source.ID != feedback.ID {
		t.Fatalf("imported=%+v", imported)
	}
	for _, path := range []string{
		filepath.Join(TaskDir(root, imported.ID), "feedback.json"),
		filepath.Join(root, imported.Evidence.Regression),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("缺少 %s: %v", path, err)
		}
	}
}

func TestRegisterProbe(t *testing.T) {
	root := t.TempDir()
	cfg := probeConfig{Version: 1, RequireProbeFor: []string{"wg/widgets/"}}
	if err := writeJSON(filepath.Join(root, "wg", "gate.probes.json"), cfg); err != nil {
		t.Fatal(err)
	}
	if _, err := NewTask(root, "probe-task", "Drawer", "探针任务", []string{"宽度 378"}); err != nil {
		t.Fatal(err)
	}
	task, probe, err := RegisterProbe(root, "wg/gate.probes.json", "probe-task", "drawer-width", []string{"wg/widgets/feedback.go"}, []string{"git", "--version"}, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if probe.ID != "drawer-width" || len(task.Probes) != 1 || task.Probes[0] != probe.ID {
		t.Fatalf("task=%+v probe=%+v", task, probe)
	}
	loaded, err := LoadProbeConfig(root, "wg/gate.probes.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Probes) != 1 || loaded.Probes[0].ID != probe.ID {
		t.Fatalf("cfg=%+v", loaded)
	}
	if _, err := os.Stat(filepath.Join(TaskDir(root, task.ID), "probe.json")); err != nil {
		t.Fatal(err)
	}
}

func TestVerifyTaskRunsRegisteredProbe(t *testing.T) {
	root := newGitRepo(t)
	if err := os.MkdirAll(filepath.Join(root, "wg", "examples", "drawer"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "wg", "examples", "drawer", "spec.json"), []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := writeJSON(filepath.Join(root, "wg", "gate.probes.json"), probeConfig{Version: 1, RequireProbeFor: []string{"wg/widgets/"}}); err != nil {
		t.Fatal(err)
	}
	if _, err := NewTask(root, "verify-task", "Drawer", "验证任务", []string{"宽度 378"}); err != nil {
		t.Fatal(err)
	}
	if _, _, err := RegisterProbe(root, "wg/gate.probes.json", "verify-task", "verify-probe", []string{"wg/widgets/feedback.go"}, []string{"git", "--version"}, time.Minute); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "wg")
	task, err := VerifyTask(root, "verify-task", "wg/gate.probes.json")
	if err != nil {
		t.Fatal(err)
	}
	if task.State != TaskVerified || task.Evidence.VerifiedTree == "" || task.Evidence.ProbeReport == "" {
		t.Fatalf("task=%+v", task)
	}
}

func TestCloseTaskWithChecker(t *testing.T) {
	t.Run("PASS closes", func(t *testing.T) {
		root, taskID := verifiedTaskForClose(t)
		if err := CloseTaskWithChecker(root, taskID, ".shots/gates", "wg/gate.probes.json", func(string, string) error { return nil }); err != nil {
			t.Fatal(err)
		}
		task, err := LoadTask(root, taskID)
		if err != nil {
			t.Fatal(err)
		}
		if task.State != TaskClosed || task.ClosedAt == nil {
			t.Fatalf("task=%+v", task)
		}
	})

	t.Run("FAIL blocks", func(t *testing.T) {
		root, taskID := verifiedTaskForClose(t)
		want := errors.New("full gate failed")
		err := CloseTaskWithChecker(root, taskID, ".shots/gates", "wg/gate.probes.json", func(string, string) error { return want })
		if err == nil || !strings.Contains(err.Error(), "full gate failed") {
			t.Fatalf("err=%v", err)
		}
		task, loadErr := LoadTask(root, taskID)
		if loadErr != nil {
			t.Fatal(loadErr)
		}
		if task.State != TaskBlocked {
			t.Fatalf("task=%+v", task)
		}
	})

	t.Run("code change requires reverify", func(t *testing.T) {
		root, taskID := verifiedTaskForClose(t)
		if err := os.WriteFile(filepath.Join(root, "file.txt"), []byte("changed\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		gitOutputForTest(t, root, "add", "file.txt")
		called := false
		err := CloseTaskWithChecker(root, taskID, ".shots/gates", "wg/gate.probes.json", func(string, string) error {
			called = true
			return nil
		})
		if err == nil || !strings.Contains(err.Error(), "代码文件已变化") || called {
			t.Fatalf("err=%v called=%v", err, called)
		}
	})
}

func verifiedTaskForClose(t *testing.T) (string, string) {
	t.Helper()
	root := newGitRepo(t)
	const taskID = "close-task"
	now := time.Now()
	task := Task{
		Version: 1, ID: taskID, Title: "关闭任务", Component: "Drawer", State: TaskVerified,
		Source: TaskSource{Type: "plan"}, Requirements: []string{"宽度 378"},
		CreatedAt: now, UpdatedAt: now,
	}
	if err := SaveTask(root, task); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "wg")
	tree := gitOutputForTest(t, root, "write-tree")
	task.Evidence.VerifiedTree = tree
	if err := SaveTask(root, task); err != nil {
		t.Fatal(err)
	}
	gitOutputForTest(t, root, "add", "wg")
	return root, taskID
}

func TestReopenTask(t *testing.T) {
	root := t.TempDir()
	now := time.Now()
	closedAt := now
	task := Task{
		Version: 1, ID: "reopen-task", Title: "重开任务", Component: "Drawer", State: TaskClosed,
		Source: TaskSource{Type: "plan"}, Requirements: []string{"宽度 378"},
		Specs: []string{"wg/examples/drawer/spec.json"}, Probes: []string{"p1"},
		Evidence: TaskEvidence{
			Regression:   "wg/tasks/reopen-task/regression.json",
			ProbeReport:  "wg/tasks/reopen-task/probe-report.json",
			GateReport:   ".shots/gates/report-abc.json",
			VerifiedTree: "abc123",
		},
		CreatedAt: now, UpdatedAt: now, ClosedAt: &closedAt,
	}
	if err := SaveTask(root, task); err != nil {
		t.Fatal(err)
	}

	got, err := ReopenTask(root, "reopen-task")
	if err != nil {
		t.Fatal(err)
	}
	if got.State != TaskImplemented {
		t.Fatalf("state=%s，want IMPLEMENTED", got.State)
	}
	if got.ClosedAt != nil || got.Evidence.VerifiedTree != "" || got.Evidence.GateReport != "" {
		t.Fatalf("关闭态证据未清空: %+v", got)
	}
	// 输入侧证据保留。
	if got.Evidence.Regression == "" || got.Evidence.ProbeReport == "" {
		t.Fatalf("输入证据被误清: %+v", got.Evidence)
	}
	// 落盘与 ledger 同步。
	loaded, err := LoadTask(root, "reopen-task")
	if err != nil {
		t.Fatal(err)
	}
	if loaded.State != TaskImplemented {
		t.Fatalf("落盘 state=%s", loaded.State)
	}
	entries, err := os.ReadFile(filepath.Join(root, "wg", "tasks", "ledger.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(entries), "IMPLEMENTED") {
		t.Fatalf("ledger 未同步: %s", entries)
	}

	// 非 CLOSED 状态 reopen 应报错。
	if _, err := ReopenTask(root, "reopen-task"); err == nil {
		t.Fatal("对 IMPLEMENTED 任务 reopen 应报错")
	}
}
