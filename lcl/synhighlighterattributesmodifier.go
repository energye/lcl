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
)

// ISynHighlighterAttributesModifier Parent: ISynHighlighterAttributes
type ISynHighlighterAttributesModifier interface {
	ISynHighlighterAttributes
	// BackAlpha
	//  Alpha = 0 means solid // 1..255 means n of 256
	BackAlpha() byte          // property BackAlpha Getter
	SetBackAlpha(value byte)  // property BackAlpha Setter
	ForeAlpha() byte          // property ForeAlpha Getter
	SetForeAlpha(value byte)  // property ForeAlpha Setter
	FrameAlpha() byte         // property FrameAlpha Getter
	SetFrameAlpha(value byte) // property FrameAlpha Setter
}

type TSynHighlighterAttributesModifier struct {
	TSynHighlighterAttributes
}

func (m *TSynHighlighterAttributesModifier) BackAlpha() byte {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterAttributesModifierAPI().SysCallN(2, 0, m.Instance())
	return byte(r)
}

func (m *TSynHighlighterAttributesModifier) SetBackAlpha(value byte) {
	if !m.IsValid() {
		return
	}
	synHighlighterAttributesModifierAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TSynHighlighterAttributesModifier) ForeAlpha() byte {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterAttributesModifierAPI().SysCallN(3, 0, m.Instance())
	return byte(r)
}

func (m *TSynHighlighterAttributesModifier) SetForeAlpha(value byte) {
	if !m.IsValid() {
		return
	}
	synHighlighterAttributesModifierAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSynHighlighterAttributesModifier) FrameAlpha() byte {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterAttributesModifierAPI().SysCallN(4, 0, m.Instance())
	return byte(r)
}

func (m *TSynHighlighterAttributesModifier) SetFrameAlpha(value byte) {
	if !m.IsValid() {
		return
	}
	synHighlighterAttributesModifierAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

// NewSynHighlighterAttributesModifier class constructor
func NewSynHighlighterAttributesModifier() ISynHighlighterAttributesModifier {
	r := synHighlighterAttributesModifierAPI().SysCallN(0)
	return AsSynHighlighterAttributesModifier(r)
}

// NewSynHighlighterAttributesModifierWithStrX2 class constructor
func NewSynHighlighterAttributesModifierWithStrX2(caption string, storedName string) ISynHighlighterAttributesModifier {
	r := synHighlighterAttributesModifierAPI().SysCallN(1, api.PasStr(caption), api.PasStr(storedName))
	return AsSynHighlighterAttributesModifier(r)
}

var (
	synHighlighterAttributesModifierOnce   base.Once
	synHighlighterAttributesModifierImport *imports.Imports = nil
)

func synHighlighterAttributesModifierAPI() *imports.Imports {
	synHighlighterAttributesModifierOnce.Do(func() {
		synHighlighterAttributesModifierImport = api.NewDefaultImports()
		synHighlighterAttributesModifierImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHighlighterAttributesModifier_Create", 0), // constructor NewSynHighlighterAttributesModifier
			/* 1 */ imports.NewTable("TSynHighlighterAttributesModifier_CreateWithStrX2", 0), // constructor NewSynHighlighterAttributesModifierWithStrX2
			/* 2 */ imports.NewTable("TSynHighlighterAttributesModifier_BackAlpha", 0), // property BackAlpha
			/* 3 */ imports.NewTable("TSynHighlighterAttributesModifier_ForeAlpha", 0), // property ForeAlpha
			/* 4 */ imports.NewTable("TSynHighlighterAttributesModifier_FrameAlpha", 0), // property FrameAlpha
		}
	})
	return synHighlighterAttributesModifierImport
}
