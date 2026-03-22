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

// ISynEditMouseAction Parent: ICollectionItem
type ISynEditMouseAction interface {
	ICollectionItem
	IsMatchingShiftState(shift types.TShiftState) bool                                                        // function
	IsMatchingClick(btn types.TSynMouseButton, cCount types.TSynMAClickCount, cDir types.TSynMAClickDir) bool // function
	IsFallback() bool                                                                                         // function
	Conflicts(other ISynEditMouseAction) bool                                                                 // function
	EqualsWithSynEditMouseActionBool(other ISynEditMouseAction, ignoreCmd bool) bool                          // function
	Clear()                                                                                                   // procedure
	Shift() types.TShiftState                                                                                 // property Shift Getter
	SetShift(value types.TShiftState)                                                                         // property Shift Setter
	// ShiftMask
	//  * ShiftMask:
	//  Only ShiftStates that are SET in the mask will be compared with "Shift"
	//  States not set in the ShiftMask will be ignored.
	ShiftMask() types.TShiftState                             // property ShiftMask Getter
	SetShiftMask(value types.TShiftState)                     // property ShiftMask Setter
	Button() types.TSynMouseButton                            // property Button Getter
	SetButton(value types.TSynMouseButton)                    // property Button Setter
	ClickCount() types.TSynMAClickCount                       // property ClickCount Getter
	SetClickCount(value types.TSynMAClickCount)               // property ClickCount Setter
	ClickDir() types.TSynMAClickDir                           // property ClickDir Getter
	SetClickDir(value types.TSynMAClickDir)                   // property ClickDir Setter
	ButtonUpRestrictions() types.TSynMAUpRestrictions         // property ButtonUpRestrictions Getter
	SetButtonUpRestrictions(value types.TSynMAUpRestrictions) // property ButtonUpRestrictions Setter
	Command() types.TSynEditorMouseCommand                    // property Command Getter
	SetCommand(value types.TSynEditorMouseCommand)            // property Command Setter
	MoveCaret() bool                                          // property MoveCaret Getter
	SetMoveCaret(value bool)                                  // property MoveCaret Setter
	IgnoreUpClick() bool                                      // property IgnoreUpClick Getter
	SetIgnoreUpClick(value bool)                              // property IgnoreUpClick Setter
	Option() types.TSynEditorMouseCommandOpt                  // property Option Getter
	SetOption(value types.TSynEditorMouseCommandOpt)          // property Option Setter
	Option2() int32                                           // property Option2 Getter
	SetOption2(value int32)                                   // property Option2 Setter
	// Priority
	//  Priority: 0 = highest / MaxInt = lowest
	Priority() types.TSynEditorMouseCommandOpt         // property Priority Getter
	SetPriority(value types.TSynEditorMouseCommandOpt) // property Priority Setter
}

type TSynEditMouseAction struct {
	TCollectionItem
}

func (m *TSynEditMouseAction) IsMatchingShiftState(shift types.TShiftState) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionAPI().SysCallN(1, m.Instance(), uintptr(shift))
	return api.GoBool(r)
}

func (m *TSynEditMouseAction) IsMatchingClick(btn types.TSynMouseButton, cCount types.TSynMAClickCount, cDir types.TSynMAClickDir) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionAPI().SysCallN(2, m.Instance(), uintptr(btn), uintptr(cCount), uintptr(cDir))
	return api.GoBool(r)
}

func (m *TSynEditMouseAction) IsFallback() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMouseAction) Conflicts(other ISynEditMouseAction) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(other))
	return api.GoBool(r)
}

func (m *TSynEditMouseAction) EqualsWithSynEditMouseActionBool(other ISynEditMouseAction, ignoreCmd bool) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(other), api.PasBool(ignoreCmd))
	return api.GoBool(r)
}

func (m *TSynEditMouseAction) Clear() {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(6, m.Instance())
}

func (m *TSynEditMouseAction) Shift() types.TShiftState {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(7, 0, m.Instance())
	return types.TShiftState(r)
}

func (m *TSynEditMouseAction) SetShift(value types.TShiftState) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) ShiftMask() types.TShiftState {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(8, 0, m.Instance())
	return types.TShiftState(r)
}

func (m *TSynEditMouseAction) SetShiftMask(value types.TShiftState) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) Button() types.TSynMouseButton {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(9, 0, m.Instance())
	return types.TSynMouseButton(r)
}

