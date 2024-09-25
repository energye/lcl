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
	. "github.com/energye/lcl/types"
)

// IVTColors Parent: IPersistent
// class to collect all switchable colors into one place
type IVTColors interface {
	IPersistent
	BackGroundColor() TColor                        // property
	HeaderFontColor() TColor                        // property
	NodeFontColor() TColor                          // property
	BorderColor() TColor                            // property
	SetBorderColor(AValue TColor)                   // property
	DisabledColor() TColor                          // property
	SetDisabledColor(AValue TColor)                 // property
	DropMarkColor() TColor                          // property
	SetDropMarkColor(AValue TColor)                 // property
	DropTargetColor() TColor                        // property
	SetDropTargetColor(AValue TColor)               // property
	DropTargetBorderColor() TColor                  // property
	SetDropTargetBorderColor(AValue TColor)         // property
	FocusedSelectionColor() TColor                  // property
	SetFocusedSelectionColor(AValue TColor)         // property
	FocusedSelectionBorderColor() TColor            // property
	SetFocusedSelectionBorderColor(AValue TColor)   // property
	GridLineColor() TColor                          // property
	SetGridLineColor(AValue TColor)                 // property
	HeaderHotColor() TColor                         // property
	SetHeaderHotColor(AValue TColor)                // property
	HotColor() TColor                               // property
	SetHotColor(AValue TColor)                      // property
	SelectionRectangleBlendColor() TColor           // property
	SetSelectionRectangleBlendColor(AValue TColor)  // property
	SelectionRectangleBorderColor() TColor          // property
	SetSelectionRectangleBorderColor(AValue TColor) // property
	SelectionTextColor() TColor                     // property
	SetSelectionTextColor(AValue TColor)            // property
	TreeLineColor() TColor                          // property
	SetTreeLineColor(AValue TColor)                 // property
	UnfocusedColor() TColor                         // property
	SetUnfocusedColor(AValue TColor)                // property
	UnfocusedSelectionColor() TColor                // property
	SetUnfocusedSelectionColor(AValue TColor)       // property
	UnfocusedSelectionBorderColor() TColor          // property
	SetUnfocusedSelectionBorderColor(AValue TColor) // property
}

// TVTColors Parent: TPersistent
// class to collect all switchable colors into one place
type TVTColors struct {
	TPersistent
}

func NewVTColors(AOwner IBaseVirtualTree) IVTColors {
	r1 := LCL().SysCallN(5857, GetObjectUintptr(AOwner))
	return AsVTColors(r1)
}

func (m *TVTColors) BackGroundColor() TColor {
	r1 := LCL().SysCallN(5854, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) HeaderFontColor() TColor {
	r1 := LCL().SysCallN(5865, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) NodeFontColor() TColor {
	r1 := LCL().SysCallN(5868, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) BorderColor() TColor {
	r1 := LCL().SysCallN(5855, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetBorderColor(AValue TColor) {
	LCL().SysCallN(5855, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DisabledColor() TColor {
	r1 := LCL().SysCallN(5858, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDisabledColor(AValue TColor) {
	LCL().SysCallN(5858, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropMarkColor() TColor {
	r1 := LCL().SysCallN(5859, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropMarkColor(AValue TColor) {
	LCL().SysCallN(5859, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropTargetColor() TColor {
	r1 := LCL().SysCallN(5861, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropTargetColor(AValue TColor) {
	LCL().SysCallN(5861, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropTargetBorderColor() TColor {
	r1 := LCL().SysCallN(5860, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropTargetBorderColor(AValue TColor) {
	LCL().SysCallN(5860, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) FocusedSelectionColor() TColor {
	r1 := LCL().SysCallN(5863, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetFocusedSelectionColor(AValue TColor) {
	LCL().SysCallN(5863, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) FocusedSelectionBorderColor() TColor {
	r1 := LCL().SysCallN(5862, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetFocusedSelectionBorderColor(AValue TColor) {
	LCL().SysCallN(5862, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) GridLineColor() TColor {
	r1 := LCL().SysCallN(5864, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetGridLineColor(AValue TColor) {
	LCL().SysCallN(5864, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) HeaderHotColor() TColor {
	r1 := LCL().SysCallN(5866, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetHeaderHotColor(AValue TColor) {
	LCL().SysCallN(5866, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) HotColor() TColor {
	r1 := LCL().SysCallN(5867, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetHotColor(AValue TColor) {
	LCL().SysCallN(5867, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionRectangleBlendColor() TColor {
	r1 := LCL().SysCallN(5869, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionRectangleBlendColor(AValue TColor) {
	LCL().SysCallN(5869, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionRectangleBorderColor() TColor {
	r1 := LCL().SysCallN(5870, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionRectangleBorderColor(AValue TColor) {
	LCL().SysCallN(5870, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionTextColor() TColor {
	r1 := LCL().SysCallN(5871, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionTextColor(AValue TColor) {
	LCL().SysCallN(5871, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) TreeLineColor() TColor {
	r1 := LCL().SysCallN(5872, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetTreeLineColor(AValue TColor) {
	LCL().SysCallN(5872, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedColor() TColor {
	r1 := LCL().SysCallN(5873, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedColor(AValue TColor) {
	LCL().SysCallN(5873, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedSelectionColor() TColor {
	r1 := LCL().SysCallN(5875, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedSelectionColor(AValue TColor) {
	LCL().SysCallN(5875, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedSelectionBorderColor() TColor {
	r1 := LCL().SysCallN(5874, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedSelectionBorderColor(AValue TColor) {
	LCL().SysCallN(5874, 1, m.Instance(), uintptr(AValue))
}

func VTColorsClass() TClass {
	ret := LCL().SysCallN(5856)
	return TClass(ret)
}
