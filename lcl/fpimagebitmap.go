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

// IFPImageBitmap Parent: ICustomBitmap
type IFPImageBitmap interface {
	ICustomBitmap
}

type TFPImageBitmap struct {
	TCustomBitmap
}

// FPImageBitmap  is static instance
var FPImageBitmap _FPImageBitmapClass

// _FPImageBitmapClass is class type defined by TFPImageBitmap
type _FPImageBitmapClass uintptr

func (_FPImageBitmapClass) IsFileExtensionSupported(fileExtension string) bool {
	r := fPImageBitmapAPI().SysCallN(1, api.PasStr(fileExtension))
	return api.GoBool(r)
}

// NewFPImageBitmap class constructor
func NewFPImageBitmap() IFPImageBitmap {
	r := fPImageBitmapAPI().SysCallN(0)
	return AsFPImageBitmap(r)
}

var (
	fPImageBitmapOnce   base.Once
	fPImageBitmapImport *imports.Imports = nil
)

func fPImageBitmapAPI() *imports.Imports {
	fPImageBitmapOnce.Do(func() {
		fPImageBitmapImport = api.NewDefaultImports()
		fPImageBitmapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPImageBitmap_Create", 0), // constructor NewFPImageBitmap
			/* 1 */ imports.NewTable("TFPImageBitmap_IsFileExtensionSupported", 0), // static function IsFileExtensionSupported
		}
	})
	return fPImageBitmapImport
}
