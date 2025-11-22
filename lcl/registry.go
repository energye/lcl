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
	CreateKeyWithUnicodeString(key string) bool                                                  // function
	CreateKeyWithString(key string) bool                                                         // function
	DeleteKeyWithUnicodeString(key string) bool                                                  // function
	DeleteKeyWithString(key string) bool                                                         // function
	DeleteValueWithUnicodeString(name string) bool                                               // function
	DeleteValueWithString(name string) bool                                                      // function
	GetDataInfoWithUnicodeStringRegDataInfo(valueName string, outValue *TRegDataInfo) bool       // function
	GetDataInfoWithStringRegDataInfo(valueName string, outValue *TRegDataInfo) bool              // function
	GetDataSizeWithUnicodeString(valueName string) int32                                         // function
	GetDataSizeWithString(valueName string) int32                                                // function
	GetDataTypeWithUnicodeString(valueName string) types.TRegDataType                            // function
	GetDataTypeWithString(valueName string) types.TRegDataType                                   // function
	GetKeyInfo(outValue *TRegKeyInfo) bool                                                       // function
	HasSubKeys() bool                                                                            // function
	KeyExistsWithUnicodeString(key string) bool                                                  // function
	KeyExistsWithString(key string) bool                                                         // function
	LoadKeyWithUnicodeStringX2(key string, fileName string) bool                                 // function
	LoadKeyWithStringX2(key string, fileName string) bool                                        // function
	OpenKeyWithUnicodeStringBool(key string, canCreate bool) bool                                // function
	OpenKeyWithStringBool(key string, canCreate bool) bool                                       // function
	OpenKeyReadOnlyWithUnicodeString(key string) bool                                            // function
	OpenKeyReadOnlyWithString(key string) bool                                                   // function
	ReadCurrencyWithUnicodeString(name string) uint64                                            // function
	ReadCurrencyWithString(name string) uint64                                                   // function
	ReadBinaryDataWithUnicodeStringPointerInt(name string, buffer *uintptr, bufSize int32) int32 // function
	ReadBinaryDataWithStringPointerInt(name string, buffer *uintptr, bufSize int32) int32        // function
	ReadBoolWithUnicodeString(name string) bool                                                  // function
	ReadBoolWithString(name string) bool                                                         // function
	ReadDateWithUnicodeString(name string) types.TDateTime                                       // function
	ReadDateWithString(name string) types.TDateTime                                              // function
	ReadDateTimeWithUnicodeString(name string) types.TDateTime                                   // function
	ReadDateTimeWithString(name string) types.TDateTime                                          // function
	ReadFloatWithUnicodeString(name string) float64                                              // function
	ReadFloatWithString(name string) float64                                                     // function
	ReadIntegerWithUnicodeString(name string) int32                                              // function
	ReadIntegerWithString(name string) int32                                                     // function
	ReadInt64WithUnicodeString(name string) int64                                                // function
	ReadInt64WithString(name string) int64                                                       // function
	ReadStringWithUnicodeString(name string) string                                              // function
	ReadStringWithString(name string) string                                                     // function
	ReadStringArray(name string) IStringArrayWrap                                                // function
	ReadTimeWithUnicodeString(name string) types.TDateTime                                       // function
	ReadTimeWithString(name string) types.TDateTime                                              // function
	RegistryConnectWithUnicodeString(uNCName string) bool                                        // function
	RegistryConnectWithString(uNCName string) bool                                               // function
	ReplaceKeyWithUnicodeStringX3(key string, fileName string, backUpFileName string) bool       // function
	ReplaceKeyWithStringX3(key string, fileName string, backUpFileName string) bool              // function
	RestoreKeyWithUnicodeStringX2(key string, fileName string) bool                              // function
	RestoreKeyWithStringX2(key string, fileName string) bool                                     // function
	SaveKeyWithUnicodeStringX2(key string, fileName string) bool                                 // function
	SaveKeyWithStringX2(key string, fileName string) bool                                        // function
	UnLoadKeyWithUnicodeString(key string) bool                                                  // function
	UnLoadKeyWithString(key string) bool                                                         // function
	ValueExistsWithUnicodeString(name string) bool                                               // function
	ValueExistsWithString(name string) bool                                                      // function
	ReadStringListWithUnicodeStringStringsBool(name string, list IStrings, forceUtf8 bool)       // procedure
	ReadStringListWithStringStrings(name string, list IStrings)                                  // procedure
	CloseKey()                                                                                   // procedure
	CloseKeyWithHKEY(key types.HKEY)                                                             // procedure
	GetKeyNames(strings IStrings)                                                                // procedure
	GetValueNames(strings IStrings)                                                              // procedure
	MoveKeyWithUnicodeStringX2Bool(oldName string, newName string, delete bool)                  // procedure
	MoveKeyWithStringX2Bool(oldName string, newName string, delete bool)                         // procedure
	RenameValueWithUnicodeStringX2(oldName string, newName string)                               // procedure
	RenameValueWithStringX2(oldName string, newName string)                                      // procedure
	WriteCurrencyWithUnicodeStringCurrency(name string, value uint64)                            // procedure
	WriteCurrencyWithStringCurrency(name string, value uint64)                                   // procedure
	WriteBinaryDataWithUnicodeStringPointerInt(name string, buffer uintptr, bufSize int32)       // procedure
	WriteBinaryDataWithStringPointerInt(name string, buffer uintptr, bufSize int32)              // procedure
	WriteBoolWithUnicodeStringBool(name string, value bool)                                      // procedure
	WriteBoolWithStringBool(name string, value bool)                                             // procedure
	WriteDateWithUnicodeStringDateTime(name string, value types.TDateTime)                       // procedure
	WriteDateWithStringDateTime(name string, value types.TDateTime)                              // procedure
	WriteDateTimeWithUnicodeStringDateTime(name string, value types.TDateTime)                   // procedure
	WriteDateTimeWithStringDateTime(name string, value types.TDateTime)                          // procedure
	WriteFloatWithUnicodeStringDouble(name string, value float64)                                // procedure
	WriteFloatWithStringDouble(name string, value float64)                                       // procedure
	WriteIntegerWithUnicodeStringInt(name string, value int32)                                   // procedure
	WriteIntegerWithStringInt(name string, value int32)                                          // procedure
	WriteInt64WithUnicodeStringInt64(name string, value int64)                                   // procedure
	WriteInt64WithStringInt64(name string, value int64)                                          // procedure
	WriteStringWithUnicodeStringX2(name string, value string)                                    // procedure
	WriteStringWithStringX2(name string, value string)                                           // procedure
	WriteExpandStringWithUnicodeStringX2(name string, value string)                              // procedure
	WriteExpandStringWithStringX2(name string, value string)                                     // procedure
	WriteStringList(name string, list IStrings, isUtf8 bool)                                     // procedure
	WriteStringArray(name string, arr IStringArrayWrap)                                          // procedure
	WriteTimeWithUnicodeStringDateTime(name string, value types.TDateTime)                       // procedure
	WriteTimeWithStringDateTime(name string, value types.TDateTime)                              // procedure
	Access() uint32                                                                              // property Access Getter
	SetAccess(value uint32)                                                                      // property Access Setter
	CurrentKey() types.HKEY                                                                      // property CurrentKey Getter
	CurrentPath() string                                                                         // property CurrentPath Getter
	LazyWrite() bool                                                                             // property LazyWrite Getter
	SetLazyWrite(value bool)                                                                     // property LazyWrite Setter
	RootKey() types.HKEY                                                                         // property RootKey Getter
	SetRootKey(value types.HKEY)                                                                 // property RootKey Setter
	StringSizeIncludesNull() bool                                                                // property StringSizeIncludesNull Getter
	LastError() int32                                                                            // property LastError Getter
	LastErrorMsg() string                                                                        // property LastErrorMsg Getter
}

