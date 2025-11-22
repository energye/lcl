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

// IPrinter Parent: IObject
type IPrinter interface {
	IObject
	WriteWithPointerIntX2(buffer uintptr, count int32, outWritten *int32) bool // function
	WriteWithAnsistring(S string) bool                                         // function
	Abort()                                                                    // procedure
	BeginDoc()                                                                 // procedure
	EndDoc()                                                                   // procedure
	NewPage()                                                                  // procedure
	BeginPage()                                                                // procedure
	EndPage()                                                                  // procedure
	Refresh()                                                                  // procedure
	SetPrinter(name string)                                                    // procedure
	RestoreDefaultBin()                                                        // procedure
	PrinterIndex() int32                                                       // property PrinterIndex Getter
	SetPrinterIndex(value int32)                                               // property PrinterIndex Setter
	PrinterName() string                                                       // property PrinterName Getter
	PaperSize() IPaperSize                                                     // property PaperSize Getter
	Orientation() types.TPrinterOrientation                                    // property Orientation Getter
	SetOrientation(value types.TPrinterOrientation)                            // property Orientation Setter
	PrinterState() types.TPrinterState                                         // property PrinterState Getter
	Copies() int32                                                             // property Copies Getter
	SetCopies(value int32)                                                     // property Copies Setter
	Printers() IStrings                                                        // property Printers Getter
	FileName() string                                                          // property FileName Getter
	SetFileName(value string)                                                  // property FileName Setter
	Fonts() IStrings                                                           // property Fonts Getter
	Canvas() ICanvas                                                           // property Canvas Getter
	CanvasClass() IPrinterCanvas                                               // property CanvasClass Getter
	SetCanvasClass(value IPrinterCanvas)                                       // property CanvasClass Setter
	PageHeight() int32                                                         // property PageHeight Getter
	PageWidth() int32                                                          // property PageWidth Getter
	PageNumber() int32                                                         // property PageNumber Getter
	Aborted() bool                                                             // property Aborted Getter
	Printing() bool                                                            // property Printing Getter
	Title() string                                                             // property Title Getter
	SetTitle(value string)                                                     // property Title Setter
	PrinterType() types.TPrinterType                                           // property PrinterType Getter
	CanPrint() bool                                                            // property CanPrint Getter
	CanRenderCopies() bool                                                     // property CanRenderCopies Getter
	XDPI() int32                                                               // property XDPI Getter
	YDPI() int32                                                               // property YDPI Getter
	RawMode() bool                                                             // property RawMode Getter
	SetRawMode(value bool)                                                     // property RawMode Setter
	DefaultBinName() string                                                    // property DefaultBinName Getter
	BinName() string                                                           // property BinName Getter
	SetBinName(value string)                                                   // property BinName Setter
	SupportedBins() IStrings                                                   // property SupportedBins Getter
}

type TPrinter struct {
	TObject
}

func (m *TPrinter) WriteWithPointerIntX2(buffer uintptr, count int32, outWritten *int32) bool {
	if !m.IsValid() {
		return false
	}
	var writtenPtr uintptr
	r := printerAPI().SysCallN(1, m.Instance(), uintptr(buffer), uintptr(count), uintptr(base.UnsafePointer(&writtenPtr)))
	*outWritten = int32(writtenPtr)
	return api.GoBool(r)
}

func (m *TPrinter) WriteWithAnsistring(S string) bool {
	if !m.IsValid() {
		return false
	}
	r := printerAPI().SysCallN(2, m.Instance(), api.PasStr(S))
	return api.GoBool(r)
}

func (m *TPrinter) Abort() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(3, m.Instance())
}

func (m *TPrinter) BeginDoc() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(4, m.Instance())
}

func (m *TPrinter) EndDoc() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(5, m.Instance())
}

func (m *TPrinter) NewPage() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(6, m.Instance())
}

func (m *TPrinter) BeginPage() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(7, m.Instance())
}

func (m *TPrinter) EndPage() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(8, m.Instance())
}

func (m *TPrinter) Refresh() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(9, m.Instance())
}

func (m *TPrinter) SetPrinter(name string) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(10, m.Instance(), api.PasStr(name))
}

func (m *TPrinter) RestoreDefaultBin() {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(11, m.Instance())
}

