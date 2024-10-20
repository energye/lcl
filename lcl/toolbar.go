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

// IToolBar Parent: IToolWindow
type IToolBar interface {
	IToolWindow
	ButtonCount() int32                                    // property
	Buttons(Index int32) IToolButton                       // property
	ButtonList() IList                                     // property
	RowCount() int32                                       // property
	ButtonDropWidth() int32                                // property
	ButtonHeight() int32                                   // property
	SetButtonHeight(AValue int32)                          // property
	ButtonWidth() int32                                    // property
	SetButtonWidth(AValue int32)                           // property
	DisabledImages() ICustomImageList                      // property
	SetDisabledImages(AValue ICustomImageList)             // property
	DragCursor() TCursor                                   // property
	SetDragCursor(AValue TCursor)                          // property
	DragKind() TDragKind                                   // property
	SetDragKind(AValue TDragKind)                          // property
	DragMode() TDragMode                                   // property
	SetDragMode(AValue TDragMode)                          // property
	DropDownWidth() int32                                  // property
	SetDropDownWidth(AValue int32)                         // property
	Flat() bool                                            // property
	SetFlat(AValue bool)                                   // property
	HotImages() ICustomImageList                           // property
	SetHotImages(AValue ICustomImageList)                  // property
	Images() ICustomImageList                              // property
	SetImages(AValue ICustomImageList)                     // property
	ImagesWidth() int32                                    // property
	SetImagesWidth(AValue int32)                           // property
	Indent() int32                                         // property
	SetIndent(AValue int32)                                // property
	List() bool                                            // property
	SetList(AValue bool)                                   // property
	ParentColor() bool                                     // property
	SetParentColor(AValue bool)                            // property
	ParentFont() bool                                      // property
	SetParentFont(AValue bool)                             // property
	ParentShowHint() bool                                  // property
	SetParentShowHint(AValue bool)                         // property
	ShowCaptions() bool                                    // property
	SetShowCaptions(AValue bool)                           // property
	Transparent() bool                                     // property
	SetTransparent(AValue bool)                            // property
	Wrapable() bool                                        // property
	SetWrapable(AValue bool)                               // property
	GetEnumeratorForToolBarEnumerator() IToolBarEnumerator // function
	SetButtonSize(NewButtonWidth, NewButtonHeight int32)   // procedure
	SetOnContextPopup(fn TContextPopupEvent)               // property event
	SetOnDblClick(fn TNotifyEvent)                         // property event
	SetOnDragDrop(fn TDragDropEvent)                       // property event
	SetOnDragOver(fn TDragOverEvent)                       // property event
	SetOnPaintButton(fn TToolBarOnPaintButton)             // property event
	SetOnEndDrag(fn TEndDragEvent)                         // property event
	SetOnMouseDown(fn TMouseEvent)                         // property event
	SetOnMouseEnter(fn TNotifyEvent)                       // property event
	SetOnMouseLeave(fn TNotifyEvent)                       // property event
	SetOnMouseMove(fn TMouseMoveEvent)                     // property event
	SetOnMouseUp(fn TMouseEvent)                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                   // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)         // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)           // property event
	SetOnStartDrag(fn TStartDragEvent)                     // property event
}

