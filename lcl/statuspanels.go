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

// IStatusPanels Parent: ICollection
type IStatusPanels interface {
	ICollection
	ItemsForStatusPanel(Index int32) IStatusPanel            // property
	SetItemsForStatusPanel(Index int32, AValue IStatusPanel) // property
	StatusBar() IStatusBar                                   // property
	AddForStatusPanel() IStatusPanel                         // function
}

// TStatusPanels Parent: TCollection
type TStatusPanels struct {
	TCollection
}

func NewStatusPanels(AStatusBar IStatusBar) IStatusPanels {
	r1 := statusPanelsImportAPI().SysCallN(2, GetObjectUintptr(AStatusBar))
	return AsStatusPanels(r1)
}

func (m *TStatusPanels) ItemsForStatusPanel(Index int32) IStatusPanel {
	r1 := statusPanelsImportAPI().SysCallN(3, 0, m.Instance(), uintptr(Index))
	return AsStatusPanel(r1)
}

func (m *TStatusPanels) SetItemsForStatusPanel(Index int32, AValue IStatusPanel) {
	statusPanelsImportAPI().SysCallN(3, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TStatusPanels) StatusBar() IStatusBar {
	r1 := statusPanelsImportAPI().SysCallN(4, m.Instance())
	return AsStatusBar(r1)
}

func (m *TStatusPanels) AddForStatusPanel() IStatusPanel {
	r1 := statusPanelsImportAPI().SysCallN(0, m.Instance())
	return AsStatusPanel(r1)
}

func StatusPanelsClass() TClass {
	ret := statusPanelsImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	statusPanelsImport       *imports.Imports = nil
	statusPanelsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StatusPanels_AddForStatusPanel", 0),
		/*1*/ imports.NewTable("StatusPanels_Class", 0),
		/*2*/ imports.NewTable("StatusPanels_Create", 0),
		/*3*/ imports.NewTable("StatusPanels_ItemsForStatusPanel", 0),
		/*4*/ imports.NewTable("StatusPanels_StatusBar", 0),
	}
)

func statusPanelsImportAPI() *imports.Imports {
	if statusPanelsImport == nil {
		statusPanelsImport = NewDefaultImports()
		statusPanelsImport.SetImportTable(statusPanelsImportTables)
		statusPanelsImportTables = nil
	}
	return statusPanelsImport
}
