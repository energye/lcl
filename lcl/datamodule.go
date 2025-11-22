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

// IDataModule Parent: IComponent
type IDataModule interface {
	IComponent
	DesignOffset() types.TPoint         // property DesignOffset Getter
	SetDesignOffset(value types.TPoint) // property DesignOffset Setter
	DesignSize() types.TPoint           // property DesignSize Getter
	SetDesignSize(value types.TPoint)   // property DesignSize Setter
	DesignPPI() int32                   // property DesignPPI Getter
	SetDesignPPI(value int32)           // property DesignPPI Setter
	OldCreateOrder() bool               // property OldCreateOrder Getter
	SetOldCreateOrder(value bool)       // property OldCreateOrder Setter
	SetOnCreate(fn TNotifyEvent)        // property event
	SetOnDestroy(fn TNotifyEvent)       // property event
}

type TDataModule struct {
	TComponent
}

func (m *TDataModule) DesignOffset() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	dataModuleAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDataModule) SetDesignOffset(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	dataModuleAPI().SysCallN(3, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDataModule) DesignSize() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	dataModuleAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDataModule) SetDesignSize(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	dataModuleAPI().SysCallN(4, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDataModule) DesignPPI() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dataModuleAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TDataModule) SetDesignPPI(value int32) {
	if !m.IsValid() {
		return
	}
	dataModuleAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TDataModule) OldCreateOrder() bool {
	if !m.IsValid() {
		return false
	}
	r := dataModuleAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDataModule) SetOldCreateOrder(value bool) {
	if !m.IsValid() {
		return
	}
	dataModuleAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TDataModule) SetOnCreate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, dataModuleAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDataModule) SetOnDestroy(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, dataModuleAPI(), api.MakeEventDataPtr(cb))
}

// NewDataModule class constructor
func NewDataModule(owner IComponent) IDataModule {
	r := dataModuleAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsDataModule(r)
}

// NewDataModuleNew class constructor
func NewDataModuleNew(owner IComponent) IDataModule {
	r := dataModuleAPI().SysCallN(1, base.GetObjectUintptr(owner))
	return AsDataModule(r)
}

// NewDataModuleNewWithComponentInt class constructor
func NewDataModuleNewWithComponentInt(owner IComponent, createMode int32) IDataModule {
	r := dataModuleAPI().SysCallN(2, base.GetObjectUintptr(owner), uintptr(createMode))
	return AsDataModule(r)
}

func TDataModuleClass() types.TClass {
	r := dataModuleAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	dataModuleOnce   base.Once
	dataModuleImport *imports.Imports = nil
)

func dataModuleAPI() *imports.Imports {
	dataModuleOnce.Do(func() {
		dataModuleImport = api.NewDefaultImports()
		dataModuleImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDataModule_Create", 0), // constructor NewDataModule
			/* 1 */ imports.NewTable("TDataModule_CreateNew", 0), // constructor NewDataModuleNew
			/* 2 */ imports.NewTable("TDataModule_CreateNewWithComponentInt", 0), // constructor NewDataModuleNewWithComponentInt
			/* 3 */ imports.NewTable("TDataModule_DesignOffset", 0), // property DesignOffset
			/* 4 */ imports.NewTable("TDataModule_DesignSize", 0), // property DesignSize
			/* 5 */ imports.NewTable("TDataModule_DesignPPI", 0), // property DesignPPI
			/* 6 */ imports.NewTable("TDataModule_OldCreateOrder", 0), // property OldCreateOrder
			/* 7 */ imports.NewTable("TDataModule_OnCreate", 0), // event OnCreate
			/* 8 */ imports.NewTable("TDataModule_OnDestroy", 0), // event OnDestroy
			/* 9 */ imports.NewTable("TDataModule_TClass", 0), // function TDataModuleClass
		}
	})
	return dataModuleImport
}
