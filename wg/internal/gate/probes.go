package gate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Probe struct {
	ID      string        `json:"id"`
	Covers  []string      `json:"covers"`
	Command []string      `json:"command"`
	Timeout time.Duration `json:"timeout"`
}

type probeConfig struct {
	Version         int      `json:"version"`
	RequireProbeFor []string `json:"requireProbeFor"`
	Probes          []Probe  `json:"probes"`
}

func LoadProbeConfig(root, file string) (probeConfig, error) {
	var cfg probeConfig
	data, err := os.ReadFile(filepath.Join(root, file))
	if err != nil {
		return cfg, fmt.Errorf("读取性能探针清单失败: %w", err)
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("解析性能探针清单失败: %w", err)
	}
	if cfg.Version != 1 {
		return cfg, fmt.Errorf("探针清单版本=%d，仅支持 1", cfg.Version)
	}
	if len(cfg.RequireProbeFor) == 0 {
		return cfg, fmt.Errorf("requireProbeFor 不能为空")
	}
	seen := map[string]bool{}
	for _, p := range cfg.Probes {
		if p.ID == "" || seen[p.ID] {
			return cfg, fmt.Errorf("探针 ID 为空或重复: %q", p.ID)
		}
		seen[p.ID] = true
		if len(p.Command) == 0 || p.Command[0] == "" {
			return cfg, fmt.Errorf("探针 %s 缺少 command", p.ID)
		}
		if len(p.Covers) == 0 {
			return cfg, fmt.Errorf("探针 %s 缺少 covers", p.ID)
		}
	}
	return cfg, nil
}

func matchesPattern(pattern, file string) bool {
	pattern = filepath.ToSlash(pattern)
	file = filepath.ToSlash(file)
	if strings.HasSuffix(pattern, "/") {
		return strings.HasPrefix(file, pattern)
	}
	ok, err := path.Match(pattern, file)
	return err == nil && ok
}

func requiresProbe(cfg probeConfig, file string) bool {
	for _, pattern := range cfg.RequireProbeFor {
		if matchesPattern(pattern, file) {
			return true
		}
	}
	return false
}

func probeCovers(p Probe, file string) bool {
	for _, pattern := range p.Covers {
		if matchesPattern(pattern, file) {
			return true
		}
	}
	return false
}

type TaskProbe struct {
	Version int           `json:"version"`
	TaskID  string        `json:"taskId"`
	ProbeID string        `json:"probeId"`
	Covers  []string      `json:"covers"`
	Command []string      `json:"command"`
	Timeout time.Duration `json:"timeout"`
}

func RegisterProbe(root, probeFile, taskID, probeID string, covers, command []string, timeout time.Duration) (Task, Probe, error) {
	if err := ValidateID(probeID); err != nil {
		return Task{}, Probe{}, err
	}
	if len(covers) == 0 || len(command) == 0 || command[0] == "" {
		return Task{}, Probe{}, fmt.Errorf("探针需要 covers 与 command")
	}
	task, err := LoadTask(root, taskID)
	if err != nil {
		return Task{}, Probe{}, err
	}
	if task.State == TaskClosed {
		return Task{}, Probe{}, fmt.Errorf("已关闭任务不能注册探针")
	}
	cfg, err := LoadProbeConfig(root, probeFile)
	if err != nil {
		return Task{}, Probe{}, err
	}
	for _, existing := range cfg.Probes {
		if existing.ID == probeID {
			return Task{}, Probe{}, fmt.Errorf("探针 ID 已存在: %s", probeID)
		}
	}
	probe := Probe{ID: probeID, Covers: append([]string(nil), covers...), Command: append([]string(nil), command...), Timeout: timeout}
	cfg.Probes = append(cfg.Probes, probe)
	if err := writeJSON(filepath.Join(root, probeFile), cfg); err != nil {
		return Task{}, Probe{}, err
	}
	name := "probe.json"
	if len(task.Probes) > 0 {
		name = "probe-" + probeID + ".json"
	}
	evidence := TaskProbe{Version: 1, TaskID: taskID, ProbeID: probeID, Covers: probe.Covers, Command: probe.Command, Timeout: timeout}
	if err := writeJSON(filepath.Join(TaskDir(root, taskID), name), evidence); err != nil {
		return Task{}, Probe{}, err
	}
	task.Probes = append(task.Probes, probeID)
	if err := SaveTask(root, task); err != nil {
		return Task{}, Probe{}, err
	}
	return task, probe, nil
}

func MissingProbeCoverage(cfg probeConfig, files []string) []string {
	var missing []string
	for _, file := range files {
		// 测试文件改动不改变运行时性能路径，不要求 §6A 探针。
		if strings.HasSuffix(file, "_test.go") {
			continue
		}
		if !requiresProbe(cfg, file) {
			continue
		}
		covered := false
		for _, p := range cfg.Probes {
			if probeCovers(p, file) {
				covered = true
				break
			}
		}
		if !covered {
			missing = append(missing, file)
		}
	}
	return missing
}
