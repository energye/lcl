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

// ICustomIniFile Parent: IObject
type ICustomIniFile interface {
	IObject
	SectionExists(section string) bool                                                                                            // function
	ReadString(section string, ident string, default_ string) string                                                              // function
	ReadInteger(section string, ident string, default_ int32) int32                                                               // function
	ReadInt64(section string, ident string, default_ int64) int64                                                                 // function
	ReadBool(section string, ident string, default_ bool) bool                                                                    // function
	ReadDate(section string, ident string, default_ types.TDateTime) types.TDateTime                                              // function
	ReadDateTime(section string, ident string, default_ types.TDateTime) types.TDateTime                                          // function
	ReadFloat(section string, ident string, default_ float64) float64                                                             // function
	ReadTime(section string, ident string, default_ types.TDateTime) types.TDateTime                                              // function
	ReadBinaryStream(section string, name string, value IStream) int32                                                            // function
	ValueExists(section string, ident string) bool                                                                                // function
	WriteString(section string, ident string, value string)                                                                       // procedure
	WriteInteger(section string, ident string, value int32)                                                                       // procedure
	WriteInt64(section string, ident string, value int64)                                                                         // procedure
	WriteBool(section string, ident string, value bool)                                                                           // procedure
	WriteDate(section string, ident string, value types.TDateTime)                                                                // procedure
	WriteDateTime(section string, ident string, value types.TDateTime)                                                            // procedure
	WriteFloat(section string, ident string, value float64)                                                                       // procedure
	WriteTime(section string, ident string, value types.TDateTime)                                                                // procedure
	WriteBinaryStream(section string, name string, value IStream)                                                                 // procedure
	ReadSection(section string, strings IStrings)                                                                                 // procedure
	ReadSections(strings IStrings)                                                                                                // procedure
	ReadSectionValuesWithStringStringsSectionValuesOptions(section string, strings IStrings, options types.TSectionValuesOptions) // procedure
	ReadSectionValuesWithStringStrings(section string, strings IStrings)                                                          // procedure
	EraseSection(section string)                                                                                                  // procedure
	DeleteKey(section string, ident string)                                                                                       // procedure
	UpdateFile()                                                                                                                  // procedure
	FileName() string                                                                                                             // property FileName Getter
	Options() types.TIniFileOptions                                                                                               // property Options Getter
	SetOptions(value types.TIniFileOptions)                                                                                       // property Options Setter
	BoolTrueStrings() IStringArrayWrap                                                                                            // property BoolTrueStrings Getter
	SetBoolTrueStrings(value IStringArrayWrap)                                                                                    // property BoolTrueStrings Setter
	BoolFalseStrings() IStringArrayWrap                                                                                           // property BoolFalseStrings Getter
	SetBoolFalseStrings(value IStringArrayWrap)                                                                                   // property BoolFalseStrings Setter
	OwnsEncoding() bool                                                                                                           // property OwnsEncoding Getter
}

type TCustomIniFile struct {
	TObject
}

func (m *TCustomIniFile) SectionExists(section string) bool {
	if !m.IsValid() {
		return false
	}
	r := customIniFileAPI().SysCallN(0, m.Instance(), api.PasStr(section))
	return api.GoBool(r)
}

func (m *TCustomIniFile) ReadString(section string, ident string, default_ string) string {
	if !m.IsValid() {
		return ""
	}
	r := customIniFileAPI().SysCallN(1, m.Instance(), api.PasStr(section), api.PasStr(ident), api.PasStr(default_))
	return api.GoStr(r)
}

func (m *TCustomIniFile) ReadInteger(section string, ident string, default_ int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customIniFileAPI().SysCallN(2, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(default_))
	return int32(r)
}

func (m *TCustomIniFile) ReadInt64(section string, ident string, default_ int64) (result int64) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(3, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&default_)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomIniFile) ReadBool(section string, ident string, default_ bool) bool {
	if !m.IsValid() {
		return false
	}
	r := customIniFileAPI().SysCallN(4, m.Instance(), api.PasStr(section), api.PasStr(ident), api.PasBool(default_))
	return api.GoBool(r)
}

