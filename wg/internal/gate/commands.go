package gate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type StepSpec struct {
	ID      string
	Args    []string
	Env     []string
	Timeout time.Duration
}

func mergeEnv(base, extra []string) []string {
	if len(extra) == 0 {
		return base
	}
	values := map[string]string{}
	for _, item := range base {
		if i := strings.IndexByte(item, '='); i >= 0 {
			values[item[:i]] = item[i+1:]
		}
	}
	for _, item := range extra {
		if i := strings.IndexByte(item, '='); i >= 0 {
			values[item[:i]] = item[i+1:]
		}
	}
	out := make([]string, 0, len(values))
	for key, value := range values {
		out = append(out, key+"="+value)
	}
	return out
}

func safeLogName(id string) string {
	return strings.NewReplacer("/", "_", "\\", "_", ":", "_").Replace(id)
}

func runStep(ctx context.Context, root, logDir string, stdout, stderr io.Writer, spec StepSpec, defaultTimeout time.Duration) Step {
	step := Step{ID: spec.ID, Command: strings.Join(spec.Args, " "), Status: StatusPass}
	if stdout == nil {
		stdout = io.Discard
	}
	if stderr == nil {
		stderr = io.Discard
	}
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		step.Status = StatusFail
		step.Detail = err.Error()
		return step
	}
	logPath := filepath.Join(logDir, safeLogName(spec.ID)+".log")
	logFile, err := os.Create(logPath)
	if err != nil {
		step.Status = StatusFail
		step.Detail = err.Error()
		return step
	}
	defer logFile.Close()
	step.LogFile = logPath

	timeout := spec.Timeout
	if timeout <= 0 {
		timeout = defaultTimeout
	}
	runCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	cmd := exec.CommandContext(runCtx, spec.Args[0], spec.Args[1:]...)
	cmd.Dir = root
	cmd.Env = mergeEnv(os.Environ(), spec.Env)
	cmd.Stdout = io.MultiWriter(stdout, logFile)
	cmd.Stderr = io.MultiWriter(stderr, logFile)
	began := time.Now()
	err = cmd.Run()
	step.DurationMS = time.Since(began).Milliseconds()
	if errors.Is(runCtx.Err(), context.DeadlineExceeded) {
		step.Status = StatusTimeout
		step.Detail = fmt.Sprintf("超过 %s", timeout)
		return step
	}
	if err != nil {
		step.Status = StatusFail
		step.Detail = err.Error()
		return step
	}
	step.Detail = "exit 0"
	return step
}

func runSteps(ctx context.Context, root, logDir string, stdout, stderr io.Writer, specs []StepSpec, timeout time.Duration) []Step {
	steps := make([]Step, 0, len(specs))
	for _, spec := range specs {
		fmt.Fprintf(stdout, "[wg-gate] RUN %s\n", spec.ID)
		step := runStep(ctx, root, logDir, stdout, stderr, spec, timeout)
		fmt.Fprintf(stdout, "[wg-gate] %s %s (%dms)\n", step.Status, spec.ID, step.DurationMS)
		steps = append(steps, step)
	}
	return steps
}

func baseSteps() []StepSpec {
	return []StepSpec{
		{ID: "build", Args: []string{"go", "build", "./wg/..."}, Timeout: 10 * time.Minute},
		{ID: "vet", Args: []string{"go", "vet", "./wg/..."}, Timeout: 10 * time.Minute},
		{ID: "checkdeps", Args: []string{"go", "run", "./wg/internal/checkdeps"}, Timeout: 5 * time.Minute},
		{ID: "test", Args: []string{"go", "test", "-count=1", "./wg/..."}, Timeout: 15 * time.Minute},
	}
}

