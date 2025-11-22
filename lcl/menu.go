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

// IMenu Parent: ILCLComponent
type IMenu interface {
	ILCLComponent
	FindItem(value uintptr, kind types.TFindItemKind) IMenuItem      // function
	GetHelpContext(value uintptr, byCommand bool) types.THelpContext // function
	IsShortcut(message *types.TLMKey) bool                           // function
	HandleAllocated() bool                                           // function
	IsRightToLeft() bool                                             // function
	UseRightToLeftAlignment() bool                                   // function
	UseRightToLeftReading() bool                                     // function
	DispatchCommand(command uint16) bool                             // function
	DestroyHandle()                                                  // procedure
	HandleNeeded()                                                   // procedure
	Handle() types.HMENU                                             // property Handle Getter
	Parent() IComponent                                              // property Parent Getter
	SetParent(value IComponent)                                      // property Parent Setter
	ShortcutHandled() bool                                           // property ShortcutHandled Getter
	SetShortcutHandled(value bool)                                   // property ShortcutHandled Setter
	AutoLineReduction() types.TMenuAutoFlag                          // property AutoLineReduction Getter
	SetAutoLineReduction(value types.TMenuAutoFlag)                  // property AutoLineReduction Setter
	BidiMode() types.TBiDiMode                                       // property BidiMode Getter
	SetBidiMode(value types.TBiDiMode)                               // property BidiMode Setter
	ParentBidiMode() bool                                            // property ParentBidiMode Getter
	SetParentBidiMode(value bool)                                    // property ParentBidiMode Setter
	Items() IMenuItem                                                // property Items Getter
	Images() ICustomImageList                                        // property Images Getter
	SetImages(value ICustomImageList)                                // property Images Setter
	ImagesWidth() int32                                              // property ImagesWidth Getter
	SetImagesWidth(value int32)                                      // property ImagesWidth Setter
	OwnerDraw() bool                                                 // property OwnerDraw Getter
	SetOwnerDraw(value bool)                                         // property OwnerDraw Setter
	SetOnDrawItem(fn TMenuDrawItemEvent)                             // property event
	SetOnMeasureItem(fn TMenuMeasureItemEvent)                       // property event
}

type TMenu struct {
	TLCLComponent
}

func (m *TMenu) FindItem(value uintptr, kind types.TFindItemKind) IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuAPI().SysCallN(1, m.Instance(), uintptr(value), uintptr(kind))
	return AsMenuItem(r)
}

func (m *TMenu) GetHelpContext(value uintptr, byCommand bool) types.THelpContext {
	if !m.IsValid() {
		return 0
	}
	r := menuAPI().SysCallN(2, m.Instance(), uintptr(value), api.PasBool(byCommand))
	return types.THelpContext(r)
}

func (m *TMenu) IsShortcut(message *types.TLMKey) bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(message)))
	return api.GoBool(r)
}

func (m *TMenu) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TMenu) IsRightToLeft() bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TMenu) UseRightToLeftAlignment() bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TMenu) UseRightToLeftReading() bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TMenu) DispatchCommand(command uint16) bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(8, m.Instance(), uintptr(command))
	return api.GoBool(r)
}

func (m *TMenu) DestroyHandle() {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(9, m.Instance())
}

func (m *TMenu) HandleNeeded() {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(10, m.Instance())
}

func (m *TMenu) Handle() types.HMENU {
	if !m.IsValid() {
		return 0
	}
	r := menuAPI().SysCallN(11, m.Instance())
	return types.HMENU(r)
}

func (m *TMenu) Parent() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := menuAPI().SysCallN(12, 0, m.Instance())
	return AsComponent(r)
}

func (m *TMenu) SetParent(value IComponent) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TMenu) ShortcutHandled() bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenu) SetShortcutHandled(value bool) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenu) AutoLineReduction() types.TMenuAutoFlag {
	if !m.IsValid() {
		return 0
	}
	r := menuAPI().SysCallN(14, 0, m.Instance())
	return types.TMenuAutoFlag(r)
}

