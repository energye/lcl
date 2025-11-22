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

// IStringTreeOptions Parent: ICustomStringTreeOptions
type IStringTreeOptions interface {
	ICustomStringTreeOptions
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
	StringOptions() types.TVTStringOptions               // property StringOptions Getter
	SetStringOptions(value types.TVTStringOptions)       // property StringOptions Setter
}

type TStringTreeOptions struct {
	TCustomStringTreeOptions
}

func (m *TStringTreeOptions) AnimationOptions() types.TVTAnimationOptions {
	if !m.IsValid() {
		return 0
	}
	r := stringTreeOptionsAPI().SysCallN(1, 0, m.Instance())
	return types.TVTAnimationOptions(r)
}

func (m *TStringTreeOptions) SetAnimationOptions(value types.TVTAnimationOptions) {
	if !m.IsValid() {
		return
	}
	stringTreeOptionsAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TStringTreeOptions) AutoOptions() types.TVTAutoOptions {
	if !m.IsValid() {
		return 0
	}
	r := stringTreeOptionsAPI().SysCallN(2, 0, m.Instance())
	return types.TVTAutoOptions(r)
}

func (m *TStringTreeOptions) SetAutoOptions(value types.TVTAutoOptions) {
	if !m.IsValid() {
		return
	}
	stringTreeOptionsAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TStringTreeOptions) ExportMode() types.TVTExportMode {
	if !m.IsValid() {
		return 0
	}
	r := stringTreeOptionsAPI().SysCallN(3, 0, m.Instance())
	return types.TVTExportMode(r)
}

func (m *TStringTreeOptions) SetExportMode(value types.TVTExportMode) {
	if !m.IsValid() {
		return
	}
	stringTreeOptionsAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TStringTreeOptions) MiscOptions() types.TVTMiscOptions {
	if !m.IsValid() {
		return 0
	}
	r := stringTreeOptionsAPI().SysCallN(4, 0, m.Instance())
	return types.TVTMiscOptions(r)
}

func (m *TStringTreeOptions) SetMiscOptions(value types.TVTMiscOptions) {
	if !m.IsValid() {
		return
	}
	stringTreeOptionsAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TStringTreeOptions) PaintOptions() types.TVTPaintOptions {
	if !m.IsValid() {
		return 0
	}
	r := stringTreeOptionsAPI().SysCallN(5, 0, m.Instance())
	return types.TVTPaintOptions(r)
}

func (m *TStringTreeOptions) SetPaintOptions(value types.TVTPaintOptions) {
	if !m.IsValid() {
		return
	}
	stringTreeOptionsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TStringTreeOptions) SelectionOptions() types.TVTSelectionOptions {
	if !m.IsValid() {
		return 0
	}
	r := stringTreeOptionsAPI().SysCallN(6, 0, m.Instance())
	return types.TVTSelectionOptions(r)
}

func (m *TStringTreeOptions) SetSelectionOptions(value types.TVTSelectionOptions) {
	if !m.IsValid() {
		return
	}
	stringTreeOptionsAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TStringTreeOptions) StringOptions() types.TVTStringOptions {
	if !m.IsValid() {
		return 0
	}
	r := stringTreeOptionsAPI().SysCallN(7, 0, m.Instance())
	return types.TVTStringOptions(r)
}

func (m *TStringTreeOptions) SetStringOptions(value types.TVTStringOptions) {
	if !m.IsValid() {
		return
	}
	stringTreeOptionsAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

// NewStringTreeOptions class constructor
func NewStringTreeOptions(owner IBaseVirtualTree) IStringTreeOptions {
	r := stringTreeOptionsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsStringTreeOptions(r)
}

var (
	stringTreeOptionsOnce   base.Once
	stringTreeOptionsImport *imports.Imports = nil
)

func stringTreeOptionsAPI() *imports.Imports {
	stringTreeOptionsOnce.Do(func() {
		stringTreeOptionsImport = api.NewDefaultImports()
		stringTreeOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStringTreeOptions_Create", 0), // constructor NewStringTreeOptions
			/* 1 */ imports.NewTable("TStringTreeOptions_AnimationOptions", 0), // property AnimationOptions
			/* 2 */ imports.NewTable("TStringTreeOptions_AutoOptions", 0), // property AutoOptions
			/* 3 */ imports.NewTable("TStringTreeOptions_ExportMode", 0), // property ExportMode
			/* 4 */ imports.NewTable("TStringTreeOptions_MiscOptions", 0), // property MiscOptions
			/* 5 */ imports.NewTable("TStringTreeOptions_PaintOptions", 0), // property PaintOptions
			/* 6 */ imports.NewTable("TStringTreeOptions_SelectionOptions", 0), // property SelectionOptions
			/* 7 */ imports.NewTable("TStringTreeOptions_StringOptions", 0), // property StringOptions
		}
	})
	return stringTreeOptionsImport
}
