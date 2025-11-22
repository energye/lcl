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

// IImageList Parent: IDragImageList
type IImageList interface {
	IDragImageList
}

type TImageList struct {
	TDragImageList
}

// NewImageList class constructor
func NewImageList(owner IComponent) IImageList {
	r := imageListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsImageList(r)
}

// NewImageListSize class constructor
func NewImageListSize(width int32, height int32) IImageList {
	r := imageListAPI().SysCallN(1, uintptr(width), uintptr(height))
	return AsImageList(r)
}

func TImageListClass() types.TClass {
	r := imageListAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	imageListOnce   base.Once
	imageListImport *imports.Imports = nil
)

func imageListAPI() *imports.Imports {
	imageListOnce.Do(func() {
		imageListImport = api.NewDefaultImports()
		imageListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TImageList_Create", 0), // constructor NewImageList
			/* 1 */ imports.NewTable("TImageList_CreateSize", 0), // constructor NewImageListSize
			/* 2 */ imports.NewTable("TImageList_TClass", 0), // function TImageListClass
		}
	})
	return imageListImport
}
