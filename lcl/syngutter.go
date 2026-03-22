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

// ISynGutter Parent: ISynGutterBase
type ISynGutter interface {
	ISynGutterBase
	HasCustomPopupMenu(outPopMenu *IPopupMenu) bool                                         // function
	DoHandleMouseAction(anAction ISynEditMouseAction, anInfo *TSynEditMouseActionInfo) bool // function
	// LineNumberPart
	//  Access to well known parts
	LineNumberPart(index int32) ISynGutterLineNumber                                                    // function
	CodeFoldPart(index int32) ISynGutterCodeFolding                                                     // function
	ChangesPart(index int32) ISynGutterChanges                                                          // function
	MarksPart(index int32) ISynGutterMarks                                                              // function
	SeparatorPart(index int32) ISynGutterSeparator                                                      // function
	LineOverviewPart(index int32) ISynGutterLineOverview                                                // function
	Paint(canvas ICanvas, surface ILazSynGutterArea, clip types.TRect, firstLine int32, lastLine int32) // procedure
	MouseDown(button types.TMouseButton, shift types.TShiftState, X int32, Y int32)                     // procedure
	MouseMove(shift types.TShiftState, X int32, Y int32)                                                // procedure
	MouseUp(button types.TMouseButton, shift types.TShiftState, X int32, Y int32)                       // procedure
	DoOnGutterClick(X int32, Y int32)                                                                   // procedure
	Cursor() types.TCursor                                                                              // property Cursor Getter
	SetCursor(value types.TCursor)                                                                      // property Cursor Setter
	SetOnGutterClick(fn TGutterClickEvent)                                                              // property event
}

type TSynGutter struct {
	TSynGutterBase
}

func (m *TSynGutter) HasCustomPopupMenu(outPopMenu *IPopupMenu) bool {
	if !m.IsValid() {
		return false
	}
	var popMenuPtr uintptr
	r := synGutterAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&popMenuPtr)))
	*outPopMenu = AsPopupMenu(popMenuPtr)
	return api.GoBool(r)
}

func (m *TSynGutter) DoHandleMouseAction(anAction ISynEditMouseAction, anInfo *TSynEditMouseActionInfo) bool {
	if !m.IsValid() {
		return false
	}
	anInfoPtr := anInfo.ToPas()
	r := synGutterAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(anAction), uintptr(base.UnsafePointer(anInfoPtr)))
	*anInfo = anInfoPtr.ToGo()
	return api.GoBool(r)
}

func (m *TSynGutter) LineNumberPart(index int32) ISynGutterLineNumber {
	if !m.IsValid() {
		return nil
	}
	r := synGutterAPI().SysCallN(3, m.Instance(), uintptr(index))
	return AsSynGutterLineNumber(r)
}

func (m *TSynGutter) CodeFoldPart(index int32) ISynGutterCodeFolding {
	if !m.IsValid() {
		return nil
	}
	r := synGutterAPI().SysCallN(4, m.Instance(), uintptr(index))
	return AsSynGutterCodeFolding(r)
}

func (m *TSynGutter) ChangesPart(index int32) ISynGutterChanges {
	if !m.IsValid() {
		return nil
	}
	r := synGutterAPI().SysCallN(5, m.Instance(), uintptr(index))
	return AsSynGutterChanges(r)
}

func (m *TSynGutter) MarksPart(index int32) ISynGutterMarks {
	if !m.IsValid() {
		return nil
	}
	r := synGutterAPI().SysCallN(6, m.Instance(), uintptr(index))
	return AsSynGutterMarks(r)
}

func (m *TSynGutter) SeparatorPart(index int32) ISynGutterSeparator {
	if !m.IsValid() {
		return nil
	}
	r := synGutterAPI().SysCallN(7, m.Instance(), uintptr(index))
	return AsSynGutterSeparator(r)
}

func (m *TSynGutter) LineOverviewPart(index int32) ISynGutterLineOverview {
	if !m.IsValid() {
		return nil
	}
	r := synGutterAPI().SysCallN(8, m.Instance(), uintptr(index))
	return AsSynGutterLineOverview(r)
}

func (m *TSynGutter) Paint(canvas ICanvas, surface ILazSynGutterArea, clip types.TRect, firstLine int32, lastLine int32) {
	if !m.IsValid() {
		return
	}
	synGutterAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(canvas), base.GetObjectUintptr(surface), uintptr(base.UnsafePointer(&clip)), uintptr(firstLine), uintptr(lastLine))
}

func (m *TSynGutter) MouseDown(button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterAPI().SysCallN(10, m.Instance(), uintptr(button), uintptr(shift), uintptr(X), uintptr(Y))
}

func (m *TSynGutter) MouseMove(shift types.TShiftState, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterAPI().SysCallN(11, m.Instance(), uintptr(shift), uintptr(X), uintptr(Y))
}

func (m *TSynGutter) MouseUp(button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterAPI().SysCallN(12, m.Instance(), uintptr(button), uintptr(shift), uintptr(X), uintptr(Y))
}

func (m *TSynGutter) DoOnGutterClick(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterAPI().SysCallN(13, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TSynGutter) Cursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := synGutterAPI().SysCallN(14, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TSynGutter) SetCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	synGutterAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutter) SetOnGutterClick(fn TGutterClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGutterClickEvent(fn)
	base.SetEvent(m, 15, synGutterAPI(), api.MakeEventDataPtr(cb))
}

// NewSynGutter class constructor
func NewSynGutter(owner ISynEditBase, side types.TSynGutterSide, textDrawer IheTextDrawer) ISynGutter {
	r := synGutterAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(side), base.GetObjectUintptr(textDrawer))
	return AsSynGutter(r)
}

var (
	synGutterOnce   base.Once
	synGutterImport *imports.Imports = nil
)

func synGutterAPI() *imports.Imports {
	synGutterOnce.Do(func() {
		synGutterImport = api.NewDefaultImports()
		synGutterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutter_Create", 0), // constructor NewSynGutter
			/* 1 */ imports.NewTable("TSynGutter_HasCustomPopupMenu", 0), // function HasCustomPopupMenu
			/* 2 */ imports.NewTable("TSynGutter_DoHandleMouseAction", 0), // function DoHandleMouseAction
			/* 3 */ imports.NewTable("TSynGutter_LineNumberPart", 0), // function LineNumberPart
			/* 4 */ imports.NewTable("TSynGutter_CodeFoldPart", 0), // function CodeFoldPart
			/* 5 */ imports.NewTable("TSynGutter_ChangesPart", 0), // function ChangesPart
			/* 6 */ imports.NewTable("TSynGutter_MarksPart", 0), // function MarksPart
			/* 7 */ imports.NewTable("TSynGutter_SeparatorPart", 0), // function SeparatorPart
			/* 8 */ imports.NewTable("TSynGutter_LineOverviewPart", 0), // function LineOverviewPart
			/* 9 */ imports.NewTable("TSynGutter_Paint", 0), // procedure Paint
			/* 10 */ imports.NewTable("TSynGutter_MouseDown", 0), // procedure MouseDown
			/* 11 */ imports.NewTable("TSynGutter_MouseMove", 0), // procedure MouseMove
			/* 12 */ imports.NewTable("TSynGutter_MouseUp", 0), // procedure MouseUp
			/* 13 */ imports.NewTable("TSynGutter_DoOnGutterClick", 0), // procedure DoOnGutterClick
			/* 14 */ imports.NewTable("TSynGutter_Cursor", 0), // property Cursor
			/* 15 */ imports.NewTable("TSynGutter_OnGutterClick", 0), // event OnGutterClick
		}
	})
	return synGutterImport
}
