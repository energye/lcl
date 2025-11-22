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

// ILazAccessibleObjectEnumerator Parent: IObject
type ILazAccessibleObjectEnumerator interface {
	IObject
	Current() ILazAccessibleObject // property Current Getter
}

type TLazAccessibleObjectEnumerator struct {
	TObject
}

func (m *TLazAccessibleObjectEnumerator) Current() ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectEnumeratorAPI().SysCallN(0, m.Instance())
	return AsLazAccessibleObject(r)
}

var (
	lazAccessibleObjectEnumeratorOnce   base.Once
	lazAccessibleObjectEnumeratorImport *imports.Imports = nil
)

func lazAccessibleObjectEnumeratorAPI() *imports.Imports {
	lazAccessibleObjectEnumeratorOnce.Do(func() {
		lazAccessibleObjectEnumeratorImport = api.NewDefaultImports()
		lazAccessibleObjectEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazAccessibleObjectEnumerator_Current", 0), // property Current
		}
	})
	return lazAccessibleObjectEnumeratorImport
}
