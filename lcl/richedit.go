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

// IRichEdit Parent: IRichMemo
type IRichEdit interface {
	IRichMemo
	FindText(searchStr string, startPos int32, length int32, options types.TSearchTypes) int32 // function
	Zoom() int32                                                                               // property Zoom Getter
	SetZoom(value int32)                                                                       // property Zoom Setter
	HideScrollBars() bool                                                                      // property HideScrollBars Getter
	SetHideScrollBars(value bool)                                                              // property HideScrollBars Setter
	PlainText() bool                                                                           // property PlainText Getter
	SetPlainText(value bool)                                                                   // property PlainText Setter
	DefAttributes() ITextAttributes                                                            // property DefAttributes Getter
	SetDefAttributes(value ITextAttributes)                                                    // property DefAttributes Setter
	SelAttributes() ITextAttributes                                                            // property SelAttributes Getter
	SetSelAttributes(value ITextAttributes)                                                    // property SelAttributes Setter
	Paragraph() IParaAttributes                                                                // property Paragraph Getter
	SetParagraph(value IParaAttributes)                                                        // property Paragraph Setter
}

type TRichEdit struct {
	TRichMemo
}

func (m *TRichEdit) FindText(searchStr string, startPos int32, length int32, options types.TSearchTypes) int32 {
	if !m.IsValid() {
		return 0
	}
	r := richEditAPI().SysCallN(1, m.Instance(), api.PasStr(searchStr), uintptr(startPos), uintptr(length), uintptr(options))
	return int32(r)
}

func (m *TRichEdit) Zoom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := richEditAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TRichEdit) SetZoom(value int32) {
	if !m.IsValid() {
		return
	}
	richEditAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TRichEdit) HideScrollBars() bool {
	if !m.IsValid() {
		return false
	}
	r := richEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRichEdit) SetHideScrollBars(value bool) {
	if !m.IsValid() {
		return
	}
	richEditAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TRichEdit) PlainText() bool {
	if !m.IsValid() {
		return false
	}
	r := richEditAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRichEdit) SetPlainText(value bool) {
	if !m.IsValid() {
		return
	}
	richEditAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TRichEdit) DefAttributes() ITextAttributes {
	if !m.IsValid() {
		return nil
	}
	r := richEditAPI().SysCallN(5, 0, m.Instance())
	return AsTextAttributes(r)
}

func (m *TRichEdit) SetDefAttributes(value ITextAttributes) {
	if !m.IsValid() {
		return
	}
	richEditAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TRichEdit) SelAttributes() ITextAttributes {
	if !m.IsValid() {
		return nil
	}
	r := richEditAPI().SysCallN(6, 0, m.Instance())
	return AsTextAttributes(r)
}

func (m *TRichEdit) SetSelAttributes(value ITextAttributes) {
	if !m.IsValid() {
		return
	}
	richEditAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TRichEdit) Paragraph() IParaAttributes {
	if !m.IsValid() {
		return nil
	}
	r := richEditAPI().SysCallN(7, 0, m.Instance())
	return AsParaAttributes(r)
}

func (m *TRichEdit) SetParagraph(value IParaAttributes) {
	if !m.IsValid() {
		return
	}
	richEditAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewRichEdit class constructor
func NewRichEdit(onwer IComponent) IRichEdit {
	r := richEditAPI().SysCallN(0, base.GetObjectUintptr(onwer))
	return AsRichEdit(r)
}

func TRichEditClass() types.TClass {
	r := richEditAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	richEditOnce   base.Once
	richEditImport *imports.Imports = nil
)

func richEditAPI() *imports.Imports {
	richEditOnce.Do(func() {
		richEditImport = api.NewDefaultImports()
		richEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRichEdit_Create", 0), // constructor NewRichEdit
			/* 1 */ imports.NewTable("TRichEdit_FindText", 0), // function FindText
			/* 2 */ imports.NewTable("TRichEdit_Zoom", 0), // property Zoom
			/* 3 */ imports.NewTable("TRichEdit_HideScrollBars", 0), // property HideScrollBars
			/* 4 */ imports.NewTable("TRichEdit_PlainText", 0), // property PlainText
			/* 5 */ imports.NewTable("TRichEdit_DefAttributes", 0), // property DefAttributes
			/* 6 */ imports.NewTable("TRichEdit_SelAttributes", 0), // property SelAttributes
			/* 7 */ imports.NewTable("TRichEdit_Paragraph", 0), // property Paragraph
			/* 8 */ imports.NewTable("TRichEdit_TClass", 0), // function TRichEditClass
		}
	})
	return richEditImport
}
