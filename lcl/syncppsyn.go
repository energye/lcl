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

// ISynCppSyn Parent: ISynCustomHighlighter
type ISynCppSyn interface {
	ISynCustomHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHCppTtkTokenKind                       // function
	ExtTokenID() types.SHCppTxtkTokenKind                      // property ExtTokenID Getter
	AsmAttri() ISynHighlighterAttributes                       // property AsmAttri Getter
	SetAsmAttri(value ISynHighlighterAttributes)               // property AsmAttri Setter
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	DirecAttri() ISynHighlighterAttributes                     // property DirecAttri Getter
	SetDirecAttri(value ISynHighlighterAttributes)             // property DirecAttri Setter
	IdentifierAttri() ISynHighlighterAttributes                // property IdentifierAttri Getter
	SetIdentifierAttri(value ISynHighlighterAttributes)        // property IdentifierAttri Setter
	InvalidAttri() ISynHighlighterAttributes                   // property InvalidAttri Getter
	SetInvalidAttri(value ISynHighlighterAttributes)           // property InvalidAttri Setter
	KeyAttri() ISynHighlighterAttributes                       // property KeyAttri Getter
	SetKeyAttri(value ISynHighlighterAttributes)               // property KeyAttri Setter
	NumberAttri() ISynHighlighterAttributes                    // property NumberAttri Getter
	SetNumberAttri(value ISynHighlighterAttributes)            // property NumberAttri Setter
	SpaceAttri() ISynHighlighterAttributes                     // property SpaceAttri Getter
	SetSpaceAttri(value ISynHighlighterAttributes)             // property SpaceAttri Setter
	StringAttri() ISynHighlighterAttributes                    // property StringAttri Getter
	SetStringAttri(value ISynHighlighterAttributes)            // property StringAttri Setter
	SymbolAttri() ISynHighlighterAttributes                    // property SymbolAttri Getter
	SetSymbolAttri(value ISynHighlighterAttributes)            // property SymbolAttri Setter
}

type TSynCppSyn struct {
	TSynCustomHighlighter
}

func (m *TSynCppSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) GetTokenID() types.SHCppTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synCppSynAPI().SysCallN(2, m.Instance())
	return types.SHCppTtkTokenKind(r)
}

func (m *TSynCppSyn) ExtTokenID() types.SHCppTxtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synCppSynAPI().SysCallN(3, m.Instance())
	return types.SHCppTxtkTokenKind(r)
}

func (m *TSynCppSyn) AsmAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetAsmAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) DirecAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetDirecAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) InvalidAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetInvalidAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) NumberAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetNumberAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) StringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCppSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCppSynAPI().SysCallN(13, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCppSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCppSynAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynCppSyn class constructor
func NewSynCppSyn(owner IComponent) ISynCppSyn {
	r := synCppSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynCppSyn(r)
}

func TSynCppSynClass() types.TClass {
	r := synCppSynAPI().SysCallN(14)
	return types.TClass(r)
}

var (
	synCppSynOnce   base.Once
	synCppSynImport *imports.Imports = nil
)

func synCppSynAPI() *imports.Imports {
	synCppSynOnce.Do(func() {
		synCppSynImport = api.NewDefaultImports()
		synCppSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCppSyn_Create", 0), // constructor NewSynCppSyn
			/* 1 */ imports.NewTable("TSynCppSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynCppSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynCppSyn_ExtTokenID", 0), // property ExtTokenID
			/* 4 */ imports.NewTable("TSynCppSyn_AsmAttri", 0), // property AsmAttri
			/* 5 */ imports.NewTable("TSynCppSyn_CommentAttri", 0), // property CommentAttri
			/* 6 */ imports.NewTable("TSynCppSyn_DirecAttri", 0), // property DirecAttri
			/* 7 */ imports.NewTable("TSynCppSyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 8 */ imports.NewTable("TSynCppSyn_InvalidAttri", 0), // property InvalidAttri
			/* 9 */ imports.NewTable("TSynCppSyn_KeyAttri", 0), // property KeyAttri
			/* 10 */ imports.NewTable("TSynCppSyn_NumberAttri", 0), // property NumberAttri
			/* 11 */ imports.NewTable("TSynCppSyn_SpaceAttri", 0), // property SpaceAttri
			/* 12 */ imports.NewTable("TSynCppSyn_StringAttri", 0), // property StringAttri
			/* 13 */ imports.NewTable("TSynCppSyn_SymbolAttri", 0), // property SymbolAttri
			/* 14 */ imports.NewTable("TSynCppSyn_TClass", 0), // function TSynCppSynClass
		}
	})
	return synCppSynImport
}
