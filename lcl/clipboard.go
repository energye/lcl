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

// IClipboard Parent: IPersistent
type IClipboard interface {
	IPersistent
	AsText() string                                           // property
	SetAsText(AValue string)                                  // property
	ClipboardType() TClipboardType                            // property
	FormatCount() int32                                       // property
	Formats(Index int32) TClipboardFormat                     // property
	AddFormat(FormatID TClipboardFormat, Stream IStream) bool // function
	AddFormat1(FormatID TClipboardFormat, Buffer []byte) bool // function
	FindPictureFormatID() TClipboardFormat                    // function
	FindFormatID(FormatName string) TClipboardFormat          // function
	GetAsHtml(ExtractFragmentOnly bool) string                // function
	GetComponent(Owner, Parent IComponent) IComponent         // function
	GetFormat(FormatID TClipboardFormat, Stream IStream) bool // function
	GetTextBuf(Buffer *string, BufSize int32) int32           // function
	HasFormat(FormatID TClipboardFormat) bool                 // function
	HasFormatName(FormatName string) bool                     // function
	HasPictureFormat() bool                                   // function
	SetComponent(Component IComponent) bool                   // function
	SetComponentAsText(Component IComponent) bool             // function
	SetFormat(FormatID TClipboardFormat, Stream IStream) bool // function
	AssignTo(Dest IPersistent)                                // procedure
	Clear()                                                   // procedure
	Close()                                                   // procedure
	SupportedFormats(List IStrings)                           // procedure
	Open()                                                    // procedure
	SetAsHtml(Html string)                                    // procedure
	SetAsHtml1(Html string, PlainText string)                 // procedure
	SetTextBuf(Buffer string)                                 // procedure
	SetOnRequest(fn TClipboardRequestEvent)                   // property event
}

// TClipboard Parent: TPersistent
type TClipboard struct {
	TPersistent
	requestPtr uintptr
}

func NewClipboard() IClipboard {
	r1 := clipboardImportAPI().SysCallN(8)
	return AsClipboard(r1)
}

func NewClipboard1(AClipboardType TClipboardType) IClipboard {
	r1 := clipboardImportAPI().SysCallN(9, uintptr(AClipboardType))
	return AsClipboard(r1)
}

func (m *TClipboard) AsText() string {
	r1 := clipboardImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TClipboard) SetAsText(AValue string) {
	clipboardImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TClipboard) ClipboardType() TClipboardType {
	r1 := clipboardImportAPI().SysCallN(6, m.Instance())
	return TClipboardType(r1)
}

func (m *TClipboard) FormatCount() int32 {
	r1 := clipboardImportAPI().SysCallN(12, m.Instance())
	return int32(r1)
}

func (m *TClipboard) Formats(Index int32) TClipboardFormat {
	r1 := clipboardImportAPI().SysCallN(13, m.Instance(), uintptr(Index))
	return TClipboardFormat(r1)
}

func (m *TClipboard) AddFormat(FormatID TClipboardFormat, Stream IStream) bool {
	r1 := clipboardImportAPI().SysCallN(0, m.Instance(), uintptr(FormatID), GetObjectUintptr(Stream))
	return GoBool(r1)
}

func (m *TClipboard) AddFormat1(FormatID TClipboardFormat, Buffer []byte) bool {
	r1 := clipboardImportAPI().SysCallN(1, m.Instance(), uintptr(FormatID), uintptr(unsafePointer(&Buffer[0])), uintptr(len(Buffer)))
	return GoBool(r1)
}

func (m *TClipboard) FindPictureFormatID() TClipboardFormat {
	r1 := clipboardImportAPI().SysCallN(11, m.Instance())
	return TClipboardFormat(r1)
}

func (m *TClipboard) FindFormatID(FormatName string) TClipboardFormat {
	r1 := clipboardImportAPI().SysCallN(10, m.Instance(), PascalStr(FormatName))
	return TClipboardFormat(r1)
}

func (m *TClipboard) GetAsHtml(ExtractFragmentOnly bool) string {
	r1 := clipboardImportAPI().SysCallN(14, m.Instance(), PascalBool(ExtractFragmentOnly))
	return GoStr(r1)
}

func (m *TClipboard) GetComponent(Owner, Parent IComponent) IComponent {
	r1 := clipboardImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(Owner), GetObjectUintptr(Parent))
	return AsComponent(r1)
}

func (m *TClipboard) GetFormat(FormatID TClipboardFormat, Stream IStream) bool {
	r1 := clipboardImportAPI().SysCallN(16, m.Instance(), uintptr(FormatID), GetObjectUintptr(Stream))
	return GoBool(r1)
}

func (m *TClipboard) GetTextBuf(Buffer *string, BufSize int32) int32 {
	r1 := sysCallGetTextBuf(clipboardImportAPI(), 17, m.Instance(), Buffer, BufSize)
	return int32(r1)
}

func (m *TClipboard) HasFormat(FormatID TClipboardFormat) bool {
	r1 := clipboardImportAPI().SysCallN(18, m.Instance(), uintptr(FormatID))
	return GoBool(r1)
}

