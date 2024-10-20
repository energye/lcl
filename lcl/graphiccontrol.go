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

// IGraphicControl Parent: IControl
type IGraphicControl interface {
	IControl
	Canvas() ICanvas // property
}

// TGraphicControl Parent: TControl
type TGraphicControl struct {
	TControl
}

func NewGraphicControl(AOwner IComponent) IGraphicControl {
	r1 := graphicControlImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsGraphicControl(r1)
}

func (m *TGraphicControl) Canvas() ICanvas {
	r1 := graphicControlImportAPI().SysCallN(0, m.Instance())
	return AsCanvas(r1)
}

func GraphicControlClass() TClass {
	ret := graphicControlImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	graphicControlImport       *imports.Imports = nil
	graphicControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GraphicControl_Canvas", 0),
		/*1*/ imports.NewTable("GraphicControl_Class", 0),
		/*2*/ imports.NewTable("GraphicControl_Create", 0),
	}
)

func graphicControlImportAPI() *imports.Imports {
	if graphicControlImport == nil {
		graphicControlImport = NewDefaultImports()
		graphicControlImport.SetImportTable(graphicControlImportTables)
		graphicControlImportTables = nil
	}
	return graphicControlImport
}
