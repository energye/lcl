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

// IEngForm Parent: IForm
type IEngForm interface {
	IForm
	ScaleForPPI(newPPI int32)                     // procedure
	ScaleForCurrentDpi()                          // procedure
	GoDeactivate()                                // procedure
	GoActivate()                                  // procedure
	InheritedWndProc(theMessage *types.TLMessage) // procedure
	EnabledMaximize(value bool)                   // procedure
	EnabledMinimize(value bool)                   // procedure
	EnabledSystemMenu(value bool)                 // procedure
	ScreenCenter()                                // procedure
	WorkAreaCenter()                              // procedure
	GoPtr() uintptr                               // property GoPtr Getter
	SetGoPtr(value uintptr)                       // property GoPtr Setter
	SetOnWndProc(fn TWndMethod)                   // property event
}

type TEngForm struct {
	TForm
}

func (m *TEngForm) ScaleForPPI(newPPI int32) {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(3, m.Instance(), uintptr(newPPI))
}

func (m *TEngForm) ScaleForCurrentDpi() {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(4, m.Instance())
}

func (m *TEngForm) GoDeactivate() {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(5, m.Instance())
}

func (m *TEngForm) GoActivate() {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(6, m.Instance())
}

func (m *TEngForm) InheritedWndProc(theMessage *types.TLMessage) {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(theMessage)))
}

func (m *TEngForm) EnabledMaximize(value bool) {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(8, m.Instance(), api.PasBool(value))
}

func (m *TEngForm) EnabledMinimize(value bool) {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(9, m.Instance(), api.PasBool(value))
}

func (m *TEngForm) EnabledSystemMenu(value bool) {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(10, m.Instance(), api.PasBool(value))
}

func (m *TEngForm) ScreenCenter() {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(11, m.Instance())
}

func (m *TEngForm) WorkAreaCenter() {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(12, m.Instance())
}

func (m *TEngForm) GoPtr() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := engFormAPI().SysCallN(13, 0, m.Instance())
	return uintptr(r)
}

func (m *TEngForm) SetGoPtr(value uintptr) {
	if !m.IsValid() {
		return
	}
	engFormAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TEngForm) SetOnWndProc(fn TWndMethod) {
	if !m.IsValid() {
		return
	}
	cb := makeTWndMethod(fn)
	base.SetEvent(m, 14, engFormAPI(), api.MakeEventDataPtr(cb))
}

// NewEngForm class constructor
func NewEngForm(owner IComponent) IEngForm {
	r := engFormAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEngForm(r)
}

// NewEngFormNew class constructor
func NewEngFormNew(owner IComponent, num int32) IEngForm {
	r := engFormAPI().SysCallN(1, base.GetObjectUintptr(owner), uintptr(num))
	return AsEngForm(r)
}

// NewEngFormFromClassName class constructor
func NewEngFormFromClassName(theOwner IComponent, className string) IEngForm {
	r := engFormAPI().SysCallN(2, base.GetObjectUintptr(theOwner), api.PasStr(className))
	return AsEngForm(r)
}

func TEngFormClass() types.TClass {
	r := engFormAPI().SysCallN(15)
	return types.TClass(r)
}

var (
	engFormOnce   base.Once
	engFormImport *imports.Imports = nil
)

func engFormAPI() *imports.Imports {
	engFormOnce.Do(func() {
		engFormImport = api.NewDefaultImports()
		engFormImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEngForm_Create", 0), // constructor NewEngForm
			/* 1 */ imports.NewTable("TEngForm_CreateNew", 0), // constructor NewEngFormNew
			/* 2 */ imports.NewTable("TEngForm_CreateFromClassName", 0), // constructor NewEngFormFromClassName
			/* 3 */ imports.NewTable("TEngForm_ScaleForPPI", 0), // procedure ScaleForPPI
			/* 4 */ imports.NewTable("TEngForm_ScaleForCurrentDpi", 0), // procedure ScaleForCurrentDpi
			/* 5 */ imports.NewTable("TEngForm_GoDeactivate", 0), // procedure GoDeactivate
			/* 6 */ imports.NewTable("TEngForm_GoActivate", 0), // procedure GoActivate
			/* 7 */ imports.NewTable("TEngForm_InheritedWndProc", 0), // procedure InheritedWndProc
			/* 8 */ imports.NewTable("TEngForm_EnabledMaximize", 0), // procedure EnabledMaximize
			/* 9 */ imports.NewTable("TEngForm_EnabledMinimize", 0), // procedure EnabledMinimize
			/* 10 */ imports.NewTable("TEngForm_EnabledSystemMenu", 0), // procedure EnabledSystemMenu
			/* 11 */ imports.NewTable("TEngForm_ScreenCenter", 0), // procedure ScreenCenter
			/* 12 */ imports.NewTable("TEngForm_WorkAreaCenter", 0), // procedure WorkAreaCenter
			/* 13 */ imports.NewTable("TEngForm_GoPtr", 0), // property GoPtr
			/* 14 */ imports.NewTable("TEngForm_OnWndProc", 0), // event OnWndProc
			/* 15 */ imports.NewTable("TEngForm_TClass", 0), // function TEngFormClass
		}
	})
	return engFormImport
}
