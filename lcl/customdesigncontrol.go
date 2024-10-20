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

// ICustomDesignControl Parent: IScrollingWinControl
type ICustomDesignControl interface {
	IScrollingWinControl
	DesignTimePPI() int32          // property
	SetDesignTimePPI(AValue int32) // property
	PixelsPerInch() int32          // property
	SetPixelsPerInch(AValue int32) // property
	Scaled() bool                  // property
	SetScaled(AValue bool)         // property
}

// TCustomDesignControl Parent: TScrollingWinControl
type TCustomDesignControl struct {
	TScrollingWinControl
}

func NewCustomDesignControl(TheOwner IComponent) ICustomDesignControl {
	r1 := customDesignControlImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsCustomDesignControl(r1)
}

func (m *TCustomDesignControl) DesignTimePPI() int32 {
	r1 := customDesignControlImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDesignControl) SetDesignTimePPI(AValue int32) {
	customDesignControlImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDesignControl) PixelsPerInch() int32 {
	r1 := customDesignControlImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDesignControl) SetPixelsPerInch(AValue int32) {
	customDesignControlImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDesignControl) Scaled() bool {
	r1 := customDesignControlImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDesignControl) SetScaled(AValue bool) {
	customDesignControlImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func CustomDesignControlClass() TClass {
	ret := customDesignControlImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customDesignControlImport       *imports.Imports = nil
	customDesignControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDesignControl_Class", 0),
		/*1*/ imports.NewTable("CustomDesignControl_Create", 0),
		/*2*/ imports.NewTable("CustomDesignControl_DesignTimePPI", 0),
		/*3*/ imports.NewTable("CustomDesignControl_PixelsPerInch", 0),
		/*4*/ imports.NewTable("CustomDesignControl_Scaled", 0),
	}
)

func customDesignControlImportAPI() *imports.Imports {
	if customDesignControlImport == nil {
		customDesignControlImport = NewDefaultImports()
		customDesignControlImport.SetImportTable(customDesignControlImportTables)
		customDesignControlImportTables = nil
	}
	return customDesignControlImport
}
