// Command checkdeps 是契约 C17 的独立入口：静态检查 wg 的 import 拓扑。
//
// 用法：
//
//	go run ./wg/internal/checkdeps          # 检查并通过退出码回报
//	go run ./wg/internal/checkdeps -v       # 打印规则表与每个包的实际依赖
//
// 退出码：0 = 无违规；1 = 有违规；2 = 检查器自身失败。
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/energye/lcl/wg/internal/deps"
)

func main() {
	verbose := flag.Bool("v", false, "打印规则表与统计")
	root := flag.String("root", "", "wg 根目录（默认从源码位置推导）")
	flag.Parse()

	dir := *root
	if dir == "" {
		dir = wgRoot()
	}

	v, err := deps.Check(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "checkdeps:", err)
		os.Exit(2)
	}

	if *verbose {
		fmt.Println("wg 根目录:", dir)
		fmt.Println("规则表:")
		for _, r := range deps.Rules() {
			fmt.Println("  ", r)
		}
		fmt.Println()
	}

	if len(v) == 0 {
		fmt.Println("checkdeps: OK（C17 单向依赖无违规）")
		return
	}
	fmt.Fprintf(os.Stderr, "checkdeps: %d 处违反 C17\n", len(v))
	for _, item := range v {
		fmt.Fprintln(os.Stderr, "  ", item.String())
	}
	os.Exit(1)
}

// wgRoot 从本文件位置推导 wg 包根目录：
// <root>/internal/checkdeps/main.go → <root>。
func wgRoot() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filepath.Dir(file)))
}
