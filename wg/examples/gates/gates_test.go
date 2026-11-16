package gates

// 功能窗枚举门禁（M7）：examples/<feature>/main.go 是 package main，不可
// import，所以门禁以子进程方式逐窗跑 --selfcheck（headless 出帧 + 像素 +
// 降级判定）。真窗的像素证据由各窗的 --shot 提供（脚本），这里守住的是
// "每一窗都能独立出图、无静默降级、ID 与目录一致"。
//
// 全量跑一遍 ~72 窗：一次性 go build 后逐个执行，秒级完成。

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/energye/lcl/wg/examples/internal/launcher"
)

// moduleRoot 返回 Go 模块根（lcl）；examplesDir 是功能窗目录。
func moduleRoot(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("无法定位测试文件")
	}
	// gates_test.go 位于 lcl/wg/examples/gates/，向上三级到模块根。
	return filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(thisFile))))
}

func examplesDir(t *testing.T) string {
	t.Helper()
	return filepath.Join(moduleRoot(t), "wg", "examples")
}

// TestEveryFeatureWindowSelfCheck 枚举全部功能窗并逐窗自检。
func TestEveryFeatureWindowSelfCheck(t *testing.T) {
	dir := examplesDir(t)

	// 1) 收集功能目录（含 main.go 的子目录）
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	var features []string
	for _, e := range entries {
		if e.IsDir() && e.Name() != "doc" && e.Name() != "internal" {
			if _, err := os.Stat(filepath.Join(dir, e.Name(), "main.go")); err == nil {
				features = append(features, e.Name())
			}
		}
	}
	if len(features) < 60 {
		t.Fatalf("功能窗数量异常：只有 %d 窗（应 ≥60）", len(features))
	}

	// 2) 一次性构建全部 main 包
	binDir := t.TempDir()
	build := exec.Command("go", "build", "-o", binDir, "./wg/examples/...")
	build.Dir = moduleRoot(t) // 模块根（lcl）
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("构建功能窗失败: %v\n%s", err, out)
	}

	// 3) 逐窗 --selfcheck
	seen := map[string]string{} // id -> 目录
	groups := map[string]bool{}
	for _, name := range features {
		t.Run(name, func(t *testing.T) {
			bin := filepath.Join(binDir, name+".exe")
			if _, err := os.Stat(bin); err != nil {
				// 非 Windows 无 .exe 后缀；按实际文件名兜底
				bin = filepath.Join(binDir, name)
				if _, err := os.Stat(bin); err != nil {
					t.Fatalf("二进制缺失：%s", bin)
				}
			}
			cmd := exec.Command(bin, "--selfcheck")
			out, err := cmd.Output()
			if err != nil {
				t.Fatalf("selfcheck 失败（exit != 0）:\n%s", out)
			}
			var rep launcher.SelfCheckReport
			if err := json.Unmarshal(out, &rep); err != nil {
				t.Fatalf("selfcheck 输出不是合法 JSON:\n%s", out)
			}
			if !rep.OK {
				t.Fatalf("自检未通过: %s", rep.Reason)
			}
			if rep.ID != name {
				t.Fatalf("meta.ID=%q 与目录名 %q 不一致", rep.ID, name)
			}
			if rep.Degraded != 0 {
				t.Fatalf("有 %d 条指令降级", rep.Degraded)
			}
			if old, dup := seen[rep.ID]; dup {
				t.Fatalf("功能 ID 重复：%s 已由 %s 注册", rep.ID, old)
			}
			seen[rep.ID] = name
			groups[rep.Group] = true
		})
	}

	// 4) 分组覆盖：九组每至少一窗
	for _, g := range []string{
		"基础组件", "输入控制", "展示组件", "数据组件", "反馈浮层",
		"图表", "布局引擎", "绘制与合成", "事件与输入",
	} {
		if !groups[g] {
			t.Errorf("分组 %q 没有任何功能窗", g)
		}
	}
}

// TestFeatureWindowMetaFormat --meta 输出必须可解析且 ID 一致（给外部脚本用）。
func TestFeatureWindowMetaFormat(t *testing.T) {
	dir := examplesDir(t)
	target := filepath.Join(dir, "button")
	bin := filepath.Join(t.TempDir(), "button.exe")
	build := exec.Command("go", "build", "-o", bin, target)
	build.Dir = moduleRoot(t)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("构建失败: %v\n%s", err, out)
	}
	out, err := exec.Command(bin, "--meta").Output()
	if err != nil {
		t.Fatalf("--meta 失败: %v\n%s", err, out)
	}
	var m struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Group string `json:"group"`
	}
	if err := json.Unmarshal(out, &m); err != nil {
		t.Fatalf("--meta 输出不是合法 JSON:\n%s", out)
	}
	if m.ID != "button" || !strings.Contains(m.Title, "Button") {
		t.Fatalf("meta 内容异常: %+v", m)
	}
}
