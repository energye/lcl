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

// ISynEditMarkup Parent: IObject
type ISynEditMarkup interface {
	IObject
	GetMarkupAttributeAtRowCol(row int32, startCol TLazSynDisplayTokenBound, anRtlInfo TLazSynDisplayRtlInfo) ISynSelectedColor                                                       // function
	GetMarkupAttributeAtWrapEnd(row int32, wrapCol TLazSynDisplayTokenBound) ISynSelectedColor                                                                                        // function
	RealEnabled() bool                                                                                                                                                                // function
	PrepareMarkupForRow(row int32)                                                                                                                                                    // procedure
	FinishMarkupForRow(row int32)                                                                                                                                                     // procedure
	BeginMarkup()                                                                                                                                                                     // procedure
	EndMarkup()                                                                                                                                                                       // procedure
	GetNextMarkupColAfterRowCol(row int32, startCol TLazSynDisplayTokenBound, anRtlInfo TLazSynDisplayRtlInfo, outNextPhys *int32, outNextLog *int32)                                 // procedure
	MergeMarkupAttributeAtRowCol(row int32, startCol TLazSynDisplayTokenBound, endCol TLazSynDisplayTokenBound, anRtlInfo TLazSynDisplayRtlInfo, markup ISynSelectedColorMergeResult) // procedure
	// MergeMarkupAttributeAtWrapEnd
	//  experimental; // params may still change
	MergeMarkupAttributeAtWrapEnd(row int32, wrapCol TLazSynDisplayTokenBound, markup ISynSelectedColorMergeResult) // procedure
	// TextChanged
	//  experimental; // params may still change
	//  Notifications about Changes to the text
	TextChanged(firstCodeLine int32, lastCodeLine int32, countDiff int32) // procedure
	TempDisable()                                                         // procedure
	TempEnable()                                                          // procedure
	IncPaintLock()                                                        // procedure
	DecPaintLock()                                                        // procedure
	MarkupInfo() ISynSelectedColor                                        // property MarkupInfo Getter
	FGColor() types.TColor                                                // property FGColor Getter
	BGColor() types.TColor                                                // property BGColor Getter
	FrameColor() types.TColor                                             // property FrameColor Getter
	FrameStyle() types.TSynLineStyle                                      // property FrameStyle Getter
	Style() types.TFontStyles                                             // property Style Getter
	Enabled() bool                                                        // property Enabled Getter
	SetEnabled(value bool)                                                // property Enabled Setter
	Lines() ISynEditStringsLinked                                         // property Lines Getter
	SetLines(value ISynEditStringsLinked)                                 // property Lines Setter
	Caret() ISynEditCaret                                                 // property Caret Getter
	SetCaret(value ISynEditCaret)                                         // property Caret Setter
	TopLine() int32                                                       // property TopLine Getter
	SetTopLine(value int32)                                               // property TopLine Setter
	LinesInWindow() int32                                                 // property LinesInWindow Getter
	SetLinesInWindow(value int32)                                         // property LinesInWindow Setter
	SetInvalidateLinesMethod(fn TInvalidateLines)                         // property event
}

type TSynEditMarkup struct {
	TObject
}

func (m *TSynEditMarkup) GetMarkupAttributeAtRowCol(row int32, startCol TLazSynDisplayTokenBound, anRtlInfo TLazSynDisplayRtlInfo) ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkupAPI().SysCallN(0, m.Instance(), uintptr(row), uintptr(base.UnsafePointer(&startCol)), uintptr(base.UnsafePointer(&anRtlInfo)))
	return AsSynSelectedColor(r)
}

func (m *TSynEditMarkup) GetMarkupAttributeAtWrapEnd(row int32, wrapCol TLazSynDisplayTokenBound) ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkupAPI().SysCallN(1, m.Instance(), uintptr(row), uintptr(base.UnsafePointer(&wrapCol)))
	return AsSynSelectedColor(r)
}

func (m *TSynEditMarkup) RealEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMarkupAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMarkup) PrepareMarkupForRow(row int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(3, m.Instance(), uintptr(row))
}

