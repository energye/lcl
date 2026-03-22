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

// ISynEditMouseActions Parent: ICollection
type ISynEditMouseActions interface {
	ICollection
	AddToSynEditMouseAction() ISynEditMouseAction                                                                                                                                                                                                                                                                                                                                                                            // function
	EqualsWithSynEditMouseActions(other ISynEditMouseActions) bool                                                                                                                                                                                                                                                                                                                                                           // function
	FindCommand(anInfo TSynEditMouseActionInfo, previous ISynEditMouseAction) ISynEditMouseAction                                                                                                                                                                                                                                                                                                                            // function
	IndexOf(mAction ISynEditMouseAction, ignoreCmd bool) int32                                                                                                                                                                                                                                                                                                                                                               // function
	AssertNoConflict(mAction ISynEditMouseAction)                                                                                                                                                                                                                                                                                                                                                                            // procedure
	ResetDefaults()                                                                                                                                                                                                                                                                                                                                                                                                          // procedure
	IncAssertLock()                                                                                                                                                                                                                                                                                                                                                                                                          // procedure
	DecAssertLock()                                                                                                                                                                                                                                                                                                                                                                                                          // procedure
	AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSStateX2SEMCOptIntX2(cmd types.TSynEditorMouseCommand, moveCaret bool, button types.TSynMouseButton, clickCount types.TSynMAClickCount, dir types.TSynMAClickDir, shift types.TShiftState, shiftMask types.TShiftState, opt types.TSynEditorMouseCommandOpt, prior int32, opt2 int32, ignoreUpClick bool)                                                          // procedure
	AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSMAURestrictionsSStateX2SEMCOptIntX2(cmd types.TSynEditorMouseCommand, moveCaret bool, button types.TSynMouseButton, clickCount types.TSynMAClickCount, dir types.TSynMAClickDir, anUpRestrict types.TSynMAUpRestrictions, shift types.TShiftState, shiftMask types.TShiftState, opt types.TSynEditorMouseCommandOpt, prior int32, opt2 int32, ignoreUpClick bool) // procedure
	ItemsWithIntToSynEditMouseAction(index int32) ISynEditMouseAction                                                                                                                                                                                                                                                                                                                                                        // property Items Getter
	SetItemsWithIntToSynEditMouseAction(index int32, value ISynEditMouseAction)                                                                                                                                                                                                                                                                                                                                              // property Items Setter
}

type TSynEditMouseActions struct {
	TCollection
}

func (m *TSynEditMouseActions) AddToSynEditMouseAction() ISynEditMouseAction {
	if !m.IsValid() {
		return nil
	}
	r := synEditMouseActionsAPI().SysCallN(1, m.Instance())
	return AsSynEditMouseAction(r)
}

func (m *TSynEditMouseActions) EqualsWithSynEditMouseActions(other ISynEditMouseActions) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditMouseActionsAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(other))
	return api.GoBool(r)
}

func (m *TSynEditMouseActions) FindCommand(anInfo TSynEditMouseActionInfo, previous ISynEditMouseAction) ISynEditMouseAction {
	if !m.IsValid() {
		return nil
	}
	anInfoPtr := anInfo.ToPas()
	r := synEditMouseActionsAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(anInfoPtr)), base.GetObjectUintptr(previous))
	return AsSynEditMouseAction(r)
}

func (m *TSynEditMouseActions) IndexOf(mAction ISynEditMouseAction, ignoreCmd bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMouseActionsAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(mAction), api.PasBool(ignoreCmd))
	return int32(r)
}

func (m *TSynEditMouseActions) AssertNoConflict(mAction ISynEditMouseAction) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionsAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(mAction))
}

func (m *TSynEditMouseActions) ResetDefaults() {
	if !m.IsValid() {
		return
	}
	synEditMouseActionsAPI().SysCallN(6, m.Instance())
}

func (m *TSynEditMouseActions) IncAssertLock() {
	if !m.IsValid() {
		return
	}
	synEditMouseActionsAPI().SysCallN(7, m.Instance())
}

func (m *TSynEditMouseActions) DecAssertLock() {
	if !m.IsValid() {
		return
	}
	synEditMouseActionsAPI().SysCallN(8, m.Instance())
}

