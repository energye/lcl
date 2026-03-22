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

// ISynBaseCompletionForm Parent: IForm
type ISynBaseCompletionForm interface {
	IForm
	ShowItemHint(index int32)   // procedure
	OnHintTimer(sender IObject) // procedure
	// CurrentEditor
	//  Must only be set via TSynCompletion.SetEditor
	CurrentEditor() ICustomSynEdit                              // property CurrentEditor Getter
	CurrentString() string                                      // property CurrentString Getter
	SetCurrentString(value string)                              // property CurrentString Setter
	ItemList() IStrings                                         // property ItemList Getter
	SetItemList(value IStrings)                                 // property ItemList Setter
	PositionToInt() int32                                       // property Position Getter
	SetPositionToInt(value int32)                               // property Position Setter
	NbLinesInWindow() int32                                     // property NbLinesInWindow Getter
	SetNbLinesInWindow(value int32)                             // property NbLinesInWindow Setter
	ClSelect() types.TColor                                     // property ClSelect Getter
	SetClSelect(value types.TColor)                             // property ClSelect Setter
	CaseSensitive() bool                                        // property CaseSensitive Getter
	SetCaseSensitive(value bool)                                // property CaseSensitive Setter
	FontHeight() int32                                          // property FontHeight Getter
	BackgroundColor() types.TColor                              // property BackgroundColor Getter
	SetBackgroundColor(value types.TColor)                      // property BackgroundColor Setter
	DrawBorderColor() types.TColor                              // property DrawBorderColor Getter
	SetDrawBorderColor(value types.TColor)                      // property DrawBorderColor Setter
	DrawBorderWidth() int32                                     // property DrawBorderWidth Getter
	SetDrawBorderWidth(value int32)                             // property DrawBorderWidth Setter
	TextColor() types.TColor                                    // property TextColor Getter
	SetTextColor(value types.TColor)                            // property TextColor Setter
	TextSelectedColor() types.TColor                            // property TextSelectedColor Getter
	SetTextSelectedColor(value types.TColor)                    // property TextSelectedColor Setter
	LongLineHintTime() int32                                    // property LongLineHintTime Getter
	SetLongLineHintTime(value int32)                            // property LongLineHintTime Setter
	LongLineHintType() types.TSynCompletionLongHintType         // property LongLineHintType Getter
	SetLongLineHintType(value types.TSynCompletionLongHintType) // property LongLineHintType Setter
	DoubleClickSelects() bool                                   // property DoubleClickSelects Getter
	SetDoubleClickSelects(value bool)                           // property DoubleClickSelects Setter
	ShowSizeDrag() bool                                         // property ShowSizeDrag Getter
	SetShowSizeDrag(value bool)                                 // property ShowSizeDrag Setter
	SetOnKeyDelete(fn TNotifyEvent)                             // property event
	SetOnPaintItem(fn TSynBaseCompletionPaintItem)              // property event
	SetOnMeasureItem(fn TSynBaseCompletionMeasureItem)          // property event
	SetOnValidate(fn TValidateEvent)                            // property event
	SetOnCancel(fn TNotifyEvent)                                // property event
	SetOnSearchPosition(fn TSynBaseCompletionSearchPosition)    // property event
	SetOnKeyCompletePrefix(fn TNotifyEvent)                     // property event
	SetOnKeyNextChar(fn TNotifyEvent)                           // property event
	SetOnKeyPrevChar(fn TNotifyEvent)                           // property event
	SetOnPositionChanged(fn TNotifyEvent)                       // property event
	SetOnDragResized(fn TNotifyEvent)                           // property event
}

type TSynBaseCompletionForm struct {
	TForm
}

func (m *TSynBaseCompletionForm) ShowItemHint(index int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(1, m.Instance(), uintptr(index))
}

func (m *TSynBaseCompletionForm) OnHintTimer(sender IObject) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TSynBaseCompletionForm) CurrentEditor() ICustomSynEdit {
	if !m.IsValid() {
		return nil
	}
	r := synBaseCompletionFormAPI().SysCallN(3, m.Instance())
	return AsCustomSynEdit(r)
}

