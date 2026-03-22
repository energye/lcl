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

// ISynPythonSyn Parent: ISynCustomHighlighter
type ISynPythonSyn interface {
	ISynCustomHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHPythonTtkTokenKind                    // function
	GetKeywordIdentifiers() IStringList                        // function
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	IdentifierAttri() ISynHighlighterAttributes                // property IdentifierAttri Getter
	SetIdentifierAttri(value ISynHighlighterAttributes)        // property IdentifierAttri Setter
	KeyAttri() ISynHighlighterAttributes                       // property KeyAttri Getter
	SetKeyAttri(value ISynHighlighterAttributes)               // property KeyAttri Setter
	NonKeyAttri() ISynHighlighterAttributes                    // property NonKeyAttri Getter
	SetNonKeyAttri(value ISynHighlighterAttributes)            // property NonKeyAttri Setter
	SystemAttri() ISynHighlighterAttributes                    // property SystemAttri Getter
	SetSystemAttri(value ISynHighlighterAttributes)            // property SystemAttri Setter
	NumberAttri() ISynHighlighterAttributes                    // property NumberAttri Getter
	SetNumberAttri(value ISynHighlighterAttributes)            // property NumberAttri Setter
	HexAttri() ISynHighlighterAttributes                       // property HexAttri Getter
	SetHexAttri(value ISynHighlighterAttributes)               // property HexAttri Setter
	OctalAttri() ISynHighlighterAttributes                     // property OctalAttri Getter
	SetOctalAttri(value ISynHighlighterAttributes)             // property OctalAttri Setter
	FloatAttri() ISynHighlighterAttributes                     // property FloatAttri Getter
	SetFloatAttri(value ISynHighlighterAttributes)             // property FloatAttri Setter
	SpaceAttri() ISynHighlighterAttributes                     // property SpaceAttri Getter
	SetSpaceAttri(value ISynHighlighterAttributes)             // property SpaceAttri Setter
	StringAttri() ISynHighlighterAttributes                    // property StringAttri Getter
	SetStringAttri(value ISynHighlighterAttributes)            // property StringAttri Setter
	DocStringAttri() ISynHighlighterAttributes                 // property DocStringAttri Getter
	SetDocStringAttri(value ISynHighlighterAttributes)         // property DocStringAttri Setter
	SymbolAttri() ISynHighlighterAttributes                    // property SymbolAttri Getter
	SetSymbolAttri(value ISynHighlighterAttributes)            // property SymbolAttri Setter
	ErrorAttri() ISynHighlighterAttributes                     // property ErrorAttri Getter
	SetErrorAttri(value ISynHighlighterAttributes)             // property ErrorAttri Setter
}

type TSynPythonSyn struct {
	TSynCustomHighlighter
}

func (m *TSynPythonSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) GetTokenID() types.SHPythonTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synPythonSynAPI().SysCallN(2, m.Instance())
	return types.SHPythonTtkTokenKind(r)
}

func (m *TSynPythonSyn) GetKeywordIdentifiers() IStringList {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(3, m.Instance())
	return AsStringList(r)
}

func (m *TSynPythonSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) NonKeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetNonKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) SystemAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetSystemAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) NumberAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetNumberAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) HexAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetHexAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) OctalAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetOctalAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) FloatAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetFloatAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(13, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) StringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(14, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) DocStringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(15, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetDocStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(16, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPythonSyn) ErrorAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synPythonSynAPI().SysCallN(17, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynPythonSyn) SetErrorAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synPythonSynAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynPythonSyn class constructor
func NewSynPythonSyn(owner IComponent) ISynPythonSyn {
	r := synPythonSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynPythonSyn(r)
}

func TSynPythonSynClass() types.TClass {
	r := synPythonSynAPI().SysCallN(18)
	return types.TClass(r)
}

var (
	synPythonSynOnce   base.Once
	synPythonSynImport *imports.Imports = nil
)

func synPythonSynAPI() *imports.Imports {
	synPythonSynOnce.Do(func() {
		synPythonSynImport = api.NewDefaultImports()
		synPythonSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynPythonSyn_Create", 0), // constructor NewSynPythonSyn
			/* 1 */ imports.NewTable("TSynPythonSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynPythonSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynPythonSyn_GetKeywordIdentifiers", 0), // function GetKeywordIdentifiers
			/* 4 */ imports.NewTable("TSynPythonSyn_CommentAttri", 0), // property CommentAttri
			/* 5 */ imports.NewTable("TSynPythonSyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 6 */ imports.NewTable("TSynPythonSyn_KeyAttri", 0), // property KeyAttri
			/* 7 */ imports.NewTable("TSynPythonSyn_NonKeyAttri", 0), // property NonKeyAttri
			/* 8 */ imports.NewTable("TSynPythonSyn_SystemAttri", 0), // property SystemAttri
			/* 9 */ imports.NewTable("TSynPythonSyn_NumberAttri", 0), // property NumberAttri
			/* 10 */ imports.NewTable("TSynPythonSyn_HexAttri", 0), // property HexAttri
			/* 11 */ imports.NewTable("TSynPythonSyn_OctalAttri", 0), // property OctalAttri
			/* 12 */ imports.NewTable("TSynPythonSyn_FloatAttri", 0), // property FloatAttri
			/* 13 */ imports.NewTable("TSynPythonSyn_SpaceAttri", 0), // property SpaceAttri
			/* 14 */ imports.NewTable("TSynPythonSyn_StringAttri", 0), // property StringAttri
			/* 15 */ imports.NewTable("TSynPythonSyn_DocStringAttri", 0), // property DocStringAttri
			/* 16 */ imports.NewTable("TSynPythonSyn_SymbolAttri", 0), // property SymbolAttri
			/* 17 */ imports.NewTable("TSynPythonSyn_ErrorAttri", 0), // property ErrorAttri
			/* 18 */ imports.NewTable("TSynPythonSyn_TClass", 0), // function TSynPythonSynClass
		}
	})
	return synPythonSynImport
}
