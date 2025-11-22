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

// IActionListEnumerator Parent: IObject
type IActionListEnumerator interface {
	IObject
	MoveNext() bool            // function
	Current() IContainedAction // property Current Getter
}

type TActionListEnumerator struct {
	TObject
}

func (m *TActionListEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := actionListEnumeratorAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TActionListEnumerator) Current() IContainedAction {
	if !m.IsValid() {
		return nil
	}
	r := actionListEnumeratorAPI().SysCallN(2, m.Instance())
	return AsContainedAction(r)
}

// NewActionListEnumerator class constructor
func NewActionListEnumerator(list ICustomActionList) IActionListEnumerator {
	r := actionListEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(list))
	return AsActionListEnumerator(r)
}

var (
	actionListEnumeratorOnce   base.Once
	actionListEnumeratorImport *imports.Imports = nil
)

func actionListEnumeratorAPI() *imports.Imports {
	actionListEnumeratorOnce.Do(func() {
		actionListEnumeratorImport = api.NewDefaultImports()
		actionListEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TActionListEnumerator_Create", 0), // constructor NewActionListEnumerator
			/* 1 */ imports.NewTable("TActionListEnumerator_MoveNext", 0), // function MoveNext
			/* 2 */ imports.NewTable("TActionListEnumerator_Current", 0), // property Current
		}
	})
	return actionListEnumeratorImport
}
