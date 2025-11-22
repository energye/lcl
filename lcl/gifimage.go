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

// IGIFImage Parent: IFPImageBitmap
type IGIFImage interface {
	IFPImageBitmap
	Interlaced() bool   // property Interlaced Getter
	BitsPerPixel() byte // property BitsPerPixel Getter
}

type TGIFImage struct {
	TFPImageBitmap
}

func (m *TGIFImage) Interlaced() bool {
	if !m.IsValid() {
		return false
	}
	r := gIFImageAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TGIFImage) BitsPerPixel() byte {
	if !m.IsValid() {
		return 0
	}
	r := gIFImageAPI().SysCallN(2, m.Instance())
	return byte(r)
}

// NewGIFImage class constructor
func NewGIFImage() IGIFImage {
	r := gIFImageAPI().SysCallN(0)
	return AsGIFImage(r)
}

var (
	gIFImageOnce   base.Once
	gIFImageImport *imports.Imports = nil
)

func gIFImageAPI() *imports.Imports {
	gIFImageOnce.Do(func() {
		gIFImageImport = api.NewDefaultImports()
		gIFImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGIFImage_Create", 0), // constructor NewGIFImage
			/* 1 */ imports.NewTable("TGIFImage_Interlaced", 0), // property Interlaced
			/* 2 */ imports.NewTable("TGIFImage_BitsPerPixel", 0), // property BitsPerPixel
		}
	})
	return gIFImageImport
}
