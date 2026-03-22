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

// ICustomSynMacroRecorder Parent: IAbstractSynHookerPlugin
type ICustomSynMacroRecorder interface {
	IAbstractSynHookerPlugin
	Error(msg string)                                                                // procedure
	RecordMacro(editor ICustomSynEdit)                                               // procedure
	PlaybackMacro(editor ICustomSynEdit)                                             // procedure
	Stop()                                                                           // procedure
	Pause()                                                                          // procedure
	Resume()                                                                         // procedure
	AssignEventsFrom(macroRecorder ICustomSynMacroRecorder)                          // procedure
	Clear()                                                                          // procedure
	AddEvent(cmd types.TSynEditorCommand, char uint16, data uintptr)                 // procedure
	InsertEvent(index int32, cmd types.TSynEditorCommand, char uint16, data uintptr) // procedure
	AddCustomEvent(event ISynMacroEvent)                                             // procedure
	InsertCustomEvent(index int32, event ISynMacroEvent)                             // procedure
	DeleteEvent(index int32)                                                         // procedure
	LoadFromStream(src IStream)                                                      // procedure
	LoadFromStreamEx(src IStream, clear bool)                                        // procedure
	SaveToStream(dest IStream)                                                       // procedure
	LoadFromFile(filename string)                                                    // procedure
	SaveToFile(filename string)                                                      // procedure
	IsEmpty() bool                                                                   // property IsEmpty Getter
	State() types.TSynMacroState                                                     // property State Getter
	EventCount() int32                                                               // property EventCount Getter
	Events(index int32) ISynMacroEvent                                               // property Events Getter
	CurrentEditor() ICustomSynEdit                                                   // property CurrentEditor Getter
	RecordShortCut() types.TShortCut                                                 // property RecordShortCut Getter
	SetRecordShortCut(value types.TShortCut)                                         // property RecordShortCut Setter
	PlaybackShortCut() types.TShortCut                                               // property PlaybackShortCut Getter
	SetPlaybackShortCut(value types.TShortCut)                                       // property PlaybackShortCut Setter
	SaveMarkerPos() bool                                                             // property SaveMarkerPos Getter
	SetSaveMarkerPos(value bool)                                                     // property SaveMarkerPos Setter
	AsString() string                                                                // property AsString Getter
	SetAsString(value string)                                                        // property AsString Setter
	MacroName() string                                                               // property MacroName Getter
	SetMacroName(value string)                                                       // property MacroName Setter
	SetOnStateChange(fn TNotifyEvent)                                                // property event
	SetOnUserCommand(fn TSynUserCommandEvent)                                        // property event
}

type TCustomSynMacroRecorder struct {
	TAbstractSynHookerPlugin
}

func (m *TCustomSynMacroRecorder) Error(msg string) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(1, m.Instance(), api.PasStr(msg))
}

func (m *TCustomSynMacroRecorder) RecordMacro(editor ICustomSynEdit) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(editor))
}

func (m *TCustomSynMacroRecorder) PlaybackMacro(editor ICustomSynEdit) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(editor))
}

func (m *TCustomSynMacroRecorder) Stop() {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(4, m.Instance())
}

func (m *TCustomSynMacroRecorder) Pause() {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(5, m.Instance())
}

func (m *TCustomSynMacroRecorder) Resume() {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(6, m.Instance())
}

func (m *TCustomSynMacroRecorder) AssignEventsFrom(macroRecorder ICustomSynMacroRecorder) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(macroRecorder))
}

func (m *TCustomSynMacroRecorder) Clear() {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(8, m.Instance())
}

func (m *TCustomSynMacroRecorder) AddEvent(cmd types.TSynEditorCommand, char uint16, data uintptr) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(9, m.Instance(), uintptr(cmd), uintptr(char), uintptr(data))
}

func (m *TCustomSynMacroRecorder) InsertEvent(index int32, cmd types.TSynEditorCommand, char uint16, data uintptr) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(10, m.Instance(), uintptr(index), uintptr(cmd), uintptr(char), uintptr(data))
}

func (m *TCustomSynMacroRecorder) AddCustomEvent(event ISynMacroEvent) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(event))
}

