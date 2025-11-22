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

// ICustomControlFilterEdit Parent: ICustomEditButton
type ICustomControlFilterEdit interface {
	ICustomEditButton
	ForceFilter(filter string) string                  // function
	InvalidateFilter()                                 // procedure
	ResetFilter()                                      // procedure
	StoreSelection()                                   // procedure
	RestoreSelection()                                 // procedure
	Filter() string                                    // property Filter Getter
	SetFilter(value string)                            // property Filter Setter
	FilterLowercase() string                           // property FilterLowercase Getter
	IdleConnected() bool                               // property IdleConnected Getter
	SetIdleConnected(value bool)                       // property IdleConnected Setter
	SortData() bool                                    // property SortData Getter
	SetSortData(value bool)                            // property SortData Setter
	SelectedPart() IObject                             // property SelectedPart Getter
	SetSelectedPart(value IObject)                     // property SelectedPart Setter
	FilterOptions() types.TFilterStringOptions         // property FilterOptions Getter
	SetFilterOptions(value types.TFilterStringOptions) // property FilterOptions Setter
	ButtonCaption() string                             // property ButtonCaption Getter
	SetButtonCaption(value string)                     // property ButtonCaption Setter
	ButtonCursor() types.TCursor                       // property ButtonCursor Getter
	SetButtonCursor(value types.TCursor)               // property ButtonCursor Setter
	ButtonHint() string                                // property ButtonHint Getter
	SetButtonHint(value string)                        // property ButtonHint Setter
	ButtonOnlyWhenFocused() bool                       // property ButtonOnlyWhenFocused Getter
	SetButtonOnlyWhenFocused(value bool)               // property ButtonOnlyWhenFocused Setter
	ButtonWidth() int32                                // property ButtonWidth Getter
	SetButtonWidth(value int32)                        // property ButtonWidth Setter
	DirectInput() bool                                 // property DirectInput Getter
	SetDirectInput(value bool)                         // property DirectInput Setter
	Flat() bool                                        // property Flat Getter
	SetFlat(value bool)                                // property Flat Setter
	FocusOnButtonClick() bool                          // property FocusOnButtonClick Getter
	SetFocusOnButtonClick(value bool)                  // property FocusOnButtonClick Setter
	AutoSelect() bool                                  // property AutoSelect Getter
	SetAutoSelect(value bool)                          // property AutoSelect Setter
	DragCursor() types.TCursor                         // property DragCursor Getter
	SetDragCursor(value types.TCursor)                 // property DragCursor Setter
	DragMode() types.TDragMode                         // property DragMode Getter
	SetDragMode(value types.TDragMode)                 // property DragMode Setter
	Glyph() IBitmap                                    // property Glyph Getter
	SetGlyph(value IBitmap)                            // property Glyph Setter
	NumGlyphs() int32                                  // property NumGlyphs Getter
	SetNumGlyphs(value int32)                          // property NumGlyphs Setter
	Images() ICustomImageList                          // property Images Getter
	SetImages(value ICustomImageList)                  // property Images Setter
	ImageIndex() int32                                 // property ImageIndex Getter
	SetImageIndex(value int32)                         // property ImageIndex Setter
	ImageWidth() int32                                 // property ImageWidth Getter
	SetImageWidth(value int32)                         // property ImageWidth Setter
	Layout() types.TLeftRight                          // property Layout Getter
	SetLayout(value types.TLeftRight)                  // property Layout Setter
	ParentFont() bool                                  // property ParentFont Getter
	SetParentFont(value bool)                          // property ParentFont Setter
	ParentShowHint() bool                              // property ParentShowHint Getter
	SetParentShowHint(value bool)                      // property ParentShowHint Setter
	Spacing() int32                                    // property Spacing Getter
	SetSpacing(value int32)                            // property Spacing Setter
	SetOnAfterFilter(fn TNotifyEvent)                  // property event
	SetOnFilterItem(fn TFilterItemEvent)               // property event
	SetOnFilterItemEx(fn TFilterItemExEvent)           // property event
	SetOnCheckItem(fn TCheckItemEvent)                 // property event
	SetOnButtonClick(fn TNotifyEvent)                  // property event
}

type TCustomControlFilterEdit struct {
	TCustomEditButton
}

