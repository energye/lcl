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

// ICustomEdit Parent: IWinControl
type ICustomEdit interface {
	IWinControl
	Clear()                              // procedure
	SelectAll()                          // procedure
	ClearSelection()                     // procedure
	CopyToClipboard()                    // procedure
	CutToClipboard()                     // procedure
	PasteFromClipboard()                 // procedure
	Undo()                               // procedure
	Alignment() types.TAlignment         // property Alignment Getter
	SetAlignment(value types.TAlignment) // property Alignment Setter
	// BorderStyle
	//  properties which are not supported by all descendents
	BorderStyle() types.TBorderStyle                       // property BorderStyle Getter
	SetBorderStyle(value types.TBorderStyle)               // property BorderStyle Setter
	CanUndo() bool                                         // property CanUndo Getter
	CaretPos() types.TPoint                                // property CaretPos Getter
	SetCaretPos(value types.TPoint)                        // property CaretPos Setter
	CharCase() types.TEditCharCase                         // property CharCase Getter
	SetCharCase(value types.TEditCharCase)                 // property CharCase Setter
	EchoMode() types.TEchoMode                             // property EchoMode Getter
	SetEchoMode(value types.TEchoMode)                     // property EchoMode Setter
	EmulatedTextHintStatus() types.TEmulatedTextHintStatus // property EmulatedTextHintStatus Getter
	HideSelection() bool                                   // property HideSelection Getter
	SetHideSelection(value bool)                           // property HideSelection Setter
	MaxLength() int32                                      // property MaxLength Getter
	SetMaxLength(value int32)                              // property MaxLength Setter
	Modified() bool                                        // property Modified Getter
	SetModified(value bool)                                // property Modified Setter
	NumbersOnly() bool                                     // property NumbersOnly Getter
	SetNumbersOnly(value bool)                             // property NumbersOnly Setter
	PasswordChar() uint16                                  // property PasswordChar Getter
	SetPasswordChar(value uint16)                          // property PasswordChar Setter
	ReadOnly() bool                                        // property ReadOnly Getter
	SetReadOnly(value bool)                                // property ReadOnly Setter
	SelLength() int32                                      // property SelLength Getter
	SetSelLength(value int32)                              // property SelLength Setter
	SelStart() int32                                       // property SelStart Getter
	SetSelStart(value int32)                               // property SelStart Setter
	SelText() string                                       // property SelText Getter
	SetSelText(value string)                               // property SelText Setter
	Text() string                                          // property Text Getter
	SetText(value string)                                  // property Text Setter
	TextHint() string                                      // property TextHint Getter
	SetTextHint(value string)                              // property TextHint Setter
	SetOnChange(fn TNotifyEvent)                           // property event
}

type TCustomEdit struct {
	TWinControl
}

func (m *TCustomEdit) Clear() {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(1, m.Instance())
}

func (m *TCustomEdit) SelectAll() {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(2, m.Instance())
}

func (m *TCustomEdit) ClearSelection() {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(3, m.Instance())
}

func (m *TCustomEdit) CopyToClipboard() {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(4, m.Instance())
}

func (m *TCustomEdit) CutToClipboard() {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(5, m.Instance())
}

func (m *TCustomEdit) PasteFromClipboard() {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(6, m.Instance())
}

func (m *TCustomEdit) Undo() {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(7, m.Instance())
}

func (m *TCustomEdit) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(8, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TCustomEdit) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(9, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCustomEdit) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) CanUndo() bool {
	if !m.IsValid() {
		return false
	}
	r := customEditAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomEdit) CaretPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomEdit) SetCaretPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomEdit) CharCase() types.TEditCharCase {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(12, 0, m.Instance())
	return types.TEditCharCase(r)
}

func (m *TCustomEdit) SetCharCase(value types.TEditCharCase) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) EchoMode() types.TEchoMode {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(13, 0, m.Instance())
	return types.TEchoMode(r)
}

func (m *TCustomEdit) SetEchoMode(value types.TEchoMode) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) EmulatedTextHintStatus() types.TEmulatedTextHintStatus {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(14, m.Instance())
	return types.TEmulatedTextHintStatus(r)
}

func (m *TCustomEdit) HideSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := customEditAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomEdit) SetHideSelection(value bool) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomEdit) MaxLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TCustomEdit) SetMaxLength(value int32) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) Modified() bool {
	if !m.IsValid() {
		return false
	}
	r := customEditAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomEdit) SetModified(value bool) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomEdit) NumbersOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := customEditAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomEdit) SetNumbersOnly(value bool) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomEdit) PasswordChar() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(19, 0, m.Instance())
	return uint16(r)
}

