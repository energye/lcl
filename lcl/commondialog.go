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

// ICommonDialog Parent: ILCLComponent
type ICommonDialog interface {
	ILCLComponent
	Execute() bool                           // function
	HandleAllocated() bool                   // function
	Close()                                  // procedure
	DoShow()                                 // procedure
	DoCanClose(canClose *bool)               // procedure
	DoClose()                                // procedure
	Handle() types.TLCLHandle                // property Handle Getter
	SetHandle(value types.TLCLHandle)        // property Handle Setter
	UserChoice() int32                       // property UserChoice Getter
	SetUserChoice(value int32)               // property UserChoice Setter
	Width() int32                            // property Width Getter
	SetWidth(value int32)                    // property Width Setter
	Height() int32                           // property Height Getter
	SetHeight(value int32)                   // property Height Setter
	HelpContext() types.THelpContext         // property HelpContext Getter
	SetHelpContext(value types.THelpContext) // property HelpContext Setter
	Title() string                           // property Title Getter
	SetTitle(value string)                   // property Title Setter
	SetOnClose(fn TNotifyEvent)              // property event
	SetOnCanClose(fn TCloseQueryEvent)       // property event
	SetOnShow(fn TNotifyEvent)               // property event
}

type TCommonDialog struct {
	TLCLComponent
}

func (m *TCommonDialog) Execute() bool {
	if !m.IsValid() {
		return false
	}
	r := commonDialogAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCommonDialog) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := commonDialogAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCommonDialog) Close() {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(3, m.Instance())
}

func (m *TCommonDialog) DoShow() {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(4, m.Instance())
}

func (m *TCommonDialog) DoCanClose(canClose *bool) {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(canClose)))
}

func (m *TCommonDialog) DoClose() {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(6, m.Instance())
}

func (m *TCommonDialog) Handle() types.TLCLHandle {
	if !m.IsValid() {
		return 0
	}
	r := commonDialogAPI().SysCallN(7, 0, m.Instance())
	return types.TLCLHandle(r)
}

func (m *TCommonDialog) SetHandle(value types.TLCLHandle) {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCommonDialog) UserChoice() int32 {
	if !m.IsValid() {
		return 0
	}
	r := commonDialogAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCommonDialog) SetUserChoice(value int32) {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCommonDialog) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := commonDialogAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TCommonDialog) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCommonDialog) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := commonDialogAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCommonDialog) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCommonDialog) HelpContext() types.THelpContext {
	if !m.IsValid() {
		return 0
	}
	r := commonDialogAPI().SysCallN(11, 0, m.Instance())
	return types.THelpContext(r)
}

func (m *TCommonDialog) SetHelpContext(value types.THelpContext) {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCommonDialog) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := commonDialogAPI().SysCallN(12, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCommonDialog) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	commonDialogAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TCommonDialog) SetOnClose(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, commonDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCommonDialog) SetOnCanClose(fn TCloseQueryEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCloseQueryEvent(fn)
	base.SetEvent(m, 14, commonDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCommonDialog) SetOnShow(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, commonDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewCommonDialog class constructor
func NewCommonDialog(theOwner IComponent) ICommonDialog {
	r := commonDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCommonDialog(r)
}

func TCommonDialogClass() types.TClass {
	r := commonDialogAPI().SysCallN(16)
	return types.TClass(r)
}

var (
	commonDialogOnce   base.Once
	commonDialogImport *imports.Imports = nil
)

func commonDialogAPI() *imports.Imports {
	commonDialogOnce.Do(func() {
		commonDialogImport = api.NewDefaultImports()
		commonDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCommonDialog_Create", 0), // constructor NewCommonDialog
			/* 1 */ imports.NewTable("TCommonDialog_Execute", 0), // function Execute
			/* 2 */ imports.NewTable("TCommonDialog_HandleAllocated", 0), // function HandleAllocated
			/* 3 */ imports.NewTable("TCommonDialog_Close", 0), // procedure Close
			/* 4 */ imports.NewTable("TCommonDialog_DoShow", 0), // procedure DoShow
			/* 5 */ imports.NewTable("TCommonDialog_DoCanClose", 0), // procedure DoCanClose
			/* 6 */ imports.NewTable("TCommonDialog_DoClose", 0), // procedure DoClose
			/* 7 */ imports.NewTable("TCommonDialog_Handle", 0), // property Handle
			/* 8 */ imports.NewTable("TCommonDialog_UserChoice", 0), // property UserChoice
			/* 9 */ imports.NewTable("TCommonDialog_Width", 0), // property Width
			/* 10 */ imports.NewTable("TCommonDialog_Height", 0), // property Height
			/* 11 */ imports.NewTable("TCommonDialog_HelpContext", 0), // property HelpContext
			/* 12 */ imports.NewTable("TCommonDialog_Title", 0), // property Title
			/* 13 */ imports.NewTable("TCommonDialog_OnClose", 0), // event OnClose
			/* 14 */ imports.NewTable("TCommonDialog_OnCanClose", 0), // event OnCanClose
			/* 15 */ imports.NewTable("TCommonDialog_OnShow", 0), // event OnShow
			/* 16 */ imports.NewTable("TCommonDialog_TClass", 0), // function TCommonDialogClass
		}
	})
	return commonDialogImport
}
