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

// IStatusPanel Parent: ICollectionItem
type IStatusPanel interface {
	ICollectionItem
	StatusBar() IStatusBar                  // function
	Alignment() types.TAlignment            // property Alignment Getter
	SetAlignment(value types.TAlignment)    // property Alignment Setter
	Bevel() types.TStatusPanelBevel         // property Bevel Getter
	SetBevel(value types.TStatusPanelBevel) // property Bevel Setter
	BidiMode() types.TBiDiMode              // property BidiMode Getter
	SetBidiMode(value types.TBiDiMode)      // property BidiMode Setter
	ParentBiDiMode() bool                   // property ParentBiDiMode Getter
	SetParentBiDiMode(value bool)           // property ParentBiDiMode Setter
	Style() types.TStatusPanelStyle         // property Style Getter
	SetStyle(value types.TStatusPanelStyle) // property Style Setter
	Text() string                           // property Text Getter
	SetText(value string)                   // property Text Setter
	Width() int32                           // property Width Getter
	SetWidth(value int32)                   // property Width Setter
}

type TStatusPanel struct {
	TCollectionItem
}

func (m *TStatusPanel) StatusBar() IStatusBar {
	if !m.IsValid() {
		return nil
	}
	r := statusPanelAPI().SysCallN(1, m.Instance())
	return AsStatusBar(r)
}

func (m *TStatusPanel) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := statusPanelAPI().SysCallN(2, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TStatusPanel) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	statusPanelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TStatusPanel) Bevel() types.TStatusPanelBevel {
	if !m.IsValid() {
		return 0
	}
	r := statusPanelAPI().SysCallN(3, 0, m.Instance())
	return types.TStatusPanelBevel(r)
}

func (m *TStatusPanel) SetBevel(value types.TStatusPanelBevel) {
	if !m.IsValid() {
		return
	}
	statusPanelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TStatusPanel) BidiMode() types.TBiDiMode {
	if !m.IsValid() {
		return 0
	}
	r := statusPanelAPI().SysCallN(4, 0, m.Instance())
	return types.TBiDiMode(r)
}

func (m *TStatusPanel) SetBidiMode(value types.TBiDiMode) {
	if !m.IsValid() {
		return
	}
	statusPanelAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TStatusPanel) ParentBiDiMode() bool {
	if !m.IsValid() {
		return false
	}
	r := statusPanelAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusPanel) SetParentBiDiMode(value bool) {
	if !m.IsValid() {
		return
	}
	statusPanelAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusPanel) Style() types.TStatusPanelStyle {
	if !m.IsValid() {
		return 0
	}
	r := statusPanelAPI().SysCallN(6, 0, m.Instance())
	return types.TStatusPanelStyle(r)
}

func (m *TStatusPanel) SetStyle(value types.TStatusPanelStyle) {
	if !m.IsValid() {
		return
	}
	statusPanelAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TStatusPanel) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := statusPanelAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TStatusPanel) SetText(value string) {
	if !m.IsValid() {
		return
	}
	statusPanelAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TStatusPanel) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := statusPanelAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TStatusPanel) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	statusPanelAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

// NewStatusPanel class constructor
func NewStatusPanel(collection ICollection) IStatusPanel {
	r := statusPanelAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsStatusPanel(r)
}

var (
	statusPanelOnce   base.Once
	statusPanelImport *imports.Imports = nil
)

func statusPanelAPI() *imports.Imports {
	statusPanelOnce.Do(func() {
		statusPanelImport = api.NewDefaultImports()
		statusPanelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStatusPanel_Create", 0), // constructor NewStatusPanel
			/* 1 */ imports.NewTable("TStatusPanel_StatusBar", 0), // function StatusBar
			/* 2 */ imports.NewTable("TStatusPanel_Alignment", 0), // property Alignment
			/* 3 */ imports.NewTable("TStatusPanel_Bevel", 0), // property Bevel
			/* 4 */ imports.NewTable("TStatusPanel_BidiMode", 0), // property BidiMode
			/* 5 */ imports.NewTable("TStatusPanel_ParentBiDiMode", 0), // property ParentBiDiMode
			/* 6 */ imports.NewTable("TStatusPanel_Style", 0), // property Style
			/* 7 */ imports.NewTable("TStatusPanel_Text", 0), // property Text
			/* 8 */ imports.NewTable("TStatusPanel_Width", 0), // property Width
		}
	})
	return statusPanelImport
}
