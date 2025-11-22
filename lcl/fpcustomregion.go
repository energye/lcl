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
	"github.com/energye/lcl/types"
)

// IFPCustomRegion Parent: IObject
type IFPCustomRegion interface {
	IObject
	GetBoundingRect() types.TRect          // function
	IsPointInRegion(X int32, Y int32) bool // function
}

type TFPCustomRegion struct {
	TObject
}

func (m *TFPCustomRegion) GetBoundingRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomRegionAPI().SysCallN(0, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPCustomRegion) IsPointInRegion(X int32, Y int32) bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomRegionAPI().SysCallN(1, m.Instance(), uintptr(X), uintptr(Y))
	return api.GoBool(r)
}

var (
	fPCustomRegionOnce   base.Once
	fPCustomRegionImport *imports.Imports = nil
)

func fPCustomRegionAPI() *imports.Imports {
	fPCustomRegionOnce.Do(func() {
		fPCustomRegionImport = api.NewDefaultImports()
		fPCustomRegionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomRegion_GetBoundingRect", 0), // function GetBoundingRect
			/* 1 */ imports.NewTable("TFPCustomRegion_IsPointInRegion", 0), // function IsPointInRegion
		}
	})
	return fPCustomRegionImport
}
