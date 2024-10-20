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

// IDividerBevel Parent: IGraphicControl
type IDividerBevel interface {
	IGraphicControl
	BevelStyle() TBevelStyle                    // property
	SetBevelStyle(AValue TBevelStyle)           // property
	BevelWidth() int32                          // property
	SetBevelWidth(AValue int32)                 // property
	CaptionSpacing() int32                      // property
	SetCaptionSpacing(AValue int32)             // property
	DragCursor() TCursor                        // property
	SetDragCursor(AValue TCursor)               // property
	DragKind() TDragKind                        // property
	SetDragKind(AValue TDragKind)               // property
	DragMode() TDragMode                        // property
	SetDragMode(AValue TDragMode)               // property
	LeftIndent() int32                          // property
	SetLeftIndent(AValue int32)                 // property
	Orientation() TTrackBarOrientation          // property
	SetOrientation(AValue TTrackBarOrientation) // property
	ParentColor() bool                          // property
	SetParentColor(AValue bool)                 // property
	ParentFont() bool                           // property
	SetParentFont(AValue bool)                  // property
	ParentShowHint() bool                       // property
	SetParentShowHint(AValue bool)              // property
	Style() TGrabStyle                          // property
	SetStyle(AValue TGrabStyle)                 // property
	Transparent() bool                          // property
	SetTransparent(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)    // property event
	SetOnDblClick(fn TNotifyEvent)              // property event
	SetOnDragDrop(fn TDragDropEvent)            // property event
	SetOnDragOver(fn TDragOverEvent)            // property event
	SetOnEndDrag(fn TEndDragEvent)              // property event
	SetOnMouseDown(fn TMouseEvent)              // property event
	SetOnMouseEnter(fn TNotifyEvent)            // property event
	SetOnMouseLeave(fn TNotifyEvent)            // property event
	SetOnMouseMove(fn TMouseMoveEvent)          // property event
	SetOnMouseUp(fn TMouseEvent)                // property event
	SetOnStartDrag(fn TStartDragEvent)          // property event
}

// TDividerBevel Parent: TGraphicControl
type TDividerBevel struct {
	TGraphicControl
	contextPopupPtr uintptr
	dblClickPtr     uintptr
	dragDropPtr     uintptr
	dragOverPtr     uintptr
	endDragPtr      uintptr
	mouseDownPtr    uintptr
	mouseEnterPtr   uintptr
	mouseLeavePtr   uintptr
	mouseMovePtr    uintptr
	mouseUpPtr      uintptr
	startDragPtr    uintptr
}

func NewDividerBevel(AOwner IComponent) IDividerBevel {
	r1 := dividerBevelImportAPI().SysCallN(4, GetObjectUintptr(AOwner))
	return AsDividerBevel(r1)
}

func (m *TDividerBevel) BevelStyle() TBevelStyle {
	r1 := dividerBevelImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TBevelStyle(r1)
}

func (m *TDividerBevel) SetBevelStyle(AValue TBevelStyle) {
	dividerBevelImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) BevelWidth() int32 {
	r1 := dividerBevelImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDividerBevel) SetBevelWidth(AValue int32) {
	dividerBevelImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) CaptionSpacing() int32 {
	r1 := dividerBevelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDividerBevel) SetCaptionSpacing(AValue int32) {
	dividerBevelImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) DragCursor() TCursor {
	r1 := dividerBevelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDividerBevel) SetDragCursor(AValue TCursor) {
	dividerBevelImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) DragKind() TDragKind {
	r1 := dividerBevelImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TDividerBevel) SetDragKind(AValue TDragKind) {
	dividerBevelImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) DragMode() TDragMode {
	r1 := dividerBevelImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDividerBevel) SetDragMode(AValue TDragMode) {
	dividerBevelImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) LeftIndent() int32 {
	r1 := dividerBevelImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDividerBevel) SetLeftIndent(AValue int32) {
	dividerBevelImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) Orientation() TTrackBarOrientation {
	r1 := dividerBevelImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TTrackBarOrientation(r1)
}

