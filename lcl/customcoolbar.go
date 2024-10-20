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

// ICustomCoolBar Parent: IToolWindow
type ICustomCoolBar interface {
	IToolWindow
	BandBorderStyle() TBorderStyle                               // property
	SetBandBorderStyle(AValue TBorderStyle)                      // property
	BandMaximize() TCoolBandMaximize                             // property
	SetBandMaximize(AValue TCoolBandMaximize)                    // property
	Bands() ICoolBands                                           // property
	SetBands(AValue ICoolBands)                                  // property
	Bitmap() IBitmap                                             // property
	SetBitmap(AValue IBitmap)                                    // property
	FixedSize() bool                                             // property
	SetFixedSize(AValue bool)                                    // property
	FixedOrder() bool                                            // property
	SetFixedOrder(AValue bool)                                   // property
	GrabStyle() TGrabStyle                                       // property
	SetGrabStyle(AValue TGrabStyle)                              // property
	GrabWidth() int32                                            // property
	SetGrabWidth(AValue int32)                                   // property
	HorizontalSpacing() int32                                    // property
	SetHorizontalSpacing(AValue int32)                           // property
	Images() ICustomImageList                                    // property
	SetImages(AValue ICustomImageList)                           // property
	ImagesWidth() int32                                          // property
	SetImagesWidth(AValue int32)                                 // property
	ShowText() bool                                              // property
	SetShowText(AValue bool)                                     // property
	Themed() bool                                                // property
	SetThemed(AValue bool)                                       // property
	Vertical() bool                                              // property
	SetVertical(AValue bool)                                     // property
	VerticalSpacing() int32                                      // property
	SetVerticalSpacing(AValue int32)                             // property
	AutosizeBands()                                              // procedure
	MouseToBandPos(X, Y int32, OutBand *int32, OutGrabber *bool) // procedure
	SetOnChange(fn TNotifyEvent)                                 // property event
}

// TCustomCoolBar Parent: TToolWindow
type TCustomCoolBar struct {
	TToolWindow
	changePtr uintptr
}

func NewCustomCoolBar(AOwner IComponent) ICustomCoolBar {
	r1 := customCoolBarImportAPI().SysCallN(6, GetObjectUintptr(AOwner))
	return AsCustomCoolBar(r1)
}

func (m *TCustomCoolBar) BandBorderStyle() TBorderStyle {
	r1 := customCoolBarImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomCoolBar) SetBandBorderStyle(AValue TBorderStyle) {
	customCoolBarImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) BandMaximize() TCoolBandMaximize {
	r1 := customCoolBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCoolBandMaximize(r1)
}

func (m *TCustomCoolBar) SetBandMaximize(AValue TCoolBandMaximize) {
	customCoolBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) Bands() ICoolBands {
	r1 := customCoolBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsCoolBands(r1)
}

func (m *TCustomCoolBar) SetBands(AValue ICoolBands) {
	customCoolBarImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCoolBar) Bitmap() IBitmap {
	r1 := customCoolBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCustomCoolBar) SetBitmap(AValue IBitmap) {
	customCoolBarImportAPI().SysCallN(4, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCoolBar) FixedSize() bool {
	r1 := customCoolBarImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetFixedSize(AValue bool) {
	customCoolBarImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) FixedOrder() bool {
	r1 := customCoolBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetFixedOrder(AValue bool) {
	customCoolBarImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) GrabStyle() TGrabStyle {
	r1 := customCoolBarImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TGrabStyle(r1)
}

func (m *TCustomCoolBar) SetGrabStyle(AValue TGrabStyle) {
	customCoolBarImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) GrabWidth() int32 {
	r1 := customCoolBarImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetGrabWidth(AValue int32) {
	customCoolBarImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) HorizontalSpacing() int32 {
	r1 := customCoolBarImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetHorizontalSpacing(AValue int32) {
	customCoolBarImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) Images() ICustomImageList {
	r1 := customCoolBarImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomCoolBar) SetImages(AValue ICustomImageList) {
	customCoolBarImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCoolBar) ImagesWidth() int32 {
	r1 := customCoolBarImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetImagesWidth(AValue int32) {
	customCoolBarImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) ShowText() bool {
	r1 := customCoolBarImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetShowText(AValue bool) {
	customCoolBarImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) Themed() bool {
	r1 := customCoolBarImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetThemed(AValue bool) {
	customCoolBarImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) Vertical() bool {
	r1 := customCoolBarImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetVertical(AValue bool) {
	customCoolBarImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) VerticalSpacing() int32 {
	r1 := customCoolBarImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetVerticalSpacing(AValue int32) {
	customCoolBarImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func CustomCoolBarClass() TClass {
	ret := customCoolBarImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TCustomCoolBar) AutosizeBands() {
	customCoolBarImportAPI().SysCallN(0, m.Instance())
}

func (m *TCustomCoolBar) MouseToBandPos(X, Y int32, OutBand *int32, OutGrabber *bool) {
	var result1 uintptr
	var result2 uintptr
	customCoolBarImportAPI().SysCallN(14, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutBand = int32(result1)
	*OutGrabber = GoBool(result2)
}

func (m *TCustomCoolBar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customCoolBarImportAPI().SysCallN(15, m.Instance(), m.changePtr)
}

var (
	customCoolBarImport       *imports.Imports = nil
	customCoolBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCoolBar_AutosizeBands", 0),
		/*1*/ imports.NewTable("CustomCoolBar_BandBorderStyle", 0),
		/*2*/ imports.NewTable("CustomCoolBar_BandMaximize", 0),
		/*3*/ imports.NewTable("CustomCoolBar_Bands", 0),
		/*4*/ imports.NewTable("CustomCoolBar_Bitmap", 0),
		/*5*/ imports.NewTable("CustomCoolBar_Class", 0),
		/*6*/ imports.NewTable("CustomCoolBar_Create", 0),
		/*7*/ imports.NewTable("CustomCoolBar_FixedOrder", 0),
		/*8*/ imports.NewTable("CustomCoolBar_FixedSize", 0),
		/*9*/ imports.NewTable("CustomCoolBar_GrabStyle", 0),
		/*10*/ imports.NewTable("CustomCoolBar_GrabWidth", 0),
		/*11*/ imports.NewTable("CustomCoolBar_HorizontalSpacing", 0),
		/*12*/ imports.NewTable("CustomCoolBar_Images", 0),
		/*13*/ imports.NewTable("CustomCoolBar_ImagesWidth", 0),
		/*14*/ imports.NewTable("CustomCoolBar_MouseToBandPos", 0),
		/*15*/ imports.NewTable("CustomCoolBar_SetOnChange", 0),
		/*16*/ imports.NewTable("CustomCoolBar_ShowText", 0),
		/*17*/ imports.NewTable("CustomCoolBar_Themed", 0),
		/*18*/ imports.NewTable("CustomCoolBar_Vertical", 0),
		/*19*/ imports.NewTable("CustomCoolBar_VerticalSpacing", 0),
	}
)

func customCoolBarImportAPI() *imports.Imports {
	if customCoolBarImport == nil {
		customCoolBarImport = NewDefaultImports()
		customCoolBarImport.SetImportTable(customCoolBarImportTables)
		customCoolBarImportTables = nil
	}
	return customCoolBarImport
}