type TRegistry struct {
	TObject
}

func (m *TRegistry) CreateKeyWithUnicodeString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(2, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) CreateKeyWithString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(3, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) DeleteKeyWithUnicodeString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(4, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) DeleteKeyWithString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(5, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) DeleteValueWithUnicodeString(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(6, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) DeleteValueWithString(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(7, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) GetDataInfoWithUnicodeStringRegDataInfo(valueName string, outValue *TRegDataInfo) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(8, m.Instance(), api.PasStr(valueName), uintptr(base.UnsafePointer(outValue)))
	return api.GoBool(r)
}

func (m *TRegistry) GetDataInfoWithStringRegDataInfo(valueName string, outValue *TRegDataInfo) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(9, m.Instance(), api.PasStr(valueName), uintptr(base.UnsafePointer(outValue)))
	return api.GoBool(r)
}

func (m *TRegistry) GetDataSizeWithUnicodeString(valueName string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(10, m.Instance(), api.PasStr(valueName))
	return int32(r)
}

func (m *TRegistry) GetDataSizeWithString(valueName string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(11, m.Instance(), api.PasStr(valueName))
	return int32(r)
}

func (m *TRegistry) GetDataTypeWithUnicodeString(valueName string) types.TRegDataType {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(12, m.Instance(), api.PasStr(valueName))
	return types.TRegDataType(r)
}

