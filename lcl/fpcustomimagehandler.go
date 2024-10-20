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

// IFPCustomImageHandler Parent: IObject
type IFPCustomImageHandler interface {
	IObject
	SetOnProgress(fn TFPImgProgressEvent) // property event
}

// TFPCustomImageHandler Parent: TObject
type TFPCustomImageHandler struct {
	TObject
	progressPtr uintptr
}

func NewFPCustomImageHandler() IFPCustomImageHandler {
	r1 := fPCustomImageHandlerImportAPI().SysCallN(1)
	return AsFPCustomImageHandler(r1)
}

func FPCustomImageHandlerClass() TClass {
	ret := fPCustomImageHandlerImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TFPCustomImageHandler) SetOnProgress(fn TFPImgProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	fPCustomImageHandlerImportAPI().SysCallN(2, m.Instance(), m.progressPtr)
}

var (
	fPCustomImageHandlerImport       *imports.Imports = nil
	fPCustomImageHandlerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomImageHandler_Class", 0),
		/*1*/ imports.NewTable("FPCustomImageHandler_Create", 0),
		/*2*/ imports.NewTable("FPCustomImageHandler_SetOnProgress", 0),
	}
)

func fPCustomImageHandlerImportAPI() *imports.Imports {
	if fPCustomImageHandlerImport == nil {
		fPCustomImageHandlerImport = NewDefaultImports()
		fPCustomImageHandlerImport.SetImportTable(fPCustomImageHandlerImportTables)
		fPCustomImageHandlerImportTables = nil
	}
	return fPCustomImageHandlerImport
}
