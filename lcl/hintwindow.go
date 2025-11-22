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

// IHintWindow Parent: ICustomForm
type IHintWindow interface {
	ICustomForm
	CalcHintRect(maxWidth int32, hint string, data uintptr) types.TRect                 // function
	OffsetHintRect(offset types.TPoint, dy int32, keepWidth bool, keepHeight bool) bool // function
	IsHintMsg(msg types.TMsg) bool                                                      // function
	ActivateHintWithString(hint string)                                                 // procedure
	ActivateHintWithRectString(rect types.TRect, hint string)                           // procedure
	ActivateWithBounds(rect types.TRect, hint string)                                   // procedure
	ActivateHintData(rect types.TRect, hint string, data uintptr)                       // procedure
	InitializeWnd()                                                                     // procedure
	ReleaseHandle()                                                                     // procedure
	Alignment() types.TAlignment                                                        // property Alignment Getter
	SetAlignment(value types.TAlignment)                                                // property Alignment Setter
	HintRect() types.TRect                                                              // property HintRect Getter
	SetHintRect(value types.TRect)                                                      // property HintRect Setter
	HintRectAdjust() types.TRect                                                        // property HintRectAdjust Getter
	SetHintRectAdjust(value types.TRect)                                                // property HintRectAdjust Setter
	HintData() uintptr                                                                  // property HintData Getter
	SetHintData(value uintptr)                                                          // property HintData Setter
	HintControl() IControl                                                              // property HintControl Getter
	SetHintControl(value IControl)                                                      // property HintControl Setter
	AutoHide() bool                                                                     // property AutoHide Getter
	SetAutoHide(value bool)                                                             // property AutoHide Setter
	HideInterval() int32                                                                // property HideInterval Getter
	SetHideInterval(value int32)                                                        // property HideInterval Setter
	SetOnMouseDown(fn TMouseEvent)                                                      // property event
	SetOnMouseUp(fn TMouseEvent)                                                        // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                                  // property event
	SetOnMouseLeave(fn TNotifyEvent)                                                    // property event
}

type THintWindow struct {
	TCustomForm
}

func (m *THintWindow) CalcHintRect(maxWidth int32, hint string, data uintptr) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(1, m.Instance(), uintptr(maxWidth), api.PasStr(hint), uintptr(data), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *THintWindow) OffsetHintRect(offset types.TPoint, dy int32, keepWidth bool, keepHeight bool) bool {
	if !m.IsValid() {
		return false
	}
	r := hintWindowAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&offset)), uintptr(dy), api.PasBool(keepWidth), api.PasBool(keepHeight))
	return api.GoBool(r)
}

func (m *THintWindow) IsHintMsg(msg types.TMsg) bool {
	if !m.IsValid() {
		return false
	}
	r := hintWindowAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&msg)))
	return api.GoBool(r)
}

func (m *THintWindow) ActivateHintWithString(hint string) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(4, m.Instance(), api.PasStr(hint))
}

func (m *THintWindow) ActivateHintWithRectString(rect types.TRect, hint string) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&rect)), api.PasStr(hint))
}

func (m *THintWindow) ActivateWithBounds(rect types.TRect, hint string) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&rect)), api.PasStr(hint))
}

func (m *THintWindow) ActivateHintData(rect types.TRect, hint string, data uintptr) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&rect)), api.PasStr(hint), uintptr(data))
}

func (m *THintWindow) InitializeWnd() {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(8, m.Instance())
}

func (m *THintWindow) ReleaseHandle() {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(9, m.Instance())
}

func (m *THintWindow) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := hintWindowAPI().SysCallN(10, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *THintWindow) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *THintWindow) HintRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *THintWindow) SetHintRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *THintWindow) HintRectAdjust() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(12, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *THintWindow) SetHintRectAdjust(value types.TRect) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(12, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *THintWindow) HintData() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := hintWindowAPI().SysCallN(13, 0, m.Instance())
	return uintptr(r)
}

