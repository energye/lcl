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

// IVirtualTreeOptions Parent: ICustomVirtualTreeOptions
type IVirtualTreeOptions interface {
	ICustomVirtualTreeOptions
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
}

// TVirtualTreeOptions Parent: TCustomVirtualTreeOptions
type TVirtualTreeOptions struct {
	TCustomVirtualTreeOptions
}

func NewVirtualTreeOptions(AOwner IBaseVirtualTree) IVirtualTreeOptions {
	r1 := virtualTreeOptionsImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsVirtualTreeOptions(r1)
}

func (m *TVirtualTreeOptions) AnimationOptions() TVTAnimationOptions {
	r1 := virtualTreeOptionsImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TVTAnimationOptions(r1)
}

func (m *TVirtualTreeOptions) SetAnimationOptions(AValue TVTAnimationOptions) {
	virtualTreeOptionsImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) AutoOptions() TVTAutoOptions {
	r1 := virtualTreeOptionsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TVTAutoOptions(r1)
}

func (m *TVirtualTreeOptions) SetAutoOptions(AValue TVTAutoOptions) {
	virtualTreeOptionsImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) ExportMode() TVTExportMode {
	r1 := virtualTreeOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TVTExportMode(r1)
}

func (m *TVirtualTreeOptions) SetExportMode(AValue TVTExportMode) {
	virtualTreeOptionsImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) MiscOptions() TVTMiscOptions {
	r1 := virtualTreeOptionsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TVTMiscOptions(r1)
}

func (m *TVirtualTreeOptions) SetMiscOptions(AValue TVTMiscOptions) {
	virtualTreeOptionsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) PaintOptions() TVTPaintOptions {
	r1 := virtualTreeOptionsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TVTPaintOptions(r1)
}

func (m *TVirtualTreeOptions) SetPaintOptions(AValue TVTPaintOptions) {
	virtualTreeOptionsImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) SelectionOptions() TVTSelectionOptions {
	r1 := virtualTreeOptionsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TVTSelectionOptions(r1)
}

func (m *TVirtualTreeOptions) SetSelectionOptions(AValue TVTSelectionOptions) {
	virtualTreeOptionsImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func VirtualTreeOptionsClass() TClass {
	ret := virtualTreeOptionsImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	virtualTreeOptionsImport       *imports.Imports = nil
	virtualTreeOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VirtualTreeOptions_AnimationOptions", 0),
		/*1*/ imports.NewTable("VirtualTreeOptions_AutoOptions", 0),
		/*2*/ imports.NewTable("VirtualTreeOptions_Class", 0),
		/*3*/ imports.NewTable("VirtualTreeOptions_Create", 0),
		/*4*/ imports.NewTable("VirtualTreeOptions_ExportMode", 0),
		/*5*/ imports.NewTable("VirtualTreeOptions_MiscOptions", 0),
		/*6*/ imports.NewTable("VirtualTreeOptions_PaintOptions", 0),
		/*7*/ imports.NewTable("VirtualTreeOptions_SelectionOptions", 0),
	}
)

func virtualTreeOptionsImportAPI() *imports.Imports {
	if virtualTreeOptionsImport == nil {
		virtualTreeOptionsImport = NewDefaultImports()
		virtualTreeOptionsImport.SetImportTable(virtualTreeOptionsImportTables)
		virtualTreeOptionsImportTables = nil
	}
	return virtualTreeOptionsImport
}