func (m *TSynBaseCompletionForm) CurrentString() string {
	if !m.IsValid() {
		return ""
	}
	r := synBaseCompletionFormAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TSynBaseCompletionForm) SetCurrentString(value string) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynBaseCompletionForm) ItemList() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synBaseCompletionFormAPI().SysCallN(5, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynBaseCompletionForm) SetItemList(value IStrings) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynBaseCompletionForm) PositionToInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletionForm) SetPositionToInt(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) NbLinesInWindow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletionForm) SetNbLinesInWindow(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) ClSelect() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynBaseCompletionForm) SetClSelect(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) CaseSensitive() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionFormAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletionForm) SetCaseSensitive(value bool) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBaseCompletionForm) FontHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletionForm) BackgroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(11, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynBaseCompletionForm) SetBackgroundColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) DrawBorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(12, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynBaseCompletionForm) SetDrawBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) DrawBorderWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletionForm) SetDrawBorderWidth(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) TextColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(14, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynBaseCompletionForm) SetTextColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) TextSelectedColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(15, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynBaseCompletionForm) SetTextSelectedColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) LongLineHintTime() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TSynBaseCompletionForm) SetLongLineHintTime(value int32) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) LongLineHintType() types.TSynCompletionLongHintType {
	if !m.IsValid() {
		return 0
	}
	r := synBaseCompletionFormAPI().SysCallN(17, 0, m.Instance())
	return types.TSynCompletionLongHintType(r)
}

