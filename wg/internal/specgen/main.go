// Command specgen 按 --meta 输出为每个功能窗生成/补齐 spec.json 的 L1 结构段。
//
// 用法（仓库根运行）：
//
//	go run ./wg/internal/specgen            # 补齐全部窗的 structure
//	go run ./wg/internal/specgen -dry-run   # 只打印，不写文件
//
// 规则（docs/wg-engine/06-COMPONENT-TEST-SPEC.md §4/§6）：
//   - 无 spec.json 的窗：新建 {id, structure:{sections, apiRows}}；
//   - 已有 spec.json：只补缺失的 structure 字段，cases/exemptions 原样保留；
//   - multi-window 不走 launcher.Main（无 --meta），跳过。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// metaOut 是 launcher --meta 的输出子集。
type metaOut struct {
	ID      string   `json:"id"`
	Demos   []string `json:"demos"`
	APIRows int      `json:"apiRows"`
}

func main() {
	dry := flag.Bool("dry-run", false, "只打印不写文件")
	flag.Parse()

	entries, err := os.ReadDir("wg/examples")
	if err != nil {
		fmt.Fprintln(os.Stderr, "specgen: 需在仓库根运行:", err)
		os.Exit(2)
	}
	count := 0
	for _, e := range entries {
		if !e.IsDir() || e.Name() == "internal" || e.Name() == "doc" || e.Name() == "multi-window" {
			continue
		}
		dir := filepath.Join("wg", "examples", e.Name())
		if _, err := os.Stat(filepath.Join(dir, "main.go")); err != nil {
			continue
		}
		if err := genOne(dir, *dry); err != nil {
			fmt.Fprintln(os.Stderr, "specgen:", e.Name(), err)
			continue
		}
		count++
	}
	fmt.Printf("specgen: %d 个功能窗已补齐 L1（dry-run=%v）\n", count, *dry)
}

func genOne(dir string, dry bool) error {
	out, err := exec.Command("go", "run", "./"+filepath.ToSlash(dir), "--meta").Output()
	if err != nil {
		return fmt.Errorf("--meta 失败: %w", err)
	}
	var m metaOut
	if err := json.Unmarshal(out, &m); err != nil {
		return fmt.Errorf("--meta 输出非法: %w", err)
	}

	specPath := filepath.Join(dir, "spec.json")
	spec := map[string]any{}
	if data, err := os.ReadFile(specPath); err == nil {
		if err := json.Unmarshal(data, &spec); err != nil {
			return fmt.Errorf("现有 spec.json 解析失败: %w", err)
		}
	}
	changed := false
	if _, ok := spec["id"]; !ok {
		spec["id"] = m.ID
		changed = true
	}
	st, _ := spec["structure"].(map[string]any)
	if st == nil {
		st = map[string]any{}
		spec["structure"] = st
	}
	if _, ok := st["sections"]; !ok && len(m.Demos) > 0 {
		st["sections"] = m.Demos
		changed = true
	}
	if _, ok := st["apiRows"]; !ok && m.APIRows > 0 {
		st["apiRows"] = m.APIRows
		changed = true
	}
	if !changed && !dry {
		return nil // 已有完整 structure：不重写（保护手写格式与键序）
	}

	merged, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return err
	}
	if dry {
		fmt.Printf("--- %s %s\n", m.ID, string(merged))
		return nil
	}
	return os.WriteFile(specPath, append(merged, '\n'), 0o644)
}