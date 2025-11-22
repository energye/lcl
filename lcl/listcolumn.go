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

// IListColumn Parent: ICollectionItem
type IListColumn interface {
	ICollectionItem
	WidthType() int32                            // property WidthType Getter
	Alignment() types.TAlignment                 // property Alignment Getter
	SetAlignment(value types.TAlignment)         // property Alignment Setter
	AutoSize() bool                              // property AutoSize Getter
	SetAutoSize(value bool)                      // property AutoSize Setter
	Caption() string                             // property Caption Getter
	SetCaption(value string)                     // property Caption Setter
	ImageIndex() int32                           // property ImageIndex Getter
	SetImageIndex(value int32)                   // property ImageIndex Setter
	MaxWidth() int32                             // property MaxWidth Getter
	SetMaxWidth(value int32)                     // property MaxWidth Setter
	MinWidth() int32                             // property MinWidth Getter
	SetMinWidth(value int32)                     // property MinWidth Setter
	Tag() uintptr                                // property Tag Getter
	SetTag(value uintptr)                        // property Tag Setter
	Visible() bool                               // property Visible Getter
	SetVisible(value bool)                       // property Visible Setter
	Width() int32                                // property Width Getter
	SetWidth(value int32)                        // property Width Setter
	SortIndicator() types.TSortIndicator         // property SortIndicator Getter
	SetSortIndicator(value types.TSortIndicator) // property SortIndicator Setter
}

type TListColumn struct {
	TCollectionItem
}

func (m *TListColumn) WidthType() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TListColumn) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(2, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TListColumn) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TListColumn) AutoSize() bool {
	if !m.IsValid() {
		return false
	}
	r := listColumnAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListColumn) SetAutoSize(value bool) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TListColumn) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := listColumnAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TListColumn) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TListColumn) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TListColumn) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TListColumn) MaxWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TListColumn) SetMaxWidth(value int32) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TListColumn) MinWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TListColumn) SetMinWidth(value int32) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TListColumn) Tag() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(8, 0, m.Instance())
	return uintptr(r)
}

func (m *TListColumn) SetTag(value uintptr) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TListColumn) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := listColumnAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListColumn) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TListColumn) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TListColumn) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TListColumn) SortIndicator() types.TSortIndicator {
	if !m.IsValid() {
		return 0
	}
	r := listColumnAPI().SysCallN(11, 0, m.Instance())
	return types.TSortIndicator(r)
}

func (m *TListColumn) SetSortIndicator(value types.TSortIndicator) {
	if !m.IsValid() {
		return
	}
	listColumnAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

// NewListColumn class constructor
func NewListColumn(collection ICollection) IListColumn {
	r := listColumnAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsListColumn(r)
}

var (
	listColumnOnce   base.Once
	listColumnImport *imports.Imports = nil
)

func listColumnAPI() *imports.Imports {
	listColumnOnce.Do(func() {
		listColumnImport = api.NewDefaultImports()
		listColumnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListColumn_Create", 0), // constructor NewListColumn
			/* 1 */ imports.NewTable("TListColumn_WidthType", 0), // property WidthType
			/* 2 */ imports.NewTable("TListColumn_Alignment", 0), // property Alignment
			/* 3 */ imports.NewTable("TListColumn_AutoSize", 0), // property AutoSize
			/* 4 */ imports.NewTable("TListColumn_Caption", 0), // property Caption
			/* 5 */ imports.NewTable("TListColumn_ImageIndex", 0), // property ImageIndex
			/* 6 */ imports.NewTable("TListColumn_MaxWidth", 0), // property MaxWidth
			/* 7 */ imports.NewTable("TListColumn_MinWidth", 0), // property MinWidth
			/* 8 */ imports.NewTable("TListColumn_Tag", 0), // property Tag
			/* 9 */ imports.NewTable("TListColumn_Visible", 0), // property Visible
			/* 10 */ imports.NewTable("TListColumn_Width", 0), // property Width
			/* 11 */ imports.NewTable("TListColumn_SortIndicator", 0), // property SortIndicator
		}
	})
	return listColumnImport
}
