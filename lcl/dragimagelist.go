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
	r1 := LCL().SysCallN(2731, GetObjectUintptr(AOwner))
	return AsDragImageList(r1)
}

func (m *TDragImageList) DragCursor() TCursor {
	r1 := LCL().SysCallN(2732, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDragImageList) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2732, 1, m.Instance(), uintptr(AValue))
}

func (m *TDragImageList) DragHotspot() (resultPoint TPoint) {
	LCL().SysCallN(2733, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDragImageList) SetDragHotspot(AValue *TPoint) {
	LCL().SysCallN(2733, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDragImageList) Dragging() bool {
	r1 := LCL().SysCallN(2737, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageList) DraggingResolution() IDragImageListResolution {
	r1 := LCL().SysCallN(2738, m.Instance())
	return AsDragImageListResolution(r1)
}

func (m *TDragImageList) ResolutionForDragImageListResolution(AImageWidth int32) IDragImageListResolution {
	r1 := LCL().SysCallN(2741, m.Instance(), uintptr(AImageWidth))
	return AsDragImageListResolution(r1)
}

func (m *TDragImageList) BeginDrag(Window HWND, X, Y int32) bool {
	r1 := LCL().SysCallN(2729, m.Instance(), uintptr(Window), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageList) DragLock(Window HWND, XPos, YPos int32) bool {
	r1 := LCL().SysCallN(2734, m.Instance(), uintptr(Window), uintptr(XPos), uintptr(YPos))
	return GoBool(r1)
}

func (m *TDragImageList) DragMove(X, Y int32) bool {
	r1 := LCL().SysCallN(2735, m.Instance(), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func (m *TDragImageList) EndDrag() bool {
	r1 := LCL().SysCallN(2739, m.Instance())
	return GoBool(r1)
}

func (m *TDragImageList) SetDragImage(Index, HotSpotX, HotSpotY int32) bool {
	r1 := LCL().SysCallN(2742, m.Instance(), uintptr(Index), uintptr(HotSpotX), uintptr(HotSpotY))
	return GoBool(r1)
}

func DragImageListClass() TClass {
	ret := LCL().SysCallN(2730)
	return TClass(ret)
}

func (m *TDragImageList) DragUnlock() {
	LCL().SysCallN(2736, m.Instance())
}

func (m *TDragImageList) HideDragImage() {
	LCL().SysCallN(2740, m.Instance())
}

func (m *TDragImageList) ShowDragImage() {
	LCL().SysCallN(2743, m.Instance())
}
