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

// IBasicAction Parent: IComponent
type IBasicAction interface {
	IComponent
	HandlesTarget(target IObject) bool        // function
	Execute() bool                            // function
	Update() bool                             // function
	UpdateTarget(target IObject)              // procedure
	ExecuteTarget(target IObject)             // procedure
	RegisterChanges(value IBasicActionLink)   // procedure
	UnRegisterChanges(value IBasicActionLink) // procedure
	ActionComponent() IComponent              // property ActionComponent Getter
	SetActionComponent(value IComponent)      // property ActionComponent Setter
	SetOnExecute(fn TNotifyEvent)             // property event
	SetOnUpdate(fn TNotifyEvent)              // property event
}

type TBasicAction struct {
	TComponent
}

func (m *TBasicAction) HandlesTarget(target IObject) bool {
	if !m.IsValid() {
		return false
	}
	r := basicActionAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(target))
	return api.GoBool(r)
}

func (m *TBasicAction) Execute() bool {
	if !m.IsValid() {
		return false
	}
	r := basicActionAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TBasicAction) Update() bool {
	if !m.IsValid() {
		return false
	}
	r := basicActionAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TBasicAction) UpdateTarget(target IObject) {
	if !m.IsValid() {
		return
	}
	basicActionAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(target))
}

func (m *TBasicAction) ExecuteTarget(target IObject) {
	if !m.IsValid() {
		return
	}
	basicActionAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(target))
}

func (m *TBasicAction) RegisterChanges(value IBasicActionLink) {
	if !m.IsValid() {
		return
	}
	basicActionAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TBasicAction) UnRegisterChanges(value IBasicActionLink) {
	if !m.IsValid() {
		return
	}
	basicActionAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TBasicAction) ActionComponent() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := basicActionAPI().SysCallN(8, 0, m.Instance())
	return AsComponent(r)
}

func (m *TBasicAction) SetActionComponent(value IComponent) {
	if !m.IsValid() {
		return
	}
	basicActionAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TBasicAction) SetOnExecute(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, basicActionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBasicAction) SetOnUpdate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, basicActionAPI(), api.MakeEventDataPtr(cb))
}

// NewBasicAction class constructor
func NewBasicAction(owner IComponent) IBasicAction {
	r := basicActionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsBasicAction(r)
}

func TBasicActionClass() types.TClass {
	r := basicActionAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	basicActionOnce   base.Once
	basicActionImport *imports.Imports = nil
)

func basicActionAPI() *imports.Imports {
	basicActionOnce.Do(func() {
		basicActionImport = api.NewDefaultImports()
		basicActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBasicAction_Create", 0), // constructor NewBasicAction
			/* 1 */ imports.NewTable("TBasicAction_HandlesTarget", 0), // function HandlesTarget
			/* 2 */ imports.NewTable("TBasicAction_Execute", 0), // function Execute
			/* 3 */ imports.NewTable("TBasicAction_Update", 0), // function Update
			/* 4 */ imports.NewTable("TBasicAction_UpdateTarget", 0), // procedure UpdateTarget
			/* 5 */ imports.NewTable("TBasicAction_ExecuteTarget", 0), // procedure ExecuteTarget
			/* 6 */ imports.NewTable("TBasicAction_RegisterChanges", 0), // procedure RegisterChanges
			/* 7 */ imports.NewTable("TBasicAction_UnRegisterChanges", 0), // procedure UnRegisterChanges
			/* 8 */ imports.NewTable("TBasicAction_ActionComponent", 0), // property ActionComponent
			/* 9 */ imports.NewTable("TBasicAction_OnExecute", 0), // event OnExecute
			/* 10 */ imports.NewTable("TBasicAction_OnUpdate", 0), // event OnUpdate
			/* 11 */ imports.NewTable("TBasicAction_TClass", 0), // function TBasicActionClass
		}
	})
	return basicActionImport
}
