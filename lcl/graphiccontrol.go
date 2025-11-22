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

// IGraphicControl Parent: IControl
type IGraphicControl interface {
	IControl
	Canvas() ICanvas // property Canvas Getter
}

type TGraphicControl struct {
	TControl
}

func (m *TGraphicControl) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := graphicControlAPI().SysCallN(1, m.Instance())
	return AsCanvas(r)
}

// NewGraphicControl class constructor
func NewGraphicControl(owner IComponent) IGraphicControl {
	r := graphicControlAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsGraphicControl(r)
}

func TGraphicControlClass() types.TClass {
	r := graphicControlAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	graphicControlOnce   base.Once
	graphicControlImport *imports.Imports = nil
)

func graphicControlAPI() *imports.Imports {
	graphicControlOnce.Do(func() {
		graphicControlImport = api.NewDefaultImports()
		graphicControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGraphicControl_Create", 0), // constructor NewGraphicControl
			/* 1 */ imports.NewTable("TGraphicControl_Canvas", 0), // property Canvas
			/* 2 */ imports.NewTable("TGraphicControl_TClass", 0), // function TGraphicControlClass
		}
	})
	return graphicControlImport
}
