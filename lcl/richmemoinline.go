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

// IRichMemoInline Parent: IObject
type IRichMemoInline interface {
	IObject
	Draw(canvas ICanvas, size types.TSize) // procedure
	SetVisible(visible bool)               // procedure
	Invalidate()                           // procedure
	Owner() ICustomRichMemo                // property Owner Getter
}

type TRichMemoInline struct {
	TObject
}

func (m *TRichMemoInline) Draw(canvas ICanvas, size types.TSize) {
	if !m.IsValid() {
		return
	}
	richMemoInlineAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&size)))
}

func (m *TRichMemoInline) SetVisible(visible bool) {
	if !m.IsValid() {
		return
	}
	richMemoInlineAPI().SysCallN(2, m.Instance(), api.PasBool(visible))
}

func (m *TRichMemoInline) Invalidate() {
	if !m.IsValid() {
		return
	}
	richMemoInlineAPI().SysCallN(3, m.Instance())
}

func (m *TRichMemoInline) Owner() ICustomRichMemo {
	if !m.IsValid() {
		return nil
	}
	r := richMemoInlineAPI().SysCallN(4, m.Instance())
	return AsCustomRichMemo(r)
}

// NewRichMemoInline class constructor
func NewRichMemoInline() IRichMemoInline {
	r := richMemoInlineAPI().SysCallN(0)
	return AsRichMemoInline(r)
}

var (
	richMemoInlineOnce   base.Once
	richMemoInlineImport *imports.Imports = nil
)

func richMemoInlineAPI() *imports.Imports {
	richMemoInlineOnce.Do(func() {
		richMemoInlineImport = api.NewDefaultImports()
		richMemoInlineImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRichMemoInline_Create", 0), // constructor NewRichMemoInline
			/* 1 */ imports.NewTable("TRichMemoInline_Draw", 0), // procedure Draw
			/* 2 */ imports.NewTable("TRichMemoInline_SetVisible", 0), // procedure SetVisible
			/* 3 */ imports.NewTable("TRichMemoInline_Invalidate", 0), // procedure Invalidate
			/* 4 */ imports.NewTable("TRichMemoInline_Owner", 0), // property Owner
		}
	})
	return richMemoInlineImport
}
