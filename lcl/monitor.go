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

// IMonitor Parent: IObject
type IMonitor interface {
	IObject
	Handle() types.HMONITOR    // property Handle Getter
	MonitorNum() int32         // property MonitorNum Getter
	Left() int32               // property Left Getter
	Height() int32             // property Height Getter
	Top() int32                // property Top Getter
	Width() int32              // property Width Getter
	BoundsRect() types.TRect   // property BoundsRect Getter
	WorkareaRect() types.TRect // property WorkareaRect Getter
	Primary() bool             // property Primary Getter
	PixelsPerInch() int32      // property PixelsPerInch Getter
}

type TMonitor struct {
	TObject
}

func (m *TMonitor) Handle() types.HMONITOR {
	if !m.IsValid() {
		return 0
	}
	r := monitorAPI().SysCallN(1, m.Instance())
	return types.HMONITOR(r)
}

func (m *TMonitor) MonitorNum() int32 {
	if !m.IsValid() {
		return 0
	}
	r := monitorAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TMonitor) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := monitorAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TMonitor) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := monitorAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TMonitor) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := monitorAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TMonitor) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := monitorAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TMonitor) BoundsRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	monitorAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TMonitor) WorkareaRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	monitorAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TMonitor) Primary() bool {
	if !m.IsValid() {
		return false
	}
	r := monitorAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TMonitor) PixelsPerInch() int32 {
	if !m.IsValid() {
		return 0
	}
	r := monitorAPI().SysCallN(10, m.Instance())
	return int32(r)
}

// NewMonitor class constructor
func NewMonitor() IMonitor {
	r := monitorAPI().SysCallN(0)
	return AsMonitor(r)
}

var (
	monitorOnce   base.Once
	monitorImport *imports.Imports = nil
)

func monitorAPI() *imports.Imports {
	monitorOnce.Do(func() {
		monitorImport = api.NewDefaultImports()
		monitorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMonitor_Create", 0), // constructor NewMonitor
			/* 1 */ imports.NewTable("TMonitor_Handle", 0), // property Handle
			/* 2 */ imports.NewTable("TMonitor_MonitorNum", 0), // property MonitorNum
			/* 3 */ imports.NewTable("TMonitor_Left", 0), // property Left
			/* 4 */ imports.NewTable("TMonitor_Height", 0), // property Height
			/* 5 */ imports.NewTable("TMonitor_Top", 0), // property Top
			/* 6 */ imports.NewTable("TMonitor_Width", 0), // property Width
			/* 7 */ imports.NewTable("TMonitor_BoundsRect", 0), // property BoundsRect
			/* 8 */ imports.NewTable("TMonitor_WorkareaRect", 0), // property WorkareaRect
			/* 9 */ imports.NewTable("TMonitor_Primary", 0), // property Primary
			/* 10 */ imports.NewTable("TMonitor_PixelsPerInch", 0), // property PixelsPerInch
		}
	})
	return monitorImport
}