func (m *TSynBaseCompletionForm) SetLongLineHintType(value types.TSynCompletionLongHintType) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TSynBaseCompletionForm) DoubleClickSelects() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionFormAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletionForm) SetDoubleClickSelects(value bool) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBaseCompletionForm) ShowSizeDrag() bool {
	if !m.IsValid() {
		return false
	}
	r := synBaseCompletionFormAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynBaseCompletionForm) SetShowSizeDrag(value bool) {
	if !m.IsValid() {
		return
	}
	synBaseCompletionFormAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynBaseCompletionForm) SetOnKeyDelete(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnPaintItem(fn TSynBaseCompletionPaintItem) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseCompletionPaintItem(fn)
	base.SetEvent(m, 21, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnMeasureItem(fn TSynBaseCompletionMeasureItem) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseCompletionMeasureItem(fn)
	base.SetEvent(m, 22, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnValidate(fn TValidateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTValidateEvent(fn)
	base.SetEvent(m, 23, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnCancel(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnSearchPosition(fn TSynBaseCompletionSearchPosition) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseCompletionSearchPosition(fn)
	base.SetEvent(m, 25, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnKeyCompletePrefix(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 26, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnKeyNextChar(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 27, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnKeyPrevChar(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnPositionChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 29, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseCompletionForm) SetOnDragResized(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 30, synBaseCompletionFormAPI(), api.MakeEventDataPtr(cb))
}

// NewSynBaseCompletionForm class constructor
func NewSynBaseCompletionForm(owner IComponent) ISynBaseCompletionForm {
	r := synBaseCompletionFormAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynBaseCompletionForm(r)
}

func TSynBaseCompletionFormClass() types.TClass {
	r := synBaseCompletionFormAPI().SysCallN(31)
	return types.TClass(r)
}

var (
	synBaseCompletionFormOnce   base.Once
	synBaseCompletionFormImport *imports.Imports = nil
)

func synBaseCompletionFormAPI() *imports.Imports {
	synBaseCompletionFormOnce.Do(func() {
		synBaseCompletionFormImport = api.NewDefaultImports()
		synBaseCompletionFormImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynBaseCompletionForm_Create", 0), // constructor NewSynBaseCompletionForm
			/* 1 */ imports.NewTable("TSynBaseCompletionForm_ShowItemHint", 0), // procedure ShowItemHint
			/* 2 */ imports.NewTable("TSynBaseCompletionForm_OnHintTimer", 0), // procedure OnHintTimer
			/* 3 */ imports.NewTable("TSynBaseCompletionForm_CurrentEditor", 0), // property CurrentEditor
			/* 4 */ imports.NewTable("TSynBaseCompletionForm_CurrentString", 0), // property CurrentString
			/* 5 */ imports.NewTable("TSynBaseCompletionForm_ItemList", 0), // property ItemList
			/* 6 */ imports.NewTable("TSynBaseCompletionForm_PositionToInt", 0), // property PositionToInt
			/* 7 */ imports.NewTable("TSynBaseCompletionForm_NbLinesInWindow", 0), // property NbLinesInWindow
			/* 8 */ imports.NewTable("TSynBaseCompletionForm_ClSelect", 0), // property ClSelect
			/* 9 */ imports.NewTable("TSynBaseCompletionForm_CaseSensitive", 0), // property CaseSensitive
			/* 10 */ imports.NewTable("TSynBaseCompletionForm_FontHeight", 0), // property FontHeight
			/* 11 */ imports.NewTable("TSynBaseCompletionForm_BackgroundColor", 0), // property BackgroundColor
			/* 12 */ imports.NewTable("TSynBaseCompletionForm_DrawBorderColor", 0), // property DrawBorderColor
			/* 13 */ imports.NewTable("TSynBaseCompletionForm_DrawBorderWidth", 0), // property DrawBorderWidth
			/* 14 */ imports.NewTable("TSynBaseCompletionForm_TextColor", 0), // property TextColor
			/* 15 */ imports.NewTable("TSynBaseCompletionForm_TextSelectedColor", 0), // property TextSelectedColor
			/* 16 */ imports.NewTable("TSynBaseCompletionForm_LongLineHintTime", 0), // property LongLineHintTime
			/* 17 */ imports.NewTable("TSynBaseCompletionForm_LongLineHintType", 0), // property LongLineHintType
			/* 18 */ imports.NewTable("TSynBaseCompletionForm_DoubleClickSelects", 0), // property DoubleClickSelects
			/* 19 */ imports.NewTable("TSynBaseCompletionForm_ShowSizeDrag", 0), // property ShowSizeDrag
			/* 20 */ imports.NewTable("TSynBaseCompletionForm_OnKeyDelete", 0), // event OnKeyDelete
			/* 21 */ imports.NewTable("TSynBaseCompletionForm_OnPaintItem", 0), // event OnPaintItem
			/* 22 */ imports.NewTable("TSynBaseCompletionForm_OnMeasureItem", 0), // event OnMeasureItem
			/* 23 */ imports.NewTable("TSynBaseCompletionForm_OnValidate", 0), // event OnValidate
			/* 24 */ imports.NewTable("TSynBaseCompletionForm_OnCancel", 0), // event OnCancel
			/* 25 */ imports.NewTable("TSynBaseCompletionForm_OnSearchPosition", 0), // event OnSearchPosition
			/* 26 */ imports.NewTable("TSynBaseCompletionForm_OnKeyCompletePrefix", 0), // event OnKeyCompletePrefix
			/* 27 */ imports.NewTable("TSynBaseCompletionForm_OnKeyNextChar", 0), // event OnKeyNextChar
			/* 28 */ imports.NewTable("TSynBaseCompletionForm_OnKeyPrevChar", 0), // event OnKeyPrevChar
			/* 29 */ imports.NewTable("TSynBaseCompletionForm_OnPositionChanged", 0), // event OnPositionChanged
			/* 30 */ imports.NewTable("TSynBaseCompletionForm_OnDragResized", 0), // event OnDragResized
			/* 31 */ imports.NewTable("TSynBaseCompletionForm_TClass", 0), // function TSynBaseCompletionFormClass
		}
	})
	return synBaseCompletionFormImport
}
