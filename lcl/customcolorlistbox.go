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

// ICustomColorListBox Parent: ICustomListBox
type ICustomColorListBox interface {
	ICustomListBox
	ColorRectWidth() int32                              // property ColorRectWidth Getter
	SetColorRectWidth(value int32)                      // property ColorRectWidth Setter
	ColorRectOffset() int32                             // property ColorRectOffset Getter
	SetColorRectOffset(value int32)                     // property ColorRectOffset Setter
	StyleToColorBoxStyle() types.TColorBoxStyle         // property Style Getter
	SetStyleToColorBoxStyle(value types.TColorBoxStyle) // property Style Setter
	Colors(index int32) types.TColor                    // property Colors Getter
	SetColors(index int32, value types.TColor)          // property Colors Setter
	ColorNames(index int32) string                      // property ColorNames Getter
	SelectedToColor() types.TColor                      // property Selected Getter
	SetSelectedToColor(value types.TColor)              // property Selected Setter
	DefaultColorColor() types.TColor                    // property DefaultColorColor Getter
	SetDefaultColorColor(value types.TColor)            // property DefaultColorColor Setter
	NoneColorColor() types.TColor                       // property NoneColorColor Getter
	SetNoneColorColor(value types.TColor)               // property NoneColorColor Setter
	ColorDialog() IColorDialog                          // property ColorDialog Getter
	SetColorDialog(value IColorDialog)                  // property ColorDialog Setter
	SetOnGetColors(fn TLBGetColorsEvent)                // property event
}

type TCustomColorListBox struct {
	TCustomListBox
}

func (m *TCustomColorListBox) ColorRectWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customColorListBoxAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TCustomColorListBox) SetColorRectWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorListBox) ColorRectOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customColorListBoxAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TCustomColorListBox) SetColorRectOffset(value int32) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorListBox) StyleToColorBoxStyle() types.TColorBoxStyle {
	if !m.IsValid() {
		return 0
	}
	r := customColorListBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TColorBoxStyle(r)
}

func (m *TCustomColorListBox) SetStyleToColorBoxStyle(value types.TColorBoxStyle) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorListBox) Colors(index int32) types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorListBoxAPI().SysCallN(4, 0, m.Instance(), uintptr(index))
	return types.TColor(r)
}

func (m *TCustomColorListBox) SetColors(index int32, value types.TColor) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(4, 1, m.Instance(), uintptr(index), uintptr(value))
}

func (m *TCustomColorListBox) ColorNames(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customColorListBoxAPI().SysCallN(5, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TCustomColorListBox) SelectedToColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorListBoxAPI().SysCallN(6, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomColorListBox) SetSelectedToColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorListBox) DefaultColorColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorListBoxAPI().SysCallN(7, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomColorListBox) SetDefaultColorColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorListBox) NoneColorColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorListBoxAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomColorListBox) SetNoneColorColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorListBox) ColorDialog() IColorDialog {
	if !m.IsValid() {
		return nil
	}
	r := customColorListBoxAPI().SysCallN(9, 0, m.Instance())
	return AsColorDialog(r)
}

func (m *TCustomColorListBox) SetColorDialog(value IColorDialog) {
	if !m.IsValid() {
		return
	}
	customColorListBoxAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomColorListBox) SetOnGetColors(fn TLBGetColorsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLBGetColorsEvent(fn)
	base.SetEvent(m, 10, customColorListBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomColorListBox class constructor
func NewCustomColorListBox(owner IComponent) ICustomColorListBox {
	r := customColorListBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomColorListBox(r)
}

func TCustomColorListBoxClass() types.TClass {
	r := customColorListBoxAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	customColorListBoxOnce   base.Once
	customColorListBoxImport *imports.Imports = nil
)

func customColorListBoxAPI() *imports.Imports {
	customColorListBoxOnce.Do(func() {
		customColorListBoxImport = api.NewDefaultImports()
		customColorListBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomColorListBox_Create", 0), // constructor NewCustomColorListBox
			/* 1 */ imports.NewTable("TCustomColorListBox_ColorRectWidth", 0), // property ColorRectWidth
			/* 2 */ imports.NewTable("TCustomColorListBox_ColorRectOffset", 0), // property ColorRectOffset
			/* 3 */ imports.NewTable("TCustomColorListBox_StyleToColorBoxStyle", 0), // property StyleToColorBoxStyle
			/* 4 */ imports.NewTable("TCustomColorListBox_Colors", 0), // property Colors
			/* 5 */ imports.NewTable("TCustomColorListBox_ColorNames", 0), // property ColorNames
			/* 6 */ imports.NewTable("TCustomColorListBox_SelectedToColor", 0), // property SelectedToColor
			/* 7 */ imports.NewTable("TCustomColorListBox_DefaultColorColor", 0), // property DefaultColorColor
			/* 8 */ imports.NewTable("TCustomColorListBox_NoneColorColor", 0), // property NoneColorColor
			/* 9 */ imports.NewTable("TCustomColorListBox_ColorDialog", 0), // property ColorDialog
			/* 10 */ imports.NewTable("TCustomColorListBox_OnGetColors", 0), // event OnGetColors
			/* 11 */ imports.NewTable("TCustomColorListBox_TClass", 0), // function TCustomColorListBoxClass
		}
	})
	return customColorListBoxImport
}
