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

// IStrings Is Abstract Class Parent: IPersistent
type IStrings interface {
	IPersistent
	AddStrings2(sArr []string)
	AddStrings3(list IStrings, clearFirst bool)
	AddPair2(name, value string, object IObject) IStrings
	LoadFromBytes(data []byte)
	AlwaysQuote() bool                                                          // property
	SetAlwaysQuote(AValue bool)                                                 // property
	Capacity() int32                                                            // property
	SetCapacity(AValue int32)                                                   // property
	CommaText() string                                                          // property
	SetCommaText(AValue string)                                                 // property
	Count() int32                                                               // property
	DelimitedText() string                                                      // property
	SetDelimitedText(AValue string)                                             // property
	Delimiter() Char                                                            // property
	SetDelimiter(AValue Char)                                                   // property
	LineBreak() string                                                          // property
	SetLineBreak(AValue string)                                                 // property
	MissingNameValueSeparatorAction() TMissingNameValueSeparatorAction          // property
	SetMissingNameValueSeparatorAction(AValue TMissingNameValueSeparatorAction) // property
	Names(Index int32) string                                                   // property
	NameValueSeparator() Char                                                   // property
	SetNameValueSeparator(AValue Char)                                          // property
	Objects(Index int32) IObject                                                // property
	SetObjects(Index int32, AValue IObject)                                     // property
	Options() TStringsOptions                                                   // property
	SetOptions(AValue TStringsOptions)                                          // property
	QuoteChar() Char                                                            // property
	SetQuoteChar(AValue Char)                                                   // property
	SkipLastLineBreak() bool                                                    // property
	SetSkipLastLineBreak(AValue bool)                                           // property
	TrailingLineBreak() bool                                                    // property
	SetTrailingLineBreak(AValue bool)                                           // property
	StrictDelimiter() bool                                                      // property
	SetStrictDelimiter(AValue bool)                                             // property
	Strings(Index int32) string                                                 // property
	SetStrings(Index int32, AValue string)                                      // property
	Text() string                                                               // property
	SetText(AValue string)                                                      // property
	TextLineBreakStyle() TTextLineBreakStyle                                    // property
	SetTextLineBreakStyle(AValue TTextLineBreakStyle)                           // property
	UseLocale() bool                                                            // property
	SetUseLocale(AValue bool)                                                   // property
	ValueFromIndex(Index int32) string                                          // property
	SetValueFromIndex(Index int32, AValue string)                               // property
	Values(Name string) string                                                  // property
	SetValues(Name string, AValue string)                                       // property
	WriteBOM() bool                                                             // property
	SetWriteBOM(AValue bool)                                                    // property
	ToObjectArray(aStart, aEnd int32) TObjectDynArray                           // function
	ToObjectArray1() TObjectDynArray                                            // function
	ToStringArray(aStart, aEnd int32) TStringDynArray                           // function
	ToStringArray1() TStringDynArray                                            // function
	Add(S string) int32                                                         // function
	AddObject(S string, AObject IObject) int32                                  // function
	AddPair(AName, AValue string) IStrings                                      // function
	AddPair1(AName, AValue string, AObject IObject) IStrings                    // function
	EqualsForBool(TheStrings IStrings) bool                                     // function
	ExtractName(S string) string                                                // function
	GetEnumerator() IStringsEnumerator                                          // function
	GetText1() string                                                           // function
	IndexOf(S string) int32                                                     // function
	IndexOf1(S string, aStart int32) int32                                      // function
	IndexOfName(Name string) int32                                              // function
	IndexOfObject(AObject IObject) int32                                        // function
	LastIndexOf(S string, aStart int32) int32                                   // function
	LastIndexOf1(S string) int32                                                // function
	Pop() string                                                                // function
	Reverse() IStrings                                                          // function
	Shift() string                                                              // function
	Slice(fromIndex int32) IStrings                                             // function
	AddStrings(TheStrings IStrings)                                             // procedure
	AddStrings1(TheStrings IStrings, ClearFirst bool)                           // procedure
	SetStrings1(TheStrings IStrings)                                            // procedure
	AddText(S string)                                                           // procedure
	AddCommaText(S string)                                                      // procedure
	AddDelimitedText(S string, ADelimiter Char, AStrictDelimiter bool)          // procedure
	AddDelimitedtext1(S string)                                                 // procedure
	Append(S string)                                                            // procedure
	BeginUpdate()                                                               // procedure
	Clear()                                                                     // procedure Is Abstract
	Delete(Index int32)                                                         // procedure Is Abstract
	EndUpdate()                                                                 // procedure
	Exchange(Index1, Index2 int32)                                              // procedure
	Fill(aValue string, aStart, aEnd int32)                                     // procedure
	GetNameValue(Index int32, OutName, OutValue *string)                        // procedure
	Insert(Index int32, S string)                                               // procedure Is Abstract
	InsertObject(Index int32, S string, AObject IObject)                        // procedure
	LoadFromFile(FileName string)                                               // procedure
	LoadFromFile1(FileName string, IgnoreEncoding bool)                         // procedure
	LoadFromStream(Stream IStream)                                              // procedure
	LoadFromStream1(Stream IStream, IgnoreEncoding bool)                        // procedure
	Move(CurIndex, NewIndex int32)                                              // procedure
	Reverse1(aList IStrings)                                                    // procedure
	SaveToFile(FileName string)                                                 // procedure
	SaveToFile1(FileName string, IgnoreEncoding bool)                           // procedure
	SaveToStream(Stream IStream)                                                // procedure
	SaveToStream1(Stream IStream, IgnoreEncoding bool)                          // procedure
	Slice1(fromIndex int32, aList IStrings)                                     // procedure
	SetText1(TheText string)                                                    // procedure
}