func (m *TCustomControlFilterEdit) ForceFilter(filter string) string {
	if !m.IsValid() {
		return ""
	}
	r := customControlFilterEditAPI().SysCallN(0, m.Instance(), api.PasStr(filter))
	return api.GoStr(r)
}

func (m *TCustomControlFilterEdit) InvalidateFilter() {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(1, m.Instance())
}

func (m *TCustomControlFilterEdit) ResetFilter() {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(2, m.Instance())
}

func (m *TCustomControlFilterEdit) StoreSelection() {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(3, m.Instance())
}

func (m *TCustomControlFilterEdit) RestoreSelection() {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(4, m.Instance())
}

func (m *TCustomControlFilterEdit) Filter() string {
	if !m.IsValid() {
		return ""
	}
	r := customControlFilterEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomControlFilterEdit) SetFilter(value string) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomControlFilterEdit) FilterLowercase() string {
	if !m.IsValid() {
		return ""
	}
	r := customControlFilterEditAPI().SysCallN(6, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomControlFilterEdit) IdleConnected() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetIdleConnected(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) SortData() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetSortData(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) SelectedPart() IObject {
	if !m.IsValid() {
		return nil
	}
	r := customControlFilterEditAPI().SysCallN(9, 0, m.Instance())
	return AsObject(r)
}

func (m *TCustomControlFilterEdit) SetSelectedPart(value IObject) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomControlFilterEdit) FilterOptions() types.TFilterStringOptions {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(10, 0, m.Instance())
	return types.TFilterStringOptions(r)
}

func (m *TCustomControlFilterEdit) SetFilterOptions(value types.TFilterStringOptions) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) ButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := customControlFilterEditAPI().SysCallN(11, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomControlFilterEdit) SetButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(11, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomControlFilterEdit) ButtonCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(12, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCustomControlFilterEdit) SetButtonCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) ButtonHint() string {
	if !m.IsValid() {
		return ""
	}
	r := customControlFilterEditAPI().SysCallN(13, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomControlFilterEdit) SetButtonHint(value string) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomControlFilterEdit) ButtonOnlyWhenFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetButtonOnlyWhenFocused(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) ButtonWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TCustomControlFilterEdit) SetButtonWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) DirectInput() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetDirectInput(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) FocusOnButtonClick() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetFocusOnButtonClick(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(20, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCustomControlFilterEdit) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(21, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCustomControlFilterEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := customControlFilterEditAPI().SysCallN(22, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TCustomControlFilterEdit) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomControlFilterEdit) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TCustomControlFilterEdit) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customControlFilterEditAPI().SysCallN(24, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomControlFilterEdit) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(24, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomControlFilterEdit) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(25, 0, m.Instance())
	return int32(r)
}

func (m *TCustomControlFilterEdit) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(26, 0, m.Instance())
	return int32(r)
}

func (m *TCustomControlFilterEdit) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) Layout() types.TLeftRight {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(27, 0, m.Instance())
	return types.TLeftRight(r)
}

