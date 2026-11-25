package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/energye/lcl/wg/internal/gate"
)

func probeCommand(args []string) error {
	if len(args) == 0 || args[0] != "new" {
		return fmt.Errorf("probe 仅支持 new")
	}
	fs := flag.NewFlagSet("probe new", flag.ExitOnError)
	taskID := fs.String("task", "", "任务 ID")
	probeID := fs.String("id", "", "探针 ID")
	command := fs.String("command", "", "可执行命令")
	timeout := fs.Duration("timeout", 2*time.Minute, "探针超时")
	var covers stringList
	fs.Var(&covers, "covers", "覆盖路径，可重复")
	_ = fs.Parse(args[1:])
	if *taskID == "" || *probeID == "" || *command == "" || len(covers) == 0 {
		return fmt.Errorf("probe new 需要 --task、--id、--command 和至少一个 --covers")
	}
	repo, err := root()
	if err != nil {
		return err
	}
	task, probe, err := gate.RegisterProbe(repo, "wg/gate.probes.json", *taskID, *probeID, covers, strings.Fields(*command), *timeout)
	if err != nil {
		return err
	}
	fmt.Printf("已生成探针 %s 并注册到任务 %s（%s）\n", probe.ID, task.ID, task.State)
	return nil
}
