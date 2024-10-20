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

// ICommonDialog Parent: ILCLComponent
type ICommonDialog interface {
	ILCLComponent
	Handle() THandle                    // property
	SetHandle(AValue THandle)           // property
	UserChoice() int32                  // property
	SetUserChoice(AValue int32)         // property
	Width() int32                       // property
	SetWidth(AValue int32)              // property
	Height() int32                      // property
	SetHeight(AValue int32)             // property
	HelpContext() THelpContext          // property
	SetHelpContext(AValue THelpContext) // property
	Title() string                      // property
	SetTitle(AValue string)             // property
	Execute() bool                      // function
	HandleAllocated() bool              // function
	Close()                             // procedure
	SetOnClose(fn TNotifyEvent)         // property event
	SetOnCanClose(fn TCloseQueryEvent)  // property event
	SetOnShow(fn TNotifyEvent)          // property event
}

// TCommonDialog Parent: TLCLComponent
type TCommonDialog struct {
	TLCLComponent
	closePtr    uintptr
	canClosePtr uintptr
	showPtr     uintptr
}

func NewCommonDialog(TheOwner IComponent) ICommonDialog {
	r1 := commonDialogImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsCommonDialog(r1)
}

func (m *TCommonDialog) Handle() THandle {
	r1 := commonDialogImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return THandle(r1)
}

func (m *TCommonDialog) SetHandle(AValue THandle) {
	commonDialogImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) UserChoice() int32 {
	r1 := commonDialogImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCommonDialog) SetUserChoice(AValue int32) {
	commonDialogImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) Width() int32 {
	r1 := commonDialogImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCommonDialog) SetWidth(AValue int32) {
	commonDialogImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) Height() int32 {
	r1 := commonDialogImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCommonDialog) SetHeight(AValue int32) {
	commonDialogImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) HelpContext() THelpContext {
	r1 := commonDialogImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TCommonDialog) SetHelpContext(AValue THelpContext) {
	commonDialogImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) Title() string {
	r1 := commonDialogImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCommonDialog) SetTitle(AValue string) {
	commonDialogImportAPI().SysCallN(11, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCommonDialog) Execute() bool {
	r1 := commonDialogImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TCommonDialog) HandleAllocated() bool {
	r1 := commonDialogImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func CommonDialogClass() TClass {
	ret := commonDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCommonDialog) Close() {
	commonDialogImportAPI().SysCallN(1, m.Instance())
}

func (m *TCommonDialog) SetOnClose(fn TNotifyEvent) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	commonDialogImportAPI().SysCallN(9, m.Instance(), m.closePtr)
}

func (m *TCommonDialog) SetOnCanClose(fn TCloseQueryEvent) {
	if m.canClosePtr != 0 {
		RemoveEventElement(m.canClosePtr)
	}
	m.canClosePtr = MakeEventDataPtr(fn)
	commonDialogImportAPI().SysCallN(8, m.Instance(), m.canClosePtr)
}

func (m *TCommonDialog) SetOnShow(fn TNotifyEvent) {
	if m.showPtr != 0 {
		RemoveEventElement(m.showPtr)
	}
	m.showPtr = MakeEventDataPtr(fn)
	commonDialogImportAPI().SysCallN(10, m.Instance(), m.showPtr)
}

var (
	commonDialogImport       *imports.Imports = nil
	commonDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CommonDialog_Class", 0),
		/*1*/ imports.NewTable("CommonDialog_Close", 0),
		/*2*/ imports.NewTable("CommonDialog_Create", 0),
		/*3*/ imports.NewTable("CommonDialog_Execute", 0),
		/*4*/ imports.NewTable("CommonDialog_Handle", 0),
		/*5*/ imports.NewTable("CommonDialog_HandleAllocated", 0),
		/*6*/ imports.NewTable("CommonDialog_Height", 0),
		/*7*/ imports.NewTable("CommonDialog_HelpContext", 0),
		/*8*/ imports.NewTable("CommonDialog_SetOnCanClose", 0),
		/*9*/ imports.NewTable("CommonDialog_SetOnClose", 0),
		/*10*/ imports.NewTable("CommonDialog_SetOnShow", 0),
		/*11*/ imports.NewTable("CommonDialog_Title", 0),
		/*12*/ imports.NewTable("CommonDialog_UserChoice", 0),
		/*13*/ imports.NewTable("CommonDialog_Width", 0),
	}
)

func commonDialogImportAPI() *imports.Imports {
	if commonDialogImport == nil {
		commonDialogImport = NewDefaultImports()
		commonDialogImport.SetImportTable(commonDialogImportTables)
		commonDialogImportTables = nil
	}
	return commonDialogImport
}
