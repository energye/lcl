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

// ICustomEdit Parent: IWinControl
type ICustomEdit interface {
	IWinControl
	Alignment() TAlignment                           // property
	SetAlignment(AValue TAlignment)                  // property
	BorderStyle() TBorderStyle                       // property
	SetBorderStyle(AValue TBorderStyle)              // property
	CanUndo() bool                                   // property
	CaretPos() (resultPoint TPoint)                  // property
	SetCaretPos(AValue *TPoint)                      // property
	CharCase() TEditCharCase                         // property
	SetCharCase(AValue TEditCharCase)                // property
	EchoMode() TEchoMode                             // property
	SetEchoMode(AValue TEchoMode)                    // property
	EmulatedTextHintStatus() TEmulatedTextHintStatus // property
	HideSelection() bool                             // property
	SetHideSelection(AValue bool)                    // property
	MaxLength() int32                                // property
	SetMaxLength(AValue int32)                       // property
	Modified() bool                                  // property
	SetModified(AValue bool)                         // property
	NumbersOnly() bool                               // property
	SetNumbersOnly(AValue bool)                      // property
	PasswordChar() Char                              // property
	SetPasswordChar(AValue Char)                     // property
	ReadOnly() bool                                  // property
	SetReadOnly(AValue bool)                         // property
	SelLength() int32                                // property
	SetSelLength(AValue int32)                       // property
	SelStart() int32                                 // property
	SetSelStart(AValue int32)                        // property
	SelText() string                                 // property
	SetSelText(AValue string)                        // property
	Text() string                                    // property
	SetText(AValue string)                           // property
	TextHint() string                                // property
	SetTextHint(AValue string)                       // property
	Clear()                                          // procedure
	SelectAll()                                      // procedure
	ClearSelection()                                 // procedure
	CopyToClipboard()                                // procedure
	CutToClipboard()                                 // procedure
	PasteFromClipboard()                             // procedure
	Undo()                                           // procedure
	SetOnChange(fn TNotifyEvent)                     // property event
}

// TCustomEdit Parent: TWinControl
type TCustomEdit struct {
	TWinControl
	changePtr uintptr
}

func NewCustomEdit(AOwner IComponent) ICustomEdit {
	r1 := customEditImportAPI().SysCallN(9, GetObjectUintptr(AOwner))
	return AsCustomEdit(r1)
}

