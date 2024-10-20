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

// IMenu Parent: ILCLComponent
type IMenu interface {
	ILCLComponent
	Handle() HMENU                                             // property
	Parent() IComponent                                        // property
	SetParent(AValue IComponent)                               // property
	ShortcutHandled() bool                                     // property
	SetShortcutHandled(AValue bool)                            // property
	BidiMode() TBiDiMode                                       // property
	SetBidiMode(AValue TBiDiMode)                              // property
	ParentBidiMode() bool                                      // property
	SetParentBidiMode(AValue bool)                             // property
	Items() IMenuItem                                          // property
	Images() ICustomImageList                                  // property
	SetImages(AValue ICustomImageList)                         // property
	ImagesWidth() int32                                        // property
	SetImagesWidth(AValue int32)                               // property
	OwnerDraw() bool                                           // property
	SetOwnerDraw(AValue bool)                                  // property
	FindItem(AValue uint32, Kind TFindItemKind) IMenuItem      // function
	GetHelpContext(AValue uint32, ByCommand bool) THelpContext // function
	IsShortcut(Message *TLMKey) bool                           // function
	HandleAllocated() bool                                     // function
	IsRightToLeft() bool                                       // function
	UseRightToLeftAlignment() bool                             // function
	UseRightToLeftReading() bool                               // function
	DispatchCommand(ACommand Word) bool                        // function
	DestroyHandle()                                            // procedure
	HandleNeeded()                                             // procedure
	SetOnDrawItem(fn TMenuDrawItemEvent)                       // property event
	SetOnMeasureItem(fn TMenuMeasureItemEvent)                 // property event
}

// TMenu Parent: TLCLComponent
type TMenu struct {
	TLCLComponent
	drawItemPtr    uintptr
	measureItemPtr uintptr
}

func NewMenu(AOwner IComponent) IMenu {
	r1 := menuImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsMenu(r1)
}

func (m *TMenu) Handle() HMENU {
	r1 := menuImportAPI().SysCallN(7, m.Instance())
	return HMENU(r1)
}

func (m *TMenu) Parent() IComponent {
	r1 := menuImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TMenu) SetParent(AValue IComponent) {
	menuImportAPI().SysCallN(16, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenu) ShortcutHandled() bool {
	r1 := menuImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetShortcutHandled(AValue bool) {
	menuImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) BidiMode() TBiDiMode {
	r1 := menuImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TMenu) SetBidiMode(AValue TBiDiMode) {
	menuImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenu) ParentBidiMode() bool {
	r1 := menuImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetParentBidiMode(AValue bool) {
	menuImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) Items() IMenuItem {
	r1 := menuImportAPI().SysCallN(14, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenu) Images() ICustomImageList {
	r1 := menuImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TMenu) SetImages(AValue ICustomImageList) {
	menuImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenu) ImagesWidth() int32 {
	r1 := menuImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMenu) SetImagesWidth(AValue int32) {
	menuImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenu) OwnerDraw() bool {
	r1 := menuImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetOwnerDraw(AValue bool) {
	menuImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) FindItem(AValue uint32, Kind TFindItemKind) IMenuItem {
	r1 := menuImportAPI().SysCallN(5, m.Instance(), uintptr(AValue), uintptr(Kind))
	return AsMenuItem(r1)
}

func (m *TMenu) GetHelpContext(AValue uint32, ByCommand bool) THelpContext {
	r1 := menuImportAPI().SysCallN(6, m.Instance(), uintptr(AValue), PascalBool(ByCommand))
	return THelpContext(r1)
}

func (m *TMenu) IsShortcut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := menuImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func (m *TMenu) HandleAllocated() bool {
	r1 := menuImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) IsRightToLeft() bool {
	r1 := menuImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) UseRightToLeftAlignment() bool {
	r1 := menuImportAPI().SysCallN(21, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) UseRightToLeftReading() bool {
	r1 := menuImportAPI().SysCallN(22, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) DispatchCommand(ACommand Word) bool {
	r1 := menuImportAPI().SysCallN(4, m.Instance(), uintptr(ACommand))
	return GoBool(r1)
}

func MenuClass() TClass {
	ret := menuImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TMenu) DestroyHandle() {
	menuImportAPI().SysCallN(3, m.Instance())
}

func (m *TMenu) HandleNeeded() {
	menuImportAPI().SysCallN(9, m.Instance())
}

func (m *TMenu) SetOnDrawItem(fn TMenuDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	menuImportAPI().SysCallN(18, m.Instance(), m.drawItemPtr)
}

func (m *TMenu) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	menuImportAPI().SysCallN(19, m.Instance(), m.measureItemPtr)
}

var (
	menuImport       *imports.Imports = nil
	menuImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Menu_BidiMode", 0),
		/*1*/ imports.NewTable("Menu_Class", 0),
		/*2*/ imports.NewTable("Menu_Create", 0),
		/*3*/ imports.NewTable("Menu_DestroyHandle", 0),
		/*4*/ imports.NewTable("Menu_DispatchCommand", 0),
		/*5*/ imports.NewTable("Menu_FindItem", 0),
		/*6*/ imports.NewTable("Menu_GetHelpContext", 0),
		/*7*/ imports.NewTable("Menu_Handle", 0),
		/*8*/ imports.NewTable("Menu_HandleAllocated", 0),
		/*9*/ imports.NewTable("Menu_HandleNeeded", 0),
		/*10*/ imports.NewTable("Menu_Images", 0),
		/*11*/ imports.NewTable("Menu_ImagesWidth", 0),
		/*12*/ imports.NewTable("Menu_IsRightToLeft", 0),
		/*13*/ imports.NewTable("Menu_IsShortcut", 0),
		/*14*/ imports.NewTable("Menu_Items", 0),
		/*15*/ imports.NewTable("Menu_OwnerDraw", 0),
		/*16*/ imports.NewTable("Menu_Parent", 0),
		/*17*/ imports.NewTable("Menu_ParentBidiMode", 0),
		/*18*/ imports.NewTable("Menu_SetOnDrawItem", 0),
		/*19*/ imports.NewTable("Menu_SetOnMeasureItem", 0),
		/*20*/ imports.NewTable("Menu_ShortcutHandled", 0),
		/*21*/ imports.NewTable("Menu_UseRightToLeftAlignment", 0),
		/*22*/ imports.NewTable("Menu_UseRightToLeftReading", 0),
	}
)

func menuImportAPI() *imports.Imports {
	if menuImport == nil {
		menuImport = NewDefaultImports()
		menuImport.SetImportTable(menuImportTables)
		menuImportTables = nil
	}
	return menuImport
}
