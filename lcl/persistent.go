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

// IPersistent Parent: IObject
type IPersistent interface {
	IObject
	GetNamePath() string       // function
	Assign(source IPersistent) // procedure
}

type TPersistent struct {
	TObject
}

func (m *TPersistent) GetNamePath() string {
	if !m.IsValid() {
		return ""
	}
	r := persistentAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TPersistent) Assign(source IPersistent) {
	if !m.IsValid() {
		return
	}
	persistentAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(source))
}

// NewPersistent class constructor
func NewPersistent() IPersistent {
	r := persistentAPI().SysCallN(0)
	return AsPersistent(r)
}

var (
	persistentOnce   base.Once
	persistentImport *imports.Imports = nil
)

func persistentAPI() *imports.Imports {
	persistentOnce.Do(func() {
		persistentImport = api.NewDefaultImports()
		persistentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPersistent_Create", 0), // constructor NewPersistent
			/* 1 */ imports.NewTable("TPersistent_GetNamePath", 0), // function GetNamePath
			/* 2 */ imports.NewTable("TPersistent_Assign", 0), // procedure Assign
		}
	})
	return persistentImport
}
