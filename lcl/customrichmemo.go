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

// ICustomRichMemo Parent: ICustomMemo
type ICustomRichMemo interface {
	ICustomMemo
	CanPaste() bool                                                                                                                                                                                                                                               // function
	GetTextAttributes(textStart int32, textParams *TFontParams) bool                                                                                                                                                                                              // function
	GetStyleRange(charOfs int32, rangeStart *int32, rangeLen *int32) bool                                                                                                                                                                                         // function
	GetParaAlignmentWithIntParaAlignment(textStart int32, align *types.TParaAlignment) bool                                                                                                                                                                       // function
	GetParaAlignmentWithInt(textStart int32) types.TParaAlignment                                                                                                                                                                                                 // function
	GetParaMetric(textStart int32, metric *TParaMetric) bool                                                                                                                                                                                                      // function
	GetParaNumbering(textStart int32, number *TParaNumbering) bool                                                                                                                                                                                                // function
	GetParaRangeWithIntParaRange(charOfs int32, paraRange *TParaRange) bool                                                                                                                                                                                       // function
	GetParaRangeWithIntX3(charOfs int32, textStart *int32, textLength *int32) bool                                                                                                                                                                                // function
	GetParaTabs(charOfs int32, stopList *TTabStopList) bool                                                                                                                                                                                                       // function
	IsLink(textStart int32) bool                                                                                                                                                                                                                                  // function
	LoadRichText(source IStream) bool                                                                                                                                                                                                                             // function
	SaveRichText(dest IStream) bool                                                                                                                                                                                                                               // function
	InDelText(uTF8Text string, insStartChar int32, replaceLength int32) int32                                                                                                                                                                                     // function
	InDelInline(inlineobj IRichMemoInline, insStartChar int32, replaceLength int32, size types.TSize) int32                                                                                                                                                       // function
	GetText(textStart int32, textLength int32) string                                                                                                                                                                                                             // function
	GetUText(textStart int32, textLength int32) string                                                                                                                                                                                                            // function
	SearchWithStringIntX2SearchOptions(niddle string, start int32, len int32, searchOpt types.TSearchOptions) int32                                                                                                                                               // function
	SearchWithStringIntX4SearchOptions(niddle string, start int32, len int32, searchOpt types.TSearchOptions, textStart *int32, textLength *int32) bool                                                                                                           // function
	Print(params TPrintParams) int32                                                                                                                                                                                                                              // function
	CharAtPos(X int32, Y int32) int32                                                                                                                                                                                                                             // function
	SetTextAttributesWithIntX2FontParams(textStart int32, textLen int32, textParams TFontParams)                                                                                                                                                                  // procedure
	SetParaAlignment(textStart int32, textLen int32, align types.TParaAlignment)                                                                                                                                                                                  // procedure
	SetParaMetric(textStart int32, textLen int32, metric TParaMetric)                                                                                                                                                                                             // procedure
	SetParaNumbering(textStart int32, textLen int32, number TParaNumbering)                                                                                                                                                                                       // procedure
	SetParaTabs(textStart int32, textLen int32, stopList TTabStopList)                                                                                                                                                                                            // procedure
	SetTextAttributesWithIntX2Font(textStart int32, textLen int32, font IFont)                                                                                                                                                                                    // procedure
	SetRangeColor(textStart int32, textLength int32, fontColor types.TColor)                                                                                                                                                                                      // procedure
	SetRangeParamsWithIntX3TextModifyMaskStringColorFontStylesX2(textStart int32, textLength int32, modifyMask types.TTextModifyMask, fontName string, fontSize int32, fontColor types.TColor, addFontStyle types.TFontStyles, removeFontStyle types.TFontStyles) // procedure
	SetRangeParamsWithIntX2TextModifyMaskFontParamsFontStylesX2(textStart int32, textLength int32, modifyMask types.TTextModifyMask, fnt TFontParams, addFontStyle types.TFontStyles, removeFontStyle types.TFontStyles)                                          // procedure
	SetRangeParaParams(textStart int32, textLength int32, modifyMask types.TParaModifyMask, paraMetric TParaMetric)                                                                                                                                               // procedure
	SetLink(textStart int32, textLength int32, isLink bool, linkRef string)                                                                                                                                                                                       // procedure
	SetSelLengthFor(aselstr string)                                                                                                                                                                                                                               // procedure
	Redo()                                                                                                                                                                                                                                                        // procedure
	ZoomFactor() float64                                                                                                                                                                                                                                          // property ZoomFactor Getter
	SetZoomFactor(value float64)                                                                                                                                                                                                                                  // property ZoomFactor Setter
	CanRedo() bool                                                                                                                                                                                                                                                // property CanRedo Getter
	Transparent() bool                                                                                                                                                                                                                                            // property Transparent Getter
	SetTransparent(value bool)                                                                                                                                                                                                                                    // property Transparent Setter
	SetOnSelectionChange(fn TNotifyEvent)                                                                                                                                                                                                                         // property event
	SetOnPrintAction(fn TPrintActionEvent)                                                                                                                                                                                                                        // property event
	SetOnLinkAction(fn TLinkActionEvent)                                                                                                                                                                                                                          // property event
}

