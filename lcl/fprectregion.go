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

// IFPRectRegion Parent: IFPCustomRegion
type IFPRectRegion interface {
	IFPCustomRegion
}

// TFPRectRegion Parent: TFPCustomRegion
type TFPRectRegion struct {
	TFPCustomRegion
}

func NewFPRectRegion() IFPRectRegion {
	r1 := fPRectRegionImportAPI().SysCallN(1)
	return AsFPRectRegion(r1)
}

func FPRectRegionClass() TClass {
	ret := fPRectRegionImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	fPRectRegionImport       *imports.Imports = nil
	fPRectRegionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPRectRegion_Class", 0),
		/*1*/ imports.NewTable("FPRectRegion_Create", 0),
	}
)

func fPRectRegionImportAPI() *imports.Imports {
	if fPRectRegionImport == nil {
		fPRectRegionImport = NewDefaultImports()
		fPRectRegionImport.SetImportTable(fPRectRegionImportTables)
		fPRectRegionImportTables = nil
	}
	return fPRectRegionImport
}
