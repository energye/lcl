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

// IBitmap Parent: IFPImageBitmap
type IBitmap interface {
	IFPImageBitmap
}

type TBitmap struct {
	TFPImageBitmap
}

// NewBitmap class constructor
func NewBitmap() IBitmap {
	r := bitmapAPI().SysCallN(0)
	return AsBitmap(r)
}

var (
	bitmapOnce   base.Once
	bitmapImport *imports.Imports = nil
)

func bitmapAPI() *imports.Imports {
	bitmapOnce.Do(func() {
		bitmapImport = api.NewDefaultImports()
		bitmapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBitmap_Create", 0), // constructor NewBitmap
		}
	})
	return bitmapImport
}
