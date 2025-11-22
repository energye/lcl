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

// ICustomProgressBar Parent: IWinControl
type ICustomProgressBar interface {
	IWinControl
	StepIt()                                            // procedure
	StepBy(delta int32)                                 // procedure
	Max() int32                                         // property Max Getter
	SetMax(value int32)                                 // property Max Setter
	Min() int32                                         // property Min Getter
	SetMin(value int32)                                 // property Min Setter
	Orientation() types.TProgressBarOrientation         // property Orientation Getter
	SetOrientation(value types.TProgressBarOrientation) // property Orientation Setter
	Position() int32                                    // property Position Getter
	SetPosition(value int32)                            // property Position Setter
	Smooth() bool                                       // property Smooth Getter
	SetSmooth(value bool)                               // property Smooth Setter
	Step() int32                                        // property Step Getter
	SetStep(value int32)                                // property Step Setter
	Style() types.TProgressBarStyle                     // property Style Getter
	SetStyle(value types.TProgressBarStyle)             // property Style Setter
	BarShowText() bool                                  // property BarShowText Getter
	SetBarShowText(value bool)                          // property BarShowText Setter
}

type TCustomProgressBar struct {
	TWinControl
}

func (m *TCustomProgressBar) StepIt() {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(1, m.Instance())
}

func (m *TCustomProgressBar) StepBy(delta int32) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(2, m.Instance(), uintptr(delta))
}

func (m *TCustomProgressBar) Max() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customProgressBarAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCustomProgressBar) SetMax(value int32) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomProgressBar) Min() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customProgressBarAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TCustomProgressBar) SetMin(value int32) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomProgressBar) Orientation() types.TProgressBarOrientation {
	if !m.IsValid() {
		return 0
	}
	r := customProgressBarAPI().SysCallN(5, 0, m.Instance())
	return types.TProgressBarOrientation(r)
}

func (m *TCustomProgressBar) SetOrientation(value types.TProgressBarOrientation) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomProgressBar) Position() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customProgressBarAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TCustomProgressBar) SetPosition(value int32) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomProgressBar) Smooth() bool {
	if !m.IsValid() {
		return false
	}
	r := customProgressBarAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomProgressBar) SetSmooth(value bool) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomProgressBar) Step() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customProgressBarAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCustomProgressBar) SetStep(value int32) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomProgressBar) Style() types.TProgressBarStyle {
	if !m.IsValid() {
		return 0
	}
	r := customProgressBarAPI().SysCallN(9, 0, m.Instance())
	return types.TProgressBarStyle(r)
}

func (m *TCustomProgressBar) SetStyle(value types.TProgressBarStyle) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomProgressBar) BarShowText() bool {
	if !m.IsValid() {
		return false
	}
	r := customProgressBarAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomProgressBar) SetBarShowText(value bool) {
	if !m.IsValid() {
		return
	}
	customProgressBarAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

// NewCustomProgressBar class constructor
func NewCustomProgressBar(owner IComponent) ICustomProgressBar {
	r := customProgressBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomProgressBar(r)
}

func TCustomProgressBarClass() types.TClass {
	r := customProgressBarAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	customProgressBarOnce   base.Once
	customProgressBarImport *imports.Imports = nil
)

func customProgressBarAPI() *imports.Imports {
	customProgressBarOnce.Do(func() {
		customProgressBarImport = api.NewDefaultImports()
		customProgressBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomProgressBar_Create", 0), // constructor NewCustomProgressBar
			/* 1 */ imports.NewTable("TCustomProgressBar_StepIt", 0), // procedure StepIt
			/* 2 */ imports.NewTable("TCustomProgressBar_StepBy", 0), // procedure StepBy
			/* 3 */ imports.NewTable("TCustomProgressBar_Max", 0), // property Max
			/* 4 */ imports.NewTable("TCustomProgressBar_Min", 0), // property Min
			/* 5 */ imports.NewTable("TCustomProgressBar_Orientation", 0), // property Orientation
			/* 6 */ imports.NewTable("TCustomProgressBar_Position", 0), // property Position
			/* 7 */ imports.NewTable("TCustomProgressBar_Smooth", 0), // property Smooth
			/* 8 */ imports.NewTable("TCustomProgressBar_Step", 0), // property Step
			/* 9 */ imports.NewTable("TCustomProgressBar_Style", 0), // property Style
			/* 10 */ imports.NewTable("TCustomProgressBar_BarShowText", 0), // property BarShowText
			/* 11 */ imports.NewTable("TCustomProgressBar_TClass", 0), // function TCustomProgressBarClass
		}
	})
	return customProgressBarImport
}
