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
	"github.com/energye/lcl/types"
)

// IIcon Parent: ICustomIcon
type IIcon interface {
	ICustomIcon
	ReleaseHandle() types.HICON  // function
	Handle() types.HICON         // property Handle Getter
	SetHandle(value types.HICON) // property Handle Setter
}

type TIcon struct {
	TCustomIcon
}

func (m *TIcon) ReleaseHandle() types.HICON {
	if !m.IsValid() {
		return 0
	}
	r := iconAPI().SysCallN(1, m.Instance())
	return types.HICON(r)
}

func (m *TIcon) Handle() types.HICON {
	if !m.IsValid() {
		return 0
	}
	r := iconAPI().SysCallN(2, 0, m.Instance())
	return types.HICON(r)
}

func (m *TIcon) SetHandle(value types.HICON) {
	if !m.IsValid() {
		return
	}
	iconAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

// NewIcon class constructor
func NewIcon() IIcon {
	r := iconAPI().SysCallN(0)
	return AsIcon(r)
}

var (
	iconOnce   base.Once
	iconImport *imports.Imports = nil
)

func iconAPI() *imports.Imports {
	iconOnce.Do(func() {
		iconImport = api.NewDefaultImports()
		iconImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TIcon_Create", 0), // constructor NewIcon
			/* 1 */ imports.NewTable("TIcon_ReleaseHandle", 0), // function ReleaseHandle
			/* 2 */ imports.NewTable("TIcon_Handle", 0), // property Handle
		}
	})
	return iconImport
}
