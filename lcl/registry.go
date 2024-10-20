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

// IRegistry Parent: IObject
type IRegistry interface {
	IObject
	Access() uint32                                            // property
	SetAccess(AValue uint32)                                   // property
	CurrentKey() HKEY                                          // property
	CurrentPath() string                                       // property
	LazyWrite() bool                                           // property
	SetLazyWrite(AValue bool)                                  // property
	RootKey() HKEY                                             // property
	SetRootKey(AValue HKEY)                                    // property
	StringSizeIncludesNull() bool                              // property
	LastError() int32                                          // property
	LastErrorMsg() string                                      // property
	CreateKey(Key string) bool                                 // function
	DeleteKey(Key string) bool                                 // function
	DeleteValue(Name string) bool                              // function
	GetDataInfo(ValueName string, OutValue *TRegDataInfo) bool // function
	GetDataSize(ValueName string) int32                        // function
	GetDataType(ValueName string) TRegDataType                 // function
	GetKeyInfo(OutValue *TRegKeyInfo) bool                     // function
	HasSubKeys() bool                                          // function
	KeyExists(Key string) bool                                 // function
	LoadKey(Key, FileName string) bool                         // function
	OpenKey(Key string, CanCreate bool) bool                   // function
	OpenKeyReadOnly(Key string) bool                           // function
	ReadCurrency(Name string) Currency                         // function
	ReadBinaryData(Name string, count int32) []byte            // function
	ReadBool(Name string) bool                                 // function
	ReadDate(Name string) TDateTime                            // function
	ReadDateTime(Name string) TDateTime                        // function
	ReadFloat(Name string) (resultDouble float64)              // function
	ReadInteger(Name string) int32                             // function
	ReadInt64(Name string) (resultInt64 int64)                 // function
	ReadString(Name string) string                             // function
	ReadStringArray(Name string) TStringArray                  // function
	ReadTime(Name string) TDateTime                            // function
	RegistryConnect(UNCName string) bool                       // function
	ReplaceKey(Key, FileName, BackUpFileName string) bool      // function
	RestoreKey(Key, FileName string) bool                      // function
	SaveKey(Key, FileName string) bool                         // function
	UnLoadKey(Key string) bool                                 // function
	ValueExists(Name string) bool                              // function
	GetKeyNames() uintptr                                      // function
	GetValueNames() uintptr                                    // function
	ReadStringList(Name string, AList IStrings)                // procedure
	CloseKey()                                                 // procedure
	CloseKey1(key HKEY)                                        // procedure
	GetKeyNames1(Strings IStrings)                             // procedure
	GetValueNames1(Strings IStrings)                           // procedure
	MoveKey(OldName, NewName string, Delete bool)              // procedure
	RenameValue(OldName, NewName string)                       // procedure
	WriteCurrency(Name string, Value Currency)                 // procedure
	WriteBinaryData(Name string, Buffer []byte)                // procedure
	WriteBool(Name string, Value bool)                         // procedure
	WriteDate(Name string, Value TDateTime)                    // procedure
	WriteDateTime(Name string, Value TDateTime)                // procedure
	WriteFloat(Name string, Value float64)                     // procedure
	WriteInteger(Name string, Value int32)                     // procedure
	WriteInt64(Name string, Value int64)                       // procedure
	WriteString(Name, Value string)                            // procedure
	WriteExpandString(Name, Value string)                      // procedure
	WriteStringArray(Name string, Arr TStringArray)            // procedure
	WriteTime(Name string, Value TDateTime)                    // procedure
}

// TRegistry Parent: TObject
type TRegistry struct {
	TObject
}

func NewRegistry() IRegistry {
	r1 := registryImportAPI().SysCallN(4)
	return AsRegistry(r1)
}

func NewRegistry1(aaccess uint32) IRegistry {
	r1 := registryImportAPI().SysCallN(5, uintptr(aaccess))
	return AsRegistry(r1)
}

func (m *TRegistry) Access() uint32 {
	r1 := registryImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TRegistry) SetAccess(AValue uint32) {
	registryImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TRegistry) CurrentKey() HKEY {
	r1 := registryImportAPI().SysCallN(7, m.Instance())
	return HKEY(r1)
}

func (m *TRegistry) CurrentPath() string {
	r1 := registryImportAPI().SysCallN(8, m.Instance())
	return GoStr(r1)
}

func (m *TRegistry) LazyWrite() bool {
	r1 := registryImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRegistry) SetLazyWrite(AValue bool) {
	registryImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRegistry) RootKey() HKEY {
	r1 := registryImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return HKEY(r1)
}

