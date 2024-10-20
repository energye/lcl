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
	"github.com/energye/lcl/api/imports"
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
	r1 := customOpenGLControlImportAPI().SysCallN(5, GetObjectUintptr(TheOwner))
	return AsCustomOpenGLControl(r1)
}

func (m *TCustomOpenGLControl) SharingControls(Index int32) ICustomOpenGLControl {
	r1 := customOpenGLControlImportAPI().SysCallN(26, m.Instance(), uintptr(Index))
	return AsCustomOpenGLControl(r1)
}

func (m *TCustomOpenGLControl) FrameDiffTimeInMSecs() int32 {
	r1 := customOpenGLControlImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TCustomOpenGLControl) SharedControl() ICustomOpenGLControl {
	r1 := customOpenGLControlImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return AsCustomOpenGLControl(r1)
}

func (m *TCustomOpenGLControl) SetSharedControl(AValue ICustomOpenGLControl) {
	customOpenGLControlImportAPI().SysCallN(24, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomOpenGLControl) AutoResizeViewport() bool {
	r1 := customOpenGLControlImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetAutoResizeViewport(AValue bool) {
	customOpenGLControlImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) DebugContext() bool {
	r1 := customOpenGLControlImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetDebugContext(AValue bool) {
	customOpenGLControlImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) RGBA() bool {
	r1 := customOpenGLControlImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetRGBA(AValue bool) {
	customOpenGLControlImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) RedBits() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetRedBits(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) GreenBits() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetGreenBits(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) BlueBits() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetBlueBits(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) OpenGLMajorVersion() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetOpenGLMajorVersion(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) OpenGLMinorVersion() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetOpenGLMinorVersion(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) MultiSampling() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetMultiSampling(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) AlphaBits() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetAlphaBits(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) DepthBits() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetDepthBits(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) StencilBits() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetStencilBits(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) AUXBuffers() uint32 {
	r1 := customOpenGLControlImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetAUXBuffers(AValue uint32) {
	customOpenGLControlImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) Options() TOpenGLControlOptions {
	r1 := customOpenGLControlImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TOpenGLControlOptions(r1)
}

func (m *TCustomOpenGLControl) SetOptions(AValue TOpenGLControlOptions) {
	customOpenGLControlImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) MakeCurrent(SaveOldToStack bool) bool {
	r1 := customOpenGLControlImportAPI().SysCallN(11, m.Instance(), PascalBool(SaveOldToStack))
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) ReleaseContext() bool {
	r1 := customOpenGLControlImportAPI().SysCallN(20, m.Instance())
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) RestoreOldOpenGLControl() bool {
	r1 := customOpenGLControlImportAPI().SysCallN(21, m.Instance())
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SharingControlCount() int32 {
	r1 := customOpenGLControlImportAPI().SysCallN(25, m.Instance())
	return int32(r1)
}

func CustomOpenGLControlClass() TClass {
	ret := customOpenGLControlImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TCustomOpenGLControl) Paint() {
	customOpenGLControlImportAPI().SysCallN(16, m.Instance())
}

func (m *TCustomOpenGLControl) RealizeBounds() {
	customOpenGLControlImportAPI().SysCallN(18, m.Instance())
}

func (m *TCustomOpenGLControl) DoOnPaint() {
	customOpenGLControlImportAPI().SysCallN(8, m.Instance())
}

func (m *TCustomOpenGLControl) SwapBuffers() {
	customOpenGLControlImportAPI().SysCallN(28, m.Instance())
}

func (m *TCustomOpenGLControl) SetOnMakeCurrent(fn TOpenGlCtrlMakeCurrentEvent) {
	if m.makeCurrentPtr != 0 {
		RemoveEventElement(m.makeCurrentPtr)
	}
	m.makeCurrentPtr = MakeEventDataPtr(fn)
	customOpenGLControlImportAPI().SysCallN(22, m.Instance(), m.makeCurrentPtr)
}

func (m *TCustomOpenGLControl) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	customOpenGLControlImportAPI().SysCallN(23, m.Instance(), m.paintPtr)
}

var (
	customOpenGLControlImport       *imports.Imports = nil
	customOpenGLControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomOpenGLControl_AUXBuffers", 0),
		/*1*/ imports.NewTable("CustomOpenGLControl_AlphaBits", 0),
		/*2*/ imports.NewTable("CustomOpenGLControl_AutoResizeViewport", 0),
		/*3*/ imports.NewTable("CustomOpenGLControl_BlueBits", 0),
		/*4*/ imports.NewTable("CustomOpenGLControl_Class", 0),
		/*5*/ imports.NewTable("CustomOpenGLControl_Create", 0),
		/*6*/ imports.NewTable("CustomOpenGLControl_DebugContext", 0),
		/*7*/ imports.NewTable("CustomOpenGLControl_DepthBits", 0),
		/*8*/ imports.NewTable("CustomOpenGLControl_DoOnPaint", 0),
		/*9*/ imports.NewTable("CustomOpenGLControl_FrameDiffTimeInMSecs", 0),
		/*10*/ imports.NewTable("CustomOpenGLControl_GreenBits", 0),
		/*11*/ imports.NewTable("CustomOpenGLControl_MakeCurrent", 0),
		/*12*/ imports.NewTable("CustomOpenGLControl_MultiSampling", 0),
		/*13*/ imports.NewTable("CustomOpenGLControl_OpenGLMajorVersion", 0),
		/*14*/ imports.NewTable("CustomOpenGLControl_OpenGLMinorVersion", 0),
		/*15*/ imports.NewTable("CustomOpenGLControl_Options", 0),
		/*16*/ imports.NewTable("CustomOpenGLControl_Paint", 0),
		/*17*/ imports.NewTable("CustomOpenGLControl_RGBA", 0),
		/*18*/ imports.NewTable("CustomOpenGLControl_RealizeBounds", 0),
		/*19*/ imports.NewTable("CustomOpenGLControl_RedBits", 0),
		/*20*/ imports.NewTable("CustomOpenGLControl_ReleaseContext", 0),
		/*21*/ imports.NewTable("CustomOpenGLControl_RestoreOldOpenGLControl", 0),
		/*22*/ imports.NewTable("CustomOpenGLControl_SetOnMakeCurrent", 0),
		/*23*/ imports.NewTable("CustomOpenGLControl_SetOnPaint", 0),
		/*24*/ imports.NewTable("CustomOpenGLControl_SharedControl", 0),
		/*25*/ imports.NewTable("CustomOpenGLControl_SharingControlCount", 0),
		/*26*/ imports.NewTable("CustomOpenGLControl_SharingControls", 0),
		/*27*/ imports.NewTable("CustomOpenGLControl_StencilBits", 0),
		/*28*/ imports.NewTable("CustomOpenGLControl_SwapBuffers", 0),
	}
)

func customOpenGLControlImportAPI() *imports.Imports {
	if customOpenGLControlImport == nil {
		customOpenGLControlImport = NewDefaultImports()
		customOpenGLControlImport.SetImportTable(customOpenGLControlImportTables)
		customOpenGLControlImportTables = nil
	}
	return customOpenGLControlImport
}
