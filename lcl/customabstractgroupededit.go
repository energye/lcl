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

// ICustomAbstractGroupedEdit Parent: ICustomControl
type ICustomAbstractGroupedEdit interface {
	ICustomControl
	Clear()                                        // procedure
	ClearSelection()                               // procedure
	CopyToClipboard()                              // procedure
	CutToClipboard()                               // procedure
	PasteFromClipboard()                           // procedure
	SelectAll()                                    // procedure
	Undo()                                         // procedure
	ValidateEdit()                                 // procedure
	AutoSizeHeightIsEditHeight() bool              // property AutoSizeHeightIsEditHeight Getter
	SetAutoSizeHeightIsEditHeight(value bool)      // property AutoSizeHeightIsEditHeight Setter
	Alignment() types.TAlignment                   // property Alignment Getter
	SetAlignment(value types.TAlignment)           // property Alignment Setter
	CanUndo() bool                                 // property CanUndo Getter
	CaretPos() types.TPoint                        // property CaretPos Getter
	SetCaretPos(value types.TPoint)                // property CaretPos Setter
	CharCase() types.TEditCharCase                 // property CharCase Getter
	SetCharCase(value types.TEditCharCase)         // property CharCase Setter
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	EchoMode() types.TEchoMode                     // property EchoMode Getter
	SetEchoMode(value types.TEchoMode)             // property EchoMode Setter
	HideSelection() bool                           // property HideSelection Getter
	SetHideSelection(value bool)                   // property HideSelection Setter
	MaxLength() int32                              // property MaxLength Getter
	SetMaxLength(value int32)                      // property MaxLength Setter
	Modified() bool                                // property Modified Getter
	SetModified(value bool)                        // property Modified Setter
	NumbersOnly() bool                             // property NumbersOnly Getter
	SetNumbersOnly(value bool)                     // property NumbersOnly Setter
	PasswordChar() uint16                          // property PasswordChar Getter
	SetPasswordChar(value uint16)                  // property PasswordChar Setter
	ReadOnly() bool                                // property ReadOnly Getter
	SetReadOnly(value bool)                        // property ReadOnly Setter
	SelLength() int32                              // property SelLength Getter
	SetSelLength(value int32)                      // property SelLength Setter
	SelStart() int32                               // property SelStart Getter
	SetSelStart(value int32)                       // property SelStart Setter
	SelText() string                               // property SelText Getter
	SetSelText(value string)                       // property SelText Setter
	Text() string                                  // property Text Getter
	SetText(value string)                          // property Text Setter
	TextHint() string                              // property TextHint Getter
	SetTextHint(value string)                      // property TextHint Setter
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

type TCustomAbstractGroupedEdit struct {
	TCustomControl
}

func (m *TCustomAbstractGroupedEdit) Clear() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(0, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ClearSelection() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(1, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CopyToClipboard() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(2, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) CutToClipboard() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(3, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) PasteFromClipboard() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(4, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) SelectAll() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(5, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) Undo() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(6, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) ValidateEdit() {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(7, m.Instance())
}

func (m *TCustomAbstractGroupedEdit) AutoSizeHeightIsEditHeight() bool {
	if !m.IsValid() {
		return false
	}
	r := customAbstractGroupedEditAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAbstractGroupedEdit) SetAutoSizeHeightIsEditHeight(value bool) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAbstractGroupedEdit) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := customAbstractGroupedEditAPI().SysCallN(9, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TCustomAbstractGroupedEdit) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAbstractGroupedEdit) CanUndo() bool {
	if !m.IsValid() {
		return false
	}
	r := customAbstractGroupedEditAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAbstractGroupedEdit) CaretPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomAbstractGroupedEdit) SetCaretPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomAbstractGroupedEdit) CharCase() types.TEditCharCase {
	if !m.IsValid() {
		return 0
	}
	r := customAbstractGroupedEditAPI().SysCallN(12, 0, m.Instance())
	return types.TEditCharCase(r)
}

