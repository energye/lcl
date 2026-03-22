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

// ISynEditUndoItem Parent: IObject
type ISynEditUndoItem interface {
	IObject
	IsEqual(anItem ISynEditUndoItem) bool // function
	IsCaretInfo() bool                    // function
	PerformUndo(caller IObject) bool      // function
}

type TSynEditUndoItem struct {
	TObject
}

func (m *TSynEditUndoItem) IsEqual(anItem ISynEditUndoItem) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoItemAPI().SysCallN(0, m.Instance(), base.GetObjectUintptr(anItem))
	return api.GoBool(r)
}

func (m *TSynEditUndoItem) IsCaretInfo() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoItemAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoItem) PerformUndo(caller IObject) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoItemAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(caller))
	return api.GoBool(r)
}

var (
	synEditUndoItemOnce   base.Once
	synEditUndoItemImport *imports.Imports = nil
)

func synEditUndoItemAPI() *imports.Imports {
	synEditUndoItemOnce.Do(func() {
		synEditUndoItemImport = api.NewDefaultImports()
		synEditUndoItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditUndoItem_IsEqual", 0), // function IsEqual
			/* 1 */ imports.NewTable("TSynEditUndoItem_IsCaretInfo", 0), // function IsCaretInfo
			/* 2 */ imports.NewTable("TSynEditUndoItem_PerformUndo", 0), // function PerformUndo
		}
	})
	return synEditUndoItemImport
}
