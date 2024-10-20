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

// IDataModule Parent: IComponent
type IDataModule interface {
	IComponent
	DesignOffset() (resultPoint TPoint) // property
	SetDesignOffset(AValue *TPoint)     // property
	DesignSize() (resultPoint TPoint)   // property
	SetDesignSize(AValue *TPoint)       // property
	DesignPPI() int32                   // property
	SetDesignPPI(AValue int32)          // property
	OldCreateOrder() bool               // property
	SetOldCreateOrder(AValue bool)      // property
	SetOnCreate(fn TNotifyEvent)        // property event
	SetOnDestroy(fn TNotifyEvent)       // property event
}

// TDataModule Parent: TComponent
type TDataModule struct {
	TComponent
	createPtr  uintptr
	destroyPtr uintptr
}

func NewDataModule(AOwner IComponent) IDataModule {
	r1 := dataModuleImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsDataModule(r1)
}

func NewDataModuleNew(AOwner IComponent) IDataModule {
	r1 := dataModuleImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsDataModule(r1)
}

func NewDataModuleNew1(AOwner IComponent, CreateMode int32) IDataModule {
	r1 := dataModuleImportAPI().SysCallN(3, GetObjectUintptr(AOwner), uintptr(CreateMode))
	return AsDataModule(r1)
}

func (m *TDataModule) DesignOffset() (resultPoint TPoint) {
	dataModuleImportAPI().SysCallN(4, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDataModule) SetDesignOffset(AValue *TPoint) {
	dataModuleImportAPI().SysCallN(4, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDataModule) DesignSize() (resultPoint TPoint) {
	dataModuleImportAPI().SysCallN(6, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDataModule) SetDesignSize(AValue *TPoint) {
	dataModuleImportAPI().SysCallN(6, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDataModule) DesignPPI() int32 {
	r1 := dataModuleImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDataModule) SetDesignPPI(AValue int32) {
	dataModuleImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TDataModule) OldCreateOrder() bool {
	r1 := dataModuleImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDataModule) SetOldCreateOrder(AValue bool) {
	dataModuleImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func DataModuleClass() TClass {
	ret := dataModuleImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TDataModule) SetOnCreate(fn TNotifyEvent) {
	if m.createPtr != 0 {
		RemoveEventElement(m.createPtr)
	}
	m.createPtr = MakeEventDataPtr(fn)
	dataModuleImportAPI().SysCallN(8, m.Instance(), m.createPtr)
}

func (m *TDataModule) SetOnDestroy(fn TNotifyEvent) {
	if m.destroyPtr != 0 {
		RemoveEventElement(m.destroyPtr)
	}
	m.destroyPtr = MakeEventDataPtr(fn)
	dataModuleImportAPI().SysCallN(9, m.Instance(), m.destroyPtr)
}

var (
	dataModuleImport       *imports.Imports = nil
	dataModuleImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DataModule_Class", 0),
		/*1*/ imports.NewTable("DataModule_Create", 0),
		/*2*/ imports.NewTable("DataModule_CreateNew", 0),
		/*3*/ imports.NewTable("DataModule_CreateNew1", 0),
		/*4*/ imports.NewTable("DataModule_DesignOffset", 0),
		/*5*/ imports.NewTable("DataModule_DesignPPI", 0),
		/*6*/ imports.NewTable("DataModule_DesignSize", 0),
		/*7*/ imports.NewTable("DataModule_OldCreateOrder", 0),
		/*8*/ imports.NewTable("DataModule_SetOnCreate", 0),
		/*9*/ imports.NewTable("DataModule_SetOnDestroy", 0),
	}
)

func dataModuleImportAPI() *imports.Imports {
	if dataModuleImport == nil {
		dataModuleImport = NewDefaultImports()
		dataModuleImport.SetImportTable(dataModuleImportTables)
		dataModuleImportTables = nil
	}
	return dataModuleImport
}
