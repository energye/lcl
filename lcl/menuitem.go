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

// IMenuItem Parent: ILCLComponent
type IMenuItem interface {
	ILCLComponent
	Find(caption string) IMenuItem                                                         // function
	GetEnumeratorToMenuItemEnumerator() IMenuItemEnumerator                                // function
	GetImageListToCustomImageList() ICustomImageList                                       // function
	GetParentMenu() IMenu                                                                  // function
	GetMergedParentMenu() IMenu                                                            // function
	GetIsRightToLeft() bool                                                                // function
	HandleAllocated() bool                                                                 // function
	HasIcon() bool                                                                         // function
	IndexOf(item IMenuItem) int32                                                          // function
	IndexOfCaption(caption string) int32                                                   // function
	VisibleIndexOf(item IMenuItem) int32                                                   // function
	IsCheckItem() bool                                                                     // function
	IsLine() bool                                                                          // function
	IsInMenuBar() bool                                                                     // function
	HasBitmap() bool                                                                       // function
	GetIconSize(dC types.HDC, dPI int32) types.TPoint                                      // function
	RethinkLines() bool                                                                    // function
	MenuVisibleIndex() int32                                                               // function
	GetImageListWithCustomImageListInt(outImages *ICustomImageList, outImagesWidth *int32) // procedure
	InitiateAction()                                                                       // procedure
	IntfDoSelect()                                                                         // procedure
	InvalidateMergedItems()                                                                // procedure
	Add(item IMenuItem)                                                                    // procedure
	AddSeparator()                                                                         // procedure
	Click()                                                                                // procedure
	Delete(index int32)                                                                    // procedure
	HandleNeeded()                                                                         // procedure
	Insert(index int32, item IMenuItem)                                                    // procedure
	RecreateHandle()                                                                       // procedure
	Remove(item IMenuItem)                                                                 // procedure
	UpdateImage(forced bool)                                                               // procedure
	UpdateImages(forced bool)                                                              // procedure
	Clear()                                                                                // procedure
	AddHandler(handlerType types.TMenuItemHandlerType, method TMethod, asFirst bool)       // procedure
	RemoveHandler(handlerType types.TMenuItemHandlerType, method TMethod)                  // procedure
	WriteDebugReport(prefix string)                                                        // procedure
	Merged() IMenuItem                                                                     // property Merged Getter
	MergedWith() IMenuItem                                                                 // property MergedWith Getter
	Count() int32                                                                          // property Count Getter
	Handle() types.HMENU                                                                   // property Handle Getter
	SetHandle(value types.HMENU)                                                           // property Handle Setter
	Items(index int32) IMenuItem                                                           // property Items Getter
	MergedItems() IMergedMenuItems                                                         // property MergedItems Getter
	MenuIndex() int32                                                                      // property MenuIndex Getter
	SetMenuIndex(value int32)                                                              // property MenuIndex Setter
	Menu() IMenu                                                                           // property Menu Getter
	Parent() IMenuItem                                                                     // property Parent Getter
	MergedParent() IMenuItem                                                               // property MergedParent Getter
	Command() uint16                                                                       // property Command Getter
	Action() IBasicAction                                                                  // property Action Getter
	SetAction(value IBasicAction)                                                          // property Action Setter
	AutoCheck() bool                                                                       // property AutoCheck Getter
	SetAutoCheck(value bool)                                                               // property AutoCheck Setter
	AutoLineReduction() types.TMenuItemAutoFlag                                            // property AutoLineReduction Getter
	SetAutoLineReduction(value types.TMenuItemAutoFlag)                                    // property AutoLineReduction Setter
	Caption() string                                                                       // property Caption Getter
	SetCaption(value string)                                                               // property Caption Setter
	Checked() bool                                                                         // property Checked Getter
	SetChecked(value bool)                                                                 // property Checked Setter
	Default() bool                                                                         // property Default Getter
	SetDefault(value bool)                                                                 // property Default Setter
	Enabled() bool                                                                         // property Enabled Getter
	SetEnabled(value bool)                                                                 // property Enabled Setter
	Bitmap() IBitmap                                                                       // property Bitmap Getter
	SetBitmap(value IBitmap)                                                               // property Bitmap Setter
	GroupIndex() byte                                                                      // property GroupIndex Getter
	SetGroupIndex(value byte)                                                              // property GroupIndex Setter
	GlyphShowMode() types.TGlyphShowMode                                                   // property GlyphShowMode Getter
	SetGlyphShowMode(value types.TGlyphShowMode)                                           // property GlyphShowMode Setter
	HelpContext() types.THelpContext                                                       // property HelpContext Getter
	SetHelpContext(value types.THelpContext)                                               // property HelpContext Setter
	Hint() string                                                                          // property Hint Getter
	SetHint(value string)                                                                  // property Hint Setter
	ImageIndex() int32                                                                     // property ImageIndex Getter
	SetImageIndex(value int32)                                                             // property ImageIndex Setter
	RadioItem() bool                                                                       // property RadioItem Getter
	SetRadioItem(value bool)                                                               // property RadioItem Setter
	RightJustify() bool                                                                    // property RightJustify Getter
	SetRightJustify(value bool)                                                            // property RightJustify Setter
	ShortCut() types.TShortCut                                                             // property ShortCut Getter
	SetShortCut(value types.TShortCut)                                                     // property ShortCut Setter
	ShortCutKey2() types.TShortCut                                                         // property ShortCutKey2 Getter
	SetShortCutKey2(value types.TShortCut)                                                 // property ShortCutKey2 Setter
	ShowAlwaysCheckable() bool                                                             // property ShowAlwaysCheckable Getter
	SetShowAlwaysCheckable(value bool)                                                     // property ShowAlwaysCheckable Setter
	SubMenuImages() ICustomImageList                                                       // property SubMenuImages Getter
	SetSubMenuImages(value ICustomImageList)                                               // property SubMenuImages Setter
	SubMenuImagesWidth() int32                                                             // property SubMenuImagesWidth Getter
	SetSubMenuImagesWidth(value int32)                                                     // property SubMenuImagesWidth Setter
	Visible() bool                                                                         // property Visible Getter
	SetVisible(value bool)                                                                 // property Visible Setter
	SetOnClick(fn TNotifyEvent)                                                            // property event
	SetOnDrawItem(fn TMenuDrawItemEvent)                                                   // property event
	SetOnMeasureItem(fn TMenuMeasureItemEvent)                                             // property event
}

