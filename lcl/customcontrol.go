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

// ICustomControl Parent: IWinControl
type ICustomControl interface {
	IWinControl
	Canvas() ICanvas                    // property
	SetCanvas(AValue ICanvas)           // property
	BorderStyle() TBorderStyle          // property
	SetBorderStyle(AValue TBorderStyle) // property
	SetOnPaint(fn TNotifyEvent)         // property event
}

// TCustomControl Parent: TWinControl
type TCustomControl struct {
	TWinControl
	paintPtr uintptr
}

func NewCustomControl(AOwner IComponent) ICustomControl {
	r1 := customControlImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsCustomControl(r1)
}

func (m *TCustomControl) Canvas() ICanvas {
	r1 := customControlImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return AsCanvas(r1)
}

func (m *TCustomControl) SetCanvas(AValue ICanvas) {
	customControlImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomControl) BorderStyle() TBorderStyle {
	r1 := customControlImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomControl) SetBorderStyle(AValue TBorderStyle) {
	customControlImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func CustomControlClass() TClass {
	ret := customControlImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomControl) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	customControlImportAPI().SysCallN(4, m.Instance(), m.paintPtr)
}

var (
	customControlImport       *imports.Imports = nil
	customControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomControl_BorderStyle", 0),
		/*1*/ imports.NewTable("CustomControl_Canvas", 0),
		/*2*/ imports.NewTable("CustomControl_Class", 0),
		/*3*/ imports.NewTable("CustomControl_Create", 0),
		/*4*/ imports.NewTable("CustomControl_SetOnPaint", 0),
	}
)

func customControlImportAPI() *imports.Imports {
	if customControlImport == nil {
		customControlImport = NewDefaultImports()
		customControlImport.SetImportTable(customControlImportTables)
		customControlImportTables = nil
	}
	return customControlImport
}
