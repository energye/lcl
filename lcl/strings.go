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
	r1 := LCL().SysCallN(5303, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetAlwaysQuote(AValue bool) {
	LCL().SysCallN(5303, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) Capacity() int32 {
	r1 := LCL().SysCallN(5306, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TStrings) SetCapacity(AValue int32) {
	LCL().SysCallN(5306, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) CommaText() string {
	r1 := LCL().SysCallN(5309, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetCommaText(AValue string) {
	LCL().SysCallN(5309, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) Count() int32 {
	r1 := LCL().SysCallN(5310, m.Instance())
	return int32(r1)
}

func (m *TStrings) DelimitedText() string {
	r1 := LCL().SysCallN(5312, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetDelimitedText(AValue string) {
	LCL().SysCallN(5312, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) Delimiter() Char {
	r1 := LCL().SysCallN(5313, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TStrings) SetDelimiter(AValue Char) {
	LCL().SysCallN(5313, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) LineBreak() string {
	r1 := LCL().SysCallN(5330, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetLineBreak(AValue string) {
	LCL().SysCallN(5330, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) MissingNameValueSeparatorAction() TMissingNameValueSeparatorAction {
	r1 := LCL().SysCallN(5335, 0, m.Instance(), 0)
	return TMissingNameValueSeparatorAction(r1)
}

func (m *TStrings) SetMissingNameValueSeparatorAction(AValue TMissingNameValueSeparatorAction) {
	LCL().SysCallN(5335, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) Names(Index int32) string {
	r1 := LCL().SysCallN(5338, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TStrings) NameValueSeparator() Char {
	r1 := LCL().SysCallN(5337, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TStrings) SetNameValueSeparator(AValue Char) {
	LCL().SysCallN(5337, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) Objects(Index int32) IObject {
	r1 := LCL().SysCallN(5339, 0, m.Instance(), uintptr(Index))
	return AsObject(r1)
}

func (m *TStrings) SetObjects(Index int32, AValue IObject) {
	LCL().SysCallN(5339, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TStrings) Options() TStringsOptions {
	r1 := LCL().SysCallN(5340, 0, m.Instance(), 0)
	return TStringsOptions(r1)
}

func (m *TStrings) SetOptions(AValue TStringsOptions) {
	LCL().SysCallN(5340, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) QuoteChar() Char {
	r1 := LCL().SysCallN(5342, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TStrings) SetQuoteChar(AValue Char) {
	LCL().SysCallN(5342, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) SkipLastLineBreak() bool {
	r1 := LCL().SysCallN(5352, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetSkipLastLineBreak(AValue bool) {
	LCL().SysCallN(5352, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) TrailingLineBreak() bool {
	r1 := LCL().SysCallN(5363, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetTrailingLineBreak(AValue bool) {
	LCL().SysCallN(5363, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) StrictDelimiter() bool {
	r1 := LCL().SysCallN(5355, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetStrictDelimiter(AValue bool) {
	LCL().SysCallN(5355, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) Strings(Index int32) string {
	r1 := LCL().SysCallN(5356, 0, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TStrings) SetStrings(Index int32, AValue string) {
	LCL().SysCallN(5356, 1, m.Instance(), uintptr(Index), PascalStr(AValue))
}

func (m *TStrings) Text() string {
	r1 := LCL().SysCallN(5357, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStrings) SetText(AValue string) {
	LCL().SysCallN(5357, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStrings) TextLineBreakStyle() TTextLineBreakStyle {
	r1 := LCL().SysCallN(5358, 0, m.Instance(), 0)
	return TTextLineBreakStyle(r1)
}

func (m *TStrings) SetTextLineBreakStyle(AValue TTextLineBreakStyle) {
	LCL().SysCallN(5358, 1, m.Instance(), uintptr(AValue))
}

func (m *TStrings) UseLocale() bool {
	r1 := LCL().SysCallN(5364, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetUseLocale(AValue bool) {
	LCL().SysCallN(5364, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) ValueFromIndex(Index int32) string {
	r1 := LCL().SysCallN(5365, 0, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TStrings) SetValueFromIndex(Index int32, AValue string) {
	LCL().SysCallN(5365, 1, m.Instance(), uintptr(Index), PascalStr(AValue))
}

func (m *TStrings) Values(Name string) string {
	r1 := LCL().SysCallN(5366, 0, m.Instance(), PascalStr(Name))
	return GoStr(r1)
}

func (m *TStrings) SetValues(Name string, AValue string) {
	LCL().SysCallN(5366, 1, m.Instance(), PascalStr(Name), PascalStr(AValue))
}

func (m *TStrings) WriteBOM() bool {
	r1 := LCL().SysCallN(5367, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStrings) SetWriteBOM(AValue bool) {
	LCL().SysCallN(5367, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStrings) ToObjectArray(aStart, aEnd int32) TObjectDynArray {
	r1 := LCL().SysCallN(5359, m.Instance(), uintptr(aStart), uintptr(aEnd))
	return TObjectDynArray(r1)
}

func (m *TStrings) ToObjectArray1() TObjectDynArray {
	r1 := LCL().SysCallN(5360, m.Instance())
	return TObjectDynArray(r1)
}

func (m *TStrings) ToStringArray(aStart, aEnd int32) TStringDynArray {
	r1 := LCL().SysCallN(5361, m.Instance(), uintptr(aStart), uintptr(aEnd))
	return TStringDynArray(r1)
}

func (m *TStrings) ToStringArray1() TStringDynArray {
	r1 := LCL().SysCallN(5362, m.Instance())
	return TStringDynArray(r1)
}

func (m *TStrings) Add(S string) int32 {
	r1 := LCL().SysCallN(5293, m.Instance(), PascalStr(S))
	return int32(r1)
}

func (m *TStrings) AddObject(S string, AObject IObject) int32 {
	r1 := LCL().SysCallN(5297, m.Instance(), PascalStr(S), GetObjectUintptr(AObject))
	return int32(r1)
}

func (m *TStrings) AddPair(AName, AValue string) IStrings {
	r1 := LCL().SysCallN(5298, m.Instance(), PascalStr(AName), PascalStr(AValue))
	return AsStrings(r1)
}

func (m *TStrings) AddPair1(AName, AValue string, AObject IObject) IStrings {
	r1 := LCL().SysCallN(5299, m.Instance(), PascalStr(AName), PascalStr(AValue), GetObjectUintptr(AObject))
	return AsStrings(r1)
}

func (m *TStrings) EqualsForBool(TheStrings IStrings) bool {
	r1 := LCL().SysCallN(5315, m.Instance(), GetObjectUintptr(TheStrings))
	return GoBool(r1)
}

func (m *TStrings) ExtractName(S string) string {
	r1 := LCL().SysCallN(5317, m.Instance(), PascalStr(S))
	return GoStr(r1)
}

func (m *TStrings) GetEnumerator() IStringsEnumerator {
	r1 := LCL().SysCallN(5319, m.Instance())
	return AsStringsEnumerator(r1)
}

func (m *TStrings) GetText1() string {
	r1 := LCL().SysCallN(5321, m.Instance())
	return GoStr(r1)
}

func (m *TStrings) IndexOf(S string) int32 {
	r1 := LCL().SysCallN(5322, m.Instance(), PascalStr(S))
	return int32(r1)
}

func (m *TStrings) IndexOf1(S string, aStart int32) int32 {
	r1 := LCL().SysCallN(5323, m.Instance(), PascalStr(S), uintptr(aStart))
	return int32(r1)
}

func (m *TStrings) IndexOfName(Name string) int32 {
	r1 := LCL().SysCallN(5324, m.Instance(), PascalStr(Name))
	return int32(r1)
}

func (m *TStrings) IndexOfObject(AObject IObject) int32 {
	r1 := LCL().SysCallN(5325, m.Instance(), GetObjectUintptr(AObject))
	return int32(r1)
}

func (m *TStrings) LastIndexOf(S string, aStart int32) int32 {
	r1 := LCL().SysCallN(5328, m.Instance(), PascalStr(S), uintptr(aStart))
	return int32(r1)
}

func (m *TStrings) LastIndexOf1(S string) int32 {
	r1 := LCL().SysCallN(5329, m.Instance(), PascalStr(S))
	return int32(r1)
}

func (m *TStrings) Pop() string {
	r1 := LCL().SysCallN(5341, m.Instance())
	return GoStr(r1)
}

func (m *TStrings) Reverse() IStrings {
	r1 := LCL().SysCallN(5343, m.Instance())
	return AsStrings(r1)
}

func (m *TStrings) Shift() string {
	r1 := LCL().SysCallN(5351, m.Instance())
	return GoStr(r1)
}

func (m *TStrings) Slice(fromIndex int32) IStrings {
	r1 := LCL().SysCallN(5353, m.Instance(), uintptr(fromIndex))
	return AsStrings(r1)
}

func StringsClass() TClass {
	ret := LCL().SysCallN(5307)
	return TClass(ret)
}

func (m *TStrings) AddStrings(TheStrings IStrings) {
	LCL().SysCallN(5300, m.Instance(), GetObjectUintptr(TheStrings))
}

func (m *TStrings) AddStrings1(TheStrings IStrings, ClearFirst bool) {
	LCL().SysCallN(5301, m.Instance(), GetObjectUintptr(TheStrings), PascalBool(ClearFirst))
}

func (m *TStrings) SetStrings1(TheStrings IStrings) {
	LCL().SysCallN(5349, m.Instance(), GetObjectUintptr(TheStrings))
}

func (m *TStrings) AddText(S string) {
	LCL().SysCallN(5302, m.Instance(), PascalStr(S))
}

func (m *TStrings) AddCommaText(S string) {
	LCL().SysCallN(5294, m.Instance(), PascalStr(S))
}

func (m *TStrings) AddDelimitedText(S string, ADelimiter Char, AStrictDelimiter bool) {
	LCL().SysCallN(5295, m.Instance(), PascalStr(S), uintptr(ADelimiter), PascalBool(AStrictDelimiter))
}

func (m *TStrings) AddDelimitedtext1(S string) {
	LCL().SysCallN(5296, m.Instance(), PascalStr(S))
}

func (m *TStrings) Append(S string) {
	LCL().SysCallN(5304, m.Instance(), PascalStr(S))
}

func (m *TStrings) BeginUpdate() {
	LCL().SysCallN(5305, m.Instance())
}

func (m *TStrings) Clear() {
	LCL().SysCallN(5308, m.Instance())
}

func (m *TStrings) Delete(Index int32) {
	LCL().SysCallN(5311, m.Instance(), uintptr(Index))
}

func (m *TStrings) EndUpdate() {
	LCL().SysCallN(5314, m.Instance())
}

func (m *TStrings) Exchange(Index1, Index2 int32) {
	LCL().SysCallN(5316, m.Instance(), uintptr(Index1), uintptr(Index2))
}

func (m *TStrings) Fill(aValue string, aStart, aEnd int32) {
	LCL().SysCallN(5318, m.Instance(), PascalStr(aValue), uintptr(aStart), uintptr(aEnd))
}

func (m *TStrings) GetNameValue(Index int32, OutName, OutValue *string) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(5320, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutName = GoStr(result1)
	*OutValue = GoStr(result2)
}

func (m *TStrings) Insert(Index int32, S string) {
	LCL().SysCallN(5326, m.Instance(), uintptr(Index), PascalStr(S))
}

func (m *TStrings) InsertObject(Index int32, S string, AObject IObject) {
	LCL().SysCallN(5327, m.Instance(), uintptr(Index), PascalStr(S), GetObjectUintptr(AObject))
}

func (m *TStrings) LoadFromFile(FileName string) {
	LCL().SysCallN(5331, m.Instance(), PascalStr(FileName))
}

func (m *TStrings) LoadFromFile1(FileName string, IgnoreEncoding bool) {
	LCL().SysCallN(5332, m.Instance(), PascalStr(FileName), PascalBool(IgnoreEncoding))
}

func (m *TStrings) LoadFromStream(Stream IStream) {
	LCL().SysCallN(5333, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TStrings) LoadFromStream1(Stream IStream, IgnoreEncoding bool) {
	LCL().SysCallN(5334, m.Instance(), GetObjectUintptr(Stream), PascalBool(IgnoreEncoding))
}

func (m *TStrings) Move(CurIndex, NewIndex int32) {
	LCL().SysCallN(5336, m.Instance(), uintptr(CurIndex), uintptr(NewIndex))
}

func (m *TStrings) Reverse1(aList IStrings) {
	LCL().SysCallN(5344, m.Instance(), GetObjectUintptr(aList))
}

func (m *TStrings) SaveToFile(FileName string) {
	LCL().SysCallN(5345, m.Instance(), PascalStr(FileName))
}

func (m *TStrings) SaveToFile1(FileName string, IgnoreEncoding bool) {
	LCL().SysCallN(5346, m.Instance(), PascalStr(FileName), PascalBool(IgnoreEncoding))
}

func (m *TStrings) SaveToStream(Stream IStream) {
	LCL().SysCallN(5347, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TStrings) SaveToStream1(Stream IStream, IgnoreEncoding bool) {
	LCL().SysCallN(5348, m.Instance(), GetObjectUintptr(Stream), PascalBool(IgnoreEncoding))
}

func (m *TStrings) Slice1(fromIndex int32, aList IStrings) {
	LCL().SysCallN(5354, m.Instance(), uintptr(fromIndex), GetObjectUintptr(aList))
}

func (m *TStrings) SetText1(TheText string) {
	LCL().SysCallN(5350, m.Instance(), PascalStr(TheText))
}
