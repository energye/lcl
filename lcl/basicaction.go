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

// IBasicAction Parent: IComponent
type IBasicAction interface {
	IComponent
	ActionComponent() IComponent              // property
	SetActionComponent(AValue IComponent)     // property
	HandlesTarget(Target IObject) bool        // function
	Execute() bool                            // function
	Update() bool                             // function
	UpdateTarget(Target IObject)              // procedure
	ExecuteTarget(Target IObject)             // procedure
	RegisterChanges(Value IBasicActionLink)   // procedure
	UnRegisterChanges(Value IBasicActionLink) // procedure
	SetOnExecute(fn TNotifyEvent)             // property event
	SetOnUpdate(fn TNotifyEvent)              // property event
}

// TBasicAction Parent: TComponent
type TBasicAction struct {
	TComponent
	executePtr uintptr
	updatePtr  uintptr
}

func NewBasicAction(AOwner IComponent) IBasicAction {
	r1 := basicActionImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsBasicAction(r1)
}

func (m *TBasicAction) ActionComponent() IComponent {
	r1 := basicActionImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TBasicAction) SetActionComponent(AValue IComponent) {
	basicActionImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBasicAction) HandlesTarget(Target IObject) bool {
	r1 := basicActionImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Target))
	return GoBool(r1)
}

func (m *TBasicAction) Execute() bool {
	r1 := basicActionImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TBasicAction) Update() bool {
	r1 := basicActionImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func BasicActionClass() TClass {
	ret := basicActionImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TBasicAction) UpdateTarget(Target IObject) {
	basicActionImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(Target))
}

func (m *TBasicAction) ExecuteTarget(Target IObject) {
	basicActionImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(Target))
}

func (m *TBasicAction) RegisterChanges(Value IBasicActionLink) {
	basicActionImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(Value))
}

func (m *TBasicAction) UnRegisterChanges(Value IBasicActionLink) {
	basicActionImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(Value))
}

func (m *TBasicAction) SetOnExecute(fn TNotifyEvent) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	basicActionImportAPI().SysCallN(7, m.Instance(), m.executePtr)
}

func (m *TBasicAction) SetOnUpdate(fn TNotifyEvent) {
	if m.updatePtr != 0 {
		RemoveEventElement(m.updatePtr)
	}
	m.updatePtr = MakeEventDataPtr(fn)
	basicActionImportAPI().SysCallN(8, m.Instance(), m.updatePtr)
}

var (
	basicActionImport       *imports.Imports = nil
	basicActionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("BasicAction_ActionComponent", 0),
		/*1*/ imports.NewTable("BasicAction_Class", 0),
		/*2*/ imports.NewTable("BasicAction_Create", 0),
		/*3*/ imports.NewTable("BasicAction_Execute", 0),
		/*4*/ imports.NewTable("BasicAction_ExecuteTarget", 0),
		/*5*/ imports.NewTable("BasicAction_HandlesTarget", 0),
		/*6*/ imports.NewTable("BasicAction_RegisterChanges", 0),
		/*7*/ imports.NewTable("BasicAction_SetOnExecute", 0),
		/*8*/ imports.NewTable("BasicAction_SetOnUpdate", 0),
		/*9*/ imports.NewTable("BasicAction_UnRegisterChanges", 0),
		/*10*/ imports.NewTable("BasicAction_Update", 0),
		/*11*/ imports.NewTable("BasicAction_UpdateTarget", 0),
	}
)

func basicActionImportAPI() *imports.Imports {
	if basicActionImport == nil {
		basicActionImport = NewDefaultImports()
		basicActionImport.SetImportTable(basicActionImportTables)
		basicActionImportTables = nil
	}
	return basicActionImport
}
