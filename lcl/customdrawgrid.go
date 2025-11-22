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

// ICustomDrawGrid Parent: ICustomGrid
type ICustomDrawGrid interface {
	ICustomGrid
	DeleteColRow(isColumn bool, index int32)                                             // procedure
	DeleteCol(index int32)                                                               // procedure
	DeleteRow(index int32)                                                               // procedure
	ExchangeColRow(isColumn bool, index int32, withIndex int32)                          // procedure
	InsertColRow(isColumn bool, index int32)                                             // procedure
	MoveColRow(isColumn bool, fromIndex int32, toIndex int32)                            // procedure
	SortColRowWithBoolInt(isColumn bool, index int32)                                    // procedure
	SortColRowWithBoolIntX3(isColumn bool, index int32, fromIndex int32, toIndex int32)  // procedure
	DefaultDrawCell(col int32, row int32, rect *types.TRect, state types.TGridDrawState) // procedure
	AllowOutboundEvents() bool                                                           // property AllowOutboundEvents Getter
	SetAllowOutboundEvents(value bool)                                                   // property AllowOutboundEvents Setter
	BorderColor() types.TColor                                                           // property BorderColor Getter
	SetBorderColor(value types.TColor)                                                   // property BorderColor Setter
	Col() int32                                                                          // property Col Getter
	SetCol(value int32)                                                                  // property Col Setter
	ColWidths(col int32) int32                                                           // property ColWidths Getter
	SetColWidths(col int32, value int32)                                                 // property ColWidths Setter
	ColRow() types.TPoint                                                                // property ColRow Getter
	SetColRow(value types.TPoint)                                                        // property ColRow Setter
	DisabledFontColor() types.TColor                                                     // property DisabledFontColor Getter
	SetDisabledFontColor(value types.TColor)                                             // property DisabledFontColor Setter
	Editor() IWinControl                                                                 // property Editor Getter
	SetEditor(value IWinControl)                                                         // property Editor Setter
	EditorBorderStyle() types.TBorderStyle                                               // property EditorBorderStyle Getter
	SetEditorBorderStyle(value types.TBorderStyle)                                       // property EditorBorderStyle Setter
	EditorMode() bool                                                                    // property EditorMode Getter
	SetEditorMode(value bool)                                                            // property EditorMode Setter
	ExtendedColSizing() bool                                                             // property ExtendedColSizing Getter
	SetExtendedColSizing(value bool)                                                     // property ExtendedColSizing Setter
	AltColorStartNormal() bool                                                           // property AltColorStartNormal Getter
	SetAltColorStartNormal(value bool)                                                   // property AltColorStartNormal Setter
	FastEditing() bool                                                                   // property FastEditing Getter
	SetFastEditing(value bool)                                                           // property FastEditing Setter
	FixedGridLineColor() types.TColor                                                    // property FixedGridLineColor Getter
	SetFixedGridLineColor(value types.TColor)                                            // property FixedGridLineColor Setter
	FocusColor() types.TColor                                                            // property FocusColor Getter
	SetFocusColor(value types.TColor)                                                    // property FocusColor Setter
	FocusRectVisible() bool                                                              // property FocusRectVisible Getter
	SetFocusRectVisible(value bool)                                                      // property FocusRectVisible Setter
	GridHeight() int32                                                                   // property GridHeight Getter
	GridWidth() int32                                                                    // property GridWidth Getter
	IsCellSelected(col int32, row int32) bool                                            // property IsCellSelected Getter
	LeftCol() int32                                                                      // property LeftCol Getter
	SetLeftCol(value int32)                                                              // property LeftCol Setter
	Row() int32                                                                          // property Row Getter
	SetRow(value int32)                                                                  // property Row Setter
	RowHeights(row int32) int32                                                          // property RowHeights Getter
	SetRowHeights(row int32, value int32)                                                // property RowHeights Setter
	SaveOptions() types.TSaveOptions                                                     // property SaveOptions Getter
	SetSaveOptions(value types.TSaveOptions)                                             // property SaveOptions Setter
	SelectedColor() types.TColor                                                         // property SelectedColor Getter
	SetSelectedColor(value types.TColor)                                                 // property SelectedColor Setter
	SelectedColumn() IGridColumn                                                         // property SelectedColumn Getter
	Selection() types.TGridRect                                                          // property Selection Getter
	SetSelection(value types.TGridRect)                                                  // property Selection Setter
	StrictSort() bool                                                                    // property StrictSort Getter
	SetStrictSort(value bool)                                                            // property StrictSort Setter
	TopRow() int32                                                                       // property TopRow Getter
	SetTopRow(value int32)                                                               // property TopRow Setter
	UseXORFeatures() bool                                                                // property UseXORFeatures Getter
	SetUseXORFeatures(value bool)                                                        // property UseXORFeatures Setter
	AutoAdvance() types.TAutoAdvance                                                     // property AutoAdvance Getter
	SetAutoAdvance(value types.TAutoAdvance)                                             // property AutoAdvance Setter
	AutoFillColumns() bool                                                               // property AutoFillColumns Getter
	SetAutoFillColumns(value bool)                                                       // property AutoFillColumns Setter
	ColCount() int32                                                                     // property ColCount Getter
	SetColCount(value int32)                                                             // property ColCount Setter
	Columns() IGridColumns                                                               // property Columns Getter
	SetColumns(value IGridColumns)                                                       // property Columns Setter
	DefaultColWidth() int32                                                              // property DefaultColWidth Getter
	SetDefaultColWidth(value int32)                                                      // property DefaultColWidth Setter
	DefaultDrawing() bool                                                                // property DefaultDrawing Getter
	SetDefaultDrawing(value bool)                                                        // property DefaultDrawing Setter
	DefaultRowHeight() int32                                                             // property DefaultRowHeight Getter
	SetDefaultRowHeight(value int32)                                                     // property DefaultRowHeight Setter
	FadeUnfocusedSelection() bool                                                        // property FadeUnfocusedSelection Getter
	SetFadeUnfocusedSelection(value bool)                                                // property FadeUnfocusedSelection Setter
	FixedColor() types.TColor                                                            // property FixedColor Getter
	SetFixedColor(value types.TColor)                                                    // property FixedColor Setter
	FixedCols() int32                                                                    // property FixedCols Getter
	SetFixedCols(value int32)                                                            // property FixedCols Setter
	FixedHotColor() types.TColor                                                         // property FixedHotColor Getter
	SetFixedHotColor(value types.TColor)                                                 // property FixedHotColor Setter
	FixedRows() int32                                                                    // property FixedRows Getter
	SetFixedRows(value int32)                                                            // property FixedRows Setter
	Flat() bool                                                                          // property Flat Getter
	SetFlat(value bool)                                                                  // property Flat Setter
	GridLineColor() types.TColor                                                         // property GridLineColor Getter
	SetGridLineColor(value types.TColor)                                                 // property GridLineColor Setter
	GridLineStyle() types.TPenStyle                                                      // property GridLineStyle Getter
	SetGridLineStyle(value types.TPenStyle)                                              // property GridLineStyle Setter
	GridLineWidth() int32                                                                // property GridLineWidth Getter
	SetGridLineWidth(value int32)                                                        // property GridLineWidth Setter
	Options() types.TGridOptions                                                         // property Options Getter
	SetOptions(value types.TGridOptions)                                                 // property Options Setter
	Options2() types.TGridOptions2                                                       // property Options2 Getter
	SetOptions2(value types.TGridOptions2)                                               // property Options2 Setter
	ParentShowHint() bool                                                                // property ParentShowHint Getter
	SetParentShowHint(value bool)                                                        // property ParentShowHint Setter
	RowCount() int32                                                                     // property RowCount Getter
	SetRowCount(value int32)                                                             // property RowCount Setter
	ScrollBars() types.TScrollStyle                                                      // property ScrollBars Getter
	SetScrollBars(value types.TScrollStyle)                                              // property ScrollBars Setter
	TabAdvance() types.TAutoAdvance                                                      // property TabAdvance Getter
	SetTabAdvance(value types.TAutoAdvance)                                              // property TabAdvance Setter
	VisibleColCount() int32                                                              // property VisibleColCount Getter
	VisibleRowCount() int32                                                              // property VisibleRowCount Getter
	SetOnAfterSelection(fn TOnSelectEvent)                                               // property event
	SetOnBeforeSelection(fn TOnSelectEvent)                                              // property event
	SetOnCompareCells(fn TOnCompareCells)                                                // property event
	SetOnContextPopup(fn TContextPopupEvent)                                             // property event
	SetOnDblClick(fn TNotifyEvent)                                                       // property event
	SetOnDragDrop(fn TDragDropEvent)                                                     // property event
	SetOnDragOver(fn TDragOverEvent)                                                     // property event
	SetOnDrawCell(fn TOnDrawCell)                                                        // property event
	SetOnButtonClick(fn TOnSelectEvent)                                                  // property event
	SetOnEndDock(fn TEndDragEvent)                                                       // property event
	SetOnEndDrag(fn TEndDragEvent)                                                       // property event
	SetOnGetEditMask(fn TGetEditEvent)                                                   // property event
	SetOnGetEditText(fn TGetEditEvent)                                                   // property event
	SetOnHeaderClick(fn THdrEvent)                                                       // property event
	SetOnHeaderSized(fn THdrEvent)                                                       // property event
	SetOnHeaderSizing(fn THeaderSizingEvent)                                             // property event
	SetOnMouseDown(fn TMouseEvent)                                                       // property event
	SetOnMouseEnter(fn TNotifyEvent)                                                     // property event
	SetOnMouseLeave(fn TNotifyEvent)                                                     // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                                   // property event
	SetOnMouseUp(fn TMouseEvent)                                                         // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                                 // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                                       // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                                         // property event
	SetOnPickListSelect(fn TNotifyEvent)                                                 // property event
	SetOnPrepareCanvas(fn TOnPrepareCanvasEvent)                                         // property event
	SetOnSelectEditor(fn TSelectEditorEvent)                                             // property event
	SetOnSelection(fn TOnSelectEvent)                                                    // property event
	SetOnSelectCell(fn TOnSelectCellEvent)                                               // property event
	SetOnSetEditText(fn TSetEditEvent)                                                   // property event
	SetOnStartDock(fn TStartDockEvent)                                                   // property event
	SetOnStartDrag(fn TStartDragEvent)                                                   // property event
	SetOnTopleftChanged(fn TNotifyEvent)                                                 // property event
	SetOnValidateEntry(fn TValidateEntryEvent)                                           // property event
}

