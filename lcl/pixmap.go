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

// IPixmap Parent: IFPImageBitmap
type IPixmap interface {
	IFPImageBitmap
}

// TPixmap Parent: TFPImageBitmap
type TPixmap struct {
	TFPImageBitmap
}

func NewPixmap() IPixmap {
	r1 := pixmapImportAPI().SysCallN(1)
	return AsPixmap(r1)
}

func PixmapClass() TClass {
	ret := pixmapImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	pixmapImport       *imports.Imports = nil
	pixmapImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Pixmap_Class", 0),
		/*1*/ imports.NewTable("Pixmap_Create", 0),
	}
)

func pixmapImportAPI() *imports.Imports {
	if pixmapImport == nil {
		pixmapImport = NewDefaultImports()
		pixmapImport.SetImportTable(pixmapImportTables)
		pixmapImportTables = nil
	}
	return pixmapImport
}
