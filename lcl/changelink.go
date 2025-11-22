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

// IChangeLink Parent: IObject
type IChangeLink interface {
	IObject
	Change()                                                       // procedure
	Sender() ICustomImageList                                      // property Sender Getter
	SetSender(value ICustomImageList)                              // property Sender Setter
	SetOnChange(fn TNotifyEvent)                                   // property event
	SetOnDestroyResolutionHandle(fn TDestroyResolutionHandleEvent) // property event
}

type TChangeLink struct {
	TObject
}

func (m *TChangeLink) Change() {
	if !m.IsValid() {
		return
	}
	changeLinkAPI().SysCallN(1, m.Instance())
}

func (m *TChangeLink) Sender() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := changeLinkAPI().SysCallN(2, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TChangeLink) SetSender(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	changeLinkAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TChangeLink) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 3, changeLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TChangeLink) SetOnDestroyResolutionHandle(fn TDestroyResolutionHandleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDestroyResolutionHandleEvent(fn)
	base.SetEvent(m, 4, changeLinkAPI(), api.MakeEventDataPtr(cb))
}

// NewChangeLink class constructor
func NewChangeLink() IChangeLink {
	r := changeLinkAPI().SysCallN(0)
	return AsChangeLink(r)
}

var (
	changeLinkOnce   base.Once
	changeLinkImport *imports.Imports = nil
)

func changeLinkAPI() *imports.Imports {
	changeLinkOnce.Do(func() {
		changeLinkImport = api.NewDefaultImports()
		changeLinkImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TChangeLink_Create", 0), // constructor NewChangeLink
			/* 1 */ imports.NewTable("TChangeLink_Change", 0), // procedure Change
			/* 2 */ imports.NewTable("TChangeLink_Sender", 0), // property Sender
			/* 3 */ imports.NewTable("TChangeLink_OnChange", 0), // event OnChange
			/* 4 */ imports.NewTable("TChangeLink_OnDestroyResolutionHandle", 0), // event OnDestroyResolutionHandle
		}
	})
	return changeLinkImport
}
