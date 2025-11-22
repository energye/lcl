//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// IPortableNetworkGraphic Parent: IFPImageBitmap
type IPortableNetworkGraphic interface {
	IFPImageBitmap
}

type TPortableNetworkGraphic struct {
	TFPImageBitmap
}

// NewPortableNetworkGraphic class constructor
func NewPortableNetworkGraphic() IPortableNetworkGraphic {
	r := portableNetworkGraphicAPI().SysCallN(0)
	return AsPortableNetworkGraphic(r)
}

var (
	portableNetworkGraphicOnce   base.Once
	portableNetworkGraphicImport *imports.Imports = nil
)

func portableNetworkGraphicAPI() *imports.Imports {
	portableNetworkGraphicOnce.Do(func() {
		portableNetworkGraphicImport = api.NewDefaultImports()
		portableNetworkGraphicImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPortableNetworkGraphic_Create", 0), // constructor NewPortableNetworkGraphic
		}
	})
	return portableNetworkGraphicImport
}
