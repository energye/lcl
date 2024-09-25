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

// IVTDragImage Parent: IObject
// Class to manage header and tree drag image during a drag'n drop operation.
type IVTDragImage interface {
	IObject
	ColorKey() TColor                                                                                             // property
	SetColorKey(AValue TColor)                                                                                    // property
	Fade() bool                                                                                                   // property
	SetFade(AValue bool)                                                                                          // property
	MoveRestriction() TVTDragMoveRestriction                                                                      // property
	SetMoveRestriction(AValue TVTDragMoveRestriction)                                                             // property
	PostBlendBias() TVTBias                                                                                       // property
	SetPostBlendBias(AValue TVTBias)                                                                              // property
	PreBlendBias() TVTBias                                                                                        // property
	SetPreBlendBias(AValue TVTBias)                                                                               // property
	Transparency() TVTTransparency                                                                                // property
	SetTransparency(AValue TVTTransparency)                                                                       // property
	Visible() bool                                                                                                // property
	DragTo(P *TPoint, ForceRepaint bool) bool                                                                     // function
	GetDragImageRect() (resultRect TRect)                                                                         // function
	WillMove(P *TPoint) bool                                                                                      // function
	EndDrag()                                                                                                     // procedure
	HideDragImage()                                                                                               // procedure
	RecaptureBackground(Tree IBaseVirtualTree, R *TRect, VisibleRegion HRGN, CaptureNCArea, ReshowDragImage bool) // procedure
	ShowDragImage()                                                                                               // procedure
}

// TVTDragImage Parent: TObject
// Class to manage header and tree drag image during a drag'n drop operation.
type TVTDragImage struct {
	TObject
}

func NewVTDragImage(AOwner IBaseVirtualTree) IVTDragImage {
	r1 := LCL().SysCallN(5878, GetObjectUintptr(AOwner))
	return AsVTDragImage(r1)
}

func (m *TVTDragImage) ColorKey() TColor {
	r1 := LCL().SysCallN(5877, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTDragImage) SetColorKey(AValue TColor) {
	LCL().SysCallN(5877, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) Fade() bool {
	r1 := LCL().SysCallN(5881, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TVTDragImage) SetFade(AValue bool) {
	LCL().SysCallN(5881, 1, m.Instance(), PascalBool(AValue))
}

func (m *TVTDragImage) MoveRestriction() TVTDragMoveRestriction {
	r1 := LCL().SysCallN(5884, 0, m.Instance(), 0)
	return TVTDragMoveRestriction(r1)
}

func (m *TVTDragImage) SetMoveRestriction(AValue TVTDragMoveRestriction) {
	LCL().SysCallN(5884, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) PostBlendBias() TVTBias {
	r1 := LCL().SysCallN(5885, 0, m.Instance(), 0)
	return TVTBias(r1)
}

func (m *TVTDragImage) SetPostBlendBias(AValue TVTBias) {
	LCL().SysCallN(5885, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) PreBlendBias() TVTBias {
	r1 := LCL().SysCallN(5886, 0, m.Instance(), 0)
	return TVTBias(r1)
}

func (m *TVTDragImage) SetPreBlendBias(AValue TVTBias) {
	LCL().SysCallN(5886, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) Transparency() TVTTransparency {
	r1 := LCL().SysCallN(5889, 0, m.Instance(), 0)
	return TVTTransparency(r1)
}

func (m *TVTDragImage) SetTransparency(AValue TVTTransparency) {
	LCL().SysCallN(5889, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) Visible() bool {
	r1 := LCL().SysCallN(5890, m.Instance())
	return GoBool(r1)
}

func (m *TVTDragImage) DragTo(P *TPoint, ForceRepaint bool) bool {
	r1 := LCL().SysCallN(5879, m.Instance(), uintptr(unsafePointer(P)), PascalBool(ForceRepaint))
	return GoBool(r1)
}

func (m *TVTDragImage) GetDragImageRect() (resultRect TRect) {
	LCL().SysCallN(5882, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TVTDragImage) WillMove(P *TPoint) bool {
	r1 := LCL().SysCallN(5891, m.Instance(), uintptr(unsafePointer(P)))
	return GoBool(r1)
}

func VTDragImageClass() TClass {
	ret := LCL().SysCallN(5876)
	return TClass(ret)
}

func (m *TVTDragImage) EndDrag() {
	LCL().SysCallN(5880, m.Instance())
}

func (m *TVTDragImage) HideDragImage() {
	LCL().SysCallN(5883, m.Instance())
}

func (m *TVTDragImage) RecaptureBackground(Tree IBaseVirtualTree, R *TRect, VisibleRegion HRGN, CaptureNCArea, ReshowDragImage bool) {
	LCL().SysCallN(5887, m.Instance(), GetObjectUintptr(Tree), uintptr(unsafePointer(R)), uintptr(VisibleRegion), PascalBool(CaptureNCArea), PascalBool(ReshowDragImage))
}

func (m *TVTDragImage) ShowDragImage() {
	LCL().SysCallN(5888, m.Instance())
}
