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

// ISizeConstraints Parent: IPersistent
type ISizeConstraints interface {
	IPersistent
	EffectiveMinWidth() int32                                               // function
	EffectiveMinHeight() int32                                              // function
	EffectiveMaxWidth() int32                                               // function
	EffectiveMaxHeight() int32                                              // function
	MinMaxWidth(width int32) int32                                          // function
	MinMaxHeight(height int32) int32                                        // function
	UpdateInterfaceConstraints()                                            // procedure
	SetInterfaceConstraints(minW int32, minH int32, maxW int32, maxH int32) // procedure
	AutoAdjustLayout(xProportion float64, yProportion float64)              // procedure
	MaxInterfaceHeight() int32                                              // property MaxInterfaceHeight Getter
	MaxInterfaceWidth() int32                                               // property MaxInterfaceWidth Getter
	MinInterfaceHeight() int32                                              // property MinInterfaceHeight Getter
	MinInterfaceWidth() int32                                               // property MinInterfaceWidth Getter
	Control() IControl                                                      // property Control Getter
	Options() types.TSizeConstraintsOptions                                 // property Options Getter
	SetOptions(value types.TSizeConstraintsOptions)                         // property Options Setter
	MaxHeight() types.TConstraintSize                                       // property MaxHeight Getter
	SetMaxHeight(value types.TConstraintSize)                               // property MaxHeight Setter
	MaxWidth() types.TConstraintSize                                        // property MaxWidth Getter
	SetMaxWidth(value types.TConstraintSize)                                // property MaxWidth Setter
	MinHeight() types.TConstraintSize                                       // property MinHeight Getter
	SetMinHeight(value types.TConstraintSize)                               // property MinHeight Setter
	MinWidth() types.TConstraintSize                                        // property MinWidth Getter
	SetMinWidth(value types.TConstraintSize)                                // property MinWidth Setter
	SetOnChange(fn TNotifyEvent)                                            // property event
}

type TSizeConstraints struct {
	TPersistent
}

func (m *TSizeConstraints) EffectiveMinWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) EffectiveMinHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) EffectiveMaxWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) EffectiveMaxHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) MinMaxWidth(width int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(5, m.Instance(), uintptr(width))
	return int32(r)
}

func (m *TSizeConstraints) MinMaxHeight(height int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(6, m.Instance(), uintptr(height))
	return int32(r)
}

func (m *TSizeConstraints) UpdateInterfaceConstraints() {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(7, m.Instance())
}

func (m *TSizeConstraints) SetInterfaceConstraints(minW int32, minH int32, maxW int32, maxH int32) {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(8, m.Instance(), uintptr(minW), uintptr(minH), uintptr(maxW), uintptr(maxH))
}

func (m *TSizeConstraints) AutoAdjustLayout(xProportion float64, yProportion float64) {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&xProportion)), uintptr(base.UnsafePointer(&yProportion)))
}

func (m *TSizeConstraints) MaxInterfaceHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) MaxInterfaceWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) MinInterfaceHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(12, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) MinInterfaceWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(13, m.Instance())
	return int32(r)
}

func (m *TSizeConstraints) Control() IControl {
	if !m.IsValid() {
		return nil
	}
	r := sizeConstraintsAPI().SysCallN(14, m.Instance())
	return AsControl(r)
}

func (m *TSizeConstraints) Options() types.TSizeConstraintsOptions {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(15, 0, m.Instance())
	return types.TSizeConstraintsOptions(r)
}

func (m *TSizeConstraints) SetOptions(value types.TSizeConstraintsOptions) {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TSizeConstraints) MaxHeight() types.TConstraintSize {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(16, 0, m.Instance())
	return types.TConstraintSize(r)
}

func (m *TSizeConstraints) SetMaxHeight(value types.TConstraintSize) {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TSizeConstraints) MaxWidth() types.TConstraintSize {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(17, 0, m.Instance())
	return types.TConstraintSize(r)
}

