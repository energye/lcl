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

// ILCLComponent Parent: IComponent
type ILCLComponent interface {
	IComponent
	LCLRefCount() int32                         // property
	RemoveAllHandlersOfObject(AnObject IObject) // procedure
	IncLCLRefCount()                            // procedure
	DecLCLRefCount()                            // procedure
}

// TLCLComponent Parent: TComponent
type TLCLComponent struct {
	TComponent
}

func NewLCLComponent(TheOwner IComponent) ILCLComponent {
	r1 := lCLComponentImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsLCLComponent(r1)
}

func (m *TLCLComponent) LCLRefCount() int32 {
	r1 := lCLComponentImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func LCLComponentClass() TClass {
	ret := lCLComponentImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TLCLComponent) RemoveAllHandlersOfObject(AnObject IObject) {
	lCLComponentImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(AnObject))
}

func (m *TLCLComponent) IncLCLRefCount() {
	lCLComponentImportAPI().SysCallN(3, m.Instance())
}

func (m *TLCLComponent) DecLCLRefCount() {
	lCLComponentImportAPI().SysCallN(2, m.Instance())
}

var (
	lCLComponentImport       *imports.Imports = nil
	lCLComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LCLComponent_Class", 0),
		/*1*/ imports.NewTable("LCLComponent_Create", 0),
		/*2*/ imports.NewTable("LCLComponent_DecLCLRefCount", 0),
		/*3*/ imports.NewTable("LCLComponent_IncLCLRefCount", 0),
		/*4*/ imports.NewTable("LCLComponent_LCLRefCount", 0),
		/*5*/ imports.NewTable("LCLComponent_RemoveAllHandlersOfObject", 0),
	}
)

func lCLComponentImportAPI() *imports.Imports {
	if lCLComponentImport == nil {
		lCLComponentImport = NewDefaultImports()
		lCLComponentImport.SetImportTable(lCLComponentImportTables)
		lCLComponentImportTables = nil
	}
	return lCLComponentImport
}
