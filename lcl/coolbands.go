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

// ICoolBands Parent: ICollection
type ICoolBands interface {
	ICollection
	AddToCoolBand() ICoolBand                               // function
	FindBand(control IControl) ICoolBand                    // function
	FindBandIndex(control IControl) int32                   // function
	ItemsWithIntToCoolBand(index int32) ICoolBand           // property Items Getter
	SetItemsWithIntToCoolBand(index int32, value ICoolBand) // property Items Setter
}

type TCoolBands struct {
	TCollection
}

func (m *TCoolBands) AddToCoolBand() ICoolBand {
	if !m.IsValid() {
		return nil
	}
	r := coolBandsAPI().SysCallN(1, m.Instance())
	return AsCoolBand(r)
}

func (m *TCoolBands) FindBand(control IControl) ICoolBand {
	if !m.IsValid() {
		return nil
	}
	r := coolBandsAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(control))
	return AsCoolBand(r)
}

func (m *TCoolBands) FindBandIndex(control IControl) int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandsAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(control))
	return int32(r)
}

func (m *TCoolBands) ItemsWithIntToCoolBand(index int32) ICoolBand {
	if !m.IsValid() {
		return nil
	}
	r := coolBandsAPI().SysCallN(4, 0, m.Instance(), uintptr(index))
	return AsCoolBand(r)
}

func (m *TCoolBands) SetItemsWithIntToCoolBand(index int32, value ICoolBand) {
	if !m.IsValid() {
		return
	}
	coolBandsAPI().SysCallN(4, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewCoolBands class constructor
func NewCoolBands(coolBar ICustomCoolBar) ICoolBands {
	r := coolBandsAPI().SysCallN(0, base.GetObjectUintptr(coolBar))
	return AsCoolBands(r)
}

var (
	coolBandsOnce   base.Once
	coolBandsImport *imports.Imports = nil
)

func coolBandsAPI() *imports.Imports {
	coolBandsOnce.Do(func() {
		coolBandsImport = api.NewDefaultImports()
		coolBandsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoolBands_Create", 0), // constructor NewCoolBands
			/* 1 */ imports.NewTable("TCoolBands_AddToCoolBand", 0), // function AddToCoolBand
			/* 2 */ imports.NewTable("TCoolBands_FindBand", 0), // function FindBand
			/* 3 */ imports.NewTable("TCoolBands_FindBandIndex", 0), // function FindBandIndex
			/* 4 */ imports.NewTable("TCoolBands_ItemsWithIntToCoolBand", 0), // property ItemsWithIntToCoolBand
		}
	})
	return coolBandsImport
}
