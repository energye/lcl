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

// ISynGutterLOvMark Parent: IObject
type ISynGutterLOvMark interface {
	IObject
	Line() int32                  // property Line Getter
	Column() int32                // property Column Getter
	Color() types.TColor          // property Color Getter
	Priority() int32              // property Priority Getter
	SetOnChange(fn TNotifyEvent)  // property event
	SetOnDestroy(fn TNotifyEvent) // property event
}

type TSynGutterLOvMark struct {
	TObject
}

func (m *TSynGutterLOvMark) Line() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLOvMarkAPI().SysCallN(0, m.Instance())
	return int32(r)
}

func (m *TSynGutterLOvMark) Column() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLOvMarkAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TSynGutterLOvMark) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLOvMarkAPI().SysCallN(2, m.Instance())
	return types.TColor(r)
}

func (m *TSynGutterLOvMark) Priority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLOvMarkAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TSynGutterLOvMark) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 4, synGutterLOvMarkAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynGutterLOvMark) SetOnDestroy(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, synGutterLOvMarkAPI(), api.MakeEventDataPtr(cb))
}

var (
	synGutterLOvMarkOnce   base.Once
	synGutterLOvMarkImport *imports.Imports = nil
)

func synGutterLOvMarkAPI() *imports.Imports {
	synGutterLOvMarkOnce.Do(func() {
		synGutterLOvMarkImport = api.NewDefaultImports()
		synGutterLOvMarkImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterLOvMark_Line", 0), // property Line
			/* 1 */ imports.NewTable("TSynGutterLOvMark_Column", 0), // property Column
			/* 2 */ imports.NewTable("TSynGutterLOvMark_Color", 0), // property Color
			/* 3 */ imports.NewTable("TSynGutterLOvMark_Priority", 0), // property Priority
			/* 4 */ imports.NewTable("TSynGutterLOvMark_OnChange", 0), // event OnChange
			/* 5 */ imports.NewTable("TSynGutterLOvMark_OnDestroy", 0), // event OnDestroy
		}
	})
	return synGutterLOvMarkImport
}
