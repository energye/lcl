//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package tool

import (
	"os"
	"runtime"
)

const (
	isWindows   = runtime.GOOS == "windows"   //support
	isLinux     = runtime.GOOS == "linux"     //support
	isDarwin    = runtime.GOOS == "darwin"    //support
	isAndroid   = runtime.GOOS == "android"   //not support
	isIos       = runtime.GOOS == "ios"       //not support
	isPlan9     = runtime.GOOS == "plan9"     //not support
	isAix       = runtime.GOOS == "aix"       //not support
	isDragonfly = runtime.GOOS == "dragonfly" //not support
	isFreebsd   = runtime.GOOS == "freebsd"   //not support
	isHurd      = runtime.GOOS == "hurd"      //not support
	isIllumos   = runtime.GOOS == "illumos"   //not support
	isJs        = runtime.GOOS == "js"        //not support
	isNacl      = runtime.GOOS == "nacl"      //not support
	isNetbsd    = runtime.GOOS == "netbsd"    //not support
	isOpenbsd   = runtime.GOOS == "openbsd"   //not support
	isSolaris   = runtime.GOOS == "solaris"   //not support
	isZos       = runtime.GOOS == "zos"       //not support
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func MkdirAll(path string) {
	var err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func IsWindows() bool {
	return isWindows
}

func IsLinux() bool {
	return isLinux
}

func IsDarwin() bool {
	return isDarwin
}

func IsPlan9() bool {
	return isPlan9
}

func ToInt(value any) int {
	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	default:
		return 0
	}
}

func ToDouble(value any) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}