func (m *TSynEditMouseActions) AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSStateX2SEMCOptIntX2(cmd types.TSynEditorMouseCommand, moveCaret bool, button types.TSynMouseButton, clickCount types.TSynMAClickCount, dir types.TSynMAClickDir, shift types.TShiftState, shiftMask types.TShiftState, opt types.TSynEditorMouseCommandOpt, prior int32, opt2 int32, ignoreUpClick bool) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionsAPI().SysCallN(9, m.Instance(), uintptr(cmd), api.PasBool(moveCaret), uintptr(button), uintptr(clickCount), uintptr(dir), uintptr(shift), uintptr(shiftMask), uintptr(opt), uintptr(prior), uintptr(opt2), api.PasBool(ignoreUpClick))
}

func (m *TSynEditMouseActions) AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSMAURestrictionsSStateX2SEMCOptIntX2(cmd types.TSynEditorMouseCommand, moveCaret bool, button types.TSynMouseButton, clickCount types.TSynMAClickCount, dir types.TSynMAClickDir, anUpRestrict types.TSynMAUpRestrictions, shift types.TShiftState, shiftMask types.TShiftState, opt types.TSynEditorMouseCommandOpt, prior int32, opt2 int32, ignoreUpClick bool) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionsAPI().SysCallN(10, m.Instance(), uintptr(cmd), api.PasBool(moveCaret), uintptr(button), uintptr(clickCount), uintptr(dir), uintptr(anUpRestrict), uintptr(shift), uintptr(shiftMask), uintptr(opt), uintptr(prior), uintptr(opt2), api.PasBool(ignoreUpClick))
}

func (m *TSynEditMouseActions) ItemsWithIntToSynEditMouseAction(index int32) ISynEditMouseAction {
	if !m.IsValid() {
		return nil
	}
	r := synEditMouseActionsAPI().SysCallN(11, 0, m.Instance(), uintptr(index))
	return AsSynEditMouseAction(r)
}

func (m *TSynEditMouseActions) SetItemsWithIntToSynEditMouseAction(index int32, value ISynEditMouseAction) {
	if !m.IsValid() {
		return
	}
	synEditMouseActionsAPI().SysCallN(11, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewSynEditMouseActions class constructor
func NewSynEditMouseActions(owner IPersistent) ISynEditMouseActions {
	r := synEditMouseActionsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynEditMouseActions(r)
}

var (
	synEditMouseActionsOnce   base.Once
	synEditMouseActionsImport *imports.Imports = nil
)

func synEditMouseActionsAPI() *imports.Imports {
	synEditMouseActionsOnce.Do(func() {
		synEditMouseActionsImport = api.NewDefaultImports()
		synEditMouseActionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMouseActions_Create", 0), // constructor NewSynEditMouseActions
			/* 1 */ imports.NewTable("TSynEditMouseActions_AddToSynEditMouseAction", 0), // function AddToSynEditMouseAction
			/* 2 */ imports.NewTable("TSynEditMouseActions_EqualsWithSynEditMouseActions", 0), // function EqualsWithSynEditMouseActions
			/* 3 */ imports.NewTable("TSynEditMouseActions_FindCommand", 0), // function FindCommand
			/* 4 */ imports.NewTable("TSynEditMouseActions_IndexOf", 0), // function IndexOf
			/* 5 */ imports.NewTable("TSynEditMouseActions_AssertNoConflict", 0), // procedure AssertNoConflict
			/* 6 */ imports.NewTable("TSynEditMouseActions_ResetDefaults", 0), // procedure ResetDefaults
			/* 7 */ imports.NewTable("TSynEditMouseActions_IncAssertLock", 0), // procedure IncAssertLock
			/* 8 */ imports.NewTable("TSynEditMouseActions_DecAssertLock", 0), // procedure DecAssertLock
			/* 9 */ imports.NewTable("TSynEditMouseActions_AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSStateX2SEMCOptIntX2", 0), // procedure AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSStateX2SEMCOptIntX2
			/* 10 */ imports.NewTable("TSynEditMouseActions_AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSMAURestrictionsSStateX2SEMCOptIntX2", 0), // procedure AddCommandWithSEMCommandBoolX2SMButtonSMACCountSMACDirSMAURestrictionsSStateX2SEMCOptIntX2
			/* 11 */ imports.NewTable("TSynEditMouseActions_ItemsWithIntToSynEditMouseAction", 0), // property ItemsWithIntToSynEditMouseAction
		}
	})
	return synEditMouseActionsImport
}
