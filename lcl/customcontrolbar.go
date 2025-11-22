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

// ICustomControlBar Parent: ICustomPanel
type ICustomControlBar interface {
	ICustomPanel
	HitTest(X int32, Y int32) IControl                           // function
	MouseToBandPos(X int32, Y int32, outGrabber *bool) ICtrlBand // function
	BeginUpdate()                                                // procedure
	EndUpdate()                                                  // procedure
	StickControls()                                              // procedure
	AutoDock() bool                                              // property AutoDock Getter
	SetAutoDock(value bool)                                      // property AutoDock Setter
	AutoDrag() bool                                              // property AutoDrag Getter
	SetAutoDrag(value bool)                                      // property AutoDrag Setter
	DrawingStyle() types.TBandDrawingStyle                       // property DrawingStyle Getter
	SetDrawingStyle(value types.TBandDrawingStyle)               // property DrawingStyle Setter
	GradientDirection() types.TGradientDirection                 // property GradientDirection Getter
	SetGradientDirection(value types.TGradientDirection)         // property GradientDirection Setter
	GradientStartColor() types.TColor                            // property GradientStartColor Getter
	SetGradientStartColor(value types.TColor)                    // property GradientStartColor Setter
	GradientEndColor() types.TColor                              // property GradientEndColor Getter
	SetGradientEndColor(value types.TColor)                      // property GradientEndColor Setter
	Picture() IPicture                                           // property Picture Getter
	SetPicture(value IPicture)                                   // property Picture Setter
	RowSize() types.TRowSize                                     // property RowSize Getter
	SetRowSize(value types.TRowSize)                             // property RowSize Setter
	RowSnap() bool                                               // property RowSnap Getter
	SetRowSnap(value bool)                                       // property RowSnap Setter
	SetOnBandDrag(fn TBandDragEvent)                             // property event
	SetOnBandInfo(fn TBandInfoEvent)                             // property event
	SetOnBandMove(fn TBandMoveEvent)                             // property event
	SetOnBandPaint(fn TBandPaintEvent)                           // property event
	SetOnCanResize(fn TCanResizeEvent)                           // property event
}

type TCustomControlBar struct {
	TCustomPanel
}

func (m *TCustomControlBar) HitTest(X int32, Y int32) IControl {
	if !m.IsValid() {
		return nil
	}
	r := customControlBarAPI().SysCallN(1, m.Instance(), uintptr(X), uintptr(Y))
	return AsControl(r)
}

func (m *TCustomControlBar) MouseToBandPos(X int32, Y int32, outGrabber *bool) ICtrlBand {
	if !m.IsValid() {
		return nil
	}
	r := customControlBarAPI().SysCallN(2, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(outGrabber)))
	return AsCtrlBand(r)
}

func (m *TCustomControlBar) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(3, m.Instance())
}

func (m *TCustomControlBar) EndUpdate() {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(4, m.Instance())
}

func (m *TCustomControlBar) StickControls() {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(5, m.Instance())
}

func (m *TCustomControlBar) AutoDock() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlBarAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlBar) SetAutoDock(value bool) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlBar) AutoDrag() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlBarAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlBar) SetAutoDrag(value bool) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlBar) DrawingStyle() types.TBandDrawingStyle {
	if !m.IsValid() {
		return 0
	}
	r := customControlBarAPI().SysCallN(8, 0, m.Instance())
	return types.TBandDrawingStyle(r)
}

func (m *TCustomControlBar) SetDrawingStyle(value types.TBandDrawingStyle) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlBar) GradientDirection() types.TGradientDirection {
	if !m.IsValid() {
		return 0
	}
	r := customControlBarAPI().SysCallN(9, 0, m.Instance())
	return types.TGradientDirection(r)
}

func (m *TCustomControlBar) SetGradientDirection(value types.TGradientDirection) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlBar) GradientStartColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customControlBarAPI().SysCallN(10, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomControlBar) SetGradientStartColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlBar) GradientEndColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customControlBarAPI().SysCallN(11, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomControlBar) SetGradientEndColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlBar) Picture() IPicture {
	if !m.IsValid() {
		return nil
	}
	r := customControlBarAPI().SysCallN(12, 0, m.Instance())
	return AsPicture(r)
}

