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

// ISynEditKeyStrokes Parent: ICollection
type ISynEditKeyStrokes interface {
	ICollection
	AddToSynEditKeyStroke() ISynEditKeyStroke                                                                                                                             // function
	FindCommand(cmd types.TSynEditorCommand) int32                                                                                                                        // function
	FindKeycode(code uint16, sS types.TShiftState) int32                                                                                                                  // function
	FindKeycodeEx(code uint16, sS types.TShiftState, data *uintptr, outIsStartOfCombo *bool, finishComboOnly bool, comboStart ISynEditKeyStrokes) types.TSynEditorCommand // function
	FindShortcut(sC types.TShortCut) int32                                                                                                                                // function
	FindShortcut2(sC types.TShortCut, sC2 types.TShortCut) int32                                                                                                          // function
	ResetKeyCombo()                                                                                                                                                       // procedure
	LoadFromStream(stream IStream)                                                                                                                                        // procedure
	ResetDefaults()                                                                                                                                                       // procedure
	SaveToStream(stream IStream)                                                                                                                                          // procedure
	ItemsWithIntToSynEditKeyStroke(index int32) ISynEditKeyStroke                                                                                                         // property Items Getter
	SetItemsWithIntToSynEditKeyStroke(index int32, value ISynEditKeyStroke)                                                                                               // property Items Setter
	PluginOffset() int32                                                                                                                                                  // property PluginOffset Getter
	SetPluginOffset(value int32)                                                                                                                                          // property PluginOffset Setter
	// UsePluginOffset
	//  only switch on while needed.
	//  So streaming will always see the constant, unmodified values
	UsePluginOffset() bool         // property UsePluginOffset Getter
	SetUsePluginOffset(value bool) // property UsePluginOffset Setter
}

type TSynEditKeyStrokes struct {
	TCollection
}

func (m *TSynEditKeyStrokes) AddToSynEditKeyStroke() ISynEditKeyStroke {
	if !m.IsValid() {
		return nil
	}
	r := synEditKeyStrokesAPI().SysCallN(1, m.Instance())
	return AsSynEditKeyStroke(r)
}

func (m *TSynEditKeyStrokes) FindCommand(cmd types.TSynEditorCommand) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokesAPI().SysCallN(2, m.Instance(), uintptr(cmd))
	return int32(r)
}

func (m *TSynEditKeyStrokes) FindKeycode(code uint16, sS types.TShiftState) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokesAPI().SysCallN(3, m.Instance(), uintptr(code), uintptr(sS))
	return int32(r)
}

func (m *TSynEditKeyStrokes) FindKeycodeEx(code uint16, sS types.TShiftState, data *uintptr, outIsStartOfCombo *bool, finishComboOnly bool, comboStart ISynEditKeyStrokes) types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	dataPtr := uintptr(*data)
	r := synEditKeyStrokesAPI().SysCallN(4, m.Instance(), uintptr(code), uintptr(sS), uintptr(base.UnsafePointer(&dataPtr)), uintptr(base.UnsafePointer(outIsStartOfCombo)), api.PasBool(finishComboOnly), base.GetObjectUintptr(comboStart))
	*data = uintptr(dataPtr)
	return types.TSynEditorCommand(r)
}

func (m *TSynEditKeyStrokes) FindShortcut(sC types.TShortCut) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokesAPI().SysCallN(5, m.Instance(), uintptr(sC))
	return int32(r)
}

func (m *TSynEditKeyStrokes) FindShortcut2(sC types.TShortCut, sC2 types.TShortCut) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokesAPI().SysCallN(6, m.Instance(), uintptr(sC), uintptr(sC2))
	return int32(r)
}

func (m *TSynEditKeyStrokes) ResetKeyCombo() {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokesAPI().SysCallN(7, m.Instance())
}

func (m *TSynEditKeyStrokes) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokesAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TSynEditKeyStrokes) ResetDefaults() {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokesAPI().SysCallN(9, m.Instance())
}

func (m *TSynEditKeyStrokes) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokesAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TSynEditKeyStrokes) ItemsWithIntToSynEditKeyStroke(index int32) ISynEditKeyStroke {
	if !m.IsValid() {
		return nil
	}
	r := synEditKeyStrokesAPI().SysCallN(11, 0, m.Instance(), uintptr(index))
	return AsSynEditKeyStroke(r)
}

func (m *TSynEditKeyStrokes) SetItemsWithIntToSynEditKeyStroke(index int32, value ISynEditKeyStroke) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokesAPI().SysCallN(11, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TSynEditKeyStrokes) PluginOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokesAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditKeyStrokes) SetPluginOffset(value int32) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokesAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStrokes) UsePluginOffset() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditKeyStrokesAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditKeyStrokes) SetUsePluginOffset(value bool) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokesAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

// NewSynEditKeyStrokes class constructor
func NewSynEditKeyStrokes(owner IPersistent) ISynEditKeyStrokes {
	r := synEditKeyStrokesAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynEditKeyStrokes(r)
}

var (
	synEditKeyStrokesOnce   base.Once
	synEditKeyStrokesImport *imports.Imports = nil
)

func synEditKeyStrokesAPI() *imports.Imports {
	synEditKeyStrokesOnce.Do(func() {
		synEditKeyStrokesImport = api.NewDefaultImports()
		synEditKeyStrokesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditKeyStrokes_Create", 0), // constructor NewSynEditKeyStrokes
			/* 1 */ imports.NewTable("TSynEditKeyStrokes_AddToSynEditKeyStroke", 0), // function AddToSynEditKeyStroke
			/* 2 */ imports.NewTable("TSynEditKeyStrokes_FindCommand", 0), // function FindCommand
			/* 3 */ imports.NewTable("TSynEditKeyStrokes_FindKeycode", 0), // function FindKeycode
			/* 4 */ imports.NewTable("TSynEditKeyStrokes_FindKeycodeEx", 0), // function FindKeycodeEx
			/* 5 */ imports.NewTable("TSynEditKeyStrokes_FindShortcut", 0), // function FindShortcut
			/* 6 */ imports.NewTable("TSynEditKeyStrokes_FindShortcut2", 0), // function FindShortcut2
			/* 7 */ imports.NewTable("TSynEditKeyStrokes_ResetKeyCombo", 0), // procedure ResetKeyCombo
			/* 8 */ imports.NewTable("TSynEditKeyStrokes_LoadFromStream", 0), // procedure LoadFromStream
			/* 9 */ imports.NewTable("TSynEditKeyStrokes_ResetDefaults", 0), // procedure ResetDefaults
			/* 10 */ imports.NewTable("TSynEditKeyStrokes_SaveToStream", 0), // procedure SaveToStream
			/* 11 */ imports.NewTable("TSynEditKeyStrokes_ItemsWithIntToSynEditKeyStroke", 0), // property ItemsWithIntToSynEditKeyStroke
			/* 12 */ imports.NewTable("TSynEditKeyStrokes_PluginOffset", 0), // property PluginOffset
			/* 13 */ imports.NewTable("TSynEditKeyStrokes_UsePluginOffset", 0), // property UsePluginOffset
		}
	})
	return synEditKeyStrokesImport
}
