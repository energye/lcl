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

// IVTHeader Parent: IPersistent
type IVTHeader interface {
	IPersistent
	AllowFocus(columnIndex int32) bool                                                                              // function
	InHeader(P types.TPoint) bool                                                                                   // function
	InHeaderSplitterArea(P types.TPoint) bool                                                                       // function
	ResizeColumns(changeBy int32, rangeStartCol int32, rangeEndCol int32, options types.TVTColumnOptions) int32     // function
	AutoFitColumns(animated bool, smartAutoFitType types.TSmartAutoFitType, rangeStartCol int32, rangeEndCol int32) // procedure
	Invalidate(column IVirtualTreeColumn, expandToBorder bool)                                                      // procedure
	LoadFromStream(stream IStream)                                                                                  // procedure
	RestoreColumns()                                                                                                // procedure
	SaveToStream(stream IStream)                                                                                    // procedure
	DragImage() IVTDragImage                                                                                        // property DragImage Getter
	States() types.THeaderStates                                                                                    // property States Getter
	Treeview() IBaseVirtualTree                                                                                     // property Treeview Getter
	UseColumns() bool                                                                                               // property UseColumns Getter
	AutoSizeIndex() int32                                                                                           // property AutoSizeIndex Getter
	SetAutoSizeIndex(value int32)                                                                                   // property AutoSizeIndex Setter
	Background() types.TColor                                                                                       // property Background Getter
	SetBackground(value types.TColor)                                                                               // property Background Setter
	Columns() IVirtualTreeColumns                                                                                   // property Columns Getter
	SetColumns(value IVirtualTreeColumns)                                                                           // property Columns Setter
	DefaultHeight() int32                                                                                           // property DefaultHeight Getter
	SetDefaultHeight(value int32)                                                                                   // property DefaultHeight Setter
	Font() IFont                                                                                                    // property Font Getter
	SetFont(value IFont)                                                                                            // property Font Setter
	FixedAreaConstraints() IVTFixedAreaConstraints                                                                  // property FixedAreaConstraints Getter
	SetFixedAreaConstraints(value IVTFixedAreaConstraints)                                                          // property FixedAreaConstraints Setter
	Height() int32                                                                                                  // property Height Getter
	SetHeight(value int32)                                                                                          // property Height Setter
	Images() ICustomImageList                                                                                       // property Images Getter
	SetImages(value ICustomImageList)                                                                               // property Images Setter
	MainColumn() int32                                                                                              // property MainColumn Getter
	SetMainColumn(value int32)                                                                                      // property MainColumn Setter
	MaxHeight() int32                                                                                               // property MaxHeight Getter
	SetMaxHeight(value int32)                                                                                       // property MaxHeight Setter
	MinHeight() int32                                                                                               // property MinHeight Getter
	SetMinHeight(value int32)                                                                                       // property MinHeight Setter
	Options() types.TVTHeaderOptions                                                                                // property Options Getter
	SetOptions(value types.TVTHeaderOptions)                                                                        // property Options Setter
	ParentFont() bool                                                                                               // property ParentFont Getter
	SetParentFont(value bool)                                                                                       // property ParentFont Setter
	PopupMenu() IPopupMenu                                                                                          // property PopupMenu Getter
	SetPopupMenu(value IPopupMenu)                                                                                  // property PopupMenu Setter
	SortColumn() int32                                                                                              // property SortColumn Getter
	SetSortColumn(value int32)                                                                                      // property SortColumn Setter
	SortDirection() types.TSortDirection                                                                            // property SortDirection Getter
	SetSortDirection(value types.TSortDirection)                                                                    // property SortDirection Setter
	Style() types.TVTHeaderStyle                                                                                    // property Style Getter
	SetStyle(value types.TVTHeaderStyle)                                                                            // property Style Setter
}

type TVTHeader struct {
	TPersistent
}

func (m *TVTHeader) AllowFocus(columnIndex int32) bool {
	if !m.IsValid() {
		return false
	}
	r := vTHeaderAPI().SysCallN(1, m.Instance(), uintptr(columnIndex))
	return api.GoBool(r)
}

