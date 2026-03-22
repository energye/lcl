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

// ISynEditKeyStroke Parent: ICollectionItem
type ISynEditKeyStroke interface {
	ICollectionItem
	// LoadFromStream
	//  TODO: add ShiftMask optional
	LoadFromStream(stream IStream) // procedure
	SaveToStream(stream IStream)   // procedure
	// Key
	//  No duplicate checking is done if assignment made via these properties!
	Key() uint16                       // property Key Getter
	SetKey(value uint16)               // property Key Setter
	Key2() uint16                      // property Key2 Getter
	SetKey2(value uint16)              // property Key2 Setter
	Shift() types.TShiftState          // property Shift Getter
	SetShift(value types.TShiftState)  // property Shift Setter
	Shift2() types.TShiftState         // property Shift2 Getter
	SetShift2(value types.TShiftState) // property Shift2 Setter
	// ShiftMask
	//  Modifier keys, that should be ignored
	//  * ShiftMask:
	//  ShiftStates that are SET in the mask will be ignored, and NOT compared with "Shift"
	ShiftMask() types.TShiftState             // property ShiftMask Getter
	SetShiftMask(value types.TShiftState)     // property ShiftMask Setter
	ShiftMask2() types.TShiftState            // property ShiftMask2 Getter
	SetShiftMask2(value types.TShiftState)    // property ShiftMask2 Setter
	Command() types.TSynEditorCommand         // property Command Getter
	SetCommand(value types.TSynEditorCommand) // property Command Setter
	ShortCut() types.TShortCut                // property ShortCut Getter
	SetShortCut(value types.TShortCut)        // property ShortCut Setter
	ShortCut2() types.TShortCut               // property ShortCut2 Getter
	SetShortCut2(value types.TShortCut)       // property ShortCut2 Setter
}

type TSynEditKeyStroke struct {
	TCollectionItem
}

func (m *TSynEditKeyStroke) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TSynEditKeyStroke) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TSynEditKeyStroke) Key() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(3, 0, m.Instance())
	return uint16(r)
}

func (m *TSynEditKeyStroke) SetKey(value uint16) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) Key2() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(4, 0, m.Instance())
	return uint16(r)
}

func (m *TSynEditKeyStroke) SetKey2(value uint16) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) Shift() types.TShiftState {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(5, 0, m.Instance())
	return types.TShiftState(r)
}

func (m *TSynEditKeyStroke) SetShift(value types.TShiftState) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) Shift2() types.TShiftState {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(6, 0, m.Instance())
	return types.TShiftState(r)
}

func (m *TSynEditKeyStroke) SetShift2(value types.TShiftState) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) ShiftMask() types.TShiftState {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(7, 0, m.Instance())
	return types.TShiftState(r)
}

func (m *TSynEditKeyStroke) SetShiftMask(value types.TShiftState) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) ShiftMask2() types.TShiftState {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(8, 0, m.Instance())
	return types.TShiftState(r)
}

func (m *TSynEditKeyStroke) SetShiftMask2(value types.TShiftState) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) Command() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(9, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynEditKeyStroke) SetCommand(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) ShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(10, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TSynEditKeyStroke) SetShortCut(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditKeyStroke) ShortCut2() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := synEditKeyStrokeAPI().SysCallN(11, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TSynEditKeyStroke) SetShortCut2(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	synEditKeyStrokeAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

// NewSynEditKeyStroke class constructor
func NewSynEditKeyStroke(collection ICollection) ISynEditKeyStroke {
	r := synEditKeyStrokeAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsSynEditKeyStroke(r)
}

var (
	synEditKeyStrokeOnce   base.Once
	synEditKeyStrokeImport *imports.Imports = nil
)

func synEditKeyStrokeAPI() *imports.Imports {
	synEditKeyStrokeOnce.Do(func() {
		synEditKeyStrokeImport = api.NewDefaultImports()
		synEditKeyStrokeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditKeyStroke_Create", 0), // constructor NewSynEditKeyStroke
			/* 1 */ imports.NewTable("TSynEditKeyStroke_LoadFromStream", 0), // procedure LoadFromStream
			/* 2 */ imports.NewTable("TSynEditKeyStroke_SaveToStream", 0), // procedure SaveToStream
			/* 3 */ imports.NewTable("TSynEditKeyStroke_Key", 0), // property Key
			/* 4 */ imports.NewTable("TSynEditKeyStroke_Key2", 0), // property Key2
			/* 5 */ imports.NewTable("TSynEditKeyStroke_Shift", 0), // property Shift
			/* 6 */ imports.NewTable("TSynEditKeyStroke_Shift2", 0), // property Shift2
			/* 7 */ imports.NewTable("TSynEditKeyStroke_ShiftMask", 0), // property ShiftMask
			/* 8 */ imports.NewTable("TSynEditKeyStroke_ShiftMask2", 0), // property ShiftMask2
			/* 9 */ imports.NewTable("TSynEditKeyStroke_Command", 0), // property Command
			/* 10 */ imports.NewTable("TSynEditKeyStroke_ShortCut", 0), // property ShortCut
			/* 11 */ imports.NewTable("TSynEditKeyStroke_ShortCut2", 0), // property ShortCut2
		}
	})
	return synEditKeyStrokeImport
}
