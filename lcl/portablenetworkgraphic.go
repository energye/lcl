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
	r1 := portableNetworkGraphicImportAPI().SysCallN(1)
	return AsPortableNetworkGraphic(r1)
}

func PortableNetworkGraphicClass() TClass {
	ret := portableNetworkGraphicImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	portableNetworkGraphicImport       *imports.Imports = nil
	portableNetworkGraphicImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PortableNetworkGraphic_Class", 0),
		/*1*/ imports.NewTable("PortableNetworkGraphic_Create", 0),
	}
)

func portableNetworkGraphicImportAPI() *imports.Imports {
	if portableNetworkGraphicImport == nil {
		portableNetworkGraphicImport = NewDefaultImports()
		portableNetworkGraphicImport.SetImportTable(portableNetworkGraphicImportTables)
		portableNetworkGraphicImportTables = nil
	}
	return portableNetworkGraphicImport
}
