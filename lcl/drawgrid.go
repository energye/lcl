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

// IDrawGrid Parent: ICustomDrawGrid
type IDrawGrid interface {
	ICustomDrawGrid
	InplaceEditor() IWinControl                          // property InplaceEditor Getter
	AlternateColor() types.TColor                        // property AlternateColor Getter
	SetAlternateColor(value types.TColor)                // property AlternateColor Setter
	AutoEdit() bool                                      // property AutoEdit Getter
	SetAutoEdit(value bool)                              // property AutoEdit Setter
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
	ExtendedSelect() bool                                // property ExtendedSelect Getter
	SetExtendedSelect(value bool)                        // property ExtendedSelect Setter
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
	TitleImageListWidth() int32                          // property TitleImageListWidth Getter
	SetTitleImageListWidth(value int32)                  // property TitleImageListWidth Setter
	TitleStyle() types.TTitleStyle                       // property TitleStyle Getter
	SetTitleStyle(value types.TTitleStyle)               // property TitleStyle Setter
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

type TDrawGrid struct {
	TCustomDrawGrid
}

func (m *TDrawGrid) InplaceEditor() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := drawGridAPI().SysCallN(1, m.Instance())
	return AsWinControl(r)
}

func (m *TDrawGrid) AlternateColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(2, 0, m.Instance())
	return types.TColor(r)
}

func (m *TDrawGrid) SetAlternateColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) AutoEdit() bool {
	if !m.IsValid() {
		return false
	}
	r := drawGridAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDrawGrid) SetAutoEdit(value bool) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TDrawGrid) ColRowDraggingCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(4, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDrawGrid) SetColRowDraggingCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) ColRowDragIndicatorColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(5, 0, m.Instance())
	return types.TColor(r)
}

func (m *TDrawGrid) SetColRowDragIndicatorColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) ColSizingCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(6, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDrawGrid) SetColSizingCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) ColumnClickSorts() bool {
	if !m.IsValid() {
		return false
	}
	r := drawGridAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDrawGrid) SetColumnClickSorts(value bool) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TDrawGrid) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(8, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDrawGrid) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(9, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TDrawGrid) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(10, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TDrawGrid) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) ExtendedSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := drawGridAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDrawGrid) SetExtendedSelect(value bool) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TDrawGrid) HeaderHotZones() types.TGridZoneSet {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(12, 0, m.Instance())
	return types.TGridZoneSet(r)
}

func (m *TDrawGrid) SetHeaderHotZones(value types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) HeaderPushZones() types.TGridZoneSet {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(13, 0, m.Instance())
	return types.TGridZoneSet(r)
}

func (m *TDrawGrid) SetHeaderPushZones(value types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) ImageIndexSortAsc() int32 {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TDrawGrid) SetImageIndexSortAsc(value int32) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) ImageIndexSortDesc() int32 {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TDrawGrid) SetImageIndexSortDesc(value int32) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) MouseWheelOption() types.TMouseWheelOption {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(16, 0, m.Instance())
	return types.TMouseWheelOption(r)
}

func (m *TDrawGrid) SetMouseWheelOption(value types.TMouseWheelOption) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := drawGridAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDrawGrid) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TDrawGrid) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := drawGridAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDrawGrid) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TDrawGrid) RangeSelectMode() types.TRangeSelectMode {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(19, 0, m.Instance())
	return types.TRangeSelectMode(r)
}

func (m *TDrawGrid) SetRangeSelectMode(value types.TRangeSelectMode) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) RowSizingCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(20, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDrawGrid) SetRowSizingCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) TitleFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := drawGridAPI().SysCallN(21, 0, m.Instance())
	return AsFont(r)
}

func (m *TDrawGrid) SetTitleFont(value IFont) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDrawGrid) TitleImageList() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := drawGridAPI().SysCallN(22, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TDrawGrid) SetTitleImageList(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDrawGrid) TitleImageListWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TDrawGrid) SetTitleImageListWidth(value int32) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) TitleStyle() types.TTitleStyle {
	if !m.IsValid() {
		return 0
	}
	r := drawGridAPI().SysCallN(24, 0, m.Instance())
	return types.TTitleStyle(r)
}

