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

// ICustomIcon Parent: IRasterImage
type ICustomIcon interface {
	IRasterImage
	GetBestIndexForSize(size types.TSize) int32                                                      // function
	ExportImage(index int32, imageClass types.TFPImageBitmapClass) IFPImageBitmap                    // function
	Add(format types.TPixelFormat, height uint16, width uint16)                                      // procedure
	AssignImage(source IRasterImage)                                                                 // procedure
	Delete(aindex int32)                                                                             // procedure
	Remove(format types.TPixelFormat, height uint16, width uint16)                                   // procedure
	GetDescription(aindex int32, outFormat *types.TPixelFormat, outHeight *uint16, outWidth *uint16) // procedure
	SetSize(width int32, height int32)                                                               // procedure
	LoadFromResourceHandle(instance types.TLCLHandle, resHandle uintptr)                             // procedure
	Sort()                                                                                           // procedure
	Current() int32                                                                                  // property Current Getter
	SetCurrent(value int32)                                                                          // property Current Setter
	Count() int32                                                                                    // property Count Getter
}

type TCustomIcon struct {
	TRasterImage
}

func (m *TCustomIcon) GetBestIndexForSize(size types.TSize) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customIconAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&size)))
	return int32(r)
}

func (m *TCustomIcon) ExportImage(index int32, imageClass types.TFPImageBitmapClass) IFPImageBitmap {
	if !m.IsValid() {
		return nil
	}
	r := customIconAPI().SysCallN(2, m.Instance(), uintptr(index), uintptr(imageClass))
	return AsFPImageBitmap(r)
}

func (m *TCustomIcon) Add(format types.TPixelFormat, height uint16, width uint16) {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(3, m.Instance(), uintptr(format), uintptr(height), uintptr(width))
}

func (m *TCustomIcon) AssignImage(source IRasterImage) {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(source))
}

func (m *TCustomIcon) Delete(aindex int32) {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(5, m.Instance(), uintptr(aindex))
}

func (m *TCustomIcon) Remove(format types.TPixelFormat, height uint16, width uint16) {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(6, m.Instance(), uintptr(format), uintptr(height), uintptr(width))
}

func (m *TCustomIcon) GetDescription(aindex int32, outFormat *types.TPixelFormat, outHeight *uint16, outWidth *uint16) {
	if !m.IsValid() {
		return
	}
	var formatPtr uintptr
	var heightPtr uintptr
	var widthPtr uintptr
	customIconAPI().SysCallN(7, m.Instance(), uintptr(aindex), uintptr(base.UnsafePointer(&formatPtr)), uintptr(base.UnsafePointer(&heightPtr)), uintptr(base.UnsafePointer(&widthPtr)))
	*outFormat = types.TPixelFormat(formatPtr)
	*outHeight = uint16(heightPtr)
	*outWidth = uint16(widthPtr)
}

func (m *TCustomIcon) SetSize(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(8, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TCustomIcon) LoadFromResourceHandle(instance types.TLCLHandle, resHandle uintptr) {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(9, m.Instance(), uintptr(instance), uintptr(resHandle))
}

func (m *TCustomIcon) Sort() {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(10, m.Instance())
}

func (m *TCustomIcon) Current() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customIconAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TCustomIcon) SetCurrent(value int32) {
	if !m.IsValid() {
		return
	}
	customIconAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomIcon) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customIconAPI().SysCallN(12, m.Instance())
	return int32(r)
}

// NewCustomIcon class constructor
func NewCustomIcon() ICustomIcon {
	r := customIconAPI().SysCallN(0)
	return AsCustomIcon(r)
}

var (
	customIconOnce   base.Once
	customIconImport *imports.Imports = nil
)

func customIconAPI() *imports.Imports {
	customIconOnce.Do(func() {
		customIconImport = api.NewDefaultImports()
		customIconImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomIcon_Create", 0), // constructor NewCustomIcon
			/* 1 */ imports.NewTable("TCustomIcon_GetBestIndexForSize", 0), // function GetBestIndexForSize
			/* 2 */ imports.NewTable("TCustomIcon_ExportImage", 0), // function ExportImage
			/* 3 */ imports.NewTable("TCustomIcon_Add", 0), // procedure Add
			/* 4 */ imports.NewTable("TCustomIcon_AssignImage", 0), // procedure AssignImage
			/* 5 */ imports.NewTable("TCustomIcon_Delete", 0), // procedure Delete
			/* 6 */ imports.NewTable("TCustomIcon_Remove", 0), // procedure Remove
			/* 7 */ imports.NewTable("TCustomIcon_GetDescription", 0), // procedure GetDescription
			/* 8 */ imports.NewTable("TCustomIcon_SetSize", 0), // procedure SetSize
			/* 9 */ imports.NewTable("TCustomIcon_LoadFromResourceHandle", 0), // procedure LoadFromResourceHandle
			/* 10 */ imports.NewTable("TCustomIcon_Sort", 0), // procedure Sort
			/* 11 */ imports.NewTable("TCustomIcon_Current", 0), // property Current
			/* 12 */ imports.NewTable("TCustomIcon_Count", 0), // property Count
		}
	})
	return customIconImport
}
