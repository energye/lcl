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

// ILazSynGutterArea Parent: ILazSynSurfaceWithText
type ILazSynGutterArea interface {
	ILazSynSurfaceWithText
	Gutter() ISynGutter         // property Gutter Getter
	SetGutter(value ISynGutter) // property Gutter Setter
	TextBounds() types.TRect    // property TextBounds Getter
}

type TLazSynGutterArea struct {
	TLazSynSurfaceWithText
}

func (m *TLazSynGutterArea) Gutter() ISynGutter {
	if !m.IsValid() {
		return nil
	}
	r := lazSynGutterAreaAPI().SysCallN(1, 0, m.Instance())
	return AsSynGutter(r)
}

func (m *TLazSynGutterArea) SetGutter(value ISynGutter) {
	if !m.IsValid() {
		return
	}
	lazSynGutterAreaAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazSynGutterArea) TextBounds() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	lazSynGutterAreaAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

// NewLazSynGutterArea class constructor
func NewLazSynGutterArea(owner IWinControl) ILazSynGutterArea {
	r := lazSynGutterAreaAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazSynGutterArea(r)
}

var (
	lazSynGutterAreaOnce   base.Once
	lazSynGutterAreaImport *imports.Imports = nil
)

func lazSynGutterAreaAPI() *imports.Imports {
	lazSynGutterAreaOnce.Do(func() {
		lazSynGutterAreaImport = api.NewDefaultImports()
		lazSynGutterAreaImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynGutterArea_Create", 0), // constructor NewLazSynGutterArea
			/* 1 */ imports.NewTable("TLazSynGutterArea_Gutter", 0), // property Gutter
			/* 2 */ imports.NewTable("TLazSynGutterArea_TextBounds", 0), // property TextBounds
		}
	})
	return lazSynGutterAreaImport
}
