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

// IControlActionLink Parent: IActionLink
type IControlActionLink interface {
	IActionLink
}

type TControlActionLink struct {
	TActionLink
}

// NewControlActionLink class constructor
func NewControlActionLink(client IObject) IControlActionLink {
	r := controlActionLinkAPI().SysCallN(0, base.GetObjectUintptr(client))
	return AsControlActionLink(r)
}

var (
	controlActionLinkOnce   base.Once
	controlActionLinkImport *imports.Imports = nil
)

func controlActionLinkAPI() *imports.Imports {
	controlActionLinkOnce.Do(func() {
		controlActionLinkImport = api.NewDefaultImports()
		controlActionLinkImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TControlActionLink_Create", 0), // constructor NewControlActionLink
		}
	})
	return controlActionLinkImport
}
