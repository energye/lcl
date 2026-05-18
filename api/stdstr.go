//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"
)

// StringToUTF8Ptr
//
// 字符串到UTF8指针
func StringToUTF8Ptr(s string) *uint8 {
	utf8StrArr := make([]uint8, len(s)+1)
	copy(utf8StrArr, s)
	return &utf8StrArr[0]
}

// PasStr string 转 底层字符串指针
func PasStr(str string) uintptr {
	if str == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(StringToUTF8Ptr(str)))
}

// GoStr 将Pascal字符串转换为Go字符串
func GoStr(str uintptr) string {
	l := StrLen(str)
	if l == 0 {
		return ""
	}
	return string(unsafe.Slice((*byte)(unsafe.Pointer(str)), int(l)))
}

type StringBuffer struct {
	Data uintptr
	Size uintptr
}

func NewStringBuffer(data, size uintptr) *StringBuffer {
	return &StringBuffer{Data: data, Size: size}
}

func (m *StringBuffer) Bytes() []byte {
	if m.Data == 0 || m.Size == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(m.Data)), m.Size)
}

func (m *StringBuffer) String() string {
	return string(m.Bytes())
}

func (m *StringBuffer) Release() {
	if m.Data == 0 {
		return
	}
	FreePointer(m.Data)
	m.Data = 0
	m.Size = 0
}
