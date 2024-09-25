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

// IMouse Parent: IObject
type IMouse interface {
	IObject
	Capture() HWND                   // property
	SetCapture(AValue HWND)          // property
	CursorPos() (resultPoint TPoint) // property
	SetCursorPos(AValue *TPoint)     // property
	IsDragging() bool                // property
	WheelScrollLines() int32         // property
	DragImmediate() bool             // property
	SetDragImmediate(AValue bool)    // property
	DragThreshold() int32            // property
	SetDragThreshold(AValue int32)   // property
}

// TMouse Parent: TObject
type TMouse struct {
	TObject
}

func NewMouse() IMouse {
	r1 := LCL().SysCallN(4381)
	return AsMouse(r1)
}

func (m *TMouse) Capture() HWND {
	r1 := LCL().SysCallN(4379, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TMouse) SetCapture(AValue HWND) {
	LCL().SysCallN(4379, 1, m.Instance(), uintptr(AValue))
}

func (m *TMouse) CursorPos() (resultPoint TPoint) {
	LCL().SysCallN(4382, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TMouse) SetCursorPos(AValue *TPoint) {
	LCL().SysCallN(4382, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TMouse) IsDragging() bool {
	r1 := LCL().SysCallN(4385, m.Instance())
	return GoBool(r1)
}

func (m *TMouse) WheelScrollLines() int32 {
	r1 := LCL().SysCallN(4386, m.Instance())
	return int32(r1)
}

func (m *TMouse) DragImmediate() bool {
	r1 := LCL().SysCallN(4383, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMouse) SetDragImmediate(AValue bool) {
	LCL().SysCallN(4383, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMouse) DragThreshold() int32 {
	r1 := LCL().SysCallN(4384, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMouse) SetDragThreshold(AValue int32) {
	LCL().SysCallN(4384, 1, m.Instance(), uintptr(AValue))
}

func MouseClass() TClass {
	ret := LCL().SysCallN(4380)
	return TClass(ret)
}
