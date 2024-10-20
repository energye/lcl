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
	r1 := streamImportAPI().SysCallN(2)
	return AsStream(r1)
}

func (m *TStream) Position() (resultInt64 int64) {
	streamImportAPI().SysCallN(4, 0, m.Instance(), uintptr(unsafePointer(&resultInt64)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) SetPosition(AValue int64) {
	streamImportAPI().SysCallN(4, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TStream) Size() (resultInt64 int64) {
	streamImportAPI().SysCallN(17, 0, m.Instance(), uintptr(unsafePointer(&resultInt64)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) SetSize(AValue int64) {
	streamImportAPI().SysCallN(17, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TStream) Read(count int32) []byte {
	_, d := sysCallBufferRead(streamImportAPI(), 5, m.Instance(), count)
	return d
}

func (m *TStream) Write(Buffer []byte) int32 {
	r1 := sysCallBufferWrite(streamImportAPI(), 18, m.Instance(), Buffer)
	return int32(r1)
}

func (m *TStream) Seek(Offset int32, Origin Word) int32 {
	r1 := streamImportAPI().SysCallN(15, m.Instance(), uintptr(Offset), uintptr(Origin))
	return int32(r1)
}

func (m *TStream) Seek1(Offset int64, Origin TSeekOrigin) (resultInt64 int64) {
	streamImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(&Offset)), uintptr(Origin), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) CopyFrom(Source IStream, Count int64) (resultInt64 int64) {
	streamImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Source), uintptr(unsafePointer(&Count)), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TStream) ReadComponent(Instance IComponent) IComponent {
	r1 := streamImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(Instance))
	return AsComponent(r1)
}

func (m *TStream) ReadComponentRes(Instance IComponent) IComponent {
	r1 := streamImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(Instance))
	return AsComponent(r1)
}

func (m *TStream) ReadByte() Byte {
	r1 := streamImportAPI().SysCallN(8, m.Instance())
	return Byte(r1)
}

func (m *TStream) ReadWord() Word {
	r1 := streamImportAPI().SysCallN(14, m.Instance())
	return Word(r1)
}

func (m *TStream) ReadDWord() uint32 {
	r1 := streamImportAPI().SysCallN(11, m.Instance())
	return uint32(r1)
}

func (m *TStream) ReadQWord() QWord {
	r1 := streamImportAPI().SysCallN(12, m.Instance())
	return QWord(r1)
}

func (m *TStream) ReadAnsiString() string {
	r1 := streamImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func StreamClass() TClass {
	ret := streamImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TStream) ReadBuffer(count int32) []byte {
	_, d := sysCallBufferRead(streamImportAPI(), 7, m.Instance(), count)
	return d
}

func (m *TStream) WriteBuffer(Buffer []byte) {
	sysCallBufferWrite(streamImportAPI(), 20, m.Instance(), Buffer)
}

func (m *TStream) WriteComponent(Instance IComponent) {
	streamImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(Instance))
}

func (m *TStream) WriteComponentRes(ResName string, Instance IComponent) {
	streamImportAPI().SysCallN(23, m.Instance(), PascalStr(ResName), GetObjectUintptr(Instance))
}

func (m *TStream) WriteDescendent(Instance, Ancestor IComponent) {
	streamImportAPI().SysCallN(25, m.Instance(), GetObjectUintptr(Instance), GetObjectUintptr(Ancestor))
}

func (m *TStream) WriteDescendentRes(ResName string, Instance, Ancestor IComponent) {
	streamImportAPI().SysCallN(26, m.Instance(), PascalStr(ResName), GetObjectUintptr(Instance), GetObjectUintptr(Ancestor))
}

func (m *TStream) WriteResourceHeader(ResName string, FixupInfo *int32) {
	var result1 uintptr
	streamImportAPI().SysCallN(28, m.Instance(), PascalStr(ResName), uintptr(unsafePointer(&result1)))
	*FixupInfo = int32(result1)
}

func (m *TStream) FixupResourceHeader(FixupInfo int32) {
	streamImportAPI().SysCallN(3, m.Instance(), uintptr(FixupInfo))
}

func (m *TStream) ReadResHeader() {
	streamImportAPI().SysCallN(13, m.Instance())
}

func (m *TStream) WriteByte(b Byte) {
	streamImportAPI().SysCallN(21, m.Instance(), uintptr(b))
}

func (m *TStream) WriteWord(w Word) {
	streamImportAPI().SysCallN(29, m.Instance(), uintptr(w))
}

func (m *TStream) WriteDWord(d uint32) {
	streamImportAPI().SysCallN(24, m.Instance(), uintptr(d))
}

func (m *TStream) WriteQWord(q QWord) {
	streamImportAPI().SysCallN(27, m.Instance(), uintptr(q))
}

func (m *TStream) WriteAnsiString(S string) {
	streamImportAPI().SysCallN(19, m.Instance(), PascalStr(S))
}

var (
	streamImport       *imports.Imports = nil
	streamImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Stream_Class", 0),
		/*1*/ imports.NewTable("Stream_CopyFrom", 0),
		/*2*/ imports.NewTable("Stream_Create", 0),
		/*3*/ imports.NewTable("Stream_FixupResourceHeader", 0),
		/*4*/ imports.NewTable("Stream_Position", 0),
		/*5*/ imports.NewTable("Stream_Read", 0),
		/*6*/ imports.NewTable("Stream_ReadAnsiString", 0),
		/*7*/ imports.NewTable("Stream_ReadBuffer", 0),
		/*8*/ imports.NewTable("Stream_ReadByte", 0),
		/*9*/ imports.NewTable("Stream_ReadComponent", 0),
		/*10*/ imports.NewTable("Stream_ReadComponentRes", 0),
		/*11*/ imports.NewTable("Stream_ReadDWord", 0),
		/*12*/ imports.NewTable("Stream_ReadQWord", 0),
		/*13*/ imports.NewTable("Stream_ReadResHeader", 0),
		/*14*/ imports.NewTable("Stream_ReadWord", 0),
		/*15*/ imports.NewTable("Stream_Seek", 0),
		/*16*/ imports.NewTable("Stream_Seek1", 0),
		/*17*/ imports.NewTable("Stream_Size", 0),
		/*18*/ imports.NewTable("Stream_Write", 0),
		/*19*/ imports.NewTable("Stream_WriteAnsiString", 0),
		/*20*/ imports.NewTable("Stream_WriteBuffer", 0),
		/*21*/ imports.NewTable("Stream_WriteByte", 0),
		/*22*/ imports.NewTable("Stream_WriteComponent", 0),
		/*23*/ imports.NewTable("Stream_WriteComponentRes", 0),
		/*24*/ imports.NewTable("Stream_WriteDWord", 0),
		/*25*/ imports.NewTable("Stream_WriteDescendent", 0),
		/*26*/ imports.NewTable("Stream_WriteDescendentRes", 0),
		/*27*/ imports.NewTable("Stream_WriteQWord", 0),
		/*28*/ imports.NewTable("Stream_WriteResourceHeader", 0),
		/*29*/ imports.NewTable("Stream_WriteWord", 0),
	}
)

func streamImportAPI() *imports.Imports {
	if streamImport == nil {
		streamImport = NewDefaultImports()
		streamImport.SetImportTable(streamImportTables)
		streamImportTables = nil
	}
	return streamImport
}
