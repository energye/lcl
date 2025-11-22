//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

import "unsafe"

type TLMessage = TMessage

type WindowPos struct {
	Hwnd            THandle
	HwndInsertAfter THandle
	X               Integer
	Y               Integer
	Cx              Integer
	Cy              Integer
	Flags           Cardinal
}

// TPaintStruct
type TPaintStruct struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     TRect
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]uint8
}

// PPaintStruct = ^TPaintStruct
type PPaintStruct uintptr

// ToGo convert TPaintStruct
func (m PPaintStruct) ToGo() *TPaintStruct {
	return (*TPaintStruct)(unsafe.Pointer(m))
}

type TSmallPoint struct {
	X int16
	Y int16
}

// TSmallPoint

func (s TSmallPoint) Empty() TSmallPoint {
	return TSmallPoint{}
}

func (s TSmallPoint) IsEqual(val TSmallPoint) bool {
	return s.X == val.X && s.Y == val.Y
}

// TLMMouse

func (m *TLMMouse) XPos() *int16 {
	return (*int16)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMMouse) YPos() *int16 {
	return (*int16)(unsafe.Pointer(&m.union[2]))
}

func (m *TLMMouse) Pos() *TSmallPoint {
	return (*TSmallPoint)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMMouse) Dummy() *int {
	return (*int)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMMouse) Result() *int {
	return (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m.union[0])) + unsafe.Sizeof(int(0))))
}

// TLMMove

func (m *TLMMove) XPos() *int16 {
	return (*int16)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMMove) YPos() *int16 {
	return (*int16)(unsafe.Pointer(&m.union[2]))
}

func (m *TLMMove) Pos() *TSmallPoint {
	return (*TSmallPoint)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMMove) Dummy() *int {
	return (*int)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMMove) Result() *int {
	return (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m.union[0])) + unsafe.Sizeof(int(0))))
}

// TLMContextMenu

func (m *TLMContextMenu) XPos() *int16 {
	return (*int16)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMContextMenu) YPos() *int16 {
	return (*int16)(unsafe.Pointer(&m.union[2]))
}

func (m *TLMContextMenu) Pos() *TSmallPoint {
	return (*TSmallPoint)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMContextMenu) Dummy() *int {
	return (*int)(unsafe.Pointer(&m.union[0]))
}

func (m *TLMContextMenu) Result() *int {
	return (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m.union[0])) + unsafe.Sizeof(int(0))))
}
