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

// IPaperSize Parent: IObject
type IPaperSize interface {
	IObject
	DefaultPapers() bool                      // property DefaultPapers Getter
	Width() int32                             // property Width Getter
	Height() int32                            // property Height Getter
	PaperName() string                        // property PaperName Getter
	SetPaperName(value string)                // property PaperName Setter
	DefaultPaperName() string                 // property DefaultPaperName Getter
	PaperRect() types.TPaperRect              // property PaperRect Getter
	SetPaperRect(value types.TPaperRect)      // property PaperRect Setter
	SupportedPapers() IStrings                // property SupportedPapers Getter
	PaperRectOf(name string) types.TPaperRect // property PaperRectOf Getter
}

type TPaperSize struct {
	TObject
}

func (m *TPaperSize) DefaultPapers() bool {
	if !m.IsValid() {
		return false
	}
	r := paperSizeAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TPaperSize) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := paperSizeAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TPaperSize) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := paperSizeAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TPaperSize) PaperName() string {
	if !m.IsValid() {
		return ""
	}
	r := paperSizeAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TPaperSize) SetPaperName(value string) {
	if !m.IsValid() {
		return
	}
	paperSizeAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TPaperSize) DefaultPaperName() string {
	if !m.IsValid() {
		return ""
	}
	r := paperSizeAPI().SysCallN(5, m.Instance())
	return api.GoStr(r)
}

func (m *TPaperSize) PaperRect() (result types.TPaperRect) {
	if !m.IsValid() {
		return
	}
	paperSizeAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPaperSize) SetPaperRect(value types.TPaperRect) {
	if !m.IsValid() {
		return
	}
	paperSizeAPI().SysCallN(6, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPaperSize) SupportedPapers() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := paperSizeAPI().SysCallN(7, m.Instance())
	return AsStrings(r)
}

func (m *TPaperSize) PaperRectOf(name string) (result types.TPaperRect) {
	if !m.IsValid() {
		return
	}
	paperSizeAPI().SysCallN(8, m.Instance(), api.PasStr(name), uintptr(base.UnsafePointer(&result)))
	return
}

// NewPaperSize class constructor
func NewPaperSize(owner IPrinter) IPaperSize {
	r := paperSizeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsPaperSize(r)
}

var (
	paperSizeOnce   base.Once
	paperSizeImport *imports.Imports = nil
)

func paperSizeAPI() *imports.Imports {
	paperSizeOnce.Do(func() {
		paperSizeImport = api.NewDefaultImports()
		paperSizeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPaperSize_Create", 0), // constructor NewPaperSize
			/* 1 */ imports.NewTable("TPaperSize_DefaultPapers", 0), // property DefaultPapers
			/* 2 */ imports.NewTable("TPaperSize_Width", 0), // property Width
			/* 3 */ imports.NewTable("TPaperSize_Height", 0), // property Height
			/* 4 */ imports.NewTable("TPaperSize_PaperName", 0), // property PaperName
			/* 5 */ imports.NewTable("TPaperSize_DefaultPaperName", 0), // property DefaultPaperName
			/* 6 */ imports.NewTable("TPaperSize_PaperRect", 0), // property PaperRect
			/* 7 */ imports.NewTable("TPaperSize_SupportedPapers", 0), // property SupportedPapers
			/* 8 */ imports.NewTable("TPaperSize_PaperRectOf", 0), // property PaperRectOf
		}
	})
	return paperSizeImport
}
