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

// IControlBorderSpacing Parent: IPersistent
type IControlBorderSpacing interface {
	IPersistent
	IsEqual(spacing IControlBorderSpacing) bool                // function
	GetSideSpace(kind types.TAnchorKind) int32                 // function
	GetSpace(kind types.TAnchorKind) int32                     // function
	AssignTo(dest IPersistent)                                 // procedure
	GetSpaceAround(spaceAround *types.TRect)                   // procedure
	AutoAdjustLayout(xProportion float64, yProportion float64) // procedure
	Control() IControl                                         // property Control Getter
	Space(kind types.TAnchorKind) int32                        // property Space Getter
	SetSpace(kind types.TAnchorKind, value int32)              // property Space Setter
	AroundLeft() int32                                         // property AroundLeft Getter
	AroundTop() int32                                          // property AroundTop Getter
	AroundRight() int32                                        // property AroundRight Getter
	AroundBottom() int32                                       // property AroundBottom Getter
	ControlLeft() int32                                        // property ControlLeft Getter
	ControlTop() int32                                         // property ControlTop Getter
	ControlWidth() int32                                       // property ControlWidth Getter
	ControlHeight() int32                                      // property ControlHeight Getter
	ControlRight() int32                                       // property ControlRight Getter
	ControlBottom() int32                                      // property ControlBottom Getter
	Left() types.TSpacingSize                                  // property Left Getter
	SetLeft(value types.TSpacingSize)                          // property Left Setter
	Top() types.TSpacingSize                                   // property Top Getter
	SetTop(value types.TSpacingSize)                           // property Top Setter
	Right() types.TSpacingSize                                 // property Right Getter
	SetRight(value types.TSpacingSize)                         // property Right Setter
	Bottom() types.TSpacingSize                                // property Bottom Getter
	SetBottom(value types.TSpacingSize)                        // property Bottom Setter
	Around() types.TSpacingSize                                // property Around Getter
	SetAround(value types.TSpacingSize)                        // property Around Setter
	InnerBorder() int32                                        // property InnerBorder Getter
	SetInnerBorder(value int32)                                // property InnerBorder Setter
	CellAlignHorizontal() types.TControlCellAlign              // property CellAlignHorizontal Getter
	SetCellAlignHorizontal(value types.TControlCellAlign)      // property CellAlignHorizontal Setter
	CellAlignVertical() types.TControlCellAlign                // property CellAlignVertical Getter
	SetCellAlignVertical(value types.TControlCellAlign)        // property CellAlignVertical Setter
	SetOnChange(fn TNotifyEvent)                               // property event
}

type TControlBorderSpacing struct {
	TPersistent
}

func (m *TControlBorderSpacing) IsEqual(spacing IControlBorderSpacing) bool {
	if !m.IsValid() {
		return false
	}
	r := controlBorderSpacingAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(spacing))
	return api.GoBool(r)
}

func (m *TControlBorderSpacing) GetSideSpace(kind types.TAnchorKind) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(2, m.Instance(), uintptr(kind))
	return int32(r)
}

func (m *TControlBorderSpacing) GetSpace(kind types.TAnchorKind) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(3, m.Instance(), uintptr(kind))
	return int32(r)
}

func (m *TControlBorderSpacing) AssignTo(dest IPersistent) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(dest))
}

func (m *TControlBorderSpacing) GetSpaceAround(spaceAround *types.TRect) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(spaceAround)))
}

func (m *TControlBorderSpacing) AutoAdjustLayout(xProportion float64, yProportion float64) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&xProportion)), uintptr(base.UnsafePointer(&yProportion)))
}

func (m *TControlBorderSpacing) Control() IControl {
	if !m.IsValid() {
		return nil
	}
	r := controlBorderSpacingAPI().SysCallN(7, m.Instance())
	return AsControl(r)
}

func (m *TControlBorderSpacing) Space(kind types.TAnchorKind) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(8, 0, m.Instance(), uintptr(kind))
	return int32(r)
}

func (m *TControlBorderSpacing) SetSpace(kind types.TAnchorKind, value int32) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(8, 1, m.Instance(), uintptr(kind), uintptr(value))
}

func (m *TControlBorderSpacing) AroundLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) AroundTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) AroundRight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) AroundBottom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(12, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) ControlLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(13, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) ControlTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(14, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) ControlWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(15, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) ControlHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(16, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) ControlRight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(17, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) ControlBottom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(18, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) Left() types.TSpacingSize {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(19, 0, m.Instance())
	return types.TSpacingSize(r)
}

func (m *TControlBorderSpacing) SetLeft(value types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) Top() types.TSpacingSize {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(20, 0, m.Instance())
	return types.TSpacingSize(r)
}

func (m *TControlBorderSpacing) SetTop(value types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) Right() types.TSpacingSize {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(21, 0, m.Instance())
	return types.TSpacingSize(r)
}

func (m *TControlBorderSpacing) SetRight(value types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) Bottom() types.TSpacingSize {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(22, 0, m.Instance())
	return types.TSpacingSize(r)
}

