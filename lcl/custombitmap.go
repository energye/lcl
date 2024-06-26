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

// ICustomBitmap Parent: IRasterImage
type ICustomBitmap interface {
	IRasterImage
	Handle() HBITMAP                        // property
	SetHandle(AValue HBITMAP)               // property
	HandleType() TBitmapHandleType          // property
	SetHandleType(AValue TBitmapHandleType) // property
	Monochrome() bool                       // property
	SetMonochrome(AValue bool)              // property
	ReleaseHandle() HBITMAP                 // function
	SetSize(AWidth, AHeight int32)          // procedure
}

// TCustomBitmap Parent: TRasterImage
type TCustomBitmap struct {
	TRasterImage
}

func NewCustomBitmap() ICustomBitmap {
	r1 := LCL().SysCallN(1327)
	return AsCustomBitmap(r1)
}

func (m *TCustomBitmap) Handle() HBITMAP {
	r1 := LCL().SysCallN(1328, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TCustomBitmap) SetHandle(AValue HBITMAP) {
	LCL().SysCallN(1328, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitmap) HandleType() TBitmapHandleType {
	r1 := LCL().SysCallN(1329, 0, m.Instance(), 0)
	return TBitmapHandleType(r1)
}

func (m *TCustomBitmap) SetHandleType(AValue TBitmapHandleType) {
	LCL().SysCallN(1329, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitmap) Monochrome() bool {
	r1 := LCL().SysCallN(1330, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomBitmap) SetMonochrome(AValue bool) {
	LCL().SysCallN(1330, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomBitmap) ReleaseHandle() HBITMAP {
	r1 := LCL().SysCallN(1331, m.Instance())
	return HBITMAP(r1)
}

func CustomBitmapClass() TClass {
	ret := LCL().SysCallN(1326)
	return TClass(ret)
}

func (m *TCustomBitmap) SetSize(AWidth, AHeight int32) {
	LCL().SysCallN(1332, m.Instance(), uintptr(AWidth), uintptr(AHeight))
}
