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

// IPrinter Parent: IObject
type IPrinter interface {
	IObject
	PrinterIndex() int32                         // property
	SetPrinterIndex(AValue int32)                // property
	PrinterName() string                         // property
	PaperSize() IPaperSize                       // property
	Orientation() TPrinterOrientation            // property
	SetOrientation(AValue TPrinterOrientation)   // property
	PrinterState() TPrinterState                 // property
	Copies() int32                               // property
	SetCopies(AValue int32)                      // property
	Printers() IStrings                          // property
	FileName() string                            // property
	SetFileName(AValue string)                   // property
	Fonts() IStrings                             // property
	Canvas() ICanvas                             // property
	CanvasClass() uintptr                        // property
	SetCanvasClass(AValue uintptr)               // property
	PageHeight() int32                           // property
	PageWidth() int32                            // property
	PageNumber() int32                           // property
	Aborted() bool                               // property
	Printing() bool                              // property
	Title() string                               // property
	SetTitle(AValue string)                      // property
	PrinterType() TPrinterType                   // property
	CanPrint() bool                              // property
	CanRenderCopies() bool                       // property
	XDPI() int32                                 // property
	YDPI() int32                                 // property
	RawMode() bool                               // property
	SetRawMode(AValue bool)                      // property
	DefaultBinName() string                      // property
	BinName() string                             // property
	SetBinName(AValue string)                    // property
	SupportedBins() IStrings                     // property
	Write(Buffer []byte, OutWritten *int32) bool // function
	Write1(s string) bool                        // function
	Abort()                                      // procedure
	BeginDoc()                                   // procedure
	EndDoc()                                     // procedure
	NewPage()                                    // procedure
	BeginPage()                                  // procedure
	EndPage()                                    // procedure
	Refresh()                                    // procedure
	SetPrinter(aName string)                     // procedure
	RestoreDefaultBin()                          // procedure
}

// TPrinter Parent: TObject
type TPrinter struct {
	TObject
}

func NewPrinter() IPrinter {
	r1 := printerImportAPI().SysCallN(11)
	return AsPrinter(r1)
}