type TCustomRichMemo struct {
	TCustomMemo
}

func (m *TCustomRichMemo) CanPaste() bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetTextAttributes(textStart int32, textParams *TFontParams) bool {
	if !m.IsValid() {
		return false
	}
	textParamsPtr := textParams.ToPas()
	r := customRichMemoAPI().SysCallN(2, m.Instance(), uintptr(textStart), uintptr(base.UnsafePointer(textParamsPtr)))
	*textParams = textParamsPtr.ToGo()
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetStyleRange(charOfs int32, rangeStart *int32, rangeLen *int32) bool {
	if !m.IsValid() {
		return false
	}
	rangeStartPtr := uintptr(*rangeStart)
	rangeLenPtr := uintptr(*rangeLen)
	r := customRichMemoAPI().SysCallN(3, m.Instance(), uintptr(charOfs), uintptr(base.UnsafePointer(&rangeStartPtr)), uintptr(base.UnsafePointer(&rangeLenPtr)))
	*rangeStart = int32(rangeStartPtr)
	*rangeLen = int32(rangeLenPtr)
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetParaAlignmentWithIntParaAlignment(textStart int32, align *types.TParaAlignment) bool {
	if !m.IsValid() {
		return false
	}
	alignPtr := uintptr(*align)
	r := customRichMemoAPI().SysCallN(4, m.Instance(), uintptr(textStart), uintptr(base.UnsafePointer(&alignPtr)))
	*align = types.TParaAlignment(alignPtr)
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetParaAlignmentWithInt(textStart int32) types.TParaAlignment {
	if !m.IsValid() {
		return 0
	}
	r := customRichMemoAPI().SysCallN(5, m.Instance(), uintptr(textStart))
	return types.TParaAlignment(r)
}

func (m *TCustomRichMemo) GetParaMetric(textStart int32, metric *TParaMetric) bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(6, m.Instance(), uintptr(textStart), uintptr(base.UnsafePointer(metric)))
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetParaNumbering(textStart int32, number *TParaNumbering) bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(7, m.Instance(), uintptr(textStart), uintptr(base.UnsafePointer(number)))
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetParaRangeWithIntParaRange(charOfs int32, paraRange *TParaRange) bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(8, m.Instance(), uintptr(charOfs), uintptr(base.UnsafePointer(paraRange)))
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetParaRangeWithIntX3(charOfs int32, textStart *int32, textLength *int32) bool {
	if !m.IsValid() {
		return false
	}
	textStartPtr := uintptr(*textStart)
	textLengthPtr := uintptr(*textLength)
	r := customRichMemoAPI().SysCallN(9, m.Instance(), uintptr(charOfs), uintptr(base.UnsafePointer(&textStartPtr)), uintptr(base.UnsafePointer(&textLengthPtr)))
	*textStart = int32(textStartPtr)
	*textLength = int32(textLengthPtr)
	return api.GoBool(r)
}

func (m *TCustomRichMemo) GetParaTabs(charOfs int32, stopList *TTabStopList) bool {
	if !m.IsValid() {
		return false
	}
	stopListPtr := stopList.ToPas()
	r := customRichMemoAPI().SysCallN(10, m.Instance(), uintptr(charOfs), uintptr(base.UnsafePointer(stopListPtr)))
	*stopList = stopListPtr.ToGo()
	return api.GoBool(r)
}

func (m *TCustomRichMemo) IsLink(textStart int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(11, m.Instance(), uintptr(textStart))
	return api.GoBool(r)
}

func (m *TCustomRichMemo) LoadRichText(source IStream) bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(source))
	return api.GoBool(r)
}

func (m *TCustomRichMemo) SaveRichText(dest IStream) bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(dest))
	return api.GoBool(r)
}

