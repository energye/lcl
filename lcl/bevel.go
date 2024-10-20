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

// IBevel Parent: IGraphicControl
type IBevel interface {
	IGraphicControl
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	Shape() TBevelShape                            // property
	SetShape(AValue TBevelShape)                   // property
	Style() TBevelStyle                            // property
	SetStyle(AValue TBevelStyle)                   // property
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnPaint(fn TNotifyEvent)                    // property event
}

// TBevel Parent: TGraphicControl
type TBevel struct {
	TGraphicControl
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	paintPtr          uintptr
}

func NewBevel(AOwner IComponent) IBevel {
	r1 := bevelImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsBevel(r1)
}

func (m *TBevel) ParentShowHint() bool {
	r1 := bevelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBevel) SetParentShowHint(AValue bool) {
	bevelImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBevel) Shape() TBevelShape {
	r1 := bevelImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TBevelShape(r1)
}

func (m *TBevel) SetShape(AValue TBevelShape) {
	bevelImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TBevel) Style() TBevelStyle {
	r1 := bevelImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TBevelStyle(r1)
}

func (m *TBevel) SetStyle(AValue TBevelStyle) {
	bevelImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func BevelClass() TClass {
	ret := bevelImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TBevel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(3, m.Instance(), m.mouseDownPtr)
}

func (m *TBevel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(4, m.Instance(), m.mouseEnterPtr)
}

func (m *TBevel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(5, m.Instance(), m.mouseLeavePtr)
}

func (m *TBevel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(6, m.Instance(), m.mouseMovePtr)
}

func (m *TBevel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(7, m.Instance(), m.mouseUpPtr)
}

func (m *TBevel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(8, m.Instance(), m.mouseWheelPtr)
}

func (m *TBevel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(9, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TBevel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(10, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TBevel) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	bevelImportAPI().SysCallN(11, m.Instance(), m.paintPtr)
}

var (
	bevelImport       *imports.Imports = nil
	bevelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Bevel_Class", 0),
		/*1*/ imports.NewTable("Bevel_Create", 0),
		/*2*/ imports.NewTable("Bevel_ParentShowHint", 0),
		/*3*/ imports.NewTable("Bevel_SetOnMouseDown", 0),
		/*4*/ imports.NewTable("Bevel_SetOnMouseEnter", 0),
		/*5*/ imports.NewTable("Bevel_SetOnMouseLeave", 0),
		/*6*/ imports.NewTable("Bevel_SetOnMouseMove", 0),
		/*7*/ imports.NewTable("Bevel_SetOnMouseUp", 0),
		/*8*/ imports.NewTable("Bevel_SetOnMouseWheel", 0),
		/*9*/ imports.NewTable("Bevel_SetOnMouseWheelDown", 0),
		/*10*/ imports.NewTable("Bevel_SetOnMouseWheelUp", 0),
		/*11*/ imports.NewTable("Bevel_SetOnPaint", 0),
		/*12*/ imports.NewTable("Bevel_Shape", 0),
		/*13*/ imports.NewTable("Bevel_Style", 0),
	}
)

func bevelImportAPI() *imports.Imports {
	if bevelImport == nil {
		bevelImport = NewDefaultImports()
		bevelImport.SetImportTable(bevelImportTables)
		bevelImportTables = nil
	}
	return bevelImport
}
