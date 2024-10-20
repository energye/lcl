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
	r1 := vTDragImageImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsVTDragImage(r1)
}

func (m *TVTDragImage) ColorKey() TColor {
	r1 := vTDragImageImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTDragImage) SetColorKey(AValue TColor) {
	vTDragImageImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) Fade() bool {
	r1 := vTDragImageImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TVTDragImage) SetFade(AValue bool) {
	vTDragImageImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TVTDragImage) MoveRestriction() TVTDragMoveRestriction {
	r1 := vTDragImageImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TVTDragMoveRestriction(r1)
}

func (m *TVTDragImage) SetMoveRestriction(AValue TVTDragMoveRestriction) {
	vTDragImageImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) PostBlendBias() TVTBias {
	r1 := vTDragImageImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TVTBias(r1)
}

func (m *TVTDragImage) SetPostBlendBias(AValue TVTBias) {
	vTDragImageImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) PreBlendBias() TVTBias {
	r1 := vTDragImageImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TVTBias(r1)
}

func (m *TVTDragImage) SetPreBlendBias(AValue TVTBias) {
	vTDragImageImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) Transparency() TVTTransparency {
	r1 := vTDragImageImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TVTTransparency(r1)
}

func (m *TVTDragImage) SetTransparency(AValue TVTTransparency) {
	vTDragImageImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTDragImage) Visible() bool {
	r1 := vTDragImageImportAPI().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TVTDragImage) DragTo(P *TPoint, ForceRepaint bool) bool {
	r1 := vTDragImageImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(P)), PascalBool(ForceRepaint))
	return GoBool(r1)
}

func (m *TVTDragImage) GetDragImageRect() (resultRect TRect) {
	vTDragImageImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TVTDragImage) WillMove(P *TPoint) bool {
	r1 := vTDragImageImportAPI().SysCallN(15, m.Instance(), uintptr(unsafePointer(P)))
	return GoBool(r1)
}

func VTDragImageClass() TClass {
	ret := vTDragImageImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TVTDragImage) EndDrag() {
	vTDragImageImportAPI().SysCallN(4, m.Instance())
}

func (m *TVTDragImage) HideDragImage() {
	vTDragImageImportAPI().SysCallN(7, m.Instance())
}

func (m *TVTDragImage) RecaptureBackground(Tree IBaseVirtualTree, R *TRect, VisibleRegion HRGN, CaptureNCArea, ReshowDragImage bool) {
	vTDragImageImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(Tree), uintptr(unsafePointer(R)), uintptr(VisibleRegion), PascalBool(CaptureNCArea), PascalBool(ReshowDragImage))
}

func (m *TVTDragImage) ShowDragImage() {
	vTDragImageImportAPI().SysCallN(12, m.Instance())
}

var (
	vTDragImageImport       *imports.Imports = nil
	vTDragImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VTDragImage_Class", 0),
		/*1*/ imports.NewTable("VTDragImage_ColorKey", 0),
		/*2*/ imports.NewTable("VTDragImage_Create", 0),
		/*3*/ imports.NewTable("VTDragImage_DragTo", 0),
		/*4*/ imports.NewTable("VTDragImage_EndDrag", 0),
		/*5*/ imports.NewTable("VTDragImage_Fade", 0),
		/*6*/ imports.NewTable("VTDragImage_GetDragImageRect", 0),
		/*7*/ imports.NewTable("VTDragImage_HideDragImage", 0),
		/*8*/ imports.NewTable("VTDragImage_MoveRestriction", 0),
		/*9*/ imports.NewTable("VTDragImage_PostBlendBias", 0),
		/*10*/ imports.NewTable("VTDragImage_PreBlendBias", 0),
		/*11*/ imports.NewTable("VTDragImage_RecaptureBackground", 0),
		/*12*/ imports.NewTable("VTDragImage_ShowDragImage", 0),
		/*13*/ imports.NewTable("VTDragImage_Transparency", 0),
		/*14*/ imports.NewTable("VTDragImage_Visible", 0),
		/*15*/ imports.NewTable("VTDragImage_WillMove", 0),
	}
)

func vTDragImageImportAPI() *imports.Imports {
	if vTDragImageImport == nil {
		vTDragImageImport = NewDefaultImports()
		vTDragImageImport.SetImportTable(vTDragImageImportTables)
		vTDragImageImportTables = nil
	}
	return vTDragImageImport
}