func (m *TCustomIniFile) ReadDate(section string, ident string, default_ types.TDateTime) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(5, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&default_)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomIniFile) ReadDateTime(section string, ident string, default_ types.TDateTime) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(6, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&default_)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomIniFile) ReadFloat(section string, ident string, default_ float64) (result float64) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(7, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&default_)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomIniFile) ReadTime(section string, ident string, default_ types.TDateTime) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(8, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&default_)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomIniFile) ReadBinaryStream(section string, name string, value IStream) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customIniFileAPI().SysCallN(9, m.Instance(), api.PasStr(section), api.PasStr(name), base.GetObjectUintptr(value))
	return int32(r)
}

func (m *TCustomIniFile) ValueExists(section string, ident string) bool {
	if !m.IsValid() {
		return false
	}
	r := customIniFileAPI().SysCallN(10, m.Instance(), api.PasStr(section), api.PasStr(ident))
	return api.GoBool(r)
}

func (m *TCustomIniFile) WriteString(section string, ident string, value string) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(11, m.Instance(), api.PasStr(section), api.PasStr(ident), api.PasStr(value))
}

func (m *TCustomIniFile) WriteInteger(section string, ident string, value int32) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(12, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(value))
}

func (m *TCustomIniFile) WriteInt64(section string, ident string, value int64) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(13, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomIniFile) WriteBool(section string, ident string, value bool) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(14, m.Instance(), api.PasStr(section), api.PasStr(ident), api.PasBool(value))
}

func (m *TCustomIniFile) WriteDate(section string, ident string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(15, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomIniFile) WriteDateTime(section string, ident string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(16, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomIniFile) WriteFloat(section string, ident string, value float64) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(17, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomIniFile) WriteTime(section string, ident string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(18, m.Instance(), api.PasStr(section), api.PasStr(ident), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomIniFile) WriteBinaryStream(section string, name string, value IStream) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(19, m.Instance(), api.PasStr(section), api.PasStr(name), base.GetObjectUintptr(value))
}

func (m *TCustomIniFile) ReadSection(section string, strings IStrings) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(20, m.Instance(), api.PasStr(section), base.GetObjectUintptr(strings))
}

func (m *TCustomIniFile) ReadSections(strings IStrings) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(strings))
}

func (m *TCustomIniFile) ReadSectionValuesWithStringStringsSectionValuesOptions(section string, strings IStrings, options types.TSectionValuesOptions) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(22, m.Instance(), api.PasStr(section), base.GetObjectUintptr(strings), uintptr(options))
}

func (m *TCustomIniFile) ReadSectionValuesWithStringStrings(section string, strings IStrings) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(23, m.Instance(), api.PasStr(section), base.GetObjectUintptr(strings))
}

func (m *TCustomIniFile) EraseSection(section string) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(24, m.Instance(), api.PasStr(section))
}

func (m *TCustomIniFile) DeleteKey(section string, ident string) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(25, m.Instance(), api.PasStr(section), api.PasStr(ident))
}

func (m *TCustomIniFile) UpdateFile() {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(26, m.Instance())
}

func (m *TCustomIniFile) FileName() string {
	if !m.IsValid() {
		return ""
	}
	r := customIniFileAPI().SysCallN(27, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomIniFile) Options() types.TIniFileOptions {
	if !m.IsValid() {
		return 0
	}
	r := customIniFileAPI().SysCallN(28, 0, m.Instance())
	return types.TIniFileOptions(r)
}

func (m *TCustomIniFile) SetOptions(value types.TIniFileOptions) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TCustomIniFile) BoolTrueStrings() IStringArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := customIniFileAPI().SysCallN(29, 0, m.Instance())
	return AsStringArrayWrap(r)
}

func (m *TCustomIniFile) SetBoolTrueStrings(value IStringArrayWrap) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(29, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomIniFile) BoolFalseStrings() IStringArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := customIniFileAPI().SysCallN(30, 0, m.Instance())
	return AsStringArrayWrap(r)
}

