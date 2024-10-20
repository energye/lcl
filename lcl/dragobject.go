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

// IDragObject Parent: IObject
type IDragObject interface {
	IObject
	AlwaysShowDragImages() bool          // property
	SetAlwaysShowDragImages(AValue bool) // property
	AutoCreated() bool                   // property
	AutoFree() bool                      // property
	Control() IControl                   // property
	SetControl(AValue IControl)          // property
	DragPos() (resultPoint TPoint)       // property
	SetDragPos(AValue *TPoint)           // property
	DragTarget() IControl                // property
	SetDragTarget(AValue IControl)       // property
	DragTargetPos() (resultPoint TPoint) // property
	SetDragTargetPos(AValue *TPoint)     // property
	Dropped() bool                       // property
	HideDragImage()                      // procedure
	ShowDragImage()                      // procedure
}

// TDragObject Parent: TObject
type TDragObject struct {
	TObject
}

func NewDragObject(AControl IControl) IDragObject {
	r1 := dragObjectImportAPI().SysCallN(6, GetObjectUintptr(AControl))
	return AsDragObject(r1)
}

func NewDragObjectAuto(AControl IControl) IDragObject {
	r1 := dragObjectImportAPI().SysCallN(1, GetObjectUintptr(AControl))
	return AsDragObject(r1)
}

func (m *TDragObject) AlwaysShowDragImages() bool {
	r1 := dragObjectImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDragObject) SetAlwaysShowDragImages(AValue bool) {
	dragObjectImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDragObject) AutoCreated() bool {
	r1 := dragObjectImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TDragObject) AutoFree() bool {
	r1 := dragObjectImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TDragObject) Control() IControl {
	r1 := dragObjectImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragObject) SetControl(AValue IControl) {
	dragObjectImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragObject) DragPos() (resultPoint TPoint) {
	dragObjectImportAPI().SysCallN(7, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragObject) SetDragPos(AValue *TPoint) {
	dragObjectImportAPI().SysCallN(7, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragObject) DragTarget() IControl {
	r1 := dragObjectImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragObject) SetDragTarget(AValue IControl) {
	dragObjectImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragObject) DragTargetPos() (resultPoint TPoint) {
	dragObjectImportAPI().SysCallN(9, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragObject) SetDragTargetPos(AValue *TPoint) {
	dragObjectImportAPI().SysCallN(9, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragObject) Dropped() bool {
	r1 := dragObjectImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func DragObjectClass() TClass {
	ret := dragObjectImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TDragObject) HideDragImage() {
	dragObjectImportAPI().SysCallN(11, m.Instance())
}

func (m *TDragObject) ShowDragImage() {
	dragObjectImportAPI().SysCallN(12, m.Instance())
}

var (
	dragObjectImport       *imports.Imports = nil
	dragObjectImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DragObject_AlwaysShowDragImages", 0),
		/*1*/ imports.NewTable("DragObject_AutoCreate", 0),
		/*2*/ imports.NewTable("DragObject_AutoCreated", 0),
		/*3*/ imports.NewTable("DragObject_AutoFree", 0),
		/*4*/ imports.NewTable("DragObject_Class", 0),
		/*5*/ imports.NewTable("DragObject_Control", 0),
		/*6*/ imports.NewTable("DragObject_Create", 0),
		/*7*/ imports.NewTable("DragObject_DragPos", 0),
		/*8*/ imports.NewTable("DragObject_DragTarget", 0),
		/*9*/ imports.NewTable("DragObject_DragTargetPos", 0),
		/*10*/ imports.NewTable("DragObject_Dropped", 0),
		/*11*/ imports.NewTable("DragObject_HideDragImage", 0),
		/*12*/ imports.NewTable("DragObject_ShowDragImage", 0),
	}
)

func dragObjectImportAPI() *imports.Imports {
	if dragObjectImport == nil {
		dragObjectImport = NewDefaultImports()
		dragObjectImport.SetImportTable(dragObjectImportTables)
		dragObjectImportTables = nil
	}
	return dragObjectImport
}
