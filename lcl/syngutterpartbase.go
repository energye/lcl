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

// ISynGutterPartBase Parent: ISynObjectListItem
type ISynGutterPartBase interface {
	ISynObjectListItem
	// HasCustomPopupMenu
	//  X/Y are relative to the gutter, not the gutter part
	HasCustomPopupMenu(outPopMenu *IPopupMenu) bool                                         // function
	DoHandleMouseAction(anAction ISynEditMouseAction, anInfo *TSynEditMouseActionInfo) bool // function
	PaintAll(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32)             // procedure
	Paint(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32)                // procedure
	ScalePPI(scaleFactor float64)                                                           // procedure
	MouseDown(button types.TMouseButton, shift types.TShiftState, X int32, Y int32)         // procedure
	MouseMove(shift types.TShiftState, X int32, Y int32)                                    // procedure
	MouseUp(button types.TMouseButton, shift types.TShiftState, X int32, Y int32)           // procedure
	ResetMouseActions()                                                                     // procedure
	DoOnGutterClick(X int32, Y int32)                                                       // procedure
	Left() int32                                                                            // property Left Getter
	Top() int32                                                                             // property Top Getter
	Height() int32                                                                          // property Height Getter
	Cursor() types.TCursor                                                                  // property Cursor Getter
	SetCursor(value types.TCursor)                                                          // property Cursor Setter
	MarkupInfo() ISynSelectedColor                                                          // property MarkupInfo Getter
	SetMarkupInfo(value ISynSelectedColor)                                                  // property MarkupInfo Setter
	MarkupInfoCurrentLine() ISynHighlighterAttributesModifier                               // property MarkupInfoCurrentLine Getter
	SetMarkupInfoCurrentLine(value ISynHighlighterAttributesModifier)                       // property MarkupInfoCurrentLine Setter
	AutoSize() bool                                                                         // property AutoSize Getter
	SetAutoSize(value bool)                                                                 // property AutoSize Setter
	Width() int32                                                                           // property Width Getter
	SetWidth(value int32)                                                                   // property Width Setter
	FullWidth() int32                                                                       // property FullWidth Getter
	LeftOffset() int32                                                                      // property LeftOffset Getter
	SetLeftOffset(value int32)                                                              // property LeftOffset Setter
	RightOffset() int32                                                                     // property RightOffset Getter
	SetRightOffset(value int32)                                                             // property RightOffset Setter
	Visible() bool                                                                          // property Visible Getter
	SetVisible(value bool)                                                                  // property Visible Setter
	MouseActions() ISynEditMouseActions                                                     // property MouseActions Getter
	SetMouseActions(value ISynEditMouseActions)                                             // property MouseActions Setter
	SetOnGutterClick(fn TGutterClickEvent)                                                  // property event
	SetOnChange(fn TNotifyEvent)                                                            // property event
}

type TSynGutterPartBase struct {
	TSynObjectListItem
}

func (m *TSynGutterPartBase) HasCustomPopupMenu(outPopMenu *IPopupMenu) bool {
	if !m.IsValid() {
		return false
	}
	var popMenuPtr uintptr
	r := synGutterPartBaseAPI().SysCallN(0, m.Instance(), uintptr(base.UnsafePointer(&popMenuPtr)))
	*outPopMenu = AsPopupMenu(popMenuPtr)
	return api.GoBool(r)
}

func (m *TSynGutterPartBase) DoHandleMouseAction(anAction ISynEditMouseAction, anInfo *TSynEditMouseActionInfo) bool {
	if !m.IsValid() {
		return false
	}
	anInfoPtr := anInfo.ToPas()
	r := synGutterPartBaseAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(anAction), uintptr(base.UnsafePointer(anInfoPtr)))
	*anInfo = anInfoPtr.ToGo()
	return api.GoBool(r)
}

func (m *TSynGutterPartBase) PaintAll(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&clip)), uintptr(firstLine), uintptr(lastLine))
}

func (m *TSynGutterPartBase) Paint(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&clip)), uintptr(firstLine), uintptr(lastLine))
}

func (m *TSynGutterPartBase) ScalePPI(scaleFactor float64) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&scaleFactor)))
}

func (m *TSynGutterPartBase) MouseDown(button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(5, m.Instance(), uintptr(button), uintptr(shift), uintptr(X), uintptr(Y))
}

func (m *TSynGutterPartBase) MouseMove(shift types.TShiftState, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(6, m.Instance(), uintptr(shift), uintptr(X), uintptr(Y))
}

func (m *TSynGutterPartBase) MouseUp(button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(7, m.Instance(), uintptr(button), uintptr(shift), uintptr(X), uintptr(Y))
}

func (m *TSynGutterPartBase) ResetMouseActions() {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(8, m.Instance())
}