func (m *TCustomIniFile) SetBoolFalseStrings(value IStringArrayWrap) {
	if !m.IsValid() {
		return
	}
	customIniFileAPI().SysCallN(30, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomIniFile) OwnsEncoding() bool {
	if !m.IsValid() {
		return false
	}
	r := customIniFileAPI().SysCallN(31, m.Instance())
	return api.GoBool(r)
}

var (
	customIniFileOnce   base.Once
	customIniFileImport *imports.Imports = nil
)

func customIniFileAPI() *imports.Imports {
	customIniFileOnce.Do(func() {
		customIniFileImport = api.NewDefaultImports()
		customIniFileImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomIniFile_SectionExists", 0), // function SectionExists
			/* 1 */ imports.NewTable("TCustomIniFile_ReadString", 0), // function ReadString
			/* 2 */ imports.NewTable("TCustomIniFile_ReadInteger", 0), // function ReadInteger
			/* 3 */ imports.NewTable("TCustomIniFile_ReadInt64", 0), // function ReadInt64
			/* 4 */ imports.NewTable("TCustomIniFile_ReadBool", 0), // function ReadBool
			/* 5 */ imports.NewTable("TCustomIniFile_ReadDate", 0), // function ReadDate
			/* 6 */ imports.NewTable("TCustomIniFile_ReadDateTime", 0), // function ReadDateTime
			/* 7 */ imports.NewTable("TCustomIniFile_ReadFloat", 0), // function ReadFloat
			/* 8 */ imports.NewTable("TCustomIniFile_ReadTime", 0), // function ReadTime
			/* 9 */ imports.NewTable("TCustomIniFile_ReadBinaryStream", 0), // function ReadBinaryStream
			/* 10 */ imports.NewTable("TCustomIniFile_ValueExists", 0), // function ValueExists
			/* 11 */ imports.NewTable("TCustomIniFile_WriteString", 0), // procedure WriteString
			/* 12 */ imports.NewTable("TCustomIniFile_WriteInteger", 0), // procedure WriteInteger
			/* 13 */ imports.NewTable("TCustomIniFile_WriteInt64", 0), // procedure WriteInt64
			/* 14 */ imports.NewTable("TCustomIniFile_WriteBool", 0), // procedure WriteBool
			/* 15 */ imports.NewTable("TCustomIniFile_WriteDate", 0), // procedure WriteDate
			/* 16 */ imports.NewTable("TCustomIniFile_WriteDateTime", 0), // procedure WriteDateTime
			/* 17 */ imports.NewTable("TCustomIniFile_WriteFloat", 0), // procedure WriteFloat
			/* 18 */ imports.NewTable("TCustomIniFile_WriteTime", 0), // procedure WriteTime
			/* 19 */ imports.NewTable("TCustomIniFile_WriteBinaryStream", 0), // procedure WriteBinaryStream
			/* 20 */ imports.NewTable("TCustomIniFile_ReadSection", 0), // procedure ReadSection
			/* 21 */ imports.NewTable("TCustomIniFile_ReadSections", 0), // procedure ReadSections
			/* 22 */ imports.NewTable("TCustomIniFile_ReadSectionValuesWithStringStringsSectionValuesOptions", 0), // procedure ReadSectionValuesWithStringStringsSectionValuesOptions
			/* 23 */ imports.NewTable("TCustomIniFile_ReadSectionValuesWithStringStrings", 0), // procedure ReadSectionValuesWithStringStrings
			/* 24 */ imports.NewTable("TCustomIniFile_EraseSection", 0), // procedure EraseSection
			/* 25 */ imports.NewTable("TCustomIniFile_DeleteKey", 0), // procedure DeleteKey
			/* 26 */ imports.NewTable("TCustomIniFile_UpdateFile", 0), // procedure UpdateFile
			/* 27 */ imports.NewTable("TCustomIniFile_FileName", 0), // property FileName
			/* 28 */ imports.NewTable("TCustomIniFile_Options", 0), // property Options
			/* 29 */ imports.NewTable("TCustomIniFile_BoolTrueStrings", 0), // property BoolTrueStrings
			/* 30 */ imports.NewTable("TCustomIniFile_BoolFalseStrings", 0), // property BoolFalseStrings
			/* 31 */ imports.NewTable("TCustomIniFile_OwnsEncoding", 0), // property OwnsEncoding
		}
	})
	return customIniFileImport
}
