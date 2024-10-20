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

// IPrinterCanvas Parent: ICanvas
type IPrinterCanvas interface {
	ICanvas
	Printer() IPrinter                         // property
	Title() string                             // property
	SetTitle(AValue string)                    // property
	PageHeight() int32                         // property
	PageWidth() int32                          // property
	PaperWidth() int32                         // property
	SetPaperWidth(AValue int32)                // property
	PaperHeight() int32                        // property
	SetPaperHeight(AValue int32)               // property
	PageNumber() int32                         // property
	TopMargin() int32                          // property
	SetTopMargin(AValue int32)                 // property
	LeftMargin() int32                         // property
	SetLeftMargin(AValue int32)                // property
	BottomMargin() int32                       // property
	SetBottomMargin(AValue int32)              // property
	RightMargin() int32                        // property
	SetRightMargin(AValue int32)               // property
	Orientation() TPrinterOrientation          // property
	SetOrientation(AValue TPrinterOrientation) // property
	XDPI() int32                               // property
	SetXDPI(AValue int32)                      // property
	YDPI() int32                               // property
	SetYDPI(AValue int32)                      // property
	BeginDoc()                                 // procedure
	NewPage()                                  // procedure
	BeginPage()                                // procedure
	EndPage()                                  // procedure
	EndDoc()                                   // procedure
}

// TPrinterCanvas Parent: TCanvas
type TPrinterCanvas struct {
	TCanvas
}

func NewPrinterCanvas(APrinter IPrinter) IPrinterCanvas {
	r1 := printerCanvasImportAPI().SysCallN(4, GetObjectUintptr(APrinter))
	return AsPrinterCanvas(r1)
}

func (m *TPrinterCanvas) Printer() IPrinter {
	r1 := printerCanvasImportAPI().SysCallN(15, m.Instance())
	return AsPrinter(r1)
}

func (m *TPrinterCanvas) Title() string {
	r1 := printerCanvasImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinterCanvas) SetTitle(AValue string) {
	printerCanvasImportAPI().SysCallN(17, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinterCanvas) PageHeight() int32 {
	r1 := printerCanvasImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TPrinterCanvas) PageWidth() int32 {
	r1 := printerCanvasImportAPI().SysCallN(12, m.Instance())
	return int32(r1)
}

func (m *TPrinterCanvas) PaperWidth() int32 {
	r1 := printerCanvasImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetPaperWidth(AValue int32) {
	printerCanvasImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) PaperHeight() int32 {
	r1 := printerCanvasImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetPaperHeight(AValue int32) {
	printerCanvasImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) PageNumber() int32 {
	r1 := printerCanvasImportAPI().SysCallN(11, m.Instance())
	return int32(r1)
}

func (m *TPrinterCanvas) TopMargin() int32 {
	r1 := printerCanvasImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetTopMargin(AValue int32) {
	printerCanvasImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) LeftMargin() int32 {
	r1 := printerCanvasImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetLeftMargin(AValue int32) {
	printerCanvasImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) BottomMargin() int32 {
	r1 := printerCanvasImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetBottomMargin(AValue int32) {
	printerCanvasImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) RightMargin() int32 {
	r1 := printerCanvasImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetRightMargin(AValue int32) {
	printerCanvasImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) Orientation() TPrinterOrientation {
	r1 := printerCanvasImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TPrinterOrientation(r1)
}

func (m *TPrinterCanvas) SetOrientation(AValue TPrinterOrientation) {
	printerCanvasImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) XDPI() int32 {
	r1 := printerCanvasImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetXDPI(AValue int32) {
	printerCanvasImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinterCanvas) YDPI() int32 {
	r1 := printerCanvasImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinterCanvas) SetYDPI(AValue int32) {
	printerCanvasImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func PrinterCanvasClass() TClass {
	ret := printerCanvasImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TPrinterCanvas) BeginDoc() {
	printerCanvasImportAPI().SysCallN(0, m.Instance())
}

func (m *TPrinterCanvas) NewPage() {
	printerCanvasImportAPI().SysCallN(8, m.Instance())
}

func (m *TPrinterCanvas) BeginPage() {
	printerCanvasImportAPI().SysCallN(1, m.Instance())
}

func (m *TPrinterCanvas) EndPage() {
	printerCanvasImportAPI().SysCallN(6, m.Instance())
}

func (m *TPrinterCanvas) EndDoc() {
	printerCanvasImportAPI().SysCallN(5, m.Instance())
}

var (
	printerCanvasImport       *imports.Imports = nil
	printerCanvasImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PrinterCanvas_BeginDoc", 0),
		/*1*/ imports.NewTable("PrinterCanvas_BeginPage", 0),
		/*2*/ imports.NewTable("PrinterCanvas_BottomMargin", 0),
		/*3*/ imports.NewTable("PrinterCanvas_Class", 0),
		/*4*/ imports.NewTable("PrinterCanvas_Create", 0),
		/*5*/ imports.NewTable("PrinterCanvas_EndDoc", 0),
		/*6*/ imports.NewTable("PrinterCanvas_EndPage", 0),
		/*7*/ imports.NewTable("PrinterCanvas_LeftMargin", 0),
		/*8*/ imports.NewTable("PrinterCanvas_NewPage", 0),
		/*9*/ imports.NewTable("PrinterCanvas_Orientation", 0),
		/*10*/ imports.NewTable("PrinterCanvas_PageHeight", 0),
		/*11*/ imports.NewTable("PrinterCanvas_PageNumber", 0),
		/*12*/ imports.NewTable("PrinterCanvas_PageWidth", 0),
		/*13*/ imports.NewTable("PrinterCanvas_PaperHeight", 0),
		/*14*/ imports.NewTable("PrinterCanvas_PaperWidth", 0),
		/*15*/ imports.NewTable("PrinterCanvas_Printer", 0),
		/*16*/ imports.NewTable("PrinterCanvas_RightMargin", 0),
		/*17*/ imports.NewTable("PrinterCanvas_Title", 0),
		/*18*/ imports.NewTable("PrinterCanvas_TopMargin", 0),
		/*19*/ imports.NewTable("PrinterCanvas_XDPI", 0),
		/*20*/ imports.NewTable("PrinterCanvas_YDPI", 0),
	}
)

func printerCanvasImportAPI() *imports.Imports {
	if printerCanvasImport == nil {
		printerCanvasImport = NewDefaultImports()
		printerCanvasImport.SetImportTable(printerCanvasImportTables)
		printerCanvasImportTables = nil
	}
	return printerCanvasImport
}