func (m *TCustomRichMemo) InDelText(uTF8Text string, insStartChar int32, replaceLength int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customRichMemoAPI().SysCallN(14, m.Instance(), api.PasStr(uTF8Text), uintptr(insStartChar), uintptr(replaceLength))
	return int32(r)
}

func (m *TCustomRichMemo) InDelInline(inlineobj IRichMemoInline, insStartChar int32, replaceLength int32, size types.TSize) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customRichMemoAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(inlineobj), uintptr(insStartChar), uintptr(replaceLength), uintptr(base.UnsafePointer(&size)))
	return int32(r)
}

func (m *TCustomRichMemo) GetText(textStart int32, textLength int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customRichMemoAPI().SysCallN(16, m.Instance(), uintptr(textStart), uintptr(textLength))
	return api.GoStr(r)
}

func (m *TCustomRichMemo) GetUText(textStart int32, textLength int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customRichMemoAPI().SysCallN(17, m.Instance(), uintptr(textStart), uintptr(textLength))
	return api.GoStr(r)
}

func (m *TCustomRichMemo) SearchWithStringIntX2SearchOptions(niddle string, start int32, len int32, searchOpt types.TSearchOptions) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customRichMemoAPI().SysCallN(18, m.Instance(), api.PasStr(niddle), uintptr(start), uintptr(len), uintptr(searchOpt))
	return int32(r)
}

func (m *TCustomRichMemo) SearchWithStringIntX4SearchOptions(niddle string, start int32, len int32, searchOpt types.TSearchOptions, textStart *int32, textLength *int32) bool {
	if !m.IsValid() {
		return false
	}
	textStartPtr := uintptr(*textStart)
	textLengthPtr := uintptr(*textLength)
	r := customRichMemoAPI().SysCallN(19, m.Instance(), api.PasStr(niddle), uintptr(start), uintptr(len), uintptr(searchOpt), uintptr(base.UnsafePointer(&textStartPtr)), uintptr(base.UnsafePointer(&textLengthPtr)))
	*textStart = int32(textStartPtr)
	*textLength = int32(textLengthPtr)
	return api.GoBool(r)
}

func (m *TCustomRichMemo) Print(params TPrintParams) int32 {
	if !m.IsValid() {
		return 0
	}
	paramsPtr := params.ToPas()
	r := customRichMemoAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(paramsPtr)))
	return int32(r)
}

func (m *TCustomRichMemo) CharAtPos(X int32, Y int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customRichMemoAPI().SysCallN(21, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r)
}

func (m *TCustomRichMemo) SetTextAttributesWithIntX2FontParams(textStart int32, textLen int32, textParams TFontParams) {
	if !m.IsValid() {
		return
	}
	textParamsPtr := textParams.ToPas()
	customRichMemoAPI().SysCallN(22, m.Instance(), uintptr(textStart), uintptr(textLen), uintptr(base.UnsafePointer(textParamsPtr)))
}

func (m *TCustomRichMemo) SetParaAlignment(textStart int32, textLen int32, align types.TParaAlignment) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(23, m.Instance(), uintptr(textStart), uintptr(textLen), uintptr(align))
}

func (m *TCustomRichMemo) SetParaMetric(textStart int32, textLen int32, metric TParaMetric) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(24, m.Instance(), uintptr(textStart), uintptr(textLen), uintptr(base.UnsafePointer(&metric)))
}

func (m *TCustomRichMemo) SetParaNumbering(textStart int32, textLen int32, number TParaNumbering) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(25, m.Instance(), uintptr(textStart), uintptr(textLen), uintptr(base.UnsafePointer(&number)))
}

func (m *TCustomRichMemo) SetParaTabs(textStart int32, textLen int32, stopList TTabStopList) {
	if !m.IsValid() {
		return
	}
	stopListPtr := stopList.ToPas()
	customRichMemoAPI().SysCallN(26, m.Instance(), uintptr(textStart), uintptr(textLen), uintptr(base.UnsafePointer(stopListPtr)))
}

func (m *TCustomRichMemo) SetTextAttributesWithIntX2Font(textStart int32, textLen int32, font IFont) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(27, m.Instance(), uintptr(textStart), uintptr(textLen), base.GetObjectUintptr(font))
}

func (m *TCustomRichMemo) SetRangeColor(textStart int32, textLength int32, fontColor types.TColor) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(28, m.Instance(), uintptr(textStart), uintptr(textLength), uintptr(fontColor))
}

