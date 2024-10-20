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

// ICustomIniFile Is Abstract Class Parent: IObject
type ICustomIniFile interface {
	IObject
	FileName() string                                                                  // property
	Options() TIniFileOptions                                                          // property
	SetOptions(AValue TIniFileOptions)                                                 // property
	BoolTrueStrings() TStringArray                                                     // property
	SetBoolTrueStrings(AValue TStringArray)                                            // property
	BoolFalseStrings() TStringArray                                                    // property
	SetBoolFalseStrings(AValue TStringArray)                                           // property
	OwnsEncoding() bool                                                                // property
	SectionExists(Section string) bool                                                 // function
	ReadString(Section, Ident, Default string) string                                  // function Is Abstract
	ReadInteger(Section, Ident string, Default int32) int32                            // function
	ReadInt64(Section, Ident string, Default int64) (resultInt64 int64)                // function
	ReadBool(Section, Ident string, Default bool) bool                                 // function
	ReadDate(Section, Ident string, Default TDateTime) TDateTime                       // function
	ReadDateTime(Section, Ident string, Default TDateTime) TDateTime                   // function
	ReadFloat(Section, Ident string, Default float64) (resultDouble float64)           // function
	ReadTime(Section, Ident string, Default TDateTime) TDateTime                       // function
	ReadBinaryStream(Section, Name string, Value IStream) int32                        // function
	ValueExists(Section, Ident string) bool                                            // function
	WriteString(Section, Ident, Value string)                                          // procedure Is Abstract
	WriteInteger(Section, Ident string, Value int32)                                   // procedure
	WriteInt64(Section, Ident string, Value int64)                                     // procedure
	WriteBool(Section, Ident string, Value bool)                                       // procedure
	WriteDate(Section, Ident string, Value TDateTime)                                  // procedure
	WriteDateTime(Section, Ident string, Value TDateTime)                              // procedure
	WriteFloat(Section, Ident string, Value float64)                                   // procedure
	WriteTime(Section, Ident string, Value TDateTime)                                  // procedure
	WriteBinaryStream(Section, Name string, Value IStream)                             // procedure
	ReadSection(Section string, Strings IStrings)                                      // procedure Is Abstract
	ReadSections(Strings IStrings)                                                     // procedure Is Abstract
	ReadSectionValues(Section string, Strings IStrings, Options TSectionValuesOptions) // procedure
	ReadSectionValues1(Section string, Strings IStrings)                               // procedure
	EraseSection(Section string)                                                       // procedure Is Abstract
	DeleteKey(Section, Ident string)                                                   // procedure Is Abstract
	UpdateFile()                                                                       // procedure Is Abstract
}

// TCustomIniFile Is Abstract Class Parent: TObject
type TCustomIniFile struct {
	TObject
}

func (m *TCustomIniFile) FileName() string {
	r1 := customIniFileImportAPI().SysCallN(5, m.Instance())
	return GoStr(r1)
}

func (m *TCustomIniFile) Options() TIniFileOptions {
	r1 := customIniFileImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TIniFileOptions(r1)
}

func (m *TCustomIniFile) SetOptions(AValue TIniFileOptions) {
	customIniFileImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIniFile) BoolTrueStrings() TStringArray {
	r1 := customIniFileImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TStringArray(r1)
}

func (m *TCustomIniFile) SetBoolTrueStrings(AValue TStringArray) {
	customIniFileImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIniFile) BoolFalseStrings() TStringArray {
	r1 := customIniFileImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TStringArray(r1)
}

func (m *TCustomIniFile) SetBoolFalseStrings(AValue TStringArray) {
	customIniFileImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIniFile) OwnsEncoding() bool {
	r1 := customIniFileImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCustomIniFile) SectionExists(Section string) bool {
	r1 := customIniFileImportAPI().SysCallN(21, m.Instance(), PascalStr(Section))
	return GoBool(r1)
}

func (m *TCustomIniFile) ReadString(Section, Ident, Default string) string {
	r1 := customIniFileImportAPI().SysCallN(19, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalStr(Default))
	return GoStr(r1)
}

