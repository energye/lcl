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

// IMainMenu Parent: IMenu
type IMainMenu interface {
	IMenu
	Merge(menu IMainMenu)             // procedure
	Unmerge(menu IMainMenu)           // procedure
	Height() int32                    // property Height Getter
	WindowHandle() types.HWND         // property WindowHandle Getter
	SetWindowHandle(value types.HWND) // property WindowHandle Setter
	SetOnChange(fn TMenuChangeEvent)  // property event
}

type TMainMenu struct {
	TMenu
}

func (m *TMainMenu) Merge(menu IMainMenu) {
	if !m.IsValid() {
		return
	}
	mainMenuAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(menu))
}

func (m *TMainMenu) Unmerge(menu IMainMenu) {
	if !m.IsValid() {
		return
	}
	mainMenuAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(menu))
}

func (m *TMainMenu) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := mainMenuAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TMainMenu) WindowHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	r := mainMenuAPI().SysCallN(4, 0, m.Instance())
	return types.HWND(r)
}

func (m *TMainMenu) SetWindowHandle(value types.HWND) {
	if !m.IsValid() {
		return
	}
	mainMenuAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TMainMenu) SetOnChange(fn TMenuChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMenuChangeEvent(fn)
	base.SetEvent(m, 5, mainMenuAPI(), api.MakeEventDataPtr(cb))
}

// NewMainMenu class constructor
func NewMainMenu(owner IComponent) IMainMenu {
	r := mainMenuAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsMainMenu(r)
}

func TMainMenuClass() types.TClass {
	r := mainMenuAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	mainMenuOnce   base.Once
	mainMenuImport *imports.Imports = nil
)

func mainMenuAPI() *imports.Imports {
	mainMenuOnce.Do(func() {
		mainMenuImport = api.NewDefaultImports()
		mainMenuImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMainMenu_Create", 0), // constructor NewMainMenu
			/* 1 */ imports.NewTable("TMainMenu_Merge", 0), // procedure Merge
			/* 2 */ imports.NewTable("TMainMenu_Unmerge", 0), // procedure Unmerge
			/* 3 */ imports.NewTable("TMainMenu_Height", 0), // property Height
			/* 4 */ imports.NewTable("TMainMenu_WindowHandle", 0), // property WindowHandle
			/* 5 */ imports.NewTable("TMainMenu_OnChange", 0), // event OnChange
			/* 6 */ imports.NewTable("TMainMenu_TClass", 0), // function TMainMenuClass
		}
	})
	return mainMenuImport
}
