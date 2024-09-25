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

// IEbEdit Parent: IGEEdit
type IEbEdit interface {
	IGEEdit
}

// TEbEdit Parent: TGEEdit
type TEbEdit struct {
	TGEEdit
}

func NewEbEdit(TheOwner IComponent) IEbEdit {
	r1 := LCL().SysCallN(2794, GetObjectUintptr(TheOwner))
	return AsEbEdit(r1)
}

func EbEditClass() TClass {
	ret := LCL().SysCallN(2793)
	return TClass(ret)
}
