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

// ILazDockPages Parent: ICustomTabControl
type ILazDockPages interface {
	ICustomTabControl
	PageWithIntToLazDockPage(index int32) ILazDockPage // property Page Getter
	ActivePageComponent() ILazDockPage                 // property ActivePageComponent Getter
	SetActivePageComponent(value ILazDockPage)         // property ActivePageComponent Setter
}

type TLazDockPages struct {
	TCustomTabControl
}

func (m *TLazDockPages) PageWithIntToLazDockPage(index int32) ILazDockPage {
	if !m.IsValid() {
		return nil
	}
	r := lazDockPagesAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsLazDockPage(r)
}

func (m *TLazDockPages) ActivePageComponent() ILazDockPage {
	if !m.IsValid() {
		return nil
	}
	r := lazDockPagesAPI().SysCallN(2, 0, m.Instance())
	return AsLazDockPage(r)
}

func (m *TLazDockPages) SetActivePageComponent(value ILazDockPage) {
	if !m.IsValid() {
		return
	}
	lazDockPagesAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewLazDockPages class constructor
func NewLazDockPages(theOwner IComponent) ILazDockPages {
	r := lazDockPagesAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsLazDockPages(r)
}

func TLazDockPagesClass() types.TClass {
	r := lazDockPagesAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	lazDockPagesOnce   base.Once
	lazDockPagesImport *imports.Imports = nil
)

func lazDockPagesAPI() *imports.Imports {
	lazDockPagesOnce.Do(func() {
		lazDockPagesImport = api.NewDefaultImports()
		lazDockPagesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazDockPages_Create", 0), // constructor NewLazDockPages
			/* 1 */ imports.NewTable("TLazDockPages_PageWithIntToLazDockPage", 0), // property PageWithIntToLazDockPage
			/* 2 */ imports.NewTable("TLazDockPages_ActivePageComponent", 0), // property ActivePageComponent
			/* 3 */ imports.NewTable("TLazDockPages_TClass", 0), // function TLazDockPagesClass
		}
	})
	return lazDockPagesImport
}