func (m *TSynEditMarkup) FinishMarkupForRow(row int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(4, m.Instance(), uintptr(row))
}

func (m *TSynEditMarkup) BeginMarkup() {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(5, m.Instance())
}

func (m *TSynEditMarkup) EndMarkup() {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(6, m.Instance())
}

func (m *TSynEditMarkup) GetNextMarkupColAfterRowCol(row int32, startCol TLazSynDisplayTokenBound, anRtlInfo TLazSynDisplayRtlInfo, outNextPhys *int32, outNextLog *int32) {
	if !m.IsValid() {
		return
	}
	var nextPhysPtr uintptr
	var nextLogPtr uintptr
	synEditMarkupAPI().SysCallN(7, m.Instance(), uintptr(row), uintptr(base.UnsafePointer(&startCol)), uintptr(base.UnsafePointer(&anRtlInfo)), uintptr(base.UnsafePointer(&nextPhysPtr)), uintptr(base.UnsafePointer(&nextLogPtr)))
	*outNextPhys = int32(nextPhysPtr)
	*outNextLog = int32(nextLogPtr)
}

func (m *TSynEditMarkup) MergeMarkupAttributeAtRowCol(row int32, startCol TLazSynDisplayTokenBound, endCol TLazSynDisplayTokenBound, anRtlInfo TLazSynDisplayRtlInfo, markup ISynSelectedColorMergeResult) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(8, m.Instance(), uintptr(row), uintptr(base.UnsafePointer(&startCol)), uintptr(base.UnsafePointer(&endCol)), uintptr(base.UnsafePointer(&anRtlInfo)), base.GetObjectUintptr(markup))
}

func (m *TSynEditMarkup) MergeMarkupAttributeAtWrapEnd(row int32, wrapCol TLazSynDisplayTokenBound, markup ISynSelectedColorMergeResult) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(9, m.Instance(), uintptr(row), uintptr(base.UnsafePointer(&wrapCol)), base.GetObjectUintptr(markup))
}

func (m *TSynEditMarkup) TextChanged(firstCodeLine int32, lastCodeLine int32, countDiff int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(10, m.Instance(), uintptr(firstCodeLine), uintptr(lastCodeLine), uintptr(countDiff))
}

func (m *TSynEditMarkup) TempDisable() {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(11, m.Instance())
}

func (m *TSynEditMarkup) TempEnable() {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(12, m.Instance())
}

func (m *TSynEditMarkup) IncPaintLock() {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(13, m.Instance())
}

func (m *TSynEditMarkup) DecPaintLock() {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(14, m.Instance())
}

func (m *TSynEditMarkup) MarkupInfo() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkupAPI().SysCallN(15, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynEditMarkup) FGColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupAPI().SysCallN(16, m.Instance())
	return types.TColor(r)
}

func (m *TSynEditMarkup) BGColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupAPI().SysCallN(17, m.Instance())
	return types.TColor(r)
}

func (m *TSynEditMarkup) FrameColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupAPI().SysCallN(18, m.Instance())
	return types.TColor(r)
}

func (m *TSynEditMarkup) FrameStyle() types.TSynLineStyle {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupAPI().SysCallN(19, m.Instance())
	return types.TSynLineStyle(r)
}

func (m *TSynEditMarkup) Style() types.TFontStyles {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupAPI().SysCallN(20, m.Instance())
	return types.TFontStyles(r)
}

func (m *TSynEditMarkup) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMarkupAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMarkup) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditMarkup) Lines() ISynEditStringsLinked {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkupAPI().SysCallN(22, 0, m.Instance())
	return AsSynEditStringsLinked(r)
}

func (m *TSynEditMarkup) SetLines(value ISynEditStringsLinked) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditMarkup) Caret() ISynEditCaret {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkupAPI().SysCallN(23, 0, m.Instance())
	return AsSynEditCaret(r)
}

