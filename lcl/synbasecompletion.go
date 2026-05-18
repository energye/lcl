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

// ISynBaseCompletion Parent: ILazSynMultiEditPlugin
type ISynBaseCompletion interface {
	ILazSynMultiEditPlugin
	IsActive() bool                                             // function
	TheForm() ISynBaseCompletionForm                            // function
	ExecuteWithStrIntX2(S string, X int32, Y int32)             // procedure
	ExecuteWithStrPoint(S string, topLeft types.TPoint)         // procedure
	ExecuteWithStrRect(S string, tokenRect types.TRect)         // procedure
	Deactivate()                                                // procedure
	CurrentString() string                                      // property CurrentString Getter
	SetCurrentString(value string)                              // property CurrentString Setter
	FontHeight() int32                                          // property FontHeight Getter
	ItemList() IStrings                                         // property ItemList Getter
	SetItemList(value IStrings)                                 // property ItemList Setter
	Position() int32                                            // property Position Getter
	SetPosition(value int32)                                    // property Position Setter
	LinesInWindow() int32                                       // property LinesInWindow Getter
	SetLinesInWindow(value int32)                               // property LinesInWindow Setter
	SelectedColor() types.TColor                                // property SelectedColor Getter
	SetSelectedColor(value types.TColor)                        // property SelectedColor Setter
	CaseSensitive() bool                                        // property CaseSensitive Getter
	SetCaseSensitive(value bool)                                // property CaseSensitive Setter
	Width() int32                                               // property Width Getter
	SetWidth(value int32)                                       // property Width Setter
	LongLineHintTime() int32                                    // property LongLineHintTime Getter
	SetLongLineHintTime(value int32)                            // property LongLineHintTime Setter
	LongLineHintType() types.TSynCompletionLongHintType         // property LongLineHintType Getter
	SetLongLineHintType(value types.TSynCompletionLongHintType) // property LongLineHintType Setter
	DoubleClickSelects() bool                                   // property DoubleClickSelects Getter
	SetDoubleClickSelects(value bool)                           // property DoubleClickSelects Setter
	ShowSizeDrag() bool                                         // property ShowSizeDrag Getter
	SetShowSizeDrag(value bool)                                 // property ShowSizeDrag Setter
	AutoUseSingleIdent() bool                                   // property AutoUseSingleIdent Getter
	SetAutoUseSingleIdent(value bool)                           // property AutoUseSingleIdent Setter
	SetOnKeyDown(fn TKeyEvent)                                  // property event
	SetOnUTF8KeyPress(fn TUTF8KeyPressEvent)                    // property event
	SetOnKeyPress(fn TKeyPressEvent)                            // property event
	SetOnKeyDelete(fn TNotifyEvent)                             // property event
	SetOnValidate(fn TValidateEvent)                            // property event
	SetOnCancel(fn TNotifyEvent)                                // property event
	SetOnBeforeExecute(fn TOnBeforeExecuteEvent)                // property event
	SetOnExecute(fn TNotifyEvent)                               // property event
	SetOnPaintItem(fn TSynBaseCompletionPaintItem)              // property event
	SetOnMeasureItem(fn TSynBaseCompletionMeasureItem)          // property event
	SetOnSearchPosition(fn TSynBaseCompletionSearchPosition)    // property event
	SetOnKeyCompletePrefix(fn TNotifyEvent)                     // property event
	SetOnKeyNextChar(fn TNotifyEvent)                           // property event
	SetOnKeyPrevChar(fn TNotifyEvent)                           // property event
	SetOnPositionChanged(fn TNotifyEvent)                       // property event
}

type TSynBaseCompletion struct {
	TLazSynMultiEditPlugin
}

func (m *TSynBaseCompletion) IsActive() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletion) TheForm() ISynBaseCompletionForm {
	if !m.IsValid() {
		return nil
	}
	r := synBaseCompletionAPI().SysCallN(2, m.Instance())
	return AsSynBaseCompletionForm(r)
}

func (m *TSynBaseCompletion) ExecuteWithStrIntX2(S string, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(3, m.Instance(), api.PasStr(S), uintptr(X), uintptr(Y))
}

