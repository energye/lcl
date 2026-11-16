// Command wg-showcase 用新引擎重建 showcase 展示台，无窗口 headless 出帧。
//
// 运行：go run ./wg/cmd/wg-showcase
// 它对应旧示例 examples/wg/showcase 的核心场景：9 卡片 + 5 类图表。
package main

import (
	"fmt"
	"os"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/migration"
)

func main() {
	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(1440, 880)
	c.MountWidget(migration.NewShowcase())

	stats, err := c.FrameOnce()
	if err != nil {
		fmt.Fprintln(os.Stderr, "出帧失败:", err)
		os.Exit(1)
	}
	fmt.Printf("wg-showcase: 帧=%d 节点=%d 层=%d 指令=%d 降级=%d 用时=%v\n",
		stats.Index+1, stats.Nodes, stats.Layers, stats.RasterOps,
		stats.Degraded, stats.FrameDuration)
}
