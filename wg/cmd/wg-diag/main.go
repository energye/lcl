// Command wg-diag 是 WG Engine 的诊断入口。
//
// 用法：
//
//	wg-diag --selftest                      # 契约 + 几何 + 缓冲 + 画布树 + 时钟 + 空帧
//	wg-diag --contracts                     # 打印 C1–C28 契约表与当期状态
//	wg-diag --headless --frames 120          # ManualClock 确定性出帧统计
//	wg-diag --stage M2 --selftest            # 以 M2 门禁标准检查（未实现断言会报 MISSING）
//	wg-diag --layout --layers --hit          # 布局/层树/命中
//	wg-diag --gesturetest --inputtest --overlaytest  # 交互链路（含浮层与组件基座）
//
// 全量门禁命令矩阵见 docs/wg-engine/05-API-ROADMAP.md §6。参数解析留在 main，
// 真正的矩阵实现放在 diag 包，便于 go test 直接调用。
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/contracts"
	"github.com/energye/lcl/wg/diag"
	"github.com/energye/lcl/wg/frame"
)

func main() {
	var (
		selftest  = flag.Bool("selftest", false, "运行自检（默认行为）")
		matrix    = flag.String("render-matrix", "", "20 图元 × 2 后端的位置比对；给目录则同时出图")
		golden    = flag.String("golden", "", "与参考图比对（目录）；首次用 -golden-update 生成")
		goldenUp  = flag.Bool("golden-update", false, "写入参考图（必须人工 review 后提交）")
		benchFlag = flag.Bool("bench", false, "1440×880 无窗口性能基准")
		damage    = flag.Bool("damage", false, "局部重绘面积断言")
		layoutF   = flag.Bool("layout", false, "20 组约束布局断言")
		layersF   = flag.Bool("layers", false, "16 组层树/缓存断言")
		schedF    = flag.Bool("schedtest", false, "14 组帧调度断言")
		hitF      = flag.Bool("hit", false, "16 组命中断言（变换/裁剪/重叠）")
		gestF     = flag.Bool("gesturetest", false, "6 组手势冲突断言（单指针单胜者）")
		inputF    = flag.Bool("inputtest", false, "22 项输入链路断言（归一化/文本/IME/捕获自愈）")
		resizeF   = flag.Bool("resizetest", false, "7 次尺寸切换断言（去抖 + 池释放顺序）")
		ovlF      = flag.Bool("overlaytest", false, "16 组浮层/组件基座断言（定位/关闭/遮罩/样式/动画）")
		idleF     = flag.Bool("idle", false, "1000 节点静止 0 帧（C14 门禁）")
		keysF     = flag.Bool("keyselftest", false, "14 项文本/按键分离断言（C28；中文上屏）")
		showCtr   = flag.Bool("contracts", false, "打印契约表")
		headless  = flag.Bool("headless", false, "无窗口确定性出帧")
		frames    = flag.Int("frames", 60, "--headless 的帧数")
		stage     = flag.String("stage", "", "按指定里程碑门禁检查（M0–M5）")
		width     = flag.Int("w", 1440, "--headless 的宽度（像素）")
		height    = flag.Int("h", 880, "--headless 的高度（像素）")
		version   = flag.Bool("version", false, "打印版本")
		quiet     = flag.Bool("quiet", false, "只打印汇总，不逐项输出")
	)
	flag.Parse()

	if *version {
		fmt.Printf("wg-diag %s (stage %s)\n", wg.Version, wg.Stage)
		return
	}

	if *stage != "" {
		s, err := parseStage(*stage)
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag:", err)
			os.Exit(2)
		}
		contracts.SetStage(s)
		fmt.Printf("门禁里程碑: %s\n", s)
	}

	ran := false
	exit := 0

	if *showCtr {
		ran = true
		printContracts(*quiet)
		if s := contracts.Summarize(contracts.CheckAll()); !s.OK() {
			exit = 1
		}
	}

	if *headless {
		ran = true
		if err := runHeadless(*width, *height, *frames); err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --headless:", err)
			exit = 1
		}
	}

	if *matrix != "" {
		ran = true
		rep := diag.RenderMatrix()
		fmt.Print(rep.String())
		if *matrix != "-" {
			if err := diag.WriteMatrixImages(*matrix); err != nil {
				fmt.Fprintln(os.Stderr, "wg-diag --render-matrix 出图:", err)
				exit = 1
			} else {
				fmt.Printf("出图目录: %s\n", *matrix)
			}
		}
		if !rep.OK() {
			exit = 1
		}
	}

	if *golden != "" {
		ran = true
		res, err := diag.Golden(*golden, *goldenUp)
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --golden:", err)
			exit = 1
		}
		fmt.Println(res.String())
		if !res.OK() {
			exit = 1
		}
	}

	if *benchFlag {
		ran = true
		res, err := diag.Bench(diag.BenchOptions{Width: *width, Height: *height, Frames: *frames})
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --bench:", err)
			exit = 1
		} else {
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "bench 未达门禁（≥59 FPS）\n")
				exit = 1
			}
		}
	}

	if *layoutF {
		ran = true
		res, err := diag.Layout()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --layout:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-32s 得到 %s 期望 %s", mark, d.Name, d.Got, d.Want)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "layout 门禁未通过（需 ≥20 组全绿）\n")
				exit = 1
			}
		}
	}

	if *layersF {
		ran = true
		res, err := diag.Layers()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --layers:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-34s %s", mark, d.Name, d.Got)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "layers 门禁未通过（需 ≥16 组全绿）\n")
				exit = 1
			}
		}
	}

	if *schedF {
		ran = true
		res, err := diag.SchedTest()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --schedtest:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-30s %s", mark, d.Name, d.Got)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "schedtest 门禁未通过（需 ≥14 组全绿）\n")
				exit = 1
			}
		}
	}

	if *hitF {
		ran = true
		res, err := diag.Hit()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --hit:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-32s %s", mark, d.Name, d.Got)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "hit 门禁未通过（需 ≥12 组全绿）\n")
				exit = 1
			}
		}
	}

	if *gestF {
		ran = true
		res, err := diag.GestureTest()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --gesturetest:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-32s %s", mark, d.Name, d.Got)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "gesturetest 门禁未通过（需 ≥6 组全绿）\n")
				exit = 1
			}
		}
	}

	if *inputF {
		ran = true
		res, err := diag.InputTest()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --inputtest:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-28s %s", mark, d.Name, d.Got)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "inputtest 门禁未通过（需 ≥22 项全绿）\n")
				exit = 1
			}
		}
	}

	if *resizeF {
		ran = true
		res, err := diag.ResizeTest()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --resizetest:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-28s %s", mark, d.Name, d.Got)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "resizetest 门禁未通过（需 ≥7 次全绿）\n")
				exit = 1
			}
		}
	}

	if *ovlF {
		ran = true
		res, err := diag.OverlayTest()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --overlaytest:", err)
			exit = 1
		} else {
			if !*quiet {
				for _, d := range res.Details {
					mark := "ok  "
					if !d.OK {
						mark = "FAIL"
					}
					fmt.Printf("%s %-32s %s", mark, d.Name, d.Got)
					if d.Note != "" {
						fmt.Printf("  ← %s", d.Note)
					}
					fmt.Println()
				}
			}
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "overlaytest 门禁未通过（需 ≥16 组全绿）\n")
				exit = 1
			}
		}
	}

	if *idleF {
		ran = true
		res, err := diag.IdleTest()
		if !*quiet {
			fmt.Printf("节点=%d 静止用帧=%d 静止时钟片=%d\n",
				res.Nodes, res.SettleFrames, res.IdleTicks)
		}
		fmt.Println(res.String())
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --idle:", err)
			exit = 1
		} else if !res.OK() {
			fmt.Fprintf(os.Stderr, "idle 门禁未通过（需 1000 节点静止 0 帧）\n")
			exit = 1
		}
	}

	if *keysF {
		ran = true
		res, err := diag.KeyTest()
		if !*quiet {
			for _, d := range res.Details {
				mark := "ok  "
				if !d.OK {
					mark = "FAIL"
				}
				fmt.Printf("%s %-28s %s", mark, d.Name, d.Got)
				if d.Note != "" {
					fmt.Printf("  ← %s", d.Note)
				}
				fmt.Println()
			}
		}
		fmt.Println(res.String())
		if err != nil || !res.OK() {
			fmt.Fprintln(os.Stderr, "keyselftest 门禁未通过（C28：文本与按键分离）")
			exit = 1
		}
	}

	if *damage {
		ran = true
		res, err := diag.Damage()
		if err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --damage:", err)
			exit = 1
		} else {
			fmt.Println(res.String())
			if !res.OK() {
				fmt.Fprintf(os.Stderr, "damage 面积超出按钮包围盒 1.5 倍\n")
				exit = 1
			}
		}
	}

	// 无参数时默认跑自检；显式 --selftest 也跑。
	if *selftest || !ran {
		if err := runSelfTest(*quiet); err != nil {
			fmt.Fprintln(os.Stderr, "wg-diag --selftest:", err)
			exit = 1
		}
	}

	os.Exit(exit)
}

