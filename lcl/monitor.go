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

// IMonitor Parent: IObject
type IMonitor interface {
	IObject
	Handle() HMONITOR                 // property
	MonitorNum() int32                // property
	Left() int32                      // property
	Height() int32                    // property
	Top() int32                       // property
	Width() int32                     // property
	BoundsRect() (resultRect TRect)   // property
	WorkareaRect() (resultRect TRect) // property
	Primary() bool                    // property
	PixelsPerInch() int32             // property
}

// TMonitor Parent: TObject
type TMonitor struct {
	TObject
}

func NewMonitor() IMonitor {
	r1 := monitorImportAPI().SysCallN(2)
	return AsMonitor(r1)
}

func (m *TMonitor) Handle() HMONITOR {
	r1 := monitorImportAPI().SysCallN(3, m.Instance())
	return HMONITOR(r1)
}

func (m *TMonitor) MonitorNum() int32 {
	r1 := monitorImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Left() int32 {
	r1 := monitorImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Height() int32 {
	r1 := monitorImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Top() int32 {
	r1 := monitorImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Width() int32 {
	r1 := monitorImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TMonitor) BoundsRect() (resultRect TRect) {
	monitorImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TMonitor) WorkareaRect() (resultRect TRect) {
	monitorImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TMonitor) Primary() bool {
	r1 := monitorImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TMonitor) PixelsPerInch() int32 {
	r1 := monitorImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func MonitorClass() TClass {
	ret := monitorImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	monitorImport       *imports.Imports = nil
	monitorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Monitor_BoundsRect", 0),
		/*1*/ imports.NewTable("Monitor_Class", 0),
		/*2*/ imports.NewTable("Monitor_Create", 0),
		/*3*/ imports.NewTable("Monitor_Handle", 0),
		/*4*/ imports.NewTable("Monitor_Height", 0),
		/*5*/ imports.NewTable("Monitor_Left", 0),
		/*6*/ imports.NewTable("Monitor_MonitorNum", 0),
		/*7*/ imports.NewTable("Monitor_PixelsPerInch", 0),
		/*8*/ imports.NewTable("Monitor_Primary", 0),
		/*9*/ imports.NewTable("Monitor_Top", 0),
		/*10*/ imports.NewTable("Monitor_Width", 0),
		/*11*/ imports.NewTable("Monitor_WorkareaRect", 0),
	}
)

func monitorImportAPI() *imports.Imports {
	if monitorImport == nil {
		monitorImport = NewDefaultImports()
		monitorImport.SetImportTable(monitorImportTables)
		monitorImportTables = nil
	}
	return monitorImport
}
