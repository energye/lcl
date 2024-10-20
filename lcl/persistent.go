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
	r1 := persistentImportAPI().SysCallN(2)
	return AsPersistent(r1)
}

func (m *TPersistent) GetNamePath() string {
	r1 := persistentImportAPI().SysCallN(6, m.Instance())
	return GoStr(r1)
}

func PersistentClass() TClass {
	ret := persistentImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TPersistent) Assign(Source IPersistent) {
	persistentImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Source))
}

func (m *TPersistent) FPOAttachObserver(AObserver IObject) {
	persistentImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TPersistent) FPODetachObserver(AObserver IObject) {
	persistentImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TPersistent) FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) {
	persistentImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(ASender), uintptr(AOperation), uintptr(Data))
}

var (
	persistentImport       *imports.Imports = nil
	persistentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Persistent_Assign", 0),
		/*1*/ imports.NewTable("Persistent_Class", 0),
		/*2*/ imports.NewTable("Persistent_Create", 0),
		/*3*/ imports.NewTable("Persistent_FPOAttachObserver", 0),
		/*4*/ imports.NewTable("Persistent_FPODetachObserver", 0),
		/*5*/ imports.NewTable("Persistent_FPONotifyObservers", 0),
		/*6*/ imports.NewTable("Persistent_GetNamePath", 0),
	}
)

func persistentImportAPI() *imports.Imports {
	if persistentImport == nil {
		persistentImport = NewDefaultImports()
		persistentImport.SetImportTable(persistentImportTables)
		persistentImportTables = nil
	}
	return persistentImport
}