func (m *TRegistry) GetDataTypeWithString(valueName string) types.TRegDataType {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(13, m.Instance(), api.PasStr(valueName))
	return types.TRegDataType(r)
}

func (m *TRegistry) GetKeyInfo(outValue *TRegKeyInfo) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(outValue)))
	return api.GoBool(r)
}

func (m *TRegistry) HasSubKeys() bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(15, m.Instance())
	return api.GoBool(r)
}

func (m *TRegistry) KeyExistsWithUnicodeString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(16, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) KeyExistsWithString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(17, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) LoadKeyWithUnicodeStringX2(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(18, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) LoadKeyWithStringX2(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(19, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) OpenKeyWithUnicodeStringBool(key string, canCreate bool) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(20, m.Instance(), api.PasStr(key), api.PasBool(canCreate))
	return api.GoBool(r)
}

func (m *TRegistry) OpenKeyWithStringBool(key string, canCreate bool) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(21, m.Instance(), api.PasStr(key), api.PasBool(canCreate))
	return api.GoBool(r)
}

func (m *TRegistry) OpenKeyReadOnlyWithUnicodeString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(22, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) OpenKeyReadOnlyWithString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(23, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) ReadCurrencyWithUnicodeString(name string) (result uint64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(24, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadCurrencyWithString(name string) (result uint64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(25, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadBinaryDataWithUnicodeStringPointerInt(name string, buffer *uintptr, bufSize int32) int32 {
	if !m.IsValid() {
		return 0
	}
	bufferPtr := uintptr(*buffer)
	r := registryAPI().SysCallN(26, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&bufferPtr)), uintptr(bufSize))
	*buffer = uintptr(bufferPtr)
	return int32(r)
}

func (m *TRegistry) ReadBinaryDataWithStringPointerInt(name string, buffer *uintptr, bufSize int32) int32 {
	if !m.IsValid() {
		return 0
	}
	bufferPtr := uintptr(*buffer)
	r := registryAPI().SysCallN(27, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&bufferPtr)), uintptr(bufSize))
	*buffer = uintptr(bufferPtr)
	return int32(r)
}

