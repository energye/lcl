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

// IPageSetupDialog Parent: ICustomPrinterSetupDialog
type IPageSetupDialog interface {
	ICustomPrinterSetupDialog
	AttachTo() ICustomForm                     // property
	SetAttachTo(AValue ICustomForm)            // property
	PageWidth() int32                          // property
	SetPageWidth(AValue int32)                 // property
	PageHeight() int32                         // property
	SetPageHeight(AValue int32)                // property
	MarginLeft() int32                         // property
	SetMarginLeft(AValue int32)                // property
	MarginTop() int32                          // property
	SetMarginTop(AValue int32)                 // property
	MarginRight() int32                        // property
	SetMarginRight(AValue int32)               // property
	MarginBottom() int32                       // property
	SetMarginBottom(AValue int32)              // property
	MinMarginLeft() int32                      // property
	SetMinMarginLeft(AValue int32)             // property
	MinMarginTop() int32                       // property
	SetMinMarginTop(AValue int32)              // property
	MinMarginRight() int32                     // property
	SetMinMarginRight(AValue int32)            // property
	MinMarginBottom() int32                    // property
	SetMinMarginBottom(AValue int32)           // property
	Options() TPageSetupDialogOptions          // property
	SetOptions(AValue TPageSetupDialogOptions) // property
	Units() TPageMeasureUnits                  // property
	SetUnits(AValue TPageMeasureUnits)         // property
	SetOnDialogResult(fn TDialogResultEvent)   // property event
}

// TPageSetupDialog Parent: TCustomPrinterSetupDialog
type TPageSetupDialog struct {
	TCustomPrinterSetupDialog
	dialogResultPtr uintptr
}

func NewPageSetupDialog(TheOwner IComponent) IPageSetupDialog {
	r1 := pageSetupDialogImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsPageSetupDialog(r1)
}

func (m *TPageSetupDialog) AttachTo() ICustomForm {
	r1 := pageSetupDialogImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TPageSetupDialog) SetAttachTo(AValue ICustomForm) {
	pageSetupDialogImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPageSetupDialog) PageWidth() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetPageWidth(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) PageHeight() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetPageHeight(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginLeft() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginLeft(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginTop() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginTop(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginRight() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginRight(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginBottom() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginBottom(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginLeft() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginLeft(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginTop() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginTop(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginRight() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginRight(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginBottom() int32 {
	r1 := pageSetupDialogImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginBottom(AValue int32) {
	pageSetupDialogImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) Options() TPageSetupDialogOptions {
	r1 := pageSetupDialogImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TPageSetupDialogOptions(r1)
}

func (m *TPageSetupDialog) SetOptions(AValue TPageSetupDialogOptions) {
	pageSetupDialogImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) Units() TPageMeasureUnits {
	r1 := pageSetupDialogImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TPageMeasureUnits(r1)
}

func (m *TPageSetupDialog) SetUnits(AValue TPageMeasureUnits) {
	pageSetupDialogImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func PageSetupDialogClass() TClass {
	ret := pageSetupDialogImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TPageSetupDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if m.dialogResultPtr != 0 {
		RemoveEventElement(m.dialogResultPtr)
	}
	m.dialogResultPtr = MakeEventDataPtr(fn)
	pageSetupDialogImportAPI().SysCallN(14, m.Instance(), m.dialogResultPtr)
}

var (
	pageSetupDialogImport       *imports.Imports = nil
	pageSetupDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PageSetupDialog_AttachTo", 0),
		/*1*/ imports.NewTable("PageSetupDialog_Class", 0),
		/*2*/ imports.NewTable("PageSetupDialog_Create", 0),
		/*3*/ imports.NewTable("PageSetupDialog_MarginBottom", 0),
		/*4*/ imports.NewTable("PageSetupDialog_MarginLeft", 0),
		/*5*/ imports.NewTable("PageSetupDialog_MarginRight", 0),
		/*6*/ imports.NewTable("PageSetupDialog_MarginTop", 0),
		/*7*/ imports.NewTable("PageSetupDialog_MinMarginBottom", 0),
		/*8*/ imports.NewTable("PageSetupDialog_MinMarginLeft", 0),
		/*9*/ imports.NewTable("PageSetupDialog_MinMarginRight", 0),
		/*10*/ imports.NewTable("PageSetupDialog_MinMarginTop", 0),
		/*11*/ imports.NewTable("PageSetupDialog_Options", 0),
		/*12*/ imports.NewTable("PageSetupDialog_PageHeight", 0),
		/*13*/ imports.NewTable("PageSetupDialog_PageWidth", 0),
		/*14*/ imports.NewTable("PageSetupDialog_SetOnDialogResult", 0),
		/*15*/ imports.NewTable("PageSetupDialog_Units", 0),
	}
)

func pageSetupDialogImportAPI() *imports.Imports {
	if pageSetupDialogImport == nil {
		pageSetupDialogImport = NewDefaultImports()
		pageSetupDialogImport.SetImportTable(pageSetupDialogImportTables)
		pageSetupDialogImportTables = nil
	}
	return pageSetupDialogImport
}
