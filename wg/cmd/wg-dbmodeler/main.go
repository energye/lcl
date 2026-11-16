// Command wg-dbmodeler 用新引擎重建 dbmodeler 建模页，headless 出帧。
//
// 运行：go run ./wg/cmd/wg-dbmodeler
// 对应旧示例 examples/wg/dbmodeler 的核心场景：表结构 + 关系画布。
package main

import (
	"fmt"
	"os"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/migration"
)

func main() {
	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(1100, 720)
	m := migration.NewDBModeler()
	c.MountWidget(m)

	stats, err := c.FrameOnce()
	if err != nil {
		fmt.Fprintln(os.Stderr, "出帧失败:", err)
		os.Exit(1)
	}
	fmt.Printf("wg-dbmodeler: 帧=%d 表=%d 字段行=%d 节点=%d 层=%d 用时=%v\n",
		stats.Index+1, len(m.Tables), len(m.Table.Rows),
		stats.Nodes, stats.Layers, stats.FrameDuration)
}
