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

// IVTVirtualNodeEnumerationWrap Parent: IObject
type IVTVirtualNodeEnumerationWrap interface {
	IObject
	GetEnumerator() IVTVirtualNodeEnumerator // function
}

type TVTVirtualNodeEnumerationWrap struct {
	TObject
}

func (m *TVTVirtualNodeEnumerationWrap) GetEnumerator() IVTVirtualNodeEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := vTVirtualNodeEnumerationWrapAPI().SysCallN(0, m.Instance())
	return AsVTVirtualNodeEnumerator(r)
}

var (
	vTVirtualNodeEnumerationWrapOnce   base.Once
	vTVirtualNodeEnumerationWrapImport *imports.Imports = nil
)

func vTVirtualNodeEnumerationWrapAPI() *imports.Imports {
	vTVirtualNodeEnumerationWrapOnce.Do(func() {
		vTVirtualNodeEnumerationWrapImport = api.NewDefaultImports()
		vTVirtualNodeEnumerationWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTVirtualNodeEnumerationWrap_GetEnumerator", 0), // function GetEnumerator
		}
	})
	return vTVirtualNodeEnumerationWrapImport
}