func (m *TCustomSynMacroRecorder) InsertCustomEvent(index int32, event ISynMacroEvent) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(12, m.Instance(), uintptr(index), base.GetObjectUintptr(event))
}

func (m *TCustomSynMacroRecorder) DeleteEvent(index int32) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(13, m.Instance(), uintptr(index))
}

func (m *TCustomSynMacroRecorder) LoadFromStream(src IStream) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(src))
}

func (m *TCustomSynMacroRecorder) LoadFromStreamEx(src IStream, clear bool) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(src), api.PasBool(clear))
}

func (m *TCustomSynMacroRecorder) SaveToStream(dest IStream) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(dest))
}

func (m *TCustomSynMacroRecorder) LoadFromFile(filename string) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(17, m.Instance(), api.PasStr(filename))
}

func (m *TCustomSynMacroRecorder) SaveToFile(filename string) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(18, m.Instance(), api.PasStr(filename))
}

func (m *TCustomSynMacroRecorder) IsEmpty() bool {
	if !m.IsValid() {
		return false
	}
	r := customSynMacroRecorderAPI().SysCallN(19, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSynMacroRecorder) State() types.TSynMacroState {
	if !m.IsValid() {
		return 0
	}
	r := customSynMacroRecorderAPI().SysCallN(20, m.Instance())
	return types.TSynMacroState(r)
}

func (m *TCustomSynMacroRecorder) EventCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynMacroRecorderAPI().SysCallN(21, m.Instance())
	return int32(r)
}

func (m *TCustomSynMacroRecorder) Events(index int32) ISynMacroEvent {
	if !m.IsValid() {
		return nil
	}
	r := customSynMacroRecorderAPI().SysCallN(22, m.Instance(), uintptr(index))
	return AsSynMacroEvent(r)
}

func (m *TCustomSynMacroRecorder) CurrentEditor() ICustomSynEdit {
	if !m.IsValid() {
		return nil
	}
	r := customSynMacroRecorderAPI().SysCallN(23, m.Instance())
	return AsCustomSynEdit(r)
}

