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

// IComponentEnumerator Parent: IObject
type IComponentEnumerator interface {
	IObject
	GetCurrent() IComponent // function
	MoveNext() bool         // function
	Current() IComponent    // property Current Getter
}

type TComponentEnumerator struct {
	TObject
}

func (m *TComponentEnumerator) GetCurrent() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := componentEnumeratorAPI().SysCallN(1, m.Instance())
	return AsComponent(r)
}

func (m *TComponentEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := componentEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TComponentEnumerator) Current() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := componentEnumeratorAPI().SysCallN(3, m.Instance())
	return AsComponent(r)
}

// NewComponentEnumerator class constructor
func NewComponentEnumerator(component IComponent) IComponentEnumerator {
	r := componentEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(component))
	return AsComponentEnumerator(r)
}

var (
	componentEnumeratorOnce   base.Once
	componentEnumeratorImport *imports.Imports = nil
)

func componentEnumeratorAPI() *imports.Imports {
	componentEnumeratorOnce.Do(func() {
		componentEnumeratorImport = api.NewDefaultImports()
		componentEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TComponentEnumerator_Create", 0), // constructor NewComponentEnumerator
			/* 1 */ imports.NewTable("TComponentEnumerator_GetCurrent", 0), // function GetCurrent
			/* 2 */ imports.NewTable("TComponentEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TComponentEnumerator_Current", 0), // property Current
		}
	})
	return componentEnumeratorImport
}