type TMenuItem struct {
	TLCLComponent
}

func (m *TMenuItem) Find(caption string) IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(1, m.Instance(), api.PasStr(caption))
	return AsMenuItem(r)
}

func (m *TMenuItem) GetEnumeratorToMenuItemEnumerator() IMenuItemEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(2, m.Instance())
	return AsMenuItemEnumerator(r)
}

func (m *TMenuItem) GetImageListToCustomImageList() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(3, m.Instance())
	return AsCustomImageList(r)
}

func (m *TMenuItem) GetParentMenu() IMenu {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(4, m.Instance())
	return AsMenu(r)
}

func (m *TMenuItem) GetMergedParentMenu() IMenu {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(5, m.Instance())
	return AsMenu(r)
}

func (m *TMenuItem) GetIsRightToLeft() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) HasIcon() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) IndexOf(item IMenuItem) int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TMenuItem) IndexOfCaption(caption string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(10, m.Instance(), api.PasStr(caption))
	return int32(r)
}

func (m *TMenuItem) VisibleIndexOf(item IMenuItem) int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TMenuItem) IsCheckItem() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) IsLine() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) IsInMenuBar() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) HasBitmap() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(15, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) GetIconSize(dC types.HDC, dPI int32) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(16, m.Instance(), uintptr(dC), uintptr(dPI), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TMenuItem) RethinkLines() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) MenuVisibleIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(18, m.Instance())
	return int32(r)
}

func (m *TMenuItem) GetImageListWithCustomImageListInt(outImages *ICustomImageList, outImagesWidth *int32) {
	if !m.IsValid() {
		return
	}
	var imagesPtr uintptr
	var imagesWidthPtr uintptr
	menuItemAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&imagesPtr)), uintptr(base.UnsafePointer(&imagesWidthPtr)))
	*outImages = AsCustomImageList(imagesPtr)
	*outImagesWidth = int32(imagesWidthPtr)
}

func (m *TMenuItem) InitiateAction() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(20, m.Instance())
}

func (m *TMenuItem) IntfDoSelect() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(21, m.Instance())
}

func (m *TMenuItem) InvalidateMergedItems() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(22, m.Instance())
}

func (m *TMenuItem) Add(item IMenuItem) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(23, m.Instance(), base.GetObjectUintptr(item))
}

func (m *TMenuItem) AddSeparator() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(24, m.Instance())
}

func (m *TMenuItem) Click() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(25, m.Instance())
}

func (m *TMenuItem) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(26, m.Instance(), uintptr(index))
}

func (m *TMenuItem) HandleNeeded() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(27, m.Instance())
}

func (m *TMenuItem) Insert(index int32, item IMenuItem) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(28, m.Instance(), uintptr(index), base.GetObjectUintptr(item))
}

