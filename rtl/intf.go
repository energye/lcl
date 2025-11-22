//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import . "github.com/energye/lcl/types"

func MakeWord(A, B uint8) uint16 {
	return uint16(A | B<<8)
}

func MakeLong(A, B uint16) int32 {
	return int32(A | B<<16)
}

func HiWord(L uint32) uint16 {
	return uint16(L >> 16)
}

func HiByte(W uint8) uint8 {
	return W >> 8
}

func MakeWParam(l, h uint16) WParam {
	return WParam(MakeLong(l, h))
}

func MakeLParam(l, h uint16) LParam {
	return LParam(MakeLong(l, h))
}

func MakeLResult(l, h uint16) LRESULT {
	return LRESULT(MakeLong(l, h))
}

func PointToLParam(P TPoint) LParam {
	return LParam((P.X & 0x0000FFFF) | (P.Y << 16))
}
