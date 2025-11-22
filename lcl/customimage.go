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

// ICustomImage Parent: IGraphicControl
type ICustomImage interface {
	IGraphicControl
	DestRect() types.TRect                              // function
	HasGraphic() bool                                   // property HasGraphic Getter
	AntialiasingMode() types.TAntialiasingMode          // property AntialiasingMode Getter
	SetAntialiasingMode(value types.TAntialiasingMode)  // property AntialiasingMode Setter
	Center() bool                                       // property Center Getter
	SetCenter(value bool)                               // property Center Setter
	KeepOriginXWhenClipped() bool                       // property KeepOriginXWhenClipped Getter
	SetKeepOriginXWhenClipped(value bool)               // property KeepOriginXWhenClipped Setter
	KeepOriginYWhenClipped() bool                       // property KeepOriginYWhenClipped Getter
	SetKeepOriginYWhenClipped(value bool)               // property KeepOriginYWhenClipped Setter
	ImageIndex() int32                                  // property ImageIndex Getter
	SetImageIndex(value int32)                          // property ImageIndex Setter
	ImageWidth() int32                                  // property ImageWidth Getter
	SetImageWidth(value int32)                          // property ImageWidth Setter
	Images() ICustomImageList                           // property Images Getter
	SetImages(value ICustomImageList)                   // property Images Setter
	Picture() IPicture                                  // property Picture Getter
	SetPicture(value IPicture)                          // property Picture Setter
	Stretch() bool                                      // property Stretch Getter
	SetStretch(value bool)                              // property Stretch Setter
	StretchOutEnabled() bool                            // property StretchOutEnabled Getter
	SetStretchOutEnabled(value bool)                    // property StretchOutEnabled Setter
	StretchInEnabled() bool                             // property StretchInEnabled Getter
	SetStretchInEnabled(value bool)                     // property StretchInEnabled Setter
	Transparent() bool                                  // property Transparent Getter
	SetTransparent(value bool)                          // property Transparent Setter
	Proportional() bool                                 // property Proportional Getter
	SetProportional(value bool)                         // property Proportional Setter
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

type TCustomImage struct {
	TGraphicControl
}

func (m *TCustomImage) DestRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomImage) HasGraphic() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) AntialiasingMode() types.TAntialiasingMode {
	if !m.IsValid() {
		return 0
	}
	r := customImageAPI().SysCallN(3, 0, m.Instance())
	return types.TAntialiasingMode(r)
}

func (m *TCustomImage) SetAntialiasingMode(value types.TAntialiasingMode) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImage) Center() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetCenter(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) KeepOriginXWhenClipped() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetKeepOriginXWhenClipped(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) KeepOriginYWhenClipped() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetKeepOriginYWhenClipped(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCustomImage) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImage) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCustomImage) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImage) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customImageAPI().SysCallN(9, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomImage) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomImage) Picture() IPicture {
	if !m.IsValid() {
		return nil
	}
	r := customImageAPI().SysCallN(10, 0, m.Instance())
	return AsPicture(r)
}

func (m *TCustomImage) SetPicture(value IPicture) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomImage) Stretch() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetStretch(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) StretchOutEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetStretchOutEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) StretchInEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetStretchInEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) Proportional() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImage) SetProportional(value bool) {
	if !m.IsValid() {
		return
	}
	customImageAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImage) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 16, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 19, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 20, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 21, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 22, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 23, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnPictureChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, customImageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImage) SetOnPaintBackground(fn TImagePaintBackgroundEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTImagePaintBackgroundEvent(fn)
	base.SetEvent(m, 25, customImageAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomImage class constructor
func NewCustomImage(owner IComponent) ICustomImage {
	r := customImageAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomImage(r)
}

func TCustomImageClass() types.TClass {
	r := customImageAPI().SysCallN(26)
	return types.TClass(r)
}

var (
	customImageOnce   base.Once
	customImageImport *imports.Imports = nil
)

func customImageAPI() *imports.Imports {
	customImageOnce.Do(func() {
		customImageImport = api.NewDefaultImports()
		customImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomImage_Create", 0), // constructor NewCustomImage
			/* 1 */ imports.NewTable("TCustomImage_DestRect", 0), // function DestRect
			/* 2 */ imports.NewTable("TCustomImage_HasGraphic", 0), // property HasGraphic
			/* 3 */ imports.NewTable("TCustomImage_AntialiasingMode", 0), // property AntialiasingMode
			/* 4 */ imports.NewTable("TCustomImage_Center", 0), // property Center
			/* 5 */ imports.NewTable("TCustomImage_KeepOriginXWhenClipped", 0), // property KeepOriginXWhenClipped
			/* 6 */ imports.NewTable("TCustomImage_KeepOriginYWhenClipped", 0), // property KeepOriginYWhenClipped
			/* 7 */ imports.NewTable("TCustomImage_ImageIndex", 0), // property ImageIndex
			/* 8 */ imports.NewTable("TCustomImage_ImageWidth", 0), // property ImageWidth
			/* 9 */ imports.NewTable("TCustomImage_Images", 0), // property Images
			/* 10 */ imports.NewTable("TCustomImage_Picture", 0), // property Picture
			/* 11 */ imports.NewTable("TCustomImage_Stretch", 0), // property Stretch
			/* 12 */ imports.NewTable("TCustomImage_StretchOutEnabled", 0), // property StretchOutEnabled
			/* 13 */ imports.NewTable("TCustomImage_StretchInEnabled", 0), // property StretchInEnabled
			/* 14 */ imports.NewTable("TCustomImage_Transparent", 0), // property Transparent
			/* 15 */ imports.NewTable("TCustomImage_Proportional", 0), // property Proportional
			/* 16 */ imports.NewTable("TCustomImage_OnMouseDown", 0), // event OnMouseDown
			/* 17 */ imports.NewTable("TCustomImage_OnMouseEnter", 0), // event OnMouseEnter
			/* 18 */ imports.NewTable("TCustomImage_OnMouseLeave", 0), // event OnMouseLeave
			/* 19 */ imports.NewTable("TCustomImage_OnMouseMove", 0), // event OnMouseMove
			/* 20 */ imports.NewTable("TCustomImage_OnMouseUp", 0), // event OnMouseUp
			/* 21 */ imports.NewTable("TCustomImage_OnMouseWheel", 0), // event OnMouseWheel
			/* 22 */ imports.NewTable("TCustomImage_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 23 */ imports.NewTable("TCustomImage_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 24 */ imports.NewTable("TCustomImage_OnPictureChanged", 0), // event OnPictureChanged
			/* 25 */ imports.NewTable("TCustomImage_OnPaintBackground", 0), // event OnPaintBackground
			/* 26 */ imports.NewTable("TCustomImage_TClass", 0), // function TCustomImageClass
		}
	})
	return customImageImport
}
