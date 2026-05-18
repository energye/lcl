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

// IRegistry Parent: IObject
type IRegistry interface {
	IObject
	CreateKey(key string) bool                                                       // function
	DeleteKey(key string) bool                                                       // function
	DeleteValue(name string) bool                                                    // function
	GetDataInfo(valueName string, outValue *TRegDataInfo) bool                       // function
	GetDataSize(valueName string) int32                                              // function
	GetDataType(valueName string) types.TRegDataType                                 // function
	GetKeyInfo(outValue *TRegKeyInfo) bool                                           // function
	HasSubKeys() bool                                                                // function
	KeyExists(key string) bool                                                       // function
	LoadKey(key string, fileName string) bool                                        // function
	OpenKey(key string, canCreate bool) bool                                         // function
	OpenKeyReadOnly(key string) bool                                                 // function
	ReadCurrency(name string) uint64                                                 // function
	ReadBinaryData(name string, buffer *uintptr, bufSize int32) int32                // function
	ReadBool(name string) bool                                                       // function
	ReadDate(name string) types.TDateTime                                            // function
	ReadDateTime(name string) types.TDateTime                                        // function
	ReadFloat(name string) float64                                                   // function
	ReadInteger(name string) int32                                                   // function
	ReadInt64(name string) int64                                                     // function
	ReadString(name string) string                                                   // function
	ReadStringArray(name string) IStringArrayWrap                                    // function
	ReadTime(name string) types.TDateTime                                            // function
	RegistryConnect(uNCName string) bool                                             // function
	ReplaceKey(key string, fileName string, backUpFileName string) bool              // function
	RestoreKey(key string, fileName string) bool                                     // function
	SaveKey(key string, fileName string) bool                                        // function
	UnLoadKey(key string) bool                                                       // function
	ValueExists(name string) bool                                                    // function
	ReadStringListWithUStringStringsBool(name string, list IStrings, forceUtf8 bool) // procedure
	ReadStringListWithStrStrings(name string, list IStrings)                         // procedure
	CloseKey()                                                                       // procedure
	CloseKeyWithHKEY(key types.HKEY)                                                 // procedure
	GetKeyNames(strings IStrings)                                                    // procedure
	GetValueNames(strings IStrings)                                                  // procedure
	MoveKey(oldName string, newName string, delete bool)                             // procedure
	RenameValue(oldName string, newName string)                                      // procedure
	WriteCurrency(name string, value uint64)                                         // procedure
	WriteBinaryData(name string, buffer uintptr, bufSize int32)                      // procedure
	WriteBool(name string, value bool)                                               // procedure
	WriteDate(name string, value types.TDateTime)                                    // procedure
	WriteDateTime(name string, value types.TDateTime)                                // procedure
	WriteFloat(name string, value float64)                                           // procedure
	WriteInteger(name string, value int32)                                           // procedure
	WriteInt64(name string, value int64)                                             // procedure
	WriteString(name string, value string)                                           // procedure
	WriteExpandString(name string, value string)                                     // procedure
	WriteStringList(name string, list IStrings, isUtf8 bool)                         // procedure
	WriteStringArray(name string, arr IStringArrayWrap)                              // procedure
	WriteTime(name string, value types.TDateTime)                                    // procedure
	Access() uint32                                                                  // property Access Getter
	SetAccess(value uint32)                                                          // property Access Setter
	CurrentKey() types.HKEY                                                          // property CurrentKey Getter
	CurrentPath() string                                                             // property CurrentPath Getter
	LazyWrite() bool                                                                 // property LazyWrite Getter
	SetLazyWrite(value bool)                                                         // property LazyWrite Setter
	RootKey() types.HKEY                                                             // property RootKey Getter
	SetRootKey(value types.HKEY)                                                     // property RootKey Setter
	StringSizeIncludesNull() bool                                                    // property StringSizeIncludesNull Getter
	LastError() int32                                                                // property LastError Getter
	LastErrorMsg() string                                                            // property LastErrorMsg Getter
}

type TRegistry struct {
	TObject
}

func (m *TRegistry) CreateKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(2, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) DeleteKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(3, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) DeleteValue(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(4, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) GetDataInfo(valueName string, outValue *TRegDataInfo) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(5, m.Instance(), api.PasStr(valueName), uintptr(base.UnsafePointer(outValue)))
	return api.GoBool(r)
}

