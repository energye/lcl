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

// ISynPluginSyncroEdit Parent: ISynPluginCustomSyncroEdit
type ISynPluginSyncroEdit interface {
	ISynPluginCustomSyncroEdit
	ScanModes() types.TSynPluginSyncroScanModes         // property ScanModes Getter
	SetScanModes(value types.TSynPluginSyncroScanModes) // property ScanModes Setter
	GutterGlyph() IBitmap                               // property GutterGlyph Getter
	SetGutterGlyph(value IBitmap)                       // property GutterGlyph Setter
	KeystrokesSelecting() ISynEditKeyStrokes            // property KeystrokesSelecting Getter
	SetKeystrokesSelecting(value ISynEditKeyStrokes)    // property KeystrokesSelecting Setter
	Keystrokes() ISynEditKeyStrokes                     // property Keystrokes Getter
	SetKeystrokes(value ISynEditKeyStrokes)             // property Keystrokes Setter
	KeystrokesOffCell() ISynEditKeyStrokes              // property KeystrokesOffCell Getter
	SetKeystrokesOffCell(value ISynEditKeyStrokes)      // property KeystrokesOffCell Setter
	SetOnModeChange(fn TNotifyEvent)                    // property event
	SetOnBeginEdit(fn TNotifyEvent)                     // property event
	SetOnEndEdit(fn TNotifyEvent)                       // property event
}

type TSynPluginSyncroEdit struct {
	TSynPluginCustomSyncroEdit
}

func (m *TSynPluginSyncroEdit) ScanModes() types.TSynPluginSyncroScanModes {
	if !m.IsValid() {
		return 0
	}
	r := synPluginSyncroEditAPI().SysCallN(1, 0, m.Instance())
	return types.TSynPluginSyncroScanModes(r)
}

func (m *TSynPluginSyncroEdit) SetScanModes(value types.TSynPluginSyncroScanModes) {
	if !m.IsValid() {
		return
	}
	synPluginSyncroEditAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TSynPluginSyncroEdit) GutterGlyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncroEditAPI().SysCallN(2, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TSynPluginSyncroEdit) SetGutterGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	synPluginSyncroEditAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncroEdit) KeystrokesSelecting() ISynEditKeyStrokes {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncroEditAPI().SysCallN(3, 0, m.Instance())
	return AsSynEditKeyStrokes(r)
}

func (m *TSynPluginSyncroEdit) SetKeystrokesSelecting(value ISynEditKeyStrokes) {
	if !m.IsValid() {
		return
	}
	synPluginSyncroEditAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncroEdit) Keystrokes() ISynEditKeyStrokes {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncroEditAPI().SysCallN(4, 0, m.Instance())
	return AsSynEditKeyStrokes(r)
}

func (m *TSynPluginSyncroEdit) SetKeystrokes(value ISynEditKeyStrokes) {
	if !m.IsValid() {
		return
	}
	synPluginSyncroEditAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncroEdit) KeystrokesOffCell() ISynEditKeyStrokes {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncroEditAPI().SysCallN(5, 0, m.Instance())
	return AsSynEditKeyStrokes(r)
}

func (m *TSynPluginSyncroEdit) SetKeystrokesOffCell(value ISynEditKeyStrokes) {
	if !m.IsValid() {
		return
	}
	synPluginSyncroEditAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncroEdit) SetOnModeChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, synPluginSyncroEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynPluginSyncroEdit) SetOnBeginEdit(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, synPluginSyncroEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynPluginSyncroEdit) SetOnEndEdit(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, synPluginSyncroEditAPI(), api.MakeEventDataPtr(cb))
}

// NewSynPluginSyncroEdit class constructor
func NewSynPluginSyncroEdit(owner IComponent) ISynPluginSyncroEdit {
	r := synPluginSyncroEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynPluginSyncroEdit(r)
}

func TSynPluginSyncroEditClass() types.TClass {
	r := synPluginSyncroEditAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	synPluginSyncroEditOnce   base.Once
	synPluginSyncroEditImport *imports.Imports = nil
)

func synPluginSyncroEditAPI() *imports.Imports {
	synPluginSyncroEditOnce.Do(func() {
		synPluginSyncroEditImport = api.NewDefaultImports()
		synPluginSyncroEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynPluginSyncroEdit_Create", 0), // constructor NewSynPluginSyncroEdit
			/* 1 */ imports.NewTable("TSynPluginSyncroEdit_ScanModes", 0), // property ScanModes
			/* 2 */ imports.NewTable("TSynPluginSyncroEdit_GutterGlyph", 0), // property GutterGlyph
			/* 3 */ imports.NewTable("TSynPluginSyncroEdit_KeystrokesSelecting", 0), // property KeystrokesSelecting
			/* 4 */ imports.NewTable("TSynPluginSyncroEdit_Keystrokes", 0), // property Keystrokes
			/* 5 */ imports.NewTable("TSynPluginSyncroEdit_KeystrokesOffCell", 0), // property KeystrokesOffCell
			/* 6 */ imports.NewTable("TSynPluginSyncroEdit_OnModeChange", 0), // event OnModeChange
			/* 7 */ imports.NewTable("TSynPluginSyncroEdit_OnBeginEdit", 0), // event OnBeginEdit
			/* 8 */ imports.NewTable("TSynPluginSyncroEdit_OnEndEdit", 0), // event OnEndEdit
			/* 9 */ imports.NewTable("TSynPluginSyncroEdit_TClass", 0), // function TSynPluginSyncroEditClass
		}
	})
	return synPluginSyncroEditImport
}