func (m *TClipboard) HasFormatName(FormatName string) bool {
	r1 := clipboardImportAPI().SysCallN(19, m.Instance(), PascalStr(FormatName))
	return GoBool(r1)
}

func (m *TClipboard) HasPictureFormat() bool {
	r1 := clipboardImportAPI().SysCallN(20, m.Instance())
	return GoBool(r1)
}

func (m *TClipboard) SetComponent(Component IComponent) bool {
	r1 := clipboardImportAPI().SysCallN(24, m.Instance(), GetObjectUintptr(Component))
	return GoBool(r1)
}

func (m *TClipboard) SetComponentAsText(Component IComponent) bool {
	r1 := clipboardImportAPI().SysCallN(25, m.Instance(), GetObjectUintptr(Component))
	return GoBool(r1)
}

func (m *TClipboard) SetFormat(FormatID TClipboardFormat, Stream IStream) bool {
	r1 := clipboardImportAPI().SysCallN(26, m.Instance(), uintptr(FormatID), GetObjectUintptr(Stream))
	return GoBool(r1)
}

func ClipboardClass() TClass {
	ret := clipboardImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TClipboard) AssignTo(Dest IPersistent) {
	clipboardImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TClipboard) Clear() {
	clipboardImportAPI().SysCallN(5, m.Instance())
}

func (m *TClipboard) Close() {
	clipboardImportAPI().SysCallN(7, m.Instance())
}

func (m *TClipboard) SupportedFormats(List IStrings) {
	clipboardImportAPI().SysCallN(29, m.Instance(), GetObjectUintptr(List))
}

func (m *TClipboard) Open() {
	clipboardImportAPI().SysCallN(21, m.Instance())
}

func (m *TClipboard) SetAsHtml(Html string) {
	clipboardImportAPI().SysCallN(22, m.Instance(), PascalStr(Html))
}

func (m *TClipboard) SetAsHtml1(Html string, PlainText string) {
	clipboardImportAPI().SysCallN(23, m.Instance(), PascalStr(Html), PascalStr(PlainText))
}

func (m *TClipboard) SetTextBuf(Buffer string) {
	clipboardImportAPI().SysCallN(28, m.Instance(), PascalStr(Buffer))
}

func (m *TClipboard) SetOnRequest(fn TClipboardRequestEvent) {
	if m.requestPtr != 0 {
		RemoveEventElement(m.requestPtr)
	}
	m.requestPtr = MakeEventDataPtr(fn)
	clipboardImportAPI().SysCallN(27, m.Instance(), m.requestPtr)
}

var (
	clipboardImport       *imports.Imports = nil
	clipboardImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Clipboard_AddFormat", 0),
		/*1*/ imports.NewTable("Clipboard_AddFormat1", 0),
		/*2*/ imports.NewTable("Clipboard_AsText", 0),
		/*3*/ imports.NewTable("Clipboard_AssignTo", 0),
		/*4*/ imports.NewTable("Clipboard_Class", 0),
		/*5*/ imports.NewTable("Clipboard_Clear", 0),
		/*6*/ imports.NewTable("Clipboard_ClipboardType", 0),
		/*7*/ imports.NewTable("Clipboard_Close", 0),
		/*8*/ imports.NewTable("Clipboard_Create", 0),
		/*9*/ imports.NewTable("Clipboard_Create1", 0),
		/*10*/ imports.NewTable("Clipboard_FindFormatID", 0),
		/*11*/ imports.NewTable("Clipboard_FindPictureFormatID", 0),
		/*12*/ imports.NewTable("Clipboard_FormatCount", 0),
		/*13*/ imports.NewTable("Clipboard_Formats", 0),
		/*14*/ imports.NewTable("Clipboard_GetAsHtml", 0),
		/*15*/ imports.NewTable("Clipboard_GetComponent", 0),
		/*16*/ imports.NewTable("Clipboard_GetFormat", 0),
		/*17*/ imports.NewTable("Clipboard_GetTextBuf", 0),
		/*18*/ imports.NewTable("Clipboard_HasFormat", 0),
		/*19*/ imports.NewTable("Clipboard_HasFormatName", 0),
		/*20*/ imports.NewTable("Clipboard_HasPictureFormat", 0),
		/*21*/ imports.NewTable("Clipboard_Open", 0),
		/*22*/ imports.NewTable("Clipboard_SetAsHtml", 0),
		/*23*/ imports.NewTable("Clipboard_SetAsHtml1", 0),
		/*24*/ imports.NewTable("Clipboard_SetComponent", 0),
		/*25*/ imports.NewTable("Clipboard_SetComponentAsText", 0),
		/*26*/ imports.NewTable("Clipboard_SetFormat", 0),
		/*27*/ imports.NewTable("Clipboard_SetOnRequest", 0),
		/*28*/ imports.NewTable("Clipboard_SetTextBuf", 0),
		/*29*/ imports.NewTable("Clipboard_SupportedFormats", 0),
	}
)

func clipboardImportAPI() *imports.Imports {
	if clipboardImport == nil {
		clipboardImport = NewDefaultImports()
		clipboardImport.SetImportTable(clipboardImportTables)
		clipboardImportTables = nil
	}
	return clipboardImport
}
