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

// ILazDockZone Parent: IDockZone
type ILazDockZone interface {
	IDockZone
	GetCaption() string                 // function
	GetParentControl() IWinControl      // function
	FreeSubComponents()                 // procedure
	Splitter() ILazDockSplitter         // property Splitter Getter
	SetSplitter(value ILazDockSplitter) // property Splitter Setter
	Pages() ILazDockPages               // property Pages Getter
	SetPages(value ILazDockPages)       // property Pages Setter
	Page() ILazDockPage                 // property Page Getter
	SetPage(value ILazDockPage)         // property Page Setter
}

type TLazDockZone struct {
	TDockZone
}

func (m *TLazDockZone) GetCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := lazDockZoneAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TLazDockZone) GetParentControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := lazDockZoneAPI().SysCallN(2, m.Instance())
	return AsWinControl(r)
}

func (m *TLazDockZone) FreeSubComponents() {
	if !m.IsValid() {
		return
	}
	lazDockZoneAPI().SysCallN(3, m.Instance())
}

func (m *TLazDockZone) Splitter() ILazDockSplitter {
	if !m.IsValid() {
		return nil
	}
	r := lazDockZoneAPI().SysCallN(4, 0, m.Instance())
	return AsLazDockSplitter(r)
}

func (m *TLazDockZone) SetSplitter(value ILazDockSplitter) {
	if !m.IsValid() {
		return
	}
	lazDockZoneAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazDockZone) Pages() ILazDockPages {
	if !m.IsValid() {
		return nil
	}
	r := lazDockZoneAPI().SysCallN(5, 0, m.Instance())
	return AsLazDockPages(r)
}

func (m *TLazDockZone) SetPages(value ILazDockPages) {
	if !m.IsValid() {
		return
	}
	lazDockZoneAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazDockZone) Page() ILazDockPage {
	if !m.IsValid() {
		return nil
	}
	r := lazDockZoneAPI().SysCallN(6, 0, m.Instance())
	return AsLazDockPage(r)
}

func (m *TLazDockZone) SetPage(value ILazDockPage) {
	if !m.IsValid() {
		return
	}
	lazDockZoneAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewLazDockZone class constructor
func NewLazDockZone(theTree IDockTree, theChildControl IControl) ILazDockZone {
	r := lazDockZoneAPI().SysCallN(0, base.GetObjectUintptr(theTree), base.GetObjectUintptr(theChildControl))
	return AsLazDockZone(r)
}

var (
	lazDockZoneOnce   base.Once
	lazDockZoneImport *imports.Imports = nil
)

func lazDockZoneAPI() *imports.Imports {
	lazDockZoneOnce.Do(func() {
		lazDockZoneImport = api.NewDefaultImports()
		lazDockZoneImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazDockZone_Create", 0), // constructor NewLazDockZone
			/* 1 */ imports.NewTable("TLazDockZone_GetCaption", 0), // function GetCaption
			/* 2 */ imports.NewTable("TLazDockZone_GetParentControl", 0), // function GetParentControl
			/* 3 */ imports.NewTable("TLazDockZone_FreeSubComponents", 0), // procedure FreeSubComponents
			/* 4 */ imports.NewTable("TLazDockZone_Splitter", 0), // property Splitter
			/* 5 */ imports.NewTable("TLazDockZone_Pages", 0), // property Pages
			/* 6 */ imports.NewTable("TLazDockZone_Page", 0), // property Page
		}
	})
	return lazDockZoneImport
}
