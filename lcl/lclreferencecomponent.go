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

// ILCLReferenceComponent Parent: ILCLComponent
type ILCLReferenceComponent interface {
	ILCLComponent
	HandleAllocated() bool    // property HandleAllocated Getter
	ReferenceAllocated() bool // property ReferenceAllocated Getter
}

type TLCLReferenceComponent struct {
	TLCLComponent
}

func (m *TLCLReferenceComponent) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := lCLReferenceComponentAPI().SysCallN(0, m.Instance())
	return api.GoBool(r)
}

func (m *TLCLReferenceComponent) ReferenceAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := lCLReferenceComponentAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

var (
	lCLReferenceComponentOnce   base.Once
	lCLReferenceComponentImport *imports.Imports = nil
)

func lCLReferenceComponentAPI() *imports.Imports {
	lCLReferenceComponentOnce.Do(func() {
		lCLReferenceComponentImport = api.NewDefaultImports()
		lCLReferenceComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLCLReferenceComponent_HandleAllocated", 0), // property HandleAllocated
			/* 1 */ imports.NewTable("TLCLReferenceComponent_ReferenceAllocated", 0), // property ReferenceAllocated
		}
	})
	return lCLReferenceComponentImport
}
