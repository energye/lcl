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

// IMonitor Parent: IObject
type IMonitor interface {
	IObject
	Handle() HMONITOR                 // property
	MonitorNum() int32                // property
	Left() int32                      // property
	Height() int32                    // property
	Top() int32                       // property
	Width() int32                     // property
	BoundsRect() (resultRect TRect)   // property
	WorkareaRect() (resultRect TRect) // property
	Primary() bool                    // property
	PixelsPerInch() int32             // property
}

// TMonitor Parent: TObject
type TMonitor struct {
	TObject
}

func NewMonitor() IMonitor {
	r1 := LCL().SysCallN(4369)
	return AsMonitor(r1)
}

func (m *TMonitor) Handle() HMONITOR {
	r1 := LCL().SysCallN(4370, m.Instance())
	return HMONITOR(r1)
}

func (m *TMonitor) MonitorNum() int32 {
	r1 := LCL().SysCallN(4373, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Left() int32 {
	r1 := LCL().SysCallN(4372, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Height() int32 {
	r1 := LCL().SysCallN(4371, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Top() int32 {
	r1 := LCL().SysCallN(4376, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Width() int32 {
	r1 := LCL().SysCallN(4377, m.Instance())
	return int32(r1)
}

func (m *TMonitor) BoundsRect() (resultRect TRect) {
	LCL().SysCallN(4367, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TMonitor) WorkareaRect() (resultRect TRect) {
	LCL().SysCallN(4378, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TMonitor) Primary() bool {
	r1 := LCL().SysCallN(4375, m.Instance())
	return GoBool(r1)
}

func (m *TMonitor) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(4374, m.Instance())
	return int32(r1)
}

func MonitorClass() TClass {
	ret := LCL().SysCallN(4368)
	return TClass(ret)
}