func (m *TVTHeader) InHeader(P types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := vTHeaderAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return api.GoBool(r)
}

func (m *TVTHeader) InHeaderSplitterArea(P types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := vTHeaderAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return api.GoBool(r)
}

func (m *TVTHeader) ResizeColumns(changeBy int32, rangeStartCol int32, rangeEndCol int32, options types.TVTColumnOptions) int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(4, m.Instance(), uintptr(changeBy), uintptr(rangeStartCol), uintptr(rangeEndCol), uintptr(options))
	return int32(r)
}

func (m *TVTHeader) AutoFitColumns(animated bool, smartAutoFitType types.TSmartAutoFitType, rangeStartCol int32, rangeEndCol int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(5, m.Instance(), api.PasBool(animated), uintptr(smartAutoFitType), uintptr(rangeStartCol), uintptr(rangeEndCol))
}

func (m *TVTHeader) Invalidate(column IVirtualTreeColumn, expandToBorder bool) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(column), api.PasBool(expandToBorder))
}

func (m *TVTHeader) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TVTHeader) RestoreColumns() {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(8, m.Instance())
}

func (m *TVTHeader) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TVTHeader) DragImage() IVTDragImage {
	if !m.IsValid() {
		return nil
	}
	r := vTHeaderAPI().SysCallN(10, m.Instance())
	return AsVTDragImage(r)
}

func (m *TVTHeader) States() types.THeaderStates {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(11, m.Instance())
	return types.THeaderStates(r)
}

func (m *TVTHeader) Treeview() IBaseVirtualTree {
	if !m.IsValid() {
		return nil
	}
	r := vTHeaderAPI().SysCallN(12, m.Instance())
	return AsBaseVirtualTree(r)
}

func (m *TVTHeader) UseColumns() bool {
	if !m.IsValid() {
		return false
	}
	r := vTHeaderAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TVTHeader) AutoSizeIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TVTHeader) SetAutoSizeIndex(value int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) Background() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(15, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTHeader) SetBackground(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) Columns() IVirtualTreeColumns {
	if !m.IsValid() {
		return nil
	}
	r := vTHeaderAPI().SysCallN(16, 0, m.Instance())
	return AsVirtualTreeColumns(r)
}

func (m *TVTHeader) SetColumns(value IVirtualTreeColumns) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TVTHeader) DefaultHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TVTHeader) SetDefaultHeight(value int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) Font() IFont {
	if !m.IsValid() {
		return nil
	}
	r := vTHeaderAPI().SysCallN(18, 0, m.Instance())
	return AsFont(r)
}

func (m *TVTHeader) SetFont(value IFont) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TVTHeader) FixedAreaConstraints() IVTFixedAreaConstraints {
	if !m.IsValid() {
		return nil
	}
	r := vTHeaderAPI().SysCallN(19, 0, m.Instance())
	return AsVTFixedAreaConstraints(r)
}

func (m *TVTHeader) SetFixedAreaConstraints(value IVTFixedAreaConstraints) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(19, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TVTHeader) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TVTHeader) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := vTHeaderAPI().SysCallN(21, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TVTHeader) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TVTHeader) MainColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TVTHeader) SetMainColumn(value int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) MaxHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TVTHeader) SetMaxHeight(value int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) MinHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TVTHeader) SetMinHeight(value int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) Options() types.TVTHeaderOptions {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(25, 0, m.Instance())
	return types.TVTHeaderOptions(r)
}

func (m *TVTHeader) SetOptions(value types.TVTHeaderOptions) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := vTHeaderAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TVTHeader) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TVTHeader) PopupMenu() IPopupMenu {
	if !m.IsValid() {
		return nil
	}
	r := vTHeaderAPI().SysCallN(27, 0, m.Instance())
	return AsPopupMenu(r)
}

func (m *TVTHeader) SetPopupMenu(value IPopupMenu) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(27, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TVTHeader) SortColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(28, 0, m.Instance())
	return int32(r)
}

func (m *TVTHeader) SetSortColumn(value int32) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) SortDirection() types.TSortDirection {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(29, 0, m.Instance())
	return types.TSortDirection(r)
}

