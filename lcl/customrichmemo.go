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

// ICustomRichMemo Parent: ICustomMemo
type ICustomRichMemo interface {
	ICustomMemo
	ZoomFactor() (resultDouble float64)                                                                                                                                   // property
	SetZoomFactor(AValue float64)                                                                                                                                         // property
	CanRedo() bool                                                                                                                                                        // property
	Transparent() bool                                                                                                                                                    // property
	SetTransparent(AValue bool)                                                                                                                                           // property
	CanPaste() bool                                                                                                                                                       // function
	GetTextAttributes(TextStart int32, TextParams *TFontParams) bool                                                                                                      // function
	GetStyleRange(CharOfs int32, RangeStart, RangeLen *int32) bool                                                                                                        // function
	GetParaAlignment(TextStart int32, AAlign *TParaAlignment) bool                                                                                                        // function
	GetParaAlignment1(TextStart int32) TParaAlignment                                                                                                                     // function
	GetParaMetric(TextStart int32, AMetric *TParaMetric) bool                                                                                                             // function
	GetParaNumbering(TextStart int32, ANumber *TParaNumbering) bool                                                                                                       // function
	GetParaRange(CharOfs int32, ParaRange *TParaRange) bool                                                                                                               // function
	GetParaRange1(CharOfs int32, TextStart, TextLength *int32) bool                                                                                                       // function
	GetParaTabs(CharOfs int32, AStopList *TTabStopList) bool                                                                                                              // function
	IsLink(TextStart int32) bool                                                                                                                                          // function
	LoadRichText(Source IStream) bool                                                                                                                                     // function
	SaveRichText(Dest IStream) bool                                                                                                                                       // function
	InDelText(UTF8Text string, InsStartChar, ReplaceLength int32) int32                                                                                                   // function
	InDelInline(inlineobj IRichMemoInline, InsStartChar, ReplaceLength int32, ASize *TSize) int32                                                                         // function
	GetTextForString(TextStart, TextLength int32) string                                                                                                                  // function
	GetUText(TextStart, TextLength int32) string                                                                                                                          // function
	Search(ANiddle string, Start, Len int32, SearchOpt TSearchOptions) int32                                                                                              // function
	Search1(ANiddle string, Start, Len int32, SearchOpt TSearchOptions, ATextStart, ATextLength *int32) bool                                                              // function
	Print(params *TPrintParams) int32                                                                                                                                     // function
	CharAtPos(x, y int32) int32                                                                                                                                           // function
	SetTextAttributes(TextStart, TextLen int32, TextParams *TFontParams)                                                                                                  // procedure
	SetParaAlignment(TextStart, TextLen int32, AAlign TParaAlignment)                                                                                                     // procedure
	SetParaMetric(TextStart, TextLen int32, AMetric *TParaMetric)                                                                                                         // procedure
	SetParaNumbering(TextStart, TextLen int32, ANumber *TParaNumbering)                                                                                                   // procedure
	SetParaTabs(TextStart, TextLen int32, AStopList *TTabStopList)                                                                                                        // procedure
	SetTextAttributes1(TextStart, TextLen int32, AFont IFont)                                                                                                             // procedure
	SetRangeColor(TextStart, TextLength int32, FontColor TColor)                                                                                                          // procedure
	SetRangeParams(TextStart, TextLength int32, ModifyMask TTextModifyMask, FontName string, FontSize int32, FontColor TColor, AddFontStyle, RemoveFontStyle TFontStyles) // procedure
	SetRangeParams1(TextStart, TextLength int32, ModifyMask TTextModifyMask, fnt *TFontParams, AddFontStyle, RemoveFontStyle TFontStyles)                                 // procedure
	SetRangeParaParams(TextStart, TextLength int32, ModifyMask TParaModifyMask, ParaMetric *TParaMetric)                                                                  // procedure
	SetLink(TextStart, TextLength int32, AIsLink bool, ALinkRef string)                                                                                                   // procedure
	SetSelLengthFor(aselstr string)                                                                                                                                       // procedure
	Redo()                                                                                                                                                                // procedure
	SetOnSelectionChange(fn TNotifyEvent)                                                                                                                                 // property event
}

// TCustomRichMemo Parent: TCustomMemo
type TCustomRichMemo struct {
	TCustomMemo
	selectionChangePtr uintptr
}

