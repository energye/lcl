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

// IStream Parent: IObject
type IStream interface {
	IObject
	Read(buffer *uintptr, count int32) int32                                     // function
	Write(buffer uintptr, count int32) int32                                     // function
	SeekWithIntWord(offset int32, origin uint16) int32                           // function
	SeekWithInt64SeekOrigin(offset int64, origin types.TSeekOrigin) int64        // function
	CopyFrom(source IStream, count int64) int64                                  // function
	ReadComponent(instance IComponent) IComponent                                // function
	ReadComponentRes(instance IComponent) IComponent                             // function
	ReadByte() byte                                                              // function
	ReadWord() uint16                                                            // function
	ReadDWord() uint32                                                           // function
	ReadQWord() uintptr                                                          // function
	ReadAnsiString() string                                                      // function
	ReadBuffer(buffer *uintptr, count int32)                                     // procedure
	WriteBuffer(buffer uintptr, count int32)                                     // procedure
	WriteComponent(instance IComponent)                                          // procedure
	WriteComponentRes(resName string, instance IComponent)                       // procedure
	WriteDescendent(instance IComponent, ancestor IComponent)                    // procedure
	WriteDescendentRes(resName string, instance IComponent, ancestor IComponent) // procedure
	WriteResourceHeader(resName string, fixupInfo *int32)                        // procedure
	FixupResourceHeader(fixupInfo int32)                                         // procedure
	ReadResHeader()                                                              // procedure
	WriteByte(B byte)                                                            // procedure
	WriteWord(W uint16)                                                          // procedure
	WriteDWord(D uint32)                                                         // procedure
	WriteQWord(Q uintptr)                                                        // procedure
	WriteAnsiString(S string)                                                    // procedure
	Position() int64                                                             // property Position Getter
	SetPosition(value int64)                                                     // property Position Setter
	Size() int64                                                                 // property Size Getter
	SetSizeToInt64(value int64)                                                  // property Size Setter
}

type TStream struct {
	TObject
}

func (m *TStream) Read(buffer *uintptr, count int32) int32 {
	if !m.IsValid() {
		return 0
	}
	bufferPtr := uintptr(*buffer)
	r := streamAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&bufferPtr)), uintptr(count))
	*buffer = uintptr(bufferPtr)
	return int32(r)
}

func (m *TStream) Write(buffer uintptr, count int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := streamAPI().SysCallN(2, m.Instance(), uintptr(buffer), uintptr(count))
	return int32(r)
}

func (m *TStream) SeekWithIntWord(offset int32, origin uint16) int32 {
	if !m.IsValid() {
		return 0
	}
	r := streamAPI().SysCallN(3, m.Instance(), uintptr(offset), uintptr(origin))
	return int32(r)
}

func (m *TStream) SeekWithInt64SeekOrigin(offset int64, origin types.TSeekOrigin) (result int64) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&offset)), uintptr(origin), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TStream) CopyFrom(source IStream, count int64) (result int64) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(source), uintptr(base.UnsafePointer(&count)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TStream) ReadComponent(instance IComponent) IComponent {
	if !m.IsValid() {
		return nil
	}
	r := streamAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(instance))
	return AsComponent(r)
}

func (m *TStream) ReadComponentRes(instance IComponent) IComponent {
	if !m.IsValid() {
		return nil
	}
	r := streamAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(instance))
	return AsComponent(r)
}

func (m *TStream) ReadByte() byte {
	if !m.IsValid() {
		return 0
	}
	r := streamAPI().SysCallN(8, m.Instance())
	return byte(r)
}

func (m *TStream) ReadWord() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := streamAPI().SysCallN(9, m.Instance())
	return uint16(r)
}

func (m *TStream) ReadDWord() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := streamAPI().SysCallN(10, m.Instance())
	return uint32(r)
}

func (m *TStream) ReadQWord() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := streamAPI().SysCallN(11, m.Instance())
	return uintptr(r)
}

func (m *TStream) ReadAnsiString() string {
	if !m.IsValid() {
		return ""
	}
	r := streamAPI().SysCallN(12, m.Instance())
	return api.GoStr(r)
}

func (m *TStream) ReadBuffer(buffer *uintptr, count int32) {
	if !m.IsValid() {
		return
	}
	bufferPtr := uintptr(*buffer)
	streamAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&bufferPtr)), uintptr(count))
	*buffer = uintptr(bufferPtr)
}

func (m *TStream) WriteBuffer(buffer uintptr, count int32) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(14, m.Instance(), uintptr(buffer), uintptr(count))
}

func (m *TStream) WriteComponent(instance IComponent) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(instance))
}

func (m *TStream) WriteComponentRes(resName string, instance IComponent) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(16, m.Instance(), api.PasStr(resName), base.GetObjectUintptr(instance))
}

func (m *TStream) WriteDescendent(instance IComponent, ancestor IComponent) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(instance), base.GetObjectUintptr(ancestor))
}

