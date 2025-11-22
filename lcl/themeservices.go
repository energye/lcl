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

// IThemeServices Parent: IObject
type IThemeServices interface {
	IObject
	// IsDisabled
	//  state helpers
	IsDisabled(details TThemedElementDetails) bool                                                                                                                                                          // function
	IsPushed(details TThemedElementDetails) bool                                                                                                                                                            // function
	IsHot(details TThemedElementDetails) bool                                                                                                                                                               // function
	IsChecked(details TThemedElementDetails) bool                                                                                                                                                           // function
	IsMixed(details TThemedElementDetails) bool                                                                                                                                                             // function
	GetElementDetailsWithThemedButton(detail types.TThemedButton) TThemedElementDetails                                                                                                                     // function
	GetElementDetailsWithThemedClock(detail types.TThemedClock) TThemedElementDetails                                                                                                                       // function
	GetElementDetailsWithThemedComboBox(detail types.TThemedComboBox) TThemedElementDetails                                                                                                                 // function
	GetElementDetailsWithThemedEdit(detail types.TThemedEdit) TThemedElementDetails                                                                                                                         // function
	GetElementDetailsWithThemedExplorerBar(detail types.TThemedExplorerBar) TThemedElementDetails                                                                                                           // function
	GetElementDetailsWithThemedHeader(detail types.TThemedHeader) TThemedElementDetails                                                                                                                     // function
	GetElementDetailsWithThemedListView(detail types.TThemedListView) TThemedElementDetails                                                                                                                 // function
	GetElementDetailsWithThemedMenu(detail types.TThemedMenu) TThemedElementDetails                                                                                                                         // function
	GetElementDetailsWithThemedPage(detail types.TThemedPage) TThemedElementDetails                                                                                                                         // function
	GetElementDetailsWithThemedProgress(detail types.TThemedProgress) TThemedElementDetails                                                                                                                 // function
	GetElementDetailsWithThemedRebar(detail types.TThemedRebar) TThemedElementDetails                                                                                                                       // function
	GetElementDetailsWithThemedScrollBar(detail types.TThemedScrollBar) TThemedElementDetails                                                                                                               // function
	GetElementDetailsWithThemedSpin(detail types.TThemedSpin) TThemedElementDetails                                                                                                                         // function
	GetElementDetailsWithThemedStartPanel(detail types.TThemedStartPanel) TThemedElementDetails                                                                                                             // function
	GetElementDetailsWithThemedStatus(detail types.TThemedStatus) TThemedElementDetails                                                                                                                     // function
	GetElementDetailsWithThemedTab(detail types.TThemedTab) TThemedElementDetails                                                                                                                           // function
	GetElementDetailsWithThemedTaskBand(detail types.TThemedTaskBand) TThemedElementDetails                                                                                                                 // function
	GetElementDetailsWithThemedTaskBar(detail types.TThemedTaskBar) TThemedElementDetails                                                                                                                   // function
	GetElementDetailsWithThemedToolBar(detail types.TThemedToolBar) TThemedElementDetails                                                                                                                   // function
	GetElementDetailsWithThemedToolTip(detail types.TThemedToolTip) TThemedElementDetails                                                                                                                   // function
	GetElementDetailsWithThemedTrackBar(detail types.TThemedTrackBar) TThemedElementDetails                                                                                                                 // function
	GetElementDetailsWithThemedTrayNotify(detail types.TThemedTrayNotify) TThemedElementDetails                                                                                                             // function
	GetElementDetailsWithThemedTreeview(detail types.TThemedTreeview) TThemedElementDetails                                                                                                                 // function
	GetElementDetailsWithThemedWindow(detail types.TThemedWindow) TThemedElementDetails                                                                                                                     // function
	GetDetailSizeForWindow(details TThemedElementDetails, window types.HWND) types.TSize                                                                                                                    // function
	GetDetailSizeForPPI(details TThemedElementDetails, pPI int32) types.TSize                                                                                                                               // function
	GetDetailRegion(dC types.HDC, details TThemedElementDetails, R types.TRect) types.HRGN                                                                                                                  // function
	GetStockImageWithIntHBitmapX2(stockID int32, outImage *types.HBitmap, outMask *types.HBitmap) bool                                                                                                      // function
	GetStockImageWithIntX3HBitmapX2(stockID int32, width int32, height int32, outImage *types.HBitmap, outMask *types.HBitmap) bool                                                                         // function
	GetOption(option types.TThemeOption) int32                                                                                                                                                              // function
	GetTextExtent(dC types.HDC, details TThemedElementDetails, S string, flags uint32, boundingRect types.TRect) types.TRect                                                                                // function
	ColorToRGB(color int32, details TThemedElementDetails) types.COLORREF                                                                                                                                   // function
	ContentRect(dC types.HDC, details TThemedElementDetails, boundingRect types.TRect) types.TRect                                                                                                          // function
	HasTransparentParts(details TThemedElementDetails) bool                                                                                                                                                 // function
	IntfDoOnThemeChange()                                                                                                                                                                                   // procedure
	DrawEdge(dC types.HDC, details TThemedElementDetails, R types.TRect, edge uint32, flags uint32, contentRect types.TRect)                                                                                // procedure
	DrawElement(dC types.HDC, details TThemedElementDetails, R types.TRect, clipRect types.TRect)                                                                                                           // procedure
	DrawIconWithHDCThemedElementDetailsRectHIMAGELISTInt(dC types.HDC, details TThemedElementDetails, R types.TRect, himl types.HIMAGELIST, index int32)                                                    // procedure
	DrawIconWithPersistentX3ThemedElementDetailsPointIntX2(canvas IPersistent, details TThemedElementDetails, P types.TPoint, imageList IPersistent, index int32, imageWidth int32, refControl IPersistent) // procedure
	DrawParentBackground(window types.HWND, target types.HDC, details TThemedElementDetails, onlyIfTransparent bool, bounds types.TRect)                                                                    // procedure
	DrawTextWithHDCThemedElementDetailsStringRectCardinalX2(dC types.HDC, details TThemedElementDetails, S string, R types.TRect, flags uint32, flags2 uint32)                                              // procedure
	DrawTextWithPersistentThemedElementDetailsStringRectCardinalX2(canvas IPersistent, details TThemedElementDetails, S string, R types.TRect, flags uint32, flags2 uint32)                                 // procedure
	PaintBorder(control IObject, eraseLRCorner bool)                                                                                                                                                        // procedure
	UpdateThemes()                                                                                                                                                                                          // procedure
	DottedBrush() types.HBRUSH                                                                                                                                                                              // property DottedBrush Getter
	ThemesAvailable() bool                                                                                                                                                                                  // property ThemesAvailable Getter
	ThemesEnabled() bool                                                                                                                                                                                    // property ThemesEnabled Getter
	SetOnThemeChange(fn TNotifyEvent)                                                                                                                                                                       // property event
}

