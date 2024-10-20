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

// IContainedAction Parent: IBasicAction
type IContainedAction interface {
	IBasicAction
	ActionList() ICustomActionList          // property
	SetActionList(AValue ICustomActionList) // property
	Index() int32                           // property
	SetIndex(AValue int32)                  // property
	Category() string                       // property
	SetCategory(AValue string)              // property
}

// TContainedAction Parent: TBasicAction
type TContainedAction struct {
	TBasicAction
}

func NewContainedAction(AOwner IComponent) IContainedAction {
	r1 := containedActionImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsContainedAction(r1)
}

func (m *TContainedAction) ActionList() ICustomActionList {
	r1 := containedActionImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsCustomActionList(r1)
}

func (m *TContainedAction) SetActionList(AValue ICustomActionList) {
	containedActionImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TContainedAction) Index() int32 {
	r1 := containedActionImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TContainedAction) SetIndex(AValue int32) {
	containedActionImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TContainedAction) Category() string {
	r1 := containedActionImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TContainedAction) SetCategory(AValue string) {
	containedActionImportAPI().SysCallN(1, 1, m.Instance(), PascalStr(AValue))
}

func ContainedActionClass() TClass {
	ret := containedActionImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	containedActionImport       *imports.Imports = nil
	containedActionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ContainedAction_ActionList", 0),
		/*1*/ imports.NewTable("ContainedAction_Category", 0),
		/*2*/ imports.NewTable("ContainedAction_Class", 0),
		/*3*/ imports.NewTable("ContainedAction_Create", 0),
		/*4*/ imports.NewTable("ContainedAction_Index", 0),
	}
)

func containedActionImportAPI() *imports.Imports {
	if containedActionImport == nil {
		containedActionImport = NewDefaultImports()
		containedActionImport.SetImportTable(containedActionImportTables)
		containedActionImportTables = nil
	}
	return containedActionImport
}
