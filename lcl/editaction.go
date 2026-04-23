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

// IEditAction Parent: IAction
type IEditAction interface {
	IAction
	// Control
	//  limits target to the specific control
	Control() ICustomEdit         // property Control Getter
	SetControl(value ICustomEdit) // property Control Setter
}

type TEditAction struct {
	TAction
}

func (m *TEditAction) Control() ICustomEdit {
	if !m.IsValid() {
		return nil
	}
	r := editActionAPI().SysCallN(1, 0, m.Instance())
	return AsCustomEdit(r)
}

func (m *TEditAction) SetControl(value ICustomEdit) {
	if !m.IsValid() {
		return
	}
	editActionAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewEditAction class constructor
func NewEditAction(owner IComponent) IEditAction {
	r := editActionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditAction(r)
}

func TEditActionClass() types.TClass {
	r := editActionAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	editActionOnce   base.Once
	editActionImport *imports.Imports = nil
)

func editActionAPI() *imports.Imports {
	editActionOnce.Do(func() {
		editActionImport = api.NewDefaultImports()
		editActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditAction_Create", 0), // constructor NewEditAction
			/* 1 */ imports.NewTable("TEditAction_Control", 0), // property Control
			/* 2 */ imports.NewTable("TEditAction_TClass", 0), // function TEditActionClass
		}
	})
	return editActionImport
}
