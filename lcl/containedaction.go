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

// IContainedAction Parent: IBasicAction
type IContainedAction interface {
	IBasicAction
	ActionList() ICustomActionList         // property ActionList Getter
	SetActionList(value ICustomActionList) // property ActionList Setter
	Index() int32                          // property Index Getter
	SetIndex(value int32)                  // property Index Setter
	Category() string                      // property Category Getter
	SetCategory(value string)              // property Category Setter
}

type TContainedAction struct {
	TBasicAction
}

func (m *TContainedAction) ActionList() ICustomActionList {
	if !m.IsValid() {
		return nil
	}
	r := containedActionAPI().SysCallN(1, 0, m.Instance())
	return AsCustomActionList(r)
}

func (m *TContainedAction) SetActionList(value ICustomActionList) {
	if !m.IsValid() {
		return
	}
	containedActionAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TContainedAction) Index() int32 {
	if !m.IsValid() {
		return 0
	}
	r := containedActionAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TContainedAction) SetIndex(value int32) {
	if !m.IsValid() {
		return
	}
	containedActionAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TContainedAction) Category() string {
	if !m.IsValid() {
		return ""
	}
	r := containedActionAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TContainedAction) SetCategory(value string) {
	if !m.IsValid() {
		return
	}
	containedActionAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

// NewContainedAction class constructor
func NewContainedAction(owner IComponent) IContainedAction {
	r := containedActionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsContainedAction(r)
}

func TContainedActionClass() types.TClass {
	r := containedActionAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	containedActionOnce   base.Once
	containedActionImport *imports.Imports = nil
)

func containedActionAPI() *imports.Imports {
	containedActionOnce.Do(func() {
		containedActionImport = api.NewDefaultImports()
		containedActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TContainedAction_Create", 0), // constructor NewContainedAction
			/* 1 */ imports.NewTable("TContainedAction_ActionList", 0), // property ActionList
			/* 2 */ imports.NewTable("TContainedAction_Index", 0), // property Index
			/* 3 */ imports.NewTable("TContainedAction_Category", 0), // property Category
			/* 4 */ imports.NewTable("TContainedAction_TClass", 0), // function TContainedActionClass
		}
	})
	return containedActionImport
}
