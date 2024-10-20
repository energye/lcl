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

// ICustomImage Parent: IGraphicControl
type ICustomImage interface {
	IGraphicControl
	HasGraphic() bool                                   // property
	AntialiasingMode() TAntialiasingMode                // property
	SetAntialiasingMode(AValue TAntialiasingMode)       // property
	Center() bool                                       // property
	SetCenter(AValue bool)                              // property
	KeepOriginXWhenClipped() bool                       // property
	SetKeepOriginXWhenClipped(AValue bool)              // property
	KeepOriginYWhenClipped() bool                       // property
	SetKeepOriginYWhenClipped(AValue bool)              // property
	ImageIndex() int32                                  // property
	SetImageIndex(AValue int32)                         // property
	ImageWidth() int32                                  // property
	SetImageWidth(AValue int32)                         // property
	Images() ICustomImageList                           // property
	SetImages(AValue ICustomImageList)                  // property
	Picture() IPicture                                  // property
	SetPicture(AValue IPicture)                         // property
	Stretch() bool                                      // property
	SetStretch(AValue bool)                             // property
	StretchOutEnabled() bool                            // property
	SetStretchOutEnabled(AValue bool)                   // property
	StretchInEnabled() bool                             // property
	SetStretchInEnabled(AValue bool)                    // property
	Transparent() bool                                  // property
	SetTransparent(AValue bool)                         // property
	Proportional() bool                                 // property
	SetProportional(AValue bool)                        // property
	DestRect() (resultRect TRect)                       // function
	SetOnMouseDown(fn TMouseEvent)                      // property event
	SetOnMouseEnter(fn TNotifyEvent)                    // property event
	SetOnMouseLeave(fn TNotifyEvent)                    // property event
	SetOnMouseMove(fn TMouseMoveEvent)                  // property event
	SetOnMouseUp(fn TMouseEvent)                        // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)      // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)        // property event
	SetOnPictureChanged(fn TNotifyEvent)                // property event
	SetOnPaintBackground(fn TImagePaintBackgroundEvent) // property event
}

// TCustomImage Parent: TGraphicControl
type TCustomImage struct {
	TGraphicControl
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	pictureChangedPtr  uintptr
	paintBackgroundPtr uintptr
}

func NewCustomImage(AOwner IComponent) ICustomImage {
	r1 := customImageImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsCustomImage(r1)
}

func (m *TCustomImage) HasGraphic() bool {
	r1 := customImageImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TCustomImage) AntialiasingMode() TAntialiasingMode {
	r1 := customImageImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAntialiasingMode(r1)
}

func (m *TCustomImage) SetAntialiasingMode(AValue TAntialiasingMode) {
	customImageImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImage) Center() bool {
	r1 := customImageImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetCenter(AValue bool) {
	customImageImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) KeepOriginXWhenClipped() bool {
	r1 := customImageImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetKeepOriginXWhenClipped(AValue bool) {
	customImageImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) KeepOriginYWhenClipped() bool {
	r1 := customImageImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetKeepOriginYWhenClipped(AValue bool) {
	customImageImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) ImageIndex() int32 {
	r1 := customImageImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImage) SetImageIndex(AValue int32) {
	customImageImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImage) ImageWidth() int32 {
	r1 := customImageImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImage) SetImageWidth(AValue int32) {
	customImageImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImage) Images() ICustomImageList {
	r1 := customImageImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomImage) SetImages(AValue ICustomImageList) {
	customImageImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomImage) Picture() IPicture {
	r1 := customImageImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TCustomImage) SetPicture(AValue IPicture) {
	customImageImportAPI().SysCallN(11, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomImage) Stretch() bool {
	r1 := customImageImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetStretch(AValue bool) {
	customImageImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) StretchOutEnabled() bool {
	r1 := customImageImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetStretchOutEnabled(AValue bool) {
	customImageImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) StretchInEnabled() bool {
	r1 := customImageImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetStretchInEnabled(AValue bool) {
	customImageImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) Transparent() bool {
	r1 := customImageImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetTransparent(AValue bool) {
	customImageImportAPI().SysCallN(26, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) Proportional() bool {
	r1 := customImageImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetProportional(AValue bool) {
	customImageImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) DestRect() (resultRect TRect) {
	customImageImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func CustomImageClass() TClass {
	ret := customImageImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomImage) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(13, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomImage) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(14, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomImage) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(15, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomImage) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(16, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomImage) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(17, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomImage) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(18, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomImage) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(19, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomImage) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(20, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomImage) SetOnPictureChanged(fn TNotifyEvent) {
	if m.pictureChangedPtr != 0 {
		RemoveEventElement(m.pictureChangedPtr)
	}
	m.pictureChangedPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(22, m.Instance(), m.pictureChangedPtr)
}

func (m *TCustomImage) SetOnPaintBackground(fn TImagePaintBackgroundEvent) {
	if m.paintBackgroundPtr != 0 {
		RemoveEventElement(m.paintBackgroundPtr)
	}
	m.paintBackgroundPtr = MakeEventDataPtr(fn)
	customImageImportAPI().SysCallN(21, m.Instance(), m.paintBackgroundPtr)
}

var (
	customImageImport       *imports.Imports = nil
	customImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomImage_AntialiasingMode", 0),
		/*1*/ imports.NewTable("CustomImage_Center", 0),
		/*2*/ imports.NewTable("CustomImage_Class", 0),
		/*3*/ imports.NewTable("CustomImage_Create", 0),
		/*4*/ imports.NewTable("CustomImage_DestRect", 0),
		/*5*/ imports.NewTable("CustomImage_HasGraphic", 0),
		/*6*/ imports.NewTable("CustomImage_ImageIndex", 0),
		/*7*/ imports.NewTable("CustomImage_ImageWidth", 0),
		/*8*/ imports.NewTable("CustomImage_Images", 0),
		/*9*/ imports.NewTable("CustomImage_KeepOriginXWhenClipped", 0),
		/*10*/ imports.NewTable("CustomImage_KeepOriginYWhenClipped", 0),
		/*11*/ imports.NewTable("CustomImage_Picture", 0),
		/*12*/ imports.NewTable("CustomImage_Proportional", 0),
		/*13*/ imports.NewTable("CustomImage_SetOnMouseDown", 0),
		/*14*/ imports.NewTable("CustomImage_SetOnMouseEnter", 0),
		/*15*/ imports.NewTable("CustomImage_SetOnMouseLeave", 0),
		/*16*/ imports.NewTable("CustomImage_SetOnMouseMove", 0),
		/*17*/ imports.NewTable("CustomImage_SetOnMouseUp", 0),
		/*18*/ imports.NewTable("CustomImage_SetOnMouseWheel", 0),
		/*19*/ imports.NewTable("CustomImage_SetOnMouseWheelDown", 0),
		/*20*/ imports.NewTable("CustomImage_SetOnMouseWheelUp", 0),
		/*21*/ imports.NewTable("CustomImage_SetOnPaintBackground", 0),
		/*22*/ imports.NewTable("CustomImage_SetOnPictureChanged", 0),
		/*23*/ imports.NewTable("CustomImage_Stretch", 0),
		/*24*/ imports.NewTable("CustomImage_StretchInEnabled", 0),
		/*25*/ imports.NewTable("CustomImage_StretchOutEnabled", 0),
		/*26*/ imports.NewTable("CustomImage_Transparent", 0),
	}
)

func customImageImportAPI() *imports.Imports {
	if customImageImport == nil {
		customImageImport = NewDefaultImports()
		customImageImport.SetImportTable(customImageImportTables)
		customImageImportTables = nil
	}
	return customImageImport
}
