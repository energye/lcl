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

// ISynManagedStorageMem Parent: ISynEditStorageMem
type ISynManagedStorageMem interface {
	ISynEditStorageMem
}

type TSynManagedStorageMem struct {
	TSynEditStorageMem
}

// NewSynManagedStorageMem class constructor
func NewSynManagedStorageMem() ISynManagedStorageMem {
	r := synManagedStorageMemAPI().SysCallN(0)
	return AsSynManagedStorageMem(r)
}

var (
	synManagedStorageMemOnce   base.Once
	synManagedStorageMemImport *imports.Imports = nil
)

func synManagedStorageMemAPI() *imports.Imports {
	synManagedStorageMemOnce.Do(func() {
		synManagedStorageMemImport = api.NewDefaultImports()
		synManagedStorageMemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynManagedStorageMem_Create", 0), // constructor NewSynManagedStorageMem
		}
	})
	return synManagedStorageMemImport
}
