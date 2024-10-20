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

// ITextAttributes Parent: IPersistent
type ITextAttributes interface {
	IPersistent
	Name() string                   // property
	SetName(AValue string)          // property
	Pitch() TFontPitch              // property
	SetPitch(AValue TFontPitch)     // property
	Charset() TFontCharSet          // property
	SetCharset(AValue TFontCharSet) // property
	Color() TColor                  // property
	SetColor(AValue TColor)         // property
	Size() int32                    // property
	SetSize(AValue int32)           // property
	Style() TFontStyles             // property
	SetStyle(AValue TFontStyles)    // property
	Height() int32                  // property
	SetHeight(AValue int32)         // property
}

// TTextAttributes Parent: TPersistent
type TTextAttributes struct {
	TPersistent
}

func NewTextAttributes(AOwner IRichMemo, AType TAttributeType) ITextAttributes {
	r1 := extAttributesImportAPI().SysCallN(3, GetObjectUintptr(AOwner), uintptr(AType))
	return AsTextAttributes(r1)
}

func (m *TTextAttributes) Name() string {
	r1 := extAttributesImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TTextAttributes) SetName(AValue string) {
	extAttributesImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TTextAttributes) Pitch() TFontPitch {
	r1 := extAttributesImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TFontPitch(r1)
}

func (m *TTextAttributes) SetPitch(AValue TFontPitch) {
	extAttributesImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Charset() TFontCharSet {
	r1 := extAttributesImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TFontCharSet(r1)
}

func (m *TTextAttributes) SetCharset(AValue TFontCharSet) {
	extAttributesImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Color() TColor {
	r1 := extAttributesImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TTextAttributes) SetColor(AValue TColor) {
	extAttributesImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Size() int32 {
	r1 := extAttributesImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTextAttributes) SetSize(AValue int32) {
	extAttributesImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Style() TFontStyles {
	r1 := extAttributesImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TFontStyles(r1)
}

func (m *TTextAttributes) SetStyle(AValue TFontStyles) {
	extAttributesImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Height() int32 {
	r1 := extAttributesImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTextAttributes) SetHeight(AValue int32) {
	extAttributesImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func TextAttributesClass() TClass {
	ret := extAttributesImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	extAttributesImport       *imports.Imports = nil
	extAttributesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TextAttributes_Charset", 0),
		/*1*/ imports.NewTable("TextAttributes_Class", 0),
		/*2*/ imports.NewTable("TextAttributes_Color", 0),
		/*3*/ imports.NewTable("TextAttributes_Create", 0),
		/*4*/ imports.NewTable("TextAttributes_Height", 0),
		/*5*/ imports.NewTable("TextAttributes_Name", 0),
		/*6*/ imports.NewTable("TextAttributes_Pitch", 0),
		/*7*/ imports.NewTable("TextAttributes_Size", 0),
		/*8*/ imports.NewTable("TextAttributes_Style", 0),
	}
)

func extAttributesImportAPI() *imports.Imports {
	if extAttributesImport == nil {
		extAttributesImport = NewDefaultImports()
		extAttributesImport.SetImportTable(extAttributesImportTables)
		extAttributesImportTables = nil
	}
	return extAttributesImport
}
