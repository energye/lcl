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

// ICustomCoolBar Parent: IToolWindow
type ICustomCoolBar interface {
	IToolWindow
	AutosizeBands()                                                    // procedure
	MouseToBandPos(X int32, Y int32, outBand *int32, outGrabber *bool) // procedure
	BandBorderStyle() types.TBorderStyle                               // property BandBorderStyle Getter
	SetBandBorderStyle(value types.TBorderStyle)                       // property BandBorderStyle Setter
	BandMaximize() types.TCoolBandMaximize                             // property BandMaximize Getter
	SetBandMaximize(value types.TCoolBandMaximize)                     // property BandMaximize Setter
	Bands() ICoolBands                                                 // property Bands Getter
	SetBands(value ICoolBands)                                         // property Bands Setter
	Bitmap() IBitmap                                                   // property Bitmap Getter
	SetBitmap(value IBitmap)                                           // property Bitmap Setter
	FixedSize() bool                                                   // property FixedSize Getter
	SetFixedSize(value bool)                                           // property FixedSize Setter
	FixedOrder() bool                                                  // property FixedOrder Getter
	SetFixedOrder(value bool)                                          // property FixedOrder Setter
	GrabStyle() types.TGrabStyle                                       // property GrabStyle Getter
	SetGrabStyle(value types.TGrabStyle)                               // property GrabStyle Setter
	GrabWidth() int32                                                  // property GrabWidth Getter
	SetGrabWidth(value int32)                                          // property GrabWidth Setter
	HorizontalSpacing() int32                                          // property HorizontalSpacing Getter
	SetHorizontalSpacing(value int32)                                  // property HorizontalSpacing Setter
	Images() ICustomImageList                                          // property Images Getter
	SetImages(value ICustomImageList)                                  // property Images Setter
	ImagesWidth() int32                                                // property ImagesWidth Getter
	SetImagesWidth(value int32)                                        // property ImagesWidth Setter
	ShowText() bool                                                    // property ShowText Getter
	SetShowText(value bool)                                            // property ShowText Setter
	Themed() bool                                                      // property Themed Getter
	SetThemed(value bool)                                              // property Themed Setter
	Vertical() bool                                                    // property Vertical Getter
	SetVertical(value bool)                                            // property Vertical Setter
	VerticalSpacing() int32                                            // property VerticalSpacing Getter
	SetVerticalSpacing(value int32)                                    // property VerticalSpacing Setter
	SetOnChange(fn TNotifyEvent)                                       // property event
}

type TCustomCoolBar struct {
	TToolWindow
}

func (m *TCustomCoolBar) AutosizeBands() {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(1, m.Instance())
}

func (m *TCustomCoolBar) MouseToBandPos(X int32, Y int32, outBand *int32, outGrabber *bool) {
	if !m.IsValid() {
		return
	}
	var bandPtr uintptr
	customCoolBarAPI().SysCallN(2, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&bandPtr)), uintptr(base.UnsafePointer(outGrabber)))
	*outBand = int32(bandPtr)
}

func (m *TCustomCoolBar) BandBorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customCoolBarAPI().SysCallN(3, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCustomCoolBar) SetBandBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCoolBar) BandMaximize() types.TCoolBandMaximize {
	if !m.IsValid() {
		return 0
	}
	r := customCoolBarAPI().SysCallN(4, 0, m.Instance())
	return types.TCoolBandMaximize(r)
}

func (m *TCustomCoolBar) SetBandMaximize(value types.TCoolBandMaximize) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCoolBar) Bands() ICoolBands {
	if !m.IsValid() {
		return nil
	}
	r := customCoolBarAPI().SysCallN(5, 0, m.Instance())
	return AsCoolBands(r)
}

func (m *TCustomCoolBar) SetBands(value ICoolBands) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomCoolBar) Bitmap() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := customCoolBarAPI().SysCallN(6, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TCustomCoolBar) SetBitmap(value IBitmap) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomCoolBar) FixedSize() bool {
	if !m.IsValid() {
		return false
	}
	r := customCoolBarAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCoolBar) SetFixedSize(value bool) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCoolBar) FixedOrder() bool {
	if !m.IsValid() {
		return false
	}
	r := customCoolBarAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCoolBar) SetFixedOrder(value bool) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCoolBar) GrabStyle() types.TGrabStyle {
	if !m.IsValid() {
		return 0
	}
	r := customCoolBarAPI().SysCallN(9, 0, m.Instance())
	return types.TGrabStyle(r)
}