func (m *TCustomControlBar) SetPicture(value IPicture) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomControlBar) RowSize() types.TRowSize {
	if !m.IsValid() {
		return 0
	}
	r := customControlBarAPI().SysCallN(13, 0, m.Instance())
	return types.TRowSize(r)
}

func (m *TCustomControlBar) SetRowSize(value types.TRowSize) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlBar) RowSnap() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlBarAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlBar) SetRowSnap(value bool) {
	if !m.IsValid() {
		return
	}
	customControlBarAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlBar) SetOnBandDrag(fn TBandDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTBandDragEvent(fn)
	base.SetEvent(m, 15, customControlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlBar) SetOnBandInfo(fn TBandInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTBandInfoEvent(fn)
	base.SetEvent(m, 16, customControlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlBar) SetOnBandMove(fn TBandMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTBandMoveEvent(fn)
	base.SetEvent(m, 17, customControlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlBar) SetOnBandPaint(fn TBandPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTBandPaintEvent(fn)
	base.SetEvent(m, 18, customControlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlBar) SetOnCanResize(fn TCanResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCanResizeEvent(fn)
	base.SetEvent(m, 19, customControlBarAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomControlBar class constructor
func NewCustomControlBar(owner IComponent) ICustomControlBar {
	r := customControlBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomControlBar(r)
}

func TCustomControlBarClass() types.TClass {
	r := customControlBarAPI().SysCallN(20)
	return types.TClass(r)
}

var (
	customControlBarOnce   base.Once
	customControlBarImport *imports.Imports = nil
)

func customControlBarAPI() *imports.Imports {
	customControlBarOnce.Do(func() {
		customControlBarImport = api.NewDefaultImports()
		customControlBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomControlBar_Create", 0), // constructor NewCustomControlBar
			/* 1 */ imports.NewTable("TCustomControlBar_HitTest", 0), // function HitTest
			/* 2 */ imports.NewTable("TCustomControlBar_MouseToBandPos", 0), // function MouseToBandPos
			/* 3 */ imports.NewTable("TCustomControlBar_BeginUpdate", 0), // procedure BeginUpdate
			/* 4 */ imports.NewTable("TCustomControlBar_EndUpdate", 0), // procedure EndUpdate
			/* 5 */ imports.NewTable("TCustomControlBar_StickControls", 0), // procedure StickControls
			/* 6 */ imports.NewTable("TCustomControlBar_AutoDock", 0), // property AutoDock
			/* 7 */ imports.NewTable("TCustomControlBar_AutoDrag", 0), // property AutoDrag
			/* 8 */ imports.NewTable("TCustomControlBar_DrawingStyle", 0), // property DrawingStyle
			/* 9 */ imports.NewTable("TCustomControlBar_GradientDirection", 0), // property GradientDirection
			/* 10 */ imports.NewTable("TCustomControlBar_GradientStartColor", 0), // property GradientStartColor
			/* 11 */ imports.NewTable("TCustomControlBar_GradientEndColor", 0), // property GradientEndColor
			/* 12 */ imports.NewTable("TCustomControlBar_Picture", 0), // property Picture
			/* 13 */ imports.NewTable("TCustomControlBar_RowSize", 0), // property RowSize
			/* 14 */ imports.NewTable("TCustomControlBar_RowSnap", 0), // property RowSnap
			/* 15 */ imports.NewTable("TCustomControlBar_OnBandDrag", 0), // event OnBandDrag
			/* 16 */ imports.NewTable("TCustomControlBar_OnBandInfo", 0), // event OnBandInfo
			/* 17 */ imports.NewTable("TCustomControlBar_OnBandMove", 0), // event OnBandMove
			/* 18 */ imports.NewTable("TCustomControlBar_OnBandPaint", 0), // event OnBandPaint
			/* 19 */ imports.NewTable("TCustomControlBar_OnCanResize", 0), // event OnCanResize
			/* 20 */ imports.NewTable("TCustomControlBar_TClass", 0), // function TCustomControlBarClass
		}
	})
	return customControlBarImport
}
