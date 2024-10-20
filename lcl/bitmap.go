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

// IBitmap Parent: IFPImageBitmap
type IBitmap interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
}

// TBitmap Parent: TFPImageBitmap
type TBitmap struct {
	TFPImageBitmap
}

func NewBitmap() IBitmap {
	r1 := bitmapImportAPI().SysCallN(1)
	return AsBitmap(r1)
}

func BitmapClass() TClass {
	ret := bitmapImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	bitmapImport       *imports.Imports = nil
	bitmapImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Bitmap_Class", 0),
		/*1*/ imports.NewTable("Bitmap_Create", 0),
	}
)

func bitmapImportAPI() *imports.Imports {
	if bitmapImport == nil {
		bitmapImport = NewDefaultImports()
		bitmapImport.SetImportTable(bitmapImportTables)
		bitmapImportTables = nil
	}
	return bitmapImport
}
