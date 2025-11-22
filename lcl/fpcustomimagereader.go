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

// IFPCustomImageReader Parent: IFPCustomImageHandler
type IFPCustomImageReader interface {
	IFPCustomImageHandler
	ImageRead(str IStream, img IFPCustomImage) IFPCustomImage // function
	// CheckContents
	//  reads image
	CheckContents(str IStream) bool // function
	// DefaultImageClass
	//  returns the size of image in stream without loading it completely. -1,-1 means this is not implemented.
	DefaultImageClass() types.TFPCustomImageClass         // property DefaultImageClass Getter
	SetDefaultImageClass(value types.TFPCustomImageClass) // property DefaultImageClass Setter
}

type TFPCustomImageReader struct {
	TFPCustomImageHandler
}

func (m *TFPCustomImageReader) ImageRead(str IStream, img IFPCustomImage) IFPCustomImage {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomImageReaderAPI().SysCallN(0, m.Instance(), base.GetObjectUintptr(str), base.GetObjectUintptr(img))
	return AsFPCustomImage(r)
}

func (m *TFPCustomImageReader) CheckContents(str IStream) bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomImageReaderAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(str))
	return api.GoBool(r)
}

func (m *TFPCustomImageReader) DefaultImageClass() types.TFPCustomImageClass {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomImageReaderAPI().SysCallN(3, 0, m.Instance())
	return types.TFPCustomImageClass(r)
}

func (m *TFPCustomImageReader) SetDefaultImageClass(value types.TFPCustomImageClass) {
	if !m.IsValid() {
		return
	}
	fPCustomImageReaderAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// FPCustomImageReader  is static instance
var FPCustomImageReader _FPCustomImageReaderClass

// _FPCustomImageReaderClass is class type defined by TFPCustomImageReader
type _FPCustomImageReaderClass uintptr

// ImageSize
//
//	Returns true if the content is readable
func (_FPCustomImageReaderClass) ImageSize(str IStream) (result types.TPoint) {
	fPCustomImageReaderAPI().SysCallN(2, base.GetObjectUintptr(str), uintptr(base.UnsafePointer(&result)))
	return
}

var (
	fPCustomImageReaderOnce   base.Once
	fPCustomImageReaderImport *imports.Imports = nil
)

func fPCustomImageReaderAPI() *imports.Imports {
	fPCustomImageReaderOnce.Do(func() {
		fPCustomImageReaderImport = api.NewDefaultImports()
		fPCustomImageReaderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomImageReader_ImageRead", 0), // function ImageRead
			/* 1 */ imports.NewTable("TFPCustomImageReader_CheckContents", 0), // function CheckContents
			/* 2 */ imports.NewTable("TFPCustomImageReader_ImageSize", 0), // static function ImageSize
			/* 3 */ imports.NewTable("TFPCustomImageReader_DefaultImageClass", 0), // property DefaultImageClass
		}
	})
	return fPCustomImageReaderImport
}
