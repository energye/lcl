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
	r1 := LCL().SysCallN(4618)
	return AsPrinter(r1)
}

func (m *TPrinter) PrinterIndex() int32 {
	r1 := LCL().SysCallN(4630, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetPrinterIndex(AValue int32) {
	LCL().SysCallN(4630, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterName() string {
	r1 := LCL().SysCallN(4631, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) PaperSize() IPaperSize {
	r1 := LCL().SysCallN(4629, m.Instance())
	return AsPaperSize(r1)
}

func (m *TPrinter) Orientation() TPrinterOrientation {
	r1 := LCL().SysCallN(4625, 0, m.Instance(), 0)
	return TPrinterOrientation(r1)
}

func (m *TPrinter) SetOrientation(AValue TPrinterOrientation) {
	LCL().SysCallN(4625, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterState() TPrinterState {
	r1 := LCL().SysCallN(4632, m.Instance())
	return TPrinterState(r1)
}

func (m *TPrinter) Copies() int32 {
	r1 := LCL().SysCallN(4617, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetCopies(AValue int32) {
	LCL().SysCallN(4617, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) Printers() IStrings {
	r1 := LCL().SysCallN(4634, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) FileName() string {
	r1 := LCL().SysCallN(4622, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetFileName(AValue string) {
	LCL().SysCallN(4622, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) Fonts() IStrings {
	r1 := LCL().SysCallN(4623, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Canvas() ICanvas {
	r1 := LCL().SysCallN(4614, m.Instance())
	return AsCanvas(r1)
}

func (m *TPrinter) CanvasClass() uintptr {
	r1 := LCL().SysCallN(4615, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TPrinter) SetCanvasClass(AValue uintptr) {
	LCL().SysCallN(4615, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PageHeight() int32 {
	r1 := LCL().SysCallN(4626, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageWidth() int32 {
	r1 := LCL().SysCallN(4628, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageNumber() int32 {
	r1 := LCL().SysCallN(4627, m.Instance())
	return int32(r1)
}

func (m *TPrinter) Aborted() bool {
	r1 := LCL().SysCallN(4608, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Printing() bool {
	r1 := LCL().SysCallN(4635, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Title() string {
	r1 := LCL().SysCallN(4641, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetTitle(AValue string) {
	LCL().SysCallN(4641, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) PrinterType() TPrinterType {
	r1 := LCL().SysCallN(4633, m.Instance())
	return TPrinterType(r1)
}

func (m *TPrinter) CanPrint() bool {
	r1 := LCL().SysCallN(4612, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) CanRenderCopies() bool {
	r1 := LCL().SysCallN(4613, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) XDPI() int32 {
	r1 := LCL().SysCallN(4644, m.Instance())
	return int32(r1)
}

func (m *TPrinter) YDPI() int32 {
	r1 := LCL().SysCallN(4645, m.Instance())
	return int32(r1)
}

func (m *TPrinter) RawMode() bool {
	r1 := LCL().SysCallN(4636, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPrinter) SetRawMode(AValue bool) {
	LCL().SysCallN(4636, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPrinter) DefaultBinName() string {
	r1 := LCL().SysCallN(4619, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) BinName() string {
	r1 := LCL().SysCallN(4611, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetBinName(AValue string) {
	LCL().SysCallN(4611, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) SupportedBins() IStrings {
	r1 := LCL().SysCallN(4640, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Write(Buffer []byte, OutWritten *int32) bool {
	var resultWritten uintptr
	r1 := LCL().SysCallN(4642, m.Instance(), uintptr(unsafePointer(&Buffer[0])), uintptr(len(Buffer)), uintptr(unsafePointer(&resultWritten)))
	*OutWritten = int32(resultWritten)
	return GoBool(r1)
}

func (m *TPrinter) Write1(s string) bool {
	r1 := LCL().SysCallN(4643, m.Instance(), PascalStr(s))
	return GoBool(r1)
}

func PrinterClass() TClass {
	ret := LCL().SysCallN(4616)
	return TClass(ret)
}

func (m *TPrinter) Abort() {
	LCL().SysCallN(4607, m.Instance())
}

func (m *TPrinter) BeginDoc() {
	LCL().SysCallN(4609, m.Instance())
}

func (m *TPrinter) EndDoc() {
	LCL().SysCallN(4620, m.Instance())
}

func (m *TPrinter) NewPage() {
	LCL().SysCallN(4624, m.Instance())
}

func (m *TPrinter) BeginPage() {
	LCL().SysCallN(4610, m.Instance())
}

func (m *TPrinter) EndPage() {
	LCL().SysCallN(4621, m.Instance())
}

func (m *TPrinter) Refresh() {
	LCL().SysCallN(4637, m.Instance())
}

func (m *TPrinter) SetPrinter(aName string) {
	LCL().SysCallN(4639, m.Instance(), PascalStr(aName))
}

func (m *TPrinter) RestoreDefaultBin() {
	LCL().SysCallN(4638, m.Instance())
}
