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

// IFPMemoryImage Parent: IFPCustomImage
type IFPMemoryImage interface {
	IFPCustomImage
}

type TFPMemoryImage struct {
	TFPCustomImage
}

// NewFPMemoryImage class constructor
func NewFPMemoryImage(width int32, height int32) IFPMemoryImage {
	r := fPMemoryImageAPI().SysCallN(0, uintptr(width), uintptr(height))
	return AsFPMemoryImage(r)
}

var (
	fPMemoryImageOnce   base.Once
	fPMemoryImageImport *imports.Imports = nil
)

func fPMemoryImageAPI() *imports.Imports {
	fPMemoryImageOnce.Do(func() {
		fPMemoryImageImport = api.NewDefaultImports()
		fPMemoryImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPMemoryImage_Create", 0), // constructor NewFPMemoryImage
		}
	})
	return fPMemoryImageImport
}
