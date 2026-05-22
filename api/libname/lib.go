//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package libname

import (
	"fmt"
	"runtime"
)

const DarwinUniversalBinaryName = "libenergy-universal.dylib" // macOS 通用二进制库名称

var (
	LibName               string // 加载完成后该变量可被置空
	UseWS                 string // 用于控制后台组件: win32/cocoa/gtk2/gtk3
	EnableUniversalBinary bool   // 用于启用通用二进制库: macOS 10.15.x
)

// GetDLLName 用于获取当前系统架构的 lib 库
func GetDLLName() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	ws, ext := "", ""
	switch goos {
	case "darwin":
		ext = "dylib"
	case "linux":
		ext = "so"
		if UseWS == "gtk3" {
			ws = "gtk3" // use gtk3
		} else {
			ws = "gtk2"
		}
		if len(ws) > 0 && ws[0] != '-' {
			ws = "-" + ws // add first str "-"
		}
	case "windows":
		ext = "dll"
	}
	// windows, macOS: dev - libenergy-[arch].xx | prod - libenergy.xx
	// linux: dev - libenergy-[arch]-[ws].xx | prod - libenergy-[ws].xx
	var name string
	if IsDev {
		name = fmt.Sprintf("libenergy-%s%s.%s", goarch, ws, ext)
	} else {
		name = fmt.Sprintf("libenergy%s.%s", ws, ext)
	}
	return name
}

// LibPath
//
//	获取 libenergy 动态库目录
//	加载优先级: LibName > 当前目录 > 环境变量(LCL_HOME) > 用户目录 (.energy) 内读取
//func LibPath(dllName string) string {
//	if LibName != "" && tool.IsExist(LibName) {
//		return LibName
//	}
//	//当前执行文件目录
//	if exeDir := path.Join(exec.Dir, dllName); tool.IsExist(exeDir) {
//		return exeDir
//	}
//	var checkenv = func(env string) (string, bool) {
//		if env = path.Join(os.Getenv(env), dllName); tool.IsExist(env) {
//			return env, true
//		}
//		return "", false
//	}
//	//环境变量
//	if libPath, ok := checkenv("LCL_HOME"); ok {
//		return libPath
//	}
//	return ""
//}
