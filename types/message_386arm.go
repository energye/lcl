//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build 386 || arm
// +build 386 arm

package types

// TMessage 消息值参见 types/messages包
type TMessage struct {
	Msg    Cardinal
	WParam WParam
	LParam LParam
	Result LRESULT
}

// TLMPaint
type TLMPaint struct {
	Msg         Cardinal
	DC          HDC
	PaintStruct PPaintStruct
	Result      LResult
}

// TLMMouse
type TLMMouse struct {
	Msg   Cardinal
	Keys  PtrInt
	union [8]byte // 变体
}

// TLMSize
type TLMSize struct {
	Msg      Cardinal
	SizeType PtrInt // see LCLType.pp (e.g. Size_Restored)
	Width    Word
	Height   Word
	Result   LResult
}

// TLMMove
type TLMMove struct {
	Msg      Cardinal
	MoveType PtrInt
	union    [8]byte
}

type TWindowPosChanged struct {
	Msg       Cardinal
	_Unused   WParam
	WindowPos WindowPos
	Result    LParam
}

// TLMKey
type TLMKey struct {
	Msg      Cardinal
	CharCode Word // LITTLE
	_Unused  Word // LITTLE
	KeyData  PtrInt
	Result   LResult
}

// TLMContextMenu
type TLMContextMenu struct {
	Msg   Cardinal
	Hwnd  HWND
	union [8]byte // 变体
}
