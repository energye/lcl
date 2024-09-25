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
	. "github.com/energye/lcl/types"
)

// ISizeConstraints Parent: IPersistent
type ISizeConstraints interface {
	IPersistent
	MaxInterfaceHeight() int32                            // property
	MaxInterfaceWidth() int32                             // property
	MinInterfaceHeight() int32                            // property
	MinInterfaceWidth() int32                             // property
	Control() IControl                                    // property
	Options() TSizeConstraintsOptions                     // property
	SetOptions(AValue TSizeConstraintsOptions)            // property
	MaxHeight() TConstraintSize                           // property
	SetMaxHeight(AValue TConstraintSize)                  // property
	MaxWidth() TConstraintSize                            // property
	SetMaxWidth(AValue TConstraintSize)                   // property
	MinHeight() TConstraintSize                           // property
	SetMinHeight(AValue TConstraintSize)                  // property
	MinWidth() TConstraintSize                            // property
	SetMinWidth(AValue TConstraintSize)                   // property
	EffectiveMinWidth() int32                             // function
	EffectiveMinHeight() int32                            // function
	EffectiveMaxWidth() int32                             // function
	EffectiveMaxHeight() int32                            // function
	MinMaxWidth(Width int32) int32                        // function
	MinMaxHeight(Height int32) int32                      // function
	UpdateInterfaceConstraints()                          // procedure
	SetInterfaceConstraints(MinW, MinH, MaxW, MaxH int32) // procedure
	AutoAdjustLayout(AXProportion, AYProportion float64)  // procedure
	SetOnChange(fn TNotifyEvent)                          // property event
}

// TSizeConstraints Parent: TPersistent
type TSizeConstraints struct {
	TPersistent
	changePtr uintptr
}

func NewSizeConstraints(AControl IControl) ISizeConstraints {
	r1 := LCL().SysCallN(5050, GetObjectUintptr(AControl))
	return AsSizeConstraints(r1)
}

func (m *TSizeConstraints) MaxInterfaceHeight() int32 {
	r1 := LCL().SysCallN(5056, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MaxInterfaceWidth() int32 {
	r1 := LCL().SysCallN(5057, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinInterfaceHeight() int32 {
	r1 := LCL().SysCallN(5060, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinInterfaceWidth() int32 {
	r1 := LCL().SysCallN(5061, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) Control() IControl {
	r1 := LCL().SysCallN(5049, m.Instance())
	return AsControl(r1)
}

func (m *TSizeConstraints) Options() TSizeConstraintsOptions {
	r1 := LCL().SysCallN(5065, 0, m.Instance(), 0)
	return TSizeConstraintsOptions(r1)
}

func (m *TSizeConstraints) SetOptions(AValue TSizeConstraintsOptions) {
	LCL().SysCallN(5065, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MaxHeight() TConstraintSize {
	r1 := LCL().SysCallN(5055, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMaxHeight(AValue TConstraintSize) {
	LCL().SysCallN(5055, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MaxWidth() TConstraintSize {
	r1 := LCL().SysCallN(5058, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMaxWidth(AValue TConstraintSize) {
	LCL().SysCallN(5058, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MinHeight() TConstraintSize {
	r1 := LCL().SysCallN(5059, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMinHeight(AValue TConstraintSize) {
	LCL().SysCallN(5059, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MinWidth() TConstraintSize {
	r1 := LCL().SysCallN(5064, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMinWidth(AValue TConstraintSize) {
	LCL().SysCallN(5064, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) EffectiveMinWidth() int32 {
	r1 := LCL().SysCallN(5054, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMinHeight() int32 {
	r1 := LCL().SysCallN(5053, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMaxWidth() int32 {
	r1 := LCL().SysCallN(5052, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMaxHeight() int32 {
	r1 := LCL().SysCallN(5051, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinMaxWidth(Width int32) int32 {
	r1 := LCL().SysCallN(5063, m.Instance(), uintptr(Width))
	return int32(r1)
}

func (m *TSizeConstraints) MinMaxHeight(Height int32) int32 {
	r1 := LCL().SysCallN(5062, m.Instance(), uintptr(Height))
	return int32(r1)
}

func SizeConstraintsClass() TClass {
	ret := LCL().SysCallN(5048)
	return TClass(ret)
}

func (m *TSizeConstraints) UpdateInterfaceConstraints() {
	LCL().SysCallN(5068, m.Instance())
}

func (m *TSizeConstraints) SetInterfaceConstraints(MinW, MinH, MaxW, MaxH int32) {
	LCL().SysCallN(5066, m.Instance(), uintptr(MinW), uintptr(MinH), uintptr(MaxW), uintptr(MaxH))
}

func (m *TSizeConstraints) AutoAdjustLayout(AXProportion, AYProportion float64) {
	LCL().SysCallN(5047, m.Instance(), uintptr(unsafePointer(&AXProportion)), uintptr(unsafePointer(&AYProportion)))
}

func (m *TSizeConstraints) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5067, m.Instance(), m.changePtr)
}
