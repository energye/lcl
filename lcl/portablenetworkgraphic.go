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

// IPortableNetworkGraphic Parent: IFPImageBitmap
type IPortableNetworkGraphic interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
}

// TPortableNetworkGraphic Parent: TFPImageBitmap
type TPortableNetworkGraphic struct {
	TFPImageBitmap
}

func NewPortableNetworkGraphic() IPortableNetworkGraphic {
	r1 := LCL().SysCallN(4614)
	return AsPortableNetworkGraphic(r1)
}

func PortableNetworkGraphicClass() TClass {
	ret := LCL().SysCallN(4613)
	return TClass(ret)
}
