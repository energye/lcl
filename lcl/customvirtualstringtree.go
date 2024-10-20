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

// ICustomVirtualStringTree Parent: IBaseVirtualTree
type ICustomVirtualStringTree interface {
	IBaseVirtualTree
	ImageText(Node IVirtualNode, Kind TVTImageKind, Column TColumnIndex) string                // property
	StaticText(Node IVirtualNode, Column TColumnIndex) string                                  // property
	Text(Node IVirtualNode, Column TColumnIndex) string                                        // property
	SetText(Node IVirtualNode, Column TColumnIndex, AValue string)                             // property
	ComputeNodeHeight(Canvas ICanvas, Node IVirtualNode, Column TColumnIndex, S string) int32  // function
	ContentToClipboard(Format TClipboardFormat, Source TVSTTextSourceType) HGLOBAL             // function
	ContentToHTML(Source TVSTTextSourceType, Caption string) string                            // function
	ContentToRTF(Source TVSTTextSourceType) string                                             // function
	ContentToAnsi(Source TVSTTextSourceType, Separator string) string                          // function
	ContentToText(Source TVSTTextSourceType, Separator string) string                          // function
	ContentToUnicode(Source TVSTTextSourceType, Separator string) string                       // function
	ContentToUTF16(Source TVSTTextSourceType, Separator string) string                         // function
	ContentToUTF8(Source TVSTTextSourceType, Separator string) string                          // function
	Path(Node IVirtualNode, Column TColumnIndex, TextType TVSTTextType, Delimiter Char) string // function
	SaveToCSVFile(FileNameWithPath string, IncludeHeading bool) bool                           // function
	ContentToCustom(Source TVSTTextSourceType)                                                 // procedure
	AddToSelection(Node IVirtualNode)                                                          // procedure
	RemoveFromSelection(Node IVirtualNode)                                                     // procedure
}

// TCustomVirtualStringTree Parent: TBaseVirtualTree
type TCustomVirtualStringTree struct {
	TBaseVirtualTree
}

func NewCustomVirtualStringTree(AOwner IComponent) ICustomVirtualStringTree {
	r1 := customVirtualStringTreeImportAPI().SysCallN(12, GetObjectUintptr(AOwner))
	return AsCustomVirtualStringTree(r1)
}

func (m *TCustomVirtualStringTree) ImageText(Node IVirtualNode, Kind TVTImageKind, Column TColumnIndex) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(Node), uintptr(Kind), uintptr(Column))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) StaticText(Node IVirtualNode, Column TColumnIndex) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) Text(Node IVirtualNode, Column TColumnIndex) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(18, 0, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) SetText(Node IVirtualNode, Column TColumnIndex, AValue string) {
	customVirtualStringTreeImportAPI().SysCallN(18, 1, m.Instance(), GetObjectUintptr(Node), uintptr(Column), PascalStr(AValue))
}

func (m *TCustomVirtualStringTree) ComputeNodeHeight(Canvas ICanvas, Node IVirtualNode, Column TColumnIndex, S string) int32 {
	r1 := customVirtualStringTreeImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(Canvas), GetObjectUintptr(Node), uintptr(Column), PascalStr(S))
	return int32(r1)
}

func (m *TCustomVirtualStringTree) ContentToClipboard(Format TClipboardFormat, Source TVSTTextSourceType) HGLOBAL {
	r1 := customVirtualStringTreeImportAPI().SysCallN(4, m.Instance(), uintptr(Format), uintptr(Source))
	return HGLOBAL(r1)
}

func (m *TCustomVirtualStringTree) ContentToHTML(Source TVSTTextSourceType, Caption string) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(6, m.Instance(), uintptr(Source), PascalStr(Caption))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToRTF(Source TVSTTextSourceType) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(7, m.Instance(), uintptr(Source))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToAnsi(Source TVSTTextSourceType, Separator string) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(3, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToText(Source TVSTTextSourceType, Separator string) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(8, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToUnicode(Source TVSTTextSourceType, Separator string) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(11, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToUTF16(Source TVSTTextSourceType, Separator string) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(9, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) ContentToUTF8(Source TVSTTextSourceType, Separator string) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(10, m.Instance(), uintptr(Source), PascalStr(Separator))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) Path(Node IVirtualNode, Column TColumnIndex, TextType TVSTTextType, Delimiter Char) string {
	r1 := customVirtualStringTreeImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(Node), uintptr(Column), uintptr(TextType), uintptr(Delimiter))
	return GoStr(r1)
}

func (m *TCustomVirtualStringTree) SaveToCSVFile(FileNameWithPath string, IncludeHeading bool) bool {
	r1 := customVirtualStringTreeImportAPI().SysCallN(16, m.Instance(), PascalStr(FileNameWithPath), PascalBool(IncludeHeading))
	return GoBool(r1)
}

func CustomVirtualStringTreeClass() TClass {
	ret := customVirtualStringTreeImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomVirtualStringTree) ContentToCustom(Source TVSTTextSourceType) {
	customVirtualStringTreeImportAPI().SysCallN(5, m.Instance(), uintptr(Source))
}

func (m *TCustomVirtualStringTree) AddToSelection(Node IVirtualNode) {
	customVirtualStringTreeImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Node))
}

func (m *TCustomVirtualStringTree) RemoveFromSelection(Node IVirtualNode) {
	customVirtualStringTreeImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(Node))
}

var (
	customVirtualStringTreeImport       *imports.Imports = nil
	customVirtualStringTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomVirtualStringTree_AddToSelection", 0),
		/*1*/ imports.NewTable("CustomVirtualStringTree_Class", 0),
		/*2*/ imports.NewTable("CustomVirtualStringTree_ComputeNodeHeight", 0),
		/*3*/ imports.NewTable("CustomVirtualStringTree_ContentToAnsi", 0),
		/*4*/ imports.NewTable("CustomVirtualStringTree_ContentToClipboard", 0),
		/*5*/ imports.NewTable("CustomVirtualStringTree_ContentToCustom", 0),
		/*6*/ imports.NewTable("CustomVirtualStringTree_ContentToHTML", 0),
		/*7*/ imports.NewTable("CustomVirtualStringTree_ContentToRTF", 0),
		/*8*/ imports.NewTable("CustomVirtualStringTree_ContentToText", 0),
		/*9*/ imports.NewTable("CustomVirtualStringTree_ContentToUTF16", 0),
		/*10*/ imports.NewTable("CustomVirtualStringTree_ContentToUTF8", 0),
		/*11*/ imports.NewTable("CustomVirtualStringTree_ContentToUnicode", 0),
		/*12*/ imports.NewTable("CustomVirtualStringTree_Create", 0),
		/*13*/ imports.NewTable("CustomVirtualStringTree_ImageText", 0),
		/*14*/ imports.NewTable("CustomVirtualStringTree_Path", 0),
		/*15*/ imports.NewTable("CustomVirtualStringTree_RemoveFromSelection", 0),
		/*16*/ imports.NewTable("CustomVirtualStringTree_SaveToCSVFile", 0),
		/*17*/ imports.NewTable("CustomVirtualStringTree_StaticText", 0),
		/*18*/ imports.NewTable("CustomVirtualStringTree_Text", 0),
	}
)

func customVirtualStringTreeImportAPI() *imports.Imports {
	if customVirtualStringTreeImport == nil {
		customVirtualStringTreeImport = NewDefaultImports()
		customVirtualStringTreeImport.SetImportTable(customVirtualStringTreeImportTables)
		customVirtualStringTreeImportTables = nil
	}
	return customVirtualStringTreeImport
}
