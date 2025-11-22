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
)

// IStatusPanels Parent: ICollection
type IStatusPanels interface {
	ICollection
	AddToStatusPanel() IStatusPanel                               // function
	ItemsWithIntToStatusPanel(index int32) IStatusPanel           // property Items Getter
	SetItemsWithIntToStatusPanel(index int32, value IStatusPanel) // property Items Setter
	StatusBar() IStatusBar                                        // property StatusBar Getter
}

type TStatusPanels struct {
	TCollection
}

func (m *TStatusPanels) AddToStatusPanel() IStatusPanel {
	if !m.IsValid() {
		return nil
	}
	r := statusPanelsAPI().SysCallN(1, m.Instance())
	return AsStatusPanel(r)
}

func (m *TStatusPanels) ItemsWithIntToStatusPanel(index int32) IStatusPanel {
	if !m.IsValid() {
		return nil
	}
	r := statusPanelsAPI().SysCallN(2, 0, m.Instance(), uintptr(index))
	return AsStatusPanel(r)
}

func (m *TStatusPanels) SetItemsWithIntToStatusPanel(index int32, value IStatusPanel) {
	if !m.IsValid() {
		return
	}
	statusPanelsAPI().SysCallN(2, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TStatusPanels) StatusBar() IStatusBar {
	if !m.IsValid() {
		return nil
	}
	r := statusPanelsAPI().SysCallN(3, m.Instance())
	return AsStatusBar(r)
}

// NewStatusPanels class constructor
func NewStatusPanels(statusBar IStatusBar) IStatusPanels {
	r := statusPanelsAPI().SysCallN(0, base.GetObjectUintptr(statusBar))
	return AsStatusPanels(r)
}

var (
	statusPanelsOnce   base.Once
	statusPanelsImport *imports.Imports = nil
)

func statusPanelsAPI() *imports.Imports {
	statusPanelsOnce.Do(func() {
		statusPanelsImport = api.NewDefaultImports()
		statusPanelsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStatusPanels_Create", 0), // constructor NewStatusPanels
			/* 1 */ imports.NewTable("TStatusPanels_AddToStatusPanel", 0), // function AddToStatusPanel
			/* 2 */ imports.NewTable("TStatusPanels_ItemsWithIntToStatusPanel", 0), // property ItemsWithIntToStatusPanel
			/* 3 */ imports.NewTable("TStatusPanels_StatusBar", 0), // property StatusBar
		}
	})
	return statusPanelsImport
}
