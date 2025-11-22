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

// IVTColors Parent: IPersistent
type IVTColors interface {
	IPersistent
	BackGroundColor() types.TColor                       // property BackGroundColor Getter
	HeaderFontColor() types.TColor                       // property HeaderFontColor Getter
	NodeFontColor() types.TColor                         // property NodeFontColor Getter
	BorderColor() types.TColor                           // property BorderColor Getter
	SetBorderColor(value types.TColor)                   // property BorderColor Setter
	DisabledColor() types.TColor                         // property DisabledColor Getter
	SetDisabledColor(value types.TColor)                 // property DisabledColor Setter
	DropMarkColor() types.TColor                         // property DropMarkColor Getter
	SetDropMarkColor(value types.TColor)                 // property DropMarkColor Setter
	DropTargetColor() types.TColor                       // property DropTargetColor Getter
	SetDropTargetColor(value types.TColor)               // property DropTargetColor Setter
	DropTargetBorderColor() types.TColor                 // property DropTargetBorderColor Getter
	SetDropTargetBorderColor(value types.TColor)         // property DropTargetBorderColor Setter
	FocusedSelectionColor() types.TColor                 // property FocusedSelectionColor Getter
	SetFocusedSelectionColor(value types.TColor)         // property FocusedSelectionColor Setter
	FocusedSelectionBorderColor() types.TColor           // property FocusedSelectionBorderColor Getter
	SetFocusedSelectionBorderColor(value types.TColor)   // property FocusedSelectionBorderColor Setter
	GridLineColor() types.TColor                         // property GridLineColor Getter
	SetGridLineColor(value types.TColor)                 // property GridLineColor Setter
	HeaderHotColor() types.TColor                        // property HeaderHotColor Getter
	SetHeaderHotColor(value types.TColor)                // property HeaderHotColor Setter
	HotColor() types.TColor                              // property HotColor Getter
	SetHotColor(value types.TColor)                      // property HotColor Setter
	SelectionRectangleBlendColor() types.TColor          // property SelectionRectangleBlendColor Getter
	SetSelectionRectangleBlendColor(value types.TColor)  // property SelectionRectangleBlendColor Setter
	SelectionRectangleBorderColor() types.TColor         // property SelectionRectangleBorderColor Getter
	SetSelectionRectangleBorderColor(value types.TColor) // property SelectionRectangleBorderColor Setter
	SelectionTextColor() types.TColor                    // property SelectionTextColor Getter
	SetSelectionTextColor(value types.TColor)            // property SelectionTextColor Setter
	TreeLineColor() types.TColor                         // property TreeLineColor Getter
	SetTreeLineColor(value types.TColor)                 // property TreeLineColor Setter
	UnfocusedColor() types.TColor                        // property UnfocusedColor Getter
	SetUnfocusedColor(value types.TColor)                // property UnfocusedColor Setter
	UnfocusedSelectionColor() types.TColor               // property UnfocusedSelectionColor Getter
	SetUnfocusedSelectionColor(value types.TColor)       // property UnfocusedSelectionColor Setter
	UnfocusedSelectionBorderColor() types.TColor         // property UnfocusedSelectionBorderColor Getter
	SetUnfocusedSelectionBorderColor(value types.TColor) // property UnfocusedSelectionBorderColor Setter
}

type TVTColors struct {
	TPersistent
}

func (m *TVTColors) BackGroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(1, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) HeaderFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(2, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) NodeFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(3, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) BorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(4, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) DisabledColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(5, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetDisabledColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) DropMarkColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(6, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetDropMarkColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) DropTargetColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(7, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetDropTargetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) DropTargetBorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetDropTargetBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) FocusedSelectionColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(9, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetFocusedSelectionColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) FocusedSelectionBorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(10, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetFocusedSelectionBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) GridLineColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(11, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetGridLineColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) HeaderHotColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(12, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetHeaderHotColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) HotColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(13, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetHotColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) SelectionRectangleBlendColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(14, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetSelectionRectangleBlendColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) SelectionRectangleBorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(15, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetSelectionRectangleBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) SelectionTextColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(16, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetSelectionTextColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) TreeLineColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(17, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetTreeLineColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) UnfocusedColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(18, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetUnfocusedColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) UnfocusedSelectionColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(19, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetUnfocusedSelectionColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TVTColors) UnfocusedSelectionBorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTColorsAPI().SysCallN(20, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTColors) SetUnfocusedSelectionBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTColorsAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

// NewVTColors class constructor
func NewVTColors(owner IBaseVirtualTree) IVTColors {
	r := vTColorsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsVTColors(r)
}

var (
	vTColorsOnce   base.Once
	vTColorsImport *imports.Imports = nil
)

func vTColorsAPI() *imports.Imports {
	vTColorsOnce.Do(func() {
		vTColorsImport = api.NewDefaultImports()
		vTColorsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTColors_Create", 0), // constructor NewVTColors
			/* 1 */ imports.NewTable("TVTColors_BackGroundColor", 0), // property BackGroundColor
			/* 2 */ imports.NewTable("TVTColors_HeaderFontColor", 0), // property HeaderFontColor
			/* 3 */ imports.NewTable("TVTColors_NodeFontColor", 0), // property NodeFontColor
			/* 4 */ imports.NewTable("TVTColors_BorderColor", 0), // property BorderColor
			/* 5 */ imports.NewTable("TVTColors_DisabledColor", 0), // property DisabledColor
			/* 6 */ imports.NewTable("TVTColors_DropMarkColor", 0), // property DropMarkColor
			/* 7 */ imports.NewTable("TVTColors_DropTargetColor", 0), // property DropTargetColor
			/* 8 */ imports.NewTable("TVTColors_DropTargetBorderColor", 0), // property DropTargetBorderColor
			/* 9 */ imports.NewTable("TVTColors_FocusedSelectionColor", 0), // property FocusedSelectionColor
			/* 10 */ imports.NewTable("TVTColors_FocusedSelectionBorderColor", 0), // property FocusedSelectionBorderColor
			/* 11 */ imports.NewTable("TVTColors_GridLineColor", 0), // property GridLineColor
			/* 12 */ imports.NewTable("TVTColors_HeaderHotColor", 0), // property HeaderHotColor
			/* 13 */ imports.NewTable("TVTColors_HotColor", 0), // property HotColor
			/* 14 */ imports.NewTable("TVTColors_SelectionRectangleBlendColor", 0), // property SelectionRectangleBlendColor
			/* 15 */ imports.NewTable("TVTColors_SelectionRectangleBorderColor", 0), // property SelectionRectangleBorderColor
			/* 16 */ imports.NewTable("TVTColors_SelectionTextColor", 0), // property SelectionTextColor
			/* 17 */ imports.NewTable("TVTColors_TreeLineColor", 0), // property TreeLineColor
			/* 18 */ imports.NewTable("TVTColors_UnfocusedColor", 0), // property UnfocusedColor
			/* 19 */ imports.NewTable("TVTColors_UnfocusedSelectionColor", 0), // property UnfocusedSelectionColor
			/* 20 */ imports.NewTable("TVTColors_UnfocusedSelectionBorderColor", 0), // property UnfocusedSelectionBorderColor
		}
	})
	return vTColorsImport
}
