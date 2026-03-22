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

// ISynJavaSyn Parent: ISynCustomHighlighter
type ISynJavaSyn interface {
	ISynCustomHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHJavaTtkTokenKind                      // function
	ExtTokenID() types.SHJavaTxtkTokenKind                     // property ExtTokenID Getter
	AnnotationAttri() ISynHighlighterAttributes                // property AnnotationAttri Getter
	SetAnnotationAttri(value ISynHighlighterAttributes)        // property AnnotationAttri Setter
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	DocumentAttri() ISynHighlighterAttributes                  // property DocumentAttri Getter
	SetDocumentAttri(value ISynHighlighterAttributes)          // property DocumentAttri Setter
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

type TSynJavaSyn struct {
	TSynCustomHighlighter
}

func (m *TSynJavaSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) GetTokenID() types.SHJavaTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synJavaSynAPI().SysCallN(2, m.Instance())
	return types.SHJavaTtkTokenKind(r)
}

func (m *TSynJavaSyn) ExtTokenID() types.SHJavaTxtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synJavaSynAPI().SysCallN(3, m.Instance())
	return types.SHJavaTxtkTokenKind(r)
}

func (m *TSynJavaSyn) AnnotationAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetAnnotationAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) DocumentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetDocumentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) InvalidAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetInvalidAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) NumberAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetNumberAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) StringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynJavaSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synJavaSynAPI().SysCallN(13, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynJavaSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synJavaSynAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynJavaSyn class constructor
func NewSynJavaSyn(owner IComponent) ISynJavaSyn {
	r := synJavaSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynJavaSyn(r)
}

func TSynJavaSynClass() types.TClass {
	r := synJavaSynAPI().SysCallN(14)
	return types.TClass(r)
}

var (
	synJavaSynOnce   base.Once
	synJavaSynImport *imports.Imports = nil
)

func synJavaSynAPI() *imports.Imports {
	synJavaSynOnce.Do(func() {
		synJavaSynImport = api.NewDefaultImports()
		synJavaSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynJavaSyn_Create", 0), // constructor NewSynJavaSyn
			/* 1 */ imports.NewTable("TSynJavaSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynJavaSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynJavaSyn_ExtTokenID", 0), // property ExtTokenID
			/* 4 */ imports.NewTable("TSynJavaSyn_AnnotationAttri", 0), // property AnnotationAttri
			/* 5 */ imports.NewTable("TSynJavaSyn_CommentAttri", 0), // property CommentAttri
			/* 6 */ imports.NewTable("TSynJavaSyn_DocumentAttri", 0), // property DocumentAttri
			/* 7 */ imports.NewTable("TSynJavaSyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 8 */ imports.NewTable("TSynJavaSyn_InvalidAttri", 0), // property InvalidAttri
			/* 9 */ imports.NewTable("TSynJavaSyn_KeyAttri", 0), // property KeyAttri
			/* 10 */ imports.NewTable("TSynJavaSyn_NumberAttri", 0), // property NumberAttri
			/* 11 */ imports.NewTable("TSynJavaSyn_SpaceAttri", 0), // property SpaceAttri
			/* 12 */ imports.NewTable("TSynJavaSyn_StringAttri", 0), // property StringAttri
			/* 13 */ imports.NewTable("TSynJavaSyn_SymbolAttri", 0), // property SymbolAttri
			/* 14 */ imports.NewTable("TSynJavaSyn_TClass", 0), // function TSynJavaSynClass
		}
	})
	return synJavaSynImport
}
