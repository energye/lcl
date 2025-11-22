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

// IXButton Parent: IGraphicControl
type IXButton interface {
	IGraphicControl
	Paint()                                 // procedure
	Resize()                                // procedure
	CaptionToString() string                // property Caption Getter
	SetCaptionToString(value string)        // property Caption Setter
	ShowCaption() bool                      // property ShowCaption Getter
	SetShowCaption(value bool)              // property ShowCaption Setter
	BackColor() types.TColor                // property BackColor Getter
	SetBackColor(value types.TColor)        // property BackColor Setter
	HoverColor() types.TColor               // property HoverColor Getter
	SetHoverColor(value types.TColor)       // property HoverColor Setter
	DownColor() types.TColor                // property DownColor Getter
	SetDownColor(value types.TColor)        // property DownColor Setter
	BorderWidth() int32                     // property BorderWidth Getter
	SetBorderWidth(value int32)             // property BorderWidth Setter
	BorderColor() types.TColor              // property BorderColor Getter
	SetBorderColor(value types.TColor)      // property BorderColor Setter
	Picture() IPicture                      // property Picture Getter
	SetPicture(value IPicture)              // property Picture Setter
	DrawMode() types.TDrawImageMode         // property DrawMode Getter
	SetDrawMode(value types.TDrawImageMode) // property DrawMode Setter
	NormalFontColor() types.TColor          // property NormalFontColor Getter
	SetNormalFontColor(value types.TColor)  // property NormalFontColor Setter
	DownFontColor() types.TColor            // property DownFontColor Getter
	SetDownFontColor(value types.TColor)    // property DownFontColor Setter
	HoverFontColor() types.TColor           // property HoverFontColor Getter
	SetHoverFontColor(value types.TColor)   // property HoverFontColor Setter
	ParentFont() bool                       // property ParentFont Getter
	SetParentFont(value bool)               // property ParentFont Setter
	ParentShowHint() bool                   // property ParentShowHint Getter
	SetParentShowHint(value bool)           // property ParentShowHint Setter
	SetOnDblClick(fn TNotifyEvent)          // property event
	SetOnMouseDown(fn TMouseEvent)          // property event
	SetOnMouseEnter(fn TNotifyEvent)        // property event
	SetOnMouseLeave(fn TNotifyEvent)        // property event
	SetOnMouseMove(fn TMouseMoveEvent)      // property event
	SetOnMouseUp(fn TMouseEvent)            // property event
}

type TXButton struct {
	TGraphicControl
}

func (m *TXButton) Paint() {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(1, m.Instance())
}

func (m *TXButton) Resize() {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(2, m.Instance())
}

func (m *TXButton) CaptionToString() string {
	if !m.IsValid() {
		return ""
	}
	r := xButtonAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TXButton) SetCaptionToString(value string) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TXButton) ShowCaption() bool {
	if !m.IsValid() {
		return false
	}
	r := xButtonAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TXButton) SetShowCaption(value bool) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TXButton) BackColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(5, 0, m.Instance())
	return types.TColor(r)
}

func (m *TXButton) SetBackColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) HoverColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(6, 0, m.Instance())
	return types.TColor(r)
}

func (m *TXButton) SetHoverColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) DownColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(7, 0, m.Instance())
	return types.TColor(r)
}

func (m *TXButton) SetDownColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) BorderWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TXButton) SetBorderWidth(value int32) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) BorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(9, 0, m.Instance())
	return types.TColor(r)
}

func (m *TXButton) SetBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) Picture() IPicture {
	if !m.IsValid() {
		return nil
	}
	r := xButtonAPI().SysCallN(10, 0, m.Instance())
	return AsPicture(r)
}

func (m *TXButton) SetPicture(value IPicture) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TXButton) DrawMode() types.TDrawImageMode {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(11, 0, m.Instance())
	return types.TDrawImageMode(r)
}

func (m *TXButton) SetDrawMode(value types.TDrawImageMode) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) NormalFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(12, 0, m.Instance())
	return types.TColor(r)
}

func (m *TXButton) SetNormalFontColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) DownFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(13, 0, m.Instance())
	return types.TColor(r)
}

func (m *TXButton) SetDownFontColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) HoverFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := xButtonAPI().SysCallN(14, 0, m.Instance())
	return types.TColor(r)
}

func (m *TXButton) SetHoverFontColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TXButton) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := xButtonAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TXButton) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TXButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := xButtonAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TXButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	xButtonAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TXButton) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, xButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TXButton) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, xButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TXButton) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, xButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TXButton) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, xButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TXButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 21, xButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TXButton) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 22, xButtonAPI(), api.MakeEventDataPtr(cb))
}

// NewXButton class constructor
func NewXButton(owner IComponent) IXButton {
	r := xButtonAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsXButton(r)
}

func TXButtonClass() types.TClass {
	r := xButtonAPI().SysCallN(23)
	return types.TClass(r)
}

var (
	xButtonOnce   base.Once
	xButtonImport *imports.Imports = nil
)

func xButtonAPI() *imports.Imports {
	xButtonOnce.Do(func() {
		xButtonImport = api.NewDefaultImports()
		xButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TXButton_Create", 0), // constructor NewXButton
			/* 1 */ imports.NewTable("TXButton_Paint", 0), // procedure Paint
			/* 2 */ imports.NewTable("TXButton_Resize", 0), // procedure Resize
			/* 3 */ imports.NewTable("TXButton_CaptionToString", 0), // property CaptionToString
			/* 4 */ imports.NewTable("TXButton_ShowCaption", 0), // property ShowCaption
			/* 5 */ imports.NewTable("TXButton_BackColor", 0), // property BackColor
			/* 6 */ imports.NewTable("TXButton_HoverColor", 0), // property HoverColor
			/* 7 */ imports.NewTable("TXButton_DownColor", 0), // property DownColor
			/* 8 */ imports.NewTable("TXButton_BorderWidth", 0), // property BorderWidth
			/* 9 */ imports.NewTable("TXButton_BorderColor", 0), // property BorderColor
			/* 10 */ imports.NewTable("TXButton_Picture", 0), // property Picture
			/* 11 */ imports.NewTable("TXButton_DrawMode", 0), // property DrawMode
			/* 12 */ imports.NewTable("TXButton_NormalFontColor", 0), // property NormalFontColor
			/* 13 */ imports.NewTable("TXButton_DownFontColor", 0), // property DownFontColor
			/* 14 */ imports.NewTable("TXButton_HoverFontColor", 0), // property HoverFontColor
			/* 15 */ imports.NewTable("TXButton_ParentFont", 0), // property ParentFont
			/* 16 */ imports.NewTable("TXButton_ParentShowHint", 0), // property ParentShowHint
			/* 17 */ imports.NewTable("TXButton_OnDblClick", 0), // event OnDblClick
			/* 18 */ imports.NewTable("TXButton_OnMouseDown", 0), // event OnMouseDown
			/* 19 */ imports.NewTable("TXButton_OnMouseEnter", 0), // event OnMouseEnter
			/* 20 */ imports.NewTable("TXButton_OnMouseLeave", 0), // event OnMouseLeave
			/* 21 */ imports.NewTable("TXButton_OnMouseMove", 0), // event OnMouseMove
			/* 22 */ imports.NewTable("TXButton_OnMouseUp", 0), // event OnMouseUp
			/* 23 */ imports.NewTable("TXButton_TClass", 0), // function TXButtonClass
		}
	})
	return xButtonImport
}