func (m *TMenuItem) RecreateHandle() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(29, m.Instance())
}

func (m *TMenuItem) Remove(item IMenuItem) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(30, m.Instance(), base.GetObjectUintptr(item))
}

func (m *TMenuItem) UpdateImage(forced bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(31, m.Instance(), api.PasBool(forced))
}

func (m *TMenuItem) UpdateImages(forced bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(32, m.Instance(), api.PasBool(forced))
}

func (m *TMenuItem) Clear() {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(33, m.Instance())
}

func (m *TMenuItem) AddHandler(handlerType types.TMenuItemHandlerType, method TMethod, asFirst bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(34, m.Instance(), uintptr(handlerType), uintptr(base.UnsafePointer(&method)), api.PasBool(asFirst))
}

func (m *TMenuItem) RemoveHandler(handlerType types.TMenuItemHandlerType, method TMethod) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(35, m.Instance(), uintptr(handlerType), uintptr(base.UnsafePointer(&method)))
}

func (m *TMenuItem) WriteDebugReport(prefix string) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(36, m.Instance(), api.PasStr(prefix))
}

func (m *TMenuItem) Merged() IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(37, m.Instance())
	return AsMenuItem(r)
}

func (m *TMenuItem) MergedWith() IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(38, m.Instance())
	return AsMenuItem(r)
}

func (m *TMenuItem) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(39, m.Instance())
	return int32(r)
}

func (m *TMenuItem) Handle() types.HMENU {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(40, 0, m.Instance())
	return types.HMENU(r)
}

func (m *TMenuItem) SetHandle(value types.HMENU) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) Items(index int32) IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(41, m.Instance(), uintptr(index))
	return AsMenuItem(r)
}

func (m *TMenuItem) MergedItems() IMergedMenuItems {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(42, m.Instance())
	return AsMergedMenuItems(r)
}

func (m *TMenuItem) MenuIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(43, 0, m.Instance())
	return int32(r)
}

func (m *TMenuItem) SetMenuIndex(value int32) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(43, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) Menu() IMenu {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(44, m.Instance())
	return AsMenu(r)
}

func (m *TMenuItem) Parent() IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(45, m.Instance())
	return AsMenuItem(r)
}

func (m *TMenuItem) MergedParent() IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(46, m.Instance())
	return AsMenuItem(r)
}

func (m *TMenuItem) Command() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(47, m.Instance())
	return uint16(r)
}

func (m *TMenuItem) Action() IBasicAction {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(48, 0, m.Instance())
	return AsBasicAction(r)
}

