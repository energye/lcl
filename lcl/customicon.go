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

// ICustomIcon Parent: IRasterImage
type ICustomIcon interface {
	IRasterImage
	Current() int32                                                                  // property
	SetCurrent(AValue int32)                                                         // property
	Count() int32                                                                    // property
	GetBestIndexForSize(ASize *TSize) int32                                          // function
	Add(AFormat TPixelFormat, AHeight, AWidth Word)                                  // procedure
	AssignImage(ASource IRasterImage)                                                // procedure
	Delete(Aindex int32)                                                             // procedure
	Remove(AFormat TPixelFormat, AHeight, AWidth Word)                               // procedure
	GetDescription(Aindex int32, OutFormat *TPixelFormat, OutHeight, OutWidth *Word) // procedure
	SetSize(AWidth, AHeight int32)                                                   // procedure
	LoadFromResourceHandle(Instance THandle, ResHandle TFPResourceHandle)            // procedure
	Sort()                                                                           // procedure
}

// TCustomIcon Parent: TRasterImage
type TCustomIcon struct {
	TRasterImage
}

func NewCustomIcon() ICustomIcon {
	r1 := customIconImportAPI().SysCallN(4)
	return AsCustomIcon(r1)
}

func (m *TCustomIcon) Current() int32 {
	r1 := customIconImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomIcon) SetCurrent(AValue int32) {
	customIconImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomIcon) Count() int32 {
	r1 := customIconImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func (m *TCustomIcon) GetBestIndexForSize(ASize *TSize) int32 {
	r1 := customIconImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(ASize)))
	return int32(r1)
}

func CustomIconClass() TClass {
	ret := customIconImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomIcon) Add(AFormat TPixelFormat, AHeight, AWidth Word) {
	customIconImportAPI().SysCallN(0, m.Instance(), uintptr(AFormat), uintptr(AHeight), uintptr(AWidth))
}

func (m *TCustomIcon) AssignImage(ASource IRasterImage) {
	customIconImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(ASource))
}

func (m *TCustomIcon) Delete(Aindex int32) {
	customIconImportAPI().SysCallN(6, m.Instance(), uintptr(Aindex))
}

func (m *TCustomIcon) Remove(AFormat TPixelFormat, AHeight, AWidth Word) {
	customIconImportAPI().SysCallN(10, m.Instance(), uintptr(AFormat), uintptr(AHeight), uintptr(AWidth))
}

func (m *TCustomIcon) GetDescription(Aindex int32, OutFormat *TPixelFormat, OutHeight, OutWidth *Word) {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	customIconImportAPI().SysCallN(8, m.Instance(), uintptr(Aindex), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)))
	*OutFormat = TPixelFormat(result1)
	*OutHeight = Word(result2)
	*OutWidth = Word(result3)
}

func (m *TCustomIcon) SetSize(AWidth, AHeight int32) {
	customIconImportAPI().SysCallN(11, m.Instance(), uintptr(AWidth), uintptr(AHeight))
}

func (m *TCustomIcon) LoadFromResourceHandle(Instance THandle, ResHandle TFPResourceHandle) {
	customIconImportAPI().SysCallN(9, m.Instance(), uintptr(Instance), uintptr(ResHandle))
}

func (m *TCustomIcon) Sort() {
	customIconImportAPI().SysCallN(12, m.Instance())
}

var (
	customIconImport       *imports.Imports = nil
	customIconImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomIcon_Add", 0),
		/*1*/ imports.NewTable("CustomIcon_AssignImage", 0),
		/*2*/ imports.NewTable("CustomIcon_Class", 0),
		/*3*/ imports.NewTable("CustomIcon_Count", 0),
		/*4*/ imports.NewTable("CustomIcon_Create", 0),
		/*5*/ imports.NewTable("CustomIcon_Current", 0),
		/*6*/ imports.NewTable("CustomIcon_Delete", 0),
		/*7*/ imports.NewTable("CustomIcon_GetBestIndexForSize", 0),
		/*8*/ imports.NewTable("CustomIcon_GetDescription", 0),
		/*9*/ imports.NewTable("CustomIcon_LoadFromResourceHandle", 0),
		/*10*/ imports.NewTable("CustomIcon_Remove", 0),
		/*11*/ imports.NewTable("CustomIcon_SetSize", 0),
		/*12*/ imports.NewTable("CustomIcon_Sort", 0),
	}
)

func customIconImportAPI() *imports.Imports {
	if customIconImport == nil {
		customIconImport = NewDefaultImports()
		customIconImport.SetImportTable(customIconImportTables)
		customIconImportTables = nil
	}
	return customIconImport
}
