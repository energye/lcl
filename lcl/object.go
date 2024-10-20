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
	"github.com/energye/lcl/api/imports"
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
	r1 := objectImportAPI().SysCallN(4)
	return AsObject(r1)
}

func (m *TObject) Equals(Obj IObject) bool {
	r1 := objectImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Obj))
	return GoBool(r1)
}

func (m *TObject) GetHashCode() uint32 {
	r1 := objectImportAPI().SysCallN(7, m.Instance())
	return uint32(r1)
}

func ObjectClass() TClass {
	ret := objectImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TObject) ToString() string {
	r1 := objectImportAPI().SysCallN(10, m.Instance())
	return GoStr(r1)
}

func (m *TObject) InheritsFrom(AClass TClass) bool {
	r1 := objectImportAPI().SysCallN(8, m.Instance(), uintptr(AClass))
	return GoBool(r1)
}

func (m *TObject) ClassType() TClass {
	r1 := objectImportAPI().SysCallN(3, m.Instance())
	return TClass(r1)
}

func (m *TObject) ClassName() string {
	r1 := objectImportAPI().SysCallN(1, m.Instance())
	return GoStr(r1)
}

func (m *TObject) ClassParent() TClass {
	r1 := objectImportAPI().SysCallN(2, m.Instance())
	return TClass(r1)
}

func (m *TObject) InstanceSize() (resultInt64 int64) {
	objectImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TObject) Free() {
	m.free(6)
}

var (
	objectImport       *imports.Imports = nil
	objectImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Object_Class", 0),
		/*1*/ imports.NewTable("Object_ClassName", 0),
		/*2*/ imports.NewTable("Object_ClassParent", 0),
		/*3*/ imports.NewTable("Object_ClassType", 0),
		/*4*/ imports.NewTable("Object_Create", 0),
		/*5*/ imports.NewTable("Object_Equals", 0),
		/*6*/ imports.NewTable("Object_Free", 0),
		/*7*/ imports.NewTable("Object_GetHashCode", 0),
		/*8*/ imports.NewTable("Object_InheritsFrom", 0),
		/*9*/ imports.NewTable("Object_InstanceSize", 0),
		/*10*/ imports.NewTable("Object_ToString", 0),
	}
)

func objectImportAPI() *imports.Imports {
	if objectImport == nil {
		objectImport = NewDefaultImports()
		objectImport.SetImportTable(objectImportTables)
		objectImportTables = nil
	}
	return objectImport
}
