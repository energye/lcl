//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"reflect"
)

// CreateForm Create a TForm.
func (m *TApplication) CreateForm(forms ...IForm) IForm {
	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()
	size := len(forms)
	if size == 0 {
		return AsForm(Application_CreateForm(m.Instance()))
	}
	for i := 0; i < size; i++ {
		form := forms[i]
		var (
			mainForm        = Application.MainForm()
			isMain          = mainForm == nil || mainForm.Instance() == 0 // 0 | nil = main
			createParamsPtr uintptr
		)
		if !isMain {
			// CreateParamsCallback
			createParamsPtr = reflect.ValueOf(form).Pointer()
		}
		// OnCreate
		if _, ok := form.(IOnCreate); ok {
			addNewFormCreate(form)
		}
		// CreateParams
		if _, ok := form.(IOnCreateParams); ok {
			addRequestCreateParamsMap(createParamsPtr, form)
		}
		formPtr := Application_CreateForm(m.Instance())
		form.SetInstance(unsafePointer(formPtr))
		if !isMain {
			Form_SetGoPtr(formPtr, createParamsPtr)
		}
		// OnAfterCreate
		if fn, ok := form.(IOnAfterCreate); ok {
			fn.FormAfterCreate(form)
		}
	}
	return nil
}

// Run the app.
func (m *TApplication) Run() {
	Application_Run(m.Instance())
}

// Initialize
//
// Initial APP information.
func (m *TApplication) Initialize() {
	CustomWidgetSetInitialization()
	Application_Initialize(m.Instance())
}

// SetRunLoopReceived 这里只是测试，实际Go并未用得着他
func (m *TApplication) SetRunLoopReceived(proc uintptr) {
	Application_SetRunLoopReceived(m.Instance(), proc)
}
