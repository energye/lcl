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

// ILCLComponent Parent: IComponent
type ILCLComponent interface {
	IComponent
	LCLRefCount() int32                         // property
	RemoveAllHandlersOfObject(AnObject IObject) // procedure
	IncLCLRefCount()                            // procedure
	DecLCLRefCount()                            // procedure
}

// TLCLComponent Parent: TComponent
type TLCLComponent struct {
	TComponent
}

func NewLCLComponent(TheOwner IComponent) ILCLComponent {
	r1 := LCL().SysCallN(3451, GetObjectUintptr(TheOwner))
	return AsLCLComponent(r1)
}

func (m *TLCLComponent) LCLRefCount() int32 {
	r1 := LCL().SysCallN(3454, m.Instance())
	return int32(r1)
}

func LCLComponentClass() TClass {
	ret := LCL().SysCallN(3450)
	return TClass(ret)
}

func (m *TLCLComponent) RemoveAllHandlersOfObject(AnObject IObject) {
	LCL().SysCallN(3455, m.Instance(), GetObjectUintptr(AnObject))
}

func (m *TLCLComponent) IncLCLRefCount() {
	LCL().SysCallN(3453, m.Instance())
}

func (m *TLCLComponent) DecLCLRefCount() {
	LCL().SysCallN(3452, m.Instance())
}
