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

// ICustomActionList Parent: ILCLComponent
type ICustomActionList interface {
	ILCLComponent
	ActionByName(actionName string) IContainedAction            // function
	GetEnumeratorToActionListEnumerator() IActionListEnumerator // function
	IndexOfName(actionName string) int32                        // function
	IsShortCut(message *types.TLMKey) bool                      // function
	Actions(index int32) IContainedAction                       // property Actions Getter
	SetActions(index int32, value IContainedAction)             // property Actions Setter
	ActionCount() int32                                         // property ActionCount Getter
	Images() ICustomImageList                                   // property Images Getter
	SetImages(value ICustomImageList)                           // property Images Setter
	State() types.TActionListState                              // property State Getter
	SetState(value types.TActionListState)                      // property State Setter
}

type TCustomActionList struct {
	TLCLComponent
}

func (m *TCustomActionList) ActionByName(actionName string) IContainedAction {
	if !m.IsValid() {
		return nil
	}
	r := customActionListAPI().SysCallN(1, m.Instance(), api.PasStr(actionName))
	return AsContainedAction(r)
}

func (m *TCustomActionList) GetEnumeratorToActionListEnumerator() IActionListEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := customActionListAPI().SysCallN(2, m.Instance())
	return AsActionListEnumerator(r)
}

func (m *TCustomActionList) IndexOfName(actionName string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customActionListAPI().SysCallN(3, m.Instance(), api.PasStr(actionName))
	return int32(r)
}

func (m *TCustomActionList) IsShortCut(message *types.TLMKey) bool {
	if !m.IsValid() {
		return false
	}
	r := customActionListAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(message)))
	return api.GoBool(r)
}

func (m *TCustomActionList) Actions(index int32) IContainedAction {
	if !m.IsValid() {
		return nil
	}
	r := customActionListAPI().SysCallN(5, 0, m.Instance(), uintptr(index))
	return AsContainedAction(r)
}

func (m *TCustomActionList) SetActions(index int32, value IContainedAction) {
	if !m.IsValid() {
		return
	}
	customActionListAPI().SysCallN(5, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TCustomActionList) ActionCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customActionListAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCustomActionList) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customActionListAPI().SysCallN(7, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomActionList) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customActionListAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomActionList) State() types.TActionListState {
	if !m.IsValid() {
		return 0
	}
	r := customActionListAPI().SysCallN(8, 0, m.Instance())
	return types.TActionListState(r)
}

func (m *TCustomActionList) SetState(value types.TActionListState) {
	if !m.IsValid() {
		return
	}
	customActionListAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

// NewCustomActionList class constructor
func NewCustomActionList(owner IComponent) ICustomActionList {
	r := customActionListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomActionList(r)
}

func TCustomActionListClass() types.TClass {
	r := customActionListAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	customActionListOnce   base.Once
	customActionListImport *imports.Imports = nil
)

func customActionListAPI() *imports.Imports {
	customActionListOnce.Do(func() {
		customActionListImport = api.NewDefaultImports()
		customActionListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomActionList_Create", 0), // constructor NewCustomActionList
			/* 1 */ imports.NewTable("TCustomActionList_ActionByName", 0), // function ActionByName
			/* 2 */ imports.NewTable("TCustomActionList_GetEnumeratorToActionListEnumerator", 0), // function GetEnumeratorToActionListEnumerator
			/* 3 */ imports.NewTable("TCustomActionList_IndexOfName", 0), // function IndexOfName
			/* 4 */ imports.NewTable("TCustomActionList_IsShortCut", 0), // function IsShortCut
			/* 5 */ imports.NewTable("TCustomActionList_Actions", 0), // property Actions
			/* 6 */ imports.NewTable("TCustomActionList_ActionCount", 0), // property ActionCount
			/* 7 */ imports.NewTable("TCustomActionList_Images", 0), // property Images
			/* 8 */ imports.NewTable("TCustomActionList_State", 0), // property State
			/* 9 */ imports.NewTable("TCustomActionList_TClass", 0), // function TCustomActionListClass
		}
	})
	return customActionListImport
}
