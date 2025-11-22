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

// IVTFixedAreaConstraints Parent: IPersistent
type IVTFixedAreaConstraints interface {
	IPersistent
	MaxHeightPercent() types.TVTConstraintPercent         // property MaxHeightPercent Getter
	SetMaxHeightPercent(value types.TVTConstraintPercent) // property MaxHeightPercent Setter
	MaxWidthPercent() types.TVTConstraintPercent          // property MaxWidthPercent Getter
	SetMaxWidthPercent(value types.TVTConstraintPercent)  // property MaxWidthPercent Setter
	MinHeightPercent() types.TVTConstraintPercent         // property MinHeightPercent Getter
	SetMinHeightPercent(value types.TVTConstraintPercent) // property MinHeightPercent Setter
	MinWidthPercent() types.TVTConstraintPercent          // property MinWidthPercent Getter
	SetMinWidthPercent(value types.TVTConstraintPercent)  // property MinWidthPercent Setter
	SetOnChange(fn TNotifyEvent)                          // property event
}

type TVTFixedAreaConstraints struct {
	TPersistent
}

func (m *TVTFixedAreaConstraints) MaxHeightPercent() types.TVTConstraintPercent {
	if !m.IsValid() {
		return 0
	}
	r := vTFixedAreaConstraintsAPI().SysCallN(1, 0, m.Instance())
	return types.TVTConstraintPercent(r)
}

func (m *TVTFixedAreaConstraints) SetMaxHeightPercent(value types.TVTConstraintPercent) {
	if !m.IsValid() {
		return
	}
	vTFixedAreaConstraintsAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TVTFixedAreaConstraints) MaxWidthPercent() types.TVTConstraintPercent {
	if !m.IsValid() {
		return 0
	}
	r := vTFixedAreaConstraintsAPI().SysCallN(2, 0, m.Instance())
	return types.TVTConstraintPercent(r)
}

func (m *TVTFixedAreaConstraints) SetMaxWidthPercent(value types.TVTConstraintPercent) {
	if !m.IsValid() {
		return
	}
	vTFixedAreaConstraintsAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TVTFixedAreaConstraints) MinHeightPercent() types.TVTConstraintPercent {
	if !m.IsValid() {
		return 0
	}
	r := vTFixedAreaConstraintsAPI().SysCallN(3, 0, m.Instance())
	return types.TVTConstraintPercent(r)
}

func (m *TVTFixedAreaConstraints) SetMinHeightPercent(value types.TVTConstraintPercent) {
	if !m.IsValid() {
		return
	}
	vTFixedAreaConstraintsAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TVTFixedAreaConstraints) MinWidthPercent() types.TVTConstraintPercent {
	if !m.IsValid() {
		return 0
	}
	r := vTFixedAreaConstraintsAPI().SysCallN(4, 0, m.Instance())
	return types.TVTConstraintPercent(r)
}

func (m *TVTFixedAreaConstraints) SetMinWidthPercent(value types.TVTConstraintPercent) {
	if !m.IsValid() {
		return
	}
	vTFixedAreaConstraintsAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TVTFixedAreaConstraints) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, vTFixedAreaConstraintsAPI(), api.MakeEventDataPtr(cb))
}

// NewVTFixedAreaConstraints class constructor
func NewVTFixedAreaConstraints(owner IVTHeader) IVTFixedAreaConstraints {
	r := vTFixedAreaConstraintsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsVTFixedAreaConstraints(r)
}

var (
	vTFixedAreaConstraintsOnce   base.Once
	vTFixedAreaConstraintsImport *imports.Imports = nil
)

func vTFixedAreaConstraintsAPI() *imports.Imports {
	vTFixedAreaConstraintsOnce.Do(func() {
		vTFixedAreaConstraintsImport = api.NewDefaultImports()
		vTFixedAreaConstraintsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTFixedAreaConstraints_Create", 0), // constructor NewVTFixedAreaConstraints
			/* 1 */ imports.NewTable("TVTFixedAreaConstraints_MaxHeightPercent", 0), // property MaxHeightPercent
			/* 2 */ imports.NewTable("TVTFixedAreaConstraints_MaxWidthPercent", 0), // property MaxWidthPercent
			/* 3 */ imports.NewTable("TVTFixedAreaConstraints_MinHeightPercent", 0), // property MinHeightPercent
			/* 4 */ imports.NewTable("TVTFixedAreaConstraints_MinWidthPercent", 0), // property MinWidthPercent
			/* 5 */ imports.NewTable("TVTFixedAreaConstraints_OnChange", 0), // event OnChange
		}
	})
	return vTFixedAreaConstraintsImport
}
