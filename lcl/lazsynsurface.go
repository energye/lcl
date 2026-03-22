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

// ILazSynSurface Parent: IObject
type ILazSynSurface interface {
	IObject
	Assign(src ILazSynSurface)                                                                         // procedure
	Paint(canvas ICanvas, clip types.TRect)                                                            // procedure
	InvalidateLines(firstTextLine types.TLineIdx, lastTextLine types.TLineIdx, screenLineOffset int32) // procedure
	SetBounds(top int32, left int32, bottom int32, right int32)                                        // procedure
	Left() int32                                                                                       // property Left Getter
	Top() int32                                                                                        // property Top Getter
	Right() int32                                                                                      // property Right Getter
	Bottom() int32                                                                                     // property Bottom Getter
	Bounds() types.TRect                                                                               // property Bounds Getter
	DisplayView() ILazSynDisplayView                                                                   // property DisplayView Getter
	SetDisplayView(value ILazSynDisplayView)                                                           // property DisplayView Setter
}

type TLazSynSurface struct {
	TObject
}

func (m *TLazSynSurface) Assign(src ILazSynSurface) {
	if !m.IsValid() {
		return
	}
	lazSynSurfaceAPI().SysCallN(0, m.Instance(), base.GetObjectUintptr(src))
}

func (m *TLazSynSurface) Paint(canvas ICanvas, clip types.TRect) {
	if !m.IsValid() {
		return
	}
	lazSynSurfaceAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&clip)))
}

func (m *TLazSynSurface) InvalidateLines(firstTextLine types.TLineIdx, lastTextLine types.TLineIdx, screenLineOffset int32) {
	if !m.IsValid() {
		return
	}
	lazSynSurfaceAPI().SysCallN(2, m.Instance(), uintptr(firstTextLine), uintptr(lastTextLine), uintptr(screenLineOffset))
}

func (m *TLazSynSurface) SetBounds(top int32, left int32, bottom int32, right int32) {
	if !m.IsValid() {
		return
	}
	lazSynSurfaceAPI().SysCallN(3, m.Instance(), uintptr(top), uintptr(left), uintptr(bottom), uintptr(right))
}

func (m *TLazSynSurface) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynSurfaceAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TLazSynSurface) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynSurfaceAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TLazSynSurface) Right() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynSurfaceAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TLazSynSurface) Bottom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynSurfaceAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TLazSynSurface) Bounds() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	lazSynSurfaceAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynSurface) DisplayView() ILazSynDisplayView {
	if !m.IsValid() {
		return nil
	}
	r := lazSynSurfaceAPI().SysCallN(9, 0, m.Instance())
	return AsLazSynDisplayView(r)
}

func (m *TLazSynSurface) SetDisplayView(value ILazSynDisplayView) {
	if !m.IsValid() {
		return
	}
	lazSynSurfaceAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

var (
	lazSynSurfaceOnce   base.Once
	lazSynSurfaceImport *imports.Imports = nil
)

func lazSynSurfaceAPI() *imports.Imports {
	lazSynSurfaceOnce.Do(func() {
		lazSynSurfaceImport = api.NewDefaultImports()
		lazSynSurfaceImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynSurface_Assign", 0), // procedure Assign
			/* 1 */ imports.NewTable("TLazSynSurface_Paint", 0), // procedure Paint
			/* 2 */ imports.NewTable("TLazSynSurface_InvalidateLines", 0), // procedure InvalidateLines
			/* 3 */ imports.NewTable("TLazSynSurface_SetBounds", 0), // procedure SetBounds
			/* 4 */ imports.NewTable("TLazSynSurface_Left", 0), // property Left
			/* 5 */ imports.NewTable("TLazSynSurface_Top", 0), // property Top
			/* 6 */ imports.NewTable("TLazSynSurface_Right", 0), // property Right
			/* 7 */ imports.NewTable("TLazSynSurface_Bottom", 0), // property Bottom
			/* 8 */ imports.NewTable("TLazSynSurface_Bounds", 0), // property Bounds
			/* 9 */ imports.NewTable("TLazSynSurface_DisplayView", 0), // property DisplayView
		}
	})
	return lazSynSurfaceImport
}
