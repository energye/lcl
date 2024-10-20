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

// ICustomHeaderControl Parent: ICustomControl
type ICustomHeaderControl interface {
	ICustomControl
	SectionFromOriginalIndex(OriginalIndex int32) IHeaderSection // property
	DragReorder() bool                                           // property
	SetDragReorder(AValue bool)                                  // property
	Images() ICustomImageList                                    // property
	SetImages(AValue ICustomImageList)                           // property
	ImagesWidth() int32                                          // property
	SetImagesWidth(AValue int32)                                 // property
	Sections() IHeaderSections                                   // property
	SetSections(AValue IHeaderSections)                          // property
	GetSectionAt(P *TPoint) int32                                // function
	Click()                                                      // procedure
	DblClick()                                                   // procedure
	Paint()                                                      // procedure
	PaintSection(Index int32)                                    // procedure
	ChangeScale(M, D int32)                                      // procedure
	SetOnSectionDrag(fn TSectionDragEvent)                       // property event
	SetOnSectionEndDrag(fn TNotifyEvent)                         // property event
	SetOnSectionClick(fn TCustomSectionNotifyEvent)              // property event
	SetOnSectionResize(fn TCustomSectionNotifyEvent)             // property event
	SetOnSectionTrack(fn TCustomSectionTrackEvent)               // property event
	SetOnSectionSeparatorDblClick(fn TCustomSectionNotifyEvent)  // property event
	SetOnCreateSectionClass(fn TCustomHCCreateSectionClassEvent) // property event
}

// TCustomHeaderControl Parent: TCustomControl
type TCustomHeaderControl struct {
	TCustomControl
	sectionDragPtr              uintptr
	sectionEndDragPtr           uintptr
	sectionClickPtr             uintptr
	sectionResizePtr            uintptr
	sectionTrackPtr             uintptr
	sectionSeparatorDblClickPtr uintptr
	createSectionClassPtr       uintptr
}

func NewCustomHeaderControl(AOwner IComponent) ICustomHeaderControl {
	r1 := customHeaderControlImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsCustomHeaderControl(r1)
}

func (m *TCustomHeaderControl) SectionFromOriginalIndex(OriginalIndex int32) IHeaderSection {
	r1 := customHeaderControlImportAPI().SysCallN(11, m.Instance(), uintptr(OriginalIndex))
	return AsHeaderSection(r1)
}

func (m *TCustomHeaderControl) DragReorder() bool {
	r1 := customHeaderControlImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomHeaderControl) SetDragReorder(AValue bool) {
	customHeaderControlImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomHeaderControl) Images() ICustomImageList {
	r1 := customHeaderControlImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomHeaderControl) SetImages(AValue ICustomImageList) {
	customHeaderControlImportAPI().SysCallN(7, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomHeaderControl) ImagesWidth() int32 {
	r1 := customHeaderControlImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomHeaderControl) SetImagesWidth(AValue int32) {
	customHeaderControlImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomHeaderControl) Sections() IHeaderSections {
	r1 := customHeaderControlImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return AsHeaderSections(r1)
}

func (m *TCustomHeaderControl) SetSections(AValue IHeaderSections) {
	customHeaderControlImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomHeaderControl) GetSectionAt(P *TPoint) int32 {
	r1 := customHeaderControlImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(P)))
	return int32(r1)
}

func CustomHeaderControlClass() TClass {
	ret := customHeaderControlImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomHeaderControl) Click() {
	customHeaderControlImportAPI().SysCallN(2, m.Instance())
}

func (m *TCustomHeaderControl) DblClick() {
	customHeaderControlImportAPI().SysCallN(4, m.Instance())
}

func (m *TCustomHeaderControl) Paint() {
	customHeaderControlImportAPI().SysCallN(9, m.Instance())
}

func (m *TCustomHeaderControl) PaintSection(Index int32) {
	customHeaderControlImportAPI().SysCallN(10, m.Instance(), uintptr(Index))
}

func (m *TCustomHeaderControl) ChangeScale(M, D int32) {
	customHeaderControlImportAPI().SysCallN(0, m.Instance(), uintptr(M), uintptr(D))
}

func (m *TCustomHeaderControl) SetOnSectionDrag(fn TSectionDragEvent) {
	if m.sectionDragPtr != 0 {
		RemoveEventElement(m.sectionDragPtr)
	}
	m.sectionDragPtr = MakeEventDataPtr(fn)
	customHeaderControlImportAPI().SysCallN(15, m.Instance(), m.sectionDragPtr)
}

