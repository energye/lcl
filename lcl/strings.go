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

// IStrings Parent: IPersistent
type IStrings interface {
	IPersistent
	ToObjectArrayWithIntX2(start int32, end int32) types.TObjectDynArray                 // function
	ToObjectArrayToObjectDynArray() types.TObjectDynArray                                // function
	ToStringArrayWithIntX2(start int32, end int32) IStringArrayWrap                      // function
	ToStringArrayToStringDynArray() IStringArrayWrap                                     // function
	Add(S string) int32                                                                  // function
	AddObject(S string, object IObject) int32                                            // function
	AddPairWithStringX2(name string, value string) IStrings                              // function
	AddPairWithStringX2Object(name string, value string, object IObject) IStrings        // function
	EqualsWithStrings(theStrings IStrings) bool                                          // function
	ExtractName(S string) string                                                         // function
	GetEnumerator() IStringsEnumerator                                                   // function
	GetText() uintptr                                                                    // function
	IndexOfWithString(S string) int32                                                    // function
	IndexOfWithStringInt(S string, start int32) int32                                    // function
	IndexOfName(name string) int32                                                       // function
	IndexOfObject(object IObject) int32                                                  // function
	LastIndexOfWithStringInt(S string, start int32) int32                                // function
	LastIndexOfWithString(S string) int32                                                // function
	Pop() string                                                                         // function
	ReverseToStrings() IStrings                                                          // function
	Shift() string                                                                       // function
	SliceWithInt(fromIndex int32) IStrings                                               // function
	AddStringsWithStrings(theStrings IStrings)                                           // procedure
	AddStringsWithStringsBool(theStrings IStrings, clearFirst bool)                      // procedure
	SetStringsWithStrings(theStrings IStrings)                                           // procedure
	AddText(S string)                                                                    // procedure
	AddCommaText(S string)                                                               // procedure
	AddDelimitedTextWithStringCharBool(S string, delimiter uint16, strictDelimiter bool) // procedure
	AddDelimitedtextWithString(S string)                                                 // procedure
	Append(S string)                                                                     // procedure
	BeginUpdate()                                                                        // procedure
	Clear()                                                                              // procedure
	Delete(index int32)                                                                  // procedure
	EndUpdate()                                                                          // procedure
	Exchange(index1 int32, index2 int32)                                                 // procedure
	Fill(value string, start int32, end int32)                                           // procedure
	GetNameValue(index int32, outName *string, outValue *string)                         // procedure
	Insert(index int32, S string)                                                        // procedure
	InsertObject(index int32, S string, object IObject)                                  // procedure
	LoadFromFileWithString(fileName string)                                              // procedure
	LoadFromFileWithStringBool(fileName string, ignoreEncoding bool)                     // procedure
	LoadFromStreamWithStream(stream IStream)                                             // procedure
	LoadFromStreamWithStreamBool(stream IStream, ignoreEncoding bool)                    // procedure
	Move(curIndex int32, newIndex int32)                                                 // procedure
	ReverseWithStrings(list IStrings)                                                    // procedure
	SaveToFileWithString(fileName string)                                                // procedure
	SaveToFileWithStringBool(fileName string, ignoreEncoding bool)                       // procedure
	SaveToStreamWithStream(stream IStream)                                               // procedure
	SaveToStreamWithStreamBool(stream IStream, ignoreEncoding bool)                      // procedure
	SliceWithIntStrings(fromIndex int32, list IStrings)                                  // procedure
	SetTextWithString(theText uintptr)                                                   // procedure
	AlwaysQuote() bool                                                                   // property AlwaysQuote Getter
	SetAlwaysQuote(value bool)                                                           // property AlwaysQuote Setter
	Capacity() int32                                                                     // property Capacity Getter
	SetCapacity(value int32)                                                             // property Capacity Setter
	CommaText() string                                                                   // property CommaText Getter
	SetCommaText(value string)                                                           // property CommaText Setter
	Count() int32                                                                        // property Count Getter
	DelimitedText() string                                                               // property DelimitedText Getter
	SetDelimitedText(value string)                                                       // property DelimitedText Setter
	Delimiter() uint16                                                                   // property Delimiter Getter
	SetDelimiter(value uint16)                                                           // property Delimiter Setter
	LineBreak() string                                                                   // property LineBreak Getter
	SetLineBreak(value string)                                                           // property LineBreak Setter
	MissingNameValueSeparatorAction() types.TMissingNameValueSeparatorAction             // property MissingNameValueSeparatorAction Getter
	SetMissingNameValueSeparatorAction(value types.TMissingNameValueSeparatorAction)     // property MissingNameValueSeparatorAction Setter
	Names(index int32) string                                                            // property Names Getter
	NameValueSeparator() uint16                                                          // property NameValueSeparator Getter
	SetNameValueSeparator(value uint16)                                                  // property NameValueSeparator Setter
	Objects(index int32) IObject                                                         // property Objects Getter
	SetObjects(index int32, value IObject)                                               // property Objects Setter
	Options() types.TStringsOptions                                                      // property Options Getter
	SetOptions(value types.TStringsOptions)                                              // property Options Setter
	QuoteChar() uint16                                                                   // property QuoteChar Getter
	SetQuoteChar(value uint16)                                                           // property QuoteChar Setter
	SkipLastLineBreak() bool                                                             // property SkipLastLineBreak Getter
	SetSkipLastLineBreak(value bool)                                                     // property SkipLastLineBreak Setter
	// TrailingLineBreak
	//  Same as SkipLastLineBreak but for Delphi compatibility. Note it has opposite meaning.
	TrailingLineBreak() bool                               // property TrailingLineBreak Getter
	SetTrailingLineBreak(value bool)                       // property TrailingLineBreak Setter
	StrictDelimiter() bool                                 // property StrictDelimiter Getter
	SetStrictDelimiter(value bool)                         // property StrictDelimiter Setter
	Strings(index int32) string                            // property Strings Getter
	SetStringsWithIntToString(index int32, value string)   // property Strings Setter
	Text() string                                          // property Text Getter
	SetTextToString(value string)                          // property Text Setter
	TextLineBreakStyle() types.TTextLineBreakStyle         // property TextLineBreakStyle Getter
	SetTextLineBreakStyle(value types.TTextLineBreakStyle) // property TextLineBreakStyle Setter
	UseLocale() bool                                       // property UseLocale Getter
	SetUseLocale(value bool)                               // property UseLocale Setter
	ValueFromIndex(index int32) string                     // property ValueFromIndex Getter
	SetValueFromIndex(index int32, value string)           // property ValueFromIndex Setter
	Values(name string) string                             // property Values Getter
	SetValues(name string, value string)                   // property Values Setter
	WriteBOM() bool                                        // property WriteBOM Getter
	SetWriteBOM(value bool)                                // property WriteBOM Setter
}

