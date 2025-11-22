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

// IParaAttributes Parent: IPersistent
type IParaAttributes interface {
	IPersistent
	Alignment() types.TAlignment              // property Alignment Getter
	SetAlignment(value types.TAlignment)      // property Alignment Setter
	FirstIndent() int32                       // property FirstIndent Getter
	SetFirstIndent(value int32)               // property FirstIndent Setter
	LeftIndent() int32                        // property LeftIndent Getter
	SetLeftIndent(value int32)                // property LeftIndent Setter
	RightIndent() int32                       // property RightIndent Getter
	SetRightIndent(value int32)               // property RightIndent Setter
	Numbering() types.TNumberingStyle         // property Numbering Getter
	SetNumbering(value types.TNumberingStyle) // property Numbering Setter
	TabCount() int32                          // property TabCount Getter
	SetTabCount(value int32)                  // property TabCount Setter
	Tab(index byte) int32                     // property Tab Getter
	SetTab(index byte, value int32)           // property Tab Setter
}

type TParaAttributes struct {
	TPersistent
}

func (m *TParaAttributes) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := paraAttributesAPI().SysCallN(1, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TParaAttributes) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	paraAttributesAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TParaAttributes) FirstIndent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := paraAttributesAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TParaAttributes) SetFirstIndent(value int32) {
	if !m.IsValid() {
		return
	}
	paraAttributesAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TParaAttributes) LeftIndent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := paraAttributesAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TParaAttributes) SetLeftIndent(value int32) {
	if !m.IsValid() {
		return
	}
	paraAttributesAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TParaAttributes) RightIndent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := paraAttributesAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TParaAttributes) SetRightIndent(value int32) {
	if !m.IsValid() {
		return
	}
	paraAttributesAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TParaAttributes) Numbering() types.TNumberingStyle {
	if !m.IsValid() {
		return 0
	}
	r := paraAttributesAPI().SysCallN(5, 0, m.Instance())
	return types.TNumberingStyle(r)
}

func (m *TParaAttributes) SetNumbering(value types.TNumberingStyle) {
	if !m.IsValid() {
		return
	}
	paraAttributesAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TParaAttributes) TabCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := paraAttributesAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TParaAttributes) SetTabCount(value int32) {
	if !m.IsValid() {
		return
	}
	paraAttributesAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TParaAttributes) Tab(index byte) int32 {
	if !m.IsValid() {
		return 0
	}
	r := paraAttributesAPI().SysCallN(7, 0, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TParaAttributes) SetTab(index byte, value int32) {
	if !m.IsValid() {
		return
	}
	paraAttributesAPI().SysCallN(7, 1, m.Instance(), uintptr(index), uintptr(value))
}

// NewParaAttributes class constructor
func NewParaAttributes(owner IRichMemo) IParaAttributes {
	r := paraAttributesAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsParaAttributes(r)
}

var (
	paraAttributesOnce   base.Once
	paraAttributesImport *imports.Imports = nil
)

func paraAttributesAPI() *imports.Imports {
	paraAttributesOnce.Do(func() {
		paraAttributesImport = api.NewDefaultImports()
		paraAttributesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TParaAttributes_Create", 0), // constructor NewParaAttributes
			/* 1 */ imports.NewTable("TParaAttributes_Alignment", 0), // property Alignment
			/* 2 */ imports.NewTable("TParaAttributes_FirstIndent", 0), // property FirstIndent
			/* 3 */ imports.NewTable("TParaAttributes_LeftIndent", 0), // property LeftIndent
			/* 4 */ imports.NewTable("TParaAttributes_RightIndent", 0), // property RightIndent
			/* 5 */ imports.NewTable("TParaAttributes_Numbering", 0), // property Numbering
			/* 6 */ imports.NewTable("TParaAttributes_TabCount", 0), // property TabCount
			/* 7 */ imports.NewTable("TParaAttributes_Tab", 0), // property Tab
		}
	})
	return paraAttributesImport
}
