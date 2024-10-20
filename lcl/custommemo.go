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

// ICustomMemo Parent: ICustomEdit
type ICustomMemo interface {
	ICustomEdit
	Lines() IStrings                        // property
	SetLines(AValue IStrings)               // property
	HorzScrollBar() IMemoScrollBar          // property
	SetHorzScrollBar(AValue IMemoScrollBar) // property
	VertScrollBar() IMemoScrollBar          // property
	SetVertScrollBar(AValue IMemoScrollBar) // property
	ScrollBars() TScrollStyle               // property
	SetScrollBars(AValue TScrollStyle)      // property
	WantReturns() bool                      // property
	SetWantReturns(AValue bool)             // property
	WantTabs() bool                         // property
	SetWantTabs(AValue bool)                // property
	WordWrap() bool                         // property
	SetWordWrap(AValue bool)                // property
	Append(AValue string)                   // procedure
}

// TCustomMemo Parent: TCustomEdit
type TCustomMemo struct {
	TCustomEdit
}

func NewCustomMemo(AOwner IComponent) ICustomMemo {
	r1 := customMemoImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsCustomMemo(r1)
}

func (m *TCustomMemo) Lines() IStrings {
	r1 := customMemoImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomMemo) SetLines(AValue IStrings) {
	customMemoImportAPI().SysCallN(4, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomMemo) HorzScrollBar() IMemoScrollBar {
	r1 := customMemoImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsMemoScrollBar(r1)
}

func (m *TCustomMemo) SetHorzScrollBar(AValue IMemoScrollBar) {
	customMemoImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomMemo) VertScrollBar() IMemoScrollBar {
	r1 := customMemoImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsMemoScrollBar(r1)
}

func (m *TCustomMemo) SetVertScrollBar(AValue IMemoScrollBar) {
	customMemoImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomMemo) ScrollBars() TScrollStyle {
	r1 := customMemoImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TCustomMemo) SetScrollBars(AValue TScrollStyle) {
	customMemoImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomMemo) WantReturns() bool {
	r1 := customMemoImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMemo) SetWantReturns(AValue bool) {
	customMemoImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomMemo) WantTabs() bool {
	r1 := customMemoImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMemo) SetWantTabs(AValue bool) {
	customMemoImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomMemo) WordWrap() bool {
	r1 := customMemoImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMemo) SetWordWrap(AValue bool) {
	customMemoImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func CustomMemoClass() TClass {
	ret := customMemoImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomMemo) Append(AValue string) {
	customMemoImportAPI().SysCallN(0, m.Instance(), PascalStr(AValue))
}

var (
	customMemoImport       *imports.Imports = nil
	customMemoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomMemo_Append", 0),
		/*1*/ imports.NewTable("CustomMemo_Class", 0),
		/*2*/ imports.NewTable("CustomMemo_Create", 0),
		/*3*/ imports.NewTable("CustomMemo_HorzScrollBar", 0),
		/*4*/ imports.NewTable("CustomMemo_Lines", 0),
		/*5*/ imports.NewTable("CustomMemo_ScrollBars", 0),
		/*6*/ imports.NewTable("CustomMemo_VertScrollBar", 0),
		/*7*/ imports.NewTable("CustomMemo_WantReturns", 0),
		/*8*/ imports.NewTable("CustomMemo_WantTabs", 0),
		/*9*/ imports.NewTable("CustomMemo_WordWrap", 0),
	}
)

func customMemoImportAPI() *imports.Imports {
	if customMemoImport == nil {
		customMemoImport = NewDefaultImports()
		customMemoImport.SetImportTable(customMemoImportTables)
		customMemoImportTables = nil
	}
	return customMemoImport
}