// TToolBar Parent: TToolWindow
type TToolBar struct {
	TToolWindow
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	paintButtonPtr    uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewToolBar(TheOwner IComponent) IToolBar {
	r1 := oolBarImportAPI().SysCallN(7, GetObjectUintptr(TheOwner))
	return AsToolBar(r1)
}

func (m *TToolBar) ButtonCount() int32 {
	r1 := oolBarImportAPI().SysCallN(0, m.Instance())
	return int32(r1)
}

func (m *TToolBar) Buttons(Index int32) IToolButton {
	r1 := oolBarImportAPI().SysCallN(5, m.Instance(), uintptr(Index))
	return AsToolButton(r1)
}

func (m *TToolBar) ButtonList() IList {
	r1 := oolBarImportAPI().SysCallN(3, m.Instance())
	return AsList(r1)
}

func (m *TToolBar) RowCount() int32 {
	r1 := oolBarImportAPI().SysCallN(23, m.Instance())
	return int32(r1)
}

func (m *TToolBar) ButtonDropWidth() int32 {
	r1 := oolBarImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TToolBar) ButtonHeight() int32 {
	r1 := oolBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetButtonHeight(AValue int32) {
	oolBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) ButtonWidth() int32 {
	r1 := oolBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetButtonWidth(AValue int32) {
	oolBarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DisabledImages() ICustomImageList {
	r1 := oolBarImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TToolBar) SetDisabledImages(AValue ICustomImageList) {
	oolBarImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolBar) DragCursor() TCursor {
	r1 := oolBarImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TToolBar) SetDragCursor(AValue TCursor) {
	oolBarImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DragKind() TDragKind {
	r1 := oolBarImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TToolBar) SetDragKind(AValue TDragKind) {
	oolBarImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DragMode() TDragMode {
	r1 := oolBarImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TToolBar) SetDragMode(AValue TDragMode) {
	oolBarImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DropDownWidth() int32 {
	r1 := oolBarImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetDropDownWidth(AValue int32) {
	oolBarImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) Flat() bool {
	r1 := oolBarImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetFlat(AValue bool) {
	oolBarImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) HotImages() ICustomImageList {
	r1 := oolBarImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TToolBar) SetHotImages(AValue ICustomImageList) {
	oolBarImportAPI().SysCallN(15, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolBar) Images() ICustomImageList {
	r1 := oolBarImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TToolBar) SetImages(AValue ICustomImageList) {
	oolBarImportAPI().SysCallN(16, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolBar) ImagesWidth() int32 {
	r1 := oolBarImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetImagesWidth(AValue int32) {
	oolBarImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) Indent() int32 {
	r1 := oolBarImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetIndent(AValue int32) {
	oolBarImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) List() bool {
	r1 := oolBarImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetList(AValue bool) {
	oolBarImportAPI().SysCallN(19, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ParentColor() bool {
	r1 := oolBarImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetParentColor(AValue bool) {
	oolBarImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ParentFont() bool {
	r1 := oolBarImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetParentFont(AValue bool) {
	oolBarImportAPI().SysCallN(21, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ParentShowHint() bool {
	r1 := oolBarImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetParentShowHint(AValue bool) {
	oolBarImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ShowCaptions() bool {
	r1 := oolBarImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetShowCaptions(AValue bool) {
	oolBarImportAPI().SysCallN(40, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) Transparent() bool {
	r1 := oolBarImportAPI().SysCallN(41, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetTransparent(AValue bool) {
	oolBarImportAPI().SysCallN(41, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) Wrapable() bool {
	r1 := oolBarImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetWrapable(AValue bool) {
	oolBarImportAPI().SysCallN(42, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) GetEnumeratorForToolBarEnumerator() IToolBarEnumerator {
	r1 := oolBarImportAPI().SysCallN(14, m.Instance())
	return AsToolBarEnumerator(r1)
}

func ToolBarClass() TClass {
	ret := oolBarImportAPI().SysCallN(6)
	return TClass(ret)
}

func (m *TToolBar) SetButtonSize(NewButtonWidth, NewButtonHeight int32) {
	oolBarImportAPI().SysCallN(24, m.Instance(), uintptr(NewButtonWidth), uintptr(NewButtonHeight))
}

func (m *TToolBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(25, m.Instance(), m.contextPopupPtr)
}

func (m *TToolBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(26, m.Instance(), m.dblClickPtr)
}

func (m *TToolBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(27, m.Instance(), m.dragDropPtr)
}

func (m *TToolBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(28, m.Instance(), m.dragOverPtr)
}

func (m *TToolBar) SetOnPaintButton(fn TToolBarOnPaintButton) {
	if m.paintButtonPtr != 0 {
		RemoveEventElement(m.paintButtonPtr)
	}
	m.paintButtonPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(38, m.Instance(), m.paintButtonPtr)
}

func (m *TToolBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(29, m.Instance(), m.endDragPtr)
}

func (m *TToolBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(30, m.Instance(), m.mouseDownPtr)
}

func (m *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(31, m.Instance(), m.mouseEnterPtr)
}

func (m *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(32, m.Instance(), m.mouseLeavePtr)
}

func (m *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(33, m.Instance(), m.mouseMovePtr)
}

func (m *TToolBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(34, m.Instance(), m.mouseUpPtr)
}

func (m *TToolBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(35, m.Instance(), m.mouseWheelPtr)
}

func (m *TToolBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(36, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TToolBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(37, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TToolBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	oolBarImportAPI().SysCallN(39, m.Instance(), m.startDragPtr)
}

var (
	oolBarImport       *imports.Imports = nil
	oolBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ToolBar_ButtonCount", 0),
		/*1*/ imports.NewTable("ToolBar_ButtonDropWidth", 0),
		/*2*/ imports.NewTable("ToolBar_ButtonHeight", 0),
		/*3*/ imports.NewTable("ToolBar_ButtonList", 0),
		/*4*/ imports.NewTable("ToolBar_ButtonWidth", 0),
		/*5*/ imports.NewTable("ToolBar_Buttons", 0),
		/*6*/ imports.NewTable("ToolBar_Class", 0),
		/*7*/ imports.NewTable("ToolBar_Create", 0),
		/*8*/ imports.NewTable("ToolBar_DisabledImages", 0),
		/*9*/ imports.NewTable("ToolBar_DragCursor", 0),
		/*10*/ imports.NewTable("ToolBar_DragKind", 0),
		/*11*/ imports.NewTable("ToolBar_DragMode", 0),
		/*12*/ imports.NewTable("ToolBar_DropDownWidth", 0),
		/*13*/ imports.NewTable("ToolBar_Flat", 0),
		/*14*/ imports.NewTable("ToolBar_GetEnumeratorForToolBarEnumerator", 0),
		/*15*/ imports.NewTable("ToolBar_HotImages", 0),
		/*16*/ imports.NewTable("ToolBar_Images", 0),
		/*17*/ imports.NewTable("ToolBar_ImagesWidth", 0),
		/*18*/ imports.NewTable("ToolBar_Indent", 0),
		/*19*/ imports.NewTable("ToolBar_List", 0),
		/*20*/ imports.NewTable("ToolBar_ParentColor", 0),
		/*21*/ imports.NewTable("ToolBar_ParentFont", 0),
		/*22*/ imports.NewTable("ToolBar_ParentShowHint", 0),
		/*23*/ imports.NewTable("ToolBar_RowCount", 0),
		/*24*/ imports.NewTable("ToolBar_SetButtonSize", 0),
		/*25*/ imports.NewTable("ToolBar_SetOnContextPopup", 0),
		/*26*/ imports.NewTable("ToolBar_SetOnDblClick", 0),
		/*27*/ imports.NewTable("ToolBar_SetOnDragDrop", 0),
		/*28*/ imports.NewTable("ToolBar_SetOnDragOver", 0),
		/*29*/ imports.NewTable("ToolBar_SetOnEndDrag", 0),
		/*30*/ imports.NewTable("ToolBar_SetOnMouseDown", 0),
		/*31*/ imports.NewTable("ToolBar_SetOnMouseEnter", 0),
		/*32*/ imports.NewTable("ToolBar_SetOnMouseLeave", 0),
		/*33*/ imports.NewTable("ToolBar_SetOnMouseMove", 0),
		/*34*/ imports.NewTable("ToolBar_SetOnMouseUp", 0),
		/*35*/ imports.NewTable("ToolBar_SetOnMouseWheel", 0),
		/*36*/ imports.NewTable("ToolBar_SetOnMouseWheelDown", 0),
		/*37*/ imports.NewTable("ToolBar_SetOnMouseWheelUp", 0),
		/*38*/ imports.NewTable("ToolBar_SetOnPaintButton", 0),
		/*39*/ imports.NewTable("ToolBar_SetOnStartDrag", 0),
		/*40*/ imports.NewTable("ToolBar_ShowCaptions", 0),
		/*41*/ imports.NewTable("ToolBar_Transparent", 0),
		/*42*/ imports.NewTable("ToolBar_Wrapable", 0),
	}
)

func oolBarImportAPI() *imports.Imports {
	if oolBarImport == nil {
		oolBarImport = NewDefaultImports()
		oolBarImport.SetImportTable(oolBarImportTables)
		oolBarImportTables = nil
	}
	return oolBarImport
}