func (m *TCustomEdit) SetPasswordChar(value uint16) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := customEditAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomEdit) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomEdit) SelLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TCustomEdit) SetSelLength(value int32) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) SelStart() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customEditAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TCustomEdit) SetSelStart(value int32) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCustomEdit) SelText() string {
	if !m.IsValid() {
		return ""
	}
	r := customEditAPI().SysCallN(23, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomEdit) SetSelText(value string) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(23, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomEdit) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := customEditAPI().SysCallN(24, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomEdit) SetText(value string) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(24, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomEdit) TextHint() string {
	if !m.IsValid() {
		return ""
	}
	r := customEditAPI().SysCallN(25, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomEdit) SetTextHint(value string) {
	if !m.IsValid() {
		return
	}
	customEditAPI().SysCallN(25, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomEdit) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 26, customEditAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomEdit class constructor
func NewCustomEdit(owner IComponent) ICustomEdit {
	r := customEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomEdit(r)
}

func TCustomEditClass() types.TClass {
	r := customEditAPI().SysCallN(27)
	return types.TClass(r)
}

var (
	customEditOnce   base.Once
	customEditImport *imports.Imports = nil
)

func customEditAPI() *imports.Imports {
	customEditOnce.Do(func() {
		customEditImport = api.NewDefaultImports()
		customEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomEdit_Create", 0), // constructor NewCustomEdit
			/* 1 */ imports.NewTable("TCustomEdit_Clear", 0), // procedure Clear
			/* 2 */ imports.NewTable("TCustomEdit_SelectAll", 0), // procedure SelectAll
			/* 3 */ imports.NewTable("TCustomEdit_ClearSelection", 0), // procedure ClearSelection
			/* 4 */ imports.NewTable("TCustomEdit_CopyToClipboard", 0), // procedure CopyToClipboard
			/* 5 */ imports.NewTable("TCustomEdit_CutToClipboard", 0), // procedure CutToClipboard
			/* 6 */ imports.NewTable("TCustomEdit_PasteFromClipboard", 0), // procedure PasteFromClipboard
			/* 7 */ imports.NewTable("TCustomEdit_Undo", 0), // procedure Undo
			/* 8 */ imports.NewTable("TCustomEdit_Alignment", 0), // property Alignment
			/* 9 */ imports.NewTable("TCustomEdit_BorderStyle", 0), // property BorderStyle
			/* 10 */ imports.NewTable("TCustomEdit_CanUndo", 0), // property CanUndo
			/* 11 */ imports.NewTable("TCustomEdit_CaretPos", 0), // property CaretPos
			/* 12 */ imports.NewTable("TCustomEdit_CharCase", 0), // property CharCase
			/* 13 */ imports.NewTable("TCustomEdit_EchoMode", 0), // property EchoMode
			/* 14 */ imports.NewTable("TCustomEdit_EmulatedTextHintStatus", 0), // property EmulatedTextHintStatus
			/* 15 */ imports.NewTable("TCustomEdit_HideSelection", 0), // property HideSelection
			/* 16 */ imports.NewTable("TCustomEdit_MaxLength", 0), // property MaxLength
			/* 17 */ imports.NewTable("TCustomEdit_Modified", 0), // property Modified
			/* 18 */ imports.NewTable("TCustomEdit_NumbersOnly", 0), // property NumbersOnly
			/* 19 */ imports.NewTable("TCustomEdit_PasswordChar", 0), // property PasswordChar
			/* 20 */ imports.NewTable("TCustomEdit_ReadOnly", 0), // property ReadOnly
			/* 21 */ imports.NewTable("TCustomEdit_SelLength", 0), // property SelLength
			/* 22 */ imports.NewTable("TCustomEdit_SelStart", 0), // property SelStart
			/* 23 */ imports.NewTable("TCustomEdit_SelText", 0), // property SelText
			/* 24 */ imports.NewTable("TCustomEdit_Text", 0), // property Text
			/* 25 */ imports.NewTable("TCustomEdit_TextHint", 0), // property TextHint
			/* 26 */ imports.NewTable("TCustomEdit_OnChange", 0), // event OnChange
			/* 27 */ imports.NewTable("TCustomEdit_TClass", 0), // function TCustomEditClass
		}
	})
	return customEditImport
}
