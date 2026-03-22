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

// ILazSynSurfaceWithText Parent: ILazSynSurface
type ILazSynSurfaceWithText interface {
	ILazSynSurface
	TextArea() ILazSynTextArea         // property TextArea Getter
	SetTextArea(value ILazSynTextArea) // property TextArea Setter
}

type TLazSynSurfaceWithText struct {
	TLazSynSurface
}

func (m *TLazSynSurfaceWithText) TextArea() ILazSynTextArea {
	if !m.IsValid() {
		return nil
	}
	r := lazSynSurfaceWithTextAPI().SysCallN(1, 0, m.Instance())
	return AsLazSynTextArea(r)
}

func (m *TLazSynSurfaceWithText) SetTextArea(value ILazSynTextArea) {
	if !m.IsValid() {
		return
	}
	lazSynSurfaceWithTextAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewLazSynSurfaceWithText class constructor
func NewLazSynSurfaceWithText(owner IWinControl) ILazSynSurfaceWithText {
	r := lazSynSurfaceWithTextAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazSynSurfaceWithText(r)
}

var (
	lazSynSurfaceWithTextOnce   base.Once
	lazSynSurfaceWithTextImport *imports.Imports = nil
)

func lazSynSurfaceWithTextAPI() *imports.Imports {
	lazSynSurfaceWithTextOnce.Do(func() {
		lazSynSurfaceWithTextImport = api.NewDefaultImports()
		lazSynSurfaceWithTextImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynSurfaceWithText_Create", 0), // constructor NewLazSynSurfaceWithText
			/* 1 */ imports.NewTable("TLazSynSurfaceWithText_TextArea", 0), // property TextArea
		}
	})
	return lazSynSurfaceWithTextImport
}
