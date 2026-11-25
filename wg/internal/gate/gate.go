package gate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const ReportVersion = 1

type Status string

const (
	StatusPass    Status = "PASS"
	StatusFail    Status = "FAIL"
	StatusTimeout Status = "TIMEOUT"
)

type Step struct {
	ID         string `json:"id"`
	Command    string `json:"command,omitempty"`
	Status     Status `json:"status"`
	DurationMS int64  `json:"durationMs"`
	Detail     string `json:"detail,omitempty"`
	LogFile    string `json:"logFile,omitempty"`
}

type AuditSummary struct {
	Expected int `json:"expected"`
	Passed   int `json:"passed"`
	Failed   int `json:"failed"`
	TimedOut int `json:"timedOut"`
}

type Report struct {
	Version      int          `json:"version"`
	TreeHash     string       `json:"treeHash"`
	BaseHEAD     string       `json:"baseHead,omitempty"`
	StartedAt    time.Time    `json:"startedAt"`
	FinishedAt   time.Time    `json:"finishedAt"`
	Status       Status       `json:"status"`
	ChangedFiles []string     `json:"changedFiles"`
	Audits       AuditSummary `json:"audits"`
	Steps        []Step       `json:"steps"`
}

type Options struct {
	ReportDir      string
	ProbeFile      string
	DefaultTimeout time.Duration
	AuditTimeout   time.Duration
	Stdout         io.Writer
	Stderr         io.Writer
}

func defaultOptions() Options {
	return Options{
		ReportDir:      ".shots/gates",
		ProbeFile:      "wg/gate.probes.json",
		DefaultTimeout: 10 * time.Minute,
		AuditTimeout:   2 * time.Minute,
		Stdout:         os.Stdout,
		Stderr:         os.Stderr,
	}
}

