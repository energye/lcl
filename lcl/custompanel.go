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

// ICustomPanel Parent: ICustomControl
type ICustomPanel interface {
	ICustomControl
	Alignment() TAlignment            // property
	SetAlignment(AValue TAlignment)   // property
	BevelColor() TColor               // property
	SetBevelColor(AValue TColor)      // property
	BevelInner() TPanelBevel          // property
	SetBevelInner(AValue TPanelBevel) // property
	BevelOuter() TPanelBevel          // property
	SetBevelOuter(AValue TPanelBevel) // property
	BevelWidth() TBevelWidth          // property
	SetBevelWidth(AValue TBevelWidth) // property
	FullRepaint() bool                // property
	SetFullRepaint(AValue bool)       // property
	ParentBackground() bool           // property
	SetParentBackground(AValue bool)  // property
	ParentColor() bool                // property
	SetParentColor(AValue bool)       // property
}

// TCustomPanel Parent: TCustomControl
type TCustomPanel struct {
	TCustomControl
}

func NewCustomPanel(TheOwner IComponent) ICustomPanel {
	r1 := customPanelImportAPI().SysCallN(6, GetObjectUintptr(TheOwner))
	return AsCustomPanel(r1)
}

func (m *TCustomPanel) Alignment() TAlignment {
	r1 := customPanelImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomPanel) SetAlignment(AValue TAlignment) {
	customPanelImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelColor() TColor {
	r1 := customPanelImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomPanel) SetBevelColor(AValue TColor) {
	customPanelImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelInner() TPanelBevel {
	r1 := customPanelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TPanelBevel(r1)
}

func (m *TCustomPanel) SetBevelInner(AValue TPanelBevel) {
	customPanelImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelOuter() TPanelBevel {
	r1 := customPanelImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TPanelBevel(r1)
}

func (m *TCustomPanel) SetBevelOuter(AValue TPanelBevel) {
	customPanelImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelWidth() TBevelWidth {
	r1 := customPanelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TBevelWidth(r1)
}

func (m *TCustomPanel) SetBevelWidth(AValue TBevelWidth) {
	customPanelImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) FullRepaint() bool {
	r1 := customPanelImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPanel) SetFullRepaint(AValue bool) {
	customPanelImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPanel) ParentBackground() bool {
	r1 := customPanelImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPanel) SetParentBackground(AValue bool) {
	customPanelImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPanel) ParentColor() bool {
	r1 := customPanelImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPanel) SetParentColor(AValue bool) {
	customPanelImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func CustomPanelClass() TClass {
	ret := customPanelImportAPI().SysCallN(5)
	return TClass(ret)
}

var (
	customPanelImport       *imports.Imports = nil
	customPanelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPanel_Alignment", 0),
		/*1*/ imports.NewTable("CustomPanel_BevelColor", 0),
		/*2*/ imports.NewTable("CustomPanel_BevelInner", 0),
		/*3*/ imports.NewTable("CustomPanel_BevelOuter", 0),
		/*4*/ imports.NewTable("CustomPanel_BevelWidth", 0),
		/*5*/ imports.NewTable("CustomPanel_Class", 0),
		/*6*/ imports.NewTable("CustomPanel_Create", 0),
		/*7*/ imports.NewTable("CustomPanel_FullRepaint", 0),
		/*8*/ imports.NewTable("CustomPanel_ParentBackground", 0),
		/*9*/ imports.NewTable("CustomPanel_ParentColor", 0),
	}
)

func customPanelImportAPI() *imports.Imports {
	if customPanelImport == nil {
		customPanelImport = NewDefaultImports()
		customPanelImport.SetImportTable(customPanelImportTables)
		customPanelImportTables = nil
	}
	return customPanelImport
}
