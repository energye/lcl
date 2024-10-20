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

// IFPCustomBrush Parent: IFPCanvasHelper
type IFPCustomBrush interface {
	IFPCanvasHelper
	Style() TFPBrushStyle           // property
	SetStyle(AValue TFPBrushStyle)  // property
	Image() IFPCustomImage          // property
	SetImage(AValue IFPCustomImage) // property
	CopyBrush() IFPCustomBrush      // function
}

// TFPCustomBrush Parent: TFPCanvasHelper
type TFPCustomBrush struct {
	TFPCanvasHelper
}

func NewFPCustomBrush() IFPCustomBrush {
	r1 := fPCustomBrushImportAPI().SysCallN(2)
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomBrush) Style() TFPBrushStyle {
	r1 := fPCustomBrushImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TFPBrushStyle(r1)
}

func (m *TFPCustomBrush) SetStyle(AValue TFPBrushStyle) {
	fPCustomBrushImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomBrush) Image() IFPCustomImage {
	r1 := fPCustomBrushImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsFPCustomImage(r1)
}

func (m *TFPCustomBrush) SetImage(AValue IFPCustomImage) {
	fPCustomBrushImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomBrush) CopyBrush() IFPCustomBrush {
	r1 := fPCustomBrushImportAPI().SysCallN(1, m.Instance())
	return AsFPCustomBrush(r1)
}

func FPCustomBrushClass() TClass {
	ret := fPCustomBrushImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	fPCustomBrushImport       *imports.Imports = nil
	fPCustomBrushImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomBrush_Class", 0),
		/*1*/ imports.NewTable("FPCustomBrush_CopyBrush", 0),
		/*2*/ imports.NewTable("FPCustomBrush_Create", 0),
		/*3*/ imports.NewTable("FPCustomBrush_Image", 0),
		/*4*/ imports.NewTable("FPCustomBrush_Style", 0),
	}
)

func fPCustomBrushImportAPI() *imports.Imports {
	if fPCustomBrushImport == nil {
		fPCustomBrushImport = NewDefaultImports()
		fPCustomBrushImport.SetImportTable(fPCustomBrushImportTables)
		fPCustomBrushImportTables = nil
	}
	return fPCustomBrushImport
}
