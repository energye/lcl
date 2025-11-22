//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// IRGBAQuadArray = array of TRGBAQuad
type IRGBAQuadArray interface {
	base.IArraySlice[TRGBAQuad]
}

// IPointArray = array of TPoint
type IPointArray interface {
	base.IArraySlice[types.TPoint]
}

// IFormatArray = array of uintptr
type IFormatArray interface {
	base.IArraySlice[uintptr]
}

// IStringArray = array of string
type IStringArray interface {
	base.IArraySlice[uintptr]
	GetValue(index int) string
}

type TStringArray struct {
	base.IArraySlice[uintptr]
}

// IFormatEtcArray = array of TPoint
type IFormatEtcArray interface {
	base.IArraySlice[TFormatEtc]
}

// NewRGBAQuadArray
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewRGBAQuadArray(count int, instance uintptr) IRGBAQuadArray {
	return base.NewStructArraySlice[TRGBAQuad](count, instance)
}

// NewPointArray
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewPointArray(count int, instance uintptr) IPointArray {
	return base.NewStructArraySlice[types.TPoint](count, instance)
}

// NewFormatArray
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewFormatArray(count int, instance uintptr) IFormatArray {
	return base.NewStructArraySlice[uintptr](count, instance)
}

// NewFormatEtcArray
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewFormatEtcArray(count int, instance uintptr) IFormatEtcArray {
	return base.NewStructArraySlice[TFormatEtc](count, instance)
}

// NewStringArray
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewStringArray(count int, instance uintptr) IStringArray {
	if instance == 0 {
		count = 0
	}
	return &TStringArray{
		IArraySlice: base.NewStructArraySlice[uintptr](count, instance),
	}
}

func (m *TStringArray) GetValue(index int) string {
	ptr := m.Get(index)
	return api.GoStr(ptr)
}