func (m *TSynBaseCompletion) ExecuteWithStrPoint(S string, topLeft types.TPoint) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(4, m.Instance(), api.PasStr(S), uintptr(base.UnsafePointer(&topLeft)))
}

func (m *TSynBaseCompletion) ExecuteWithStrRect(S string, tokenRect types.TRect) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(5, m.Instance(), api.PasStr(S), uintptr(base.UnsafePointer(&tokenRect)))
}

func (m *TSynBaseCompletion) Deactivate() {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(6, m.Instance())
}

func (m *TSynBaseCompletion) CurrentString() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synBaseCompletionAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynBaseCompletion) SetCurrentString(value string) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynBaseCompletion) FontHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletion) ItemList() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synBaseCompletionAPI().SysCallN(9, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynBaseCompletion) SetItemList(value IStrings) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynBaseCompletion) Position() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletion) SetPosition(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletion) LinesInWindow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletion) SetLinesInWindow(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletion) SelectedColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionAPI().SysCallN(12, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynBaseCompletion) SetSelectedColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletion) CaseSensitive() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletion) SetCaseSensitive(value bool) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBaseCompletion) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletion) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletion) LongLineHintTime() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletion) SetLongLineHintTime(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletion) LongLineHintType() types.TSynCompletionLongHintType {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionAPI().SysCallN(16, 0, m.Instance())
	return types.TSynCompletionLongHintType(r)
}

