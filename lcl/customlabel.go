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

// ICustomLabel Parent: IGraphicControl
type ICustomLabel interface {
	IGraphicControl
	CalcFittingFontHeight(theText string, maxWidth int32, maxHeight int32, outFontHeight *int32, outNeededWidth *int32, outNeededHeight *int32) bool // function
	AdjustFontForOptimalFill() bool                                                                                                                  // function
	Paint()                                                                                                                                          // procedure
}

type TCustomLabel struct {
	TGraphicControl
}

func (m *TCustomLabel) CalcFittingFontHeight(theText string, maxWidth int32, maxHeight int32, outFontHeight *int32, outNeededWidth *int32, outNeededHeight *int32) bool {
	if !m.IsValid() {
		return false
	}
	var fontHeightPtr uintptr
	var neededWidthPtr uintptr
	var neededHeightPtr uintptr
	r := customLabelAPI().SysCallN(1, m.Instance(), api.PasStr(theText), uintptr(maxWidth), uintptr(maxHeight), uintptr(base.UnsafePointer(&fontHeightPtr)), uintptr(base.UnsafePointer(&neededWidthPtr)), uintptr(base.UnsafePointer(&neededHeightPtr)))
	*outFontHeight = int32(fontHeightPtr)
	*outNeededWidth = int32(neededWidthPtr)
	*outNeededHeight = int32(neededHeightPtr)
	return api.GoBool(r)
}

func (m *TCustomLabel) AdjustFontForOptimalFill() bool {
	if !m.IsValid() {
		return false
	}
	r := customLabelAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomLabel) Paint() {
	if !m.IsValid() {
		return
	}
	customLabelAPI().SysCallN(3, m.Instance())
}

// NewCustomLabel class constructor
func NewCustomLabel(theOwner IComponent) ICustomLabel {
	r := customLabelAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomLabel(r)
}

func TCustomLabelClass() types.TClass {
	r := customLabelAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	customLabelOnce   base.Once
	customLabelImport *imports.Imports = nil
)

func customLabelAPI() *imports.Imports {
	customLabelOnce.Do(func() {
		customLabelImport = api.NewDefaultImports()
		customLabelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomLabel_Create", 0), // constructor NewCustomLabel
			/* 1 */ imports.NewTable("TCustomLabel_CalcFittingFontHeight", 0), // function CalcFittingFontHeight
			/* 2 */ imports.NewTable("TCustomLabel_AdjustFontForOptimalFill", 0), // function AdjustFontForOptimalFill
			/* 3 */ imports.NewTable("TCustomLabel_Paint", 0), // procedure Paint
			/* 4 */ imports.NewTable("TCustomLabel_TClass", 0), // function TCustomLabelClass
		}
	})
	return customLabelImport
}
