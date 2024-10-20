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

// IShape Parent: IGraphicControl
type IShape interface {
	IGraphicControl
	Brush() IBrush                                  // property
	SetBrush(AValue IBrush)                         // property
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	Pen() IPen                                      // property
	SetPen(AValue IPen)                             // property
	Shape() TShapeType                              // property
	SetShape(AValue TShapeType)                     // property
	Paint()                                         // procedure
	StyleChanged(Sender IObject)                    // procedure
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
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
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDock(fn TStartDockEvent)              // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TShape Parent: TGraphicControl
type TShape struct {
	TGraphicControl
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDockPtr         uintptr
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
	paintPtr           uintptr
	startDockPtr       uintptr
	startDragPtr       uintptr
}

func NewShape(TheOwner IComponent) IShape {
	r1 := shapeImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsShape(r1)
}

func (m *TShape) Brush() IBrush {
	r1 := shapeImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsBrush(r1)
}

func (m *TShape) SetBrush(AValue IBrush) {
	shapeImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TShape) DragCursor() TCursor {
	r1 := shapeImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TShape) SetDragCursor(AValue TCursor) {
	shapeImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) DragKind() TDragKind {
	r1 := shapeImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TShape) SetDragKind(AValue TDragKind) {
	shapeImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) DragMode() TDragMode {
	r1 := shapeImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TShape) SetDragMode(AValue TDragMode) {
	shapeImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) ParentShowHint() bool {
	r1 := shapeImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TShape) SetParentShowHint(AValue bool) {
	shapeImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TShape) Pen() IPen {
	r1 := shapeImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsPen(r1)
}

func (m *TShape) SetPen(AValue IPen) {
	shapeImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TShape) Shape() TShapeType {
	r1 := shapeImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return TShapeType(r1)
}

func (m *TShape) SetShape(AValue TShapeType) {
	shapeImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func ShapeClass() TClass {
	ret := shapeImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TShape) Paint() {
	shapeImportAPI().SysCallN(6, m.Instance())
}

func (m *TShape) StyleChanged(Sender IObject) {
	shapeImportAPI().SysCallN(28, m.Instance(), GetObjectUintptr(Sender))
}

func (m *TShape) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(9, m.Instance(), m.dragDropPtr)
}

func (m *TShape) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(10, m.Instance(), m.dragOverPtr)
}

func (m *TShape) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(11, m.Instance(), m.endDockPtr)
}

func (m *TShape) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(12, m.Instance(), m.endDragPtr)
}

func (m *TShape) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(13, m.Instance(), m.mouseDownPtr)
}

func (m *TShape) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(14, m.Instance(), m.mouseEnterPtr)
}

func (m *TShape) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(15, m.Instance(), m.mouseLeavePtr)
}

func (m *TShape) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(16, m.Instance(), m.mouseMovePtr)
}

func (m *TShape) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(17, m.Instance(), m.mouseUpPtr)
}

func (m *TShape) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(18, m.Instance(), m.mouseWheelPtr)
}

func (m *TShape) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(19, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TShape) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(23, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TShape) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(20, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TShape) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(21, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TShape) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(22, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TShape) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(24, m.Instance(), m.paintPtr)
}

func (m *TShape) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(25, m.Instance(), m.startDockPtr)
}

func (m *TShape) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	shapeImportAPI().SysCallN(26, m.Instance(), m.startDragPtr)
}

var (
	shapeImport       *imports.Imports = nil
	shapeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Shape_Brush", 0),
		/*1*/ imports.NewTable("Shape_Class", 0),
		/*2*/ imports.NewTable("Shape_Create", 0),
		/*3*/ imports.NewTable("Shape_DragCursor", 0),
		/*4*/ imports.NewTable("Shape_DragKind", 0),
		/*5*/ imports.NewTable("Shape_DragMode", 0),
		/*6*/ imports.NewTable("Shape_Paint", 0),
		/*7*/ imports.NewTable("Shape_ParentShowHint", 0),
		/*8*/ imports.NewTable("Shape_Pen", 0),
		/*9*/ imports.NewTable("Shape_SetOnDragDrop", 0),
		/*10*/ imports.NewTable("Shape_SetOnDragOver", 0),
		/*11*/ imports.NewTable("Shape_SetOnEndDock", 0),
		/*12*/ imports.NewTable("Shape_SetOnEndDrag", 0),
		/*13*/ imports.NewTable("Shape_SetOnMouseDown", 0),
		/*14*/ imports.NewTable("Shape_SetOnMouseEnter", 0),
		/*15*/ imports.NewTable("Shape_SetOnMouseLeave", 0),
		/*16*/ imports.NewTable("Shape_SetOnMouseMove", 0),
		/*17*/ imports.NewTable("Shape_SetOnMouseUp", 0),
		/*18*/ imports.NewTable("Shape_SetOnMouseWheel", 0),
		/*19*/ imports.NewTable("Shape_SetOnMouseWheelDown", 0),
		/*20*/ imports.NewTable("Shape_SetOnMouseWheelHorz", 0),
		/*21*/ imports.NewTable("Shape_SetOnMouseWheelLeft", 0),
		/*22*/ imports.NewTable("Shape_SetOnMouseWheelRight", 0),
		/*23*/ imports.NewTable("Shape_SetOnMouseWheelUp", 0),
		/*24*/ imports.NewTable("Shape_SetOnPaint", 0),
		/*25*/ imports.NewTable("Shape_SetOnStartDock", 0),
		/*26*/ imports.NewTable("Shape_SetOnStartDrag", 0),
		/*27*/ imports.NewTable("Shape_Shape", 0),
		/*28*/ imports.NewTable("Shape_StyleChanged", 0),
	}
)

func shapeImportAPI() *imports.Imports {
	if shapeImport == nil {
		shapeImport = NewDefaultImports()
		shapeImport.SetImportTable(shapeImportTables)
		shapeImportTables = nil
	}
	return shapeImport
}
