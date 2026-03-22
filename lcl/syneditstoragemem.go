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

// ISynEditStorageMem Parent: IObject
type ISynEditStorageMem interface {
	IObject
	InsertRows(index int32, count int32) // procedure
	DeleteRows(index int32, count int32) // procedure
	Capacity() int32                     // property Capacity Getter
	SetCapacity(value int32)             // property Capacity Setter
	// Count
	//  Capacity must be maintained by owner (Shrink)
	Count() int32         // property Count Getter
	SetCount(value int32) // property Count Setter
}

type TSynEditStorageMem struct {
	TObject
}

func (m *TSynEditStorageMem) InsertRows(index int32, count int32) {
	if !m.IsValid() {
		return
	}
	synEditStorageMemAPI().SysCallN(1, m.Instance(), uintptr(index), uintptr(count))
}

func (m *TSynEditStorageMem) DeleteRows(index int32, count int32) {
	if !m.IsValid() {
		return
	}
	synEditStorageMemAPI().SysCallN(2, m.Instance(), uintptr(index), uintptr(count))
}

func (m *TSynEditStorageMem) Capacity() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStorageMemAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditStorageMem) SetCapacity(value int32) {
	if !m.IsValid() {
		return
	}
	synEditStorageMemAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditStorageMem) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStorageMemAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditStorageMem) SetCount(value int32) {
	if !m.IsValid() {
		return
	}
	synEditStorageMemAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

// NewSynEditStorageMem class constructor
func NewSynEditStorageMem() ISynEditStorageMem {
	r := synEditStorageMemAPI().SysCallN(0)
	return AsSynEditStorageMem(r)
}

var (
	synEditStorageMemOnce   base.Once
	synEditStorageMemImport *imports.Imports = nil
)

func synEditStorageMemAPI() *imports.Imports {
	synEditStorageMemOnce.Do(func() {
		synEditStorageMemImport = api.NewDefaultImports()
		synEditStorageMemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditStorageMem_Create", 0), // constructor NewSynEditStorageMem
			/* 1 */ imports.NewTable("TSynEditStorageMem_InsertRows", 0), // procedure InsertRows
			/* 2 */ imports.NewTable("TSynEditStorageMem_DeleteRows", 0), // procedure DeleteRows
			/* 3 */ imports.NewTable("TSynEditStorageMem_Capacity", 0), // property Capacity
			/* 4 */ imports.NewTable("TSynEditStorageMem_Count", 0), // property Count
		}
	})
	return synEditStorageMemImport
}
