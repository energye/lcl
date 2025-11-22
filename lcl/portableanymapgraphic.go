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

// IPortableAnyMapGraphic Parent: IFPImageBitmap
type IPortableAnyMapGraphic interface {
	IFPImageBitmap
}

type TPortableAnyMapGraphic struct {
	TFPImageBitmap
}

// NewPortableAnyMapGraphic class constructor
func NewPortableAnyMapGraphic() IPortableAnyMapGraphic {
	r := portableAnyMapGraphicAPI().SysCallN(0)
	return AsPortableAnyMapGraphic(r)
}

var (
	portableAnyMapGraphicOnce   base.Once
	portableAnyMapGraphicImport *imports.Imports = nil
)

func portableAnyMapGraphicAPI() *imports.Imports {
	portableAnyMapGraphicOnce.Do(func() {
		portableAnyMapGraphicImport = api.NewDefaultImports()
		portableAnyMapGraphicImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPortableAnyMapGraphic_Create", 0), // constructor NewPortableAnyMapGraphic
		}
	})
	return portableAnyMapGraphicImport
}
