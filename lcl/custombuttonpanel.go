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

// ICustomButtonPanel Parent: ICustomPanel
type ICustomButtonPanel interface {
	ICustomPanel
	OKButton() IPanelBitBtn                      // property OKButton Getter
	HelpButton() IPanelBitBtn                    // property HelpButton Getter
	CloseButton() IPanelBitBtn                   // property CloseButton Getter
	CancelButton() IPanelBitBtn                  // property CancelButton Getter
	ButtonOrder() types.TButtonOrder             // property ButtonOrder Getter
	SetButtonOrder(value types.TButtonOrder)     // property ButtonOrder Setter
	DefaultButton() types.TPanelButtonEx         // property DefaultButton Getter
	SetDefaultButton(value types.TPanelButtonEx) // property DefaultButton Setter
	ShowButtons() types.TPanelButtons            // property ShowButtons Getter
	SetShowButtons(value types.TPanelButtons)    // property ShowButtons Setter
	ShowGlyphs() types.TPanelButtons             // property ShowGlyphs Getter
	SetShowGlyphs(value types.TPanelButtons)     // property ShowGlyphs Setter
	ShowBevel() bool                             // property ShowBevel Getter
	SetShowBevel(value bool)                     // property ShowBevel Setter
	Spacing() types.TSpacingSize                 // property Spacing Getter
	SetSpacing(value types.TSpacingSize)         // property Spacing Setter
}

type TCustomButtonPanel struct {
	TCustomPanel
}

func (m *TCustomButtonPanel) OKButton() IPanelBitBtn {
	if !m.IsValid() {
		return nil
	}
	r := customButtonPanelAPI().SysCallN(1, m.Instance())
	return AsPanelBitBtn(r)
}

func (m *TCustomButtonPanel) HelpButton() IPanelBitBtn {
	if !m.IsValid() {
		return nil
	}
	r := customButtonPanelAPI().SysCallN(2, m.Instance())
	return AsPanelBitBtn(r)
}

func (m *TCustomButtonPanel) CloseButton() IPanelBitBtn {
	if !m.IsValid() {
		return nil
	}
	r := customButtonPanelAPI().SysCallN(3, m.Instance())
	return AsPanelBitBtn(r)
}

func (m *TCustomButtonPanel) CancelButton() IPanelBitBtn {
	if !m.IsValid() {
		return nil
	}
	r := customButtonPanelAPI().SysCallN(4, m.Instance())
	return AsPanelBitBtn(r)
}

func (m *TCustomButtonPanel) ButtonOrder() types.TButtonOrder {
	if !m.IsValid() {
		return 0
	}
	r := customButtonPanelAPI().SysCallN(5, 0, m.Instance())
	return types.TButtonOrder(r)
}

func (m *TCustomButtonPanel) SetButtonOrder(value types.TButtonOrder) {
	if !m.IsValid() {
		return
	}
	customButtonPanelAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomButtonPanel) DefaultButton() types.TPanelButtonEx {
	if !m.IsValid() {
		return 0
	}
	r := customButtonPanelAPI().SysCallN(6, 0, m.Instance())
	return types.TPanelButtonEx(r)
}

func (m *TCustomButtonPanel) SetDefaultButton(value types.TPanelButtonEx) {
	if !m.IsValid() {
		return
	}
	customButtonPanelAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomButtonPanel) ShowButtons() types.TPanelButtons {
	if !m.IsValid() {
		return 0
	}
	r := customButtonPanelAPI().SysCallN(7, 0, m.Instance())
	return types.TPanelButtons(r)
}

func (m *TCustomButtonPanel) SetShowButtons(value types.TPanelButtons) {
	if !m.IsValid() {
		return
	}
	customButtonPanelAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomButtonPanel) ShowGlyphs() types.TPanelButtons {
	if !m.IsValid() {
		return 0
	}
	r := customButtonPanelAPI().SysCallN(8, 0, m.Instance())
	return types.TPanelButtons(r)
}

func (m *TCustomButtonPanel) SetShowGlyphs(value types.TPanelButtons) {
	if !m.IsValid() {
		return
	}
	customButtonPanelAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomButtonPanel) ShowBevel() bool {
	if !m.IsValid() {
		return false
	}
	r := customButtonPanelAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomButtonPanel) SetShowBevel(value bool) {
	if !m.IsValid() {
		return
	}
	customButtonPanelAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomButtonPanel) Spacing() types.TSpacingSize {
	if !m.IsValid() {
		return 0
	}
	r := customButtonPanelAPI().SysCallN(10, 0, m.Instance())
	return types.TSpacingSize(r)
}

func (m *TCustomButtonPanel) SetSpacing(value types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	customButtonPanelAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

// NewCustomButtonPanel class constructor
func NewCustomButtonPanel(owner IComponent) ICustomButtonPanel {
	r := customButtonPanelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomButtonPanel(r)
}

func TCustomButtonPanelClass() types.TClass {
	r := customButtonPanelAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	customButtonPanelOnce   base.Once
	customButtonPanelImport *imports.Imports = nil
)

func customButtonPanelAPI() *imports.Imports {
	customButtonPanelOnce.Do(func() {
		customButtonPanelImport = api.NewDefaultImports()
		customButtonPanelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomButtonPanel_Create", 0), // constructor NewCustomButtonPanel
			/* 1 */ imports.NewTable("TCustomButtonPanel_OKButton", 0), // property OKButton
			/* 2 */ imports.NewTable("TCustomButtonPanel_HelpButton", 0), // property HelpButton
			/* 3 */ imports.NewTable("TCustomButtonPanel_CloseButton", 0), // property CloseButton
			/* 4 */ imports.NewTable("TCustomButtonPanel_CancelButton", 0), // property CancelButton
			/* 5 */ imports.NewTable("TCustomButtonPanel_ButtonOrder", 0), // property ButtonOrder
			/* 6 */ imports.NewTable("TCustomButtonPanel_DefaultButton", 0), // property DefaultButton
			/* 7 */ imports.NewTable("TCustomButtonPanel_ShowButtons", 0), // property ShowButtons
			/* 8 */ imports.NewTable("TCustomButtonPanel_ShowGlyphs", 0), // property ShowGlyphs
			/* 9 */ imports.NewTable("TCustomButtonPanel_ShowBevel", 0), // property ShowBevel
			/* 10 */ imports.NewTable("TCustomButtonPanel_Spacing", 0), // property Spacing
			/* 11 */ imports.NewTable("TCustomButtonPanel_TClass", 0), // function TCustomButtonPanelClass
		}
	})
	return customButtonPanelImport
}
