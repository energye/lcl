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

// ICustomGraphicControl Parent: IGraphicControl
type ICustomGraphicControl interface {
	IGraphicControl
	InheritedWndProc(theMessage *types.TLMessage) // procedure
	// ActionLink
	//  optional properties (not every descendent supports them)
	ActionLink() IControlActionLink                    // property ActionLink Getter
	SetActionLink(value IControlActionLink)            // property ActionLink Setter
	DesktopFont() bool                                 // property DesktopFont Getter
	SetDesktopFont(value bool)                         // property DesktopFont Setter
	DragCursor() types.TCursor                         // property DragCursor Getter
	SetDragCursor(value types.TCursor)                 // property DragCursor Setter
	DragKind() types.TDragKind                         // property DragKind Getter
	SetDragKind(value types.TDragKind)                 // property DragKind Setter
	DragMode() types.TDragMode                         // property DragMode Getter
	SetDragMode(value types.TDragMode)                 // property DragMode Setter
	MouseCapture() bool                                // property MouseCapture Getter
	SetMouseCapture(value bool)                        // property MouseCapture Setter
	ParentBackground() bool                            // property ParentBackground Getter
	SetParentBackground(value bool)                    // property ParentBackground Setter
	ParentColor() bool                                 // property ParentColor Getter
	SetParentColor(value bool)                         // property ParentColor Setter
	ParentFont() bool                                  // property ParentFont Getter
	SetParentFont(value bool)                          // property ParentFont Setter
	ParentShowHint() bool                              // property ParentShowHint Getter
	SetParentShowHint(value bool)                      // property ParentShowHint Setter
	SessionProperties() string                         // property SessionProperties Getter
	SetSessionProperties(value string)                 // property SessionProperties Setter
	Text() string                                      // property Text Getter
	SetText(value string)                              // property Text Setter
	SetOnPaint(fn TNotifyEvent)                        // property event
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnContextPopup(fn TContextPopupEvent)           // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnTripleClick(fn TNotifyEvent)                  // property event
	SetOnQuadClick(fn TNotifyEvent)                    // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnEndDock(fn TEndDragEvent)                     // property event
	SetOnEndDrag(fn TEndDragEvent)                     // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)                  // property event
	SetOnTextChanged(fn TCustomGraphicNotify)          // property event
	SetOnWndProc(fn TWndMethod)                        // property event
	SetOnDestroy(fn TCustomGraphicNotify)              // property event
}

type TCustomGraphicControl struct {
	TGraphicControl
}

func (m *TCustomGraphicControl) InheritedWndProc(theMessage *types.TLMessage) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(theMessage)))
}

func (m *TCustomGraphicControl) ActionLink() IControlActionLink {
	if !m.IsValid() {
		return nil
	}
	r := customGraphicControlAPI().SysCallN(2, 0, m.Instance())
	return AsControlActionLink(r)
}

