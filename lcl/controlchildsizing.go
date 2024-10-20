//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IControlChildSizing Parent: IPersistent
type IControlChildSizing interface {
	IPersistent
	Control() IWinControl                                 // property
	LeftRightSpacing() int32                              // property
	SetLeftRightSpacing(AValue int32)                     // property
	TopBottomSpacing() int32                              // property
	SetTopBottomSpacing(AValue int32)                     // property
	HorizontalSpacing() int32                             // property
	SetHorizontalSpacing(AValue int32)                    // property
	VerticalSpacing() int32                               // property
	SetVerticalSpacing(AValue int32)                      // property
	EnlargeHorizontal() TChildControlResizeStyle          // property
	SetEnlargeHorizontal(AValue TChildControlResizeStyle) // property
	EnlargeVertical() TChildControlResizeStyle            // property
	SetEnlargeVertical(AValue TChildControlResizeStyle)   // property
	ShrinkHorizontal() TChildControlResizeStyle           // property
	SetShrinkHorizontal(AValue TChildControlResizeStyle)  // property
	ShrinkVertical() TChildControlResizeStyle             // property
	SetShrinkVertical(AValue TChildControlResizeStyle)    // property
	Layout() TControlChildrenLayout                       // property
	SetLayout(AValue TControlChildrenLayout)              // property
	ControlsPerLine() int32                               // property
	SetControlsPerLine(AValue int32)                      // property
	IsEqual(Sizing IControlChildSizing) bool              // function
	AssignTo(Dest IPersistent)                            // procedure
	SetGridSpacing(Spacing int32)                         // procedure
	SetOnChange(fn TNotifyEvent)                          // property event
}

// TControlChildSizing Parent: TPersistent
type TControlChildSizing struct {
	TPersistent
	changePtr uintptr
}

func NewControlChildSizing(OwnerControl IWinControl) IControlChildSizing {
	r1 := controlChildSizingImportAPI().SysCallN(4, GetObjectUintptr(OwnerControl))
	return AsControlChildSizing(r1)
}

func (m *TControlChildSizing) Control() IWinControl {
	r1 := controlChildSizingImportAPI().SysCallN(2, m.Instance())
	return AsWinControl(r1)
}

func (m *TControlChildSizing) LeftRightSpacing() int32 {
	r1 := controlChildSizingImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetLeftRightSpacing(AValue int32) {
	controlChildSizingImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) TopBottomSpacing() int32 {
	r1 := controlChildSizingImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetTopBottomSpacing(AValue int32) {
	controlChildSizingImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) HorizontalSpacing() int32 {
	r1 := controlChildSizingImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetHorizontalSpacing(AValue int32) {
	controlChildSizingImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) VerticalSpacing() int32 {
	r1 := controlChildSizingImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetVerticalSpacing(AValue int32) {
	controlChildSizingImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) EnlargeHorizontal() TChildControlResizeStyle {
	r1 := controlChildSizingImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetEnlargeHorizontal(AValue TChildControlResizeStyle) {
	controlChildSizingImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) EnlargeVertical() TChildControlResizeStyle {
	r1 := controlChildSizingImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetEnlargeVertical(AValue TChildControlResizeStyle) {
	controlChildSizingImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) ShrinkHorizontal() TChildControlResizeStyle {
	r1 := controlChildSizingImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetShrinkHorizontal(AValue TChildControlResizeStyle) {
	controlChildSizingImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) ShrinkVertical() TChildControlResizeStyle {
	r1 := controlChildSizingImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TChildControlResizeStyle(r1)
}

func (m *TControlChildSizing) SetShrinkVertical(AValue TChildControlResizeStyle) {
	controlChildSizingImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) Layout() TControlChildrenLayout {
	r1 := controlChildSizingImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TControlChildrenLayout(r1)
}

func (m *TControlChildSizing) SetLayout(AValue TControlChildrenLayout) {
	controlChildSizingImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) ControlsPerLine() int32 {
	r1 := controlChildSizingImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlChildSizing) SetControlsPerLine(AValue int32) {
	controlChildSizingImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlChildSizing) IsEqual(Sizing IControlChildSizing) bool {
	r1 := controlChildSizingImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(Sizing))
	return GoBool(r1)
}

func ControlChildSizingClass() TClass {
	ret := controlChildSizingImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TControlChildSizing) AssignTo(Dest IPersistent) {
	controlChildSizingImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TControlChildSizing) SetGridSpacing(Spacing int32) {
	controlChildSizingImportAPI().SysCallN(11, m.Instance(), uintptr(Spacing))
}

func (m *TControlChildSizing) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	controlChildSizingImportAPI().SysCallN(12, m.Instance(), m.changePtr)
}

var (
	controlChildSizingImport       *imports.Imports = nil
	controlChildSizingImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ControlChildSizing_AssignTo", 0),
		/*1*/ imports.NewTable("ControlChildSizing_Class", 0),
		/*2*/ imports.NewTable("ControlChildSizing_Control", 0),
		/*3*/ imports.NewTable("ControlChildSizing_ControlsPerLine", 0),
		/*4*/ imports.NewTable("ControlChildSizing_Create", 0),
		/*5*/ imports.NewTable("ControlChildSizing_EnlargeHorizontal", 0),
		/*6*/ imports.NewTable("ControlChildSizing_EnlargeVertical", 0),
		/*7*/ imports.NewTable("ControlChildSizing_HorizontalSpacing", 0),
		/*8*/ imports.NewTable("ControlChildSizing_IsEqual", 0),
		/*9*/ imports.NewTable("ControlChildSizing_Layout", 0),
		/*10*/ imports.NewTable("ControlChildSizing_LeftRightSpacing", 0),
		/*11*/ imports.NewTable("ControlChildSizing_SetGridSpacing", 0),
		/*12*/ imports.NewTable("ControlChildSizing_SetOnChange", 0),
		/*13*/ imports.NewTable("ControlChildSizing_ShrinkHorizontal", 0),
		/*14*/ imports.NewTable("ControlChildSizing_ShrinkVertical", 0),
		/*15*/ imports.NewTable("ControlChildSizing_TopBottomSpacing", 0),
		/*16*/ imports.NewTable("ControlChildSizing_VerticalSpacing", 0),
	}
)

func controlChildSizingImportAPI() *imports.Imports {
	if controlChildSizingImport == nil {
		controlChildSizingImport = NewDefaultImports()
		controlChildSizingImport.SetImportTable(controlChildSizingImportTables)
		controlChildSizingImportTables = nil
	}
	return controlChildSizingImport
}
