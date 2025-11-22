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

// IToolBarEnumerator Parent: IObject
type IToolBarEnumerator interface {
	IObject
	MoveNext() bool       // function
	Current() IToolButton // property Current Getter
}

type TToolBarEnumerator struct {
	TObject
}

func (m *TToolBarEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarEnumeratorAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBarEnumerator) Current() IToolButton {
	if !m.IsValid() {
		return nil
	}
	r := toolBarEnumeratorAPI().SysCallN(2, m.Instance())
	return AsToolButton(r)
}

// NewToolBarEnumerator class constructor
func NewToolBarEnumerator(toolBar IToolBar) IToolBarEnumerator {
	r := toolBarEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(toolBar))
	return AsToolBarEnumerator(r)
}

var (
	toolBarEnumeratorOnce   base.Once
	toolBarEnumeratorImport *imports.Imports = nil
)

func toolBarEnumeratorAPI() *imports.Imports {
	toolBarEnumeratorOnce.Do(func() {
		toolBarEnumeratorImport = api.NewDefaultImports()
		toolBarEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TToolBarEnumerator_Create", 0), // constructor NewToolBarEnumerator
			/* 1 */ imports.NewTable("TToolBarEnumerator_MoveNext", 0), // function MoveNext
			/* 2 */ imports.NewTable("TToolBarEnumerator_Current", 0), // property Current
		}
	})
	return toolBarEnumeratorImport
}