func (m *TRegistry) GetDataSize(valueName string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(6, m.Instance(), api.PasStr(valueName))
	return int32(r)
}

func (m *TRegistry) GetDataType(valueName string) types.TRegDataType {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(7, m.Instance(), api.PasStr(valueName))
	return types.TRegDataType(r)
}

func (m *TRegistry) GetKeyInfo(outValue *TRegKeyInfo) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(outValue)))
	return api.GoBool(r)
}

func (m *TRegistry) HasSubKeys() bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TRegistry) KeyExists(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(10, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) LoadKey(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(11, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) OpenKey(key string, canCreate bool) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(12, m.Instance(), api.PasStr(key), api.PasBool(canCreate))
	return api.GoBool(r)
}

func (m *TRegistry) OpenKeyReadOnly(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(13, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) ReadCurrency(name string) (result uint64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(14, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadBinaryData(name string, buffer *uintptr, bufSize int32) int32 {
	if !m.IsValid() {
		return 0
	}
	bufferPtr := uintptr(*buffer)
	r := registryAPI().SysCallN(15, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&bufferPtr)), uintptr(bufSize))
	*buffer = uintptr(bufferPtr)
	return int32(r)
}

func (m *TRegistry) ReadBool(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(16, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) ReadDate(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(17, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadDateTime(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(18, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadFloat(name string) (result float64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(19, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadInteger(name string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(20, m.Instance(), api.PasStr(name))
	return int32(r)
}

func (m *TRegistry) ReadInt64(name string) (result int64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(21, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadString(name string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	registryAPI().SysCallN(22, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TRegistry) ReadStringArray(name string) IStringArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := registryAPI().SysCallN(23, m.Instance(), api.PasStr(name))
	return AsStringArrayWrap(r)
}

func (m *TRegistry) ReadTime(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(24, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) RegistryConnect(uNCName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(25, m.Instance(), api.PasStr(uNCName))
	return api.GoBool(r)
}

func (m *TRegistry) ReplaceKey(key string, fileName string, backUpFileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(26, m.Instance(), api.PasStr(key), api.PasStr(fileName), api.PasStr(backUpFileName))
	return api.GoBool(r)
}

func (m *TRegistry) RestoreKey(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(27, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) SaveKey(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(28, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) UnLoadKey(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(29, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) ValueExists(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(30, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) ReadStringListWithUStringStringsBool(name string, list IStrings, forceUtf8 bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(31, m.Instance(), api.PasStr(name), base.GetObjectUintptr(list), api.PasBool(forceUtf8))
}

func (m *TRegistry) ReadStringListWithStrStrings(name string, list IStrings) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(32, m.Instance(), api.PasStr(name), base.GetObjectUintptr(list))
}

func (m *TRegistry) CloseKey() {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(33, m.Instance())
}

func (m *TRegistry) CloseKeyWithHKEY(key types.HKEY) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(34, m.Instance(), uintptr(key))
}

func (m *TRegistry) GetKeyNames(strings IStrings) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(35, m.Instance(), base.GetObjectUintptr(strings))
}

func (m *TRegistry) GetValueNames(strings IStrings) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(36, m.Instance(), base.GetObjectUintptr(strings))
}

func (m *TRegistry) MoveKey(oldName string, newName string, delete bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(37, m.Instance(), api.PasStr(oldName), api.PasStr(newName), api.PasBool(delete))
}

func (m *TRegistry) RenameValue(oldName string, newName string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(38, m.Instance(), api.PasStr(oldName), api.PasStr(newName))
}

func (m *TRegistry) WriteCurrency(name string, value uint64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(39, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteBinaryData(name string, buffer uintptr, bufSize int32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(40, m.Instance(), api.PasStr(name), uintptr(buffer), uintptr(bufSize))
}

func (m *TRegistry) WriteBool(name string, value bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(41, m.Instance(), api.PasStr(name), api.PasBool(value))
}

func (m *TRegistry) WriteDate(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(42, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteDateTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(43, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteFloat(name string, value float64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(44, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteInteger(name string, value int32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(45, m.Instance(), api.PasStr(name), uintptr(value))
}

func (m *TRegistry) WriteInt64(name string, value int64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(46, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteString(name string, value string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(47, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TRegistry) WriteExpandString(name string, value string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(48, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TRegistry) WriteStringList(name string, list IStrings, isUtf8 bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(49, m.Instance(), api.PasStr(name), base.GetObjectUintptr(list), api.PasBool(isUtf8))
}

func (m *TRegistry) WriteStringArray(name string, arr IStringArrayWrap) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(50, m.Instance(), api.PasStr(name), base.GetObjectUintptr(arr))
}

func (m *TRegistry) WriteTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(51, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) Access() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(52, 0, m.Instance())
	return uint32(r)
}

func (m *TRegistry) SetAccess(value uint32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(52, 1, m.Instance(), uintptr(value))
}

func (m *TRegistry) CurrentKey() types.HKEY {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(53, m.Instance())
	return types.HKEY(r)
}

func (m *TRegistry) CurrentPath() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	registryAPI().SysCallN(54, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TRegistry) LazyWrite() bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(55, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRegistry) SetLazyWrite(value bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(55, 1, m.Instance(), api.PasBool(value))
}

func (m *TRegistry) RootKey() types.HKEY {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(56, 0, m.Instance())
	return types.HKEY(r)
}

func (m *TRegistry) SetRootKey(value types.HKEY) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(56, 1, m.Instance(), uintptr(value))
}

func (m *TRegistry) StringSizeIncludesNull() bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(57, m.Instance())
	return api.GoBool(r)
}

func (m *TRegistry) LastError() int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(58, m.Instance())
	return int32(r)
}

func (m *TRegistry) LastErrorMsg() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	registryAPI().SysCallN(59, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

// NewRegistry class constructor
func NewRegistry() IRegistry {
	r := registryAPI().SysCallN(0)
	return AsRegistry(r)
}

// NewRegistryWithLongword class constructor
func NewRegistryWithLongword(aaccess uint32) IRegistry {
	r := registryAPI().SysCallN(1, uintptr(aaccess))
	return AsRegistry(r)
}

var (
	registryOnce   base.Once
	registryImport *imports.Imports = nil
)

func registryAPI() *imports.Imports {
	registryOnce.Do(func() {
		registryImport = api.NewDefaultImports()
		registryImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRegistry_Create", 0), // constructor NewRegistry
			/* 1 */ imports.NewTable("TRegistry_CreateWithLongword", 0), // constructor NewRegistryWithLongword
			/* 2 */ imports.NewTable("TRegistry_CreateKey", 0), // function CreateKey
			/* 3 */ imports.NewTable("TRegistry_DeleteKey", 0), // function DeleteKey
			/* 4 */ imports.NewTable("TRegistry_DeleteValue", 0), // function DeleteValue
			/* 5 */ imports.NewTable("TRegistry_GetDataInfo", 0), // function GetDataInfo
			/* 6 */ imports.NewTable("TRegistry_GetDataSize", 0), // function GetDataSize
			/* 7 */ imports.NewTable("TRegistry_GetDataType", 0), // function GetDataType
			/* 8 */ imports.NewTable("TRegistry_GetKeyInfo", 0), // function GetKeyInfo
			/* 9 */ imports.NewTable("TRegistry_HasSubKeys", 0), // function HasSubKeys
			/* 10 */ imports.NewTable("TRegistry_KeyExists", 0), // function KeyExists
			/* 11 */ imports.NewTable("TRegistry_LoadKey", 0), // function LoadKey
			/* 12 */ imports.NewTable("TRegistry_OpenKey", 0), // function OpenKey
			/* 13 */ imports.NewTable("TRegistry_OpenKeyReadOnly", 0), // function OpenKeyReadOnly
			/* 14 */ imports.NewTable("TRegistry_ReadCurrency", 0), // function ReadCurrency
			/* 15 */ imports.NewTable("TRegistry_ReadBinaryData", 0), // function ReadBinaryData
			/* 16 */ imports.NewTable("TRegistry_ReadBool", 0), // function ReadBool
			/* 17 */ imports.NewTable("TRegistry_ReadDate", 0), // function ReadDate
			/* 18 */ imports.NewTable("TRegistry_ReadDateTime", 0), // function ReadDateTime
			/* 19 */ imports.NewTable("TRegistry_ReadFloat", 0), // function ReadFloat
			/* 20 */ imports.NewTable("TRegistry_ReadInteger", 0), // function ReadInteger
			/* 21 */ imports.NewTable("TRegistry_ReadInt64", 0), // function ReadInt64
			/* 22 */ imports.NewTable("TRegistry_ReadString", 0), // function ReadString
			/* 23 */ imports.NewTable("TRegistry_ReadStringArray", 0), // function ReadStringArray
			/* 24 */ imports.NewTable("TRegistry_ReadTime", 0), // function ReadTime
			/* 25 */ imports.NewTable("TRegistry_RegistryConnect", 0), // function RegistryConnect
			/* 26 */ imports.NewTable("TRegistry_ReplaceKey", 0), // function ReplaceKey
			/* 27 */ imports.NewTable("TRegistry_RestoreKey", 0), // function RestoreKey
			/* 28 */ imports.NewTable("TRegistry_SaveKey", 0), // function SaveKey
			/* 29 */ imports.NewTable("TRegistry_UnLoadKey", 0), // function UnLoadKey
			/* 30 */ imports.NewTable("TRegistry_ValueExists", 0), // function ValueExists
			/* 31 */ imports.NewTable("TRegistry_ReadStringListWithUStringStringsBool", 0), // procedure ReadStringListWithUStringStringsBool
			/* 32 */ imports.NewTable("TRegistry_ReadStringListWithStrStrings", 0), // procedure ReadStringListWithStrStrings
			/* 33 */ imports.NewTable("TRegistry_CloseKey", 0), // procedure CloseKey
			/* 34 */ imports.NewTable("TRegistry_CloseKeyWithHKEY", 0), // procedure CloseKeyWithHKEY
			/* 35 */ imports.NewTable("TRegistry_GetKeyNames", 0), // procedure GetKeyNames
			/* 36 */ imports.NewTable("TRegistry_GetValueNames", 0), // procedure GetValueNames
			/* 37 */ imports.NewTable("TRegistry_MoveKey", 0), // procedure MoveKey
			/* 38 */ imports.NewTable("TRegistry_RenameValue", 0), // procedure RenameValue
			/* 39 */ imports.NewTable("TRegistry_WriteCurrency", 0), // procedure WriteCurrency
			/* 40 */ imports.NewTable("TRegistry_WriteBinaryData", 0), // procedure WriteBinaryData
			/* 41 */ imports.NewTable("TRegistry_WriteBool", 0), // procedure WriteBool
			/* 42 */ imports.NewTable("TRegistry_WriteDate", 0), // procedure WriteDate
			/* 43 */ imports.NewTable("TRegistry_WriteDateTime", 0), // procedure WriteDateTime
			/* 44 */ imports.NewTable("TRegistry_WriteFloat", 0), // procedure WriteFloat
			/* 45 */ imports.NewTable("TRegistry_WriteInteger", 0), // procedure WriteInteger
			/* 46 */ imports.NewTable("TRegistry_WriteInt64", 0), // procedure WriteInt64
			/* 47 */ imports.NewTable("TRegistry_WriteString", 0), // procedure WriteString
			/* 48 */ imports.NewTable("TRegistry_WriteExpandString", 0), // procedure WriteExpandString
			/* 49 */ imports.NewTable("TRegistry_WriteStringList", 0), // procedure WriteStringList
			/* 50 */ imports.NewTable("TRegistry_WriteStringArray", 0), // procedure WriteStringArray
			/* 51 */ imports.NewTable("TRegistry_WriteTime", 0), // procedure WriteTime
			/* 52 */ imports.NewTable("TRegistry_Access", 0), // property Access
			/* 53 */ imports.NewTable("TRegistry_CurrentKey", 0), // property CurrentKey
			/* 54 */ imports.NewTable("TRegistry_CurrentPath", 0), // property CurrentPath
			/* 55 */ imports.NewTable("TRegistry_LazyWrite", 0), // property LazyWrite
			/* 56 */ imports.NewTable("TRegistry_RootKey", 0), // property RootKey
			/* 57 */ imports.NewTable("TRegistry_StringSizeIncludesNull", 0), // property StringSizeIncludesNull
			/* 58 */ imports.NewTable("TRegistry_LastError", 0), // property LastError
			/* 59 */ imports.NewTable("TRegistry_LastErrorMsg", 0), // property LastErrorMsg
		}
	})
	return registryImport
}
