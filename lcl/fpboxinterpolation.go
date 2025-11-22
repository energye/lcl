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

// IFPBoxInterpolation Parent: IFPCustomInterpolation
type IFPBoxInterpolation interface {
	IFPCustomInterpolation
	Execute(X int32, Y int32, W int32, H int32) // procedure
}

type TFPBoxInterpolation struct {
	TFPCustomInterpolation
}

func (m *TFPBoxInterpolation) Execute(X int32, Y int32, W int32, H int32) {
	if !m.IsValid() {
		return
	}
	fPBoxInterpolationAPI().SysCallN(0, m.Instance(), uintptr(X), uintptr(Y), uintptr(W), uintptr(H))
}

var (
	fPBoxInterpolationOnce   base.Once
	fPBoxInterpolationImport *imports.Imports = nil
)

func fPBoxInterpolationAPI() *imports.Imports {
	fPBoxInterpolationOnce.Do(func() {
		fPBoxInterpolationImport = api.NewDefaultImports()
		fPBoxInterpolationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPBoxInterpolation_Execute", 0), // procedure Execute
		}
	})
	return fPBoxInterpolationImport
}
