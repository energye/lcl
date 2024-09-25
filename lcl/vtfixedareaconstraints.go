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

// IVTFixedAreaConstraints Parent: IPersistent
type IVTFixedAreaConstraints interface {
	IPersistent
	MaxHeightPercent() TVTConstraintPercent          // property
	SetMaxHeightPercent(AValue TVTConstraintPercent) // property
	MaxWidthPercent() TVTConstraintPercent           // property
	SetMaxWidthPercent(AValue TVTConstraintPercent)  // property
	MinHeightPercent() TVTConstraintPercent          // property
	SetMinHeightPercent(AValue TVTConstraintPercent) // property
	MinWidthPercent() TVTConstraintPercent           // property
	SetMinWidthPercent(AValue TVTConstraintPercent)  // property
	SetOnChange(fn TNotifyEvent)                     // property event
}

// TVTFixedAreaConstraints Parent: TPersistent
type TVTFixedAreaConstraints struct {
	TPersistent
	changePtr uintptr
}

func NewVTFixedAreaConstraints(AOwner IVTHeader) IVTFixedAreaConstraints {
	r1 := LCL().SysCallN(5893, GetObjectUintptr(AOwner))
	return AsVTFixedAreaConstraints(r1)
}

func (m *TVTFixedAreaConstraints) MaxHeightPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5894, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMaxHeightPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5894, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MaxWidthPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5895, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMaxWidthPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5895, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MinHeightPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5896, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMinHeightPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5896, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MinWidthPercent() TVTConstraintPercent {
	r1 := LCL().SysCallN(5897, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMinWidthPercent(AValue TVTConstraintPercent) {
	LCL().SysCallN(5897, 1, m.Instance(), uintptr(AValue))
}

func VTFixedAreaConstraintsClass() TClass {
	ret := LCL().SysCallN(5892)
	return TClass(ret)
}

func (m *TVTFixedAreaConstraints) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5898, m.Instance(), m.changePtr)
}
