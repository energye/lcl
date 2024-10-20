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

// IMaskEdit Parent: ICustomMaskEdit
type IMaskEdit interface {
	ICustomMaskEdit
	IsMasked() bool                                // property
	EditText() string                              // property
	SetEditText(AValue string)                     // property
	AutoSelect() bool                              // property
	SetAutoSelect(AValue bool)                     // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	EditMask() string                              // property
	SetEditMask(AValue string)                     // property
	SpaceChar() Char                               // property
	SetSpaceChar(AValue Char)                      // property
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TMaskEdit Parent: TCustomMaskEdit
type TMaskEdit struct {
	TCustomMaskEdit
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	editingDonePtr    uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewMaskEdit(TheOwner IComponent) IMaskEdit {
	r1 := maskEditImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsMaskEdit(r1)
}

func (m *TMaskEdit) IsMasked() bool {
	r1 := maskEditImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TMaskEdit) EditText() string {
	r1 := maskEditImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMaskEdit) SetEditText(AValue string) {
	maskEditImportAPI().SysCallN(7, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMaskEdit) AutoSelect() bool {
	r1 := maskEditImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetAutoSelect(AValue bool) {
	maskEditImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) DragCursor() TCursor {
	r1 := maskEditImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TMaskEdit) SetDragCursor(AValue TCursor) {
	maskEditImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TMaskEdit) DragKind() TDragKind {
	r1 := maskEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TMaskEdit) SetDragKind(AValue TDragKind) {
	maskEditImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TMaskEdit) DragMode() TDragMode {
	r1 := maskEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TMaskEdit) SetDragMode(AValue TDragMode) {
	maskEditImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TMaskEdit) ParentColor() bool {
	r1 := maskEditImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetParentColor(AValue bool) {
	maskEditImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) ParentFont() bool {
	r1 := maskEditImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetParentFont(AValue bool) {
	maskEditImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) ParentShowHint() bool {
	r1 := maskEditImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetParentShowHint(AValue bool) {
	maskEditImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) EditMask() string {
	r1 := maskEditImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMaskEdit) SetEditMask(AValue string) {
	maskEditImportAPI().SysCallN(6, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMaskEdit) SpaceChar() Char {
	r1 := maskEditImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TMaskEdit) SetSpaceChar(AValue Char) {
	maskEditImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func MaskEditClass() TClass {
	ret := maskEditImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TMaskEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(12, m.Instance(), m.dblClickPtr)
}

func (m *TMaskEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(13, m.Instance(), m.dragDropPtr)
}

func (m *TMaskEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(14, m.Instance(), m.dragOverPtr)
}

func (m *TMaskEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(15, m.Instance(), m.editingDonePtr)
}

func (m *TMaskEdit) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(16, m.Instance(), m.endDockPtr)
}

func (m *TMaskEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(17, m.Instance(), m.endDragPtr)
}

func (m *TMaskEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TMaskEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(19, m.Instance(), m.mouseEnterPtr)
}

func (m *TMaskEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(20, m.Instance(), m.mouseLeavePtr)
}

func (m *TMaskEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(21, m.Instance(), m.mouseMovePtr)
}

func (m *TMaskEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(22, m.Instance(), m.mouseUpPtr)
}

func (m *TMaskEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(23, m.Instance(), m.mouseWheelPtr)
}

func (m *TMaskEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(24, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TMaskEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(25, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TMaskEdit) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(26, m.Instance(), m.startDockPtr)
}

func (m *TMaskEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	maskEditImportAPI().SysCallN(27, m.Instance(), m.startDragPtr)
}

var (
	maskEditImport       *imports.Imports = nil
	maskEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MaskEdit_AutoSelect", 0),
		/*1*/ imports.NewTable("MaskEdit_Class", 0),
		/*2*/ imports.NewTable("MaskEdit_Create", 0),
		/*3*/ imports.NewTable("MaskEdit_DragCursor", 0),
		/*4*/ imports.NewTable("MaskEdit_DragKind", 0),
		/*5*/ imports.NewTable("MaskEdit_DragMode", 0),
		/*6*/ imports.NewTable("MaskEdit_EditMask", 0),
		/*7*/ imports.NewTable("MaskEdit_EditText", 0),
		/*8*/ imports.NewTable("MaskEdit_IsMasked", 0),
		/*9*/ imports.NewTable("MaskEdit_ParentColor", 0),
		/*10*/ imports.NewTable("MaskEdit_ParentFont", 0),
		/*11*/ imports.NewTable("MaskEdit_ParentShowHint", 0),
		/*12*/ imports.NewTable("MaskEdit_SetOnDblClick", 0),
		/*13*/ imports.NewTable("MaskEdit_SetOnDragDrop", 0),
		/*14*/ imports.NewTable("MaskEdit_SetOnDragOver", 0),
		/*15*/ imports.NewTable("MaskEdit_SetOnEditingDone", 0),
		/*16*/ imports.NewTable("MaskEdit_SetOnEndDock", 0),
		/*17*/ imports.NewTable("MaskEdit_SetOnEndDrag", 0),
		/*18*/ imports.NewTable("MaskEdit_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("MaskEdit_SetOnMouseEnter", 0),
		/*20*/ imports.NewTable("MaskEdit_SetOnMouseLeave", 0),
		/*21*/ imports.NewTable("MaskEdit_SetOnMouseMove", 0),
		/*22*/ imports.NewTable("MaskEdit_SetOnMouseUp", 0),
		/*23*/ imports.NewTable("MaskEdit_SetOnMouseWheel", 0),
		/*24*/ imports.NewTable("MaskEdit_SetOnMouseWheelDown", 0),
		/*25*/ imports.NewTable("MaskEdit_SetOnMouseWheelUp", 0),
		/*26*/ imports.NewTable("MaskEdit_SetOnStartDock", 0),
		/*27*/ imports.NewTable("MaskEdit_SetOnStartDrag", 0),
		/*28*/ imports.NewTable("MaskEdit_SpaceChar", 0),
	}
)

func maskEditImportAPI() *imports.Imports {
	if maskEditImport == nil {
		maskEditImport = NewDefaultImports()
		maskEditImport.SetImportTable(maskEditImportTables)
		maskEditImportTables = nil
	}
	return maskEditImport
}
