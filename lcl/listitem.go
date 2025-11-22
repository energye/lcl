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

// IListItem Parent: IPersistent
type IListItem interface {
	IPersistent
	DisplayRect(code types.TDisplayCode) types.TRect                       // function
	DisplayRectSubItem(subItem int32, code types.TDisplayCode) types.TRect // function
	EditCaption() bool                                                     // function
	GetStates() types.TListItemStates                                      // function
	Delete()                                                               // procedure
	MakeVisible(partialOK bool)                                            // procedure
	Caption() string                                                       // property Caption Getter
	SetCaption(value string)                                               // property Caption Setter
	Checked() bool                                                         // property Checked Getter
	SetChecked(value bool)                                                 // property Checked Setter
	Cut() bool                                                             // property Cut Getter
	SetCut(value bool)                                                     // property Cut Setter
	Data() uintptr                                                         // property Data Getter
	SetData(value uintptr)                                                 // property Data Setter
	DropTarget() bool                                                      // property DropTarget Getter
	SetDropTarget(value bool)                                              // property DropTarget Setter
	Focused() bool                                                         // property Focused Getter
	SetFocused(value bool)                                                 // property Focused Setter
	Index() int32                                                          // property Index Getter
	ImageIndex() int32                                                     // property ImageIndex Getter
	SetImageIndex(value int32)                                             // property ImageIndex Setter
	Left() int32                                                           // property Left Getter
	SetLeft(value int32)                                                   // property Left Setter
	ListView() ICustomListView                                             // property ListView Getter
	Owner() IListItems                                                     // property Owner Getter
	Position() types.TPoint                                                // property Position Getter
	SetPosition(value types.TPoint)                                        // property Position Setter
	Selected() bool                                                        // property Selected Getter
	SetSelected(value bool)                                                // property Selected Setter
	StateIndex() int32                                                     // property StateIndex Getter
	SetStateIndex(value int32)                                             // property StateIndex Setter
	SubItems() IStrings                                                    // property SubItems Getter
	SetSubItems(value IStrings)                                            // property SubItems Setter
	SubItemImages(index int32) int32                                       // property SubItemImages Getter
	SetSubItemImages(index int32, value int32)                             // property SubItemImages Setter
	Top() int32                                                            // property Top Getter
	SetTop(value int32)                                                    // property Top Setter
}

type TListItem struct {
	TPersistent
}

func (m *TListItem) DisplayRect(code types.TDisplayCode) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(1, m.Instance(), uintptr(code), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TListItem) DisplayRectSubItem(subItem int32, code types.TDisplayCode) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(2, m.Instance(), uintptr(subItem), uintptr(code), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TListItem) EditCaption() bool {
	if !m.IsValid() {
		return false
	}
	r := listItemAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TListItem) GetStates() types.TListItemStates {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(4, m.Instance())
	return types.TListItemStates(r)
}

func (m *TListItem) Delete() {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(5, m.Instance())
}

func (m *TListItem) MakeVisible(partialOK bool) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(6, m.Instance(), api.PasBool(partialOK))
}

func (m *TListItem) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := listItemAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TListItem) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TListItem) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := listItemAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListItem) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TListItem) Cut() bool {
	if !m.IsValid() {
		return false
	}
	r := listItemAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListItem) SetCut(value bool) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TListItem) Data() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(10, 0, m.Instance())
	return uintptr(r)
}

func (m *TListItem) SetData(value uintptr) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TListItem) DropTarget() bool {
	if !m.IsValid() {
		return false
	}
	r := listItemAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListItem) SetDropTarget(value bool) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TListItem) Focused() bool {
	if !m.IsValid() {
		return false
	}
	r := listItemAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListItem) SetFocused(value bool) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TListItem) Index() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(13, m.Instance())
	return int32(r)
}

func (m *TListItem) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TListItem) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TListItem) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TListItem) SetLeft(value int32) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TListItem) ListView() ICustomListView {
	if !m.IsValid() {
		return nil
	}
	r := listItemAPI().SysCallN(16, m.Instance())
	return AsCustomListView(r)
}

func (m *TListItem) Owner() IListItems {
	if !m.IsValid() {
		return nil
	}
	r := listItemAPI().SysCallN(17, m.Instance())
	return AsListItems(r)
}

func (m *TListItem) Position() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(18, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TListItem) SetPosition(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(18, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TListItem) Selected() bool {
	if !m.IsValid() {
		return false
	}
	r := listItemAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListItem) SetSelected(value bool) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TListItem) StateIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TListItem) SetStateIndex(value int32) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TListItem) SubItems() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := listItemAPI().SysCallN(21, 0, m.Instance())
	return AsStrings(r)
}

func (m *TListItem) SetSubItems(value IStrings) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TListItem) SubItemImages(index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(22, 0, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TListItem) SetSubItemImages(index int32, value int32) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(22, 1, m.Instance(), uintptr(index), uintptr(value))
}

func (m *TListItem) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TListItem) SetTop(value int32) {
	if !m.IsValid() {
		return
	}
	listItemAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

// NewListItem class constructor
func NewListItem(owner IListItems) IListItem {
	r := listItemAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsListItem(r)
}

var (
	listItemOnce   base.Once
	listItemImport *imports.Imports = nil
)

func listItemAPI() *imports.Imports {
	listItemOnce.Do(func() {
		listItemImport = api.NewDefaultImports()
		listItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListItem_Create", 0), // constructor NewListItem
			/* 1 */ imports.NewTable("TListItem_DisplayRect", 0), // function DisplayRect
			/* 2 */ imports.NewTable("TListItem_DisplayRectSubItem", 0), // function DisplayRectSubItem
			/* 3 */ imports.NewTable("TListItem_EditCaption", 0), // function EditCaption
			/* 4 */ imports.NewTable("TListItem_GetStates", 0), // function GetStates
			/* 5 */ imports.NewTable("TListItem_Delete", 0), // procedure Delete
			/* 6 */ imports.NewTable("TListItem_MakeVisible", 0), // procedure MakeVisible
			/* 7 */ imports.NewTable("TListItem_Caption", 0), // property Caption
			/* 8 */ imports.NewTable("TListItem_Checked", 0), // property Checked
			/* 9 */ imports.NewTable("TListItem_Cut", 0), // property Cut
			/* 10 */ imports.NewTable("TListItem_Data", 0), // property Data
			/* 11 */ imports.NewTable("TListItem_DropTarget", 0), // property DropTarget
			/* 12 */ imports.NewTable("TListItem_Focused", 0), // property Focused
			/* 13 */ imports.NewTable("TListItem_Index", 0), // property Index
			/* 14 */ imports.NewTable("TListItem_ImageIndex", 0), // property ImageIndex
			/* 15 */ imports.NewTable("TListItem_Left", 0), // property Left
			/* 16 */ imports.NewTable("TListItem_ListView", 0), // property ListView
			/* 17 */ imports.NewTable("TListItem_Owner", 0), // property Owner
			/* 18 */ imports.NewTable("TListItem_Position", 0), // property Position
			/* 19 */ imports.NewTable("TListItem_Selected", 0), // property Selected
			/* 20 */ imports.NewTable("TListItem_StateIndex", 0), // property StateIndex
			/* 21 */ imports.NewTable("TListItem_SubItems", 0), // property SubItems
			/* 22 */ imports.NewTable("TListItem_SubItemImages", 0), // property SubItemImages
			/* 23 */ imports.NewTable("TListItem_Top", 0), // property Top
		}
	})
	return listItemImport
}
