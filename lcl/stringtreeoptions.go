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

// IStringTreeOptions Parent: ICustomStringTreeOptions
type IStringTreeOptions interface {
	ICustomStringTreeOptions
	AnimationOptions() TVTAnimationOptions          // property
	SetAnimationOptions(AValue TVTAnimationOptions) // property
	AutoOptions() TVTAutoOptions                    // property
	SetAutoOptions(AValue TVTAutoOptions)           // property
	ExportMode() TVTExportMode                      // property
	SetExportMode(AValue TVTExportMode)             // property
	MiscOptions() TVTMiscOptions                    // property
	SetMiscOptions(AValue TVTMiscOptions)           // property
	PaintOptions() TVTPaintOptions                  // property
	SetPaintOptions(AValue TVTPaintOptions)         // property
	SelectionOptions() TVTSelectionOptions          // property
	SetSelectionOptions(AValue TVTSelectionOptions) // property
	StringOptions() TVTStringOptions                // property
	SetStringOptions(AValue TVTStringOptions)       // property
}

// TStringTreeOptions Parent: TCustomStringTreeOptions
type TStringTreeOptions struct {
	TCustomStringTreeOptions
}

func NewStringTreeOptions(AOwner IBaseVirtualTree) IStringTreeOptions {
	r1 := stringTreeOptionsImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsStringTreeOptions(r1)
}

func (m *TStringTreeOptions) AnimationOptions() TVTAnimationOptions {
	r1 := stringTreeOptionsImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TVTAnimationOptions(r1)
}

func (m *TStringTreeOptions) SetAnimationOptions(AValue TVTAnimationOptions) {
	stringTreeOptionsImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) AutoOptions() TVTAutoOptions {
	r1 := stringTreeOptionsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TVTAutoOptions(r1)
}

func (m *TStringTreeOptions) SetAutoOptions(AValue TVTAutoOptions) {
	stringTreeOptionsImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) ExportMode() TVTExportMode {
	r1 := stringTreeOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TVTExportMode(r1)
}

func (m *TStringTreeOptions) SetExportMode(AValue TVTExportMode) {
	stringTreeOptionsImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) MiscOptions() TVTMiscOptions {
	r1 := stringTreeOptionsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TVTMiscOptions(r1)
}

func (m *TStringTreeOptions) SetMiscOptions(AValue TVTMiscOptions) {
	stringTreeOptionsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) PaintOptions() TVTPaintOptions {
	r1 := stringTreeOptionsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TVTPaintOptions(r1)
}

func (m *TStringTreeOptions) SetPaintOptions(AValue TVTPaintOptions) {
	stringTreeOptionsImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) SelectionOptions() TVTSelectionOptions {
	r1 := stringTreeOptionsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TVTSelectionOptions(r1)
}

func (m *TStringTreeOptions) SetSelectionOptions(AValue TVTSelectionOptions) {
	stringTreeOptionsImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringTreeOptions) StringOptions() TVTStringOptions {
	r1 := stringTreeOptionsImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TVTStringOptions(r1)
}

func (m *TStringTreeOptions) SetStringOptions(AValue TVTStringOptions) {
	stringTreeOptionsImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func StringTreeOptionsClass() TClass {
	ret := stringTreeOptionsImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	stringTreeOptionsImport       *imports.Imports = nil
	stringTreeOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StringTreeOptions_AnimationOptions", 0),
		/*1*/ imports.NewTable("StringTreeOptions_AutoOptions", 0),
		/*2*/ imports.NewTable("StringTreeOptions_Class", 0),
		/*3*/ imports.NewTable("StringTreeOptions_Create", 0),
		/*4*/ imports.NewTable("StringTreeOptions_ExportMode", 0),
		/*5*/ imports.NewTable("StringTreeOptions_MiscOptions", 0),
		/*6*/ imports.NewTable("StringTreeOptions_PaintOptions", 0),
		/*7*/ imports.NewTable("StringTreeOptions_SelectionOptions", 0),
		/*8*/ imports.NewTable("StringTreeOptions_StringOptions", 0),
	}
)

func stringTreeOptionsImportAPI() *imports.Imports {
	if stringTreeOptionsImport == nil {
		stringTreeOptionsImport = NewDefaultImports()
		stringTreeOptionsImport.SetImportTable(stringTreeOptionsImportTables)
		stringTreeOptionsImportTables = nil
	}
	return stringTreeOptionsImport
}
