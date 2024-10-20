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

// IControlBorderSpacing Parent: IPersistent
type IControlBorderSpacing interface {
	IPersistent
	Control() IControl                                   // property
	Space(Kind TAnchorKind) int32                        // property
	SetSpace(Kind TAnchorKind, AValue int32)             // property
	AroundLeft() int32                                   // property
	AroundTop() int32                                    // property
	AroundRight() int32                                  // property
	AroundBottom() int32                                 // property
	ControlLeft() int32                                  // property
	ControlTop() int32                                   // property
	ControlWidth() int32                                 // property
	ControlHeight() int32                                // property
	ControlRight() int32                                 // property
	ControlBottom() int32                                // property
	Left() TSpacingSize                                  // property
	SetLeft(AValue TSpacingSize)                         // property
	Top() TSpacingSize                                   // property
	SetTop(AValue TSpacingSize)                          // property
	Right() TSpacingSize                                 // property
	SetRight(AValue TSpacingSize)                        // property
	Bottom() TSpacingSize                                // property
	SetBottom(AValue TSpacingSize)                       // property
	Around() TSpacingSize                                // property
	SetAround(AValue TSpacingSize)                       // property
	InnerBorder() int32                                  // property
	SetInnerBorder(AValue int32)                         // property
	CellAlignHorizontal() TControlCellAlign              // property
	SetCellAlignHorizontal(AValue TControlCellAlign)     // property
	CellAlignVertical() TControlCellAlign                // property
	SetCellAlignVertical(AValue TControlCellAlign)       // property
	IsEqual(Spacing IControlBorderSpacing) bool          // function
	GetSideSpace(Kind TAnchorKind) int32                 // function
	GetSpace1(Kind TAnchorKind) int32                    // function
	AssignTo(Dest IPersistent)                           // procedure
	GetSpaceAround(SpaceAround *TRect)                   // procedure
	AutoAdjustLayout(AXProportion, AYProportion float64) // procedure
	SetOnChange(fn TNotifyEvent)                         // property event
}

// TControlBorderSpacing Parent: TPersistent
type TControlBorderSpacing struct {
	TPersistent
	changePtr uintptr
}

func NewControlBorderSpacing(OwnerControl IControl, ADefault *TControlBorderSpacingDefault) IControlBorderSpacing {
	r1 := controlBorderSpacingImportAPI().SysCallN(18, GetObjectUintptr(OwnerControl), uintptr(unsafePointer(ADefault)))
	return AsControlBorderSpacing(r1)
}

func (m *TControlBorderSpacing) Control() IControl {
	r1 := controlBorderSpacingImportAPI().SysCallN(11, m.Instance())
	return AsControl(r1)
}

