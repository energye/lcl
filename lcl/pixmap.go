//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// IPixmap Parent: IFPImageBitmap
type IPixmap interface {
	IFPImageBitmap
}

type TPixmap struct {
	TFPImageBitmap
}

// NewPixmap class constructor
func NewPixmap() IPixmap {
	r := pixmapAPI().SysCallN(0)
	return AsPixmap(r)
}

var (
	pixmapOnce   base.Once
	pixmapImport *imports.Imports = nil
)

func pixmapAPI() *imports.Imports {
	pixmapOnce.Do(func() {
		pixmapImport = api.NewDefaultImports()
		pixmapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPixmap_Create", 0), // constructor NewPixmap
		}
	})
	return pixmapImport
}
