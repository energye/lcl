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

// IFPListEnumerator Parent: IObject
type IFPListEnumerator interface {
	IObject
	GetCurrent() uintptr // function
	MoveNext() bool      // function
	Current() uintptr    // property Current Getter
}

type TFPListEnumerator struct {
	TObject
}

func (m *TFPListEnumerator) GetCurrent() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := fPListEnumeratorAPI().SysCallN(1, m.Instance())
	return uintptr(r)
}

func (m *TFPListEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := fPListEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TFPListEnumerator) Current() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := fPListEnumeratorAPI().SysCallN(3, m.Instance())
	return uintptr(r)
}

// NewFPListEnumerator class constructor
func NewFPListEnumerator(list IFPList) IFPListEnumerator {
	r := fPListEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(list))
	return AsFPListEnumerator(r)
}

var (
	fPListEnumeratorOnce   base.Once
	fPListEnumeratorImport *imports.Imports = nil
)

func fPListEnumeratorAPI() *imports.Imports {
	fPListEnumeratorOnce.Do(func() {
		fPListEnumeratorImport = api.NewDefaultImports()
		fPListEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPListEnumerator_Create", 0), // constructor NewFPListEnumerator
			/* 1 */ imports.NewTable("TFPListEnumerator_GetCurrent", 0), // function GetCurrent
			/* 2 */ imports.NewTable("TFPListEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TFPListEnumerator_Current", 0), // property Current
		}
	})
	return fPListEnumeratorImport
}
