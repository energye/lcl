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

// ICustomOpenGLControl Parent: IWinControl
type ICustomOpenGLControl interface {
	IWinControl
	MakeCurrent(saveOldToStack bool) bool             // function
	ReleaseContext() bool                             // function
	RestoreOldOpenGLControl() bool                    // function
	SharingControlCount() int32                       // function
	Paint()                                           // procedure
	RealizeBounds()                                   // procedure
	DoOnPaint()                                       // procedure
	SwapBuffers()                                     // procedure
	SharingControls(index int32) ICustomOpenGLControl // property SharingControls Getter
	FrameDiffTimeInMSecs() int32                      // property FrameDiffTimeInMSecs Getter
	SharedControl() ICustomOpenGLControl              // property SharedControl Getter
	SetSharedControl(value ICustomOpenGLControl)      // property SharedControl Setter
	AutoResizeViewport() bool                         // property AutoResizeViewport Getter
	SetAutoResizeViewport(value bool)                 // property AutoResizeViewport Setter
	DebugContext() bool                               // property DebugContext Getter
	SetDebugContext(value bool)                       // property DebugContext Setter
	RGBA() bool                                       // property RGBA Getter
	SetRGBA(value bool)                               // property RGBA Setter
	OpenGLMajorVersion() uint32                       // property OpenGLMajorVersion Getter
	SetOpenGLMajorVersion(value uint32)               // property OpenGLMajorVersion Setter
	OpenGLMinorVersion() uint32                       // property OpenGLMinorVersion Getter
	SetOpenGLMinorVersion(value uint32)               // property OpenGLMinorVersion Setter
	// MultiSampling
	//  Number of samples per pixel, for OpenGL multi-sampling (anti-aliasing).
	//  Value <= 1 means that we use 1 sample per pixel, which means no anti-aliasing.
	//  Higher values mean anti-aliasing. Exactly which values are supported
	//  depends on GPU, common modern GPUs support values like 2 and 4.
	//  If this is > 1, and we will not be able to create OpenGL
	//  with multi-sampling, we will fallback to normal non-multi-sampled context.
	//  You can query OpenGL values GL_SAMPLE_BUFFERS_ARB and GL_SAMPLES_ARB
	//  (see ARB_multisample extension) to see how many samples have been
	//  actually allocated for your context.
	MultiSampling() uint32                           // property MultiSampling Getter
	SetMultiSampling(value uint32)                   // property MultiSampling Setter
	AlphaBits() uint32                               // property AlphaBits Getter
	SetAlphaBits(value uint32)                       // property AlphaBits Setter
	DepthBits() uint32                               // property DepthBits Getter
	SetDepthBits(value uint32)                       // property DepthBits Setter
	StencilBits() uint32                             // property StencilBits Getter
	SetStencilBits(value uint32)                     // property StencilBits Setter
	AUXBuffers() uint32                              // property AUXBuffers Getter
	SetAUXBuffers(value uint32)                      // property AUXBuffers Setter
	Options() types.TOpenGLControlOptions            // property Options Getter
	SetOptions(value types.TOpenGLControlOptions)    // property Options Setter
	SetOnMakeCurrent(fn TOpenGlCtrlMakeCurrentEvent) // property event
	SetOnPaint(fn TNotifyEvent)                      // property event
}

type TCustomOpenGLControl struct {
	TWinControl
}

func (m *TCustomOpenGLControl) MakeCurrent(saveOldToStack bool) bool {
	if !m.IsValid() {
		return false
	}
	r := customOpenGLControlAPI().SysCallN(1, m.Instance(), api.PasBool(saveOldToStack))
	return api.GoBool(r)
}

func (m *TCustomOpenGLControl) ReleaseContext() bool {
	if !m.IsValid() {
		return false
	}
	r := customOpenGLControlAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomOpenGLControl) RestoreOldOpenGLControl() bool {
	if !m.IsValid() {
		return false
	}
	r := customOpenGLControlAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomOpenGLControl) SharingControlCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCustomOpenGLControl) Paint() {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(5, m.Instance())
}

func (m *TCustomOpenGLControl) RealizeBounds() {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(6, m.Instance())
}

func (m *TCustomOpenGLControl) DoOnPaint() {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(7, m.Instance())
}

func (m *TCustomOpenGLControl) SwapBuffers() {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(8, m.Instance())
}

func (m *TCustomOpenGLControl) SharingControls(index int32) ICustomOpenGLControl {
	if !m.IsValid() {
		return nil
	}
	r := customOpenGLControlAPI().SysCallN(9, m.Instance(), uintptr(index))
	return AsCustomOpenGLControl(r)
}

func (m *TCustomOpenGLControl) FrameDiffTimeInMSecs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TCustomOpenGLControl) SharedControl() ICustomOpenGLControl {
	if !m.IsValid() {
		return nil
	}
	r := customOpenGLControlAPI().SysCallN(11, 0, m.Instance())
	return AsCustomOpenGLControl(r)
}

