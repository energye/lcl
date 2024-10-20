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

// IListItem Parent: IPersistent
type IListItem interface {
	IPersistent
	Caption() string                                                        // property
	SetCaption(AValue string)                                               // property
	Checked() bool                                                          // property
	SetChecked(AValue bool)                                                 // property
	Cut() bool                                                              // property
	SetCut(AValue bool)                                                     // property
	Data() uintptr                                                          // property
	SetData(AValue uintptr)                                                 // property
	DropTarget() bool                                                       // property
	SetDropTarget(AValue bool)                                              // property
	Focused() bool                                                          // property
	SetFocused(AValue bool)                                                 // property
	Index() int32                                                           // property
	ImageIndex() TImageIndex                                                // property
	SetImageIndex(AValue TImageIndex)                                       // property
	Left() int32                                                            // property
	SetLeft(AValue int32)                                                   // property
	ListView() ICustomListView                                              // property
	Owner() IListItems                                                      // property
	Position() (resultPoint TPoint)                                         // property
	SetPosition(AValue *TPoint)                                             // property
	Selected() bool                                                         // property
	SetSelected(AValue bool)                                                // property
	StateIndex() TImageIndex                                                // property
	SetStateIndex(AValue TImageIndex)                                       // property
	SubItems() IStrings                                                     // property
	SetSubItems(AValue IStrings)                                            // property
	SubItemImages(AIndex int32) int32                                       // property
	SetSubItemImages(AIndex int32, AValue int32)                            // property
	Top() int32                                                             // property
	SetTop(AValue int32)                                                    // property
	DisplayRect(Code TDisplayCode) (resultRect TRect)                       // function
	DisplayRectSubItem(subItem int32, Code TDisplayCode) (resultRect TRect) // function
	EditCaption() bool                                                      // function
	GetStates() TListItemStates                                             // function
	Delete()                                                                // procedure
	MakeVisible(PartialOK bool)                                             // procedure
}

// TListItem Parent: TPersistent
type TListItem struct {
	TPersistent
}

func NewListItem(AOwner IListItems) IListItem {
	r1 := listItemImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsListItem(r1)
}

func (m *TListItem) Caption() string {
	r1 := listItemImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TListItem) SetCaption(AValue string) {
	listItemImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TListItem) Checked() bool {
	r1 := listItemImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetChecked(AValue bool) {
	listItemImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Cut() bool {
	r1 := listItemImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetCut(AValue bool) {
	listItemImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Data() uintptr {
	r1 := listItemImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TListItem) SetData(AValue uintptr) {
	listItemImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) DropTarget() bool {
	r1 := listItemImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetDropTarget(AValue bool) {
	listItemImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Focused() bool {
	r1 := listItemImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetFocused(AValue bool) {
	listItemImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Index() int32 {
	r1 := listItemImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TListItem) ImageIndex() TImageIndex {
	r1 := listItemImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListItem) SetImageIndex(AValue TImageIndex) {
	listItemImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) Left() int32 {
	r1 := listItemImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListItem) SetLeft(AValue int32) {
	listItemImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) ListView() ICustomListView {
	r1 := listItemImportAPI().SysCallN(16, m.Instance())
	return AsCustomListView(r1)
}

func (m *TListItem) Owner() IListItems {
	r1 := listItemImportAPI().SysCallN(18, m.Instance())
	return AsListItems(r1)
}

func (m *TListItem) Position() (resultPoint TPoint) {
	listItemImportAPI().SysCallN(19, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TListItem) SetPosition(AValue *TPoint) {
	listItemImportAPI().SysCallN(19, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TListItem) Selected() bool {
	r1 := listItemImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetSelected(AValue bool) {
	listItemImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) StateIndex() TImageIndex {
	r1 := listItemImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListItem) SetStateIndex(AValue TImageIndex) {
	listItemImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) SubItems() IStrings {
	r1 := listItemImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TListItem) SetSubItems(AValue IStrings) {
	listItemImportAPI().SysCallN(23, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListItem) SubItemImages(AIndex int32) int32 {
	r1 := listItemImportAPI().SysCallN(22, 0, m.Instance(), uintptr(AIndex))
	return int32(r1)
}

func (m *TListItem) SetSubItemImages(AIndex int32, AValue int32) {
	listItemImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AIndex), uintptr(AValue))
}

func (m *TListItem) Top() int32 {
	r1 := listItemImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListItem) SetTop(AValue int32) {
	listItemImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) DisplayRect(Code TDisplayCode) (resultRect TRect) {
	listItemImportAPI().SysCallN(7, m.Instance(), uintptr(Code), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TListItem) DisplayRectSubItem(subItem int32, Code TDisplayCode) (resultRect TRect) {
	listItemImportAPI().SysCallN(8, m.Instance(), uintptr(subItem), uintptr(Code), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TListItem) EditCaption() bool {
	r1 := listItemImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TListItem) GetStates() TListItemStates {
	r1 := listItemImportAPI().SysCallN(12, m.Instance())
	return TListItemStates(r1)
}

func ListItemClass() TClass {
	ret := listItemImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TListItem) Delete() {
	listItemImportAPI().SysCallN(6, m.Instance())
}

func (m *TListItem) MakeVisible(PartialOK bool) {
	listItemImportAPI().SysCallN(17, m.Instance(), PascalBool(PartialOK))
}

var (
	listItemImport       *imports.Imports = nil
	listItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListItem_Caption", 0),
		/*1*/ imports.NewTable("ListItem_Checked", 0),
		/*2*/ imports.NewTable("ListItem_Class", 0),
		/*3*/ imports.NewTable("ListItem_Create", 0),
		/*4*/ imports.NewTable("ListItem_Cut", 0),
		/*5*/ imports.NewTable("ListItem_Data", 0),
		/*6*/ imports.NewTable("ListItem_Delete", 0),
		/*7*/ imports.NewTable("ListItem_DisplayRect", 0),
		/*8*/ imports.NewTable("ListItem_DisplayRectSubItem", 0),
		/*9*/ imports.NewTable("ListItem_DropTarget", 0),
		/*10*/ imports.NewTable("ListItem_EditCaption", 0),
		/*11*/ imports.NewTable("ListItem_Focused", 0),
		/*12*/ imports.NewTable("ListItem_GetStates", 0),
		/*13*/ imports.NewTable("ListItem_ImageIndex", 0),
		/*14*/ imports.NewTable("ListItem_Index", 0),
		/*15*/ imports.NewTable("ListItem_Left", 0),
		/*16*/ imports.NewTable("ListItem_ListView", 0),
		/*17*/ imports.NewTable("ListItem_MakeVisible", 0),
		/*18*/ imports.NewTable("ListItem_Owner", 0),
		/*19*/ imports.NewTable("ListItem_Position", 0),
		/*20*/ imports.NewTable("ListItem_Selected", 0),
		/*21*/ imports.NewTable("ListItem_StateIndex", 0),
		/*22*/ imports.NewTable("ListItem_SubItemImages", 0),
		/*23*/ imports.NewTable("ListItem_SubItems", 0),
		/*24*/ imports.NewTable("ListItem_Top", 0),
	}
)

func listItemImportAPI() *imports.Imports {
	if listItemImport == nil {
		listItemImport = NewDefaultImports()
		listItemImport.SetImportTable(listItemImportTables)
		listItemImportTables = nil
	}
	return listItemImport
}