func (m *TCustomGraphicControl) SetActionLink(value IControlActionLink) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomGraphicControl) DesktopFont() bool {
	if !m.IsValid() {
		return false
	}
	r := customGraphicControlAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGraphicControl) SetDesktopFont(value bool) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomGraphicControl) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := customGraphicControlAPI().SysCallN(4, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCustomGraphicControl) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomGraphicControl) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := customGraphicControlAPI().SysCallN(5, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TCustomGraphicControl) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomGraphicControl) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := customGraphicControlAPI().SysCallN(6, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCustomGraphicControl) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomGraphicControl) MouseCapture() bool {
	if !m.IsValid() {
		return false
	}
	r := customGraphicControlAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGraphicControl) SetMouseCapture(value bool) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomGraphicControl) ParentBackground() bool {
	if !m.IsValid() {
		return false
	}
	r := customGraphicControlAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGraphicControl) SetParentBackground(value bool) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomGraphicControl) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := customGraphicControlAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGraphicControl) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomGraphicControl) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := customGraphicControlAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGraphicControl) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomGraphicControl) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := customGraphicControlAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGraphicControl) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomGraphicControl) SessionProperties() string {
	if !m.IsValid() {
		return ""
	}
	r := customGraphicControlAPI().SysCallN(12, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomGraphicControl) SetSessionProperties(value string) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomGraphicControl) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := customGraphicControlAPI().SysCallN(13, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomGraphicControl) SetText(value string) {
	if !m.IsValid() {
		return
	}
	customGraphicControlAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomGraphicControl) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 15, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 16, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnTripleClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnQuadClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 20, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 21, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 22, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 23, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 24, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 25, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 26, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 27, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 29, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 30, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 31, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 32, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 33, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 34, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 35, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 36, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 37, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnTextChanged(fn TCustomGraphicNotify) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomGraphicNotify(fn)
	base.SetEvent(m, 38, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnWndProc(fn TWndMethod) {
	if !m.IsValid() {
		return
	}
	cb := makeTWndMethod(fn)
	base.SetEvent(m, 39, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomGraphicControl) SetOnDestroy(fn TCustomGraphicNotify) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomGraphicNotify(fn)
	base.SetEvent(m, 40, customGraphicControlAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomGraphicControl class constructor
func NewCustomGraphicControl(owner IComponent) ICustomGraphicControl {
	r := customGraphicControlAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomGraphicControl(r)
}

func TCustomGraphicControlClass() types.TClass {
	r := customGraphicControlAPI().SysCallN(41)
	return types.TClass(r)
}

var (
	customGraphicControlOnce   base.Once
	customGraphicControlImport *imports.Imports = nil
)

func customGraphicControlAPI() *imports.Imports {
	customGraphicControlOnce.Do(func() {
		customGraphicControlImport = api.NewDefaultImports()
		customGraphicControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomGraphicControl_Create", 0), // constructor NewCustomGraphicControl
			/* 1 */ imports.NewTable("TCustomGraphicControl_InheritedWndProc", 0), // procedure InheritedWndProc
			/* 2 */ imports.NewTable("TCustomGraphicControl_ActionLink", 0), // property ActionLink
			/* 3 */ imports.NewTable("TCustomGraphicControl_DesktopFont", 0), // property DesktopFont
			/* 4 */ imports.NewTable("TCustomGraphicControl_DragCursor", 0), // property DragCursor
			/* 5 */ imports.NewTable("TCustomGraphicControl_DragKind", 0), // property DragKind
			/* 6 */ imports.NewTable("TCustomGraphicControl_DragMode", 0), // property DragMode
			/* 7 */ imports.NewTable("TCustomGraphicControl_MouseCapture", 0), // property MouseCapture
			/* 8 */ imports.NewTable("TCustomGraphicControl_ParentBackground", 0), // property ParentBackground
			/* 9 */ imports.NewTable("TCustomGraphicControl_ParentColor", 0), // property ParentColor
			/* 10 */ imports.NewTable("TCustomGraphicControl_ParentFont", 0), // property ParentFont
			/* 11 */ imports.NewTable("TCustomGraphicControl_ParentShowHint", 0), // property ParentShowHint
			/* 12 */ imports.NewTable("TCustomGraphicControl_SessionProperties", 0), // property SessionProperties
			/* 13 */ imports.NewTable("TCustomGraphicControl_Text", 0), // property Text
			/* 14 */ imports.NewTable("TCustomGraphicControl_OnPaint", 0), // event OnPaint
			/* 15 */ imports.NewTable("TCustomGraphicControl_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 16 */ imports.NewTable("TCustomGraphicControl_OnContextPopup", 0), // event OnContextPopup
			/* 17 */ imports.NewTable("TCustomGraphicControl_OnDblClick", 0), // event OnDblClick
			/* 18 */ imports.NewTable("TCustomGraphicControl_OnTripleClick", 0), // event OnTripleClick
			/* 19 */ imports.NewTable("TCustomGraphicControl_OnQuadClick", 0), // event OnQuadClick
			/* 20 */ imports.NewTable("TCustomGraphicControl_OnDragDrop", 0), // event OnDragDrop
			/* 21 */ imports.NewTable("TCustomGraphicControl_OnDragOver", 0), // event OnDragOver
			/* 22 */ imports.NewTable("TCustomGraphicControl_OnEndDock", 0), // event OnEndDock
			/* 23 */ imports.NewTable("TCustomGraphicControl_OnEndDrag", 0), // event OnEndDrag
			/* 24 */ imports.NewTable("TCustomGraphicControl_OnMouseDown", 0), // event OnMouseDown
			/* 25 */ imports.NewTable("TCustomGraphicControl_OnMouseMove", 0), // event OnMouseMove
			/* 26 */ imports.NewTable("TCustomGraphicControl_OnMouseUp", 0), // event OnMouseUp
			/* 27 */ imports.NewTable("TCustomGraphicControl_OnMouseEnter", 0), // event OnMouseEnter
			/* 28 */ imports.NewTable("TCustomGraphicControl_OnMouseLeave", 0), // event OnMouseLeave
			/* 29 */ imports.NewTable("TCustomGraphicControl_OnMouseWheel", 0), // event OnMouseWheel
			/* 30 */ imports.NewTable("TCustomGraphicControl_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 31 */ imports.NewTable("TCustomGraphicControl_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 32 */ imports.NewTable("TCustomGraphicControl_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 33 */ imports.NewTable("TCustomGraphicControl_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 34 */ imports.NewTable("TCustomGraphicControl_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 35 */ imports.NewTable("TCustomGraphicControl_OnStartDock", 0), // event OnStartDock
			/* 36 */ imports.NewTable("TCustomGraphicControl_OnStartDrag", 0), // event OnStartDrag
			/* 37 */ imports.NewTable("TCustomGraphicControl_OnEditingDone", 0), // event OnEditingDone
			/* 38 */ imports.NewTable("TCustomGraphicControl_OnTextChanged", 0), // event OnTextChanged
			/* 39 */ imports.NewTable("TCustomGraphicControl_OnWndProc", 0), // event OnWndProc
			/* 40 */ imports.NewTable("TCustomGraphicControl_OnDestroy", 0), // event OnDestroy
			/* 41 */ imports.NewTable("TCustomGraphicControl_TClass", 0), // function TCustomGraphicControlClass
		}
	})
	return customGraphicControlImport
}
