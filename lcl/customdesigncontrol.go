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
	"github.com/energye/lcl/types"
)

// ICustomDesignControl Parent: IScrollingWinControl
type ICustomDesignControl interface {
	IScrollingWinControl
	DesignTimePPI() int32         // property DesignTimePPI Getter
	SetDesignTimePPI(value int32) // property DesignTimePPI Setter
	PixelsPerInch() int32         // property PixelsPerInch Getter
	SetPixelsPerInch(value int32) // property PixelsPerInch Setter
	Scaled() bool                 // property Scaled Getter
	SetScaled(value bool)         // property Scaled Setter
}

type TCustomDesignControl struct {
	TScrollingWinControl
}

func (m *TCustomDesignControl) DesignTimePPI() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDesignControlAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDesignControl) SetDesignTimePPI(value int32) {
	if !m.IsValid() {
		return
	}
	customDesignControlAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDesignControl) PixelsPerInch() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDesignControlAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDesignControl) SetPixelsPerInch(value int32) {
	if !m.IsValid() {
		return
	}
	customDesignControlAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDesignControl) Scaled() bool {
	if !m.IsValid() {
		return false
	}
	r := customDesignControlAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDesignControl) SetScaled(value bool) {
	if !m.IsValid() {
		return
	}
	customDesignControlAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

// NewCustomDesignControl class constructor
func NewCustomDesignControl(theOwner IComponent) ICustomDesignControl {
	r := customDesignControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomDesignControl(r)
}

func TCustomDesignControlClass() types.TClass {
	r := customDesignControlAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	customDesignControlOnce   base.Once
	customDesignControlImport *imports.Imports = nil
)

func customDesignControlAPI() *imports.Imports {
	customDesignControlOnce.Do(func() {
		customDesignControlImport = api.NewDefaultImports()
		customDesignControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDesignControl_Create", 0), // constructor NewCustomDesignControl
			/* 1 */ imports.NewTable("TCustomDesignControl_DesignTimePPI", 0), // property DesignTimePPI
			/* 2 */ imports.NewTable("TCustomDesignControl_PixelsPerInch", 0), // property PixelsPerInch
			/* 3 */ imports.NewTable("TCustomDesignControl_Scaled", 0), // property Scaled
			/* 4 */ imports.NewTable("TCustomDesignControl_TClass", 0), // function TCustomDesignControlClass
		}
	})
	return customDesignControlImport
}
