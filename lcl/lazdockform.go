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

// ILazDockForm Parent: ICustomForm
type ILazDockForm interface {
	ICustomForm
	MainControl() IControl                                       // property
	SetMainControl(AValue IControl)                              // property
	FindMainControlCandidate() IControl                          // function
	FindHeader(x, y int32, OutPart *TLazDockHeaderPart) IControl // function
	IsDockedControl(Control IControl) bool                       // function
	ControlHasTitle(Control IControl) bool                       // function
	GetTitleRect(Control IControl) (resultRect TRect)            // function
	GetTitleOrientation(Control IControl) TDockOrientation       // function
	UpdateCaption()                                              // procedure
}

// TLazDockForm Parent: TCustomForm
type TLazDockForm struct {
	TCustomForm
}

func NewLazDockForm(AOwner IComponent) ILazDockForm {
	r1 := lazDockFormImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsLazDockForm(r1)
}

func (m *TLazDockForm) MainControl() IControl {
	r1 := lazDockFormImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TLazDockForm) SetMainControl(AValue IControl) {
	lazDockFormImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockForm) FindMainControlCandidate() IControl {
	r1 := lazDockFormImportAPI().SysCallN(4, m.Instance())
	return AsControl(r1)
}

func (m *TLazDockForm) FindHeader(x, y int32, OutPart *TLazDockHeaderPart) IControl {
	var result1 uintptr
	r1 := lazDockFormImportAPI().SysCallN(3, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(&result1)))
	*OutPart = TLazDockHeaderPart(result1)
	return AsControl(r1)
}

func (m *TLazDockForm) IsDockedControl(Control IControl) bool {
	r1 := lazDockFormImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TLazDockForm) ControlHasTitle(Control IControl) bool {
	r1 := lazDockFormImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TLazDockForm) GetTitleRect(Control IControl) (resultRect TRect) {
	lazDockFormImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(Control), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TLazDockForm) GetTitleOrientation(Control IControl) TDockOrientation {
	r1 := lazDockFormImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Control))
	return TDockOrientation(r1)
}

func LazDockFormClass() TClass {
	ret := lazDockFormImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TLazDockForm) UpdateCaption() {
	lazDockFormImportAPI().SysCallN(9, m.Instance())
}

var (
	lazDockFormImport       *imports.Imports = nil
	lazDockFormImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazDockForm_Class", 0),
		/*1*/ imports.NewTable("LazDockForm_ControlHasTitle", 0),
		/*2*/ imports.NewTable("LazDockForm_Create", 0),
		/*3*/ imports.NewTable("LazDockForm_FindHeader", 0),
		/*4*/ imports.NewTable("LazDockForm_FindMainControlCandidate", 0),
		/*5*/ imports.NewTable("LazDockForm_GetTitleOrientation", 0),
		/*6*/ imports.NewTable("LazDockForm_GetTitleRect", 0),
		/*7*/ imports.NewTable("LazDockForm_IsDockedControl", 0),
		/*8*/ imports.NewTable("LazDockForm_MainControl", 0),
		/*9*/ imports.NewTable("LazDockForm_UpdateCaption", 0),
	}
)

func lazDockFormImportAPI() *imports.Imports {
	if lazDockFormImport == nil {
		lazDockFormImport = NewDefaultImports()
		lazDockFormImport.SetImportTable(lazDockFormImportTables)
		lazDockFormImportTables = nil
	}
	return lazDockFormImport
}
