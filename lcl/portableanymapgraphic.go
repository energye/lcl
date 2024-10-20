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
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IPortableAnyMapGraphic Parent: IFPImageBitmap
type IPortableAnyMapGraphic interface {
	IFPImageBitmap
}

// TPortableAnyMapGraphic Parent: TFPImageBitmap
type TPortableAnyMapGraphic struct {
	TFPImageBitmap
}

func NewPortableAnyMapGraphic() IPortableAnyMapGraphic {
	r1 := portableAnyMapGraphicImportAPI().SysCallN(1)
	return AsPortableAnyMapGraphic(r1)
}

func PortableAnyMapGraphicClass() TClass {
	ret := portableAnyMapGraphicImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	portableAnyMapGraphicImport       *imports.Imports = nil
	portableAnyMapGraphicImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PortableAnyMapGraphic_Class", 0),
		/*1*/ imports.NewTable("PortableAnyMapGraphic_Create", 0),
	}
)

func portableAnyMapGraphicImportAPI() *imports.Imports {
	if portableAnyMapGraphicImport == nil {
		portableAnyMapGraphicImport = NewDefaultImports()
		portableAnyMapGraphicImport.SetImportTable(portableAnyMapGraphicImportTables)
		portableAnyMapGraphicImportTables = nil
	}
	return portableAnyMapGraphicImport
}
