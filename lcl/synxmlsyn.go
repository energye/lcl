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

// ISynXMLSyn Parent: ISynCustomXmlHighlighter
type ISynXMLSyn interface {
	ISynCustomXmlHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes       // function
	GetTokenID() types.SHXMLTtkTokenKind                             // function
	ElementAttri() ISynHighlighterAttributes                         // property ElementAttri Getter
	SetElementAttri(value ISynHighlighterAttributes)                 // property ElementAttri Setter
	AttributeAttri() ISynHighlighterAttributes                       // property AttributeAttri Getter
	SetAttributeAttri(value ISynHighlighterAttributes)               // property AttributeAttri Setter
	NamespaceAttributeAttri() ISynHighlighterAttributes              // property NamespaceAttributeAttri Getter
	SetNamespaceAttributeAttri(value ISynHighlighterAttributes)      // property NamespaceAttributeAttri Setter
	AttributeValueAttri() ISynHighlighterAttributes                  // property AttributeValueAttri Getter
	SetAttributeValueAttri(value ISynHighlighterAttributes)          // property AttributeValueAttri Setter
	NamespaceAttributeValueAttri() ISynHighlighterAttributes         // property NamespaceAttributeValueAttri Getter
	SetNamespaceAttributeValueAttri(value ISynHighlighterAttributes) // property NamespaceAttributeValueAttri Setter
	TextAttri() ISynHighlighterAttributes                            // property TextAttri Getter
	SetTextAttri(value ISynHighlighterAttributes)                    // property TextAttri Setter
	CDATAAttri() ISynHighlighterAttributes                           // property CDATAAttri Getter
	SetCDATAAttri(value ISynHighlighterAttributes)                   // property CDATAAttri Setter
	EntityRefAttri() ISynHighlighterAttributes                       // property EntityRefAttri Getter
	SetEntityRefAttri(value ISynHighlighterAttributes)               // property EntityRefAttri Setter
	ProcessingInstructionAttri() ISynHighlighterAttributes           // property ProcessingInstructionAttri Getter
	SetProcessingInstructionAttri(value ISynHighlighterAttributes)   // property ProcessingInstructionAttri Setter
	CommentAttri() ISynHighlighterAttributes                         // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)                 // property CommentAttri Setter
	DocTypeAttri() ISynHighlighterAttributes                         // property DocTypeAttri Getter
	SetDocTypeAttri(value ISynHighlighterAttributes)                 // property DocTypeAttri Setter
	SpaceAttri() ISynHighlighterAttributes                           // property SpaceAttri Getter
	SetSpaceAttri(value ISynHighlighterAttributes)                   // property SpaceAttri Setter
	SymbolAttri() ISynHighlighterAttributes                          // property SymbolAttri Getter
	SetSymbolAttri(value ISynHighlighterAttributes)                  // property SymbolAttri Setter
	WantBracesParsed() bool                                          // property WantBracesParsed Getter
	SetWantBracesParsed(value bool)                                  // property WantBracesParsed Setter
}

type TSynXMLSyn struct {
	TSynCustomXmlHighlighter
}

func (m *TSynXMLSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) GetTokenID() types.SHXMLTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synXMLSynAPI().SysCallN(2, m.Instance())
	return types.SHXMLTtkTokenKind(r)
}

func (m *TSynXMLSyn) ElementAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(3, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetElementAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) AttributeAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetAttributeAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) NamespaceAttributeAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetNamespaceAttributeAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) AttributeValueAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetAttributeValueAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) NamespaceAttributeValueAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetNamespaceAttributeValueAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) TextAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetTextAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) CDATAAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetCDATAAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) EntityRefAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetEntityRefAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) ProcessingInstructionAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetProcessingInstructionAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) DocTypeAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(13, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetDocTypeAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(14, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synXMLSynAPI().SysCallN(15, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynXMLSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynXMLSyn) WantBracesParsed() bool {
	if !m.IsValid() {
		return false
	}
	r := synXMLSynAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynXMLSyn) SetWantBracesParsed(value bool) {
	if !m.IsValid() {
		return
	}
	synXMLSynAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

// NewSynXMLSyn class constructor
func NewSynXMLSyn(owner IComponent) ISynXMLSyn {
	r := synXMLSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynXMLSyn(r)
}

func TSynXMLSynClass() types.TClass {
	r := synXMLSynAPI().SysCallN(17)
	return types.TClass(r)
}

var (
	synXMLSynOnce   base.Once
	synXMLSynImport *imports.Imports = nil
)

func synXMLSynAPI() *imports.Imports {
	synXMLSynOnce.Do(func() {
		synXMLSynImport = api.NewDefaultImports()
		synXMLSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynXMLSyn_Create", 0), // constructor NewSynXMLSyn
			/* 1 */ imports.NewTable("TSynXMLSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynXMLSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynXMLSyn_ElementAttri", 0), // property ElementAttri
			/* 4 */ imports.NewTable("TSynXMLSyn_AttributeAttri", 0), // property AttributeAttri
			/* 5 */ imports.NewTable("TSynXMLSyn_NamespaceAttributeAttri", 0), // property NamespaceAttributeAttri
			/* 6 */ imports.NewTable("TSynXMLSyn_AttributeValueAttri", 0), // property AttributeValueAttri
			/* 7 */ imports.NewTable("TSynXMLSyn_NamespaceAttributeValueAttri", 0), // property NamespaceAttributeValueAttri
			/* 8 */ imports.NewTable("TSynXMLSyn_TextAttri", 0), // property TextAttri
			/* 9 */ imports.NewTable("TSynXMLSyn_CDATAAttri", 0), // property CDATAAttri
			/* 10 */ imports.NewTable("TSynXMLSyn_EntityRefAttri", 0), // property EntityRefAttri
			/* 11 */ imports.NewTable("TSynXMLSyn_ProcessingInstructionAttri", 0), // property ProcessingInstructionAttri
			/* 12 */ imports.NewTable("TSynXMLSyn_CommentAttri", 0), // property CommentAttri
			/* 13 */ imports.NewTable("TSynXMLSyn_DocTypeAttri", 0), // property DocTypeAttri
			/* 14 */ imports.NewTable("TSynXMLSyn_SpaceAttri", 0), // property SpaceAttri
			/* 15 */ imports.NewTable("TSynXMLSyn_SymbolAttri", 0), // property SymbolAttri
			/* 16 */ imports.NewTable("TSynXMLSyn_WantBracesParsed", 0), // property WantBracesParsed
			/* 17 */ imports.NewTable("TSynXMLSyn_TClass", 0), // function TSynXMLSynClass
		}
	})
	return synXMLSynImport
}