func (m *TSynEditMouseAction) SetButton(value types.TSynMouseButton) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) ClickCount() types.TSynMAClickCount {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(10, 0, m.Instance())
	return types.TSynMAClickCount(r)
}

func (m *TSynEditMouseAction) SetClickCount(value types.TSynMAClickCount) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) ClickDir() types.TSynMAClickDir {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(11, 0, m.Instance())
	return types.TSynMAClickDir(r)
}

func (m *TSynEditMouseAction) SetClickDir(value types.TSynMAClickDir) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) ButtonUpRestrictions() types.TSynMAUpRestrictions {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(12, 0, m.Instance())
	return types.TSynMAUpRestrictions(r)
}

func (m *TSynEditMouseAction) SetButtonUpRestrictions(value types.TSynMAUpRestrictions) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) Command() types.TSynEditorMouseCommand {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(13, 0, m.Instance())
	return types.TSynEditorMouseCommand(r)
}

func (m *TSynEditMouseAction) SetCommand(value types.TSynEditorMouseCommand) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) MoveCaret() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMouseAction) SetMoveCaret(value bool) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditMouseAction) IgnoreUpClick() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditMouseAction) SetIgnoreUpClick(value bool) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditMouseAction) Option() types.TSynEditorMouseCommandOpt {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(16, 0, m.Instance())
	return types.TSynEditorMouseCommandOpt(r)
}

func (m *TSynEditMouseAction) SetOption(value types.TSynEditorMouseCommandOpt) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) Option2() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditMouseAction) SetOption2(value int32) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditMouseAction) Priority() types.TSynEditorMouseCommandOpt {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionAPI().SysCallN(18, 0, m.Instance())
	return types.TSynEditorMouseCommandOpt(r)
}

func (m *TSynEditMouseAction) SetPriority(value types.TSynEditorMouseCommandOpt) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

// NewSynEditMouseAction class constructor
func NewSynEditMouseAction(collection ICollection) ISynEditMouseAction {
	r := synEditMouseActionAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsSynEditMouseAction(r)
}

var (
	synEditMouseActionOnce   base.Once
	synEditMouseActionImport *imports.Imports = nil
)

func synEditMouseActionAPI() *imports.Imports {
	synEditMouseActionOnce.Do(func() {
		synEditMouseActionImport = api.NewDefaultImports()
		synEditMouseActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMouseAction_Create", 0), // constructor NewSynEditMouseAction
			/* 1 */ imports.NewTable("TSynEditMouseAction_IsMatchingShiftState", 0), // function IsMatchingShiftState
			/* 2 */ imports.NewTable("TSynEditMouseAction_IsMatchingClick", 0), // function IsMatchingClick
			/* 3 */ imports.NewTable("TSynEditMouseAction_IsFallback", 0), // function IsFallback
			/* 4 */ imports.NewTable("TSynEditMouseAction_Conflicts", 0), // function Conflicts
			/* 5 */ imports.NewTable("TSynEditMouseAction_EqualsWithSynEditMouseActionBool", 0), // function EqualsWithSynEditMouseActionBool
			/* 6 */ imports.NewTable("TSynEditMouseAction_Clear", 0), // procedure Clear
			/* 7 */ imports.NewTable("TSynEditMouseAction_Shift", 0), // property Shift
			/* 8 */ imports.NewTable("TSynEditMouseAction_ShiftMask", 0), // property ShiftMask
			/* 9 */ imports.NewTable("TSynEditMouseAction_Button", 0), // property Button
			/* 10 */ imports.NewTable("TSynEditMouseAction_ClickCount", 0), // property ClickCount
			/* 11 */ imports.NewTable("TSynEditMouseAction_ClickDir", 0), // property ClickDir
			/* 12 */ imports.NewTable("TSynEditMouseAction_ButtonUpRestrictions", 0), // property ButtonUpRestrictions
			/* 13 */ imports.NewTable("TSynEditMouseAction_Command", 0), // property Command
			/* 14 */ imports.NewTable("TSynEditMouseAction_MoveCaret", 0), // property MoveCaret
			/* 15 */ imports.NewTable("TSynEditMouseAction_IgnoreUpClick", 0), // property IgnoreUpClick
			/* 16 */ imports.NewTable("TSynEditMouseAction_Option", 0), // property Option
			/* 17 */ imports.NewTable("TSynEditMouseAction_Option2", 0), // property Option2
			/* 18 */ imports.NewTable("TSynEditMouseAction_Priority", 0), // property Priority
		}
	})
	return synEditMouseActionImport
}
