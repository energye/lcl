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

// IClipboard Parent: IPersistent
type IClipboard interface {
	IPersistent
	AddFormatWithClipboardFormatStream(formatID types.TClipboardFormat, stream IStream) bool                  // function
	AddFormatWithClipboardFormatPointerInt(formatID types.TClipboardFormat, buffer *uintptr, size int32) bool // function
	FindPictureFormatID() types.TClipboardFormat                                                              // function
	FindFormatID(formatName string) types.TClipboardFormat                                                    // function
	GetAsHtml(extractFragmentOnly bool) string                                                                // function
	GetComponentWithComponentX2(owner IComponent, parent IComponent) IComponent                               // function
	GetFormat(formatID types.TClipboardFormat, stream IStream) bool                                           // function
	GetTextBuf(buffer uintptr, bufSize int32) int32                                                           // function
	HasFormat(formatID types.TClipboardFormat) bool                                                           // function
	HasFormatName(formatName string) bool                                                                     // function
	HasPictureFormat() bool                                                                                   // function
	SetComponent(component IComponent) bool                                                                   // function
	SetComponentAsText(component IComponent) bool                                                             // function
	SetFormat(formatID types.TClipboardFormat, stream IStream) bool                                           // function
	SetSupportedFormats(formatCount int32, formatList types.TClipboardFormat) bool                            // function
	AssignTo(dest IPersistent)                                                                                // procedure
	Clear()                                                                                                   // procedure
	Close()                                                                                                   // procedure
	SupportedFormatsWithStrings(list IStrings)                                                                // procedure
	SupportedFormatsWithIntPClipboardFormat(formatCount *int32, formatList *types.TClipboardFormat)           // procedure
	Open()                                                                                                    // procedure
	SetAsHtmlWithString(html string)                                                                          // procedure
	SetAsHtmlWithStringX2(html string, plainText string)                                                      // procedure
	SetTextBuf(buffer uintptr)                                                                                // procedure
	AsText() string                                                                                           // property AsText Getter
	SetAsText(value string)                                                                                   // property AsText Setter
	ClipboardType() types.TClipboardType                                                                      // property ClipboardType Getter
	FormatCount() int32                                                                                       // property FormatCount Getter
	Formats(index int32) types.TClipboardFormat                                                               // property Formats Getter
	SetOnRequest(fn TClipboardRequestEvent)                                                                   // property event
}

type TClipboard struct {
	TPersistent
}

func (m *TClipboard) AddFormatWithClipboardFormatStream(formatID types.TClipboardFormat, stream IStream) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(2, m.Instance(), uintptr(formatID), base.GetObjectUintptr(stream))
	return api.GoBool(r)
}

func (m *TClipboard) AddFormatWithClipboardFormatPointerInt(formatID types.TClipboardFormat, buffer *uintptr, size int32) bool {
	if !m.IsValid() {
		return false
	}
	bufferPtr := uintptr(*buffer)
	r := clipboardAPI().SysCallN(3, m.Instance(), uintptr(formatID), uintptr(base.UnsafePointer(&bufferPtr)), uintptr(size))
	*buffer = uintptr(bufferPtr)
	return api.GoBool(r)
}

func (m *TClipboard) FindPictureFormatID() types.TClipboardFormat {
	if !m.IsValid() {
		return 0
	}
	r := clipboardAPI().SysCallN(4, m.Instance())
	return types.TClipboardFormat(r)
}

func (m *TClipboard) FindFormatID(formatName string) types.TClipboardFormat {
	if !m.IsValid() {
		return 0
	}
	r := clipboardAPI().SysCallN(5, m.Instance(), api.PasStr(formatName))
	return types.TClipboardFormat(r)
}

func (m *TClipboard) GetAsHtml(extractFragmentOnly bool) string {
	if !m.IsValid() {
		return ""
	}
	r := clipboardAPI().SysCallN(6, m.Instance(), api.PasBool(extractFragmentOnly))
	return api.GoStr(r)
}

func (m *TClipboard) GetComponentWithComponentX2(owner IComponent, parent IComponent) IComponent {
	if !m.IsValid() {
		return nil
	}
	r := clipboardAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(owner), base.GetObjectUintptr(parent))
	return AsComponent(r)
}

