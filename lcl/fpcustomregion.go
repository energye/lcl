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

// IFPCustomRegion Is Abstract Class Parent: IObject
type IFPCustomRegion interface {
	IObject
	GetBoundingRect() (resultRect TRect) // function Is Abstract
	IsPointInRegion(AX, AY int32) bool   // function Is Abstract
}

// TFPCustomRegion Is Abstract Class Parent: TObject
type TFPCustomRegion struct {
	TObject
}

func (m *TFPCustomRegion) GetBoundingRect() (resultRect TRect) {
	fPCustomRegionImportAPI().SysCallN(1, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TFPCustomRegion) IsPointInRegion(AX, AY int32) bool {
	r1 := fPCustomRegionImportAPI().SysCallN(2, m.Instance(), uintptr(AX), uintptr(AY))
	return GoBool(r1)
}

func FPCustomRegionClass() TClass {
	ret := fPCustomRegionImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	fPCustomRegionImport       *imports.Imports = nil
	fPCustomRegionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomRegion_Class", 0),
		/*1*/ imports.NewTable("FPCustomRegion_GetBoundingRect", 0),
		/*2*/ imports.NewTable("FPCustomRegion_IsPointInRegion", 0),
	}
)

func fPCustomRegionImportAPI() *imports.Imports {
	if fPCustomRegionImport == nil {
		fPCustomRegionImport = NewDefaultImports()
		fPCustomRegionImport.SetImportTable(fPCustomRegionImportTables)
		fPCustomRegionImportTables = nil
	}
	return fPCustomRegionImport
}