func (m *TSynGutterPartBase) DoOnGutterClick(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(9, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TSynGutterPartBase) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TSynGutterPartBase) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TSynGutterPartBase) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(12, m.Instance())
	return int32(r)
}

func (m *TSynGutterPartBase) Cursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(13, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TSynGutterPartBase) SetCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterPartBase) MarkupInfo() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synGutterPartBaseAPI().SysCallN(14, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynGutterPartBase) SetMarkupInfo(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterPartBase) MarkupInfoCurrentLine() ISynHighlighterAttributesModifier {
	if !m.IsValid() {
		return nil
	}
	r := synGutterPartBaseAPI().SysCallN(15, 0, m.Instance())
	return AsSynHighlighterAttributesModifier(r)
}

func (m *TSynGutterPartBase) SetMarkupInfoCurrentLine(value ISynHighlighterAttributesModifier) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterPartBase) AutoSize() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterPartBaseAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterPartBase) SetAutoSize(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynGutterPartBase) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterPartBase) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterPartBase) FullWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(18, m.Instance())
	return int32(r)
}

func (m *TSynGutterPartBase) LeftOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterPartBase) SetLeftOffset(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterPartBase) RightOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartBaseAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterPartBase) SetRightOffset(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterPartBase) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterPartBaseAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterPartBase) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynGutterPartBase) MouseActions() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synGutterPartBaseAPI().SysCallN(22, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynGutterPartBase) SetMouseActions(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synGutterPartBaseAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterPartBase) SetOnGutterClick(fn TGutterClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGutterClickEvent(fn)
	base.SetEvent(m, 23, synGutterPartBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynGutterPartBase) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, synGutterPartBaseAPI(), api.MakeEventDataPtr(cb))
}

var (
	synGutterPartBaseOnce   base.Once
	synGutterPartBaseImport *imports.Imports = nil
)

func synGutterPartBaseAPI() *imports.Imports {
	synGutterPartBaseOnce.Do(func() {
		synGutterPartBaseImport = api.NewDefaultImports()
		synGutterPartBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterPartBase_HasCustomPopupMenu", 0), // function HasCustomPopupMenu
			/* 1 */ imports.NewTable("TSynGutterPartBase_DoHandleMouseAction", 0), // function DoHandleMouseAction
			/* 2 */ imports.NewTable("TSynGutterPartBase_PaintAll", 0), // procedure PaintAll
			/* 3 */ imports.NewTable("TSynGutterPartBase_Paint", 0), // procedure Paint
			/* 4 */ imports.NewTable("TSynGutterPartBase_ScalePPI", 0), // procedure ScalePPI
			/* 5 */ imports.NewTable("TSynGutterPartBase_MouseDown", 0), // procedure MouseDown
			/* 6 */ imports.NewTable("TSynGutterPartBase_MouseMove", 0), // procedure MouseMove
			/* 7 */ imports.NewTable("TSynGutterPartBase_MouseUp", 0), // procedure MouseUp
			/* 8 */ imports.NewTable("TSynGutterPartBase_ResetMouseActions", 0), // procedure ResetMouseActions
			/* 9 */ imports.NewTable("TSynGutterPartBase_DoOnGutterClick", 0), // procedure DoOnGutterClick
			/* 10 */ imports.NewTable("TSynGutterPartBase_Left", 0), // property Left
			/* 11 */ imports.NewTable("TSynGutterPartBase_Top", 0), // property Top
			/* 12 */ imports.NewTable("TSynGutterPartBase_Height", 0), // property Height
			/* 13 */ imports.NewTable("TSynGutterPartBase_Cursor", 0), // property Cursor
			/* 14 */ imports.NewTable("TSynGutterPartBase_MarkupInfo", 0), // property MarkupInfo
			/* 15 */ imports.NewTable("TSynGutterPartBase_MarkupInfoCurrentLine", 0), // property MarkupInfoCurrentLine
			/* 16 */ imports.NewTable("TSynGutterPartBase_AutoSize", 0), // property AutoSize
			/* 17 */ imports.NewTable("TSynGutterPartBase_Width", 0), // property Width
			/* 18 */ imports.NewTable("TSynGutterPartBase_FullWidth", 0), // property FullWidth
			/* 19 */ imports.NewTable("TSynGutterPartBase_LeftOffset", 0), // property LeftOffset
			/* 20 */ imports.NewTable("TSynGutterPartBase_RightOffset", 0), // property RightOffset
			/* 21 */ imports.NewTable("TSynGutterPartBase_Visible", 0), // property Visible
			/* 22 */ imports.NewTable("TSynGutterPartBase_MouseActions", 0), // property MouseActions
			/* 23 */ imports.NewTable("TSynGutterPartBase_OnGutterClick", 0), // event OnGutterClick
			/* 24 */ imports.NewTable("TSynGutterPartBase_OnChange", 0), // event OnChange
		}
	})
	return synGutterPartBaseImport
}
