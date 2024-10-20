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

// ILabel Parent: ICustomLabel
type ILabel interface {
	ICustomLabel
	Alignment() TAlignment                          // property
	SetAlignment(AValue TAlignment)                 // property
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	FocusControl() IWinControl                      // property
	SetFocusControl(AValue IWinControl)             // property
	Layout() TTextLayout                            // property
	SetLayout(AValue TTextLayout)                   // property
	OptimalFill() bool                              // property
	SetOptimalFill(AValue bool)                     // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	ShowAccelChar() bool                            // property
	SetShowAccelChar(AValue bool)                   // property
	Transparent() bool                              // property
	SetTransparent(AValue bool)                     // property
	WordWrap() bool                                 // property
	SetWordWrap(AValue bool)                        // property
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

// TLabel Parent: TCustomLabel
type TLabel struct {
	TCustomLabel
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

func NewLabel(TheOwner IComponent) ILabel {
	r1 := labelImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsLabel(r1)
}

func (m *TLabel) Alignment() TAlignment {
	r1 := labelImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLabel) SetAlignment(AValue TAlignment) {
	labelImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) DragCursor() TCursor {
	r1 := labelImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLabel) SetDragCursor(AValue TCursor) {
	labelImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) DragKind() TDragKind {
	r1 := labelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLabel) SetDragKind(AValue TDragKind) {
	labelImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) DragMode() TDragMode {
	r1 := labelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLabel) SetDragMode(AValue TDragMode) {
	labelImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) FocusControl() IWinControl {
	r1 := labelImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TLabel) SetFocusControl(AValue IWinControl) {
	labelImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLabel) Layout() TTextLayout {
	r1 := labelImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TLabel) SetLayout(AValue TTextLayout) {
	labelImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) OptimalFill() bool {
	r1 := labelImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetOptimalFill(AValue bool) {
	labelImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ParentColor() bool {
	r1 := labelImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetParentColor(AValue bool) {
	labelImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ParentFont() bool {
	r1 := labelImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetParentFont(AValue bool) {
	labelImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ParentShowHint() bool {
	r1 := labelImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetParentShowHint(AValue bool) {
	labelImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ShowAccelChar() bool {
	r1 := labelImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetShowAccelChar(AValue bool) {
	labelImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) Transparent() bool {
	r1 := labelImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetTransparent(AValue bool) {
	labelImportAPI().SysCallN(30, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) WordWrap() bool {
	r1 := labelImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetWordWrap(AValue bool) {
	labelImportAPI().SysCallN(31, 1, m.Instance(), PascalBool(AValue))
}

func LabelClass() TClass {
	ret := labelImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TLabel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(12, m.Instance(), m.contextPopupPtr)
}

func (m *TLabel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(13, m.Instance(), m.dblClickPtr)
}

func (m *TLabel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(14, m.Instance(), m.dragDropPtr)
}

func (m *TLabel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(15, m.Instance(), m.dragOverPtr)
}

func (m *TLabel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(16, m.Instance(), m.endDragPtr)
}

func (m *TLabel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(17, m.Instance(), m.mouseDownPtr)
}

func (m *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(18, m.Instance(), m.mouseEnterPtr)
}

func (m *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(19, m.Instance(), m.mouseLeavePtr)
}

func (m *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(20, m.Instance(), m.mouseMovePtr)
}

func (m *TLabel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(21, m.Instance(), m.mouseUpPtr)
}

func (m *TLabel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(22, m.Instance(), m.mouseWheelPtr)
}

func (m *TLabel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(23, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TLabel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(27, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TLabel) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(24, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TLabel) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(25, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TLabel) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(26, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TLabel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	labelImportAPI().SysCallN(28, m.Instance(), m.startDragPtr)
}

var (
	labelImport       *imports.Imports = nil
	labelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Label_Alignment", 0),
		/*1*/ imports.NewTable("Label_Class", 0),
		/*2*/ imports.NewTable("Label_Create", 0),
		/*3*/ imports.NewTable("Label_DragCursor", 0),
		/*4*/ imports.NewTable("Label_DragKind", 0),
		/*5*/ imports.NewTable("Label_DragMode", 0),
		/*6*/ imports.NewTable("Label_FocusControl", 0),
		/*7*/ imports.NewTable("Label_Layout", 0),
		/*8*/ imports.NewTable("Label_OptimalFill", 0),
		/*9*/ imports.NewTable("Label_ParentColor", 0),
		/*10*/ imports.NewTable("Label_ParentFont", 0),
		/*11*/ imports.NewTable("Label_ParentShowHint", 0),
		/*12*/ imports.NewTable("Label_SetOnContextPopup", 0),
		/*13*/ imports.NewTable("Label_SetOnDblClick", 0),
		/*14*/ imports.NewTable("Label_SetOnDragDrop", 0),
		/*15*/ imports.NewTable("Label_SetOnDragOver", 0),
		/*16*/ imports.NewTable("Label_SetOnEndDrag", 0),
		/*17*/ imports.NewTable("Label_SetOnMouseDown", 0),
		/*18*/ imports.NewTable("Label_SetOnMouseEnter", 0),
		/*19*/ imports.NewTable("Label_SetOnMouseLeave", 0),
		/*20*/ imports.NewTable("Label_SetOnMouseMove", 0),
		/*21*/ imports.NewTable("Label_SetOnMouseUp", 0),
		/*22*/ imports.NewTable("Label_SetOnMouseWheel", 0),
		/*23*/ imports.NewTable("Label_SetOnMouseWheelDown", 0),
		/*24*/ imports.NewTable("Label_SetOnMouseWheelHorz", 0),
		/*25*/ imports.NewTable("Label_SetOnMouseWheelLeft", 0),
		/*26*/ imports.NewTable("Label_SetOnMouseWheelRight", 0),
		/*27*/ imports.NewTable("Label_SetOnMouseWheelUp", 0),
		/*28*/ imports.NewTable("Label_SetOnStartDrag", 0),
		/*29*/ imports.NewTable("Label_ShowAccelChar", 0),
		/*30*/ imports.NewTable("Label_Transparent", 0),
		/*31*/ imports.NewTable("Label_WordWrap", 0),
	}
)

func labelImportAPI() *imports.Imports {
	if labelImport == nil {
		labelImport = NewDefaultImports()
		labelImport.SetImportTable(labelImportTables)
		labelImportTables = nil
	}
	return labelImport
}
