//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IListColumn Parent: ICollectionItem
type IListColumn interface {
	ICollectionItem
	WidthType() int32                       // property
	Alignment() TAlignment                  // property
	SetAlignment(AValue TAlignment)         // property
	AutoSize() bool                         // property
	SetAutoSize(AValue bool)                // property
	Caption() string                        // property
	SetCaption(AValue string)               // property
	ImageIndex() TImageIndex                // property
	SetImageIndex(AValue TImageIndex)       // property
	MaxWidth() int32                        // property
	SetMaxWidth(AValue int32)               // property
	MinWidth() int32                        // property
	SetMinWidth(AValue int32)               // property
	Tag() uint32                            // property
	SetTag(AValue uint32)                   // property
	Visible() bool                          // property
	SetVisible(AValue bool)                 // property
	Width() int32                           // property
	SetWidth(AValue int32)                  // property
	SortIndicator() TSortIndicator          // property
	SetSortIndicator(AValue TSortIndicator) // property
}

// TListColumn Parent: TCollectionItem
type TListColumn struct {
	TCollectionItem
}

func NewListColumn(ACollection ICollection) IListColumn {
	r1 := listColumnImportAPI().SysCallN(4, GetObjectUintptr(ACollection))
	return AsListColumn(r1)
}

func (m *TListColumn) WidthType() int32 {
	r1 := listColumnImportAPI().SysCallN(12, m.Instance())
	return int32(r1)
}

func (m *TListColumn) Alignment() TAlignment {
	r1 := listColumnImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TListColumn) SetAlignment(AValue TAlignment) {
	listColumnImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) AutoSize() bool {
	r1 := listColumnImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListColumn) SetAutoSize(AValue bool) {
	listColumnImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListColumn) Caption() string {
	r1 := listColumnImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TListColumn) SetCaption(AValue string) {
	listColumnImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TListColumn) ImageIndex() TImageIndex {
	r1 := listColumnImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListColumn) SetImageIndex(AValue TImageIndex) {
	listColumnImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) MaxWidth() int32 {
	r1 := listColumnImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListColumn) SetMaxWidth(AValue int32) {
	listColumnImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) MinWidth() int32 {
	r1 := listColumnImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListColumn) SetMinWidth(AValue int32) {
	listColumnImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) Tag() uint32 {
	r1 := listColumnImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TListColumn) SetTag(AValue uint32) {
	listColumnImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) Visible() bool {
	r1 := listColumnImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListColumn) SetVisible(AValue bool) {
	listColumnImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListColumn) Width() int32 {
	r1 := listColumnImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListColumn) SetWidth(AValue int32) {
	listColumnImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TListColumn) SortIndicator() TSortIndicator {
	r1 := listColumnImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TSortIndicator(r1)
}

func (m *TListColumn) SetSortIndicator(AValue TSortIndicator) {
	listColumnImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func ListColumnClass() TClass {
	ret := listColumnImportAPI().SysCallN(3)
	return TClass(ret)
}

var (
	listColumnImport       *imports.Imports = nil
	listColumnImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListColumn_Alignment", 0),
		/*1*/ imports.NewTable("ListColumn_AutoSize", 0),
		/*2*/ imports.NewTable("ListColumn_Caption", 0),
		/*3*/ imports.NewTable("ListColumn_Class", 0),
		/*4*/ imports.NewTable("ListColumn_Create", 0),
		/*5*/ imports.NewTable("ListColumn_ImageIndex", 0),
		/*6*/ imports.NewTable("ListColumn_MaxWidth", 0),
		/*7*/ imports.NewTable("ListColumn_MinWidth", 0),
		/*8*/ imports.NewTable("ListColumn_SortIndicator", 0),
		/*9*/ imports.NewTable("ListColumn_Tag", 0),
		/*10*/ imports.NewTable("ListColumn_Visible", 0),
		/*11*/ imports.NewTable("ListColumn_Width", 0),
		/*12*/ imports.NewTable("ListColumn_WidthType", 0),
	}
)

func listColumnImportAPI() *imports.Imports {
	if listColumnImport == nil {
		listColumnImport = NewDefaultImports()
		listColumnImport.SetImportTable(listColumnImportTables)
		listColumnImportTables = nil
	}
	return listColumnImport
}