func (m *TControlBorderSpacing) Space(Kind TAnchorKind) int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(27, 0, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func (m *TControlBorderSpacing) SetSpace(Kind TAnchorKind, AValue int32) {
	controlBorderSpacingImportAPI().SysCallN(27, 1, m.Instance(), uintptr(Kind), uintptr(AValue))
}

func (m *TControlBorderSpacing) AroundLeft() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundTop() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundRight() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) AroundBottom() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlLeft() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlTop() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(16, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlWidth() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(17, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlHeight() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(13, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlRight() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(15, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) ControlBottom() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(12, m.Instance())
	return int32(r1)
}

func (m *TControlBorderSpacing) Left() TSpacingSize {
	r1 := controlBorderSpacingImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetLeft(AValue TSpacingSize) {
	controlBorderSpacingImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Top() TSpacingSize {
	r1 := controlBorderSpacingImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetTop(AValue TSpacingSize) {
	controlBorderSpacingImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Right() TSpacingSize {
	r1 := controlBorderSpacingImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetRight(AValue TSpacingSize) {
	controlBorderSpacingImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Bottom() TSpacingSize {
	r1 := controlBorderSpacingImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetBottom(AValue TSpacingSize) {
	controlBorderSpacingImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) Around() TSpacingSize {
	r1 := controlBorderSpacingImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TSpacingSize(r1)
}

func (m *TControlBorderSpacing) SetAround(AValue TSpacingSize) {
	controlBorderSpacingImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) InnerBorder() int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlBorderSpacing) SetInnerBorder(AValue int32) {
	controlBorderSpacingImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) CellAlignHorizontal() TControlCellAlign {
	r1 := controlBorderSpacingImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TControlCellAlign(r1)
}

func (m *TControlBorderSpacing) SetCellAlignHorizontal(AValue TControlCellAlign) {
	controlBorderSpacingImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) CellAlignVertical() TControlCellAlign {
	r1 := controlBorderSpacingImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TControlCellAlign(r1)
}

func (m *TControlBorderSpacing) SetCellAlignVertical(AValue TControlCellAlign) {
	controlBorderSpacingImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlBorderSpacing) IsEqual(Spacing IControlBorderSpacing) bool {
	r1 := controlBorderSpacingImportAPI().SysCallN(23, m.Instance(), GetObjectUintptr(Spacing))
	return GoBool(r1)
}

func (m *TControlBorderSpacing) GetSideSpace(Kind TAnchorKind) int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(19, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func (m *TControlBorderSpacing) GetSpace1(Kind TAnchorKind) int32 {
	r1 := controlBorderSpacingImportAPI().SysCallN(20, m.Instance(), uintptr(Kind))
	return int32(r1)
}

func ControlBorderSpacingClass() TClass {
	ret := controlBorderSpacingImportAPI().SysCallN(10)
	return TClass(ret)
}

func (m *TControlBorderSpacing) AssignTo(Dest IPersistent) {
	controlBorderSpacingImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TControlBorderSpacing) GetSpaceAround(SpaceAround *TRect) {
	var result0 uintptr
	controlBorderSpacingImportAPI().SysCallN(21, m.Instance(), uintptr(unsafePointer(&result0)))
	*SpaceAround = *(*TRect)(getPointer(result0))
}

func (m *TControlBorderSpacing) AutoAdjustLayout(AXProportion, AYProportion float64) {
	controlBorderSpacingImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&AXProportion)), uintptr(unsafePointer(&AYProportion)))
}

func (m *TControlBorderSpacing) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	controlBorderSpacingImportAPI().SysCallN(26, m.Instance(), m.changePtr)
}

var (
	controlBorderSpacingImport       *imports.Imports = nil
	controlBorderSpacingImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ControlBorderSpacing_Around", 0),
		/*1*/ imports.NewTable("ControlBorderSpacing_AroundBottom", 0),
		/*2*/ imports.NewTable("ControlBorderSpacing_AroundLeft", 0),
		/*3*/ imports.NewTable("ControlBorderSpacing_AroundRight", 0),
		/*4*/ imports.NewTable("ControlBorderSpacing_AroundTop", 0),
		/*5*/ imports.NewTable("ControlBorderSpacing_AssignTo", 0),
		/*6*/ imports.NewTable("ControlBorderSpacing_AutoAdjustLayout", 0),
		/*7*/ imports.NewTable("ControlBorderSpacing_Bottom", 0),
		/*8*/ imports.NewTable("ControlBorderSpacing_CellAlignHorizontal", 0),
		/*9*/ imports.NewTable("ControlBorderSpacing_CellAlignVertical", 0),
		/*10*/ imports.NewTable("ControlBorderSpacing_Class", 0),
		/*11*/ imports.NewTable("ControlBorderSpacing_Control", 0),
		/*12*/ imports.NewTable("ControlBorderSpacing_ControlBottom", 0),
		/*13*/ imports.NewTable("ControlBorderSpacing_ControlHeight", 0),
		/*14*/ imports.NewTable("ControlBorderSpacing_ControlLeft", 0),
		/*15*/ imports.NewTable("ControlBorderSpacing_ControlRight", 0),
		/*16*/ imports.NewTable("ControlBorderSpacing_ControlTop", 0),
		/*17*/ imports.NewTable("ControlBorderSpacing_ControlWidth", 0),
		/*18*/ imports.NewTable("ControlBorderSpacing_Create", 0),
		/*19*/ imports.NewTable("ControlBorderSpacing_GetSideSpace", 0),
		/*20*/ imports.NewTable("ControlBorderSpacing_GetSpace1", 0),
		/*21*/ imports.NewTable("ControlBorderSpacing_GetSpaceAround", 0),
		/*22*/ imports.NewTable("ControlBorderSpacing_InnerBorder", 0),
		/*23*/ imports.NewTable("ControlBorderSpacing_IsEqual", 0),
		/*24*/ imports.NewTable("ControlBorderSpacing_Left", 0),
		/*25*/ imports.NewTable("ControlBorderSpacing_Right", 0),
		/*26*/ imports.NewTable("ControlBorderSpacing_SetOnChange", 0),
		/*27*/ imports.NewTable("ControlBorderSpacing_Space", 0),
		/*28*/ imports.NewTable("ControlBorderSpacing_Top", 0),
	}
)

func controlBorderSpacingImportAPI() *imports.Imports {
	if controlBorderSpacingImport == nil {
		controlBorderSpacingImport = NewDefaultImports()
		controlBorderSpacingImport.SetImportTable(controlBorderSpacingImportTables)
		controlBorderSpacingImportTables = nil
	}
	return controlBorderSpacingImport
}
