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

// IBrush Parent: IFPCustomBrush
type IBrush interface {
	IFPCustomBrush
	EqualsBrush(brush IBrush) bool // function
	Bitmap() ICustomBitmap         // property Bitmap Getter
	SetBitmap(value ICustomBitmap) // property Bitmap Setter
	Color() types.TColor           // property Color Getter
	SetColor(value types.TColor)   // property Color Setter
}

type TBrush struct {
	TFPCustomBrush
}

func (m *TBrush) EqualsBrush(brush IBrush) bool {
	if !m.IsValid() {
		return false
	}
	r := brushAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(brush))
	return api.GoBool(r)
}

func (m *TBrush) Bitmap() ICustomBitmap {
	if !m.IsValid() {
		return nil
	}
	r := brushAPI().SysCallN(2, 0, m.Instance())
	return AsCustomBitmap(r)
}

func (m *TBrush) SetBitmap(value ICustomBitmap) {
	if !m.IsValid() {
		return
	}
	brushAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TBrush) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := brushAPI().SysCallN(3, 0, m.Instance())
	return types.TColor(r)
}

func (m *TBrush) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	brushAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// NewBrush class constructor
func NewBrush() IBrush {
	r := brushAPI().SysCallN(0)
	return AsBrush(r)
}

var (
	brushOnce   base.Once
	brushImport *imports.Imports = nil
)

func brushAPI() *imports.Imports {
	brushOnce.Do(func() {
		brushImport = api.NewDefaultImports()
		brushImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBrush_Create", 0), // constructor NewBrush
			/* 1 */ imports.NewTable("TBrush_EqualsBrush", 0), // function EqualsBrush
			/* 2 */ imports.NewTable("TBrush_Bitmap", 0), // property Bitmap
			/* 3 */ imports.NewTable("TBrush_Color", 0), // property Color
		}
	})
	return brushImport
}
