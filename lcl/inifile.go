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

// IIniFile Parent: ICustomIniFile
type IIniFile interface {
	ICustomIniFile
	ReadSectionRaw(section string, strings IStrings) // procedure
	Stream() IStream                                 // property Stream Getter
	CacheUpdates() bool                              // property CacheUpdates Getter
	SetCacheUpdates(value bool)                      // property CacheUpdates Setter
	WriteBOM() bool                                  // property WriteBOM Getter
	SetWriteBOM(value bool)                          // property WriteBOM Setter
}

type TIniFile struct {
	TCustomIniFile
}

func (m *TIniFile) ReadSectionRaw(section string, strings IStrings) {
	if !m.IsValid() {
		return
	}
	iniFileAPI().SysCallN(3, m.Instance(), api.PasStr(section), base.GetObjectUintptr(strings))
}

func (m *TIniFile) Stream() IStream {
	if !m.IsValid() {
		return nil
	}
	r := iniFileAPI().SysCallN(4, m.Instance())
	return AsStream(r)
}

func (m *TIniFile) CacheUpdates() bool {
	if !m.IsValid() {
		return false
	}
	r := iniFileAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TIniFile) SetCacheUpdates(value bool) {
	if !m.IsValid() {
		return
	}
	iniFileAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TIniFile) WriteBOM() bool {
	if !m.IsValid() {
		return false
	}
	r := iniFileAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TIniFile) SetWriteBOM(value bool) {
	if !m.IsValid() {
		return
	}
	iniFileAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

// NewIniFile class constructor
func NewIniFile(fileName string, options types.TIniFileOptions) IIniFile {
	r := iniFileAPI().SysCallN(0, api.PasStr(fileName), uintptr(options))
	return AsIniFile(r)
}

// NewIniFileWithStreamIniFileoptions class constructor
func NewIniFileWithStreamIniFileoptions(stream IStream, options types.TIniFileOptions) IIniFile {
	r := iniFileAPI().SysCallN(1, base.GetObjectUintptr(stream), uintptr(options))
	return AsIniFile(r)
}

// NewIniFileWithStreamBool class constructor
func NewIniFileWithStreamBool(stream IStream, escapeLineFeeds bool) IIniFile {
	r := iniFileAPI().SysCallN(2, base.GetObjectUintptr(stream), api.PasBool(escapeLineFeeds))
	return AsIniFile(r)
}

var (
	iniFileOnce   base.Once
	iniFileImport *imports.Imports = nil
)

func iniFileAPI() *imports.Imports {
	iniFileOnce.Do(func() {
		iniFileImport = api.NewDefaultImports()
		iniFileImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TIniFile_Create", 0), // constructor NewIniFile
			/* 1 */ imports.NewTable("TIniFile_CreateWithStreamIniFileoptions", 0), // constructor NewIniFileWithStreamIniFileoptions
			/* 2 */ imports.NewTable("TIniFile_CreateWithStreamBool", 0), // constructor NewIniFileWithStreamBool
			/* 3 */ imports.NewTable("TIniFile_ReadSectionRaw", 0), // procedure ReadSectionRaw
			/* 4 */ imports.NewTable("TIniFile_Stream", 0), // property Stream
			/* 5 */ imports.NewTable("TIniFile_CacheUpdates", 0), // property CacheUpdates
			/* 6 */ imports.NewTable("TIniFile_WriteBOM", 0), // property WriteBOM
		}
	})
	return iniFileImport
}
