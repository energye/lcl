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

// ICustomLabel Parent: IGraphicControl
type ICustomLabel interface {
	IGraphicControl
	CalcFittingFontHeight(TheText string, MaxWidth, MaxHeight int32, OutFontHeight, OutNeededWidth, OutNeededHeight *int32) bool // function
	AdjustFontForOptimalFill() bool                                                                                              // function
	Paint()                                                                                                                      // procedure
}

// TCustomLabel Parent: TGraphicControl
type TCustomLabel struct {
	TGraphicControl
}

func NewCustomLabel(TheOwner IComponent) ICustomLabel {
	r1 := customLabelImportAPI().SysCallN(3, GetObjectUintptr(TheOwner))
	return AsCustomLabel(r1)
}

func (m *TCustomLabel) CalcFittingFontHeight(TheText string, MaxWidth, MaxHeight int32, OutFontHeight, OutNeededWidth, OutNeededHeight *int32) bool {
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := customLabelImportAPI().SysCallN(1, m.Instance(), PascalStr(TheText), uintptr(MaxWidth), uintptr(MaxHeight), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*OutFontHeight = int32(result2)
	*OutNeededWidth = int32(result3)
	*OutNeededHeight = int32(result4)
	return GoBool(r1)
}

func (m *TCustomLabel) AdjustFontForOptimalFill() bool {
	r1 := customLabelImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func CustomLabelClass() TClass {
	ret := customLabelImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomLabel) Paint() {
	customLabelImportAPI().SysCallN(4, m.Instance())
}

var (
	customLabelImport       *imports.Imports = nil
	customLabelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomLabel_AdjustFontForOptimalFill", 0),
		/*1*/ imports.NewTable("CustomLabel_CalcFittingFontHeight", 0),
		/*2*/ imports.NewTable("CustomLabel_Class", 0),
		/*3*/ imports.NewTable("CustomLabel_Create", 0),
		/*4*/ imports.NewTable("CustomLabel_Paint", 0),
	}
)

func customLabelImportAPI() *imports.Imports {
	if customLabelImport == nil {
		customLabelImport = NewDefaultImports()
		customLabelImport.SetImportTable(customLabelImportTables)
		customLabelImportTables = nil
	}
	return customLabelImport
}
