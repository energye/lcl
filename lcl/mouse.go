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
	r1 := mouseImportAPI().SysCallN(2)
	return AsMouse(r1)
}

func (m *TMouse) Capture() HWND {
	r1 := mouseImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TMouse) SetCapture(AValue HWND) {
	mouseImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TMouse) CursorPos() (resultPoint TPoint) {
	mouseImportAPI().SysCallN(3, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TMouse) SetCursorPos(AValue *TPoint) {
	mouseImportAPI().SysCallN(3, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TMouse) IsDragging() bool {
	r1 := mouseImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TMouse) WheelScrollLines() int32 {
	r1 := mouseImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TMouse) DragImmediate() bool {
	r1 := mouseImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMouse) SetDragImmediate(AValue bool) {
	mouseImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMouse) DragThreshold() int32 {
	r1 := mouseImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMouse) SetDragThreshold(AValue int32) {
	mouseImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func MouseClass() TClass {
	ret := mouseImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	mouseImport       *imports.Imports = nil
	mouseImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Mouse_Capture", 0),
		/*1*/ imports.NewTable("Mouse_Class", 0),
		/*2*/ imports.NewTable("Mouse_Create", 0),
		/*3*/ imports.NewTable("Mouse_CursorPos", 0),
		/*4*/ imports.NewTable("Mouse_DragImmediate", 0),
		/*5*/ imports.NewTable("Mouse_DragThreshold", 0),
		/*6*/ imports.NewTable("Mouse_IsDragging", 0),
		/*7*/ imports.NewTable("Mouse_WheelScrollLines", 0),
	}
)

func mouseImportAPI() *imports.Imports {
	if mouseImport == nil {
		mouseImport = NewDefaultImports()
		mouseImport.SetImportTable(mouseImportTables)
		mouseImportTables = nil
	}
	return mouseImport
}
