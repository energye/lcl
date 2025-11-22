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

// IStringGrid Parent: ICustomStringGrid
type IStringGrid interface {
	ICustomStringGrid
	Modified() bool                                      // property Modified Getter
	SetModified(value bool)                              // property Modified Setter
	InplaceEditor() IWinControl                          // property InplaceEditor Getter
	AlternateColor() types.TColor                        // property AlternateColor Getter
	SetAlternateColor(value types.TColor)                // property AlternateColor Setter
	AutoEdit() bool                                      // property AutoEdit Getter
	SetAutoEdit(value bool)                              // property AutoEdit Setter
	CellHintPriority() types.TCellHintPriority           // property CellHintPriority Getter
	SetCellHintPriority(value types.TCellHintPriority)   // property CellHintPriority Setter
	ColRowDraggingCursor() types.TCursor                 // property ColRowDraggingCursor Getter
	SetColRowDraggingCursor(value types.TCursor)         // property ColRowDraggingCursor Setter
	ColRowDragIndicatorColor() types.TColor              // property ColRowDragIndicatorColor Getter
	SetColRowDragIndicatorColor(value types.TColor)      // property ColRowDragIndicatorColor Setter
	ColSizingCursor() types.TCursor                      // property ColSizingCursor Getter
	SetColSizingCursor(value types.TCursor)              // property ColSizingCursor Setter
	ColumnClickSorts() bool                              // property ColumnClickSorts Getter
	SetColumnClickSorts(value bool)                      // property ColumnClickSorts Setter
	DragCursor() types.TCursor                           // property DragCursor Getter
	SetDragCursor(value types.TCursor)                   // property DragCursor Setter
	DragKind() types.TDragKind                           // property DragKind Getter
	SetDragKind(value types.TDragKind)                   // property DragKind Setter
	DragMode() types.TDragMode                           // property DragMode Getter
	SetDragMode(value types.TDragMode)                   // property DragMode Setter
	HeaderHotZones() types.TGridZoneSet                  // property HeaderHotZones Getter
	SetHeaderHotZones(value types.TGridZoneSet)          // property HeaderHotZones Setter
	HeaderPushZones() types.TGridZoneSet                 // property HeaderPushZones Getter
	SetHeaderPushZones(value types.TGridZoneSet)         // property HeaderPushZones Setter
	ImageIndexSortAsc() int32                            // property ImageIndexSortAsc Getter
	SetImageIndexSortAsc(value int32)                    // property ImageIndexSortAsc Setter
	ImageIndexSortDesc() int32                           // property ImageIndexSortDesc Getter
	SetImageIndexSortDesc(value int32)                   // property ImageIndexSortDesc Setter
	MouseWheelOption() types.TMouseWheelOption           // property MouseWheelOption Getter
	SetMouseWheelOption(value types.TMouseWheelOption)   // property MouseWheelOption Setter
	ParentColor() bool                                   // property ParentColor Getter
	SetParentColor(value bool)                           // property ParentColor Setter
	ParentFont() bool                                    // property ParentFont Getter
	SetParentFont(value bool)                            // property ParentFont Setter
	RangeSelectMode() types.TRangeSelectMode             // property RangeSelectMode Getter
	SetRangeSelectMode(value types.TRangeSelectMode)     // property RangeSelectMode Setter
	RowSizingCursor() types.TCursor                      // property RowSizingCursor Getter
	SetRowSizingCursor(value types.TCursor)              // property RowSizingCursor Setter
	TitleFont() IFont                                    // property TitleFont Getter
	SetTitleFont(value IFont)                            // property TitleFont Setter
	TitleImageList() ICustomImageList                    // property TitleImageList Getter
	SetTitleImageList(value ICustomImageList)            // property TitleImageList Setter
	TitleStyle() types.TTitleStyle                       // property TitleStyle Getter
	SetTitleStyle(value types.TTitleStyle)               // property TitleStyle Setter
	SetOnCellProcess(fn TCellProcessEvent)               // property event
	SetOnCheckboxToggled(fn TToggledCheckboxEvent)       // property event
	SetOnEditingDone(fn TNotifyEvent)                    // property event
	SetOnGetCellHint(fn TGetCellHintEvent)               // property event
	SetOnGetCheckboxState(fn TGetCheckboxStateEvent)     // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)      // property event
	SetOnSetCheckboxState(fn TSetCheckboxStateEvent)     // property event
	SetOnUserCheckboxBitmap(fn TUserCheckboxBitmapEvent) // property event
	SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent)   // property event
}

