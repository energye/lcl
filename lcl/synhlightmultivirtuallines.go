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

// ISynHLightMultiVirtualLines Parent: ISynEditStringsBase
type ISynHLightMultiVirtualLines interface {
	ISynEditStringsBase
	Debug()                                                                                                                  // procedure
	PrepareRegionScan(startLineIdx int32)                                                                                    // procedure
	FinishRegionScan(endLineIdx int32)                                                                                       // procedure
	RegionScanUpdateFirstRegionEnd(anEndPoint types.TPoint, tokenEndPos int32)                                               // procedure
	RegionScanUpdateOrInsertRegion(startPoint types.TPoint, anEndPoint types.TPoint, tokenStartPos int32, tokenEndPos int32) // procedure
	RegionScanUpdateLastRegionStart(startPoint types.TPoint, tokenStartPos int32, lineIndex int32)                           // procedure
	RealLinesInserted(index int32, count int32)                                                                              // procedure
	RealLinesDeleted(index int32, count int32)                                                                               // procedure
	RealLinesChanged(index int32, count int32)                                                                               // procedure
	ResetHLChangedLines()                                                                                                    // procedure
	FirstHLChangedLine() int32                                                                                               // property FirstHLChangedLine Getter
	LastHLChangedLine() int32                                                                                                // property LastHLChangedLine Getter
	SectionList() ISynHLightMultiSectionList                                                                                 // property SectionList Getter
	Scheme() ISynHighlighterMultiScheme                                                                                      // property Scheme Getter
	SetScheme(value ISynHighlighterMultiScheme)                                                                              // property Scheme Setter
}

type TSynHLightMultiVirtualLines struct {
	TSynEditStringsBase
}

func (m *TSynHLightMultiVirtualLines) Debug() {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(1, m.Instance())
}

func (m *TSynHLightMultiVirtualLines) PrepareRegionScan(startLineIdx int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(2, m.Instance(), uintptr(startLineIdx))
}

func (m *TSynHLightMultiVirtualLines) FinishRegionScan(endLineIdx int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(3, m.Instance(), uintptr(endLineIdx))
}

func (m *TSynHLightMultiVirtualLines) RegionScanUpdateFirstRegionEnd(anEndPoint types.TPoint, tokenEndPos int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&anEndPoint)), uintptr(tokenEndPos))
}

func (m *TSynHLightMultiVirtualLines) RegionScanUpdateOrInsertRegion(startPoint types.TPoint, anEndPoint types.TPoint, tokenStartPos int32, tokenEndPos int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&startPoint)), uintptr(base.UnsafePointer(&anEndPoint)), uintptr(tokenStartPos), uintptr(tokenEndPos))
}

func (m *TSynHLightMultiVirtualLines) RegionScanUpdateLastRegionStart(startPoint types.TPoint, tokenStartPos int32, lineIndex int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&startPoint)), uintptr(tokenStartPos), uintptr(lineIndex))
}

func (m *TSynHLightMultiVirtualLines) RealLinesInserted(index int32, count int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(7, m.Instance(), uintptr(index), uintptr(count))
}

func (m *TSynHLightMultiVirtualLines) RealLinesDeleted(index int32, count int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(8, m.Instance(), uintptr(index), uintptr(count))
}

func (m *TSynHLightMultiVirtualLines) RealLinesChanged(index int32, count int32) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(9, m.Instance(), uintptr(index), uintptr(count))
}

func (m *TSynHLightMultiVirtualLines) ResetHLChangedLines() {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(10, m.Instance())
}

func (m *TSynHLightMultiVirtualLines) FirstHLChangedLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHLightMultiVirtualLinesAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TSynHLightMultiVirtualLines) LastHLChangedLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHLightMultiVirtualLinesAPI().SysCallN(12, m.Instance())
	return int32(r)
}

func (m *TSynHLightMultiVirtualLines) SectionList() ISynHLightMultiSectionList {
	if !m.IsValid() {
		return nil
	}
	r := synHLightMultiVirtualLinesAPI().SysCallN(13, m.Instance())
	return AsSynHLightMultiSectionList(r)
}

func (m *TSynHLightMultiVirtualLines) Scheme() ISynHighlighterMultiScheme {
	if !m.IsValid() {
		return nil
	}
	r := synHLightMultiVirtualLinesAPI().SysCallN(14, 0, m.Instance())
	return AsSynHighlighterMultiScheme(r)
}

func (m *TSynHLightMultiVirtualLines) SetScheme(value ISynHighlighterMultiScheme) {
	if !m.IsValid() {
		return
	}
	synHLightMultiVirtualLinesAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynHLightMultiVirtualLines class constructor
func NewSynHLightMultiVirtualLines(lines ISynEditStringsBase) ISynHLightMultiVirtualLines {
	r := synHLightMultiVirtualLinesAPI().SysCallN(0, base.GetObjectUintptr(lines))
	return AsSynHLightMultiVirtualLines(r)
}

var (
	synHLightMultiVirtualLinesOnce   base.Once
	synHLightMultiVirtualLinesImport *imports.Imports = nil
)

func synHLightMultiVirtualLinesAPI() *imports.Imports {
	synHLightMultiVirtualLinesOnce.Do(func() {
		synHLightMultiVirtualLinesImport = api.NewDefaultImports()
		synHLightMultiVirtualLinesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHLightMultiVirtualLines_Create", 0), // constructor NewSynHLightMultiVirtualLines
			/* 1 */ imports.NewTable("TSynHLightMultiVirtualLines_Debug", 0), // procedure Debug
			/* 2 */ imports.NewTable("TSynHLightMultiVirtualLines_PrepareRegionScan", 0), // procedure PrepareRegionScan
			/* 3 */ imports.NewTable("TSynHLightMultiVirtualLines_FinishRegionScan", 0), // procedure FinishRegionScan
			/* 4 */ imports.NewTable("TSynHLightMultiVirtualLines_RegionScanUpdateFirstRegionEnd", 0), // procedure RegionScanUpdateFirstRegionEnd
			/* 5 */ imports.NewTable("TSynHLightMultiVirtualLines_RegionScanUpdateOrInsertRegion", 0), // procedure RegionScanUpdateOrInsertRegion
			/* 6 */ imports.NewTable("TSynHLightMultiVirtualLines_RegionScanUpdateLastRegionStart", 0), // procedure RegionScanUpdateLastRegionStart
			/* 7 */ imports.NewTable("TSynHLightMultiVirtualLines_RealLinesInserted", 0), // procedure RealLinesInserted
			/* 8 */ imports.NewTable("TSynHLightMultiVirtualLines_RealLinesDeleted", 0), // procedure RealLinesDeleted
			/* 9 */ imports.NewTable("TSynHLightMultiVirtualLines_RealLinesChanged", 0), // procedure RealLinesChanged
			/* 10 */ imports.NewTable("TSynHLightMultiVirtualLines_ResetHLChangedLines", 0), // procedure ResetHLChangedLines
			/* 11 */ imports.NewTable("TSynHLightMultiVirtualLines_FirstHLChangedLine", 0), // property FirstHLChangedLine
			/* 12 */ imports.NewTable("TSynHLightMultiVirtualLines_LastHLChangedLine", 0), // property LastHLChangedLine
			/* 13 */ imports.NewTable("TSynHLightMultiVirtualLines_SectionList", 0), // property SectionList
			/* 14 */ imports.NewTable("TSynHLightMultiVirtualLines_Scheme", 0), // property Scheme
		}
	})
	return synHLightMultiVirtualLinesImport
}
