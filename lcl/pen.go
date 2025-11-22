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

// IPen Parent: IFPCustomPen
type IPen interface {
	IFPCustomPen
	GetPattern() IUint32ArrayWrap                         // function
	SetPatternWithPenPattern(pattern IUint32ArrayWrap)    // procedure
	Color() types.TColor                                  // property Color Getter
	SetColor(value types.TColor)                          // property Color Setter
	Cosmetic() bool                                       // property Cosmetic Getter
	SetCosmetic(value bool)                               // property Cosmetic Setter
	EndCapToPenEndCap() types.TPenEndCap                  // property EndCap Getter
	SetEndCapToPenEndCap(value types.TPenEndCap)          // property EndCap Setter
	JoinStyleToPenJoinStyle() types.TPenJoinStyle         // property JoinStyle Getter
	SetJoinStyleToPenJoinStyle(value types.TPenJoinStyle) // property JoinStyle Setter
}

type TPen struct {
	TFPCustomPen
}

func (m *TPen) GetPattern() IUint32ArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := penAPI().SysCallN(1, m.Instance())
	return AsUint32ArrayWrap(r)
}

func (m *TPen) SetPatternWithPenPattern(pattern IUint32ArrayWrap) {
	if !m.IsValid() {
		return
	}
	penAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(pattern))
}

func (m *TPen) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := penAPI().SysCallN(3, 0, m.Instance())
	return types.TColor(r)
}

func (m *TPen) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	penAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TPen) Cosmetic() bool {
	if !m.IsValid() {
		return false
	}
	r := penAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPen) SetCosmetic(value bool) {
	if !m.IsValid() {
		return
	}
	penAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TPen) EndCapToPenEndCap() types.TPenEndCap {
	if !m.IsValid() {
		return 0
	}
	r := penAPI().SysCallN(5, 0, m.Instance())
	return types.TPenEndCap(r)
}

func (m *TPen) SetEndCapToPenEndCap(value types.TPenEndCap) {
	if !m.IsValid() {
		return
	}
	penAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TPen) JoinStyleToPenJoinStyle() types.TPenJoinStyle {
	if !m.IsValid() {
		return 0
	}
	r := penAPI().SysCallN(6, 0, m.Instance())
	return types.TPenJoinStyle(r)
}

func (m *TPen) SetJoinStyleToPenJoinStyle(value types.TPenJoinStyle) {
	if !m.IsValid() {
		return
	}
	penAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

// NewPen class constructor
func NewPen() IPen {
	r := penAPI().SysCallN(0)
	return AsPen(r)
}

var (
	penOnce   base.Once
	penImport *imports.Imports = nil
)

func penAPI() *imports.Imports {
	penOnce.Do(func() {
		penImport = api.NewDefaultImports()
		penImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPen_Create", 0), // constructor NewPen
			/* 1 */ imports.NewTable("TPen_GetPattern", 0), // function GetPattern
			/* 2 */ imports.NewTable("TPen_SetPatternWithPenPattern", 0), // procedure SetPatternWithPenPattern
			/* 3 */ imports.NewTable("TPen_Color", 0), // property Color
			/* 4 */ imports.NewTable("TPen_Cosmetic", 0), // property Cosmetic
			/* 5 */ imports.NewTable("TPen_EndCapToPenEndCap", 0), // property EndCapToPenEndCap
			/* 6 */ imports.NewTable("TPen_JoinStyleToPenJoinStyle", 0), // property JoinStyleToPenJoinStyle
		}
	})
	return penImport
}
