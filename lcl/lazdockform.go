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

// ILazDockForm Parent: ICustomForm
type ILazDockForm interface {
	ICustomForm
	FindMainControlCandidate() IControl                                      // function
	FindHeader(X int32, Y int32, outPart *types.TLazDockHeaderPart) IControl // function
	IsDockedControl(control IControl) bool                                   // function
	ControlHasTitle(control IControl) bool                                   // function
	GetTitleRect(control IControl) types.TRect                               // function
	GetTitleOrientation(control IControl) types.TDockOrientation             // function
	UpdateCaption()                                                          // procedure
	MainControl() IControl                                                   // property MainControl Getter
	SetMainControl(value IControl)                                           // property MainControl Setter
}

type TLazDockForm struct {
	TCustomForm
}

func (m *TLazDockForm) FindMainControlCandidate() IControl {
	if !m.IsValid() {
		return nil
	}
	r := lazDockFormAPI().SysCallN(1, m.Instance())
	return AsControl(r)
}

func (m *TLazDockForm) FindHeader(X int32, Y int32, outPart *types.TLazDockHeaderPart) IControl {
	if !m.IsValid() {
		return nil
	}
	var partPtr uintptr
	r := lazDockFormAPI().SysCallN(2, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&partPtr)))
	*outPart = types.TLazDockHeaderPart(partPtr)
	return AsControl(r)
}

func (m *TLazDockForm) IsDockedControl(control IControl) bool {
	if !m.IsValid() {
		return false
	}
	r := lazDockFormAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(control))
	return api.GoBool(r)
}

func (m *TLazDockForm) ControlHasTitle(control IControl) bool {
	if !m.IsValid() {
		return false
	}
	r := lazDockFormAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(control))
	return api.GoBool(r)
}

func (m *TLazDockForm) GetTitleRect(control IControl) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	lazDockFormAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazDockForm) GetTitleOrientation(control IControl) types.TDockOrientation {
	if !m.IsValid() {
		return 0
	}
	r := lazDockFormAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(control))
	return types.TDockOrientation(r)
}

func (m *TLazDockForm) UpdateCaption() {
	if !m.IsValid() {
		return
	}
	lazDockFormAPI().SysCallN(7, m.Instance())
}

func (m *TLazDockForm) MainControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := lazDockFormAPI().SysCallN(9, 0, m.Instance())
	return AsControl(r)
}

func (m *TLazDockForm) SetMainControl(value IControl) {
	if !m.IsValid() {
		return
	}
	lazDockFormAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

// LazDockForm  is static instance
var LazDockForm _LazDockFormClass

// _LazDockFormClass is class type defined by TLazDockForm
type _LazDockFormClass uintptr

func (_LazDockFormClass) UpdateMainControlInParents(startControl IControl) {
	lazDockFormAPI().SysCallN(8, base.GetObjectUintptr(startControl))
}

// NewLazDockForm class constructor
func NewLazDockForm(owner IComponent) ILazDockForm {
	r := lazDockFormAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazDockForm(r)
}

func TLazDockFormClass() types.TClass {
	r := lazDockFormAPI().SysCallN(10)
	return types.TClass(r)
}

var (
	lazDockFormOnce   base.Once
	lazDockFormImport *imports.Imports = nil
)

func lazDockFormAPI() *imports.Imports {
	lazDockFormOnce.Do(func() {
		lazDockFormImport = api.NewDefaultImports()
		lazDockFormImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazDockForm_Create", 0), // constructor NewLazDockForm
			/* 1 */ imports.NewTable("TLazDockForm_FindMainControlCandidate", 0), // function FindMainControlCandidate
			/* 2 */ imports.NewTable("TLazDockForm_FindHeader", 0), // function FindHeader
			/* 3 */ imports.NewTable("TLazDockForm_IsDockedControl", 0), // function IsDockedControl
			/* 4 */ imports.NewTable("TLazDockForm_ControlHasTitle", 0), // function ControlHasTitle
			/* 5 */ imports.NewTable("TLazDockForm_GetTitleRect", 0), // function GetTitleRect
			/* 6 */ imports.NewTable("TLazDockForm_GetTitleOrientation", 0), // function GetTitleOrientation
			/* 7 */ imports.NewTable("TLazDockForm_UpdateCaption", 0), // procedure UpdateCaption
			/* 8 */ imports.NewTable("TLazDockForm_UpdateMainControlInParents", 0), // static procedure UpdateMainControlInParents
			/* 9 */ imports.NewTable("TLazDockForm_MainControl", 0), // property MainControl
			/* 10 */ imports.NewTable("TLazDockForm_TClass", 0), // function TLazDockFormClass
		}
	})
	return lazDockFormImport
}
