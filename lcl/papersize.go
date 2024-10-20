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

// IPaperSize Parent: IObject
type IPaperSize interface {
	IObject
	DefaultPapers() bool                                   // property
	Width() int32                                          // property
	Height() int32                                         // property
	PaperName() string                                     // property
	SetPaperName(AValue string)                            // property
	DefaultPaperName() string                              // property
	PaperRect() (resultPaperRect TPaperRect)               // property
	SetPaperRect(AValue *TPaperRect)                       // property
	SupportedPapers() IStrings                             // property
	PaperRectOf(aName string) (resultPaperRect TPaperRect) // property
}

// TPaperSize Parent: TObject
type TPaperSize struct {
	TObject
}

func NewPaperSize(aOwner IPrinter) IPaperSize {
	r1 := paperSizeImportAPI().SysCallN(1, GetObjectUintptr(aOwner))
	return AsPaperSize(r1)
}

func (m *TPaperSize) DefaultPapers() bool {
	r1 := paperSizeImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func (m *TPaperSize) Width() int32 {
	r1 := paperSizeImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TPaperSize) Height() int32 {
	r1 := paperSizeImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TPaperSize) PaperName() string {
	r1 := paperSizeImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPaperSize) SetPaperName(AValue string) {
	paperSizeImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPaperSize) DefaultPaperName() string {
	r1 := paperSizeImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TPaperSize) PaperRect() (resultPaperRect TPaperRect) {
	r1 := paperSizeImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return *(*TPaperRect)(getPointer(r1))
}

func (m *TPaperSize) SetPaperRect(AValue *TPaperRect) {
	paperSizeImportAPI().SysCallN(6, 1, m.Instance(), uintptr(unsafePointer(AValue)))
}

func (m *TPaperSize) SupportedPapers() IStrings {
	r1 := paperSizeImportAPI().SysCallN(8, m.Instance())
	return AsStrings(r1)
}

func (m *TPaperSize) PaperRectOf(aName string) (resultPaperRect TPaperRect) {
	r1 := paperSizeImportAPI().SysCallN(7, m.Instance(), PascalStr(aName))
	return *(*TPaperRect)(getPointer(r1))
}

func PaperSizeClass() TClass {
	ret := paperSizeImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	paperSizeImport       *imports.Imports = nil
	paperSizeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PaperSize_Class", 0),
		/*1*/ imports.NewTable("PaperSize_Create", 0),
		/*2*/ imports.NewTable("PaperSize_DefaultPaperName", 0),
		/*3*/ imports.NewTable("PaperSize_DefaultPapers", 0),
		/*4*/ imports.NewTable("PaperSize_Height", 0),
		/*5*/ imports.NewTable("PaperSize_PaperName", 0),
		/*6*/ imports.NewTable("PaperSize_PaperRect", 0),
		/*7*/ imports.NewTable("PaperSize_PaperRectOf", 0),
		/*8*/ imports.NewTable("PaperSize_SupportedPapers", 0),
		/*9*/ imports.NewTable("PaperSize_Width", 0),
	}
)

func paperSizeImportAPI() *imports.Imports {
	if paperSizeImport == nil {
		paperSizeImport = NewDefaultImports()
		paperSizeImport.SetImportTable(paperSizeImportTables)
		paperSizeImportTables = nil
	}
	return paperSizeImport
}
