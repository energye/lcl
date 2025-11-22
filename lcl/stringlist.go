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

// IStringList Parent: IStrings
type IStringList interface {
	IStrings
	Find(S string, outIndex *int32) bool        // function
	Sort()                                      // procedure
	Duplicates() types.TDuplicates              // property Duplicates Getter
	SetDuplicates(value types.TDuplicates)      // property Duplicates Setter
	Sorted() bool                               // property Sorted Getter
	SetSorted(value bool)                       // property Sorted Setter
	CaseSensitive() bool                        // property CaseSensitive Getter
	SetCaseSensitive(value bool)                // property CaseSensitive Setter
	OwnsObjects() bool                          // property OwnsObjects Getter
	SetOwnsObjects(value bool)                  // property OwnsObjects Setter
	SortStyle() types.TStringsSortStyle         // property SortStyle Getter
	SetSortStyle(value types.TStringsSortStyle) // property SortStyle Setter
	SetOnChange(fn TNotifyEvent)                // property event
	SetOnChanging(fn TNotifyEvent)              // property event
}

type TStringList struct {
	TStrings
}

func (m *TStringList) Find(S string, outIndex *int32) bool {
	if !m.IsValid() {
		return false
	}
	var indexPtr uintptr
	r := stringListAPI().SysCallN(1, m.Instance(), api.PasStr(S), uintptr(base.UnsafePointer(&indexPtr)))
	*outIndex = int32(indexPtr)
	return api.GoBool(r)
}

func (m *TStringList) Sort() {
	if !m.IsValid() {
		return
	}
	stringListAPI().SysCallN(2, m.Instance())
}

func (m *TStringList) Duplicates() types.TDuplicates {
	if !m.IsValid() {
		return 0
	}
	r := stringListAPI().SysCallN(3, 0, m.Instance())
	return types.TDuplicates(r)
}

func (m *TStringList) SetDuplicates(value types.TDuplicates) {
	if !m.IsValid() {
		return
	}
	stringListAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TStringList) Sorted() bool {
	if !m.IsValid() {
		return false
	}
	r := stringListAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringList) SetSorted(value bool) {
	if !m.IsValid() {
		return
	}
	stringListAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringList) CaseSensitive() bool {
	if !m.IsValid() {
		return false
	}
	r := stringListAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringList) SetCaseSensitive(value bool) {
	if !m.IsValid() {
		return
	}
	stringListAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringList) OwnsObjects() bool {
	if !m.IsValid() {
		return false
	}
	r := stringListAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStringList) SetOwnsObjects(value bool) {
	if !m.IsValid() {
		return
	}
	stringListAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TStringList) SortStyle() types.TStringsSortStyle {
	if !m.IsValid() {
		return 0
	}
	r := stringListAPI().SysCallN(7, 0, m.Instance())
	return types.TStringsSortStyle(r)
}

func (m *TStringList) SetSortStyle(value types.TStringsSortStyle) {
	if !m.IsValid() {
		return
	}
	stringListAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TStringList) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, stringListAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStringList) SetOnChanging(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, stringListAPI(), api.MakeEventDataPtr(cb))
}

// NewStringList class constructor
func NewStringList() IStringList {
	r := stringListAPI().SysCallN(0)
	return AsStringList(r)
}

var (
	stringListOnce   base.Once
	stringListImport *imports.Imports = nil
)

func stringListAPI() *imports.Imports {
	stringListOnce.Do(func() {
		stringListImport = api.NewDefaultImports()
		stringListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStringList_Create", 0), // constructor NewStringList
			/* 1 */ imports.NewTable("TStringList_Find", 0), // function Find
			/* 2 */ imports.NewTable("TStringList_Sort", 0), // procedure Sort
			/* 3 */ imports.NewTable("TStringList_Duplicates", 0), // property Duplicates
			/* 4 */ imports.NewTable("TStringList_Sorted", 0), // property Sorted
			/* 5 */ imports.NewTable("TStringList_CaseSensitive", 0), // property CaseSensitive
			/* 6 */ imports.NewTable("TStringList_OwnsObjects", 0), // property OwnsObjects
			/* 7 */ imports.NewTable("TStringList_SortStyle", 0), // property SortStyle
			/* 8 */ imports.NewTable("TStringList_OnChange", 0), // event OnChange
			/* 9 */ imports.NewTable("TStringList_OnChanging", 0), // event OnChanging
		}
	})
	return stringListImport
}