func (m *TRegistry) ReadBoolWithUnicodeString(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(28, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) ReadBoolWithString(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(29, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) ReadDateWithUnicodeString(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(30, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadDateWithString(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(31, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadDateTimeWithUnicodeString(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(32, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadDateTimeWithString(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(33, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadFloatWithUnicodeString(name string) (result float64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(34, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadFloatWithString(name string) (result float64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(35, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadIntegerWithUnicodeString(name string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(36, m.Instance(), api.PasStr(name))
	return int32(r)
}

func (m *TRegistry) ReadIntegerWithString(name string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(37, m.Instance(), api.PasStr(name))
	return int32(r)
}

func (m *TRegistry) ReadInt64WithUnicodeString(name string) (result int64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(38, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadInt64WithString(name string) (result int64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(39, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadStringWithUnicodeString(name string) string {
	if !m.IsValid() {
		return ""
	}
	r := registryAPI().SysCallN(40, m.Instance(), api.PasStr(name))
	return api.GoStr(r)
}

func (m *TRegistry) ReadStringWithString(name string) string {
	if !m.IsValid() {
		return ""
	}
	r := registryAPI().SysCallN(41, m.Instance(), api.PasStr(name))
	return api.GoStr(r)
}

func (m *TRegistry) ReadStringArray(name string) IStringArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := registryAPI().SysCallN(42, m.Instance(), api.PasStr(name))
	return AsStringArrayWrap(r)
}

func (m *TRegistry) ReadTimeWithUnicodeString(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(43, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) ReadTimeWithString(name string) (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(44, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TRegistry) RegistryConnectWithUnicodeString(uNCName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(45, m.Instance(), api.PasStr(uNCName))
	return api.GoBool(r)
}

func (m *TRegistry) RegistryConnectWithString(uNCName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(46, m.Instance(), api.PasStr(uNCName))
	return api.GoBool(r)
}

func (m *TRegistry) ReplaceKeyWithUnicodeStringX3(key string, fileName string, backUpFileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(47, m.Instance(), api.PasStr(key), api.PasStr(fileName), api.PasStr(backUpFileName))
	return api.GoBool(r)
}

func (m *TRegistry) ReplaceKeyWithStringX3(key string, fileName string, backUpFileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(48, m.Instance(), api.PasStr(key), api.PasStr(fileName), api.PasStr(backUpFileName))
	return api.GoBool(r)
}

func (m *TRegistry) RestoreKeyWithUnicodeStringX2(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(49, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) RestoreKeyWithStringX2(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(50, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) SaveKeyWithUnicodeStringX2(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(51, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) SaveKeyWithStringX2(key string, fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(52, m.Instance(), api.PasStr(key), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TRegistry) UnLoadKeyWithUnicodeString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(53, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) UnLoadKeyWithString(key string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(54, m.Instance(), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TRegistry) ValueExistsWithUnicodeString(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(55, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) ValueExistsWithString(name string) bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(56, m.Instance(), api.PasStr(name))
	return api.GoBool(r)
}

func (m *TRegistry) ReadStringListWithUnicodeStringStringsBool(name string, list IStrings, forceUtf8 bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(57, m.Instance(), api.PasStr(name), base.GetObjectUintptr(list), api.PasBool(forceUtf8))
}

func (m *TRegistry) ReadStringListWithStringStrings(name string, list IStrings) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(58, m.Instance(), api.PasStr(name), base.GetObjectUintptr(list))
}

func (m *TRegistry) CloseKey() {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(59, m.Instance())
}

func (m *TRegistry) CloseKeyWithHKEY(key types.HKEY) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(60, m.Instance(), uintptr(key))
}

func (m *TRegistry) GetKeyNames(strings IStrings) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(61, m.Instance(), base.GetObjectUintptr(strings))
}

func (m *TRegistry) GetValueNames(strings IStrings) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(62, m.Instance(), base.GetObjectUintptr(strings))
}

func (m *TRegistry) MoveKeyWithUnicodeStringX2Bool(oldName string, newName string, delete bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(63, m.Instance(), api.PasStr(oldName), api.PasStr(newName), api.PasBool(delete))
}

func (m *TRegistry) MoveKeyWithStringX2Bool(oldName string, newName string, delete bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(64, m.Instance(), api.PasStr(oldName), api.PasStr(newName), api.PasBool(delete))
}

func (m *TRegistry) RenameValueWithUnicodeStringX2(oldName string, newName string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(65, m.Instance(), api.PasStr(oldName), api.PasStr(newName))
}

func (m *TRegistry) RenameValueWithStringX2(oldName string, newName string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(66, m.Instance(), api.PasStr(oldName), api.PasStr(newName))
}

func (m *TRegistry) WriteCurrencyWithUnicodeStringCurrency(name string, value uint64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(67, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteCurrencyWithStringCurrency(name string, value uint64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(68, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteBinaryDataWithUnicodeStringPointerInt(name string, buffer uintptr, bufSize int32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(69, m.Instance(), api.PasStr(name), uintptr(buffer), uintptr(bufSize))
}

func (m *TRegistry) WriteBinaryDataWithStringPointerInt(name string, buffer uintptr, bufSize int32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(70, m.Instance(), api.PasStr(name), uintptr(buffer), uintptr(bufSize))
}

func (m *TRegistry) WriteBoolWithUnicodeStringBool(name string, value bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(71, m.Instance(), api.PasStr(name), api.PasBool(value))
}

func (m *TRegistry) WriteBoolWithStringBool(name string, value bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(72, m.Instance(), api.PasStr(name), api.PasBool(value))
}

func (m *TRegistry) WriteDateWithUnicodeStringDateTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(73, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteDateWithStringDateTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(74, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteDateTimeWithUnicodeStringDateTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(75, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteDateTimeWithStringDateTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(76, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteFloatWithUnicodeStringDouble(name string, value float64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(77, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteFloatWithStringDouble(name string, value float64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(78, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteIntegerWithUnicodeStringInt(name string, value int32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(79, m.Instance(), api.PasStr(name), uintptr(value))
}

func (m *TRegistry) WriteIntegerWithStringInt(name string, value int32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(80, m.Instance(), api.PasStr(name), uintptr(value))
}

func (m *TRegistry) WriteInt64WithUnicodeStringInt64(name string, value int64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(81, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteInt64WithStringInt64(name string, value int64) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(82, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteStringWithUnicodeStringX2(name string, value string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(83, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TRegistry) WriteStringWithStringX2(name string, value string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(84, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TRegistry) WriteExpandStringWithUnicodeStringX2(name string, value string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(85, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TRegistry) WriteExpandStringWithStringX2(name string, value string) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(86, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TRegistry) WriteStringList(name string, list IStrings, isUtf8 bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(87, m.Instance(), api.PasStr(name), base.GetObjectUintptr(list), api.PasBool(isUtf8))
}

func (m *TRegistry) WriteStringArray(name string, arr IStringArrayWrap) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(88, m.Instance(), api.PasStr(name), base.GetObjectUintptr(arr))
}

func (m *TRegistry) WriteTimeWithUnicodeStringDateTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(89, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) WriteTimeWithStringDateTime(name string, value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(90, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&value)))
}

func (m *TRegistry) Access() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(91, 0, m.Instance())
	return uint32(r)
}

func (m *TRegistry) SetAccess(value uint32) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(91, 1, m.Instance(), uintptr(value))
}

func (m *TRegistry) CurrentKey() types.HKEY {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(92, m.Instance())
	return types.HKEY(r)
}

func (m *TRegistry) CurrentPath() string {
	if !m.IsValid() {
		return ""
	}
	r := registryAPI().SysCallN(93, m.Instance())
	return api.GoStr(r)
}

func (m *TRegistry) LazyWrite() bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(94, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRegistry) SetLazyWrite(value bool) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(94, 1, m.Instance(), api.PasBool(value))
}

func (m *TRegistry) RootKey() types.HKEY {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(95, 0, m.Instance())
	return types.HKEY(r)
}

func (m *TRegistry) SetRootKey(value types.HKEY) {
	if !m.IsValid() {
		return
	}
	registryAPI().SysCallN(95, 1, m.Instance(), uintptr(value))
}

func (m *TRegistry) StringSizeIncludesNull() bool {
	if !m.IsValid() {
		return false
	}
	r := registryAPI().SysCallN(96, m.Instance())
	return api.GoBool(r)
}

func (m *TRegistry) LastError() int32 {
	if !m.IsValid() {
		return 0
	}
	r := registryAPI().SysCallN(97, m.Instance())
	return int32(r)
}

func (m *TRegistry) LastErrorMsg() string {
	if !m.IsValid() {
		return ""
	}
	r := registryAPI().SysCallN(98, m.Instance())
	return api.GoStr(r)
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
			/* 2 */ imports.NewTable("TRegistry_CreateKeyWithUnicodeString", 0), // function CreateKeyWithUnicodeString
			/* 3 */ imports.NewTable("TRegistry_CreateKeyWithString", 0), // function CreateKeyWithString
			/* 4 */ imports.NewTable("TRegistry_DeleteKeyWithUnicodeString", 0), // function DeleteKeyWithUnicodeString
			/* 5 */ imports.NewTable("TRegistry_DeleteKeyWithString", 0), // function DeleteKeyWithString
			/* 6 */ imports.NewTable("TRegistry_DeleteValueWithUnicodeString", 0), // function DeleteValueWithUnicodeString
			/* 7 */ imports.NewTable("TRegistry_DeleteValueWithString", 0), // function DeleteValueWithString
			/* 8 */ imports.NewTable("TRegistry_GetDataInfoWithUnicodeStringRegDataInfo", 0), // function GetDataInfoWithUnicodeStringRegDataInfo
			/* 9 */ imports.NewTable("TRegistry_GetDataInfoWithStringRegDataInfo", 0), // function GetDataInfoWithStringRegDataInfo
			/* 10 */ imports.NewTable("TRegistry_GetDataSizeWithUnicodeString", 0), // function GetDataSizeWithUnicodeString
			/* 11 */ imports.NewTable("TRegistry_GetDataSizeWithString", 0), // function GetDataSizeWithString
			/* 12 */ imports.NewTable("TRegistry_GetDataTypeWithUnicodeString", 0), // function GetDataTypeWithUnicodeString
			/* 13 */ imports.NewTable("TRegistry_GetDataTypeWithString", 0), // function GetDataTypeWithString
			/* 14 */ imports.NewTable("TRegistry_GetKeyInfo", 0), // function GetKeyInfo
			/* 15 */ imports.NewTable("TRegistry_HasSubKeys", 0), // function HasSubKeys
			/* 16 */ imports.NewTable("TRegistry_KeyExistsWithUnicodeString", 0), // function KeyExistsWithUnicodeString
			/* 17 */ imports.NewTable("TRegistry_KeyExistsWithString", 0), // function KeyExistsWithString
			/* 18 */ imports.NewTable("TRegistry_LoadKeyWithUnicodeStringX2", 0), // function LoadKeyWithUnicodeStringX2
			/* 19 */ imports.NewTable("TRegistry_LoadKeyWithStringX2", 0), // function LoadKeyWithStringX2
			/* 20 */ imports.NewTable("TRegistry_OpenKeyWithUnicodeStringBool", 0), // function OpenKeyWithUnicodeStringBool
			/* 21 */ imports.NewTable("TRegistry_OpenKeyWithStringBool", 0), // function OpenKeyWithStringBool
			/* 22 */ imports.NewTable("TRegistry_OpenKeyReadOnlyWithUnicodeString", 0), // function OpenKeyReadOnlyWithUnicodeString
			/* 23 */ imports.NewTable("TRegistry_OpenKeyReadOnlyWithString", 0), // function OpenKeyReadOnlyWithString
			/* 24 */ imports.NewTable("TRegistry_ReadCurrencyWithUnicodeString", 0), // function ReadCurrencyWithUnicodeString
			/* 25 */ imports.NewTable("TRegistry_ReadCurrencyWithString", 0), // function ReadCurrencyWithString
			/* 26 */ imports.NewTable("TRegistry_ReadBinaryDataWithUnicodeStringPointerInt", 0), // function ReadBinaryDataWithUnicodeStringPointerInt
			/* 27 */ imports.NewTable("TRegistry_ReadBinaryDataWithStringPointerInt", 0), // function ReadBinaryDataWithStringPointerInt
			/* 28 */ imports.NewTable("TRegistry_ReadBoolWithUnicodeString", 0), // function ReadBoolWithUnicodeString
			/* 29 */ imports.NewTable("TRegistry_ReadBoolWithString", 0), // function ReadBoolWithString
			/* 30 */ imports.NewTable("TRegistry_ReadDateWithUnicodeString", 0), // function ReadDateWithUnicodeString
			/* 31 */ imports.NewTable("TRegistry_ReadDateWithString", 0), // function ReadDateWithString
			/* 32 */ imports.NewTable("TRegistry_ReadDateTimeWithUnicodeString", 0), // function ReadDateTimeWithUnicodeString
			/* 33 */ imports.NewTable("TRegistry_ReadDateTimeWithString", 0), // function ReadDateTimeWithString
			/* 34 */ imports.NewTable("TRegistry_ReadFloatWithUnicodeString", 0), // function ReadFloatWithUnicodeString
			/* 35 */ imports.NewTable("TRegistry_ReadFloatWithString", 0), // function ReadFloatWithString
			/* 36 */ imports.NewTable("TRegistry_ReadIntegerWithUnicodeString", 0), // function ReadIntegerWithUnicodeString
			/* 37 */ imports.NewTable("TRegistry_ReadIntegerWithString", 0), // function ReadIntegerWithString
			/* 38 */ imports.NewTable("TRegistry_ReadInt64WithUnicodeString", 0), // function ReadInt64WithUnicodeString
			/* 39 */ imports.NewTable("TRegistry_ReadInt64WithString", 0), // function ReadInt64WithString
			/* 40 */ imports.NewTable("TRegistry_ReadStringWithUnicodeString", 0), // function ReadStringWithUnicodeString
			/* 41 */ imports.NewTable("TRegistry_ReadStringWithString", 0), // function ReadStringWithString
			/* 42 */ imports.NewTable("TRegistry_ReadStringArray", 0), // function ReadStringArray
			/* 43 */ imports.NewTable("TRegistry_ReadTimeWithUnicodeString", 0), // function ReadTimeWithUnicodeString
			/* 44 */ imports.NewTable("TRegistry_ReadTimeWithString", 0), // function ReadTimeWithString
			/* 45 */ imports.NewTable("TRegistry_RegistryConnectWithUnicodeString", 0), // function RegistryConnectWithUnicodeString
			/* 46 */ imports.NewTable("TRegistry_RegistryConnectWithString", 0), // function RegistryConnectWithString
			/* 47 */ imports.NewTable("TRegistry_ReplaceKeyWithUnicodeStringX3", 0), // function ReplaceKeyWithUnicodeStringX3
			/* 48 */ imports.NewTable("TRegistry_ReplaceKeyWithStringX3", 0), // function ReplaceKeyWithStringX3
			/* 49 */ imports.NewTable("TRegistry_RestoreKeyWithUnicodeStringX2", 0), // function RestoreKeyWithUnicodeStringX2
			/* 50 */ imports.NewTable("TRegistry_RestoreKeyWithStringX2", 0), // function RestoreKeyWithStringX2
			/* 51 */ imports.NewTable("TRegistry_SaveKeyWithUnicodeStringX2", 0), // function SaveKeyWithUnicodeStringX2
			/* 52 */ imports.NewTable("TRegistry_SaveKeyWithStringX2", 0), // function SaveKeyWithStringX2
			/* 53 */ imports.NewTable("TRegistry_UnLoadKeyWithUnicodeString", 0), // function UnLoadKeyWithUnicodeString
			/* 54 */ imports.NewTable("TRegistry_UnLoadKeyWithString", 0), // function UnLoadKeyWithString
			/* 55 */ imports.NewTable("TRegistry_ValueExistsWithUnicodeString", 0), // function ValueExistsWithUnicodeString
			/* 56 */ imports.NewTable("TRegistry_ValueExistsWithString", 0), // function ValueExistsWithString
			/* 57 */ imports.NewTable("TRegistry_ReadStringListWithUnicodeStringStringsBool", 0), // procedure ReadStringListWithUnicodeStringStringsBool
			/* 58 */ imports.NewTable("TRegistry_ReadStringListWithStringStrings", 0), // procedure ReadStringListWithStringStrings
			/* 59 */ imports.NewTable("TRegistry_CloseKey", 0), // procedure CloseKey
			/* 60 */ imports.NewTable("TRegistry_CloseKeyWithHKEY", 0), // procedure CloseKeyWithHKEY
			/* 61 */ imports.NewTable("TRegistry_GetKeyNames", 0), // procedure GetKeyNames
			/* 62 */ imports.NewTable("TRegistry_GetValueNames", 0), // procedure GetValueNames
			/* 63 */ imports.NewTable("TRegistry_MoveKeyWithUnicodeStringX2Bool", 0), // procedure MoveKeyWithUnicodeStringX2Bool
			/* 64 */ imports.NewTable("TRegistry_MoveKeyWithStringX2Bool", 0), // procedure MoveKeyWithStringX2Bool
			/* 65 */ imports.NewTable("TRegistry_RenameValueWithUnicodeStringX2", 0), // procedure RenameValueWithUnicodeStringX2
			/* 66 */ imports.NewTable("TRegistry_RenameValueWithStringX2", 0), // procedure RenameValueWithStringX2
			/* 67 */ imports.NewTable("TRegistry_WriteCurrencyWithUnicodeStringCurrency", 0), // procedure WriteCurrencyWithUnicodeStringCurrency
			/* 68 */ imports.NewTable("TRegistry_WriteCurrencyWithStringCurrency", 0), // procedure WriteCurrencyWithStringCurrency
			/* 69 */ imports.NewTable("TRegistry_WriteBinaryDataWithUnicodeStringPointerInt", 0), // procedure WriteBinaryDataWithUnicodeStringPointerInt
			/* 70 */ imports.NewTable("TRegistry_WriteBinaryDataWithStringPointerInt", 0), // procedure WriteBinaryDataWithStringPointerInt
			/* 71 */ imports.NewTable("TRegistry_WriteBoolWithUnicodeStringBool", 0), // procedure WriteBoolWithUnicodeStringBool
			/* 72 */ imports.NewTable("TRegistry_WriteBoolWithStringBool", 0), // procedure WriteBoolWithStringBool
			/* 73 */ imports.NewTable("TRegistry_WriteDateWithUnicodeStringDateTime", 0), // procedure WriteDateWithUnicodeStringDateTime
			/* 74 */ imports.NewTable("TRegistry_WriteDateWithStringDateTime", 0), // procedure WriteDateWithStringDateTime
			/* 75 */ imports.NewTable("TRegistry_WriteDateTimeWithUnicodeStringDateTime", 0), // procedure WriteDateTimeWithUnicodeStringDateTime
			/* 76 */ imports.NewTable("TRegistry_WriteDateTimeWithStringDateTime", 0), // procedure WriteDateTimeWithStringDateTime
			/* 77 */ imports.NewTable("TRegistry_WriteFloatWithUnicodeStringDouble", 0), // procedure WriteFloatWithUnicodeStringDouble
			/* 78 */ imports.NewTable("TRegistry_WriteFloatWithStringDouble", 0), // procedure WriteFloatWithStringDouble
			/* 79 */ imports.NewTable("TRegistry_WriteIntegerWithUnicodeStringInt", 0), // procedure WriteIntegerWithUnicodeStringInt
			/* 80 */ imports.NewTable("TRegistry_WriteIntegerWithStringInt", 0), // procedure WriteIntegerWithStringInt
			/* 81 */ imports.NewTable("TRegistry_WriteInt64WithUnicodeStringInt64", 0), // procedure WriteInt64WithUnicodeStringInt64
			/* 82 */ imports.NewTable("TRegistry_WriteInt64WithStringInt64", 0), // procedure WriteInt64WithStringInt64
			/* 83 */ imports.NewTable("TRegistry_WriteStringWithUnicodeStringX2", 0), // procedure WriteStringWithUnicodeStringX2
			/* 84 */ imports.NewTable("TRegistry_WriteStringWithStringX2", 0), // procedure WriteStringWithStringX2
			/* 85 */ imports.NewTable("TRegistry_WriteExpandStringWithUnicodeStringX2", 0), // procedure WriteExpandStringWithUnicodeStringX2
			/* 86 */ imports.NewTable("TRegistry_WriteExpandStringWithStringX2", 0), // procedure WriteExpandStringWithStringX2
			/* 87 */ imports.NewTable("TRegistry_WriteStringList", 0), // procedure WriteStringList
			/* 88 */ imports.NewTable("TRegistry_WriteStringArray", 0), // procedure WriteStringArray
			/* 89 */ imports.NewTable("TRegistry_WriteTimeWithUnicodeStringDateTime", 0), // procedure WriteTimeWithUnicodeStringDateTime
			/* 90 */ imports.NewTable("TRegistry_WriteTimeWithStringDateTime", 0), // procedure WriteTimeWithStringDateTime
			/* 91 */ imports.NewTable("TRegistry_Access", 0), // property Access
			/* 92 */ imports.NewTable("TRegistry_CurrentKey", 0), // property CurrentKey
			/* 93 */ imports.NewTable("TRegistry_CurrentPath", 0), // property CurrentPath
			/* 94 */ imports.NewTable("TRegistry_LazyWrite", 0), // property LazyWrite
			/* 95 */ imports.NewTable("TRegistry_RootKey", 0), // property RootKey
			/* 96 */ imports.NewTable("TRegistry_StringSizeIncludesNull", 0), // property StringSizeIncludesNull
			/* 97 */ imports.NewTable("TRegistry_LastError", 0), // property LastError
			/* 98 */ imports.NewTable("TRegistry_LastErrorMsg", 0), // property LastErrorMsg
		}
	})
	return registryImport
}
