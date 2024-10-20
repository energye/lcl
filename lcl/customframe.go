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

// ICustomFrame Parent: ICustomDesignControl
type ICustomFrame interface {
	ICustomDesignControl
	ParentBackground() bool          // property
	SetParentBackground(AValue bool) // property
}

// TCustomFrame Parent: TCustomDesignControl
type TCustomFrame struct {
	TCustomDesignControl
}

func NewCustomFrame(AOwner IComponent) ICustomFrame {
	r1 := customFrameImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomFrame(r1)
}

func (m *TCustomFrame) ParentBackground() bool {
	r1 := customFrameImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFrame) SetParentBackground(AValue bool) {
	customFrameImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func CustomFrameClass() TClass {
	ret := customFrameImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customFrameImport       *imports.Imports = nil
	customFrameImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomFrame_Class", 0),
		/*1*/ imports.NewTable("CustomFrame_Create", 0),
		/*2*/ imports.NewTable("CustomFrame_ParentBackground", 0),
	}
)

func customFrameImportAPI() *imports.Imports {
	if customFrameImport == nil {
		customFrameImport = NewDefaultImports()
		customFrameImport.SetImportTable(customFrameImportTables)
		customFrameImportTables = nil
	}
	return customFrameImport
}
