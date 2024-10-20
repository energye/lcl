//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IRichMemoInline Parent: IObject
type IRichMemoInline interface {
	IObject
	Owner() ICustomRichMemo            // property
	Draw(Canvas ICanvas, ASize *TSize) // procedure
	SetVisible(AVisible bool)          // procedure
	Invalidate()                       // procedure
}

// TRichMemoInline Parent: TObject
type TRichMemoInline struct {
	TObject
}

func NewRichMemoInline() IRichMemoInline {
	r1 := richMemoInlineImportAPI().SysCallN(1)
	return AsRichMemoInline(r1)
}

func (m *TRichMemoInline) Owner() ICustomRichMemo {
	r1 := richMemoInlineImportAPI().SysCallN(4, m.Instance())
	return AsCustomRichMemo(r1)
}

func RichMemoInlineClass() TClass {
	ret := richMemoInlineImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TRichMemoInline) Draw(Canvas ICanvas, ASize *TSize) {
	richMemoInlineImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(Canvas), uintptr(unsafePointer(ASize)))
}

func (m *TRichMemoInline) SetVisible(AVisible bool) {
	richMemoInlineImportAPI().SysCallN(5, m.Instance(), PascalBool(AVisible))
}

func (m *TRichMemoInline) Invalidate() {
	richMemoInlineImportAPI().SysCallN(3, m.Instance())
}

var (
	richMemoInlineImport       *imports.Imports = nil
	richMemoInlineImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("RichMemoInline_Class", 0),
		/*1*/ imports.NewTable("RichMemoInline_Create", 0),
		/*2*/ imports.NewTable("RichMemoInline_Draw", 0),
		/*3*/ imports.NewTable("RichMemoInline_Invalidate", 0),
		/*4*/ imports.NewTable("RichMemoInline_Owner", 0),
		/*5*/ imports.NewTable("RichMemoInline_SetVisible", 0),
	}
)

func richMemoInlineImportAPI() *imports.Imports {
	if richMemoInlineImport == nil {
		richMemoInlineImport = NewDefaultImports()
		richMemoInlineImport.SetImportTable(richMemoInlineImportTables)
		richMemoInlineImportTables = nil
	}
	return richMemoInlineImport
}