func (m *TClipboard) GetFormat(formatID types.TClipboardFormat, stream IStream) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(8, m.Instance(), uintptr(formatID), base.GetObjectUintptr(stream))
	return api.GoBool(r)
}

func (m *TClipboard) GetTextBuf(buffer uintptr, bufSize int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := clipboardAPI().SysCallN(9, m.Instance(), uintptr(buffer), uintptr(bufSize))
	return int32(r)
}

func (m *TClipboard) HasFormat(formatID types.TClipboardFormat) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(10, m.Instance(), uintptr(formatID))
	return api.GoBool(r)
}

func (m *TClipboard) HasFormatName(formatName string) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(11, m.Instance(), api.PasStr(formatName))
	return api.GoBool(r)
}

func (m *TClipboard) HasPictureFormat() bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TClipboard) SetComponent(component IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(component))
	return api.GoBool(r)
}

func (m *TClipboard) SetComponentAsText(component IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(component))
	return api.GoBool(r)
}

func (m *TClipboard) SetFormat(formatID types.TClipboardFormat, stream IStream) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(15, m.Instance(), uintptr(formatID), base.GetObjectUintptr(stream))
	return api.GoBool(r)
}

func (m *TClipboard) SetSupportedFormats(formatCount int32, formatList types.TClipboardFormat) bool {
	if !m.IsValid() {
		return false
	}
	r := clipboardAPI().SysCallN(16, m.Instance(), uintptr(formatCount), uintptr(formatList))
	return api.GoBool(r)
}

func (m *TClipboard) AssignTo(dest IPersistent) {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(dest))
}

func (m *TClipboard) Clear() {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(18, m.Instance())
}

func (m *TClipboard) Close() {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(19, m.Instance())
}

func (m *TClipboard) SupportedFormatsWithStrings(list IStrings) {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(list))
}

func (m *TClipboard) SupportedFormatsWithIntPClipboardFormat(formatCount *int32, formatList *types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	formatCountPtr := uintptr(*formatCount)
	formatListPtr := uintptr(*formatList)
	clipboardAPI().SysCallN(21, m.Instance(), uintptr(base.UnsafePointer(&formatCountPtr)), uintptr(base.UnsafePointer(&formatListPtr)))
	*formatCount = int32(formatCountPtr)
	*formatList = types.TClipboardFormat(formatListPtr)
}

func (m *TClipboard) Open() {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(22, m.Instance())
}

func (m *TClipboard) SetAsHtmlWithString(html string) {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(23, m.Instance(), api.PasStr(html))
}

func (m *TClipboard) SetAsHtmlWithStringX2(html string, plainText string) {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(24, m.Instance(), api.PasStr(html), api.PasStr(plainText))
}

func (m *TClipboard) SetTextBuf(buffer uintptr) {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(25, m.Instance(), uintptr(buffer))
}

func (m *TClipboard) AsText() string {
	if !m.IsValid() {
		return ""
	}
	r := clipboardAPI().SysCallN(26, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TClipboard) SetAsText(value string) {
	if !m.IsValid() {
		return
	}
	clipboardAPI().SysCallN(26, 1, m.Instance(), api.PasStr(value))
}

func (m *TClipboard) ClipboardType() types.TClipboardType {
	if !m.IsValid() {
		return 0
	}
	r := clipboardAPI().SysCallN(27, m.Instance())
	return types.TClipboardType(r)
}

func (m *TClipboard) FormatCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := clipboardAPI().SysCallN(28, m.Instance())
	return int32(r)
}

func (m *TClipboard) Formats(index int32) types.TClipboardFormat {
	if !m.IsValid() {
		return 0
	}
	r := clipboardAPI().SysCallN(29, m.Instance(), uintptr(index))
	return types.TClipboardFormat(r)
}

func (m *TClipboard) SetOnRequest(fn TClipboardRequestEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTClipboardRequestEvent(fn)
	base.SetEvent(m, 30, clipboardAPI(), api.MakeEventDataPtr(cb))
}

