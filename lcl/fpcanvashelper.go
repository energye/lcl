//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IFPCanvasHelper Parent: IPersistent
type IFPCanvasHelper interface {
	IPersistent
	Allocated() bool                                          // property
	FixedCanvas() bool                                        // property
	Canvas() IFPCustomCanvas                                  // property
	FPColor() (resultFPColor TFPColor)                        // property
	SetFPColor(AValue *TFPColor)                              // property
	DelayAllocate() bool                                      // property
	SetDelayAllocate(AValue bool)                             // property
	AllocateResources(ACanvas IFPCustomCanvas, CanDelay bool) // procedure
	DeallocateResources()                                     // procedure
	SetOnChanging(fn TNotifyEvent)                            // property event
	SetOnChange(fn TNotifyEvent)                              // property event
}

// TFPCanvasHelper Parent: TPersistent
type TFPCanvasHelper struct {
	TPersistent
	changingPtr uintptr
	changePtr   uintptr
}

func NewFPCanvasHelper() IFPCanvasHelper {
	r1 := fPCanvasHelperImportAPI().SysCallN(4)
	return AsFPCanvasHelper(r1)
}

func (m *TFPCanvasHelper) Allocated() bool {
	r1 := fPCanvasHelperImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TFPCanvasHelper) FixedCanvas() bool {
	r1 := fPCanvasHelperImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TFPCanvasHelper) Canvas() IFPCustomCanvas {
	r1 := fPCanvasHelperImportAPI().SysCallN(2, m.Instance())
	return AsFPCustomCanvas(r1)
}

func (m *TFPCanvasHelper) FPColor() (resultFPColor TFPColor) {
	r1 := fPCanvasHelperImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPCanvasHelper) SetFPColor(AValue *TFPColor) {
	fPCanvasHelperImportAPI().SysCallN(7, 1, m.Instance(), uintptr(unsafePointer(AValue)))
}

func (m *TFPCanvasHelper) DelayAllocate() bool {
	r1 := fPCanvasHelperImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCanvasHelper) SetDelayAllocate(AValue bool) {
	fPCanvasHelperImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func FPCanvasHelperClass() TClass {
	ret := fPCanvasHelperImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TFPCanvasHelper) AllocateResources(ACanvas IFPCustomCanvas, CanDelay bool) {
	fPCanvasHelperImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(ACanvas), PascalBool(CanDelay))
}

func (m *TFPCanvasHelper) DeallocateResources() {
	fPCanvasHelperImportAPI().SysCallN(5, m.Instance())
}

func (m *TFPCanvasHelper) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	fPCanvasHelperImportAPI().SysCallN(10, m.Instance(), m.changingPtr)
}

func (m *TFPCanvasHelper) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	fPCanvasHelperImportAPI().SysCallN(9, m.Instance(), m.changePtr)
}

var (
	fPCanvasHelperImport       *imports.Imports = nil
	fPCanvasHelperImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCanvasHelper_AllocateResources", 0),
		/*1*/ imports.NewTable("FPCanvasHelper_Allocated", 0),
		/*2*/ imports.NewTable("FPCanvasHelper_Canvas", 0),
		/*3*/ imports.NewTable("FPCanvasHelper_Class", 0),
		/*4*/ imports.NewTable("FPCanvasHelper_Create", 0),
		/*5*/ imports.NewTable("FPCanvasHelper_DeallocateResources", 0),
		/*6*/ imports.NewTable("FPCanvasHelper_DelayAllocate", 0),
		/*7*/ imports.NewTable("FPCanvasHelper_FPColor", 0),
		/*8*/ imports.NewTable("FPCanvasHelper_FixedCanvas", 0),
		/*9*/ imports.NewTable("FPCanvasHelper_SetOnChange", 0),
		/*10*/ imports.NewTable("FPCanvasHelper_SetOnChanging", 0),
	}
)

func fPCanvasHelperImportAPI() *imports.Imports {
	if fPCanvasHelperImport == nil {
		fPCanvasHelperImport = NewDefaultImports()
		fPCanvasHelperImport.SetImportTable(fPCanvasHelperImportTables)
		fPCanvasHelperImportTables = nil
	}
	return fPCanvasHelperImport
}
