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

// IStatusPanel Parent: ICollectionItem
type IStatusPanel interface {
	ICollectionItem
	Alignment() TAlignment             // property
	SetAlignment(AValue TAlignment)    // property
	Bevel() TStatusPanelBevel          // property
	SetBevel(AValue TStatusPanelBevel) // property
	BidiMode() TBiDiMode               // property
	SetBidiMode(AValue TBiDiMode)      // property
	ParentBiDiMode() bool              // property
	SetParentBiDiMode(AValue bool)     // property
	Style() TStatusPanelStyle          // property
	SetStyle(AValue TStatusPanelStyle) // property
	Text() string                      // property
	SetText(AValue string)             // property
	Width() int32                      // property
	SetWidth(AValue int32)             // property
	StatusBar() IStatusBar             // function
}

// TStatusPanel Parent: TCollectionItem
type TStatusPanel struct {
	TCollectionItem
}

func NewStatusPanel(ACollection ICollection) IStatusPanel {
	r1 := statusPanelImportAPI().SysCallN(4, GetObjectUintptr(ACollection))
	return AsStatusPanel(r1)
}

func (m *TStatusPanel) Alignment() TAlignment {
	r1 := statusPanelImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TStatusPanel) SetAlignment(AValue TAlignment) {
	statusPanelImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) Bevel() TStatusPanelBevel {
	r1 := statusPanelImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TStatusPanelBevel(r1)
}

func (m *TStatusPanel) SetBevel(AValue TStatusPanelBevel) {
	statusPanelImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) BidiMode() TBiDiMode {
	r1 := statusPanelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TStatusPanel) SetBidiMode(AValue TBiDiMode) {
	statusPanelImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) ParentBiDiMode() bool {
	r1 := statusPanelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusPanel) SetParentBiDiMode(AValue bool) {
	statusPanelImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusPanel) Style() TStatusPanelStyle {
	r1 := statusPanelImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TStatusPanelStyle(r1)
}

func (m *TStatusPanel) SetStyle(AValue TStatusPanelStyle) {
	statusPanelImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) Text() string {
	r1 := statusPanelImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStatusPanel) SetText(AValue string) {
	statusPanelImportAPI().SysCallN(8, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStatusPanel) Width() int32 {
	r1 := statusPanelImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TStatusPanel) SetWidth(AValue int32) {
	statusPanelImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) StatusBar() IStatusBar {
	r1 := statusPanelImportAPI().SysCallN(6, m.Instance())
	return AsStatusBar(r1)
}

func StatusPanelClass() TClass {
	ret := statusPanelImportAPI().SysCallN(3)
	return TClass(ret)
}

var (
	statusPanelImport       *imports.Imports = nil
	statusPanelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StatusPanel_Alignment", 0),
		/*1*/ imports.NewTable("StatusPanel_Bevel", 0),
		/*2*/ imports.NewTable("StatusPanel_BidiMode", 0),
		/*3*/ imports.NewTable("StatusPanel_Class", 0),
		/*4*/ imports.NewTable("StatusPanel_Create", 0),
		/*5*/ imports.NewTable("StatusPanel_ParentBiDiMode", 0),
		/*6*/ imports.NewTable("StatusPanel_StatusBar", 0),
		/*7*/ imports.NewTable("StatusPanel_Style", 0),
		/*8*/ imports.NewTable("StatusPanel_Text", 0),
		/*9*/ imports.NewTable("StatusPanel_Width", 0),
	}
)

func statusPanelImportAPI() *imports.Imports {
	if statusPanelImport == nil {
		statusPanelImport = NewDefaultImports()
		statusPanelImport.SetImportTable(statusPanelImportTables)
		statusPanelImportTables = nil
	}
	return statusPanelImport
}