type TStrings struct {
	TPersistent
}

func (m *TStrings) ToObjectArrayWithIntX2(start int32, end int32) types.TObjectDynArray {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(0, m.Instance(), uintptr(start), uintptr(end))
	return types.TObjectDynArray(r)
}

func (m *TStrings) ToObjectArrayToObjectDynArray() types.TObjectDynArray {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(1, m.Instance())
	return types.TObjectDynArray(r)
}

func (m *TStrings) ToStringArrayWithIntX2(start int32, end int32) IStringArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(2, m.Instance(), uintptr(start), uintptr(end))
	return AsStringArrayWrap(r)
}

func (m *TStrings) ToStringArrayToStringDynArray() IStringArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(3, m.Instance())
	return AsStringArrayWrap(r)
}

func (m *TStrings) Add(S string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(4, m.Instance(), api.PasStr(S))
	return int32(r)
}

func (m *TStrings) AddObject(S string, object IObject) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(5, m.Instance(), api.PasStr(S), base.GetObjectUintptr(object))
	return int32(r)
}

func (m *TStrings) AddPairWithStringX2(name string, value string) IStrings {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(6, m.Instance(), api.PasStr(name), api.PasStr(value))
	return AsStrings(r)
}

func (m *TStrings) AddPairWithStringX2Object(name string, value string, object IObject) IStrings {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(7, m.Instance(), api.PasStr(name), api.PasStr(value), base.GetObjectUintptr(object))
	return AsStrings(r)
}

