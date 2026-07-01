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
	"sync"

	"github.com/energye/lcl/api"
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
	Screen      IScreen    // Screen
	Mouse       IMouse     // Mouse
	Clipboard   IClipboard // Clipboard
	Printer     IPrinter   // Printer
	initOnce    sync.Once  //
)

func init() {
	//runtime.LockOSThread() // macOS
}

// Init LCL
func Init() {
	runtime.LockOSThread()
	api.SetOnReleaseCallback(func() {
		runtime.UnlockOSThread()
	})
	initOnce.Do(func() {
		defer func() {
			if err := recover(); err != nil {
				println(err)
				os.Exit(1)
			}
		}()

		// Initialize runtime libenergy
		initialize.Initialize()

		// Register LCL object event callbacks
		api.SetEventCallback(eventCallback, api.EctLCL)
		// Unregister LCL object event callbacks
		api.SetEventCallback(removeEventCallback, api.EctLCLRemove)
		// Register message event callbacks
		api.SetEventCallback(messageCallback, api.EctMessage)
		// Register window creation callback
		api.SetEventCallback(createParamsCallback, api.EctCreateParams)
		// Register window creation parameters callback
		api.SetEventCallback(formCreateCallback, api.EctFormCreate)
		// Register main thread synchronization callback
		api.SetEventCallback(threadSyncCallback, api.EctUIThreadSync)
		// Register asynchronous main thread callback
		api.SetEventCallback(threadAsyncCallback, api.EctUIThreadAsync)

		// Get fixed global instance
		Application = AsApp(api.Application_Instance())
		Screen = AsScreen(api.Screen_Instance())
		Mouse = AsMouse(api.Mouse_Instance())
		Clipboard = AsClipboard(api.Clipboard_Instance())
		Printer = AsPrinter(api.Printer_Instance())
	})
}
