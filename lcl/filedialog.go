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

// IFileDialog Parent: ICommonDialog
type IFileDialog interface {
	ICommonDialog
	DoTypeChange()                            // procedure
	IntfFileTypeChanged(newFilterIndex int32) // procedure
	Files() IStrings                          // property Files Getter
	HistoryList() IStrings                    // property HistoryList Getter
	SetHistoryList(value IStrings)            // property HistoryList Setter
	DefaultExt() string                       // property DefaultExt Getter
	SetDefaultExt(value string)               // property DefaultExt Setter
	FileName() string                         // property FileName Getter
	SetFileName(value string)                 // property FileName Setter
	Filter() string                           // property Filter Getter
	SetFilter(value string)                   // property Filter Setter
	FilterIndex() int32                       // property FilterIndex Getter
	SetFilterIndex(value int32)               // property FilterIndex Setter
	InitialDir() string                       // property InitialDir Getter
	SetInitialDir(value string)               // property InitialDir Setter
	SetOnHelpClicked(fn TNotifyEvent)         // property event
	SetOnTypeChange(fn TNotifyEvent)          // property event
}

type TFileDialog struct {
	TCommonDialog
}

func (m *TFileDialog) DoTypeChange() {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(3, m.Instance())
}

func (m *TFileDialog) IntfFileTypeChanged(newFilterIndex int32) {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(4, m.Instance(), uintptr(newFilterIndex))
}

func (m *TFileDialog) Files() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := fileDialogAPI().SysCallN(5, m.Instance())
	return AsStrings(r)
}

func (m *TFileDialog) HistoryList() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := fileDialogAPI().SysCallN(6, 0, m.Instance())
	return AsStrings(r)
}

func (m *TFileDialog) SetHistoryList(value IStrings) {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFileDialog) DefaultExt() string {
	if !m.IsValid() {
		return ""
	}
	r := fileDialogAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileDialog) SetDefaultExt(value string) {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileDialog) FileName() string {
	if !m.IsValid() {
		return ""
	}
	r := fileDialogAPI().SysCallN(8, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileDialog) SetFileName(value string) {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileDialog) Filter() string {
	if !m.IsValid() {
		return ""
	}
	r := fileDialogAPI().SysCallN(9, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileDialog) SetFilter(value string) {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileDialog) FilterIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fileDialogAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TFileDialog) SetFilterIndex(value int32) {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TFileDialog) InitialDir() string {
	if !m.IsValid() {
		return ""
	}
	r := fileDialogAPI().SysCallN(11, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileDialog) SetInitialDir(value string) {
	if !m.IsValid() {
		return
	}
	fileDialogAPI().SysCallN(11, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileDialog) SetOnHelpClicked(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, fileDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFileDialog) SetOnTypeChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, fileDialogAPI(), api.MakeEventDataPtr(cb))
}

// FileDialog  is static instance
var FileDialog _FileDialogClass

// _FileDialogClass is class type defined by TFileDialog
type _FileDialogClass uintptr

func (_FileDialogClass) FindMaskInFilter(filter string, mask string) int32 {
	r := fileDialogAPI().SysCallN(1, api.PasStr(filter), api.PasStr(mask))
	return int32(r)
}

func (_FileDialogClass) ExtractAllFilterMasks(filter string, skipAllFilesMask bool) string {
	r := fileDialogAPI().SysCallN(2, api.PasStr(filter), api.PasBool(skipAllFilesMask))
	return api.GoStr(r)
}

// NewFileDialog class constructor
func NewFileDialog(theOwner IComponent) IFileDialog {
	r := fileDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsFileDialog(r)
}

func TFileDialogClass() types.TClass {
	r := fileDialogAPI().SysCallN(14)
	return types.TClass(r)
}

var (
	fileDialogOnce   base.Once
	fileDialogImport *imports.Imports = nil
)

func fileDialogAPI() *imports.Imports {
	fileDialogOnce.Do(func() {
		fileDialogImport = api.NewDefaultImports()
		fileDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFileDialog_Create", 0), // constructor NewFileDialog
			/* 1 */ imports.NewTable("TFileDialog_FindMaskInFilter", 0), // static function FindMaskInFilter
			/* 2 */ imports.NewTable("TFileDialog_ExtractAllFilterMasks", 0), // static function ExtractAllFilterMasks
			/* 3 */ imports.NewTable("TFileDialog_DoTypeChange", 0), // procedure DoTypeChange
			/* 4 */ imports.NewTable("TFileDialog_IntfFileTypeChanged", 0), // procedure IntfFileTypeChanged
			/* 5 */ imports.NewTable("TFileDialog_Files", 0), // property Files
			/* 6 */ imports.NewTable("TFileDialog_HistoryList", 0), // property HistoryList
			/* 7 */ imports.NewTable("TFileDialog_DefaultExt", 0), // property DefaultExt
			/* 8 */ imports.NewTable("TFileDialog_FileName", 0), // property FileName
			/* 9 */ imports.NewTable("TFileDialog_Filter", 0), // property Filter
			/* 10 */ imports.NewTable("TFileDialog_FilterIndex", 0), // property FilterIndex
			/* 11 */ imports.NewTable("TFileDialog_InitialDir", 0), // property InitialDir
			/* 12 */ imports.NewTable("TFileDialog_OnHelpClicked", 0), // event OnHelpClicked
			/* 13 */ imports.NewTable("TFileDialog_OnTypeChange", 0), // event OnTypeChange
			/* 14 */ imports.NewTable("TFileDialog_TClass", 0), // function TFileDialogClass
		}
	})
	return fileDialogImport
}