func (m *TCustomAbstractGroupedEdit) SetCharCase(value types.TEditCharCase) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAbstractGroupedEdit) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := customAbstractGroupedEditAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAbstractGroupedEdit) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAbstractGroupedEdit) EchoMode() types.TEchoMode {
	if !m.IsValid() {
		return 0
	}
	r := customAbstractGroupedEditAPI().SysCallN(14, 0, m.Instance())
	return types.TEchoMode(r)
}

func (m *TCustomAbstractGroupedEdit) SetEchoMode(value types.TEchoMode) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAbstractGroupedEdit) HideSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := customAbstractGroupedEditAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAbstractGroupedEdit) SetHideSelection(value bool) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAbstractGroupedEdit) MaxLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customAbstractGroupedEditAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TCustomAbstractGroupedEdit) SetMaxLength(value int32) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAbstractGroupedEdit) Modified() bool {
	if !m.IsValid() {
		return false
	}
	r := customAbstractGroupedEditAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAbstractGroupedEdit) SetModified(value bool) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAbstractGroupedEdit) NumbersOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := customAbstractGroupedEditAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAbstractGroupedEdit) SetNumbersOnly(value bool) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAbstractGroupedEdit) PasswordChar() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := customAbstractGroupedEditAPI().SysCallN(19, 0, m.Instance())
	return uint16(r)
}

func (m *TCustomAbstractGroupedEdit) SetPasswordChar(value uint16) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAbstractGroupedEdit) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := customAbstractGroupedEditAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAbstractGroupedEdit) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAbstractGroupedEdit) SelLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customAbstractGroupedEditAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TCustomAbstractGroupedEdit) SetSelLength(value int32) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAbstractGroupedEdit) SelStart() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customAbstractGroupedEditAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TCustomAbstractGroupedEdit) SetSelStart(value int32) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAbstractGroupedEdit) SelText() string {
	if !m.IsValid() {
		return ""
	}
	r := customAbstractGroupedEditAPI().SysCallN(23, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomAbstractGroupedEdit) SetSelText(value string) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(23, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomAbstractGroupedEdit) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := customAbstractGroupedEditAPI().SysCallN(24, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomAbstractGroupedEdit) SetText(value string) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(24, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomAbstractGroupedEdit) TextHint() string {
	if !m.IsValid() {
		return ""
	}
	r := customAbstractGroupedEditAPI().SysCallN(25, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomAbstractGroupedEdit) SetTextHint(value string) {
	if !m.IsValid() {
		return
	}
	customAbstractGroupedEditAPI().SysCallN(25, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomAbstractGroupedEdit) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 26, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 27, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 29, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 30, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 31, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 32, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 33, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 34, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 35, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 36, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 37, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 38, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 39, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 40, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomAbstractGroupedEdit) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 41, customAbstractGroupedEditAPI(), api.MakeEventDataPtr(cb))
}

var (
	customAbstractGroupedEditOnce   base.Once
	customAbstractGroupedEditImport *imports.Imports = nil
)

