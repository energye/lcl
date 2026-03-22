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

// ISynJScriptSyn Parent: ISynCustomFoldHighlighter
type ISynJScriptSyn interface {
	ISynCustomFoldHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHJScriptTtkTokenKind                   // function
	BracketAttri() ISynHighlighterAttributes                   // property BracketAttri Getter
	SetBracketAttri(value ISynHighlighterAttributes)           // property BracketAttri Setter
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	IdentifierAttri() ISynHighlighterAttributes                // property IdentifierAttri Getter
	SetIdentifierAttri(value ISynHighlighterAttributes)        // property IdentifierAttri Setter
	KeyAttri() ISynHighlighterAttributes                       // property KeyAttri Getter
	SetKeyAttri(value ISynHighlighterAttributes)               // property KeyAttri Setter
	NonReservedKeyAttri() ISynHighlighterAttributes            // property NonReservedKeyAttri Getter
	SetNonReservedKeyAttri(value ISynHighlighterAttributes)    // property NonReservedKeyAttri Setter
	EventAttri() ISynHighlighterAttributes                     // property EventAttri Getter
	SetEventAttri(value ISynHighlighterAttributes)             // property EventAttri Setter
	NumberAttri() ISynHighlighterAttributes                    // property NumberAttri Getter
	SetNumberAttri(value ISynHighlighterAttributes)            // property NumberAttri Setter
	SpaceAttri() ISynHighlighterAttributes                     // property SpaceAttri Getter
	SetSpaceAttri(value ISynHighlighterAttributes)             // property SpaceAttri Setter
	StringAttri() ISynHighlighterAttributes                    // property StringAttri Getter
	SetStringAttri(value ISynHighlighterAttributes)            // property StringAttri Setter
	SymbolAttri() ISynHighlighterAttributes                    // property SymbolAttri Getter
	SetSymbolAttri(value ISynHighlighterAttributes)            // property SymbolAttri Setter
}

type TSynJScriptSyn struct {
	TSynCustomFoldHighlighter
}

func (m *TSynJScriptSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) GetTokenID() types.SHJScriptTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synJScriptSynAPI().SysCallN(2, m.Instance())
	return types.SHJScriptTtkTokenKind(r)
}

func (m *TSynJScriptSyn) BracketAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(3, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetBracketAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) NonReservedKeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetNonReservedKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) EventAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetEventAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) NumberAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetNumberAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) StringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJScriptSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJScriptSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJScriptSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJScriptSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynJScriptSyn class constructor
func NewSynJScriptSyn(owner IComponent) ISynJScriptSyn {
	r := synJScriptSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynJScriptSyn(r)
}

func TSynJScriptSynClass() types.TClass {
	r := synJScriptSynAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	synJScriptSynOnce   base.Once
	synJScriptSynImport *imports.Imports = nil
)

func synJScriptSynAPI() *imports.Imports {
	synJScriptSynOnce.Do(func() {
		synJScriptSynImport = api.NewDefaultImports()
		synJScriptSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynJScriptSyn_Create", 0), // constructor NewSynJScriptSyn
			/* 1 */ imports.NewTable("TSynJScriptSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynJScriptSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynJScriptSyn_BracketAttri", 0), // property BracketAttri
			/* 4 */ imports.NewTable("TSynJScriptSyn_CommentAttri", 0), // property CommentAttri
			/* 5 */ imports.NewTable("TSynJScriptSyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 6 */ imports.NewTable("TSynJScriptSyn_KeyAttri", 0), // property KeyAttri
			/* 7 */ imports.NewTable("TSynJScriptSyn_NonReservedKeyAttri", 0), // property NonReservedKeyAttri
			/* 8 */ imports.NewTable("TSynJScriptSyn_EventAttri", 0), // property EventAttri
			/* 9 */ imports.NewTable("TSynJScriptSyn_NumberAttri", 0), // property NumberAttri
			/* 10 */ imports.NewTable("TSynJScriptSyn_SpaceAttri", 0), // property SpaceAttri
			/* 11 */ imports.NewTable("TSynJScriptSyn_StringAttri", 0), // property StringAttri
			/* 12 */ imports.NewTable("TSynJScriptSyn_SymbolAttri", 0), // property SymbolAttri
			/* 13 */ imports.NewTable("TSynJScriptSyn_TClass", 0), // function TSynJScriptSynClass
		}
	})
	return synJScriptSynImport
}
