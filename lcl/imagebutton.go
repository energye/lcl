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

// IImageButton Parent: IGraphicControl
type IImageButton interface {
	IGraphicControl
	Caption() string                         // property
	SetCaption(AValue string)                // property
	DragCursor() TCursor                     // property
	SetDragCursor(AValue TCursor)            // property
	DragKind() TDragKind                     // property
	SetDragKind(AValue TDragKind)            // property
	DragMode() TDragMode                     // property
	SetDragMode(AValue TDragMode)            // property
	ImageCount() int32                       // property
	SetImageCount(AValue int32)              // property
	Orientation() TImageOrientation          // property
	SetOrientation(AValue TImageOrientation) // property
	ModalResult() TModalResult               // property
	SetModalResult(AValue TModalResult)      // property
	ParentShowHint() bool                    // property
	SetParentShowHint(AValue bool)           // property
	ParentFont() bool                        // property
	SetParentFont(AValue bool)               // property
	Picture() IPicture                       // property
	SetPicture(AValue IPicture)              // property
	ShowCaption() bool                       // property
	SetShowCaption(AValue bool)              // property
	Wordwarp() bool                          // property
	SetWordwarp(AValue bool)                 // property
	Click()                                  // procedure
	SetOnContextPopup(fn TContextPopupEvent) // property event
	SetOnDblClick(fn TNotifyEvent)           // property event
	SetOnDragDrop(fn TDragDropEvent)         // property event
	SetOnDragOver(fn TDragOverEvent)         // property event
	SetOnEndDock(fn TEndDragEvent)           // property event
	SetOnEndDrag(fn TEndDragEvent)           // property event
	SetOnMouseDown(fn TMouseEvent)           // property event
	SetOnMouseEnter(fn TNotifyEvent)         // property event
	SetOnMouseLeave(fn TNotifyEvent)         // property event
	SetOnMouseMove(fn TMouseMoveEvent)       // property event
	SetOnMouseUp(fn TMouseEvent)             // property event
}

// TImageButton Parent: TGraphicControl
type TImageButton struct {
	TGraphicControl
	contextPopupPtr uintptr
	dblClickPtr     uintptr
	dragDropPtr     uintptr
	dragOverPtr     uintptr
	endDockPtr      uintptr
	endDragPtr      uintptr
	mouseDownPtr    uintptr
	mouseEnterPtr   uintptr
	mouseLeavePtr   uintptr
	mouseMovePtr    uintptr
	mouseUpPtr      uintptr
}

func NewImageButton(AOwner IComponent) IImageButton {
	r1 := mageButtonImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsImageButton(r1)
}

func (m *TImageButton) Caption() string {
	r1 := mageButtonImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TImageButton) SetCaption(AValue string) {
	mageButtonImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TImageButton) DragCursor() TCursor {
	r1 := mageButtonImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TImageButton) SetDragCursor(AValue TCursor) {
	mageButtonImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) DragKind() TDragKind {
	r1 := mageButtonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TImageButton) SetDragKind(AValue TDragKind) {
	mageButtonImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) DragMode() TDragMode {
	r1 := mageButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TImageButton) SetDragMode(AValue TDragMode) {
	mageButtonImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) ImageCount() int32 {
	r1 := mageButtonImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TImageButton) SetImageCount(AValue int32) {
	mageButtonImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) Orientation() TImageOrientation {
	r1 := mageButtonImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TImageOrientation(r1)
}

