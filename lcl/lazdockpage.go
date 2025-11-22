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
	"github.com/energye/lcl/types"
)

// ILazDockPage Parent: ICustomPage
type ILazDockPage interface {
	ICustomPage
	DockZone() IDockZone        // property DockZone Getter
	PageControl() ILazDockPages // property PageControl Getter
}

type TLazDockPage struct {
	TCustomPage
}

func (m *TLazDockPage) DockZone() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := lazDockPageAPI().SysCallN(1, m.Instance())
	return AsDockZone(r)
}

func (m *TLazDockPage) PageControl() ILazDockPages {
	if !m.IsValid() {
		return nil
	}
	r := lazDockPageAPI().SysCallN(2, m.Instance())
	return AsLazDockPages(r)
}

// NewLazDockPage class constructor
func NewLazDockPage(theOwner IComponent) ILazDockPage {
	r := lazDockPageAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsLazDockPage(r)
}

func TLazDockPageClass() types.TClass {
	r := lazDockPageAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	lazDockPageOnce   base.Once
	lazDockPageImport *imports.Imports = nil
)

func lazDockPageAPI() *imports.Imports {
	lazDockPageOnce.Do(func() {
		lazDockPageImport = api.NewDefaultImports()
		lazDockPageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazDockPage_Create", 0), // constructor NewLazDockPage
			/* 1 */ imports.NewTable("TLazDockPage_DockZone", 0), // property DockZone
			/* 2 */ imports.NewTable("TLazDockPage_PageControl", 0), // property PageControl
			/* 3 */ imports.NewTable("TLazDockPage_TClass", 0), // function TLazDockPageClass
		}
	})
	return lazDockPageImport
}
