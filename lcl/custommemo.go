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

// ICustomMemo Parent: ICustomEdit
type ICustomMemo interface {
	ICustomEdit
	Append(value string)                    // procedure
	Lines() IStrings                        // property Lines Getter
	SetLines(value IStrings)                // property Lines Setter
	HorzScrollBar() IMemoScrollbar          // property HorzScrollBar Getter
	SetHorzScrollBar(value IMemoScrollbar)  // property HorzScrollBar Setter
	VertScrollBar() IMemoScrollbar          // property VertScrollBar Getter
	SetVertScrollBar(value IMemoScrollbar)  // property VertScrollBar Setter
	ScrollBars() types.TScrollStyle         // property ScrollBars Getter
	SetScrollBars(value types.TScrollStyle) // property ScrollBars Setter
	WantReturns() bool                      // property WantReturns Getter
	SetWantReturns(value bool)              // property WantReturns Setter
	WantTabs() bool                         // property WantTabs Getter
	SetWantTabs(value bool)                 // property WantTabs Setter
	WordWrap() bool                         // property WordWrap Getter
	SetWordWrap(value bool)                 // property WordWrap Setter
}

type TCustomMemo struct {
	TCustomEdit
}

func (m *TCustomMemo) Append(value string) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(1, m.Instance(), api.PasStr(value))
}

func (m *TCustomMemo) Lines() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customMemoAPI().SysCallN(2, 0, m.Instance())
	return AsStrings(r)
}

func (m *TCustomMemo) SetLines(value IStrings) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomMemo) HorzScrollBar() IMemoScrollbar {
	if !m.IsValid() {
		return nil
	}
	r := customMemoAPI().SysCallN(3, 0, m.Instance())
	return AsMemoScrollbar(r)
}

func (m *TCustomMemo) SetHorzScrollBar(value IMemoScrollbar) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomMemo) VertScrollBar() IMemoScrollbar {
	if !m.IsValid() {
		return nil
	}
	r := customMemoAPI().SysCallN(4, 0, m.Instance())
	return AsMemoScrollbar(r)
}

func (m *TCustomMemo) SetVertScrollBar(value IMemoScrollbar) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomMemo) ScrollBars() types.TScrollStyle {
	if !m.IsValid() {
		return 0
	}
	r := customMemoAPI().SysCallN(5, 0, m.Instance())
	return types.TScrollStyle(r)
}

func (m *TCustomMemo) SetScrollBars(value types.TScrollStyle) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomMemo) WantReturns() bool {
	if !m.IsValid() {
		return false
	}
	r := customMemoAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomMemo) SetWantReturns(value bool) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomMemo) WantTabs() bool {
	if !m.IsValid() {
		return false
	}
	r := customMemoAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomMemo) SetWantTabs(value bool) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomMemo) WordWrap() bool {
	if !m.IsValid() {
		return false
	}
	r := customMemoAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomMemo) SetWordWrap(value bool) {
	if !m.IsValid() {
		return
	}
	customMemoAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

// NewCustomMemo class constructor
func NewCustomMemo(owner IComponent) ICustomMemo {
	r := customMemoAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomMemo(r)
}

func TCustomMemoClass() types.TClass {
	r := customMemoAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	customMemoOnce   base.Once
	customMemoImport *imports.Imports = nil
)

func customMemoAPI() *imports.Imports {
	customMemoOnce.Do(func() {
		customMemoImport = api.NewDefaultImports()
		customMemoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomMemo_Create", 0), // constructor NewCustomMemo
			/* 1 */ imports.NewTable("TCustomMemo_Append", 0), // procedure Append
			/* 2 */ imports.NewTable("TCustomMemo_Lines", 0), // property Lines
			/* 3 */ imports.NewTable("TCustomMemo_HorzScrollBar", 0), // property HorzScrollBar
			/* 4 */ imports.NewTable("TCustomMemo_VertScrollBar", 0), // property VertScrollBar
			/* 5 */ imports.NewTable("TCustomMemo_ScrollBars", 0), // property ScrollBars
			/* 6 */ imports.NewTable("TCustomMemo_WantReturns", 0), // property WantReturns
			/* 7 */ imports.NewTable("TCustomMemo_WantTabs", 0), // property WantTabs
			/* 8 */ imports.NewTable("TCustomMemo_WordWrap", 0), // property WordWrap
			/* 9 */ imports.NewTable("TCustomMemo_TClass", 0), // function TCustomMemoClass
		}
	})
	return customMemoImport
}
