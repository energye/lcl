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

// IFPCustomBrush Parent: IFPCanvasHelper
type IFPCustomBrush interface {
	IFPCanvasHelper
	CopyBrush() IFPCustomBrush          // function
	Style() types.TFPBrushStyle         // property Style Getter
	SetStyle(value types.TFPBrushStyle) // property Style Setter
	Image() IFPCustomImage              // property Image Getter
	SetImage(value IFPCustomImage)      // property Image Setter
	Pattern() IBrushPatternWrap         // property Pattern Getter
	SetPattern(value IBrushPatternWrap) // property Pattern Setter
}

type TFPCustomBrush struct {
	TFPCanvasHelper
}

func (m *TFPCustomBrush) CopyBrush() IFPCustomBrush {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomBrushAPI().SysCallN(1, m.Instance())
	return AsFPCustomBrush(r)
}

func (m *TFPCustomBrush) Style() types.TFPBrushStyle {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomBrushAPI().SysCallN(2, 0, m.Instance())
	return types.TFPBrushStyle(r)
}

func (m *TFPCustomBrush) SetStyle(value types.TFPBrushStyle) {
	if !m.IsValid() {
		return
	}
	fPCustomBrushAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomBrush) Image() IFPCustomImage {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomBrushAPI().SysCallN(3, 0, m.Instance())
	return AsFPCustomImage(r)
}

func (m *TFPCustomBrush) SetImage(value IFPCustomImage) {
	if !m.IsValid() {
		return
	}
	fPCustomBrushAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFPCustomBrush) Pattern() IBrushPatternWrap {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomBrushAPI().SysCallN(4, 0, m.Instance())
	return AsBrushPatternWrap(r)
}

func (m *TFPCustomBrush) SetPattern(value IBrushPatternWrap) {
	if !m.IsValid() {
		return
	}
	fPCustomBrushAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewFPCustomBrush class constructor
func NewFPCustomBrush() IFPCustomBrush {
	r := fPCustomBrushAPI().SysCallN(0)
	return AsFPCustomBrush(r)
}

var (
	fPCustomBrushOnce   base.Once
	fPCustomBrushImport *imports.Imports = nil
)

func fPCustomBrushAPI() *imports.Imports {
	fPCustomBrushOnce.Do(func() {
		fPCustomBrushImport = api.NewDefaultImports()
		fPCustomBrushImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomBrush_Create", 0), // constructor NewFPCustomBrush
			/* 1 */ imports.NewTable("TFPCustomBrush_CopyBrush", 0), // function CopyBrush
			/* 2 */ imports.NewTable("TFPCustomBrush_Style", 0), // property Style
			/* 3 */ imports.NewTable("TFPCustomBrush_Image", 0), // property Image
			/* 4 */ imports.NewTable("TFPCustomBrush_Pattern", 0), // property Pattern
		}
	})
	return fPCustomBrushImport
}
