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

// ISynCustomHighlighter Parent: IComponent
type ISynCustomHighlighter interface {
	IComponent
	AddSpecialAttribute(caption string, storedName string) ISynHighlighterAttributes // function
	GetEol() bool                                                                    // function
	GetRange() uintptr                                                               // function
	GetToken() string                                                                // function
	GetEndOfLineAttribute() ISynHighlighterAttributes                                // function
	GetTokenAttribute() ISynHighlighterAttributes                                    // function
	GetTokenKind() int32                                                             // function
	GetTokenPos() int32                                                              // function
	GetTokenLen() int32                                                              // function
	IsKeyword(keyword string) bool                                                   // function
	// IdleScanRanges
	//  * IdleScanRanges
	//  Scan in small chunks during OnIdle; Return True, if more work avail
	//  This method is still under development. It may be changed, removed, un-virtualized, or anything.
	//  In future SynEdit & HL may have other IDLE tasks, and if and when that happens, there will be new ways to control this
	IdleScanRanges() bool                                     // function
	NeedScan() bool                                           // function
	UseUserSettings(settingIndex int32) bool                  // function
	LoadFromRegistry(rootKey types.HKEY, key string) bool     // function
	SaveToRegistry(rootKey types.HKEY, key string) bool       // function
	LoadFromFile(fileName string) bool                        // function
	SaveToFile(fileName string) bool                          // function
	DefHighlightChange(sender IObject)                        // procedure
	BeginUpdate()                                             // procedure
	EndUpdate()                                               // procedure
	AttachToLines(lines ISynEditStringsBase)                  // procedure
	DetachFromLines(lines ISynEditStringsBase)                // procedure
	GetTokenEx(outTokenStart *uintptr, outTokenLength *int32) // procedure
	Next()                                                    // procedure
	NextToEol()                                               // procedure
	StartAtLineIndex(lineNumber int32)                        // procedure
	ContinueNextLine()                                        // procedure
	ScanRanges()                                              // procedure
	ScanAllRanges()                                           // procedure
	SetRange(value uintptr)                                   // procedure
	ResetRange()                                              // procedure
	SetLine(newValue string, lineNumber int32)                // procedure
	EnumUserSettings(settings IStrings)                       // procedure
	AttributeChangeNeedScan() bool                            // property AttributeChangeNeedScan Getter
	DrawDivider(index int32) TSynDividerDrawConfigSetting     // property DrawDivider Getter
	CurrentLines() ISynEditStringsBase                        // property CurrentLines Getter
	SetCurrentLines(value ISynEditStringsBase)                // property CurrentLines Setter
	LineIndex() int32                                         // property LineIndex Getter
	IdentChars() types.TSynIdentChars                         // property IdentChars Getter
	WordBreakChars() types.TSynIdentChars                     // property WordBreakChars Getter
	SetWordBreakChars(value types.TSynIdentChars)             // property WordBreakChars Setter
	LanguageName() string                                     // property LanguageName Getter
	AttrCount() int32                                         // property AttrCount Getter
	Attribute(idx int32) ISynHighlighterAttributes            // property Attribute Getter
	Capabilities() types.TSynHighlighterCapabilities          // property Capabilities Getter
	SampleSource() string                                     // property SampleSource Getter
	SetSampleSource(value string)                             // property SampleSource Setter
	DividerDrawConfig(index int32) ISynDividerDrawConfig      // property DividerDrawConfig Getter
	DividerDrawConfigCount() int32                            // property DividerDrawConfigCount Getter
	DefaultFilter() string                                    // property DefaultFilter Getter
	SetDefaultFilter(value string)                            // property DefaultFilter Setter
	Enabled() bool                                            // property Enabled Getter
	SetEnabled(value bool)                                    // property Enabled Setter
}

type TSynCustomHighlighter struct {
	TComponent
}

func (m *TSynCustomHighlighter) AddSpecialAttribute(caption string, storedName string) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCustomHighlighterAPI().SysCallN(2, m.Instance(), api.PasStr(caption), api.PasStr(storedName))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCustomHighlighter) GetEol() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) GetRange() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(4, m.Instance())
	return uintptr(r)
}

