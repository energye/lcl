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
	r1 := LCL().SysCallN(6043, GetObjectUintptr(AOwner))
	return AsVirtualTreeOptions(r1)
}

func (m *TVirtualTreeOptions) AnimationOptions() TVTAnimationOptions {
	r1 := LCL().SysCallN(6040, 0, m.Instance(), 0)
	return TVTAnimationOptions(r1)
}

func (m *TVirtualTreeOptions) SetAnimationOptions(AValue TVTAnimationOptions) {
	LCL().SysCallN(6040, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) AutoOptions() TVTAutoOptions {
	r1 := LCL().SysCallN(6041, 0, m.Instance(), 0)
	return TVTAutoOptions(r1)
}

func (m *TVirtualTreeOptions) SetAutoOptions(AValue TVTAutoOptions) {
	LCL().SysCallN(6041, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) ExportMode() TVTExportMode {
	r1 := LCL().SysCallN(6044, 0, m.Instance(), 0)
	return TVTExportMode(r1)
}

func (m *TVirtualTreeOptions) SetExportMode(AValue TVTExportMode) {
	LCL().SysCallN(6044, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) MiscOptions() TVTMiscOptions {
	r1 := LCL().SysCallN(6045, 0, m.Instance(), 0)
	return TVTMiscOptions(r1)
}

func (m *TVirtualTreeOptions) SetMiscOptions(AValue TVTMiscOptions) {
	LCL().SysCallN(6045, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) PaintOptions() TVTPaintOptions {
	r1 := LCL().SysCallN(6046, 0, m.Instance(), 0)
	return TVTPaintOptions(r1)
}

func (m *TVirtualTreeOptions) SetPaintOptions(AValue TVTPaintOptions) {
	LCL().SysCallN(6046, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeOptions) SelectionOptions() TVTSelectionOptions {
	r1 := LCL().SysCallN(6047, 0, m.Instance(), 0)
	return TVTSelectionOptions(r1)
}

func (m *TVirtualTreeOptions) SetSelectionOptions(AValue TVTSelectionOptions) {
	LCL().SysCallN(6047, 1, m.Instance(), uintptr(AValue))
}

func VirtualTreeOptionsClass() TClass {
	ret := LCL().SysCallN(6042)
	return TClass(ret)
}
