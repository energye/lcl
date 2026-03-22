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

// ISynEditMark Parent: IObject
type ISynEditMark interface {
	IObject
	IncChangeLock()                  // procedure
	DecChangeLock()                  // procedure
	OwnerEdit() ISynEditBase         // property OwnerEdit Getter
	SetOwnerEdit(value ISynEditBase) // property OwnerEdit Setter
	OldLine() int32                  // property OldLine Getter
	Line() int32                     // property Line Getter
	SetLine(value int32)             // property Line Setter
	Column() int32                   // property Column Getter
	SetColumn(value int32)           // property Column Setter
	Priority() int32                 // property Priority Getter
	SetPriority(value int32)         // property Priority Setter
	Visible() bool                   // property Visible Getter
	SetVisible(value bool)           // property Visible Setter
	BookmarkNumber() int32           // property BookmarkNumber Getter
	SetBookmarkNumber(value int32)   // property BookmarkNumber Setter
	IsBookmark() bool                // property IsBookmark Getter
	// InternalImage
	//  InternalImage: Use Internal bookmark image 0..9;
	//  Ignore "BookMarkOpt.BookmarkImages" or "ImageList"
	InternalImage() bool         // property InternalImage Getter
	SetInternalImage(value bool) // property InternalImage Setter
	// ImageIndex
	//  ImageIndex: Index in "BookMarkOpt.BookmarkImages" or "ImageList"
	ImageIndex() int32         // property ImageIndex Getter
	SetImageIndex(value int32) // property ImageIndex Setter
	// ImageList
	//  ImageList: If assigned, then use instead of "BookMarkOpt.BookmarkImages"
	//  Must have same width as "BookMarkOpt.BookmarkImages"
	ImageList() ICustomImageList         // property ImageList Getter
	SetImageList(value ICustomImageList) // property ImageList Setter
}

type TSynEditMark struct {
	TObject
}

func (m *TSynEditMark) IncChangeLock() {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(1, m.Instance())
}

func (m *TSynEditMark) DecChangeLock() {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(2, m.Instance())
}

func (m *TSynEditMark) OwnerEdit() ISynEditBase {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkAPI().SysCallN(3, 0, m.Instance())
	return AsSynEditBase(r)
}

func (m *TSynEditMark) SetOwnerEdit(value ISynEditBase) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditMark) OldLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TSynEditMark) Line() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMark) SetLine(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMark) Column() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMark) SetColumn(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMark) Priority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMark) SetPriority(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMark) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMarkAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMark) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditMark) BookmarkNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMark) SetBookmarkNumber(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMark) IsBookmark() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMarkAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMark) InternalImage() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMarkAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMark) SetInternalImage(value bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditMark) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMark) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMark) ImageList() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkAPI().SysCallN(13, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TSynEditMark) SetImageList(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	synEditMarkAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynEditMark class constructor
func NewSynEditMark(synEdit ISynEditBase) ISynEditMark {
	r := synEditMarkAPI().SysCallN(0, base.GetObjectUintptr(synEdit))
	return AsSynEditMark(r)
}

var (
	synEditMarkOnce   base.Once
	synEditMarkImport *imports.Imports = nil
)

func synEditMarkAPI() *imports.Imports {
	synEditMarkOnce.Do(func() {
		synEditMarkImport = api.NewDefaultImports()
		synEditMarkImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMark_Create", 0), // constructor NewSynEditMark
			/* 1 */ imports.NewTable("TSynEditMark_IncChangeLock", 0), // procedure IncChangeLock
			/* 2 */ imports.NewTable("TSynEditMark_DecChangeLock", 0), // procedure DecChangeLock
			/* 3 */ imports.NewTable("TSynEditMark_OwnerEdit", 0), // property OwnerEdit
			/* 4 */ imports.NewTable("TSynEditMark_OldLine", 0), // property OldLine
			/* 5 */ imports.NewTable("TSynEditMark_Line", 0), // property Line
			/* 6 */ imports.NewTable("TSynEditMark_Column", 0), // property Column
			/* 7 */ imports.NewTable("TSynEditMark_Priority", 0), // property Priority
			/* 8 */ imports.NewTable("TSynEditMark_Visible", 0), // property Visible
			/* 9 */ imports.NewTable("TSynEditMark_BookmarkNumber", 0), // property BookmarkNumber
			/* 10 */ imports.NewTable("TSynEditMark_IsBookmark", 0), // property IsBookmark
			/* 11 */ imports.NewTable("TSynEditMark_InternalImage", 0), // property InternalImage
			/* 12 */ imports.NewTable("TSynEditMark_ImageIndex", 0), // property ImageIndex
			/* 13 */ imports.NewTable("TSynEditMark_ImageList", 0), // property ImageList
		}
	})
	return synEditMarkImport
}
