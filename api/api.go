//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/libname"
)

// EventCallbackType 事件回调类型
// 普通事件: LCL, CEF, Webview2, Webkit2 事件
// 消息事件: 特殊消息
// 异常事件: 异常消息
// UI线程回调: UI线程同步, UI线程异步
// 窗口创建事件: FormCreate
// 创建参数事件: CreateParams
type EventCallbackType int32

const (
	EctLCL              EventCallbackType = iota // LCL 对象事件回调
	EctCEF                                       // CEF 对象事件回调
	EctWV                                        // Webview 对象事件回调
	EctLCLRemove                                 // LCL 对象事件回调移除
	EctCEFRemove                                 // CEF 对象事件回调移除
	EctWVRemove                                  // Webview 对象事件回调移除
	EctMessage                                   // 消息事件回调
	EctExceptionHandler                          // 异常事件回调
	EctUIThreadSync                              // 在UI线程运行事件 同步
	EctUIThreadAsync                             // 在UI线程运行事件 异步(实际是队列)
	EctFormCreate                                // 窗口创建时回调
	EctCreateParams                              // 创建窗口参数回调
)

var (
	lib imports.DLL // 导入 API
	// 释放回调函数
	releaseCallback func()
	// 加载Lib回调函数
	loadLibCallback func() (imports.DLL, error) // 自定义加载 lib 回调函数
)

// SetOnLoadLibCallback
//
//	设置用于加载lib动态库的回调函数
//	调用 inits.Init 之前设置
func SetOnLoadLibCallback(fn func() (imports.DLL, error)) {
	if loadLibCallback == nil {
		loadLibCallback = fn
	}
}

func loadLib() imports.DLL {
	if loadLibCallback != nil {
		lib, _ := loadLibCallback()
		return lib
	} else {
		lib, err := imports.NewDLL(libname.LibName)
		if lib == 0 && err != nil {
			panic(err)
		}
		return lib
	}
}

func closeLib() {
	if lib != 0 {
		lib.Release()
		lib = 0
	}
}

// NewDefaultImports 创建默认导入表调用
func NewDefaultImports() *imports.Imports {
	if lib == 0 {
		lib = loadLib()
	}
	result := new(imports.Imports)
	result.Dll = lib
	result.NextType()
	return result
}

// Init 初始化 dll API
func Init() {
	// 加载动态链接库
	lib = loadLib()
}

// LibRelease app run end
func LibRelease() {
	if releaseCallback != nil {
		releaseCallback()
	}
	WidgetSetFinalization()
	// 运行结束后就结束close掉lib，不然他不会关掉的
	closeLib()
}

// SetReleaseCallback 应用运行结束后释放资源之前执行
func SetReleaseCallback(fn func()) {
	if releaseCallback == nil {
		releaseCallback = fn
	}
}
