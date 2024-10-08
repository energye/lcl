//----------------------------------------
// The code is automatically generated by the generate-tool.
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
)

// CreateResForm : 从资源中创建Form，不使用Application.CreateForm
//func CreateResForm(owner IComponent, fields ...interface{}) {
//	resObjectBuild(1, owner, 0, fields...)
//}

func (m *TForm) EnabledMaximize(v bool) {
	bi := m.BorderIcons()
	if v {
		if !bi.In(types.BiMaximize) {
			bi = bi.Include(types.BiMaximize)
			m.SetBorderIcons(bi)
		}
	} else {
		if bi.In(types.BiMaximize) {
			bi = bi.Exclude(types.BiMaximize)
			m.SetBorderIcons(bi)
		}
	}
}

func (m *TForm) EnabledMinimize(v bool) {
	bi := m.BorderIcons()
	if v {
		if !bi.In(types.BiMinimize) {
			bi = bi.Include(types.BiMinimize)
			m.SetBorderIcons(bi)
		}
	} else {
		if bi.In(types.BiMinimize) {
			bi = bi.Exclude(types.BiMinimize)
			m.SetBorderIcons(bi)
		}
	}
}

func (m *TForm) EnabledSystemMenu(v bool) {
	bi := m.BorderIcons()
	if v {
		if !bi.In(types.BiSystemMenu) {
			bi = bi.Include(types.BiSystemMenu)
			m.SetBorderIcons(bi)
		}
	} else {
		if bi.In(types.BiSystemMenu) {
			bi = bi.Exclude(types.BiSystemMenu)
			m.SetBorderIcons(bi)
		}
	}
}

func (m *TForm) ScreenCenter() {
	m.SetLeft((Screen.Width() - m.Width()) / 2)
	m.SetTop((Screen.Height() - m.Height()) / 2)
}

func (m *TForm) WorkAreaCenter() {
	m.SetLeft((Screen.WorkAreaWidth() - m.Width()) / 2)
	m.SetTop((Screen.WorkAreaHeight() - m.Height()) / 2)
}

func (m *TForm) ScaleForPPI(newPPI int32) {
	psi := m.PixelsPerInch()
	if newPPI >= 30 && newPPI > psi {
		m.AutoAdjustLayout(types.LapAutoAdjustForDPI, psi, newPPI, MulDiv(m.Width(), newPPI, psi), MulDiv(m.Height(), newPPI, psi))
	}
}

func (m *TForm) ScaleForCurrentDpi() {
	if !m.Scaled() {
		m.SetScaled(true)
	} else {
		psi := m.PixelsPerInch()
		monPsi := m.Monitor().PixelsPerInch()
		if psi != monPsi {
			m.AutoAdjustLayout(types.LapAutoAdjustForDPI, psi, monPsi, MulDiv(m.Width(), monPsi, psi), MulDiv(m.Height(), monPsi, psi))
		}
	}
}

func (m *TForm) InheritedWndProc(heMessage *types.TMessage) {
	api.Form_InheritedWndProc(m.Instance(), heMessage)
}

// SetOnWndProc : 窗口消息过程
func (m *TForm) SetOnWndProc(fn TWndProcEvent) {
	api.Form_SetOnWndProc(m.Instance(), fn)
}

func (m *TForm) SetGoPtr(ptr uintptr) {
	api.Form_SetGoPtr(m.Instance(), ptr)
}

// ScaleSelf : 这个方法主要是用于当不使用资源窗口创建时用，这个方法要用于设置了Width, Height或者ClientWidth、ClientHeight之后
func (m *TForm) ScaleSelf() {
	if Application.Scaled() {
		m.SetClientWidth(int32(float64(m.ClientWidth()) * (float64(Screen.PixelsPerInch()) / 96.0)))
		m.SetClientHeight(int32(float64(m.ClientHeight()) * (float64(Screen.PixelsPerInch()) / 96.0)))
	}
}

// IOnCreate TForm OnCreate
type IOnCreate interface {
	FormCreate(sender IObject)
}

// IOnAfterCreate TForm OnAfterCreate
type IOnAfterCreate interface {
	FormAfterCreate(sender IObject)
}

// IOnCreateParams TForm CreateParams
type IOnCreateParams interface {
	CreateParams(params *types.TCreateParams)
}
