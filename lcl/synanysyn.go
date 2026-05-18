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

// ISynAnySyn Parent: ISynCustomHighlighter
type ISynAnySyn interface {
	ISynCustomHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHAnyTtkTokenKind                       // function
	IsConstant(constant string) bool                           // function
	LoadHighLighter(file string)                               // procedure
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	Comments() types.CommentStyles                             // property Comments Getter
	SetComments(value types.CommentStyles)                     // property Comments Setter
	DetectPreprocessor() bool                                  // property DetectPreprocessor Getter
	SetDetectPreprocessor(value bool)                          // property DetectPreprocessor Setter
	IdentifierAttri() ISynHighlighterAttributes                // property IdentifierAttri Getter
	SetIdentifierAttri(value ISynHighlighterAttributes)        // property IdentifierAttri Setter
	IdentifierChars() string                                   // property IdentifierChars Getter
	SetIdentifierChars(value string)                           // property IdentifierChars Setter
	KeyAttri() ISynHighlighterAttributes                       // property KeyAttri Getter
	SetKeyAttri(value ISynHighlighterAttributes)               // property KeyAttri Setter
	ConstantAttri() ISynHighlighterAttributes                  // property ConstantAttri Getter
	SetConstantAttri(value ISynHighlighterAttributes)          // property ConstantAttri Setter
	ObjectAttri() ISynHighlighterAttributes                    // property ObjectAttri Getter
	SetObjectAttri(value ISynHighlighterAttributes)            // property ObjectAttri Setter
	EntityAttri() ISynHighlighterAttributes                    // property EntityAttri Getter
	SetEntityAttri(value ISynHighlighterAttributes)            // property EntityAttri Setter
	VariableAttri() ISynHighlighterAttributes                  // property VariableAttri Getter
	SetVariableAttri(value ISynHighlighterAttributes)          // property VariableAttri Setter
	DotAttri() ISynHighlighterAttributes                       // property DotAttri Getter
	SetDotAttri(value ISynHighlighterAttributes)               // property DotAttri Setter
	KeyWords() IStrings                                        // property KeyWords Getter
	SetKeyWords(value IStrings)                                // property KeyWords Setter
	Constants() IStrings                                       // property Constants Getter
	SetConstants(value IStrings)                               // property Constants Setter
	Objects() IStrings                                         // property Objects Getter
	SetObjects(value IStrings)                                 // property Objects Setter
	NumberAttri() ISynHighlighterAttributes                    // property NumberAttri Getter
	SetNumberAttri(value ISynHighlighterAttributes)            // property NumberAttri Setter
	PreprocessorAttri() ISynHighlighterAttributes              // property PreprocessorAttri Getter
	SetPreprocessorAttri(value ISynHighlighterAttributes)      // property PreprocessorAttri Setter
	SpaceAttri() ISynHighlighterAttributes                     // property SpaceAttri Getter
	SetSpaceAttri(value ISynHighlighterAttributes)             // property SpaceAttri Setter
	StringAttri() ISynHighlighterAttributes                    // property StringAttri Getter
	SetStringAttri(value ISynHighlighterAttributes)            // property StringAttri Setter
	SymbolAttri() ISynHighlighterAttributes                    // property SymbolAttri Getter
	SetSymbolAttri(value ISynHighlighterAttributes)            // property SymbolAttri Setter
	StringDelim() types.TStringDelim                           // property StringDelim Getter
	SetStringDelim(value types.TStringDelim)                   // property StringDelim Setter
	Markup() bool                                              // property Markup Getter
	SetMarkup(value bool)                                      // property Markup Setter
	Entity() bool                                              // property Entity Getter
	SetEntity(value bool)                                      // property Entity Setter
	DollarVariables() bool                                     // property DollarVariables Getter
	SetDollarVariables(value bool)                             // property DollarVariables Setter
	ActiveDot() bool                                           // property ActiveDot Getter
	SetActiveDot(value bool)                                   // property ActiveDot Setter
}

type TSynAnySyn struct {
	TSynCustomHighlighter
}

func (m *TSynAnySyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) GetTokenID() types.SHAnyTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synAnySynAPI().SysCallN(2, m.Instance())
	return types.SHAnyTtkTokenKind(r)
}

func (m *TSynAnySyn) IsConstant(constant string) bool {
	if !m.IsValid() {
		return false
	}
	r := synAnySynAPI().SysCallN(3, m.Instance(), api.PasStr(constant))
	return api.GoBool(r)
}

func (m *TSynAnySyn) LoadHighLighter(file string) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(4, m.Instance(), api.PasStr(file))
}

func (m *TSynAnySyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) Comments() types.CommentStyles {
	if !m.IsValid() {
		return 0
	}
	r := synAnySynAPI().SysCallN(6, 0, m.Instance())
	return types.CommentStyles(r)
}

