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

// ICustomGroupBox Parent: IWinControl
type ICustomGroupBox interface {
	IWinControl
	ParentBackground() bool          // property
	SetParentBackground(AValue bool) // property
}

// TCustomGroupBox Parent: TWinControl
type TCustomGroupBox struct {
	TWinControl
}

func NewCustomGroupBox(AOwner IComponent) ICustomGroupBox {
	r1 := customGroupBoxImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomGroupBox(r1)
}

func (m *TCustomGroupBox) ParentBackground() bool {
	r1 := customGroupBoxImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomGroupBox) SetParentBackground(AValue bool) {
	customGroupBoxImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func CustomGroupBoxClass() TClass {
	ret := customGroupBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customGroupBoxImport       *imports.Imports = nil
	customGroupBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomGroupBox_Class", 0),
		/*1*/ imports.NewTable("CustomGroupBox_Create", 0),
		/*2*/ imports.NewTable("CustomGroupBox_ParentBackground", 0),
	}
)

func customGroupBoxImportAPI() *imports.Imports {
	if customGroupBoxImport == nil {
		customGroupBoxImport = NewDefaultImports()
		customGroupBoxImport.SetImportTable(customGroupBoxImportTables)
		customGroupBoxImportTables = nil
	}
	return customGroupBoxImport
}
