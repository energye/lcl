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

// IDragDockObject Parent: IDragObject
type IDragDockObject interface {
	IDragObject
	DockOffset() (resultPoint TPoint)  // property
	SetDockOffset(AValue *TPoint)      // property
	DockRect() (resultRect TRect)      // property
	SetDockRect(AValue *TRect)         // property
	DropAlign() TAlign                 // property
	SetDropAlign(AValue TAlign)        // property
	DropOnControl() IControl           // property
	SetDropOnControl(AValue IControl)  // property
	Floating() bool                    // property
	SetFloating(AValue bool)           // property
	IncreaseDockArea() bool            // property
	EraseDockRect() (resultRect TRect) // property
	SetEraseDockRect(AValue *TRect)    // property
}

// TDragDockObject Parent: TDragObject
type TDragDockObject struct {
	TDragObject
}

func NewDragDockObject(AControl IControl) IDragDockObject {
	r1 := dragDockObjectImportAPI().SysCallN(1, GetObjectUintptr(AControl))
	return AsDragDockObject(r1)
}

func (m *TDragDockObject) DockOffset() (resultPoint TPoint) {
	dragDockObjectImportAPI().SysCallN(2, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragDockObject) SetDockOffset(AValue *TPoint) {
	dragDockObjectImportAPI().SysCallN(2, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragDockObject) DockRect() (resultRect TRect) {
	dragDockObjectImportAPI().SysCallN(3, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TDragDockObject) SetDockRect(AValue *TRect) {
	dragDockObjectImportAPI().SysCallN(3, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragDockObject) DropAlign() TAlign {
	r1 := dragDockObjectImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TAlign(r1)
}

func (m *TDragDockObject) SetDropAlign(AValue TAlign) {
	dragDockObjectImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TDragDockObject) DropOnControl() IControl {
	r1 := dragDockObjectImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragDockObject) SetDropOnControl(AValue IControl) {
	dragDockObjectImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragDockObject) Floating() bool {
	r1 := dragDockObjectImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDragDockObject) SetFloating(AValue bool) {
	dragDockObjectImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDragDockObject) IncreaseDockArea() bool {
	r1 := dragDockObjectImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TDragDockObject) EraseDockRect() (resultRect TRect) {
	dragDockObjectImportAPI().SysCallN(6, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TDragDockObject) SetEraseDockRect(AValue *TRect) {
	dragDockObjectImportAPI().SysCallN(6, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func DragDockObjectClass() TClass {
	ret := dragDockObjectImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	dragDockObjectImport       *imports.Imports = nil
	dragDockObjectImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DragDockObject_Class", 0),
		/*1*/ imports.NewTable("DragDockObject_Create", 0),
		/*2*/ imports.NewTable("DragDockObject_DockOffset", 0),
		/*3*/ imports.NewTable("DragDockObject_DockRect", 0),
		/*4*/ imports.NewTable("DragDockObject_DropAlign", 0),
		/*5*/ imports.NewTable("DragDockObject_DropOnControl", 0),
		/*6*/ imports.NewTable("DragDockObject_EraseDockRect", 0),
		/*7*/ imports.NewTable("DragDockObject_Floating", 0),
		/*8*/ imports.NewTable("DragDockObject_IncreaseDockArea", 0),
	}
)

func dragDockObjectImportAPI() *imports.Imports {
	if dragDockObjectImport == nil {
		dragDockObjectImport = NewDefaultImports()
		dragDockObjectImport.SetImportTable(dragDockObjectImportTables)
		dragDockObjectImportTables = nil
	}
	return dragDockObjectImport
}
