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

// INotifyEventDelegate Parent: IObject
type INotifyEventDelegate interface {
	IObject
	SetOnNotify(fn TNotifyEvent) // property event
}

type TNotifyEventDelegate struct {
	TObject
}

func (m *TNotifyEventDelegate) SetOnNotify(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 0, notifyEventDelegateAPI(), api.MakeEventDataPtr(cb))
}

var (
	notifyEventDelegateOnce   base.Once
	notifyEventDelegateImport *imports.Imports = nil
)

func notifyEventDelegateAPI() *imports.Imports {
	notifyEventDelegateOnce.Do(func() {
		notifyEventDelegateImport = api.NewDefaultImports()
		notifyEventDelegateImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TNotifyEventDelegate_OnNotify", 0), // event OnNotify
		}
	})
	return notifyEventDelegateImport
}
