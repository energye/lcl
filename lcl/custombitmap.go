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

// ICustomBitmap Parent: IRasterImage
type ICustomBitmap interface {
	IRasterImage
	Handle() HBITMAP                        // property
	SetHandle(AValue HBITMAP)               // property
	HandleType() TBitmapHandleType          // property
	SetHandleType(AValue TBitmapHandleType) // property
	Monochrome() bool                       // property
	SetMonochrome(AValue bool)              // property
	ReleaseHandle() HBITMAP                 // function
	SetSize(AWidth, AHeight int32)          // procedure
}

// TCustomBitmap Parent: TRasterImage
type TCustomBitmap struct {
	TRasterImage
}

func NewCustomBitmap() ICustomBitmap {
	r1 := customBitmapImportAPI().SysCallN(1)
	return AsCustomBitmap(r1)
}

func (m *TCustomBitmap) Handle() HBITMAP {
	r1 := customBitmapImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TCustomBitmap) SetHandle(AValue HBITMAP) {
	customBitmapImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitmap) HandleType() TBitmapHandleType {
	r1 := customBitmapImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TBitmapHandleType(r1)
}

func (m *TCustomBitmap) SetHandleType(AValue TBitmapHandleType) {
	customBitmapImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitmap) Monochrome() bool {
	r1 := customBitmapImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomBitmap) SetMonochrome(AValue bool) {
	customBitmapImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomBitmap) ReleaseHandle() HBITMAP {
	r1 := customBitmapImportAPI().SysCallN(5, m.Instance())
	return HBITMAP(r1)
}

func CustomBitmapClass() TClass {
	ret := customBitmapImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomBitmap) SetSize(AWidth, AHeight int32) {
	customBitmapImportAPI().SysCallN(6, m.Instance(), uintptr(AWidth), uintptr(AHeight))
}

var (
	customBitmapImport       *imports.Imports = nil
	customBitmapImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomBitmap_Class", 0),
		/*1*/ imports.NewTable("CustomBitmap_Create", 0),
		/*2*/ imports.NewTable("CustomBitmap_Handle", 0),
		/*3*/ imports.NewTable("CustomBitmap_HandleType", 0),
		/*4*/ imports.NewTable("CustomBitmap_Monochrome", 0),
		/*5*/ imports.NewTable("CustomBitmap_ReleaseHandle", 0),
		/*6*/ imports.NewTable("CustomBitmap_SetSize", 0),
	}
)

func customBitmapImportAPI() *imports.Imports {
	if customBitmapImport == nil {
		customBitmapImport = NewDefaultImports()
		customBitmapImport.SetImportTable(customBitmapImportTables)
		customBitmapImportTables = nil
	}
	return customBitmapImport
}
