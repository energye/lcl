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
	r1 := LCL().SysCallN(4661)
	return AsPrinter(r1)
}

func (m *TPrinter) PrinterIndex() int32 {
	r1 := LCL().SysCallN(4673, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetPrinterIndex(AValue int32) {
	LCL().SysCallN(4673, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterName() string {
	r1 := LCL().SysCallN(4674, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) PaperSize() IPaperSize {
	r1 := LCL().SysCallN(4672, m.Instance())
	return AsPaperSize(r1)
}

func (m *TPrinter) Orientation() TPrinterOrientation {
	r1 := LCL().SysCallN(4668, 0, m.Instance(), 0)
	return TPrinterOrientation(r1)
}

func (m *TPrinter) SetOrientation(AValue TPrinterOrientation) {
	LCL().SysCallN(4668, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PrinterState() TPrinterState {
	r1 := LCL().SysCallN(4675, m.Instance())
	return TPrinterState(r1)
}

func (m *TPrinter) Copies() int32 {
	r1 := LCL().SysCallN(4660, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPrinter) SetCopies(AValue int32) {
	LCL().SysCallN(4660, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) Printers() IStrings {
	r1 := LCL().SysCallN(4677, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) FileName() string {
	r1 := LCL().SysCallN(4665, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetFileName(AValue string) {
	LCL().SysCallN(4665, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) Fonts() IStrings {
	r1 := LCL().SysCallN(4666, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Canvas() ICanvas {
	r1 := LCL().SysCallN(4657, m.Instance())
	return AsCanvas(r1)
}

func (m *TPrinter) CanvasClass() uintptr {
	r1 := LCL().SysCallN(4658, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TPrinter) SetCanvasClass(AValue uintptr) {
	LCL().SysCallN(4658, 1, m.Instance(), uintptr(AValue))
}

func (m *TPrinter) PageHeight() int32 {
	r1 := LCL().SysCallN(4669, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageWidth() int32 {
	r1 := LCL().SysCallN(4671, m.Instance())
	return int32(r1)
}

func (m *TPrinter) PageNumber() int32 {
	r1 := LCL().SysCallN(4670, m.Instance())
	return int32(r1)
}

func (m *TPrinter) Aborted() bool {
	r1 := LCL().SysCallN(4651, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Printing() bool {
	r1 := LCL().SysCallN(4678, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) Title() string {
	r1 := LCL().SysCallN(4684, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetTitle(AValue string) {
	LCL().SysCallN(4684, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) PrinterType() TPrinterType {
	r1 := LCL().SysCallN(4676, m.Instance())
	return TPrinterType(r1)
}

func (m *TPrinter) CanPrint() bool {
	r1 := LCL().SysCallN(4655, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) CanRenderCopies() bool {
	r1 := LCL().SysCallN(4656, m.Instance())
	return GoBool(r1)
}

func (m *TPrinter) XDPI() int32 {
	r1 := LCL().SysCallN(4687, m.Instance())
	return int32(r1)
}

func (m *TPrinter) YDPI() int32 {
	r1 := LCL().SysCallN(4688, m.Instance())
	return int32(r1)
}

func (m *TPrinter) RawMode() bool {
	r1 := LCL().SysCallN(4679, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPrinter) SetRawMode(AValue bool) {
	LCL().SysCallN(4679, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPrinter) DefaultBinName() string {
	r1 := LCL().SysCallN(4662, m.Instance())
	return GoStr(r1)
}

func (m *TPrinter) BinName() string {
	r1 := LCL().SysCallN(4654, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPrinter) SetBinName(AValue string) {
	LCL().SysCallN(4654, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPrinter) SupportedBins() IStrings {
	r1 := LCL().SysCallN(4683, m.Instance())
	return AsStrings(r1)
}

func (m *TPrinter) Write(Buffer []byte, OutWritten *int32) bool {
	var resultWritten uintptr
	r1 := LCL().SysCallN(4685, m.Instance(), uintptr(unsafePointer(&Buffer[0])), uintptr(len(Buffer)), uintptr(unsafePointer(&resultWritten)))
	*OutWritten = int32(resultWritten)
	return GoBool(r1)
}

func (m *TPrinter) Write1(s string) bool {
	r1 := LCL().SysCallN(4686, m.Instance(), PascalStr(s))
	return GoBool(r1)
}

func PrinterClass() TClass {
	ret := LCL().SysCallN(4659)
	return TClass(ret)
}

func (m *TPrinter) Abort() {
	LCL().SysCallN(4650, m.Instance())
}

func (m *TPrinter) BeginDoc() {
	LCL().SysCallN(4652, m.Instance())
}

func (m *TPrinter) EndDoc() {
	LCL().SysCallN(4663, m.Instance())
}

func (m *TPrinter) NewPage() {
	LCL().SysCallN(4667, m.Instance())
}

func (m *TPrinter) BeginPage() {
	LCL().SysCallN(4653, m.Instance())
}

func (m *TPrinter) EndPage() {
	LCL().SysCallN(4664, m.Instance())
}

func (m *TPrinter) Refresh() {
	LCL().SysCallN(4680, m.Instance())
}

func (m *TPrinter) SetPrinter(aName string) {
	LCL().SysCallN(4682, m.Instance(), PascalStr(aName))
}

func (m *TPrinter) RestoreDefaultBin() {
	LCL().SysCallN(4681, m.Instance())
}