func (m *TMenuItem) SetAction(value IBasicAction) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(48, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TMenuItem) AutoCheck() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(49, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetAutoCheck(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(49, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) AutoLineReduction() types.TMenuItemAutoFlag {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(50, 0, m.Instance())
	return types.TMenuItemAutoFlag(r)
}

func (m *TMenuItem) SetAutoLineReduction(value types.TMenuItemAutoFlag) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(50, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := menuItemAPI().SysCallN(51, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TMenuItem) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(51, 1, m.Instance(), api.PasStr(value))
}

func (m *TMenuItem) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(52, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(52, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) Default() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(53, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetDefault(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(53, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(54, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(54, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) Bitmap() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(55, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TMenuItem) SetBitmap(value IBitmap) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(55, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TMenuItem) GroupIndex() byte {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(56, 0, m.Instance())
	return byte(r)
}

func (m *TMenuItem) SetGroupIndex(value byte) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(56, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) GlyphShowMode() types.TGlyphShowMode {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(57, 0, m.Instance())
	return types.TGlyphShowMode(r)
}

func (m *TMenuItem) SetGlyphShowMode(value types.TGlyphShowMode) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(57, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) HelpContext() types.THelpContext {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(58, 0, m.Instance())
	return types.THelpContext(r)
}

func (m *TMenuItem) SetHelpContext(value types.THelpContext) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(58, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) Hint() string {
	if !m.IsValid() {
		return ""
	}
	r := menuItemAPI().SysCallN(59, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TMenuItem) SetHint(value string) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(59, 1, m.Instance(), api.PasStr(value))
}

func (m *TMenuItem) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(60, 0, m.Instance())
	return int32(r)
}

func (m *TMenuItem) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(60, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) RadioItem() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(61, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetRadioItem(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(61, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) RightJustify() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(62, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetRightJustify(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(62, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) ShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(63, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TMenuItem) SetShortCut(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(63, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) ShortCutKey2() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(64, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TMenuItem) SetShortCutKey2(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(64, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) ShowAlwaysCheckable() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(65, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetShowAlwaysCheckable(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(65, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) SubMenuImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := menuItemAPI().SysCallN(66, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TMenuItem) SetSubMenuImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(66, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TMenuItem) SubMenuImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuItemAPI().SysCallN(67, 0, m.Instance())
	return int32(r)
}

func (m *TMenuItem) SetSubMenuImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(67, 1, m.Instance(), uintptr(value))
}

func (m *TMenuItem) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemAPI().SysCallN(68, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItem) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	menuItemAPI().SysCallN(68, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenuItem) SetOnClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 69, menuItemAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMenuItem) SetOnDrawItem(fn TMenuDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMenuDrawItemEvent(fn)
	base.SetEvent(m, 70, menuItemAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMenuItem) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMenuMeasureItemEvent(fn)
	base.SetEvent(m, 71, menuItemAPI(), api.MakeEventDataPtr(cb))
}

// NewMenuItem class constructor
func NewMenuItem(theOwner IComponent) IMenuItem {
	r := menuItemAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsMenuItem(r)
}

func TMenuItemClass() types.TClass {
	r := menuItemAPI().SysCallN(72)
	return types.TClass(r)
}

var (
	menuItemOnce   base.Once
	menuItemImport *imports.Imports = nil
)

func menuItemAPI() *imports.Imports {
	menuItemOnce.Do(func() {
		menuItemImport = api.NewDefaultImports()
		menuItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMenuItem_Create", 0), // constructor NewMenuItem
			/* 1 */ imports.NewTable("TMenuItem_Find", 0), // function Find
			/* 2 */ imports.NewTable("TMenuItem_GetEnumeratorToMenuItemEnumerator", 0), // function GetEnumeratorToMenuItemEnumerator
			/* 3 */ imports.NewTable("TMenuItem_GetImageListToCustomImageList", 0), // function GetImageListToCustomImageList
			/* 4 */ imports.NewTable("TMenuItem_GetParentMenu", 0), // function GetParentMenu
			/* 5 */ imports.NewTable("TMenuItem_GetMergedParentMenu", 0), // function GetMergedParentMenu
			/* 6 */ imports.NewTable("TMenuItem_GetIsRightToLeft", 0), // function GetIsRightToLeft
			/* 7 */ imports.NewTable("TMenuItem_HandleAllocated", 0), // function HandleAllocated
			/* 8 */ imports.NewTable("TMenuItem_HasIcon", 0), // function HasIcon
			/* 9 */ imports.NewTable("TMenuItem_IndexOf", 0), // function IndexOf
			/* 10 */ imports.NewTable("TMenuItem_IndexOfCaption", 0), // function IndexOfCaption
			/* 11 */ imports.NewTable("TMenuItem_VisibleIndexOf", 0), // function VisibleIndexOf
			/* 12 */ imports.NewTable("TMenuItem_IsCheckItem", 0), // function IsCheckItem
			/* 13 */ imports.NewTable("TMenuItem_IsLine", 0), // function IsLine
			/* 14 */ imports.NewTable("TMenuItem_IsInMenuBar", 0), // function IsInMenuBar
			/* 15 */ imports.NewTable("TMenuItem_HasBitmap", 0), // function HasBitmap
			/* 16 */ imports.NewTable("TMenuItem_GetIconSize", 0), // function GetIconSize
			/* 17 */ imports.NewTable("TMenuItem_RethinkLines", 0), // function RethinkLines
			/* 18 */ imports.NewTable("TMenuItem_MenuVisibleIndex", 0), // function MenuVisibleIndex
			/* 19 */ imports.NewTable("TMenuItem_GetImageListWithCustomImageListInt", 0), // procedure GetImageListWithCustomImageListInt
			/* 20 */ imports.NewTable("TMenuItem_InitiateAction", 0), // procedure InitiateAction
			/* 21 */ imports.NewTable("TMenuItem_IntfDoSelect", 0), // procedure IntfDoSelect
			/* 22 */ imports.NewTable("TMenuItem_InvalidateMergedItems", 0), // procedure InvalidateMergedItems
			/* 23 */ imports.NewTable("TMenuItem_Add", 0), // procedure Add
			/* 24 */ imports.NewTable("TMenuItem_AddSeparator", 0), // procedure AddSeparator
			/* 25 */ imports.NewTable("TMenuItem_Click", 0), // procedure Click
			/* 26 */ imports.NewTable("TMenuItem_Delete", 0), // procedure Delete
			/* 27 */ imports.NewTable("TMenuItem_HandleNeeded", 0), // procedure HandleNeeded
			/* 28 */ imports.NewTable("TMenuItem_Insert", 0), // procedure Insert
			/* 29 */ imports.NewTable("TMenuItem_RecreateHandle", 0), // procedure RecreateHandle
			/* 30 */ imports.NewTable("TMenuItem_Remove", 0), // procedure Remove
			/* 31 */ imports.NewTable("TMenuItem_UpdateImage", 0), // procedure UpdateImage
			/* 32 */ imports.NewTable("TMenuItem_UpdateImages", 0), // procedure UpdateImages
			/* 33 */ imports.NewTable("TMenuItem_Clear", 0), // procedure Clear
			/* 34 */ imports.NewTable("TMenuItem_AddHandler", 0), // procedure AddHandler
			/* 35 */ imports.NewTable("TMenuItem_RemoveHandler", 0), // procedure RemoveHandler
			/* 36 */ imports.NewTable("TMenuItem_WriteDebugReport", 0), // procedure WriteDebugReport
			/* 37 */ imports.NewTable("TMenuItem_Merged", 0), // property Merged
			/* 38 */ imports.NewTable("TMenuItem_MergedWith", 0), // property MergedWith
			/* 39 */ imports.NewTable("TMenuItem_Count", 0), // property Count
			/* 40 */ imports.NewTable("TMenuItem_Handle", 0), // property Handle
			/* 41 */ imports.NewTable("TMenuItem_Items", 0), // property Items
			/* 42 */ imports.NewTable("TMenuItem_MergedItems", 0), // property MergedItems
			/* 43 */ imports.NewTable("TMenuItem_MenuIndex", 0), // property MenuIndex
			/* 44 */ imports.NewTable("TMenuItem_Menu", 0), // property Menu
			/* 45 */ imports.NewTable("TMenuItem_Parent", 0), // property Parent
			/* 46 */ imports.NewTable("TMenuItem_MergedParent", 0), // property MergedParent
			/* 47 */ imports.NewTable("TMenuItem_Command", 0), // property Command
			/* 48 */ imports.NewTable("TMenuItem_Action", 0), // property Action
			/* 49 */ imports.NewTable("TMenuItem_AutoCheck", 0), // property AutoCheck
			/* 50 */ imports.NewTable("TMenuItem_AutoLineReduction", 0), // property AutoLineReduction
			/* 51 */ imports.NewTable("TMenuItem_Caption", 0), // property Caption
			/* 52 */ imports.NewTable("TMenuItem_Checked", 0), // property Checked
			/* 53 */ imports.NewTable("TMenuItem_Default", 0), // property Default
			/* 54 */ imports.NewTable("TMenuItem_Enabled", 0), // property Enabled
			/* 55 */ imports.NewTable("TMenuItem_Bitmap", 0), // property Bitmap
			/* 56 */ imports.NewTable("TMenuItem_GroupIndex", 0), // property GroupIndex
			/* 57 */ imports.NewTable("TMenuItem_GlyphShowMode", 0), // property GlyphShowMode
			/* 58 */ imports.NewTable("TMenuItem_HelpContext", 0), // property HelpContext
			/* 59 */ imports.NewTable("TMenuItem_Hint", 0), // property Hint
			/* 60 */ imports.NewTable("TMenuItem_ImageIndex", 0), // property ImageIndex
			/* 61 */ imports.NewTable("TMenuItem_RadioItem", 0), // property RadioItem
			/* 62 */ imports.NewTable("TMenuItem_RightJustify", 0), // property RightJustify
			/* 63 */ imports.NewTable("TMenuItem_ShortCut", 0), // property ShortCut
			/* 64 */ imports.NewTable("TMenuItem_ShortCutKey2", 0), // property ShortCutKey2
			/* 65 */ imports.NewTable("TMenuItem_ShowAlwaysCheckable", 0), // property ShowAlwaysCheckable
			/* 66 */ imports.NewTable("TMenuItem_SubMenuImages", 0), // property SubMenuImages
			/* 67 */ imports.NewTable("TMenuItem_SubMenuImagesWidth", 0), // property SubMenuImagesWidth
			/* 68 */ imports.NewTable("TMenuItem_Visible", 0), // property Visible
			/* 69 */ imports.NewTable("TMenuItem_OnClick", 0), // event OnClick
			/* 70 */ imports.NewTable("TMenuItem_OnDrawItem", 0), // event OnDrawItem
			/* 71 */ imports.NewTable("TMenuItem_OnMeasureItem", 0), // event OnMeasureItem
			/* 72 */ imports.NewTable("TMenuItem_TClass", 0), // function TMenuItemClass
		}
	})
	return menuItemImport
}
