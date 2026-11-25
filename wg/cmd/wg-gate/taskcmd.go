package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/energye/lcl/wg/internal/gate"
)

type stringList []string

func (s *stringList) String() string { return strings.Join(*s, ", ") }
func (s *stringList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func taskCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("task 需要子命令：new 或 status")
	}
	switch args[0] {
	case "new":
		return taskNew(args[1:])
	case "status":
		return taskStatus(args[1:])
	case "verify":
		return taskVerify(args[1:])
	case "close":
		return taskClose(args[1:])
	default:
		return fmt.Errorf("未知 task 子命令 %q", args[0])
	}
}

func taskNew(args []string) error {
	fs := flag.NewFlagSet("task new", flag.ExitOnError)
	id := fs.String("id", "", "任务 ID")
	component := fs.String("component", "", "组件名")
	title := fs.String("title", "", "任务标题")
	var requirements stringList
	fs.Var(&requirements, "requirement", "验收要求，可重复")
	_ = fs.Parse(args)
	if *id == "" || *component == "" || *title == "" || len(requirements) == 0 {
		return fmt.Errorf("task new 需要 --id、--component、--title 和至少一个 --requirement")
	}
	repo, err := root()
	if err != nil {
		return err
	}
	task, err := gate.NewTask(repo, *id, *component, *title, requirements)
	if err != nil {
		return err
	}
	fmt.Printf("已创建任务 %s（%s）: %s\n", task.ID, task.State, gate.TaskPath(repo, task.ID))
	return nil
}

func taskStatus(args []string) error {
	fs := flag.NewFlagSet("task status", flag.ExitOnError)
	id := fs.String("id", "", "任务 ID")
	_ = fs.Parse(args)
	if *id == "" {
		return fmt.Errorf("task status 需要 --id")
	}
	repo, err := root()
	if err != nil {
		return err
	}
	task, err := gate.LoadTask(repo, *id)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func taskVerify(args []string) error {
	fs := flag.NewFlagSet("task verify", flag.ExitOnError)
	id := fs.String("id", "", "任务 ID")
	probeFile := fs.String("probe-file", "wg/gate.probes.json", "性能探针清单")
	_ = fs.Parse(args)
	if *id == "" {
		return fmt.Errorf("task verify 需要 --id")
	}
	repo, err := root()
	if err != nil {
		return err
	}
	task, err := gate.VerifyTask(repo, *id, *probeFile)
	if err != nil {
		return err
	}
	fmt.Printf("任务 %s 已进入 %s，tree=%s\n", task.ID, task.State, task.Evidence.VerifiedTree)
	return nil
}

func taskClose(args []string) error {
	fs := flag.NewFlagSet("task close", flag.ExitOnError)
	id := fs.String("id", "", "任务 ID")
	reportDir := fs.String("report-dir", ".shots/gates", "全量门禁报告目录")
	probeFile := fs.String("probe-file", "wg/gate.probes.json", "性能探针清单")
	_ = fs.Parse(args)
	if *id == "" {
		return fmt.Errorf("task close 需要 --id")
	}
	repo, err := root()
	if err != nil {
		return err
	}
	if err := gate.CloseTask(repo, *id, *reportDir, *probeFile); err != nil {
		return err
	}
	fmt.Printf("任务 %s 已 CLOSED\n", *id)
	return nil
}

func feedbackCommand(args []string) error {
	if len(args) == 0 || args[0] != "import" {
		return fmt.Errorf("feedback 仅支持 import")
	}
	fs := flag.NewFlagSet("feedback import", flag.ExitOnError)
	taskID := fs.String("task", "", "任务 ID")
	file := fs.String("file", "", "反馈 JSON 文件")
	_ = fs.Parse(args[1:])
	if *taskID == "" || *file == "" {
		return fmt.Errorf("feedback import 需要 --task 与 --file")
	}
	repo, err := root()
	if err != nil {
		return err
	}
	task, err := gate.ImportFeedback(repo, *taskID, *file)
	if err != nil {
		return err
	}
	fmt.Printf("已导入反馈 %s，任务 %s 进入 %s\n", task.Source.ID, task.ID, task.State)
	return nil
}