func (m *TCustomOpenGLControl) SetSharedControl(value ICustomOpenGLControl) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomOpenGLControl) AutoResizeViewport() bool {
	if !m.IsValid() {
		return false
	}
	r := customOpenGLControlAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomOpenGLControl) SetAutoResizeViewport(value bool) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomOpenGLControl) DebugContext() bool {
	if !m.IsValid() {
		return false
	}
	r := customOpenGLControlAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomOpenGLControl) SetDebugContext(value bool) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomOpenGLControl) RGBA() bool {
	if !m.IsValid() {
		return false
	}
	r := customOpenGLControlAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomOpenGLControl) SetRGBA(value bool) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomOpenGLControl) OpenGLMajorVersion() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(15, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomOpenGLControl) SetOpenGLMajorVersion(value uint32) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) OpenGLMinorVersion() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(16, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomOpenGLControl) SetOpenGLMinorVersion(value uint32) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) MultiSampling() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(17, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomOpenGLControl) SetMultiSampling(value uint32) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) AlphaBits() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(18, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomOpenGLControl) SetAlphaBits(value uint32) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) DepthBits() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(19, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomOpenGLControl) SetDepthBits(value uint32) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) StencilBits() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(20, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomOpenGLControl) SetStencilBits(value uint32) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) AUXBuffers() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(21, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomOpenGLControl) SetAUXBuffers(value uint32) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) Options() types.TOpenGLControlOptions {
	if !m.IsValid() {
		return 0
	}
	r := customOpenGLControlAPI().SysCallN(22, 0, m.Instance())
	return types.TOpenGLControlOptions(r)
}

func (m *TCustomOpenGLControl) SetOptions(value types.TOpenGLControlOptions) {
	if !m.IsValid() {
		return
	}
	customOpenGLControlAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCustomOpenGLControl) SetOnMakeCurrent(fn TOpenGlCtrlMakeCurrentEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOpenGlCtrlMakeCurrentEvent(fn)
	base.SetEvent(m, 23, customOpenGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomOpenGLControl) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, customOpenGLControlAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomOpenGLControl class constructor
func NewCustomOpenGLControl(theOwner IComponent) ICustomOpenGLControl {
	r := customOpenGLControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomOpenGLControl(r)
}

func TCustomOpenGLControlClass() types.TClass {
	r := customOpenGLControlAPI().SysCallN(25)
	return types.TClass(r)
}

var (
	customOpenGLControlOnce   base.Once
	customOpenGLControlImport *imports.Imports = nil
)

func customOpenGLControlAPI() *imports.Imports {
	customOpenGLControlOnce.Do(func() {
		customOpenGLControlImport = api.NewDefaultImports()
		customOpenGLControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomOpenGLControl_Create", 0), // constructor NewCustomOpenGLControl
			/* 1 */ imports.NewTable("TCustomOpenGLControl_MakeCurrent", 0), // function MakeCurrent
			/* 2 */ imports.NewTable("TCustomOpenGLControl_ReleaseContext", 0), // function ReleaseContext
			/* 3 */ imports.NewTable("TCustomOpenGLControl_RestoreOldOpenGLControl", 0), // function RestoreOldOpenGLControl
			/* 4 */ imports.NewTable("TCustomOpenGLControl_SharingControlCount", 0), // function SharingControlCount
			/* 5 */ imports.NewTable("TCustomOpenGLControl_Paint", 0), // procedure Paint
			/* 6 */ imports.NewTable("TCustomOpenGLControl_RealizeBounds", 0), // procedure RealizeBounds
			/* 7 */ imports.NewTable("TCustomOpenGLControl_DoOnPaint", 0), // procedure DoOnPaint
			/* 8 */ imports.NewTable("TCustomOpenGLControl_SwapBuffers", 0), // procedure SwapBuffers
			/* 9 */ imports.NewTable("TCustomOpenGLControl_SharingControls", 0), // property SharingControls
			/* 10 */ imports.NewTable("TCustomOpenGLControl_FrameDiffTimeInMSecs", 0), // property FrameDiffTimeInMSecs
			/* 11 */ imports.NewTable("TCustomOpenGLControl_SharedControl", 0), // property SharedControl
			/* 12 */ imports.NewTable("TCustomOpenGLControl_AutoResizeViewport", 0), // property AutoResizeViewport
			/* 13 */ imports.NewTable("TCustomOpenGLControl_DebugContext", 0), // property DebugContext
			/* 14 */ imports.NewTable("TCustomOpenGLControl_RGBA", 0), // property RGBA
			/* 15 */ imports.NewTable("TCustomOpenGLControl_OpenGLMajorVersion", 0), // property OpenGLMajorVersion
			/* 16 */ imports.NewTable("TCustomOpenGLControl_OpenGLMinorVersion", 0), // property OpenGLMinorVersion
			/* 17 */ imports.NewTable("TCustomOpenGLControl_MultiSampling", 0), // property MultiSampling
			/* 18 */ imports.NewTable("TCustomOpenGLControl_AlphaBits", 0), // property AlphaBits
			/* 19 */ imports.NewTable("TCustomOpenGLControl_DepthBits", 0), // property DepthBits
			/* 20 */ imports.NewTable("TCustomOpenGLControl_StencilBits", 0), // property StencilBits
			/* 21 */ imports.NewTable("TCustomOpenGLControl_AUXBuffers", 0), // property AUXBuffers
			/* 22 */ imports.NewTable("TCustomOpenGLControl_Options", 0), // property Options
			/* 23 */ imports.NewTable("TCustomOpenGLControl_OnMakeCurrent", 0), // event OnMakeCurrent
			/* 24 */ imports.NewTable("TCustomOpenGLControl_OnPaint", 0), // event OnPaint
			/* 25 */ imports.NewTable("TCustomOpenGLControl_TClass", 0), // function TCustomOpenGLControlClass
		}
	})
	return customOpenGLControlImport
}
