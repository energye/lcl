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
)

// IWindowMagnetOptions Parent: IPersistent
type IWindowMagnetOptions interface {
	IPersistent
	AssignTo(dest IPersistent)    // procedure
	SnapToMonitor() bool          // property SnapToMonitor Getter
	SetSnapToMonitor(value bool)  // property SnapToMonitor Setter
	SnapToForms() bool            // property SnapToForms Getter
	SetSnapToForms(value bool)    // property SnapToForms Setter
	SnapFormTarget() bool         // property SnapFormTarget Getter
	SetSnapFormTarget(value bool) // property SnapFormTarget Setter
	Distance() int32              // property Distance Getter
	SetDistance(value int32)      // property Distance Setter
}

type TWindowMagnetOptions struct {
	TPersistent
}

func (m *TWindowMagnetOptions) AssignTo(dest IPersistent) {
	if !m.IsValid() {
		return
	}
	windowMagnetOptionsAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(dest))
}

func (m *TWindowMagnetOptions) SnapToMonitor() bool {
	if !m.IsValid() {
		return false
	}
	r := windowMagnetOptionsAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWindowMagnetOptions) SetSnapToMonitor(value bool) {
	if !m.IsValid() {
		return
	}
	windowMagnetOptionsAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TWindowMagnetOptions) SnapToForms() bool {
	if !m.IsValid() {
		return false
	}
	r := windowMagnetOptionsAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWindowMagnetOptions) SetSnapToForms(value bool) {
	if !m.IsValid() {
		return
	}
	windowMagnetOptionsAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TWindowMagnetOptions) SnapFormTarget() bool {
	if !m.IsValid() {
		return false
	}
	r := windowMagnetOptionsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWindowMagnetOptions) SetSnapFormTarget(value bool) {
	if !m.IsValid() {
		return
	}
	windowMagnetOptionsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TWindowMagnetOptions) Distance() int32 {
	if !m.IsValid() {
		return 0
	}
	r := windowMagnetOptionsAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TWindowMagnetOptions) SetDistance(value int32) {
	if !m.IsValid() {
		return
	}
	windowMagnetOptionsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

// NewWindowMagnetOptions class constructor
func NewWindowMagnetOptions() IWindowMagnetOptions {
	r := windowMagnetOptionsAPI().SysCallN(0)
	return AsWindowMagnetOptions(r)
}

var (
	windowMagnetOptionsOnce   base.Once
	windowMagnetOptionsImport *imports.Imports = nil
)

func windowMagnetOptionsAPI() *imports.Imports {
	windowMagnetOptionsOnce.Do(func() {
		windowMagnetOptionsImport = api.NewDefaultImports()
		windowMagnetOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWindowMagnetOptions_Create", 0), // constructor NewWindowMagnetOptions
			/* 1 */ imports.NewTable("TWindowMagnetOptions_AssignTo", 0), // procedure AssignTo
			/* 2 */ imports.NewTable("TWindowMagnetOptions_SnapToMonitor", 0), // property SnapToMonitor
			/* 3 */ imports.NewTable("TWindowMagnetOptions_SnapToForms", 0), // property SnapToForms
			/* 4 */ imports.NewTable("TWindowMagnetOptions_SnapFormTarget", 0), // property SnapFormTarget
			/* 5 */ imports.NewTable("TWindowMagnetOptions_Distance", 0), // property Distance
		}
	})
	return windowMagnetOptionsImport
}