func (m *TCustomIniFile) ReadInteger(Section, Ident string, Default int32) int32 {
	r1 := customIniFileImportAPI().SysCallN(14, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return int32(r1)
}

func (m *TCustomIniFile) ReadInt64(Section, Ident string, Default int64) (resultInt64 int64) {
	customIniFileImportAPI().SysCallN(13, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Default)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCustomIniFile) ReadBool(Section, Ident string, Default bool) bool {
	r1 := customIniFileImportAPI().SysCallN(9, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalBool(Default))
	return GoBool(r1)
}

func (m *TCustomIniFile) ReadDate(Section, Ident string, Default TDateTime) TDateTime {
	r1 := customIniFileImportAPI().SysCallN(10, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return TDateTime(r1)
}

func (m *TCustomIniFile) ReadDateTime(Section, Ident string, Default TDateTime) TDateTime {
	r1 := customIniFileImportAPI().SysCallN(11, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return TDateTime(r1)
}

func (m *TCustomIniFile) ReadFloat(Section, Ident string, Default float64) (resultDouble float64) {
	customIniFileImportAPI().SysCallN(12, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Default)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomIniFile) ReadTime(Section, Ident string, Default TDateTime) TDateTime {
	r1 := customIniFileImportAPI().SysCallN(20, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return TDateTime(r1)
}

func (m *TCustomIniFile) ReadBinaryStream(Section, Name string, Value IStream) int32 {
	r1 := customIniFileImportAPI().SysCallN(8, m.Instance(), PascalStr(Section), PascalStr(Name), GetObjectUintptr(Value))
	return int32(r1)
}

func (m *TCustomIniFile) ValueExists(Section, Ident string) bool {
	r1 := customIniFileImportAPI().SysCallN(23, m.Instance(), PascalStr(Section), PascalStr(Ident))
	return GoBool(r1)
}

func CustomIniFileClass() TClass {
	ret := customIniFileImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomIniFile) WriteString(Section, Ident, Value string) {
	customIniFileImportAPI().SysCallN(31, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalStr(Value))
}

func (m *TCustomIniFile) WriteInteger(Section, Ident string, Value int32) {
	customIniFileImportAPI().SysCallN(30, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteInt64(Section, Ident string, Value int64) {
	customIniFileImportAPI().SysCallN(29, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Value)))
}

func (m *TCustomIniFile) WriteBool(Section, Ident string, Value bool) {
	customIniFileImportAPI().SysCallN(25, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalBool(Value))
}

func (m *TCustomIniFile) WriteDate(Section, Ident string, Value TDateTime) {
	customIniFileImportAPI().SysCallN(26, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteDateTime(Section, Ident string, Value TDateTime) {
	customIniFileImportAPI().SysCallN(27, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteFloat(Section, Ident string, Value float64) {
	customIniFileImportAPI().SysCallN(28, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Value)))
}

func (m *TCustomIniFile) WriteTime(Section, Ident string, Value TDateTime) {
	customIniFileImportAPI().SysCallN(32, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteBinaryStream(Section, Name string, Value IStream) {
	customIniFileImportAPI().SysCallN(24, m.Instance(), PascalStr(Section), PascalStr(Name), GetObjectUintptr(Value))
}

func (m *TCustomIniFile) ReadSection(Section string, Strings IStrings) {
	customIniFileImportAPI().SysCallN(15, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings))
}

func (m *TCustomIniFile) ReadSections(Strings IStrings) {
	customIniFileImportAPI().SysCallN(18, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TCustomIniFile) ReadSectionValues(Section string, Strings IStrings, Options TSectionValuesOptions) {
	customIniFileImportAPI().SysCallN(16, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings), uintptr(Options))
}

func (m *TCustomIniFile) ReadSectionValues1(Section string, Strings IStrings) {
	customIniFileImportAPI().SysCallN(17, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings))
}

func (m *TCustomIniFile) EraseSection(Section string) {
	customIniFileImportAPI().SysCallN(4, m.Instance(), PascalStr(Section))
}

func (m *TCustomIniFile) DeleteKey(Section, Ident string) {
	customIniFileImportAPI().SysCallN(3, m.Instance(), PascalStr(Section), PascalStr(Ident))
}

func (m *TCustomIniFile) UpdateFile() {
	customIniFileImportAPI().SysCallN(22, m.Instance())
}

var (
	customIniFileImport       *imports.Imports = nil
	customIniFileImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomIniFile_BoolFalseStrings", 0),
		/*1*/ imports.NewTable("CustomIniFile_BoolTrueStrings", 0),
		/*2*/ imports.NewTable("CustomIniFile_Class", 0),
		/*3*/ imports.NewTable("CustomIniFile_DeleteKey", 0),
		/*4*/ imports.NewTable("CustomIniFile_EraseSection", 0),
		/*5*/ imports.NewTable("CustomIniFile_FileName", 0),
		/*6*/ imports.NewTable("CustomIniFile_Options", 0),
		/*7*/ imports.NewTable("CustomIniFile_OwnsEncoding", 0),
		/*8*/ imports.NewTable("CustomIniFile_ReadBinaryStream", 0),
		/*9*/ imports.NewTable("CustomIniFile_ReadBool", 0),
		/*10*/ imports.NewTable("CustomIniFile_ReadDate", 0),
		/*11*/ imports.NewTable("CustomIniFile_ReadDateTime", 0),
		/*12*/ imports.NewTable("CustomIniFile_ReadFloat", 0),
		/*13*/ imports.NewTable("CustomIniFile_ReadInt64", 0),
		/*14*/ imports.NewTable("CustomIniFile_ReadInteger", 0),
		/*15*/ imports.NewTable("CustomIniFile_ReadSection", 0),
		/*16*/ imports.NewTable("CustomIniFile_ReadSectionValues", 0),
		/*17*/ imports.NewTable("CustomIniFile_ReadSectionValues1", 0),
		/*18*/ imports.NewTable("CustomIniFile_ReadSections", 0),
		/*19*/ imports.NewTable("CustomIniFile_ReadString", 0),
		/*20*/ imports.NewTable("CustomIniFile_ReadTime", 0),
		/*21*/ imports.NewTable("CustomIniFile_SectionExists", 0),
		/*22*/ imports.NewTable("CustomIniFile_UpdateFile", 0),
		/*23*/ imports.NewTable("CustomIniFile_ValueExists", 0),
		/*24*/ imports.NewTable("CustomIniFile_WriteBinaryStream", 0),
		/*25*/ imports.NewTable("CustomIniFile_WriteBool", 0),
		/*26*/ imports.NewTable("CustomIniFile_WriteDate", 0),
		/*27*/ imports.NewTable("CustomIniFile_WriteDateTime", 0),
		/*28*/ imports.NewTable("CustomIniFile_WriteFloat", 0),
		/*29*/ imports.NewTable("CustomIniFile_WriteInt64", 0),
		/*30*/ imports.NewTable("CustomIniFile_WriteInteger", 0),
		/*31*/ imports.NewTable("CustomIniFile_WriteString", 0),
		/*32*/ imports.NewTable("CustomIniFile_WriteTime", 0),
	}
)

func customIniFileImportAPI() *imports.Imports {
	if customIniFileImport == nil {
		customIniFileImport = NewDefaultImports()
		customIniFileImport.SetImportTable(customIniFileImportTables)
		customIniFileImportTables = nil
	}
	return customIniFileImport
}
