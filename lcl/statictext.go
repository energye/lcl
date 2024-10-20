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

// IStaticText Parent: ICustomStaticText
type IStaticText interface {
	ICustomStaticText
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TStaticText Parent: TCustomStaticText
type TStaticText struct {
	TCustomStaticText
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewStaticText(AOwner IComponent) IStaticText {
	r1 := staticTextImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsStaticText(r1)
}

func (m *TStaticText) DragCursor() TCursor {
	r1 := staticTextImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStaticText) SetDragCursor(AValue TCursor) {
	staticTextImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TStaticText) DragKind() TDragKind {
	r1 := staticTextImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStaticText) SetDragKind(AValue TDragKind) {
	staticTextImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TStaticText) DragMode() TDragMode {
	r1 := staticTextImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStaticText) SetDragMode(AValue TDragMode) {
	staticTextImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TStaticText) ParentFont() bool {
	r1 := staticTextImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStaticText) SetParentFont(AValue bool) {
	staticTextImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStaticText) ParentColor() bool {
	r1 := staticTextImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStaticText) SetParentColor(AValue bool) {
	staticTextImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStaticText) ParentShowHint() bool {
	r1 := staticTextImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStaticText) SetParentShowHint(AValue bool) {
	staticTextImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func StaticTextClass() TClass {
	ret := staticTextImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TStaticText) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(8, m.Instance(), m.contextPopupPtr)
}

func (m *TStaticText) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(9, m.Instance(), m.dblClickPtr)
}

func (m *TStaticText) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TStaticText) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TStaticText) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(12, m.Instance(), m.endDragPtr)
}

func (m *TStaticText) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(13, m.Instance(), m.mouseDownPtr)
}

func (m *TStaticText) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(14, m.Instance(), m.mouseEnterPtr)
}

func (m *TStaticText) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(15, m.Instance(), m.mouseLeavePtr)
}

func (m *TStaticText) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(16, m.Instance(), m.mouseMovePtr)
}

func (m *TStaticText) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(17, m.Instance(), m.mouseUpPtr)
}

func (m *TStaticText) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(18, m.Instance(), m.mouseWheelPtr)
}

func (m *TStaticText) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(19, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TStaticText) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(23, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TStaticText) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(20, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TStaticText) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(21, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TStaticText) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(22, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TStaticText) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	staticTextImportAPI().SysCallN(24, m.Instance(), m.startDragPtr)
}

var (
	staticTextImport       *imports.Imports = nil
	staticTextImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StaticText_Class", 0),
		/*1*/ imports.NewTable("StaticText_Create", 0),
		/*2*/ imports.NewTable("StaticText_DragCursor", 0),
		/*3*/ imports.NewTable("StaticText_DragKind", 0),
		/*4*/ imports.NewTable("StaticText_DragMode", 0),
		/*5*/ imports.NewTable("StaticText_ParentColor", 0),
		/*6*/ imports.NewTable("StaticText_ParentFont", 0),
		/*7*/ imports.NewTable("StaticText_ParentShowHint", 0),
		/*8*/ imports.NewTable("StaticText_SetOnContextPopup", 0),
		/*9*/ imports.NewTable("StaticText_SetOnDblClick", 0),
		/*10*/ imports.NewTable("StaticText_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("StaticText_SetOnDragOver", 0),
		/*12*/ imports.NewTable("StaticText_SetOnEndDrag", 0),
		/*13*/ imports.NewTable("StaticText_SetOnMouseDown", 0),
		/*14*/ imports.NewTable("StaticText_SetOnMouseEnter", 0),
		/*15*/ imports.NewTable("StaticText_SetOnMouseLeave", 0),
		/*16*/ imports.NewTable("StaticText_SetOnMouseMove", 0),
		/*17*/ imports.NewTable("StaticText_SetOnMouseUp", 0),
		/*18*/ imports.NewTable("StaticText_SetOnMouseWheel", 0),
		/*19*/ imports.NewTable("StaticText_SetOnMouseWheelDown", 0),
		/*20*/ imports.NewTable("StaticText_SetOnMouseWheelHorz", 0),
		/*21*/ imports.NewTable("StaticText_SetOnMouseWheelLeft", 0),
		/*22*/ imports.NewTable("StaticText_SetOnMouseWheelRight", 0),
		/*23*/ imports.NewTable("StaticText_SetOnMouseWheelUp", 0),
		/*24*/ imports.NewTable("StaticText_SetOnStartDrag", 0),
	}
)

func staticTextImportAPI() *imports.Imports {
	if staticTextImport == nil {
		staticTextImport = NewDefaultImports()
		staticTextImport.SetImportTable(staticTextImportTables)
		staticTextImportTables = nil
	}
	return staticTextImport
}
