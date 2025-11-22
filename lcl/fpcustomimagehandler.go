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

// IFPCustomImageHandler Parent: IObject
type IFPCustomImageHandler interface {
	IObject
	SetOnProgress(fn TFPImgProgressEvent) // property event
}

type TFPCustomImageHandler struct {
	TObject
}

func (m *TFPCustomImageHandler) SetOnProgress(fn TFPImgProgressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTFPImgProgressEvent(fn)
	base.SetEvent(m, 1, fPCustomImageHandlerAPI(), api.MakeEventDataPtr(cb))
}

// NewFPCustomImageHandler class constructor
func NewFPCustomImageHandler() IFPCustomImageHandler {
	r := fPCustomImageHandlerAPI().SysCallN(0)
	return AsFPCustomImageHandler(r)
}

var (
	fPCustomImageHandlerOnce   base.Once
	fPCustomImageHandlerImport *imports.Imports = nil
)

func fPCustomImageHandlerAPI() *imports.Imports {
	fPCustomImageHandlerOnce.Do(func() {
		fPCustomImageHandlerImport = api.NewDefaultImports()
		fPCustomImageHandlerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomImageHandler_Create", 0), // constructor NewFPCustomImageHandler
			/* 1 */ imports.NewTable("TFPCustomImageHandler_OnProgress", 0), // event OnProgress
		}
	})
	return fPCustomImageHandlerImport
}
