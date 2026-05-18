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

// ISynGutterLineNumber Parent: ISynGutterPartBase
type ISynGutterLineNumber interface {
	ISynGutterPartBase
	FormatLineNumber(line int32, isDot bool) string                                             // function
	PaintWithCanvasRectIntX2(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32) // procedure
	DigitCount() int32                                                                          // property DigitCount Getter
	SetDigitCount(value int32)                                                                  // property DigitCount Setter
	ShowOnlyLineNumbersMultiplesOf() int32                                                      // property ShowOnlyLineNumbersMultiplesOf Getter
	SetShowOnlyLineNumbersMultiplesOf(value int32)                                              // property ShowOnlyLineNumbersMultiplesOf Setter
	ZeroStart() bool                                                                            // property ZeroStart Getter
	SetZeroStart(value bool)                                                                    // property ZeroStart Setter
	LeadingZeros() bool                                                                         // property LeadingZeros Getter
	SetLeadingZeros(value bool)                                                                 // property LeadingZeros Setter
	SetOnFormatLineNumber(fn TSynEditGetGutterLineTextEvent)                                    // property event
}

type TSynGutterLineNumber struct {
	TSynGutterPartBase
}

func (m *TSynGutterLineNumber) FormatLineNumber(line int32, isDot bool) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	synGutterLineNumberAPI().SysCallN(1, m.Instance(), uintptr(line), api.PasBool(isDot), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynGutterLineNumber) PaintWithCanvasRectIntX2(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32) {
	if !m.IsValid() {
		return
	}
	synGutterLineNumberAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&clip)), uintptr(firstLine), uintptr(lastLine))
}

func (m *TSynGutterLineNumber) DigitCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLineNumberAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterLineNumber) SetDigitCount(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterLineNumberAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterLineNumber) ShowOnlyLineNumbersMultiplesOf() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLineNumberAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterLineNumber) SetShowOnlyLineNumbersMultiplesOf(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterLineNumberAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterLineNumber) ZeroStart() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterLineNumberAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterLineNumber) SetZeroStart(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterLineNumberAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynGutterLineNumber) LeadingZeros() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterLineNumberAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterLineNumber) SetLeadingZeros(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterLineNumberAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynGutterLineNumber) SetOnFormatLineNumber(fn TSynEditGetGutterLineTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynEditGetGutterLineTextEvent(fn)
	base.SetEvent(m, 7, synGutterLineNumberAPI(), api.MakeEventDataPtr(cb))
}

// NewSynGutterLineNumber class constructor
func NewSynGutterLineNumber(owner IComponent) ISynGutterLineNumber {
	r := synGutterLineNumberAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterLineNumber(r)
}

func TSynGutterLineNumberClass() types.TClass {
	r := synGutterLineNumberAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	synGutterLineNumberOnce   base.Once
	synGutterLineNumberImport *imports.Imports = nil
)

func synGutterLineNumberAPI() *imports.Imports {
	synGutterLineNumberOnce.Do(func() {
		synGutterLineNumberImport = api.NewDefaultImports()
		synGutterLineNumberImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterLineNumber_Create", 0), // constructor NewSynGutterLineNumber
			/* 1 */ imports.NewTable("TSynGutterLineNumber_FormatLineNumber", 0), // function FormatLineNumber
			/* 2 */ imports.NewTable("TSynGutterLineNumber_PaintWithCanvasRectIntX2", 0), // procedure PaintWithCanvasRectIntX2
			/* 3 */ imports.NewTable("TSynGutterLineNumber_DigitCount", 0), // property DigitCount
			/* 4 */ imports.NewTable("TSynGutterLineNumber_ShowOnlyLineNumbersMultiplesOf", 0), // property ShowOnlyLineNumbersMultiplesOf
			/* 5 */ imports.NewTable("TSynGutterLineNumber_ZeroStart", 0), // property ZeroStart
			/* 6 */ imports.NewTable("TSynGutterLineNumber_LeadingZeros", 0), // property LeadingZeros
			/* 7 */ imports.NewTable("TSynGutterLineNumber_OnFormatLineNumber", 0), // event OnFormatLineNumber
			/* 8 */ imports.NewTable("TSynGutterLineNumber_TClass", 0), // function TSynGutterLineNumberClass
		}
	})
	return synGutterLineNumberImport
}
