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
)

// IBeautifierSetIndentCallback Parent: IObject
type IBeautifierSetIndentCallback interface {
	IObject
	// Callback
	//  LinePos: 1-based, the line that should be changed
	//  Indent: New indent in spaces (Logical = Physical
	//  RelativeToLinePos: Indent specifies +/- offset from indent on RTLine (0: for absolute indent)
	//  IndentChars:
	//  String used to build indent; maybe empty, single char, or string (usually 1 tab or 1 space)
	//  The String will be repeated and cut as needed, then filled with spaces at the end
	//  NOTE: If this is specified the TSynBeautifierIndentType is ignored
	//  IndentCharsFromLinePos:
	//  Use tab/space mix from this Line for indent (if specified > 0)
	//  "IndentChars" will only be used, if the found tab/space mix is to short
	//  NOTE: If this is specified the TSynBeautifierIndentType is ignored
	Callback(linePos int32, indent int32, relativeToLinePos int32, indentChars string, indentCharsFromLinePos int32) // procedure
}

type TBeautifierSetIndentCallback struct {
	TObject
}

func (m *TBeautifierSetIndentCallback) Callback(linePos int32, indent int32, relativeToLinePos int32, indentChars string, indentCharsFromLinePos int32) {
	if !m.IsValid() {
		return
	}
	beautifierSetIndentCallbackAPI().SysCallN(0, m.Instance(), uintptr(linePos), uintptr(indent), uintptr(relativeToLinePos), api.PasStr(indentChars), uintptr(indentCharsFromLinePos))
}

var (
	beautifierSetIndentCallbackOnce   base.Once
	beautifierSetIndentCallbackImport *imports.Imports = nil
)

func beautifierSetIndentCallbackAPI() *imports.Imports {
	beautifierSetIndentCallbackOnce.Do(func() {
		beautifierSetIndentCallbackImport = api.NewDefaultImports()
		beautifierSetIndentCallbackImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBeautifierSetIndentCallback_Callback", 0), // procedure Callback
		}
	})
	return beautifierSetIndentCallbackImport
}