func (m *TStrings) EqualsWithStrings(theStrings IStrings) bool {
	if !m.IsValid() {
		return false
	}
	r := stringsAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(theStrings))
	return api.GoBool(r)
}

func (m *TStrings) ExtractName(S string) string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(9, m.Instance(), api.PasStr(S))
	return api.GoStr(r)
}

func (m *TStrings) GetEnumerator() IStringsEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(10, m.Instance())
	return AsStringsEnumerator(r)
}

func (m *TStrings) GetText() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(11, m.Instance())
	return uintptr(r)
}

func (m *TStrings) IndexOfWithString(S string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(12, m.Instance(), api.PasStr(S))
	return int32(r)
}

func (m *TStrings) IndexOfWithStringInt(S string, start int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(13, m.Instance(), api.PasStr(S), uintptr(start))
	return int32(r)
}

func (m *TStrings) IndexOfName(name string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(14, m.Instance(), api.PasStr(name))
	return int32(r)
}

func (m *TStrings) IndexOfObject(object IObject) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(object))
	return int32(r)
}

func (m *TStrings) LastIndexOfWithStringInt(S string, start int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(16, m.Instance(), api.PasStr(S), uintptr(start))
	return int32(r)
}

func (m *TStrings) LastIndexOfWithString(S string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(17, m.Instance(), api.PasStr(S))
	return int32(r)
}

func (m *TStrings) Pop() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(18, m.Instance())
	return api.GoStr(r)
}

func (m *TStrings) ReverseToStrings() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(19, m.Instance())
	return AsStrings(r)
}

func (m *TStrings) Shift() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(20, m.Instance())
	return api.GoStr(r)
}

func (m *TStrings) SliceWithInt(fromIndex int32) IStrings {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(21, m.Instance(), uintptr(fromIndex))
	return AsStrings(r)
}

func (m *TStrings) AddStringsWithStrings(theStrings IStrings) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(theStrings))
}

func (m *TStrings) AddStringsWithStringsBool(theStrings IStrings, clearFirst bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(23, m.Instance(), base.GetObjectUintptr(theStrings), api.PasBool(clearFirst))
}

func (m *TStrings) SetStringsWithStrings(theStrings IStrings) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(24, m.Instance(), base.GetObjectUintptr(theStrings))
}

func (m *TStrings) AddText(S string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(25, m.Instance(), api.PasStr(S))
}

func (m *TStrings) AddCommaText(S string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(26, m.Instance(), api.PasStr(S))
}

func (m *TStrings) AddDelimitedTextWithStringCharBool(S string, delimiter uint16, strictDelimiter bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(27, m.Instance(), api.PasStr(S), uintptr(delimiter), api.PasBool(strictDelimiter))
}

func (m *TStrings) AddDelimitedtextWithString(S string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(28, m.Instance(), api.PasStr(S))
}

func (m *TStrings) Append(S string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(29, m.Instance(), api.PasStr(S))
}

func (m *TStrings) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(30, m.Instance())
}

func (m *TStrings) Clear() {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(31, m.Instance())
}

func (m *TStrings) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(32, m.Instance(), uintptr(index))
}

func (m *TStrings) EndUpdate() {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(33, m.Instance())
}

func (m *TStrings) Exchange(index1 int32, index2 int32) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(34, m.Instance(), uintptr(index1), uintptr(index2))
}

func (m *TStrings) Fill(value string, start int32, end int32) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(35, m.Instance(), api.PasStr(value), uintptr(start), uintptr(end))
}

