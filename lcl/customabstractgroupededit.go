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

// ICustomAbstractGroupedEdit Is Abstract Class Parent: ICustomControl
type ICustomAbstractGroupedEdit interface {
	ICustomControl
	AutoSizeHeightIsEditHeight() bool              // property
	SetAutoSizeHeightIsEditHeight(AValue bool)     // property
	Alignment() TAlignment                         // property
	SetAlignment(AValue TAlignment)                // property
	CanUndo() bool                                 // property
	CaretPos() (resultPoint TPoint)                // property
	SetCaretPos(AValue *TPoint)                    // property
	CharCase() TEditCharCase                       // property
	SetCharCase(AValue TEditCharCase)              // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	EchoMode() TEchoMode                           // property
	SetEchoMode(AValue TEchoMode)                  // property
	HideSelection() bool                           // property
	SetHideSelection(AValue bool)                  // property
	MaxLength() int32                              // property
	SetMaxLength(AValue int32)                     // property
	Modified() bool                                // property
	SetModified(AValue bool)                       // property
	NumbersOnly() bool                             // property
	SetNumbersOnly(AValue bool)                    // property
	PasswordChar() Char                            // property
	SetPasswordChar(AValue Char)                   // property
	ReadOnly() bool                                // property
	SetReadOnly(AValue bool)                       // property
	SelLength() int32                              // property
	SetSelLength(AValue int32)                     // property
	SelStart() int32                               // property
	SetSelStart(AValue int32)                      // property
	SelText() string                               // property
	SetSelText(AValue string)                      // property
	Text() string                                  // property
	SetText(AValue string)                         // property
	TextHint() string                              // property
	SetTextHint(AValue string)                     // property
	Clear()                                        // procedure
	ClearSelection()                               // procedure
	CopyToClipboard()                              // procedure
	CutToClipboard()                               // procedure
	PasteFromClipboard()                           // procedure
	SelectAll()                                    // procedure
	Undo()                                         // procedure
	ValidateEdit()                                 // procedure
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TCustomAbstractGroupedEdit Is Abstract Class Parent: TCustomControl
type TCustomAbstractGroupedEdit struct {
	TCustomControl
	changePtr         uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	editingDonePtr    uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseWheelPtr     uintptr
	mouseWheelUpPtr   uintptr
	mouseWheelDownPtr uintptr
	mouseUpPtr        uintptr
	startDragPtr      uintptr
}

func (m *TCustomAbstractGroupedEdit) AutoSizeHeightIsEditHeight() bool {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetAutoSizeHeightIsEditHeight(AValue bool) {
	customAbstractGroupedEditImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) Alignment() TAlignment {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomAbstractGroupedEdit) SetAlignment(AValue TAlignment) {
	customAbstractGroupedEditImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) CanUndo() bool {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) CaretPos() (resultPoint TPoint) {
	customAbstractGroupedEditImportAPI().SysCallN(3, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomAbstractGroupedEdit) SetCaretPos(AValue *TPoint) {
	customAbstractGroupedEditImportAPI().SysCallN(3, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomAbstractGroupedEdit) CharCase() TEditCharCase {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomAbstractGroupedEdit) SetCharCase(AValue TEditCharCase) {
	customAbstractGroupedEditImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) ParentColor() bool {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetParentColor(AValue bool) {
	customAbstractGroupedEditImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) EchoMode() TEchoMode {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TEchoMode(r1)
}

func (m *TCustomAbstractGroupedEdit) SetEchoMode(AValue TEchoMode) {
	customAbstractGroupedEditImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) HideSelection() bool {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetHideSelection(AValue bool) {
	customAbstractGroupedEditImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) MaxLength() int32 {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetMaxLength(AValue int32) {
	customAbstractGroupedEditImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) Modified() bool {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetModified(AValue bool) {
	customAbstractGroupedEditImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) NumbersOnly() bool {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetNumbersOnly(AValue bool) {
	customAbstractGroupedEditImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) PasswordChar() Char {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomAbstractGroupedEdit) SetPasswordChar(AValue Char) {
	customAbstractGroupedEditImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) ReadOnly() bool {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAbstractGroupedEdit) SetReadOnly(AValue bool) {
	customAbstractGroupedEditImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelLength() int32 {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelLength(AValue int32) {
	customAbstractGroupedEditImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelStart() int32 {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelStart(AValue int32) {
	customAbstractGroupedEditImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAbstractGroupedEdit) SelText() string {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetSelText(AValue string) {
	customAbstractGroupedEditImportAPI().SysCallN(21, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAbstractGroupedEdit) Text() string {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetText(AValue string) {
	customAbstractGroupedEditImportAPI().SysCallN(39, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAbstractGroupedEdit) TextHint() string {
	r1 := customAbstractGroupedEditImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAbstractGroupedEdit) SetTextHint(AValue string) {
	customAbstractGroupedEditImportAPI().SysCallN(40, 1, m.Instance(), PascalStr(AValue))
}

func CustomAbstractGroupedEditClass() TClass {
	ret := customAbstractGroupedEditImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TCustomAbstractGroupedEdit) Clear() {
	customAbstractGroupedEditImportAPI().SysCallN(6, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ClearSelection() {
	customAbstractGroupedEditImportAPI().SysCallN(7, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CopyToClipboard() {
	customAbstractGroupedEditImportAPI().SysCallN(8, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CutToClipboard() {
	customAbstractGroupedEditImportAPI().SysCallN(9, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) PasteFromClipboard() {
	customAbstractGroupedEditImportAPI().SysCallN(17, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) SelectAll() {
	customAbstractGroupedEditImportAPI().SysCallN(22, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) Undo() {
	customAbstractGroupedEditImportAPI().SysCallN(41, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ValidateEdit() {
	customAbstractGroupedEditImportAPI().SysCallN(42, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(23, m.Instance(), m.changePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(24, m.Instance(), m.contextPopupPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(25, m.Instance(), m.dblClickPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(26, m.Instance(), m.dragDropPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(27, m.Instance(), m.dragOverPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(28, m.Instance(), m.editingDonePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(29, m.Instance(), m.endDragPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(30, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(31, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(32, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(33, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(35, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(37, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(36, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(34, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomAbstractGroupedEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	customAbstractGroupedEditImportAPI().SysCallN(38, m.Instance(), m.startDragPtr)
}

var (
	customAbstractGroupedEditImport       *imports.Imports = nil
	customAbstractGroupedEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomAbstractGroupedEdit_Alignment", 0),
		/*1*/ imports.NewTable("CustomAbstractGroupedEdit_AutoSizeHeightIsEditHeight", 0),
		/*2*/ imports.NewTable("CustomAbstractGroupedEdit_CanUndo", 0),
		/*3*/ imports.NewTable("CustomAbstractGroupedEdit_CaretPos", 0),
		/*4*/ imports.NewTable("CustomAbstractGroupedEdit_CharCase", 0),
		/*5*/ imports.NewTable("CustomAbstractGroupedEdit_Class", 0),
		/*6*/ imports.NewTable("CustomAbstractGroupedEdit_Clear", 0),
		/*7*/ imports.NewTable("CustomAbstractGroupedEdit_ClearSelection", 0),
		/*8*/ imports.NewTable("CustomAbstractGroupedEdit_CopyToClipboard", 0),
		/*9*/ imports.NewTable("CustomAbstractGroupedEdit_CutToClipboard", 0),
		/*10*/ imports.NewTable("CustomAbstractGroupedEdit_EchoMode", 0),
		/*11*/ imports.NewTable("CustomAbstractGroupedEdit_HideSelection", 0),
		/*12*/ imports.NewTable("CustomAbstractGroupedEdit_MaxLength", 0),
		/*13*/ imports.NewTable("CustomAbstractGroupedEdit_Modified", 0),
		/*14*/ imports.NewTable("CustomAbstractGroupedEdit_NumbersOnly", 0),
		/*15*/ imports.NewTable("CustomAbstractGroupedEdit_ParentColor", 0),
		/*16*/ imports.NewTable("CustomAbstractGroupedEdit_PasswordChar", 0),
		/*17*/ imports.NewTable("CustomAbstractGroupedEdit_PasteFromClipboard", 0),
		/*18*/ imports.NewTable("CustomAbstractGroupedEdit_ReadOnly", 0),
		/*19*/ imports.NewTable("CustomAbstractGroupedEdit_SelLength", 0),
		/*20*/ imports.NewTable("CustomAbstractGroupedEdit_SelStart", 0),
		/*21*/ imports.NewTable("CustomAbstractGroupedEdit_SelText", 0),
		/*22*/ imports.NewTable("CustomAbstractGroupedEdit_SelectAll", 0),
		/*23*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnChange", 0),
		/*24*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnContextPopup", 0),
		/*25*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnDblClick", 0),
		/*26*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnDragDrop", 0),
		/*27*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnDragOver", 0),
		/*28*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnEditingDone", 0),
		/*29*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnEndDrag", 0),
		/*30*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseDown", 0),
		/*31*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseEnter", 0),
		/*32*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseLeave", 0),
		/*33*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseMove", 0),
		/*34*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseUp", 0),
		/*35*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseWheel", 0),
		/*36*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseWheelDown", 0),
		/*37*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnMouseWheelUp", 0),
		/*38*/ imports.NewTable("CustomAbstractGroupedEdit_SetOnStartDrag", 0),
		/*39*/ imports.NewTable("CustomAbstractGroupedEdit_Text", 0),
		/*40*/ imports.NewTable("CustomAbstractGroupedEdit_TextHint", 0),
		/*41*/ imports.NewTable("CustomAbstractGroupedEdit_Undo", 0),
		/*42*/ imports.NewTable("CustomAbstractGroupedEdit_ValidateEdit", 0),
	}
)

func customAbstractGroupedEditImportAPI() *imports.Imports {
	if customAbstractGroupedEditImport == nil {
		customAbstractGroupedEditImport = NewDefaultImports()
		customAbstractGroupedEditImport.SetImportTable(customAbstractGroupedEditImportTables)
		customAbstractGroupedEditImportTables = nil
	}
	return customAbstractGroupedEditImport
}