func (m *TDividerBevel) SetOrientation(AValue TTrackBarOrientation) {
	dividerBevelImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) ParentColor() bool {
	r1 := dividerBevelImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetParentColor(AValue bool) {
	dividerBevelImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDividerBevel) ParentFont() bool {
	r1 := dividerBevelImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetParentFont(AValue bool) {
	dividerBevelImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDividerBevel) ParentShowHint() bool {
	r1 := dividerBevelImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetParentShowHint(AValue bool) {
	dividerBevelImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDividerBevel) Style() TGrabStyle {
	r1 := dividerBevelImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return TGrabStyle(r1)
}

func (m *TDividerBevel) SetStyle(AValue TGrabStyle) {
	dividerBevelImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) Transparent() bool {
	r1 := dividerBevelImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetTransparent(AValue bool) {
	dividerBevelImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func DividerBevelClass() TClass {
	ret := dividerBevelImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TDividerBevel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(13, m.Instance(), m.contextPopupPtr)
}

func (m *TDividerBevel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(14, m.Instance(), m.dblClickPtr)
}

func (m *TDividerBevel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(15, m.Instance(), m.dragDropPtr)
}

func (m *TDividerBevel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(16, m.Instance(), m.dragOverPtr)
}

func (m *TDividerBevel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(17, m.Instance(), m.endDragPtr)
}

func (m *TDividerBevel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TDividerBevel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(19, m.Instance(), m.mouseEnterPtr)
}

func (m *TDividerBevel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(20, m.Instance(), m.mouseLeavePtr)
}

func (m *TDividerBevel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(21, m.Instance(), m.mouseMovePtr)
}

func (m *TDividerBevel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(22, m.Instance(), m.mouseUpPtr)
}

func (m *TDividerBevel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	dividerBevelImportAPI().SysCallN(23, m.Instance(), m.startDragPtr)
}

var (
	dividerBevelImport       *imports.Imports = nil
	dividerBevelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DividerBevel_BevelStyle", 0),
		/*1*/ imports.NewTable("DividerBevel_BevelWidth", 0),
		/*2*/ imports.NewTable("DividerBevel_CaptionSpacing", 0),
		/*3*/ imports.NewTable("DividerBevel_Class", 0),
		/*4*/ imports.NewTable("DividerBevel_Create", 0),
		/*5*/ imports.NewTable("DividerBevel_DragCursor", 0),
		/*6*/ imports.NewTable("DividerBevel_DragKind", 0),
		/*7*/ imports.NewTable("DividerBevel_DragMode", 0),
		/*8*/ imports.NewTable("DividerBevel_LeftIndent", 0),
		/*9*/ imports.NewTable("DividerBevel_Orientation", 0),
		/*10*/ imports.NewTable("DividerBevel_ParentColor", 0),
		/*11*/ imports.NewTable("DividerBevel_ParentFont", 0),
		/*12*/ imports.NewTable("DividerBevel_ParentShowHint", 0),
		/*13*/ imports.NewTable("DividerBevel_SetOnContextPopup", 0),
		/*14*/ imports.NewTable("DividerBevel_SetOnDblClick", 0),
		/*15*/ imports.NewTable("DividerBevel_SetOnDragDrop", 0),
		/*16*/ imports.NewTable("DividerBevel_SetOnDragOver", 0),
		/*17*/ imports.NewTable("DividerBevel_SetOnEndDrag", 0),
		/*18*/ imports.NewTable("DividerBevel_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("DividerBevel_SetOnMouseEnter", 0),
		/*20*/ imports.NewTable("DividerBevel_SetOnMouseLeave", 0),
		/*21*/ imports.NewTable("DividerBevel_SetOnMouseMove", 0),
		/*22*/ imports.NewTable("DividerBevel_SetOnMouseUp", 0),
		/*23*/ imports.NewTable("DividerBevel_SetOnStartDrag", 0),
		/*24*/ imports.NewTable("DividerBevel_Style", 0),
		/*25*/ imports.NewTable("DividerBevel_Transparent", 0),
	}
)

func dividerBevelImportAPI() *imports.Imports {
	if dividerBevelImport == nil {
		dividerBevelImport = NewDefaultImports()
		dividerBevelImport.SetImportTable(dividerBevelImportTables)
		dividerBevelImportTables = nil
	}
	return dividerBevelImport
}
