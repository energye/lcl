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

// IMainMenu Parent: IMenu
type IMainMenu interface {
	IMenu
	Height() int32                   // property
	WindowHandle() HWND              // property
	SetWindowHandle(AValue HWND)     // property
	Merge(Menu IMainMenu)            // procedure
	Unmerge(Menu IMainMenu)          // procedure
	SetOnChange(fn TMenuChangeEvent) // property event
}

// TMainMenu Parent: TMenu
type TMainMenu struct {
	TMenu
	changePtr uintptr
}

func NewMainMenu(AOwner IComponent) IMainMenu {
	r1 := mainMenuImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsMainMenu(r1)
}

func (m *TMainMenu) Height() int32 {
	r1 := mainMenuImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TMainMenu) WindowHandle() HWND {
	r1 := mainMenuImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TMainMenu) SetWindowHandle(AValue HWND) {
	mainMenuImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func MainMenuClass() TClass {
	ret := mainMenuImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TMainMenu) Merge(Menu IMainMenu) {
	mainMenuImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(Menu))
}

func (m *TMainMenu) Unmerge(Menu IMainMenu) {
	mainMenuImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Menu))
}

func (m *TMainMenu) SetOnChange(fn TMenuChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	mainMenuImportAPI().SysCallN(4, m.Instance(), m.changePtr)
}

var (
	mainMenuImport       *imports.Imports = nil
	mainMenuImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MainMenu_Class", 0),
		/*1*/ imports.NewTable("MainMenu_Create", 0),
		/*2*/ imports.NewTable("MainMenu_Height", 0),
		/*3*/ imports.NewTable("MainMenu_Merge", 0),
		/*4*/ imports.NewTable("MainMenu_SetOnChange", 0),
		/*5*/ imports.NewTable("MainMenu_Unmerge", 0),
		/*6*/ imports.NewTable("MainMenu_WindowHandle", 0),
	}
)

func mainMenuImportAPI() *imports.Imports {
	if mainMenuImport == nil {
		mainMenuImport = NewDefaultImports()
		mainMenuImport.SetImportTable(mainMenuImportTables)
		mainMenuImportTables = nil
	}
	return mainMenuImport
}