func (m *TImageButton) SetOrientation(AValue TImageOrientation) {
	mageButtonImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) ModalResult() TModalResult {
	r1 := mageButtonImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TImageButton) SetModalResult(AValue TModalResult) {
	mageButtonImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) ParentShowHint() bool {
	r1 := mageButtonImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetParentShowHint(AValue bool) {
	mageButtonImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TImageButton) ParentFont() bool {
	r1 := mageButtonImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetParentFont(AValue bool) {
	mageButtonImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TImageButton) Picture() IPicture {
	r1 := mageButtonImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TImageButton) SetPicture(AValue IPicture) {
	mageButtonImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TImageButton) ShowCaption() bool {
	r1 := mageButtonImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetShowCaption(AValue bool) {
	mageButtonImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TImageButton) Wordwarp() bool {
	r1 := mageButtonImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetWordwarp(AValue bool) {
	mageButtonImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func ImageButtonClass() TClass {
	ret := mageButtonImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TImageButton) Click() {
	mageButtonImportAPI().SysCallN(2, m.Instance())
}

func (m *TImageButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(13, m.Instance(), m.contextPopupPtr)
}

func (m *TImageButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(14, m.Instance(), m.dblClickPtr)
}

func (m *TImageButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(15, m.Instance(), m.dragDropPtr)
}

func (m *TImageButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(16, m.Instance(), m.dragOverPtr)
}

func (m *TImageButton) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(17, m.Instance(), m.endDockPtr)
}

func (m *TImageButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(18, m.Instance(), m.endDragPtr)
}

func (m *TImageButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(19, m.Instance(), m.mouseDownPtr)
}

func (m *TImageButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(20, m.Instance(), m.mouseEnterPtr)
}

func (m *TImageButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(21, m.Instance(), m.mouseLeavePtr)
}

func (m *TImageButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(22, m.Instance(), m.mouseMovePtr)
}

func (m *TImageButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	mageButtonImportAPI().SysCallN(23, m.Instance(), m.mouseUpPtr)
}

var (
	mageButtonImport       *imports.Imports = nil
	mageButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ImageButton_Caption", 0),
		/*1*/ imports.NewTable("ImageButton_Class", 0),
		/*2*/ imports.NewTable("ImageButton_Click", 0),
		/*3*/ imports.NewTable("ImageButton_Create", 0),
		/*4*/ imports.NewTable("ImageButton_DragCursor", 0),
		/*5*/ imports.NewTable("ImageButton_DragKind", 0),
		/*6*/ imports.NewTable("ImageButton_DragMode", 0),
		/*7*/ imports.NewTable("ImageButton_ImageCount", 0),
		/*8*/ imports.NewTable("ImageButton_ModalResult", 0),
		/*9*/ imports.NewTable("ImageButton_Orientation", 0),
		/*10*/ imports.NewTable("ImageButton_ParentFont", 0),
		/*11*/ imports.NewTable("ImageButton_ParentShowHint", 0),
		/*12*/ imports.NewTable("ImageButton_Picture", 0),
		/*13*/ imports.NewTable("ImageButton_SetOnContextPopup", 0),
		/*14*/ imports.NewTable("ImageButton_SetOnDblClick", 0),
		/*15*/ imports.NewTable("ImageButton_SetOnDragDrop", 0),
		/*16*/ imports.NewTable("ImageButton_SetOnDragOver", 0),
		/*17*/ imports.NewTable("ImageButton_SetOnEndDock", 0),
		/*18*/ imports.NewTable("ImageButton_SetOnEndDrag", 0),
		/*19*/ imports.NewTable("ImageButton_SetOnMouseDown", 0),
		/*20*/ imports.NewTable("ImageButton_SetOnMouseEnter", 0),
		/*21*/ imports.NewTable("ImageButton_SetOnMouseLeave", 0),
		/*22*/ imports.NewTable("ImageButton_SetOnMouseMove", 0),
		/*23*/ imports.NewTable("ImageButton_SetOnMouseUp", 0),
		/*24*/ imports.NewTable("ImageButton_ShowCaption", 0),
		/*25*/ imports.NewTable("ImageButton_Wordwarp", 0),
	}
)

func mageButtonImportAPI() *imports.Imports {
	if mageButtonImport == nil {
		mageButtonImport = NewDefaultImports()
		mageButtonImport.SetImportTable(mageButtonImportTables)
		mageButtonImportTables = nil
	}
	return mageButtonImport
}
