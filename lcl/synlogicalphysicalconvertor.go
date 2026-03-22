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

// ISynLogicalPhysicalConvertor Parent: IObject
type ISynLogicalPhysicalConvertor interface {
	IObject
	// PhysicalToLogicalWithIntX2
	//  Line is 0-based // Column is 1-based
	PhysicalToLogicalWithIntX2(index int32, column int32) int32                                                                                                    // function
	PhysicalToLogicalWithIntX3SPCSideSLPFlags(index int32, column int32, outColOffset *int32, charSide types.TSynPhysCharSide, flags types.TSynLogPhysFlags) int32 // function
	// LogicalToPhysicalWithIntX2
	//  ACharPos 1=before 1st char
	LogicalToPhysicalWithIntX2(index int32, bytePos int32) int32                                                                                                // function
	LogicalToPhysicalWithIntX3SLCSideSLPFlags(index int32, bytePos int32, colOffset *int32, charSide types.TSynLogCharSide, flags types.TSynLogPhysFlags) int32 // function
	CurrentLine() int32                                                                                                                                         // property CurrentLine Getter
	SetCurrentLine(value int32)                                                                                                                                 // property CurrentLine Setter
	CurrentWidths() PPhysicalCharWidth                                                                                                                          // property CurrentWidths Getter
	CurrentWidthsDirect() IByteArrayWrap                                                                                                                        // property CurrentWidthsDirect Getter
	CurrentWidthsCount() int32                                                                                                                                  // property CurrentWidthsCount Getter
	// AdjustedLogToPhysOrigin
	//  properties set, if lpfAdjustTo... is used
	//  By LogToPhys
	AdjustedLogToPhysOrigin() int32 // property AdjustedLogToPhysOrigin Getter
	// AdjustedPhysToLogOrigin
	//  By PhysToLog
	AdjustedPhysToLogOrigin() int32    // property AdjustedPhysToLogOrigin Getter
	UnAdjustedPhysToLogResult() int32  // property UnAdjustedPhysToLogResult Getter
	UnAdjustedPhysToLogColOffs() int32 // property UnAdjustedPhysToLogColOffs Getter
}

type TSynLogicalPhysicalConvertor struct {
	TObject
}

func (m *TSynLogicalPhysicalConvertor) PhysicalToLogicalWithIntX2(index int32, column int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(1, m.Instance(), uintptr(index), uintptr(column))
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) PhysicalToLogicalWithIntX3SPCSideSLPFlags(index int32, column int32, outColOffset *int32, charSide types.TSynPhysCharSide, flags types.TSynLogPhysFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	var colOffsetPtr uintptr
	r := synLogicalPhysicalConvertorAPI().SysCallN(2, m.Instance(), uintptr(index), uintptr(column), uintptr(base.UnsafePointer(&colOffsetPtr)), uintptr(charSide), uintptr(flags))
	*outColOffset = int32(colOffsetPtr)
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) LogicalToPhysicalWithIntX2(index int32, bytePos int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(bytePos))
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) LogicalToPhysicalWithIntX3SLCSideSLPFlags(index int32, bytePos int32, colOffset *int32, charSide types.TSynLogCharSide, flags types.TSynLogPhysFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	colOffsetPtr := uintptr(*colOffset)
	r := synLogicalPhysicalConvertorAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(bytePos), uintptr(base.UnsafePointer(&colOffsetPtr)), uintptr(charSide), uintptr(flags))
	*colOffset = int32(colOffsetPtr)
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) CurrentLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) SetCurrentLine(value int32) {
	if !m.IsValid() {
		return
	}
	synLogicalPhysicalConvertorAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TSynLogicalPhysicalConvertor) CurrentWidths() PPhysicalCharWidth {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(6, m.Instance())
	return PPhysicalCharWidth(r)
}

func (m *TSynLogicalPhysicalConvertor) CurrentWidthsDirect() IByteArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(7, m.Instance())
	return AsByteArrayWrap(r)
}

func (m *TSynLogicalPhysicalConvertor) CurrentWidthsCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) AdjustedLogToPhysOrigin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) AdjustedPhysToLogOrigin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) UnAdjustedPhysToLogResult() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TSynLogicalPhysicalConvertor) UnAdjustedPhysToLogColOffs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synLogicalPhysicalConvertorAPI().SysCallN(12, m.Instance())
	return int32(r)
}

// NewSynLogicalPhysicalConvertor class constructor
func NewSynLogicalPhysicalConvertor(lines ISynEditStrings) ISynLogicalPhysicalConvertor {
	r := synLogicalPhysicalConvertorAPI().SysCallN(0, base.GetObjectUintptr(lines))
	return AsSynLogicalPhysicalConvertor(r)
}

var (
	synLogicalPhysicalConvertorOnce   base.Once
	synLogicalPhysicalConvertorImport *imports.Imports = nil
)

func synLogicalPhysicalConvertorAPI() *imports.Imports {
	synLogicalPhysicalConvertorOnce.Do(func() {
		synLogicalPhysicalConvertorImport = api.NewDefaultImports()
		synLogicalPhysicalConvertorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynLogicalPhysicalConvertor_Create", 0), // constructor NewSynLogicalPhysicalConvertor
			/* 1 */ imports.NewTable("TSynLogicalPhysicalConvertor_PhysicalToLogicalWithIntX2", 0), // function PhysicalToLogicalWithIntX2
			/* 2 */ imports.NewTable("TSynLogicalPhysicalConvertor_PhysicalToLogicalWithIntX3SPCSideSLPFlags", 0), // function PhysicalToLogicalWithIntX3SPCSideSLPFlags
			/* 3 */ imports.NewTable("TSynLogicalPhysicalConvertor_LogicalToPhysicalWithIntX2", 0), // function LogicalToPhysicalWithIntX2
			/* 4 */ imports.NewTable("TSynLogicalPhysicalConvertor_LogicalToPhysicalWithIntX3SLCSideSLPFlags", 0), // function LogicalToPhysicalWithIntX3SLCSideSLPFlags
			/* 5 */ imports.NewTable("TSynLogicalPhysicalConvertor_CurrentLine", 0), // property CurrentLine
			/* 6 */ imports.NewTable("TSynLogicalPhysicalConvertor_CurrentWidths", 0), // property CurrentWidths
			/* 7 */ imports.NewTable("TSynLogicalPhysicalConvertor_CurrentWidthsDirect", 0), // property CurrentWidthsDirect
			/* 8 */ imports.NewTable("TSynLogicalPhysicalConvertor_CurrentWidthsCount", 0), // property CurrentWidthsCount
			/* 9 */ imports.NewTable("TSynLogicalPhysicalConvertor_AdjustedLogToPhysOrigin", 0), // property AdjustedLogToPhysOrigin
			/* 10 */ imports.NewTable("TSynLogicalPhysicalConvertor_AdjustedPhysToLogOrigin", 0), // property AdjustedPhysToLogOrigin
			/* 11 */ imports.NewTable("TSynLogicalPhysicalConvertor_UnAdjustedPhysToLogResult", 0), // property UnAdjustedPhysToLogResult
			/* 12 */ imports.NewTable("TSynLogicalPhysicalConvertor_UnAdjustedPhysToLogColOffs", 0), // property UnAdjustedPhysToLogColOffs
		}
	})
	return synLogicalPhysicalConvertorImport
}
