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

// IMenuItem Parent: ILCLComponent
type IMenuItem interface {
	ILCLComponent
	Merged() IMenuItem                                                // property
	MergedWith() IMenuItem                                            // property
	Count() int32                                                     // property
	Handle() HMENU                                                    // property
	SetHandle(AValue HMENU)                                           // property
	Items(Index int32) IMenuItem                                      // property
	MergedItems() IMergedMenuItems                                    // property
	MenuIndex() int32                                                 // property
	SetMenuIndex(AValue int32)                                        // property
	Menu() IMenu                                                      // property
	Parent() IMenuItem                                                // property
	MergedParent() IMenuItem                                          // property
	Command() Word                                                    // property
	Action() IBasicAction                                             // property
	SetAction(AValue IBasicAction)                                    // property
	AutoCheck() bool                                                  // property
	SetAutoCheck(AValue bool)                                         // property
	Caption() string                                                  // property
	SetCaption(AValue string)                                         // property
	Checked() bool                                                    // property
	SetChecked(AValue bool)                                           // property
	Default() bool                                                    // property
	SetDefault(AValue bool)                                           // property
	Enabled() bool                                                    // property
	SetEnabled(AValue bool)                                           // property
	Bitmap() IBitmap                                                  // property
	SetBitmap(AValue IBitmap)                                         // property
	GroupIndex() Byte                                                 // property
	SetGroupIndex(AValue Byte)                                        // property
	GlyphShowMode() TGlyphShowMode                                    // property
	SetGlyphShowMode(AValue TGlyphShowMode)                           // property
	HelpContext() THelpContext                                        // property
	SetHelpContext(AValue THelpContext)                               // property
	Hint() string                                                     // property
	SetHint(AValue string)                                            // property
	ImageIndex() TImageIndex                                          // property
	SetImageIndex(AValue TImageIndex)                                 // property
	RadioItem() bool                                                  // property
	SetRadioItem(AValue bool)                                         // property
	RightJustify() bool                                               // property
	SetRightJustify(AValue bool)                                      // property
	ShortCut() TShortCut                                              // property
	SetShortCut(AValue TShortCut)                                     // property
	ShortCutKey2() TShortCut                                          // property
	SetShortCutKey2(AValue TShortCut)                                 // property
	ShowAlwaysCheckable() bool                                        // property
	SetShowAlwaysCheckable(AValue bool)                               // property
	SubMenuImages() ICustomImageList                                  // property
	SetSubMenuImages(AValue ICustomImageList)                         // property
	SubMenuImagesWidth() int32                                        // property
	SetSubMenuImagesWidth(AValue int32)                               // property
	Visible() bool                                                    // property
	SetVisible(AValue bool)                                           // property
	Find(ACaption string) IMenuItem                                   // function
	GetEnumeratorForMenuItemEnumerator() IMenuItemEnumerator          // function
	GetImageList() ICustomImageList                                   // function
	GetParentMenu() IMenu                                             // function
	GetMergedParentMenu() IMenu                                       // function
	GetIsRightToLeft() bool                                           // function
	HandleAllocated() bool                                            // function
	HasIcon() bool                                                    // function
	IndexOf(Item IMenuItem) int32                                     // function
	IndexOfCaption(ACaption string) int32                             // function
	VisibleIndexOf(Item IMenuItem) int32                              // function
	IsCheckItem() bool                                                // function
	IsLine() bool                                                     // function
	IsInMenuBar() bool                                                // function
	HasBitmap() bool                                                  // function
	GetIconSize(ADC HDC, DPI int32) (resultPoint TPoint)              // function
	MenuVisibleIndex() int32                                          // function
	GetImageList1(OutImages *ICustomImageList, OutImagesWidth *int32) // procedure
	InitiateAction()                                                  // procedure
	IntfDoSelect()                                                    // procedure
	InvalidateMergedItems()                                           // procedure
	Add(Item IMenuItem)                                               // procedure
	AddSeparator()                                                    // procedure
	Click()                                                           // procedure
	Delete(Index int32)                                               // procedure
	HandleNeeded()                                                    // procedure
	Insert(Index int32, Item IMenuItem)                               // procedure
	RecreateHandle()                                                  // procedure
	Remove(Item IMenuItem)                                            // procedure
	UpdateImage(forced bool)                                          // procedure
	UpdateImages(forced bool)                                         // procedure
	Clear()                                                           // procedure
	WriteDebugReport(Prefix string)                                   // procedure
	SetOnClick(fn TNotifyEvent)                                       // property event
	SetOnDrawItem(fn TMenuDrawItemEvent)                              // property event
	SetOnMeasureItem(fn TMenuMeasureItemEvent)                        // property event
}