func (m *TPrinter) PrinterIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TPrinter) SetPrinterIndex(value int32) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TPrinter) PrinterName() string {
	if !m.IsValid() {
		return ""
	}
	r := printerAPI().SysCallN(13, m.Instance())
	return api.GoStr(r)
}

func (m *TPrinter) PaperSize() IPaperSize {
	if !m.IsValid() {
		return nil
	}
	r := printerAPI().SysCallN(14, m.Instance())
	return AsPaperSize(r)
}

func (m *TPrinter) Orientation() types.TPrinterOrientation {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(15, 0, m.Instance())
	return types.TPrinterOrientation(r)
}

func (m *TPrinter) SetOrientation(value types.TPrinterOrientation) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TPrinter) PrinterState() types.TPrinterState {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(16, m.Instance())
	return types.TPrinterState(r)
}

func (m *TPrinter) Copies() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TPrinter) SetCopies(value int32) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TPrinter) Printers() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := printerAPI().SysCallN(18, m.Instance())
	return AsStrings(r)
}

func (m *TPrinter) FileName() string {
	if !m.IsValid() {
		return ""
	}
	r := printerAPI().SysCallN(19, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TPrinter) SetFileName(value string) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(19, 1, m.Instance(), api.PasStr(value))
}

func (m *TPrinter) Fonts() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := printerAPI().SysCallN(20, m.Instance())
	return AsStrings(r)
}

func (m *TPrinter) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := printerAPI().SysCallN(21, m.Instance())
	return AsCanvas(r)
}

func (m *TPrinter) CanvasClass() IPrinterCanvas {
	if !m.IsValid() {
		return nil
	}
	r := printerAPI().SysCallN(22, 0, m.Instance())
	return AsPrinterCanvas(r)
}

func (m *TPrinter) SetCanvasClass(value IPrinterCanvas) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPrinter) PageHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(23, m.Instance())
	return int32(r)
}

func (m *TPrinter) PageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(24, m.Instance())
	return int32(r)
}

func (m *TPrinter) PageNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(25, m.Instance())
	return int32(r)
}

func (m *TPrinter) Aborted() bool {
	if !m.IsValid() {
		return false
	}
	r := printerAPI().SysCallN(26, m.Instance())
	return api.GoBool(r)
}

func (m *TPrinter) Printing() bool {
	if !m.IsValid() {
		return false
	}
	r := printerAPI().SysCallN(27, m.Instance())
	return api.GoBool(r)
}

func (m *TPrinter) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := printerAPI().SysCallN(28, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TPrinter) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(28, 1, m.Instance(), api.PasStr(value))
}

func (m *TPrinter) PrinterType() types.TPrinterType {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(29, m.Instance())
	return types.TPrinterType(r)
}

func (m *TPrinter) CanPrint() bool {
	if !m.IsValid() {
		return false
	}
	r := printerAPI().SysCallN(30, m.Instance())
	return api.GoBool(r)
}

func (m *TPrinter) CanRenderCopies() bool {
	if !m.IsValid() {
		return false
	}
	r := printerAPI().SysCallN(31, m.Instance())
	return api.GoBool(r)
}

func (m *TPrinter) XDPI() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(32, m.Instance())
	return int32(r)
}

func (m *TPrinter) YDPI() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerAPI().SysCallN(33, m.Instance())
	return int32(r)
}

func (m *TPrinter) RawMode() bool {
	if !m.IsValid() {
		return false
	}
	r := printerAPI().SysCallN(34, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPrinter) SetRawMode(value bool) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(34, 1, m.Instance(), api.PasBool(value))
}

func (m *TPrinter) DefaultBinName() string {
	if !m.IsValid() {
		return ""
	}
	r := printerAPI().SysCallN(35, m.Instance())
	return api.GoStr(r)
}

func (m *TPrinter) BinName() string {
	if !m.IsValid() {
		return ""
	}
	r := printerAPI().SysCallN(36, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TPrinter) SetBinName(value string) {
	if !m.IsValid() {
		return
	}
	printerAPI().SysCallN(36, 1, m.Instance(), api.PasStr(value))
}

func (m *TPrinter) SupportedBins() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := printerAPI().SysCallN(37, m.Instance())
	return AsStrings(r)
}

