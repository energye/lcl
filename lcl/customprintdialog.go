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

// ICustomPrintDialog Parent: ICommonDialog
type ICustomPrintDialog interface {
	ICommonDialog
	Collate() bool                         // property
	SetCollate(AValue bool)                // property
	Copies() int32                         // property
	SetCopies(AValue int32)                // property
	FromPage() int32                       // property
	SetFromPage(AValue int32)              // property
	MinPage() int32                        // property
	SetMinPage(AValue int32)               // property
	MaxPage() int32                        // property
	SetMaxPage(AValue int32)               // property
	Options() TPrintDialogOptions          // property
	SetOptions(AValue TPrintDialogOptions) // property
	PrintToFile() bool                     // property
	SetPrintToFile(AValue bool)            // property
	PrintRange() TPrintRange               // property
	SetPrintRange(AValue TPrintRange)      // property
	ToPage() int32                         // property
	SetToPage(AValue int32)                // property
}

// TCustomPrintDialog Parent: TCommonDialog
type TCustomPrintDialog struct {
	TCommonDialog
}

func NewCustomPrintDialog(TheOwner IComponent) ICustomPrintDialog {
	r1 := customPrintDialogImportAPI().SysCallN(3, GetObjectUintptr(TheOwner))
	return AsCustomPrintDialog(r1)
}

func (m *TCustomPrintDialog) Collate() bool {
	r1 := customPrintDialogImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPrintDialog) SetCollate(AValue bool) {
	customPrintDialogImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPrintDialog) Copies() int32 {
	r1 := customPrintDialogImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetCopies(AValue int32) {
	customPrintDialogImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) FromPage() int32 {
	r1 := customPrintDialogImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetFromPage(AValue int32) {
	customPrintDialogImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) MinPage() int32 {
	r1 := customPrintDialogImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetMinPage(AValue int32) {
	customPrintDialogImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) MaxPage() int32 {
	r1 := customPrintDialogImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetMaxPage(AValue int32) {
	customPrintDialogImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) Options() TPrintDialogOptions {
	r1 := customPrintDialogImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TPrintDialogOptions(r1)
}

func (m *TCustomPrintDialog) SetOptions(AValue TPrintDialogOptions) {
	customPrintDialogImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) PrintToFile() bool {
	r1 := customPrintDialogImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPrintDialog) SetPrintToFile(AValue bool) {
	customPrintDialogImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPrintDialog) PrintRange() TPrintRange {
	r1 := customPrintDialogImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TPrintRange(r1)
}

func (m *TCustomPrintDialog) SetPrintRange(AValue TPrintRange) {
	customPrintDialogImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) ToPage() int32 {
	r1 := customPrintDialogImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetToPage(AValue int32) {
	customPrintDialogImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func CustomPrintDialogClass() TClass {
	ret := customPrintDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customPrintDialogImport       *imports.Imports = nil
	customPrintDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPrintDialog_Class", 0),
		/*1*/ imports.NewTable("CustomPrintDialog_Collate", 0),
		/*2*/ imports.NewTable("CustomPrintDialog_Copies", 0),
		/*3*/ imports.NewTable("CustomPrintDialog_Create", 0),
		/*4*/ imports.NewTable("CustomPrintDialog_FromPage", 0),
		/*5*/ imports.NewTable("CustomPrintDialog_MaxPage", 0),
		/*6*/ imports.NewTable("CustomPrintDialog_MinPage", 0),
		/*7*/ imports.NewTable("CustomPrintDialog_Options", 0),
		/*8*/ imports.NewTable("CustomPrintDialog_PrintRange", 0),
		/*9*/ imports.NewTable("CustomPrintDialog_PrintToFile", 0),
		/*10*/ imports.NewTable("CustomPrintDialog_ToPage", 0),
	}
)

func customPrintDialogImportAPI() *imports.Imports {
	if customPrintDialogImport == nil {
		customPrintDialogImport = NewDefaultImports()
		customPrintDialogImport.SetImportTable(customPrintDialogImportTables)
		customPrintDialogImportTables = nil
	}
	return customPrintDialogImport
}