func parseStage(s string) (contracts.Stage, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "M0":
		return contracts.M0, nil
	case "M1":
		return contracts.M1, nil
	case "M2":
		return contracts.M2, nil
	case "M3":
		return contracts.M3, nil
	case "M4":
		return contracts.M4, nil
	case "M5":
		return contracts.M5, nil
	default:
		return contracts.M0, fmt.Errorf("未知里程碑 %q（可选 M0–M5）", s)
	}
}

// runSelfTest 执行自检并输出结果。
func runSelfTest(quiet bool) error {
	began := time.Now()
	results := wg.SelfTest()

	groups := make([]string, 0, 8)
	seen := map[string]bool{}
	for _, r := range results {
		if !seen[r.Group] {
			seen[r.Group] = true
			groups = append(groups, r.Group)
		}
	}

	if !quiet {
		fmt.Printf("wg-diag selftest  %s (stage %s)\n\n", wg.Version, wg.Stage)
		for _, g := range groups {
			fmt.Printf("── %s\n", g)
			for _, r := range results {
				if r.Group != g {
					continue
				}
				fmt.Println("  " + r.String())
			}
		}
		fmt.Println()
	}

	fails := 0
	for _, r := range results {
		if !r.OK {
			fails++
		}
	}
	fmt.Printf("selftest: %d 项，失败 %d，用时 %.2fms\n",
		len(results), fails, float64(time.Since(began).Microseconds())/1000)
	if fails > 0 {
		if r, ok := wg.FirstFailure(results); ok {
			return fmt.Errorf("%s / %s: %s", r.Group, r.Name, r.Detail)
		}
		return fmt.Errorf("%d 项失败", fails)
	}
	return nil
}

