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

// ISynCssSyn Parent: ISynCustomHighlighter
type ISynCssSyn interface {
	ISynCustomHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHCssTtkTokenKind                       // function
	KeyHash2(toHash uintptr) int32                             // function
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	IdentifierAttri() ISynHighlighterAttributes                // property IdentifierAttri Getter
	SetIdentifierAttri(value ISynHighlighterAttributes)        // property IdentifierAttri Setter
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
	MeasurementUnitAttri() ISynHighlighterAttributes           // property MeasurementUnitAttri Getter
	SetMeasurementUnitAttri(value ISynHighlighterAttributes)   // property MeasurementUnitAttri Setter
	SelectorAttri() ISynHighlighterAttributes                  // property SelectorAttri Getter
	SetSelectorAttri(value ISynHighlighterAttributes)          // property SelectorAttri Setter
}

type TSynCssSyn struct {
	TSynCustomHighlighter
}

func (m *TSynCssSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) GetTokenID() types.SHCssTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synCssSynAPI().SysCallN(2, m.Instance())
	return types.SHCssTtkTokenKind(r)
}

func (m *TSynCssSyn) KeyHash2(toHash uintptr) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCssSynAPI().SysCallN(3, m.Instance(), uintptr(toHash))
	return int32(r)
}

func (m *TSynCssSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) NumberAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetNumberAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) StringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) MeasurementUnitAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetMeasurementUnitAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynCssSyn) SelectorAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synCssSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynCssSyn) SetSelectorAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synCssSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynCssSyn class constructor
func NewSynCssSyn(owner IComponent) ISynCssSyn {
	r := synCssSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynCssSyn(r)
}

func TSynCssSynClass() types.TClass {
	r := synCssSynAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	synCssSynOnce   base.Once
	synCssSynImport *imports.Imports = nil
)

func synCssSynAPI() *imports.Imports {
	synCssSynOnce.Do(func() {
		synCssSynImport = api.NewDefaultImports()
		synCssSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCssSyn_Create", 0), // constructor NewSynCssSyn
			/* 1 */ imports.NewTable("TSynCssSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynCssSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynCssSyn_KeyHash2", 0), // function KeyHash2
			/* 4 */ imports.NewTable("TSynCssSyn_CommentAttri", 0), // property CommentAttri
			/* 5 */ imports.NewTable("TSynCssSyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 6 */ imports.NewTable("TSynCssSyn_KeyAttri", 0), // property KeyAttri
			/* 7 */ imports.NewTable("TSynCssSyn_NumberAttri", 0), // property NumberAttri
			/* 8 */ imports.NewTable("TSynCssSyn_SpaceAttri", 0), // property SpaceAttri
			/* 9 */ imports.NewTable("TSynCssSyn_StringAttri", 0), // property StringAttri
			/* 10 */ imports.NewTable("TSynCssSyn_SymbolAttri", 0), // property SymbolAttri
			/* 11 */ imports.NewTable("TSynCssSyn_MeasurementUnitAttri", 0), // property MeasurementUnitAttri
			/* 12 */ imports.NewTable("TSynCssSyn_SelectorAttri", 0), // property SelectorAttri
			/* 13 */ imports.NewTable("TSynCssSyn_TClass", 0), // function TSynCssSynClass
		}
	})
	return synCssSynImport
}
