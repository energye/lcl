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

// IFPCustomInterpolation Parent: IObject
type IFPCustomInterpolation interface {
	IObject
	Canvas() IFPCustomCanvas // property Canvas Getter
	Image() IFPCustomImage   // property Image Getter
}

type TFPCustomInterpolation struct {
	TObject
}

func (m *TFPCustomInterpolation) Canvas() IFPCustomCanvas {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomInterpolationAPI().SysCallN(0, m.Instance())
	return AsFPCustomCanvas(r)
}

func (m *TFPCustomInterpolation) Image() IFPCustomImage {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomInterpolationAPI().SysCallN(1, m.Instance())
	return AsFPCustomImage(r)
}

var (
	fPCustomInterpolationOnce   base.Once
	fPCustomInterpolationImport *imports.Imports = nil
)

func fPCustomInterpolationAPI() *imports.Imports {
	fPCustomInterpolationOnce.Do(func() {
		fPCustomInterpolationImport = api.NewDefaultImports()
		fPCustomInterpolationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomInterpolation_Canvas", 0), // property Canvas
			/* 1 */ imports.NewTable("TFPCustomInterpolation_Image", 0), // property Image
		}
	})
	return fPCustomInterpolationImport
}
