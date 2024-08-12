//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/types"
	"sync"
)

var (
	createParamsMap sync.Map
	newFormCreate   IForm
)

func addNewFormCreate(form IForm) {
	newFormCreate = form
}

func addRequestCreateParamsMap(ptr uintptr, proc IForm) {
	createParamsMap.Store(ptr, proc)
}

func requestCallCreateParamsCallbackProc(ptr uintptr, sender, params uintptr) uintptr {
	if val, ok := createParamsMap.Load(ptr); ok {
		if form, ok := val.(IForm); ok {
			form.SetInstance(unsafePointer(sender))
			switch form.(type) {
			case IOnCreateParams:
				form.(IOnCreateParams).CreateParams((*types.TCreateParams)(unsafePointer(params)))
			}
		}
		createParamsMap.Delete(ptr)
	}
	return 0
}

func requestCallFormCreateCallbackProc(ptr uintptr, sender uintptr) uintptr {
	if newFormCreate != nil {
		currentForm := newFormCreate
		newFormCreate = nil
		currentForm.SetInstance(unsafePointer(sender))
		switch currentForm.(type) {
		case IOnCreate:
			currentForm.(IOnCreate).FormCreate(currentForm)
		}
	}
	return 0
}