type TCustomDrawGrid struct {
	TCustomGrid
}

func (m *TCustomDrawGrid) DeleteColRow(isColumn bool, index int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(1, m.Instance(), api.PasBool(isColumn), uintptr(index))
}

func (m *TCustomDrawGrid) DeleteCol(index int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(2, m.Instance(), uintptr(index))
}

func (m *TCustomDrawGrid) DeleteRow(index int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(3, m.Instance(), uintptr(index))
}

func (m *TCustomDrawGrid) ExchangeColRow(isColumn bool, index int32, withIndex int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(4, m.Instance(), api.PasBool(isColumn), uintptr(index), uintptr(withIndex))
}

func (m *TCustomDrawGrid) InsertColRow(isColumn bool, index int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(5, m.Instance(), api.PasBool(isColumn), uintptr(index))
}

func (m *TCustomDrawGrid) MoveColRow(isColumn bool, fromIndex int32, toIndex int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(6, m.Instance(), api.PasBool(isColumn), uintptr(fromIndex), uintptr(toIndex))
}

func (m *TCustomDrawGrid) SortColRowWithBoolInt(isColumn bool, index int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(7, m.Instance(), api.PasBool(isColumn), uintptr(index))
}

func (m *TCustomDrawGrid) SortColRowWithBoolIntX3(isColumn bool, index int32, fromIndex int32, toIndex int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(8, m.Instance(), api.PasBool(isColumn), uintptr(index), uintptr(fromIndex), uintptr(toIndex))
}

