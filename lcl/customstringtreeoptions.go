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

// ICustomStringTreeOptions Parent: ICustomVirtualTreeOptions
type ICustomStringTreeOptions interface {
	ICustomVirtualTreeOptions
}

// TCustomStringTreeOptions Parent: TCustomVirtualTreeOptions
type TCustomStringTreeOptions struct {
	TCustomVirtualTreeOptions
}

func NewCustomStringTreeOptions(AOwner IBaseVirtualTree) ICustomStringTreeOptions {
	r1 := LCL().SysCallN(2321, GetObjectUintptr(AOwner))
	return AsCustomStringTreeOptions(r1)
}

func CustomStringTreeOptionsClass() TClass {
	ret := LCL().SysCallN(2320)
	return TClass(ret)
}
