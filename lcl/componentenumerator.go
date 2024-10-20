//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IComponentEnumerator Parent: IObject
type IComponentEnumerator interface {
	IObject
	Current() IComponent    // property
	GetCurrent() IComponent // function
	MoveNext() bool         // function
}

// TComponentEnumerator Parent: TObject
type TComponentEnumerator struct {
	TObject
}

func NewComponentEnumerator(AComponent IComponent) IComponentEnumerator {
	r1 := componentEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AComponent))
	return AsComponentEnumerator(r1)
}

func (m *TComponentEnumerator) Current() IComponent {
	r1 := componentEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsComponent(r1)
}

func (m *TComponentEnumerator) GetCurrent() IComponent {
	r1 := componentEnumeratorImportAPI().SysCallN(3, m.Instance())
	return AsComponent(r1)
}

func (m *TComponentEnumerator) MoveNext() bool {
	r1 := componentEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func ComponentEnumeratorClass() TClass {
	ret := componentEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	componentEnumeratorImport       *imports.Imports = nil
	componentEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ComponentEnumerator_Class", 0),
		/*1*/ imports.NewTable("ComponentEnumerator_Create", 0),
		/*2*/ imports.NewTable("ComponentEnumerator_Current", 0),
		/*3*/ imports.NewTable("ComponentEnumerator_GetCurrent", 0),
		/*4*/ imports.NewTable("ComponentEnumerator_MoveNext", 0),
	}
)

func componentEnumeratorImportAPI() *imports.Imports {
	if componentEnumeratorImport == nil {
		componentEnumeratorImport = NewDefaultImports()
		componentEnumeratorImport.SetImportTable(componentEnumeratorImportTables)
		componentEnumeratorImportTables = nil
	}
	return componentEnumeratorImport
}
