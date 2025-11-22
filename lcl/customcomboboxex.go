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

// ICustomComboBoxEx Parent: ICustomComboBox
type ICustomComboBoxEx interface {
	ICustomComboBox
	AddToInt() int32                                                                                                    // function
	AddWithStringIntImageIndexX3(caption string, indent int32, imgIdx int32, overlayImgIdx int32, selectedImgIdx int32) // procedure
	AssignItemsExWithStrings(items IStrings)                                                                            // procedure
	AssignItemsExWithComboExItems(itemsEx IComboExItems)                                                                // procedure
	Delete(index int32)                                                                                                 // procedure
	DeleteSelected()                                                                                                    // procedure
	Insert(index int32, caption string, indent int32, imgIdx int32, overlayImgIdx int32, selectedImgIdx int32)          // procedure
	AutoCompleteOptions() types.TAutoCompleteOptions                                                                    // property AutoCompleteOptions Getter
	SetAutoCompleteOptions(value types.TAutoCompleteOptions)                                                            // property AutoCompleteOptions Setter
	Images() ICustomImageList                                                                                           // property Images Getter
	SetImages(value ICustomImageList)                                                                                   // property Images Setter
	ImagesWidth() int32                                                                                                 // property ImagesWidth Getter
	SetImagesWidth(value int32)                                                                                         // property ImagesWidth Setter
	ItemsEx() IComboExItems                                                                                             // property ItemsEx Getter
	SetItemsEx(value IComboExItems)                                                                                     // property ItemsEx Setter
	StyleToComboBoxExStyle() types.TComboBoxExStyle                                                                     // property Style Getter
	SetStyleToComboBoxExStyle(value types.TComboBoxExStyle)                                                             // property Style Setter
	StyleEx() types.TComboBoxExStyles                                                                                   // property StyleEx Getter
	SetStyleEx(value types.TComboBoxExStyles)                                                                           // property StyleEx Setter
}

type TCustomComboBoxEx struct {
	TCustomComboBox
}

func (m *TCustomComboBoxEx) AddToInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxExAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TCustomComboBoxEx) AddWithStringIntImageIndexX3(caption string, indent int32, imgIdx int32, overlayImgIdx int32, selectedImgIdx int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(2, m.Instance(), api.PasStr(caption), uintptr(indent), uintptr(imgIdx), uintptr(overlayImgIdx), uintptr(selectedImgIdx))
}

func (m *TCustomComboBoxEx) AssignItemsExWithStrings(items IStrings) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(items))
}

func (m *TCustomComboBoxEx) AssignItemsExWithComboExItems(itemsEx IComboExItems) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(itemsEx))
}

func (m *TCustomComboBoxEx) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(5, m.Instance(), uintptr(index))
}

func (m *TCustomComboBoxEx) DeleteSelected() {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(6, m.Instance())
}

func (m *TCustomComboBoxEx) Insert(index int32, caption string, indent int32, imgIdx int32, overlayImgIdx int32, selectedImgIdx int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(7, m.Instance(), uintptr(index), api.PasStr(caption), uintptr(indent), uintptr(imgIdx), uintptr(overlayImgIdx), uintptr(selectedImgIdx))
}

func (m *TCustomComboBoxEx) AutoCompleteOptions() types.TAutoCompleteOptions {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxExAPI().SysCallN(8, 0, m.Instance())
	return types.TAutoCompleteOptions(r)
}

func (m *TCustomComboBoxEx) SetAutoCompleteOptions(value types.TAutoCompleteOptions) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBoxEx) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customComboBoxExAPI().SysCallN(9, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomComboBoxEx) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomComboBoxEx) ImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxExAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCustomComboBoxEx) SetImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBoxEx) ItemsEx() IComboExItems {
	if !m.IsValid() {
		return nil
	}
	r := customComboBoxExAPI().SysCallN(11, 0, m.Instance())
	return AsComboExItems(r)
}

func (m *TCustomComboBoxEx) SetItemsEx(value IComboExItems) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomComboBoxEx) StyleToComboBoxExStyle() types.TComboBoxExStyle {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxExAPI().SysCallN(12, 0, m.Instance())
	return types.TComboBoxExStyle(r)
}

func (m *TCustomComboBoxEx) SetStyleToComboBoxExStyle(value types.TComboBoxExStyle) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBoxEx) StyleEx() types.TComboBoxExStyles {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxExAPI().SysCallN(13, 0, m.Instance())
	return types.TComboBoxExStyles(r)
}

func (m *TCustomComboBoxEx) SetStyleEx(value types.TComboBoxExStyles) {
	if !m.IsValid() {
		return
	}
	customComboBoxExAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

// NewCustomComboBoxEx class constructor
func NewCustomComboBoxEx(theOwner IComponent) ICustomComboBoxEx {
	r := customComboBoxExAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomComboBoxEx(r)
}

func TCustomComboBoxExClass() types.TClass {
	r := customComboBoxExAPI().SysCallN(14)
	return types.TClass(r)
}

var (
	customComboBoxExOnce   base.Once
	customComboBoxExImport *imports.Imports = nil
)

func customComboBoxExAPI() *imports.Imports {
	customComboBoxExOnce.Do(func() {
		customComboBoxExImport = api.NewDefaultImports()
		customComboBoxExImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomComboBoxEx_Create", 0), // constructor NewCustomComboBoxEx
			/* 1 */ imports.NewTable("TCustomComboBoxEx_AddToInt", 0), // function AddToInt
			/* 2 */ imports.NewTable("TCustomComboBoxEx_AddWithStringIntImageIndexX3", 0), // procedure AddWithStringIntImageIndexX3
			/* 3 */ imports.NewTable("TCustomComboBoxEx_AssignItemsExWithStrings", 0), // procedure AssignItemsExWithStrings
			/* 4 */ imports.NewTable("TCustomComboBoxEx_AssignItemsExWithComboExItems", 0), // procedure AssignItemsExWithComboExItems
			/* 5 */ imports.NewTable("TCustomComboBoxEx_Delete", 0), // procedure Delete
			/* 6 */ imports.NewTable("TCustomComboBoxEx_DeleteSelected", 0), // procedure DeleteSelected
			/* 7 */ imports.NewTable("TCustomComboBoxEx_Insert", 0), // procedure Insert
			/* 8 */ imports.NewTable("TCustomComboBoxEx_AutoCompleteOptions", 0), // property AutoCompleteOptions
			/* 9 */ imports.NewTable("TCustomComboBoxEx_Images", 0), // property Images
			/* 10 */ imports.NewTable("TCustomComboBoxEx_ImagesWidth", 0), // property ImagesWidth
			/* 11 */ imports.NewTable("TCustomComboBoxEx_ItemsEx", 0), // property ItemsEx
			/* 12 */ imports.NewTable("TCustomComboBoxEx_StyleToComboBoxExStyle", 0), // property StyleToComboBoxExStyle
			/* 13 */ imports.NewTable("TCustomComboBoxEx_StyleEx", 0), // property StyleEx
			/* 14 */ imports.NewTable("TCustomComboBoxEx_TClass", 0), // function TCustomComboBoxExClass
		}
	})
	return customComboBoxExImport
}
