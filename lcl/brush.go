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

// IBrush Parent: IFPCustomBrush
type IBrush interface {
	IFPCustomBrush
	Bitmap() ICustomBitmap          // property
	SetBitmap(AValue ICustomBitmap) // property
	Color() TColor                  // property
	SetColor(AValue TColor)         // property
	EqualsBrush(ABrush IBrush) bool // function
}

// TBrush Parent: TFPCustomBrush
type TBrush struct {
	TFPCustomBrush
}

func NewBrush() IBrush {
	r1 := brushImportAPI().SysCallN(3)
	return AsBrush(r1)
}

func (m *TBrush) Bitmap() ICustomBitmap {
	r1 := brushImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsCustomBitmap(r1)
}

func (m *TBrush) SetBitmap(AValue ICustomBitmap) {
	brushImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBrush) Color() TColor {
	r1 := brushImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TBrush) SetColor(AValue TColor) {
	brushImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TBrush) EqualsBrush(ABrush IBrush) bool {
	r1 := brushImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(ABrush))
	return GoBool(r1)
}

func BrushClass() TClass {
	ret := brushImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	brushImport       *imports.Imports = nil
	brushImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Brush_Bitmap", 0),
		/*1*/ imports.NewTable("Brush_Class", 0),
		/*2*/ imports.NewTable("Brush_Color", 0),
		/*3*/ imports.NewTable("Brush_Create", 0),
		/*4*/ imports.NewTable("Brush_EqualsBrush", 0),
	}
)

func brushImportAPI() *imports.Imports {
	if brushImport == nil {
		brushImport = NewDefaultImports()
		brushImport.SetImportTable(brushImportTables)
		brushImportTables = nil
	}
	return brushImport
}
