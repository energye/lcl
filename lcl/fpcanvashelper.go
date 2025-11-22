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
)

// IFPCanvasHelper Parent: IPersistent
type IFPCanvasHelper interface {
	IPersistent
	AllocateResources(canvas IFPCustomCanvas, canDelay bool) // procedure
	// DeallocateResources
	//  free all resource used by this helper
	DeallocateResources() // procedure
	Allocated() bool      // property Allocated Getter
	// FixedCanvas
	//  properties cannot be changed when allocated
	FixedCanvas() bool // property FixedCanvas Getter
	// Canvas
	//  Canvas for which the helper is allocated
	Canvas() IFPCustomCanvas // property Canvas Getter
	// FPColor
	//  color of the helper
	FPColor() TFPColor             // property FPColor Getter
	SetFPColor(value TFPColor)     // property FPColor Setter
	DelayAllocate() bool           // property DelayAllocate Getter
	SetDelayAllocate(value bool)   // property DelayAllocate Setter
	SetOnChanging(fn TNotifyEvent) // property event
	SetOnChange(fn TNotifyEvent)   // property event
}

type TFPCanvasHelper struct {
	TPersistent
}

func (m *TFPCanvasHelper) AllocateResources(canvas IFPCustomCanvas, canDelay bool) {
	if !m.IsValid() {
		return
	}
	fPCanvasHelperAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(canvas), api.PasBool(canDelay))
}

func (m *TFPCanvasHelper) DeallocateResources() {
	if !m.IsValid() {
		return
	}
	fPCanvasHelperAPI().SysCallN(2, m.Instance())
}

func (m *TFPCanvasHelper) Allocated() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCanvasHelperAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCanvasHelper) FixedCanvas() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCanvasHelperAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCanvasHelper) Canvas() IFPCustomCanvas {
	if !m.IsValid() {
		return nil
	}
	r := fPCanvasHelperAPI().SysCallN(5, m.Instance())
	return AsFPCustomCanvas(r)
}

func (m *TFPCanvasHelper) FPColor() (result TFPColor) {
	if !m.IsValid() {
		return
	}
	fPCanvasHelperAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPCanvasHelper) SetFPColor(value TFPColor) {
	if !m.IsValid() {
		return
	}
	fPCanvasHelperAPI().SysCallN(6, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TFPCanvasHelper) DelayAllocate() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCanvasHelperAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCanvasHelper) SetDelayAllocate(value bool) {
	if !m.IsValid() {
		return
	}
	fPCanvasHelperAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCanvasHelper) SetOnChanging(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, fPCanvasHelperAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFPCanvasHelper) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, fPCanvasHelperAPI(), api.MakeEventDataPtr(cb))
}

// NewFPCanvasHelper class constructor
func NewFPCanvasHelper() IFPCanvasHelper {
	r := fPCanvasHelperAPI().SysCallN(0)
	return AsFPCanvasHelper(r)
}

var (
	fPCanvasHelperOnce   base.Once
	fPCanvasHelperImport *imports.Imports = nil
)

func fPCanvasHelperAPI() *imports.Imports {
	fPCanvasHelperOnce.Do(func() {
		fPCanvasHelperImport = api.NewDefaultImports()
		fPCanvasHelperImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCanvasHelper_Create", 0), // constructor NewFPCanvasHelper
			/* 1 */ imports.NewTable("TFPCanvasHelper_AllocateResources", 0), // procedure AllocateResources
			/* 2 */ imports.NewTable("TFPCanvasHelper_DeallocateResources", 0), // procedure DeallocateResources
			/* 3 */ imports.NewTable("TFPCanvasHelper_Allocated", 0), // property Allocated
			/* 4 */ imports.NewTable("TFPCanvasHelper_FixedCanvas", 0), // property FixedCanvas
			/* 5 */ imports.NewTable("TFPCanvasHelper_Canvas", 0), // property Canvas
			/* 6 */ imports.NewTable("TFPCanvasHelper_FPColor", 0), // property FPColor
			/* 7 */ imports.NewTable("TFPCanvasHelper_DelayAllocate", 0), // property DelayAllocate
			/* 8 */ imports.NewTable("TFPCanvasHelper_OnChanging", 0), // event OnChanging
			/* 9 */ imports.NewTable("TFPCanvasHelper_OnChange", 0), // event OnChange
		}
	})
	return fPCanvasHelperImport
}
