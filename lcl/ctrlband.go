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

// ICtrlBand Parent: IObject
type ICtrlBand interface {
	IObject
	BandRect() types.TRect         // property BandRect Getter
	SetBandRect(value types.TRect) // property BandRect Setter
	Bottom() int32                 // property Bottom Getter
	Control() IControl             // property Control Getter
	SetControl(value IControl)     // property Control Setter
	ControlHeight() int32          // property ControlHeight Getter
	SetControlHeight(value int32)  // property ControlHeight Setter
	ControlLeft() int32            // property ControlLeft Getter
	SetControlLeft(value int32)    // property ControlLeft Setter
	ControlTop() int32             // property ControlTop Getter
	SetControlTop(value int32)     // property ControlTop Setter
	ControlWidth() int32           // property ControlWidth Getter
	SetControlWidth(value int32)   // property ControlWidth Setter
	ControlVisible() bool          // property ControlVisible Getter
	SetControlVisible(value bool)  // property ControlVisible Setter
	Height() int32                 // property Height Getter
	SetHeight(value int32)         // property Height Setter
	InitLeft() int32               // property InitLeft Getter
	SetInitLeft(value int32)       // property InitLeft Setter
	InitTop() int32                // property InitTop Getter
	SetInitTop(value int32)        // property InitTop Setter
	Left() int32                   // property Left Getter
	SetLeft(value int32)           // property Left Setter
	Right() int32                  // property Right Getter
	SetRight(value int32)          // property Right Setter
	Top() int32                    // property Top Getter
	SetTop(value int32)            // property Top Setter
	Visible() bool                 // property Visible Getter
	SetVisible(value bool)         // property Visible Setter
	Width() int32                  // property Width Getter
	SetWidth(value int32)          // property Width Setter
}

type TCtrlBand struct {
	TObject
}

func (m *TCtrlBand) BandRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(0, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCtrlBand) SetBandRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(0, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCtrlBand) Bottom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) Control() IControl {
	if !m.IsValid() {
		return nil
	}
	r := ctrlBandAPI().SysCallN(2, 0, m.Instance())
	return AsControl(r)
}

func (m *TCtrlBand) SetControl(value IControl) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCtrlBand) ControlHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetControlHeight(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) ControlLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetControlLeft(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) ControlTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetControlTop(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) ControlWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetControlWidth(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) ControlVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := ctrlBandAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCtrlBand) SetControlVisible(value bool) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCtrlBand) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) InitLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetInitLeft(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) InitTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetInitTop(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetLeft(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) Right() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetRight(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetTop(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCtrlBand) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := ctrlBandAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCtrlBand) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCtrlBand) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := ctrlBandAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TCtrlBand) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	ctrlBandAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

var (
	ctrlBandOnce   base.Once
	ctrlBandImport *imports.Imports = nil
)

func ctrlBandAPI() *imports.Imports {
	ctrlBandOnce.Do(func() {
		ctrlBandImport = api.NewDefaultImports()
		ctrlBandImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCtrlBand_BandRect", 0), // property BandRect
			/* 1 */ imports.NewTable("TCtrlBand_Bottom", 0), // property Bottom
			/* 2 */ imports.NewTable("TCtrlBand_Control", 0), // property Control
			/* 3 */ imports.NewTable("TCtrlBand_ControlHeight", 0), // property ControlHeight
			/* 4 */ imports.NewTable("TCtrlBand_ControlLeft", 0), // property ControlLeft
			/* 5 */ imports.NewTable("TCtrlBand_ControlTop", 0), // property ControlTop
			/* 6 */ imports.NewTable("TCtrlBand_ControlWidth", 0), // property ControlWidth
			/* 7 */ imports.NewTable("TCtrlBand_ControlVisible", 0), // property ControlVisible
			/* 8 */ imports.NewTable("TCtrlBand_Height", 0), // property Height
			/* 9 */ imports.NewTable("TCtrlBand_InitLeft", 0), // property InitLeft
			/* 10 */ imports.NewTable("TCtrlBand_InitTop", 0), // property InitTop
			/* 11 */ imports.NewTable("TCtrlBand_Left", 0), // property Left
			/* 12 */ imports.NewTable("TCtrlBand_Right", 0), // property Right
			/* 13 */ imports.NewTable("TCtrlBand_Top", 0), // property Top
			/* 14 */ imports.NewTable("TCtrlBand_Visible", 0), // property Visible
			/* 15 */ imports.NewTable("TCtrlBand_Width", 0), // property Width
		}
	})
	return ctrlBandImport
}
