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

// ISynGutterBase Parent: IPersistent
type ISynGutterBase interface {
	IPersistent
	RecalcBounds()                         // procedure
	ScalePPI(scaleFactor float64)          // procedure
	DoAutoSize()                           // procedure
	ResetMouseActions()                    // procedure
	Left() int32                           // property Left Getter
	Top() int32                            // property Top Getter
	Height() int32                         // property Height Getter
	Width() int32                          // property Width Getter
	SetWidth(value int32)                  // property Width Setter
	Side() types.TSynGutterSide            // property Side Getter
	AutoSize() bool                        // property AutoSize Getter
	SetAutoSize(value bool)                // property AutoSize Setter
	Visible() bool                         // property Visible Getter
	SetVisible(value bool)                 // property Visible Setter
	LeftOffset() int32                     // property LeftOffset Getter
	SetLeftOffset(value int32)             // property LeftOffset Setter
	RightOffset() int32                    // property RightOffset Getter
	SetRightOffset(value int32)            // property RightOffset Setter
	Parts() ISynGutterPartListBase         // property Parts Getter
	SetParts(value ISynGutterPartListBase) // property Parts Setter
	// SynEdit
	//  properties available for the GutterPartClasses
	SynEdit() ISynEditBase                                       // property SynEdit Getter
	TextDrawer() IheTextDrawer                                   // property TextDrawer Getter
	Color() types.TColor                                         // property Color Getter
	SetColor(value types.TColor)                                 // property Color Setter
	CurrentLineColor() ISynHighlighterAttributesModifier         // property CurrentLineColor Getter
	SetCurrentLineColor(value ISynHighlighterAttributesModifier) // property CurrentLineColor Setter
	MouseActions() ISynEditMouseActions                          // property MouseActions Getter
	SetMouseActions(value ISynEditMouseActions)                  // property MouseActions Setter
	SetOnChange(fn TNotifyEvent)                                 // property event
	SetOnResize(fn TNotifyEvent)                                 // property event
}

type TSynGutterBase struct {
	TPersistent
}

func (m *TSynGutterBase) RecalcBounds() {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(0, m.Instance())
}

func (m *TSynGutterBase) ScalePPI(scaleFactor float64) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)))
}

func (m *TSynGutterBase) DoAutoSize() {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(2, m.Instance())
}

func (m *TSynGutterBase) ResetMouseActions() {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(3, m.Instance())
}

func (m *TSynGutterBase) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TSynGutterBase) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TSynGutterBase) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TSynGutterBase) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterBase) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterBase) Side() types.TSynGutterSide {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(8, m.Instance())
	return types.TSynGutterSide(r)
}

func (m *TSynGutterBase) AutoSize() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterBaseAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterBase) SetAutoSize(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynGutterBase) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterBaseAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterBase) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynGutterBase) LeftOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterBase) SetLeftOffset(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterBase) RightOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterBase) SetRightOffset(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterBase) Parts() ISynGutterPartListBase {
	if !m.IsValid() {
		return nil
	}
	r := synGutterBaseAPI().SysCallN(13, 0, m.Instance())
	return AsSynGutterPartListBase(r)
}

func (m *TSynGutterBase) SetParts(value ISynGutterPartListBase) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterBase) SynEdit() ISynEditBase {
	if !m.IsValid() {
		return nil
	}
	r := synGutterBaseAPI().SysCallN(14, m.Instance())
	return AsSynEditBase(r)
}

func (m *TSynGutterBase) TextDrawer() IheTextDrawer {
	if !m.IsValid() {
		return nil
	}
	r := synGutterBaseAPI().SysCallN(15, m.Instance())
	return AsHeTextDrawer(r)
}

func (m *TSynGutterBase) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synGutterBaseAPI().SysCallN(16, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynGutterBase) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterBase) CurrentLineColor() ISynHighlighterAttributesModifier {
	if !m.IsValid() {
		return nil
	}
	r := synGutterBaseAPI().SysCallN(17, 0, m.Instance())
	return AsSynHighlighterAttributesModifier(r)
}

func (m *TSynGutterBase) SetCurrentLineColor(value ISynHighlighterAttributesModifier) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterBase) MouseActions() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synGutterBaseAPI().SysCallN(18, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynGutterBase) SetMouseActions(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synGutterBaseAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterBase) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, synGutterBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynGutterBase) SetOnResize(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, synGutterBaseAPI(), api.MakeEventDataPtr(cb))
}

var (
	synGutterBaseOnce   base.Once
	synGutterBaseImport *imports.Imports = nil
)

func synGutterBaseAPI() *imports.Imports {
	synGutterBaseOnce.Do(func() {
		synGutterBaseImport = api.NewDefaultImports()
		synGutterBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterBase_RecalcBounds", 0), // procedure RecalcBounds
			/* 1 */ imports.NewTable("TSynGutterBase_ScalePPI", 0), // procedure ScalePPI
			/* 2 */ imports.NewTable("TSynGutterBase_DoAutoSize", 0), // procedure DoAutoSize
			/* 3 */ imports.NewTable("TSynGutterBase_ResetMouseActions", 0), // procedure ResetMouseActions
			/* 4 */ imports.NewTable("TSynGutterBase_Left", 0), // property Left
			/* 5 */ imports.NewTable("TSynGutterBase_Top", 0), // property Top
			/* 6 */ imports.NewTable("TSynGutterBase_Height", 0), // property Height
			/* 7 */ imports.NewTable("TSynGutterBase_Width", 0), // property Width
			/* 8 */ imports.NewTable("TSynGutterBase_Side", 0), // property Side
			/* 9 */ imports.NewTable("TSynGutterBase_AutoSize", 0), // property AutoSize
			/* 10 */ imports.NewTable("TSynGutterBase_Visible", 0), // property Visible
			/* 11 */ imports.NewTable("TSynGutterBase_LeftOffset", 0), // property LeftOffset
			/* 12 */ imports.NewTable("TSynGutterBase_RightOffset", 0), // property RightOffset
			/* 13 */ imports.NewTable("TSynGutterBase_Parts", 0), // property Parts
			/* 14 */ imports.NewTable("TSynGutterBase_SynEdit", 0), // property SynEdit
			/* 15 */ imports.NewTable("TSynGutterBase_TextDrawer", 0), // property TextDrawer
			/* 16 */ imports.NewTable("TSynGutterBase_Color", 0), // property Color
			/* 17 */ imports.NewTable("TSynGutterBase_CurrentLineColor", 0), // property CurrentLineColor
			/* 18 */ imports.NewTable("TSynGutterBase_MouseActions", 0), // property MouseActions
			/* 19 */ imports.NewTable("TSynGutterBase_OnChange", 0), // event OnChange
			/* 20 */ imports.NewTable("TSynGutterBase_OnResize", 0), // event OnResize
		}
	})
	return synGutterBaseImport
}
