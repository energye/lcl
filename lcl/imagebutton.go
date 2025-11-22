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

// IImageButton Parent: IGraphicControl
type IImageButton interface {
	IGraphicControl
	Click()                                       // procedure
	CaptionToString() string                      // property Caption Getter
	SetCaptionToString(value string)              // property Caption Setter
	DragCursor() types.TCursor                    // property DragCursor Getter
	SetDragCursor(value types.TCursor)            // property DragCursor Setter
	DragKind() types.TDragKind                    // property DragKind Getter
	SetDragKind(value types.TDragKind)            // property DragKind Setter
	DragMode() types.TDragMode                    // property DragMode Getter
	SetDragMode(value types.TDragMode)            // property DragMode Setter
	ImageCount() int32                            // property ImageCount Getter
	SetImageCount(value int32)                    // property ImageCount Setter
	Orientation() types.TImageOrientation         // property Orientation Getter
	SetOrientation(value types.TImageOrientation) // property Orientation Setter
	ModalResult() types.TModalResult              // property ModalResult Getter
	SetModalResult(value types.TModalResult)      // property ModalResult Setter
	ParentShowHint() bool                         // property ParentShowHint Getter
	SetParentShowHint(value bool)                 // property ParentShowHint Setter
	ParentFont() bool                             // property ParentFont Getter
	SetParentFont(value bool)                     // property ParentFont Setter
	Picture() IPicture                            // property Picture Getter
	SetPicture(value IPicture)                    // property Picture Setter
	ShowCaption() bool                            // property ShowCaption Getter
	SetShowCaption(value bool)                    // property ShowCaption Setter
	Wordwarp() bool                               // property Wordwarp Getter
	SetWordwarp(value bool)                       // property Wordwarp Setter
	SetOnContextPopup(fn TContextPopupEvent)      // property event
	SetOnDblClick(fn TNotifyEvent)                // property event
	SetOnDragDrop(fn TDragDropEvent)              // property event
	SetOnDragOver(fn TDragOverEvent)              // property event
	SetOnEndDock(fn TEndDragEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                // property event
	SetOnMouseDown(fn TMouseEvent)                // property event
	SetOnMouseEnter(fn TNotifyEvent)              // property event
	SetOnMouseLeave(fn TNotifyEvent)              // property event
	SetOnMouseMove(fn TMouseMoveEvent)            // property event
	SetOnMouseUp(fn TMouseEvent)                  // property event
}

type TImageButton struct {
	TGraphicControl
}

func (m *TImageButton) Click() {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(1, m.Instance())
}

func (m *TImageButton) CaptionToString() string {
	if !m.IsValid() {
		return ""
	}
	r := imageButtonAPI().SysCallN(2, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TImageButton) SetCaptionToString(value string) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(2, 1, m.Instance(), api.PasStr(value))
}

func (m *TImageButton) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := imageButtonAPI().SysCallN(3, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TImageButton) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TImageButton) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := imageButtonAPI().SysCallN(4, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TImageButton) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TImageButton) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := imageButtonAPI().SysCallN(5, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TImageButton) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TImageButton) ImageCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := imageButtonAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TImageButton) SetImageCount(value int32) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TImageButton) Orientation() types.TImageOrientation {
	if !m.IsValid() {
		return 0
	}
	r := imageButtonAPI().SysCallN(7, 0, m.Instance())
	return types.TImageOrientation(r)
}

func (m *TImageButton) SetOrientation(value types.TImageOrientation) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TImageButton) ModalResult() types.TModalResult {
	if !m.IsValid() {
		return 0
	}
	r := imageButtonAPI().SysCallN(8, 0, m.Instance())
	return types.TModalResult(r)
}

func (m *TImageButton) SetModalResult(value types.TModalResult) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TImageButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := imageButtonAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TImageButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TImageButton) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := imageButtonAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TImageButton) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TImageButton) Picture() IPicture {
	if !m.IsValid() {
		return nil
	}
	r := imageButtonAPI().SysCallN(11, 0, m.Instance())
	return AsPicture(r)
}

