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

// IFPCustomPen Parent: IFPCanvasHelper
type IFPCustomPen interface {
	IFPCanvasHelper
	CopyPen() IFPCustomPen // function
	// Style
	//  Creates a copy of the pen with all properties the same, but not allocated
	Style() types.TFPPenStyle                                 // property Style Getter
	SetStyle(value types.TFPPenStyle)                         // property Style Setter
	Width() int32                                             // property Width Getter
	SetWidth(value int32)                                     // property Width Setter
	Mode() types.TFPPenMode                                   // property Mode Getter
	SetMode(value types.TFPPenMode)                           // property Mode Setter
	Pattern() uint32                                          // property Pattern Getter
	SetPatternToLongword(value uint32)                        // property Pattern Setter
	EndCapToFPPenEndCap() types.TFPPenEndCap                  // property EndCap Getter
	SetEndCapToFPPenEndCap(value types.TFPPenEndCap)          // property EndCap Setter
	JoinStyleToFPPenJoinStyle() types.TFPPenJoinStyle         // property JoinStyle Getter
	SetJoinStyleToFPPenJoinStyle(value types.TFPPenJoinStyle) // property JoinStyle Setter
}

type TFPCustomPen struct {
	TFPCanvasHelper
}

func (m *TFPCustomPen) CopyPen() IFPCustomPen {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomPenAPI().SysCallN(1, m.Instance())
	return AsFPCustomPen(r)
}

func (m *TFPCustomPen) Style() types.TFPPenStyle {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomPenAPI().SysCallN(2, 0, m.Instance())
	return types.TFPPenStyle(r)
}

func (m *TFPCustomPen) SetStyle(value types.TFPPenStyle) {
	if !m.IsValid() {
		return
	}
	fPCustomPenAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomPen) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomPenAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TFPCustomPen) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomPenAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomPen) Mode() types.TFPPenMode {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomPenAPI().SysCallN(4, 0, m.Instance())
	return types.TFPPenMode(r)
}

func (m *TFPCustomPen) SetMode(value types.TFPPenMode) {
	if !m.IsValid() {
		return
	}
	fPCustomPenAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomPen) Pattern() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomPenAPI().SysCallN(5, 0, m.Instance())
	return uint32(r)
}

func (m *TFPCustomPen) SetPatternToLongword(value uint32) {
	if !m.IsValid() {
		return
	}
	fPCustomPenAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomPen) EndCapToFPPenEndCap() types.TFPPenEndCap {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomPenAPI().SysCallN(6, 0, m.Instance())
	return types.TFPPenEndCap(r)
}

func (m *TFPCustomPen) SetEndCapToFPPenEndCap(value types.TFPPenEndCap) {
	if !m.IsValid() {
		return
	}
	fPCustomPenAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomPen) JoinStyleToFPPenJoinStyle() types.TFPPenJoinStyle {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomPenAPI().SysCallN(7, 0, m.Instance())
	return types.TFPPenJoinStyle(r)
}

func (m *TFPCustomPen) SetJoinStyleToFPPenJoinStyle(value types.TFPPenJoinStyle) {
	if !m.IsValid() {
		return
	}
	fPCustomPenAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

// NewFPCustomPen class constructor
func NewFPCustomPen() IFPCustomPen {
	r := fPCustomPenAPI().SysCallN(0)
	return AsFPCustomPen(r)
}

var (
	fPCustomPenOnce   base.Once
	fPCustomPenImport *imports.Imports = nil
)

func fPCustomPenAPI() *imports.Imports {
	fPCustomPenOnce.Do(func() {
		fPCustomPenImport = api.NewDefaultImports()
		fPCustomPenImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomPen_Create", 0), // constructor NewFPCustomPen
			/* 1 */ imports.NewTable("TFPCustomPen_CopyPen", 0), // function CopyPen
			/* 2 */ imports.NewTable("TFPCustomPen_Style", 0), // property Style
			/* 3 */ imports.NewTable("TFPCustomPen_Width", 0), // property Width
			/* 4 */ imports.NewTable("TFPCustomPen_Mode", 0), // property Mode
			/* 5 */ imports.NewTable("TFPCustomPen_Pattern", 0), // property Pattern
			/* 6 */ imports.NewTable("TFPCustomPen_EndCapToFPPenEndCap", 0), // property EndCapToFPPenEndCap
			/* 7 */ imports.NewTable("TFPCustomPen_JoinStyleToFPPenJoinStyle", 0), // property JoinStyleToFPPenJoinStyle
		}
	})
	return fPCustomPenImport
}