// printContracts 打印契约表。
func printContracts(quiet bool) {
	st := contracts.Stage(contracts.Current)
	results := contracts.CheckAll()
	sort.SliceStable(results, func(i, j int) bool { return results[i].Since < results[j].Since })

	fmt.Printf("契约表（当前门禁 %s）\n", st)
	for _, r := range results {
		if quiet && r.Status == contracts.Pending {
			continue
		}
		fmt.Println("  " + r.String())
	}
	s := contracts.Summarize(results)
	fmt.Printf("契约汇总: %s\n\n", s)
}

// runHeadless 用 ManualClock 跑确定性帧序列。
func runHeadless(w, h, frames int) error {
	if frames <= 0 {
		return fmt.Errorf("--frames 必须 > 0")
	}
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(w, h)

	clk := frame.NewManualClock(0, frame.DefaultStep)
	c.SetClock(clk)

	deltas := make([]time.Duration, 0, frames)
	dirty := 0
	var last frame.Stats
	for i := 0; i < frames; i++ {
		st, err := c.FrameOnce()
		if err != nil {
			return fmt.Errorf("第 %d 帧: %w", i, err)
		}
		if st.Dirty {
			dirty++
		}
		if st.Delta > 0 {
			deltas = append(deltas, st.Delta)
		}
		last = st
		// M0 没有光栅与呈现，FrameOnce 不清脏区。这里手动清一次，
		// 模拟"该帧已经提交"，否则后续帧会被同一块脏区一直拖着，
		// 帧统计就失去意义（M1 接入光栅后由合成器在提交后清除）。
		if root := c.Root(); root != nil {
			root.ClearDamageTree()
			root.ClearStructureDirty()
		}
		clk.Advance()
	}

	fmt.Printf("headless: %dx%d, %d 帧, 脏帧 %d\n", w, h, frames, dirty)
	fmt.Printf("  末帧: %s\n", last)
	fmt.Printf("  帧间隔中位数: %v\n", frame.Median(deltas))
	if p95 := frame.Percentile(deltas, 95); p95 > 0 {
		fmt.Printf("  帧间隔 P95  : %v\n", p95)
	}
	return nil
}