// TMenuItem Parent: TLCLComponent
type TMenuItem struct {
	TLCLComponent
	clickPtr       uintptr
	drawItemPtr    uintptr
	measureItemPtr uintptr
}

func NewMenuItem(TheOwner IComponent) IMenuItem {
	r1 := menuItemImportAPI().SysCallN(12, GetObjectUintptr(TheOwner))
	return AsMenuItem(r1)
}

func (m *TMenuItem) Merged() IMenuItem {
	r1 := menuItemImportAPI().SysCallN(47, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) MergedWith() IMenuItem {
	r1 := menuItemImportAPI().SysCallN(50, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) Count() int32 {
	r1 := menuItemImportAPI().SysCallN(11, m.Instance())
	return int32(r1)
}

func (m *TMenuItem) Handle() HMENU {
	r1 := menuItemImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return HMENU(r1)
}

func (m *TMenuItem) SetHandle(AValue HMENU) {
	menuItemImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Items(Index int32) IMenuItem {
	r1 := menuItemImportAPI().SysCallN(43, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func (m *TMenuItem) MergedItems() IMergedMenuItems {
	r1 := menuItemImportAPI().SysCallN(48, m.Instance())
	return AsMergedMenuItems(r1)
}

func (m *TMenuItem) MenuIndex() int32 {
	r1 := menuItemImportAPI().SysCallN(45, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMenuItem) SetMenuIndex(AValue int32) {
	menuItemImportAPI().SysCallN(45, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Menu() IMenu {
	r1 := menuItemImportAPI().SysCallN(44, m.Instance())
	return AsMenu(r1)
}

func (m *TMenuItem) Parent() IMenuItem {
	r1 := menuItemImportAPI().SysCallN(51, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) MergedParent() IMenuItem {
	r1 := menuItemImportAPI().SysCallN(49, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) Command() Word {
	r1 := menuItemImportAPI().SysCallN(10, m.Instance())
	return Word(r1)
}

func (m *TMenuItem) Action() IBasicAction {
	r1 := menuItemImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsBasicAction(r1)
}

func (m *TMenuItem) SetAction(AValue IBasicAction) {
	menuItemImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenuItem) AutoCheck() bool {
	r1 := menuItemImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetAutoCheck(AValue bool) {
	menuItemImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Caption() string {
	r1 := menuItemImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMenuItem) SetCaption(AValue string) {
	menuItemImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMenuItem) Checked() bool {
	r1 := menuItemImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetChecked(AValue bool) {
	menuItemImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Default() bool {
	r1 := menuItemImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetDefault(AValue bool) {
	menuItemImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Enabled() bool {
	r1 := menuItemImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetEnabled(AValue bool) {
	menuItemImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Bitmap() IBitmap {
	r1 := menuItemImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TMenuItem) SetBitmap(AValue IBitmap) {
	menuItemImportAPI().SysCallN(4, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenuItem) GroupIndex() Byte {
	r1 := menuItemImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TMenuItem) SetGroupIndex(AValue Byte) {
	menuItemImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) GlyphShowMode() TGlyphShowMode {
	r1 := menuItemImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return TGlyphShowMode(r1)
}

func (m *TMenuItem) SetGlyphShowMode(AValue TGlyphShowMode) {
	menuItemImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) HelpContext() THelpContext {
	r1 := menuItemImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TMenuItem) SetHelpContext(AValue THelpContext) {
	menuItemImportAPI().SysCallN(31, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Hint() string {
	r1 := menuItemImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMenuItem) SetHint(AValue string) {
	menuItemImportAPI().SysCallN(32, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMenuItem) ImageIndex() TImageIndex {
	r1 := menuItemImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TMenuItem) SetImageIndex(AValue TImageIndex) {
	menuItemImportAPI().SysCallN(33, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) RadioItem() bool {
	r1 := menuItemImportAPI().SysCallN(52, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetRadioItem(AValue bool) {
	menuItemImportAPI().SysCallN(52, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) RightJustify() bool {
	r1 := menuItemImportAPI().SysCallN(55, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetRightJustify(AValue bool) {
	menuItemImportAPI().SysCallN(55, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) ShortCut() TShortCut {
	r1 := menuItemImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return TShortCut(r1)
}

func (m *TMenuItem) SetShortCut(AValue TShortCut) {
	menuItemImportAPI().SysCallN(59, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) ShortCutKey2() TShortCut {
	r1 := menuItemImportAPI().SysCallN(60, 0, m.Instance(), 0)
	return TShortCut(r1)
}

func (m *TMenuItem) SetShortCutKey2(AValue TShortCut) {
	menuItemImportAPI().SysCallN(60, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) ShowAlwaysCheckable() bool {
	r1 := menuItemImportAPI().SysCallN(61, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetShowAlwaysCheckable(AValue bool) {
	menuItemImportAPI().SysCallN(61, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) SubMenuImages() ICustomImageList {
	r1 := menuItemImportAPI().SysCallN(62, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TMenuItem) SetSubMenuImages(AValue ICustomImageList) {
	menuItemImportAPI().SysCallN(62, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenuItem) SubMenuImagesWidth() int32 {
	r1 := menuItemImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMenuItem) SetSubMenuImagesWidth(AValue int32) {
	menuItemImportAPI().SysCallN(63, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Visible() bool {
	r1 := menuItemImportAPI().SysCallN(66, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetVisible(AValue bool) {
	menuItemImportAPI().SysCallN(66, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Find(ACaption string) IMenuItem {
	r1 := menuItemImportAPI().SysCallN(16, m.Instance(), PascalStr(ACaption))
	return AsMenuItem(r1)
}

func (m *TMenuItem) GetEnumeratorForMenuItemEnumerator() IMenuItemEnumerator {
	r1 := menuItemImportAPI().SysCallN(17, m.Instance())
	return AsMenuItemEnumerator(r1)
}

func (m *TMenuItem) GetImageList() ICustomImageList {
	r1 := menuItemImportAPI().SysCallN(19, m.Instance())
	return AsCustomImageList(r1)
}

func (m *TMenuItem) GetParentMenu() IMenu {
	r1 := menuItemImportAPI().SysCallN(23, m.Instance())
	return AsMenu(r1)
}

func (m *TMenuItem) GetMergedParentMenu() IMenu {
	r1 := menuItemImportAPI().SysCallN(22, m.Instance())
	return AsMenu(r1)
}

func (m *TMenuItem) GetIsRightToLeft() bool {
	r1 := menuItemImportAPI().SysCallN(21, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) HandleAllocated() bool {
	r1 := menuItemImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) HasIcon() bool {
	r1 := menuItemImportAPI().SysCallN(30, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) IndexOf(Item IMenuItem) int32 {
	r1 := menuItemImportAPI().SysCallN(34, m.Instance(), GetObjectUintptr(Item))
	return int32(r1)
}

func (m *TMenuItem) IndexOfCaption(ACaption string) int32 {
	r1 := menuItemImportAPI().SysCallN(35, m.Instance(), PascalStr(ACaption))
	return int32(r1)
}

func (m *TMenuItem) VisibleIndexOf(Item IMenuItem) int32 {
	r1 := menuItemImportAPI().SysCallN(67, m.Instance(), GetObjectUintptr(Item))
	return int32(r1)
}

func (m *TMenuItem) IsCheckItem() bool {
	r1 := menuItemImportAPI().SysCallN(40, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) IsLine() bool {
	r1 := menuItemImportAPI().SysCallN(42, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) IsInMenuBar() bool {
	r1 := menuItemImportAPI().SysCallN(41, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) HasBitmap() bool {
	r1 := menuItemImportAPI().SysCallN(29, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) GetIconSize(ADC HDC, DPI int32) (resultPoint TPoint) {
	menuItemImportAPI().SysCallN(18, m.Instance(), uintptr(ADC), uintptr(DPI), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TMenuItem) MenuVisibleIndex() int32 {
	r1 := menuItemImportAPI().SysCallN(46, m.Instance())
	return int32(r1)
}

func MenuItemClass() TClass {
	ret := menuItemImportAPI().SysCallN(7)
	return TClass(ret)
}

func (m *TMenuItem) GetImageList1(OutImages *ICustomImageList, OutImagesWidth *int32) {
	var result0 uintptr
	var result1 uintptr
	menuItemImportAPI().SysCallN(20, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*OutImages = AsCustomImageList(result0)
	*OutImagesWidth = int32(result1)
}

func (m *TMenuItem) InitiateAction() {
	menuItemImportAPI().SysCallN(36, m.Instance())
}

func (m *TMenuItem) IntfDoSelect() {
	menuItemImportAPI().SysCallN(38, m.Instance())
}

func (m *TMenuItem) InvalidateMergedItems() {
	menuItemImportAPI().SysCallN(39, m.Instance())
}

func (m *TMenuItem) Add(Item IMenuItem) {
	menuItemImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Item))
}

func (m *TMenuItem) AddSeparator() {
	menuItemImportAPI().SysCallN(2, m.Instance())
}

func (m *TMenuItem) Click() {
	menuItemImportAPI().SysCallN(9, m.Instance())
}

func (m *TMenuItem) Delete(Index int32) {
	menuItemImportAPI().SysCallN(14, m.Instance(), uintptr(Index))
}

func (m *TMenuItem) HandleNeeded() {
	menuItemImportAPI().SysCallN(28, m.Instance())
}

func (m *TMenuItem) Insert(Index int32, Item IMenuItem) {
	menuItemImportAPI().SysCallN(37, m.Instance(), uintptr(Index), GetObjectUintptr(Item))
}

func (m *TMenuItem) RecreateHandle() {
	menuItemImportAPI().SysCallN(53, m.Instance())
}

func (m *TMenuItem) Remove(Item IMenuItem) {
	menuItemImportAPI().SysCallN(54, m.Instance(), GetObjectUintptr(Item))
}

func (m *TMenuItem) UpdateImage(forced bool) {
	menuItemImportAPI().SysCallN(64, m.Instance(), PascalBool(forced))
}

func (m *TMenuItem) UpdateImages(forced bool) {
	menuItemImportAPI().SysCallN(65, m.Instance(), PascalBool(forced))
}

func (m *TMenuItem) Clear() {
	menuItemImportAPI().SysCallN(8, m.Instance())
}

func (m *TMenuItem) WriteDebugReport(Prefix string) {
	menuItemImportAPI().SysCallN(68, m.Instance(), PascalStr(Prefix))
}

func (m *TMenuItem) SetOnClick(fn TNotifyEvent) {
	if m.clickPtr != 0 {
		RemoveEventElement(m.clickPtr)
	}
	m.clickPtr = MakeEventDataPtr(fn)
	menuItemImportAPI().SysCallN(56, m.Instance(), m.clickPtr)
}

func (m *TMenuItem) SetOnDrawItem(fn TMenuDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	menuItemImportAPI().SysCallN(57, m.Instance(), m.drawItemPtr)
}

func (m *TMenuItem) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	menuItemImportAPI().SysCallN(58, m.Instance(), m.measureItemPtr)
}

var (
	menuItemImport       *imports.Imports = nil
	menuItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MenuItem_Action", 0),
		/*1*/ imports.NewTable("MenuItem_Add", 0),
		/*2*/ imports.NewTable("MenuItem_AddSeparator", 0),
		/*3*/ imports.NewTable("MenuItem_AutoCheck", 0),
		/*4*/ imports.NewTable("MenuItem_Bitmap", 0),
		/*5*/ imports.NewTable("MenuItem_Caption", 0),
		/*6*/ imports.NewTable("MenuItem_Checked", 0),
		/*7*/ imports.NewTable("MenuItem_Class", 0),
		/*8*/ imports.NewTable("MenuItem_Clear", 0),
		/*9*/ imports.NewTable("MenuItem_Click", 0),
		/*10*/ imports.NewTable("MenuItem_Command", 0),
		/*11*/ imports.NewTable("MenuItem_Count", 0),
		/*12*/ imports.NewTable("MenuItem_Create", 0),
		/*13*/ imports.NewTable("MenuItem_Default", 0),
		/*14*/ imports.NewTable("MenuItem_Delete", 0),
		/*15*/ imports.NewTable("MenuItem_Enabled", 0),
		/*16*/ imports.NewTable("MenuItem_Find", 0),
		/*17*/ imports.NewTable("MenuItem_GetEnumeratorForMenuItemEnumerator", 0),
		/*18*/ imports.NewTable("MenuItem_GetIconSize", 0),
		/*19*/ imports.NewTable("MenuItem_GetImageList", 0),
		/*20*/ imports.NewTable("MenuItem_GetImageList1", 0),
		/*21*/ imports.NewTable("MenuItem_GetIsRightToLeft", 0),
		/*22*/ imports.NewTable("MenuItem_GetMergedParentMenu", 0),
		/*23*/ imports.NewTable("MenuItem_GetParentMenu", 0),
		/*24*/ imports.NewTable("MenuItem_GlyphShowMode", 0),
		/*25*/ imports.NewTable("MenuItem_GroupIndex", 0),
		/*26*/ imports.NewTable("MenuItem_Handle", 0),
		/*27*/ imports.NewTable("MenuItem_HandleAllocated", 0),
		/*28*/ imports.NewTable("MenuItem_HandleNeeded", 0),
		/*29*/ imports.NewTable("MenuItem_HasBitmap", 0),
		/*30*/ imports.NewTable("MenuItem_HasIcon", 0),
		/*31*/ imports.NewTable("MenuItem_HelpContext", 0),
		/*32*/ imports.NewTable("MenuItem_Hint", 0),
		/*33*/ imports.NewTable("MenuItem_ImageIndex", 0),
		/*34*/ imports.NewTable("MenuItem_IndexOf", 0),
		/*35*/ imports.NewTable("MenuItem_IndexOfCaption", 0),
		/*36*/ imports.NewTable("MenuItem_InitiateAction", 0),
		/*37*/ imports.NewTable("MenuItem_Insert", 0),
		/*38*/ imports.NewTable("MenuItem_IntfDoSelect", 0),
		/*39*/ imports.NewTable("MenuItem_InvalidateMergedItems", 0),
		/*40*/ imports.NewTable("MenuItem_IsCheckItem", 0),
		/*41*/ imports.NewTable("MenuItem_IsInMenuBar", 0),
		/*42*/ imports.NewTable("MenuItem_IsLine", 0),
		/*43*/ imports.NewTable("MenuItem_Items", 0),
		/*44*/ imports.NewTable("MenuItem_Menu", 0),
		/*45*/ imports.NewTable("MenuItem_MenuIndex", 0),
		/*46*/ imports.NewTable("MenuItem_MenuVisibleIndex", 0),
		/*47*/ imports.NewTable("MenuItem_Merged", 0),
		/*48*/ imports.NewTable("MenuItem_MergedItems", 0),
		/*49*/ imports.NewTable("MenuItem_MergedParent", 0),
		/*50*/ imports.NewTable("MenuItem_MergedWith", 0),
		/*51*/ imports.NewTable("MenuItem_Parent", 0),
		/*52*/ imports.NewTable("MenuItem_RadioItem", 0),
		/*53*/ imports.NewTable("MenuItem_RecreateHandle", 0),
		/*54*/ imports.NewTable("MenuItem_Remove", 0),
		/*55*/ imports.NewTable("MenuItem_RightJustify", 0),
		/*56*/ imports.NewTable("MenuItem_SetOnClick", 0),
		/*57*/ imports.NewTable("MenuItem_SetOnDrawItem", 0),
		/*58*/ imports.NewTable("MenuItem_SetOnMeasureItem", 0),
		/*59*/ imports.NewTable("MenuItem_ShortCut", 0),
		/*60*/ imports.NewTable("MenuItem_ShortCutKey2", 0),
		/*61*/ imports.NewTable("MenuItem_ShowAlwaysCheckable", 0),
		/*62*/ imports.NewTable("MenuItem_SubMenuImages", 0),
		/*63*/ imports.NewTable("MenuItem_SubMenuImagesWidth", 0),
		/*64*/ imports.NewTable("MenuItem_UpdateImage", 0),
		/*65*/ imports.NewTable("MenuItem_UpdateImages", 0),
		/*66*/ imports.NewTable("MenuItem_Visible", 0),
		/*67*/ imports.NewTable("MenuItem_VisibleIndexOf", 0),
		/*68*/ imports.NewTable("MenuItem_WriteDebugReport", 0),
	}
)

func menuItemImportAPI() *imports.Imports {
	if menuItemImport == nil {
		menuItemImport = NewDefaultImports()
		menuItemImport.SetImportTable(menuItemImportTables)
		menuItemImportTables = nil
	}
	return menuItemImport
}
