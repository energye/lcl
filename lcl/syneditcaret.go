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

// ISynEditCaret Parent: ISynEditBaseCaret
type ISynEditCaret interface {
	ISynEditBaseCaret
	WasAtLineChar(point types.TPoint) bool                // function
	WasAtLineByte(point types.TPoint) bool                // function
	MoveHoriz(count int32) bool                           // function
	AssignFromWithSynEditBaseCaret(src ISynEditBaseCaret) // procedure
	IncForcePastEOL()                                     // procedure
	DecForcePastEOL()                                     // procedure
	IncForceAdjustToNextChar()                            // procedure
	DecForceAdjustToNextChar()                            // procedure
	IncAutoMoveOnEdit()                                   // procedure
	DecAutoMoveOnEdit()                                   // procedure
	ChangeOnTouch()                                       // procedure
	Touch(changeOnTouch bool)                             // procedure
	OldLinePos() int32                                    // property OldLinePos Getter
	OldCharPos() int32                                    // property OldCharPos Getter
	OldLineCharPos() types.TPoint                         // property OldLineCharPos Getter
	OldLineBytePos() types.TPoint                         // property OldLineBytePos Getter
	OldFullLogicalPos() TLogCaretPoint                    // property OldFullLogicalPos Getter
	SkipTabs() bool                                       // property SkipTabs Getter
	SetSkipTabs(value bool)                               // property SkipTabs Setter
	AllowPastEOL() bool                                   // property AllowPastEOL Getter
	SetAllowPastEOL(value bool)                           // property AllowPastEOL Setter
	KeepCaretX() bool                                     // property KeepCaretX Getter
	SetKeepCaretX(value bool)                             // property KeepCaretX Setter
	KeepCaretXPos() int32                                 // property KeepCaretXPos Getter
	SetKeepCaretXPos(value int32)                         // property KeepCaretXPos Setter
	SetMaxLeftChar(fn TMaxLeftCharFunc)                   // property event
}

type TSynEditCaret struct {
	TSynEditBaseCaret
}

func (m *TSynEditCaret) WasAtLineChar(point types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditCaretAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&point)))
	return api.GoBool(r)
}

func (m *TSynEditCaret) WasAtLineByte(point types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditCaretAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&point)))
	return api.GoBool(r)
}

func (m *TSynEditCaret) MoveHoriz(count int32) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditCaretAPI().SysCallN(3, m.Instance(), uintptr(count))
	return api.GoBool(r)
}

func (m *TSynEditCaret) AssignFromWithSynEditBaseCaret(src ISynEditBaseCaret) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(src))
}

func (m *TSynEditCaret) IncForcePastEOL() {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(5, m.Instance())
}

func (m *TSynEditCaret) DecForcePastEOL() {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(6, m.Instance())
}

func (m *TSynEditCaret) IncForceAdjustToNextChar() {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(7, m.Instance())
}

func (m *TSynEditCaret) DecForceAdjustToNextChar() {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(8, m.Instance())
}

func (m *TSynEditCaret) IncAutoMoveOnEdit() {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(9, m.Instance())
}

func (m *TSynEditCaret) DecAutoMoveOnEdit() {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(10, m.Instance())
}

func (m *TSynEditCaret) ChangeOnTouch() {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(11, m.Instance())
}

func (m *TSynEditCaret) Touch(changeOnTouch bool) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(12, m.Instance(), api.PasBool(changeOnTouch))
}

func (m *TSynEditCaret) OldLinePos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditCaretAPI().SysCallN(13, m.Instance())
	return int32(r)
}

func (m *TSynEditCaret) OldCharPos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditCaretAPI().SysCallN(14, m.Instance())
	return int32(r)
}