type TThemeServices struct {
	TObject
}

func (m *TThemeServices) IsDisabled(details TThemedElementDetails) bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&details)))
	return api.GoBool(r)
}

func (m *TThemeServices) IsPushed(details TThemedElementDetails) bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&details)))
	return api.GoBool(r)
}

func (m *TThemeServices) IsHot(details TThemedElementDetails) bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&details)))
	return api.GoBool(r)
}

func (m *TThemeServices) IsChecked(details TThemedElementDetails) bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&details)))
	return api.GoBool(r)
}

func (m *TThemeServices) IsMixed(details TThemedElementDetails) bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&details)))
	return api.GoBool(r)
}

func (m *TThemeServices) GetElementDetailsWithThemedButton(detail types.TThemedButton) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(6, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedClock(detail types.TThemedClock) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(7, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedComboBox(detail types.TThemedComboBox) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(8, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedEdit(detail types.TThemedEdit) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(9, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedExplorerBar(detail types.TThemedExplorerBar) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(10, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedHeader(detail types.TThemedHeader) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(11, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedListView(detail types.TThemedListView) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(12, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedMenu(detail types.TThemedMenu) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(13, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedPage(detail types.TThemedPage) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(14, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedProgress(detail types.TThemedProgress) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(15, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedRebar(detail types.TThemedRebar) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(16, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedScrollBar(detail types.TThemedScrollBar) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(17, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedSpin(detail types.TThemedSpin) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(18, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedStartPanel(detail types.TThemedStartPanel) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(19, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedStatus(detail types.TThemedStatus) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(20, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedTab(detail types.TThemedTab) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(21, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedTaskBand(detail types.TThemedTaskBand) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(22, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedTaskBar(detail types.TThemedTaskBar) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(23, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedToolBar(detail types.TThemedToolBar) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(24, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedToolTip(detail types.TThemedToolTip) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(25, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedTrackBar(detail types.TThemedTrackBar) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(26, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedTrayNotify(detail types.TThemedTrayNotify) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(27, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedTreeview(detail types.TThemedTreeview) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(28, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetElementDetailsWithThemedWindow(detail types.TThemedWindow) (result TThemedElementDetails) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(29, m.Instance(), uintptr(detail), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetDetailSizeForWindow(details TThemedElementDetails, window types.HWND) (result types.TSize) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(30, m.Instance(), uintptr(base.UnsafePointer(&details)), uintptr(window), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetDetailSizeForPPI(details TThemedElementDetails, pPI int32) (result types.TSize) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(31, m.Instance(), uintptr(base.UnsafePointer(&details)), uintptr(pPI), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) GetDetailRegion(dC types.HDC, details TThemedElementDetails, R types.TRect) types.HRGN {
	if !m.IsValid() {
		return 0
	}
	r := themeServicesAPI().SysCallN(32, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&details)), uintptr(base.UnsafePointer(&R)))
	return types.HRGN(r)
}

func (m *TThemeServices) GetStockImageWithIntHBitmapX2(stockID int32, outImage *types.HBitmap, outMask *types.HBitmap) bool {
	if !m.IsValid() {
		return false
	}
	var imagePtr uintptr
	var maskPtr uintptr
	r := themeServicesAPI().SysCallN(33, m.Instance(), uintptr(stockID), uintptr(base.UnsafePointer(&imagePtr)), uintptr(base.UnsafePointer(&maskPtr)))
	*outImage = types.HBitmap(imagePtr)
	*outMask = types.HBitmap(maskPtr)
	return api.GoBool(r)
}

func (m *TThemeServices) GetStockImageWithIntX3HBitmapX2(stockID int32, width int32, height int32, outImage *types.HBitmap, outMask *types.HBitmap) bool {
	if !m.IsValid() {
		return false
	}
	var imagePtr uintptr
	var maskPtr uintptr
	r := themeServicesAPI().SysCallN(34, m.Instance(), uintptr(stockID), uintptr(width), uintptr(height), uintptr(base.UnsafePointer(&imagePtr)), uintptr(base.UnsafePointer(&maskPtr)))
	*outImage = types.HBitmap(imagePtr)
	*outMask = types.HBitmap(maskPtr)
	return api.GoBool(r)
}

func (m *TThemeServices) GetOption(option types.TThemeOption) int32 {
	if !m.IsValid() {
		return 0
	}
	r := themeServicesAPI().SysCallN(35, m.Instance(), uintptr(option))
	return int32(r)
}

func (m *TThemeServices) GetTextExtent(dC types.HDC, details TThemedElementDetails, S string, flags uint32, boundingRect types.TRect) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(36, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&details)), api.PasStr(S), uintptr(flags), uintptr(base.UnsafePointer(&boundingRect)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) ColorToRGB(color int32, details TThemedElementDetails) types.COLORREF {
	if !m.IsValid() {
		return 0
	}
	r := themeServicesAPI().SysCallN(37, m.Instance(), uintptr(color), uintptr(base.UnsafePointer(&details)))
	return types.COLORREF(r)
}

func (m *TThemeServices) ContentRect(dC types.HDC, details TThemedElementDetails, boundingRect types.TRect) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(38, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&details)), uintptr(base.UnsafePointer(&boundingRect)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TThemeServices) HasTransparentParts(details TThemedElementDetails) bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(39, m.Instance(), uintptr(base.UnsafePointer(&details)))
	return api.GoBool(r)
}

func (m *TThemeServices) IntfDoOnThemeChange() {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(40, m.Instance())
}

func (m *TThemeServices) DrawEdge(dC types.HDC, details TThemedElementDetails, R types.TRect, edge uint32, flags uint32, contentRect types.TRect) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(41, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&details)), uintptr(base.UnsafePointer(&R)), uintptr(edge), uintptr(flags), uintptr(base.UnsafePointer(&contentRect)))
}

func (m *TThemeServices) DrawElement(dC types.HDC, details TThemedElementDetails, R types.TRect, clipRect types.TRect) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(42, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&details)), uintptr(base.UnsafePointer(&R)), uintptr(base.UnsafePointer(&clipRect)))
}

func (m *TThemeServices) DrawIconWithHDCThemedElementDetailsRectHIMAGELISTInt(dC types.HDC, details TThemedElementDetails, R types.TRect, himl types.HIMAGELIST, index int32) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(43, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&details)), uintptr(base.UnsafePointer(&R)), uintptr(himl), uintptr(index))
}

func (m *TThemeServices) DrawIconWithPersistentX3ThemedElementDetailsPointIntX2(canvas IPersistent, details TThemedElementDetails, P types.TPoint, imageList IPersistent, index int32, imageWidth int32, refControl IPersistent) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(44, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&details)), uintptr(base.UnsafePointer(&P)), base.GetObjectUintptr(imageList), uintptr(index), uintptr(imageWidth), base.GetObjectUintptr(refControl))
}

func (m *TThemeServices) DrawParentBackground(window types.HWND, target types.HDC, details TThemedElementDetails, onlyIfTransparent bool, bounds types.TRect) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(45, m.Instance(), uintptr(window), uintptr(target), uintptr(base.UnsafePointer(&details)), api.PasBool(onlyIfTransparent), uintptr(base.UnsafePointer(&bounds)))
}

func (m *TThemeServices) DrawTextWithHDCThemedElementDetailsStringRectCardinalX2(dC types.HDC, details TThemedElementDetails, S string, R types.TRect, flags uint32, flags2 uint32) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(46, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&details)), api.PasStr(S), uintptr(base.UnsafePointer(&R)), uintptr(flags), uintptr(flags2))
}

func (m *TThemeServices) DrawTextWithPersistentThemedElementDetailsStringRectCardinalX2(canvas IPersistent, details TThemedElementDetails, S string, R types.TRect, flags uint32, flags2 uint32) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(47, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&details)), api.PasStr(S), uintptr(base.UnsafePointer(&R)), uintptr(flags), uintptr(flags2))
}

func (m *TThemeServices) PaintBorder(control IObject, eraseLRCorner bool) {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(48, m.Instance(), base.GetObjectUintptr(control), api.PasBool(eraseLRCorner))
}

func (m *TThemeServices) UpdateThemes() {
	if !m.IsValid() {
		return
	}
	themeServicesAPI().SysCallN(49, m.Instance())
}

func (m *TThemeServices) DottedBrush() types.HBRUSH {
	if !m.IsValid() {
		return 0
	}
	r := themeServicesAPI().SysCallN(50, m.Instance())
	return types.HBRUSH(r)
}

func (m *TThemeServices) ThemesAvailable() bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(51, m.Instance())
	return api.GoBool(r)
}

