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

// IHeaderSection Parent: ICollectionItem
type IHeaderSection interface {
	ICollectionItem
	Left() int32                         // property
	Right() int32                        // property
	State() THeaderSectionState          // property
	SetState(AValue THeaderSectionState) // property
	Alignment() TAlignment               // property
	SetAlignment(AValue TAlignment)      // property
	ImageIndex() TImageIndex             // property
	SetImageIndex(AValue TImageIndex)    // property
	MaxWidth() int32                     // property
	SetMaxWidth(AValue int32)            // property
	MinWidth() int32                     // property
	SetMinWidth(AValue int32)            // property
	Text() string                        // property
	SetText(AValue string)               // property
	Width() int32                        // property
	SetWidth(AValue int32)               // property
	Visible() bool                       // property
	SetVisible(AValue bool)              // property
	OriginalIndex() int32                // property
}

// THeaderSection Parent: TCollectionItem
type THeaderSection struct {
	TCollectionItem
}

func NewHeaderSection(ACollection ICollection) IHeaderSection {
	r1 := headerSectionImportAPI().SysCallN(2, GetObjectUintptr(ACollection))
	return AsHeaderSection(r1)
}

func (m *THeaderSection) Left() int32 {
	r1 := headerSectionImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *THeaderSection) Right() int32 {
	r1 := headerSectionImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func (m *THeaderSection) State() THeaderSectionState {
	r1 := headerSectionImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return THeaderSectionState(r1)
}

func (m *THeaderSection) SetState(AValue THeaderSectionState) {
	headerSectionImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) Alignment() TAlignment {
	r1 := headerSectionImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *THeaderSection) SetAlignment(AValue TAlignment) {
	headerSectionImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) ImageIndex() TImageIndex {
	r1 := headerSectionImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *THeaderSection) SetImageIndex(AValue TImageIndex) {
	headerSectionImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) MaxWidth() int32 {
	r1 := headerSectionImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *THeaderSection) SetMaxWidth(AValue int32) {
	headerSectionImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) MinWidth() int32 {
	r1 := headerSectionImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *THeaderSection) SetMinWidth(AValue int32) {
	headerSectionImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) Text() string {
	r1 := headerSectionImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *THeaderSection) SetText(AValue string) {
	headerSectionImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *THeaderSection) Width() int32 {
	r1 := headerSectionImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *THeaderSection) SetWidth(AValue int32) {
	headerSectionImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) Visible() bool {
	r1 := headerSectionImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *THeaderSection) SetVisible(AValue bool) {
	headerSectionImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *THeaderSection) OriginalIndex() int32 {
	r1 := headerSectionImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func HeaderSectionClass() TClass {
	ret := headerSectionImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	headerSectionImport       *imports.Imports = nil
	headerSectionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("HeaderSection_Alignment", 0),
		/*1*/ imports.NewTable("HeaderSection_Class", 0),
		/*2*/ imports.NewTable("HeaderSection_Create", 0),
		/*3*/ imports.NewTable("HeaderSection_ImageIndex", 0),
		/*4*/ imports.NewTable("HeaderSection_Left", 0),
		/*5*/ imports.NewTable("HeaderSection_MaxWidth", 0),
		/*6*/ imports.NewTable("HeaderSection_MinWidth", 0),
		/*7*/ imports.NewTable("HeaderSection_OriginalIndex", 0),
		/*8*/ imports.NewTable("HeaderSection_Right", 0),
		/*9*/ imports.NewTable("HeaderSection_State", 0),
		/*10*/ imports.NewTable("HeaderSection_Text", 0),
		/*11*/ imports.NewTable("HeaderSection_Visible", 0),
		/*12*/ imports.NewTable("HeaderSection_Width", 0),
	}
)

func headerSectionImportAPI() *imports.Imports {
	if headerSectionImport == nil {
		headerSectionImport = NewDefaultImports()
		headerSectionImport.SetImportTable(headerSectionImportTables)
		headerSectionImportTables = nil
	}
	return headerSectionImport
}
