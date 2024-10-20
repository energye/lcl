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

// ICustomComboBoxEx Parent: ICustomComboBox
type ICustomComboBoxEx interface {
	ICustomComboBox
	AutoCompleteOptions() TAutoCompleteOptions                                                                                         // property
	SetAutoCompleteOptions(AValue TAutoCompleteOptions)                                                                                // property
	Images() ICustomImageList                                                                                                          // property
	SetImages(AValue ICustomImageList)                                                                                                 // property
	ImagesWidth() int32                                                                                                                // property
	SetImagesWidth(AValue int32)                                                                                                       // property
	ItemsEx() IComboExItems                                                                                                            // property
	SetItemsEx(AValue IComboExItems)                                                                                                   // property
	StyleForComboBoxExStyle() TComboBoxExStyle                                                                                         // property
	SetStyleForComboBoxExStyle(AValue TComboBoxExStyle)                                                                                // property
	StyleEx() TComboBoxExStyles                                                                                                        // property
	SetStyleEx(AValue TComboBoxExStyles)                                                                                               // property
	Add() int32                                                                                                                        // function
	Add1(ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex)                 // procedure
	AssignItemsEx(AItems IStrings)                                                                                                     // procedure
	AssignItemsEx1(AItemsEx IComboExItems)                                                                                             // procedure
	Delete(AIndex int32)                                                                                                               // procedure
	DeleteSelected()                                                                                                                   // procedure
	Insert(AIndex int32, ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex) // procedure
}

// TCustomComboBoxEx Parent: TCustomComboBox
type TCustomComboBoxEx struct {
	TCustomComboBox
}

func NewCustomComboBoxEx(TheOwner IComponent) ICustomComboBoxEx {
	r1 := customComboBoxExImportAPI().SysCallN(6, GetObjectUintptr(TheOwner))
	return AsCustomComboBoxEx(r1)
}

func (m *TCustomComboBoxEx) AutoCompleteOptions() TAutoCompleteOptions {
	r1 := customComboBoxExImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TAutoCompleteOptions(r1)
}

func (m *TCustomComboBoxEx) SetAutoCompleteOptions(AValue TAutoCompleteOptions) {
	customComboBoxExImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) Images() ICustomImageList {
	r1 := customComboBoxExImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomComboBoxEx) SetImages(AValue ICustomImageList) {
	customComboBoxExImportAPI().SysCallN(9, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomComboBoxEx) ImagesWidth() int32 {
	r1 := customComboBoxExImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBoxEx) SetImagesWidth(AValue int32) {
	customComboBoxExImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) ItemsEx() IComboExItems {
	r1 := customComboBoxExImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return AsComboExItems(r1)
}

func (m *TCustomComboBoxEx) SetItemsEx(AValue IComboExItems) {
	customComboBoxExImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomComboBoxEx) StyleForComboBoxExStyle() TComboBoxExStyle {
	r1 := customComboBoxExImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TComboBoxExStyle(r1)
}

func (m *TCustomComboBoxEx) SetStyleForComboBoxExStyle(AValue TComboBoxExStyle) {
	customComboBoxExImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) StyleEx() TComboBoxExStyles {
	r1 := customComboBoxExImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TComboBoxExStyles(r1)
}

func (m *TCustomComboBoxEx) SetStyleEx(AValue TComboBoxExStyles) {
	customComboBoxExImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBoxEx) Add() int32 {
	r1 := customComboBoxExImportAPI().SysCallN(0, m.Instance())
	return int32(r1)
}

func CustomComboBoxExClass() TClass {
	ret := customComboBoxExImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TCustomComboBoxEx) Add1(ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex) {
	customComboBoxExImportAPI().SysCallN(1, m.Instance(), PascalStr(ACaption), uintptr(AIndent), uintptr(AImgIdx), uintptr(AOverlayImgIdx), uintptr(ASelectedImgIdx))
}

func (m *TCustomComboBoxEx) AssignItemsEx(AItems IStrings) {
	customComboBoxExImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(AItems))
}

func (m *TCustomComboBoxEx) AssignItemsEx1(AItemsEx IComboExItems) {
	customComboBoxExImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(AItemsEx))
}

func (m *TCustomComboBoxEx) Delete(AIndex int32) {
	customComboBoxExImportAPI().SysCallN(7, m.Instance(), uintptr(AIndex))
}

func (m *TCustomComboBoxEx) DeleteSelected() {
	customComboBoxExImportAPI().SysCallN(8, m.Instance())
}

func (m *TCustomComboBoxEx) Insert(AIndex int32, ACaption string, AIndent int32, AImgIdx TImageIndex, AOverlayImgIdx TImageIndex, ASelectedImgIdx TImageIndex) {
	customComboBoxExImportAPI().SysCallN(11, m.Instance(), uintptr(AIndex), PascalStr(ACaption), uintptr(AIndent), uintptr(AImgIdx), uintptr(AOverlayImgIdx), uintptr(ASelectedImgIdx))
}

var (
	customComboBoxExImport       *imports.Imports = nil
	customComboBoxExImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomComboBoxEx_Add", 0),
		/*1*/ imports.NewTable("CustomComboBoxEx_Add1", 0),
		/*2*/ imports.NewTable("CustomComboBoxEx_AssignItemsEx", 0),
		/*3*/ imports.NewTable("CustomComboBoxEx_AssignItemsEx1", 0),
		/*4*/ imports.NewTable("CustomComboBoxEx_AutoCompleteOptions", 0),
		/*5*/ imports.NewTable("CustomComboBoxEx_Class", 0),
		/*6*/ imports.NewTable("CustomComboBoxEx_Create", 0),
		/*7*/ imports.NewTable("CustomComboBoxEx_Delete", 0),
		/*8*/ imports.NewTable("CustomComboBoxEx_DeleteSelected", 0),
		/*9*/ imports.NewTable("CustomComboBoxEx_Images", 0),
		/*10*/ imports.NewTable("CustomComboBoxEx_ImagesWidth", 0),
		/*11*/ imports.NewTable("CustomComboBoxEx_Insert", 0),
		/*12*/ imports.NewTable("CustomComboBoxEx_ItemsEx", 0),
		/*13*/ imports.NewTable("CustomComboBoxEx_StyleEx", 0),
		/*14*/ imports.NewTable("CustomComboBoxEx_StyleForComboBoxExStyle", 0),
	}
)

func customComboBoxExImportAPI() *imports.Imports {
	if customComboBoxExImport == nil {
		customComboBoxExImport = NewDefaultImports()
		customComboBoxExImport.SetImportTable(customComboBoxExImportTables)
		customComboBoxExImportTables = nil
	}
	return customComboBoxExImport
}
