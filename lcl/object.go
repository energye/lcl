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
)

// IObject Parent: base.IBase
type IObject interface {
	base.IBase
	Equals(obj IObject) bool // function
	ToString() string        // function
	Free()                   // procedure
	// Dispatch
	//  message handling routines
	Dispatch(message *uintptr) // procedure
}

type TObject struct {
	base.TBase
}

func (m *TObject) Equals(obj IObject) bool {
	if !m.IsValid() {
		return false
	}
	r := objectAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(obj))
	return api.GoBool(r)
}

func (m *TObject) ToString() string {
	if !m.IsValid() {
		return ""
	}
	r := objectAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
}

func (m *TObject) Free() {
	if !m.IsValid() {
		return
	}
	objectAPI().SysCallN(3, m.Instance())
	m.TBase.Free()
}

func (m *TObject) Dispatch(message *uintptr) {
	if !m.IsValid() {
		return
	}
	messagePtr := uintptr(*message)
	objectAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&messagePtr)))
	*message = uintptr(messagePtr)
}

// NewObject class constructor
func NewObject() IObject {
	r := objectAPI().SysCallN(0)
	return AsObject(r)
}

var (
	objectOnce   base.Once
	objectImport *imports.Imports = nil
)

func objectAPI() *imports.Imports {
	objectOnce.Do(func() {
		objectImport = api.NewDefaultImports()
		objectImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TObject_Create", 0), // constructor NewObject
			/* 1 */ imports.NewTable("TObject_Equals", 0), // function Equals
			/* 2 */ imports.NewTable("TObject_ToString", 0), // function ToString
			/* 3 */ imports.NewTable("TObject_Free", 0), // procedure Free
			/* 4 */ imports.NewTable("TObject_Dispatch", 0), // procedure Dispatch
		}
	})
	return objectImport
}
