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

// IShortCutList Parent: IStringList
type IShortCutList interface {
	IStringList
	ShortCuts(Index int32) TShortCut          // property
	IndexOfShortCut(Shortcut TShortCut) int32 // function
}

// TShortCutList Parent: TStringList
type TShortCutList struct {
	TStringList
}

func NewShortCutList() IShortCutList {
	r1 := LCL().SysCallN(5044)
	return AsShortCutList(r1)
}

func (m *TShortCutList) ShortCuts(Index int32) TShortCut {
	r1 := LCL().SysCallN(5046, m.Instance(), uintptr(Index))
	return TShortCut(r1)
}

func (m *TShortCutList) IndexOfShortCut(Shortcut TShortCut) int32 {
	r1 := LCL().SysCallN(5045, m.Instance(), uintptr(Shortcut))
	return int32(r1)
}

func ShortCutListClass() TClass {
	ret := LCL().SysCallN(5043)
	return TClass(ret)
}
