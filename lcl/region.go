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

// IRegion Parent: IGraphicsObject
type IRegion interface {
	IGraphicsObject
	// AddRectangle
	//  Convenience routines to add elements to the region
	AddRectangle(x1 int32, y1 int32, x2 int32, y2 int32) // procedure
	ClipRect() types.TRect                               // property ClipRect Getter
	SetClipRect(value types.TRect)                       // property ClipRect Setter
}

type TRegion struct {
	TGraphicsObject
}

func (m *TRegion) AddRectangle(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	regionAPI().SysCallN(1, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TRegion) ClipRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	regionAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegion) SetClipRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	regionAPI().SysCallN(2, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

// NewRegion class constructor
func NewRegion() IRegion {
	r := regionAPI().SysCallN(0)
	return AsRegion(r)
}

var (
	regionOnce   base.Once
	regionImport *imports.Imports = nil
)

func regionAPI() *imports.Imports {
	regionOnce.Do(func() {
		regionImport = api.NewDefaultImports()
		regionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRegion_Create", 0), // constructor NewRegion
			/* 1 */ imports.NewTable("TRegion_AddRectangle", 0), // procedure AddRectangle
			/* 2 */ imports.NewTable("TRegion_ClipRect", 0), // property ClipRect
		}
	})
	return regionImport
}
