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

// IThemeServices Parent: IObject
// TThemeServices is a small foot print class to provide the user with pure
// Windows XP theme related abilities like painting elements and text or
// retrieving certain info.
type IThemeServices interface {
	IObject
	DottedBrush() HBRUSH                                                                                                                                       // property
	ThemesAvailable() bool                                                                                                                                     // property
	ThemesEnabled() bool                                                                                                                                       // property
	IsDisabled(Details *TThemedElementDetails) bool                                                                                                            // function
	IsPushed(Details *TThemedElementDetails) bool                                                                                                              // function
	IsHot(Details *TThemedElementDetails) bool                                                                                                                 // function
	IsChecked(Details *TThemedElementDetails) bool                                                                                                             // function
	IsMixed(Details *TThemedElementDetails) bool                                                                                                               // function
	GetElementDetails(Detail TThemedButton) (resultThemedElementDetails TThemedElementDetails)                                                                 // function
	GetElementDetails1(Detail TThemedClock) (resultThemedElementDetails TThemedElementDetails)                                                                 // function
	GetElementDetails2(Detail TThemedComboBox) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails3(Detail TThemedEdit) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails4(Detail TThemedExplorerBar) (resultThemedElementDetails TThemedElementDetails)                                                           // function
	GetElementDetails5(Detail TThemedHeader) (resultThemedElementDetails TThemedElementDetails)                                                                // function
	GetElementDetails6(Detail TThemedListView) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails7(Detail TThemedMenu) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails8(Detail TThemedPage) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails9(Detail TThemedProgress) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails10(Detail TThemedRebar) (resultThemedElementDetails TThemedElementDetails)                                                                // function
	GetElementDetails11(Detail TThemedScrollBar) (resultThemedElementDetails TThemedElementDetails)                                                            // function
	GetElementDetails12(Detail TThemedSpin) (resultThemedElementDetails TThemedElementDetails)                                                                 // function
	GetElementDetails13(Detail TThemedStartPanel) (resultThemedElementDetails TThemedElementDetails)                                                           // function
	GetElementDetails14(Detail TThemedStatus) (resultThemedElementDetails TThemedElementDetails)                                                               // function
	GetElementDetails15(Detail TThemedTab) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails16(Detail TThemedTaskBand) (resultThemedElementDetails TThemedElementDetails)                                                             // function
	GetElementDetails17(Detail TThemedTaskBar) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails18(Detail TThemedToolBar) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails19(Detail TThemedToolTip) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails20(Detail TThemedTrackBar) (resultThemedElementDetails TThemedElementDetails)                                                             // function
	GetElementDetails21(Detail TThemedTrayNotify) (resultThemedElementDetails TThemedElementDetails)                                                           // function
	GetElementDetails22(Detail TThemedTreeview) (resultThemedElementDetails TThemedElementDetails)                                                             // function
	GetElementDetails23(Detail TThemedWindow) (resultThemedElementDetails TThemedElementDetails)                                                               // function
	GetDetailSizeForWindow(Details *TThemedElementDetails, AWindow HWND) (resultSize TSize)                                                                    // function
	GetDetailSizeForPPI(Details *TThemedElementDetails, PPI int32) (resultSize TSize)                                                                          // function
	GetDetailRegion(DC HDC, Details *TThemedElementDetails, R *TRect) HRGN                                                                                     // function
	GetStockImage(StockID int32, OutImage, OutMask *HBITMAP) bool                                                                                              // function
	GetStockImage1(StockID int32, AWidth, AHeight int32, OutImage, OutMask *HBITMAP) bool                                                                      // function
	GetOption(AOption TThemeOption) int32                                                                                                                      // function
	GetTextExtent(DC HDC, Details *TThemedElementDetails, S string, Flags uint32, BoundingRect *TRect) (resultRect TRect)                                      // function
	ColorToRGB(Color int32, Details *PThemedElementDetails) COLORREF                                                                                           // function
	ContentRect(DC HDC, Details *TThemedElementDetails, BoundingRect *TRect) (resultRect TRect)                                                                // function
	HasTransparentParts(Details *TThemedElementDetails) bool                                                                                                   // function
	IntfDoOnThemeChange()                                                                                                                                      // procedure
	DrawEdge(DC HDC, Details *TThemedElementDetails, R *TRect, Edge, Flags uint32, AContentRect *TRect)                                                        // procedure
	DrawElement(DC HDC, Details *TThemedElementDetails, R *TRect, ClipRect *TRect)                                                                             // procedure
	DrawIcon(DC HDC, Details *TThemedElementDetails, R *TRect, himl HIMAGELIST, Index int32)                                                                   // procedure
	DrawIcon1(ACanvas IPersistent, Details *TThemedElementDetails, P *TPoint, AImageList IPersistent, Index int32, AImageWidth int32, ARefControl IPersistent) // procedure
	DrawParentBackground(Window HWND, Target HDC, Details *PThemedElementDetails, OnlyIfTransparent bool, Bounds *TRect)                                       // procedure
	DrawText(DC HDC, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32)                                                                 // procedure
	DrawText1(ACanvas IPersistent, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32)                                                   // procedure
	PaintBorder(Control IObject, EraseLRCorner bool)                                                                                                           // procedure
	UpdateThemes()                                                                                                                                             // procedure
	SetOnThemeChange(fn TNotifyEvent)                                                                                                                          // property event
}

