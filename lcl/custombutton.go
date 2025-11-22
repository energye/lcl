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

// ICustomButton Parent: IButtonControl
type ICustomButton interface {
	IButtonControl
	Click()                                  // procedure
	Active() bool                            // property Active Getter
	Default() bool                           // property Default Getter
	SetDefault(value bool)                   // property Default Setter
	ModalResult() types.TModalResult         // property ModalResult Getter
	SetModalResult(value types.TModalResult) // property ModalResult Setter
	ShortCut() types.TShortCut               // property ShortCut Getter
	ShortCutKey2() types.TShortCut           // property ShortCutKey2 Getter
	Cancel() bool                            // property Cancel Getter
	SetCancel(value bool)                    // property Cancel Setter
}

type TCustomButton struct {
	TButtonControl
}

func (m *TCustomButton) Click() {
	if !m.IsValid() {
		return
	}
	customButtonAPI().SysCallN(1, m.Instance())
}

func (m *TCustomButton) Active() bool {
	if !m.IsValid() {
		return false
	}
	r := customButtonAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomButton) Default() bool {
	if !m.IsValid() {
		return false
	}
	r := customButtonAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomButton) SetDefault(value bool) {
	if !m.IsValid() {
		return
	}
	customButtonAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomButton) ModalResult() types.TModalResult {
	if !m.IsValid() {
		return 0
	}
	r := customButtonAPI().SysCallN(4, 0, m.Instance())
	return types.TModalResult(r)
}

func (m *TCustomButton) SetModalResult(value types.TModalResult) {
	if !m.IsValid() {
		return
	}
	customButtonAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomButton) ShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := customButtonAPI().SysCallN(5, m.Instance())
	return types.TShortCut(r)
}

func (m *TCustomButton) ShortCutKey2() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := customButtonAPI().SysCallN(6, m.Instance())
	return types.TShortCut(r)
}

func (m *TCustomButton) Cancel() bool {
	if !m.IsValid() {
		return false
	}
	r := customButtonAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomButton) SetCancel(value bool) {
	if !m.IsValid() {
		return
	}
	customButtonAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

// NewCustomButton class constructor
func NewCustomButton(theOwner IComponent) ICustomButton {
	r := customButtonAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomButton(r)
}

func TCustomButtonClass() types.TClass {
	r := customButtonAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	customButtonOnce   base.Once
	customButtonImport *imports.Imports = nil
)

func customButtonAPI() *imports.Imports {
	customButtonOnce.Do(func() {
		customButtonImport = api.NewDefaultImports()
		customButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomButton_Create", 0), // constructor NewCustomButton
			/* 1 */ imports.NewTable("TCustomButton_Click", 0), // procedure Click
			/* 2 */ imports.NewTable("TCustomButton_Active", 0), // property Active
			/* 3 */ imports.NewTable("TCustomButton_Default", 0), // property Default
			/* 4 */ imports.NewTable("TCustomButton_ModalResult", 0), // property ModalResult
			/* 5 */ imports.NewTable("TCustomButton_ShortCut", 0), // property ShortCut
			/* 6 */ imports.NewTable("TCustomButton_ShortCutKey2", 0), // property ShortCutKey2
			/* 7 */ imports.NewTable("TCustomButton_Cancel", 0), // property Cancel
			/* 8 */ imports.NewTable("TCustomButton_TClass", 0), // function TCustomButtonClass
		}
	})
	return customButtonImport
}
