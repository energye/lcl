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

// IControlScrollBar Parent: IPersistent
type IControlScrollBar interface {
	IPersistent
	IsScrollBarVisible() bool               // function
	ScrollPos() int32                       // function
	GetOtherScrollBar() IControlScrollBar   // function
	ControlSize() int32                     // function
	ClientSize() int32                      // function
	ClientSizeWithBar() int32               // function
	ClientSizeWithoutBar() int32            // function
	Kind() types.TScrollBarKind             // property Kind Getter
	Size() int32                            // property Size Getter
	Increment() types.TScrollBarInc         // property Increment Getter
	SetIncrement(value types.TScrollBarInc) // property Increment Setter
	Page() types.TScrollBarInc              // property Page Getter
	SetPage(value types.TScrollBarInc)      // property Page Setter
	Smooth() bool                           // property Smooth Getter
	SetSmooth(value bool)                   // property Smooth Setter
	Position() int32                        // property Position Getter
	SetPosition(value int32)                // property Position Setter
	Range() int32                           // property Range Getter
	SetRange(value int32)                   // property Range Setter
	Tracking() bool                         // property Tracking Getter
	SetTracking(value bool)                 // property Tracking Setter
	Visible() bool                          // property Visible Getter
	SetVisible(value bool)                  // property Visible Setter
}

type TControlScrollBar struct {
	TPersistent
}

func (m *TControlScrollBar) IsScrollBarVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := controlScrollBarAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TControlScrollBar) ScrollPos() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) GetOtherScrollBar() IControlScrollBar {
	if !m.IsValid() {
		return nil
	}
	r := controlScrollBarAPI().SysCallN(3, m.Instance())
	return AsControlScrollBar(r)
}

func (m *TControlScrollBar) ControlSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) ClientSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) ClientSizeWithBar() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) ClientSizeWithoutBar() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) Kind() types.TScrollBarKind {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(8, m.Instance())
	return types.TScrollBarKind(r)
}

func (m *TControlScrollBar) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) Increment() types.TScrollBarInc {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(10, 0, m.Instance())
	return types.TScrollBarInc(r)
}

func (m *TControlScrollBar) SetIncrement(value types.TScrollBarInc) {
	if !m.IsValid() {
		return
	}
	controlScrollBarAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TControlScrollBar) Page() types.TScrollBarInc {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(11, 0, m.Instance())
	return types.TScrollBarInc(r)
}

func (m *TControlScrollBar) SetPage(value types.TScrollBarInc) {
	if !m.IsValid() {
		return
	}
	controlScrollBarAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TControlScrollBar) Smooth() bool {
	if !m.IsValid() {
		return false
	}
	r := controlScrollBarAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControlScrollBar) SetSmooth(value bool) {
	if !m.IsValid() {
		return
	}
	controlScrollBarAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TControlScrollBar) Position() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) SetPosition(value int32) {
	if !m.IsValid() {
		return
	}
	controlScrollBarAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TControlScrollBar) Range() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlScrollBarAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TControlScrollBar) SetRange(value int32) {
	if !m.IsValid() {
		return
	}
	controlScrollBarAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TControlScrollBar) Tracking() bool {
	if !m.IsValid() {
		return false
	}
	r := controlScrollBarAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControlScrollBar) SetTracking(value bool) {
	if !m.IsValid() {
		return
	}
	controlScrollBarAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TControlScrollBar) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := controlScrollBarAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControlScrollBar) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	controlScrollBarAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

// NewControlScrollBar class constructor
func NewControlScrollBar(control IWinControl, kind types.TScrollBarKind) IControlScrollBar {
	r := controlScrollBarAPI().SysCallN(0, base.GetObjectUintptr(control), uintptr(kind))
	return AsControlScrollBar(r)
}

var (
	controlScrollBarOnce   base.Once
	controlScrollBarImport *imports.Imports = nil
)

func controlScrollBarAPI() *imports.Imports {
	controlScrollBarOnce.Do(func() {
		controlScrollBarImport = api.NewDefaultImports()
		controlScrollBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TControlScrollBar_Create", 0), // constructor NewControlScrollBar
			/* 1 */ imports.NewTable("TControlScrollBar_IsScrollBarVisible", 0), // function IsScrollBarVisible
			/* 2 */ imports.NewTable("TControlScrollBar_ScrollPos", 0), // function ScrollPos
			/* 3 */ imports.NewTable("TControlScrollBar_GetOtherScrollBar", 0), // function GetOtherScrollBar
			/* 4 */ imports.NewTable("TControlScrollBar_ControlSize", 0), // function ControlSize
			/* 5 */ imports.NewTable("TControlScrollBar_ClientSize", 0), // function ClientSize
			/* 6 */ imports.NewTable("TControlScrollBar_ClientSizeWithBar", 0), // function ClientSizeWithBar
			/* 7 */ imports.NewTable("TControlScrollBar_ClientSizeWithoutBar", 0), // function ClientSizeWithoutBar
			/* 8 */ imports.NewTable("TControlScrollBar_Kind", 0), // property Kind
			/* 9 */ imports.NewTable("TControlScrollBar_Size", 0), // property Size
			/* 10 */ imports.NewTable("TControlScrollBar_Increment", 0), // property Increment
			/* 11 */ imports.NewTable("TControlScrollBar_Page", 0), // property Page
			/* 12 */ imports.NewTable("TControlScrollBar_Smooth", 0), // property Smooth
			/* 13 */ imports.NewTable("TControlScrollBar_Position", 0), // property Position
			/* 14 */ imports.NewTable("TControlScrollBar_Range", 0), // property Range
			/* 15 */ imports.NewTable("TControlScrollBar_Tracking", 0), // property Tracking
			/* 16 */ imports.NewTable("TControlScrollBar_Visible", 0), // property Visible
		}
	})
	return controlScrollBarImport
}