func (m *TStrings) GetNameValue(index int32, outName *string, outValue *string) {
	if !m.IsValid() {
		return
	}
	var namePtr uintptr
	var valuePtr uintptr
	stringsAPI().SysCallN(36, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&namePtr)), uintptr(base.UnsafePointer(&valuePtr)))
	*outName = api.GoStr(namePtr)
	*outValue = api.GoStr(valuePtr)
}

func (m *TStrings) Insert(index int32, S string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(37, m.Instance(), uintptr(index), api.PasStr(S))
}

func (m *TStrings) InsertObject(index int32, S string, object IObject) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(38, m.Instance(), uintptr(index), api.PasStr(S), base.GetObjectUintptr(object))
}

func (m *TStrings) LoadFromFileWithString(fileName string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(39, m.Instance(), api.PasStr(fileName))
}

func (m *TStrings) LoadFromFileWithStringBool(fileName string, ignoreEncoding bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(40, m.Instance(), api.PasStr(fileName), api.PasBool(ignoreEncoding))
}

func (m *TStrings) LoadFromStreamWithStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(41, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TStrings) LoadFromStreamWithStreamBool(stream IStream, ignoreEncoding bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(42, m.Instance(), base.GetObjectUintptr(stream), api.PasBool(ignoreEncoding))
}

func (m *TStrings) Move(curIndex int32, newIndex int32) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(43, m.Instance(), uintptr(curIndex), uintptr(newIndex))
}

func (m *TStrings) ReverseWithStrings(list IStrings) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(44, m.Instance(), base.GetObjectUintptr(list))
}

func (m *TStrings) SaveToFileWithString(fileName string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(45, m.Instance(), api.PasStr(fileName))
}

func (m *TStrings) SaveToFileWithStringBool(fileName string, ignoreEncoding bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(46, m.Instance(), api.PasStr(fileName), api.PasBool(ignoreEncoding))
}

func (m *TStrings) SaveToStreamWithStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(47, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TStrings) SaveToStreamWithStreamBool(stream IStream, ignoreEncoding bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(48, m.Instance(), base.GetObjectUintptr(stream), api.PasBool(ignoreEncoding))
}

func (m *TStrings) SliceWithIntStrings(fromIndex int32, list IStrings) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(49, m.Instance(), uintptr(fromIndex), base.GetObjectUintptr(list))
}

func (m *TStrings) SetTextWithString(theText uintptr) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(50, m.Instance(), uintptr(theText))
}

func (m *TStrings) AlwaysQuote() bool {
	if !m.IsValid() {
		return false
	}
	r := stringsAPI().SysCallN(51, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStrings) SetAlwaysQuote(value bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(51, 1, m.Instance(), api.PasBool(value))
}

func (m *TStrings) Capacity() int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(52, 0, m.Instance())
	return int32(r)
}

func (m *TStrings) SetCapacity(value int32) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(52, 1, m.Instance(), uintptr(value))
}

func (m *TStrings) CommaText() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(53, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TStrings) SetCommaText(value string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(53, 1, m.Instance(), api.PasStr(value))
}

func (m *TStrings) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(54, m.Instance())
	return int32(r)
}

func (m *TStrings) DelimitedText() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(55, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TStrings) SetDelimitedText(value string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(55, 1, m.Instance(), api.PasStr(value))
}

func (m *TStrings) Delimiter() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(56, 0, m.Instance())
	return uint16(r)
}

func (m *TStrings) SetDelimiter(value uint16) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(56, 1, m.Instance(), uintptr(value))
}

func (m *TStrings) LineBreak() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(57, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TStrings) SetLineBreak(value string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(57, 1, m.Instance(), api.PasStr(value))
}

func (m *TStrings) MissingNameValueSeparatorAction() types.TMissingNameValueSeparatorAction {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(58, 0, m.Instance())
	return types.TMissingNameValueSeparatorAction(r)
}

func (m *TStrings) SetMissingNameValueSeparatorAction(value types.TMissingNameValueSeparatorAction) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(58, 1, m.Instance(), uintptr(value))
}

func (m *TStrings) Names(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(59, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TStrings) NameValueSeparator() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(60, 0, m.Instance())
	return uint16(r)
}