func Run(ctx context.Context, root, tree string, opts Options) (*Report, error) {
	if opts.ReportDir == "" {
		opts.ReportDir = defaultOptions().ReportDir
	}
	if opts.ProbeFile == "" {
		opts.ProbeFile = defaultOptions().ProbeFile
	}
	if opts.DefaultTimeout <= 0 {
		opts.DefaultTimeout = defaultOptions().DefaultTimeout
	}
	if opts.AuditTimeout <= 0 {
		opts.AuditTimeout = defaultOptions().AuditTimeout
	}
	if opts.Stdout == nil {
		opts.Stdout = io.Discard
	}
	if opts.Stderr == nil {
		opts.Stderr = io.Discard
	}
	if err := EnsureStaged(root); err != nil {
		return nil, err
	}
	if tree == "" {
		var err error
		tree, err = StagedTree(root)
		if err != nil {
			return nil, err
		}
	}
	files, err := ChangedFiles(root, tree)
	if err != nil {
		return nil, err
	}
	features, err := FeatureWindows(root)
	if err != nil {
		return nil, err
	}
	cfg, cfgErr := LoadProbeConfig(root, opts.ProbeFile)
	head, _ := HeadTree(root)
	report := &Report{
		Version:      ReportVersion,
		TreeHash:     tree,
		BaseHEAD:     head,
		StartedAt:    time.Now(),
		ChangedFiles: files,
		Audits:       AuditSummary{Expected: len(features)},
	}
	logDir := filepath.Join(opts.ReportDir, "logs", tree)
	if !filepath.IsAbs(logDir) {
		logDir = filepath.Join(root, logDir)
	}

	report.Steps = append(report.Steps, runSteps(ctx, root, logDir, opts.Stdout, opts.Stderr, baseSteps(), opts.DefaultTimeout)...)

	tempDir, err := os.MkdirTemp("", "wg-gate-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)
	diagBin := filepath.Join(tempDir, exeName("wg-diag"))
	diagBuild := runStep(ctx, root, logDir, opts.Stdout, opts.Stderr, StepSpec{
		ID: "diag:build", Args: []string{"go", "build", "-buildvcs=false", "-o", diagBin, "./wg/cmd/wg-diag"}, Timeout: 5 * time.Minute,
	}, opts.DefaultTimeout)
	report.Steps = append(report.Steps, diagBuild)
	if diagBuild.Status == StatusPass {
		report.Steps = append(report.Steps, runSteps(ctx, root, logDir, opts.Stdout, opts.Stderr, diagSteps(diagBin), opts.DefaultTimeout)...)
	}

	binDir := filepath.Join(tempDir, "examples")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		return nil, err
	}
	auditBuild := runStep(ctx, root, logDir, opts.Stdout, opts.Stderr, StepSpec{
		ID: "audit:build", Args: []string{"go", "build", "-buildvcs=false", "-o", binDir, "./wg/examples/..."}, Timeout: 10 * time.Minute,
	}, opts.DefaultTimeout)
	report.Steps = append(report.Steps, auditBuild)
	if auditBuild.Status == StatusPass {
		for _, feature := range features {
			fmt.Fprintf(opts.Stdout, "[wg-gate] RUN audit:%s\n", feature)
			step, _ := runAudit(ctx, root, filepath.Join(binDir, exeName(feature)), feature, logDir, opts.Stdout, opts.Stderr, opts.AuditTimeout)
			fmt.Fprintf(opts.Stdout, "[wg-gate] %s audit:%s (%dms)\n", step.Status, feature, step.DurationMS)
			report.Steps = append(report.Steps, step)
			switch step.Status {
			case StatusPass:
				report.Audits.Passed++
			case StatusTimeout:
				report.Audits.TimedOut++
			default:
				report.Audits.Failed++
			}
		}
	}

	if cfgErr != nil {
		report.Steps = append(report.Steps, Step{ID: "probes:config", Status: StatusFail, Detail: cfgErr.Error()})
	} else {
		missing := MissingProbeCoverage(cfg, files)
		for _, file := range missing {
			report.Steps = append(report.Steps, Step{
				ID: "probe-coverage:" + file, Status: StatusFail,
				Detail: "改动路径需要 §6A 性能探针，但 gate.probes.json 未覆盖",
			})
		}
		if len(cfg.Probes) == 0 {
			detail := "没有注册性能探针，且本次没有性能相关改动"
			if len(missing) > 0 {
				detail = "存在未注册性能探针的改动"
			}
			status := StatusPass
			if len(missing) > 0 {
				status = StatusFail
			}
			report.Steps = append(report.Steps, Step{ID: "probes", Status: status, Detail: detail})
		} else {
			for _, probe := range cfg.Probes {
				step := runStep(ctx, root, logDir, opts.Stdout, opts.Stderr, StepSpec{
					ID: "probe:" + probe.ID, Args: probe.Command, Timeout: probe.Timeout,
				}, opts.DefaultTimeout)
				report.Steps = append(report.Steps, step)
			}
		}
	}

	report.FinishedAt = time.Now()
	report.Status = StatusFail
	if allPass(report.Steps) && report.Audits.Expected == report.Audits.Passed {
		report.Status = StatusPass
	}
	if err := WriteReport(root, opts.ReportDir, report); err != nil {
		return report, err
	}
	return report, nil
}

func allPass(steps []Step) bool {
	if len(steps) == 0 {
		return false
	}
	for _, step := range steps {
		if step.Status != StatusPass {
			return false
		}
	}
	return true
}

func ReportPath(root, dir, tree string) string {
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(root, dir)
	}
	return filepath.Join(dir, "report-"+tree+".json")
}

