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

// IStringsEnumerator Parent: IObject
type IStringsEnumerator interface {
	IObject
	Current() string    // property
	GetCurrent() string // function
	MoveNext() bool     // function
}

// TStringsEnumerator Parent: TObject
type TStringsEnumerator struct {
	TObject
}

func NewStringsEnumerator(AStrings IStrings) IStringsEnumerator {
	r1 := LCL().SysCallN(5289, GetObjectUintptr(AStrings))
	return AsStringsEnumerator(r1)
}

func (m *TStringsEnumerator) Current() string {
	r1 := LCL().SysCallN(5290, m.Instance())
	return GoStr(r1)
}

func (m *TStringsEnumerator) GetCurrent() string {
	r1 := LCL().SysCallN(5291, m.Instance())
	return GoStr(r1)
}

func (m *TStringsEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(5292, m.Instance())
	return GoBool(r1)
}

func StringsEnumeratorClass() TClass {
	ret := LCL().SysCallN(5288)
	return TClass(ret)
}
