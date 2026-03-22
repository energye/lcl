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

// ISynEditBaseCaret Parent: ISynEditPointBase
type ISynEditBaseCaret interface {
	ISynEditPointBase
	IsAtLineChar(point types.TPoint) bool                   // function
	IsAtLineByte(point types.TPoint, byteOffset int32) bool // function
	IsAtPos(caret ISynEditCaret) bool                       // function
	AssignFromWithSynEditBaseCaret(src ISynEditBaseCaret)   // procedure
	Invalidate()                                            // procedure
	InvalidateBytePos()                                     // procedure
	InvalidateCharPos()                                     // procedure
	LinePos() int32                                         // property LinePos Getter
	SetLinePos(value int32)                                 // property LinePos Setter
	CharPos() int32                                         // property CharPos Getter
	SetCharPos(value int32)                                 // property CharPos Setter
	LineCharPos() types.TPoint                              // property LineCharPos Getter
	SetLineCharPos(value types.TPoint)                      // property LineCharPos Setter
	BytePos() int32                                         // property BytePos Getter
	SetBytePos(value int32)                                 // property BytePos Setter
	BytePosOffset() int32                                   // property BytePosOffset Getter
	SetBytePosOffset(value int32)                           // property BytePosOffset Setter
	LineBytePos() types.TPoint                              // property LineBytePos Getter
	SetLineBytePos(value types.TPoint)                      // property LineBytePos Setter
	FullLogicalPos() TLogCaretPoint                         // property FullLogicalPos Getter
	SetFullLogicalPos(value TLogCaretPoint)                 // property FullLogicalPos Setter
	ViewedLineCharPos() types.TPoint                        // property ViewedLineCharPos Getter
	SetViewedLineCharPos(value types.TPoint)                // property ViewedLineCharPos Setter
	ViewedLinePos() types.TLinePos                          // property ViewedLinePos Getter
	SetViewedLinePos(value types.TLinePos)                  // property ViewedLinePos Setter
	LineText() string                                       // property LineText Getter
	SetLineText(value string)                               // property LineText Setter
}

type TSynEditBaseCaret struct {
	TSynEditPointBase
}

func (m *TSynEditBaseCaret) IsAtLineChar(point types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseCaretAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&point)))
	return api.GoBool(r)
}

func (m *TSynEditBaseCaret) IsAtLineByte(point types.TPoint, byteOffset int32) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseCaretAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&point)), uintptr(byteOffset))
	return api.GoBool(r)
}

func (m *TSynEditBaseCaret) IsAtPos(caret ISynEditCaret) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseCaretAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(caret))
	return api.GoBool(r)
}

func (m *TSynEditBaseCaret) AssignFromWithSynEditBaseCaret(src ISynEditBaseCaret) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(src))
}

func (m *TSynEditBaseCaret) Invalidate() {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(5, m.Instance())
}

func (m *TSynEditBaseCaret) InvalidateBytePos() {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(6, m.Instance())
}

func (m *TSynEditBaseCaret) InvalidateCharPos() {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(7, m.Instance())
}

func (m *TSynEditBaseCaret) LinePos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseCaretAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBaseCaret) SetLinePos(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBaseCaret) CharPos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseCaretAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBaseCaret) SetCharPos(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBaseCaret) LineCharPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBaseCaret) SetLineCharPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBaseCaret) BytePos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseCaretAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBaseCaret) SetBytePos(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBaseCaret) BytePosOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseCaretAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBaseCaret) SetBytePosOffset(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBaseCaret) LineBytePos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(13, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBaseCaret) SetLineBytePos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(13, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBaseCaret) FullLogicalPos() (result TLogCaretPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(14, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBaseCaret) SetFullLogicalPos(value TLogCaretPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(14, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBaseCaret) ViewedLineCharPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(15, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBaseCaret) SetViewedLineCharPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(15, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBaseCaret) ViewedLinePos() types.TLinePos {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseCaretAPI().SysCallN(16, 0, m.Instance())
	return types.TLinePos(r)
}

func (m *TSynEditBaseCaret) SetViewedLinePos(value types.TLinePos) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBaseCaret) LineText() string {
	if !m.IsValid() {
		return ""
	}
	r := synEditBaseCaretAPI().SysCallN(17, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TSynEditBaseCaret) SetLineText(value string) {
	if !m.IsValid() {
		return
	}
	synEditBaseCaretAPI().SysCallN(17, 1, m.Instance(), api.PasStr(value))
}

// NewSynEditBaseCaret class constructor
func NewSynEditBaseCaret() ISynEditBaseCaret {
	r := synEditBaseCaretAPI().SysCallN(0)
	return AsSynEditBaseCaret(r)
}

var (
	synEditBaseCaretOnce   base.Once
	synEditBaseCaretImport *imports.Imports = nil
)

func synEditBaseCaretAPI() *imports.Imports {
	synEditBaseCaretOnce.Do(func() {
		synEditBaseCaretImport = api.NewDefaultImports()
		synEditBaseCaretImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditBaseCaret_Create", 0), // constructor NewSynEditBaseCaret
			/* 1 */ imports.NewTable("TSynEditBaseCaret_IsAtLineChar", 0), // function IsAtLineChar
			/* 2 */ imports.NewTable("TSynEditBaseCaret_IsAtLineByte", 0), // function IsAtLineByte
			/* 3 */ imports.NewTable("TSynEditBaseCaret_IsAtPos", 0), // function IsAtPos
			/* 4 */ imports.NewTable("TSynEditBaseCaret_AssignFromWithSynEditBaseCaret", 0), // procedure AssignFromWithSynEditBaseCaret
			/* 5 */ imports.NewTable("TSynEditBaseCaret_Invalidate", 0), // procedure Invalidate
			/* 6 */ imports.NewTable("TSynEditBaseCaret_InvalidateBytePos", 0), // procedure InvalidateBytePos
			/* 7 */ imports.NewTable("TSynEditBaseCaret_InvalidateCharPos", 0), // procedure InvalidateCharPos
			/* 8 */ imports.NewTable("TSynEditBaseCaret_LinePos", 0), // property LinePos
			/* 9 */ imports.NewTable("TSynEditBaseCaret_CharPos", 0), // property CharPos
			/* 10 */ imports.NewTable("TSynEditBaseCaret_LineCharPos", 0), // property LineCharPos
			/* 11 */ imports.NewTable("TSynEditBaseCaret_BytePos", 0), // property BytePos
			/* 12 */ imports.NewTable("TSynEditBaseCaret_BytePosOffset", 0), // property BytePosOffset
			/* 13 */ imports.NewTable("TSynEditBaseCaret_LineBytePos", 0), // property LineBytePos
			/* 14 */ imports.NewTable("TSynEditBaseCaret_FullLogicalPos", 0), // property FullLogicalPos
			/* 15 */ imports.NewTable("TSynEditBaseCaret_ViewedLineCharPos", 0), // property ViewedLineCharPos
			/* 16 */ imports.NewTable("TSynEditBaseCaret_ViewedLinePos", 0), // property ViewedLinePos
			/* 17 */ imports.NewTable("TSynEditBaseCaret_LineText", 0), // property LineText
		}
	})
	return synEditBaseCaretImport
}