func (m *TSizeConstraints) SetMaxWidth(value types.TConstraintSize) {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TSizeConstraints) MinHeight() types.TConstraintSize {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(18, 0, m.Instance())
	return types.TConstraintSize(r)
}

func (m *TSizeConstraints) SetMinHeight(value types.TConstraintSize) {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TSizeConstraints) MinWidth() types.TConstraintSize {
	if !m.IsValid() {
		return 0
	}
	r := sizeConstraintsAPI().SysCallN(19, 0, m.Instance())
	return types.TConstraintSize(r)
}

func (m *TSizeConstraints) SetMinWidth(value types.TConstraintSize) {
	if !m.IsValid() {
		return
	}
	sizeConstraintsAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TSizeConstraints) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, sizeConstraintsAPI(), api.MakeEventDataPtr(cb))
}

// NewSizeConstraints class constructor
func NewSizeConstraints(control IControl) ISizeConstraints {
	r := sizeConstraintsAPI().SysCallN(0, base.GetObjectUintptr(control))
	return AsSizeConstraints(r)
}

var (
	sizeConstraintsOnce   base.Once
	sizeConstraintsImport *imports.Imports = nil
)

func sizeConstraintsAPI() *imports.Imports {
	sizeConstraintsOnce.Do(func() {
		sizeConstraintsImport = api.NewDefaultImports()
		sizeConstraintsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSizeConstraints_Create", 0), // constructor NewSizeConstraints
			/* 1 */ imports.NewTable("TSizeConstraints_EffectiveMinWidth", 0), // function EffectiveMinWidth
			/* 2 */ imports.NewTable("TSizeConstraints_EffectiveMinHeight", 0), // function EffectiveMinHeight
			/* 3 */ imports.NewTable("TSizeConstraints_EffectiveMaxWidth", 0), // function EffectiveMaxWidth
			/* 4 */ imports.NewTable("TSizeConstraints_EffectiveMaxHeight", 0), // function EffectiveMaxHeight
			/* 5 */ imports.NewTable("TSizeConstraints_MinMaxWidth", 0), // function MinMaxWidth
			/* 6 */ imports.NewTable("TSizeConstraints_MinMaxHeight", 0), // function MinMaxHeight
			/* 7 */ imports.NewTable("TSizeConstraints_UpdateInterfaceConstraints", 0), // procedure UpdateInterfaceConstraints
			/* 8 */ imports.NewTable("TSizeConstraints_SetInterfaceConstraints", 0), // procedure SetInterfaceConstraints
			/* 9 */ imports.NewTable("TSizeConstraints_AutoAdjustLayout", 0), // procedure AutoAdjustLayout
			/* 10 */ imports.NewTable("TSizeConstraints_MaxInterfaceHeight", 0), // property MaxInterfaceHeight
			/* 11 */ imports.NewTable("TSizeConstraints_MaxInterfaceWidth", 0), // property MaxInterfaceWidth
			/* 12 */ imports.NewTable("TSizeConstraints_MinInterfaceHeight", 0), // property MinInterfaceHeight
			/* 13 */ imports.NewTable("TSizeConstraints_MinInterfaceWidth", 0), // property MinInterfaceWidth
			/* 14 */ imports.NewTable("TSizeConstraints_Control", 0), // property Control
			/* 15 */ imports.NewTable("TSizeConstraints_Options", 0), // property Options
			/* 16 */ imports.NewTable("TSizeConstraints_MaxHeight", 0), // property MaxHeight
			/* 17 */ imports.NewTable("TSizeConstraints_MaxWidth", 0), // property MaxWidth
			/* 18 */ imports.NewTable("TSizeConstraints_MinHeight", 0), // property MinHeight
			/* 19 */ imports.NewTable("TSizeConstraints_MinWidth", 0), // property MinWidth
			/* 20 */ imports.NewTable("TSizeConstraints_OnChange", 0), // event OnChange
		}
	})
	return sizeConstraintsImport
}
