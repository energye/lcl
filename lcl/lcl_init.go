//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import (
	"os"
	"runtime"

	"github.com/energye/lcl/api"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/internal/initialize"
)

var (
	// Application
	//
	//  The TApplication singleton.
	//  is a unit global variable with the
	//  instance representing the currently executing application.
	//  The value for the variable is assigned in the initialization section for the unit when the application is started.
	//  It is freed in the finalization section when the application is terminated.
	//  Use the
	//  TApplicationProperties
	//  component to provide design-time configuration settings for properties (including event handlers) which are automatically applied to Application at run-time.
	Application IApp
	Screen      IScreen    // 屏幕
	Mouse       IMouse     // 鼠标
	Clipboard   IClipboard // 剪切板
	Printer     IPrinter   // 打印机
)

// Init LCL
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	defer func() {
		if err := recover(); err != nil {
			println(err)
			os.Exit(1)
		}
	}()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// 初始化
	initialize.Initialize(libs, resources)

	// 注册 LCL 对象事件回调
	api.SetEventCallback(eventCallback, api.EctLCL)
	// 注册 LCL 对象事件回调移除
	api.SetEventCallback(removeEventCallback, api.EctLCLRemove)
	// 注册 消息事件回调
	api.SetEventCallback(messageCallback, api.EctMessage)
	// 注册 窗口创建时回调
	api.SetEventCallback(createParamsCallback, api.EctCreateParams)
	// 注册 创建窗口参数回调
	api.SetEventCallback(formCreateCallback, api.EctFormCreate)
	// 注册 主线程同步 回调
	api.SetEventCallback(threadSyncCallback, api.EctUIThreadSync)
	// 注册 主线程回调 异步
	api.SetEventCallback(threadAsyncCallback, api.EctUIThreadAsync)

	// 获取固定全局实例
	Application = AsApp(api.Application_Instance())
	Screen = AsScreen(api.Screen_Instance())
	Mouse = AsMouse(api.Mouse_Instance())
	Clipboard = AsClipboard(api.Clipboard_Instance())
	Printer = AsPrinter(api.Printer_Instance())
}