// NewPrinter class constructor
func NewPrinter() IPrinter {
	r := printerAPI().SysCallN(0)
	return AsPrinter(r)
}

var (
	printerOnce   base.Once
	printerImport *imports.Imports = nil
)

func printerAPI() *imports.Imports {
	printerOnce.Do(func() {
		printerImport = api.NewDefaultImports()
		printerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPrinter_Create", 0), // constructor NewPrinter
			/* 1 */ imports.NewTable("TPrinter_WriteWithPointerIntX2", 0), // function WriteWithPointerIntX2
			/* 2 */ imports.NewTable("TPrinter_WriteWithAnsistring", 0), // function WriteWithAnsistring
			/* 3 */ imports.NewTable("TPrinter_Abort", 0), // procedure Abort
			/* 4 */ imports.NewTable("TPrinter_BeginDoc", 0), // procedure BeginDoc
			/* 5 */ imports.NewTable("TPrinter_EndDoc", 0), // procedure EndDoc
			/* 6 */ imports.NewTable("TPrinter_NewPage", 0), // procedure NewPage
			/* 7 */ imports.NewTable("TPrinter_BeginPage", 0), // procedure BeginPage
			/* 8 */ imports.NewTable("TPrinter_EndPage", 0), // procedure EndPage
			/* 9 */ imports.NewTable("TPrinter_Refresh", 0), // procedure Refresh
			/* 10 */ imports.NewTable("TPrinter_SetPrinter", 0), // procedure SetPrinter
			/* 11 */ imports.NewTable("TPrinter_RestoreDefaultBin", 0), // procedure RestoreDefaultBin
			/* 12 */ imports.NewTable("TPrinter_PrinterIndex", 0), // property PrinterIndex
			/* 13 */ imports.NewTable("TPrinter_PrinterName", 0), // property PrinterName
			/* 14 */ imports.NewTable("TPrinter_PaperSize", 0), // property PaperSize
			/* 15 */ imports.NewTable("TPrinter_Orientation", 0), // property Orientation
			/* 16 */ imports.NewTable("TPrinter_PrinterState", 0), // property PrinterState
			/* 17 */ imports.NewTable("TPrinter_Copies", 0), // property Copies
			/* 18 */ imports.NewTable("TPrinter_Printers", 0), // property Printers
			/* 19 */ imports.NewTable("TPrinter_FileName", 0), // property FileName
			/* 20 */ imports.NewTable("TPrinter_Fonts", 0), // property Fonts
			/* 21 */ imports.NewTable("TPrinter_Canvas", 0), // property Canvas
			/* 22 */ imports.NewTable("TPrinter_CanvasClass", 0), // property CanvasClass
			/* 23 */ imports.NewTable("TPrinter_PageHeight", 0), // property PageHeight
			/* 24 */ imports.NewTable("TPrinter_PageWidth", 0), // property PageWidth
			/* 25 */ imports.NewTable("TPrinter_PageNumber", 0), // property PageNumber
			/* 26 */ imports.NewTable("TPrinter_Aborted", 0), // property Aborted
			/* 27 */ imports.NewTable("TPrinter_Printing", 0), // property Printing
			/* 28 */ imports.NewTable("TPrinter_Title", 0), // property Title
			/* 29 */ imports.NewTable("TPrinter_PrinterType", 0), // property PrinterType
			/* 30 */ imports.NewTable("TPrinter_CanPrint", 0), // property CanPrint
			/* 31 */ imports.NewTable("TPrinter_CanRenderCopies", 0), // property CanRenderCopies
			/* 32 */ imports.NewTable("TPrinter_XDPI", 0), // property XDPI
			/* 33 */ imports.NewTable("TPrinter_YDPI", 0), // property YDPI
			/* 34 */ imports.NewTable("TPrinter_RawMode", 0), // property RawMode
			/* 35 */ imports.NewTable("TPrinter_DefaultBinName", 0), // property DefaultBinName
			/* 36 */ imports.NewTable("TPrinter_BinName", 0), // property BinName
			/* 37 */ imports.NewTable("TPrinter_SupportedBins", 0), // property SupportedBins
		}
	})
	return printerImport
}