// TStrings Is Abstract Class Parent: TPersistent
type TStrings struct {
	TPersistent
}

func (m *TStrings) AlwaysQuote() bool {
	r1 := stringsImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetAlwaysQuote(AValue bool) {
	stringsImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) Capacity() int32 {
	r1 := stringsImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TStrings) SetCapacity(AValue int32) {
	stringsImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) CommaText() string {
	r1 := stringsImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetCommaText(AValue string) {
	stringsImportAPI().SysCallN(16, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) Count() int32 {
	r1 := stringsImportAPI().SysCallN(17, m.Instance())
	return int32(r1)
}

func (m *TStrings) DelimitedText() string {
	r1 := stringsImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetDelimitedText(AValue string) {
	stringsImportAPI().SysCallN(19, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) Delimiter() Char {
	r1 := stringsImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TStrings) SetDelimiter(AValue Char) {
	stringsImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) LineBreak() string {
	r1 := stringsImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetLineBreak(AValue string) {
	stringsImportAPI().SysCallN(37, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) MissingNameValueSeparatorAction() TMissingNameValueSeparatorAction {
	r1 := stringsImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return TMissingNameValueSeparatorAction(r1)
}

func (m *TStrings) SetMissingNameValueSeparatorAction(AValue TMissingNameValueSeparatorAction) {
	stringsImportAPI().SysCallN(42, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) Names(Index int32) string {
	r1 := stringsImportAPI().SysCallN(45, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TStrings) NameValueSeparator() Char {
	r1 := stringsImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TStrings) SetNameValueSeparator(AValue Char) {
	stringsImportAPI().SysCallN(44, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) Objects(Index int32) IObject {
	r1 := stringsImportAPI().SysCallN(46, 0, m.Instance(), uintptr(Index))
	return AsObject(r1)
}

func (m *TStrings) SetObjects(Index int32, AValue IObject) {
	stringsImportAPI().SysCallN(46, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TStrings) Options() TStringsOptions {
	r1 := stringsImportAPI().SysCallN(47, 0, m.Instance(), 0)
	return TStringsOptions(r1)
}

func (m *TStrings) SetOptions(AValue TStringsOptions) {
	stringsImportAPI().SysCallN(47, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) QuoteChar() Char {
	r1 := stringsImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TStrings) SetQuoteChar(AValue Char) {
	stringsImportAPI().SysCallN(49, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) SkipLastLineBreak() bool {
	r1 := stringsImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetSkipLastLineBreak(AValue bool) {
	stringsImportAPI().SysCallN(59, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) TrailingLineBreak() bool {
	r1 := stringsImportAPI().SysCallN(70, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetTrailingLineBreak(AValue bool) {
	stringsImportAPI().SysCallN(70, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) StrictDelimiter() bool {
	r1 := stringsImportAPI().SysCallN(62, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetStrictDelimiter(AValue bool) {
	stringsImportAPI().SysCallN(62, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) Strings(Index int32) string {
	r1 := stringsImportAPI().SysCallN(63, 0, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TStrings) SetStrings(Index int32, AValue string) {
	stringsImportAPI().SysCallN(63, 1, m.Instance(), uintptr(Index), PascalStr(AValue))
}

func (m *TStrings) Text() string {
	r1 := stringsImportAPI().SysCallN(64, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetText(AValue string) {
	stringsImportAPI().SysCallN(64, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) TextLineBreakStyle() TTextLineBreakStyle {
	r1 := stringsImportAPI().SysCallN(65, 0, m.Instance(), 0)
	return TTextLineBreakStyle(r1)
}

func (m *TStrings) SetTextLineBreakStyle(AValue TTextLineBreakStyle) {
	stringsImportAPI().SysCallN(65, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) UseLocale() bool {
	r1 := stringsImportAPI().SysCallN(71, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetUseLocale(AValue bool) {
	stringsImportAPI().SysCallN(71, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) ValueFromIndex(Index int32) string {
	r1 := stringsImportAPI().SysCallN(72, 0, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TStrings) SetValueFromIndex(Index int32, AValue string) {
	stringsImportAPI().SysCallN(72, 1, m.Instance(), uintptr(Index), PascalStr(AValue))
}

func (m *TStrings) Values(Name string) string {
	r1 := stringsImportAPI().SysCallN(73, 0, m.Instance(), PascalStr(Name))
	return GoStr(r1)
}

func (m *TStrings) SetValues(Name string, AValue string) {
	stringsImportAPI().SysCallN(73, 1, m.Instance(), PascalStr(Name), PascalStr(AValue))
}

func (m *TStrings) WriteBOM() bool {
	r1 := stringsImportAPI().SysCallN(74, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetWriteBOM(AValue bool) {
	stringsImportAPI().SysCallN(74, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) ToObjectArray(aStart, aEnd int32) TObjectDynArray {
	r1 := stringsImportAPI().SysCallN(66, m.Instance(), uintptr(aStart), uintptr(aEnd))
	return TObjectDynArray(r1)
}

func (m *TStrings) ToObjectArray1() TObjectDynArray {
	r1 := stringsImportAPI().SysCallN(67, m.Instance())
	return TObjectDynArray(r1)
}

func (m *TStrings) ToStringArray(aStart, aEnd int32) TStringDynArray {
	r1 := stringsImportAPI().SysCallN(68, m.Instance(), uintptr(aStart), uintptr(aEnd))
	return TStringDynArray(r1)
}

func (m *TStrings) ToStringArray1() TStringDynArray {
	r1 := stringsImportAPI().SysCallN(69, m.Instance())
	return TStringDynArray(r1)
}

func (m *TStrings) Add(S string) int32 {
	r1 := stringsImportAPI().SysCallN(0, m.Instance(), PascalStr(S))
	return int32(r1)
}

func (m *TStrings) AddObject(S string, AObject IObject) int32 {
	r1 := stringsImportAPI().SysCallN(4, m.Instance(), PascalStr(S), GetObjectUintptr(AObject))
	return int32(r1)
}

func (m *TStrings) AddPair(AName, AValue string) IStrings {
	r1 := stringsImportAPI().SysCallN(5, m.Instance(), PascalStr(AName), PascalStr(AValue))
	return AsStrings(r1)
}

func (m *TStrings) AddPair1(AName, AValue string, AObject IObject) IStrings {
	r1 := stringsImportAPI().SysCallN(6, m.Instance(), PascalStr(AName), PascalStr(AValue), GetObjectUintptr(AObject))
	return AsStrings(r1)
}

func (m *TStrings) EqualsForBool(TheStrings IStrings) bool {
	r1 := stringsImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(TheStrings))
	return GoBool(r1)
}

func (m *TStrings) ExtractName(S string) string {
	r1 := stringsImportAPI().SysCallN(24, m.Instance(), PascalStr(S))
	return GoStr(r1)
}

func (m *TStrings) GetEnumerator() IStringsEnumerator {
	r1 := stringsImportAPI().SysCallN(26, m.Instance())
	return AsStringsEnumerator(r1)
}

func (m *TStrings) GetText1() string {
	r1 := stringsImportAPI().SysCallN(28, m.Instance())
	return GoStr(r1)
}

func (m *TStrings) IndexOf(S string) int32 {
	r1 := stringsImportAPI().SysCallN(29, m.Instance(), PascalStr(S))
	return int32(r1)
}

func (m *TStrings) IndexOf1(S string, aStart int32) int32 {
	r1 := stringsImportAPI().SysCallN(30, m.Instance(), PascalStr(S), uintptr(aStart))
	return int32(r1)
}

func (m *TStrings) IndexOfName(Name string) int32 {
	r1 := stringsImportAPI().SysCallN(31, m.Instance(), PascalStr(Name))
	return int32(r1)
}

func (m *TStrings) IndexOfObject(AObject IObject) int32 {
	r1 := stringsImportAPI().SysCallN(32, m.Instance(), GetObjectUintptr(AObject))
	return int32(r1)
}

func (m *TStrings) LastIndexOf(S string, aStart int32) int32 {
	r1 := stringsImportAPI().SysCallN(35, m.Instance(), PascalStr(S), uintptr(aStart))
	return int32(r1)
}

func (m *TStrings) LastIndexOf1(S string) int32 {
	r1 := stringsImportAPI().SysCallN(36, m.Instance(), PascalStr(S))
	return int32(r1)
}

func (m *TStrings) Pop() string {
	r1 := stringsImportAPI().SysCallN(48, m.Instance())
	return GoStr(r1)
}

func (m *TStrings) Reverse() IStrings {
	r1 := stringsImportAPI().SysCallN(50, m.Instance())
	return AsStrings(r1)
}

func (m *TStrings) Shift() string {
	r1 := stringsImportAPI().SysCallN(58, m.Instance())
	return GoStr(r1)
}

func (m *TStrings) Slice(fromIndex int32) IStrings {
	r1 := stringsImportAPI().SysCallN(60, m.Instance(), uintptr(fromIndex))
	return AsStrings(r1)
}

func StringsClass() TClass {
	ret := stringsImportAPI().SysCallN(14)
	return TClass(ret)
}

func (m *TStrings) AddStrings(TheStrings IStrings) {
	stringsImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(TheStrings))
}

func (m *TStrings) AddStrings1(TheStrings IStrings, ClearFirst bool) {
	stringsImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(TheStrings), PascalBool(ClearFirst))
}

func (m *TStrings) SetStrings1(TheStrings IStrings) {
	stringsImportAPI().SysCallN(56, m.Instance(), GetObjectUintptr(TheStrings))
}

func (m *TStrings) AddText(S string) {
	stringsImportAPI().SysCallN(9, m.Instance(), PascalStr(S))
}

func (m *TStrings) AddCommaText(S string) {
	stringsImportAPI().SysCallN(1, m.Instance(), PascalStr(S))
}

func (m *TStrings) AddDelimitedText(S string, ADelimiter Char, AStrictDelimiter bool) {
	stringsImportAPI().SysCallN(2, m.Instance(), PascalStr(S), uintptr(ADelimiter), PascalBool(AStrictDelimiter))
}

func (m *TStrings) AddDelimitedtext1(S string) {
	stringsImportAPI().SysCallN(3, m.Instance(), PascalStr(S))
}

func (m *TStrings) Append(S string) {
	stringsImportAPI().SysCallN(11, m.Instance(), PascalStr(S))
}

func (m *TStrings) BeginUpdate() {
	stringsImportAPI().SysCallN(12, m.Instance())
}

func (m *TStrings) Clear() {
	stringsImportAPI().SysCallN(15, m.Instance())
}

func (m *TStrings) Delete(Index int32) {
	stringsImportAPI().SysCallN(18, m.Instance(), uintptr(Index))
}

func (m *TStrings) EndUpdate() {
	stringsImportAPI().SysCallN(21, m.Instance())
}

func (m *TStrings) Exchange(Index1, Index2 int32) {
	stringsImportAPI().SysCallN(23, m.Instance(), uintptr(Index1), uintptr(Index2))
}

func (m *TStrings) Fill(aValue string, aStart, aEnd int32) {
	stringsImportAPI().SysCallN(25, m.Instance(), PascalStr(aValue), uintptr(aStart), uintptr(aEnd))
}

func (m *TStrings) GetNameValue(Index int32, OutName, OutValue *string) {
	var result1 uintptr
	var result2 uintptr
	stringsImportAPI().SysCallN(27, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutName = GoStr(result1)
	*OutValue = GoStr(result2)
}

func (m *TStrings) Insert(Index int32, S string) {
	stringsImportAPI().SysCallN(33, m.Instance(), uintptr(Index), PascalStr(S))
}

func (m *TStrings) InsertObject(Index int32, S string, AObject IObject) {
	stringsImportAPI().SysCallN(34, m.Instance(), uintptr(Index), PascalStr(S), GetObjectUintptr(AObject))
}

func (m *TStrings) LoadFromFile(FileName string) {
	stringsImportAPI().SysCallN(38, m.Instance(), PascalStr(FileName))
}

func (m *TStrings) LoadFromFile1(FileName string, IgnoreEncoding bool) {
	stringsImportAPI().SysCallN(39, m.Instance(), PascalStr(FileName), PascalBool(IgnoreEncoding))
}

func (m *TStrings) LoadFromStream(Stream IStream) {
	stringsImportAPI().SysCallN(40, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TStrings) LoadFromStream1(Stream IStream, IgnoreEncoding bool) {
	stringsImportAPI().SysCallN(41, m.Instance(), GetObjectUintptr(Stream), PascalBool(IgnoreEncoding))
}

func (m *TStrings) Move(CurIndex, NewIndex int32) {
	stringsImportAPI().SysCallN(43, m.Instance(), uintptr(CurIndex), uintptr(NewIndex))
}

func (m *TStrings) Reverse1(aList IStrings) {
	stringsImportAPI().SysCallN(51, m.Instance(), GetObjectUintptr(aList))
}

func (m *TStrings) SaveToFile(FileName string) {
	stringsImportAPI().SysCallN(52, m.Instance(), PascalStr(FileName))
}

func (m *TStrings) SaveToFile1(FileName string, IgnoreEncoding bool) {
	stringsImportAPI().SysCallN(53, m.Instance(), PascalStr(FileName), PascalBool(IgnoreEncoding))
}

func (m *TStrings) SaveToStream(Stream IStream) {
	stringsImportAPI().SysCallN(54, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TStrings) SaveToStream1(Stream IStream, IgnoreEncoding bool) {
	stringsImportAPI().SysCallN(55, m.Instance(), GetObjectUintptr(Stream), PascalBool(IgnoreEncoding))
}

func (m *TStrings) Slice1(fromIndex int32, aList IStrings) {
	stringsImportAPI().SysCallN(61, m.Instance(), uintptr(fromIndex), GetObjectUintptr(aList))
}

func (m *TStrings) SetText1(TheText string) {
	stringsImportAPI().SysCallN(57, m.Instance(), PascalStr(TheText))
}

var (
	stringsImport       *imports.Imports = nil
	stringsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Strings_Add", 0),
		/*1*/ imports.NewTable("Strings_AddCommaText", 0),
		/*2*/ imports.NewTable("Strings_AddDelimitedText", 0),
		/*3*/ imports.NewTable("Strings_AddDelimitedtext1", 0),
		/*4*/ imports.NewTable("Strings_AddObject", 0),
		/*5*/ imports.NewTable("Strings_AddPair", 0),
		/*6*/ imports.NewTable("Strings_AddPair1", 0),
		/*7*/ imports.NewTable("Strings_AddStrings", 0),
		/*8*/ imports.NewTable("Strings_AddStrings1", 0),
		/*9*/ imports.NewTable("Strings_AddText", 0),
		/*10*/ imports.NewTable("Strings_AlwaysQuote", 0),
		/*11*/ imports.NewTable("Strings_Append", 0),
		/*12*/ imports.NewTable("Strings_BeginUpdate", 0),
		/*13*/ imports.NewTable("Strings_Capacity", 0),
		/*14*/ imports.NewTable("Strings_Class", 0),
		/*15*/ imports.NewTable("Strings_Clear", 0),
		/*16*/ imports.NewTable("Strings_CommaText", 0),
		/*17*/ imports.NewTable("Strings_Count", 0),
		/*18*/ imports.NewTable("Strings_Delete", 0),
		/*19*/ imports.NewTable("Strings_DelimitedText", 0),
		/*20*/ imports.NewTable("Strings_Delimiter", 0),
		/*21*/ imports.NewTable("Strings_EndUpdate", 0),
		/*22*/ imports.NewTable("Strings_EqualsForBool", 0),
		/*23*/ imports.NewTable("Strings_Exchange", 0),
		/*24*/ imports.NewTable("Strings_ExtractName", 0),
		/*25*/ imports.NewTable("Strings_Fill", 0),
		/*26*/ imports.NewTable("Strings_GetEnumerator", 0),
		/*27*/ imports.NewTable("Strings_GetNameValue", 0),
		/*28*/ imports.NewTable("Strings_GetText1", 0),
		/*29*/ imports.NewTable("Strings_IndexOf", 0),
		/*30*/ imports.NewTable("Strings_IndexOf1", 0),
		/*31*/ imports.NewTable("Strings_IndexOfName", 0),
		/*32*/ imports.NewTable("Strings_IndexOfObject", 0),
		/*33*/ imports.NewTable("Strings_Insert", 0),
		/*34*/ imports.NewTable("Strings_InsertObject", 0),
		/*35*/ imports.NewTable("Strings_LastIndexOf", 0),
		/*36*/ imports.NewTable("Strings_LastIndexOf1", 0),
		/*37*/ imports.NewTable("Strings_LineBreak", 0),
		/*38*/ imports.NewTable("Strings_LoadFromFile", 0),
		/*39*/ imports.NewTable("Strings_LoadFromFile1", 0),
		/*40*/ imports.NewTable("Strings_LoadFromStream", 0),
		/*41*/ imports.NewTable("Strings_LoadFromStream1", 0),
		/*42*/ imports.NewTable("Strings_MissingNameValueSeparatorAction", 0),
		/*43*/ imports.NewTable("Strings_Move", 0),
		/*44*/ imports.NewTable("Strings_NameValueSeparator", 0),
		/*45*/ imports.NewTable("Strings_Names", 0),
		/*46*/ imports.NewTable("Strings_Objects", 0),
		/*47*/ imports.NewTable("Strings_Options", 0),
		/*48*/ imports.NewTable("Strings_Pop", 0),
		/*49*/ imports.NewTable("Strings_QuoteChar", 0),
		/*50*/ imports.NewTable("Strings_Reverse", 0),
		/*51*/ imports.NewTable("Strings_Reverse1", 0),
		/*52*/ imports.NewTable("Strings_SaveToFile", 0),
		/*53*/ imports.NewTable("Strings_SaveToFile1", 0),
		/*54*/ imports.NewTable("Strings_SaveToStream", 0),
		/*55*/ imports.NewTable("Strings_SaveToStream1", 0),
		/*56*/ imports.NewTable("Strings_SetStrings1", 0),
		/*57*/ imports.NewTable("Strings_SetText1", 0),
		/*58*/ imports.NewTable("Strings_Shift", 0),
		/*59*/ imports.NewTable("Strings_SkipLastLineBreak", 0),
		/*60*/ imports.NewTable("Strings_Slice", 0),
		/*61*/ imports.NewTable("Strings_Slice1", 0),
		/*62*/ imports.NewTable("Strings_StrictDelimiter", 0),
		/*63*/ imports.NewTable("Strings_Strings", 0),
		/*64*/ imports.NewTable("Strings_Text", 0),
		/*65*/ imports.NewTable("Strings_TextLineBreakStyle", 0),
		/*66*/ imports.NewTable("Strings_ToObjectArray", 0),
		/*67*/ imports.NewTable("Strings_ToObjectArray1", 0),
		/*68*/ imports.NewTable("Strings_ToStringArray", 0),
		/*69*/ imports.NewTable("Strings_ToStringArray1", 0),
		/*70*/ imports.NewTable("Strings_TrailingLineBreak", 0),
		/*71*/ imports.NewTable("Strings_UseLocale", 0),
		/*72*/ imports.NewTable("Strings_ValueFromIndex", 0),
		/*73*/ imports.NewTable("Strings_Values", 0),
		/*74*/ imports.NewTable("Strings_WriteBOM", 0),
	}
)

func stringsImportAPI() *imports.Imports {
	if stringsImport == nil {
		stringsImport = NewDefaultImports()
		stringsImport.SetImportTable(stringsImportTables)
		stringsImportTables = nil
	}
	return stringsImport
}
