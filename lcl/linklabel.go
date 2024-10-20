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

// ILinkLabel Parent: ICustomLabel
type ILinkLabel interface {
	ICustomLabel
	Alignment() TAlignment                   // property
	SetAlignment(AValue TAlignment)          // property
	DragCursor() TCursor                     // property
	SetDragCursor(AValue TCursor)            // property
	DragKind() TDragKind                     // property
	SetDragKind(AValue TDragKind)            // property
	DragMode() TDragMode                     // property
	SetDragMode(AValue TDragMode)            // property
	ParentColor() bool                       // property
	SetParentColor(AValue bool)              // property
	ParentFont() bool                        // property
	SetParentFont(AValue bool)               // property
	ParentShowHint() bool                    // property
	SetParentShowHint(AValue bool)           // property
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
	SetOnStartDock(fn TStartDockEvent)       // property event
	SetOnStartDrag(fn TStartDragEvent)       // property event
	SetOnLinkClick(fn TSysLinkEvent)         // property event
}

// TLinkLabel Parent: TCustomLabel
type TLinkLabel struct {
	TCustomLabel
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
	startDockPtr    uintptr
	startDragPtr    uintptr
	linkClickPtr    uintptr
}

func NewLinkLabel(AOwner IComponent) ILinkLabel {
	r1 := linkLabelImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsLinkLabel(r1)
}

func (m *TLinkLabel) Alignment() TAlignment {
	r1 := linkLabelImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLinkLabel) SetAlignment(AValue TAlignment) {
	linkLabelImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) DragCursor() TCursor {
	r1 := linkLabelImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLinkLabel) SetDragCursor(AValue TCursor) {
	linkLabelImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) DragKind() TDragKind {
	r1 := linkLabelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLinkLabel) SetDragKind(AValue TDragKind) {
	linkLabelImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) DragMode() TDragMode {
	r1 := linkLabelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLinkLabel) SetDragMode(AValue TDragMode) {
	linkLabelImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TLinkLabel) ParentColor() bool {
	r1 := linkLabelImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLinkLabel) SetParentColor(AValue bool) {
	linkLabelImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLinkLabel) ParentFont() bool {
	r1 := linkLabelImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLinkLabel) SetParentFont(AValue bool) {
	linkLabelImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLinkLabel) ParentShowHint() bool {
	r1 := linkLabelImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLinkLabel) SetParentShowHint(AValue bool) {
	linkLabelImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func LinkLabelClass() TClass {
	ret := linkLabelImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TLinkLabel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(9, m.Instance(), m.contextPopupPtr)
}

func (m *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(10, m.Instance(), m.dblClickPtr)
}

func (m *TLinkLabel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(11, m.Instance(), m.dragDropPtr)
}

func (m *TLinkLabel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(12, m.Instance(), m.dragOverPtr)
}

func (m *TLinkLabel) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(13, m.Instance(), m.endDockPtr)
}

func (m *TLinkLabel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(14, m.Instance(), m.endDragPtr)
}

func (m *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(16, m.Instance(), m.mouseDownPtr)
}

func (m *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(17, m.Instance(), m.mouseEnterPtr)
}

func (m *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(18, m.Instance(), m.mouseLeavePtr)
}

func (m *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(19, m.Instance(), m.mouseMovePtr)
}

func (m *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(20, m.Instance(), m.mouseUpPtr)
}

func (m *TLinkLabel) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(21, m.Instance(), m.startDockPtr)
}

func (m *TLinkLabel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(22, m.Instance(), m.startDragPtr)
}

func (m *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
	if m.linkClickPtr != 0 {
		RemoveEventElement(m.linkClickPtr)
	}
	m.linkClickPtr = MakeEventDataPtr(fn)
	linkLabelImportAPI().SysCallN(15, m.Instance(), m.linkClickPtr)
}

var (
	linkLabelImport       *imports.Imports = nil
	linkLabelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LinkLabel_Alignment", 0),
		/*1*/ imports.NewTable("LinkLabel_Class", 0),
		/*2*/ imports.NewTable("LinkLabel_Create", 0),
		/*3*/ imports.NewTable("LinkLabel_DragCursor", 0),
		/*4*/ imports.NewTable("LinkLabel_DragKind", 0),
		/*5*/ imports.NewTable("LinkLabel_DragMode", 0),
		/*6*/ imports.NewTable("LinkLabel_ParentColor", 0),
		/*7*/ imports.NewTable("LinkLabel_ParentFont", 0),
		/*8*/ imports.NewTable("LinkLabel_ParentShowHint", 0),
		/*9*/ imports.NewTable("LinkLabel_SetOnContextPopup", 0),
		/*10*/ imports.NewTable("LinkLabel_SetOnDblClick", 0),
		/*11*/ imports.NewTable("LinkLabel_SetOnDragDrop", 0),
		/*12*/ imports.NewTable("LinkLabel_SetOnDragOver", 0),
		/*13*/ imports.NewTable("LinkLabel_SetOnEndDock", 0),
		/*14*/ imports.NewTable("LinkLabel_SetOnEndDrag", 0),
		/*15*/ imports.NewTable("LinkLabel_SetOnLinkClick", 0),
		/*16*/ imports.NewTable("LinkLabel_SetOnMouseDown", 0),
		/*17*/ imports.NewTable("LinkLabel_SetOnMouseEnter", 0),
		/*18*/ imports.NewTable("LinkLabel_SetOnMouseLeave", 0),
		/*19*/ imports.NewTable("LinkLabel_SetOnMouseMove", 0),
		/*20*/ imports.NewTable("LinkLabel_SetOnMouseUp", 0),
		/*21*/ imports.NewTable("LinkLabel_SetOnStartDock", 0),
		/*22*/ imports.NewTable("LinkLabel_SetOnStartDrag", 0),
	}
)

func linkLabelImportAPI() *imports.Imports {
	if linkLabelImport == nil {
		linkLabelImport = NewDefaultImports()
		linkLabelImport.SetImportTable(linkLabelImportTables)
		linkLabelImportTables = nil
	}
	return linkLabelImport
}
