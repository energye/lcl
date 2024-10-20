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
	r1 := customRichMemoImportAPI().SysCallN(4, GetObjectUintptr(AOwner))
	return AsCustomRichMemo(r1)
}

func (m *TCustomRichMemo) ZoomFactor() (resultDouble float64) {
	customRichMemoImportAPI().SysCallN(39, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomRichMemo) SetZoomFactor(AValue float64) {
	customRichMemoImportAPI().SysCallN(39, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCustomRichMemo) CanRedo() bool {
	r1 := customRichMemoImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCustomRichMemo) Transparent() bool {
	r1 := customRichMemoImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomRichMemo) SetTransparent(AValue bool) {
	customRichMemoImportAPI().SysCallN(38, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomRichMemo) CanPaste() bool {
	r1 := customRichMemoImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetTextAttributes(TextStart int32, TextParams *TFontParams) bool {
	var result1 uintptr
	r1 := customRichMemoImportAPI().SysCallN(13, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*TextParams = *(*TFontParams)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetStyleRange(CharOfs int32, RangeStart, RangeLen *int32) bool {
	var result1 uintptr
	var result2 uintptr
	r1 := customRichMemoImportAPI().SysCallN(12, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*RangeStart = int32(result1)
	*RangeLen = int32(result2)
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaAlignment(TextStart int32, AAlign *TParaAlignment) bool {
	var result1 uintptr
	r1 := customRichMemoImportAPI().SysCallN(5, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*AAlign = TParaAlignment(result1)
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaAlignment1(TextStart int32) TParaAlignment {
	r1 := customRichMemoImportAPI().SysCallN(6, m.Instance(), uintptr(TextStart))
	return TParaAlignment(r1)
}

func (m *TCustomRichMemo) GetParaMetric(TextStart int32, AMetric *TParaMetric) bool {
	var result1 uintptr
	r1 := customRichMemoImportAPI().SysCallN(7, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*AMetric = *(*TParaMetric)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaNumbering(TextStart int32, ANumber *TParaNumbering) bool {
	var result1 uintptr
	r1 := customRichMemoImportAPI().SysCallN(8, m.Instance(), uintptr(TextStart), uintptr(unsafePointer(&result1)))
	*ANumber = *(*TParaNumbering)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaRange(CharOfs int32, ParaRange *TParaRange) bool {
	var result1 uintptr
	r1 := customRichMemoImportAPI().SysCallN(9, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)))
	*ParaRange = *(*TParaRange)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaRange1(CharOfs int32, TextStart, TextLength *int32) bool {
	var result1 uintptr
	var result2 uintptr
	r1 := customRichMemoImportAPI().SysCallN(10, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*TextStart = int32(result1)
	*TextLength = int32(result2)
	return GoBool(r1)
}

func (m *TCustomRichMemo) GetParaTabs(CharOfs int32, AStopList *TTabStopList) bool {
	var result1 uintptr
	r1 := customRichMemoImportAPI().SysCallN(11, m.Instance(), uintptr(CharOfs), uintptr(unsafePointer(&result1)))
	*AStopList = *(*TTabStopList)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomRichMemo) IsLink(TextStart int32) bool {
	r1 := customRichMemoImportAPI().SysCallN(18, m.Instance(), uintptr(TextStart))
	return GoBool(r1)
}

func (m *TCustomRichMemo) LoadRichText(Source IStream) bool {
	r1 := customRichMemoImportAPI().SysCallN(19, m.Instance(), GetObjectUintptr(Source))
	return GoBool(r1)
}

func (m *TCustomRichMemo) SaveRichText(Dest IStream) bool {
	r1 := customRichMemoImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(Dest))
	return GoBool(r1)
}

func (m *TCustomRichMemo) InDelText(UTF8Text string, InsStartChar, ReplaceLength int32) int32 {
	r1 := customRichMemoImportAPI().SysCallN(17, m.Instance(), PascalStr(UTF8Text), uintptr(InsStartChar), uintptr(ReplaceLength))
	return int32(r1)
}

func (m *TCustomRichMemo) InDelInline(inlineobj IRichMemoInline, InsStartChar, ReplaceLength int32, ASize *TSize) int32 {
	r1 := customRichMemoImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(inlineobj), uintptr(InsStartChar), uintptr(ReplaceLength), uintptr(unsafePointer(ASize)))
	return int32(r1)
}

func (m *TCustomRichMemo) GetTextForString(TextStart, TextLength int32) string {
	r1 := customRichMemoImportAPI().SysCallN(14, m.Instance(), uintptr(TextStart), uintptr(TextLength))
	return GoStr(r1)
}

func (m *TCustomRichMemo) GetUText(TextStart, TextLength int32) string {
	r1 := customRichMemoImportAPI().SysCallN(15, m.Instance(), uintptr(TextStart), uintptr(TextLength))
	return GoStr(r1)
}

func (m *TCustomRichMemo) Search(ANiddle string, Start, Len int32, SearchOpt TSearchOptions) int32 {
	r1 := customRichMemoImportAPI().SysCallN(23, m.Instance(), PascalStr(ANiddle), uintptr(Start), uintptr(Len), uintptr(SearchOpt))
	return int32(r1)
}

func (m *TCustomRichMemo) Search1(ANiddle string, Start, Len int32, SearchOpt TSearchOptions, ATextStart, ATextLength *int32) bool {
	var result3 uintptr
	var result4 uintptr
	r1 := customRichMemoImportAPI().SysCallN(24, m.Instance(), PascalStr(ANiddle), uintptr(Start), uintptr(Len), uintptr(SearchOpt), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*ATextStart = int32(result3)
	*ATextLength = int32(result4)
	return GoBool(r1)
}

func (m *TCustomRichMemo) Print(params *TPrintParams) int32 {
	r1 := customRichMemoImportAPI().SysCallN(20, m.Instance(), uintptr(unsafePointer(params)))
	return int32(r1)
}

func (m *TCustomRichMemo) CharAtPos(x, y int32) int32 {
	r1 := customRichMemoImportAPI().SysCallN(2, m.Instance(), uintptr(x), uintptr(y))
	return int32(r1)
}

func CustomRichMemoClass() TClass {
	ret := customRichMemoImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TCustomRichMemo) SetTextAttributes(TextStart, TextLen int32, TextParams *TFontParams) {
	customRichMemoImportAPI().SysCallN(36, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(TextParams)))
}

func (m *TCustomRichMemo) SetParaAlignment(TextStart, TextLen int32, AAlign TParaAlignment) {
	customRichMemoImportAPI().SysCallN(27, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(AAlign))
}

func (m *TCustomRichMemo) SetParaMetric(TextStart, TextLen int32, AMetric *TParaMetric) {
	customRichMemoImportAPI().SysCallN(28, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(AMetric)))
}

func (m *TCustomRichMemo) SetParaNumbering(TextStart, TextLen int32, ANumber *TParaNumbering) {
	customRichMemoImportAPI().SysCallN(29, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(ANumber)))
}

func (m *TCustomRichMemo) SetParaTabs(TextStart, TextLen int32, AStopList *TTabStopList) {
	customRichMemoImportAPI().SysCallN(30, m.Instance(), uintptr(TextStart), uintptr(TextLen), uintptr(unsafePointer(AStopList)))
}

func (m *TCustomRichMemo) SetTextAttributes1(TextStart, TextLen int32, AFont IFont) {
	customRichMemoImportAPI().SysCallN(37, m.Instance(), uintptr(TextStart), uintptr(TextLen), GetObjectUintptr(AFont))
}

func (m *TCustomRichMemo) SetRangeColor(TextStart, TextLength int32, FontColor TColor) {
	customRichMemoImportAPI().SysCallN(31, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(FontColor))
}

func (m *TCustomRichMemo) SetRangeParams(TextStart, TextLength int32, ModifyMask TTextModifyMask, FontName string, FontSize int32, FontColor TColor, AddFontStyle, RemoveFontStyle TFontStyles) {
	customRichMemoImportAPI().SysCallN(33, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(ModifyMask), PascalStr(FontName), uintptr(FontSize), uintptr(FontColor), uintptr(AddFontStyle), uintptr(RemoveFontStyle))
}

func (m *TCustomRichMemo) SetRangeParams1(TextStart, TextLength int32, ModifyMask TTextModifyMask, fnt *TFontParams, AddFontStyle, RemoveFontStyle TFontStyles) {
	customRichMemoImportAPI().SysCallN(34, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(ModifyMask), uintptr(unsafePointer(fnt)), uintptr(AddFontStyle), uintptr(RemoveFontStyle))
}

func (m *TCustomRichMemo) SetRangeParaParams(TextStart, TextLength int32, ModifyMask TParaModifyMask, ParaMetric *TParaMetric) {
	customRichMemoImportAPI().SysCallN(32, m.Instance(), uintptr(TextStart), uintptr(TextLength), uintptr(ModifyMask), uintptr(unsafePointer(ParaMetric)))
}

func (m *TCustomRichMemo) SetLink(TextStart, TextLength int32, AIsLink bool, ALinkRef string) {
	customRichMemoImportAPI().SysCallN(25, m.Instance(), uintptr(TextStart), uintptr(TextLength), PascalBool(AIsLink), PascalStr(ALinkRef))
}

func (m *TCustomRichMemo) SetSelLengthFor(aselstr string) {
	customRichMemoImportAPI().SysCallN(35, m.Instance(), PascalStr(aselstr))
}

func (m *TCustomRichMemo) Redo() {
	customRichMemoImportAPI().SysCallN(21, m.Instance())
}

func (m *TCustomRichMemo) SetOnSelectionChange(fn TNotifyEvent) {
	if m.selectionChangePtr != 0 {
		RemoveEventElement(m.selectionChangePtr)
	}
	m.selectionChangePtr = MakeEventDataPtr(fn)
	customRichMemoImportAPI().SysCallN(26, m.Instance(), m.selectionChangePtr)
}

var (
	customRichMemoImport       *imports.Imports = nil
	customRichMemoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRichMemo_CanPaste", 0),
		/*1*/ imports.NewTable("CustomRichMemo_CanRedo", 0),
		/*2*/ imports.NewTable("CustomRichMemo_CharAtPos", 0),
		/*3*/ imports.NewTable("CustomRichMemo_Class", 0),
		/*4*/ imports.NewTable("CustomRichMemo_Create", 0),
		/*5*/ imports.NewTable("CustomRichMemo_GetParaAlignment", 0),
		/*6*/ imports.NewTable("CustomRichMemo_GetParaAlignment1", 0),
		/*7*/ imports.NewTable("CustomRichMemo_GetParaMetric", 0),
		/*8*/ imports.NewTable("CustomRichMemo_GetParaNumbering", 0),
		/*9*/ imports.NewTable("CustomRichMemo_GetParaRange", 0),
		/*10*/ imports.NewTable("CustomRichMemo_GetParaRange1", 0),
		/*11*/ imports.NewTable("CustomRichMemo_GetParaTabs", 0),
		/*12*/ imports.NewTable("CustomRichMemo_GetStyleRange", 0),
		/*13*/ imports.NewTable("CustomRichMemo_GetTextAttributes", 0),
		/*14*/ imports.NewTable("CustomRichMemo_GetTextForString", 0),
		/*15*/ imports.NewTable("CustomRichMemo_GetUText", 0),
		/*16*/ imports.NewTable("CustomRichMemo_InDelInline", 0),
		/*17*/ imports.NewTable("CustomRichMemo_InDelText", 0),
		/*18*/ imports.NewTable("CustomRichMemo_IsLink", 0),
		/*19*/ imports.NewTable("CustomRichMemo_LoadRichText", 0),
		/*20*/ imports.NewTable("CustomRichMemo_Print", 0),
		/*21*/ imports.NewTable("CustomRichMemo_Redo", 0),
		/*22*/ imports.NewTable("CustomRichMemo_SaveRichText", 0),
		/*23*/ imports.NewTable("CustomRichMemo_Search", 0),
		/*24*/ imports.NewTable("CustomRichMemo_Search1", 0),
		/*25*/ imports.NewTable("CustomRichMemo_SetLink", 0),
		/*26*/ imports.NewTable("CustomRichMemo_SetOnSelectionChange", 0),
		/*27*/ imports.NewTable("CustomRichMemo_SetParaAlignment", 0),
		/*28*/ imports.NewTable("CustomRichMemo_SetParaMetric", 0),
		/*29*/ imports.NewTable("CustomRichMemo_SetParaNumbering", 0),
		/*30*/ imports.NewTable("CustomRichMemo_SetParaTabs", 0),
		/*31*/ imports.NewTable("CustomRichMemo_SetRangeColor", 0),
		/*32*/ imports.NewTable("CustomRichMemo_SetRangeParaParams", 0),
		/*33*/ imports.NewTable("CustomRichMemo_SetRangeParams", 0),
		/*34*/ imports.NewTable("CustomRichMemo_SetRangeParams1", 0),
		/*35*/ imports.NewTable("CustomRichMemo_SetSelLengthFor", 0),
		/*36*/ imports.NewTable("CustomRichMemo_SetTextAttributes", 0),
		/*37*/ imports.NewTable("CustomRichMemo_SetTextAttributes1", 0),
		/*38*/ imports.NewTable("CustomRichMemo_Transparent", 0),
		/*39*/ imports.NewTable("CustomRichMemo_ZoomFactor", 0),
	}
)

func customRichMemoImportAPI() *imports.Imports {
	if customRichMemoImport == nil {
		customRichMemoImport = NewDefaultImports()
		customRichMemoImport.SetImportTable(customRichMemoImportTables)
		customRichMemoImportTables = nil
	}
	return customRichMemoImport
}