type TStringGrid struct {
	TCustomStringGrid
}

func (m *TStringGrid) Modified() bool {
	if !m.IsValid() {
		return false
	}
	r := stringGridAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringGrid) SetModified(value bool) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringGrid) InplaceEditor() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := stringGridAPI().SysCallN(2, m.Instance())
	return AsWinControl(r)
}

func (m *TStringGrid) AlternateColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(3, 0, m.Instance())
	return types.TColor(r)
}

func (m *TStringGrid) SetAlternateColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) AutoEdit() bool {
	if !m.IsValid() {
		return false
	}
	r := stringGridAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringGrid) SetAutoEdit(value bool) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringGrid) CellHintPriority() types.TCellHintPriority {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(5, 0, m.Instance())
	return types.TCellHintPriority(r)
}

func (m *TStringGrid) SetCellHintPriority(value types.TCellHintPriority) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) ColRowDraggingCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(6, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TStringGrid) SetColRowDraggingCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) ColRowDragIndicatorColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(7, 0, m.Instance())
	return types.TColor(r)
}

func (m *TStringGrid) SetColRowDragIndicatorColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) ColSizingCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(8, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TStringGrid) SetColSizingCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) ColumnClickSorts() bool {
	if !m.IsValid() {
		return false
	}
	r := stringGridAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringGrid) SetColumnClickSorts(value bool) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringGrid) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(10, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TStringGrid) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(11, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TStringGrid) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(12, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TStringGrid) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) HeaderHotZones() types.TGridZoneSet {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(13, 0, m.Instance())
	return types.TGridZoneSet(r)
}

func (m *TStringGrid) SetHeaderHotZones(value types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) HeaderPushZones() types.TGridZoneSet {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(14, 0, m.Instance())
	return types.TGridZoneSet(r)
}

func (m *TStringGrid) SetHeaderPushZones(value types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) ImageIndexSortAsc() int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TStringGrid) SetImageIndexSortAsc(value int32) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) ImageIndexSortDesc() int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TStringGrid) SetImageIndexSortDesc(value int32) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) MouseWheelOption() types.TMouseWheelOption {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(17, 0, m.Instance())
	return types.TMouseWheelOption(r)
}

func (m *TStringGrid) SetMouseWheelOption(value types.TMouseWheelOption) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := stringGridAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringGrid) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringGrid) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := stringGridAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringGrid) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringGrid) RangeSelectMode() types.TRangeSelectMode {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(20, 0, m.Instance())
	return types.TRangeSelectMode(r)
}

func (m *TStringGrid) SetRangeSelectMode(value types.TRangeSelectMode) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) RowSizingCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(21, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TStringGrid) SetRowSizingCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) TitleFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := stringGridAPI().SysCallN(22, 0, m.Instance())
	return AsFont(r)
}