func customAbstractGroupedEditAPI() *imports.Imports {
	customAbstractGroupedEditOnce.Do(func() {
		customAbstractGroupedEditImport = api.NewDefaultImports()
		customAbstractGroupedEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomAbstractGroupedEdit_Clear", 0), // procedure Clear
			/* 1 */ imports.NewTable("TCustomAbstractGroupedEdit_ClearSelection", 0), // procedure ClearSelection
			/* 2 */ imports.NewTable("TCustomAbstractGroupedEdit_CopyToClipboard", 0), // procedure CopyToClipboard
			/* 3 */ imports.NewTable("TCustomAbstractGroupedEdit_CutToClipboard", 0), // procedure CutToClipboard
			/* 4 */ imports.NewTable("TCustomAbstractGroupedEdit_PasteFromClipboard", 0), // procedure PasteFromClipboard
			/* 5 */ imports.NewTable("TCustomAbstractGroupedEdit_SelectAll", 0), // procedure SelectAll
			/* 6 */ imports.NewTable("TCustomAbstractGroupedEdit_Undo", 0), // procedure Undo
			/* 7 */ imports.NewTable("TCustomAbstractGroupedEdit_ValidateEdit", 0), // procedure ValidateEdit
			/* 8 */ imports.NewTable("TCustomAbstractGroupedEdit_AutoSizeHeightIsEditHeight", 0), // property AutoSizeHeightIsEditHeight
			/* 9 */ imports.NewTable("TCustomAbstractGroupedEdit_Alignment", 0), // property Alignment
			/* 10 */ imports.NewTable("TCustomAbstractGroupedEdit_CanUndo", 0), // property CanUndo
			/* 11 */ imports.NewTable("TCustomAbstractGroupedEdit_CaretPos", 0), // property CaretPos
			/* 12 */ imports.NewTable("TCustomAbstractGroupedEdit_CharCase", 0), // property CharCase
			/* 13 */ imports.NewTable("TCustomAbstractGroupedEdit_ParentColor", 0), // property ParentColor
			/* 14 */ imports.NewTable("TCustomAbstractGroupedEdit_EchoMode", 0), // property EchoMode
			/* 15 */ imports.NewTable("TCustomAbstractGroupedEdit_HideSelection", 0), // property HideSelection
			/* 16 */ imports.NewTable("TCustomAbstractGroupedEdit_MaxLength", 0), // property MaxLength
			/* 17 */ imports.NewTable("TCustomAbstractGroupedEdit_Modified", 0), // property Modified
			/* 18 */ imports.NewTable("TCustomAbstractGroupedEdit_NumbersOnly", 0), // property NumbersOnly
			/* 19 */ imports.NewTable("TCustomAbstractGroupedEdit_PasswordChar", 0), // property PasswordChar
			/* 20 */ imports.NewTable("TCustomAbstractGroupedEdit_ReadOnly", 0), // property ReadOnly
			/* 21 */ imports.NewTable("TCustomAbstractGroupedEdit_SelLength", 0), // property SelLength
			/* 22 */ imports.NewTable("TCustomAbstractGroupedEdit_SelStart", 0), // property SelStart
			/* 23 */ imports.NewTable("TCustomAbstractGroupedEdit_SelText", 0), // property SelText
			/* 24 */ imports.NewTable("TCustomAbstractGroupedEdit_Text", 0), // property Text
			/* 25 */ imports.NewTable("TCustomAbstractGroupedEdit_TextHint", 0), // property TextHint
			/* 26 */ imports.NewTable("TCustomAbstractGroupedEdit_OnChange", 0), // event OnChange
			/* 27 */ imports.NewTable("TCustomAbstractGroupedEdit_OnContextPopup", 0), // event OnContextPopup
			/* 28 */ imports.NewTable("TCustomAbstractGroupedEdit_OnDblClick", 0), // event OnDblClick
			/* 29 */ imports.NewTable("TCustomAbstractGroupedEdit_OnDragDrop", 0), // event OnDragDrop
			/* 30 */ imports.NewTable("TCustomAbstractGroupedEdit_OnDragOver", 0), // event OnDragOver
			/* 31 */ imports.NewTable("TCustomAbstractGroupedEdit_OnEditingDone", 0), // event OnEditingDone
			/* 32 */ imports.NewTable("TCustomAbstractGroupedEdit_OnEndDrag", 0), // event OnEndDrag
			/* 33 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseDown", 0), // event OnMouseDown
			/* 34 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseEnter", 0), // event OnMouseEnter
			/* 35 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseLeave", 0), // event OnMouseLeave
			/* 36 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseMove", 0), // event OnMouseMove
			/* 37 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseWheel", 0), // event OnMouseWheel
			/* 38 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 39 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 40 */ imports.NewTable("TCustomAbstractGroupedEdit_OnMouseUp", 0), // event OnMouseUp
			/* 41 */ imports.NewTable("TCustomAbstractGroupedEdit_OnStartDrag", 0), // event OnStartDrag
		}
	})
	return customAbstractGroupedEditImport
}
