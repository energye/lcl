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

// ICustomPage Parent: IWinControl
type ICustomPage interface {
	IWinControl
	PageIndex() int32                 // property
	SetPageIndex(AValue int32)        // property
	TabVisible() bool                 // property
	SetTabVisible(AValue bool)        // property
	ImageIndex() TImageIndex          // property
	SetImageIndex(AValue TImageIndex) // property
	CanTab() bool                     // function
	VisibleIndex() int32              // function
	SetOnHide(fn TNotifyEvent)        // property event
	SetOnShow(fn TNotifyEvent)        // property event
}

// TCustomPage Parent: TWinControl
type TCustomPage struct {
	TWinControl
	hidePtr uintptr
	showPtr uintptr
}

func NewCustomPage(TheOwner IComponent) ICustomPage {
	r1 := customPageImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsCustomPage(r1)
}

func (m *TCustomPage) PageIndex() int32 {
	r1 := customPageImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPage) SetPageIndex(AValue int32) {
	customPageImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPage) TabVisible() bool {
	r1 := customPageImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPage) SetTabVisible(AValue bool) {
	customPageImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPage) ImageIndex() TImageIndex {
	r1 := customPageImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomPage) SetImageIndex(AValue TImageIndex) {
	customPageImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPage) CanTab() bool {
	r1 := customPageImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCustomPage) VisibleIndex() int32 {
	r1 := customPageImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func CustomPageClass() TClass {
	ret := customPageImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomPage) SetOnHide(fn TNotifyEvent) {
	if m.hidePtr != 0 {
		RemoveEventElement(m.hidePtr)
	}
	m.hidePtr = MakeEventDataPtr(fn)
	customPageImportAPI().SysCallN(5, m.Instance(), m.hidePtr)
}

func (m *TCustomPage) SetOnShow(fn TNotifyEvent) {
	if m.showPtr != 0 {
		RemoveEventElement(m.showPtr)
	}
	m.showPtr = MakeEventDataPtr(fn)
	customPageImportAPI().SysCallN(6, m.Instance(), m.showPtr)
}

var (
	customPageImport       *imports.Imports = nil
	customPageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPage_CanTab", 0),
		/*1*/ imports.NewTable("CustomPage_Class", 0),
		/*2*/ imports.NewTable("CustomPage_Create", 0),
		/*3*/ imports.NewTable("CustomPage_ImageIndex", 0),
		/*4*/ imports.NewTable("CustomPage_PageIndex", 0),
		/*5*/ imports.NewTable("CustomPage_SetOnHide", 0),
		/*6*/ imports.NewTable("CustomPage_SetOnShow", 0),
		/*7*/ imports.NewTable("CustomPage_TabVisible", 0),
		/*8*/ imports.NewTable("CustomPage_VisibleIndex", 0),
	}
)

func customPageImportAPI() *imports.Imports {
	if customPageImport == nil {
		customPageImport = NewDefaultImports()
		customPageImport.SetImportTable(customPageImportTables)
		customPageImportTables = nil
	}
	return customPageImport
}
