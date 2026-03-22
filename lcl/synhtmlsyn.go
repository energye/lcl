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

// ISynHTMLSyn Parent: ISynCustomXmlHighlighter
type ISynHTMLSyn interface {
	ISynCustomXmlHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHHTMLTtkTokenKind                      // function
	AndAttri() ISynHighlighterAttributes                       // property AndAttri Getter
	SetAndAttri(value ISynHighlighterAttributes)               // property AndAttri Setter
	ASPAttri() ISynHighlighterAttributes                       // property ASPAttri Getter
	SetASPAttri(value ISynHighlighterAttributes)               // property ASPAttri Setter
	CDATAAttri() ISynHighlighterAttributes                     // property CDATAAttri Getter
	SetCDATAAttri(value ISynHighlighterAttributes)             // property CDATAAttri Setter
	DOCTYPEAttri() ISynHighlighterAttributes                   // property DOCTYPEAttri Getter
	SetDOCTYPEAttri(value ISynHighlighterAttributes)           // property DOCTYPEAttri Setter
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	IdentifierAttri() ISynHighlighterAttributes                // property IdentifierAttri Getter
	SetIdentifierAttri(value ISynHighlighterAttributes)        // property IdentifierAttri Setter
	KeyAttri() ISynHighlighterAttributes                       // property KeyAttri Getter
	SetKeyAttri(value ISynHighlighterAttributes)               // property KeyAttri Setter
	SpaceAttri() ISynHighlighterAttributes                     // property SpaceAttri Getter
	SetSpaceAttri(value ISynHighlighterAttributes)             // property SpaceAttri Setter
	SymbolAttri() ISynHighlighterAttributes                    // property SymbolAttri Getter
	SetSymbolAttri(value ISynHighlighterAttributes)            // property SymbolAttri Setter
	TextAttri() ISynHighlighterAttributes                      // property TextAttri Getter
	SetTextAttri(value ISynHighlighterAttributes)              // property TextAttri Setter
	UndefKeyAttri() ISynHighlighterAttributes                  // property UndefKeyAttri Getter
	SetUndefKeyAttri(value ISynHighlighterAttributes)          // property UndefKeyAttri Setter
	ValueAttri() ISynHighlighterAttributes                     // property ValueAttri Getter
	SetValueAttri(value ISynHighlighterAttributes)             // property ValueAttri Setter
	Mode() types.TSynHTMLSynMode                               // property Mode Getter
	SetMode(value types.TSynHTMLSynMode)                       // property Mode Setter
}

type TSynHTMLSyn struct {
	TSynCustomXmlHighlighter
}

func (m *TSynHTMLSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) GetTokenID() types.SHHTMLTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synHTMLSynAPI().SysCallN(2, m.Instance())
	return types.SHHTMLTtkTokenKind(r)
}

func (m *TSynHTMLSyn) AndAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(3, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetAndAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) ASPAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetASPAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) CDATAAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetCDATAAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) DOCTYPEAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetDOCTYPEAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) TextAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetTextAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) UndefKeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(13, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetUndefKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) ValueAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHTMLSynAPI().SysCallN(14, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHTMLSyn) SetValueAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHTMLSyn) Mode() types.TSynHTMLSynMode {
	if !m.IsValid() {
		return 0
	}
	r := synHTMLSynAPI().SysCallN(15, 0, m.Instance())
	return types.TSynHTMLSynMode(r)
}

func (m *TSynHTMLSyn) SetMode(value types.TSynHTMLSynMode) {
	if !m.IsValid() {
		return
	}
	synHTMLSynAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

// NewSynHTMLSyn class constructor
func NewSynHTMLSyn(owner IComponent) ISynHTMLSyn {
	r := synHTMLSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynHTMLSyn(r)
}

func TSynHTMLSynClass() types.TClass {
	r := synHTMLSynAPI().SysCallN(16)
	return types.TClass(r)
}

var (
	synHTMLSynOnce   base.Once
	synHTMLSynImport *imports.Imports = nil
)

func synHTMLSynAPI() *imports.Imports {
	synHTMLSynOnce.Do(func() {
		synHTMLSynImport = api.NewDefaultImports()
		synHTMLSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHTMLSyn_Create", 0), // constructor NewSynHTMLSyn
			/* 1 */ imports.NewTable("TSynHTMLSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynHTMLSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynHTMLSyn_AndAttri", 0), // property AndAttri
			/* 4 */ imports.NewTable("TSynHTMLSyn_ASPAttri", 0), // property ASPAttri
			/* 5 */ imports.NewTable("TSynHTMLSyn_CDATAAttri", 0), // property CDATAAttri
			/* 6 */ imports.NewTable("TSynHTMLSyn_DOCTYPEAttri", 0), // property DOCTYPEAttri
			/* 7 */ imports.NewTable("TSynHTMLSyn_CommentAttri", 0), // property CommentAttri
			/* 8 */ imports.NewTable("TSynHTMLSyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 9 */ imports.NewTable("TSynHTMLSyn_KeyAttri", 0), // property KeyAttri
			/* 10 */ imports.NewTable("TSynHTMLSyn_SpaceAttri", 0), // property SpaceAttri
			/* 11 */ imports.NewTable("TSynHTMLSyn_SymbolAttri", 0), // property SymbolAttri
			/* 12 */ imports.NewTable("TSynHTMLSyn_TextAttri", 0), // property TextAttri
			/* 13 */ imports.NewTable("TSynHTMLSyn_UndefKeyAttri", 0), // property UndefKeyAttri
			/* 14 */ imports.NewTable("TSynHTMLSyn_ValueAttri", 0), // property ValueAttri
			/* 15 */ imports.NewTable("TSynHTMLSyn_Mode", 0), // property Mode
			/* 16 */ imports.NewTable("TSynHTMLSyn_TClass", 0), // function TSynHTMLSynClass
		}
	})
	return synHTMLSynImport
}