func (m *TRegistry) SetRootKey(AValue HKEY) {
	registryImportAPI().SysCallN(44, 1, m.Instance(), uintptr(AValue))
}

func (m *TRegistry) StringSizeIncludesNull() bool {
	r1 := registryImportAPI().SysCallN(46, m.Instance())
	return GoBool(r1)
}

func (m *TRegistry) LastError() int32 {
	r1 := registryImportAPI().SysCallN(21, m.Instance())
	return int32(r1)
}

func (m *TRegistry) LastErrorMsg() string {
	r1 := registryImportAPI().SysCallN(22, m.Instance())
	return GoStr(r1)
}

func (m *TRegistry) CreateKey(Key string) bool {
	r1 := registryImportAPI().SysCallN(6, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) DeleteKey(Key string) bool {
	r1 := registryImportAPI().SysCallN(9, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) DeleteValue(Name string) bool {
	r1 := registryImportAPI().SysCallN(10, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) GetDataInfo(ValueName string, OutValue *TRegDataInfo) bool {
	var result1 uintptr
	r1 := registryImportAPI().SysCallN(11, m.Instance(), PascalStr(ValueName), uintptr(unsafePointer(&result1)))
	*OutValue = *(*TRegDataInfo)(getPointer(result1))
	return GoBool(r1)
}

func (m *TRegistry) GetDataSize(ValueName string) int32 {
	r1 := registryImportAPI().SysCallN(12, m.Instance(), PascalStr(ValueName))
	return int32(r1)
}

func (m *TRegistry) GetDataType(ValueName string) TRegDataType {
	r1 := registryImportAPI().SysCallN(13, m.Instance(), PascalStr(ValueName))
	return TRegDataType(r1)
}

func (m *TRegistry) GetKeyInfo(OutValue *TRegKeyInfo) bool {
	var result0 uintptr
	r1 := registryImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(&result0)))
	*OutValue = *(*TRegKeyInfo)(getPointer(result0))
	return GoBool(r1)
}

func (m *TRegistry) HasSubKeys() bool {
	r1 := registryImportAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TRegistry) KeyExists(Key string) bool {
	r1 := registryImportAPI().SysCallN(20, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) LoadKey(Key, FileName string) bool {
	r1 := registryImportAPI().SysCallN(24, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) OpenKey(Key string, CanCreate bool) bool {
	r1 := registryImportAPI().SysCallN(26, m.Instance(), PascalStr(Key), PascalBool(CanCreate))
	return GoBool(r1)
}

func (m *TRegistry) OpenKeyReadOnly(Key string) bool {
	r1 := registryImportAPI().SysCallN(27, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) ReadCurrency(Name string) Currency {
	r1 := registryImportAPI().SysCallN(30, m.Instance(), PascalStr(Name))
	return Currency(r1)
}

func (m *TRegistry) ReadBinaryData(Name string, count int32) []byte {
	if count <= 0 {
		return nil
	}
	buffer := make([]byte, count)
	registryImportAPI().SysCallN(28, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&buffer[0])), uintptr(count))
	return buffer
}

func (m *TRegistry) ReadBool(Name string) bool {
	r1 := registryImportAPI().SysCallN(29, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) ReadDate(Name string) TDateTime {
	r1 := registryImportAPI().SysCallN(31, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) ReadDateTime(Name string) TDateTime {
	r1 := registryImportAPI().SysCallN(32, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) ReadFloat(Name string) (resultDouble float64) {
	registryImportAPI().SysCallN(33, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TRegistry) ReadInteger(Name string) int32 {
	r1 := registryImportAPI().SysCallN(35, m.Instance(), PascalStr(Name))
	return int32(r1)
}

func (m *TRegistry) ReadInt64(Name string) (resultInt64 int64) {
	registryImportAPI().SysCallN(34, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TRegistry) ReadString(Name string) string {
	r1 := registryImportAPI().SysCallN(36, m.Instance(), PascalStr(Name))
	return GoStr(r1)
}

func (m *TRegistry) ReadStringArray(Name string) TStringArray {
	r1 := registryImportAPI().SysCallN(37, m.Instance(), PascalStr(Name))
	return TStringArray(r1)
}

func (m *TRegistry) ReadTime(Name string) TDateTime {
	r1 := registryImportAPI().SysCallN(39, m.Instance(), PascalStr(Name))
	return TDateTime(r1)
}

func (m *TRegistry) RegistryConnect(UNCName string) bool {
	r1 := registryImportAPI().SysCallN(40, m.Instance(), PascalStr(UNCName))
	return GoBool(r1)
}

func (m *TRegistry) ReplaceKey(Key, FileName, BackUpFileName string) bool {
	r1 := registryImportAPI().SysCallN(42, m.Instance(), PascalStr(Key), PascalStr(FileName), PascalStr(BackUpFileName))
	return GoBool(r1)
}

func (m *TRegistry) RestoreKey(Key, FileName string) bool {
	r1 := registryImportAPI().SysCallN(43, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) SaveKey(Key, FileName string) bool {
	r1 := registryImportAPI().SysCallN(45, m.Instance(), PascalStr(Key), PascalStr(FileName))
	return GoBool(r1)
}

func (m *TRegistry) UnLoadKey(Key string) bool {
	r1 := registryImportAPI().SysCallN(47, m.Instance(), PascalStr(Key))
	return GoBool(r1)
}

func (m *TRegistry) ValueExists(Name string) bool {
	r1 := registryImportAPI().SysCallN(48, m.Instance(), PascalStr(Name))
	return GoBool(r1)
}

func (m *TRegistry) GetKeyNames() uintptr {
	r1 := registryImportAPI().SysCallN(15, m.Instance())
	return uintptr(r1)
}

func (m *TRegistry) GetValueNames() uintptr {
	r1 := registryImportAPI().SysCallN(17, m.Instance())
	return uintptr(r1)
}

func RegistryClass() TClass {
	ret := registryImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TRegistry) ReadStringList(Name string, AList IStrings) {
	registryImportAPI().SysCallN(38, m.Instance(), PascalStr(Name), GetObjectUintptr(AList))
}

func (m *TRegistry) CloseKey() {
	registryImportAPI().SysCallN(2, m.Instance())
}

func (m *TRegistry) CloseKey1(key HKEY) {
	registryImportAPI().SysCallN(3, m.Instance(), uintptr(key))
}

func (m *TRegistry) GetKeyNames1(Strings IStrings) {
	registryImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TRegistry) GetValueNames1(Strings IStrings) {
	registryImportAPI().SysCallN(18, m.Instance(), GetObjectUintptr(Strings))
}

func (m *TRegistry) MoveKey(OldName, NewName string, Delete bool) {
	registryImportAPI().SysCallN(25, m.Instance(), PascalStr(OldName), PascalStr(NewName), PascalBool(Delete))
}

func (m *TRegistry) RenameValue(OldName, NewName string) {
	registryImportAPI().SysCallN(41, m.Instance(), PascalStr(OldName), PascalStr(NewName))
}

func (m *TRegistry) WriteCurrency(Name string, Value Currency) {
	registryImportAPI().SysCallN(51, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteBinaryData(Name string, Buffer []byte) {
	registryImportAPI().SysCallN(49, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&Buffer[0])), uintptr(len(Buffer)))
}

func (m *TRegistry) WriteBool(Name string, Value bool) {
	registryImportAPI().SysCallN(50, m.Instance(), PascalStr(Name), PascalBool(Value))
}

func (m *TRegistry) WriteDate(Name string, Value TDateTime) {
	registryImportAPI().SysCallN(52, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteDateTime(Name string, Value TDateTime) {
	registryImportAPI().SysCallN(53, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteFloat(Name string, Value float64) {
	registryImportAPI().SysCallN(55, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&Value)))
}

func (m *TRegistry) WriteInteger(Name string, Value int32) {
	registryImportAPI().SysCallN(57, m.Instance(), PascalStr(Name), uintptr(Value))
}

func (m *TRegistry) WriteInt64(Name string, Value int64) {
	registryImportAPI().SysCallN(56, m.Instance(), PascalStr(Name), uintptr(unsafePointer(&Value)))
}

func (m *TRegistry) WriteString(Name, Value string) {
	registryImportAPI().SysCallN(58, m.Instance(), PascalStr(Name), PascalStr(Value))
}

func (m *TRegistry) WriteExpandString(Name, Value string) {
	registryImportAPI().SysCallN(54, m.Instance(), PascalStr(Name), PascalStr(Value))
}

func (m *TRegistry) WriteStringArray(Name string, Arr TStringArray) {
	registryImportAPI().SysCallN(59, m.Instance(), PascalStr(Name), uintptr(Arr))
}

func (m *TRegistry) WriteTime(Name string, Value TDateTime) {
	registryImportAPI().SysCallN(60, m.Instance(), PascalStr(Name), uintptr(Value))
}

var (
	registryImport       *imports.Imports = nil
	registryImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Registry_Access", 0),
		/*1*/ imports.NewTable("Registry_Class", 0),
		/*2*/ imports.NewTable("Registry_CloseKey", 0),
		/*3*/ imports.NewTable("Registry_CloseKey1", 0),
		/*4*/ imports.NewTable("Registry_Create", 0),
		/*5*/ imports.NewTable("Registry_Create1", 0),
		/*6*/ imports.NewTable("Registry_CreateKey", 0),
		/*7*/ imports.NewTable("Registry_CurrentKey", 0),
		/*8*/ imports.NewTable("Registry_CurrentPath", 0),
		/*9*/ imports.NewTable("Registry_DeleteKey", 0),
		/*10*/ imports.NewTable("Registry_DeleteValue", 0),
		/*11*/ imports.NewTable("Registry_GetDataInfo", 0),
		/*12*/ imports.NewTable("Registry_GetDataSize", 0),
		/*13*/ imports.NewTable("Registry_GetDataType", 0),
		/*14*/ imports.NewTable("Registry_GetKeyInfo", 0),
		/*15*/ imports.NewTable("Registry_GetKeyNames", 0),
		/*16*/ imports.NewTable("Registry_GetKeyNames1", 0),
		/*17*/ imports.NewTable("Registry_GetValueNames", 0),
		/*18*/ imports.NewTable("Registry_GetValueNames1", 0),
		/*19*/ imports.NewTable("Registry_HasSubKeys", 0),
		/*20*/ imports.NewTable("Registry_KeyExists", 0),
		/*21*/ imports.NewTable("Registry_LastError", 0),
		/*22*/ imports.NewTable("Registry_LastErrorMsg", 0),
		/*23*/ imports.NewTable("Registry_LazyWrite", 0),
		/*24*/ imports.NewTable("Registry_LoadKey", 0),
		/*25*/ imports.NewTable("Registry_MoveKey", 0),
		/*26*/ imports.NewTable("Registry_OpenKey", 0),
		/*27*/ imports.NewTable("Registry_OpenKeyReadOnly", 0),
		/*28*/ imports.NewTable("Registry_ReadBinaryData", 0),
		/*29*/ imports.NewTable("Registry_ReadBool", 0),
		/*30*/ imports.NewTable("Registry_ReadCurrency", 0),
		/*31*/ imports.NewTable("Registry_ReadDate", 0),
		/*32*/ imports.NewTable("Registry_ReadDateTime", 0),
		/*33*/ imports.NewTable("Registry_ReadFloat", 0),
		/*34*/ imports.NewTable("Registry_ReadInt64", 0),
		/*35*/ imports.NewTable("Registry_ReadInteger", 0),
		/*36*/ imports.NewTable("Registry_ReadString", 0),
		/*37*/ imports.NewTable("Registry_ReadStringArray", 0),
		/*38*/ imports.NewTable("Registry_ReadStringList", 0),
		/*39*/ imports.NewTable("Registry_ReadTime", 0),
		/*40*/ imports.NewTable("Registry_RegistryConnect", 0),
		/*41*/ imports.NewTable("Registry_RenameValue", 0),
		/*42*/ imports.NewTable("Registry_ReplaceKey", 0),
		/*43*/ imports.NewTable("Registry_RestoreKey", 0),
		/*44*/ imports.NewTable("Registry_RootKey", 0),
		/*45*/ imports.NewTable("Registry_SaveKey", 0),
		/*46*/ imports.NewTable("Registry_StringSizeIncludesNull", 0),
		/*47*/ imports.NewTable("Registry_UnLoadKey", 0),
		/*48*/ imports.NewTable("Registry_ValueExists", 0),
		/*49*/ imports.NewTable("Registry_WriteBinaryData", 0),
		/*50*/ imports.NewTable("Registry_WriteBool", 0),
		/*51*/ imports.NewTable("Registry_WriteCurrency", 0),
		/*52*/ imports.NewTable("Registry_WriteDate", 0),
		/*53*/ imports.NewTable("Registry_WriteDateTime", 0),
		/*54*/ imports.NewTable("Registry_WriteExpandString", 0),
		/*55*/ imports.NewTable("Registry_WriteFloat", 0),
		/*56*/ imports.NewTable("Registry_WriteInt64", 0),
		/*57*/ imports.NewTable("Registry_WriteInteger", 0),
		/*58*/ imports.NewTable("Registry_WriteString", 0),
		/*59*/ imports.NewTable("Registry_WriteStringArray", 0),
		/*60*/ imports.NewTable("Registry_WriteTime", 0),
	}
)

func registryImportAPI() *imports.Imports {
	if registryImport == nil {
		registryImport = NewDefaultImports()
		registryImport.SetImportTable(registryImportTables)
		registryImportTables = nil
	}
	return registryImport
}
