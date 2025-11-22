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

// IWinControlEnumerator Parent: IObject
type IWinControlEnumerator interface {
	IObject
	GetEnumerator() IWinControlEnumerator // function
	MoveNext() bool                       // function
	Current() IControl                    // property Current Getter
}

type TWinControlEnumerator struct {
	TObject
}

func (m *TWinControlEnumerator) GetEnumerator() IWinControlEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := winControlEnumeratorAPI().SysCallN(1, m.Instance())
	return AsWinControlEnumerator(r)
}

func (m *TWinControlEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControlEnumerator) Current() IControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlEnumeratorAPI().SysCallN(3, m.Instance())
	return AsControl(r)
}

// NewWinControlEnumerator class constructor
func NewWinControlEnumerator(parent IWinControl, lowToHigh bool) IWinControlEnumerator {
	r := winControlEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(parent), api.PasBool(lowToHigh))
	return AsWinControlEnumerator(r)
}

var (
	winControlEnumeratorOnce   base.Once
	winControlEnumeratorImport *imports.Imports = nil
)

func winControlEnumeratorAPI() *imports.Imports {
	winControlEnumeratorOnce.Do(func() {
		winControlEnumeratorImport = api.NewDefaultImports()
		winControlEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWinControlEnumerator_Create", 0), // constructor NewWinControlEnumerator
			/* 1 */ imports.NewTable("TWinControlEnumerator_GetEnumerator", 0), // function GetEnumerator
			/* 2 */ imports.NewTable("TWinControlEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TWinControlEnumerator_Current", 0), // property Current
		}
	})
	return winControlEnumeratorImport
}
