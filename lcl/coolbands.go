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

// ICoolBands Parent: ICollection
type ICoolBands interface {
	ICollection
	ItemsForCoolBand(Index int32) ICoolBand            // property
	SetItemsForCoolBand(Index int32, AValue ICoolBand) // property
	AddForCoolBand() ICoolBand                         // function
	FindBand(AControl IControl) ICoolBand              // function
	FindBandIndex(AControl IControl) int32             // function
}

// TCoolBands Parent: TCollection
type TCoolBands struct {
	TCollection
}

func NewCoolBands(ACoolBar ICustomCoolBar) ICoolBands {
	r1 := coolBandsImportAPI().SysCallN(2, GetObjectUintptr(ACoolBar))
	return AsCoolBands(r1)
}

func (m *TCoolBands) ItemsForCoolBand(Index int32) ICoolBand {
	r1 := coolBandsImportAPI().SysCallN(5, 0, m.Instance(), uintptr(Index))
	return AsCoolBand(r1)
}

func (m *TCoolBands) SetItemsForCoolBand(Index int32, AValue ICoolBand) {
	coolBandsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TCoolBands) AddForCoolBand() ICoolBand {
	r1 := coolBandsImportAPI().SysCallN(0, m.Instance())
	return AsCoolBand(r1)
}

func (m *TCoolBands) FindBand(AControl IControl) ICoolBand {
	r1 := coolBandsImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(AControl))
	return AsCoolBand(r1)
}

func (m *TCoolBands) FindBandIndex(AControl IControl) int32 {
	r1 := coolBandsImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func CoolBandsClass() TClass {
	ret := coolBandsImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	coolBandsImport       *imports.Imports = nil
	coolBandsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoolBands_AddForCoolBand", 0),
		/*1*/ imports.NewTable("CoolBands_Class", 0),
		/*2*/ imports.NewTable("CoolBands_Create", 0),
		/*3*/ imports.NewTable("CoolBands_FindBand", 0),
		/*4*/ imports.NewTable("CoolBands_FindBandIndex", 0),
		/*5*/ imports.NewTable("CoolBands_ItemsForCoolBand", 0),
	}
)

func coolBandsImportAPI() *imports.Imports {
	if coolBandsImport == nil {
		coolBandsImport = NewDefaultImports()
		coolBandsImport.SetImportTable(coolBandsImportTables)
		coolBandsImportTables = nil
	}
	return coolBandsImport
}
