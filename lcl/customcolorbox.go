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

// ICustomColorBox Parent: ICustomComboBox
type ICustomColorBox interface {
	ICustomComboBox
	ColorRectWidth() int32                              // property ColorRectWidth Getter
	SetColorRectWidth(value int32)                      // property ColorRectWidth Setter
	ColorRectOffset() int32                             // property ColorRectOffset Getter
	SetColorRectOffset(value int32)                     // property ColorRectOffset Setter
	StyleToColorBoxStyle() types.TColorBoxStyle         // property Style Getter
	SetStyleToColorBoxStyle(value types.TColorBoxStyle) // property Style Setter
	Colors(index int32) types.TColor                    // property Colors Getter
	ColorNames(index int32) string                      // property ColorNames Getter
	Selected() types.TColor                             // property Selected Getter
	SetSelected(value types.TColor)                     // property Selected Setter
	DefaultColorColor() types.TColor                    // property DefaultColorColor Getter
	SetDefaultColorColor(value types.TColor)            // property DefaultColorColor Setter
	NoneColorColor() types.TColor                       // property NoneColorColor Getter
	SetNoneColorColor(value types.TColor)               // property NoneColorColor Setter
	ColorDialog() IColorDialog                          // property ColorDialog Getter
	SetColorDialog(value IColorDialog)                  // property ColorDialog Setter
	SetOnGetColors(fn TGetColorsEvent)                  // property event
}

type TCustomColorBox struct {
	TCustomComboBox
}

func (m *TCustomColorBox) ColorRectWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customColorBoxAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TCustomColorBox) SetColorRectWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customColorBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorBox) ColorRectOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customColorBoxAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TCustomColorBox) SetColorRectOffset(value int32) {
	if !m.IsValid() {
		return
	}
	customColorBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorBox) StyleToColorBoxStyle() types.TColorBoxStyle {
	if !m.IsValid() {
		return 0
	}
	r := customColorBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TColorBoxStyle(r)
}

func (m *TCustomColorBox) SetStyleToColorBoxStyle(value types.TColorBoxStyle) {
	if !m.IsValid() {
		return
	}
	customColorBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorBox) Colors(index int32) types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorBoxAPI().SysCallN(4, m.Instance(), uintptr(index))
	return types.TColor(r)
}

func (m *TCustomColorBox) ColorNames(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customColorBoxAPI().SysCallN(5, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TCustomColorBox) Selected() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorBoxAPI().SysCallN(6, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomColorBox) SetSelected(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customColorBoxAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorBox) DefaultColorColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorBoxAPI().SysCallN(7, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomColorBox) SetDefaultColorColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customColorBoxAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorBox) NoneColorColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customColorBoxAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomColorBox) SetNoneColorColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customColorBoxAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomColorBox) ColorDialog() IColorDialog {
	if !m.IsValid() {
		return nil
	}
	r := customColorBoxAPI().SysCallN(9, 0, m.Instance())
	return AsColorDialog(r)
}

func (m *TCustomColorBox) SetColorDialog(value IColorDialog) {
	if !m.IsValid() {
		return
	}
	customColorBoxAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomColorBox) SetOnGetColors(fn TGetColorsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetColorsEvent(fn)
	base.SetEvent(m, 10, customColorBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomColorBox class constructor
func NewCustomColorBox(owner IComponent) ICustomColorBox {
	r := customColorBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomColorBox(r)
}

func TCustomColorBoxClass() types.TClass {
	r := customColorBoxAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	customColorBoxOnce   base.Once
	customColorBoxImport *imports.Imports = nil
)

func customColorBoxAPI() *imports.Imports {
	customColorBoxOnce.Do(func() {
		customColorBoxImport = api.NewDefaultImports()
		customColorBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomColorBox_Create", 0), // constructor NewCustomColorBox
			/* 1 */ imports.NewTable("TCustomColorBox_ColorRectWidth", 0), // property ColorRectWidth
			/* 2 */ imports.NewTable("TCustomColorBox_ColorRectOffset", 0), // property ColorRectOffset
			/* 3 */ imports.NewTable("TCustomColorBox_StyleToColorBoxStyle", 0), // property StyleToColorBoxStyle
			/* 4 */ imports.NewTable("TCustomColorBox_Colors", 0), // property Colors
			/* 5 */ imports.NewTable("TCustomColorBox_ColorNames", 0), // property ColorNames
			/* 6 */ imports.NewTable("TCustomColorBox_Selected", 0), // property Selected
			/* 7 */ imports.NewTable("TCustomColorBox_DefaultColorColor", 0), // property DefaultColorColor
			/* 8 */ imports.NewTable("TCustomColorBox_NoneColorColor", 0), // property NoneColorColor
			/* 9 */ imports.NewTable("TCustomColorBox_ColorDialog", 0), // property ColorDialog
			/* 10 */ imports.NewTable("TCustomColorBox_OnGetColors", 0), // event OnGetColors
			/* 11 */ imports.NewTable("TCustomColorBox_TClass", 0), // function TCustomColorBoxClass
		}
	})
	return customColorBoxImport
}
