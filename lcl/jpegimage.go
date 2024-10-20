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

// IJPEGImage Parent: IFPImageBitmap
type IJPEGImage interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	CompressionQuality() TJPEGQualityRange          // property
	SetCompressionQuality(AValue TJPEGQualityRange) // property
	GrayScale() bool                                // property
	MinHeight() int32                               // property
	SetMinHeight(AValue int32)                      // property
	MinWidth() int32                                // property
	SetMinWidth(AValue int32)                       // property
	ProgressiveEncoding() bool                      // property
	SetProgressiveEncoding(AValue bool)             // property
	Performance() TJPEGPerformance                  // property
	SetPerformance(AValue TJPEGPerformance)         // property
	Scale() TJPEGScale                              // property
	SetScale(AValue TJPEGScale)                     // property
	Smoothing() bool                                // property
	SetSmoothing(AValue bool)                       // property
	Compress()                                      // procedure
}

// TJPEGImage Parent: TFPImageBitmap
type TJPEGImage struct {
	TFPImageBitmap
}

func NewJPEGImage() IJPEGImage {
	r1 := jPEGImageImportAPI().SysCallN(3)
	return AsJPEGImage(r1)
}

func (m *TJPEGImage) CompressionQuality() TJPEGQualityRange {
	r1 := jPEGImageImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TJPEGQualityRange(r1)
}

func (m *TJPEGImage) SetCompressionQuality(AValue TJPEGQualityRange) {
	jPEGImageImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) GrayScale() bool {
	r1 := jPEGImageImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TJPEGImage) MinHeight() int32 {
	r1 := jPEGImageImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TJPEGImage) SetMinHeight(AValue int32) {
	jPEGImageImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) MinWidth() int32 {
	r1 := jPEGImageImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TJPEGImage) SetMinWidth(AValue int32) {
	jPEGImageImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) ProgressiveEncoding() bool {
	r1 := jPEGImageImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TJPEGImage) SetProgressiveEncoding(AValue bool) {
	jPEGImageImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TJPEGImage) Performance() TJPEGPerformance {
	r1 := jPEGImageImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TJPEGPerformance(r1)
}

func (m *TJPEGImage) SetPerformance(AValue TJPEGPerformance) {
	jPEGImageImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) Scale() TJPEGScale {
	r1 := jPEGImageImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TJPEGScale(r1)
}

func (m *TJPEGImage) SetScale(AValue TJPEGScale) {
	jPEGImageImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) Smoothing() bool {
	r1 := jPEGImageImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TJPEGImage) SetSmoothing(AValue bool) {
	jPEGImageImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func JPEGImageClass() TClass {
	ret := jPEGImageImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TJPEGImage) Compress() {
	jPEGImageImportAPI().SysCallN(1, m.Instance())
}

var (
	jPEGImageImport       *imports.Imports = nil
	jPEGImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("JPEGImage_Class", 0),
		/*1*/ imports.NewTable("JPEGImage_Compress", 0),
		/*2*/ imports.NewTable("JPEGImage_CompressionQuality", 0),
		/*3*/ imports.NewTable("JPEGImage_Create", 0),
		/*4*/ imports.NewTable("JPEGImage_GrayScale", 0),
		/*5*/ imports.NewTable("JPEGImage_MinHeight", 0),
		/*6*/ imports.NewTable("JPEGImage_MinWidth", 0),
		/*7*/ imports.NewTable("JPEGImage_Performance", 0),
		/*8*/ imports.NewTable("JPEGImage_ProgressiveEncoding", 0),
		/*9*/ imports.NewTable("JPEGImage_Scale", 0),
		/*10*/ imports.NewTable("JPEGImage_Smoothing", 0),
	}
)

func jPEGImageImportAPI() *imports.Imports {
	if jPEGImageImport == nil {
		jPEGImageImport = NewDefaultImports()
		jPEGImageImport.SetImportTable(jPEGImageImportTables)
		jPEGImageImportTables = nil
	}
	return jPEGImageImport
}