// NewClipboard class constructor
func NewClipboard() IClipboard {
	r := clipboardAPI().SysCallN(0)
	return AsClipboard(r)
}

// NewClipboardWithClipboardType class constructor
func NewClipboardWithClipboardType(clipboardType types.TClipboardType) IClipboard {
	r := clipboardAPI().SysCallN(1, uintptr(clipboardType))
	return AsClipboard(r)
}

var (
	clipboardOnce   base.Once
	clipboardImport *imports.Imports = nil
)

func clipboardAPI() *imports.Imports {
	clipboardOnce.Do(func() {
		clipboardImport = api.NewDefaultImports()
		clipboardImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TClipboard_Create", 0), // constructor NewClipboard
			/* 1 */ imports.NewTable("TClipboard_CreateWithClipboardType", 0), // constructor NewClipboardWithClipboardType
			/* 2 */ imports.NewTable("TClipboard_AddFormatWithClipboardFormatStream", 0), // function AddFormatWithClipboardFormatStream
			/* 3 */ imports.NewTable("TClipboard_AddFormatWithClipboardFormatPointerInt", 0), // function AddFormatWithClipboardFormatPointerInt
			/* 4 */ imports.NewTable("TClipboard_FindPictureFormatID", 0), // function FindPictureFormatID
			/* 5 */ imports.NewTable("TClipboard_FindFormatID", 0), // function FindFormatID
			/* 6 */ imports.NewTable("TClipboard_GetAsHtml", 0), // function GetAsHtml
			/* 7 */ imports.NewTable("TClipboard_GetComponentWithComponentX2", 0), // function GetComponentWithComponentX2
			/* 8 */ imports.NewTable("TClipboard_GetFormat", 0), // function GetFormat
			/* 9 */ imports.NewTable("TClipboard_GetTextBuf", 0), // function GetTextBuf
			/* 10 */ imports.NewTable("TClipboard_HasFormat", 0), // function HasFormat
			/* 11 */ imports.NewTable("TClipboard_HasFormatName", 0), // function HasFormatName
			/* 12 */ imports.NewTable("TClipboard_HasPictureFormat", 0), // function HasPictureFormat
			/* 13 */ imports.NewTable("TClipboard_SetComponent", 0), // function SetComponent
			/* 14 */ imports.NewTable("TClipboard_SetComponentAsText", 0), // function SetComponentAsText
			/* 15 */ imports.NewTable("TClipboard_SetFormat", 0), // function SetFormat
			/* 16 */ imports.NewTable("TClipboard_SetSupportedFormats", 0), // function SetSupportedFormats
			/* 17 */ imports.NewTable("TClipboard_AssignTo", 0), // procedure AssignTo
			/* 18 */ imports.NewTable("TClipboard_Clear", 0), // procedure Clear
			/* 19 */ imports.NewTable("TClipboard_Close", 0), // procedure Close
			/* 20 */ imports.NewTable("TClipboard_SupportedFormatsWithStrings", 0), // procedure SupportedFormatsWithStrings
			/* 21 */ imports.NewTable("TClipboard_SupportedFormatsWithIntPClipboardFormat", 0), // procedure SupportedFormatsWithIntPClipboardFormat
			/* 22 */ imports.NewTable("TClipboard_Open", 0), // procedure Open
			/* 23 */ imports.NewTable("TClipboard_SetAsHtmlWithString", 0), // procedure SetAsHtmlWithString
			/* 24 */ imports.NewTable("TClipboard_SetAsHtmlWithStringX2", 0), // procedure SetAsHtmlWithStringX2
			/* 25 */ imports.NewTable("TClipboard_SetTextBuf", 0), // procedure SetTextBuf
			/* 26 */ imports.NewTable("TClipboard_AsText", 0), // property AsText
			/* 27 */ imports.NewTable("TClipboard_ClipboardType", 0), // property ClipboardType
			/* 28 */ imports.NewTable("TClipboard_FormatCount", 0), // property FormatCount
			/* 29 */ imports.NewTable("TClipboard_Formats", 0), // property Formats
			/* 30 */ imports.NewTable("TClipboard_OnRequest", 0), // event OnRequest
		}
	})
	return clipboardImport
}