func (m *TSynEditCaret) OldLineCharPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditCaret) OldLineBytePos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditCaret) OldFullLogicalPos() (result TLogCaretPoint) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditCaret) SkipTabs() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditCaretAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditCaret) SetSkipTabs(value bool) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditCaret) AllowPastEOL() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditCaretAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditCaret) SetAllowPastEOL(value bool) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditCaret) KeepCaretX() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditCaretAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditCaret) SetKeepCaretX(value bool) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditCaret) KeepCaretXPos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditCaretAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditCaret) SetKeepCaretXPos(value int32) {
	if !m.IsValid() {
		return
	}
	synEditCaretAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditCaret) SetMaxLeftChar(fn TMaxLeftCharFunc) {
	if !m.IsValid() {
		return
	}
	cb := makeTMaxLeftCharFunc(fn)
	base.SetEvent(m, 22, synEditCaretAPI(), api.MakeEventDataPtr(cb))
}

// NewSynEditCaret class constructor
func NewSynEditCaret() ISynEditCaret {
	r := synEditCaretAPI().SysCallN(0)
	return AsSynEditCaret(r)
}

var (
	synEditCaretOnce   base.Once
	synEditCaretImport *imports.Imports = nil
)

func synEditCaretAPI() *imports.Imports {
	synEditCaretOnce.Do(func() {
		synEditCaretImport = api.NewDefaultImports()
		synEditCaretImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditCaret_Create", 0), // constructor NewSynEditCaret
			/* 1 */ imports.NewTable("TSynEditCaret_WasAtLineChar", 0), // function WasAtLineChar
			/* 2 */ imports.NewTable("TSynEditCaret_WasAtLineByte", 0), // function WasAtLineByte
			/* 3 */ imports.NewTable("TSynEditCaret_MoveHoriz", 0), // function MoveHoriz
			/* 4 */ imports.NewTable("TSynEditCaret_AssignFromWithSynEditBaseCaret", 0), // procedure AssignFromWithSynEditBaseCaret
			/* 5 */ imports.NewTable("TSynEditCaret_IncForcePastEOL", 0), // procedure IncForcePastEOL
			/* 6 */ imports.NewTable("TSynEditCaret_DecForcePastEOL", 0), // procedure DecForcePastEOL
			/* 7 */ imports.NewTable("TSynEditCaret_IncForceAdjustToNextChar", 0), // procedure IncForceAdjustToNextChar
			/* 8 */ imports.NewTable("TSynEditCaret_DecForceAdjustToNextChar", 0), // procedure DecForceAdjustToNextChar
			/* 9 */ imports.NewTable("TSynEditCaret_IncAutoMoveOnEdit", 0), // procedure IncAutoMoveOnEdit
			/* 10 */ imports.NewTable("TSynEditCaret_DecAutoMoveOnEdit", 0), // procedure DecAutoMoveOnEdit
			/* 11 */ imports.NewTable("TSynEditCaret_ChangeOnTouch", 0), // procedure ChangeOnTouch
			/* 12 */ imports.NewTable("TSynEditCaret_Touch", 0), // procedure Touch
			/* 13 */ imports.NewTable("TSynEditCaret_OldLinePos", 0), // property OldLinePos
			/* 14 */ imports.NewTable("TSynEditCaret_OldCharPos", 0), // property OldCharPos
			/* 15 */ imports.NewTable("TSynEditCaret_OldLineCharPos", 0), // property OldLineCharPos
			/* 16 */ imports.NewTable("TSynEditCaret_OldLineBytePos", 0), // property OldLineBytePos
			/* 17 */ imports.NewTable("TSynEditCaret_OldFullLogicalPos", 0), // property OldFullLogicalPos
			/* 18 */ imports.NewTable("TSynEditCaret_SkipTabs", 0), // property SkipTabs
			/* 19 */ imports.NewTable("TSynEditCaret_AllowPastEOL", 0), // property AllowPastEOL
			/* 20 */ imports.NewTable("TSynEditCaret_KeepCaretX", 0), // property KeepCaretX
			/* 21 */ imports.NewTable("TSynEditCaret_KeepCaretXPos", 0), // property KeepCaretXPos
			/* 22 */ imports.NewTable("TSynEditCaret_MaxLeftChar", 0), // event MaxLeftChar
		}
	})
	return synEditCaretImport
}
