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

// IJPEGImage Parent: IFPImageBitmap
type IJPEGImage interface {
	IFPImageBitmap
	Compress()                                           // procedure
	CompressionQuality() types.TJPEGQualityRange         // property CompressionQuality Getter
	SetCompressionQuality(value types.TJPEGQualityRange) // property CompressionQuality Setter
	GrayScale() bool                                     // property GrayScale Getter
	SetGrayScale(value bool)                             // property GrayScale Setter
	MinHeight() int32                                    // property MinHeight Getter
	SetMinHeight(value int32)                            // property MinHeight Setter
	MinWidth() int32                                     // property MinWidth Getter
	SetMinWidth(value int32)                             // property MinWidth Setter
	ProgressiveEncoding() bool                           // property ProgressiveEncoding Getter
	SetProgressiveEncoding(value bool)                   // property ProgressiveEncoding Setter
	Performance() types.TJPEGPerformance                 // property Performance Getter
	SetPerformance(value types.TJPEGPerformance)         // property Performance Setter
	Scale() types.TJPEGScale                             // property Scale Getter
	SetScale(value types.TJPEGScale)                     // property Scale Setter
	Smoothing() bool                                     // property Smoothing Getter
	SetSmoothing(value bool)                             // property Smoothing Setter
}

type TJPEGImage struct {
	TFPImageBitmap
}

func (m *TJPEGImage) Compress() {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(1, m.Instance())
}

func (m *TJPEGImage) CompressionQuality() types.TJPEGQualityRange {
	if !m.IsValid() {
		return 0
	}
	r := jPEGImageAPI().SysCallN(2, 0, m.Instance())
	return types.TJPEGQualityRange(r)
}

func (m *TJPEGImage) SetCompressionQuality(value types.TJPEGQualityRange) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TJPEGImage) GrayScale() bool {
	if !m.IsValid() {
		return false
	}
	r := jPEGImageAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TJPEGImage) SetGrayScale(value bool) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TJPEGImage) MinHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := jPEGImageAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TJPEGImage) SetMinHeight(value int32) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TJPEGImage) MinWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := jPEGImageAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TJPEGImage) SetMinWidth(value int32) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TJPEGImage) ProgressiveEncoding() bool {
	if !m.IsValid() {
		return false
	}
	r := jPEGImageAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TJPEGImage) SetProgressiveEncoding(value bool) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TJPEGImage) Performance() types.TJPEGPerformance {
	if !m.IsValid() {
		return 0
	}
	r := jPEGImageAPI().SysCallN(7, 0, m.Instance())
	return types.TJPEGPerformance(r)
}

func (m *TJPEGImage) SetPerformance(value types.TJPEGPerformance) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TJPEGImage) Scale() types.TJPEGScale {
	if !m.IsValid() {
		return 0
	}
	r := jPEGImageAPI().SysCallN(8, 0, m.Instance())
	return types.TJPEGScale(r)
}

func (m *TJPEGImage) SetScale(value types.TJPEGScale) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TJPEGImage) Smoothing() bool {
	if !m.IsValid() {
		return false
	}
	r := jPEGImageAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TJPEGImage) SetSmoothing(value bool) {
	if !m.IsValid() {
		return
	}
	jPEGImageAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

// NewJPEGImage class constructor
func NewJPEGImage() IJPEGImage {
	r := jPEGImageAPI().SysCallN(0)
	return AsJPEGImage(r)
}

var (
	jPEGImageOnce   base.Once
	jPEGImageImport *imports.Imports = nil
)

func jPEGImageAPI() *imports.Imports {
	jPEGImageOnce.Do(func() {
		jPEGImageImport = api.NewDefaultImports()
		jPEGImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TJPEGImage_Create", 0), // constructor NewJPEGImage
			/* 1 */ imports.NewTable("TJPEGImage_Compress", 0), // procedure Compress
			/* 2 */ imports.NewTable("TJPEGImage_CompressionQuality", 0), // property CompressionQuality
			/* 3 */ imports.NewTable("TJPEGImage_GrayScale", 0), // property GrayScale
			/* 4 */ imports.NewTable("TJPEGImage_MinHeight", 0), // property MinHeight
			/* 5 */ imports.NewTable("TJPEGImage_MinWidth", 0), // property MinWidth
			/* 6 */ imports.NewTable("TJPEGImage_ProgressiveEncoding", 0), // property ProgressiveEncoding
			/* 7 */ imports.NewTable("TJPEGImage_Performance", 0), // property Performance
			/* 8 */ imports.NewTable("TJPEGImage_Scale", 0), // property Scale
			/* 9 */ imports.NewTable("TJPEGImage_Smoothing", 0), // property Smoothing
		}
	})
	return jPEGImageImport
}
