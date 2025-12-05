//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build amd64 || arm64 || loong64
// +build amd64 arm64 loong64

package types

// TMessage 消息值参见 types/messages包
type TMessage struct {
	Msg        Cardinal
	_UnusedMsg Cardinal
	WParam     WParam
	LParam     LParam
	Result     LRESULT
}

// TLMPaint
type TLMPaint struct {
	Msg         Cardinal
	_UnusedMsg  Cardinal // cpu64
	DC          HDC
	PaintStruct PPaintStruct // call ToGo() *TPaintStruct
	Result      LResult
}

// TLMMouse
type TLMMouse struct {
	Msg   Cardinal
	Keys  PtrInt
	union [16]byte // 变体
}

// TLMSize
type TLMSize struct {
	Msg        Cardinal
	_UnusedMsg Cardinal // cpu64
	SizeType   PtrInt   // see LCLType.pp (e.g. Size_Restored)
	Width      Word
	Height     Word
	_Unused    LongInt // cpu64
	Result     LResult
}

// TLMMove
type TLMMove struct {
	Msg        Cardinal
	_UnusedMsg Cardinal // cpu64
	MoveType   PtrInt   // 0 = update, 1 = force RequestAlign, 128 = Source is Interface (Widget has moved)
	union      [16]byte // 变体
}

type TWindowPosChanged struct {
	Msg        Cardinal
	_UnusedMsg Cardinal
	_Unused    WParam
	WindowPos  WindowPos
	Result     LParam
}

// TLMKey
type TLMKey struct {
	Msg        Cardinal
	_UnusedMsg Cardinal // cpu64
	CharCode   Word     // LITTLE
	_Unused    Word     // LITTLE
	_Unused2   LongInt  // cpu64
	KeyData    PtrInt
	Result     LResult
}

// TLMContextMenu
type TLMContextMenu struct {
	Msg        Cardinal
	_UnusedMsg Cardinal // cpu64
	Hwnd       HWND
	union      [16]byte // 变体
}

// TLMSetFocus
type TLMSetFocus struct {
	Msg        Cardinal
	_UnusedMsg Cardinal // cpu64
	FocusedWnd HWND
	_Unused    LParam
	Result     LRESULT
}
