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

// ISynSelectedColorMergeResult Parent: ISynSelectedColor
type ISynSelectedColorMergeResult interface {
	ISynSelectedColor
	InitMergeInfo()                                                                                                                               // procedure
	ProcessMergeInfo()                                                                                                                            // procedure
	CleanupMergeInfo()                                                                                                                            // procedure
	MergeWithSynHighlighterAttributesModifier(other ISynHighlighterAttributesModifier)                                                            // procedure
	MergeWithSHAModifierLSDTBoundX2(other ISynHighlighterAttributesModifier, leftCol TLazSynDisplayTokenBound, rightCol TLazSynDisplayTokenBound) // procedure
	MergeFrames(other ISynHighlighterAttributesModifier, leftCol TLazSynDisplayTokenBound, rightCol TLazSynDisplayTokenBound)                     // procedure
	FrameSideColors(side types.TLazSynBorderSide) types.TColor                                                                                    // property FrameSideColors Getter
	FrameSideStyles(side types.TLazSynBorderSide) types.TSynLineStyle                                                                             // property FrameSideStyles Getter
	// CurrentStartX
	//  boundaries for current paint
	CurrentStartX() TLazSynDisplayTokenBound         // property CurrentStartX Getter
	SetCurrentStartX(value TLazSynDisplayTokenBound) // property CurrentStartX Setter
	CurrentEndX() TLazSynDisplayTokenBound           // property CurrentEndX Getter
	SetCurrentEndX(value TLazSynDisplayTokenBound)   // property CurrentEndX Setter
}

type TSynSelectedColorMergeResult struct {
	TSynSelectedColor
}

func (m *TSynSelectedColorMergeResult) InitMergeInfo() {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(2, m.Instance())
}

func (m *TSynSelectedColorMergeResult) ProcessMergeInfo() {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(3, m.Instance())
}

func (m *TSynSelectedColorMergeResult) CleanupMergeInfo() {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(4, m.Instance())
}

func (m *TSynSelectedColorMergeResult) MergeWithSynHighlighterAttributesModifier(other ISynHighlighterAttributesModifier) {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(other))
}

func (m *TSynSelectedColorMergeResult) MergeWithSHAModifierLSDTBoundX2(other ISynHighlighterAttributesModifier, leftCol TLazSynDisplayTokenBound, rightCol TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(other), uintptr(base.UnsafePointer(&leftCol)), uintptr(base.UnsafePointer(&rightCol)))
}

func (m *TSynSelectedColorMergeResult) MergeFrames(other ISynHighlighterAttributesModifier, leftCol TLazSynDisplayTokenBound, rightCol TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(other), uintptr(base.UnsafePointer(&leftCol)), uintptr(base.UnsafePointer(&rightCol)))
}

func (m *TSynSelectedColorMergeResult) FrameSideColors(side types.TLazSynBorderSide) types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synSelectedColorMergeResultAPI().SysCallN(8, m.Instance(), uintptr(side))
	return types.TColor(r)
}

func (m *TSynSelectedColorMergeResult) FrameSideStyles(side types.TLazSynBorderSide) types.TSynLineStyle {
	if !m.IsValid() {
		return 0
	}
	r := synSelectedColorMergeResultAPI().SysCallN(9, m.Instance(), uintptr(side))
	return types.TSynLineStyle(r)
}

func (m *TSynSelectedColorMergeResult) CurrentStartX() (result TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynSelectedColorMergeResult) SetCurrentStartX(value TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynSelectedColorMergeResult) CurrentEndX() (result TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynSelectedColorMergeResult) SetCurrentEndX(value TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorMergeResultAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

// NewSynSelectedColorMergeResult class constructor
func NewSynSelectedColorMergeResult() ISynSelectedColorMergeResult {
	r := synSelectedColorMergeResultAPI().SysCallN(0)
	return AsSynSelectedColorMergeResult(r)
}

// NewSynSelectedColorMergeResultWithStrX2 class constructor
func NewSynSelectedColorMergeResultWithStrX2(caption string, storedName string) ISynSelectedColorMergeResult {
	r := synSelectedColorMergeResultAPI().SysCallN(1, api.PasStr(caption), api.PasStr(storedName))
	return AsSynSelectedColorMergeResult(r)
}

var (
	synSelectedColorMergeResultOnce   base.Once
	synSelectedColorMergeResultImport *imports.Imports = nil
)

func synSelectedColorMergeResultAPI() *imports.Imports {
	synSelectedColorMergeResultOnce.Do(func() {
		synSelectedColorMergeResultImport = api.NewDefaultImports()
		synSelectedColorMergeResultImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynSelectedColorMergeResult_Create", 0), // constructor NewSynSelectedColorMergeResult
			/* 1 */ imports.NewTable("TSynSelectedColorMergeResult_CreateWithStrX2", 0), // constructor NewSynSelectedColorMergeResultWithStrX2
			/* 2 */ imports.NewTable("TSynSelectedColorMergeResult_InitMergeInfo", 0), // procedure InitMergeInfo
			/* 3 */ imports.NewTable("TSynSelectedColorMergeResult_ProcessMergeInfo", 0), // procedure ProcessMergeInfo
			/* 4 */ imports.NewTable("TSynSelectedColorMergeResult_CleanupMergeInfo", 0), // procedure CleanupMergeInfo
			/* 5 */ imports.NewTable("TSynSelectedColorMergeResult_MergeWithSynHighlighterAttributesModifier", 0), // procedure MergeWithSynHighlighterAttributesModifier
			/* 6 */ imports.NewTable("TSynSelectedColorMergeResult_MergeWithSHAModifierLSDTBoundX2", 0), // procedure MergeWithSHAModifierLSDTBoundX2
			/* 7 */ imports.NewTable("TSynSelectedColorMergeResult_MergeFrames", 0), // procedure MergeFrames
			/* 8 */ imports.NewTable("TSynSelectedColorMergeResult_FrameSideColors", 0), // property FrameSideColors
			/* 9 */ imports.NewTable("TSynSelectedColorMergeResult_FrameSideStyles", 0), // property FrameSideStyles
			/* 10 */ imports.NewTable("TSynSelectedColorMergeResult_CurrentStartX", 0), // property CurrentStartX
			/* 11 */ imports.NewTable("TSynSelectedColorMergeResult_CurrentEndX", 0), // property CurrentEndX
		}
	})
	return synSelectedColorMergeResultImport
}
