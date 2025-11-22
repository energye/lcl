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

// ICustomPrintDialog Parent: ICommonDialog
type ICustomPrintDialog interface {
	ICommonDialog
	Collate() bool                              // property Collate Getter
	SetCollate(value bool)                      // property Collate Setter
	Copies() int32                              // property Copies Getter
	SetCopies(value int32)                      // property Copies Setter
	FromPage() int32                            // property FromPage Getter
	SetFromPage(value int32)                    // property FromPage Setter
	MinPage() int32                             // property MinPage Getter
	SetMinPage(value int32)                     // property MinPage Setter
	MaxPage() int32                             // property MaxPage Getter
	SetMaxPage(value int32)                     // property MaxPage Setter
	Options() types.TPrintDialogOptions         // property Options Getter
	SetOptions(value types.TPrintDialogOptions) // property Options Setter
	PrintToFile() bool                          // property PrintToFile Getter
	SetPrintToFile(value bool)                  // property PrintToFile Setter
	PrintRange() types.TPrintRange              // property PrintRange Getter
	SetPrintRange(value types.TPrintRange)      // property PrintRange Setter
	ToPage() int32                              // property ToPage Getter
	SetToPage(value int32)                      // property ToPage Setter
}

type TCustomPrintDialog struct {
	TCommonDialog
}

func (m *TCustomPrintDialog) Collate() bool {
	if !m.IsValid() {
		return false
	}
	r := customPrintDialogAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomPrintDialog) SetCollate(value bool) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomPrintDialog) Copies() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPrintDialogAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPrintDialog) SetCopies(value int32) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPrintDialog) FromPage() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPrintDialogAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPrintDialog) SetFromPage(value int32) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPrintDialog) MinPage() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPrintDialogAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPrintDialog) SetMinPage(value int32) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPrintDialog) MaxPage() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPrintDialogAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPrintDialog) SetMaxPage(value int32) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPrintDialog) Options() types.TPrintDialogOptions {
	if !m.IsValid() {
		return 0
	}
	r := customPrintDialogAPI().SysCallN(6, 0, m.Instance())
	return types.TPrintDialogOptions(r)
}

func (m *TCustomPrintDialog) SetOptions(value types.TPrintDialogOptions) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPrintDialog) PrintToFile() bool {
	if !m.IsValid() {
		return false
	}
	r := customPrintDialogAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomPrintDialog) SetPrintToFile(value bool) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomPrintDialog) PrintRange() types.TPrintRange {
	if !m.IsValid() {
		return 0
	}
	r := customPrintDialogAPI().SysCallN(8, 0, m.Instance())
	return types.TPrintRange(r)
}

func (m *TCustomPrintDialog) SetPrintRange(value types.TPrintRange) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPrintDialog) ToPage() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPrintDialogAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPrintDialog) SetToPage(value int32) {
	if !m.IsValid() {
		return
	}
	customPrintDialogAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

// NewCustomPrintDialog class constructor
func NewCustomPrintDialog(theOwner IComponent) ICustomPrintDialog {
	r := customPrintDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomPrintDialog(r)
}

func TCustomPrintDialogClass() types.TClass {
	r := customPrintDialogAPI().SysCallN(10)
	return types.TClass(r)
}

var (
	customPrintDialogOnce   base.Once
	customPrintDialogImport *imports.Imports = nil
)

func customPrintDialogAPI() *imports.Imports {
	customPrintDialogOnce.Do(func() {
		customPrintDialogImport = api.NewDefaultImports()
		customPrintDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPrintDialog_Create", 0), // constructor NewCustomPrintDialog
			/* 1 */ imports.NewTable("TCustomPrintDialog_Collate", 0), // property Collate
			/* 2 */ imports.NewTable("TCustomPrintDialog_Copies", 0), // property Copies
			/* 3 */ imports.NewTable("TCustomPrintDialog_FromPage", 0), // property FromPage
			/* 4 */ imports.NewTable("TCustomPrintDialog_MinPage", 0), // property MinPage
			/* 5 */ imports.NewTable("TCustomPrintDialog_MaxPage", 0), // property MaxPage
			/* 6 */ imports.NewTable("TCustomPrintDialog_Options", 0), // property Options
			/* 7 */ imports.NewTable("TCustomPrintDialog_PrintToFile", 0), // property PrintToFile
			/* 8 */ imports.NewTable("TCustomPrintDialog_PrintRange", 0), // property PrintRange
			/* 9 */ imports.NewTable("TCustomPrintDialog_ToPage", 0), // property ToPage
			/* 10 */ imports.NewTable("TCustomPrintDialog_TClass", 0), // function TCustomPrintDialogClass
		}
	})
	return customPrintDialogImport
}