func (m *TStrings) SetNameValueSeparator(value uint16) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(60, 1, m.Instance(), uintptr(value))
}

func (m *TStrings) Objects(index int32) IObject {
	if !m.IsValid() {
		return nil
	}
	r := stringsAPI().SysCallN(61, 0, m.Instance(), uintptr(index))
	return AsObject(r)
}

func (m *TStrings) SetObjects(index int32, value IObject) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(61, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TStrings) Options() types.TStringsOptions {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(62, 0, m.Instance())
	return types.TStringsOptions(r)
}

func (m *TStrings) SetOptions(value types.TStringsOptions) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(62, 1, m.Instance(), uintptr(value))
}

func (m *TStrings) QuoteChar() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(63, 0, m.Instance())
	return uint16(r)
}

func (m *TStrings) SetQuoteChar(value uint16) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(63, 1, m.Instance(), uintptr(value))
}

func (m *TStrings) SkipLastLineBreak() bool {
	if !m.IsValid() {
		return false
	}
	r := stringsAPI().SysCallN(64, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStrings) SetSkipLastLineBreak(value bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(64, 1, m.Instance(), api.PasBool(value))
}

func (m *TStrings) TrailingLineBreak() bool {
	if !m.IsValid() {
		return false
	}
	r := stringsAPI().SysCallN(65, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStrings) SetTrailingLineBreak(value bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(65, 1, m.Instance(), api.PasBool(value))
}

func (m *TStrings) StrictDelimiter() bool {
	if !m.IsValid() {
		return false
	}
	r := stringsAPI().SysCallN(66, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStrings) SetStrictDelimiter(value bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(66, 1, m.Instance(), api.PasBool(value))
}

func (m *TStrings) Strings(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(67, 0, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TStrings) SetStringsWithIntToString(index int32, value string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(67, 1, m.Instance(), uintptr(index), api.PasStr(value))
}

func (m *TStrings) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(68, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TStrings) SetTextToString(value string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(68, 1, m.Instance(), api.PasStr(value))
}

func (m *TStrings) TextLineBreakStyle() types.TTextLineBreakStyle {
	if !m.IsValid() {
		return 0
	}
	r := stringsAPI().SysCallN(69, 0, m.Instance())
	return types.TTextLineBreakStyle(r)
}

func (m *TStrings) SetTextLineBreakStyle(value types.TTextLineBreakStyle) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(69, 1, m.Instance(), uintptr(value))
}

func (m *TStrings) UseLocale() bool {
	if !m.IsValid() {
		return false
	}
	r := stringsAPI().SysCallN(70, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStrings) SetUseLocale(value bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(70, 1, m.Instance(), api.PasBool(value))
}

func (m *TStrings) ValueFromIndex(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(71, 0, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TStrings) SetValueFromIndex(index int32, value string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(71, 1, m.Instance(), uintptr(index), api.PasStr(value))
}

func (m *TStrings) Values(name string) string {
	if !m.IsValid() {
		return ""
	}
	r := stringsAPI().SysCallN(72, 0, m.Instance(), api.PasStr(name))
	return api.GoStr(r)
}

func (m *TStrings) SetValues(name string, value string) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(72, 1, m.Instance(), api.PasStr(name), api.PasStr(value))
}

func (m *TStrings) WriteBOM() bool {
	if !m.IsValid() {
		return false
	}
	r := stringsAPI().SysCallN(73, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStrings) SetWriteBOM(value bool) {
	if !m.IsValid() {
		return
	}
	stringsAPI().SysCallN(73, 1, m.Instance(), api.PasBool(value))
}

var (
	stringsOnce   base.Once
	stringsImport *imports.Imports = nil
)

func stringsAPI() *imports.Imports {
	stringsOnce.Do(func() {
		stringsImport = api.NewDefaultImports()
		stringsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStrings_ToObjectArrayWithIntX2", 0), // function ToObjectArrayWithIntX2
			/* 1 */ imports.NewTable("TStrings_ToObjectArrayToObjectDynArray", 0), // function ToObjectArrayToObjectDynArray
			/* 2 */ imports.NewTable("TStrings_ToStringArrayWithIntX2", 0), // function ToStringArrayWithIntX2
			/* 3 */ imports.NewTable("TStrings_ToStringArrayToStringDynArray", 0), // function ToStringArrayToStringDynArray
			/* 4 */ imports.NewTable("TStrings_Add", 0), // function Add
			/* 5 */ imports.NewTable("TStrings_AddObject", 0), // function AddObject
			/* 6 */ imports.NewTable("TStrings_AddPairWithStringX2", 0), // function AddPairWithStringX2
			/* 7 */ imports.NewTable("TStrings_AddPairWithStringX2Object", 0), // function AddPairWithStringX2Object
			/* 8 */ imports.NewTable("TStrings_EqualsWithStrings", 0), // function EqualsWithStrings
			/* 9 */ imports.NewTable("TStrings_ExtractName", 0), // function ExtractName
			/* 10 */ imports.NewTable("TStrings_GetEnumerator", 0), // function GetEnumerator
			/* 11 */ imports.NewTable("TStrings_GetText", 0), // function GetText
			/* 12 */ imports.NewTable("TStrings_IndexOfWithString", 0), // function IndexOfWithString
			/* 13 */ imports.NewTable("TStrings_IndexOfWithStringInt", 0), // function IndexOfWithStringInt
			/* 14 */ imports.NewTable("TStrings_IndexOfName", 0), // function IndexOfName
			/* 15 */ imports.NewTable("TStrings_IndexOfObject", 0), // function IndexOfObject
			/* 16 */ imports.NewTable("TStrings_LastIndexOfWithStringInt", 0), // function LastIndexOfWithStringInt
			/* 17 */ imports.NewTable("TStrings_LastIndexOfWithString", 0), // function LastIndexOfWithString
			/* 18 */ imports.NewTable("TStrings_Pop", 0), // function Pop
			/* 19 */ imports.NewTable("TStrings_ReverseToStrings", 0), // function ReverseToStrings
			/* 20 */ imports.NewTable("TStrings_Shift", 0), // function Shift
			/* 21 */ imports.NewTable("TStrings_SliceWithInt", 0), // function SliceWithInt
			/* 22 */ imports.NewTable("TStrings_AddStringsWithStrings", 0), // procedure AddStringsWithStrings
			/* 23 */ imports.NewTable("TStrings_AddStringsWithStringsBool", 0), // procedure AddStringsWithStringsBool
			/* 24 */ imports.NewTable("TStrings_SetStringsWithStrings", 0), // procedure SetStringsWithStrings
			/* 25 */ imports.NewTable("TStrings_AddText", 0), // procedure AddText
			/* 26 */ imports.NewTable("TStrings_AddCommaText", 0), // procedure AddCommaText
			/* 27 */ imports.NewTable("TStrings_AddDelimitedTextWithStringCharBool", 0), // procedure AddDelimitedTextWithStringCharBool
			/* 28 */ imports.NewTable("TStrings_AddDelimitedtextWithString", 0), // procedure AddDelimitedtextWithString
			/* 29 */ imports.NewTable("TStrings_Append", 0), // procedure Append
			/* 30 */ imports.NewTable("TStrings_BeginUpdate", 0), // procedure BeginUpdate
			/* 31 */ imports.NewTable("TStrings_Clear", 0), // procedure Clear
			/* 32 */ imports.NewTable("TStrings_Delete", 0), // procedure Delete
			/* 33 */ imports.NewTable("TStrings_EndUpdate", 0), // procedure EndUpdate
			/* 34 */ imports.NewTable("TStrings_Exchange", 0), // procedure Exchange
			/* 35 */ imports.NewTable("TStrings_Fill", 0), // procedure Fill
			/* 36 */ imports.NewTable("TStrings_GetNameValue", 0), // procedure GetNameValue
			/* 37 */ imports.NewTable("TStrings_Insert", 0), // procedure Insert
			/* 38 */ imports.NewTable("TStrings_InsertObject", 0), // procedure InsertObject
			/* 39 */ imports.NewTable("TStrings_LoadFromFileWithString", 0), // procedure LoadFromFileWithString
			/* 40 */ imports.NewTable("TStrings_LoadFromFileWithStringBool", 0), // procedure LoadFromFileWithStringBool
			/* 41 */ imports.NewTable("TStrings_LoadFromStreamWithStream", 0), // procedure LoadFromStreamWithStream
			/* 42 */ imports.NewTable("TStrings_LoadFromStreamWithStreamBool", 0), // procedure LoadFromStreamWithStreamBool
			/* 43 */ imports.NewTable("TStrings_Move", 0), // procedure Move
			/* 44 */ imports.NewTable("TStrings_ReverseWithStrings", 0), // procedure ReverseWithStrings
			/* 45 */ imports.NewTable("TStrings_SaveToFileWithString", 0), // procedure SaveToFileWithString
			/* 46 */ imports.NewTable("TStrings_SaveToFileWithStringBool", 0), // procedure SaveToFileWithStringBool
			/* 47 */ imports.NewTable("TStrings_SaveToStreamWithStream", 0), // procedure SaveToStreamWithStream
			/* 48 */ imports.NewTable("TStrings_SaveToStreamWithStreamBool", 0), // procedure SaveToStreamWithStreamBool
			/* 49 */ imports.NewTable("TStrings_SliceWithIntStrings", 0), // procedure SliceWithIntStrings
			/* 50 */ imports.NewTable("TStrings_SetTextWithString", 0), // procedure SetTextWithString
			/* 51 */ imports.NewTable("TStrings_AlwaysQuote", 0), // property AlwaysQuote
			/* 52 */ imports.NewTable("TStrings_Capacity", 0), // property Capacity
			/* 53 */ imports.NewTable("TStrings_CommaText", 0), // property CommaText
			/* 54 */ imports.NewTable("TStrings_Count", 0), // property Count
			/* 55 */ imports.NewTable("TStrings_DelimitedText", 0), // property DelimitedText
			/* 56 */ imports.NewTable("TStrings_Delimiter", 0), // property Delimiter
			/* 57 */ imports.NewTable("TStrings_LineBreak", 0), // property LineBreak
			/* 58 */ imports.NewTable("TStrings_MissingNameValueSeparatorAction", 0), // property MissingNameValueSeparatorAction
			/* 59 */ imports.NewTable("TStrings_Names", 0), // property Names
			/* 60 */ imports.NewTable("TStrings_NameValueSeparator", 0), // property NameValueSeparator
			/* 61 */ imports.NewTable("TStrings_Objects", 0), // property Objects
			/* 62 */ imports.NewTable("TStrings_Options", 0), // property Options
			/* 63 */ imports.NewTable("TStrings_QuoteChar", 0), // property QuoteChar
			/* 64 */ imports.NewTable("TStrings_SkipLastLineBreak", 0), // property SkipLastLineBreak
			/* 65 */ imports.NewTable("TStrings_TrailingLineBreak", 0), // property TrailingLineBreak
			/* 66 */ imports.NewTable("TStrings_StrictDelimiter", 0), // property StrictDelimiter
			/* 67 */ imports.NewTable("TStrings_Strings", 0), // property Strings
			/* 68 */ imports.NewTable("TStrings_Text", 0), // property Text
			/* 69 */ imports.NewTable("TStrings_TextLineBreakStyle", 0), // property TextLineBreakStyle
			/* 70 */ imports.NewTable("TStrings_UseLocale", 0), // property UseLocale
			/* 71 */ imports.NewTable("TStrings_ValueFromIndex", 0), // property ValueFromIndex
			/* 72 */ imports.NewTable("TStrings_Values", 0), // property Values
			/* 73 */ imports.NewTable("TStrings_WriteBOM", 0), // property WriteBOM
		}
	})
	return stringsImport
}
