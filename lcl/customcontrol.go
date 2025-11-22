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

// ICustomControl Parent: IWinControl
type ICustomControl interface {
	IWinControl
	Canvas() ICanvas         // property Canvas Getter
	SetCanvas(value ICanvas) // property Canvas Setter
	// BorderStyleToBorderStyle
	//  properties which are not supported by all descendents
	BorderStyleToBorderStyle() types.TBorderStyle         // property BorderStyle Getter
	SetBorderStyleToBorderStyle(value types.TBorderStyle) // property BorderStyle Setter
	SetOnPaint(fn TNotifyEvent)                           // property event
}

type TCustomControl struct {
	TWinControl
}

func (m *TCustomControl) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := customControlAPI().SysCallN(1, 0, m.Instance())
	return AsCanvas(r)
}

func (m *TCustomControl) SetCanvas(value ICanvas) {
	if !m.IsValid() {
		return
	}
	customControlAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomControl) BorderStyleToBorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customControlAPI().SysCallN(2, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCustomControl) SetBorderStyleToBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	customControlAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControl) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 3, customControlAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomControl class constructor
func NewCustomControl(owner IComponent) ICustomControl {
	r := customControlAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomControl(r)
}

func TCustomControlClass() types.TClass {
	r := customControlAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	customControlOnce   base.Once
	customControlImport *imports.Imports = nil
)

func customControlAPI() *imports.Imports {
	customControlOnce.Do(func() {
		customControlImport = api.NewDefaultImports()
		customControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomControl_Create", 0), // constructor NewCustomControl
			/* 1 */ imports.NewTable("TCustomControl_Canvas", 0), // property Canvas
			/* 2 */ imports.NewTable("TCustomControl_BorderStyleToBorderStyle", 0), // property BorderStyleToBorderStyle
			/* 3 */ imports.NewTable("TCustomControl_OnPaint", 0), // event OnPaint
			/* 4 */ imports.NewTable("TCustomControl_TClass", 0), // function TCustomControlClass
		}
	})
	return customControlImport
}
