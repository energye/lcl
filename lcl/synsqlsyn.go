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

// ISynSQLSyn Parent: ISynCustomHighlighter
type ISynSQLSyn interface {
	ISynCustomHighlighter
	GetDefaultAttribute(index int32) ISynHighlighterAttributes // function
	GetTokenID() types.SHSQLTtkTokenKind                       // function
	CommentAttri() ISynHighlighterAttributes                   // property CommentAttri Getter
	SetCommentAttri(value ISynHighlighterAttributes)           // property CommentAttri Setter
	DataTypeAttri() ISynHighlighterAttributes                  // property DataTypeAttri Getter
	SetDataTypeAttri(value ISynHighlighterAttributes)          // property DataTypeAttri Setter
	DefaultPackageAttri() ISynHighlighterAttributes            // property DefaultPackageAttri Getter
	SetDefaultPackageAttri(value ISynHighlighterAttributes)    // property DefaultPackageAttri Setter
	ExceptionAttri() ISynHighlighterAttributes                 // property ExceptionAttri Getter
	SetExceptionAttri(value ISynHighlighterAttributes)         // property ExceptionAttri Setter
	FunctionAttri() ISynHighlighterAttributes                  // property FunctionAttri Getter
	SetFunctionAttri(value ISynHighlighterAttributes)          // property FunctionAttri Setter
	IdentifierAttri() ISynHighlighterAttributes                // property IdentifierAttri Getter
	SetIdentifierAttri(value ISynHighlighterAttributes)        // property IdentifierAttri Setter
	KeyAttri() ISynHighlighterAttributes                       // property KeyAttri Getter
	SetKeyAttri(value ISynHighlighterAttributes)               // property KeyAttri Setter
	NumberAttri() ISynHighlighterAttributes                    // property NumberAttri Getter
	SetNumberAttri(value ISynHighlighterAttributes)            // property NumberAttri Setter
	PLSQLAttri() ISynHighlighterAttributes                     // property PLSQLAttri Getter
	SetPLSQLAttri(value ISynHighlighterAttributes)             // property PLSQLAttri Setter
	SpaceAttri() ISynHighlighterAttributes                     // property SpaceAttri Getter
	SetSpaceAttri(value ISynHighlighterAttributes)             // property SpaceAttri Setter
	SQLPlusAttri() ISynHighlighterAttributes                   // property SQLPlusAttri Getter
	SetSQLPlusAttri(value ISynHighlighterAttributes)           // property SQLPlusAttri Setter
	StringAttri() ISynHighlighterAttributes                    // property StringAttri Getter
	SetStringAttri(value ISynHighlighterAttributes)            // property StringAttri Setter
	SymbolAttri() ISynHighlighterAttributes                    // property SymbolAttri Getter
	SetSymbolAttri(value ISynHighlighterAttributes)            // property SymbolAttri Setter
	TableNameAttri() ISynHighlighterAttributes                 // property TableNameAttri Getter
	SetTableNameAttri(value ISynHighlighterAttributes)         // property TableNameAttri Setter
	TableNames() IStrings                                      // property TableNames Getter
	SetTableNames(value IStrings)                              // property TableNames Setter
	VariableAttri() ISynHighlighterAttributes                  // property VariableAttri Getter
	SetVariableAttri(value ISynHighlighterAttributes)          // property VariableAttri Setter
	SQLDialect() types.TSQLDialect                             // property SQLDialect Getter
	SetSQLDialect(value types.TSQLDialect)                     // property SQLDialect Setter
}

type TSynSQLSyn struct {
	TSynCustomHighlighter
}

func (m *TSynSQLSyn) GetDefaultAttribute(index int32) ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) GetTokenID() types.SHSQLTtkTokenKind {
	if !m.IsValid() {
		return 0
	}
	r := synSQLSynAPI().SysCallN(2, m.Instance())
	return types.SHSQLTtkTokenKind(r)
}

