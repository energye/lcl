//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// IActionList Parent: ICustomActionList
type IActionList interface {
	ICustomActionList
	SetOnChange(fn TNotifyEvent)  // property event
	SetOnExecute(fn TActionEvent) // property event
	SetOnUpdate(fn TActionEvent)  // property event
}

type TActionList struct {
	TCustomActionList
}

func (m *TActionList) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 1, actionListAPI(), api.MakeEventDataPtr(cb))
}

func (m *TActionList) SetOnExecute(fn TActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTActionEvent(fn)
	base.SetEvent(m, 2, actionListAPI(), api.MakeEventDataPtr(cb))
}

func (m *TActionList) SetOnUpdate(fn TActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTActionEvent(fn)
	base.SetEvent(m, 3, actionListAPI(), api.MakeEventDataPtr(cb))
}

// NewActionList class constructor
func NewActionList(owner IComponent) IActionList {
	r := actionListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsActionList(r)
}

func TActionListClass() types.TClass {
	r := actionListAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	actionListOnce   base.Once
	actionListImport *imports.Imports = nil
)

func actionListAPI() *imports.Imports {
	actionListOnce.Do(func() {
		actionListImport = api.NewDefaultImports()
		actionListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TActionList_Create", 0), // constructor NewActionList
			/* 1 */ imports.NewTable("TActionList_OnChange", 0), // event OnChange
			/* 2 */ imports.NewTable("TActionList_OnExecute", 0), // event OnExecute
			/* 3 */ imports.NewTable("TActionList_OnUpdate", 0), // event OnUpdate
			/* 4 */ imports.NewTable("TActionList_TClass", 0), // function TActionListClass
		}
	})
	return actionListImport
}
