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

// IDockManager Is Abstract Class Parent: IPersistent
type IDockManager interface {
	IPersistent
	GetDockEdge(ADockObject IDragDockObject) bool                                  // function
	AutoFreeByControl() bool                                                       // function
	IsEnabledControl(Control IControl) bool                                        // function
	CanBeDoubleDocked() bool                                                       // function
	BeginUpdate()                                                                  // procedure
	EndUpdate()                                                                    // procedure
	GetControlBounds(Control IControl, OutControlBounds *TRect)                    // procedure Is Abstract
	InsertControl(ADockObject IDragDockObject)                                     // procedure
	InsertControl1(Control IControl, InsertAt TAlign, DropCtl IControl)            // procedure Is Abstract
	LoadFromStream(Stream IStream)                                                 // procedure Is Abstract
	PaintSite(DC HDC)                                                              // procedure
	MessageHandler(Sender IControl, Message *TLMessage)                            // procedure
	PositionDockRect(ADockObject IDragDockObject)                                  // procedure
	PositionDockRect1(Client, DropCtl IControl, DropAlign TAlign, DockRect *TRect) // procedure Is Abstract
	RemoveControl(Control IControl)                                                // procedure Is Abstract
	ResetBounds(Force bool)                                                        // procedure Is Abstract
	SaveToStream(Stream IStream)                                                   // procedure Is Abstract
	SetReplacingControl(Control IControl)                                          // procedure
}

// TDockManager Is Abstract Class Parent: TPersistent
type TDockManager struct {
	TPersistent
}

func (m *TDockManager) GetDockEdge(ADockObject IDragDockObject) bool {
	r1 := LCL().SysCallN(2659, m.Instance(), GetObjectUintptr(ADockObject))
	return GoBool(r1)
}

func (m *TDockManager) AutoFreeByControl() bool {
	r1 := LCL().SysCallN(2653, m.Instance())
	return GoBool(r1)
}

func (m *TDockManager) IsEnabledControl(Control IControl) bool {
	r1 := LCL().SysCallN(2662, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TDockManager) CanBeDoubleDocked() bool {
	r1 := LCL().SysCallN(2655, m.Instance())
	return GoBool(r1)
}

func DockManagerClass() TClass {
	ret := LCL().SysCallN(2656)
	return TClass(ret)
}

func (m *TDockManager) BeginUpdate() {
	LCL().SysCallN(2654, m.Instance())
}

func (m *TDockManager) EndUpdate() {
	LCL().SysCallN(2657, m.Instance())
}

func (m *TDockManager) GetControlBounds(Control IControl, OutControlBounds *TRect) {
	var result1 uintptr
	LCL().SysCallN(2658, m.Instance(), GetObjectUintptr(Control), uintptr(unsafePointer(&result1)))
	*OutControlBounds = *(*TRect)(getPointer(result1))
}

func (m *TDockManager) InsertControl(ADockObject IDragDockObject) {
	LCL().SysCallN(2660, m.Instance(), GetObjectUintptr(ADockObject))
}

func (m *TDockManager) InsertControl1(Control IControl, InsertAt TAlign, DropCtl IControl) {
	LCL().SysCallN(2661, m.Instance(), GetObjectUintptr(Control), uintptr(InsertAt), GetObjectUintptr(DropCtl))
}

func (m *TDockManager) LoadFromStream(Stream IStream) {
	LCL().SysCallN(2663, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TDockManager) PaintSite(DC HDC) {
	LCL().SysCallN(2665, m.Instance(), uintptr(DC))
}

func (m *TDockManager) MessageHandler(Sender IControl, Message *TLMessage) {
	var result1 uintptr
	LCL().SysCallN(2664, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*Message = *(*TLMessage)(getPointer(result1))
}

func (m *TDockManager) PositionDockRect(ADockObject IDragDockObject) {
	LCL().SysCallN(2666, m.Instance(), GetObjectUintptr(ADockObject))
}

func (m *TDockManager) PositionDockRect1(Client, DropCtl IControl, DropAlign TAlign, DockRect *TRect) {
	var result2 uintptr
	LCL().SysCallN(2667, m.Instance(), GetObjectUintptr(Client), GetObjectUintptr(DropCtl), uintptr(DropAlign), uintptr(unsafePointer(&result2)))
	*DockRect = *(*TRect)(getPointer(result2))
}

func (m *TDockManager) RemoveControl(Control IControl) {
	LCL().SysCallN(2668, m.Instance(), GetObjectUintptr(Control))
}

func (m *TDockManager) ResetBounds(Force bool) {
	LCL().SysCallN(2669, m.Instance(), PascalBool(Force))
}

func (m *TDockManager) SaveToStream(Stream IStream) {
	LCL().SysCallN(2670, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TDockManager) SetReplacingControl(Control IControl) {
	LCL().SysCallN(2671, m.Instance(), GetObjectUintptr(Control))
}