func (m *TCustomRichMemo) SetRangeParamsWithIntX3TextModifyMaskStringColorFontStylesX2(textStart int32, textLength int32, modifyMask types.TTextModifyMask, fontName string, fontSize int32, fontColor types.TColor, addFontStyle types.TFontStyles, removeFontStyle types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(29, m.Instance(), uintptr(textStart), uintptr(textLength), uintptr(modifyMask), api.PasStr(fontName), uintptr(fontSize), uintptr(fontColor), uintptr(addFontStyle), uintptr(removeFontStyle))
}

func (m *TCustomRichMemo) SetRangeParamsWithIntX2TextModifyMaskFontParamsFontStylesX2(textStart int32, textLength int32, modifyMask types.TTextModifyMask, fnt TFontParams, addFontStyle types.TFontStyles, removeFontStyle types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	fntPtr := fnt.ToPas()
	customRichMemoAPI().SysCallN(30, m.Instance(), uintptr(textStart), uintptr(textLength), uintptr(modifyMask), uintptr(base.UnsafePointer(fntPtr)), uintptr(addFontStyle), uintptr(removeFontStyle))
}

func (m *TCustomRichMemo) SetRangeParaParams(textStart int32, textLength int32, modifyMask types.TParaModifyMask, paraMetric TParaMetric) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(31, m.Instance(), uintptr(textStart), uintptr(textLength), uintptr(modifyMask), uintptr(base.UnsafePointer(&paraMetric)))
}

func (m *TCustomRichMemo) SetLink(textStart int32, textLength int32, isLink bool, linkRef string) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(32, m.Instance(), uintptr(textStart), uintptr(textLength), api.PasBool(isLink), api.PasStr(linkRef))
}

func (m *TCustomRichMemo) SetSelLengthFor(aselstr string) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(33, m.Instance(), api.PasStr(aselstr))
}

func (m *TCustomRichMemo) Redo() {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(34, m.Instance())
}