func (m *TSynBaseCompletion) SetLongLineHintType(value types.TSynCompletionLongHintType) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletion) DoubleClickSelects() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletion) SetDoubleClickSelects(value bool) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBaseCompletion) ShowSizeDrag() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletion) SetShowSizeDrag(value bool) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBaseCompletion) AutoUseSingleIdent() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletion) SetAutoUseSingleIdent(value bool) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBaseCompletion) SetOnKeyDown(fn TKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTKeyEvent(fn)
	base.SetEvent(m, 20, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnUTF8KeyPress(fn TUTF8KeyPressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUTF8KeyPressEvent(fn)
	base.SetEvent(m, 21, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnKeyPress(fn TKeyPressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTKeyPressEvent(fn)
	base.SetEvent(m, 22, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnKeyDelete(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 23, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnValidate(fn TValidateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTValidateEvent(fn)
	base.SetEvent(m, 24, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnCancel(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 25, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnBeforeExecute(fn TOnBeforeExecuteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnBeforeExecuteEvent(fn)
	base.SetEvent(m, 26, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnExecute(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 27, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnPaintItem(fn TSynBaseCompletionPaintItem) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseCompletionPaintItem(fn)
	base.SetEvent(m, 28, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnMeasureItem(fn TSynBaseCompletionMeasureItem) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseCompletionMeasureItem(fn)
	base.SetEvent(m, 29, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnSearchPosition(fn TSynBaseCompletionSearchPosition) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseCompletionSearchPosition(fn)
	base.SetEvent(m, 30, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnKeyCompletePrefix(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 31, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnKeyNextChar(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 32, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnKeyPrevChar(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 33, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletion) SetOnPositionChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 34, synBaseCompletionAPI(), api.MakeEventDataPtr(cb))
}

// NewSynBaseCompletion class constructor
func NewSynBaseCompletion(owner IComponent) ISynBaseCompletion {
	r := synBaseCompletionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynBaseCompletion(r)
}

func TSynBaseCompletionClass() types.TClass {
	r := synBaseCompletionAPI().SysCallN(35)
	return types.TClass(r)
}

var (
	synBaseCompletionOnce   base.Once
	synBaseCompletionImport *imports.Imports = nil
)

func synBaseCompletionAPI() *imports.Imports {
	synBaseCompletionOnce.Do(func() {
		synBaseCompletionImport = api.NewDefaultImports()
		synBaseCompletionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynBaseCompletion_Create", 0), // constructor NewSynBaseCompletion
			/* 1 */ imports.NewTable("TSynBaseCompletion_IsActive", 0), // function IsActive
			/* 2 */ imports.NewTable("TSynBaseCompletion_TheForm", 0), // function TheForm
			/* 3 */ imports.NewTable("TSynBaseCompletion_ExecuteWithStrIntX2", 0), // procedure ExecuteWithStrIntX2
			/* 4 */ imports.NewTable("TSynBaseCompletion_ExecuteWithStrPoint", 0), // procedure ExecuteWithStrPoint
			/* 5 */ imports.NewTable("TSynBaseCompletion_ExecuteWithStrRect", 0), // procedure ExecuteWithStrRect
			/* 6 */ imports.NewTable("TSynBaseCompletion_Deactivate", 0), // procedure Deactivate
			/* 7 */ imports.NewTable("TSynBaseCompletion_CurrentString", 0), // property CurrentString
			/* 8 */ imports.NewTable("TSynBaseCompletion_FontHeight", 0), // property FontHeight
			/* 9 */ imports.NewTable("TSynBaseCompletion_ItemList", 0), // property ItemList
			/* 10 */ imports.NewTable("TSynBaseCompletion_Position", 0), // property Position
			/* 11 */ imports.NewTable("TSynBaseCompletion_LinesInWindow", 0), // property LinesInWindow
			/* 12 */ imports.NewTable("TSynBaseCompletion_SelectedColor", 0), // property SelectedColor
			/* 13 */ imports.NewTable("TSynBaseCompletion_CaseSensitive", 0), // property CaseSensitive
			/* 14 */ imports.NewTable("TSynBaseCompletion_Width", 0), // property Width
			/* 15 */ imports.NewTable("TSynBaseCompletion_LongLineHintTime", 0), // property LongLineHintTime
			/* 16 */ imports.NewTable("TSynBaseCompletion_LongLineHintType", 0), // property LongLineHintType
			/* 17 */ imports.NewTable("TSynBaseCompletion_DoubleClickSelects", 0), // property DoubleClickSelects
			/* 18 */ imports.NewTable("TSynBaseCompletion_ShowSizeDrag", 0), // property ShowSizeDrag
			/* 19 */ imports.NewTable("TSynBaseCompletion_AutoUseSingleIdent", 0), // property AutoUseSingleIdent
			/* 20 */ imports.NewTable("TSynBaseCompletion_OnKeyDown", 0), // event OnKeyDown
			/* 21 */ imports.NewTable("TSynBaseCompletion_OnUTF8KeyPress", 0), // event OnUTF8KeyPress
			/* 22 */ imports.NewTable("TSynBaseCompletion_OnKeyPress", 0), // event OnKeyPress
			/* 23 */ imports.NewTable("TSynBaseCompletion_OnKeyDelete", 0), // event OnKeyDelete
			/* 24 */ imports.NewTable("TSynBaseCompletion_OnValidate", 0), // event OnValidate
			/* 25 */ imports.NewTable("TSynBaseCompletion_OnCancel", 0), // event OnCancel
			/* 26 */ imports.NewTable("TSynBaseCompletion_OnBeforeExecute", 0), // event OnBeforeExecute
			/* 27 */ imports.NewTable("TSynBaseCompletion_OnExecute", 0), // event OnExecute
			/* 28 */ imports.NewTable("TSynBaseCompletion_OnPaintItem", 0), // event OnPaintItem
			/* 29 */ imports.NewTable("TSynBaseCompletion_OnMeasureItem", 0), // event OnMeasureItem
			/* 30 */ imports.NewTable("TSynBaseCompletion_OnSearchPosition", 0), // event OnSearchPosition
			/* 31 */ imports.NewTable("TSynBaseCompletion_OnKeyCompletePrefix", 0), // event OnKeyCompletePrefix
			/* 32 */ imports.NewTable("TSynBaseCompletion_OnKeyNextChar", 0), // event OnKeyNextChar
			/* 33 */ imports.NewTable("TSynBaseCompletion_OnKeyPrevChar", 0), // event OnKeyPrevChar
			/* 34 */ imports.NewTable("TSynBaseCompletion_OnPositionChanged", 0), // event OnPositionChanged
			/* 35 */ imports.NewTable("TSynBaseCompletion_TClass", 0), // function TSynBaseCompletionClass
		}
	})
	return synBaseCompletionImport
}