func (m *THintWindow) SetHintData(value uintptr) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *THintWindow) HintControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := hintWindowAPI().SysCallN(14, 0, m.Instance())
	return AsControl(r)
}

func (m *THintWindow) SetHintControl(value IControl) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *THintWindow) AutoHide() bool {
	if !m.IsValid() {
		return false
	}
	r := hintWindowAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *THintWindow) SetAutoHide(value bool) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *THintWindow) HideInterval() int32 {
	if !m.IsValid() {
		return 0
	}
	r := hintWindowAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *THintWindow) SetHideInterval(value int32) {
	if !m.IsValid() {
		return
	}
	hintWindowAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *THintWindow) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, hintWindowAPI(), api.MakeEventDataPtr(cb))
}

func (m *THintWindow) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, hintWindowAPI(), api.MakeEventDataPtr(cb))
}

func (m *THintWindow) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 19, hintWindowAPI(), api.MakeEventDataPtr(cb))
}

func (m *THintWindow) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, hintWindowAPI(), api.MakeEventDataPtr(cb))
}

// NewHintWindow class constructor
func NewHintWindow(owner IComponent) IHintWindow {
	r := hintWindowAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsHintWindow(r)
}

func THintWindowClass() types.TClass {
	r := hintWindowAPI().SysCallN(21)
	return types.TClass(r)
}

var (
	hintWindowOnce   base.Once
	hintWindowImport *imports.Imports = nil
)

func hintWindowAPI() *imports.Imports {
	hintWindowOnce.Do(func() {
		hintWindowImport = api.NewDefaultImports()
		hintWindowImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THintWindow_Create", 0), // constructor NewHintWindow
			/* 1 */ imports.NewTable("THintWindow_CalcHintRect", 0), // function CalcHintRect
			/* 2 */ imports.NewTable("THintWindow_OffsetHintRect", 0), // function OffsetHintRect
			/* 3 */ imports.NewTable("THintWindow_IsHintMsg", 0), // function IsHintMsg
			/* 4 */ imports.NewTable("THintWindow_ActivateHintWithString", 0), // procedure ActivateHintWithString
			/* 5 */ imports.NewTable("THintWindow_ActivateHintWithRectString", 0), // procedure ActivateHintWithRectString
			/* 6 */ imports.NewTable("THintWindow_ActivateWithBounds", 0), // procedure ActivateWithBounds
			/* 7 */ imports.NewTable("THintWindow_ActivateHintData", 0), // procedure ActivateHintData
			/* 8 */ imports.NewTable("THintWindow_InitializeWnd", 0), // procedure InitializeWnd
			/* 9 */ imports.NewTable("THintWindow_ReleaseHandle", 0), // procedure ReleaseHandle
			/* 10 */ imports.NewTable("THintWindow_Alignment", 0), // property Alignment
			/* 11 */ imports.NewTable("THintWindow_HintRect", 0), // property HintRect
			/* 12 */ imports.NewTable("THintWindow_HintRectAdjust", 0), // property HintRectAdjust
			/* 13 */ imports.NewTable("THintWindow_HintData", 0), // property HintData
			/* 14 */ imports.NewTable("THintWindow_HintControl", 0), // property HintControl
			/* 15 */ imports.NewTable("THintWindow_AutoHide", 0), // property AutoHide
			/* 16 */ imports.NewTable("THintWindow_HideInterval", 0), // property HideInterval
			/* 17 */ imports.NewTable("THintWindow_OnMouseDown", 0), // event OnMouseDown
			/* 18 */ imports.NewTable("THintWindow_OnMouseUp", 0), // event OnMouseUp
			/* 19 */ imports.NewTable("THintWindow_OnMouseMove", 0), // event OnMouseMove
			/* 20 */ imports.NewTable("THintWindow_OnMouseLeave", 0), // event OnMouseLeave
			/* 21 */ imports.NewTable("THintWindow_TClass", 0), // function THintWindowClass
		}
	})
	return hintWindowImport
}
