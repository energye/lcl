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

// IGraphicsObject Parent: IPersistent
type IGraphicsObject interface {
	IPersistent
	SetOnChanging(fn TNotifyEvent) // property event
	SetOnChange(fn TNotifyEvent)   // property event
}

// TGraphicsObject Parent: TPersistent
type TGraphicsObject struct {
	TPersistent
	changingPtr uintptr
	changePtr   uintptr
}

func NewGraphicsObject() IGraphicsObject {
	r1 := graphicsObjectImportAPI().SysCallN(1)
	return AsGraphicsObject(r1)
}

func GraphicsObjectClass() TClass {
	ret := graphicsObjectImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TGraphicsObject) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	graphicsObjectImportAPI().SysCallN(3, m.Instance(), m.changingPtr)
}

func (m *TGraphicsObject) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	graphicsObjectImportAPI().SysCallN(2, m.Instance(), m.changePtr)
}

var (
	graphicsObjectImport       *imports.Imports = nil
	graphicsObjectImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GraphicsObject_Class", 0),
		/*1*/ imports.NewTable("GraphicsObject_Create", 0),
		/*2*/ imports.NewTable("GraphicsObject_SetOnChange", 0),
		/*3*/ imports.NewTable("GraphicsObject_SetOnChanging", 0),
	}
)

func graphicsObjectImportAPI() *imports.Imports {
	if graphicsObjectImport == nil {
		graphicsObjectImport = NewDefaultImports()
		graphicsObjectImport.SetImportTable(graphicsObjectImportTables)
		graphicsObjectImportTables = nil
	}
	return graphicsObjectImport
}