func (m *TCustomRichMemo) ZoomFactor() (result float64) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(35, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomRichMemo) SetZoomFactor(value float64) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(35, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomRichMemo) CanRedo() bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(36, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomRichMemo) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := customRichMemoAPI().SysCallN(37, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomRichMemo) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	customRichMemoAPI().SysCallN(37, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomRichMemo) SetOnSelectionChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 38, customRichMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomRichMemo) SetOnPrintAction(fn TPrintActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTPrintActionEvent(fn)
	base.SetEvent(m, 39, customRichMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomRichMemo) SetOnLinkAction(fn TLinkActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLinkActionEvent(fn)
	base.SetEvent(m, 40, customRichMemoAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomRichMemo class constructor
func NewCustomRichMemo(owner IComponent) ICustomRichMemo {
	r := customRichMemoAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomRichMemo(r)
}

func TCustomRichMemoClass() types.TClass {
	r := customRichMemoAPI().SysCallN(41)
	return types.TClass(r)
}

var (
	customRichMemoOnce   base.Once
	customRichMemoImport *imports.Imports = nil
)

func customRichMemoAPI() *imports.Imports {
	customRichMemoOnce.Do(func() {
		customRichMemoImport = api.NewDefaultImports()
		customRichMemoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomRichMemo_Create", 0), // constructor NewCustomRichMemo
			/* 1 */ imports.NewTable("TCustomRichMemo_CanPaste", 0), // function CanPaste
			/* 2 */ imports.NewTable("TCustomRichMemo_GetTextAttributes", 0), // function GetTextAttributes
			/* 3 */ imports.NewTable("TCustomRichMemo_GetStyleRange", 0), // function GetStyleRange
			/* 4 */ imports.NewTable("TCustomRichMemo_GetParaAlignmentWithIntParaAlignment", 0), // function GetParaAlignmentWithIntParaAlignment
			/* 5 */ imports.NewTable("TCustomRichMemo_GetParaAlignmentWithInt", 0), // function GetParaAlignmentWithInt
			/* 6 */ imports.NewTable("TCustomRichMemo_GetParaMetric", 0), // function GetParaMetric
			/* 7 */ imports.NewTable("TCustomRichMemo_GetParaNumbering", 0), // function GetParaNumbering
			/* 8 */ imports.NewTable("TCustomRichMemo_GetParaRangeWithIntParaRange", 0), // function GetParaRangeWithIntParaRange
			/* 9 */ imports.NewTable("TCustomRichMemo_GetParaRangeWithIntX3", 0), // function GetParaRangeWithIntX3
			/* 10 */ imports.NewTable("TCustomRichMemo_GetParaTabs", 0), // function GetParaTabs
			/* 11 */ imports.NewTable("TCustomRichMemo_isLink", 0), // function IsLink
			/* 12 */ imports.NewTable("TCustomRichMemo_LoadRichText", 0), // function LoadRichText
			/* 13 */ imports.NewTable("TCustomRichMemo_SaveRichText", 0), // function SaveRichText
			/* 14 */ imports.NewTable("TCustomRichMemo_InDelText", 0), // function InDelText
			/* 15 */ imports.NewTable("TCustomRichMemo_InDelInline", 0), // function InDelInline
			/* 16 */ imports.NewTable("TCustomRichMemo_GetText", 0), // function GetText
			/* 17 */ imports.NewTable("TCustomRichMemo_GetUText", 0), // function GetUText
			/* 18 */ imports.NewTable("TCustomRichMemo_SearchWithStringIntX2SearchOptions", 0), // function SearchWithStringIntX2SearchOptions
			/* 19 */ imports.NewTable("TCustomRichMemo_SearchWithStringIntX4SearchOptions", 0), // function SearchWithStringIntX4SearchOptions
			/* 20 */ imports.NewTable("TCustomRichMemo_Print", 0), // function Print
			/* 21 */ imports.NewTable("TCustomRichMemo_CharAtPos", 0), // function CharAtPos
			/* 22 */ imports.NewTable("TCustomRichMemo_SetTextAttributesWithIntX2FontParams", 0), // procedure SetTextAttributesWithIntX2FontParams
			/* 23 */ imports.NewTable("TCustomRichMemo_SetParaAlignment", 0), // procedure SetParaAlignment
			/* 24 */ imports.NewTable("TCustomRichMemo_SetParaMetric", 0), // procedure SetParaMetric
			/* 25 */ imports.NewTable("TCustomRichMemo_SetParaNumbering", 0), // procedure SetParaNumbering
			/* 26 */ imports.NewTable("TCustomRichMemo_SetParaTabs", 0), // procedure SetParaTabs
			/* 27 */ imports.NewTable("TCustomRichMemo_SetTextAttributesWithIntX2Font", 0), // procedure SetTextAttributesWithIntX2Font
			/* 28 */ imports.NewTable("TCustomRichMemo_SetRangeColor", 0), // procedure SetRangeColor
			/* 29 */ imports.NewTable("TCustomRichMemo_SetRangeParamsWithIntX3TextModifyMaskStringColorFontStylesX2", 0), // procedure SetRangeParamsWithIntX3TextModifyMaskStringColorFontStylesX2
			/* 30 */ imports.NewTable("TCustomRichMemo_SetRangeParamsWithIntX2TextModifyMaskFontParamsFontStylesX2", 0), // procedure SetRangeParamsWithIntX2TextModifyMaskFontParamsFontStylesX2
			/* 31 */ imports.NewTable("TCustomRichMemo_SetRangeParaParams", 0), // procedure SetRangeParaParams
			/* 32 */ imports.NewTable("TCustomRichMemo_SetLink", 0), // procedure SetLink
			/* 33 */ imports.NewTable("TCustomRichMemo_SetSelLengthFor", 0), // procedure SetSelLengthFor
			/* 34 */ imports.NewTable("TCustomRichMemo_Redo", 0), // procedure Redo
			/* 35 */ imports.NewTable("TCustomRichMemo_ZoomFactor", 0), // property ZoomFactor
			/* 36 */ imports.NewTable("TCustomRichMemo_CanRedo", 0), // property CanRedo
			/* 37 */ imports.NewTable("TCustomRichMemo_Transparent", 0), // property Transparent
			/* 38 */ imports.NewTable("TCustomRichMemo_OnSelectionChange", 0), // event OnSelectionChange
			/* 39 */ imports.NewTable("TCustomRichMemo_OnPrintAction", 0), // event OnPrintAction
			/* 40 */ imports.NewTable("TCustomRichMemo_OnLinkAction", 0), // event OnLinkAction
			/* 41 */ imports.NewTable("TCustomRichMemo_TClass", 0), // function TCustomRichMemoClass
		}
	})
	return customRichMemoImport
}