func (m *TImageButton) SetPicture(value IPicture) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TImageButton) ShowCaption() bool {
	if !m.IsValid() {
		return false
	}
	r := imageButtonAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TImageButton) SetShowCaption(value bool) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TImageButton) Wordwarp() bool {
	if !m.IsValid() {
		return false
	}
	r := imageButtonAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TImageButton) SetWordwarp(value bool) {
	if !m.IsValid() {
		return
	}
	imageButtonAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TImageButton) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 14, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 16, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 17, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 18, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 19, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 20, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 21, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 22, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 23, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImageButton) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 24, imageButtonAPI(), api.MakeEventDataPtr(cb))
}

// NewImageButton class constructor
func NewImageButton(owner IComponent) IImageButton {
	r := imageButtonAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsImageButton(r)
}

func TImageButtonClass() types.TClass {
	r := imageButtonAPI().SysCallN(25)
	return types.TClass(r)
}

var (
	imageButtonOnce   base.Once
	imageButtonImport *imports.Imports = nil
)

func imageButtonAPI() *imports.Imports {
	imageButtonOnce.Do(func() {
		imageButtonImport = api.NewDefaultImports()
		imageButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TImageButton_Create", 0), // constructor NewImageButton
			/* 1 */ imports.NewTable("TImageButton_Click", 0), // procedure Click
			/* 2 */ imports.NewTable("TImageButton_CaptionToString", 0), // property CaptionToString
			/* 3 */ imports.NewTable("TImageButton_DragCursor", 0), // property DragCursor
			/* 4 */ imports.NewTable("TImageButton_DragKind", 0), // property DragKind
			/* 5 */ imports.NewTable("TImageButton_DragMode", 0), // property DragMode
			/* 6 */ imports.NewTable("TImageButton_ImageCount", 0), // property ImageCount
			/* 7 */ imports.NewTable("TImageButton_Orientation", 0), // property Orientation
			/* 8 */ imports.NewTable("TImageButton_ModalResult", 0), // property ModalResult
			/* 9 */ imports.NewTable("TImageButton_ParentShowHint", 0), // property ParentShowHint
			/* 10 */ imports.NewTable("TImageButton_ParentFont", 0), // property ParentFont
			/* 11 */ imports.NewTable("TImageButton_Picture", 0), // property Picture
			/* 12 */ imports.NewTable("TImageButton_ShowCaption", 0), // property ShowCaption
			/* 13 */ imports.NewTable("TImageButton_Wordwarp", 0), // property Wordwarp
			/* 14 */ imports.NewTable("TImageButton_OnContextPopup", 0), // event OnContextPopup
			/* 15 */ imports.NewTable("TImageButton_OnDblClick", 0), // event OnDblClick
			/* 16 */ imports.NewTable("TImageButton_OnDragDrop", 0), // event OnDragDrop
			/* 17 */ imports.NewTable("TImageButton_OnDragOver", 0), // event OnDragOver
			/* 18 */ imports.NewTable("TImageButton_OnEndDock", 0), // event OnEndDock
			/* 19 */ imports.NewTable("TImageButton_OnEndDrag", 0), // event OnEndDrag
			/* 20 */ imports.NewTable("TImageButton_OnMouseDown", 0), // event OnMouseDown
			/* 21 */ imports.NewTable("TImageButton_OnMouseEnter", 0), // event OnMouseEnter
			/* 22 */ imports.NewTable("TImageButton_OnMouseLeave", 0), // event OnMouseLeave
			/* 23 */ imports.NewTable("TImageButton_OnMouseMove", 0), // event OnMouseMove
			/* 24 */ imports.NewTable("TImageButton_OnMouseUp", 0), // event OnMouseUp
			/* 25 */ imports.NewTable("TImageButton_TClass", 0), // function TImageButtonClass
		}
	})
	return imageButtonImport
}