func WriteReport(root, dir string, report *Report) error {
	base := dir
	if !filepath.IsAbs(base) {
		base = filepath.Join(root, base)
	}
	if err := os.MkdirAll(base, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	paths := []string{ReportPath(root, dir, report.TreeHash), filepath.Join(base, "latest.json")}
	for _, path := range paths {
		tmp := path + ".tmp"
		if err := os.WriteFile(tmp, data, 0o644); err != nil {
			return err
		}
		if err := os.Rename(tmp, path); err != nil {
			_ = os.Remove(path)
			if err := os.Rename(tmp, path); err != nil {
				return err
			}
		}
	}
	return nil
}

func Check(root, tree string, opts Options) error {
	if opts.ReportDir == "" {
		opts.ReportDir = defaultOptions().ReportDir
	}
	if opts.ProbeFile == "" {
		opts.ProbeFile = defaultOptions().ProbeFile
	}
	data, err := os.ReadFile(ReportPath(root, opts.ReportDir, tree))
	if err != nil {
		latest := opts.ReportDir
		if !filepath.IsAbs(latest) {
			latest = filepath.Join(root, latest)
		}
		data, err = os.ReadFile(filepath.Join(latest, "latest.json"))
		if err != nil {
			return fmt.Errorf("门禁报告缺失或过期：%w\n先执行 go run ./wg/cmd/wg-gate run", err)
		}
	}
	var report Report
	if err := json.Unmarshal(data, &report); err != nil {
		return fmt.Errorf("门禁报告损坏: %w", err)
	}
	if report.Version != ReportVersion {
		return fmt.Errorf("报告版本=%d，期望 %d", report.Version, ReportVersion)
	}
	if report.TreeHash != tree {
		if len(report.TreeHash) != 40 {
			return fmt.Errorf("报告 tree=%q 非法，当前 staged tree=%s", report.TreeHash, tree)
		}
		taskOnly, err := OnlyTaskChanges(root, report.TreeHash, tree)
		if err != nil {
			return err
		}
		if !taskOnly {
			return fmt.Errorf("报告 tree=%s，当前 staged tree=%s；代码文件已变化", report.TreeHash, tree)
		}
	}
	files, err := ChangedFiles(root, tree)
	if err != nil {
		return err
	}
	reportFiles := WithoutTaskFiles(report.ChangedFiles)
	currentFiles := WithoutTaskFiles(files)
	if strings.Join(reportFiles, "\n") != strings.Join(currentFiles, "\n") {
		return fmt.Errorf("报告改动文件清单与当前 staged tree 不一致")
	}
	if report.Status != StatusPass {
		return fmt.Errorf("门禁报告状态=%s", report.Status)
	}
	for _, step := range report.Steps {
		if step.Status != StatusPass {
			return fmt.Errorf("门禁步骤 %s=%s: %s", step.ID, step.Status, step.Detail)
		}
	}
	features, err := FeatureWindows(root)
	if err != nil {
		return err
	}
	if report.Audits.Expected != len(features) || report.Audits.Passed != len(features) || report.Audits.Failed != 0 || report.Audits.TimedOut != 0 {
		return fmt.Errorf("审计汇总不完整: expected=%d passed=%d failed=%d timeout=%d，当前窗口=%d",
			report.Audits.Expected, report.Audits.Passed, report.Audits.Failed, report.Audits.TimedOut, len(features))
	}
	cfg, err := LoadProbeConfig(root, opts.ProbeFile)
	if err != nil {
		return err
	}
	required := map[string]bool{}
	for _, spec := range baseSteps() {
		required[spec.ID] = true
	}
	for _, spec := range diagSteps("wg-diag") {
		required[spec.ID] = true
	}
	for _, feature := range features {
		required["audit:"+feature] = true
	}
	if len(cfg.Probes) == 0 {
		required["probes"] = true
	} else {
		for _, probe := range cfg.Probes {
			required["probe:"+probe.ID] = true
		}
	}
	seen := map[string]bool{}
	for _, step := range report.Steps {
		if seen[step.ID] {
			return fmt.Errorf("报告包含重复步骤 %s", step.ID)
		}
		seen[step.ID] = true
	}
	var missing []string
	for id := range required {
		if !seen[id] {
			missing = append(missing, id)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		return fmt.Errorf("报告缺少必需步骤: %s", strings.Join(missing, ", "))
	}
	if len(MissingProbeCoverage(cfg, currentFiles)) > 0 {
		return fmt.Errorf("存在未注册 §6A 性能探针的改动")
	}
	return nil
}