func (m *TPrinter) PrinterIndex() int32 {
	r1 := printerImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetPrinterIndex(AValue int32) {
	printerImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterName() string {
	r1 := printerImportAPI().SysCallN(24, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) PaperSize() IPaperSize {
	r1 := printerImportAPI().SysCallN(22, m.Instance())
	return AsPaperSize(r1)
}

func (m *TPrinter) Orientation() TPrinterOrientation {
	r1 := printerImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TPrinterOrientation(r1)
}

func (m *TPrinter) SetOrientation(AValue TPrinterOrientation) {
	printerImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterState() TPrinterState {
	r1 := printerImportAPI().SysCallN(25, m.Instance())
	return TPrinterState(r1)
}

func (m *TPrinter) Copies() int32 {
	r1 := printerImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetCopies(AValue int32) {
	printerImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) Printers() IStrings {
	r1 := printerImportAPI().SysCallN(27, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) FileName() string {
	r1 := printerImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetFileName(AValue string) {
	printerImportAPI().SysCallN(15, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) Fonts() IStrings {
	r1 := printerImportAPI().SysCallN(16, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Canvas() ICanvas {
	r1 := printerImportAPI().SysCallN(7, m.Instance())
	return AsCanvas(r1)
}

func (m *TPrinter) CanvasClass() uintptr {
	r1 := printerImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TPrinter) SetCanvasClass(AValue uintptr) {
	printerImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PageHeight() int32 {
	r1 := printerImportAPI().SysCallN(19, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageWidth() int32 {
	r1 := printerImportAPI().SysCallN(21, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageNumber() int32 {
	r1 := printerImportAPI().SysCallN(20, m.Instance())
	return int32(r1)
}

func (m *TPrinter) Aborted() bool {
	r1 := printerImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Printing() bool {
	r1 := printerImportAPI().SysCallN(28, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Title() string {
	r1 := printerImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetTitle(AValue string) {
	printerImportAPI().SysCallN(34, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) PrinterType() TPrinterType {
	r1 := printerImportAPI().SysCallN(26, m.Instance())
	return TPrinterType(r1)
}

func (m *TPrinter) CanPrint() bool {
	r1 := printerImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) CanRenderCopies() bool {
	r1 := printerImportAPI().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) XDPI() int32 {
	r1 := printerImportAPI().SysCallN(37, m.Instance())
	return int32(r1)
}

func (m *TPrinter) YDPI() int32 {
	r1 := printerImportAPI().SysCallN(38, m.Instance())
	return int32(r1)
}

func (m *TPrinter) RawMode() bool {
	r1 := printerImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPrinter) SetRawMode(AValue bool) {
	printerImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPrinter) DefaultBinName() string {
	r1 := printerImportAPI().SysCallN(12, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) BinName() string {
	r1 := printerImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetBinName(AValue string) {
	printerImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) SupportedBins() IStrings {
	r1 := printerImportAPI().SysCallN(33, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Write(Buffer []byte, OutWritten *int32) bool {
	var resultWritten uintptr
	r1 := printerImportAPI().SysCallN(35, m.Instance(), uintptr(unsafePointer(&Buffer[0])), uintptr(len(Buffer)), uintptr(unsafePointer(&resultWritten)))
	*OutWritten = int32(resultWritten)
	return GoBool(r1)
}

func (m *TPrinter) Write1(s string) bool {
	r1 := printerImportAPI().SysCallN(36, m.Instance(), PascalStr(s))
	return GoBool(r1)
}

func PrinterClass() TClass {
	ret := printerImportAPI().SysCallN(9)
	return TClass(ret)
}

func (m *TPrinter) Abort() {
	printerImportAPI().SysCallN(0, m.Instance())
}

func (m *TPrinter) BeginDoc() {
	printerImportAPI().SysCallN(2, m.Instance())
}

func (m *TPrinter) EndDoc() {
	printerImportAPI().SysCallN(13, m.Instance())
}

func (m *TPrinter) NewPage() {
	printerImportAPI().SysCallN(17, m.Instance())
}

func (m *TPrinter) BeginPage() {
	printerImportAPI().SysCallN(3, m.Instance())
}

func (m *TPrinter) EndPage() {
	printerImportAPI().SysCallN(14, m.Instance())
}

func (m *TPrinter) Refresh() {
	printerImportAPI().SysCallN(30, m.Instance())
}

func (m *TPrinter) SetPrinter(aName string) {
	printerImportAPI().SysCallN(32, m.Instance(), PascalStr(aName))
}

func (m *TPrinter) RestoreDefaultBin() {
	printerImportAPI().SysCallN(31, m.Instance())
}

var (
	printerImport       *imports.Imports = nil
	printerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Printer_Abort", 0),
		/*1*/ imports.NewTable("Printer_Aborted", 0),
		/*2*/ imports.NewTable("Printer_BeginDoc", 0),
		/*3*/ imports.NewTable("Printer_BeginPage", 0),
		/*4*/ imports.NewTable("Printer_BinName", 0),
		/*5*/ imports.NewTable("Printer_CanPrint", 0),
		/*6*/ imports.NewTable("Printer_CanRenderCopies", 0),
		/*7*/ imports.NewTable("Printer_Canvas", 0),
		/*8*/ imports.NewTable("Printer_CanvasClass", 0),
		/*9*/ imports.NewTable("Printer_Class", 0),
		/*10*/ imports.NewTable("Printer_Copies", 0),
		/*11*/ imports.NewTable("Printer_Create", 0),
		/*12*/ imports.NewTable("Printer_DefaultBinName", 0),
		/*13*/ imports.NewTable("Printer_EndDoc", 0),
		/*14*/ imports.NewTable("Printer_EndPage", 0),
		/*15*/ imports.NewTable("Printer_FileName", 0),
		/*16*/ imports.NewTable("Printer_Fonts", 0),
		/*17*/ imports.NewTable("Printer_NewPage", 0),
		/*18*/ imports.NewTable("Printer_Orientation", 0),
		/*19*/ imports.NewTable("Printer_PageHeight", 0),
		/*20*/ imports.NewTable("Printer_PageNumber", 0),
		/*21*/ imports.NewTable("Printer_PageWidth", 0),
		/*22*/ imports.NewTable("Printer_PaperSize", 0),
		/*23*/ imports.NewTable("Printer_PrinterIndex", 0),
		/*24*/ imports.NewTable("Printer_PrinterName", 0),
		/*25*/ imports.NewTable("Printer_PrinterState", 0),
		/*26*/ imports.NewTable("Printer_PrinterType", 0),
		/*27*/ imports.NewTable("Printer_Printers", 0),
		/*28*/ imports.NewTable("Printer_Printing", 0),
		/*29*/ imports.NewTable("Printer_RawMode", 0),
		/*30*/ imports.NewTable("Printer_Refresh", 0),
		/*31*/ imports.NewTable("Printer_RestoreDefaultBin", 0),
		/*32*/ imports.NewTable("Printer_SetPrinter", 0),
		/*33*/ imports.NewTable("Printer_SupportedBins", 0),
		/*34*/ imports.NewTable("Printer_Title", 0),
		/*35*/ imports.NewTable("Printer_Write", 0),
		/*36*/ imports.NewTable("Printer_Write1", 0),
		/*37*/ imports.NewTable("Printer_XDPI", 0),
		/*38*/ imports.NewTable("Printer_YDPI", 0),
	}
)

func printerImportAPI() *imports.Imports {
	if printerImport == nil {
		printerImport = NewDefaultImports()
		printerImport.SetImportTable(printerImportTables)
		printerImportTables = nil
	}
	return printerImport
}