func (m *TDrawGrid) SetTitleStyle(value types.TTitleStyle) {
	if !m.IsValid() {
		return
	}
	drawGridAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TDrawGrid) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTToggledCheckboxEvent(fn)
	base.SetEvent(m, 25, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 26, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnGetCellHint(fn TGetCellHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetCellHintEvent(fn)
	base.SetEvent(m, 27, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnGetCheckboxState(fn TGetCheckboxStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetCheckboxStateEvent(fn)
	base.SetEvent(m, 28, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 29, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 30, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 31, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnSetCheckboxState(fn TSetCheckboxStateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSetCheckboxStateEvent(fn)
	base.SetEvent(m, 32, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnUserCheckboxBitmap(fn TUserCheckboxBitmapEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUserCheckboxBitmapEvent(fn)
	base.SetEvent(m, 33, drawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDrawGrid) SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUserCheckBoxImageEvent(fn)
	base.SetEvent(m, 34, drawGridAPI(), api.MakeEventDataPtr(cb))
}

// NewDrawGrid class constructor
func NewDrawGrid(owner IComponent) IDrawGrid {
	r := drawGridAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsDrawGrid(r)
}

func TDrawGridClass() types.TClass {
	r := drawGridAPI().SysCallN(35)
	return types.TClass(r)
}

var (
	drawGridOnce   base.Once
	drawGridImport *imports.Imports = nil
)

func drawGridAPI() *imports.Imports {
	drawGridOnce.Do(func() {
		drawGridImport = api.NewDefaultImports()
		drawGridImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDrawGrid_Create", 0), // constructor NewDrawGrid
			/* 1 */ imports.NewTable("TDrawGrid_InplaceEditor", 0), // property InplaceEditor
			/* 2 */ imports.NewTable("TDrawGrid_AlternateColor", 0), // property AlternateColor
			/* 3 */ imports.NewTable("TDrawGrid_AutoEdit", 0), // property AutoEdit
			/* 4 */ imports.NewTable("TDrawGrid_ColRowDraggingCursor", 0), // property ColRowDraggingCursor
			/* 5 */ imports.NewTable("TDrawGrid_ColRowDragIndicatorColor", 0), // property ColRowDragIndicatorColor
			/* 6 */ imports.NewTable("TDrawGrid_ColSizingCursor", 0), // property ColSizingCursor
			/* 7 */ imports.NewTable("TDrawGrid_ColumnClickSorts", 0), // property ColumnClickSorts
			/* 8 */ imports.NewTable("TDrawGrid_DragCursor", 0), // property DragCursor
			/* 9 */ imports.NewTable("TDrawGrid_DragKind", 0), // property DragKind
			/* 10 */ imports.NewTable("TDrawGrid_DragMode", 0), // property DragMode
			/* 11 */ imports.NewTable("TDrawGrid_ExtendedSelect", 0), // property ExtendedSelect
			/* 12 */ imports.NewTable("TDrawGrid_HeaderHotZones", 0), // property HeaderHotZones
			/* 13 */ imports.NewTable("TDrawGrid_HeaderPushZones", 0), // property HeaderPushZones
			/* 14 */ imports.NewTable("TDrawGrid_ImageIndexSortAsc", 0), // property ImageIndexSortAsc
			/* 15 */ imports.NewTable("TDrawGrid_ImageIndexSortDesc", 0), // property ImageIndexSortDesc
			/* 16 */ imports.NewTable("TDrawGrid_MouseWheelOption", 0), // property MouseWheelOption
			/* 17 */ imports.NewTable("TDrawGrid_ParentColor", 0), // property ParentColor
			/* 18 */ imports.NewTable("TDrawGrid_ParentFont", 0), // property ParentFont
			/* 19 */ imports.NewTable("TDrawGrid_RangeSelectMode", 0), // property RangeSelectMode
			/* 20 */ imports.NewTable("TDrawGrid_RowSizingCursor", 0), // property RowSizingCursor
			/* 21 */ imports.NewTable("TDrawGrid_TitleFont", 0), // property TitleFont
			/* 22 */ imports.NewTable("TDrawGrid_TitleImageList", 0), // property TitleImageList
			/* 23 */ imports.NewTable("TDrawGrid_TitleImageListWidth", 0), // property TitleImageListWidth
			/* 24 */ imports.NewTable("TDrawGrid_TitleStyle", 0), // property TitleStyle
			/* 25 */ imports.NewTable("TDrawGrid_OnCheckboxToggled", 0), // event OnCheckboxToggled
			/* 26 */ imports.NewTable("TDrawGrid_OnEditingDone", 0), // event OnEditingDone
			/* 27 */ imports.NewTable("TDrawGrid_OnGetCellHint", 0), // event OnGetCellHint
			/* 28 */ imports.NewTable("TDrawGrid_OnGetCheckboxState", 0), // event OnGetCheckboxState
			/* 29 */ imports.NewTable("TDrawGrid_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 30 */ imports.NewTable("TDrawGrid_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 31 */ imports.NewTable("TDrawGrid_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 32 */ imports.NewTable("TDrawGrid_OnSetCheckboxState", 0), // event OnSetCheckboxState
			/* 33 */ imports.NewTable("TDrawGrid_OnUserCheckboxBitmap", 0), // event OnUserCheckboxBitmap
			/* 34 */ imports.NewTable("TDrawGrid_OnUserCheckboxImage", 0), // event OnUserCheckboxImage
			/* 35 */ imports.NewTable("TDrawGrid_TClass", 0), // function TDrawGridClass
		}
	})
	return drawGridImport
}
