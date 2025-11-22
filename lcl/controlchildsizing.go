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

// IControlChildSizing Parent: IPersistent
type IControlChildSizing interface {
	IPersistent
	IsEqual(sizing IControlChildSizing) bool                   // function
	AssignTo(dest IPersistent)                                 // procedure
	SetGridSpacing(spacing int32)                              // procedure
	Control() IWinControl                                      // property Control Getter
	LeftRightSpacing() int32                                   // property LeftRightSpacing Getter
	SetLeftRightSpacing(value int32)                           // property LeftRightSpacing Setter
	TopBottomSpacing() int32                                   // property TopBottomSpacing Getter
	SetTopBottomSpacing(value int32)                           // property TopBottomSpacing Setter
	HorizontalSpacing() int32                                  // property HorizontalSpacing Getter
	SetHorizontalSpacing(value int32)                          // property HorizontalSpacing Setter
	VerticalSpacing() int32                                    // property VerticalSpacing Getter
	SetVerticalSpacing(value int32)                            // property VerticalSpacing Setter
	EnlargeHorizontal() types.TChildControlResizeStyle         // property EnlargeHorizontal Getter
	SetEnlargeHorizontal(value types.TChildControlResizeStyle) // property EnlargeHorizontal Setter
	EnlargeVertical() types.TChildControlResizeStyle           // property EnlargeVertical Getter
	SetEnlargeVertical(value types.TChildControlResizeStyle)   // property EnlargeVertical Setter
	ShrinkHorizontal() types.TChildControlResizeStyle          // property ShrinkHorizontal Getter
	SetShrinkHorizontal(value types.TChildControlResizeStyle)  // property ShrinkHorizontal Setter
	ShrinkVertical() types.TChildControlResizeStyle            // property ShrinkVertical Getter
	SetShrinkVertical(value types.TChildControlResizeStyle)    // property ShrinkVertical Setter
	Layout() types.TControlChildrenLayout                      // property Layout Getter
	SetLayout(value types.TControlChildrenLayout)              // property Layout Setter
	ControlsPerLine() int32                                    // property ControlsPerLine Getter
	SetControlsPerLine(value int32)                            // property ControlsPerLine Setter
	SetOnChange(fn TNotifyEvent)                               // property event
}

type TControlChildSizing struct {
	TPersistent
}

func (m *TControlChildSizing) IsEqual(sizing IControlChildSizing) bool {
	if !m.IsValid() {
		return false
	}
	r := controlChildSizingAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(sizing))
	return api.GoBool(r)
}

func (m *TControlChildSizing) AssignTo(dest IPersistent) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(dest))
}

func (m *TControlChildSizing) SetGridSpacing(spacing int32) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(3, m.Instance(), uintptr(spacing))
}

func (m *TControlChildSizing) Control() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := controlChildSizingAPI().SysCallN(4, m.Instance())
	return AsWinControl(r)
}

func (m *TControlChildSizing) LeftRightSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TControlChildSizing) SetLeftRightSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) TopBottomSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TControlChildSizing) SetTopBottomSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) HorizontalSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TControlChildSizing) SetHorizontalSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) VerticalSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TControlChildSizing) SetVerticalSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) EnlargeHorizontal() types.TChildControlResizeStyle {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(9, 0, m.Instance())
	return types.TChildControlResizeStyle(r)
}

func (m *TControlChildSizing) SetEnlargeHorizontal(value types.TChildControlResizeStyle) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) EnlargeVertical() types.TChildControlResizeStyle {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(10, 0, m.Instance())
	return types.TChildControlResizeStyle(r)
}

func (m *TControlChildSizing) SetEnlargeVertical(value types.TChildControlResizeStyle) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) ShrinkHorizontal() types.TChildControlResizeStyle {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(11, 0, m.Instance())
	return types.TChildControlResizeStyle(r)
}

func (m *TControlChildSizing) SetShrinkHorizontal(value types.TChildControlResizeStyle) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) ShrinkVertical() types.TChildControlResizeStyle {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(12, 0, m.Instance())
	return types.TChildControlResizeStyle(r)
}

func (m *TControlChildSizing) SetShrinkVertical(value types.TChildControlResizeStyle) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) Layout() types.TControlChildrenLayout {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(13, 0, m.Instance())
	return types.TControlChildrenLayout(r)
}

func (m *TControlChildSizing) SetLayout(value types.TControlChildrenLayout) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) ControlsPerLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlChildSizingAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TControlChildSizing) SetControlsPerLine(value int32) {
	if !m.IsValid() {
		return
	}
	controlChildSizingAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TControlChildSizing) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, controlChildSizingAPI(), api.MakeEventDataPtr(cb))
}

// NewControlChildSizing class constructor
func NewControlChildSizing(ownerControl IWinControl) IControlChildSizing {
	r := controlChildSizingAPI().SysCallN(0, base.GetObjectUintptr(ownerControl))
	return AsControlChildSizing(r)
}

var (
	controlChildSizingOnce   base.Once
	controlChildSizingImport *imports.Imports = nil
)

func controlChildSizingAPI() *imports.Imports {
	controlChildSizingOnce.Do(func() {
		controlChildSizingImport = api.NewDefaultImports()
		controlChildSizingImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TControlChildSizing_Create", 0), // constructor NewControlChildSizing
			/* 1 */ imports.NewTable("TControlChildSizing_IsEqual", 0), // function IsEqual
			/* 2 */ imports.NewTable("TControlChildSizing_AssignTo", 0), // procedure AssignTo
			/* 3 */ imports.NewTable("TControlChildSizing_SetGridSpacing", 0), // procedure SetGridSpacing
			/* 4 */ imports.NewTable("TControlChildSizing_Control", 0), // property Control
			/* 5 */ imports.NewTable("TControlChildSizing_LeftRightSpacing", 0), // property LeftRightSpacing
			/* 6 */ imports.NewTable("TControlChildSizing_TopBottomSpacing", 0), // property TopBottomSpacing
			/* 7 */ imports.NewTable("TControlChildSizing_HorizontalSpacing", 0), // property HorizontalSpacing
			/* 8 */ imports.NewTable("TControlChildSizing_VerticalSpacing", 0), // property VerticalSpacing
			/* 9 */ imports.NewTable("TControlChildSizing_EnlargeHorizontal", 0), // property EnlargeHorizontal
			/* 10 */ imports.NewTable("TControlChildSizing_EnlargeVertical", 0), // property EnlargeVertical
			/* 11 */ imports.NewTable("TControlChildSizing_ShrinkHorizontal", 0), // property ShrinkHorizontal
			/* 12 */ imports.NewTable("TControlChildSizing_ShrinkVertical", 0), // property ShrinkVertical
			/* 13 */ imports.NewTable("TControlChildSizing_Layout", 0), // property Layout
			/* 14 */ imports.NewTable("TControlChildSizing_ControlsPerLine", 0), // property ControlsPerLine
			/* 15 */ imports.NewTable("TControlChildSizing_OnChange", 0), // event OnChange
		}
	})
	return controlChildSizingImport
}
