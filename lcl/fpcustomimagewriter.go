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

// IFPCustomImageWriter Parent: IFPCustomImageHandler
type IFPCustomImageWriter interface {
	IFPCustomImageHandler
	ImageWrite(str IStream, img IFPCustomImage) // procedure
}

type TFPCustomImageWriter struct {
	TFPCustomImageHandler
}

func (m *TFPCustomImageWriter) ImageWrite(str IStream, img IFPCustomImage) {
	if !m.IsValid() {
		return
	}
	fPCustomImageWriterAPI().SysCallN(0, m.Instance(), base.GetObjectUintptr(str), base.GetObjectUintptr(img))
}

var (
	fPCustomImageWriterOnce   base.Once
	fPCustomImageWriterImport *imports.Imports = nil
)

func fPCustomImageWriterAPI() *imports.Imports {
	fPCustomImageWriterOnce.Do(func() {
		fPCustomImageWriterImport = api.NewDefaultImports()
		fPCustomImageWriterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomImageWriter_ImageWrite", 0), // procedure ImageWrite
		}
	})
	return fPCustomImageWriterImport
}
