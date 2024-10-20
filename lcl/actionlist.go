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
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IActionList Parent: ICustomActionList
type IActionList interface {
	ICustomActionList
	SetOnChange(fn TNotifyEvent)  // property event
	SetOnExecute(fn TActionEvent) // property event
	SetOnUpdate(fn TActionEvent)  // property event
}

// TActionList Parent: TCustomActionList
type TActionList struct {
	TCustomActionList
	changePtr  uintptr
	executePtr uintptr
	updatePtr  uintptr
}

func NewActionList(AOwner IComponent) IActionList {
	r1 := actionListImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsActionList(r1)
}

func ActionListClass() TClass {
	ret := actionListImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TActionList) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	actionListImportAPI().SysCallN(2, m.Instance(), m.changePtr)
}

func (m *TActionList) SetOnExecute(fn TActionEvent) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	actionListImportAPI().SysCallN(3, m.Instance(), m.executePtr)
}

func (m *TActionList) SetOnUpdate(fn TActionEvent) {
	if m.updatePtr != 0 {
		RemoveEventElement(m.updatePtr)
	}
	m.updatePtr = MakeEventDataPtr(fn)
	actionListImportAPI().SysCallN(4, m.Instance(), m.updatePtr)
}

var (
	actionListImport       *imports.Imports = nil
	actionListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ActionList_Class", 0),
		/*1*/ imports.NewTable("ActionList_Create", 0),
		/*2*/ imports.NewTable("ActionList_SetOnChange", 0),
		/*3*/ imports.NewTable("ActionList_SetOnExecute", 0),
		/*4*/ imports.NewTable("ActionList_SetOnUpdate", 0),
	}
)

func actionListImportAPI() *imports.Imports {
	if actionListImport == nil {
		actionListImport = NewDefaultImports()
		actionListImport.SetImportTable(actionListImportTables)
		actionListImportTables = nil
	}
	return actionListImport
}