func (m *TStringGrid) SetTitleFont(value IFont) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TStringGrid) TitleImageList() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := stringGridAPI().SysCallN(23, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TStringGrid) SetTitleImageList(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(23, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TStringGrid) TitleStyle() types.TTitleStyle {
	if !m.IsValid() {
		return 0
	}
	r := stringGridAPI().SysCallN(24, 0, m.Instance())
	return types.TTitleStyle(r)
}

func (m *TStringGrid) SetTitleStyle(value types.TTitleStyle) {
	if !m.IsValid() {
		return
	}
	stringGridAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TStringGrid) SetOnCellProcess(fn TCellProcessEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCellProcessEvent(fn)
	base.SetEvent(m, 25, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTToggledCheckboxEvent(fn)
	base.SetEvent(m, 26, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 27, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnGetCellHint(fn TGetCellHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetCellHintEvent(fn)
	base.SetEvent(m, 28, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnGetCheckboxState(fn TGetCheckboxStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetCheckboxStateEvent(fn)
	base.SetEvent(m, 29, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 30, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 31, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 32, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnSetCheckboxState(fn TSetCheckboxStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSetCheckboxStateEvent(fn)
	base.SetEvent(m, 33, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnUserCheckboxBitmap(fn TUserCheckboxBitmapEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUserCheckboxBitmapEvent(fn)
	base.SetEvent(m, 34, stringGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringGrid) SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUserCheckBoxImageEvent(fn)
	base.SetEvent(m, 35, stringGridAPI(), api.MakeEventDataPtr(cb))
}

// NewStringGrid class constructor
func NewStringGrid(owner IComponent) IStringGrid {
	r := stringGridAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsStringGrid(r)
}

func TStringGridClass() types.TClass {
	r := stringGridAPI().SysCallN(36)
	return types.TClass(r)
}

var (
	stringGridOnce   base.Once
	stringGridImport *imports.Imports = nil
)

func stringGridAPI() *imports.Imports {
	stringGridOnce.Do(func() {
		stringGridImport = api.NewDefaultImports()
		stringGridImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStringGrid_Create", 0), // constructor NewStringGrid
			/* 1 */ imports.NewTable("TStringGrid_Modified", 0), // property Modified
			/* 2 */ imports.NewTable("TStringGrid_InplaceEditor", 0), // property InplaceEditor
			/* 3 */ imports.NewTable("TStringGrid_AlternateColor", 0), // property AlternateColor
			/* 4 */ imports.NewTable("TStringGrid_AutoEdit", 0), // property AutoEdit
			/* 5 */ imports.NewTable("TStringGrid_CellHintPriority", 0), // property CellHintPriority
			/* 6 */ imports.NewTable("TStringGrid_ColRowDraggingCursor", 0), // property ColRowDraggingCursor
			/* 7 */ imports.NewTable("TStringGrid_ColRowDragIndicatorColor", 0), // property ColRowDragIndicatorColor
			/* 8 */ imports.NewTable("TStringGrid_ColSizingCursor", 0), // property ColSizingCursor
			/* 9 */ imports.NewTable("TStringGrid_ColumnClickSorts", 0), // property ColumnClickSorts
			/* 10 */ imports.NewTable("TStringGrid_DragCursor", 0), // property DragCursor
			/* 11 */ imports.NewTable("TStringGrid_DragKind", 0), // property DragKind
			/* 12 */ imports.NewTable("TStringGrid_DragMode", 0), // property DragMode
			/* 13 */ imports.NewTable("TStringGrid_HeaderHotZones", 0), // property HeaderHotZones
			/* 14 */ imports.NewTable("TStringGrid_HeaderPushZones", 0), // property HeaderPushZones
			/* 15 */ imports.NewTable("TStringGrid_ImageIndexSortAsc", 0), // property ImageIndexSortAsc
			/* 16 */ imports.NewTable("TStringGrid_ImageIndexSortDesc", 0), // property ImageIndexSortDesc
			/* 17 */ imports.NewTable("TStringGrid_MouseWheelOption", 0), // property MouseWheelOption
			/* 18 */ imports.NewTable("TStringGrid_ParentColor", 0), // property ParentColor
			/* 19 */ imports.NewTable("TStringGrid_ParentFont", 0), // property ParentFont
			/* 20 */ imports.NewTable("TStringGrid_RangeSelectMode", 0), // property RangeSelectMode
			/* 21 */ imports.NewTable("TStringGrid_RowSizingCursor", 0), // property RowSizingCursor
			/* 22 */ imports.NewTable("TStringGrid_TitleFont", 0), // property TitleFont
			/* 23 */ imports.NewTable("TStringGrid_TitleImageList", 0), // property TitleImageList
			/* 24 */ imports.NewTable("TStringGrid_TitleStyle", 0), // property TitleStyle
			/* 25 */ imports.NewTable("TStringGrid_OnCellProcess", 0), // event OnCellProcess
			/* 26 */ imports.NewTable("TStringGrid_OnCheckboxToggled", 0), // event OnCheckboxToggled
			/* 27 */ imports.NewTable("TStringGrid_OnEditingDone", 0), // event OnEditingDone
			/* 28 */ imports.NewTable("TStringGrid_OnGetCellHint", 0), // event OnGetCellHint
			/* 29 */ imports.NewTable("TStringGrid_OnGetCheckboxState", 0), // event OnGetCheckboxState
			/* 30 */ imports.NewTable("TStringGrid_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 31 */ imports.NewTable("TStringGrid_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 32 */ imports.NewTable("TStringGrid_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 33 */ imports.NewTable("TStringGrid_OnSetCheckboxState", 0), // event OnSetCheckboxState
			/* 34 */ imports.NewTable("TStringGrid_OnUserCheckboxBitmap", 0), // event OnUserCheckboxBitmap
			/* 35 */ imports.NewTable("TStringGrid_OnUserCheckboxImage", 0), // event OnUserCheckboxImage
			/* 36 */ imports.NewTable("TStringGrid_TClass", 0), // function TStringGridClass
		}
	})
	return stringGridImport
}
