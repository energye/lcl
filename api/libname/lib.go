//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package libname

import (
	"github.com/energye/lcl/tools"
	"github.com/energye/lcl/tools/exec"
	"os"
	"path"
	"runtime"
)

const libName = "liblcl"

var (
	LibName          string // 加载完成后该变量被置空
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

func GetDLLName() string {
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}

// GetLibPath
//
//	获取 liblcl 动态库目录
//	加载优先级: LibName > 当前目录 > 环境变量(LCL_HOME | LCLCEF_HOME | LCLWV2_HOME | LCLWK2_HOME | ENERGY_HOME) > 用户目录(.energy)文件内读取
func GetLibPath(dllName string) string {
	if LibName != "" && tools.IsExist(LibName) {
		return LibName
	}
	//当前执行文件目录
	if exeDir := path.Join(exec.Dir, dllName); tools.IsExist(exeDir) {
		return exeDir
	}
	var checkenv = func(env string) (string, bool) {
		if env = path.Join(os.Getenv(env), dllName); tools.IsExist(env) {
			return env, true
		}
		return "", false
	}
	//环境变量
	if libPath, ok := checkenv("LCL_HOME"); ok {
		return libPath
	}
	if libPath, ok := checkenv("LCLCEF_HOME"); ok {
		return libPath
	}
	if libPath, ok := checkenv("LCLWV2_HOME"); ok {
		return libPath
	}
	if libPath, ok := checkenv("LCLWK2_HOME"); ok {
		return libPath
	}
	if libPath, ok := checkenv("ENERGY_HOME"); ok {
		return libPath
	}
	return ""
}
