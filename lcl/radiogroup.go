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

// IRadioGroup Parent: ICustomRadioGroup
type IRadioGroup interface {
	ICustomRadioGroup
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TRadioGroup Parent: TCustomRadioGroup
type TRadioGroup struct {
	TCustomRadioGroup
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
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

func NewRadioGroup(TheOwner IComponent) IRadioGroup {
	r1 := radioGroupImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsRadioGroup(r1)
}

func (m *TRadioGroup) DragCursor() TCursor {
	r1 := radioGroupImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TRadioGroup) SetDragCursor(AValue TCursor) {
	radioGroupImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioGroup) DragMode() TDragMode {
	r1 := radioGroupImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TRadioGroup) SetDragMode(AValue TDragMode) {
	radioGroupImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioGroup) ParentFont() bool {
	r1 := radioGroupImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioGroup) SetParentFont(AValue bool) {
	radioGroupImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioGroup) ParentColor() bool {
	r1 := radioGroupImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioGroup) SetParentColor(AValue bool) {
	radioGroupImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioGroup) ParentShowHint() bool {
	r1 := radioGroupImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioGroup) SetParentShowHint(AValue bool) {
	radioGroupImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func RadioGroupClass() TClass {
	ret := radioGroupImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TRadioGroup) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(7, m.Instance(), m.dblClickPtr)
}

func (m *TRadioGroup) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(8, m.Instance(), m.dragDropPtr)
}

func (m *TRadioGroup) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(9, m.Instance(), m.dragOverPtr)
}

func (m *TRadioGroup) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(10, m.Instance(), m.endDragPtr)
}

func (m *TRadioGroup) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(11, m.Instance(), m.mouseDownPtr)
}

func (m *TRadioGroup) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(12, m.Instance(), m.mouseEnterPtr)
}

func (m *TRadioGroup) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(13, m.Instance(), m.mouseLeavePtr)
}

func (m *TRadioGroup) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(14, m.Instance(), m.mouseMovePtr)
}

func (m *TRadioGroup) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(15, m.Instance(), m.mouseUpPtr)
}

func (m *TRadioGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(16, m.Instance(), m.mouseWheelPtr)
}

func (m *TRadioGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(17, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TRadioGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(18, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TRadioGroup) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	radioGroupImportAPI().SysCallN(19, m.Instance(), m.startDragPtr)
}

var (
	radioGroupImport       *imports.Imports = nil
	radioGroupImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("RadioGroup_Class", 0),
		/*1*/ imports.NewTable("RadioGroup_Create", 0),
		/*2*/ imports.NewTable("RadioGroup_DragCursor", 0),
		/*3*/ imports.NewTable("RadioGroup_DragMode", 0),
		/*4*/ imports.NewTable("RadioGroup_ParentColor", 0),
		/*5*/ imports.NewTable("RadioGroup_ParentFont", 0),
		/*6*/ imports.NewTable("RadioGroup_ParentShowHint", 0),
		/*7*/ imports.NewTable("RadioGroup_SetOnDblClick", 0),
		/*8*/ imports.NewTable("RadioGroup_SetOnDragDrop", 0),
		/*9*/ imports.NewTable("RadioGroup_SetOnDragOver", 0),
		/*10*/ imports.NewTable("RadioGroup_SetOnEndDrag", 0),
		/*11*/ imports.NewTable("RadioGroup_SetOnMouseDown", 0),
		/*12*/ imports.NewTable("RadioGroup_SetOnMouseEnter", 0),
		/*13*/ imports.NewTable("RadioGroup_SetOnMouseLeave", 0),
		/*14*/ imports.NewTable("RadioGroup_SetOnMouseMove", 0),
		/*15*/ imports.NewTable("RadioGroup_SetOnMouseUp", 0),
		/*16*/ imports.NewTable("RadioGroup_SetOnMouseWheel", 0),
		/*17*/ imports.NewTable("RadioGroup_SetOnMouseWheelDown", 0),
		/*18*/ imports.NewTable("RadioGroup_SetOnMouseWheelUp", 0),
		/*19*/ imports.NewTable("RadioGroup_SetOnStartDrag", 0),
	}
)

func radioGroupImportAPI() *imports.Imports {
	if radioGroupImport == nil {
		radioGroupImport = NewDefaultImports()
		radioGroupImport.SetImportTable(radioGroupImportTables)
		radioGroupImportTables = nil
	}
	return radioGroupImport
}
