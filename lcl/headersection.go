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

// IHeaderSection Parent: ICollectionItem
type IHeaderSection interface {
	ICollectionItem
	Left() int32                              // property Left Getter
	Right() int32                             // property Right Getter
	State() types.THeaderSectionState         // property State Getter
	SetState(value types.THeaderSectionState) // property State Setter
	// OriginalIndex
	//  index which doesn't change when the user reorders the sections
	OriginalIndex() int32                // property OriginalIndex Getter
	Alignment() types.TAlignment         // property Alignment Getter
	SetAlignment(value types.TAlignment) // property Alignment Setter
	ImageIndex() int32                   // property ImageIndex Getter
	SetImageIndex(value int32)           // property ImageIndex Setter
	MaxWidth() int32                     // property MaxWidth Getter
	SetMaxWidth(value int32)             // property MaxWidth Setter
	MinWidth() int32                     // property MinWidth Getter
	SetMinWidth(value int32)             // property MinWidth Setter
	Text() string                        // property Text Getter
	SetText(value string)                // property Text Setter
	Width() int32                        // property Width Getter
	SetWidth(value int32)                // property Width Setter
	Visible() bool                       // property Visible Getter
	SetVisible(value bool)               // property Visible Setter
}

type THeaderSection struct {
	TCollectionItem
}

func (m *THeaderSection) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *THeaderSection) Right() int32 {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *THeaderSection) State() types.THeaderSectionState {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(3, 0, m.Instance())
	return types.THeaderSectionState(r)
}

func (m *THeaderSection) SetState(value types.THeaderSectionState) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *THeaderSection) OriginalIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *THeaderSection) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(5, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *THeaderSection) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *THeaderSection) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *THeaderSection) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *THeaderSection) MaxWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *THeaderSection) SetMaxWidth(value int32) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *THeaderSection) MinWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *THeaderSection) SetMinWidth(value int32) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *THeaderSection) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := headerSectionAPI().SysCallN(9, 0, m.Instance())
	return api.GoStr(r)
}

func (m *THeaderSection) SetText(value string) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *THeaderSection) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := headerSectionAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *THeaderSection) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *THeaderSection) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := headerSectionAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *THeaderSection) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	headerSectionAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

// NewHeaderSection class constructor
func NewHeaderSection(collection ICollection) IHeaderSection {
	r := headerSectionAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsHeaderSection(r)
}

var (
	headerSectionOnce   base.Once
	headerSectionImport *imports.Imports = nil
)

func headerSectionAPI() *imports.Imports {
	headerSectionOnce.Do(func() {
		headerSectionImport = api.NewDefaultImports()
		headerSectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THeaderSection_Create", 0), // constructor NewHeaderSection
			/* 1 */ imports.NewTable("THeaderSection_Left", 0), // property Left
			/* 2 */ imports.NewTable("THeaderSection_Right", 0), // property Right
			/* 3 */ imports.NewTable("THeaderSection_State", 0), // property State
			/* 4 */ imports.NewTable("THeaderSection_OriginalIndex", 0), // property OriginalIndex
			/* 5 */ imports.NewTable("THeaderSection_Alignment", 0), // property Alignment
			/* 6 */ imports.NewTable("THeaderSection_ImageIndex", 0), // property ImageIndex
			/* 7 */ imports.NewTable("THeaderSection_MaxWidth", 0), // property MaxWidth
			/* 8 */ imports.NewTable("THeaderSection_MinWidth", 0), // property MinWidth
			/* 9 */ imports.NewTable("THeaderSection_Text", 0), // property Text
			/* 10 */ imports.NewTable("THeaderSection_Width", 0), // property Width
			/* 11 */ imports.NewTable("THeaderSection_Visible", 0), // property Visible
		}
	})
	return headerSectionImport
}
