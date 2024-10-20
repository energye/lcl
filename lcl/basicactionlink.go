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

// IBasicActionLink Parent: IObject
type IBasicActionLink interface {
	IObject
	Action() IBasicAction               // property
	SetAction(AValue IBasicAction)      // property
	Execute(AComponent IComponent) bool // function
	Update() bool                       // function
	SetOnChange(fn TNotifyEvent)        // property event
}

// TBasicActionLink Parent: TObject
type TBasicActionLink struct {
	TObject
	changePtr uintptr
}

func NewBasicActionLink(AClient IObject) IBasicActionLink {
	r1 := basicActionLinkImportAPI().SysCallN(2, GetObjectUintptr(AClient))
	return AsBasicActionLink(r1)
}

func (m *TBasicActionLink) Action() IBasicAction {
	r1 := basicActionLinkImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsBasicAction(r1)
}

func (m *TBasicActionLink) SetAction(AValue IBasicAction) {
	basicActionLinkImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBasicActionLink) Execute(AComponent IComponent) bool {
	r1 := basicActionLinkImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(AComponent))
	return GoBool(r1)
}

func (m *TBasicActionLink) Update() bool {
	r1 := basicActionLinkImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func BasicActionLinkClass() TClass {
	ret := basicActionLinkImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TBasicActionLink) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	basicActionLinkImportAPI().SysCallN(4, m.Instance(), m.changePtr)
}

var (
	basicActionLinkImport       *imports.Imports = nil
	basicActionLinkImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("BasicActionLink_Action", 0),
		/*1*/ imports.NewTable("BasicActionLink_Class", 0),
		/*2*/ imports.NewTable("BasicActionLink_Create", 0),
		/*3*/ imports.NewTable("BasicActionLink_Execute", 0),
		/*4*/ imports.NewTable("BasicActionLink_SetOnChange", 0),
		/*5*/ imports.NewTable("BasicActionLink_Update", 0),
	}
)

func basicActionLinkImportAPI() *imports.Imports {
	if basicActionLinkImport == nil {
		basicActionLinkImport = NewDefaultImports()
		basicActionLinkImport.SetImportTable(basicActionLinkImportTables)
		basicActionLinkImportTables = nil
	}
	return basicActionLinkImport
}
