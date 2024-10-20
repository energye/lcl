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
	r1 := vTColorsImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsVTColors(r1)
}

func (m *TVTColors) BackGroundColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(0, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) HeaderFontColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(11, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) NodeFontColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(14, m.Instance())
	return TColor(r1)
}

func (m *TVTColors) BorderColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetBorderColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DisabledColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDisabledColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropMarkColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropMarkColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropTargetColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropTargetColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) DropTargetBorderColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetDropTargetBorderColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) FocusedSelectionColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetFocusedSelectionColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) FocusedSelectionBorderColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetFocusedSelectionBorderColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) GridLineColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetGridLineColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) HeaderHotColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetHeaderHotColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) HotColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetHotColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionRectangleBlendColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionRectangleBlendColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionRectangleBorderColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionRectangleBorderColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) SelectionTextColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetSelectionTextColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) TreeLineColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetTreeLineColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedSelectionColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedSelectionColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTColors) UnfocusedSelectionBorderColor() TColor {
	r1 := vTColorsImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTColors) SetUnfocusedSelectionBorderColor(AValue TColor) {
	vTColorsImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func VTColorsClass() TClass {
	ret := vTColorsImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	vTColorsImport       *imports.Imports = nil
	vTColorsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VTColors_BackGroundColor", 0),
		/*1*/ imports.NewTable("VTColors_BorderColor", 0),
		/*2*/ imports.NewTable("VTColors_Class", 0),
		/*3*/ imports.NewTable("VTColors_Create", 0),
		/*4*/ imports.NewTable("VTColors_DisabledColor", 0),
		/*5*/ imports.NewTable("VTColors_DropMarkColor", 0),
		/*6*/ imports.NewTable("VTColors_DropTargetBorderColor", 0),
		/*7*/ imports.NewTable("VTColors_DropTargetColor", 0),
		/*8*/ imports.NewTable("VTColors_FocusedSelectionBorderColor", 0),
		/*9*/ imports.NewTable("VTColors_FocusedSelectionColor", 0),
		/*10*/ imports.NewTable("VTColors_GridLineColor", 0),
		/*11*/ imports.NewTable("VTColors_HeaderFontColor", 0),
		/*12*/ imports.NewTable("VTColors_HeaderHotColor", 0),
		/*13*/ imports.NewTable("VTColors_HotColor", 0),
		/*14*/ imports.NewTable("VTColors_NodeFontColor", 0),
		/*15*/ imports.NewTable("VTColors_SelectionRectangleBlendColor", 0),
		/*16*/ imports.NewTable("VTColors_SelectionRectangleBorderColor", 0),
		/*17*/ imports.NewTable("VTColors_SelectionTextColor", 0),
		/*18*/ imports.NewTable("VTColors_TreeLineColor", 0),
		/*19*/ imports.NewTable("VTColors_UnfocusedColor", 0),
		/*20*/ imports.NewTable("VTColors_UnfocusedSelectionBorderColor", 0),
		/*21*/ imports.NewTable("VTColors_UnfocusedSelectionColor", 0),
	}
)

func vTColorsImportAPI() *imports.Imports {
	if vTColorsImport == nil {
		vTColorsImport = NewDefaultImports()
		vTColorsImport.SetImportTable(vTColorsImportTables)
		vTColorsImportTables = nil
	}
	return vTColorsImport
}