func (m *TSynEditMarkup) SetCaret(value ISynEditCaret) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(23, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditMarkup) TopLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMarkup) SetTopLine(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMarkup) LinesInWindow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupAPI().SysCallN(25, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMarkup) SetLinesInWindow(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkupAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMarkup) SetInvalidateLinesMethod(fn TInvalidateLines) {
	if !m.IsValid() {
		return
	}
	cb := makeTInvalidateLines(fn)
	base.SetEvent(m, 26, synEditMarkupAPI(), api.MakeEventDataPtr(cb))
}

var (
	synEditMarkupOnce   base.Once
	synEditMarkupImport *imports.Imports = nil
)

func synEditMarkupAPI() *imports.Imports {
	synEditMarkupOnce.Do(func() {
		synEditMarkupImport = api.NewDefaultImports()
		synEditMarkupImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMarkup_GetMarkupAttributeAtRowCol", 0), // function GetMarkupAttributeAtRowCol
			/* 1 */ imports.NewTable("TSynEditMarkup_GetMarkupAttributeAtWrapEnd", 0), // function GetMarkupAttributeAtWrapEnd
			/* 2 */ imports.NewTable("TSynEditMarkup_RealEnabled", 0), // function RealEnabled
			/* 3 */ imports.NewTable("TSynEditMarkup_PrepareMarkupForRow", 0), // procedure PrepareMarkupForRow
			/* 4 */ imports.NewTable("TSynEditMarkup_FinishMarkupForRow", 0), // procedure FinishMarkupForRow
			/* 5 */ imports.NewTable("TSynEditMarkup_BeginMarkup", 0), // procedure BeginMarkup
			/* 6 */ imports.NewTable("TSynEditMarkup_EndMarkup", 0), // procedure EndMarkup
			/* 7 */ imports.NewTable("TSynEditMarkup_GetNextMarkupColAfterRowCol", 0), // procedure GetNextMarkupColAfterRowCol
			/* 8 */ imports.NewTable("TSynEditMarkup_MergeMarkupAttributeAtRowCol", 0), // procedure MergeMarkupAttributeAtRowCol
			/* 9 */ imports.NewTable("TSynEditMarkup_MergeMarkupAttributeAtWrapEnd", 0), // procedure MergeMarkupAttributeAtWrapEnd
			/* 10 */ imports.NewTable("TSynEditMarkup_TextChanged", 0), // procedure TextChanged
			/* 11 */ imports.NewTable("TSynEditMarkup_TempDisable", 0), // procedure TempDisable
			/* 12 */ imports.NewTable("TSynEditMarkup_TempEnable", 0), // procedure TempEnable
			/* 13 */ imports.NewTable("TSynEditMarkup_IncPaintLock", 0), // procedure IncPaintLock
			/* 14 */ imports.NewTable("TSynEditMarkup_DecPaintLock", 0), // procedure DecPaintLock
			/* 15 */ imports.NewTable("TSynEditMarkup_MarkupInfo", 0), // property MarkupInfo
			/* 16 */ imports.NewTable("TSynEditMarkup_FGColor", 0), // property FGColor
			/* 17 */ imports.NewTable("TSynEditMarkup_BGColor", 0), // property BGColor
			/* 18 */ imports.NewTable("TSynEditMarkup_FrameColor", 0), // property FrameColor
			/* 19 */ imports.NewTable("TSynEditMarkup_FrameStyle", 0), // property FrameStyle
			/* 20 */ imports.NewTable("TSynEditMarkup_Style", 0), // property Style
			/* 21 */ imports.NewTable("TSynEditMarkup_Enabled", 0), // property Enabled
			/* 22 */ imports.NewTable("TSynEditMarkup_Lines", 0), // property Lines
			/* 23 */ imports.NewTable("TSynEditMarkup_Caret", 0), // property Caret
			/* 24 */ imports.NewTable("TSynEditMarkup_TopLine", 0), // property TopLine
			/* 25 */ imports.NewTable("TSynEditMarkup_LinesInWindow", 0), // property LinesInWindow
			/* 26 */ imports.NewTable("TSynEditMarkup_InvalidateLinesMethod", 0), // event InvalidateLinesMethod
		}
	})
	return synEditMarkupImport
}
