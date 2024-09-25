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

// IDirectoryEdit Parent: ICustomEditButton
type IDirectoryEdit interface {
	ICustomEditButton
	AutoSelected() bool                           // property
	SetAutoSelected(AValue bool)                  // property
	Directory() string                            // property
	SetDirectory(AValue string)                   // property
	RootDir() string                              // property
	SetRootDir(AValue string)                     // property
	DialogTitle() string                          // property
	SetDialogTitle(AValue string)                 // property
	DialogOptions() TOpenOptions                  // property
	SetDialogOptions(AValue TOpenOptions)         // property
	ShowHidden() bool                             // property
	SetShowHidden(AValue bool)                    // property
	ButtonCaption() string                        // property
	SetButtonCaption(AValue string)               // property
	ButtonCursor() TCursor                        // property
	SetButtonCursor(AValue TCursor)               // property
	ButtonHint() string                           // property
	SetButtonHint(AValue string)                  // property
	ButtonOnlyWhenFocused() bool                  // property
	SetButtonOnlyWhenFocused(AValue bool)         // property
	ButtonWidth() int32                           // property
	SetButtonWidth(AValue int32)                  // property
	DirectInput() bool                            // property
	SetDirectInput(AValue bool)                   // property
	Glyph() IBitmap                               // property
	SetGlyph(AValue IBitmap)                      // property
	NumGlyphs() int32                             // property
	SetNumGlyphs(AValue int32)                    // property
	Images() ICustomImageList                     // property
	SetImages(AValue ICustomImageList)            // property
	ImageIndex() TImageIndex                      // property
	SetImageIndex(AValue TImageIndex)             // property
	ImageWidth() int32                            // property
	SetImageWidth(AValue int32)                   // property
	Flat() bool                                   // property
	SetFlat(AValue bool)                          // property
	FocusOnButtonClick() bool                     // property
	SetFocusOnButtonClick(AValue bool)            // property
	AutoSelect() bool                             // property
	SetAutoSelect(AValue bool)                    // property
	DragCursor() TCursor                          // property
	SetDragCursor(AValue TCursor)                 // property
	DragMode() TDragMode                          // property
	SetDragMode(AValue TDragMode)                 // property
	Layout() TLeftRight                           // property
	SetLayout(AValue TLeftRight)                  // property
	ParentFont() bool                             // property
	SetParentFont(AValue bool)                    // property
	ParentShowHint() bool                         // property
	SetParentShowHint(AValue bool)                // property
	Spacing() int32                               // property
	SetSpacing(AValue int32)                      // property
	RunDialog()                                   // procedure
	SetOnAcceptDirectory(fn TAcceptFileNameEvent) // property event
	SetOnButtonClick(fn TNotifyEvent)             // property event
}

// TDirectoryEdit Parent: TCustomEditButton
type TDirectoryEdit struct {
	TCustomEditButton
	acceptDirectoryPtr uintptr
	buttonClickPtr     uintptr
}

func NewDirectoryEdit(AOwner IComponent) IDirectoryEdit {
	r1 := LCL().SysCallN(2604, GetObjectUintptr(AOwner))
	return AsDirectoryEdit(r1)
}

func (m *TDirectoryEdit) AutoSelected() bool {
	r1 := LCL().SysCallN(2597, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelected(AValue bool) {
	LCL().SysCallN(2597, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Directory() string {
	r1 := LCL().SysCallN(2608, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDirectory(AValue string) {
	LCL().SysCallN(2608, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) RootDir() string {
	r1 := LCL().SysCallN(2621, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetRootDir(AValue string) {
	LCL().SysCallN(2621, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogTitle() string {
	r1 := LCL().SysCallN(2606, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDialogTitle(AValue string) {
	LCL().SysCallN(2606, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogOptions() TOpenOptions {
	r1 := LCL().SysCallN(2605, 0, m.Instance(), 0)
	return TOpenOptions(r1)
}

func (m *TDirectoryEdit) SetDialogOptions(AValue TOpenOptions) {
	LCL().SysCallN(2605, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ShowHidden() bool {
	r1 := LCL().SysCallN(2625, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetShowHidden(AValue bool) {
	LCL().SysCallN(2625, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonCaption() string {
	r1 := LCL().SysCallN(2598, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonCaption(AValue string) {
	LCL().SysCallN(2598, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonCursor() TCursor {
	r1 := LCL().SysCallN(2599, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetButtonCursor(AValue TCursor) {
	LCL().SysCallN(2599, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ButtonHint() string {
	r1 := LCL().SysCallN(2600, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonHint(AValue string) {
	LCL().SysCallN(2600, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonOnlyWhenFocused() bool {
	r1 := LCL().SysCallN(2601, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetButtonOnlyWhenFocused(AValue bool) {
	LCL().SysCallN(2601, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonWidth() int32 {
	r1 := LCL().SysCallN(2602, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetButtonWidth(AValue int32) {
	LCL().SysCallN(2602, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DirectInput() bool {
	r1 := LCL().SysCallN(2607, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetDirectInput(AValue bool) {
	LCL().SysCallN(2607, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Glyph() IBitmap {
	r1 := LCL().SysCallN(2613, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TDirectoryEdit) SetGlyph(AValue IBitmap) {
	LCL().SysCallN(2613, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) NumGlyphs() int32 {
	r1 := LCL().SysCallN(2618, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetNumGlyphs(AValue int32) {
	LCL().SysCallN(2618, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Images() ICustomImageList {
	r1 := LCL().SysCallN(2616, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TDirectoryEdit) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(2616, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(2614, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDirectoryEdit) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(2614, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ImageWidth() int32 {
	r1 := LCL().SysCallN(2615, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetImageWidth(AValue int32) {
	LCL().SysCallN(2615, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Flat() bool {
	r1 := LCL().SysCallN(2611, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFlat(AValue bool) {
	LCL().SysCallN(2611, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) FocusOnButtonClick() bool {
	r1 := LCL().SysCallN(2612, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFocusOnButtonClick(AValue bool) {
	LCL().SysCallN(2612, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(2596, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(2596, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) DragCursor() TCursor {
	r1 := LCL().SysCallN(2609, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2609, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DragMode() TDragMode {
	r1 := LCL().SysCallN(2610, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDirectoryEdit) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(2610, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Layout() TLeftRight {
	r1 := LCL().SysCallN(2617, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TDirectoryEdit) SetLayout(AValue TLeftRight) {
	LCL().SysCallN(2617, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ParentFont() bool {
	r1 := LCL().SysCallN(2619, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(2619, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(2620, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2620, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Spacing() int32 {
	r1 := LCL().SysCallN(2626, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetSpacing(AValue int32) {
	LCL().SysCallN(2626, 1, m.Instance(), uintptr(AValue))
}

func DirectoryEditClass() TClass {
	ret := LCL().SysCallN(2603)
	return TClass(ret)
}

func (m *TDirectoryEdit) RunDialog() {
	LCL().SysCallN(2622, m.Instance())
}

func (m *TDirectoryEdit) SetOnAcceptDirectory(fn TAcceptFileNameEvent) {
	if m.acceptDirectoryPtr != 0 {
		RemoveEventElement(m.acceptDirectoryPtr)
	}
	m.acceptDirectoryPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2623, m.Instance(), m.acceptDirectoryPtr)
}

func (m *TDirectoryEdit) SetOnButtonClick(fn TNotifyEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2624, m.Instance(), m.buttonClickPtr)
}
