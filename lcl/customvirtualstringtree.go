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

// ICustomVirtualStringTree Parent: IBaseVirtualTree
type ICustomVirtualStringTree interface {
	IBaseVirtualTree
	ComputeNodeHeight(canvas ICanvas, node types.PVirtualNode, column int32, S string) int32          // function
	ContentToClipboard(format types.TClipboardFormat, source types.TVSTTextSourceType) types.HGLOBAL  // function
	ContentToHTML(source types.TVSTTextSourceType, caption string) string                             // function
	ContentToRTF(source types.TVSTTextSourceType) string                                              // function
	ContentToAnsi(source types.TVSTTextSourceType, separator string) string                           // function
	ContentToText(source types.TVSTTextSourceType, separator string) string                           // function
	ContentToUnicode(source types.TVSTTextSourceType, separator string) string                        // function
	ContentToUTF16(source types.TVSTTextSourceType, separator string) string                          // function
	ContentToUTF8(source types.TVSTTextSourceType, separator string) string                           // function
	Path(node types.PVirtualNode, column int32, textType types.TVSTTextType, delimiter uint16) string // function
	SaveToCSVFile(fileNameWithPath string, includeHeading bool) bool                                  // function
	ContentToCustom(source types.TVSTTextSourceType)                                                  // procedure
	AddToSelection(node types.PVirtualNode)                                                           // procedure
	RemoveFromSelection(node types.PVirtualNode)                                                      // procedure
	ImageText(node types.PVirtualNode, kind types.TVTImageKind, column int32) string                  // property ImageText Getter
	StaticText(node types.PVirtualNode, column int32) string                                          // property StaticText Getter
	Text(node types.PVirtualNode, column int32) string                                                // property Text Getter
	SetText(node types.PVirtualNode, column int32, value string)                                      // property Text Setter
}

type TCustomVirtualStringTree struct {
	TBaseVirtualTree
}

func (m *TCustomVirtualStringTree) ComputeNodeHeight(canvas ICanvas, node types.PVirtualNode, column int32, S string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customVirtualStringTreeAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(canvas), uintptr(node), uintptr(column), api.PasStr(S))
	return int32(r)
}

func (m *TCustomVirtualStringTree) ContentToClipboard(format types.TClipboardFormat, source types.TVSTTextSourceType) types.HGLOBAL {
	if !m.IsValid() {
		return 0
	}
	r := customVirtualStringTreeAPI().SysCallN(2, m.Instance(), uintptr(format), uintptr(source))
	return types.HGLOBAL(r)
}

func (m *TCustomVirtualStringTree) ContentToHTML(source types.TVSTTextSourceType, caption string) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(3, m.Instance(), uintptr(source), api.PasStr(caption))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) ContentToRTF(source types.TVSTTextSourceType) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(4, m.Instance(), uintptr(source))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) ContentToAnsi(source types.TVSTTextSourceType, separator string) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(5, m.Instance(), uintptr(source), api.PasStr(separator))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) ContentToText(source types.TVSTTextSourceType, separator string) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(6, m.Instance(), uintptr(source), api.PasStr(separator))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) ContentToUnicode(source types.TVSTTextSourceType, separator string) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(7, m.Instance(), uintptr(source), api.PasStr(separator))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) ContentToUTF16(source types.TVSTTextSourceType, separator string) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(8, m.Instance(), uintptr(source), api.PasStr(separator))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) ContentToUTF8(source types.TVSTTextSourceType, separator string) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(9, m.Instance(), uintptr(source), api.PasStr(separator))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) Path(node types.PVirtualNode, column int32, textType types.TVSTTextType, delimiter uint16) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(10, m.Instance(), uintptr(node), uintptr(column), uintptr(textType), uintptr(delimiter))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) SaveToCSVFile(fileNameWithPath string, includeHeading bool) bool {
	if !m.IsValid() {
		return false
	}
	r := customVirtualStringTreeAPI().SysCallN(11, m.Instance(), api.PasStr(fileNameWithPath), api.PasBool(includeHeading))
	return api.GoBool(r)
}

func (m *TCustomVirtualStringTree) ContentToCustom(source types.TVSTTextSourceType) {
	if !m.IsValid() {
		return
	}
	customVirtualStringTreeAPI().SysCallN(12, m.Instance(), uintptr(source))
}

