package gate

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type TaskProbeResult struct {
	ProbeID    string `json:"probeId"`
	Status     Status `json:"status"`
	DurationMS int64  `json:"durationMs"`
	Detail     string `json:"detail,omitempty"`
	LogFile    string `json:"logFile,omitempty"`
}

type TaskProbeReport struct {
	Version  int               `json:"version"`
	TaskID   string            `json:"taskId"`
	TreeHash string            `json:"treeHash"`
	Status   Status            `json:"status"`
	Results  []TaskProbeResult `json:"results"`
}

func VerifyTask(root, taskID, probeFile string) (Task, error) {
	task, err := LoadTask(root, taskID)
	if err != nil {
		return Task{}, err
	}
	if task.State != TaskNew && task.State != TaskSpecRed && task.State != TaskImplemented && task.State != TaskVerified && task.State != TaskBlocked {
		return Task{}, fmt.Errorf("任务状态 %s 不允许 verify", task.State)
	}
	if len(task.Requirements) == 0 {
		return Task{}, fmt.Errorf("任务缺少 requirements")
	}
	if len(task.Specs) == 0 {
		return Task{}, fmt.Errorf("任务缺少 spec 路径")
	}
	for _, spec := range task.Specs {
		if _, err := os.Stat(filepath.Join(root, spec)); err != nil {
			return Task{}, fmt.Errorf("spec 不存在: %s", spec)
		}
	}
	if task.Source.Type == "human-feedback" {
		if task.Evidence.Regression == "" {
			return Task{}, fmt.Errorf("人工反馈任务缺少 regression.json")
		}
		if _, err := os.Stat(filepath.Join(root, task.Evidence.Regression)); err != nil {
			return Task{}, fmt.Errorf("回归初稿不存在: %s", task.Evidence.Regression)
		}
		if _, err := os.Stat(filepath.Join(TaskDir(root, taskID), "feedback.json")); err != nil {
			return Task{}, fmt.Errorf("反馈快照不存在")
		}
	}
	if len(task.Probes) == 0 {
		return Task{}, fmt.Errorf("任务尚未注册 §6A 探针")
	}
	cfg, err := LoadProbeConfig(root, probeFile)
	if err != nil {
		return Task{}, err
	}
	registered := map[string]Probe{}
	for _, probe := range cfg.Probes {
		registered[probe.ID] = probe
	}
	for _, id := range task.Probes {
		if _, ok := registered[id]; !ok {
			return Task{}, fmt.Errorf("任务探针 %s 未注册", id)
		}
	}
	if err := EnsureStaged(root); err != nil {
		return Task{}, err
	}
	tree, err := StagedTree(root)
	if err != nil {
		return Task{}, err
	}
	logDir := filepath.Join(TaskDir(root, taskID), "logs", tree)
	report := TaskProbeReport{Version: 1, TaskID: taskID, TreeHash: tree, Status: StatusPass}
	for _, id := range task.Probes {
		probe := registered[id]
		step := runStep(context.Background(), root, logDir, io.Discard, io.Discard, StepSpec{
			ID: "probe:" + id, Args: probe.Command, Timeout: probe.Timeout,
		}, 2*time.Minute)
		report.Results = append(report.Results, TaskProbeResult{
			ProbeID: id, Status: step.Status, DurationMS: step.DurationMS, Detail: step.Detail, LogFile: step.LogFile,
		})
		if step.Status != StatusPass {
			report.Status = StatusFail
		}
	}
	reportPath := filepath.ToSlash(filepath.Join("wg", "tasks", taskID, "probe-report.json"))
	if err := writeJSON(filepath.Join(root, reportPath), report); err != nil {
		return Task{}, err
	}
	task.Evidence.ProbeReport = reportPath
	if report.Status != StatusPass {
		task.State = TaskBlocked
		task.BlockedReason = "任务探针未通过，详见 " + reportPath
		if err := SaveTask(root, task); err != nil {
			return Task{}, err
		}
		return task, fmt.Errorf("%s", task.BlockedReason)
	}
	task.State = TaskVerified
	task.Evidence.VerifiedTree = tree
	task.BlockedReason = ""
	if err := SaveTask(root, task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func CloseTask(root, taskID, reportDir, probeFile string) error {
	return CloseTaskWithChecker(root, taskID, reportDir, probeFile, func(root, tree string) error {
		return Check(root, tree, Options{ReportDir: reportDir, ProbeFile: probeFile})
	})
}

func CloseTaskWithChecker(root, taskID, reportDir, probeFile string, checker func(string, string) error) error {
	task, err := LoadTask(root, taskID)
	if err != nil {
		return err
	}
	if task.State != TaskVerified {
		return fmt.Errorf("任务状态 %s，必须先 verify", task.State)
	}
	if err := EnsureStaged(root); err != nil {
		return err
	}
	tree, err := StagedTree(root)
	if err != nil {
		return err
	}
	if task.Evidence.VerifiedTree != tree {
		taskOnly, err := OnlyTaskChanges(root, task.Evidence.VerifiedTree, tree)
		if err != nil {
			return err
		}
		if !taskOnly {
			return fmt.Errorf("verify 后代码文件已变化，请重新 verify")
		}
	}
	if err := checker(root, tree); err != nil {
		task.State = TaskBlocked
		task.BlockedReason = err.Error()
		if saveErr := SaveTask(root, task); saveErr != nil {
			return saveErr
		}
		return err
	}
	now := time.Now()
	task.State = TaskClosed
	task.ClosedAt = &now
	task.Evidence.GateReport = ReportPath(root, reportDir, tree)
	task.BlockedReason = ""
	if err := SaveTask(root, task); err != nil {
		return err
	}
	return nil
}
