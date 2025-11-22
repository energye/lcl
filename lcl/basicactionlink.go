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

// IBasicActionLink Parent: IObject
type IBasicActionLink interface {
	IObject
	Execute(component IComponent) bool // function
	Update() bool                      // function
	Action() IBasicAction              // property Action Getter
	SetAction(value IBasicAction)      // property Action Setter
	SetOnChange(fn TNotifyEvent)       // property event
}

type TBasicActionLink struct {
	TObject
}

func (m *TBasicActionLink) Execute(component IComponent) bool {
	if !m.IsValid() {
		return false
	}
	r := basicActionLinkAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(component))
	return api.GoBool(r)
}

func (m *TBasicActionLink) Update() bool {
	if !m.IsValid() {
		return false
	}
	r := basicActionLinkAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TBasicActionLink) Action() IBasicAction {
	if !m.IsValid() {
		return nil
	}
	r := basicActionLinkAPI().SysCallN(3, 0, m.Instance())
	return AsBasicAction(r)
}

func (m *TBasicActionLink) SetAction(value IBasicAction) {
	if !m.IsValid() {
		return
	}
	basicActionLinkAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TBasicActionLink) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 4, basicActionLinkAPI(), api.MakeEventDataPtr(cb))
}

// NewBasicActionLink class constructor
func NewBasicActionLink(client IObject) IBasicActionLink {
	r := basicActionLinkAPI().SysCallN(0, base.GetObjectUintptr(client))
	return AsBasicActionLink(r)
}

var (
	basicActionLinkOnce   base.Once
	basicActionLinkImport *imports.Imports = nil
)

func basicActionLinkAPI() *imports.Imports {
	basicActionLinkOnce.Do(func() {
		basicActionLinkImport = api.NewDefaultImports()
		basicActionLinkImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBasicActionLink_Create", 0), // constructor NewBasicActionLink
			/* 1 */ imports.NewTable("TBasicActionLink_Execute", 0), // function Execute
			/* 2 */ imports.NewTable("TBasicActionLink_Update", 0), // function Update
			/* 3 */ imports.NewTable("TBasicActionLink_Action", 0), // property Action
			/* 4 */ imports.NewTable("TBasicActionLink_OnChange", 0), // event OnChange
		}
	})
	return basicActionLinkImport
}
