//go:build darwin && !prod
// +build darwin,!prod

package gen

import (
	"github.com/energye/lcl/config"
	"github.com/energye/lcl/internal/initialize/macapp"
	"os"
)

func GenMacOSApp(exeFileDir, cefSubProcess string) {
	// 生成 macOS app
	// 设置执行文件绝对目录
	os.Args[0] = exeFileDir
	config.Get().IsCEF = true
	macapp.MacApp.SetBrowseSubprocessPath(cefSubProcess)
	macapp.Init()
}
