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

// ICustomVTEditLink Parent: IInterfacedObject
type ICustomVTEditLink interface {
	IInterfacedObject
	SetOnBeginEdit(fn TVTBeginCancelEndEditEvent)  // property event
	SetOnCancelEdit(fn TVTBeginCancelEndEditEvent) // property event
	SetOnEndEdit(fn TVTBeginCancelEndEditEvent)    // property event
	SetOnPrepareEdit(fn TVTPrepareEditEvent)       // property event
	SetOnGetBounds(fn TVTGetBoundsEvent)           // property event
	SetOnProcessMessage(fn TVTProcessMessageEvent) // property event
	SetOnSetBounds(fn TVTSetBoundsEvent)           // property event
	SetOnDestroy(fn TNotifyEvent)                  // property event
	AsIntfVTEditLink() uintptr
}

type TCustomVTEditLink struct {
	TInterfacedObject
}

func (m *TCustomVTEditLink) SetOnBeginEdit(fn TVTBeginCancelEndEditEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeginCancelEndEditEvent(fn)
	base.SetEvent(m, 1, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) SetOnCancelEdit(fn TVTBeginCancelEndEditEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeginCancelEndEditEvent(fn)
	base.SetEvent(m, 2, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) SetOnEndEdit(fn TVTBeginCancelEndEditEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeginCancelEndEditEvent(fn)
	base.SetEvent(m, 3, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) SetOnPrepareEdit(fn TVTPrepareEditEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPrepareEditEvent(fn)
	base.SetEvent(m, 4, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) SetOnGetBounds(fn TVTGetBoundsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetBoundsEvent(fn)
	base.SetEvent(m, 5, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) SetOnProcessMessage(fn TVTProcessMessageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTProcessMessageEvent(fn)
	base.SetEvent(m, 6, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) SetOnSetBounds(fn TVTSetBoundsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSetBoundsEvent(fn)
	base.SetEvent(m, 7, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) SetOnDestroy(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, customVTEditLinkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomVTEditLink) AsIntfVTEditLink() uintptr {
	return m.GetIntfPointer(0)
}

// NewCustomVTEditLink class constructor
func NewCustomVTEditLink() ICustomVTEditLink {
	var vTEditLinkPtr uintptr // IVTEditLink
	r := customVTEditLinkAPI().SysCallN(0, uintptr(base.UnsafePointer(&vTEditLinkPtr)))
	ret := AsCustomVTEditLink(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, vTEditLinkPtr)
	}
	return ret
}

var (
	customVTEditLinkOnce   base.Once
	customVTEditLinkImport *imports.Imports = nil
)

func customVTEditLinkAPI() *imports.Imports {
	customVTEditLinkOnce.Do(func() {
		customVTEditLinkImport = api.NewDefaultImports()
		customVTEditLinkImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomVTEditLink_Create", 0), // constructor NewCustomVTEditLink
			/* 1 */ imports.NewTable("TCustomVTEditLink_OnBeginEdit", 0), // event OnBeginEdit
			/* 2 */ imports.NewTable("TCustomVTEditLink_OnCancelEdit", 0), // event OnCancelEdit
			/* 3 */ imports.NewTable("TCustomVTEditLink_OnEndEdit", 0), // event OnEndEdit
			/* 4 */ imports.NewTable("TCustomVTEditLink_OnPrepareEdit", 0), // event OnPrepareEdit
			/* 5 */ imports.NewTable("TCustomVTEditLink_OnGetBounds", 0), // event OnGetBounds
			/* 6 */ imports.NewTable("TCustomVTEditLink_OnProcessMessage", 0), // event OnProcessMessage
			/* 7 */ imports.NewTable("TCustomVTEditLink_OnSetBounds", 0), // event OnSetBounds
			/* 8 */ imports.NewTable("TCustomVTEditLink_OnDestroy", 0), // event OnDestroy
		}
	})
	return customVTEditLinkImport
}
