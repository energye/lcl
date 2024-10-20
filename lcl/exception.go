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

// IException Parent: IObject
type IException interface {
	IObject
	HelpContext() int32          // property
	SetHelpContext(AValue int32) // property
	Message() string             // property
	SetMessage(AValue string)    // property
}

// Exception Parent: TObject
type Exception struct {
	TObject
}

func NewException(msg string) IException {
	r1 := exceptionImportAPI().SysCallN(1, PascalStr(msg))
	return AsException(r1)
}

func NewExceptionHelp(Msg string, AHelpContext int32) IException {
	r1 := exceptionImportAPI().SysCallN(2, PascalStr(Msg), uintptr(AHelpContext))
	return AsException(r1)
}

func (m *Exception) HelpContext() int32 {
	r1 := exceptionImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *Exception) SetHelpContext(AValue int32) {
	exceptionImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *Exception) Message() string {
	r1 := exceptionImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *Exception) SetMessage(AValue string) {
	exceptionImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func ExceptionClass() TClass {
	ret := exceptionImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	exceptionImport       *imports.Imports = nil
	exceptionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Exception_Class", 0),
		/*1*/ imports.NewTable("Exception_Create", 0),
		/*2*/ imports.NewTable("Exception_CreateHelp", 0),
		/*3*/ imports.NewTable("Exception_HelpContext", 0),
		/*4*/ imports.NewTable("Exception_Message", 0),
	}
)

func exceptionImportAPI() *imports.Imports {
	if exceptionImport == nil {
		exceptionImport = NewDefaultImports()
		exceptionImport.SetImportTable(exceptionImportTables)
		exceptionImportTables = nil
	}
	return exceptionImport
}
