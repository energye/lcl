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

// IMemoryStream Parent: ICustomMemoryStream
type IMemoryStream interface {
	ICustomMemoryStream
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Clear()                        // procedure
	LoadFromStream(Stream IStream) // procedure
	LoadFromFile(FileName string)  // procedure
}

// TMemoryStream Parent: TCustomMemoryStream
type TMemoryStream struct {
	TCustomMemoryStream
}

func NewMemoryStream() IMemoryStream {
	r1 := LCL().SysCallN(4262)
	return AsMemoryStream(r1)
}

func MemoryStreamClass() TClass {
	ret := LCL().SysCallN(4260)
	return TClass(ret)
}

func (m *TMemoryStream) Clear() {
	LCL().SysCallN(4261, m.Instance())
}

func (m *TMemoryStream) LoadFromStream(Stream IStream) {
	LCL().SysCallN(4264, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TMemoryStream) LoadFromFile(FileName string) {
	LCL().SysCallN(4263, m.Instance(), PascalStr(FileName))
}
