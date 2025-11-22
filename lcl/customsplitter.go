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

// ICustomSplitter Parent: ICustomControl
type ICustomSplitter interface {
	ICustomControl
	GetOtherResizeControl() IControl                         // function
	GetSplitterPosition() int32                              // function
	AnchorSplitter(kind types.TAnchorKind, control IControl) // procedure
	MoveSplitter(offset int32)                               // procedure
	SetSplitterPosition(newPosition int32)                   // procedure
	ResizeControl() IControl                                 // property ResizeControl Getter
	SetResizeControl(value IControl)                         // property ResizeControl Setter
	AutoSnap() bool                                          // property AutoSnap Getter
	SetAutoSnap(value bool)                                  // property AutoSnap Setter
	Beveled() bool                                           // property Beveled Getter
	SetBeveled(value bool)                                   // property Beveled Setter
	MinSize() int32                                          // property MinSize Getter
	SetMinSize(value int32)                                  // property MinSize Setter
	ResizeAnchor() types.TAnchorKind                         // property ResizeAnchor Getter
	SetResizeAnchor(value types.TAnchorKind)                 // property ResizeAnchor Setter
	ResizeStyle() types.TResizeStyle                         // property ResizeStyle Getter
	SetResizeStyle(value types.TResizeStyle)                 // property ResizeStyle Setter
	SetOnCanOffset(fn TCanOffsetEvent)                       // property event
	SetOnCanResize(fn TCanResizeEvent)                       // property event
	SetOnMoved(fn TNotifyEvent)                              // property event
}

type TCustomSplitter struct {
	TCustomControl
}

func (m *TCustomSplitter) GetOtherResizeControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := customSplitterAPI().SysCallN(1, m.Instance())
	return AsControl(r)
}

func (m *TCustomSplitter) GetSplitterPosition() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSplitterAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCustomSplitter) AnchorSplitter(kind types.TAnchorKind, control IControl) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(3, m.Instance(), uintptr(kind), base.GetObjectUintptr(control))
}

func (m *TCustomSplitter) MoveSplitter(offset int32) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(4, m.Instance(), uintptr(offset))
}

func (m *TCustomSplitter) SetSplitterPosition(newPosition int32) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(5, m.Instance(), uintptr(newPosition))
}

func (m *TCustomSplitter) ResizeControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := customSplitterAPI().SysCallN(6, 0, m.Instance())
	return AsControl(r)
}

func (m *TCustomSplitter) SetResizeControl(value IControl) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSplitter) AutoSnap() bool {
	if !m.IsValid() {
		return false
	}
	r := customSplitterAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSplitter) SetAutoSnap(value bool) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSplitter) Beveled() bool {
	if !m.IsValid() {
		return false
	}
	r := customSplitterAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSplitter) SetBeveled(value bool) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSplitter) MinSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSplitterAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSplitter) SetMinSize(value int32) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSplitter) ResizeAnchor() types.TAnchorKind {
	if !m.IsValid() {
		return 0
	}
	r := customSplitterAPI().SysCallN(10, 0, m.Instance())
	return types.TAnchorKind(r)
}

func (m *TCustomSplitter) SetResizeAnchor(value types.TAnchorKind) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSplitter) ResizeStyle() types.TResizeStyle {
	if !m.IsValid() {
		return 0
	}
	r := customSplitterAPI().SysCallN(11, 0, m.Instance())
	return types.TResizeStyle(r)
}

func (m *TCustomSplitter) SetResizeStyle(value types.TResizeStyle) {
	if !m.IsValid() {
		return
	}
	customSplitterAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSplitter) SetOnCanOffset(fn TCanOffsetEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCanOffsetEvent(fn)
	base.SetEvent(m, 12, customSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSplitter) SetOnCanResize(fn TCanResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCanResizeEvent(fn)
	base.SetEvent(m, 13, customSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSplitter) SetOnMoved(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, customSplitterAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomSplitter class constructor
func NewCustomSplitter(theOwner IComponent) ICustomSplitter {
	r := customSplitterAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomSplitter(r)
}

func TCustomSplitterClass() types.TClass {
	r := customSplitterAPI().SysCallN(15)
	return types.TClass(r)
}

var (
	customSplitterOnce   base.Once
	customSplitterImport *imports.Imports = nil
)

func customSplitterAPI() *imports.Imports {
	customSplitterOnce.Do(func() {
		customSplitterImport = api.NewDefaultImports()
		customSplitterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomSplitter_Create", 0), // constructor NewCustomSplitter
			/* 1 */ imports.NewTable("TCustomSplitter_GetOtherResizeControl", 0), // function GetOtherResizeControl
			/* 2 */ imports.NewTable("TCustomSplitter_GetSplitterPosition", 0), // function GetSplitterPosition
			/* 3 */ imports.NewTable("TCustomSplitter_AnchorSplitter", 0), // procedure AnchorSplitter
			/* 4 */ imports.NewTable("TCustomSplitter_MoveSplitter", 0), // procedure MoveSplitter
			/* 5 */ imports.NewTable("TCustomSplitter_SetSplitterPosition", 0), // procedure SetSplitterPosition
			/* 6 */ imports.NewTable("TCustomSplitter_ResizeControl", 0), // property ResizeControl
			/* 7 */ imports.NewTable("TCustomSplitter_AutoSnap", 0), // property AutoSnap
			/* 8 */ imports.NewTable("TCustomSplitter_Beveled", 0), // property Beveled
			/* 9 */ imports.NewTable("TCustomSplitter_MinSize", 0), // property MinSize
			/* 10 */ imports.NewTable("TCustomSplitter_ResizeAnchor", 0), // property ResizeAnchor
			/* 11 */ imports.NewTable("TCustomSplitter_ResizeStyle", 0), // property ResizeStyle
			/* 12 */ imports.NewTable("TCustomSplitter_OnCanOffset", 0), // event OnCanOffset
			/* 13 */ imports.NewTable("TCustomSplitter_OnCanResize", 0), // event OnCanResize
			/* 14 */ imports.NewTable("TCustomSplitter_OnMoved", 0), // event OnMoved
			/* 15 */ imports.NewTable("TCustomSplitter_TClass", 0), // function TCustomSplitterClass
		}
	})
	return customSplitterImport
}
