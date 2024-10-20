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

// ICustomActionList Parent: ILCLComponent
type ICustomActionList interface {
	ILCLComponent
	Actions(Index int32) IContainedAction                        // property
	SetActions(Index int32, AValue IContainedAction)             // property
	ActionCount() int32                                          // property
	Images() ICustomImageList                                    // property
	SetImages(AValue ICustomImageList)                           // property
	State() TActionListState                                     // property
	SetState(AValue TActionListState)                            // property
	ActionByName(ActionName string) IContainedAction             // function
	GetEnumeratorForActionListEnumerator() IActionListEnumerator // function
	IndexOfName(ActionName string) int32                         // function
	IsShortCut(Message *TLMKey) bool                             // function
}

// TCustomActionList Parent: TLCLComponent
type TCustomActionList struct {
	TLCLComponent
}

func NewCustomActionList(AOwner IComponent) ICustomActionList {
	r1 := customActionListImportAPI().SysCallN(4, GetObjectUintptr(AOwner))
	return AsCustomActionList(r1)
}

func (m *TCustomActionList) Actions(Index int32) IContainedAction {
	r1 := customActionListImportAPI().SysCallN(2, 0, m.Instance(), uintptr(Index))
	return AsContainedAction(r1)
}

func (m *TCustomActionList) SetActions(Index int32, AValue IContainedAction) {
	customActionListImportAPI().SysCallN(2, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TCustomActionList) ActionCount() int32 {
	r1 := customActionListImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TCustomActionList) Images() ICustomImageList {
	r1 := customActionListImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomActionList) SetImages(AValue ICustomImageList) {
	customActionListImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomActionList) State() TActionListState {
	r1 := customActionListImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TActionListState(r1)
}

func (m *TCustomActionList) SetState(AValue TActionListState) {
	customActionListImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomActionList) ActionByName(ActionName string) IContainedAction {
	r1 := customActionListImportAPI().SysCallN(0, m.Instance(), PascalStr(ActionName))
	return AsContainedAction(r1)
}

func (m *TCustomActionList) GetEnumeratorForActionListEnumerator() IActionListEnumerator {
	r1 := customActionListImportAPI().SysCallN(5, m.Instance())
	return AsActionListEnumerator(r1)
}

func (m *TCustomActionList) IndexOfName(ActionName string) int32 {
	r1 := customActionListImportAPI().SysCallN(7, m.Instance(), PascalStr(ActionName))
	return int32(r1)
}

func (m *TCustomActionList) IsShortCut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := customActionListImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func CustomActionListClass() TClass {
	ret := customActionListImportAPI().SysCallN(3)
	return TClass(ret)
}

var (
	customActionListImport       *imports.Imports = nil
	customActionListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomActionList_ActionByName", 0),
		/*1*/ imports.NewTable("CustomActionList_ActionCount", 0),
		/*2*/ imports.NewTable("CustomActionList_Actions", 0),
		/*3*/ imports.NewTable("CustomActionList_Class", 0),
		/*4*/ imports.NewTable("CustomActionList_Create", 0),
		/*5*/ imports.NewTable("CustomActionList_GetEnumeratorForActionListEnumerator", 0),
		/*6*/ imports.NewTable("CustomActionList_Images", 0),
		/*7*/ imports.NewTable("CustomActionList_IndexOfName", 0),
		/*8*/ imports.NewTable("CustomActionList_IsShortCut", 0),
		/*9*/ imports.NewTable("CustomActionList_State", 0),
	}
)

func customActionListImportAPI() *imports.Imports {
	if customActionListImport == nil {
		customActionListImport = NewDefaultImports()
		customActionListImport.SetImportTable(customActionListImportTables)
		customActionListImportTables = nil
	}
	return customActionListImport
}