func (m *TSynCustomHighlighter) GetToken() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	synCustomHighlighterAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynCustomHighlighter) GetEndOfLineAttribute() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCustomHighlighterAPI().SysCallN(6, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCustomHighlighter) GetTokenAttribute() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCustomHighlighterAPI().SysCallN(7, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCustomHighlighter) GetTokenKind() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TSynCustomHighlighter) GetTokenPos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TSynCustomHighlighter) GetTokenLen() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TSynCustomHighlighter) IsKeyword(keyword string) bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(11, m.Instance(), api.PasStr(keyword))
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) IdleScanRanges() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) NeedScan() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) UseUserSettings(settingIndex int32) bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(14, m.Instance(), uintptr(settingIndex))
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) LoadFromRegistry(rootKey types.HKEY, key string) bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(15, m.Instance(), uintptr(rootKey), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) SaveToRegistry(rootKey types.HKEY, key string) bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(16, m.Instance(), uintptr(rootKey), api.PasStr(key))
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) LoadFromFile(fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(17, m.Instance(), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) SaveToFile(fileName string) bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(18, m.Instance(), api.PasStr(fileName))
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) DefHighlightChange(sender IObject) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TSynCustomHighlighter) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(20, m.Instance())
}

func (m *TSynCustomHighlighter) EndUpdate() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(21, m.Instance())
}

func (m *TSynCustomHighlighter) AttachToLines(lines ISynEditStringsBase) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(lines))
}

func (m *TSynCustomHighlighter) DetachFromLines(lines ISynEditStringsBase) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(23, m.Instance(), base.GetObjectUintptr(lines))
}

func (m *TSynCustomHighlighter) GetTokenEx(outTokenStart *uintptr, outTokenLength *int32) {
	if !m.IsValid() {
		return
	}
	var tokenStartPtr uintptr
	var tokenLengthPtr uintptr
	synCustomHighlighterAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&tokenStartPtr)), uintptr(base.UnsafePointer(&tokenLengthPtr)))
	*outTokenStart = uintptr(tokenStartPtr)
	*outTokenLength = int32(tokenLengthPtr)
}

func (m *TSynCustomHighlighter) Next() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(25, m.Instance())
}

func (m *TSynCustomHighlighter) NextToEol() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(26, m.Instance())
}

func (m *TSynCustomHighlighter) StartAtLineIndex(lineNumber int32) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(27, m.Instance(), uintptr(lineNumber))
}

func (m *TSynCustomHighlighter) ContinueNextLine() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(28, m.Instance())
}

func (m *TSynCustomHighlighter) ScanRanges() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(29, m.Instance())
}

func (m *TSynCustomHighlighter) ScanAllRanges() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(30, m.Instance())
}

func (m *TSynCustomHighlighter) SetRange(value uintptr) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(31, m.Instance(), uintptr(value))
}

func (m *TSynCustomHighlighter) ResetRange() {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(32, m.Instance())
}

func (m *TSynCustomHighlighter) SetLine(newValue string, lineNumber int32) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(33, m.Instance(), api.PasStr(newValue), uintptr(lineNumber))
}

func (m *TSynCustomHighlighter) EnumUserSettings(settings IStrings) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(34, m.Instance(), base.GetObjectUintptr(settings))
}

func (m *TSynCustomHighlighter) AttributeChangeNeedScan() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(35, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) DrawDivider(index int32) (result TSynDividerDrawConfigSetting) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(36, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynCustomHighlighter) CurrentLines() ISynEditStringsBase {
	if !m.IsValid() {
		return nil
	}
	r := synCustomHighlighterAPI().SysCallN(37, 0, m.Instance())
	return AsSynEditStringsBase(r)
}

func (m *TSynCustomHighlighter) SetCurrentLines(value ISynEditStringsBase) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(37, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCustomHighlighter) LineIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(38, m.Instance())
	return int32(r)
}

func (m *TSynCustomHighlighter) IdentChars() types.TSynIdentChars {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(39, m.Instance())
	return types.TSynIdentChars(r)
}

func (m *TSynCustomHighlighter) WordBreakChars() types.TSynIdentChars {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(40, 0, m.Instance())
	return types.TSynIdentChars(r)
}

func (m *TSynCustomHighlighter) SetWordBreakChars(value types.TSynIdentChars) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TSynCustomHighlighter) LanguageName() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synCustomHighlighterAPI().SysCallN(41, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynCustomHighlighter) AttrCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(42, m.Instance())
	return int32(r)
}

