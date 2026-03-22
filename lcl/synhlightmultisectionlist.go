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
)

// ISynHLightMultiSectionList Parent: ISynEditStorageMem
type ISynHLightMultiSectionList interface {
	ISynEditStorageMem
	IndexOfFirstSectionAtLineIdx(lineIdx int32, charPos int32, useNext bool) int32 // function
	IndexOfFirstSectionAtVirtualIdx(lineIdx int32, getLastSection bool) int32      // function
	VirtualIdxToRealIdx(vLineIdx int32) int32                                      // function
	Debug()                                                                        // procedure
	Insert(anIndex int32, anSection TSynHLightMultiVirtualSection)                 // procedure
	Delete(anIndex int32)                                                          // procedure
	Sections(index int32) TSynHLightMultiVirtualSection                            // property Sections Getter
	SetSections(index int32, value TSynHLightMultiVirtualSection)                  // property Sections Setter
	PSections(index int32) TSynHLightMultiVirtualSection                           // property PSections Getter
}

type TSynHLightMultiSectionList struct {
	TSynEditStorageMem
}

func (m *TSynHLightMultiSectionList) IndexOfFirstSectionAtLineIdx(lineIdx int32, charPos int32, useNext bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHLightMultiSectionListAPI().SysCallN(1, m.Instance(), uintptr(lineIdx), uintptr(charPos), api.PasBool(useNext))
	return int32(r)
}

func (m *TSynHLightMultiSectionList) IndexOfFirstSectionAtVirtualIdx(lineIdx int32, getLastSection bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHLightMultiSectionListAPI().SysCallN(2, m.Instance(), uintptr(lineIdx), api.PasBool(getLastSection))
	return int32(r)
}

func (m *TSynHLightMultiSectionList) VirtualIdxToRealIdx(vLineIdx int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHLightMultiSectionListAPI().SysCallN(3, m.Instance(), uintptr(vLineIdx))
	return int32(r)
}

func (m *TSynHLightMultiSectionList) Debug() {
	if !m.IsValid() {
		return
	}
	synHLightMultiSectionListAPI().SysCallN(4, m.Instance())
}

func (m *TSynHLightMultiSectionList) Insert(anIndex int32, anSection TSynHLightMultiVirtualSection) {
	if !m.IsValid() {
		return
	}
	synHLightMultiSectionListAPI().SysCallN(5, m.Instance(), uintptr(anIndex), uintptr(base.UnsafePointer(&anSection)))
}

func (m *TSynHLightMultiSectionList) Delete(anIndex int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiSectionListAPI().SysCallN(6, m.Instance(), uintptr(anIndex))
}

func (m *TSynHLightMultiSectionList) Sections(index int32) (result TSynHLightMultiVirtualSection) {
	if !m.IsValid() {
		return
	}
	synHLightMultiSectionListAPI().SysCallN(7, 0, m.Instance(), uintptr(index), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynHLightMultiSectionList) SetSections(index int32, value TSynHLightMultiVirtualSection) {
	if !m.IsValid() {
		return
	}
	synHLightMultiSectionListAPI().SysCallN(7, 1, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynHLightMultiSectionList) PSections(index int32) (result TSynHLightMultiVirtualSection) {
	if !m.IsValid() {
		return
	}
	synHLightMultiSectionListAPI().SysCallN(8, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

// NewSynHLightMultiSectionList class constructor
func NewSynHLightMultiSectionList() ISynHLightMultiSectionList {
	r := synHLightMultiSectionListAPI().SysCallN(0)
	return AsSynHLightMultiSectionList(r)
}

var (
	synHLightMultiSectionListOnce   base.Once
	synHLightMultiSectionListImport *imports.Imports = nil
)

func synHLightMultiSectionListAPI() *imports.Imports {
	synHLightMultiSectionListOnce.Do(func() {
		synHLightMultiSectionListImport = api.NewDefaultImports()
		synHLightMultiSectionListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHLightMultiSectionList_Create", 0), // constructor NewSynHLightMultiSectionList
			/* 1 */ imports.NewTable("TSynHLightMultiSectionList_IndexOfFirstSectionAtLineIdx", 0), // function IndexOfFirstSectionAtLineIdx
			/* 2 */ imports.NewTable("TSynHLightMultiSectionList_IndexOfFirstSectionAtVirtualIdx", 0), // function IndexOfFirstSectionAtVirtualIdx
			/* 3 */ imports.NewTable("TSynHLightMultiSectionList_VirtualIdxToRealIdx", 0), // function VirtualIdxToRealIdx
			/* 4 */ imports.NewTable("TSynHLightMultiSectionList_Debug", 0), // procedure Debug
			/* 5 */ imports.NewTable("TSynHLightMultiSectionList_Insert", 0), // procedure Insert
			/* 6 */ imports.NewTable("TSynHLightMultiSectionList_Delete", 0), // procedure Delete
			/* 7 */ imports.NewTable("TSynHLightMultiSectionList_Sections", 0), // property Sections
			/* 8 */ imports.NewTable("TSynHLightMultiSectionList_PSections", 0), // property PSections
		}
	})
	return synHLightMultiSectionListImport
}
