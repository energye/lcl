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

// IIconOptions Parent: IPersistent
type IIconOptions interface {
	IPersistent
	Arrangement() TIconArrangement          // property
	SetArrangement(AValue TIconArrangement) // property
	AutoArrange() bool                      // property
	SetAutoArrange(AValue bool)             // property
	WrapText() bool                         // property
	SetWrapText(AValue bool)                // property
}

// TIconOptions Parent: TPersistent
type TIconOptions struct {
	TPersistent
}

func NewIconOptions(AOwner ICustomListView) IIconOptions {
	r1 := conOptionsImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsIconOptions(r1)
}

func (m *TIconOptions) Arrangement() TIconArrangement {
	r1 := conOptionsImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TIconArrangement(r1)
}

func (m *TIconOptions) SetArrangement(AValue TIconArrangement) {
	conOptionsImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TIconOptions) AutoArrange() bool {
	r1 := conOptionsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIconOptions) SetAutoArrange(AValue bool) {
	conOptionsImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TIconOptions) WrapText() bool {
	r1 := conOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIconOptions) SetWrapText(AValue bool) {
	conOptionsImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func IconOptionsClass() TClass {
	ret := conOptionsImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	conOptionsImport       *imports.Imports = nil
	conOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("IconOptions_Arrangement", 0),
		/*1*/ imports.NewTable("IconOptions_AutoArrange", 0),
		/*2*/ imports.NewTable("IconOptions_Class", 0),
		/*3*/ imports.NewTable("IconOptions_Create", 0),
		/*4*/ imports.NewTable("IconOptions_WrapText", 0),
	}
)

func conOptionsImportAPI() *imports.Imports {
	if conOptionsImport == nil {
		conOptionsImport = NewDefaultImports()
		conOptionsImport.SetImportTable(conOptionsImportTables)
		conOptionsImportTables = nil
	}
	return conOptionsImport
}