func (m *TCustomSynMacroRecorder) RecordShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := customSynMacroRecorderAPI().SysCallN(24, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TCustomSynMacroRecorder) SetRecordShortCut(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynMacroRecorder) PlaybackShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := customSynMacroRecorderAPI().SysCallN(25, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TCustomSynMacroRecorder) SetPlaybackShortCut(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynMacroRecorder) SaveMarkerPos() bool {
	if !m.IsValid() {
		return false
	}
	r := customSynMacroRecorderAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSynMacroRecorder) SetSaveMarkerPos(value bool) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSynMacroRecorder) AsString() string {
	if !m.IsValid() {
		return ""
	}
	r := customSynMacroRecorderAPI().SysCallN(27, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomSynMacroRecorder) SetAsString(value string) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(27, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomSynMacroRecorder) MacroName() string {
	if !m.IsValid() {
		return ""
	}
	r := customSynMacroRecorderAPI().SysCallN(28, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomSynMacroRecorder) SetMacroName(value string) {
	if !m.IsValid() {
		return
	}
	customSynMacroRecorderAPI().SysCallN(28, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomSynMacroRecorder) SetOnStateChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 29, customSynMacroRecorderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynMacroRecorder) SetOnUserCommand(fn TSynUserCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynUserCommandEvent(fn)
	base.SetEvent(m, 30, customSynMacroRecorderAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomSynMacroRecorder class constructor
func NewCustomSynMacroRecorder(owner IComponent) ICustomSynMacroRecorder {
	r := customSynMacroRecorderAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomSynMacroRecorder(r)
}

func TCustomSynMacroRecorderClass() types.TClass {
	r := customSynMacroRecorderAPI().SysCallN(31)
	return types.TClass(r)
}

var (
	customSynMacroRecorderOnce   base.Once
	customSynMacroRecorderImport *imports.Imports = nil
)

func customSynMacroRecorderAPI() *imports.Imports {
	customSynMacroRecorderOnce.Do(func() {
		customSynMacroRecorderImport = api.NewDefaultImports()
		customSynMacroRecorderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomSynMacroRecorder_Create", 0), // constructor NewCustomSynMacroRecorder
			/* 1 */ imports.NewTable("TCustomSynMacroRecorder_Error", 0), // procedure Error
			/* 2 */ imports.NewTable("TCustomSynMacroRecorder_RecordMacro", 0), // procedure RecordMacro
			/* 3 */ imports.NewTable("TCustomSynMacroRecorder_PlaybackMacro", 0), // procedure PlaybackMacro
			/* 4 */ imports.NewTable("TCustomSynMacroRecorder_Stop", 0), // procedure Stop
			/* 5 */ imports.NewTable("TCustomSynMacroRecorder_Pause", 0), // procedure Pause
			/* 6 */ imports.NewTable("TCustomSynMacroRecorder_Resume", 0), // procedure Resume
			/* 7 */ imports.NewTable("TCustomSynMacroRecorder_AssignEventsFrom", 0), // procedure AssignEventsFrom
			/* 8 */ imports.NewTable("TCustomSynMacroRecorder_Clear", 0), // procedure Clear
			/* 9 */ imports.NewTable("TCustomSynMacroRecorder_AddEvent", 0), // procedure AddEvent
			/* 10 */ imports.NewTable("TCustomSynMacroRecorder_InsertEvent", 0), // procedure InsertEvent
			/* 11 */ imports.NewTable("TCustomSynMacroRecorder_AddCustomEvent", 0), // procedure AddCustomEvent
			/* 12 */ imports.NewTable("TCustomSynMacroRecorder_InsertCustomEvent", 0), // procedure InsertCustomEvent
			/* 13 */ imports.NewTable("TCustomSynMacroRecorder_DeleteEvent", 0), // procedure DeleteEvent
			/* 14 */ imports.NewTable("TCustomSynMacroRecorder_LoadFromStream", 0), // procedure LoadFromStream
			/* 15 */ imports.NewTable("TCustomSynMacroRecorder_LoadFromStreamEx", 0), // procedure LoadFromStreamEx
			/* 16 */ imports.NewTable("TCustomSynMacroRecorder_SaveToStream", 0), // procedure SaveToStream
			/* 17 */ imports.NewTable("TCustomSynMacroRecorder_LoadFromFile", 0), // procedure LoadFromFile
			/* 18 */ imports.NewTable("TCustomSynMacroRecorder_SaveToFile", 0), // procedure SaveToFile
			/* 19 */ imports.NewTable("TCustomSynMacroRecorder_IsEmpty", 0), // property IsEmpty
			/* 20 */ imports.NewTable("TCustomSynMacroRecorder_State", 0), // property State
			/* 21 */ imports.NewTable("TCustomSynMacroRecorder_EventCount", 0), // property EventCount
			/* 22 */ imports.NewTable("TCustomSynMacroRecorder_Events", 0), // property Events
			/* 23 */ imports.NewTable("TCustomSynMacroRecorder_CurrentEditor", 0), // property CurrentEditor
			/* 24 */ imports.NewTable("TCustomSynMacroRecorder_RecordShortCut", 0), // property RecordShortCut
			/* 25 */ imports.NewTable("TCustomSynMacroRecorder_PlaybackShortCut", 0), // property PlaybackShortCut
			/* 26 */ imports.NewTable("TCustomSynMacroRecorder_SaveMarkerPos", 0), // property SaveMarkerPos
			/* 27 */ imports.NewTable("TCustomSynMacroRecorder_AsString", 0), // property AsString
			/* 28 */ imports.NewTable("TCustomSynMacroRecorder_MacroName", 0), // property MacroName
			/* 29 */ imports.NewTable("TCustomSynMacroRecorder_OnStateChange", 0), // event OnStateChange
			/* 30 */ imports.NewTable("TCustomSynMacroRecorder_OnUserCommand", 0), // event OnUserCommand
			/* 31 */ imports.NewTable("TCustomSynMacroRecorder_TClass", 0), // function TCustomSynMacroRecorderClass
		}
	})
	return customSynMacroRecorderImport
}