func (m *TCustomEdit) Alignment() TAlignment {
	r1 := customEditImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomEdit) SetAlignment(AValue TAlignment) {
	customEditImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) BorderStyle() TBorderStyle {
	r1 := customEditImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomEdit) SetBorderStyle(AValue TBorderStyle) {
	customEditImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) CanUndo() bool {
	r1 := customEditImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TCustomEdit) CaretPos() (resultPoint TPoint) {
	customEditImportAPI().SysCallN(3, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomEdit) SetCaretPos(AValue *TPoint) {
	customEditImportAPI().SysCallN(3, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomEdit) CharCase() TEditCharCase {
	r1 := customEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomEdit) SetCharCase(AValue TEditCharCase) {
	customEditImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) EchoMode() TEchoMode {
	r1 := customEditImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TEchoMode(r1)
}

func (m *TCustomEdit) SetEchoMode(AValue TEchoMode) {
	customEditImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) EmulatedTextHintStatus() TEmulatedTextHintStatus {
	r1 := customEditImportAPI().SysCallN(12, m.Instance())
	return TEmulatedTextHintStatus(r1)
}

func (m *TCustomEdit) HideSelection() bool {
	r1 := customEditImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetHideSelection(AValue bool) {
	customEditImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) MaxLength() int32 {
	r1 := customEditImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetMaxLength(AValue int32) {
	customEditImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) Modified() bool {
	r1 := customEditImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetModified(AValue bool) {
	customEditImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) NumbersOnly() bool {
	r1 := customEditImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetNumbersOnly(AValue bool) {
	customEditImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) PasswordChar() Char {
	r1 := customEditImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomEdit) SetPasswordChar(AValue Char) {
	customEditImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) ReadOnly() bool {
	r1 := customEditImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomEdit) SetReadOnly(AValue bool) {
	customEditImportAPI().SysCallN(19, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomEdit) SelLength() int32 {
	r1 := customEditImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetSelLength(AValue int32) {
	customEditImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) SelStart() int32 {
	r1 := customEditImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomEdit) SetSelStart(AValue int32) {
	customEditImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomEdit) SelText() string {
	r1 := customEditImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetSelText(AValue string) {
	customEditImportAPI().SysCallN(22, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomEdit) Text() string {
	r1 := customEditImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetText(AValue string) {
	customEditImportAPI().SysCallN(25, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomEdit) TextHint() string {
	r1 := customEditImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomEdit) SetTextHint(AValue string) {
	customEditImportAPI().SysCallN(26, 1, m.Instance(), PascalStr(AValue))
}

func CustomEditClass() TClass {
	ret := customEditImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TCustomEdit) Clear() {
	customEditImportAPI().SysCallN(6, m.Instance())
}

func (m *TCustomEdit) SelectAll() {
	customEditImportAPI().SysCallN(23, m.Instance())
}

func (m *TCustomEdit) ClearSelection() {
	customEditImportAPI().SysCallN(7, m.Instance())
}

func (m *TCustomEdit) CopyToClipboard() {
	customEditImportAPI().SysCallN(8, m.Instance())
}

func (m *TCustomEdit) CutToClipboard() {
	customEditImportAPI().SysCallN(10, m.Instance())
}

func (m *TCustomEdit) PasteFromClipboard() {
	customEditImportAPI().SysCallN(18, m.Instance())
}

func (m *TCustomEdit) Undo() {
	customEditImportAPI().SysCallN(27, m.Instance())
}

func (m *TCustomEdit) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customEditImportAPI().SysCallN(24, m.Instance(), m.changePtr)
}

var (
	customEditImport       *imports.Imports = nil
	customEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomEdit_Alignment", 0),
		/*1*/ imports.NewTable("CustomEdit_BorderStyle", 0),
		/*2*/ imports.NewTable("CustomEdit_CanUndo", 0),
		/*3*/ imports.NewTable("CustomEdit_CaretPos", 0),
		/*4*/ imports.NewTable("CustomEdit_CharCase", 0),
		/*5*/ imports.NewTable("CustomEdit_Class", 0),
		/*6*/ imports.NewTable("CustomEdit_Clear", 0),
		/*7*/ imports.NewTable("CustomEdit_ClearSelection", 0),
		/*8*/ imports.NewTable("CustomEdit_CopyToClipboard", 0),
		/*9*/ imports.NewTable("CustomEdit_Create", 0),
		/*10*/ imports.NewTable("CustomEdit_CutToClipboard", 0),
		/*11*/ imports.NewTable("CustomEdit_EchoMode", 0),
		/*12*/ imports.NewTable("CustomEdit_EmulatedTextHintStatus", 0),
		/*13*/ imports.NewTable("CustomEdit_HideSelection", 0),
		/*14*/ imports.NewTable("CustomEdit_MaxLength", 0),
		/*15*/ imports.NewTable("CustomEdit_Modified", 0),
		/*16*/ imports.NewTable("CustomEdit_NumbersOnly", 0),
		/*17*/ imports.NewTable("CustomEdit_PasswordChar", 0),
		/*18*/ imports.NewTable("CustomEdit_PasteFromClipboard", 0),
		/*19*/ imports.NewTable("CustomEdit_ReadOnly", 0),
		/*20*/ imports.NewTable("CustomEdit_SelLength", 0),
		/*21*/ imports.NewTable("CustomEdit_SelStart", 0),
		/*22*/ imports.NewTable("CustomEdit_SelText", 0),
		/*23*/ imports.NewTable("CustomEdit_SelectAll", 0),
		/*24*/ imports.NewTable("CustomEdit_SetOnChange", 0),
		/*25*/ imports.NewTable("CustomEdit_Text", 0),
		/*26*/ imports.NewTable("CustomEdit_TextHint", 0),
		/*27*/ imports.NewTable("CustomEdit_Undo", 0),
	}
)

func customEditImportAPI() *imports.Imports {
	if customEditImport == nil {
		customEditImport = NewDefaultImports()
		customEditImport.SetImportTable(customEditImportTables)
		customEditImportTables = nil
	}
	return customEditImport
}
