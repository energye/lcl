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

// IDragImageList Parent: ICustomImageList
type IDragImageList interface {
	ICustomImageList
	DragCursor() TCursor                                                             // property
	SetDragCursor(AValue TCursor)                                                    // property
	DragHotspot() (resultPoint TPoint)                                               // property
	SetDragHotspot(AValue *TPoint)                                                   // property
	Dragging() bool                                                                  // property
	DraggingResolution() IDragImageListResolution                                    // property
	ResolutionForDragImageListResolution(AImageWidth int32) IDragImageListResolution // property
	BeginDrag(Window HWND, X, Y int32) bool                                          // function
	DragLock(Window HWND, XPos, YPos int32) bool                                     // function
	DragMove(X, Y int32) bool                                                        // function
	EndDrag() bool                                                                   // function
	SetDragImage(Index, HotSpotX, HotSpotY int32) bool                               // function
	DragUnlock()                                                                     // procedure
	HideDragImage()                                                                  // procedure
	ShowDragImage()                                                                  // procedure
}

// TDragImageList Parent: TCustomImageList
type TDragImageList struct {
	TCustomImageList
}

func NewDragImageList(AOwner IComponent) IDragImageList {
	r1 := dragImageListImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsDragImageList(r1)
}

func (m *TDragImageList) DragCursor() TCursor {
	r1 := dragImageListImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDragImageList) SetDragCursor(AValue TCursor) {
	dragImageListImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TDragImageList) DragHotspot() (resultPoint TPoint) {
	dragImageListImportAPI().SysCallN(4, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragImageList) SetDragHotspot(AValue *TPoint) {
	dragImageListImportAPI().SysCallN(4, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragImageList) Dragging() bool {
	r1 := dragImageListImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageList) DraggingResolution() IDragImageListResolution {
	r1 := dragImageListImportAPI().SysCallN(9, m.Instance())
	return AsDragImageListResolution(r1)
}

func (m *TDragImageList) ResolutionForDragImageListResolution(AImageWidth int32) IDragImageListResolution {
	r1 := dragImageListImportAPI().SysCallN(12, m.Instance(), uintptr(AImageWidth))
	return AsDragImageListResolution(r1)
}

func (m *TDragImageList) BeginDrag(Window HWND, X, Y int32) bool {
	r1 := dragImageListImportAPI().SysCallN(0, m.Instance(), uintptr(Window), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageList) DragLock(Window HWND, XPos, YPos int32) bool {
	r1 := dragImageListImportAPI().SysCallN(5, m.Instance(), uintptr(Window), uintptr(XPos), uintptr(YPos))
	return GoBool(r1)
}

func (m *TDragImageList) DragMove(X, Y int32) bool {
	r1 := dragImageListImportAPI().SysCallN(6, m.Instance(), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageList) EndDrag() bool {
	r1 := dragImageListImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageList) SetDragImage(Index, HotSpotX, HotSpotY int32) bool {
	r1 := dragImageListImportAPI().SysCallN(13, m.Instance(), uintptr(Index), uintptr(HotSpotX), uintptr(HotSpotY))
	return GoBool(r1)
}

func DragImageListClass() TClass {
	ret := dragImageListImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TDragImageList) DragUnlock() {
	dragImageListImportAPI().SysCallN(7, m.Instance())
}

func (m *TDragImageList) HideDragImage() {
	dragImageListImportAPI().SysCallN(11, m.Instance())
}

func (m *TDragImageList) ShowDragImage() {
	dragImageListImportAPI().SysCallN(14, m.Instance())
}

var (
	dragImageListImport       *imports.Imports = nil
	dragImageListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DragImageList_BeginDrag", 0),
		/*1*/ imports.NewTable("DragImageList_Class", 0),
		/*2*/ imports.NewTable("DragImageList_Create", 0),
		/*3*/ imports.NewTable("DragImageList_DragCursor", 0),
		/*4*/ imports.NewTable("DragImageList_DragHotspot", 0),
		/*5*/ imports.NewTable("DragImageList_DragLock", 0),
		/*6*/ imports.NewTable("DragImageList_DragMove", 0),
		/*7*/ imports.NewTable("DragImageList_DragUnlock", 0),
		/*8*/ imports.NewTable("DragImageList_Dragging", 0),
		/*9*/ imports.NewTable("DragImageList_DraggingResolution", 0),
		/*10*/ imports.NewTable("DragImageList_EndDrag", 0),
		/*11*/ imports.NewTable("DragImageList_HideDragImage", 0),
		/*12*/ imports.NewTable("DragImageList_ResolutionForDragImageListResolution", 0),
		/*13*/ imports.NewTable("DragImageList_SetDragImage", 0),
		/*14*/ imports.NewTable("DragImageList_ShowDragImage", 0),
	}
)

func dragImageListImportAPI() *imports.Imports {
	if dragImageListImport == nil {
		dragImageListImport = NewDefaultImports()
		dragImageListImport.SetImportTable(dragImageListImportTables)
		dragImageListImportTables = nil
	}
	return dragImageListImport
}
