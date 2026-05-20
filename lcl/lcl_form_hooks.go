// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

// :predefine:

package lcl

import (
	"github.com/energye/lcl/types"
)

type IFormHook interface{}

type IFormCreateHook interface {
	OnFormCreate(sender IObject)
}

type IFormShowHook interface {
	OnFormShow(sender IObject)
}

type IFormCloseQueryHook interface {
	OnFormCloseQuery(sender IObject, canClose *bool) bool
}

type IFormCloseHook interface {
	OnFormClose(sender IObject, closeAction *types.TCloseAction) bool
}

func CallFormCreate(target IFormHook, sender IObject) {
	if target != nil {
		if hook, ok := target.(IFormCreateHook); ok {
			hook.OnFormCreate(sender)
		}
	}
}

func CallFormShow(target IFormHook, sender IObject) {
	if target != nil {
		if hook, ok := target.(IFormShowHook); ok {
			hook.OnFormShow(sender)
		}
	}
}

func CallFormCloseQuery(target IFormHook, sender IObject, canClose *bool) bool {
	if target != nil {
		if hook, ok := target.(IFormCloseQueryHook); ok {
			return hook.OnFormCloseQuery(sender, canClose)
		}
	}
	return false
}

func CallFormClose(target IFormHook, sender IObject, closeAction *types.TCloseAction) bool {
	if target != nil {
		if hook, ok := target.(IFormCloseHook); ok {
			return hook.OnFormClose(sender, closeAction)
		}
	}
	return false
}
