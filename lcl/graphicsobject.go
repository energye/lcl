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

// IGraphicsObject Parent: IPersistent
type IGraphicsObject interface {
	IPersistent
	SetOnChanging(fn TNotifyEvent) // property event
	SetOnChange(fn TNotifyEvent)   // property event
}

type TGraphicsObject struct {
	TPersistent
}

func (m *TGraphicsObject) SetOnChanging(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 1, graphicsObjectAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGraphicsObject) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 2, graphicsObjectAPI(), api.MakeEventDataPtr(cb))
}

// NewGraphicsObject class constructor
func NewGraphicsObject() IGraphicsObject {
	r := graphicsObjectAPI().SysCallN(0)
	return AsGraphicsObject(r)
}

var (
	graphicsObjectOnce   base.Once
	graphicsObjectImport *imports.Imports = nil
)

func graphicsObjectAPI() *imports.Imports {
	graphicsObjectOnce.Do(func() {
		graphicsObjectImport = api.NewDefaultImports()
		graphicsObjectImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGraphicsObject_Create", 0), // constructor NewGraphicsObject
			/* 1 */ imports.NewTable("TGraphicsObject_OnChanging", 0), // event OnChanging
			/* 2 */ imports.NewTable("TGraphicsObject_OnChange", 0), // event OnChange
		}
	})
	return graphicsObjectImport
}
