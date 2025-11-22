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

// IColorButton Parent: ICustomSpeedButton
type IColorButton interface {
	ICustomSpeedButton
	BorderWidth() int32                            // property BorderWidth Getter
	SetBorderWidth(value int32)                    // property BorderWidth Setter
	ButtonColorAutoSize() bool                     // property ButtonColorAutoSize Getter
	SetButtonColorAutoSize(value bool)             // property ButtonColorAutoSize Setter
	ButtonColorSize() int32                        // property ButtonColorSize Getter
	SetButtonColorSize(value int32)                // property ButtonColorSize Setter
	ButtonColor() types.TColor                     // property ButtonColor Getter
	SetButtonColor(value types.TColor)             // property ButtonColor Setter
	ColorDialog() IColorDialog                     // property ColorDialog Getter
	SetColorDialog(value IColorDialog)             // property ColorDialog Setter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnColorChanged(fn TNotifyEvent)             // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnPaint(fn TNotifyEvent)                    // property event
}

type TColorButton struct {
	TCustomSpeedButton
}

func (m *TColorButton) BorderWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := colorButtonAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TColorButton) SetBorderWidth(value int32) {
	if !m.IsValid() {
		return
	}
	colorButtonAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TColorButton) ButtonColorAutoSize() bool {
	if !m.IsValid() {
		return false
	}
	r := colorButtonAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TColorButton) SetButtonColorAutoSize(value bool) {
	if !m.IsValid() {
		return
	}
	colorButtonAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TColorButton) ButtonColorSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := colorButtonAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TColorButton) SetButtonColorSize(value int32) {
	if !m.IsValid() {
		return
	}
	colorButtonAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TColorButton) ButtonColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := colorButtonAPI().SysCallN(4, 0, m.Instance())
	return types.TColor(r)
}

func (m *TColorButton) SetButtonColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	colorButtonAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TColorButton) ColorDialog() IColorDialog {
	if !m.IsValid() {
		return nil
	}
	r := colorButtonAPI().SysCallN(5, 0, m.Instance())
	return AsColorDialog(r)
}

func (m *TColorButton) SetColorDialog(value IColorDialog) {
	if !m.IsValid() {
		return
	}
	colorButtonAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TColorButton) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := colorButtonAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TColorButton) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	colorButtonAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TColorButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := colorButtonAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TColorButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	colorButtonAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TColorButton) SetOnColorChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 10, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 13, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorButton) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, colorButtonAPI(), api.MakeEventDataPtr(cb))
}

// NewColorButton class constructor
func NewColorButton(anOwner IComponent) IColorButton {
	r := colorButtonAPI().SysCallN(0, base.GetObjectUintptr(anOwner))
	return AsColorButton(r)
}

func TColorButtonClass() types.TClass {
	r := colorButtonAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	colorButtonOnce   base.Once
	colorButtonImport *imports.Imports = nil
)

func colorButtonAPI() *imports.Imports {
	colorButtonOnce.Do(func() {
		colorButtonImport = api.NewDefaultImports()
		colorButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TColorButton_Create", 0), // constructor NewColorButton
			/* 1 */ imports.NewTable("TColorButton_BorderWidth", 0), // property BorderWidth
			/* 2 */ imports.NewTable("TColorButton_ButtonColorAutoSize", 0), // property ButtonColorAutoSize
			/* 3 */ imports.NewTable("TColorButton_ButtonColorSize", 0), // property ButtonColorSize
			/* 4 */ imports.NewTable("TColorButton_ButtonColor", 0), // property ButtonColor
			/* 5 */ imports.NewTable("TColorButton_ColorDialog", 0), // property ColorDialog
			/* 6 */ imports.NewTable("TColorButton_ParentFont", 0), // property ParentFont
			/* 7 */ imports.NewTable("TColorButton_ParentShowHint", 0), // property ParentShowHint
			/* 8 */ imports.NewTable("TColorButton_OnColorChanged", 0), // event OnColorChanged
			/* 9 */ imports.NewTable("TColorButton_OnDblClick", 0), // event OnDblClick
			/* 10 */ imports.NewTable("TColorButton_OnMouseDown", 0), // event OnMouseDown
			/* 11 */ imports.NewTable("TColorButton_OnMouseEnter", 0), // event OnMouseEnter
			/* 12 */ imports.NewTable("TColorButton_OnMouseLeave", 0), // event OnMouseLeave
			/* 13 */ imports.NewTable("TColorButton_OnMouseMove", 0), // event OnMouseMove
			/* 14 */ imports.NewTable("TColorButton_OnMouseUp", 0), // event OnMouseUp
			/* 15 */ imports.NewTable("TColorButton_OnMouseWheel", 0), // event OnMouseWheel
			/* 16 */ imports.NewTable("TColorButton_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 17 */ imports.NewTable("TColorButton_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 18 */ imports.NewTable("TColorButton_OnPaint", 0), // event OnPaint
			/* 19 */ imports.NewTable("TColorButton_TClass", 0), // function TColorButtonClass
		}
	})
	return colorButtonImport
}
