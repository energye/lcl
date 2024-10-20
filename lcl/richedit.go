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

// IRichEdit Parent: IRichMemo
type IRichEdit interface {
	IRichMemo
	Zoom() int32                                                                             // property
	SetZoom(AValue int32)                                                                    // property
	HideScrollBars() bool                                                                    // property
	SetHideScrollBars(AValue bool)                                                           // property
	PlainText() bool                                                                         // property
	SetPlainText(AValue bool)                                                                // property
	DefAttributes() ITextAttributes                                                          // property
	SetDefAttributes(AValue ITextAttributes)                                                 // property
	SelAttributes() ITextAttributes                                                          // property
	SetSelAttributes(AValue ITextAttributes)                                                 // property
	Paragraph() IParaAttributes                                                              // property
	SetParagraph(AValue IParaAttributes)                                                     // property
	FindText(ASearchStr string, AStartPos int32, ALength int32, AOptions TSearchTypes) int32 // function
}

// TRichEdit Parent: TRichMemo
type TRichEdit struct {
	TRichMemo
}

func NewRichEdit(AOnwer IComponent) IRichEdit {
	r1 := richEditImportAPI().SysCallN(1, GetObjectUintptr(AOnwer))
	return AsRichEdit(r1)
}

func (m *TRichEdit) Zoom() int32 {
	r1 := richEditImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TRichEdit) SetZoom(AValue int32) {
	richEditImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichEdit) HideScrollBars() bool {
	r1 := richEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichEdit) SetHideScrollBars(AValue bool) {
	richEditImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichEdit) PlainText() bool {
	r1 := richEditImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichEdit) SetPlainText(AValue bool) {
	richEditImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichEdit) DefAttributes() ITextAttributes {
	r1 := richEditImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsTextAttributes(r1)
}

func (m *TRichEdit) SetDefAttributes(AValue ITextAttributes) {
	richEditImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TRichEdit) SelAttributes() ITextAttributes {
	r1 := richEditImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return AsTextAttributes(r1)
}

func (m *TRichEdit) SetSelAttributes(AValue ITextAttributes) {
	richEditImportAPI().SysCallN(7, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TRichEdit) Paragraph() IParaAttributes {
	r1 := richEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsParaAttributes(r1)
}

func (m *TRichEdit) SetParagraph(AValue IParaAttributes) {
	richEditImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TRichEdit) FindText(ASearchStr string, AStartPos int32, ALength int32, AOptions TSearchTypes) int32 {
	r1 := richEditImportAPI().SysCallN(3, m.Instance(), PascalStr(ASearchStr), uintptr(AStartPos), uintptr(ALength), uintptr(AOptions))
	return int32(r1)
}

func RichEditClass() TClass {
	ret := richEditImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	richEditImport       *imports.Imports = nil
	richEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("RichEdit_Class", 0),
		/*1*/ imports.NewTable("RichEdit_Create", 0),
		/*2*/ imports.NewTable("RichEdit_DefAttributes", 0),
		/*3*/ imports.NewTable("RichEdit_FindText", 0),
		/*4*/ imports.NewTable("RichEdit_HideScrollBars", 0),
		/*5*/ imports.NewTable("RichEdit_Paragraph", 0),
		/*6*/ imports.NewTable("RichEdit_PlainText", 0),
		/*7*/ imports.NewTable("RichEdit_SelAttributes", 0),
		/*8*/ imports.NewTable("RichEdit_Zoom", 0),
	}
)

func richEditImportAPI() *imports.Imports {
	if richEditImport == nil {
		richEditImport = NewDefaultImports()
		richEditImport.SetImportTable(richEditImportTables)
		richEditImportTables = nil
	}
	return richEditImport
}
