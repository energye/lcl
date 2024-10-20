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

// IColorButton Parent: ICustomSpeedButton
type IColorButton interface {
	ICustomSpeedButton
	BorderWidth() int32                            // property
	SetBorderWidth(AValue int32)                   // property
	ButtonColorAutoSize() bool                     // property
	SetButtonColorAutoSize(AValue bool)            // property
	ButtonColorSize() int32                        // property
	SetButtonColorSize(AValue int32)               // property
	ButtonColor() TColor                           // property
	SetButtonColor(AValue TColor)                  // property
	ColorDialog() IColorDialog                     // property
	SetColorDialog(AValue IColorDialog)            // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnColorChanged(fn TNotifyEvent)             // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnPaint(fn TNotifyEvent)                    // property event
}

// TColorButton Parent: TCustomSpeedButton
type TColorButton struct {
	TCustomSpeedButton
	colorChangedPtr   uintptr
	dblClickPtr       uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	paintPtr          uintptr
}

func NewColorButton(AnOwner IComponent) IColorButton {
	r1 := colorButtonImportAPI().SysCallN(6, GetObjectUintptr(AnOwner))
	return AsColorButton(r1)
}

func (m *TColorButton) BorderWidth() int32 {
	r1 := colorButtonImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorButton) SetBorderWidth(AValue int32) {
	colorButtonImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorButton) ButtonColorAutoSize() bool {
	r1 := colorButtonImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorButton) SetButtonColorAutoSize(AValue bool) {
	colorButtonImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorButton) ButtonColorSize() int32 {
	r1 := colorButtonImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorButton) SetButtonColorSize(AValue int32) {
	colorButtonImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorButton) ButtonColor() TColor {
	r1 := colorButtonImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TColorButton) SetButtonColor(AValue TColor) {
	colorButtonImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorButton) ColorDialog() IColorDialog {
	r1 := colorButtonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsColorDialog(r1)
}

func (m *TColorButton) SetColorDialog(AValue IColorDialog) {
	colorButtonImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TColorButton) ParentFont() bool {
	r1 := colorButtonImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorButton) SetParentFont(AValue bool) {
	colorButtonImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorButton) ParentShowHint() bool {
	r1 := colorButtonImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorButton) SetParentShowHint(AValue bool) {
	colorButtonImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func ColorButtonClass() TClass {
	ret := colorButtonImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TColorButton) SetOnColorChanged(fn TNotifyEvent) {
	if m.colorChangedPtr != 0 {
		RemoveEventElement(m.colorChangedPtr)
	}
	m.colorChangedPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(9, m.Instance(), m.colorChangedPtr)
}

func (m *TColorButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(10, m.Instance(), m.dblClickPtr)
}

func (m *TColorButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(11, m.Instance(), m.mouseDownPtr)
}

func (m *TColorButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(12, m.Instance(), m.mouseEnterPtr)
}

func (m *TColorButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(13, m.Instance(), m.mouseLeavePtr)
}

func (m *TColorButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(14, m.Instance(), m.mouseMovePtr)
}

func (m *TColorButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(15, m.Instance(), m.mouseUpPtr)
}

func (m *TColorButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(16, m.Instance(), m.mouseWheelPtr)
}

func (m *TColorButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(17, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TColorButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(18, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TColorButton) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	colorButtonImportAPI().SysCallN(19, m.Instance(), m.paintPtr)
}

var (
	colorButtonImport       *imports.Imports = nil
	colorButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ColorButton_BorderWidth", 0),
		/*1*/ imports.NewTable("ColorButton_ButtonColor", 0),
		/*2*/ imports.NewTable("ColorButton_ButtonColorAutoSize", 0),
		/*3*/ imports.NewTable("ColorButton_ButtonColorSize", 0),
		/*4*/ imports.NewTable("ColorButton_Class", 0),
		/*5*/ imports.NewTable("ColorButton_ColorDialog", 0),
		/*6*/ imports.NewTable("ColorButton_Create", 0),
		/*7*/ imports.NewTable("ColorButton_ParentFont", 0),
		/*8*/ imports.NewTable("ColorButton_ParentShowHint", 0),
		/*9*/ imports.NewTable("ColorButton_SetOnColorChanged", 0),
		/*10*/ imports.NewTable("ColorButton_SetOnDblClick", 0),
		/*11*/ imports.NewTable("ColorButton_SetOnMouseDown", 0),
		/*12*/ imports.NewTable("ColorButton_SetOnMouseEnter", 0),
		/*13*/ imports.NewTable("ColorButton_SetOnMouseLeave", 0),
		/*14*/ imports.NewTable("ColorButton_SetOnMouseMove", 0),
		/*15*/ imports.NewTable("ColorButton_SetOnMouseUp", 0),
		/*16*/ imports.NewTable("ColorButton_SetOnMouseWheel", 0),
		/*17*/ imports.NewTable("ColorButton_SetOnMouseWheelDown", 0),
		/*18*/ imports.NewTable("ColorButton_SetOnMouseWheelUp", 0),
		/*19*/ imports.NewTable("ColorButton_SetOnPaint", 0),
	}
)

func colorButtonImportAPI() *imports.Imports {
	if colorButtonImport == nil {
		colorButtonImport = NewDefaultImports()
		colorButtonImport.SetImportTable(colorButtonImportTables)
		colorButtonImportTables = nil
	}
	return colorButtonImport
}