func (m *TCustomDrawGrid) DefaultDrawCell(col int32, row int32, rect *types.TRect, state types.TGridDrawState) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(9, m.Instance(), uintptr(col), uintptr(row), uintptr(base.UnsafePointer(rect)), uintptr(state))
}

func (m *TCustomDrawGrid) AllowOutboundEvents() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetAllowOutboundEvents(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) BorderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(11, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetBorderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) Col() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetCol(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) ColWidths(col int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(13, 0, m.Instance(), uintptr(col))
	return int32(r)
}

func (m *TCustomDrawGrid) SetColWidths(col int32, value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(13, 1, m.Instance(), uintptr(col), uintptr(value))
}

func (m *TCustomDrawGrid) ColRow() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(14, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomDrawGrid) SetColRow(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(14, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomDrawGrid) DisabledFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(15, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetDisabledFontColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) Editor() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := customDrawGridAPI().SysCallN(16, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TCustomDrawGrid) SetEditor(value IWinControl) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomDrawGrid) EditorBorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(17, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCustomDrawGrid) SetEditorBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) EditorMode() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetEditorMode(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) ExtendedColSizing() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetExtendedColSizing(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) AltColorStartNormal() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetAltColorStartNormal(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) FastEditing() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetFastEditing(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) FixedGridLineColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(22, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetFixedGridLineColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) FocusColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(23, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetFocusColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) FocusRectVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetFocusRectVisible(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) GridHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(25, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) GridWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(26, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) IsCellSelected(col int32, row int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(27, m.Instance(), uintptr(col), uintptr(row))
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) LeftCol() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(28, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetLeftCol(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) Row() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(29, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetRow(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) RowHeights(row int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(30, 0, m.Instance(), uintptr(row))
	return int32(r)
}

func (m *TCustomDrawGrid) SetRowHeights(row int32, value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(30, 1, m.Instance(), uintptr(row), uintptr(value))
}

func (m *TCustomDrawGrid) SaveOptions() types.TSaveOptions {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(31, 0, m.Instance())
	return types.TSaveOptions(r)
}

func (m *TCustomDrawGrid) SetSaveOptions(value types.TSaveOptions) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) SelectedColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(32, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetSelectedColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(32, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) SelectedColumn() IGridColumn {
	if !m.IsValid() {
		return nil
	}
	r := customDrawGridAPI().SysCallN(33, m.Instance())
	return AsGridColumn(r)
}

func (m *TCustomDrawGrid) Selection() (result types.TGridRect) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(34, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomDrawGrid) SetSelection(value types.TGridRect) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(34, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomDrawGrid) StrictSort() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(35, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetStrictSort(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(35, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) TopRow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(36, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetTopRow(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) UseXORFeatures() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(37, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetUseXORFeatures(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(37, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) AutoAdvance() types.TAutoAdvance {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(38, 0, m.Instance())
	return types.TAutoAdvance(r)
}

func (m *TCustomDrawGrid) SetAutoAdvance(value types.TAutoAdvance) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(38, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) AutoFillColumns() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(39, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetAutoFillColumns(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(39, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) ColCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(40, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetColCount(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) Columns() IGridColumns {
	if !m.IsValid() {
		return nil
	}
	r := customDrawGridAPI().SysCallN(41, 0, m.Instance())
	return AsGridColumns(r)
}

func (m *TCustomDrawGrid) SetColumns(value IGridColumns) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(41, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomDrawGrid) DefaultColWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(42, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetDefaultColWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(42, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) DefaultDrawing() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(43, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetDefaultDrawing(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(43, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) DefaultRowHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(44, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetDefaultRowHeight(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(44, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) FadeUnfocusedSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(45, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetFadeUnfocusedSelection(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(45, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) FixedColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(46, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetFixedColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(46, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) FixedCols() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(47, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetFixedCols(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(47, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) FixedHotColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(48, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetFixedHotColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(48, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) FixedRows() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(49, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetFixedRows(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(49, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(50, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(50, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) GridLineColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(51, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomDrawGrid) SetGridLineColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(51, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) GridLineStyle() types.TPenStyle {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(52, 0, m.Instance())
	return types.TPenStyle(r)
}

func (m *TCustomDrawGrid) SetGridLineStyle(value types.TPenStyle) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(52, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) GridLineWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(53, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetGridLineWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(53, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) Options() types.TGridOptions {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(54, 0, m.Instance())
	return types.TGridOptions(r)
}

func (m *TCustomDrawGrid) SetOptions(value types.TGridOptions) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(54, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) Options2() types.TGridOptions2 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(55, 0, m.Instance())
	return types.TGridOptions2(r)
}

func (m *TCustomDrawGrid) SetOptions2(value types.TGridOptions2) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(55, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := customDrawGridAPI().SysCallN(56, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDrawGrid) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(56, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomDrawGrid) RowCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(57, 0, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetRowCount(value int32) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(57, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) ScrollBars() types.TScrollStyle {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(58, 0, m.Instance())
	return types.TScrollStyle(r)
}

func (m *TCustomDrawGrid) SetScrollBars(value types.TScrollStyle) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(58, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) TabAdvance() types.TAutoAdvance {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(59, 0, m.Instance())
	return types.TAutoAdvance(r)
}

func (m *TCustomDrawGrid) SetTabAdvance(value types.TAutoAdvance) {
	if !m.IsValid() {
		return
	}
	customDrawGridAPI().SysCallN(59, 1, m.Instance(), uintptr(value))
}

func (m *TCustomDrawGrid) VisibleColCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(60, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) VisibleRowCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customDrawGridAPI().SysCallN(61, m.Instance())
	return int32(r)
}

func (m *TCustomDrawGrid) SetOnAfterSelection(fn TOnSelectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSelectEvent(fn)
	base.SetEvent(m, 62, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnBeforeSelection(fn TOnSelectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSelectEvent(fn)
	base.SetEvent(m, 63, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnCompareCells(fn TOnCompareCells) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCompareCells(fn)
	base.SetEvent(m, 64, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 65, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 66, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 67, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 68, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnDrawCell(fn TOnDrawCell) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnDrawCell(fn)
	base.SetEvent(m, 69, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnButtonClick(fn TOnSelectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSelectEvent(fn)
	base.SetEvent(m, 70, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 71, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 72, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnGetEditMask(fn TGetEditEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetEditEvent(fn)
	base.SetEvent(m, 73, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnGetEditText(fn TGetEditEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetEditEvent(fn)
	base.SetEvent(m, 74, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnHeaderClick(fn THdrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTHdrEvent(fn)
	base.SetEvent(m, 75, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnHeaderSized(fn THdrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTHdrEvent(fn)
	base.SetEvent(m, 76, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnHeaderSizing(fn THeaderSizingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTHeaderSizingEvent(fn)
	base.SetEvent(m, 77, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 78, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 79, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 80, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 81, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 82, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 83, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 84, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 85, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnPickListSelect(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 86, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnPrepareCanvas(fn TOnPrepareCanvasEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnPrepareCanvasEvent(fn)
	base.SetEvent(m, 87, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnSelectEditor(fn TSelectEditorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSelectEditorEvent(fn)
	base.SetEvent(m, 88, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnSelection(fn TOnSelectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSelectEvent(fn)
	base.SetEvent(m, 89, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnSelectCell(fn TOnSelectCellEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnSelectCellEvent(fn)
	base.SetEvent(m, 90, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnSetEditText(fn TSetEditEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSetEditEvent(fn)
	base.SetEvent(m, 91, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 92, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 93, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnTopleftChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 94, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomDrawGrid) SetOnValidateEntry(fn TValidateEntryEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTValidateEntryEvent(fn)
	base.SetEvent(m, 95, customDrawGridAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomDrawGrid class constructor
func NewCustomDrawGrid(owner IComponent) ICustomDrawGrid {
	r := customDrawGridAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomDrawGrid(r)
}

func TCustomDrawGridClass() types.TClass {
	r := customDrawGridAPI().SysCallN(96)
	return types.TClass(r)
}

var (
	customDrawGridOnce   base.Once
	customDrawGridImport *imports.Imports = nil
)

func customDrawGridAPI() *imports.Imports {
	customDrawGridOnce.Do(func() {
		customDrawGridImport = api.NewDefaultImports()
		customDrawGridImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDrawGrid_Create", 0), // constructor NewCustomDrawGrid
			/* 1 */ imports.NewTable("TCustomDrawGrid_DeleteColRow", 0), // procedure DeleteColRow
			/* 2 */ imports.NewTable("TCustomDrawGrid_DeleteCol", 0), // procedure DeleteCol
			/* 3 */ imports.NewTable("TCustomDrawGrid_DeleteRow", 0), // procedure DeleteRow
			/* 4 */ imports.NewTable("TCustomDrawGrid_ExchangeColRow", 0), // procedure ExchangeColRow
			/* 5 */ imports.NewTable("TCustomDrawGrid_InsertColRow", 0), // procedure InsertColRow
			/* 6 */ imports.NewTable("TCustomDrawGrid_MoveColRow", 0), // procedure MoveColRow
			/* 7 */ imports.NewTable("TCustomDrawGrid_SortColRowWithBoolInt", 0), // procedure SortColRowWithBoolInt
			/* 8 */ imports.NewTable("TCustomDrawGrid_SortColRowWithBoolIntX3", 0), // procedure SortColRowWithBoolIntX3
			/* 9 */ imports.NewTable("TCustomDrawGrid_DefaultDrawCell", 0), // procedure DefaultDrawCell
			/* 10 */ imports.NewTable("TCustomDrawGrid_AllowOutboundEvents", 0), // property AllowOutboundEvents
			/* 11 */ imports.NewTable("TCustomDrawGrid_BorderColor", 0), // property BorderColor
			/* 12 */ imports.NewTable("TCustomDrawGrid_Col", 0), // property Col
			/* 13 */ imports.NewTable("TCustomDrawGrid_ColWidths", 0), // property ColWidths
			/* 14 */ imports.NewTable("TCustomDrawGrid_ColRow", 0), // property ColRow
			/* 15 */ imports.NewTable("TCustomDrawGrid_DisabledFontColor", 0), // property DisabledFontColor
			/* 16 */ imports.NewTable("TCustomDrawGrid_Editor", 0), // property Editor
			/* 17 */ imports.NewTable("TCustomDrawGrid_EditorBorderStyle", 0), // property EditorBorderStyle
			/* 18 */ imports.NewTable("TCustomDrawGrid_EditorMode", 0), // property EditorMode
			/* 19 */ imports.NewTable("TCustomDrawGrid_ExtendedColSizing", 0), // property ExtendedColSizing
			/* 20 */ imports.NewTable("TCustomDrawGrid_AltColorStartNormal", 0), // property AltColorStartNormal
			/* 21 */ imports.NewTable("TCustomDrawGrid_FastEditing", 0), // property FastEditing
			/* 22 */ imports.NewTable("TCustomDrawGrid_FixedGridLineColor", 0), // property FixedGridLineColor
			/* 23 */ imports.NewTable("TCustomDrawGrid_FocusColor", 0), // property FocusColor
			/* 24 */ imports.NewTable("TCustomDrawGrid_FocusRectVisible", 0), // property FocusRectVisible
			/* 25 */ imports.NewTable("TCustomDrawGrid_GridHeight", 0), // property GridHeight
			/* 26 */ imports.NewTable("TCustomDrawGrid_GridWidth", 0), // property GridWidth
			/* 27 */ imports.NewTable("TCustomDrawGrid_IsCellSelected", 0), // property IsCellSelected
			/* 28 */ imports.NewTable("TCustomDrawGrid_LeftCol", 0), // property LeftCol
			/* 29 */ imports.NewTable("TCustomDrawGrid_Row", 0), // property Row
			/* 30 */ imports.NewTable("TCustomDrawGrid_RowHeights", 0), // property RowHeights
			/* 31 */ imports.NewTable("TCustomDrawGrid_SaveOptions", 0), // property SaveOptions
			/* 32 */ imports.NewTable("TCustomDrawGrid_SelectedColor", 0), // property SelectedColor
			/* 33 */ imports.NewTable("TCustomDrawGrid_SelectedColumn", 0), // property SelectedColumn
			/* 34 */ imports.NewTable("TCustomDrawGrid_Selection", 0), // property Selection
			/* 35 */ imports.NewTable("TCustomDrawGrid_StrictSort", 0), // property StrictSort
			/* 36 */ imports.NewTable("TCustomDrawGrid_TopRow", 0), // property TopRow
			/* 37 */ imports.NewTable("TCustomDrawGrid_UseXORFeatures", 0), // property UseXORFeatures
			/* 38 */ imports.NewTable("TCustomDrawGrid_AutoAdvance", 0), // property AutoAdvance
			/* 39 */ imports.NewTable("TCustomDrawGrid_AutoFillColumns", 0), // property AutoFillColumns
			/* 40 */ imports.NewTable("TCustomDrawGrid_ColCount", 0), // property ColCount
			/* 41 */ imports.NewTable("TCustomDrawGrid_Columns", 0), // property Columns
			/* 42 */ imports.NewTable("TCustomDrawGrid_DefaultColWidth", 0), // property DefaultColWidth
			/* 43 */ imports.NewTable("TCustomDrawGrid_DefaultDrawing", 0), // property DefaultDrawing
			/* 44 */ imports.NewTable("TCustomDrawGrid_DefaultRowHeight", 0), // property DefaultRowHeight
			/* 45 */ imports.NewTable("TCustomDrawGrid_FadeUnfocusedSelection", 0), // property FadeUnfocusedSelection
			/* 46 */ imports.NewTable("TCustomDrawGrid_FixedColor", 0), // property FixedColor
			/* 47 */ imports.NewTable("TCustomDrawGrid_FixedCols", 0), // property FixedCols
			/* 48 */ imports.NewTable("TCustomDrawGrid_FixedHotColor", 0), // property FixedHotColor
			/* 49 */ imports.NewTable("TCustomDrawGrid_FixedRows", 0), // property FixedRows
			/* 50 */ imports.NewTable("TCustomDrawGrid_Flat", 0), // property Flat
			/* 51 */ imports.NewTable("TCustomDrawGrid_GridLineColor", 0), // property GridLineColor
			/* 52 */ imports.NewTable("TCustomDrawGrid_GridLineStyle", 0), // property GridLineStyle
			/* 53 */ imports.NewTable("TCustomDrawGrid_GridLineWidth", 0), // property GridLineWidth
			/* 54 */ imports.NewTable("TCustomDrawGrid_Options", 0), // property Options
			/* 55 */ imports.NewTable("TCustomDrawGrid_Options2", 0), // property Options2
			/* 56 */ imports.NewTable("TCustomDrawGrid_ParentShowHint", 0), // property ParentShowHint
			/* 57 */ imports.NewTable("TCustomDrawGrid_RowCount", 0), // property RowCount
			/* 58 */ imports.NewTable("TCustomDrawGrid_ScrollBars", 0), // property ScrollBars
			/* 59 */ imports.NewTable("TCustomDrawGrid_TabAdvance", 0), // property TabAdvance
			/* 60 */ imports.NewTable("TCustomDrawGrid_VisibleColCount", 0), // property VisibleColCount
			/* 61 */ imports.NewTable("TCustomDrawGrid_VisibleRowCount", 0), // property VisibleRowCount
			/* 62 */ imports.NewTable("TCustomDrawGrid_OnAfterSelection", 0), // event OnAfterSelection
			/* 63 */ imports.NewTable("TCustomDrawGrid_OnBeforeSelection", 0), // event OnBeforeSelection
			/* 64 */ imports.NewTable("TCustomDrawGrid_OnCompareCells", 0), // event OnCompareCells
			/* 65 */ imports.NewTable("TCustomDrawGrid_OnContextPopup", 0), // event OnContextPopup
			/* 66 */ imports.NewTable("TCustomDrawGrid_OnDblClick", 0), // event OnDblClick
			/* 67 */ imports.NewTable("TCustomDrawGrid_OnDragDrop", 0), // event OnDragDrop
			/* 68 */ imports.NewTable("TCustomDrawGrid_OnDragOver", 0), // event OnDragOver
			/* 69 */ imports.NewTable("TCustomDrawGrid_OnDrawCell", 0), // event OnDrawCell
			/* 70 */ imports.NewTable("TCustomDrawGrid_OnButtonClick", 0), // event OnButtonClick
			/* 71 */ imports.NewTable("TCustomDrawGrid_OnEndDock", 0), // event OnEndDock
			/* 72 */ imports.NewTable("TCustomDrawGrid_OnEndDrag", 0), // event OnEndDrag
			/* 73 */ imports.NewTable("TCustomDrawGrid_OnGetEditMask", 0), // event OnGetEditMask
			/* 74 */ imports.NewTable("TCustomDrawGrid_OnGetEditText", 0), // event OnGetEditText
			/* 75 */ imports.NewTable("TCustomDrawGrid_OnHeaderClick", 0), // event OnHeaderClick
			/* 76 */ imports.NewTable("TCustomDrawGrid_OnHeaderSized", 0), // event OnHeaderSized
			/* 77 */ imports.NewTable("TCustomDrawGrid_OnHeaderSizing", 0), // event OnHeaderSizing
			/* 78 */ imports.NewTable("TCustomDrawGrid_OnMouseDown", 0), // event OnMouseDown
			/* 79 */ imports.NewTable("TCustomDrawGrid_OnMouseEnter", 0), // event OnMouseEnter
			/* 80 */ imports.NewTable("TCustomDrawGrid_OnMouseLeave", 0), // event OnMouseLeave
			/* 81 */ imports.NewTable("TCustomDrawGrid_OnMouseMove", 0), // event OnMouseMove
			/* 82 */ imports.NewTable("TCustomDrawGrid_OnMouseUp", 0), // event OnMouseUp
			/* 83 */ imports.NewTable("TCustomDrawGrid_OnMouseWheel", 0), // event OnMouseWheel
			/* 84 */ imports.NewTable("TCustomDrawGrid_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 85 */ imports.NewTable("TCustomDrawGrid_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 86 */ imports.NewTable("TCustomDrawGrid_OnPickListSelect", 0), // event OnPickListSelect
			/* 87 */ imports.NewTable("TCustomDrawGrid_OnPrepareCanvas", 0), // event OnPrepareCanvas
			/* 88 */ imports.NewTable("TCustomDrawGrid_OnSelectEditor", 0), // event OnSelectEditor
			/* 89 */ imports.NewTable("TCustomDrawGrid_OnSelection", 0), // event OnSelection
			/* 90 */ imports.NewTable("TCustomDrawGrid_OnSelectCell", 0), // event OnSelectCell
			/* 91 */ imports.NewTable("TCustomDrawGrid_OnSetEditText", 0), // event OnSetEditText
			/* 92 */ imports.NewTable("TCustomDrawGrid_OnStartDock", 0), // event OnStartDock
			/* 93 */ imports.NewTable("TCustomDrawGrid_OnStartDrag", 0), // event OnStartDrag
			/* 94 */ imports.NewTable("TCustomDrawGrid_OnTopleftChanged", 0), // event OnTopleftChanged
			/* 95 */ imports.NewTable("TCustomDrawGrid_OnValidateEntry", 0), // event OnValidateEntry
			/* 96 */ imports.NewTable("TCustomDrawGrid_TClass", 0), // function TCustomDrawGridClass
		}
	})
	return customDrawGridImport
}
