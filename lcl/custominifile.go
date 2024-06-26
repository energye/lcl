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
	r1 := LCL().SysCallN(1951, m.Instance())
	return GoStr(r1)
}

func (m *TCustomIniFile) Options() TIniFileOptions {
	r1 := LCL().SysCallN(1952, 0, m.Instance(), 0)
	return TIniFileOptions(r1)
}

func (m *TCustomIniFile) SetOptions(AValue TIniFileOptions) {
	LCL().SysCallN(1952, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIniFile) BoolTrueStrings() TStringArray {
	r1 := LCL().SysCallN(1947, 0, m.Instance(), 0)
	return TStringArray(r1)
}

func (m *TCustomIniFile) SetBoolTrueStrings(AValue TStringArray) {
	LCL().SysCallN(1947, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIniFile) BoolFalseStrings() TStringArray {
	r1 := LCL().SysCallN(1946, 0, m.Instance(), 0)
	return TStringArray(r1)
}

func (m *TCustomIniFile) SetBoolFalseStrings(AValue TStringArray) {
	LCL().SysCallN(1946, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIniFile) OwnsEncoding() bool {
	r1 := LCL().SysCallN(1953, m.Instance())
	return GoBool(r1)
}

func (m *TCustomIniFile) SectionExists(Section string) bool {
	r1 := LCL().SysCallN(1967, m.Instance(), PascalStr(Section))
	return GoBool(r1)
}

func (m *TCustomIniFile) ReadString(Section, Ident, Default string) string {
	r1 := LCL().SysCallN(1965, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalStr(Default))
	return GoStr(r1)
}

func (m *TCustomIniFile) ReadInteger(Section, Ident string, Default int32) int32 {
	r1 := LCL().SysCallN(1960, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return int32(r1)
}

func (m *TCustomIniFile) ReadInt64(Section, Ident string, Default int64) (resultInt64 int64) {
	LCL().SysCallN(1959, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Default)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCustomIniFile) ReadBool(Section, Ident string, Default bool) bool {
	r1 := LCL().SysCallN(1955, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalBool(Default))
	return GoBool(r1)
}

func (m *TCustomIniFile) ReadDate(Section, Ident string, Default TDateTime) TDateTime {
	r1 := LCL().SysCallN(1956, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return TDateTime(r1)
}

func (m *TCustomIniFile) ReadDateTime(Section, Ident string, Default TDateTime) TDateTime {
	r1 := LCL().SysCallN(1957, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return TDateTime(r1)
}

func (m *TCustomIniFile) ReadFloat(Section, Ident string, Default float64) (resultDouble float64) {
	LCL().SysCallN(1958, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Default)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomIniFile) ReadTime(Section, Ident string, Default TDateTime) TDateTime {
	r1 := LCL().SysCallN(1966, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Default))
	return TDateTime(r1)
}

func (m *TCustomIniFile) ReadBinaryStream(Section, Name string, Value IStream) int32 {
	r1 := LCL().SysCallN(1954, m.Instance(), PascalStr(Section), PascalStr(Name), GetObjectUintptr(Value))
	return int32(r1)
}

func (m *TCustomIniFile) ValueExists(Section, Ident string) bool {
	r1 := LCL().SysCallN(1969, m.Instance(), PascalStr(Section), PascalStr(Ident))
	return GoBool(r1)
}

func CustomIniFileClass() TClass {
	ret := LCL().SysCallN(1948)
	return TClass(ret)
}

func (m *TCustomIniFile) WriteString(Section, Ident, Value string) {
	LCL().SysCallN(1977, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalStr(Value))
}

func (m *TCustomIniFile) WriteInteger(Section, Ident string, Value int32) {
	LCL().SysCallN(1976, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteInt64(Section, Ident string, Value int64) {
	LCL().SysCallN(1975, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Value)))
}

func (m *TCustomIniFile) WriteBool(Section, Ident string, Value bool) {
	LCL().SysCallN(1971, m.Instance(), PascalStr(Section), PascalStr(Ident), PascalBool(Value))
}

func (m *TCustomIniFile) WriteDate(Section, Ident string, Value TDateTime) {
	LCL().SysCallN(1972, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteDateTime(Section, Ident string, Value TDateTime) {
	LCL().SysCallN(1973, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteFloat(Section, Ident string, Value float64) {
	LCL().SysCallN(1974, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(unsafePointer(&Value)))
}

func (m *TCustomIniFile) WriteTime(Section, Ident string, Value TDateTime) {
	LCL().SysCallN(1978, m.Instance(), PascalStr(Section), PascalStr(Ident), uintptr(Value))
}

func (m *TCustomIniFile) WriteBinaryStream(Section, Name string, Value IStream) {
	LCL().SysCallN(1970, m.Instance(), PascalStr(Section), PascalStr(Name), GetObjectUintptr(Value))
}

func (m *TCustomIniFile) ReadSection(Section string, Strings IStrings) {
	LCL().SysCallN(1961, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings))
}

func (m *TCustomIniFile) ReadSections(Strings IStrings) {
	LCL().SysCallN(1964, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TCustomIniFile) ReadSectionValues(Section string, Strings IStrings, Options TSectionValuesOptions) {
	LCL().SysCallN(1962, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings), uintptr(Options))
}

func (m *TCustomIniFile) ReadSectionValues1(Section string, Strings IStrings) {
	LCL().SysCallN(1963, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings))
}

func (m *TCustomIniFile) EraseSection(Section string) {
	LCL().SysCallN(1950, m.Instance(), PascalStr(Section))
}

func (m *TCustomIniFile) DeleteKey(Section, Ident string) {
	LCL().SysCallN(1949, m.Instance(), PascalStr(Section), PascalStr(Ident))
}

func (m *TCustomIniFile) UpdateFile() {
	LCL().SysCallN(1968, m.Instance())
}
