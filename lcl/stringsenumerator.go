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

// IStringsEnumerator Parent: IObject
type IStringsEnumerator interface {
	IObject
	GetCurrent() string // function
	MoveNext() bool     // function
	Current() string    // property Current Getter
}

type TStringsEnumerator struct {
	TObject
}

func (m *TStringsEnumerator) GetCurrent() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsEnumeratorAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TStringsEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := stringsEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TStringsEnumerator) Current() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsEnumeratorAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

// NewStringsEnumerator class constructor
func NewStringsEnumerator(strings IStrings) IStringsEnumerator {
	r := stringsEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(strings))
	return AsStringsEnumerator(r)
}

var (
	stringsEnumeratorOnce   base.Once
	stringsEnumeratorImport *imports.Imports = nil
)

func stringsEnumeratorAPI() *imports.Imports {
	stringsEnumeratorOnce.Do(func() {
		stringsEnumeratorImport = api.NewDefaultImports()
		stringsEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStringsEnumerator_Create", 0), // constructor NewStringsEnumerator
			/* 1 */ imports.NewTable("TStringsEnumerator_GetCurrent", 0), // function GetCurrent
			/* 2 */ imports.NewTable("TStringsEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TStringsEnumerator_Current", 0), // property Current
		}
	})
	return stringsEnumeratorImport
}
