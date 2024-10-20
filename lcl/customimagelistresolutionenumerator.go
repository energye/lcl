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

// ICustomImageListResolutionEnumerator Parent: IObject
type ICustomImageListResolutionEnumerator interface {
	IObject
	Current() ICustomImageListResolution                 // property
	GetEnumerator() ICustomImageListResolutionEnumerator // function
	MoveNext() bool                                      // function
}

// TCustomImageListResolutionEnumerator Parent: TObject
type TCustomImageListResolutionEnumerator struct {
	TObject
}

func NewCustomImageListResolutionEnumerator(AImgList ICustomImageList, ADesc bool) ICustomImageListResolutionEnumerator {
	r1 := customImageListResolutionEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AImgList), PascalBool(ADesc))
	return AsCustomImageListResolutionEnumerator(r1)
}

func (m *TCustomImageListResolutionEnumerator) Current() ICustomImageListResolution {
	r1 := customImageListResolutionEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageListResolutionEnumerator) GetEnumerator() ICustomImageListResolutionEnumerator {
	r1 := customImageListResolutionEnumeratorImportAPI().SysCallN(3, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func (m *TCustomImageListResolutionEnumerator) MoveNext() bool {
	r1 := customImageListResolutionEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func CustomImageListResolutionEnumeratorClass() TClass {
	ret := customImageListResolutionEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customImageListResolutionEnumeratorImport       *imports.Imports = nil
	customImageListResolutionEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomImageListResolutionEnumerator_Class", 0),
		/*1*/ imports.NewTable("CustomImageListResolutionEnumerator_Create", 0),
		/*2*/ imports.NewTable("CustomImageListResolutionEnumerator_Current", 0),
		/*3*/ imports.NewTable("CustomImageListResolutionEnumerator_GetEnumerator", 0),
		/*4*/ imports.NewTable("CustomImageListResolutionEnumerator_MoveNext", 0),
	}
)

func customImageListResolutionEnumeratorImportAPI() *imports.Imports {
	if customImageListResolutionEnumeratorImport == nil {
		customImageListResolutionEnumeratorImport = NewDefaultImports()
		customImageListResolutionEnumeratorImport.SetImportTable(customImageListResolutionEnumeratorImportTables)
		customImageListResolutionEnumeratorImportTables = nil
	}
	return customImageListResolutionEnumeratorImport
}
