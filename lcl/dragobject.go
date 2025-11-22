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

// IDragObject Parent: IObject
type IDragObject interface {
	IObject
	HideDragImage()                      // procedure
	ShowDragImage()                      // procedure
	AlwaysShowDragImages() bool          // property AlwaysShowDragImages Getter
	SetAlwaysShowDragImages(value bool)  // property AlwaysShowDragImages Setter
	AutoCreated() bool                   // property AutoCreated Getter
	AutoFree() bool                      // property AutoFree Getter
	Control() IControl                   // property Control Getter
	SetControl(value IControl)           // property Control Setter
	DragPos() types.TPoint               // property DragPos Getter
	SetDragPos(value types.TPoint)       // property DragPos Setter
	DragTarget() IControl                // property DragTarget Getter
	SetDragTarget(value IControl)        // property DragTarget Setter
	DragTargetPos() types.TPoint         // property DragTargetPos Getter
	SetDragTargetPos(value types.TPoint) // property DragTargetPos Setter
	Dropped() bool                       // property Dropped Getter
}

type TDragObject struct {
	TObject
}

func (m *TDragObject) HideDragImage() {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(2, m.Instance())
}

func (m *TDragObject) ShowDragImage() {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(3, m.Instance())
}

func (m *TDragObject) AlwaysShowDragImages() bool {
	if !m.IsValid() {
		return false
	}
	r := dragObjectAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDragObject) SetAlwaysShowDragImages(value bool) {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TDragObject) AutoCreated() bool {
	if !m.IsValid() {
		return false
	}
	r := dragObjectAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TDragObject) AutoFree() bool {
	if !m.IsValid() {
		return false
	}
	r := dragObjectAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TDragObject) Control() IControl {
	if !m.IsValid() {
		return nil
	}
	r := dragObjectAPI().SysCallN(7, 0, m.Instance())
	return AsControl(r)
}

func (m *TDragObject) SetControl(value IControl) {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDragObject) DragPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDragObject) SetDragPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDragObject) DragTarget() IControl {
	if !m.IsValid() {
		return nil
	}
	r := dragObjectAPI().SysCallN(9, 0, m.Instance())
	return AsControl(r)
}

func (m *TDragObject) SetDragTarget(value IControl) {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDragObject) DragTargetPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDragObject) SetDragTargetPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	dragObjectAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDragObject) Dropped() bool {
	if !m.IsValid() {
		return false
	}
	r := dragObjectAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

// NewDragObject class constructor
func NewDragObject(control IControl) IDragObject {
	r := dragObjectAPI().SysCallN(0, base.GetObjectUintptr(control))
	return AsDragObject(r)
}

// NewDragObjecteate class constructor
func NewDragObjecteate(control IControl) IDragObject {
	r := dragObjectAPI().SysCallN(1, base.GetObjectUintptr(control))
	return AsDragObject(r)
}

var (
	dragObjectOnce   base.Once
	dragObjectImport *imports.Imports = nil
)

func dragObjectAPI() *imports.Imports {
	dragObjectOnce.Do(func() {
		dragObjectImport = api.NewDefaultImports()
		dragObjectImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDragObject_Create", 0), // constructor NewDragObject
			/* 1 */ imports.NewTable("TDragObject_AutoCreate", 0), // constructor NewDragObjecteate
			/* 2 */ imports.NewTable("TDragObject_HideDragImage", 0), // procedure HideDragImage
			/* 3 */ imports.NewTable("TDragObject_ShowDragImage", 0), // procedure ShowDragImage
			/* 4 */ imports.NewTable("TDragObject_AlwaysShowDragImages", 0), // property AlwaysShowDragImages
			/* 5 */ imports.NewTable("TDragObject_AutoCreated", 0), // property AutoCreated
			/* 6 */ imports.NewTable("TDragObject_AutoFree", 0), // property AutoFree
			/* 7 */ imports.NewTable("TDragObject_Control", 0), // property Control
			/* 8 */ imports.NewTable("TDragObject_DragPos", 0), // property DragPos
			/* 9 */ imports.NewTable("TDragObject_DragTarget", 0), // property DragTarget
			/* 10 */ imports.NewTable("TDragObject_DragTargetPos", 0), // property DragTargetPos
			/* 11 */ imports.NewTable("TDragObject_Dropped", 0), // property Dropped
		}
	})
	return dragObjectImport
}