func (m *TVTHeader) SetSortDirection(value types.TSortDirection) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeader) Style() types.TVTHeaderStyle {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderAPI().SysCallN(30, 0, m.Instance())
	return types.TVTHeaderStyle(r)
}

func (m *TVTHeader) SetStyle(value types.TVTHeaderStyle) {
	if !m.IsValid() {
		return
	}
	vTHeaderAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

// NewVTHeader class constructor
func NewVTHeader(owner IBaseVirtualTree) IVTHeader {
	r := vTHeaderAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsVTHeader(r)
}

var (
	vTHeaderOnce   base.Once
	vTHeaderImport *imports.Imports = nil
)

func vTHeaderAPI() *imports.Imports {
	vTHeaderOnce.Do(func() {
		vTHeaderImport = api.NewDefaultImports()
		vTHeaderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTHeader_Create", 0), // constructor NewVTHeader
			/* 1 */ imports.NewTable("TVTHeader_AllowFocus", 0), // function AllowFocus
			/* 2 */ imports.NewTable("TVTHeader_InHeader", 0), // function InHeader
			/* 3 */ imports.NewTable("TVTHeader_InHeaderSplitterArea", 0), // function InHeaderSplitterArea
			/* 4 */ imports.NewTable("TVTHeader_ResizeColumns", 0), // function ResizeColumns
			/* 5 */ imports.NewTable("TVTHeader_AutoFitColumns", 0), // procedure AutoFitColumns
			/* 6 */ imports.NewTable("TVTHeader_Invalidate", 0), // procedure Invalidate
			/* 7 */ imports.NewTable("TVTHeader_LoadFromStream", 0), // procedure LoadFromStream
			/* 8 */ imports.NewTable("TVTHeader_RestoreColumns", 0), // procedure RestoreColumns
			/* 9 */ imports.NewTable("TVTHeader_SaveToStream", 0), // procedure SaveToStream
			/* 10 */ imports.NewTable("TVTHeader_DragImage", 0), // property DragImage
			/* 11 */ imports.NewTable("TVTHeader_States", 0), // property States
			/* 12 */ imports.NewTable("TVTHeader_Treeview", 0), // property Treeview
			/* 13 */ imports.NewTable("TVTHeader_UseColumns", 0), // property UseColumns
			/* 14 */ imports.NewTable("TVTHeader_AutoSizeIndex", 0), // property AutoSizeIndex
			/* 15 */ imports.NewTable("TVTHeader_Background", 0), // property Background
			/* 16 */ imports.NewTable("TVTHeader_Columns", 0), // property Columns
			/* 17 */ imports.NewTable("TVTHeader_DefaultHeight", 0), // property DefaultHeight
			/* 18 */ imports.NewTable("TVTHeader_Font", 0), // property Font
			/* 19 */ imports.NewTable("TVTHeader_FixedAreaConstraints", 0), // property FixedAreaConstraints
			/* 20 */ imports.NewTable("TVTHeader_Height", 0), // property Height
			/* 21 */ imports.NewTable("TVTHeader_Images", 0), // property Images
			/* 22 */ imports.NewTable("TVTHeader_MainColumn", 0), // property MainColumn
			/* 23 */ imports.NewTable("TVTHeader_MaxHeight", 0), // property MaxHeight
			/* 24 */ imports.NewTable("TVTHeader_MinHeight", 0), // property MinHeight
			/* 25 */ imports.NewTable("TVTHeader_Options", 0), // property Options
			/* 26 */ imports.NewTable("TVTHeader_ParentFont", 0), // property ParentFont
			/* 27 */ imports.NewTable("TVTHeader_PopupMenu", 0), // property PopupMenu
			/* 28 */ imports.NewTable("TVTHeader_SortColumn", 0), // property SortColumn
			/* 29 */ imports.NewTable("TVTHeader_SortDirection", 0), // property SortDirection
			/* 30 */ imports.NewTable("TVTHeader_Style", 0), // property Style
		}
	})
	return vTHeaderImport
}
