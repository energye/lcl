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

// ILazDockSplitter Parent: ICustomSplitter
type ILazDockSplitter interface {
	ICustomSplitter
}

// TLazDockSplitter Parent: TCustomSplitter
type TLazDockSplitter struct {
	TCustomSplitter
}

func NewLazDockSplitter(AOwner IComponent) ILazDockSplitter {
	r1 := LCL().SysCallN(3563, GetObjectUintptr(AOwner))
	return AsLazDockSplitter(r1)
}

func LazDockSplitterClass() TClass {
	ret := LCL().SysCallN(3562)
	return TClass(ret)
}
