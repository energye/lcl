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

// IATGauge Parent: IGraphicControl
type IATGauge interface {
	IGraphicControl
	AddProgress(value int32)                       // procedure
	PercentDone() int32                            // property PercentDone Getter
	Theme() types.PATFlatTheme                     // property Theme Getter
	SetTheme(value types.PATFlatTheme)             // property Theme Setter
	BorderStyle() types.TBorderStyle               // property BorderStyle Getter
	SetBorderStyle(value types.TBorderStyle)       // property BorderStyle Setter
	DoubleBuffered() bool                          // property DoubleBuffered Getter
	SetDoubleBuffered(value bool)                  // property DoubleBuffered Setter
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	KindToATGaugeKind() types.TATGaugeKind         // property Kind Getter
	SetKindToATGaugeKind(value types.TATGaugeKind) // property Kind Setter
	Progress() int32                               // property Progress Getter
	SetProgress(value int32)                       // property Progress Setter
	MinValue() int32                               // property MinValue Getter
	SetMinValue(value int32)                       // property MinValue Setter
	MaxValue() int32                               // property MaxValue Getter
	SetMaxValue(value int32)                       // property MaxValue Setter
	ShowText() bool                                // property ShowText Getter
	SetShowText(value bool)                        // property ShowText Setter
	ShowTextInverted() bool                        // property ShowTextInverted Getter
	SetShowTextInverted(value bool)                // property ShowTextInverted Setter
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
}

type TATGauge struct {
	TGraphicControl
}

func (m *TATGauge) AddProgress(value int32) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(1, m.Instance(), uintptr(value))
}

func (m *TATGauge) PercentDone() int32 {
	if !m.IsValid() {
		return 0
	}
	r := aTGaugeAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TATGauge) Theme() types.PATFlatTheme {
	if !m.IsValid() {
		return 0
	}
	r := aTGaugeAPI().SysCallN(3, 0, m.Instance())
	return types.PATFlatTheme(r)
}

func (m *TATGauge) SetTheme(value types.PATFlatTheme) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(3, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TATGauge) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := aTGaugeAPI().SysCallN(4, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TATGauge) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TATGauge) DoubleBuffered() bool {
	if !m.IsValid() {
		return false
	}
	r := aTGaugeAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TATGauge) SetDoubleBuffered(value bool) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TATGauge) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := aTGaugeAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TATGauge) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TATGauge) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := aTGaugeAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TATGauge) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TATGauge) KindToATGaugeKind() types.TATGaugeKind {
	if !m.IsValid() {
		return 0
	}
	r := aTGaugeAPI().SysCallN(8, 0, m.Instance())
	return types.TATGaugeKind(r)
}

func (m *TATGauge) SetKindToATGaugeKind(value types.TATGaugeKind) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TATGauge) Progress() int32 {
	if !m.IsValid() {
		return 0
	}
	r := aTGaugeAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TATGauge) SetProgress(value int32) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TATGauge) MinValue() int32 {
	if !m.IsValid() {
		return 0
	}
	r := aTGaugeAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TATGauge) SetMinValue(value int32) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TATGauge) MaxValue() int32 {
	if !m.IsValid() {
		return 0
	}
	r := aTGaugeAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TATGauge) SetMaxValue(value int32) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TATGauge) ShowText() bool {
	if !m.IsValid() {
		return false
	}
	r := aTGaugeAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TATGauge) SetShowText(value bool) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TATGauge) ShowTextInverted() bool {
	if !m.IsValid() {
		return false
	}
	r := aTGaugeAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TATGauge) SetShowTextInverted(value bool) {
	if !m.IsValid() {
		return
	}
	aTGaugeAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TATGauge) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 15, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 16, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 18, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 21, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 22, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TATGauge) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 23, aTGaugeAPI(), api.MakeEventDataPtr(cb))
}

// NewATGauge class constructor
func NewATGauge(owner IComponent) IATGauge {
	r := aTGaugeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsATGauge(r)
}

func TATGaugeClass() types.TClass {
	r := aTGaugeAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	aTGaugeOnce   base.Once
	aTGaugeImport *imports.Imports = nil
)

func aTGaugeAPI() *imports.Imports {
	aTGaugeOnce.Do(func() {
		aTGaugeImport = api.NewDefaultImports()
		aTGaugeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TATGauge_Create", 0), // constructor NewATGauge
			/* 1 */ imports.NewTable("TATGauge_AddProgress", 0), // procedure AddProgress
			/* 2 */ imports.NewTable("TATGauge_PercentDone", 0), // property PercentDone
			/* 3 */ imports.NewTable("TATGauge_Theme", 0), // property Theme
			/* 4 */ imports.NewTable("TATGauge_BorderStyle", 0), // property BorderStyle
			/* 5 */ imports.NewTable("TATGauge_DoubleBuffered", 0), // property DoubleBuffered
			/* 6 */ imports.NewTable("TATGauge_ParentColor", 0), // property ParentColor
			/* 7 */ imports.NewTable("TATGauge_ParentShowHint", 0), // property ParentShowHint
			/* 8 */ imports.NewTable("TATGauge_KindToATGaugeKind", 0), // property KindToATGaugeKind
			/* 9 */ imports.NewTable("TATGauge_Progress", 0), // property Progress
			/* 10 */ imports.NewTable("TATGauge_MinValue", 0), // property MinValue
			/* 11 */ imports.NewTable("TATGauge_MaxValue", 0), // property MaxValue
			/* 12 */ imports.NewTable("TATGauge_ShowText", 0), // property ShowText
			/* 13 */ imports.NewTable("TATGauge_ShowTextInverted", 0), // property ShowTextInverted
			/* 14 */ imports.NewTable("TATGauge_OnDblClick", 0), // event OnDblClick
			/* 15 */ imports.NewTable("TATGauge_OnContextPopup", 0), // event OnContextPopup
			/* 16 */ imports.NewTable("TATGauge_OnMouseDown", 0), // event OnMouseDown
			/* 17 */ imports.NewTable("TATGauge_OnMouseUp", 0), // event OnMouseUp
			/* 18 */ imports.NewTable("TATGauge_OnMouseMove", 0), // event OnMouseMove
			/* 19 */ imports.NewTable("TATGauge_OnMouseEnter", 0), // event OnMouseEnter
			/* 20 */ imports.NewTable("TATGauge_OnMouseLeave", 0), // event OnMouseLeave
			/* 21 */ imports.NewTable("TATGauge_OnMouseWheel", 0), // event OnMouseWheel
			/* 22 */ imports.NewTable("TATGauge_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 23 */ imports.NewTable("TATGauge_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 24 */ imports.NewTable("TATGauge_TClass", 0), // function TATGaugeClass
		}
	})
	return aTGaugeImport
}
