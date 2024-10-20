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

// IParaAttributes Parent: IPersistent
type IParaAttributes interface {
	IPersistent
	Alignment() TAlignment               // property
	SetAlignment(AValue TAlignment)      // property
	FirstIndent() int32                  // property
	SetFirstIndent(AValue int32)         // property
	LeftIndent() int32                   // property
	SetLeftIndent(AValue int32)          // property
	RightIndent() int32                  // property
	SetRightIndent(AValue int32)         // property
	Numbering() TNumberingStyle          // property
	SetNumbering(AValue TNumberingStyle) // property
	TabCount() int32                     // property
	SetTabCount(AValue int32)            // property
	Tab(Index Byte) int32                // property
	SetTab(Index Byte, AValue int32)     // property
}

// TParaAttributes Parent: TPersistent
type TParaAttributes struct {
	TPersistent
}

func NewParaAttributes(AOwner IRichMemo) IParaAttributes {
	r1 := paraAttributesImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsParaAttributes(r1)
}

func (m *TParaAttributes) Alignment() TAlignment {
	r1 := paraAttributesImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TParaAttributes) SetAlignment(AValue TAlignment) {
	paraAttributesImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) FirstIndent() int32 {
	r1 := paraAttributesImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetFirstIndent(AValue int32) {
	paraAttributesImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) LeftIndent() int32 {
	r1 := paraAttributesImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetLeftIndent(AValue int32) {
	paraAttributesImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) RightIndent() int32 {
	r1 := paraAttributesImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetRightIndent(AValue int32) {
	paraAttributesImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) Numbering() TNumberingStyle {
	r1 := paraAttributesImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TNumberingStyle(r1)
}

func (m *TParaAttributes) SetNumbering(AValue TNumberingStyle) {
	paraAttributesImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) TabCount() int32 {
	r1 := paraAttributesImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TParaAttributes) SetTabCount(AValue int32) {
	paraAttributesImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TParaAttributes) Tab(Index Byte) int32 {
	r1 := paraAttributesImportAPI().SysCallN(7, 0, m.Instance(), uintptr(Index))
	return int32(r1)
}

func (m *TParaAttributes) SetTab(Index Byte, AValue int32) {
	paraAttributesImportAPI().SysCallN(7, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func ParaAttributesClass() TClass {
	ret := paraAttributesImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	paraAttributesImport       *imports.Imports = nil
	paraAttributesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ParaAttributes_Alignment", 0),
		/*1*/ imports.NewTable("ParaAttributes_Class", 0),
		/*2*/ imports.NewTable("ParaAttributes_Create", 0),
		/*3*/ imports.NewTable("ParaAttributes_FirstIndent", 0),
		/*4*/ imports.NewTable("ParaAttributes_LeftIndent", 0),
		/*5*/ imports.NewTable("ParaAttributes_Numbering", 0),
		/*6*/ imports.NewTable("ParaAttributes_RightIndent", 0),
		/*7*/ imports.NewTable("ParaAttributes_Tab", 0),
		/*8*/ imports.NewTable("ParaAttributes_TabCount", 0),
	}
)

func paraAttributesImportAPI() *imports.Imports {
	if paraAttributesImport == nil {
		paraAttributesImport = NewDefaultImports()
		paraAttributesImport.SetImportTable(paraAttributesImportTables)
		paraAttributesImportTables = nil
	}
	return paraAttributesImport
}
