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

// ICustomDrawGrid Parent: ICustomGrid
type ICustomDrawGrid interface {
	ICustomGrid
	AllowOutboundEvents() bool                                             // property
	SetAllowOutboundEvents(AValue bool)                                    // property
	BorderColor() TColor                                                   // property
	SetBorderColor(AValue TColor)                                          // property
	Col() int32                                                            // property
	SetCol(AValue int32)                                                   // property
	ColWidths(aCol int32) int32                                            // property
	SetColWidths(aCol int32, AValue int32)                                 // property
	ColRow() (resultPoint TPoint)                                          // property
	SetColRow(AValue *TPoint)                                              // property
	DisabledFontColor() TColor                                             // property
	SetDisabledFontColor(AValue TColor)                                    // property
	Editor() IWinControl                                                   // property
	SetEditor(AValue IWinControl)                                          // property
	EditorBorderStyle() TBorderStyle                                       // property
	SetEditorBorderStyle(AValue TBorderStyle)                              // property
	EditorMode() bool                                                      // property
	SetEditorMode(AValue bool)                                             // property
	ExtendedColSizing() bool                                               // property
	SetExtendedColSizing(AValue bool)                                      // property
	AltColorStartNormal() bool                                             // property
	SetAltColorStartNormal(AValue bool)                                    // property
	FastEditing() bool                                                     // property
	SetFastEditing(AValue bool)                                            // property
	FixedGridLineColor() TColor                                            // property
	SetFixedGridLineColor(AValue TColor)                                   // property
	FocusColor() TColor                                                    // property
	SetFocusColor(AValue TColor)                                           // property
	FocusRectVisible() bool                                                // property
	SetFocusRectVisible(AValue bool)                                       // property
	GridHeight() int32                                                     // property
	GridWidth() int32                                                      // property
	IsCellSelected(aCol, aRow int32) bool                                  // property
	LeftCol() int32                                                        // property
	SetLeftCol(AValue int32)                                               // property
	Row() int32                                                            // property
	SetRow(AValue int32)                                                   // property
	RowHeights(aRow int32) int32                                           // property
	SetRowHeights(aRow int32, AValue int32)                                // property
	SaveOptions() TSaveOptions                                             // property
	SetSaveOptions(AValue TSaveOptions)                                    // property
	SelectedColor() TColor                                                 // property
	SetSelectedColor(AValue TColor)                                        // property
	SelectedColumn() IGridColumn                                           // property
	Selection() (resultGridRect TGridRect)                                 // property
	SetSelection(AValue *TGridRect)                                        // property
	StrictSort() bool                                                      // property
	SetStrictSort(AValue bool)                                             // property
	TopRow() int32                                                         // property
	SetTopRow(AValue int32)                                                // property
	UseXORFeatures() bool                                                  // property
	SetUseXORFeatures(AValue bool)                                         // property
	AutoAdvance() TAutoAdvance                                             // property
	SetAutoAdvance(AValue TAutoAdvance)                                    // property
	AutoFillColumns() bool                                                 // property
	SetAutoFillColumns(AValue bool)                                        // property
	ColCount() int32                                                       // property
	SetColCount(AValue int32)                                              // property
	Columns() IGridColumns                                                 // property
	SetColumns(AValue IGridColumns)                                        // property
	DefaultColWidth() int32                                                // property
	SetDefaultColWidth(AValue int32)                                       // property
	DefaultDrawing() bool                                                  // property
	SetDefaultDrawing(AValue bool)                                         // property
	DefaultRowHeight() int32                                               // property
	SetDefaultRowHeight(AValue int32)                                      // property
	FadeUnfocusedSelection() bool                                          // property
	SetFadeUnfocusedSelection(AValue bool)                                 // property
	FixedColor() TColor                                                    // property
	SetFixedColor(AValue TColor)                                           // property
	FixedCols() int32                                                      // property
	SetFixedCols(AValue int32)                                             // property
	FixedHotColor() TColor                                                 // property
	SetFixedHotColor(AValue TColor)                                        // property
	FixedRows() int32                                                      // property
	SetFixedRows(AValue int32)                                             // property
	Flat() bool                                                            // property
	SetFlat(AValue bool)                                                   // property
	GridLineColor() TColor                                                 // property
	SetGridLineColor(AValue TColor)                                        // property
	GridLineStyle() TPenStyle                                              // property
	SetGridLineStyle(AValue TPenStyle)                                     // property
	GridLineWidth() int32                                                  // property
	SetGridLineWidth(AValue int32)                                         // property
	Options() TGridOptions                                                 // property
	SetOptions(AValue TGridOptions)                                        // property
	Options2() TGridOptions2                                               // property
	SetOptions2(AValue TGridOptions2)                                      // property
	ParentShowHint() bool                                                  // property
	SetParentShowHint(AValue bool)                                         // property
	RowCount() int32                                                       // property
	SetRowCount(AValue int32)                                              // property
	ScrollBars() TScrollStyle                                              // property
	SetScrollBars(AValue TScrollStyle)                                     // property
	TabAdvance() TAutoAdvance                                              // property
	SetTabAdvance(AValue TAutoAdvance)                                     // property
	VisibleColCount() int32                                                // property
	VisibleRowCount() int32                                                // property
	DeleteColRow(IsColumn bool, index int32)                               // procedure
	DeleteCol(Index int32)                                                 // procedure
	DeleteRow(Index int32)                                                 // procedure
	ExchangeColRow(IsColumn bool, index, WithIndex int32)                  // procedure
	InsertColRow(IsColumn bool, index int32)                               // procedure
	MoveColRow(IsColumn bool, FromIndex, ToIndex int32)                    // procedure
	SortColRow(IsColumn bool, index int32)                                 // procedure
	SortColRow1(IsColumn bool, Index, FromIndex, ToIndex int32)            // procedure
	DefaultDrawCell(aCol, aRow int32, aRect *TRect, aState TGridDrawState) // procedure
	SetOnAfterSelection(fn TOnSelectEvent)                                 // property event
	SetOnBeforeSelection(fn TOnSelectEvent)                                // property event
	SetOnColRowDeleted(fn TGridOperationEvent)                             // property event
	SetOnColRowExchanged(fn TGridOperationEvent)                           // property event
	SetOnColRowInserted(fn TGridOperationEvent)                            // property event
	SetOnColRowMoved(fn TGridOperationEvent)                               // property event
	SetOnCompareCells(fn TOnCompareCells)                                  // property event
	SetOnContextPopup(fn TContextPopupEvent)                               // property event
	SetOnDblClick(fn TNotifyEvent)                                         // property event
	SetOnDragDrop(fn TDragDropEvent)                                       // property event
	SetOnDragOver(fn TDragOverEvent)                                       // property event
	SetOnDrawCell(fn TOnDrawCell)                                          // property event
	SetOnButtonClick(fn TOnSelectEvent)                                    // property event
	SetOnEndDock(fn TEndDragEvent)                                         // property event
	SetOnEndDrag(fn TEndDragEvent)                                         // property event
	SetOnGetEditMask(fn TGetEditEvent)                                     // property event
	SetOnGetEditText(fn TGetEditEvent)                                     // property event
	SetOnHeaderClick(fn THdrEvent)                                         // property event
	SetOnHeaderSized(fn THdrEvent)                                         // property event
	SetOnHeaderSizing(fn THeaderSizingEvent)                               // property event
	SetOnMouseDown(fn TMouseEvent)                                         // property event
	SetOnMouseEnter(fn TNotifyEvent)                                       // property event
	SetOnMouseLeave(fn TNotifyEvent)                                       // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                     // property event
	SetOnMouseUp(fn TMouseEvent)                                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                   // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                         // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                           // property event
	SetOnPickListSelect(fn TNotifyEvent)                                   // property event
	SetOnPrepareCanvas(fn TOnPrepareCanvasEvent)                           // property event
	SetOnSelectEditor(fn TSelectEditorEvent)                               // property event
	SetOnSelection(fn TOnSelectEvent)                                      // property event
	SetOnSelectCell(fn TOnSelectCellEvent)                                 // property event
	SetOnSetEditText(fn TSetEditEvent)                                     // property event
	SetOnStartDock(fn TStartDockEvent)                                     // property event
	SetOnStartDrag(fn TStartDragEvent)                                     // property event
	SetOnTopleftChanged(fn TNotifyEvent)                                   // property event
	SetOnValidateEntry(fn TValidateEntryEvent)                             // property event
}

