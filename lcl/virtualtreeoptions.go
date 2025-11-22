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

// IVirtualTreeOptions Parent: ICustomVirtualTreeOptions
type IVirtualTreeOptions interface {
	ICustomVirtualTreeOptions
	AnimationOptions() types.TVTAnimationOptions         // property AnimationOptions Getter
	SetAnimationOptions(value types.TVTAnimationOptions) // property AnimationOptions Setter
	AutoOptions() types.TVTAutoOptions                   // property AutoOptions Getter
	SetAutoOptions(value types.TVTAutoOptions)           // property AutoOptions Setter
	ExportMode() types.TVTExportMode                     // property ExportMode Getter
	SetExportMode(value types.TVTExportMode)             // property ExportMode Setter
	MiscOptions() types.TVTMiscOptions                   // property MiscOptions Getter
	SetMiscOptions(value types.TVTMiscOptions)           // property MiscOptions Setter
	PaintOptions() types.TVTPaintOptions                 // property PaintOptions Getter
	SetPaintOptions(value types.TVTPaintOptions)         // property PaintOptions Setter
	SelectionOptions() types.TVTSelectionOptions         // property SelectionOptions Getter
	SetSelectionOptions(value types.TVTSelectionOptions) // property SelectionOptions Setter
}

type TVirtualTreeOptions struct {
	TCustomVirtualTreeOptions
}

func (m *TVirtualTreeOptions) AnimationOptions() types.TVTAnimationOptions {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeOptionsAPI().SysCallN(1, 0, m.Instance())
	return types.TVTAnimationOptions(r)
}

func (m *TVirtualTreeOptions) SetAnimationOptions(value types.TVTAnimationOptions) {
	if !m.IsValid() {
		return
	}
	virtualTreeOptionsAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeOptions) AutoOptions() types.TVTAutoOptions {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeOptionsAPI().SysCallN(2, 0, m.Instance())
	return types.TVTAutoOptions(r)
}

func (m *TVirtualTreeOptions) SetAutoOptions(value types.TVTAutoOptions) {
	if !m.IsValid() {
		return
	}
	virtualTreeOptionsAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeOptions) ExportMode() types.TVTExportMode {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeOptionsAPI().SysCallN(3, 0, m.Instance())
	return types.TVTExportMode(r)
}

func (m *TVirtualTreeOptions) SetExportMode(value types.TVTExportMode) {
	if !m.IsValid() {
		return
	}
	virtualTreeOptionsAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeOptions) MiscOptions() types.TVTMiscOptions {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeOptionsAPI().SysCallN(4, 0, m.Instance())
	return types.TVTMiscOptions(r)
}

func (m *TVirtualTreeOptions) SetMiscOptions(value types.TVTMiscOptions) {
	if !m.IsValid() {
		return
	}
	virtualTreeOptionsAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeOptions) PaintOptions() types.TVTPaintOptions {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeOptionsAPI().SysCallN(5, 0, m.Instance())
	return types.TVTPaintOptions(r)
}

func (m *TVirtualTreeOptions) SetPaintOptions(value types.TVTPaintOptions) {
	if !m.IsValid() {
		return
	}
	virtualTreeOptionsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeOptions) SelectionOptions() types.TVTSelectionOptions {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeOptionsAPI().SysCallN(6, 0, m.Instance())
	return types.TVTSelectionOptions(r)
}

func (m *TVirtualTreeOptions) SetSelectionOptions(value types.TVTSelectionOptions) {
	if !m.IsValid() {
		return
	}
	virtualTreeOptionsAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

// NewVirtualTreeOptions class constructor
func NewVirtualTreeOptions(owner IBaseVirtualTree) IVirtualTreeOptions {
	r := virtualTreeOptionsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsVirtualTreeOptions(r)
}

var (
	virtualTreeOptionsOnce   base.Once
	virtualTreeOptionsImport *imports.Imports = nil
)

func virtualTreeOptionsAPI() *imports.Imports {
	virtualTreeOptionsOnce.Do(func() {
		virtualTreeOptionsImport = api.NewDefaultImports()
		virtualTreeOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVirtualTreeOptions_Create", 0), // constructor NewVirtualTreeOptions
			/* 1 */ imports.NewTable("TVirtualTreeOptions_AnimationOptions", 0), // property AnimationOptions
			/* 2 */ imports.NewTable("TVirtualTreeOptions_AutoOptions", 0), // property AutoOptions
			/* 3 */ imports.NewTable("TVirtualTreeOptions_ExportMode", 0), // property ExportMode
			/* 4 */ imports.NewTable("TVirtualTreeOptions_MiscOptions", 0), // property MiscOptions
			/* 5 */ imports.NewTable("TVirtualTreeOptions_PaintOptions", 0), // property PaintOptions
			/* 6 */ imports.NewTable("TVirtualTreeOptions_SelectionOptions", 0), // property SelectionOptions
		}
	})
	return virtualTreeOptionsImport
}
