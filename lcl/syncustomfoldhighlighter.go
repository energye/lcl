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

// ISynCustomFoldHighlighter Parent: ISynCustomHighlighter
type ISynCustomFoldHighlighter interface {
	ISynCustomHighlighter
	// FoldBlockOpeningCountWithLineIdxSynFoldBlockFilter
	//  Info about Folds
	FoldBlockOpeningCountWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32 // function
	FoldBlockClosingCountWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32 // function
	FoldBlockEndLevelWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32     // function
	FoldBlockMinLevelWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32     // function
	// FoldBlockNestedTypesWithLIdxIntPointerSFBFilter
	//  * All nested FoldType (cfbtBegin) if available. Similar to TopCodeFoldBlockType
	//  - Index=0 is most outer / Index=FoldBlockEndLevel is most inner (TopCodeFoldBlockType 0=inner)
	//  - False, if it can not be determined for the filter settings
	FoldBlockNestedTypesWithLIdxIntPointerSFBFilter(lineIndex types.TLineIdx, nestIndex int32, outType_ *uintptr, filter TSynFoldBlockFilter) bool                              // function
	FoldBlockOpeningCountWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32                                            // function
	FoldBlockClosingCountWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32                                            // function
	FoldBlockEndLevelWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32                                                // function
	FoldBlockMinLevelWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32                                                // function
	FoldBlockNestedTypesWithLIdxIntX2PointerSFBFFlags(lineIndex types.TLineIdx, nestIndex int32, outType_ *uintptr, foldGroup int32, flags types.TSynFoldBlockFilterFlags) bool // function
	FoldTypeCount() int32                                                                                                                                                       // function
	FoldTypeAtNodeIndex(lineIndex int32, foldIndex int32, useCloseNodes bool) int32                                                                                             // function
	FindNextLineWithMinFoldLevelWithLIdxIntSFBFilter(lineIndex types.TLineIdx, searchLevel int32, filter TSynFoldBlockFilter) int32                                             // function
	FindNextLineWithMinFoldLevelWithLIdxIntX2SFBFFlags(lineIndex types.TLineIdx, searchLevel int32, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32                // function
	FoldEndLineWithIntX2(lineIndex int32, foldIndex int32) int32                                                                                                                // function
	FoldEndLineWithIntX2SFBFilter(lineIndex int32, foldIndex int32, filter TSynFoldBlockFilter) int32                                                                           // function
	FoldEndLineWithIntX3SFBFFlags(lineIndex int32, foldIndex int32, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32                                                // function
	// FoldLineLength
	//  AFoldGroup: integer = 0; AFlags: TSynFoldBlockFilterFlags = []): integer; overload;
	FoldLineLength(lineIndex int32, foldIndex int32) int32                  // function
	PerformScan(startIndex int32, endIndex int32, forceEndIndex bool) int32 // function
	DoCurrentLinesChanged()                                                 // procedure
	// FoldNodeInfo
	//  All fold-nodes
	//  FoldNodeInfo: Returns a shared object
	//  Adding RefCount, will prevent others from getting further copies, but not from using copies they already have.
	//  If not adding refcount, the object should not be stored/re-used
	//  Not adding ref-count, should only be done for CountEx, NodeInfoEx
	FoldNodeInfo(line types.TLineIdx) ILazSynFoldNodeInfoList // property FoldNodeInfo Getter
	FoldConfig(index int32) ISynCustomFoldConfig              // property FoldConfig Getter
	SetFoldConfig(index int32, value ISynCustomFoldConfig)    // property FoldConfig Setter
	FoldConfigCount() int32                                   // property FoldConfigCount Getter
}

type TSynCustomFoldHighlighter struct {
	TSynCustomHighlighter
}

func (m *TSynCustomFoldHighlighter) FoldBlockOpeningCountWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(1, m.Instance(), uintptr(lineIndex), uintptr(base.UnsafePointer(&filter)))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockClosingCountWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(2, m.Instance(), uintptr(lineIndex), uintptr(base.UnsafePointer(&filter)))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockEndLevelWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(3, m.Instance(), uintptr(lineIndex), uintptr(base.UnsafePointer(&filter)))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockMinLevelWithLineIdxSynFoldBlockFilter(lineIndex types.TLineIdx, filter TSynFoldBlockFilter) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(4, m.Instance(), uintptr(lineIndex), uintptr(base.UnsafePointer(&filter)))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockNestedTypesWithLIdxIntPointerSFBFilter(lineIndex types.TLineIdx, nestIndex int32, outType_ *uintptr, filter TSynFoldBlockFilter) bool {
	if !m.IsValid() {
		return false
	}
	var type_Ptr uintptr
	r := synCustomFoldHighlighterAPI().SysCallN(5, m.Instance(), uintptr(lineIndex), uintptr(nestIndex), uintptr(base.UnsafePointer(&type_Ptr)), uintptr(base.UnsafePointer(&filter)))
	*outType_ = uintptr(type_Ptr)
	return api.GoBool(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockOpeningCountWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(6, m.Instance(), uintptr(lineIndex), uintptr(foldGroup), uintptr(flags))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockClosingCountWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(7, m.Instance(), uintptr(lineIndex), uintptr(foldGroup), uintptr(flags))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockEndLevelWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(8, m.Instance(), uintptr(lineIndex), uintptr(foldGroup), uintptr(flags))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockMinLevelWithLIdxIntSFBFFlags(lineIndex types.TLineIdx, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(9, m.Instance(), uintptr(lineIndex), uintptr(foldGroup), uintptr(flags))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldBlockNestedTypesWithLIdxIntX2PointerSFBFFlags(lineIndex types.TLineIdx, nestIndex int32, outType_ *uintptr, foldGroup int32, flags types.TSynFoldBlockFilterFlags) bool {
	if !m.IsValid() {
		return false
	}
	var type_Ptr uintptr
	r := synCustomFoldHighlighterAPI().SysCallN(10, m.Instance(), uintptr(lineIndex), uintptr(nestIndex), uintptr(base.UnsafePointer(&type_Ptr)), uintptr(foldGroup), uintptr(flags))
	*outType_ = uintptr(type_Ptr)
	return api.GoBool(r)
}

func (m *TSynCustomFoldHighlighter) FoldTypeCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldTypeAtNodeIndex(lineIndex int32, foldIndex int32, useCloseNodes bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(12, m.Instance(), uintptr(lineIndex), uintptr(foldIndex), api.PasBool(useCloseNodes))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FindNextLineWithMinFoldLevelWithLIdxIntSFBFilter(lineIndex types.TLineIdx, searchLevel int32, filter TSynFoldBlockFilter) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(13, m.Instance(), uintptr(lineIndex), uintptr(searchLevel), uintptr(base.UnsafePointer(&filter)))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FindNextLineWithMinFoldLevelWithLIdxIntX2SFBFFlags(lineIndex types.TLineIdx, searchLevel int32, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(14, m.Instance(), uintptr(lineIndex), uintptr(searchLevel), uintptr(foldGroup), uintptr(flags))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldEndLineWithIntX2(lineIndex int32, foldIndex int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(15, m.Instance(), uintptr(lineIndex), uintptr(foldIndex))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldEndLineWithIntX2SFBFilter(lineIndex int32, foldIndex int32, filter TSynFoldBlockFilter) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(16, m.Instance(), uintptr(lineIndex), uintptr(foldIndex), uintptr(base.UnsafePointer(&filter)))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldEndLineWithIntX3SFBFFlags(lineIndex int32, foldIndex int32, foldGroup int32, flags types.TSynFoldBlockFilterFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(17, m.Instance(), uintptr(lineIndex), uintptr(foldIndex), uintptr(foldGroup), uintptr(flags))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) FoldLineLength(lineIndex int32, foldIndex int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(18, m.Instance(), uintptr(lineIndex), uintptr(foldIndex))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) PerformScan(startIndex int32, endIndex int32, forceEndIndex bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(19, m.Instance(), uintptr(startIndex), uintptr(endIndex), api.PasBool(forceEndIndex))
	return int32(r)
}

func (m *TSynCustomFoldHighlighter) DoCurrentLinesChanged() {
	if !m.IsValid() {
		return
	}
	synCustomFoldHighlighterAPI().SysCallN(20, m.Instance())
}

func (m *TSynCustomFoldHighlighter) FoldNodeInfo(line types.TLineIdx) ILazSynFoldNodeInfoList {
	if !m.IsValid() {
		return nil
	}
	r := synCustomFoldHighlighterAPI().SysCallN(21, m.Instance(), uintptr(line))
	return AsLazSynFoldNodeInfoList(r)
}

func (m *TSynCustomFoldHighlighter) FoldConfig(index int32) ISynCustomFoldConfig {
	if !m.IsValid() {
		return nil
	}
	r := synCustomFoldHighlighterAPI().SysCallN(22, 0, m.Instance(), uintptr(index))
	return AsSynCustomFoldConfig(r)
}

func (m *TSynCustomFoldHighlighter) SetFoldConfig(index int32, value ISynCustomFoldConfig) {
	if !m.IsValid() {
		return
	}
	synCustomFoldHighlighterAPI().SysCallN(22, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TSynCustomFoldHighlighter) FoldConfigCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldHighlighterAPI().SysCallN(23, m.Instance())
	return int32(r)
}

// NewSynCustomFoldHighlighter class constructor
func NewSynCustomFoldHighlighter(owner IComponent) ISynCustomFoldHighlighter {
	r := synCustomFoldHighlighterAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynCustomFoldHighlighter(r)
}

func TSynCustomFoldHighlighterClass() types.TClass {
	r := synCustomFoldHighlighterAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	synCustomFoldHighlighterOnce   base.Once
	synCustomFoldHighlighterImport *imports.Imports = nil
)

func synCustomFoldHighlighterAPI() *imports.Imports {
	synCustomFoldHighlighterOnce.Do(func() {
		synCustomFoldHighlighterImport = api.NewDefaultImports()
		synCustomFoldHighlighterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCustomFoldHighlighter_Create", 0), // constructor NewSynCustomFoldHighlighter
			/* 1 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockOpeningCountWithLineIdxSynFoldBlockFilter", 0), // function FoldBlockOpeningCountWithLineIdxSynFoldBlockFilter
			/* 2 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockClosingCountWithLineIdxSynFoldBlockFilter", 0), // function FoldBlockClosingCountWithLineIdxSynFoldBlockFilter
			/* 3 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockEndLevelWithLineIdxSynFoldBlockFilter", 0), // function FoldBlockEndLevelWithLineIdxSynFoldBlockFilter
			/* 4 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockMinLevelWithLineIdxSynFoldBlockFilter", 0), // function FoldBlockMinLevelWithLineIdxSynFoldBlockFilter
			/* 5 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockNestedTypesWithLIdxIntPointerSFBFilter", 0), // function FoldBlockNestedTypesWithLIdxIntPointerSFBFilter
			/* 6 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockOpeningCountWithLIdxIntSFBFFlags", 0), // function FoldBlockOpeningCountWithLIdxIntSFBFFlags
			/* 7 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockClosingCountWithLIdxIntSFBFFlags", 0), // function FoldBlockClosingCountWithLIdxIntSFBFFlags
			/* 8 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockEndLevelWithLIdxIntSFBFFlags", 0), // function FoldBlockEndLevelWithLIdxIntSFBFFlags
			/* 9 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockMinLevelWithLIdxIntSFBFFlags", 0), // function FoldBlockMinLevelWithLIdxIntSFBFFlags
			/* 10 */ imports.NewTable("TSynCustomFoldHighlighter_FoldBlockNestedTypesWithLIdxIntX2PointerSFBFFlags", 0), // function FoldBlockNestedTypesWithLIdxIntX2PointerSFBFFlags
			/* 11 */ imports.NewTable("TSynCustomFoldHighlighter_FoldTypeCount", 0), // function FoldTypeCount
			/* 12 */ imports.NewTable("TSynCustomFoldHighlighter_FoldTypeAtNodeIndex", 0), // function FoldTypeAtNodeIndex
			/* 13 */ imports.NewTable("TSynCustomFoldHighlighter_FindNextLineWithMinFoldLevelWithLIdxIntSFBFilter", 0), // function FindNextLineWithMinFoldLevelWithLIdxIntSFBFilter
			/* 14 */ imports.NewTable("TSynCustomFoldHighlighter_FindNextLineWithMinFoldLevelWithLIdxIntX2SFBFFlags", 0), // function FindNextLineWithMinFoldLevelWithLIdxIntX2SFBFFlags
			/* 15 */ imports.NewTable("TSynCustomFoldHighlighter_FoldEndLineWithIntX2", 0), // function FoldEndLineWithIntX2
			/* 16 */ imports.NewTable("TSynCustomFoldHighlighter_FoldEndLineWithIntX2SFBFilter", 0), // function FoldEndLineWithIntX2SFBFilter
			/* 17 */ imports.NewTable("TSynCustomFoldHighlighter_FoldEndLineWithIntX3SFBFFlags", 0), // function FoldEndLineWithIntX3SFBFFlags
			/* 18 */ imports.NewTable("TSynCustomFoldHighlighter_FoldLineLength", 0), // function FoldLineLength
			/* 19 */ imports.NewTable("TSynCustomFoldHighlighter_PerformScan", 0), // function PerformScan
			/* 20 */ imports.NewTable("TSynCustomFoldHighlighter_DoCurrentLinesChanged", 0), // procedure DoCurrentLinesChanged
			/* 21 */ imports.NewTable("TSynCustomFoldHighlighter_FoldNodeInfo", 0), // property FoldNodeInfo
			/* 22 */ imports.NewTable("TSynCustomFoldHighlighter_FoldConfig", 0), // property FoldConfig
			/* 23 */ imports.NewTable("TSynCustomFoldHighlighter_FoldConfigCount", 0), // property FoldConfigCount
			/* 24 */ imports.NewTable("TSynCustomFoldHighlighter_TClass", 0), // function TSynCustomFoldHighlighterClass
		}
	})
	return synCustomFoldHighlighterImport
}