func (m *TSynCustomHighlighter) Attribute(idx int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCustomHighlighterAPI().SysCallN(43, m.Instance(), uintptr(idx))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCustomHighlighter) Capabilities() types.TSynHighlighterCapabilities {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(44, m.Instance())
	return types.TSynHighlighterCapabilities(r)
}

func (m *TSynCustomHighlighter) SampleSource() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synCustomHighlighterAPI().SysCallN(45, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynCustomHighlighter) SetSampleSource(value string) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(45, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynCustomHighlighter) DividerDrawConfig(index int32) ISynDividerDrawConfig {
	if !m.IsValid() {
		return nil
	}
	r := synCustomHighlighterAPI().SysCallN(46, m.Instance(), uintptr(index))
	return AsSynDividerDrawConfig(r)
}

func (m *TSynCustomHighlighter) DividerDrawConfigCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomHighlighterAPI().SysCallN(47, m.Instance())
	return int32(r)
}

func (m *TSynCustomHighlighter) DefaultFilter() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synCustomHighlighterAPI().SysCallN(48, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynCustomHighlighter) SetDefaultFilter(value string) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(48, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynCustomHighlighter) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomHighlighterAPI().SysCallN(49, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomHighlighter) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	synCustomHighlighterAPI().SysCallN(49, 1, m.Instance(), api.PasBool(value))
}

// SynCustomHighlighter  is static instance
var SynCustomHighlighter _SynCustomHighlighterClass

// _SynCustomHighlighterClass is class type defined by TSynCustomHighlighter
type _SynCustomHighlighterClass uintptr

func (_SynCustomHighlighterClass) GetCapabilities() types.TSynHighlighterCapabilities {
	r := synCustomHighlighterAPI().SysCallN(0)
	return types.TSynHighlighterCapabilities(r)
}

func (_SynCustomHighlighterClass) GetLanguageName() string {
	r := synCustomHighlighterAPI().SysCallN(1)
	return api.GoStr(r)
}

var (
	synCustomHighlighterOnce   base.Once
	synCustomHighlighterImport *imports.Imports = nil
)

