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

// IRegion Parent: IGraphicsObject
type IRegion interface {
	IGraphicsObject
	ClipRect() (resultRect TRect)      // property
	SetClipRect(AValue *TRect)         // property
	AddRectangle(X1, Y1, X2, Y2 int32) // procedure
}

// TRegion Parent: TGraphicsObject
type TRegion struct {
	TGraphicsObject
}

func NewRegion() IRegion {
	r1 := regionImportAPI().SysCallN(3)
	return AsRegion(r1)
}

func (m *TRegion) ClipRect() (resultRect TRect) {
	regionImportAPI().SysCallN(2, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TRegion) SetClipRect(AValue *TRect) {
	regionImportAPI().SysCallN(2, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func RegionClass() TClass {
	ret := regionImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TRegion) AddRectangle(X1, Y1, X2, Y2 int32) {
	regionImportAPI().SysCallN(0, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

var (
	regionImport       *imports.Imports = nil
	regionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Region_AddRectangle", 0),
		/*1*/ imports.NewTable("Region_Class", 0),
		/*2*/ imports.NewTable("Region_ClipRect", 0),
		/*3*/ imports.NewTable("Region_Create", 0),
	}
)

func regionImportAPI() *imports.Imports {
	if regionImport == nil {
		regionImport = NewDefaultImports()
		regionImport.SetImportTable(regionImportTables)
		regionImportTables = nil
	}
	return regionImport
}