func (m *TStream) WriteDescendentRes(resName string, instance IComponent, ancestor IComponent) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(18, m.Instance(), api.PasStr(resName), base.GetObjectUintptr(instance), base.GetObjectUintptr(ancestor))
}

func (m *TStream) WriteResourceHeader(resName string, fixupInfo *int32) {
	if !m.IsValid() {
		return
	}
	fixupInfoPtr := uintptr(*fixupInfo)
	streamAPI().SysCallN(19, m.Instance(), api.PasStr(resName), uintptr(base.UnsafePointer(&fixupInfoPtr)))
	*fixupInfo = int32(fixupInfoPtr)
}

func (m *TStream) FixupResourceHeader(fixupInfo int32) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(20, m.Instance(), uintptr(fixupInfo))
}

func (m *TStream) ReadResHeader() {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(21, m.Instance())
}

func (m *TStream) WriteByte(B byte) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(22, m.Instance(), uintptr(B))
}

func (m *TStream) WriteWord(W uint16) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(23, m.Instance(), uintptr(W))
}

func (m *TStream) WriteDWord(D uint32) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(24, m.Instance(), uintptr(D))
}

func (m *TStream) WriteQWord(Q uintptr) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(25, m.Instance(), uintptr(Q))
}

func (m *TStream) WriteAnsiString(S string) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(26, m.Instance(), api.PasStr(S))
}

func (m *TStream) Position() (result int64) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(27, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TStream) SetPosition(value int64) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(27, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TStream) Size() (result int64) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(28, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TStream) SetSizeToInt64(value int64) {
	if !m.IsValid() {
		return
	}
	streamAPI().SysCallN(28, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

// NewStream class constructor
func NewStream() IStream {
	r := streamAPI().SysCallN(0)
	return AsStream(r)
}

var (
	streamOnce   base.Once
	streamImport *imports.Imports = nil
)

func streamAPI() *imports.Imports {
	streamOnce.Do(func() {
		streamImport = api.NewDefaultImports()
		streamImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStream_Create", 0), // constructor NewStream
			/* 1 */ imports.NewTable("TStream_Read", 0), // function Read
			/* 2 */ imports.NewTable("TStream_Write", 0), // function Write
			/* 3 */ imports.NewTable("TStream_SeekWithIntWord", 0), // function SeekWithIntWord
			/* 4 */ imports.NewTable("TStream_SeekWithInt64SeekOrigin", 0), // function SeekWithInt64SeekOrigin
			/* 5 */ imports.NewTable("TStream_CopyFrom", 0), // function CopyFrom
			/* 6 */ imports.NewTable("TStream_ReadComponent", 0), // function ReadComponent
			/* 7 */ imports.NewTable("TStream_ReadComponentRes", 0), // function ReadComponentRes
			/* 8 */ imports.NewTable("TStream_ReadByte", 0), // function ReadByte
			/* 9 */ imports.NewTable("TStream_ReadWord", 0), // function ReadWord
			/* 10 */ imports.NewTable("TStream_ReadDWord", 0), // function ReadDWord
			/* 11 */ imports.NewTable("TStream_ReadQWord", 0), // function ReadQWord
			/* 12 */ imports.NewTable("TStream_ReadAnsiString", 0), // function ReadAnsiString
			/* 13 */ imports.NewTable("TStream_ReadBuffer", 0), // procedure ReadBuffer
			/* 14 */ imports.NewTable("TStream_WriteBuffer", 0), // procedure WriteBuffer
			/* 15 */ imports.NewTable("TStream_WriteComponent", 0), // procedure WriteComponent
			/* 16 */ imports.NewTable("TStream_WriteComponentRes", 0), // procedure WriteComponentRes
			/* 17 */ imports.NewTable("TStream_WriteDescendent", 0), // procedure WriteDescendent
			/* 18 */ imports.NewTable("TStream_WriteDescendentRes", 0), // procedure WriteDescendentRes
			/* 19 */ imports.NewTable("TStream_WriteResourceHeader", 0), // procedure WriteResourceHeader
			/* 20 */ imports.NewTable("TStream_FixupResourceHeader", 0), // procedure FixupResourceHeader
			/* 21 */ imports.NewTable("TStream_ReadResHeader", 0), // procedure ReadResHeader
			/* 22 */ imports.NewTable("TStream_WriteByte", 0), // procedure WriteByte
			/* 23 */ imports.NewTable("TStream_WriteWord", 0), // procedure WriteWord
			/* 24 */ imports.NewTable("TStream_WriteDWord", 0), // procedure WriteDWord
			/* 25 */ imports.NewTable("TStream_WriteQWord", 0), // procedure WriteQWord
			/* 26 */ imports.NewTable("TStream_WriteAnsiString", 0), // procedure WriteAnsiString
			/* 27 */ imports.NewTable("TStream_Position", 0), // property Position
			/* 28 */ imports.NewTable("TStream_Size", 0), // property Size
		}
	})
	return streamImport
}
