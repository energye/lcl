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
	r1 := LCL().SysCallN(4466, GetObjectUintptr(TheOwner))
	return AsPageSetupDialog(r1)
}

func (m *TPageSetupDialog) AttachTo() ICustomForm {
	r1 := LCL().SysCallN(4464, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TPageSetupDialog) SetAttachTo(AValue ICustomForm) {
	LCL().SysCallN(4464, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPageSetupDialog) PageWidth() int32 {
	r1 := LCL().SysCallN(4477, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetPageWidth(AValue int32) {
	LCL().SysCallN(4477, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) PageHeight() int32 {
	r1 := LCL().SysCallN(4476, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetPageHeight(AValue int32) {
	LCL().SysCallN(4476, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginLeft() int32 {
	r1 := LCL().SysCallN(4468, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginLeft(AValue int32) {
	LCL().SysCallN(4468, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginTop() int32 {
	r1 := LCL().SysCallN(4470, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginTop(AValue int32) {
	LCL().SysCallN(4470, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginRight() int32 {
	r1 := LCL().SysCallN(4469, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginRight(AValue int32) {
	LCL().SysCallN(4469, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MarginBottom() int32 {
	r1 := LCL().SysCallN(4467, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMarginBottom(AValue int32) {
	LCL().SysCallN(4467, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginLeft() int32 {
	r1 := LCL().SysCallN(4472, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginLeft(AValue int32) {
	LCL().SysCallN(4472, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginTop() int32 {
	r1 := LCL().SysCallN(4474, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginTop(AValue int32) {
	LCL().SysCallN(4474, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginRight() int32 {
	r1 := LCL().SysCallN(4473, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginRight(AValue int32) {
	LCL().SysCallN(4473, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) MinMarginBottom() int32 {
	r1 := LCL().SysCallN(4471, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageSetupDialog) SetMinMarginBottom(AValue int32) {
	LCL().SysCallN(4471, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) Options() TPageSetupDialogOptions {
	r1 := LCL().SysCallN(4475, 0, m.Instance(), 0)
	return TPageSetupDialogOptions(r1)
}

func (m *TPageSetupDialog) SetOptions(AValue TPageSetupDialogOptions) {
	LCL().SysCallN(4475, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageSetupDialog) Units() TPageMeasureUnits {
	r1 := LCL().SysCallN(4479, 0, m.Instance(), 0)
	return TPageMeasureUnits(r1)
}

func (m *TPageSetupDialog) SetUnits(AValue TPageMeasureUnits) {
	LCL().SysCallN(4479, 1, m.Instance(), uintptr(AValue))
}

func PageSetupDialogClass() TClass {
	ret := LCL().SysCallN(4465)
	return TClass(ret)
}

func (m *TPageSetupDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if m.dialogResultPtr != 0 {
		RemoveEventElement(m.dialogResultPtr)
	}
	m.dialogResultPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4478, m.Instance(), m.dialogResultPtr)
}
