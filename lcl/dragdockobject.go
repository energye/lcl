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

// IDragDockObject Parent: IDragObject
type IDragDockObject interface {
	IDragObject
	DockOffset() types.TPoint           // property DockOffset Getter
	SetDockOffset(value types.TPoint)   // property DockOffset Setter
	DockRect() types.TRect              // property DockRect Getter
	SetDockRect(value types.TRect)      // property DockRect Setter
	DropAlign() types.TAlign            // property DropAlign Getter
	SetDropAlign(value types.TAlign)    // property DropAlign Setter
	DropOnControl() IControl            // property DropOnControl Getter
	SetDropOnControl(value IControl)    // property DropOnControl Setter
	Floating() bool                     // property Floating Getter
	SetFloating(value bool)             // property Floating Setter
	IncreaseDockArea() bool             // property IncreaseDockArea Getter
	EraseDockRect() types.TRect         // property EraseDockRect Getter
	SetEraseDockRect(value types.TRect) // property EraseDockRect Setter
}

type TDragDockObject struct {
	TDragObject
}

func (m *TDragDockObject) DockOffset() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDragDockObject) SetDockOffset(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(2, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDragDockObject) DockRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDragDockObject) SetDockRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(3, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDragDockObject) DropAlign() types.TAlign {
	if !m.IsValid() {
		return 0
	}
	r := dragDockObjectAPI().SysCallN(4, 0, m.Instance())
	return types.TAlign(r)
}

func (m *TDragDockObject) SetDropAlign(value types.TAlign) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TDragDockObject) DropOnControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := dragDockObjectAPI().SysCallN(5, 0, m.Instance())
	return AsControl(r)
}

func (m *TDragDockObject) SetDropOnControl(value IControl) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDragDockObject) Floating() bool {
	if !m.IsValid() {
		return false
	}
	r := dragDockObjectAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDragDockObject) SetFloating(value bool) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TDragDockObject) IncreaseDockArea() bool {
	if !m.IsValid() {
		return false
	}
	r := dragDockObjectAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TDragDockObject) EraseDockRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDragDockObject) SetEraseDockRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	dragDockObjectAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

// NewDragDockObject class constructor
func NewDragDockObject(control IControl) IDragDockObject {
	r := dragDockObjectAPI().SysCallN(0, base.GetObjectUintptr(control))
	return AsDragDockObject(r)
}

// NewDragDockObjecteate class constructor
func NewDragDockObjecteate(control IControl) IDragDockObject {
	r := dragDockObjectAPI().SysCallN(1, base.GetObjectUintptr(control))
	return AsDragDockObject(r)
}

var (
	dragDockObjectOnce   base.Once
	dragDockObjectImport *imports.Imports = nil
)

func dragDockObjectAPI() *imports.Imports {
	dragDockObjectOnce.Do(func() {
		dragDockObjectImport = api.NewDefaultImports()
		dragDockObjectImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDragDockObject_Create", 0), // constructor NewDragDockObject
			/* 1 */ imports.NewTable("TDragDockObject_AutoCreate", 0), // constructor NewDragDockObjecteate
			/* 2 */ imports.NewTable("TDragDockObject_DockOffset", 0), // property DockOffset
			/* 3 */ imports.NewTable("TDragDockObject_DockRect", 0), // property DockRect
			/* 4 */ imports.NewTable("TDragDockObject_DropAlign", 0), // property DropAlign
			/* 5 */ imports.NewTable("TDragDockObject_DropOnControl", 0), // property DropOnControl
			/* 6 */ imports.NewTable("TDragDockObject_Floating", 0), // property Floating
			/* 7 */ imports.NewTable("TDragDockObject_IncreaseDockArea", 0), // property IncreaseDockArea
			/* 8 */ imports.NewTable("TDragDockObject_EraseDockRect", 0), // property EraseDockRect
		}
	})
	return dragDockObjectImport
}
