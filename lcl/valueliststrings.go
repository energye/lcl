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

// IValueListStrings Parent: IStringList
type IValueListStrings interface {
	IStringList
}

// TValueListStrings Parent: TStringList
type TValueListStrings struct {
	TStringList
}

func NewValueListStrings(AOwner IValueListEditor) IValueListStrings {
	r1 := LCL().SysCallN(5936, GetObjectUintptr(AOwner))
	return AsValueListStrings(r1)
}

func ValueListStringsClass() TClass {
	ret := LCL().SysCallN(5935)
	return TClass(ret)
}
