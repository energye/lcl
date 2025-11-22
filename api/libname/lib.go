//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package libname

import (
	"runtime"
)

const libName = "liblcl"

var LibName string // 加载完成后该变量可被置空

// GetDLLName 用于获取当前系统架构的 lib 库
func GetDLLName() string {
	platformExtNames := map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}

// LibPath
//
//	获取 liblcl 动态库目录
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
