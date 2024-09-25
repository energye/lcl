//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	. "github.com/energye/lcl/types"
)

// ICustomSplitter Parent: ICustomControl
type ICustomSplitter interface {
	ICustomControl
	ResizeControl() IControl                            // property
	SetResizeControl(AValue IControl)                   // property
	AutoSnap() bool                                     // property
	SetAutoSnap(AValue bool)                            // property
	Beveled() bool                                      // property
	SetBeveled(AValue bool)                             // property
	MinSize() int32                                     // property
	SetMinSize(AValue int32)                            // property
	ResizeAnchor() TAnchorKind                          // property
	SetResizeAnchor(AValue TAnchorKind)                 // property
	ResizeStyle() TResizeStyle                          // property
	SetResizeStyle(AValue TResizeStyle)                 // property
	GetOtherResizeControl() IControl                    // function
	GetSplitterPosition() int32                         // function
	AnchorSplitter(Kind TAnchorKind, AControl IControl) // procedure
	MoveSplitter(Offset int32)                          // procedure
	SetSplitterPosition(NewPosition int32)              // procedure
	SetOnCanOffset(fn TCanOffsetEvent)                  // property event
	SetOnCanResize(fn TCanResizeEvent)                  // property event
	SetOnMoved(fn TNotifyEvent)                         // property event
}

// TCustomSplitter Parent: TCustomControl
type TCustomSplitter struct {
	TCustomControl
	canOffsetPtr uintptr
	canResizePtr uintptr
	movedPtr     uintptr
}

func NewCustomSplitter(TheOwner IComponent) ICustomSplitter {
	r1 := LCL().SysCallN(2281, GetObjectUintptr(TheOwner))
	return AsCustomSplitter(r1)
}

func (m *TCustomSplitter) ResizeControl() IControl {
	r1 := LCL().SysCallN(2287, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomSplitter) SetResizeControl(AValue IControl) {
	LCL().SysCallN(2287, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSplitter) AutoSnap() bool {
	r1 := LCL().SysCallN(2278, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetAutoSnap(AValue bool) {
	LCL().SysCallN(2278, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) Beveled() bool {
	r1 := LCL().SysCallN(2279, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetBeveled(AValue bool) {
	LCL().SysCallN(2279, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) MinSize() int32 {
	r1 := LCL().SysCallN(2284, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSplitter) SetMinSize(AValue int32) {
	LCL().SysCallN(2284, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeAnchor() TAnchorKind {
	r1 := LCL().SysCallN(2286, 0, m.Instance(), 0)
	return TAnchorKind(r1)
}

func (m *TCustomSplitter) SetResizeAnchor(AValue TAnchorKind) {
	LCL().SysCallN(2286, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeStyle() TResizeStyle {
	r1 := LCL().SysCallN(2288, 0, m.Instance(), 0)
	return TResizeStyle(r1)
}

func (m *TCustomSplitter) SetResizeStyle(AValue TResizeStyle) {
	LCL().SysCallN(2288, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) GetOtherResizeControl() IControl {
	r1 := LCL().SysCallN(2282, m.Instance())
	return AsControl(r1)
}

func (m *TCustomSplitter) GetSplitterPosition() int32 {
	r1 := LCL().SysCallN(2283, m.Instance())
	return int32(r1)
}

func CustomSplitterClass() TClass {
	ret := LCL().SysCallN(2280)
	return TClass(ret)
}

func (m *TCustomSplitter) AnchorSplitter(Kind TAnchorKind, AControl IControl) {
	LCL().SysCallN(2277, m.Instance(), uintptr(Kind), GetObjectUintptr(AControl))
}

func (m *TCustomSplitter) MoveSplitter(Offset int32) {
	LCL().SysCallN(2285, m.Instance(), uintptr(Offset))
}

func (m *TCustomSplitter) SetSplitterPosition(NewPosition int32) {
	LCL().SysCallN(2292, m.Instance(), uintptr(NewPosition))
}

func (m *TCustomSplitter) SetOnCanOffset(fn TCanOffsetEvent) {
	if m.canOffsetPtr != 0 {
		RemoveEventElement(m.canOffsetPtr)
	}
	m.canOffsetPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2289, m.Instance(), m.canOffsetPtr)
}

func (m *TCustomSplitter) SetOnCanResize(fn TCanResizeEvent) {
	if m.canResizePtr != 0 {
		RemoveEventElement(m.canResizePtr)
	}
	m.canResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2290, m.Instance(), m.canResizePtr)
}

func (m *TCustomSplitter) SetOnMoved(fn TNotifyEvent) {
	if m.movedPtr != 0 {
		RemoveEventElement(m.movedPtr)
	}
	m.movedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2291, m.Instance(), m.movedPtr)
}