func (m *TThemeServices) ThemesEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := themeServicesAPI().SysCallN(52, m.Instance())
	return api.GoBool(r)
}

func (m *TThemeServices) SetOnThemeChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 53, themeServicesAPI(), api.MakeEventDataPtr(cb))
}

// NewThemeServices class constructor
func NewThemeServices() IThemeServices {
	r := themeServicesAPI().SysCallN(0)
	return AsThemeServices(r)
}

var (
	themeServicesOnce   base.Once
	themeServicesImport *imports.Imports = nil
)

func themeServicesAPI() *imports.Imports {
	themeServicesOnce.Do(func() {
		themeServicesImport = api.NewDefaultImports()
		themeServicesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TThemeServices_Create", 0), // constructor NewThemeServices
			/* 1 */ imports.NewTable("TThemeServices_IsDisabled", 0), // function IsDisabled
			/* 2 */ imports.NewTable("TThemeServices_IsPushed", 0), // function IsPushed
			/* 3 */ imports.NewTable("TThemeServices_IsHot", 0), // function IsHot
			/* 4 */ imports.NewTable("TThemeServices_IsChecked", 0), // function IsChecked
			/* 5 */ imports.NewTable("TThemeServices_IsMixed", 0), // function IsMixed
			/* 6 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedButton", 0), // function GetElementDetailsWithThemedButton
			/* 7 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedClock", 0), // function GetElementDetailsWithThemedClock
			/* 8 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedComboBox", 0), // function GetElementDetailsWithThemedComboBox
			/* 9 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedEdit", 0), // function GetElementDetailsWithThemedEdit
			/* 10 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedExplorerBar", 0), // function GetElementDetailsWithThemedExplorerBar
			/* 11 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedHeader", 0), // function GetElementDetailsWithThemedHeader
			/* 12 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedListView", 0), // function GetElementDetailsWithThemedListView
			/* 13 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedMenu", 0), // function GetElementDetailsWithThemedMenu
			/* 14 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedPage", 0), // function GetElementDetailsWithThemedPage
			/* 15 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedProgress", 0), // function GetElementDetailsWithThemedProgress
			/* 16 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedRebar", 0), // function GetElementDetailsWithThemedRebar
			/* 17 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedScrollBar", 0), // function GetElementDetailsWithThemedScrollBar
			/* 18 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedSpin", 0), // function GetElementDetailsWithThemedSpin
			/* 19 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedStartPanel", 0), // function GetElementDetailsWithThemedStartPanel
			/* 20 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedStatus", 0), // function GetElementDetailsWithThemedStatus
			/* 21 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedTab", 0), // function GetElementDetailsWithThemedTab
			/* 22 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedTaskBand", 0), // function GetElementDetailsWithThemedTaskBand
			/* 23 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedTaskBar", 0), // function GetElementDetailsWithThemedTaskBar
			/* 24 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedToolBar", 0), // function GetElementDetailsWithThemedToolBar
			/* 25 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedToolTip", 0), // function GetElementDetailsWithThemedToolTip
			/* 26 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedTrackBar", 0), // function GetElementDetailsWithThemedTrackBar
			/* 27 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedTrayNotify", 0), // function GetElementDetailsWithThemedTrayNotify
			/* 28 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedTreeview", 0), // function GetElementDetailsWithThemedTreeview
			/* 29 */ imports.NewTable("TThemeServices_GetElementDetailsWithThemedWindow", 0), // function GetElementDetailsWithThemedWindow
			/* 30 */ imports.NewTable("TThemeServices_GetDetailSizeForWindow", 0), // function GetDetailSizeForWindow
			/* 31 */ imports.NewTable("TThemeServices_GetDetailSizeForPPI", 0), // function GetDetailSizeForPPI
			/* 32 */ imports.NewTable("TThemeServices_GetDetailRegion", 0), // function GetDetailRegion
			/* 33 */ imports.NewTable("TThemeServices_GetStockImageWithIntHBitmapX2", 0), // function GetStockImageWithIntHBitmapX2
			/* 34 */ imports.NewTable("TThemeServices_GetStockImageWithIntX3HBitmapX2", 0), // function GetStockImageWithIntX3HBitmapX2
			/* 35 */ imports.NewTable("TThemeServices_GetOption", 0), // function GetOption
			/* 36 */ imports.NewTable("TThemeServices_GetTextExtent", 0), // function GetTextExtent
			/* 37 */ imports.NewTable("TThemeServices_ColorToRGB", 0), // function ColorToRGB
			/* 38 */ imports.NewTable("TThemeServices_ContentRect", 0), // function ContentRect
			/* 39 */ imports.NewTable("TThemeServices_HasTransparentParts", 0), // function HasTransparentParts
			/* 40 */ imports.NewTable("TThemeServices_IntfDoOnThemeChange", 0), // procedure IntfDoOnThemeChange
			/* 41 */ imports.NewTable("TThemeServices_DrawEdge", 0), // procedure DrawEdge
			/* 42 */ imports.NewTable("TThemeServices_DrawElement", 0), // procedure DrawElement
			/* 43 */ imports.NewTable("TThemeServices_DrawIconWithHDCThemedElementDetailsRectHIMAGELISTInt", 0), // procedure DrawIconWithHDCThemedElementDetailsRectHIMAGELISTInt
			/* 44 */ imports.NewTable("TThemeServices_DrawIconWithPersistentX3ThemedElementDetailsPointIntX2", 0), // procedure DrawIconWithPersistentX3ThemedElementDetailsPointIntX2
			/* 45 */ imports.NewTable("TThemeServices_DrawParentBackground", 0), // procedure DrawParentBackground
			/* 46 */ imports.NewTable("TThemeServices_DrawTextWithHDCThemedElementDetailsStringRectCardinalX2", 0), // procedure DrawTextWithHDCThemedElementDetailsStringRectCardinalX2
			/* 47 */ imports.NewTable("TThemeServices_DrawTextWithPersistentThemedElementDetailsStringRectCardinalX2", 0), // procedure DrawTextWithPersistentThemedElementDetailsStringRectCardinalX2
			/* 48 */ imports.NewTable("TThemeServices_PaintBorder", 0), // procedure PaintBorder
			/* 49 */ imports.NewTable("TThemeServices_UpdateThemes", 0), // procedure UpdateThemes
			/* 50 */ imports.NewTable("TThemeServices_DottedBrush", 0), // property DottedBrush
			/* 51 */ imports.NewTable("TThemeServices_ThemesAvailable", 0), // property ThemesAvailable
			/* 52 */ imports.NewTable("TThemeServices_ThemesEnabled", 0), // property ThemesEnabled
			/* 53 */ imports.NewTable("TThemeServices_OnThemeChange", 0), // event OnThemeChange
		}
	})
	return themeServicesImport
}
