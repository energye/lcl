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

// IChangeLink Parent: IObject
type IChangeLink interface {
	IObject
	Sender() ICustomImageList                                      // property
	SetSender(AValue ICustomImageList)                             // property
	Change()                                                       // procedure
	SetOnChange(fn TNotifyEvent)                                   // property event
	SetOnDestroyResolutionHandle(fn TDestroyResolutionHandleEvent) // property event
}

// TChangeLink Parent: TObject
type TChangeLink struct {
	TObject
	changePtr                  uintptr
	destroyResolutionHandlePtr uintptr
}

func NewChangeLink() IChangeLink {
	r1 := changeLinkImportAPI().SysCallN(2)
	return AsChangeLink(r1)
}

func (m *TChangeLink) Sender() ICustomImageList {
	r1 := changeLinkImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TChangeLink) SetSender(AValue ICustomImageList) {
	changeLinkImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func ChangeLinkClass() TClass {
	ret := changeLinkImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TChangeLink) Change() {
	changeLinkImportAPI().SysCallN(0, m.Instance())
}

func (m *TChangeLink) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	changeLinkImportAPI().SysCallN(4, m.Instance(), m.changePtr)
}

func (m *TChangeLink) SetOnDestroyResolutionHandle(fn TDestroyResolutionHandleEvent) {
	if m.destroyResolutionHandlePtr != 0 {
		RemoveEventElement(m.destroyResolutionHandlePtr)
	}
	m.destroyResolutionHandlePtr = MakeEventDataPtr(fn)
	changeLinkImportAPI().SysCallN(5, m.Instance(), m.destroyResolutionHandlePtr)
}

var (
	changeLinkImport       *imports.Imports = nil
	changeLinkImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ChangeLink_Change", 0),
		/*1*/ imports.NewTable("ChangeLink_Class", 0),
		/*2*/ imports.NewTable("ChangeLink_Create", 0),
		/*3*/ imports.NewTable("ChangeLink_Sender", 0),
		/*4*/ imports.NewTable("ChangeLink_SetOnChange", 0),
		/*5*/ imports.NewTable("ChangeLink_SetOnDestroyResolutionHandle", 0),
	}
)

func changeLinkImportAPI() *imports.Imports {
	if changeLinkImport == nil {
		changeLinkImport = NewDefaultImports()
		changeLinkImport.SetImportTable(changeLinkImportTables)
		changeLinkImportTables = nil
	}
	return changeLinkImport
}
