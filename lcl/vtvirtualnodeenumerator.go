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
	"github.com/energye/lcl/types"
)

// IVTVirtualNodeEnumerator Parent: IObject
type IVTVirtualNodeEnumerator interface {
	IObject
	MoveNext() bool              // function
	Current() types.PVirtualNode // property Current Getter
}

type TVTVirtualNodeEnumerator struct {
	TObject
}

func (m *TVTVirtualNodeEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := vTVirtualNodeEnumeratorAPI().SysCallN(0, m.Instance())
	return api.GoBool(r)
}

func (m *TVTVirtualNodeEnumerator) Current() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := vTVirtualNodeEnumeratorAPI().SysCallN(1, m.Instance())
	return types.PVirtualNode(r)
}

var (
	vTVirtualNodeEnumeratorOnce   base.Once
	vTVirtualNodeEnumeratorImport *imports.Imports = nil
)

func vTVirtualNodeEnumeratorAPI() *imports.Imports {
	vTVirtualNodeEnumeratorOnce.Do(func() {
		vTVirtualNodeEnumeratorImport = api.NewDefaultImports()
		vTVirtualNodeEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTVirtualNodeEnumerator_MoveNext", 0), // function MoveNext
			/* 1 */ imports.NewTable("TVTVirtualNodeEnumerator_Current", 0), // property Current
		}
	})
	return vTVirtualNodeEnumeratorImport
}
