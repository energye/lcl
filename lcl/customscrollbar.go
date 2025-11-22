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

// ICustomScrollBar Parent: IWinControl
type ICustomScrollBar interface {
	IWinControl
	SetParamsWithIntX4(position int32, min int32, max int32, pageSize int32) // procedure
	SetParamsWithIntX3(position int32, min int32, max int32)                 // procedure
	Kind() types.TScrollBarKind                                              // property Kind Getter
	SetKind(value types.TScrollBarKind)                                      // property Kind Setter
	LargeChange() types.TScrollBarInc                                        // property LargeChange Getter
	SetLargeChange(value types.TScrollBarInc)                                // property LargeChange Setter
	Max() int32                                                              // property Max Getter
	SetMax(value int32)                                                      // property Max Setter
	Min() int32                                                              // property Min Getter
	SetMin(value int32)                                                      // property Min Setter
	PageSize() int32                                                         // property PageSize Getter
	SetPageSize(value int32)                                                 // property PageSize Setter
	Position() int32                                                         // property Position Getter
	SetPosition(value int32)                                                 // property Position Setter
	SmallChange() types.TScrollBarInc                                        // property SmallChange Getter
	SetSmallChange(value types.TScrollBarInc)                                // property SmallChange Setter
	SetOnChange(fn TNotifyEvent)                                             // property event
	SetOnScroll(fn TScrollEvent)                                             // property event
}

type TCustomScrollBar struct {
	TWinControl
}

func (m *TCustomScrollBar) SetParamsWithIntX4(position int32, min int32, max int32, pageSize int32) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(1, m.Instance(), uintptr(position), uintptr(min), uintptr(max), uintptr(pageSize))
}

func (m *TCustomScrollBar) SetParamsWithIntX3(position int32, min int32, max int32) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(2, m.Instance(), uintptr(position), uintptr(min), uintptr(max))
}

func (m *TCustomScrollBar) Kind() types.TScrollBarKind {
	if !m.IsValid() {
		return 0
	}
	r := customScrollBarAPI().SysCallN(3, 0, m.Instance())
	return types.TScrollBarKind(r)
}

func (m *TCustomScrollBar) SetKind(value types.TScrollBarKind) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomScrollBar) LargeChange() types.TScrollBarInc {
	if !m.IsValid() {
		return 0
	}
	r := customScrollBarAPI().SysCallN(4, 0, m.Instance())
	return types.TScrollBarInc(r)
}

func (m *TCustomScrollBar) SetLargeChange(value types.TScrollBarInc) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomScrollBar) Max() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customScrollBarAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TCustomScrollBar) SetMax(value int32) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomScrollBar) Min() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customScrollBarAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TCustomScrollBar) SetMin(value int32) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomScrollBar) PageSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customScrollBarAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCustomScrollBar) SetPageSize(value int32) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomScrollBar) Position() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customScrollBarAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCustomScrollBar) SetPosition(value int32) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomScrollBar) SmallChange() types.TScrollBarInc {
	if !m.IsValid() {
		return 0
	}
	r := customScrollBarAPI().SysCallN(9, 0, m.Instance())
	return types.TScrollBarInc(r)
}

func (m *TCustomScrollBar) SetSmallChange(value types.TScrollBarInc) {
	if !m.IsValid() {
		return
	}
	customScrollBarAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomScrollBar) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, customScrollBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomScrollBar) SetOnScroll(fn TScrollEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTScrollEvent(fn)
	base.SetEvent(m, 11, customScrollBarAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomScrollBar class constructor
func NewCustomScrollBar(owner IComponent) ICustomScrollBar {
	r := customScrollBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomScrollBar(r)
}

func TCustomScrollBarClass() types.TClass {
	r := customScrollBarAPI().SysCallN(12)
	return types.TClass(r)
}

var (
	customScrollBarOnce   base.Once
	customScrollBarImport *imports.Imports = nil
)

func customScrollBarAPI() *imports.Imports {
	customScrollBarOnce.Do(func() {
		customScrollBarImport = api.NewDefaultImports()
		customScrollBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomScrollBar_Create", 0), // constructor NewCustomScrollBar
			/* 1 */ imports.NewTable("TCustomScrollBar_SetParamsWithIntX4", 0), // procedure SetParamsWithIntX4
			/* 2 */ imports.NewTable("TCustomScrollBar_SetParamsWithIntX3", 0), // procedure SetParamsWithIntX3
			/* 3 */ imports.NewTable("TCustomScrollBar_Kind", 0), // property Kind
			/* 4 */ imports.NewTable("TCustomScrollBar_LargeChange", 0), // property LargeChange
			/* 5 */ imports.NewTable("TCustomScrollBar_Max", 0), // property Max
			/* 6 */ imports.NewTable("TCustomScrollBar_Min", 0), // property Min
			/* 7 */ imports.NewTable("TCustomScrollBar_PageSize", 0), // property PageSize
			/* 8 */ imports.NewTable("TCustomScrollBar_Position", 0), // property Position
			/* 9 */ imports.NewTable("TCustomScrollBar_SmallChange", 0), // property SmallChange
			/* 10 */ imports.NewTable("TCustomScrollBar_OnChange", 0), // event OnChange
			/* 11 */ imports.NewTable("TCustomScrollBar_OnScroll", 0), // event OnScroll
			/* 12 */ imports.NewTable("TCustomScrollBar_TClass", 0), // function TCustomScrollBarClass
		}
	})
	return customScrollBarImport
}
