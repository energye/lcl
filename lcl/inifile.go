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

// IIniFile Parent: ICustomIniFile
type IIniFile interface {
	ICustomIniFile
	Stream() IStream                                 // property
	CacheUpdates() bool                              // property
	SetCacheUpdates(AValue bool)                     // property
	WriteBOM() bool                                  // property
	SetWriteBOM(AValue bool)                         // property
	ReadSectionRaw(Section string, Strings IStrings) // procedure
}

// TIniFile Parent: TCustomIniFile
type TIniFile struct {
	TCustomIniFile
}

func NewIniFile(AFileName string, AOptions TIniFileOptions) IIniFile {
	r1 := niFileImportAPI().SysCallN(2, PascalStr(AFileName), uintptr(AOptions))
	return AsIniFile(r1)
}

func NewIniFile1(AStream IStream, AOptions TIniFileOptions) IIniFile {
	r1 := niFileImportAPI().SysCallN(3, GetObjectUintptr(AStream), uintptr(AOptions))
	return AsIniFile(r1)
}

func (m *TIniFile) Stream() IStream {
	r1 := niFileImportAPI().SysCallN(5, m.Instance())
	return AsStream(r1)
}

func (m *TIniFile) CacheUpdates() bool {
	r1 := niFileImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIniFile) SetCacheUpdates(AValue bool) {
	niFileImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TIniFile) WriteBOM() bool {
	r1 := niFileImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIniFile) SetWriteBOM(AValue bool) {
	niFileImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func IniFileClass() TClass {
	ret := niFileImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TIniFile) ReadSectionRaw(Section string, Strings IStrings) {
	niFileImportAPI().SysCallN(4, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings))
}

var (
	niFileImport       *imports.Imports = nil
	niFileImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("IniFile_CacheUpdates", 0),
		/*1*/ imports.NewTable("IniFile_Class", 0),
		/*2*/ imports.NewTable("IniFile_Create", 0),
		/*3*/ imports.NewTable("IniFile_Create1", 0),
		/*4*/ imports.NewTable("IniFile_ReadSectionRaw", 0),
		/*5*/ imports.NewTable("IniFile_Stream", 0),
		/*6*/ imports.NewTable("IniFile_WriteBOM", 0),
	}
)

func niFileImportAPI() *imports.Imports {
	if niFileImport == nil {
		niFileImport = NewDefaultImports()
		niFileImport.SetImportTable(niFileImportTables)
		niFileImportTables = nil
	}
	return niFileImport
}
