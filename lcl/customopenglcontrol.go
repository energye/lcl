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

// ICustomOpenGLControl Parent: IWinControl
type ICustomOpenGLControl interface {
	IWinControl
	SharingControls(Index int32) ICustomOpenGLControl // property
	FrameDiffTimeInMSecs() int32                      // property
	SharedControl() ICustomOpenGLControl              // property
	SetSharedControl(AValue ICustomOpenGLControl)     // property
	AutoResizeViewport() bool                         // property
	SetAutoResizeViewport(AValue bool)                // property
	DebugContext() bool                               // property
	SetDebugContext(AValue bool)                      // property
	RGBA() bool                                       // property
	SetRGBA(AValue bool)                              // property
	RedBits() uint32                                  // property
	SetRedBits(AValue uint32)                         // property
	GreenBits() uint32                                // property
	SetGreenBits(AValue uint32)                       // property
	BlueBits() uint32                                 // property
	SetBlueBits(AValue uint32)                        // property
	OpenGLMajorVersion() uint32                       // property
	SetOpenGLMajorVersion(AValue uint32)              // property
	OpenGLMinorVersion() uint32                       // property
	SetOpenGLMinorVersion(AValue uint32)              // property
	MultiSampling() uint32                            // property
	SetMultiSampling(AValue uint32)                   // property
	AlphaBits() uint32                                // property
	SetAlphaBits(AValue uint32)                       // property
	DepthBits() uint32                                // property
	SetDepthBits(AValue uint32)                       // property
	StencilBits() uint32                              // property
	SetStencilBits(AValue uint32)                     // property
	AUXBuffers() uint32                               // property
	SetAUXBuffers(AValue uint32)                      // property
	Options() TOpenGLControlOptions                   // property
	SetOptions(AValue TOpenGLControlOptions)          // property
	MakeCurrent(SaveOldToStack bool) bool             // function
	ReleaseContext() bool                             // function
	RestoreOldOpenGLControl() bool                    // function
	SharingControlCount() int32                       // function
	Paint()                                           // procedure
	RealizeBounds()                                   // procedure
	DoOnPaint()                                       // procedure
	SwapBuffers()                                     // procedure
	SetOnMakeCurrent(fn TOpenGlCtrlMakeCurrentEvent)  // property event
	SetOnPaint(fn TNotifyEvent)                       // property event
}

// TCustomOpenGLControl Parent: TWinControl
type TCustomOpenGLControl struct {
	TWinControl
	makeCurrentPtr uintptr
	paintPtr       uintptr
}

func NewCustomOpenGLControl(TheOwner IComponent) ICustomOpenGLControl {
	r1 := LCL().SysCallN(2112, GetObjectUintptr(TheOwner))
	return AsCustomOpenGLControl(r1)
}

func (m *TCustomOpenGLControl) SharingControls(Index int32) ICustomOpenGLControl {
	r1 := LCL().SysCallN(2133, m.Instance(), uintptr(Index))
	return AsCustomOpenGLControl(r1)
}

func (m *TCustomOpenGLControl) FrameDiffTimeInMSecs() int32 {
	r1 := LCL().SysCallN(2116, m.Instance())
	return int32(r1)
}

func (m *TCustomOpenGLControl) SharedControl() ICustomOpenGLControl {
	r1 := LCL().SysCallN(2131, 0, m.Instance(), 0)
	return AsCustomOpenGLControl(r1)
}

func (m *TCustomOpenGLControl) SetSharedControl(AValue ICustomOpenGLControl) {
	LCL().SysCallN(2131, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomOpenGLControl) AutoResizeViewport() bool {
	r1 := LCL().SysCallN(2109, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetAutoResizeViewport(AValue bool) {
	LCL().SysCallN(2109, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) DebugContext() bool {
	r1 := LCL().SysCallN(2113, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetDebugContext(AValue bool) {
	LCL().SysCallN(2113, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) RGBA() bool {
	r1 := LCL().SysCallN(2124, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetRGBA(AValue bool) {
	LCL().SysCallN(2124, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) RedBits() uint32 {
	r1 := LCL().SysCallN(2126, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetRedBits(AValue uint32) {
	LCL().SysCallN(2126, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) GreenBits() uint32 {
	r1 := LCL().SysCallN(2117, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetGreenBits(AValue uint32) {
	LCL().SysCallN(2117, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) BlueBits() uint32 {
	r1 := LCL().SysCallN(2110, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetBlueBits(AValue uint32) {
	LCL().SysCallN(2110, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) OpenGLMajorVersion() uint32 {
	r1 := LCL().SysCallN(2120, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetOpenGLMajorVersion(AValue uint32) {
	LCL().SysCallN(2120, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) OpenGLMinorVersion() uint32 {
	r1 := LCL().SysCallN(2121, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetOpenGLMinorVersion(AValue uint32) {
	LCL().SysCallN(2121, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) MultiSampling() uint32 {
	r1 := LCL().SysCallN(2119, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetMultiSampling(AValue uint32) {
	LCL().SysCallN(2119, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) AlphaBits() uint32 {
	r1 := LCL().SysCallN(2108, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetAlphaBits(AValue uint32) {
	LCL().SysCallN(2108, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) DepthBits() uint32 {
	r1 := LCL().SysCallN(2114, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetDepthBits(AValue uint32) {
	LCL().SysCallN(2114, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) StencilBits() uint32 {
	r1 := LCL().SysCallN(2134, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetStencilBits(AValue uint32) {
	LCL().SysCallN(2134, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) AUXBuffers() uint32 {
	r1 := LCL().SysCallN(2107, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetAUXBuffers(AValue uint32) {
	LCL().SysCallN(2107, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) Options() TOpenGLControlOptions {
	r1 := LCL().SysCallN(2122, 0, m.Instance(), 0)
	return TOpenGLControlOptions(r1)
}

func (m *TCustomOpenGLControl) SetOptions(AValue TOpenGLControlOptions) {
	LCL().SysCallN(2122, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) MakeCurrent(SaveOldToStack bool) bool {
	r1 := LCL().SysCallN(2118, m.Instance(), PascalBool(SaveOldToStack))
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) ReleaseContext() bool {
	r1 := LCL().SysCallN(2127, m.Instance())
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) RestoreOldOpenGLControl() bool {
	r1 := LCL().SysCallN(2128, m.Instance())
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SharingControlCount() int32 {
	r1 := LCL().SysCallN(2132, m.Instance())
	return int32(r1)
}

func CustomOpenGLControlClass() TClass {
	ret := LCL().SysCallN(2111)
	return TClass(ret)
}

func (m *TCustomOpenGLControl) Paint() {
	LCL().SysCallN(2123, m.Instance())
}

func (m *TCustomOpenGLControl) RealizeBounds() {
	LCL().SysCallN(2125, m.Instance())
}

func (m *TCustomOpenGLControl) DoOnPaint() {
	LCL().SysCallN(2115, m.Instance())
}

func (m *TCustomOpenGLControl) SwapBuffers() {
	LCL().SysCallN(2135, m.Instance())
}

func (m *TCustomOpenGLControl) SetOnMakeCurrent(fn TOpenGlCtrlMakeCurrentEvent) {
	if m.makeCurrentPtr != 0 {
		RemoveEventElement(m.makeCurrentPtr)
	}
	m.makeCurrentPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2129, m.Instance(), m.makeCurrentPtr)
}

func (m *TCustomOpenGLControl) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2130, m.Instance(), m.paintPtr)
}
