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

// ICalculatorDialog Parent: IExtCommonDialog
type ICalculatorDialog interface {
	IExtCommonDialog
	CalcDisplay() float64                              // property CalcDisplay Getter
	Memory() float64                                   // property Memory Getter
	BeepOnError() bool                                 // property BeepOnError Getter
	SetBeepOnError(value bool)                         // property BeepOnError Setter
	CalculatorLayout() types.TCalculatorLayout         // property CalculatorLayout Getter
	SetCalculatorLayout(value types.TCalculatorLayout) // property CalculatorLayout Setter
	Precision() byte                                   // property Precision Getter
	SetPrecision(value byte)                           // property Precision Setter
	Value() float64                                    // property Value Getter
	SetValue(value float64)                            // property Value Setter
	DialogScale() int32                                // property DialogScale Getter
	SetDialogScale(value int32)                        // property DialogScale Setter
	ColorBtnDigits() types.TColor                      // property ColorBtnDigits Getter
	SetColorBtnDigits(value types.TColor)              // property ColorBtnDigits Setter
	ColorBtnMemory() types.TColor                      // property ColorBtnMemory Getter
	SetColorBtnMemory(value types.TColor)              // property ColorBtnMemory Setter
	ColorBtnOk() types.TColor                          // property ColorBtnOk Getter
	SetColorBtnOk(value types.TColor)                  // property ColorBtnOk Setter
	ColorBtnCancel() types.TColor                      // property ColorBtnCancel Getter
	SetColorBtnCancel(value types.TColor)              // property ColorBtnCancel Setter
	ColorBtnClear() types.TColor                       // property ColorBtnClear Getter
	SetColorBtnClear(value types.TColor)               // property ColorBtnClear Setter
	ColorBtnOthers() types.TColor                      // property ColorBtnOthers Getter
	SetColorBtnOthers(value types.TColor)              // property ColorBtnOthers Setter
	ColorDisplayText() types.TColor                    // property ColorDisplayText Getter
	SetColorDisplayText(value types.TColor)            // property ColorDisplayText Setter
	ColorDisplayBack() types.TColor                    // property ColorDisplayBack Getter
	SetColorDisplayBack(value types.TColor)            // property ColorDisplayBack Setter
	SetOnCalcKey(fn TKeyPressEvent)                    // property event
	SetOnChange(fn TNotifyEvent)                       // property event
	SetOnDisplayChange(fn TNotifyEvent)                // property event
}

type TCalculatorDialog struct {
	TExtCommonDialog
}

func (m *TCalculatorDialog) CalcDisplay() (result float64) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCalculatorDialog) Memory() (result float64) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCalculatorDialog) BeepOnError() bool {
	if !m.IsValid() {
		return false
	}
	r := calculatorDialogAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCalculatorDialog) SetBeepOnError(value bool) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCalculatorDialog) CalculatorLayout() types.TCalculatorLayout {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(4, 0, m.Instance())
	return types.TCalculatorLayout(r)
}

func (m *TCalculatorDialog) SetCalculatorLayout(value types.TCalculatorLayout) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) Precision() byte {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(5, 0, m.Instance())
	return byte(r)
}

func (m *TCalculatorDialog) SetPrecision(value byte) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) Value() (result float64) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCalculatorDialog) SetValue(value float64) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(6, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCalculatorDialog) DialogScale() int32 {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCalculatorDialog) SetDialogScale(value int32) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorBtnDigits() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorBtnDigits(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorBtnMemory() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(9, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorBtnMemory(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorBtnOk() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(10, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorBtnOk(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorBtnCancel() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(11, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorBtnCancel(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorBtnClear() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(12, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorBtnClear(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorBtnOthers() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(13, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorBtnOthers(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorDisplayText() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(14, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorDisplayText(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) ColorDisplayBack() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := calculatorDialogAPI().SysCallN(15, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCalculatorDialog) SetColorDisplayBack(value types.TColor) {
	if !m.IsValid() {
		return
	}
	calculatorDialogAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCalculatorDialog) SetOnCalcKey(fn TKeyPressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTKeyPressEvent(fn)
	base.SetEvent(m, 16, calculatorDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalculatorDialog) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, calculatorDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalculatorDialog) SetOnDisplayChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, calculatorDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewCalculatorDialog class constructor
func NewCalculatorDialog(owner IComponent) ICalculatorDialog {
	r := calculatorDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCalculatorDialog(r)
}

func TCalculatorDialogClass() types.TClass {
	r := calculatorDialogAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	calculatorDialogOnce   base.Once
	calculatorDialogImport *imports.Imports = nil
)

func calculatorDialogAPI() *imports.Imports {
	calculatorDialogOnce.Do(func() {
		calculatorDialogImport = api.NewDefaultImports()
		calculatorDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCalculatorDialog_Create", 0), // constructor NewCalculatorDialog
			/* 1 */ imports.NewTable("TCalculatorDialog_CalcDisplay", 0), // property CalcDisplay
			/* 2 */ imports.NewTable("TCalculatorDialog_Memory", 0), // property Memory
			/* 3 */ imports.NewTable("TCalculatorDialog_BeepOnError", 0), // property BeepOnError
			/* 4 */ imports.NewTable("TCalculatorDialog_CalculatorLayout", 0), // property CalculatorLayout
			/* 5 */ imports.NewTable("TCalculatorDialog_Precision", 0), // property Precision
			/* 6 */ imports.NewTable("TCalculatorDialog_Value", 0), // property Value
			/* 7 */ imports.NewTable("TCalculatorDialog_DialogScale", 0), // property DialogScale
			/* 8 */ imports.NewTable("TCalculatorDialog_ColorBtnDigits", 0), // property ColorBtnDigits
			/* 9 */ imports.NewTable("TCalculatorDialog_ColorBtnMemory", 0), // property ColorBtnMemory
			/* 10 */ imports.NewTable("TCalculatorDialog_ColorBtnOk", 0), // property ColorBtnOk
			/* 11 */ imports.NewTable("TCalculatorDialog_ColorBtnCancel", 0), // property ColorBtnCancel
			/* 12 */ imports.NewTable("TCalculatorDialog_ColorBtnClear", 0), // property ColorBtnClear
			/* 13 */ imports.NewTable("TCalculatorDialog_ColorBtnOthers", 0), // property ColorBtnOthers
			/* 14 */ imports.NewTable("TCalculatorDialog_ColorDisplayText", 0), // property ColorDisplayText
			/* 15 */ imports.NewTable("TCalculatorDialog_ColorDisplayBack", 0), // property ColorDisplayBack
			/* 16 */ imports.NewTable("TCalculatorDialog_OnCalcKey", 0), // event OnCalcKey
			/* 17 */ imports.NewTable("TCalculatorDialog_OnChange", 0), // event OnChange
			/* 18 */ imports.NewTable("TCalculatorDialog_OnDisplayChange", 0), // event OnDisplayChange
			/* 19 */ imports.NewTable("TCalculatorDialog_TClass", 0), // function TCalculatorDialogClass
		}
	})
	return calculatorDialogImport
}
