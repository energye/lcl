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
	r1 := LCL().SysCallN(4340, GetObjectUintptr(AOwner))
	return AsMenu(r1)
}

func (m *TMenu) Handle() HMENU {
	r1 := LCL().SysCallN(4345, m.Instance())
	return HMENU(r1)
}

func (m *TMenu) Parent() IComponent {
	r1 := LCL().SysCallN(4354, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TMenu) SetParent(AValue IComponent) {
	LCL().SysCallN(4354, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenu) ShortcutHandled() bool {
	r1 := LCL().SysCallN(4358, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetShortcutHandled(AValue bool) {
	LCL().SysCallN(4358, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) BidiMode() TBiDiMode {
	r1 := LCL().SysCallN(4338, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TMenu) SetBidiMode(AValue TBiDiMode) {
	LCL().SysCallN(4338, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenu) ParentBidiMode() bool {
	r1 := LCL().SysCallN(4355, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetParentBidiMode(AValue bool) {
	LCL().SysCallN(4355, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) Items() IMenuItem {
	r1 := LCL().SysCallN(4352, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenu) Images() ICustomImageList {
	r1 := LCL().SysCallN(4348, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TMenu) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(4348, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenu) ImagesWidth() int32 {
	r1 := LCL().SysCallN(4349, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMenu) SetImagesWidth(AValue int32) {
	LCL().SysCallN(4349, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenu) OwnerDraw() bool {
	r1 := LCL().SysCallN(4353, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenu) SetOwnerDraw(AValue bool) {
	LCL().SysCallN(4353, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenu) FindItem(AValue uint32, Kind TFindItemKind) IMenuItem {
	r1 := LCL().SysCallN(4343, m.Instance(), uintptr(AValue), uintptr(Kind))
	return AsMenuItem(r1)
}

func (m *TMenu) GetHelpContext(AValue uint32, ByCommand bool) THelpContext {
	r1 := LCL().SysCallN(4344, m.Instance(), uintptr(AValue), PascalBool(ByCommand))
	return THelpContext(r1)
}

func (m *TMenu) IsShortcut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(4351, m.Instance(), uintptr(unsafePointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func (m *TMenu) HandleAllocated() bool {
	r1 := LCL().SysCallN(4346, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) IsRightToLeft() bool {
	r1 := LCL().SysCallN(4350, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) UseRightToLeftAlignment() bool {
	r1 := LCL().SysCallN(4359, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) UseRightToLeftReading() bool {
	r1 := LCL().SysCallN(4360, m.Instance())
	return GoBool(r1)
}

func (m *TMenu) DispatchCommand(ACommand Word) bool {
	r1 := LCL().SysCallN(4342, m.Instance(), uintptr(ACommand))
	return GoBool(r1)
}

func MenuClass() TClass {
	ret := LCL().SysCallN(4339)
	return TClass(ret)
}

func (m *TMenu) DestroyHandle() {
	LCL().SysCallN(4341, m.Instance())
}

func (m *TMenu) HandleNeeded() {
	LCL().SysCallN(4347, m.Instance())
}

func (m *TMenu) SetOnDrawItem(fn TMenuDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4356, m.Instance(), m.drawItemPtr)
}

func (m *TMenu) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4357, m.Instance(), m.measureItemPtr)
}
