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

// IXButton Parent: IGraphicControl
type IXButton interface {
	IGraphicControl
	Caption() string                   // property
	SetCaption(AValue string)          // property
	ShowCaption() bool                 // property
	SetShowCaption(AValue bool)        // property
	BackColor() TColor                 // property
	SetBackColor(AValue TColor)        // property
	HoverColor() TColor                // property
	SetHoverColor(AValue TColor)       // property
	DownColor() TColor                 // property
	SetDownColor(AValue TColor)        // property
	BorderWidth() int32                // property
	SetBorderWidth(AValue int32)       // property
	BorderColor() TColor               // property
	SetBorderColor(AValue TColor)      // property
	Picture() IPicture                 // property
	SetPicture(AValue IPicture)        // property
	DrawMode() TDrawImageMode          // property
	SetDrawMode(AValue TDrawImageMode) // property
	NormalFontColor() TColor           // property
	SetNormalFontColor(AValue TColor)  // property
	DownFontColor() TColor             // property
	SetDownFontColor(AValue TColor)    // property
	HoverFontColor() TColor            // property
	SetHoverFontColor(AValue TColor)   // property
	ParentFont() bool                  // property
	SetParentFont(AValue bool)         // property
	ParentShowHint() bool              // property
	SetParentShowHint(AValue bool)     // property
	Paint()                            // procedure
	Resize()                           // procedure
	SetOnDblClick(fn TNotifyEvent)     // property event
	SetOnMouseDown(fn TMouseEvent)     // property event
	SetOnMouseEnter(fn TNotifyEvent)   // property event
	SetOnMouseLeave(fn TNotifyEvent)   // property event
	SetOnMouseMove(fn TMouseMoveEvent) // property event
	SetOnMouseUp(fn TMouseEvent)       // property event
}

// TXButton Parent: TGraphicControl
type TXButton struct {
	TGraphicControl
	dblClickPtr   uintptr
	mouseDownPtr  uintptr
	mouseEnterPtr uintptr
	mouseLeavePtr uintptr
	mouseMovePtr  uintptr
	mouseUpPtr    uintptr
}

func NewXButton(AOwner IComponent) IXButton {
	r1 := xButtonImportAPI().SysCallN(5, GetObjectUintptr(AOwner))
	return AsXButton(r1)
}

func (m *TXButton) Caption() string {
	r1 := xButtonImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TXButton) SetCaption(AValue string) {
	xButtonImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func (m *TXButton) ShowCaption() bool {
	r1 := xButtonImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetShowCaption(AValue bool) {
	xButtonImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) BackColor() TColor {
	r1 := xButtonImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBackColor(AValue TColor) {
	xButtonImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverColor() TColor {
	r1 := xButtonImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverColor(AValue TColor) {
	xButtonImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownColor() TColor {
	r1 := xButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownColor(AValue TColor) {
	xButtonImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderWidth() int32 {
	r1 := xButtonImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TXButton) SetBorderWidth(AValue int32) {
	xButtonImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderColor() TColor {
	r1 := xButtonImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBorderColor(AValue TColor) {
	xButtonImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) Picture() IPicture {
	r1 := xButtonImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TXButton) SetPicture(AValue IPicture) {
	xButtonImportAPI().SysCallN(15, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TXButton) DrawMode() TDrawImageMode {
	r1 := xButtonImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TDrawImageMode(r1)
}

func (m *TXButton) SetDrawMode(AValue TDrawImageMode) {
	xButtonImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) NormalFontColor() TColor {
	r1 := xButtonImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetNormalFontColor(AValue TColor) {
	xButtonImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownFontColor() TColor {
	r1 := xButtonImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownFontColor(AValue TColor) {
	xButtonImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverFontColor() TColor {
	r1 := xButtonImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverFontColor(AValue TColor) {
	xButtonImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) ParentFont() bool {
	r1 := xButtonImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentFont(AValue bool) {
	xButtonImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) ParentShowHint() bool {
	r1 := xButtonImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentShowHint(AValue bool) {
	xButtonImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func XButtonClass() TClass {
	ret := xButtonImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TXButton) Paint() {
	xButtonImportAPI().SysCallN(12, m.Instance())
}

func (m *TXButton) Resize() {
	xButtonImportAPI().SysCallN(16, m.Instance())
}

func (m *TXButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	xButtonImportAPI().SysCallN(17, m.Instance(), m.dblClickPtr)
}

func (m *TXButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	xButtonImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TXButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	xButtonImportAPI().SysCallN(19, m.Instance(), m.mouseEnterPtr)
}

func (m *TXButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	xButtonImportAPI().SysCallN(20, m.Instance(), m.mouseLeavePtr)
}

func (m *TXButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	xButtonImportAPI().SysCallN(21, m.Instance(), m.mouseMovePtr)
}

func (m *TXButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	xButtonImportAPI().SysCallN(22, m.Instance(), m.mouseUpPtr)
}

var (
	xButtonImport       *imports.Imports = nil
	xButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("XButton_BackColor", 0),
		/*1*/ imports.NewTable("XButton_BorderColor", 0),
		/*2*/ imports.NewTable("XButton_BorderWidth", 0),
		/*3*/ imports.NewTable("XButton_Caption", 0),
		/*4*/ imports.NewTable("XButton_Class", 0),
		/*5*/ imports.NewTable("XButton_Create", 0),
		/*6*/ imports.NewTable("XButton_DownColor", 0),
		/*7*/ imports.NewTable("XButton_DownFontColor", 0),
		/*8*/ imports.NewTable("XButton_DrawMode", 0),
		/*9*/ imports.NewTable("XButton_HoverColor", 0),
		/*10*/ imports.NewTable("XButton_HoverFontColor", 0),
		/*11*/ imports.NewTable("XButton_NormalFontColor", 0),
		/*12*/ imports.NewTable("XButton_Paint", 0),
		/*13*/ imports.NewTable("XButton_ParentFont", 0),
		/*14*/ imports.NewTable("XButton_ParentShowHint", 0),
		/*15*/ imports.NewTable("XButton_Picture", 0),
		/*16*/ imports.NewTable("XButton_Resize", 0),
		/*17*/ imports.NewTable("XButton_SetOnDblClick", 0),
		/*18*/ imports.NewTable("XButton_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("XButton_SetOnMouseEnter", 0),
		/*20*/ imports.NewTable("XButton_SetOnMouseLeave", 0),
		/*21*/ imports.NewTable("XButton_SetOnMouseMove", 0),
		/*22*/ imports.NewTable("XButton_SetOnMouseUp", 0),
		/*23*/ imports.NewTable("XButton_ShowCaption", 0),
	}
)

func xButtonImportAPI() *imports.Imports {
	if xButtonImport == nil {
		xButtonImport = NewDefaultImports()
		xButtonImport.SetImportTable(xButtonImportTables)
		xButtonImportTables = nil
	}
	return xButtonImport
}
