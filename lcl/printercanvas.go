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

// IPrinterCanvas Parent: ICanvas
type IPrinterCanvas interface {
	ICanvas
	BeginDoc()                                      // procedure
	NewPage()                                       // procedure
	BeginPage()                                     // procedure
	EndPage()                                       // procedure
	EndDoc()                                        // procedure
	Printer() IPrinter                              // property Printer Getter
	Title() string                                  // property Title Getter
	SetTitle(value string)                          // property Title Setter
	PageHeight() int32                              // property PageHeight Getter
	PageWidth() int32                               // property PageWidth Getter
	PaperWidth() int32                              // property PaperWidth Getter
	SetPaperWidth(value int32)                      // property PaperWidth Setter
	PaperHeight() int32                             // property PaperHeight Getter
	SetPaperHeight(value int32)                     // property PaperHeight Setter
	PageNumber() int32                              // property PageNumber Getter
	TopMargin() int32                               // property TopMargin Getter
	SetTopMargin(value int32)                       // property TopMargin Setter
	LeftMargin() int32                              // property LeftMargin Getter
	SetLeftMargin(value int32)                      // property LeftMargin Setter
	BottomMargin() int32                            // property BottomMargin Getter
	SetBottomMargin(value int32)                    // property BottomMargin Setter
	RightMargin() int32                             // property RightMargin Getter
	SetRightMargin(value int32)                     // property RightMargin Setter
	Orientation() types.TPrinterOrientation         // property Orientation Getter
	SetOrientation(value types.TPrinterOrientation) // property Orientation Setter
	XDPI() int32                                    // property XDPI Getter
	SetXDPI(value int32)                            // property XDPI Setter
	YDPI() int32                                    // property YDPI Getter
	SetYDPI(value int32)                            // property YDPI Setter
}

type TPrinterCanvas struct {
	TCanvas
}

func (m *TPrinterCanvas) BeginDoc() {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(1, m.Instance())
}

func (m *TPrinterCanvas) NewPage() {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(2, m.Instance())
}

func (m *TPrinterCanvas) BeginPage() {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(3, m.Instance())
}

func (m *TPrinterCanvas) EndPage() {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(4, m.Instance())
}

func (m *TPrinterCanvas) EndDoc() {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(5, m.Instance())
}

func (m *TPrinterCanvas) Printer() IPrinter {
	if !m.IsValid() {
		return nil
	}
	r := printerCanvasAPI().SysCallN(6, m.Instance())
	return AsPrinter(r)
}

func (m *TPrinterCanvas) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := printerCanvasAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TPrinterCanvas) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TPrinterCanvas) PageHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) PageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) PaperWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetPaperWidth(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) PaperHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetPaperHeight(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) PageNumber() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(12, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) TopMargin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetTopMargin(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) LeftMargin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetLeftMargin(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) BottomMargin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetBottomMargin(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) RightMargin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetRightMargin(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) Orientation() types.TPrinterOrientation {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(17, 0, m.Instance())
	return types.TPrinterOrientation(r)
}

func (m *TPrinterCanvas) SetOrientation(value types.TPrinterOrientation) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) XDPI() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetXDPI(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TPrinterCanvas) YDPI() int32 {
	if !m.IsValid() {
		return 0
	}
	r := printerCanvasAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TPrinterCanvas) SetYDPI(value int32) {
	if !m.IsValid() {
		return
	}
	printerCanvasAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

// NewPrinterCanvas class constructor
func NewPrinterCanvas(printer IPrinter) IPrinterCanvas {
	r := printerCanvasAPI().SysCallN(0, base.GetObjectUintptr(printer))
	return AsPrinterCanvas(r)
}

var (
	printerCanvasOnce   base.Once
	printerCanvasImport *imports.Imports = nil
)

func printerCanvasAPI() *imports.Imports {
	printerCanvasOnce.Do(func() {
		printerCanvasImport = api.NewDefaultImports()
		printerCanvasImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPrinterCanvas_Create", 0), // constructor NewPrinterCanvas
			/* 1 */ imports.NewTable("TPrinterCanvas_BeginDoc", 0), // procedure BeginDoc
			/* 2 */ imports.NewTable("TPrinterCanvas_NewPage", 0), // procedure NewPage
			/* 3 */ imports.NewTable("TPrinterCanvas_BeginPage", 0), // procedure BeginPage
			/* 4 */ imports.NewTable("TPrinterCanvas_EndPage", 0), // procedure EndPage
			/* 5 */ imports.NewTable("TPrinterCanvas_EndDoc", 0), // procedure EndDoc
			/* 6 */ imports.NewTable("TPrinterCanvas_Printer", 0), // property Printer
			/* 7 */ imports.NewTable("TPrinterCanvas_Title", 0), // property Title
			/* 8 */ imports.NewTable("TPrinterCanvas_PageHeight", 0), // property PageHeight
			/* 9 */ imports.NewTable("TPrinterCanvas_PageWidth", 0), // property PageWidth
			/* 10 */ imports.NewTable("TPrinterCanvas_PaperWidth", 0), // property PaperWidth
			/* 11 */ imports.NewTable("TPrinterCanvas_PaperHeight", 0), // property PaperHeight
			/* 12 */ imports.NewTable("TPrinterCanvas_PageNumber", 0), // property PageNumber
			/* 13 */ imports.NewTable("TPrinterCanvas_TopMargin", 0), // property TopMargin
			/* 14 */ imports.NewTable("TPrinterCanvas_LeftMargin", 0), // property LeftMargin
			/* 15 */ imports.NewTable("TPrinterCanvas_BottomMargin", 0), // property BottomMargin
			/* 16 */ imports.NewTable("TPrinterCanvas_RightMargin", 0), // property RightMargin
			/* 17 */ imports.NewTable("TPrinterCanvas_Orientation", 0), // property Orientation
			/* 18 */ imports.NewTable("TPrinterCanvas_XDPI", 0), // property XDPI
			/* 19 */ imports.NewTable("TPrinterCanvas_YDPI", 0), // property YDPI
		}
	})
	return printerCanvasImport
}
