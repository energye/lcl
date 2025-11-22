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

// ICustomHeaderControl Parent: ICustomControl
type ICustomHeaderControl interface {
	ICustomControl
	GetSectionAt(P types.TPoint) int32                           // function
	Click()                                                      // procedure
	DblClick()                                                   // procedure
	Paint()                                                      // procedure
	PaintSection(index int32)                                    // procedure
	ChangeScale(M int32, D int32)                                // procedure
	SectionFromOriginalIndex(originalIndex int32) IHeaderSection // property SectionFromOriginalIndex Getter
	DragReorder() bool                                           // property DragReorder Getter
	SetDragReorder(value bool)                                   // property DragReorder Setter
	Images() ICustomImageList                                    // property Images Getter
	SetImages(value ICustomImageList)                            // property Images Setter
	ImagesWidth() int32                                          // property ImagesWidth Getter
	SetImagesWidth(value int32)                                  // property ImagesWidth Setter
	Sections() IHeaderSections                                   // property Sections Getter
	SetSections(value IHeaderSections)                           // property Sections Setter
	SetOnSectionDrag(fn TSectionDragEvent)                       // property event
	SetOnSectionEndDrag(fn TNotifyEvent)                         // property event
	SetOnSectionClick(fn TCustomSectionNotifyEvent)              // property event
	SetOnSectionResize(fn TCustomSectionNotifyEvent)             // property event
	SetOnSectionTrack(fn TCustomSectionTrackEvent)               // property event
	SetOnSectionSeparatorDblClick(fn TCustomSectionNotifyEvent)  // property event
	SetOnCreateSectionClass(fn TCustomHCCreateSectionClassEvent) // property event
}

type TCustomHeaderControl struct {
	TCustomControl
}

func (m *TCustomHeaderControl) GetSectionAt(P types.TPoint) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customHeaderControlAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return int32(r)
}

func (m *TCustomHeaderControl) Click() {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(2, m.Instance())
}

func (m *TCustomHeaderControl) DblClick() {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(3, m.Instance())
}

func (m *TCustomHeaderControl) Paint() {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(4, m.Instance())
}

func (m *TCustomHeaderControl) PaintSection(index int32) {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(5, m.Instance(), uintptr(index))
}

func (m *TCustomHeaderControl) ChangeScale(M int32, D int32) {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(6, m.Instance(), uintptr(M), uintptr(D))
}

func (m *TCustomHeaderControl) SectionFromOriginalIndex(originalIndex int32) IHeaderSection {
	if !m.IsValid() {
		return nil
	}
	r := customHeaderControlAPI().SysCallN(7, m.Instance(), uintptr(originalIndex))
	return AsHeaderSection(r)
}

func (m *TCustomHeaderControl) DragReorder() bool {
	if !m.IsValid() {
		return false
	}
	r := customHeaderControlAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomHeaderControl) SetDragReorder(value bool) {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomHeaderControl) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customHeaderControlAPI().SysCallN(9, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomHeaderControl) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomHeaderControl) ImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customHeaderControlAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCustomHeaderControl) SetImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomHeaderControl) Sections() IHeaderSections {
	if !m.IsValid() {
		return nil
	}
	r := customHeaderControlAPI().SysCallN(11, 0, m.Instance())
	return AsHeaderSections(r)
}

func (m *TCustomHeaderControl) SetSections(value IHeaderSections) {
	if !m.IsValid() {
		return
	}
	customHeaderControlAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomHeaderControl) SetOnSectionDrag(fn TSectionDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSectionDragEvent(fn)
	base.SetEvent(m, 12, customHeaderControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomHeaderControl) SetOnSectionEndDrag(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, customHeaderControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomHeaderControl) SetOnSectionClick(fn TCustomSectionNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomSectionNotifyEvent(fn)
	base.SetEvent(m, 14, customHeaderControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomHeaderControl) SetOnSectionResize(fn TCustomSectionNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomSectionNotifyEvent(fn)
	base.SetEvent(m, 15, customHeaderControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomHeaderControl) SetOnSectionTrack(fn TCustomSectionTrackEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomSectionTrackEvent(fn)
	base.SetEvent(m, 16, customHeaderControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomHeaderControl) SetOnSectionSeparatorDblClick(fn TCustomSectionNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomSectionNotifyEvent(fn)
	base.SetEvent(m, 17, customHeaderControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomHeaderControl) SetOnCreateSectionClass(fn TCustomHCCreateSectionClassEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomHCCreateSectionClassEvent(fn)
	base.SetEvent(m, 18, customHeaderControlAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomHeaderControl class constructor
func NewCustomHeaderControl(owner IComponent) ICustomHeaderControl {
	r := customHeaderControlAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomHeaderControl(r)
}

func TCustomHeaderControlClass() types.TClass {
	r := customHeaderControlAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	customHeaderControlOnce   base.Once
	customHeaderControlImport *imports.Imports = nil
)

func customHeaderControlAPI() *imports.Imports {
	customHeaderControlOnce.Do(func() {
		customHeaderControlImport = api.NewDefaultImports()
		customHeaderControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomHeaderControl_Create", 0), // constructor NewCustomHeaderControl
			/* 1 */ imports.NewTable("TCustomHeaderControl_GetSectionAt", 0), // function GetSectionAt
			/* 2 */ imports.NewTable("TCustomHeaderControl_Click", 0), // procedure Click
			/* 3 */ imports.NewTable("TCustomHeaderControl_DblClick", 0), // procedure DblClick
			/* 4 */ imports.NewTable("TCustomHeaderControl_Paint", 0), // procedure Paint
			/* 5 */ imports.NewTable("TCustomHeaderControl_PaintSection", 0), // procedure PaintSection
			/* 6 */ imports.NewTable("TCustomHeaderControl_ChangeScale", 0), // procedure ChangeScale
			/* 7 */ imports.NewTable("TCustomHeaderControl_SectionFromOriginalIndex", 0), // property SectionFromOriginalIndex
			/* 8 */ imports.NewTable("TCustomHeaderControl_DragReorder", 0), // property DragReorder
			/* 9 */ imports.NewTable("TCustomHeaderControl_Images", 0), // property Images
			/* 10 */ imports.NewTable("TCustomHeaderControl_ImagesWidth", 0), // property ImagesWidth
			/* 11 */ imports.NewTable("TCustomHeaderControl_Sections", 0), // property Sections
			/* 12 */ imports.NewTable("TCustomHeaderControl_OnSectionDrag", 0), // event OnSectionDrag
			/* 13 */ imports.NewTable("TCustomHeaderControl_OnSectionEndDrag", 0), // event OnSectionEndDrag
			/* 14 */ imports.NewTable("TCustomHeaderControl_OnSectionClick", 0), // event OnSectionClick
			/* 15 */ imports.NewTable("TCustomHeaderControl_OnSectionResize", 0), // event OnSectionResize
			/* 16 */ imports.NewTable("TCustomHeaderControl_OnSectionTrack", 0), // event OnSectionTrack
			/* 17 */ imports.NewTable("TCustomHeaderControl_OnSectionSeparatorDblClick", 0), // event OnSectionSeparatorDblClick
			/* 18 */ imports.NewTable("TCustomHeaderControl_OnCreateSectionClass", 0), // event OnCreateSectionClass
			/* 19 */ imports.NewTable("TCustomHeaderControl_TClass", 0), // function TCustomHeaderControlClass
		}
	})
	return customHeaderControlImport
}
