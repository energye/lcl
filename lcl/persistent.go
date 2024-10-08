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

// IPersistent Parent: IObject
type IPersistent interface {
	IObject
	GetNamePath() string                                                               // function
	Assign(Source IPersistent)                                                         // procedure
	FPOAttachObserver(AObserver IObject)                                               // procedure
	FPODetachObserver(AObserver IObject)                                               // procedure
	FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) // procedure
}

// TPersistent Parent: TObject
type TPersistent struct {
	TObject
}

func NewPersistent() IPersistent {
	r1 := LCL().SysCallN(4565)
	return AsPersistent(r1)
}

func (m *TPersistent) GetNamePath() string {
	r1 := LCL().SysCallN(4569, m.Instance())
	return GoStr(r1)
}

func PersistentClass() TClass {
	ret := LCL().SysCallN(4564)
	return TClass(ret)
}

func (m *TPersistent) Assign(Source IPersistent) {
	LCL().SysCallN(4563, m.Instance(), GetObjectUintptr(Source))
}

func (m *TPersistent) FPOAttachObserver(AObserver IObject) {
	LCL().SysCallN(4566, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TPersistent) FPODetachObserver(AObserver IObject) {
	LCL().SysCallN(4567, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TPersistent) FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) {
	LCL().SysCallN(4568, m.Instance(), GetObjectUintptr(ASender), uintptr(AOperation), uintptr(Data))
}
