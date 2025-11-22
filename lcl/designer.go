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

// IDesigner Parent: IIDesigner
type IDesigner interface {
	IIDesigner
	SetOnIsDesignMsg(fn TDesignerIsDesignMsgEvent)                         // property event
	SetOnUTF8KeyPress(fn TDesignerUTF8KeyPressEvent)                       // property event
	SetOnModified(fn TDesignerModifiedEvent)                               // property event
	SetOnNotification(fn TDesignerNotificationEvent)                       // property event
	SetOnPaintGrid(fn TDesignerPaintGridEvent)                             // property event
	SetOnValidateRename(fn TDesignerValidateRenameEvent)                   // property event
	SetOnGetShiftState(fn TDesignerGetShiftStateEvent)                     // property event
	SetOnSelectOnlyThisComponent(fn TDesignerSelectOnlyThisComponentEvent) // property event
	SetOnUniqueName(fn TDesignerUniqueNameEvent)                           // property event
	SetOnPrepareFreeDesigner(fn TDesignerPrepareFreeDesignerEvent)         // property event
}

type TDesigner struct {
	TIDesigner
}

func (m *TDesigner) SetOnIsDesignMsg(fn TDesignerIsDesignMsgEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerIsDesignMsgEvent(fn)
	base.SetEvent(m, 1, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnUTF8KeyPress(fn TDesignerUTF8KeyPressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerUTF8KeyPressEvent(fn)
	base.SetEvent(m, 2, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnModified(fn TDesignerModifiedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerModifiedEvent(fn)
	base.SetEvent(m, 3, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnNotification(fn TDesignerNotificationEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerNotificationEvent(fn)
	base.SetEvent(m, 4, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnPaintGrid(fn TDesignerPaintGridEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerPaintGridEvent(fn)
	base.SetEvent(m, 5, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnValidateRename(fn TDesignerValidateRenameEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerValidateRenameEvent(fn)
	base.SetEvent(m, 6, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnGetShiftState(fn TDesignerGetShiftStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerGetShiftStateEvent(fn)
	base.SetEvent(m, 7, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnSelectOnlyThisComponent(fn TDesignerSelectOnlyThisComponentEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerSelectOnlyThisComponentEvent(fn)
	base.SetEvent(m, 8, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnUniqueName(fn TDesignerUniqueNameEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerUniqueNameEvent(fn)
	base.SetEvent(m, 9, designerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDesigner) SetOnPrepareFreeDesigner(fn TDesignerPrepareFreeDesignerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDesignerPrepareFreeDesignerEvent(fn)
	base.SetEvent(m, 10, designerAPI(), api.MakeEventDataPtr(cb))
}

// NewDesigner class constructor
func NewDesigner() IDesigner {
	r := designerAPI().SysCallN(0)
	return AsDesigner(r)
}

var (
	designerOnce   base.Once
	designerImport *imports.Imports = nil
)

func designerAPI() *imports.Imports {
	designerOnce.Do(func() {
		designerImport = api.NewDefaultImports()
		designerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDesigner_Create", 0), // constructor NewDesigner
			/* 1 */ imports.NewTable("TDesigner_OnIsDesignMsg", 0), // event OnIsDesignMsg
			/* 2 */ imports.NewTable("TDesigner_OnUTF8KeyPress", 0), // event OnUTF8KeyPress
			/* 3 */ imports.NewTable("TDesigner_OnModified", 0), // event OnModified
			/* 4 */ imports.NewTable("TDesigner_OnNotification", 0), // event OnNotification
			/* 5 */ imports.NewTable("TDesigner_OnPaintGrid", 0), // event OnPaintGrid
			/* 6 */ imports.NewTable("TDesigner_OnValidateRename", 0), // event OnValidateRename
			/* 7 */ imports.NewTable("TDesigner_OnGetShiftState", 0), // event OnGetShiftState
			/* 8 */ imports.NewTable("TDesigner_OnSelectOnlyThisComponent", 0), // event OnSelectOnlyThisComponent
			/* 9 */ imports.NewTable("TDesigner_OnUniqueName", 0), // event OnUniqueName
			/* 10 */ imports.NewTable("TDesigner_OnPrepareFreeDesigner", 0), // event OnPrepareFreeDesigner
		}
	})
	return designerImport
}