// TCustomDrawGrid Parent: TCustomGrid
type TCustomDrawGrid struct {
	TCustomGrid
	afterSelectionPtr  uintptr
	beforeSelectionPtr uintptr
	colRowDeletedPtr   uintptr
	colRowExchangedPtr uintptr
	colRowInsertedPtr  uintptr
	colRowMovedPtr     uintptr
	compareCellsPtr    uintptr
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	drawCellPtr        uintptr
	buttonClickPtr     uintptr
	endDockPtr         uintptr
	endDragPtr         uintptr
	getEditMaskPtr     uintptr
	getEditTextPtr     uintptr
	headerClickPtr     uintptr
	headerSizedPtr     uintptr
	headerSizingPtr    uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	pickListSelectPtr  uintptr
	prepareCanvasPtr   uintptr
	selectEditorPtr    uintptr
	selectionPtr       uintptr
	selectCellPtr      uintptr
	setEditTextPtr     uintptr
	startDockPtr       uintptr
	startDragPtr       uintptr
	topleftChangedPtr  uintptr
	validateEntryPtr   uintptr
}

func NewCustomDrawGrid(AOwner IComponent) ICustomDrawGrid {
	r1 := customDrawGridImportAPI().SysCallN(11, GetObjectUintptr(AOwner))
	return AsCustomDrawGrid(r1)
}

