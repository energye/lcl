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

// IFPCustomImageReader Is Abstract Class Parent: IFPCustomImageHandler
type IFPCustomImageReader interface {
	IFPCustomImageHandler
	DefaultImageClass() TFPCustomImageClass                   // property
	SetDefaultImageClass(AValue TFPCustomImageClass)          // property
	ImageRead(Str IStream, Img IFPCustomImage) IFPCustomImage // function
	CheckContents(Str IStream) bool                           // function
}

// TFPCustomImageReader Is Abstract Class Parent: TFPCustomImageHandler
type TFPCustomImageReader struct {
	TFPCustomImageHandler
}

func (m *TFPCustomImageReader) DefaultImageClass() TFPCustomImageClass {
	r1 := fPCustomImageReaderImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TFPCustomImageClass(r1)
}

func (m *TFPCustomImageReader) SetDefaultImageClass(AValue TFPCustomImageClass) {
	fPCustomImageReaderImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomImageReader) ImageRead(Str IStream, Img IFPCustomImage) IFPCustomImage {
	r1 := fPCustomImageReaderImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Img))
	return AsFPCustomImage(r1)
}

func (m *TFPCustomImageReader) CheckContents(Str IStream) bool {
	r1 := fPCustomImageReaderImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Str))
	return GoBool(r1)
}

func FPCustomImageReaderClass() TClass {
	ret := fPCustomImageReaderImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	fPCustomImageReaderImport       *imports.Imports = nil
	fPCustomImageReaderImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomImageReader_CheckContents", 0),
		/*1*/ imports.NewTable("FPCustomImageReader_Class", 0),
		/*2*/ imports.NewTable("FPCustomImageReader_DefaultImageClass", 0),
		/*3*/ imports.NewTable("FPCustomImageReader_ImageRead", 0),
	}
)

func fPCustomImageReaderImportAPI() *imports.Imports {
	if fPCustomImageReaderImport == nil {
		fPCustomImageReaderImport = NewDefaultImports()
		fPCustomImageReaderImport.SetImportTable(fPCustomImageReaderImportTables)
		fPCustomImageReaderImportTables = nil
	}
	return fPCustomImageReaderImport
}
