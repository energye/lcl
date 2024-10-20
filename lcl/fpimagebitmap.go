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

// IFPImageBitmap Parent: ICustomBitmap
type IFPImageBitmap interface {
	ICustomBitmap
}

// TFPImageBitmap Parent: TCustomBitmap
type TFPImageBitmap struct {
	TCustomBitmap
}

func NewFPImageBitmap() IFPImageBitmap {
	r1 := fPImageBitmapImportAPI().SysCallN(1)
	return AsFPImageBitmap(r1)
}

func FPImageBitmapClass() TClass {
	ret := fPImageBitmapImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	fPImageBitmapImport       *imports.Imports = nil
	fPImageBitmapImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPImageBitmap_Class", 0),
		/*1*/ imports.NewTable("FPImageBitmap_Create", 0),
	}
)

func fPImageBitmapImportAPI() *imports.Imports {
	if fPImageBitmapImport == nil {
		fPImageBitmapImport = NewDefaultImports()
		fPImageBitmapImport.SetImportTable(fPImageBitmapImportTables)
		fPImageBitmapImportTables = nil
	}
	return fPImageBitmapImport
}
