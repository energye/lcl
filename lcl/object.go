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

// IObject Root Interface
type IObject interface {
	Instance() uintptr
	UnsafeAddr() unsafePointer
	IsValid() bool
	Is() TIs
	SetInstance(instance unsafePointer)
	FreeAndNil()
	Nil()
	Equals(Obj IObject) bool           // function
	GetHashCode() uint32               // function
	ToString() string                  // function
	InheritsFrom(AClass TClass) bool   // function
	ClassType() TClass                 // function
	ClassName() string                 // function
	ClassParent() TClass               // function
	InstanceSize() (resultInt64 int64) // function
	Free()                             // procedure
}

// TObject Root Object
type TObject struct {
	instance unsafePointer
}

func NewObject() IObject {
	r1 := LCL().SysCallN(4391)
	return AsObject(r1)
}

func (m *TObject) Equals(Obj IObject) bool {
	r1 := LCL().SysCallN(4392, m.Instance(), GetObjectUintptr(Obj))
	return GoBool(r1)
}

func (m *TObject) GetHashCode() uint32 {
	r1 := LCL().SysCallN(4394, m.Instance())
	return uint32(r1)
}

func ObjectClass() TClass {
	ret := LCL().SysCallN(4387)
	return TClass(ret)
}

func (m *TObject) ToString() string {
	r1 := LCL().SysCallN(4397, m.Instance())
	return GoStr(r1)
}

func (m *TObject) InheritsFrom(AClass TClass) bool {
	r1 := LCL().SysCallN(4395, m.Instance(), uintptr(AClass))
	return GoBool(r1)
}

func (m *TObject) ClassType() TClass {
	r1 := LCL().SysCallN(4390, m.Instance())
	return TClass(r1)
}

func (m *TObject) ClassName() string {
	r1 := LCL().SysCallN(4388, m.Instance())
	return GoStr(r1)
}

func (m *TObject) ClassParent() TClass {
	r1 := LCL().SysCallN(4389, m.Instance())
	return TClass(r1)
}

func (m *TObject) InstanceSize() (resultInt64 int64) {
	LCL().SysCallN(4396, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TObject) Free() {
	m.free(4393)
}
