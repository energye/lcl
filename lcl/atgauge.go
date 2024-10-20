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

// IATGauge Parent: IGraphicControl
type IATGauge interface {
	IGraphicControl
	PercentDone() int32                            // property
	Theme() (resultPATFlatTheme TATFlatTheme)      // property
	SetTheme(AValue *TATFlatTheme)                 // property
	BorderStyle() TBorderStyle                     // property
	SetBorderStyle(AValue TBorderStyle)            // property
	DoubleBuffered() bool                          // property
	SetDoubleBuffered(AValue bool)                 // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	Kind() TATGaugeKind                            // property
	SetKind(AValue TATGaugeKind)                   // property
	Progress() int32                               // property
	SetProgress(AValue int32)                      // property
	MinValue() int32                               // property
	SetMinValue(AValue int32)                      // property
	MaxValue() int32                               // property
	SetMaxValue(AValue int32)                      // property
	ShowText() bool                                // property
	SetShowText(AValue bool)                       // property
	ShowTextInverted() bool                        // property
	SetShowTextInverted(AValue bool)               // property
	AddProgress(AValue int32)                      // procedure
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
}

// TATGauge Parent: TGraphicControl
type TATGauge struct {
	TGraphicControl
	dblClickPtr       uintptr
	contextPopupPtr   uintptr
	mouseDownPtr      uintptr
	mouseUpPtr        uintptr
	mouseMovePtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
}

func NewATGauge(AOwner IComponent) IATGauge {
	r1 := aTGaugeImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsATGauge(r1)
}

func (m *TATGauge) PercentDone() int32 {
	r1 := aTGaugeImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TATGauge) Theme() (resultPATFlatTheme TATFlatTheme) {
	r1 := aTGaugeImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return *(*TATFlatTheme)(getPointer(r1))
}

func (m *TATGauge) SetTheme(AValue *TATFlatTheme) {
	aTGaugeImportAPI().SysCallN(24, 1, m.Instance(), uintptr(unsafePointer(AValue)))
}

func (m *TATGauge) BorderStyle() TBorderStyle {
	r1 := aTGaugeImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TATGauge) SetBorderStyle(AValue TBorderStyle) {
	aTGaugeImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) DoubleBuffered() bool {
	r1 := aTGaugeImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetDoubleBuffered(AValue bool) {
	aTGaugeImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) ParentColor() bool {
	r1 := aTGaugeImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetParentColor(AValue bool) {
	aTGaugeImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) ParentShowHint() bool {
	r1 := aTGaugeImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetParentShowHint(AValue bool) {
	aTGaugeImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) Kind() TATGaugeKind {
	r1 := aTGaugeImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TATGaugeKind(r1)
}

func (m *TATGauge) SetKind(AValue TATGaugeKind) {
	aTGaugeImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) Progress() int32 {
	r1 := aTGaugeImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TATGauge) SetProgress(AValue int32) {
	aTGaugeImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) MinValue() int32 {
	r1 := aTGaugeImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TATGauge) SetMinValue(AValue int32) {
	aTGaugeImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) MaxValue() int32 {
	r1 := aTGaugeImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TATGauge) SetMaxValue(AValue int32) {
	aTGaugeImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) ShowText() bool {
	r1 := aTGaugeImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetShowText(AValue bool) {
	aTGaugeImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) ShowTextInverted() bool {
	r1 := aTGaugeImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetShowTextInverted(AValue bool) {
	aTGaugeImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func ATGaugeClass() TClass {
	ret := aTGaugeImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TATGauge) AddProgress(AValue int32) {
	aTGaugeImportAPI().SysCallN(0, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(13, m.Instance(), m.dblClickPtr)
}

func (m *TATGauge) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(12, m.Instance(), m.contextPopupPtr)
}

func (m *TATGauge) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(14, m.Instance(), m.mouseDownPtr)
}

func (m *TATGauge) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(18, m.Instance(), m.mouseUpPtr)
}

func (m *TATGauge) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(17, m.Instance(), m.mouseMovePtr)
}

func (m *TATGauge) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(15, m.Instance(), m.mouseEnterPtr)
}

func (m *TATGauge) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(16, m.Instance(), m.mouseLeavePtr)
}

func (m *TATGauge) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(19, m.Instance(), m.mouseWheelPtr)
}

func (m *TATGauge) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(20, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TATGauge) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	aTGaugeImportAPI().SysCallN(21, m.Instance(), m.mouseWheelUpPtr)
}

var (
	aTGaugeImport       *imports.Imports = nil
	aTGaugeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ATGauge_AddProgress", 0),
		/*1*/ imports.NewTable("ATGauge_BorderStyle", 0),
		/*2*/ imports.NewTable("ATGauge_Class", 0),
		/*3*/ imports.NewTable("ATGauge_Create", 0),
		/*4*/ imports.NewTable("ATGauge_DoubleBuffered", 0),
		/*5*/ imports.NewTable("ATGauge_Kind", 0),
		/*6*/ imports.NewTable("ATGauge_MaxValue", 0),
		/*7*/ imports.NewTable("ATGauge_MinValue", 0),
		/*8*/ imports.NewTable("ATGauge_ParentColor", 0),
		/*9*/ imports.NewTable("ATGauge_ParentShowHint", 0),
		/*10*/ imports.NewTable("ATGauge_PercentDone", 0),
		/*11*/ imports.NewTable("ATGauge_Progress", 0),
		/*12*/ imports.NewTable("ATGauge_SetOnContextPopup", 0),
		/*13*/ imports.NewTable("ATGauge_SetOnDblClick", 0),
		/*14*/ imports.NewTable("ATGauge_SetOnMouseDown", 0),
		/*15*/ imports.NewTable("ATGauge_SetOnMouseEnter", 0),
		/*16*/ imports.NewTable("ATGauge_SetOnMouseLeave", 0),
		/*17*/ imports.NewTable("ATGauge_SetOnMouseMove", 0),
		/*18*/ imports.NewTable("ATGauge_SetOnMouseUp", 0),
		/*19*/ imports.NewTable("ATGauge_SetOnMouseWheel", 0),
		/*20*/ imports.NewTable("ATGauge_SetOnMouseWheelDown", 0),
		/*21*/ imports.NewTable("ATGauge_SetOnMouseWheelUp", 0),
		/*22*/ imports.NewTable("ATGauge_ShowText", 0),
		/*23*/ imports.NewTable("ATGauge_ShowTextInverted", 0),
		/*24*/ imports.NewTable("ATGauge_Theme", 0),
	}
)

func aTGaugeImportAPI() *imports.Imports {
	if aTGaugeImport == nil {
		aTGaugeImport = NewDefaultImports()
		aTGaugeImport.SetImportTable(aTGaugeImportTables)
		aTGaugeImportTables = nil
	}
	return aTGaugeImport
}
