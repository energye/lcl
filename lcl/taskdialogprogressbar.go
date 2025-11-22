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

// ITaskDialogProgressBar Parent: IPersistent
type ITaskDialogProgressBar interface {
	IPersistent
	Initialize()                            // procedure
	SetRange(min int32, max int32)          // procedure
	MarqueeSpeed() uint32                   // property MarqueeSpeed Getter
	SetMarqueeSpeed(value uint32)           // property MarqueeSpeed Setter
	Max() int32                             // property Max Getter
	SetMax(value int32)                     // property Max Setter
	Min() int32                             // property Min Getter
	SetMin(value int32)                     // property Min Setter
	Position() int32                        // property Position Getter
	SetPosition(value int32)                // property Position Setter
	State() types.TProgressBarState         // property State Getter
	SetState(value types.TProgressBarState) // property State Setter
}

type TTaskDialogProgressBar struct {
	TPersistent
}

func (m *TTaskDialogProgressBar) Initialize() {
	if !m.IsValid() {
		return
	}
	taskDialogProgressBarAPI().SysCallN(1, m.Instance())
}

func (m *TTaskDialogProgressBar) SetRange(min int32, max int32) {
	if !m.IsValid() {
		return
	}
	taskDialogProgressBarAPI().SysCallN(2, m.Instance(), uintptr(min), uintptr(max))
}

func (m *TTaskDialogProgressBar) MarqueeSpeed() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := taskDialogProgressBarAPI().SysCallN(3, 0, m.Instance())
	return uint32(r)
}

func (m *TTaskDialogProgressBar) SetMarqueeSpeed(value uint32) {
	if !m.IsValid() {
		return
	}
	taskDialogProgressBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TTaskDialogProgressBar) Max() int32 {
	if !m.IsValid() {
		return 0
	}
	r := taskDialogProgressBarAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TTaskDialogProgressBar) SetMax(value int32) {
	if !m.IsValid() {
		return
	}
	taskDialogProgressBarAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TTaskDialogProgressBar) Min() int32 {
	if !m.IsValid() {
		return 0
	}
	r := taskDialogProgressBarAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TTaskDialogProgressBar) SetMin(value int32) {
	if !m.IsValid() {
		return
	}
	taskDialogProgressBarAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TTaskDialogProgressBar) Position() int32 {
	if !m.IsValid() {
		return 0
	}
	r := taskDialogProgressBarAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TTaskDialogProgressBar) SetPosition(value int32) {
	if !m.IsValid() {
		return
	}
	taskDialogProgressBarAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TTaskDialogProgressBar) State() types.TProgressBarState {
	if !m.IsValid() {
		return 0
	}
	r := taskDialogProgressBarAPI().SysCallN(7, 0, m.Instance())
	return types.TProgressBarState(r)
}

func (m *TTaskDialogProgressBar) SetState(value types.TProgressBarState) {
	if !m.IsValid() {
		return
	}
	taskDialogProgressBarAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

// NewTaskDialogProgressBar class constructor
func NewTaskDialogProgressBar(dialog ICustomTaskDialog) ITaskDialogProgressBar {
	r := taskDialogProgressBarAPI().SysCallN(0, base.GetObjectUintptr(dialog))
	return AsTaskDialogProgressBar(r)
}

var (
	taskDialogProgressBarOnce   base.Once
	taskDialogProgressBarImport *imports.Imports = nil
)

func taskDialogProgressBarAPI() *imports.Imports {
	taskDialogProgressBarOnce.Do(func() {
		taskDialogProgressBarImport = api.NewDefaultImports()
		taskDialogProgressBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTaskDialogProgressBar_Create", 0), // constructor NewTaskDialogProgressBar
			/* 1 */ imports.NewTable("TTaskDialogProgressBar_Initialize", 0), // procedure Initialize
			/* 2 */ imports.NewTable("TTaskDialogProgressBar_SetRange", 0), // procedure SetRange
			/* 3 */ imports.NewTable("TTaskDialogProgressBar_MarqueeSpeed", 0), // property MarqueeSpeed
			/* 4 */ imports.NewTable("TTaskDialogProgressBar_Max", 0), // property Max
			/* 5 */ imports.NewTable("TTaskDialogProgressBar_Min", 0), // property Min
			/* 6 */ imports.NewTable("TTaskDialogProgressBar_Position", 0), // property Position
			/* 7 */ imports.NewTable("TTaskDialogProgressBar_State", 0), // property State
		}
	})
	return taskDialogProgressBarImport
}
