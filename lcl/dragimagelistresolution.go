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

// IDragImageListResolution Parent: ICustomImageListResolution
type IDragImageListResolution interface {
	ICustomImageListResolution
	DragHotspot() (resultPoint TPoint)           // property
	SetDragHotspot(AValue *TPoint)               // property
	Dragging() bool                              // property
	BeginDrag(Window HWND, X, Y int32) bool      // function
	DragLock(Window HWND, XPos, YPos int32) bool // function
	DragMove(X, Y int32) bool                    // function
	EndDrag() bool                               // function
	DragUnlock()                                 // procedure
	HideDragImage()                              // procedure
	ShowDragImage()                              // procedure
}

// TDragImageListResolution Parent: TCustomImageListResolution
type TDragImageListResolution struct {
	TCustomImageListResolution
}

func NewDragImageListResolution(TheOwner IComponent) IDragImageListResolution {
	r1 := dragImageListResolutionImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsDragImageListResolution(r1)
}

func (m *TDragImageListResolution) DragHotspot() (resultPoint TPoint) {
	dragImageListResolutionImportAPI().SysCallN(3, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragImageListResolution) SetDragHotspot(AValue *TPoint) {
	dragImageListResolutionImportAPI().SysCallN(3, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragImageListResolution) Dragging() bool {
	r1 := dragImageListResolutionImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageListResolution) BeginDrag(Window HWND, X, Y int32) bool {
	r1 := dragImageListResolutionImportAPI().SysCallN(0, m.Instance(), uintptr(Window), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageListResolution) DragLock(Window HWND, XPos, YPos int32) bool {
	r1 := dragImageListResolutionImportAPI().SysCallN(4, m.Instance(), uintptr(Window), uintptr(XPos), uintptr(YPos))
	return GoBool(r1)
}

func (m *TDragImageListResolution) DragMove(X, Y int32) bool {
	r1 := dragImageListResolutionImportAPI().SysCallN(5, m.Instance(), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageListResolution) EndDrag() bool {
	r1 := dragImageListResolutionImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func DragImageListResolutionClass() TClass {
	ret := dragImageListResolutionImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TDragImageListResolution) DragUnlock() {
	dragImageListResolutionImportAPI().SysCallN(6, m.Instance())
}

func (m *TDragImageListResolution) HideDragImage() {
	dragImageListResolutionImportAPI().SysCallN(9, m.Instance())
}

func (m *TDragImageListResolution) ShowDragImage() {
	dragImageListResolutionImportAPI().SysCallN(10, m.Instance())
}

var (
	dragImageListResolutionImport       *imports.Imports = nil
	dragImageListResolutionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DragImageListResolution_BeginDrag", 0),
		/*1*/ imports.NewTable("DragImageListResolution_Class", 0),
		/*2*/ imports.NewTable("DragImageListResolution_Create", 0),
		/*3*/ imports.NewTable("DragImageListResolution_DragHotspot", 0),
		/*4*/ imports.NewTable("DragImageListResolution_DragLock", 0),
		/*5*/ imports.NewTable("DragImageListResolution_DragMove", 0),
		/*6*/ imports.NewTable("DragImageListResolution_DragUnlock", 0),
		/*7*/ imports.NewTable("DragImageListResolution_Dragging", 0),
		/*8*/ imports.NewTable("DragImageListResolution_EndDrag", 0),
		/*9*/ imports.NewTable("DragImageListResolution_HideDragImage", 0),
		/*10*/ imports.NewTable("DragImageListResolution_ShowDragImage", 0),
	}
)

func dragImageListResolutionImportAPI() *imports.Imports {
	if dragImageListResolutionImport == nil {
		dragImageListResolutionImport = NewDefaultImports()
		dragImageListResolutionImport.SetImportTable(dragImageListResolutionImportTables)
		dragImageListResolutionImportTables = nil
	}
	return dragImageListResolutionImport
}
