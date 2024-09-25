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
	r1 := LCL().SysCallN(2750, GetObjectUintptr(AControl))
	return AsDragObject(r1)
}

func NewDragObjectAuto(AControl IControl) IDragObject {
	r1 := LCL().SysCallN(2745, GetObjectUintptr(AControl))
	return AsDragObject(r1)
}

func (m *TDragObject) AlwaysShowDragImages() bool {
	r1 := LCL().SysCallN(2744, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDragObject) SetAlwaysShowDragImages(AValue bool) {
	LCL().SysCallN(2744, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDragObject) AutoCreated() bool {
	r1 := LCL().SysCallN(2746, m.Instance())
	return GoBool(r1)
}

func (m *TDragObject) AutoFree() bool {
	r1 := LCL().SysCallN(2747, m.Instance())
	return GoBool(r1)
}

func (m *TDragObject) Control() IControl {
	r1 := LCL().SysCallN(2749, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragObject) SetControl(AValue IControl) {
	LCL().SysCallN(2749, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragObject) DragPos() (resultPoint TPoint) {
	LCL().SysCallN(2751, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragObject) SetDragPos(AValue *TPoint) {
	LCL().SysCallN(2751, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragObject) DragTarget() IControl {
	r1 := LCL().SysCallN(2752, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TDragObject) SetDragTarget(AValue IControl) {
	LCL().SysCallN(2752, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDragObject) DragTargetPos() (resultPoint TPoint) {
	LCL().SysCallN(2753, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragObject) SetDragTargetPos(AValue *TPoint) {
	LCL().SysCallN(2753, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragObject) Dropped() bool {
	r1 := LCL().SysCallN(2754, m.Instance())
	return GoBool(r1)
}

func DragObjectClass() TClass {
	ret := LCL().SysCallN(2748)
	return TClass(ret)
}

func (m *TDragObject) HideDragImage() {
	LCL().SysCallN(2755, m.Instance())
}

func (m *TDragObject) ShowDragImage() {
	LCL().SysCallN(2756, m.Instance())
}
