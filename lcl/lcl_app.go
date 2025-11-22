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
	"reflect"
	"runtime"

	"github.com/energye/lcl/api"

	"github.com/energye/lcl/types"

	"github.com/energye/lcl/base"
)

type IApp interface {
	IApplication
	Initialize()
	NewForms(forms ...IEngForm)
	NewForm(form IEngForm)
	Run()
}

type TApp struct {
	TApplication
}

func (m *TApp) NewForms(forms ...IEngForm) {
	for _, form := range forms {
		m.NewForm(form)
	}
}

func (m *TApp) Run() {
	defer func() {
		for i := int(Application.ComponentCount()) - 1; i >= 0; i-- {
			comp := Application.Components(int32(i))
			comp.Free()
		}
		if runtime.GOOS == "linux" {
			api.WidgetSetFinalization()
		}
		api.LibRelease()
	}()
	m.TApplication.Run()
}

func (m *TApp) Initialize() {
	if runtime.GOOS == "linux" {
		api.WidgetSetInitialization()
	}
	m.TApplication.Initialize()
	api.Application_Initialize(m.Instance())
}

func (m *TApp) NewForm(form IEngForm) {
	var (
		result uintptr
		class  = AsComponent(TEngFormClass())
	)
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
	// Create new form
	m.CreateForm(class, &result)
	form.SetInstance(base.UnsafePointer(result))
	if !isMain {
		form.SetGoPtr(createParamsPtr)
	}
	// call OnAfterCreate
	if fn, ok := form.(IOnAfterCreate); ok {
		fn.FormAfterCreate(form)
	}
	// 注册一些默认事件
	if fn, ok := form.(IOnCloseQuery); ok {
		form.SetOnCloseQuery(func(sender IObject, canClose *bool) {
			fn.OnCloseQuery(sender, canClose)
		})
	}
	if fn, ok := form.(IOnClose); ok {
		form.SetOnClose(func(sender IObject, closeAction *types.TCloseAction) {
			fn.OnClose(sender, closeAction)
		})
	}
}

// AsApp Convert a pointer object to an existing class object
func AsApp(obj any) IApp {
	var result *TApp
	instance := base.GetInstance(obj)
	if instance == nil {
		return result
	}
	result = new(TApp)
	base.SetObjectInstance(result, instance)
	return result
}

// RunApp
//
// simplify running.
func RunApp(forms ...IEngForm) {
	Application.Initialize()
	Application.SetMainFormOnTaskBar(true)
	for i := 0; i < len(forms); i++ {
		Application.NewForm(forms[i])
	}
	Application.Run()
}
