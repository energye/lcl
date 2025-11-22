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

// ICustomGroupBox Parent: IWinControl
type ICustomGroupBox interface {
	IWinControl
	ParentBackground() bool         // property ParentBackground Getter
	SetParentBackground(value bool) // property ParentBackground Setter
}

type TCustomGroupBox struct {
	TWinControl
}

func (m *TCustomGroupBox) ParentBackground() bool {
	if !m.IsValid() {
		return false
	}
	r := customGroupBoxAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGroupBox) SetParentBackground(value bool) {
	if !m.IsValid() {
		return
	}
	customGroupBoxAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

// NewCustomGroupBox class constructor
func NewCustomGroupBox(owner IComponent) ICustomGroupBox {
	r := customGroupBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomGroupBox(r)
}

func TCustomGroupBoxClass() types.TClass {
	r := customGroupBoxAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	customGroupBoxOnce   base.Once
	customGroupBoxImport *imports.Imports = nil
)

func customGroupBoxAPI() *imports.Imports {
	customGroupBoxOnce.Do(func() {
		customGroupBoxImport = api.NewDefaultImports()
		customGroupBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomGroupBox_Create", 0), // constructor NewCustomGroupBox
			/* 1 */ imports.NewTable("TCustomGroupBox_ParentBackground", 0), // property ParentBackground
			/* 2 */ imports.NewTable("TCustomGroupBox_TClass", 0), // function TCustomGroupBoxClass
		}
	})
	return customGroupBoxImport
}
