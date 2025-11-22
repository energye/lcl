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

// ICustomPanel Parent: ICustomControl
type ICustomPanel interface {
	ICustomControl
	Alignment() types.TAlignment           // property Alignment Getter
	SetAlignment(value types.TAlignment)   // property Alignment Setter
	BevelColor() types.TColor              // property BevelColor Getter
	SetBevelColor(value types.TColor)      // property BevelColor Setter
	BevelInner() types.TPanelBevel         // property BevelInner Getter
	SetBevelInner(value types.TPanelBevel) // property BevelInner Setter
	BevelOuter() types.TPanelBevel         // property BevelOuter Getter
	SetBevelOuter(value types.TPanelBevel) // property BevelOuter Setter
	BevelWidth() types.TBevelWidth         // property BevelWidth Getter
	SetBevelWidth(value types.TBevelWidth) // property BevelWidth Setter
	FullRepaint() bool                     // property FullRepaint Getter
	SetFullRepaint(value bool)             // property FullRepaint Setter
	ParentBackground() bool                // property ParentBackground Getter
	SetParentBackground(value bool)        // property ParentBackground Setter
	ParentColor() bool                     // property ParentColor Getter
	SetParentColor(value bool)             // property ParentColor Setter
}

type TCustomPanel struct {
	TCustomControl
}

func (m *TCustomPanel) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := customPanelAPI().SysCallN(1, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TCustomPanel) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPanel) BevelColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customPanelAPI().SysCallN(2, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomPanel) SetBevelColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPanel) BevelInner() types.TPanelBevel {
	if !m.IsValid() {
		return 0
	}
	r := customPanelAPI().SysCallN(3, 0, m.Instance())
	return types.TPanelBevel(r)
}

func (m *TCustomPanel) SetBevelInner(value types.TPanelBevel) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPanel) BevelOuter() types.TPanelBevel {
	if !m.IsValid() {
		return 0
	}
	r := customPanelAPI().SysCallN(4, 0, m.Instance())
	return types.TPanelBevel(r)
}

func (m *TCustomPanel) SetBevelOuter(value types.TPanelBevel) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPanel) BevelWidth() types.TBevelWidth {
	if !m.IsValid() {
		return 0
	}
	r := customPanelAPI().SysCallN(5, 0, m.Instance())
	return types.TBevelWidth(r)
}

func (m *TCustomPanel) SetBevelWidth(value types.TBevelWidth) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomPanel) FullRepaint() bool {
	if !m.IsValid() {
		return false
	}
	r := customPanelAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomPanel) SetFullRepaint(value bool) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomPanel) ParentBackground() bool {
	if !m.IsValid() {
		return false
	}
	r := customPanelAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomPanel) SetParentBackground(value bool) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomPanel) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := customPanelAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomPanel) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	customPanelAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

// NewCustomPanel class constructor
func NewCustomPanel(theOwner IComponent) ICustomPanel {
	r := customPanelAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomPanel(r)
}

func TCustomPanelClass() types.TClass {
	r := customPanelAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	customPanelOnce   base.Once
	customPanelImport *imports.Imports = nil
)

func customPanelAPI() *imports.Imports {
	customPanelOnce.Do(func() {
		customPanelImport = api.NewDefaultImports()
		customPanelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPanel_Create", 0), // constructor NewCustomPanel
			/* 1 */ imports.NewTable("TCustomPanel_Alignment", 0), // property Alignment
			/* 2 */ imports.NewTable("TCustomPanel_BevelColor", 0), // property BevelColor
			/* 3 */ imports.NewTable("TCustomPanel_BevelInner", 0), // property BevelInner
			/* 4 */ imports.NewTable("TCustomPanel_BevelOuter", 0), // property BevelOuter
			/* 5 */ imports.NewTable("TCustomPanel_BevelWidth", 0), // property BevelWidth
			/* 6 */ imports.NewTable("TCustomPanel_FullRepaint", 0), // property FullRepaint
			/* 7 */ imports.NewTable("TCustomPanel_ParentBackground", 0), // property ParentBackground
			/* 8 */ imports.NewTable("TCustomPanel_ParentColor", 0), // property ParentColor
			/* 9 */ imports.NewTable("TCustomPanel_TClass", 0), // function TCustomPanelClass
		}
	})
	return customPanelImport
}
