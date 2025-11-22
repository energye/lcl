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

// ICustomPairSplitter Parent: IWinControl
type ICustomPairSplitter interface {
	IWinControl
	ChildClassAllowed(childClass types.TClass) bool // function
	CreateWnd()                                     // procedure
	UpdatePosition()                                // procedure
	CreateSides()                                   // procedure
	Loaded()                                        // procedure
	Sides(index int32) IPairSplitterSide            // property Sides Getter
	SplitterType() types.TPairSplitterType          // property SplitterType Getter
	SetSplitterType(value types.TPairSplitterType)  // property SplitterType Setter
	Position() int32                                // property Position Getter
	SetPosition(value int32)                        // property Position Setter
}

type TCustomPairSplitter struct {
	TWinControl
}

func (m *TCustomPairSplitter) ChildClassAllowed(childClass types.TClass) bool {
	if !m.IsValid() {
		return false
	}
	r := customPairSplitterAPI().SysCallN(1, m.Instance(), uintptr(childClass))
	return api.GoBool(r)
}

func (m *TCustomPairSplitter) CreateWnd() {
	if !m.IsValid() {
		return
	}
	customPairSplitterAPI().SysCallN(2, m.Instance())
}

func (m *TCustomPairSplitter) UpdatePosition() {
	if !m.IsValid() {
		return
	}
	customPairSplitterAPI().SysCallN(3, m.Instance())
}

func (m *TCustomPairSplitter) CreateSides() {
	if !m.IsValid() {
		return
	}
	customPairSplitterAPI().SysCallN(4, m.Instance())
}

func (m *TCustomPairSplitter) Loaded() {
	if !m.IsValid() {
		return
	}
	customPairSplitterAPI().SysCallN(5, m.Instance())
}

func (m *TCustomPairSplitter) Sides(index int32) IPairSplitterSide {
	if !m.IsValid() {
		return nil
	}
	r := customPairSplitterAPI().SysCallN(6, m.Instance(), uintptr(index))
	return AsPairSplitterSide(r)
}

func (m *TCustomPairSplitter) SplitterType() types.TPairSplitterType {
	if !m.IsValid() {
		return 0
	}
	r := customPairSplitterAPI().SysCallN(7, 0, m.Instance())
	return types.TPairSplitterType(r)
}

func (m *TCustomPairSplitter) SetSplitterType(value types.TPairSplitterType) {
	if !m.IsValid() {
		return
	}
	customPairSplitterAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPairSplitter) Position() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customPairSplitterAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCustomPairSplitter) SetPosition(value int32) {
	if !m.IsValid() {
		return
	}
	customPairSplitterAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

// NewCustomPairSplitter class constructor
func NewCustomPairSplitter(theOwner IComponent) ICustomPairSplitter {
	r := customPairSplitterAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomPairSplitter(r)
}

func TCustomPairSplitterClass() types.TClass {
	r := customPairSplitterAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	customPairSplitterOnce   base.Once
	customPairSplitterImport *imports.Imports = nil
)

func customPairSplitterAPI() *imports.Imports {
	customPairSplitterOnce.Do(func() {
		customPairSplitterImport = api.NewDefaultImports()
		customPairSplitterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPairSplitter_Create", 0), // constructor NewCustomPairSplitter
			/* 1 */ imports.NewTable("TCustomPairSplitter_ChildClassAllowed", 0), // function ChildClassAllowed
			/* 2 */ imports.NewTable("TCustomPairSplitter_CreateWnd", 0), // procedure CreateWnd
			/* 3 */ imports.NewTable("TCustomPairSplitter_UpdatePosition", 0), // procedure UpdatePosition
			/* 4 */ imports.NewTable("TCustomPairSplitter_CreateSides", 0), // procedure CreateSides
			/* 5 */ imports.NewTable("TCustomPairSplitter_Loaded", 0), // procedure Loaded
			/* 6 */ imports.NewTable("TCustomPairSplitter_Sides", 0), // property Sides
			/* 7 */ imports.NewTable("TCustomPairSplitter_SplitterType", 0), // property SplitterType
			/* 8 */ imports.NewTable("TCustomPairSplitter_Position", 0), // property Position
			/* 9 */ imports.NewTable("TCustomPairSplitter_TClass", 0), // function TCustomPairSplitterClass
		}
	})
	return customPairSplitterImport
}
