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

// ICustomImageListResolutionEnumerator Parent: IObject
type ICustomImageListResolutionEnumerator interface {
	IObject
	GetEnumerator() ICustomImageListResolutionEnumerator // function
	MoveNext() bool                                      // function
	Current() ICustomImageListResolution                 // property Current Getter
}

type TCustomImageListResolutionEnumerator struct {
	TObject
}

func (m *TCustomImageListResolutionEnumerator) GetEnumerator() ICustomImageListResolutionEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := customImageListResolutionEnumeratorAPI().SysCallN(1, m.Instance())
	return AsCustomImageListResolutionEnumerator(r)
}

func (m *TCustomImageListResolutionEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageListResolutionEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImageListResolutionEnumerator) Current() ICustomImageListResolution {
	if !m.IsValid() {
		return nil
	}
	r := customImageListResolutionEnumeratorAPI().SysCallN(3, m.Instance())
	return AsCustomImageListResolution(r)
}

// NewCustomImageListResolutionEnumerator class constructor
func NewCustomImageListResolutionEnumerator(imgList ICustomImageList, desc bool) ICustomImageListResolutionEnumerator {
	r := customImageListResolutionEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(imgList), api.PasBool(desc))
	return AsCustomImageListResolutionEnumerator(r)
}

var (
	customImageListResolutionEnumeratorOnce   base.Once
	customImageListResolutionEnumeratorImport *imports.Imports = nil
)

func customImageListResolutionEnumeratorAPI() *imports.Imports {
	customImageListResolutionEnumeratorOnce.Do(func() {
		customImageListResolutionEnumeratorImport = api.NewDefaultImports()
		customImageListResolutionEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomImageListResolutionEnumerator_Create", 0), // constructor NewCustomImageListResolutionEnumerator
			/* 1 */ imports.NewTable("TCustomImageListResolutionEnumerator_GetEnumerator", 0), // function GetEnumerator
			/* 2 */ imports.NewTable("TCustomImageListResolutionEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TCustomImageListResolutionEnumerator_Current", 0), // property Current
		}
	})
	return customImageListResolutionEnumeratorImport
}