func (m *TSynAnySyn) SetComments(value types.CommentStyles) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TSynAnySyn) DetectPreprocessor() bool {
	if !m.IsValid() {
		return false
	}
	r := synAnySynAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynAnySyn) SetDetectPreprocessor(value bool) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynAnySyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) IdentifierChars() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synAnySynAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynAnySyn) SetIdentifierChars(value string) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynAnySyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) ConstantAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetConstantAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) ObjectAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetObjectAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) EntityAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(13, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetEntityAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) VariableAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(14, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetVariableAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) DotAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(15, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetDotAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) KeyWords() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(16, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynAnySyn) SetKeyWords(value IStrings) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) Constants() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(17, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynAnySyn) SetConstants(value IStrings) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) Objects() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(18, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynAnySyn) SetObjects(value IStrings) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) NumberAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(19, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetNumberAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(19, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) PreprocessorAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(20, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetPreprocessorAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(20, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(21, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) StringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(22, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synAnySynAPI().SysCallN(23, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynAnySyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(23, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAnySyn) StringDelim() types.TStringDelim {
	if !m.IsValid() {
		return 0
	}
	r := synAnySynAPI().SysCallN(24, 0, m.Instance())
	return types.TStringDelim(r)
}

func (m *TSynAnySyn) SetStringDelim(value types.TStringDelim) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TSynAnySyn) Markup() bool {
	if !m.IsValid() {
		return false
	}
	r := synAnySynAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynAnySyn) SetMarkup(value bool) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynAnySyn) Entity() bool {
	if !m.IsValid() {
		return false
	}
	r := synAnySynAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynAnySyn) SetEntity(value bool) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynAnySyn) DollarVariables() bool {
	if !m.IsValid() {
		return false
	}
	r := synAnySynAPI().SysCallN(27, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynAnySyn) SetDollarVariables(value bool) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(27, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynAnySyn) ActiveDot() bool {
	if !m.IsValid() {
		return false
	}
	r := synAnySynAPI().SysCallN(28, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynAnySyn) SetActiveDot(value bool) {
	if !m.IsValid() {
		return
	}
	synAnySynAPI().SysCallN(28, 1, m.Instance(), api.PasBool(value))
}

// NewSynAnySyn class constructor
func NewSynAnySyn(owner IComponent) ISynAnySyn {
	r := synAnySynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynAnySyn(r)
}

func TSynAnySynClass() types.TClass {
	r := synAnySynAPI().SysCallN(29)
	return types.TClass(r)
}

var (
	synAnySynOnce   base.Once
	synAnySynImport *imports.Imports = nil
)

func synAnySynAPI() *imports.Imports {
	synAnySynOnce.Do(func() {
		synAnySynImport = api.NewDefaultImports()
		synAnySynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynAnySyn_Create", 0), // constructor NewSynAnySyn
			/* 1 */ imports.NewTable("TSynAnySyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynAnySyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynAnySyn_IsConstant", 0), // function IsConstant
			/* 4 */ imports.NewTable("TSynAnySyn_LoadHighLighter", 0), // procedure LoadHighLighter
			/* 5 */ imports.NewTable("TSynAnySyn_CommentAttri", 0), // property CommentAttri
			/* 6 */ imports.NewTable("TSynAnySyn_Comments", 0), // property Comments
			/* 7 */ imports.NewTable("TSynAnySyn_DetectPreprocessor", 0), // property DetectPreprocessor
			/* 8 */ imports.NewTable("TSynAnySyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 9 */ imports.NewTable("TSynAnySyn_IdentifierChars", 0), // property IdentifierChars
			/* 10 */ imports.NewTable("TSynAnySyn_KeyAttri", 0), // property KeyAttri
			/* 11 */ imports.NewTable("TSynAnySyn_ConstantAttri", 0), // property ConstantAttri
			/* 12 */ imports.NewTable("TSynAnySyn_ObjectAttri", 0), // property ObjectAttri
			/* 13 */ imports.NewTable("TSynAnySyn_EntityAttri", 0), // property EntityAttri
			/* 14 */ imports.NewTable("TSynAnySyn_VariableAttri", 0), // property VariableAttri
			/* 15 */ imports.NewTable("TSynAnySyn_DotAttri", 0), // property DotAttri
			/* 16 */ imports.NewTable("TSynAnySyn_KeyWords", 0), // property KeyWords
			/* 17 */ imports.NewTable("TSynAnySyn_Constants", 0), // property Constants
			/* 18 */ imports.NewTable("TSynAnySyn_Objects", 0), // property Objects
			/* 19 */ imports.NewTable("TSynAnySyn_NumberAttri", 0), // property NumberAttri
			/* 20 */ imports.NewTable("TSynAnySyn_PreprocessorAttri", 0), // property PreprocessorAttri
			/* 21 */ imports.NewTable("TSynAnySyn_SpaceAttri", 0), // property SpaceAttri
			/* 22 */ imports.NewTable("TSynAnySyn_StringAttri", 0), // property StringAttri
			/* 23 */ imports.NewTable("TSynAnySyn_SymbolAttri", 0), // property SymbolAttri
			/* 24 */ imports.NewTable("TSynAnySyn_StringDelim", 0), // property StringDelim
			/* 25 */ imports.NewTable("TSynAnySyn_Markup", 0), // property Markup
			/* 26 */ imports.NewTable("TSynAnySyn_Entity", 0), // property Entity
			/* 27 */ imports.NewTable("TSynAnySyn_DollarVariables", 0), // property DollarVariables
			/* 28 */ imports.NewTable("TSynAnySyn_ActiveDot", 0), // property ActiveDot
			/* 29 */ imports.NewTable("TSynAnySyn_TClass", 0), // function TSynAnySynClass
		}
	})
	return synAnySynImport
}
