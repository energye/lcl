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

// IFontDialog Parent: ICommonDialog
type IFontDialog interface {
	ICommonDialog
	Font() IFont                          // property
	SetFont(AValue IFont)                 // property
	MinFontSize() int32                   // property
	SetMinFontSize(AValue int32)          // property
	MaxFontSize() int32                   // property
	SetMaxFontSize(AValue int32)          // property
	Options() TFontDialogOptions          // property
	SetOptions(AValue TFontDialogOptions) // property
	PreviewText() string                  // property
	SetPreviewText(AValue string)         // property
	ApplyClicked()                        // procedure
	SetOnApplyClicked(fn TNotifyEvent)    // property event
}

// TFontDialog Parent: TCommonDialog
type TFontDialog struct {
	TCommonDialog
	applyClickedPtr uintptr
}

func NewFontDialog(AOwner IComponent) IFontDialog {
	r1 := fontDialogImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsFontDialog(r1)
}

func (m *TFontDialog) Font() IFont {
	r1 := fontDialogImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TFontDialog) SetFont(AValue IFont) {
	fontDialogImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFontDialog) MinFontSize() int32 {
	r1 := fontDialogImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFontDialog) SetMinFontSize(AValue int32) {
	fontDialogImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) MaxFontSize() int32 {
	r1 := fontDialogImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFontDialog) SetMaxFontSize(AValue int32) {
	fontDialogImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) Options() TFontDialogOptions {
	r1 := fontDialogImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TFontDialogOptions(r1)
}

func (m *TFontDialog) SetOptions(AValue TFontDialogOptions) {
	fontDialogImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) PreviewText() string {
	r1 := fontDialogImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFontDialog) SetPreviewText(AValue string) {
	fontDialogImportAPI().SysCallN(7, 1, m.Instance(), PascalStr(AValue))
}

func FontDialogClass() TClass {
	ret := fontDialogImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TFontDialog) ApplyClicked() {
	fontDialogImportAPI().SysCallN(0, m.Instance())
}

func (m *TFontDialog) SetOnApplyClicked(fn TNotifyEvent) {
	if m.applyClickedPtr != 0 {
		RemoveEventElement(m.applyClickedPtr)
	}
	m.applyClickedPtr = MakeEventDataPtr(fn)
	fontDialogImportAPI().SysCallN(8, m.Instance(), m.applyClickedPtr)
}

var (
	fontDialogImport       *imports.Imports = nil
	fontDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FontDialog_ApplyClicked", 0),
		/*1*/ imports.NewTable("FontDialog_Class", 0),
		/*2*/ imports.NewTable("FontDialog_Create", 0),
		/*3*/ imports.NewTable("FontDialog_Font", 0),
		/*4*/ imports.NewTable("FontDialog_MaxFontSize", 0),
		/*5*/ imports.NewTable("FontDialog_MinFontSize", 0),
		/*6*/ imports.NewTable("FontDialog_Options", 0),
		/*7*/ imports.NewTable("FontDialog_PreviewText", 0),
		/*8*/ imports.NewTable("FontDialog_SetOnApplyClicked", 0),
	}
)

func fontDialogImportAPI() *imports.Imports {
	if fontDialogImport == nil {
		fontDialogImport = NewDefaultImports()
		fontDialogImport.SetImportTable(fontDialogImportTables)
		fontDialogImportTables = nil
	}
	return fontDialogImport
}
