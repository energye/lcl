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

// IGIFImage Parent: IFPImageBitmap
type IGIFImage interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Interlaced() bool   // property
	BitsPerPixel() byte // property
}

// TGIFImage Parent: TFPImageBitmap
type TGIFImage struct {
	TFPImageBitmap
}

func NewGIFImage() IGIFImage {
	r1 := gIFImageImportAPI().SysCallN(2)
	return AsGIFImage(r1)
}

func (m *TGIFImage) Interlaced() bool {
	r1 := gIFImageImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TGIFImage) BitsPerPixel() byte {
	r1 := gIFImageImportAPI().SysCallN(0, m.Instance())
	return byte(r1)
}

func GIFImageClass() TClass {
	ret := gIFImageImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	gIFImageImport       *imports.Imports = nil
	gIFImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GIFImage_BitsPerPixel", 0),
		/*1*/ imports.NewTable("GIFImage_Class", 0),
		/*2*/ imports.NewTable("GIFImage_Create", 0),
		/*3*/ imports.NewTable("GIFImage_Interlaced", 0),
	}
)

func gIFImageImportAPI() *imports.Imports {
	if gIFImageImport == nil {
		gIFImageImport = NewDefaultImports()
		gIFImageImport.SetImportTable(gIFImageImportTables)
		gIFImageImportTables = nil
	}
	return gIFImageImport
}