func (m *TSynSQLSyn) CommentAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(3, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetCommentAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) DataTypeAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetDataTypeAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) DefaultPackageAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(5, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetDefaultPackageAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) ExceptionAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(6, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetExceptionAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) FunctionAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(7, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetFunctionAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) IdentifierAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(8, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetIdentifierAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) KeyAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(9, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetKeyAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) NumberAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(10, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetNumberAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) PLSQLAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetPLSQLAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) SpaceAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(12, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetSpaceAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) SQLPlusAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(13, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetSQLPlusAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) StringAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(14, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetStringAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) SymbolAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(15, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetSymbolAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) TableNameAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(16, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetTableNameAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) TableNames() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(17, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynSQLSyn) SetTableNames(value IStrings) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) VariableAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synSQLSynAPI().SysCallN(18, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynSQLSyn) SetVariableAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynSQLSyn) SQLDialect() types.TSQLDialect {
	if !m.IsValid() {
		return 0
	}
	r := synSQLSynAPI().SysCallN(19, 0, m.Instance())
	return types.TSQLDialect(r)
}

func (m *TSynSQLSyn) SetSQLDialect(value types.TSQLDialect) {
	if !m.IsValid() {
		return
	}
	synSQLSynAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

// NewSynSQLSyn class constructor
func NewSynSQLSyn(owner IComponent) ISynSQLSyn {
	r := synSQLSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynSQLSyn(r)
}

func TSynSQLSynClass() types.TClass {
	r := synSQLSynAPI().SysCallN(20)
	return types.TClass(r)
}

var (
	synSQLSynOnce   base.Once
	synSQLSynImport *imports.Imports = nil
)

func synSQLSynAPI() *imports.Imports {
	synSQLSynOnce.Do(func() {
		synSQLSynImport = api.NewDefaultImports()
		synSQLSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynSQLSyn_Create", 0), // constructor NewSynSQLSyn
			/* 1 */ imports.NewTable("TSynSQLSyn_GetDefaultAttribute", 0), // function GetDefaultAttribute
			/* 2 */ imports.NewTable("TSynSQLSyn_GetTokenID", 0), // function GetTokenID
			/* 3 */ imports.NewTable("TSynSQLSyn_CommentAttri", 0), // property CommentAttri
			/* 4 */ imports.NewTable("TSynSQLSyn_DataTypeAttri", 0), // property DataTypeAttri
			/* 5 */ imports.NewTable("TSynSQLSyn_DefaultPackageAttri", 0), // property DefaultPackageAttri
			/* 6 */ imports.NewTable("TSynSQLSyn_ExceptionAttri", 0), // property ExceptionAttri
			/* 7 */ imports.NewTable("TSynSQLSyn_FunctionAttri", 0), // property FunctionAttri
			/* 8 */ imports.NewTable("TSynSQLSyn_IdentifierAttri", 0), // property IdentifierAttri
			/* 9 */ imports.NewTable("TSynSQLSyn_KeyAttri", 0), // property KeyAttri
			/* 10 */ imports.NewTable("TSynSQLSyn_NumberAttri", 0), // property NumberAttri
			/* 11 */ imports.NewTable("TSynSQLSyn_PLSQLAttri", 0), // property PLSQLAttri
			/* 12 */ imports.NewTable("TSynSQLSyn_SpaceAttri", 0), // property SpaceAttri
			/* 13 */ imports.NewTable("TSynSQLSyn_SQLPlusAttri", 0), // property SQLPlusAttri
			/* 14 */ imports.NewTable("TSynSQLSyn_StringAttri", 0), // property StringAttri
			/* 15 */ imports.NewTable("TSynSQLSyn_SymbolAttri", 0), // property SymbolAttri
			/* 16 */ imports.NewTable("TSynSQLSyn_TableNameAttri", 0), // property TableNameAttri
			/* 17 */ imports.NewTable("TSynSQLSyn_TableNames", 0), // property TableNames
			/* 18 */ imports.NewTable("TSynSQLSyn_VariableAttri", 0), // property VariableAttri
			/* 19 */ imports.NewTable("TSynSQLSyn_SQLDialect", 0), // property SQLDialect
			/* 20 */ imports.NewTable("TSynSQLSyn_TClass", 0), // function TSynSQLSynClass
		}
	})
	return synSQLSynImport
}
