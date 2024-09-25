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

// IGridColumnTitle Parent: IPersistent
type IGridColumnTitle interface {
	IPersistent
	Column() IGridColumn                             // property
	Alignment() TAlignment                           // property
	SetAlignment(AValue TAlignment)                  // property
	Caption() string                                 // property
	SetCaption(AValue string)                        // property
	Color() TColor                                   // property
	SetColor(AValue TColor)                          // property
	Font() IFont                                     // property
	SetFont(AValue IFont)                            // property
	ImageIndex() TImageIndex                         // property
	SetImageIndex(AValue TImageIndex)                // property
	ImageLayout() TButtonLayout                      // property
	SetImageLayout(AValue TButtonLayout)             // property
	Layout() TTextLayout                             // property
	SetLayout(AValue TTextLayout)                    // property
	MultiLine() bool                                 // property
	SetMultiLine(AValue bool)                        // property
	PrefixOption() TPrefixOption                     // property
	SetPrefixOption(AValue TPrefixOption)            // property
	IsDefault() bool                                 // function
	FillTitleDefaultFont()                           // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)          // procedure
	ScaleFontsPPI(AToPPI int32, AProportion float64) // procedure
}

// TGridColumnTitle Parent: TPersistent
type TGridColumnTitle struct {
	TPersistent
}

func NewGridColumnTitle(TheColumn IGridColumn) IGridColumnTitle {
	r1 := LCL().SysCallN(3240, GetObjectUintptr(TheColumn))
	return AsGridColumnTitle(r1)
}

func (m *TGridColumnTitle) Column() IGridColumn {
	r1 := LCL().SysCallN(3239, m.Instance())
	return AsGridColumn(r1)
}

func (m *TGridColumnTitle) Alignment() TAlignment {
	r1 := LCL().SysCallN(3235, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TGridColumnTitle) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3235, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Caption() string {
	r1 := LCL().SysCallN(3236, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumnTitle) SetCaption(AValue string) {
	LCL().SysCallN(3236, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumnTitle) Color() TColor {
	r1 := LCL().SysCallN(3238, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGridColumnTitle) SetColor(AValue TColor) {
	LCL().SysCallN(3238, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Font() IFont {
	r1 := LCL().SysCallN(3243, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TGridColumnTitle) SetFont(AValue IFont) {
	LCL().SysCallN(3243, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumnTitle) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(3244, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TGridColumnTitle) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(3244, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) ImageLayout() TButtonLayout {
	r1 := LCL().SysCallN(3245, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TGridColumnTitle) SetImageLayout(AValue TButtonLayout) {
	LCL().SysCallN(3245, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Layout() TTextLayout {
	r1 := LCL().SysCallN(3247, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TGridColumnTitle) SetLayout(AValue TTextLayout) {
	LCL().SysCallN(3247, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) MultiLine() bool {
	r1 := LCL().SysCallN(3248, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumnTitle) SetMultiLine(AValue bool) {
	LCL().SysCallN(3248, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumnTitle) PrefixOption() TPrefixOption {
	r1 := LCL().SysCallN(3249, 0, m.Instance(), 0)
	return TPrefixOption(r1)
}

func (m *TGridColumnTitle) SetPrefixOption(AValue TPrefixOption) {
	LCL().SysCallN(3249, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) IsDefault() bool {
	r1 := LCL().SysCallN(3246, m.Instance())
	return GoBool(r1)
}

func GridColumnTitleClass() TClass {
	ret := LCL().SysCallN(3237)
	return TClass(ret)
}

func (m *TGridColumnTitle) FillTitleDefaultFont() {
	LCL().SysCallN(3241, m.Instance())
}

func (m *TGridColumnTitle) FixDesignFontsPPI(ADesignTimePPI int32) {
	LCL().SysCallN(3242, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TGridColumnTitle) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	LCL().SysCallN(3250, m.Instance(), uintptr(AToPPI), uintptr(unsafePointer(&AProportion)))
}