func (m *TCustomHeaderControl) SetOnSectionEndDrag(fn TNotifyEvent) {
	if m.sectionEndDragPtr != 0 {
		RemoveEventElement(m.sectionEndDragPtr)
	}
	m.sectionEndDragPtr = MakeEventDataPtr(fn)
	customHeaderControlImportAPI().SysCallN(16, m.Instance(), m.sectionEndDragPtr)
}

func (m *TCustomHeaderControl) SetOnSectionClick(fn TCustomSectionNotifyEvent) {
	if m.sectionClickPtr != 0 {
		RemoveEventElement(m.sectionClickPtr)
	}
	m.sectionClickPtr = MakeEventDataPtr(fn)
	customHeaderControlImportAPI().SysCallN(14, m.Instance(), m.sectionClickPtr)
}

func (m *TCustomHeaderControl) SetOnSectionResize(fn TCustomSectionNotifyEvent) {
	if m.sectionResizePtr != 0 {
		RemoveEventElement(m.sectionResizePtr)
	}
	m.sectionResizePtr = MakeEventDataPtr(fn)
	customHeaderControlImportAPI().SysCallN(17, m.Instance(), m.sectionResizePtr)
}

func (m *TCustomHeaderControl) SetOnSectionTrack(fn TCustomSectionTrackEvent) {
	if m.sectionTrackPtr != 0 {
		RemoveEventElement(m.sectionTrackPtr)
	}
	m.sectionTrackPtr = MakeEventDataPtr(fn)
	customHeaderControlImportAPI().SysCallN(19, m.Instance(), m.sectionTrackPtr)
}

func (m *TCustomHeaderControl) SetOnSectionSeparatorDblClick(fn TCustomSectionNotifyEvent) {
	if m.sectionSeparatorDblClickPtr != 0 {
		RemoveEventElement(m.sectionSeparatorDblClickPtr)
	}
	m.sectionSeparatorDblClickPtr = MakeEventDataPtr(fn)
	customHeaderControlImportAPI().SysCallN(18, m.Instance(), m.sectionSeparatorDblClickPtr)
}

func (m *TCustomHeaderControl) SetOnCreateSectionClass(fn TCustomHCCreateSectionClassEvent) {
	if m.createSectionClassPtr != 0 {
		RemoveEventElement(m.createSectionClassPtr)
	}
	m.createSectionClassPtr = MakeEventDataPtr(fn)
	customHeaderControlImportAPI().SysCallN(13, m.Instance(), m.createSectionClassPtr)
}

var (
	customHeaderControlImport       *imports.Imports = nil
	customHeaderControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomHeaderControl_ChangeScale", 0),
		/*1*/ imports.NewTable("CustomHeaderControl_Class", 0),
		/*2*/ imports.NewTable("CustomHeaderControl_Click", 0),
		/*3*/ imports.NewTable("CustomHeaderControl_Create", 0),
		/*4*/ imports.NewTable("CustomHeaderControl_DblClick", 0),
		/*5*/ imports.NewTable("CustomHeaderControl_DragReorder", 0),
		/*6*/ imports.NewTable("CustomHeaderControl_GetSectionAt", 0),
		/*7*/ imports.NewTable("CustomHeaderControl_Images", 0),
		/*8*/ imports.NewTable("CustomHeaderControl_ImagesWidth", 0),
		/*9*/ imports.NewTable("CustomHeaderControl_Paint", 0),
		/*10*/ imports.NewTable("CustomHeaderControl_PaintSection", 0),
		/*11*/ imports.NewTable("CustomHeaderControl_SectionFromOriginalIndex", 0),
		/*12*/ imports.NewTable("CustomHeaderControl_Sections", 0),
		/*13*/ imports.NewTable("CustomHeaderControl_SetOnCreateSectionClass", 0),
		/*14*/ imports.NewTable("CustomHeaderControl_SetOnSectionClick", 0),
		/*15*/ imports.NewTable("CustomHeaderControl_SetOnSectionDrag", 0),
		/*16*/ imports.NewTable("CustomHeaderControl_SetOnSectionEndDrag", 0),
		/*17*/ imports.NewTable("CustomHeaderControl_SetOnSectionResize", 0),
		/*18*/ imports.NewTable("CustomHeaderControl_SetOnSectionSeparatorDblClick", 0),
		/*19*/ imports.NewTable("CustomHeaderControl_SetOnSectionTrack", 0),
	}
)

func customHeaderControlImportAPI() *imports.Imports {
	if customHeaderControlImport == nil {
		customHeaderControlImport = NewDefaultImports()
		customHeaderControlImport.SetImportTable(customHeaderControlImportTables)
		customHeaderControlImportTables = nil
	}
	return customHeaderControlImport
}
