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

// ICustomBitmap Parent: IRasterImage
type ICustomBitmap interface {
	IRasterImage
	ReleaseHandle() types.HBitmap                // function
	SetSize(width int32, height int32)           // procedure
	Handle() types.HBitmap                       // property Handle Getter
	SetHandle(value types.HBitmap)               // property Handle Setter
	HandleType() types.TBitmapHandleType         // property HandleType Getter
	SetHandleType(value types.TBitmapHandleType) // property HandleType Setter
	Monochrome() bool                            // property Monochrome Getter
	SetMonochrome(value bool)                    // property Monochrome Setter
}

type TCustomBitmap struct {
	TRasterImage
}

func (m *TCustomBitmap) ReleaseHandle() types.HBitmap {
	if !m.IsValid() {
		return 0
	}
	r := customBitmapAPI().SysCallN(1, m.Instance())
	return types.HBitmap(r)
}

func (m *TCustomBitmap) SetSize(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	customBitmapAPI().SysCallN(2, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TCustomBitmap) Handle() types.HBitmap {
	if !m.IsValid() {
		return 0
	}
	r := customBitmapAPI().SysCallN(3, 0, m.Instance())
	return types.HBitmap(r)
}

func (m *TCustomBitmap) SetHandle(value types.HBitmap) {
	if !m.IsValid() {
		return
	}
	customBitmapAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitmap) HandleType() types.TBitmapHandleType {
	if !m.IsValid() {
		return 0
	}
	r := customBitmapAPI().SysCallN(4, 0, m.Instance())
	return types.TBitmapHandleType(r)
}

func (m *TCustomBitmap) SetHandleType(value types.TBitmapHandleType) {
	if !m.IsValid() {
		return
	}
	customBitmapAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitmap) Monochrome() bool {
	if !m.IsValid() {
		return false
	}
	r := customBitmapAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomBitmap) SetMonochrome(value bool) {
	if !m.IsValid() {
		return
	}
	customBitmapAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

// NewCustomBitmap class constructor
func NewCustomBitmap() ICustomBitmap {
	r := customBitmapAPI().SysCallN(0)
	return AsCustomBitmap(r)
}

var (
	customBitmapOnce   base.Once
	customBitmapImport *imports.Imports = nil
)

func customBitmapAPI() *imports.Imports {
	customBitmapOnce.Do(func() {
		customBitmapImport = api.NewDefaultImports()
		customBitmapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomBitmap_Create", 0), // constructor NewCustomBitmap
			/* 1 */ imports.NewTable("TCustomBitmap_ReleaseHandle", 0), // function ReleaseHandle
			/* 2 */ imports.NewTable("TCustomBitmap_SetSize", 0), // procedure SetSize
			/* 3 */ imports.NewTable("TCustomBitmap_Handle", 0), // property Handle
			/* 4 */ imports.NewTable("TCustomBitmap_HandleType", 0), // property HandleType
			/* 5 */ imports.NewTable("TCustomBitmap_Monochrome", 0), // property Monochrome
		}
	})
	return customBitmapImport
}
