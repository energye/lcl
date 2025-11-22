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

// IPageSetupDialog Parent: ICustomPrinterSetupDialog
type IPageSetupDialog interface {
	ICustomPrinterSetupDialog
	AttachTo() ICustomForm                          // property AttachTo Getter
	SetAttachTo(value ICustomForm)                  // property AttachTo Setter
	PageWidth() int32                               // property PageWidth Getter
	SetPageWidth(value int32)                       // property PageWidth Setter
	PageHeight() int32                              // property PageHeight Getter
	SetPageHeight(value int32)                      // property PageHeight Setter
	MarginLeft() int32                              // property MarginLeft Getter
	SetMarginLeft(value int32)                      // property MarginLeft Setter
	MarginTop() int32                               // property MarginTop Getter
	SetMarginTop(value int32)                       // property MarginTop Setter
	MarginRight() int32                             // property MarginRight Getter
	SetMarginRight(value int32)                     // property MarginRight Setter
	MarginBottom() int32                            // property MarginBottom Getter
	SetMarginBottom(value int32)                    // property MarginBottom Setter
	MinMarginLeft() int32                           // property MinMarginLeft Getter
	SetMinMarginLeft(value int32)                   // property MinMarginLeft Setter
	MinMarginTop() int32                            // property MinMarginTop Getter
	SetMinMarginTop(value int32)                    // property MinMarginTop Setter
	MinMarginRight() int32                          // property MinMarginRight Getter
	SetMinMarginRight(value int32)                  // property MinMarginRight Setter
	MinMarginBottom() int32                         // property MinMarginBottom Getter
	SetMinMarginBottom(value int32)                 // property MinMarginBottom Setter
	Options() types.TPageSetupDialogOptions         // property Options Getter
	SetOptions(value types.TPageSetupDialogOptions) // property Options Setter
	Units() types.TPageMeasureUnits                 // property Units Getter
	SetUnits(value types.TPageMeasureUnits)         // property Units Setter
	SetOnDialogResult(fn TDialogResultEvent)        // property event
}

type TPageSetupDialog struct {
	TCustomPrinterSetupDialog
}

func (m *TPageSetupDialog) AttachTo() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := pageSetupDialogAPI().SysCallN(1, 0, m.Instance())
	return AsCustomForm(r)
}

func (m *TPageSetupDialog) SetAttachTo(value ICustomForm) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPageSetupDialog) PageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetPageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) PageHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetPageHeight(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MarginLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMarginLeft(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MarginTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMarginTop(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MarginRight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMarginRight(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MarginBottom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMarginBottom(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MinMarginLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMinMarginLeft(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MinMarginTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMinMarginTop(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MinMarginRight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMinMarginRight(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) MinMarginBottom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TPageSetupDialog) SetMinMarginBottom(value int32) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) Options() types.TPageSetupDialogOptions {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(12, 0, m.Instance())
	return types.TPageSetupDialogOptions(r)
}

func (m *TPageSetupDialog) SetOptions(value types.TPageSetupDialogOptions) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) Units() types.TPageMeasureUnits {
	if !m.IsValid() {
		return 0
	}
	r := pageSetupDialogAPI().SysCallN(13, 0, m.Instance())
	return types.TPageMeasureUnits(r)
}

func (m *TPageSetupDialog) SetUnits(value types.TPageMeasureUnits) {
	if !m.IsValid() {
		return
	}
	pageSetupDialogAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TPageSetupDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDialogResultEvent(fn)
	base.SetEvent(m, 14, pageSetupDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewPageSetupDialog class constructor
func NewPageSetupDialog(theOwner IComponent) IPageSetupDialog {
	r := pageSetupDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPageSetupDialog(r)
}

func TPageSetupDialogClass() types.TClass {
	r := pageSetupDialogAPI().SysCallN(15)
	return types.TClass(r)
}

var (
	pageSetupDialogOnce   base.Once
	pageSetupDialogImport *imports.Imports = nil
)

func pageSetupDialogAPI() *imports.Imports {
	pageSetupDialogOnce.Do(func() {
		pageSetupDialogImport = api.NewDefaultImports()
		pageSetupDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPageSetupDialog_Create", 0), // constructor NewPageSetupDialog
			/* 1 */ imports.NewTable("TPageSetupDialog_AttachTo", 0), // property AttachTo
			/* 2 */ imports.NewTable("TPageSetupDialog_PageWidth", 0), // property PageWidth
			/* 3 */ imports.NewTable("TPageSetupDialog_PageHeight", 0), // property PageHeight
			/* 4 */ imports.NewTable("TPageSetupDialog_MarginLeft", 0), // property MarginLeft
			/* 5 */ imports.NewTable("TPageSetupDialog_MarginTop", 0), // property MarginTop
			/* 6 */ imports.NewTable("TPageSetupDialog_MarginRight", 0), // property MarginRight
			/* 7 */ imports.NewTable("TPageSetupDialog_MarginBottom", 0), // property MarginBottom
			/* 8 */ imports.NewTable("TPageSetupDialog_MinMarginLeft", 0), // property MinMarginLeft
			/* 9 */ imports.NewTable("TPageSetupDialog_MinMarginTop", 0), // property MinMarginTop
			/* 10 */ imports.NewTable("TPageSetupDialog_MinMarginRight", 0), // property MinMarginRight
			/* 11 */ imports.NewTable("TPageSetupDialog_MinMarginBottom", 0), // property MinMarginBottom
			/* 12 */ imports.NewTable("TPageSetupDialog_Options", 0), // property Options
			/* 13 */ imports.NewTable("TPageSetupDialog_Units", 0), // property Units
			/* 14 */ imports.NewTable("TPageSetupDialog_OnDialogResult", 0), // event OnDialogResult
			/* 15 */ imports.NewTable("TPageSetupDialog_TClass", 0), // function TPageSetupDialogClass
		}
	})
	return pageSetupDialogImport
}