func (m *TCustomCoolBar) SetGrabStyle(value types.TGrabStyle) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCoolBar) GrabWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCoolBarAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCustomCoolBar) SetGrabWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCoolBar) HorizontalSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCoolBarAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TCustomCoolBar) SetHorizontalSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCoolBar) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customCoolBarAPI().SysCallN(12, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomCoolBar) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomCoolBar) ImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCoolBarAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TCustomCoolBar) SetImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCoolBar) ShowText() bool {
	if !m.IsValid() {
		return false
	}
	r := customCoolBarAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCoolBar) SetShowText(value bool) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCoolBar) Themed() bool {
	if !m.IsValid() {
		return false
	}
	r := customCoolBarAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCoolBar) SetThemed(value bool) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCoolBar) Vertical() bool {
	if !m.IsValid() {
		return false
	}
	r := customCoolBarAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCoolBar) SetVertical(value bool) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCoolBar) VerticalSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCoolBarAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TCustomCoolBar) SetVerticalSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	customCoolBarAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCoolBar) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, customCoolBarAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomCoolBar class constructor
func NewCustomCoolBar(owner IComponent) ICustomCoolBar {
	r := customCoolBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomCoolBar(r)
}

func TCustomCoolBarClass() types.TClass {
	r := customCoolBarAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	customCoolBarOnce   base.Once
	customCoolBarImport *imports.Imports = nil
)

func customCoolBarAPI() *imports.Imports {
	customCoolBarOnce.Do(func() {
		customCoolBarImport = api.NewDefaultImports()
		customCoolBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCoolBar_Create", 0), // constructor NewCustomCoolBar
			/* 1 */ imports.NewTable("TCustomCoolBar_AutosizeBands", 0), // procedure AutosizeBands
			/* 2 */ imports.NewTable("TCustomCoolBar_MouseToBandPos", 0), // procedure MouseToBandPos
			/* 3 */ imports.NewTable("TCustomCoolBar_BandBorderStyle", 0), // property BandBorderStyle
			/* 4 */ imports.NewTable("TCustomCoolBar_BandMaximize", 0), // property BandMaximize
			/* 5 */ imports.NewTable("TCustomCoolBar_Bands", 0), // property Bands
			/* 6 */ imports.NewTable("TCustomCoolBar_Bitmap", 0), // property Bitmap
			/* 7 */ imports.NewTable("TCustomCoolBar_FixedSize", 0), // property FixedSize
			/* 8 */ imports.NewTable("TCustomCoolBar_FixedOrder", 0), // property FixedOrder
			/* 9 */ imports.NewTable("TCustomCoolBar_GrabStyle", 0), // property GrabStyle
			/* 10 */ imports.NewTable("TCustomCoolBar_GrabWidth", 0), // property GrabWidth
			/* 11 */ imports.NewTable("TCustomCoolBar_HorizontalSpacing", 0), // property HorizontalSpacing
			/* 12 */ imports.NewTable("TCustomCoolBar_Images", 0), // property Images
			/* 13 */ imports.NewTable("TCustomCoolBar_ImagesWidth", 0), // property ImagesWidth
			/* 14 */ imports.NewTable("TCustomCoolBar_ShowText", 0), // property ShowText
			/* 15 */ imports.NewTable("TCustomCoolBar_Themed", 0), // property Themed
			/* 16 */ imports.NewTable("TCustomCoolBar_Vertical", 0), // property Vertical
			/* 17 */ imports.NewTable("TCustomCoolBar_VerticalSpacing", 0), // property VerticalSpacing
			/* 18 */ imports.NewTable("TCustomCoolBar_OnChange", 0), // event OnChange
			/* 19 */ imports.NewTable("TCustomCoolBar_TClass", 0), // function TCustomCoolBarClass
		}
	})
	return customCoolBarImport
}
