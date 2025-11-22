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

// IVTDragImage Parent: IObject
type IVTDragImage interface {
	IObject
	DragTo(P types.TPoint, forceRepaint bool) bool                                                                                // function
	GetDragImageRect() types.TRect                                                                                                // function
	WillMove(P types.TPoint) bool                                                                                                 // function
	EndDrag()                                                                                                                     // procedure
	HideDragImage()                                                                                                               // procedure
	RecaptureBackground(tree IBaseVirtualTree, R types.TRect, visibleRegion types.HRGN, captureNCArea bool, reshowDragImage bool) // procedure
	ShowDragImage()                                                                                                               // procedure
	ColorKey() types.TColor                                                                                                       // property ColorKey Getter
	SetColorKey(value types.TColor)                                                                                               // property ColorKey Setter
	Fade() bool                                                                                                                   // property Fade Getter
	SetFade(value bool)                                                                                                           // property Fade Setter
	MoveRestriction() types.TVTDragMoveRestriction                                                                                // property MoveRestriction Getter
	SetMoveRestriction(value types.TVTDragMoveRestriction)                                                                        // property MoveRestriction Setter
	PostBlendBias() types.TVTBias                                                                                                 // property PostBlendBias Getter
	SetPostBlendBias(value types.TVTBias)                                                                                         // property PostBlendBias Setter
	PreBlendBias() types.TVTBias                                                                                                  // property PreBlendBias Getter
	SetPreBlendBias(value types.TVTBias)                                                                                          // property PreBlendBias Setter
	Transparency() types.TVTTransparency                                                                                          // property Transparency Getter
	SetTransparency(value types.TVTTransparency)                                                                                  // property Transparency Setter
	Visible() bool                                                                                                                // property Visible Getter
}

type TVTDragImage struct {
	TObject
}

func (m *TVTDragImage) DragTo(P types.TPoint, forceRepaint bool) bool {
	if !m.IsValid() {
		return false
	}
	r := vTDragImageAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&P)), api.PasBool(forceRepaint))
	return api.GoBool(r)
}

func (m *TVTDragImage) GetDragImageRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TVTDragImage) WillMove(P types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := vTDragImageAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return api.GoBool(r)
}

func (m *TVTDragImage) EndDrag() {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(4, m.Instance())
}

func (m *TVTDragImage) HideDragImage() {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(5, m.Instance())
}

func (m *TVTDragImage) RecaptureBackground(tree IBaseVirtualTree, R types.TRect, visibleRegion types.HRGN, captureNCArea bool, reshowDragImage bool) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(tree), uintptr(base.UnsafePointer(&R)), uintptr(visibleRegion), api.PasBool(captureNCArea), api.PasBool(reshowDragImage))
}

func (m *TVTDragImage) ShowDragImage() {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(7, m.Instance())
}

func (m *TVTDragImage) ColorKey() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := vTDragImageAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVTDragImage) SetColorKey(value types.TColor) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TVTDragImage) Fade() bool {
	if !m.IsValid() {
		return false
	}
	r := vTDragImageAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TVTDragImage) SetFade(value bool) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TVTDragImage) MoveRestriction() types.TVTDragMoveRestriction {
	if !m.IsValid() {
		return 0
	}
	r := vTDragImageAPI().SysCallN(10, 0, m.Instance())
	return types.TVTDragMoveRestriction(r)
}

func (m *TVTDragImage) SetMoveRestriction(value types.TVTDragMoveRestriction) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TVTDragImage) PostBlendBias() types.TVTBias {
	if !m.IsValid() {
		return 0
	}
	r := vTDragImageAPI().SysCallN(11, 0, m.Instance())
	return types.TVTBias(r)
}

func (m *TVTDragImage) SetPostBlendBias(value types.TVTBias) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TVTDragImage) PreBlendBias() types.TVTBias {
	if !m.IsValid() {
		return 0
	}
	r := vTDragImageAPI().SysCallN(12, 0, m.Instance())
	return types.TVTBias(r)
}

func (m *TVTDragImage) SetPreBlendBias(value types.TVTBias) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TVTDragImage) Transparency() types.TVTTransparency {
	if !m.IsValid() {
		return 0
	}
	r := vTDragImageAPI().SysCallN(13, 0, m.Instance())
	return types.TVTTransparency(r)
}

func (m *TVTDragImage) SetTransparency(value types.TVTTransparency) {
	if !m.IsValid() {
		return
	}
	vTDragImageAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TVTDragImage) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := vTDragImageAPI().SysCallN(14, m.Instance())
	return api.GoBool(r)
}

// NewVTDragImage class constructor
func NewVTDragImage(owner IBaseVirtualTree) IVTDragImage {
	r := vTDragImageAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsVTDragImage(r)
}

var (
	vTDragImageOnce   base.Once
	vTDragImageImport *imports.Imports = nil
)

func vTDragImageAPI() *imports.Imports {
	vTDragImageOnce.Do(func() {
		vTDragImageImport = api.NewDefaultImports()
		vTDragImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTDragImage_Create", 0), // constructor NewVTDragImage
			/* 1 */ imports.NewTable("TVTDragImage_DragTo", 0), // function DragTo
			/* 2 */ imports.NewTable("TVTDragImage_GetDragImageRect", 0), // function GetDragImageRect
			/* 3 */ imports.NewTable("TVTDragImage_WillMove", 0), // function WillMove
			/* 4 */ imports.NewTable("TVTDragImage_EndDrag", 0), // procedure EndDrag
			/* 5 */ imports.NewTable("TVTDragImage_HideDragImage", 0), // procedure HideDragImage
			/* 6 */ imports.NewTable("TVTDragImage_RecaptureBackground", 0), // procedure RecaptureBackground
			/* 7 */ imports.NewTable("TVTDragImage_ShowDragImage", 0), // procedure ShowDragImage
			/* 8 */ imports.NewTable("TVTDragImage_ColorKey", 0), // property ColorKey
			/* 9 */ imports.NewTable("TVTDragImage_Fade", 0), // property Fade
			/* 10 */ imports.NewTable("TVTDragImage_MoveRestriction", 0), // property MoveRestriction
			/* 11 */ imports.NewTable("TVTDragImage_PostBlendBias", 0), // property PostBlendBias
			/* 12 */ imports.NewTable("TVTDragImage_PreBlendBias", 0), // property PreBlendBias
			/* 13 */ imports.NewTable("TVTDragImage_Transparency", 0), // property Transparency
			/* 14 */ imports.NewTable("TVTDragImage_Visible", 0), // property Visible
		}
	})
	return vTDragImageImport
}
