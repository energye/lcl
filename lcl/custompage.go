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

// ICustomPage Parent: IWinControl
type ICustomPage interface {
	IWinControl
	CanTab() bool              // function
	VisibleIndex() int32       // function
	PageIndex() int32          // property PageIndex Getter
	SetPageIndex(value int32)  // property PageIndex Setter
	TabVisible() bool          // property TabVisible Getter
	SetTabVisible(value bool)  // property TabVisible Setter
	ImageIndex() int32         // property ImageIndex Getter
	SetImageIndex(value int32) // property ImageIndex Setter
	SetOnHide(fn TNotifyEvent) // property event
	SetOnShow(fn TNotifyEvent) // property event
}

type TCustomPage struct {
	TWinControl
}

func (m *TCustomPage) CanTab() bool {
	if !m.IsValid() {
		return false
	}
	r := customPageAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomPage) VisibleIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPageAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCustomPage) PageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPageAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPage) SetPageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customPageAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPage) TabVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := customPageAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomPage) SetTabVisible(value bool) {
	if !m.IsValid() {
		return
	}
	customPageAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomPage) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPageAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPage) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customPageAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPage) SetOnHide(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, customPageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomPage) SetOnShow(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, customPageAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomPage class constructor
func NewCustomPage(theOwner IComponent) ICustomPage {
	r := customPageAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomPage(r)
}

func TCustomPageClass() types.TClass {
	r := customPageAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	customPageOnce   base.Once
	customPageImport *imports.Imports = nil
)

func customPageAPI() *imports.Imports {
	customPageOnce.Do(func() {
		customPageImport = api.NewDefaultImports()
		customPageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPage_Create", 0), // constructor NewCustomPage
			/* 1 */ imports.NewTable("TCustomPage_CanTab", 0), // function CanTab
			/* 2 */ imports.NewTable("TCustomPage_VisibleIndex", 0), // function VisibleIndex
			/* 3 */ imports.NewTable("TCustomPage_PageIndex", 0), // property PageIndex
			/* 4 */ imports.NewTable("TCustomPage_TabVisible", 0), // property TabVisible
			/* 5 */ imports.NewTable("TCustomPage_ImageIndex", 0), // property ImageIndex
			/* 6 */ imports.NewTable("TCustomPage_OnHide", 0), // event OnHide
			/* 7 */ imports.NewTable("TCustomPage_OnShow", 0), // event OnShow
			/* 8 */ imports.NewTable("TCustomPage_TClass", 0), // function TCustomPageClass
		}
	})
	return customPageImport
}
