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

// IInterfacedObject Parent: IObject
type IInterfacedObject interface {
	IObject
	RefCount() int32 // property RefCount Getter
	// Release 当前实例以接口形式返回时, 释放当前对象的引用, 适用于 Interface 类型指针
	Release()
}

type TInterfacedObject struct {
	TObject
}

func (m *TInterfacedObject) RefCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := interfacedObjectAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TInterfacedObject) Release() {
	if !m.IsValid() {
		return
	}
	var data = m.Instance()
	Release(uintptr(base.UnsafePointer(&data)))
	m.TBase.Free()
}

// NewInterfacedObject class constructor
func NewInterfacedObject() IInterfacedObject {
	r := interfacedObjectAPI().SysCallN(0)
	return AsInterfacedObject(r)
}

var (
	interfacedObjectOnce   base.Once
	interfacedObjectImport *imports.Imports = nil
)

func interfacedObjectAPI() *imports.Imports {
	interfacedObjectOnce.Do(func() {
		interfacedObjectImport = api.NewDefaultImports()
		interfacedObjectImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TInterfacedObject_Create", 0), // constructor NewInterfacedObject
			/* 1 */ imports.NewTable("TInterfacedObject_RefCount", 0), // property RefCount
		}
	})
	return interfacedObjectImport
}