// TThemeServices Parent: TObject
// TThemeServices is a small foot print class to provide the user with pure
// Windows XP theme related abilities like painting elements and text or
// retrieving certain info.
type TThemeServices struct {
	TObject
	themeChangePtr uintptr
}

func NewThemeServices() IThemeServices {
	r1 := hemeServicesImportAPI().SysCallN(3)
	return AsThemeServices(r1)
}

func (m *TThemeServices) DottedBrush() HBRUSH {
	r1 := hemeServicesImportAPI().SysCallN(4, m.Instance())
	return HBRUSH(r1)
}

func (m *TThemeServices) ThemesAvailable() bool {
	r1 := hemeServicesImportAPI().SysCallN(52, m.Instance())
	return GoBool(r1)
}

func (m *TThemeServices) ThemesEnabled() bool {
	r1 := hemeServicesImportAPI().SysCallN(53, m.Instance())
	return GoBool(r1)
}

func (m *TThemeServices) IsDisabled(Details *TThemedElementDetails) bool {
	r1 := hemeServicesImportAPI().SysCallN(46, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsPushed(Details *TThemedElementDetails) bool {
	r1 := hemeServicesImportAPI().SysCallN(49, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsHot(Details *TThemedElementDetails) bool {
	r1 := hemeServicesImportAPI().SysCallN(47, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsChecked(Details *TThemedElementDetails) bool {
	r1 := hemeServicesImportAPI().SysCallN(45, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsMixed(Details *TThemedElementDetails) bool {
	r1 := hemeServicesImportAPI().SysCallN(48, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) GetElementDetails(Detail TThemedButton) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(15, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails1(Detail TThemedClock) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(16, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails2(Detail TThemedComboBox) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(27, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails3(Detail TThemedEdit) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(32, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails4(Detail TThemedExplorerBar) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(33, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails5(Detail TThemedHeader) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(34, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails6(Detail TThemedListView) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(35, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails7(Detail TThemedMenu) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(36, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails8(Detail TThemedPage) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(37, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails9(Detail TThemedProgress) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(38, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails10(Detail TThemedRebar) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(17, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails11(Detail TThemedScrollBar) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(18, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails12(Detail TThemedSpin) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(19, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails13(Detail TThemedStartPanel) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(20, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails14(Detail TThemedStatus) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(21, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails15(Detail TThemedTab) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(22, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails16(Detail TThemedTaskBand) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(23, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails17(Detail TThemedTaskBar) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(24, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails18(Detail TThemedToolBar) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(25, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails19(Detail TThemedToolTip) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(26, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails20(Detail TThemedTrackBar) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(28, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails21(Detail TThemedTrayNotify) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(29, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails22(Detail TThemedTreeview) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(30, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails23(Detail TThemedWindow) (resultThemedElementDetails TThemedElementDetails) {
	hemeServicesImportAPI().SysCallN(31, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetDetailSizeForWindow(Details *TThemedElementDetails, AWindow HWND) (resultSize TSize) {
	hemeServicesImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(Details)), uintptr(AWindow), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TThemeServices) GetDetailSizeForPPI(Details *TThemedElementDetails, PPI int32) (resultSize TSize) {
	hemeServicesImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(Details)), uintptr(PPI), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TThemeServices) GetDetailRegion(DC HDC, Details *TThemedElementDetails, R *TRect) HRGN {
	r1 := hemeServicesImportAPI().SysCallN(12, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)))
	return HRGN(r1)
}

func (m *TThemeServices) GetStockImage(StockID int32, OutImage, OutMask *HBITMAP) bool {
	var result1 uintptr
	var result2 uintptr
	r1 := hemeServicesImportAPI().SysCallN(40, m.Instance(), uintptr(StockID), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutImage = HBITMAP(result1)
	*OutMask = HBITMAP(result2)
	return GoBool(r1)
}

func (m *TThemeServices) GetStockImage1(StockID int32, AWidth, AHeight int32, OutImage, OutMask *HBITMAP) bool {
	var result2 uintptr
	var result3 uintptr
	r1 := hemeServicesImportAPI().SysCallN(41, m.Instance(), uintptr(StockID), uintptr(AWidth), uintptr(AHeight), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)))
	*OutImage = HBITMAP(result2)
	*OutMask = HBITMAP(result3)
	return GoBool(r1)
}

func (m *TThemeServices) GetOption(AOption TThemeOption) int32 {
	r1 := hemeServicesImportAPI().SysCallN(39, m.Instance(), uintptr(AOption))
	return int32(r1)
}

func (m *TThemeServices) GetTextExtent(DC HDC, Details *TThemedElementDetails, S string, Flags uint32, BoundingRect *TRect) (resultRect TRect) {
	hemeServicesImportAPI().SysCallN(42, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), PascalStr(S), uintptr(Flags), uintptr(unsafePointer(BoundingRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TThemeServices) ColorToRGB(Color int32, Details *PThemedElementDetails) COLORREF {
	r1 := hemeServicesImportAPI().SysCallN(1, m.Instance(), uintptr(Color), uintptr(unsafePointer(Details)))
	return COLORREF(r1)
}

func (m *TThemeServices) ContentRect(DC HDC, Details *TThemedElementDetails, BoundingRect *TRect) (resultRect TRect) {
	hemeServicesImportAPI().SysCallN(2, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(BoundingRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TThemeServices) HasTransparentParts(Details *TThemedElementDetails) bool {
	r1 := hemeServicesImportAPI().SysCallN(43, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func ThemeServicesClass() TClass {
	ret := hemeServicesImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TThemeServices) IntfDoOnThemeChange() {
	hemeServicesImportAPI().SysCallN(44, m.Instance())
}

func (m *TThemeServices) DrawEdge(DC HDC, Details *TThemedElementDetails, R *TRect, Edge, Flags uint32, AContentRect *TRect) {
	hemeServicesImportAPI().SysCallN(5, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)), uintptr(Edge), uintptr(Flags), uintptr(unsafePointer(AContentRect)))
}

func (m *TThemeServices) DrawElement(DC HDC, Details *TThemedElementDetails, R *TRect, ClipRect *TRect) {
	hemeServicesImportAPI().SysCallN(6, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)), uintptr(unsafePointer(ClipRect)))
}

func (m *TThemeServices) DrawIcon(DC HDC, Details *TThemedElementDetails, R *TRect, himl HIMAGELIST, Index int32) {
	hemeServicesImportAPI().SysCallN(7, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)), uintptr(himl), uintptr(Index))
}

func (m *TThemeServices) DrawIcon1(ACanvas IPersistent, Details *TThemedElementDetails, P *TPoint, AImageList IPersistent, Index int32, AImageWidth int32, ARefControl IPersistent) {
	hemeServicesImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(ACanvas), uintptr(unsafePointer(Details)), uintptr(unsafePointer(P)), GetObjectUintptr(AImageList), uintptr(Index), uintptr(AImageWidth), GetObjectUintptr(ARefControl))
}

func (m *TThemeServices) DrawParentBackground(Window HWND, Target HDC, Details *PThemedElementDetails, OnlyIfTransparent bool, Bounds *TRect) {
	hemeServicesImportAPI().SysCallN(9, m.Instance(), uintptr(Window), uintptr(Target), uintptr(unsafePointer(Details)), PascalBool(OnlyIfTransparent), uintptr(unsafePointer(Bounds)))
}

func (m *TThemeServices) DrawText(DC HDC, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32) {
	hemeServicesImportAPI().SysCallN(10, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), PascalStr(S), uintptr(unsafePointer(R)), uintptr(Flags), uintptr(Flags2))
}

func (m *TThemeServices) DrawText1(ACanvas IPersistent, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32) {
	hemeServicesImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(ACanvas), uintptr(unsafePointer(Details)), PascalStr(S), uintptr(unsafePointer(R)), uintptr(Flags), uintptr(Flags2))
}

func (m *TThemeServices) PaintBorder(Control IObject, EraseLRCorner bool) {
	hemeServicesImportAPI().SysCallN(50, m.Instance(), GetObjectUintptr(Control), PascalBool(EraseLRCorner))
}

func (m *TThemeServices) UpdateThemes() {
	hemeServicesImportAPI().SysCallN(54, m.Instance())
}

func (m *TThemeServices) SetOnThemeChange(fn TNotifyEvent) {
	if m.themeChangePtr != 0 {
		RemoveEventElement(m.themeChangePtr)
	}
	m.themeChangePtr = MakeEventDataPtr(fn)
	hemeServicesImportAPI().SysCallN(51, m.Instance(), m.themeChangePtr)
}

var (
	hemeServicesImport       *imports.Imports = nil
	hemeServicesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ThemeServices_Class", 0),
		/*1*/ imports.NewTable("ThemeServices_ColorToRGB", 0),
		/*2*/ imports.NewTable("ThemeServices_ContentRect", 0),
		/*3*/ imports.NewTable("ThemeServices_Create", 0),
		/*4*/ imports.NewTable("ThemeServices_DottedBrush", 0),
		/*5*/ imports.NewTable("ThemeServices_DrawEdge", 0),
		/*6*/ imports.NewTable("ThemeServices_DrawElement", 0),
		/*7*/ imports.NewTable("ThemeServices_DrawIcon", 0),
		/*8*/ imports.NewTable("ThemeServices_DrawIcon1", 0),
		/*9*/ imports.NewTable("ThemeServices_DrawParentBackground", 0),
		/*10*/ imports.NewTable("ThemeServices_DrawText", 0),
		/*11*/ imports.NewTable("ThemeServices_DrawText1", 0),
		/*12*/ imports.NewTable("ThemeServices_GetDetailRegion", 0),
		/*13*/ imports.NewTable("ThemeServices_GetDetailSizeForPPI", 0),
		/*14*/ imports.NewTable("ThemeServices_GetDetailSizeForWindow", 0),
		/*15*/ imports.NewTable("ThemeServices_GetElementDetails", 0),
		/*16*/ imports.NewTable("ThemeServices_GetElementDetails1", 0),
		/*17*/ imports.NewTable("ThemeServices_GetElementDetails10", 0),
		/*18*/ imports.NewTable("ThemeServices_GetElementDetails11", 0),
		/*19*/ imports.NewTable("ThemeServices_GetElementDetails12", 0),
		/*20*/ imports.NewTable("ThemeServices_GetElementDetails13", 0),
		/*21*/ imports.NewTable("ThemeServices_GetElementDetails14", 0),
		/*22*/ imports.NewTable("ThemeServices_GetElementDetails15", 0),
		/*23*/ imports.NewTable("ThemeServices_GetElementDetails16", 0),
		/*24*/ imports.NewTable("ThemeServices_GetElementDetails17", 0),
		/*25*/ imports.NewTable("ThemeServices_GetElementDetails18", 0),
		/*26*/ imports.NewTable("ThemeServices_GetElementDetails19", 0),
		/*27*/ imports.NewTable("ThemeServices_GetElementDetails2", 0),
		/*28*/ imports.NewTable("ThemeServices_GetElementDetails20", 0),
		/*29*/ imports.NewTable("ThemeServices_GetElementDetails21", 0),
		/*30*/ imports.NewTable("ThemeServices_GetElementDetails22", 0),
		/*31*/ imports.NewTable("ThemeServices_GetElementDetails23", 0),
		/*32*/ imports.NewTable("ThemeServices_GetElementDetails3", 0),
		/*33*/ imports.NewTable("ThemeServices_GetElementDetails4", 0),
		/*34*/ imports.NewTable("ThemeServices_GetElementDetails5", 0),
		/*35*/ imports.NewTable("ThemeServices_GetElementDetails6", 0),
		/*36*/ imports.NewTable("ThemeServices_GetElementDetails7", 0),
		/*37*/ imports.NewTable("ThemeServices_GetElementDetails8", 0),
		/*38*/ imports.NewTable("ThemeServices_GetElementDetails9", 0),
		/*39*/ imports.NewTable("ThemeServices_GetOption", 0),
		/*40*/ imports.NewTable("ThemeServices_GetStockImage", 0),
		/*41*/ imports.NewTable("ThemeServices_GetStockImage1", 0),
		/*42*/ imports.NewTable("ThemeServices_GetTextExtent", 0),
		/*43*/ imports.NewTable("ThemeServices_HasTransparentParts", 0),
		/*44*/ imports.NewTable("ThemeServices_IntfDoOnThemeChange", 0),
		/*45*/ imports.NewTable("ThemeServices_IsChecked", 0),
		/*46*/ imports.NewTable("ThemeServices_IsDisabled", 0),
		/*47*/ imports.NewTable("ThemeServices_IsHot", 0),
		/*48*/ imports.NewTable("ThemeServices_IsMixed", 0),
		/*49*/ imports.NewTable("ThemeServices_IsPushed", 0),
		/*50*/ imports.NewTable("ThemeServices_PaintBorder", 0),
		/*51*/ imports.NewTable("ThemeServices_SetOnThemeChange", 0),
		/*52*/ imports.NewTable("ThemeServices_ThemesAvailable", 0),
		/*53*/ imports.NewTable("ThemeServices_ThemesEnabled", 0),
		/*54*/ imports.NewTable("ThemeServices_UpdateThemes", 0),
	}
)

func hemeServicesImportAPI() *imports.Imports {
	if hemeServicesImport == nil {
		hemeServicesImport = NewDefaultImports()
		hemeServicesImport.SetImportTable(hemeServicesImportTables)
		hemeServicesImportTables = nil
	}
	return hemeServicesImport
}
