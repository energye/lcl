// Command wg-imedemo 用新引擎重建 canvas_ime 输入法演示，headless 出帧。
//
// 运行：go run ./wg/cmd/wg-imedemo
// 对应旧示例 examples/wg/canvas_ime 的核心场景：自绘输入框 + 按钮 + 状态面板。
package main

import (
	"fmt"
	"os"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/migration"
)

func main() {
	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(960, 640)
	m := migration.NewCanvasIME()
	c.MountWidget(m)

	// 走一遍 C28 文本链路：组合串上屏 → 中文替换。
	m.Submit.OnClick()

	stats, err := c.FrameOnce()
	if err != nil {
		fmt.Fprintln(os.Stderr, "出帧失败:", err)
		os.Exit(1)
	}
	fmt.Printf("wg-imedemo: 帧=%d 节点=%d 文本行=%d 输入=%q 用时=%v\n",
		stats.Index+1, stats.Nodes, stats.TextLines,
		m.Input.Editor().Value().Text, stats.FrameDuration)
}