func (m *TMenu) SetAutoLineReduction(value types.TMenuAutoFlag) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TMenu) BidiMode() types.TBiDiMode {
	if !m.IsValid() {
		return 0
	}
	r := menuAPI().SysCallN(15, 0, m.Instance())
	return types.TBiDiMode(r)
}

func (m *TMenu) SetBidiMode(value types.TBiDiMode) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TMenu) ParentBidiMode() bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenu) SetParentBidiMode(value bool) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenu) Items() IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuAPI().SysCallN(17, m.Instance())
	return AsMenuItem(r)
}

func (m *TMenu) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := menuAPI().SysCallN(18, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TMenu) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TMenu) ImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := menuAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TMenu) SetImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TMenu) OwnerDraw() bool {
	if !m.IsValid() {
		return false
	}
	r := menuAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMenu) SetOwnerDraw(value bool) {
	if !m.IsValid() {
		return
	}
	menuAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TMenu) SetOnDrawItem(fn TMenuDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMenuDrawItemEvent(fn)
	base.SetEvent(m, 21, menuAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMenu) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMenuMeasureItemEvent(fn)
	base.SetEvent(m, 22, menuAPI(), api.MakeEventDataPtr(cb))
}

// NewMenu class constructor
func NewMenu(owner IComponent) IMenu {
	r := menuAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsMenu(r)
}

func TMenuClass() types.TClass {
	r := menuAPI().SysCallN(23)
	return types.TClass(r)
}

var (
	menuOnce   base.Once
	menuImport *imports.Imports = nil
)

func menuAPI() *imports.Imports {
	menuOnce.Do(func() {
		menuImport = api.NewDefaultImports()
		menuImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMenu_Create", 0), // constructor NewMenu
			/* 1 */ imports.NewTable("TMenu_FindItem", 0), // function FindItem
			/* 2 */ imports.NewTable("TMenu_GetHelpContext", 0), // function GetHelpContext
			/* 3 */ imports.NewTable("TMenu_IsShortcut", 0), // function IsShortcut
			/* 4 */ imports.NewTable("TMenu_HandleAllocated", 0), // function HandleAllocated
			/* 5 */ imports.NewTable("TMenu_IsRightToLeft", 0), // function IsRightToLeft
			/* 6 */ imports.NewTable("TMenu_UseRightToLeftAlignment", 0), // function UseRightToLeftAlignment
			/* 7 */ imports.NewTable("TMenu_UseRightToLeftReading", 0), // function UseRightToLeftReading
			/* 8 */ imports.NewTable("TMenu_DispatchCommand", 0), // function DispatchCommand
			/* 9 */ imports.NewTable("TMenu_DestroyHandle", 0), // procedure DestroyHandle
			/* 10 */ imports.NewTable("TMenu_HandleNeeded", 0), // procedure HandleNeeded
			/* 11 */ imports.NewTable("TMenu_Handle", 0), // property Handle
			/* 12 */ imports.NewTable("TMenu_Parent", 0), // property Parent
			/* 13 */ imports.NewTable("TMenu_ShortcutHandled", 0), // property ShortcutHandled
			/* 14 */ imports.NewTable("TMenu_AutoLineReduction", 0), // property AutoLineReduction
			/* 15 */ imports.NewTable("TMenu_BidiMode", 0), // property BidiMode
			/* 16 */ imports.NewTable("TMenu_ParentBidiMode", 0), // property ParentBidiMode
			/* 17 */ imports.NewTable("TMenu_Items", 0), // property Items
			/* 18 */ imports.NewTable("TMenu_Images", 0), // property Images
			/* 19 */ imports.NewTable("TMenu_ImagesWidth", 0), // property ImagesWidth
			/* 20 */ imports.NewTable("TMenu_OwnerDraw", 0), // property OwnerDraw
			/* 21 */ imports.NewTable("TMenu_OnDrawItem", 0), // event OnDrawItem
			/* 22 */ imports.NewTable("TMenu_OnMeasureItem", 0), // event OnMeasureItem
			/* 23 */ imports.NewTable("TMenu_TClass", 0), // function TMenuClass
		}
	})
	return menuImport
}
