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

// IStream Parent: IObject
type IStream interface {
	IObject
	Position() (resultInt64 int64)                                    // property
	SetPosition(AValue int64)                                         // property
	Size() (resultInt64 int64)                                        // property
	SetSize(AValue int64)                                             // property
	Read(count int32) []byte                                          // function
	Write(Buffer []byte) int32                                        // function
	Seek(Offset int32, Origin Word) int32                             // function
	Seek1(Offset int64, Origin TSeekOrigin) (resultInt64 int64)       // function
	CopyFrom(Source IStream, Count int64) (resultInt64 int64)         // function
	ReadComponent(Instance IComponent) IComponent                     // function
	ReadComponentRes(Instance IComponent) IComponent                  // function
	ReadByte() Byte                                                   // function
	ReadWord() Word                                                   // function
	ReadDWord() uint32                                                // function
	ReadQWord() QWord                                                 // function
	ReadAnsiString() string                                           // function
	ReadBuffer(count int32) []byte                                    // procedure
	WriteBuffer(Buffer []byte)                                        // procedure
	WriteComponent(Instance IComponent)                               // procedure
	WriteComponentRes(ResName string, Instance IComponent)            // procedure
	WriteDescendent(Instance, Ancestor IComponent)                    // procedure
	WriteDescendentRes(ResName string, Instance, Ancestor IComponent) // procedure
	WriteResourceHeader(ResName string, FixupInfo *int32)             // procedure
	FixupResourceHeader(FixupInfo int32)                              // procedure
	ReadResHeader()                                                   // procedure
	WriteByte(b Byte)                                                 // procedure
	WriteWord(w Word)                                                 // procedure
	WriteDWord(d uint32)                                              // procedure
	WriteQWord(q QWord)                                               // procedure
	WriteAnsiString(S string)                                         // procedure
}

// TStream Parent: TObject
type TStream struct {
	TObject
}

func NewStream() IStream {
	r1 := LCL().SysCallN(5202)
	return AsStream(r1)
}

func (m *TStream) Position() (resultInt64 int64) {
	LCL().SysCallN(5204, 0, m.Instance(), uintptr(unsafePointer(&resultInt64)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) SetPosition(AValue int64) {
	LCL().SysCallN(5204, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TStream) Size() (resultInt64 int64) {
	LCL().SysCallN(5217, 0, m.Instance(), uintptr(unsafePointer(&resultInt64)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) SetSize(AValue int64) {
	LCL().SysCallN(5217, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TStream) Read(count int32) []byte {
	_, d := sysCallBufferRead(5205, m.Instance(), count)
	return d
}

func (m *TStream) Write(Buffer []byte) int32 {
	r1 := sysCallBufferWrite(5218, m.Instance(), Buffer)
	return int32(r1)
}

func (m *TStream) Seek(Offset int32, Origin Word) int32 {
	r1 := LCL().SysCallN(5215, m.Instance(), uintptr(Offset), uintptr(Origin))
	return int32(r1)
}

func (m *TStream) Seek1(Offset int64, Origin TSeekOrigin) (resultInt64 int64) {
	LCL().SysCallN(5216, m.Instance(), uintptr(unsafePointer(&Offset)), uintptr(Origin), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) CopyFrom(Source IStream, Count int64) (resultInt64 int64) {
	LCL().SysCallN(5201, m.Instance(), GetObjectUintptr(Source), uintptr(unsafePointer(&Count)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) ReadComponent(Instance IComponent) IComponent {
	r1 := LCL().SysCallN(5209, m.Instance(), GetObjectUintptr(Instance))
	return AsComponent(r1)
}

func (m *TStream) ReadComponentRes(Instance IComponent) IComponent {
	r1 := LCL().SysCallN(5210, m.Instance(), GetObjectUintptr(Instance))
	return AsComponent(r1)
}

func (m *TStream) ReadByte() Byte {
	r1 := LCL().SysCallN(5208, m.Instance())
	return Byte(r1)
}

func (m *TStream) ReadWord() Word {
	r1 := LCL().SysCallN(5214, m.Instance())
	return Word(r1)
}

func (m *TStream) ReadDWord() uint32 {
	r1 := LCL().SysCallN(5211, m.Instance())
	return uint32(r1)
}

func (m *TStream) ReadQWord() QWord {
	r1 := LCL().SysCallN(5212, m.Instance())
	return QWord(r1)
}

func (m *TStream) ReadAnsiString() string {
	r1 := LCL().SysCallN(5206, m.Instance())
	return GoStr(r1)
}

func StreamClass() TClass {
	ret := LCL().SysCallN(5200)
	return TClass(ret)
}

func (m *TStream) ReadBuffer(count int32) []byte {
	_, d := sysCallBufferRead(5207, m.Instance(), count)
	return d
}

func (m *TStream) WriteBuffer(Buffer []byte) {
	sysCallBufferWrite(5220, m.Instance(), Buffer)
}

func (m *TStream) WriteComponent(Instance IComponent) {
	LCL().SysCallN(5222, m.Instance(), GetObjectUintptr(Instance))
}

func (m *TStream) WriteComponentRes(ResName string, Instance IComponent) {
	LCL().SysCallN(5223, m.Instance(), PascalStr(ResName), GetObjectUintptr(Instance))
}

func (m *TStream) WriteDescendent(Instance, Ancestor IComponent) {
	LCL().SysCallN(5225, m.Instance(), GetObjectUintptr(Instance), GetObjectUintptr(Ancestor))
}

func (m *TStream) WriteDescendentRes(ResName string, Instance, Ancestor IComponent) {
	LCL().SysCallN(5226, m.Instance(), PascalStr(ResName), GetObjectUintptr(Instance), GetObjectUintptr(Ancestor))
}

func (m *TStream) WriteResourceHeader(ResName string, FixupInfo *int32) {
	var result1 uintptr
	LCL().SysCallN(5228, m.Instance(), PascalStr(ResName), uintptr(unsafePointer(&result1)))
	*FixupInfo = int32(result1)
}

func (m *TStream) FixupResourceHeader(FixupInfo int32) {
	LCL().SysCallN(5203, m.Instance(), uintptr(FixupInfo))
}

func (m *TStream) ReadResHeader() {
	LCL().SysCallN(5213, m.Instance())
}

func (m *TStream) WriteByte(b Byte) {
	LCL().SysCallN(5221, m.Instance(), uintptr(b))
}

func (m *TStream) WriteWord(w Word) {
	LCL().SysCallN(5229, m.Instance(), uintptr(w))
}

func (m *TStream) WriteDWord(d uint32) {
	LCL().SysCallN(5224, m.Instance(), uintptr(d))
}

func (m *TStream) WriteQWord(q QWord) {
	LCL().SysCallN(5227, m.Instance(), uintptr(q))
}

func (m *TStream) WriteAnsiString(S string) {
	LCL().SysCallN(5219, m.Instance(), PascalStr(S))
}