func synCustomHighlighterAPI() *imports.Imports {
	synCustomHighlighterOnce.Do(func() {
		synCustomHighlighterImport = api.NewDefaultImports()
		synCustomHighlighterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCustomHighlighter_GetCapabilities", 0), // static function GetCapabilities
			/* 1 */ imports.NewTable("TSynCustomHighlighter_GetLanguageName", 0), // static function GetLanguageName
			/* 2 */ imports.NewTable("TSynCustomHighlighter_AddSpecialAttribute", 0), // function AddSpecialAttribute
			/* 3 */ imports.NewTable("TSynCustomHighlighter_GetEol", 0), // function GetEol
			/* 4 */ imports.NewTable("TSynCustomHighlighter_GetRange", 0), // function GetRange
			/* 5 */ imports.NewTable("TSynCustomHighlighter_GetToken", 0), // function GetToken
			/* 6 */ imports.NewTable("TSynCustomHighlighter_GetEndOfLineAttribute", 0), // function GetEndOfLineAttribute
			/* 7 */ imports.NewTable("TSynCustomHighlighter_GetTokenAttribute", 0), // function GetTokenAttribute
			/* 8 */ imports.NewTable("TSynCustomHighlighter_GetTokenKind", 0), // function GetTokenKind
			/* 9 */ imports.NewTable("TSynCustomHighlighter_GetTokenPos", 0), // function GetTokenPos
			/* 10 */ imports.NewTable("TSynCustomHighlighter_GetTokenLen", 0), // function GetTokenLen
			/* 11 */ imports.NewTable("TSynCustomHighlighter_IsKeyword", 0), // function IsKeyword
			/* 12 */ imports.NewTable("TSynCustomHighlighter_IdleScanRanges", 0), // function IdleScanRanges
			/* 13 */ imports.NewTable("TSynCustomHighlighter_NeedScan", 0), // function NeedScan
			/* 14 */ imports.NewTable("TSynCustomHighlighter_UseUserSettings", 0), // function UseUserSettings
			/* 15 */ imports.NewTable("TSynCustomHighlighter_LoadFromRegistry", 0), // function LoadFromRegistry
			/* 16 */ imports.NewTable("TSynCustomHighlighter_SaveToRegistry", 0), // function SaveToRegistry
			/* 17 */ imports.NewTable("TSynCustomHighlighter_LoadFromFile", 0), // function LoadFromFile
			/* 18 */ imports.NewTable("TSynCustomHighlighter_SaveToFile", 0), // function SaveToFile
			/* 19 */ imports.NewTable("TSynCustomHighlighter_DefHighlightChange", 0), // procedure DefHighlightChange
			/* 20 */ imports.NewTable("TSynCustomHighlighter_BeginUpdate", 0), // procedure BeginUpdate
			/* 21 */ imports.NewTable("TSynCustomHighlighter_EndUpdate", 0), // procedure EndUpdate
			/* 22 */ imports.NewTable("TSynCustomHighlighter_AttachToLines", 0), // procedure AttachToLines
			/* 23 */ imports.NewTable("TSynCustomHighlighter_DetachFromLines", 0), // procedure DetachFromLines
			/* 24 */ imports.NewTable("TSynCustomHighlighter_GetTokenEx", 0), // procedure GetTokenEx
			/* 25 */ imports.NewTable("TSynCustomHighlighter_Next", 0), // procedure Next
			/* 26 */ imports.NewTable("TSynCustomHighlighter_NextToEol", 0), // procedure NextToEol
			/* 27 */ imports.NewTable("TSynCustomHighlighter_StartAtLineIndex", 0), // procedure StartAtLineIndex
			/* 28 */ imports.NewTable("TSynCustomHighlighter_ContinueNextLine", 0), // procedure ContinueNextLine
			/* 29 */ imports.NewTable("TSynCustomHighlighter_ScanRanges", 0), // procedure ScanRanges
			/* 30 */ imports.NewTable("TSynCustomHighlighter_ScanAllRanges", 0), // procedure ScanAllRanges
			/* 31 */ imports.NewTable("TSynCustomHighlighter_SetRange", 0), // procedure SetRange
			/* 32 */ imports.NewTable("TSynCustomHighlighter_ResetRange", 0), // procedure ResetRange
			/* 33 */ imports.NewTable("TSynCustomHighlighter_SetLine", 0), // procedure SetLine
			/* 34 */ imports.NewTable("TSynCustomHighlighter_EnumUserSettings", 0), // procedure EnumUserSettings
			/* 35 */ imports.NewTable("TSynCustomHighlighter_AttributeChangeNeedScan", 0), // property AttributeChangeNeedScan
			/* 36 */ imports.NewTable("TSynCustomHighlighter_DrawDivider", 0), // property DrawDivider
			/* 37 */ imports.NewTable("TSynCustomHighlighter_CurrentLines", 0), // property CurrentLines
			/* 38 */ imports.NewTable("TSynCustomHighlighter_LineIndex", 0), // property LineIndex
			/* 39 */ imports.NewTable("TSynCustomHighlighter_IdentChars", 0), // property IdentChars
			/* 40 */ imports.NewTable("TSynCustomHighlighter_WordBreakChars", 0), // property WordBreakChars
			/* 41 */ imports.NewTable("TSynCustomHighlighter_LanguageName", 0), // property LanguageName
			/* 42 */ imports.NewTable("TSynCustomHighlighter_AttrCount", 0), // property AttrCount
			/* 43 */ imports.NewTable("TSynCustomHighlighter_Attribute", 0), // property Attribute
			/* 44 */ imports.NewTable("TSynCustomHighlighter_Capabilities", 0), // property Capabilities
			/* 45 */ imports.NewTable("TSynCustomHighlighter_SampleSource", 0), // property SampleSource
			/* 46 */ imports.NewTable("TSynCustomHighlighter_DividerDrawConfig", 0), // property DividerDrawConfig
			/* 47 */ imports.NewTable("TSynCustomHighlighter_DividerDrawConfigCount", 0), // property DividerDrawConfigCount
			/* 48 */ imports.NewTable("TSynCustomHighlighter_DefaultFilter", 0), // property DefaultFilter
			/* 49 */ imports.NewTable("TSynCustomHighlighter_Enabled", 0), // property Enabled
		}
	})
	return synCustomHighlighterImport
}