func (m *TCustomVirtualStringTree) AddToSelection(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	customVirtualStringTreeAPI().SysCallN(13, m.Instance(), uintptr(node))
}

func (m *TCustomVirtualStringTree) RemoveFromSelection(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	customVirtualStringTreeAPI().SysCallN(14, m.Instance(), uintptr(node))
}

func (m *TCustomVirtualStringTree) ImageText(node types.PVirtualNode, kind types.TVTImageKind, column int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(15, m.Instance(), uintptr(node), uintptr(kind), uintptr(column))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) StaticText(node types.PVirtualNode, column int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(16, m.Instance(), uintptr(node), uintptr(column))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) Text(node types.PVirtualNode, column int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customVirtualStringTreeAPI().SysCallN(17, 0, m.Instance(), uintptr(node), uintptr(column))
	return api.GoStr(r)
}

func (m *TCustomVirtualStringTree) SetText(node types.PVirtualNode, column int32, value string) {
	if !m.IsValid() {
		return
	}
	customVirtualStringTreeAPI().SysCallN(17, 1, m.Instance(), uintptr(node), uintptr(column), api.PasStr(value))
}

// NewCustomVirtualStringTree class constructor
func NewCustomVirtualStringTree(owner IComponent) ICustomVirtualStringTree {
	r := customVirtualStringTreeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomVirtualStringTree(r)
}

func TCustomVirtualStringTreeClass() types.TClass {
	r := customVirtualStringTreeAPI().SysCallN(18)
	return types.TClass(r)
}

var (
	customVirtualStringTreeOnce   base.Once
	customVirtualStringTreeImport *imports.Imports = nil
)

func customVirtualStringTreeAPI() *imports.Imports {
	customVirtualStringTreeOnce.Do(func() {
		customVirtualStringTreeImport = api.NewDefaultImports()
		customVirtualStringTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomVirtualStringTree_Create", 0), // constructor NewCustomVirtualStringTree
			/* 1 */ imports.NewTable("TCustomVirtualStringTree_ComputeNodeHeight", 0), // function ComputeNodeHeight
			/* 2 */ imports.NewTable("TCustomVirtualStringTree_ContentToClipboard", 0), // function ContentToClipboard
			/* 3 */ imports.NewTable("TCustomVirtualStringTree_ContentToHTML", 0), // function ContentToHTML
			/* 4 */ imports.NewTable("TCustomVirtualStringTree_ContentToRTF", 0), // function ContentToRTF
			/* 5 */ imports.NewTable("TCustomVirtualStringTree_ContentToAnsi", 0), // function ContentToAnsi
			/* 6 */ imports.NewTable("TCustomVirtualStringTree_ContentToText", 0), // function ContentToText
			/* 7 */ imports.NewTable("TCustomVirtualStringTree_ContentToUnicode", 0), // function ContentToUnicode
			/* 8 */ imports.NewTable("TCustomVirtualStringTree_ContentToUTF16", 0), // function ContentToUTF16
			/* 9 */ imports.NewTable("TCustomVirtualStringTree_ContentToUTF8", 0), // function ContentToUTF8
			/* 10 */ imports.NewTable("TCustomVirtualStringTree_Path", 0), // function Path
			/* 11 */ imports.NewTable("TCustomVirtualStringTree_SaveToCSVFile", 0), // function SaveToCSVFile
			/* 12 */ imports.NewTable("TCustomVirtualStringTree_ContentToCustom", 0), // procedure ContentToCustom
			/* 13 */ imports.NewTable("TCustomVirtualStringTree_AddToSelection", 0), // procedure AddToSelection
			/* 14 */ imports.NewTable("TCustomVirtualStringTree_RemoveFromSelection", 0), // procedure RemoveFromSelection
			/* 15 */ imports.NewTable("TCustomVirtualStringTree_ImageText", 0), // property ImageText
			/* 16 */ imports.NewTable("TCustomVirtualStringTree_StaticText", 0), // property StaticText
			/* 17 */ imports.NewTable("TCustomVirtualStringTree_Text", 0), // property Text
			/* 18 */ imports.NewTable("TCustomVirtualStringTree_TClass", 0), // function TCustomVirtualStringTreeClass
		}
	})
	return customVirtualStringTreeImport
}