func (m *TCustomControlFilterEdit) SetLayout(value types.TLeftRight) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(28, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(28, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := customControlFilterEditAPI().SysCallN(29, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomControlFilterEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(29, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomControlFilterEdit) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customControlFilterEditAPI().SysCallN(30, 0, m.Instance())
	return int32(r)
}

func (m *TCustomControlFilterEdit) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	customControlFilterEditAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TCustomControlFilterEdit) SetOnAfterFilter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 31, customControlFilterEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlFilterEdit) SetOnFilterItem(fn TFilterItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTFilterItemEvent(fn)
	base.SetEvent(m, 32, customControlFilterEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlFilterEdit) SetOnFilterItemEx(fn TFilterItemExEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTFilterItemExEvent(fn)
	base.SetEvent(m, 33, customControlFilterEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlFilterEdit) SetOnCheckItem(fn TCheckItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCheckItemEvent(fn)
	base.SetEvent(m, 34, customControlFilterEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomControlFilterEdit) SetOnButtonClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 35, customControlFilterEditAPI(), api.MakeEventDataPtr(cb))
}

var (
	customControlFilterEditOnce   base.Once
	customControlFilterEditImport *imports.Imports = nil
)

func customControlFilterEditAPI() *imports.Imports {
	customControlFilterEditOnce.Do(func() {
		customControlFilterEditImport = api.NewDefaultImports()
		customControlFilterEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomControlFilterEdit_ForceFilter", 0), // function ForceFilter
			/* 1 */ imports.NewTable("TCustomControlFilterEdit_InvalidateFilter", 0), // procedure InvalidateFilter
			/* 2 */ imports.NewTable("TCustomControlFilterEdit_ResetFilter", 0), // procedure ResetFilter
			/* 3 */ imports.NewTable("TCustomControlFilterEdit_StoreSelection", 0), // procedure StoreSelection
			/* 4 */ imports.NewTable("TCustomControlFilterEdit_RestoreSelection", 0), // procedure RestoreSelection
			/* 5 */ imports.NewTable("TCustomControlFilterEdit_Filter", 0), // property Filter
			/* 6 */ imports.NewTable("TCustomControlFilterEdit_FilterLowercase", 0), // property FilterLowercase
			/* 7 */ imports.NewTable("TCustomControlFilterEdit_IdleConnected", 0), // property IdleConnected
			/* 8 */ imports.NewTable("TCustomControlFilterEdit_SortData", 0), // property SortData
			/* 9 */ imports.NewTable("TCustomControlFilterEdit_SelectedPart", 0), // property SelectedPart
			/* 10 */ imports.NewTable("TCustomControlFilterEdit_FilterOptions", 0), // property FilterOptions
			/* 11 */ imports.NewTable("TCustomControlFilterEdit_ButtonCaption", 0), // property ButtonCaption
			/* 12 */ imports.NewTable("TCustomControlFilterEdit_ButtonCursor", 0), // property ButtonCursor
			/* 13 */ imports.NewTable("TCustomControlFilterEdit_ButtonHint", 0), // property ButtonHint
			/* 14 */ imports.NewTable("TCustomControlFilterEdit_ButtonOnlyWhenFocused", 0), // property ButtonOnlyWhenFocused
			/* 15 */ imports.NewTable("TCustomControlFilterEdit_ButtonWidth", 0), // property ButtonWidth
			/* 16 */ imports.NewTable("TCustomControlFilterEdit_DirectInput", 0), // property DirectInput
			/* 17 */ imports.NewTable("TCustomControlFilterEdit_Flat", 0), // property Flat
			/* 18 */ imports.NewTable("TCustomControlFilterEdit_FocusOnButtonClick", 0), // property FocusOnButtonClick
			/* 19 */ imports.NewTable("TCustomControlFilterEdit_AutoSelect", 0), // property AutoSelect
			/* 20 */ imports.NewTable("TCustomControlFilterEdit_DragCursor", 0), // property DragCursor
			/* 21 */ imports.NewTable("TCustomControlFilterEdit_DragMode", 0), // property DragMode
			/* 22 */ imports.NewTable("TCustomControlFilterEdit_Glyph", 0), // property Glyph
			/* 23 */ imports.NewTable("TCustomControlFilterEdit_NumGlyphs", 0), // property NumGlyphs
			/* 24 */ imports.NewTable("TCustomControlFilterEdit_Images", 0), // property Images
			/* 25 */ imports.NewTable("TCustomControlFilterEdit_ImageIndex", 0), // property ImageIndex
			/* 26 */ imports.NewTable("TCustomControlFilterEdit_ImageWidth", 0), // property ImageWidth
			/* 27 */ imports.NewTable("TCustomControlFilterEdit_Layout", 0), // property Layout
			/* 28 */ imports.NewTable("TCustomControlFilterEdit_ParentFont", 0), // property ParentFont
			/* 29 */ imports.NewTable("TCustomControlFilterEdit_ParentShowHint", 0), // property ParentShowHint
			/* 30 */ imports.NewTable("TCustomControlFilterEdit_Spacing", 0), // property Spacing
			/* 31 */ imports.NewTable("TCustomControlFilterEdit_OnAfterFilter", 0), // event OnAfterFilter
			/* 32 */ imports.NewTable("TCustomControlFilterEdit_OnFilterItem", 0), // event OnFilterItem
			/* 33 */ imports.NewTable("TCustomControlFilterEdit_OnFilterItemEx", 0), // event OnFilterItemEx
			/* 34 */ imports.NewTable("TCustomControlFilterEdit_OnCheckItem", 0), // event OnCheckItem
			/* 35 */ imports.NewTable("TCustomControlFilterEdit_OnButtonClick", 0), // event OnButtonClick
		}
	})
	return customControlFilterEditImport
}
