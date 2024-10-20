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

// IImageList Parent: IDragImageList
type IImageList interface {
	IDragImageList
}

// TImageList Parent: TDragImageList
type TImageList struct {
	TDragImageList
}

func NewImageList(AOwner IComponent) IImageList {
	r1 := mageListImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsImageList(r1)
}

func ImageListClass() TClass {
	ret := mageListImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	mageListImport       *imports.Imports = nil
	mageListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ImageList_Class", 0),
		/*1*/ imports.NewTable("ImageList_Create", 0),
	}
)

func mageListImportAPI() *imports.Imports {
	if mageListImport == nil {
		mageListImport = NewDefaultImports()
		mageListImport.SetImportTable(mageListImportTables)
		mageListImportTables = nil
	}
	return mageListImport
}
