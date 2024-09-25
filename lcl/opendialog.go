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
	. "github.com/energye/lcl/types"
)

// IOpenDialog Parent: IFileDialog
type IOpenDialog interface {
	IFileDialog
	Options() TOpenOptions                          // property
	SetOptions(AValue TOpenOptions)                 // property
	DoFolderChange()                                // procedure
	DoSelectionChange()                             // procedure
	IntfSetOption(AOption TOpenOption, AValue bool) // procedure
	SetOnFolderChange(fn TNotifyEvent)              // property event
	SetOnSelectionChange(fn TNotifyEvent)           // property event
}

// TOpenDialog Parent: TFileDialog
type TOpenDialog struct {
	TFileDialog
	folderChangePtr    uintptr
	selectionChangePtr uintptr
}

func NewOpenDialog(TheOwner IComponent) IOpenDialog {
	r1 := LCL().SysCallN(4399, GetObjectUintptr(TheOwner))
	return AsOpenDialog(r1)
}

func (m *TOpenDialog) Options() TOpenOptions {
	r1 := LCL().SysCallN(4403, 0, m.Instance(), 0)
	return TOpenOptions(r1)
}

func (m *TOpenDialog) SetOptions(AValue TOpenOptions) {
	LCL().SysCallN(4403, 1, m.Instance(), uintptr(AValue))
}

func OpenDialogClass() TClass {
	ret := LCL().SysCallN(4398)
	return TClass(ret)
}

func (m *TOpenDialog) DoFolderChange() {
	LCL().SysCallN(4400, m.Instance())
}

func (m *TOpenDialog) DoSelectionChange() {
	LCL().SysCallN(4401, m.Instance())
}

func (m *TOpenDialog) IntfSetOption(AOption TOpenOption, AValue bool) {
	LCL().SysCallN(4402, m.Instance(), uintptr(AOption), PascalBool(AValue))
}

func (m *TOpenDialog) SetOnFolderChange(fn TNotifyEvent) {
	if m.folderChangePtr != 0 {
		RemoveEventElement(m.folderChangePtr)
	}
	m.folderChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4404, m.Instance(), m.folderChangePtr)
}

func (m *TOpenDialog) SetOnSelectionChange(fn TNotifyEvent) {
	if m.selectionChangePtr != 0 {
		RemoveEventElement(m.selectionChangePtr)
	}
	m.selectionChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4405, m.Instance(), m.selectionChangePtr)
}
