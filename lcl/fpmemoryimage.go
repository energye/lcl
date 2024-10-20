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

// IFPMemoryImage Parent: IFPCustomImage
type IFPMemoryImage interface {
	IFPCustomImage
}

// TFPMemoryImage Parent: TFPCustomImage
type TFPMemoryImage struct {
	TFPCustomImage
}

func NewFPMemoryImage(AWidth, AHeight int32) IFPMemoryImage {
	r1 := fPMemoryImageImportAPI().SysCallN(1, uintptr(AWidth), uintptr(AHeight))
	return AsFPMemoryImage(r1)
}

func FPMemoryImageClass() TClass {
	ret := fPMemoryImageImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	fPMemoryImageImport       *imports.Imports = nil
	fPMemoryImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPMemoryImage_Class", 0),
		/*1*/ imports.NewTable("FPMemoryImage_Create", 0),
	}
)

func fPMemoryImageImportAPI() *imports.Imports {
	if fPMemoryImageImport == nil {
		fPMemoryImageImport = NewDefaultImports()
		fPMemoryImageImport.SetImportTable(fPMemoryImageImportTables)
		fPMemoryImageImportTables = nil
	}
	return fPMemoryImageImport
}