func (m *TCustomDrawGrid) AllowOutboundEvents() bool {
	r1 := customDrawGridImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAllowOutboundEvents(AValue bool) {
	customDrawGridImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) BorderColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetBorderColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Col() int32 {
	r1 := customDrawGridImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetCol(AValue int32) {
	customDrawGridImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ColWidths(aCol int32) int32 {
	r1 := customDrawGridImportAPI().SysCallN(9, 0, m.Instance(), uintptr(aCol))
	return int32(r1)
}

func (m *TCustomDrawGrid) SetColWidths(aCol int32, AValue int32) {
	customDrawGridImportAPI().SysCallN(9, 1, m.Instance(), uintptr(aCol), uintptr(AValue))
}

func (m *TCustomDrawGrid) ColRow() (resultPoint TPoint) {
	customDrawGridImportAPI().SysCallN(8, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomDrawGrid) SetColRow(AValue *TPoint) {
	customDrawGridImportAPI().SysCallN(8, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomDrawGrid) DisabledFontColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetDisabledFontColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Editor() IWinControl {
	r1 := customDrawGridImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomDrawGrid) SetEditor(AValue IWinControl) {
	customDrawGridImportAPI().SysCallN(20, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomDrawGrid) EditorBorderStyle() TBorderStyle {
	r1 := customDrawGridImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomDrawGrid) SetEditorBorderStyle(AValue TBorderStyle) {
	customDrawGridImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) EditorMode() bool {
	r1 := customDrawGridImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetEditorMode(AValue bool) {
	customDrawGridImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) ExtendedColSizing() bool {
	r1 := customDrawGridImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetExtendedColSizing(AValue bool) {
	customDrawGridImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) AltColorStartNormal() bool {
	r1 := customDrawGridImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAltColorStartNormal(AValue bool) {
	customDrawGridImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FastEditing() bool {
	r1 := customDrawGridImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFastEditing(AValue bool) {
	customDrawGridImportAPI().SysCallN(26, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FixedGridLineColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedGridLineColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FocusColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFocusColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(33, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FocusRectVisible() bool {
	r1 := customDrawGridImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFocusRectVisible(AValue bool) {
	customDrawGridImportAPI().SysCallN(34, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) GridHeight() int32 {
	r1 := customDrawGridImportAPI().SysCallN(35, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) GridWidth() int32 {
	r1 := customDrawGridImportAPI().SysCallN(39, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) IsCellSelected(aCol, aRow int32) bool {
	r1 := customDrawGridImportAPI().SysCallN(41, m.Instance(), uintptr(aCol), uintptr(aRow))
	return GoBool(r1)
}

func (m *TCustomDrawGrid) LeftCol() int32 {
	r1 := customDrawGridImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetLeftCol(AValue int32) {
	customDrawGridImportAPI().SysCallN(42, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Row() int32 {
	r1 := customDrawGridImportAPI().SysCallN(47, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRow(AValue int32) {
	customDrawGridImportAPI().SysCallN(47, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) RowHeights(aRow int32) int32 {
	r1 := customDrawGridImportAPI().SysCallN(49, 0, m.Instance(), uintptr(aRow))
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRowHeights(aRow int32, AValue int32) {
	customDrawGridImportAPI().SysCallN(49, 1, m.Instance(), uintptr(aRow), uintptr(AValue))
}

func (m *TCustomDrawGrid) SaveOptions() TSaveOptions {
	r1 := customDrawGridImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return TSaveOptions(r1)
}

func (m *TCustomDrawGrid) SetSaveOptions(AValue TSaveOptions) {
	customDrawGridImportAPI().SysCallN(50, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) SelectedColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(52, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetSelectedColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(52, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) SelectedColumn() IGridColumn {
	r1 := customDrawGridImportAPI().SysCallN(53, m.Instance())
	return AsGridColumn(r1)
}

func (m *TCustomDrawGrid) Selection() (resultGridRect TGridRect) {
	customDrawGridImportAPI().SysCallN(54, 0, m.Instance(), uintptr(unsafePointer(&resultGridRect)), uintptr(unsafePointer(&resultGridRect)))
	return
}

func (m *TCustomDrawGrid) SetSelection(AValue *TGridRect) {
	customDrawGridImportAPI().SysCallN(54, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomDrawGrid) StrictSort() bool {
	r1 := customDrawGridImportAPI().SysCallN(95, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetStrictSort(AValue bool) {
	customDrawGridImportAPI().SysCallN(95, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) TopRow() int32 {
	r1 := customDrawGridImportAPI().SysCallN(97, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetTopRow(AValue int32) {
	customDrawGridImportAPI().SysCallN(97, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) UseXORFeatures() bool {
	r1 := customDrawGridImportAPI().SysCallN(98, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetUseXORFeatures(AValue bool) {
	customDrawGridImportAPI().SysCallN(98, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) AutoAdvance() TAutoAdvance {
	r1 := customDrawGridImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TAutoAdvance(r1)
}

func (m *TCustomDrawGrid) SetAutoAdvance(AValue TAutoAdvance) {
	customDrawGridImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) AutoFillColumns() bool {
	r1 := customDrawGridImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAutoFillColumns(AValue bool) {
	customDrawGridImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) ColCount() int32 {
	r1 := customDrawGridImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetColCount(AValue int32) {
	customDrawGridImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Columns() IGridColumns {
	r1 := customDrawGridImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsGridColumns(r1)
}

func (m *TCustomDrawGrid) SetColumns(AValue IGridColumns) {
	customDrawGridImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomDrawGrid) DefaultColWidth() int32 {
	r1 := customDrawGridImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetDefaultColWidth(AValue int32) {
	customDrawGridImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) DefaultDrawing() bool {
	r1 := customDrawGridImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetDefaultDrawing(AValue bool) {
	customDrawGridImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) DefaultRowHeight() int32 {
	r1 := customDrawGridImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetDefaultRowHeight(AValue int32) {
	customDrawGridImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FadeUnfocusedSelection() bool {
	r1 := customDrawGridImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFadeUnfocusedSelection(AValue bool) {
	customDrawGridImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FixedColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedCols() int32 {
	r1 := customDrawGridImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetFixedCols(AValue int32) {
	customDrawGridImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedHotColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedHotColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(30, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedRows() int32 {
	r1 := customDrawGridImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetFixedRows(AValue int32) {
	customDrawGridImportAPI().SysCallN(31, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Flat() bool {
	r1 := customDrawGridImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFlat(AValue bool) {
	customDrawGridImportAPI().SysCallN(32, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) GridLineColor() TColor {
	r1 := customDrawGridImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetGridLineColor(AValue TColor) {
	customDrawGridImportAPI().SysCallN(36, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) GridLineStyle() TPenStyle {
	r1 := customDrawGridImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return TPenStyle(r1)
}

func (m *TCustomDrawGrid) SetGridLineStyle(AValue TPenStyle) {
	customDrawGridImportAPI().SysCallN(37, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) GridLineWidth() int32 {
	r1 := customDrawGridImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetGridLineWidth(AValue int32) {
	customDrawGridImportAPI().SysCallN(38, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Options() TGridOptions {
	r1 := customDrawGridImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return TGridOptions(r1)
}

func (m *TCustomDrawGrid) SetOptions(AValue TGridOptions) {
	customDrawGridImportAPI().SysCallN(44, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Options2() TGridOptions2 {
	r1 := customDrawGridImportAPI().SysCallN(45, 0, m.Instance(), 0)
	return TGridOptions2(r1)
}

func (m *TCustomDrawGrid) SetOptions2(AValue TGridOptions2) {
	customDrawGridImportAPI().SysCallN(45, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ParentShowHint() bool {
	r1 := customDrawGridImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetParentShowHint(AValue bool) {
	customDrawGridImportAPI().SysCallN(46, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) RowCount() int32 {
	r1 := customDrawGridImportAPI().SysCallN(48, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRowCount(AValue int32) {
	customDrawGridImportAPI().SysCallN(48, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ScrollBars() TScrollStyle {
	r1 := customDrawGridImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TCustomDrawGrid) SetScrollBars(AValue TScrollStyle) {
	customDrawGridImportAPI().SysCallN(51, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) TabAdvance() TAutoAdvance {
	r1 := customDrawGridImportAPI().SysCallN(96, 0, m.Instance(), 0)
	return TAutoAdvance(r1)
}

func (m *TCustomDrawGrid) SetTabAdvance(AValue TAutoAdvance) {
	customDrawGridImportAPI().SysCallN(96, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) VisibleColCount() int32 {
	r1 := customDrawGridImportAPI().SysCallN(99, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) VisibleRowCount() int32 {
	r1 := customDrawGridImportAPI().SysCallN(100, m.Instance())
	return int32(r1)
}

func CustomDrawGridClass() TClass {
	ret := customDrawGridImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TCustomDrawGrid) DeleteColRow(IsColumn bool, index int32) {
	customDrawGridImportAPI().SysCallN(17, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) DeleteCol(Index int32) {
	customDrawGridImportAPI().SysCallN(16, m.Instance(), uintptr(Index))
}

func (m *TCustomDrawGrid) DeleteRow(Index int32) {
	customDrawGridImportAPI().SysCallN(18, m.Instance(), uintptr(Index))
}

func (m *TCustomDrawGrid) ExchangeColRow(IsColumn bool, index, WithIndex int32) {
	customDrawGridImportAPI().SysCallN(23, m.Instance(), PascalBool(IsColumn), uintptr(index), uintptr(WithIndex))
}

func (m *TCustomDrawGrid) InsertColRow(IsColumn bool, index int32) {
	customDrawGridImportAPI().SysCallN(40, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) MoveColRow(IsColumn bool, FromIndex, ToIndex int32) {
	customDrawGridImportAPI().SysCallN(43, m.Instance(), PascalBool(IsColumn), uintptr(FromIndex), uintptr(ToIndex))
}

func (m *TCustomDrawGrid) SortColRow(IsColumn bool, index int32) {
	customDrawGridImportAPI().SysCallN(93, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) SortColRow1(IsColumn bool, Index, FromIndex, ToIndex int32) {
	customDrawGridImportAPI().SysCallN(94, m.Instance(), PascalBool(IsColumn), uintptr(Index), uintptr(FromIndex), uintptr(ToIndex))
}

func (m *TCustomDrawGrid) DefaultDrawCell(aCol, aRow int32, aRect *TRect, aState TGridDrawState) {
	var result1 uintptr
	customDrawGridImportAPI().SysCallN(13, m.Instance(), uintptr(aCol), uintptr(aRow), uintptr(unsafePointer(&result1)), uintptr(aState))
	*aRect = *(*TRect)(getPointer(result1))
}

func (m *TCustomDrawGrid) SetOnAfterSelection(fn TOnSelectEvent) {
	if m.afterSelectionPtr != 0 {
		RemoveEventElement(m.afterSelectionPtr)
	}
	m.afterSelectionPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(55, m.Instance(), m.afterSelectionPtr)
}

func (m *TCustomDrawGrid) SetOnBeforeSelection(fn TOnSelectEvent) {
	if m.beforeSelectionPtr != 0 {
		RemoveEventElement(m.beforeSelectionPtr)
	}
	m.beforeSelectionPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(56, m.Instance(), m.beforeSelectionPtr)
}

func (m *TCustomDrawGrid) SetOnColRowDeleted(fn TGridOperationEvent) {
	if m.colRowDeletedPtr != 0 {
		RemoveEventElement(m.colRowDeletedPtr)
	}
	m.colRowDeletedPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(58, m.Instance(), m.colRowDeletedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowExchanged(fn TGridOperationEvent) {
	if m.colRowExchangedPtr != 0 {
		RemoveEventElement(m.colRowExchangedPtr)
	}
	m.colRowExchangedPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(59, m.Instance(), m.colRowExchangedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowInserted(fn TGridOperationEvent) {
	if m.colRowInsertedPtr != 0 {
		RemoveEventElement(m.colRowInsertedPtr)
	}
	m.colRowInsertedPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(60, m.Instance(), m.colRowInsertedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowMoved(fn TGridOperationEvent) {
	if m.colRowMovedPtr != 0 {
		RemoveEventElement(m.colRowMovedPtr)
	}
	m.colRowMovedPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(61, m.Instance(), m.colRowMovedPtr)
}

func (m *TCustomDrawGrid) SetOnCompareCells(fn TOnCompareCells) {
	if m.compareCellsPtr != 0 {
		RemoveEventElement(m.compareCellsPtr)
	}
	m.compareCellsPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(62, m.Instance(), m.compareCellsPtr)
}

func (m *TCustomDrawGrid) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(63, m.Instance(), m.contextPopupPtr)
}

func (m *TCustomDrawGrid) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(64, m.Instance(), m.dblClickPtr)
}

func (m *TCustomDrawGrid) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(65, m.Instance(), m.dragDropPtr)
}

func (m *TCustomDrawGrid) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(66, m.Instance(), m.dragOverPtr)
}

func (m *TCustomDrawGrid) SetOnDrawCell(fn TOnDrawCell) {
	if m.drawCellPtr != 0 {
		RemoveEventElement(m.drawCellPtr)
	}
	m.drawCellPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(67, m.Instance(), m.drawCellPtr)
}

func (m *TCustomDrawGrid) SetOnButtonClick(fn TOnSelectEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(57, m.Instance(), m.buttonClickPtr)
}

func (m *TCustomDrawGrid) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(68, m.Instance(), m.endDockPtr)
}

func (m *TCustomDrawGrid) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(69, m.Instance(), m.endDragPtr)
}

func (m *TCustomDrawGrid) SetOnGetEditMask(fn TGetEditEvent) {
	if m.getEditMaskPtr != 0 {
		RemoveEventElement(m.getEditMaskPtr)
	}
	m.getEditMaskPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(70, m.Instance(), m.getEditMaskPtr)
}

func (m *TCustomDrawGrid) SetOnGetEditText(fn TGetEditEvent) {
	if m.getEditTextPtr != 0 {
		RemoveEventElement(m.getEditTextPtr)
	}
	m.getEditTextPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(71, m.Instance(), m.getEditTextPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderClick(fn THdrEvent) {
	if m.headerClickPtr != 0 {
		RemoveEventElement(m.headerClickPtr)
	}
	m.headerClickPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(72, m.Instance(), m.headerClickPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderSized(fn THdrEvent) {
	if m.headerSizedPtr != 0 {
		RemoveEventElement(m.headerSizedPtr)
	}
	m.headerSizedPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(73, m.Instance(), m.headerSizedPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderSizing(fn THeaderSizingEvent) {
	if m.headerSizingPtr != 0 {
		RemoveEventElement(m.headerSizingPtr)
	}
	m.headerSizingPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(74, m.Instance(), m.headerSizingPtr)
}

func (m *TCustomDrawGrid) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(75, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomDrawGrid) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(76, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomDrawGrid) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(77, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomDrawGrid) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(78, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomDrawGrid) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(79, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(80, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(81, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(82, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomDrawGrid) SetOnPickListSelect(fn TNotifyEvent) {
	if m.pickListSelectPtr != 0 {
		RemoveEventElement(m.pickListSelectPtr)
	}
	m.pickListSelectPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(83, m.Instance(), m.pickListSelectPtr)
}

func (m *TCustomDrawGrid) SetOnPrepareCanvas(fn TOnPrepareCanvasEvent) {
	if m.prepareCanvasPtr != 0 {
		RemoveEventElement(m.prepareCanvasPtr)
	}
	m.prepareCanvasPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(84, m.Instance(), m.prepareCanvasPtr)
}

func (m *TCustomDrawGrid) SetOnSelectEditor(fn TSelectEditorEvent) {
	if m.selectEditorPtr != 0 {
		RemoveEventElement(m.selectEditorPtr)
	}
	m.selectEditorPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(86, m.Instance(), m.selectEditorPtr)
}

func (m *TCustomDrawGrid) SetOnSelection(fn TOnSelectEvent) {
	if m.selectionPtr != 0 {
		RemoveEventElement(m.selectionPtr)
	}
	m.selectionPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(87, m.Instance(), m.selectionPtr)
}

func (m *TCustomDrawGrid) SetOnSelectCell(fn TOnSelectCellEvent) {
	if m.selectCellPtr != 0 {
		RemoveEventElement(m.selectCellPtr)
	}
	m.selectCellPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(85, m.Instance(), m.selectCellPtr)
}

func (m *TCustomDrawGrid) SetOnSetEditText(fn TSetEditEvent) {
	if m.setEditTextPtr != 0 {
		RemoveEventElement(m.setEditTextPtr)
	}
	m.setEditTextPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(88, m.Instance(), m.setEditTextPtr)
}

func (m *TCustomDrawGrid) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(89, m.Instance(), m.startDockPtr)
}

func (m *TCustomDrawGrid) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(90, m.Instance(), m.startDragPtr)
}

func (m *TCustomDrawGrid) SetOnTopleftChanged(fn TNotifyEvent) {
	if m.topleftChangedPtr != 0 {
		RemoveEventElement(m.topleftChangedPtr)
	}
	m.topleftChangedPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(91, m.Instance(), m.topleftChangedPtr)
}

func (m *TCustomDrawGrid) SetOnValidateEntry(fn TValidateEntryEvent) {
	if m.validateEntryPtr != 0 {
		RemoveEventElement(m.validateEntryPtr)
	}
	m.validateEntryPtr = MakeEventDataPtr(fn)
	customDrawGridImportAPI().SysCallN(92, m.Instance(), m.validateEntryPtr)
}

var (
	customDrawGridImport       *imports.Imports = nil
	customDrawGridImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDrawGrid_AllowOutboundEvents", 0),
		/*1*/ imports.NewTable("CustomDrawGrid_AltColorStartNormal", 0),
		/*2*/ imports.NewTable("CustomDrawGrid_AutoAdvance", 0),
		/*3*/ imports.NewTable("CustomDrawGrid_AutoFillColumns", 0),
		/*4*/ imports.NewTable("CustomDrawGrid_BorderColor", 0),
		/*5*/ imports.NewTable("CustomDrawGrid_Class", 0),
		/*6*/ imports.NewTable("CustomDrawGrid_Col", 0),
		/*7*/ imports.NewTable("CustomDrawGrid_ColCount", 0),
		/*8*/ imports.NewTable("CustomDrawGrid_ColRow", 0),
		/*9*/ imports.NewTable("CustomDrawGrid_ColWidths", 0),
		/*10*/ imports.NewTable("CustomDrawGrid_Columns", 0),
		/*11*/ imports.NewTable("CustomDrawGrid_Create", 0),
		/*12*/ imports.NewTable("CustomDrawGrid_DefaultColWidth", 0),
		/*13*/ imports.NewTable("CustomDrawGrid_DefaultDrawCell", 0),
		/*14*/ imports.NewTable("CustomDrawGrid_DefaultDrawing", 0),
		/*15*/ imports.NewTable("CustomDrawGrid_DefaultRowHeight", 0),
		/*16*/ imports.NewTable("CustomDrawGrid_DeleteCol", 0),
		/*17*/ imports.NewTable("CustomDrawGrid_DeleteColRow", 0),
		/*18*/ imports.NewTable("CustomDrawGrid_DeleteRow", 0),
		/*19*/ imports.NewTable("CustomDrawGrid_DisabledFontColor", 0),
		/*20*/ imports.NewTable("CustomDrawGrid_Editor", 0),
		/*21*/ imports.NewTable("CustomDrawGrid_EditorBorderStyle", 0),
		/*22*/ imports.NewTable("CustomDrawGrid_EditorMode", 0),
		/*23*/ imports.NewTable("CustomDrawGrid_ExchangeColRow", 0),
		/*24*/ imports.NewTable("CustomDrawGrid_ExtendedColSizing", 0),
		/*25*/ imports.NewTable("CustomDrawGrid_FadeUnfocusedSelection", 0),
		/*26*/ imports.NewTable("CustomDrawGrid_FastEditing", 0),
		/*27*/ imports.NewTable("CustomDrawGrid_FixedColor", 0),
		/*28*/ imports.NewTable("CustomDrawGrid_FixedCols", 0),
		/*29*/ imports.NewTable("CustomDrawGrid_FixedGridLineColor", 0),
		/*30*/ imports.NewTable("CustomDrawGrid_FixedHotColor", 0),
		/*31*/ imports.NewTable("CustomDrawGrid_FixedRows", 0),
		/*32*/ imports.NewTable("CustomDrawGrid_Flat", 0),
		/*33*/ imports.NewTable("CustomDrawGrid_FocusColor", 0),
		/*34*/ imports.NewTable("CustomDrawGrid_FocusRectVisible", 0),
		/*35*/ imports.NewTable("CustomDrawGrid_GridHeight", 0),
		/*36*/ imports.NewTable("CustomDrawGrid_GridLineColor", 0),
		/*37*/ imports.NewTable("CustomDrawGrid_GridLineStyle", 0),
		/*38*/ imports.NewTable("CustomDrawGrid_GridLineWidth", 0),
		/*39*/ imports.NewTable("CustomDrawGrid_GridWidth", 0),
		/*40*/ imports.NewTable("CustomDrawGrid_InsertColRow", 0),
		/*41*/ imports.NewTable("CustomDrawGrid_IsCellSelected", 0),
		/*42*/ imports.NewTable("CustomDrawGrid_LeftCol", 0),
		/*43*/ imports.NewTable("CustomDrawGrid_MoveColRow", 0),
		/*44*/ imports.NewTable("CustomDrawGrid_Options", 0),
		/*45*/ imports.NewTable("CustomDrawGrid_Options2", 0),
		/*46*/ imports.NewTable("CustomDrawGrid_ParentShowHint", 0),
		/*47*/ imports.NewTable("CustomDrawGrid_Row", 0),
		/*48*/ imports.NewTable("CustomDrawGrid_RowCount", 0),
		/*49*/ imports.NewTable("CustomDrawGrid_RowHeights", 0),
		/*50*/ imports.NewTable("CustomDrawGrid_SaveOptions", 0),
		/*51*/ imports.NewTable("CustomDrawGrid_ScrollBars", 0),
		/*52*/ imports.NewTable("CustomDrawGrid_SelectedColor", 0),
		/*53*/ imports.NewTable("CustomDrawGrid_SelectedColumn", 0),
		/*54*/ imports.NewTable("CustomDrawGrid_Selection", 0),
		/*55*/ imports.NewTable("CustomDrawGrid_SetOnAfterSelection", 0),
		/*56*/ imports.NewTable("CustomDrawGrid_SetOnBeforeSelection", 0),
		/*57*/ imports.NewTable("CustomDrawGrid_SetOnButtonClick", 0),
		/*58*/ imports.NewTable("CustomDrawGrid_SetOnColRowDeleted", 0),
		/*59*/ imports.NewTable("CustomDrawGrid_SetOnColRowExchanged", 0),
		/*60*/ imports.NewTable("CustomDrawGrid_SetOnColRowInserted", 0),
		/*61*/ imports.NewTable("CustomDrawGrid_SetOnColRowMoved", 0),
		/*62*/ imports.NewTable("CustomDrawGrid_SetOnCompareCells", 0),
		/*63*/ imports.NewTable("CustomDrawGrid_SetOnContextPopup", 0),
		/*64*/ imports.NewTable("CustomDrawGrid_SetOnDblClick", 0),
		/*65*/ imports.NewTable("CustomDrawGrid_SetOnDragDrop", 0),
		/*66*/ imports.NewTable("CustomDrawGrid_SetOnDragOver", 0),
		/*67*/ imports.NewTable("CustomDrawGrid_SetOnDrawCell", 0),
		/*68*/ imports.NewTable("CustomDrawGrid_SetOnEndDock", 0),
		/*69*/ imports.NewTable("CustomDrawGrid_SetOnEndDrag", 0),
		/*70*/ imports.NewTable("CustomDrawGrid_SetOnGetEditMask", 0),
		/*71*/ imports.NewTable("CustomDrawGrid_SetOnGetEditText", 0),
		/*72*/ imports.NewTable("CustomDrawGrid_SetOnHeaderClick", 0),
		/*73*/ imports.NewTable("CustomDrawGrid_SetOnHeaderSized", 0),
		/*74*/ imports.NewTable("CustomDrawGrid_SetOnHeaderSizing", 0),
		/*75*/ imports.NewTable("CustomDrawGrid_SetOnMouseDown", 0),
		/*76*/ imports.NewTable("CustomDrawGrid_SetOnMouseEnter", 0),
		/*77*/ imports.NewTable("CustomDrawGrid_SetOnMouseLeave", 0),
		/*78*/ imports.NewTable("CustomDrawGrid_SetOnMouseMove", 0),
		/*79*/ imports.NewTable("CustomDrawGrid_SetOnMouseUp", 0),
		/*80*/ imports.NewTable("CustomDrawGrid_SetOnMouseWheel", 0),
		/*81*/ imports.NewTable("CustomDrawGrid_SetOnMouseWheelDown", 0),
		/*82*/ imports.NewTable("CustomDrawGrid_SetOnMouseWheelUp", 0),
		/*83*/ imports.NewTable("CustomDrawGrid_SetOnPickListSelect", 0),
		/*84*/ imports.NewTable("CustomDrawGrid_SetOnPrepareCanvas", 0),
		/*85*/ imports.NewTable("CustomDrawGrid_SetOnSelectCell", 0),
		/*86*/ imports.NewTable("CustomDrawGrid_SetOnSelectEditor", 0),
		/*87*/ imports.NewTable("CustomDrawGrid_SetOnSelection", 0),
		/*88*/ imports.NewTable("CustomDrawGrid_SetOnSetEditText", 0),
		/*89*/ imports.NewTable("CustomDrawGrid_SetOnStartDock", 0),
		/*90*/ imports.NewTable("CustomDrawGrid_SetOnStartDrag", 0),
		/*91*/ imports.NewTable("CustomDrawGrid_SetOnTopleftChanged", 0),
		/*92*/ imports.NewTable("CustomDrawGrid_SetOnValidateEntry", 0),
		/*93*/ imports.NewTable("CustomDrawGrid_SortColRow", 0),
		/*94*/ imports.NewTable("CustomDrawGrid_SortColRow1", 0),
		/*95*/ imports.NewTable("CustomDrawGrid_StrictSort", 0),
		/*96*/ imports.NewTable("CustomDrawGrid_TabAdvance", 0),
		/*97*/ imports.NewTable("CustomDrawGrid_TopRow", 0),
		/*98*/ imports.NewTable("CustomDrawGrid_UseXORFeatures", 0),
		/*99*/ imports.NewTable("CustomDrawGrid_VisibleColCount", 0),
		/*100*/ imports.NewTable("CustomDrawGrid_VisibleRowCount", 0),
	}
)

func customDrawGridImportAPI() *imports.Imports {
	if customDrawGridImport == nil {
		customDrawGridImport = NewDefaultImports()
		customDrawGridImport.SetImportTable(customDrawGridImportTables)
		customDrawGridImportTables = nil
	}
	return customDrawGridImport
}