func (m *TControlBorderSpacing) SetBottom(value types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) Around() types.TSpacingSize {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(23, 0, m.Instance())
	return types.TSpacingSize(r)
}

func (m *TControlBorderSpacing) SetAround(value types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) InnerBorder() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TControlBorderSpacing) SetInnerBorder(value int32) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) CellAlignHorizontal() types.TControlCellAlign {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(25, 0, m.Instance())
	return types.TControlCellAlign(r)
}

func (m *TControlBorderSpacing) SetCellAlignHorizontal(value types.TControlCellAlign) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) CellAlignVertical() types.TControlCellAlign {
	if !m.IsValid() {
		return 0
	}
	r := controlBorderSpacingAPI().SysCallN(26, 0, m.Instance())
	return types.TControlCellAlign(r)
}

func (m *TControlBorderSpacing) SetCellAlignVertical(value types.TControlCellAlign) {
	if !m.IsValid() {
		return
	}
	controlBorderSpacingAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TControlBorderSpacing) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 27, controlBorderSpacingAPI(), api.MakeEventDataPtr(cb))
}

// NewControlBorderSpacing class constructor
func NewControlBorderSpacing(ownerControl IControl, default_ TControlBorderSpacingDefault) IControlBorderSpacing {
	r := controlBorderSpacingAPI().SysCallN(0, base.GetObjectUintptr(ownerControl), uintptr(base.UnsafePointer(&default_)))
	return AsControlBorderSpacing(r)
}

var (
	controlBorderSpacingOnce   base.Once
	controlBorderSpacingImport *imports.Imports = nil
)

func controlBorderSpacingAPI() *imports.Imports {
	controlBorderSpacingOnce.Do(func() {
		controlBorderSpacingImport = api.NewDefaultImports()
		controlBorderSpacingImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TControlBorderSpacing_Create", 0), // constructor NewControlBorderSpacing
			/* 1 */ imports.NewTable("TControlBorderSpacing_IsEqual", 0), // function IsEqual
			/* 2 */ imports.NewTable("TControlBorderSpacing_GetSideSpace", 0), // function GetSideSpace
			/* 3 */ imports.NewTable("TControlBorderSpacing_GetSpace", 0), // function GetSpace
			/* 4 */ imports.NewTable("TControlBorderSpacing_AssignTo", 0), // procedure AssignTo
			/* 5 */ imports.NewTable("TControlBorderSpacing_GetSpaceAround", 0), // procedure GetSpaceAround
			/* 6 */ imports.NewTable("TControlBorderSpacing_AutoAdjustLayout", 0), // procedure AutoAdjustLayout
			/* 7 */ imports.NewTable("TControlBorderSpacing_Control", 0), // property Control
			/* 8 */ imports.NewTable("TControlBorderSpacing_Space", 0), // property Space
			/* 9 */ imports.NewTable("TControlBorderSpacing_AroundLeft", 0), // property AroundLeft
			/* 10 */ imports.NewTable("TControlBorderSpacing_AroundTop", 0), // property AroundTop
			/* 11 */ imports.NewTable("TControlBorderSpacing_AroundRight", 0), // property AroundRight
			/* 12 */ imports.NewTable("TControlBorderSpacing_AroundBottom", 0), // property AroundBottom
			/* 13 */ imports.NewTable("TControlBorderSpacing_ControlLeft", 0), // property ControlLeft
			/* 14 */ imports.NewTable("TControlBorderSpacing_ControlTop", 0), // property ControlTop
			/* 15 */ imports.NewTable("TControlBorderSpacing_ControlWidth", 0), // property ControlWidth
			/* 16 */ imports.NewTable("TControlBorderSpacing_ControlHeight", 0), // property ControlHeight
			/* 17 */ imports.NewTable("TControlBorderSpacing_ControlRight", 0), // property ControlRight
			/* 18 */ imports.NewTable("TControlBorderSpacing_ControlBottom", 0), // property ControlBottom
			/* 19 */ imports.NewTable("TControlBorderSpacing_Left", 0), // property Left
			/* 20 */ imports.NewTable("TControlBorderSpacing_Top", 0), // property Top
			/* 21 */ imports.NewTable("TControlBorderSpacing_Right", 0), // property Right
			/* 22 */ imports.NewTable("TControlBorderSpacing_Bottom", 0), // property Bottom
			/* 23 */ imports.NewTable("TControlBorderSpacing_Around", 0), // property Around
			/* 24 */ imports.NewTable("TControlBorderSpacing_InnerBorder", 0), // property InnerBorder
			/* 25 */ imports.NewTable("TControlBorderSpacing_CellAlignHorizontal", 0), // property CellAlignHorizontal
			/* 26 */ imports.NewTable("TControlBorderSpacing_CellAlignVertical", 0), // property CellAlignVertical
			/* 27 */ imports.NewTable("TControlBorderSpacing_OnChange", 0), // event OnChange
		}
	})
	return controlBorderSpacingImport
}
