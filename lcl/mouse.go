//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// IMouse Parent: IObject
type IMouse interface {
	IObject
	Capture() types.HWND             // property Capture Getter
	SetCapture(value types.HWND)     // property Capture Setter
	CursorPos() types.TPoint         // property CursorPos Getter
	SetCursorPos(value types.TPoint) // property CursorPos Setter
	IsDragging() bool                // property IsDragging Getter
	WheelScrollLines() int32         // property WheelScrollLines Getter
	DragImmediate() bool             // property DragImmediate Getter
	SetDragImmediate(value bool)     // property DragImmediate Setter
	DragThreshold() int32            // property DragThreshold Getter
	SetDragThreshold(value int32)    // property DragThreshold Setter
}

type TMouse struct {
	TObject
}

func (m *TMouse) Capture() types.HWND {
	if !m.IsValid() {
		return 0
	}
	r := mouseAPI().SysCallN(0, 0, m.Instance())
	return types.HWND(r)
}

func (m *TMouse) SetCapture(value types.HWND) {
	if !m.IsValid() {
		return
	}
	mouseAPI().SysCallN(0, 1, m.Instance(), uintptr(value))
}

func (m *TMouse) CursorPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	mouseAPI().SysCallN(1, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TMouse) SetCursorPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	mouseAPI().SysCallN(1, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TMouse) IsDragging() bool {
	if !m.IsValid() {
		return false
	}
	r := mouseAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TMouse) WheelScrollLines() int32 {
	if !m.IsValid() {
		return 0
	}
	r := mouseAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TMouse) DragImmediate() bool {
	if !m.IsValid() {
		return false
	}
	r := mouseAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMouse) SetDragImmediate(value bool) {
	if !m.IsValid() {
		return
	}
	mouseAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TMouse) DragThreshold() int32 {
	if !m.IsValid() {
		return 0
	}
	r := mouseAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TMouse) SetDragThreshold(value int32) {
	if !m.IsValid() {
		return
	}
	mouseAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

var (
	mouseOnce   base.Once
	mouseImport *imports.Imports = nil
)

func mouseAPI() *imports.Imports {
	mouseOnce.Do(func() {
		mouseImport = api.NewDefaultImports()
		mouseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMouse_Capture", 0), // property Capture
			/* 1 */ imports.NewTable("TMouse_CursorPos", 0), // property CursorPos
			/* 2 */ imports.NewTable("TMouse_IsDragging", 0), // property IsDragging
			/* 3 */ imports.NewTable("TMouse_WheelScrollLines", 0), // property WheelScrollLines
			/* 4 */ imports.NewTable("TMouse_DragImmediate", 0), // property DragImmediate
			/* 5 */ imports.NewTable("TMouse_DragThreshold", 0), // property DragThreshold
		}
	})
	return mouseImport
}
