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

// ISynBookMarkOpt Parent: IPersistent
type ISynBookMarkOpt interface {
	IPersistent
	BookmarkImages() ICustomImageList         // property BookmarkImages Getter
	SetBookmarkImages(value ICustomImageList) // property BookmarkImages Setter
	DrawBookmarksFirst() bool                 // property DrawBookmarksFirst Getter
	SetDrawBookmarksFirst(value bool)         // property DrawBookmarksFirst Setter
	EnableKeys() bool                         // property EnableKeys Getter
	SetEnableKeys(value bool)                 // property EnableKeys Setter
	GlyphsVisible() bool                      // property GlyphsVisible Getter
	SetGlyphsVisible(value bool)              // property GlyphsVisible Setter
	LeftMargin() int32                        // property LeftMargin Getter
	SetLeftMargin(value int32)                // property LeftMargin Setter
	Xoffset() int32                           // property Xoffset Getter
	SetXoffset(value int32)                   // property Xoffset Setter
	SetOnChange(fn TNotifyEvent)              // property event
}

type TSynBookMarkOpt struct {
	TPersistent
}

func (m *TSynBookMarkOpt) BookmarkImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := synBookMarkOptAPI().SysCallN(1, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TSynBookMarkOpt) SetBookmarkImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	synBookMarkOptAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynBookMarkOpt) DrawBookmarksFirst() bool {
	if !m.IsValid() {
		return false
	}
	r := synBookMarkOptAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBookMarkOpt) SetDrawBookmarksFirst(value bool) {
	if !m.IsValid() {
		return
	}
	synBookMarkOptAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBookMarkOpt) EnableKeys() bool {
	if !m.IsValid() {
		return false
	}
	r := synBookMarkOptAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBookMarkOpt) SetEnableKeys(value bool) {
	if !m.IsValid() {
		return
	}
	synBookMarkOptAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBookMarkOpt) GlyphsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := synBookMarkOptAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBookMarkOpt) SetGlyphsVisible(value bool) {
	if !m.IsValid() {
		return
	}
	synBookMarkOptAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBookMarkOpt) LeftMargin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBookMarkOptAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TSynBookMarkOpt) SetLeftMargin(value int32) {
	if !m.IsValid() {
		return
	}
	synBookMarkOptAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TSynBookMarkOpt) Xoffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBookMarkOptAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TSynBookMarkOpt) SetXoffset(value int32) {
	if !m.IsValid() {
		return
	}
	synBookMarkOptAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TSynBookMarkOpt) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, synBookMarkOptAPI(), api.MakeEventDataPtr(cb))
}

// NewSynBookMarkOpt class constructor
func NewSynBookMarkOpt(owner IComponent) ISynBookMarkOpt {
	r := synBookMarkOptAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynBookMarkOpt(r)
}

var (
	synBookMarkOptOnce   base.Once
	synBookMarkOptImport *imports.Imports = nil
)

func synBookMarkOptAPI() *imports.Imports {
	synBookMarkOptOnce.Do(func() {
		synBookMarkOptImport = api.NewDefaultImports()
		synBookMarkOptImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynBookMarkOpt_Create", 0), // constructor NewSynBookMarkOpt
			/* 1 */ imports.NewTable("TSynBookMarkOpt_BookmarkImages", 0), // property BookmarkImages
			/* 2 */ imports.NewTable("TSynBookMarkOpt_DrawBookmarksFirst", 0), // property DrawBookmarksFirst
			/* 3 */ imports.NewTable("TSynBookMarkOpt_EnableKeys", 0), // property EnableKeys
			/* 4 */ imports.NewTable("TSynBookMarkOpt_GlyphsVisible", 0), // property GlyphsVisible
			/* 5 */ imports.NewTable("TSynBookMarkOpt_LeftMargin", 0), // property LeftMargin
			/* 6 */ imports.NewTable("TSynBookMarkOpt_Xoffset", 0), // property Xoffset
			/* 7 */ imports.NewTable("TSynBookMarkOpt_OnChange", 0), // event OnChange
		}
	})
	return synBookMarkOptImport
}
