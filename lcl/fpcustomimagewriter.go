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

// IFPCustomImageWriter Is Abstract Class Parent: IFPCustomImageHandler
type IFPCustomImageWriter interface {
	IFPCustomImageHandler
	ImageWrite(Str IStream, Img IFPCustomImage) // procedure
}

// TFPCustomImageWriter Is Abstract Class Parent: TFPCustomImageHandler
type TFPCustomImageWriter struct {
	TFPCustomImageHandler
}

func FPCustomImageWriterClass() TClass {
	ret := fPCustomImageWriterImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TFPCustomImageWriter) ImageWrite(Str IStream, Img IFPCustomImage) {
	fPCustomImageWriterImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Img))
}

var (
	fPCustomImageWriterImport       *imports.Imports = nil
	fPCustomImageWriterImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomImageWriter_Class", 0),
		/*1*/ imports.NewTable("FPCustomImageWriter_ImageWrite", 0),
	}
)

func fPCustomImageWriterImportAPI() *imports.Imports {
	if fPCustomImageWriterImport == nil {
		fPCustomImageWriterImport = NewDefaultImports()
		fPCustomImageWriterImport.SetImportTable(fPCustomImageWriterImportTables)
		fPCustomImageWriterImportTables = nil
	}
	return fPCustomImageWriterImport
}
