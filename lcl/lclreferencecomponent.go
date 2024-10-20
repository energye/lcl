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

// ILCLReferenceComponent Is Abstract Class Parent: ILCLComponent
// A base class for all components having a handle
type ILCLReferenceComponent interface {
	ILCLComponent
	HandleAllocated() bool    // property
	ReferenceAllocated() bool // property
}

// TLCLReferenceComponent Is Abstract Class Parent: TLCLComponent
// A base class for all components having a handle
type TLCLReferenceComponent struct {
	TLCLComponent
}

func (m *TLCLReferenceComponent) HandleAllocated() bool {
	r1 := lCLReferenceComponentImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TLCLReferenceComponent) ReferenceAllocated() bool {
	r1 := lCLReferenceComponentImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func LCLReferenceComponentClass() TClass {
	ret := lCLReferenceComponentImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	lCLReferenceComponentImport       *imports.Imports = nil
	lCLReferenceComponentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LCLReferenceComponent_Class", 0),
		/*1*/ imports.NewTable("LCLReferenceComponent_HandleAllocated", 0),
		/*2*/ imports.NewTable("LCLReferenceComponent_ReferenceAllocated", 0),
	}
)

func lCLReferenceComponentImportAPI() *imports.Imports {
	if lCLReferenceComponentImport == nil {
		lCLReferenceComponentImport = NewDefaultImports()
		lCLReferenceComponentImport.SetImportTable(lCLReferenceComponentImportTables)
		lCLReferenceComponentImportTables = nil
	}
	return lCLReferenceComponentImport
}