func NewCustomRichMemo(AOwner IComponent) ICustomRichMemo {
	r1 := LCL().SysCallN(2196, GetObjectUintptr(AOwner))
	return AsCustomRichMemo(r1)
}

func (m *TCustomRichMemo) ZoomFactor() (resultDouble float64) {
	LCL().SysCallN(2231, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomRichMemo) SetZoomFactor(AValue float64) {
	LCL().SysCallN(2231, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCustomRichMemo) CanRedo() bool {
	r1 := LCL().SysCallN(2193, m.Instance())
	return GoBool(r1)
}

func (m *TCustomRichMemo) Transparent() bool {
	r1 := LCL().SysCallN(2230, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomRichMemo) SetTransparent(AValue bool) {
	LCL().SysCallN(2230, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomRichMemo) CanPaste() bool {
	r1 := LCL().SysCallN(2192, m.Instance())
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetTextAttributes(TextStart int32, TextParams *TFontParams) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(2205, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*TextParams = *(*TFontParams)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetStyleRange(CharOfs int32, RangeStart, RangeLen *int32) bool {
	var result1 uintptr
	var result2 uintptr
	r1 := LCL().SysCallN(2204, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*RangeStart = int32(result1)
	*RangeLen = int32(result2)
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaAlignment(TextStart int32, AAlign *TParaAlignment) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(2197, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*AAlign = TParaAlignment(result1)
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaAlignment1(TextStart int32) TParaAlignment {
	r1 := LCL().SysCallN(2198, m.Instance(), uintptr(TextStart))
	return TParaAlignment(r1)
}

func (m *TCustomRichMemo) GetParaMetric(TextStart int32, AMetric *TParaMetric) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(2199, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*AMetric = *(*TParaMetric)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaNumbering(TextStart int32, ANumber *TParaNumbering) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(2200, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*ANumber = *(*TParaNumbering)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaRange(CharOfs int32, ParaRange *TParaRange) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(2201, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)))
	*ParaRange = *(*TParaRange)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaRange1(CharOfs int32, TextStart, TextLength *int32) bool {
	var result1 uintptr
	var result2 uintptr
	r1 := LCL().SysCallN(2202, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*TextStart = int32(result1)
	*TextLength = int32(result2)
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaTabs(CharOfs int32, AStopList *TTabStopList) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(2203, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)))
	*AStopList = *(*TTabStopList)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) IsLink(TextStart int32) bool {
	r1 := LCL().SysCallN(2210, m.Instance(), uintptr(TextStart))
	return GoBool(r1)
}

func (m *TCustomRichMemo) LoadRichText(Source IStream) bool {
	r1 := LCL().SysCallN(2211, m.Instance(), GetObjectUintptr(Source))
	return GoBool(r1)
}

func (m *TCustomRichMemo) SaveRichText(Dest IStream) bool {
	r1 := LCL().SysCallN(2214, m.Instance(), GetObjectUintptr(Dest))
	return GoBool(r1)
}

func (m *TCustomRichMemo) InDelText(UTF8Text string, InsStartChar, ReplaceLength int32) int32 {
	r1 := LCL().SysCallN(2209, m.Instance(), PascalStr(UTF8Text), uintptr(InsStartChar), uintptr(ReplaceLength))
	return int32(r1)
}

func (m *TCustomRichMemo) InDelInline(inlineobj IRichMemoInline, InsStartChar, ReplaceLength int32, ASize *TSize) int32 {
	r1 := LCL().SysCallN(2208, m.Instance(), GetObjectUintptr(inlineobj), uintptr(InsStartChar), uintptr(ReplaceLength), uintptr(unsafePointer(ASize)))
	return int32(r1)
}

func (m *TCustomRichMemo) GetTextForString(TextStart, TextLength int32) string {
	r1 := LCL().SysCallN(2206, m.Instance(), uintptr(TextStart), uintptr(TextLength))
	return GoStr(r1)
}

func (m *TCustomRichMemo) GetUText(TextStart, TextLength int32) string {
	r1 := LCL().SysCallN(2207, m.Instance(), uintptr(TextStart), uintptr(TextLength))
	return GoStr(r1)
}

func (m *TCustomRichMemo) Search(ANiddle string, Start, Len int32, SearchOpt TSearchOptions) int32 {
	r1 := LCL().SysCallN(2215, m.Instance(), PascalStr(ANiddle), uintptr(Start), uintptr(Len), uintptr(SearchOpt))
	return int32(r1)
}

func (m *TCustomRichMemo) Search1(ANiddle string, Start, Len int32, SearchOpt TSearchOptions, ATextStart, ATextLength *int32) bool {
	var result3 uintptr
	var result4 uintptr
	r1 := LCL().SysCallN(2216, m.Instance(), PascalStr(ANiddle), uintptr(Start), uintptr(Len), uintptr(SearchOpt), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*ATextStart = int32(result3)
	*ATextLength = int32(result4)
	return GoBool(r1)
}

func (m *TCustomRichMemo) Print(params *TPrintParams) int32 {
	r1 := LCL().SysCallN(2212, m.Instance(), uintptr(unsafePointer(params)))
	return int32(r1)
}

func (m *TCustomRichMemo) CharAtPos(x, y int32) int32 {
	r1 := LCL().SysCallN(2194, m.Instance(), uintptr(x), uintptr(y))
	return int32(r1)
}

func CustomRichMemoClass() TClass {
	ret := LCL().SysCallN(2195)
	return TClass(ret)
}

func (m *TCustomRichMemo) SetTextAttributes(TextStart, TextLen int32, TextParams *TFontParams) {
	LCL().SysCallN(2228, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(TextParams)))
}

func (m *TCustomRichMemo) SetParaAlignment(TextStart, TextLen int32, AAlign TParaAlignment) {
	LCL().SysCallN(2219, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(AAlign))
}

func (m *TCustomRichMemo) SetParaMetric(TextStart, TextLen int32, AMetric *TParaMetric) {
	LCL().SysCallN(2220, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(AMetric)))
}

func (m *TCustomRichMemo) SetParaNumbering(TextStart, TextLen int32, ANumber *TParaNumbering) {
	LCL().SysCallN(2221, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(ANumber)))
}

func (m *TCustomRichMemo) SetParaTabs(TextStart, TextLen int32, AStopList *TTabStopList) {
	LCL().SysCallN(2222, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(AStopList)))
}

func (m *TCustomRichMemo) SetTextAttributes1(TextStart, TextLen int32, AFont IFont) {
	LCL().SysCallN(2229, m.Instance(), uintptr(TextStart), uintptr(TextLen), GetObjectUintptr(AFont))
}

func (m *TCustomRichMemo) SetRangeColor(TextStart, TextLength int32, FontColor TColor) {
	LCL().SysCallN(2223, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(FontColor))
}

func (m *TCustomRichMemo) SetRangeParams(TextStart, TextLength int32, ModifyMask TTextModifyMask, FontName string, FontSize int32, FontColor TColor, AddFontStyle, RemoveFontStyle TFontStyles) {
	LCL().SysCallN(2225, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(ModifyMask), PascalStr(FontName), uintptr(FontSize), uintptr(FontColor), uintptr(AddFontStyle), uintptr(RemoveFontStyle))
}

func (m *TCustomRichMemo) SetRangeParams1(TextStart, TextLength int32, ModifyMask TTextModifyMask, fnt *TFontParams, AddFontStyle, RemoveFontStyle TFontStyles) {
	LCL().SysCallN(2226, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(ModifyMask), uintptr(unsafePointer(fnt)), uintptr(AddFontStyle), uintptr(RemoveFontStyle))
}

func (m *TCustomRichMemo) SetRangeParaParams(TextStart, TextLength int32, ModifyMask TParaModifyMask, ParaMetric *TParaMetric) {
	LCL().SysCallN(2224, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(ModifyMask), uintptr(unsafePointer(ParaMetric)))
}

func (m *TCustomRichMemo) SetLink(TextStart, TextLength int32, AIsLink bool, ALinkRef string) {
	LCL().SysCallN(2217, m.Instance(), uintptr(TextStart), uintptr(TextLength), PascalBool(AIsLink), PascalStr(ALinkRef))
}

func (m *TCustomRichMemo) SetSelLengthFor(aselstr string) {
	LCL().SysCallN(2227, m.Instance(), PascalStr(aselstr))
}

func (m *TCustomRichMemo) Redo() {
	LCL().SysCallN(2213, m.Instance())
}

func (m *TCustomRichMemo) SetOnSelectionChange(fn TNotifyEvent) {
	if m.selectionChangePtr != 0 {
		RemoveEventElement(m.selectionChangePtr)
	}
	m.selectionChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2218, m.Instance(), m.selectionChangePtr)
}
