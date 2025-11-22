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

// IFontDialog Parent: ICommonDialog
type IFontDialog interface {
	ICommonDialog
	ApplyClicked()                             // procedure
	Font() IFont                               // property Font Getter
	SetFont(value IFont)                       // property Font Setter
	MinFontSize() int32                        // property MinFontSize Getter
	SetMinFontSize(value int32)                // property MinFontSize Setter
	MaxFontSize() int32                        // property MaxFontSize Getter
	SetMaxFontSize(value int32)                // property MaxFontSize Setter
	Options() types.TFontDialogOptions         // property Options Getter
	SetOptions(value types.TFontDialogOptions) // property Options Setter
	PreviewText() string                       // property PreviewText Getter
	SetPreviewText(value string)               // property PreviewText Setter
	SetOnApplyClicked(fn TNotifyEvent)         // property event
}

type TFontDialog struct {
	TCommonDialog
}

func (m *TFontDialog) ApplyClicked() {
	if !m.IsValid() {
		return
	}
	fontDialogAPI().SysCallN(1, m.Instance())
}

func (m *TFontDialog) Font() IFont {
	if !m.IsValid() {
		return nil
	}
	r := fontDialogAPI().SysCallN(2, 0, m.Instance())
	return AsFont(r)
}

func (m *TFontDialog) SetFont(value IFont) {
	if !m.IsValid() {
		return
	}
	fontDialogAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFontDialog) MinFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fontDialogAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TFontDialog) SetMinFontSize(value int32) {
	if !m.IsValid() {
		return
	}
	fontDialogAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TFontDialog) MaxFontSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fontDialogAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TFontDialog) SetMaxFontSize(value int32) {
	if !m.IsValid() {
		return
	}
	fontDialogAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TFontDialog) Options() types.TFontDialogOptions {
	if !m.IsValid() {
		return 0
	}
	r := fontDialogAPI().SysCallN(5, 0, m.Instance())
	return types.TFontDialogOptions(r)
}

func (m *TFontDialog) SetOptions(value types.TFontDialogOptions) {
	if !m.IsValid() {
		return
	}
	fontDialogAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TFontDialog) PreviewText() string {
	if !m.IsValid() {
		return ""
	}
	r := fontDialogAPI().SysCallN(6, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFontDialog) SetPreviewText(value string) {
	if !m.IsValid() {
		return
	}
	fontDialogAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TFontDialog) SetOnApplyClicked(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, fontDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewFontDialog class constructor
func NewFontDialog(owner IComponent) IFontDialog {
	r := fontDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsFontDialog(r)
}

func TFontDialogClass() types.TClass {
	r := fontDialogAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	fontDialogOnce   base.Once
	fontDialogImport *imports.Imports = nil
)

func fontDialogAPI() *imports.Imports {
	fontDialogOnce.Do(func() {
		fontDialogImport = api.NewDefaultImports()
		fontDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFontDialog_Create", 0), // constructor NewFontDialog
			/* 1 */ imports.NewTable("TFontDialog_ApplyClicked", 0), // procedure ApplyClicked
			/* 2 */ imports.NewTable("TFontDialog_Font", 0), // property Font
			/* 3 */ imports.NewTable("TFontDialog_MinFontSize", 0), // property MinFontSize
			/* 4 */ imports.NewTable("TFontDialog_MaxFontSize", 0), // property MaxFontSize
			/* 5 */ imports.NewTable("TFontDialog_Options", 0), // property Options
			/* 6 */ imports.NewTable("TFontDialog_PreviewText", 0), // property PreviewText
			/* 7 */ imports.NewTable("TFontDialog_OnApplyClicked", 0), // event OnApplyClicked
			/* 8 */ imports.NewTable("TFontDialog_TClass", 0), // function TFontDialogClass
		}
	})
	return fontDialogImport
}