func diagSteps(diag string) []StepSpec {
	commands := []struct {
		id   string
		args []string
	}{
		{id: "contracts", args: []string{"--contracts"}},
		{id: "selftest", args: []string{"--selftest"}},
		{id: "layout", args: []string{"--layout"}},
		{id: "layers", args: []string{"--layers"}},
		{id: "schedtest", args: []string{"--schedtest"}},
		{id: "hit", args: []string{"--hit"}},
		{id: "gesturetest", args: []string{"--gesturetest"}},
		{id: "inputtest", args: []string{"--inputtest"}},
		{id: "resizetest", args: []string{"--resizetest"}},
		{id: "overlaytest", args: []string{"--overlaytest"}},
		{id: "idle", args: []string{"--idle"}},
		{id: "keyselftest", args: []string{"--keyselftest"}},
		{id: "formtest", args: []string{"--formtest"}},
		{id: "migrationtest", args: []string{"--migrationtest"}},
		{id: "damage", args: []string{"--damage"}},
		{id: "render-matrix", args: []string{"--render-matrix", "-"}},
		{id: "golden", args: []string{"--golden", "wg/testdata"}},
		{id: "bench", args: []string{"--bench", "--frames", "120"}},
		{id: "headless", args: []string{"--headless", "--frames", "120"}},
	}
	steps := make([]StepSpec, 0, len(commands)+4)
	for _, command := range commands {
		steps = append(steps, StepSpec{
			ID:      "diag:" + command.id,
			Args:    append([]string{diag}, command.args...),
			Timeout: 5 * time.Minute,
		})
	}
	// --bench 走扁平 painter 路径（SetPainter→Draw），真实窗口走场景树
	// DrawTree。因此：WG_AA 对 bench 有真实对照意义（关闭 AA 直接降光栅成本）；
	// WG_LAYERS/WG_CACHE 在扁平路径下不适用（扁平路径不建层、不查层缓存），
	// 保留这两步仅作"改动后仍 ≥59 FPS"的回归冒烟，不构成层/缓存 A/B 证据。
	steps = append(steps,
		StepSpec{ID: "diag:bench-aa-off", Args: []string{diag, "--bench", "--frames", "120"}, Env: []string{"WG_AA=0"}, Timeout: 5 * time.Minute},
		StepSpec{ID: "diag:bench-layers-off", Args: []string{diag, "--bench", "--frames", "120"}, Env: []string{"WG_LAYERS=0"}, Timeout: 5 * time.Minute},
		StepSpec{ID: "diag:bench-cache-off", Args: []string{diag, "--bench", "--frames", "120"}, Env: []string{"WG_CACHE=0"}, Timeout: 5 * time.Minute},
		StepSpec{ID: "diag:bench-damage-on", Args: []string{diag, "--bench", "--frames", "120"}, Env: []string{"WG_DAMAGE=1"}, Timeout: 5 * time.Minute},
	)
	return steps
}

type auditReport struct {
	Spec       string `json:"spec"`
	Pages      int    `json:"pages"`
	Elements   int    `json:"elements"`
	Failed     int    `json:"failed"`
	SpecFailed int    `json:"specFailed"`
	Findings   []struct {
		OK bool `json:"ok"`
	} `json:"findings"`
}

func runAudit(ctx context.Context, root, bin, id, logDir string, stdout, stderr io.Writer, timeout time.Duration) (Step, auditReport) {
	spec := StepSpec{ID: "audit:" + id, Args: []string{bin, "--audit"}, Timeout: timeout}
	step := runStep(ctx, root, logDir, stdout, stderr, spec, timeout)
	var rep auditReport
	if step.Status != StatusPass {
		return step, rep
	}
	if id == "multi-window" {
		step.Detail = "exit 0（multi-window 生命周期审计，无 report.json）"
		return step, rep
	}
	data, err := os.ReadFile(filepath.Join(root, ".shots", "audit", id, "report.json"))
	if err != nil {
		step.Status = StatusFail
		step.Detail = "缺少 report.json: " + err.Error()
		return step, rep
	}
	if err := json.Unmarshal(data, &rep); err != nil {
		step.Status = StatusFail
		step.Detail = "report.json 解析失败: " + err.Error()
		return step, rep
	}
	// 与审计进程退出码一致：有声明式 spec 时取 specFailed，启发式
	// 误报（如启用态随点击翻转）不阻塞门禁；无 spec 才用启发式失败数。
	gateFailed := rep.Failed
	if rep.Spec != "" {
		gateFailed = rep.SpecFailed
	}
	if gateFailed != 0 {
		step.Status = StatusFail
		step.Detail = fmt.Sprintf("failed=%d specFailed=%d", rep.Failed, rep.SpecFailed)
		return step, rep
	}
	step.Detail = fmt.Sprintf("pages=%d elements=%d findings=%d", rep.Pages, rep.Elements, len(rep.Findings))
	return step, rep
}
