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

// IRichMemoInline Parent: IObject
type IRichMemoInline interface {
	IObject
	Owner() ICustomRichMemo            // property
	Draw(Canvas ICanvas, ASize *TSize) // procedure
	SetVisible(AVisible bool)          // procedure
	Invalidate()                       // procedure
}

// TRichMemoInline Parent: TObject
type TRichMemoInline struct {
	TObject
}

func NewRichMemoInline() IRichMemoInline {
	r1 := LCL().SysCallN(4859)
	return AsRichMemoInline(r1)
}

func (m *TRichMemoInline) Owner() ICustomRichMemo {
	r1 := LCL().SysCallN(4862, m.Instance())
	return AsCustomRichMemo(r1)
}

func RichMemoInlineClass() TClass {
	ret := LCL().SysCallN(4858)
	return TClass(ret)
}

func (m *TRichMemoInline) Draw(Canvas ICanvas, ASize *TSize) {
	LCL().SysCallN(4860, m.Instance(), GetObjectUintptr(Canvas), uintptr(unsafePointer(ASize)))
}

func (m *TRichMemoInline) SetVisible(AVisible bool) {
	LCL().SysCallN(4863, m.Instance(), PascalBool(AVisible))
}

func (m *TRichMemoInline) Invalidate() {
	LCL().SysCallN(4861, m.Instance())
}
