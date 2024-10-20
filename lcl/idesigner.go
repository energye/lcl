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

// IIDesigner Is Abstract Class Parent: IObject
type IIDesigner interface {
	IObject
	LookupRoot() IComponent                                        // property
	DefaultFormBoundsValid() bool                                  // property
	SetDefaultFormBoundsValid(AValue bool)                         // property
	IsDesignMsg(Sender IControl, Message *TLMessage) bool          // function Is Abstract
	GetShiftState() TShiftState                                    // function Is Abstract
	UniqueName(BaseName string) string                             // function Is Abstract
	UTF8KeyPress(UTF8Key *TUTF8Char)                               // procedure Is Abstract
	Modified()                                                     // procedure Is Abstract
	Notification(AComponent IComponent, Operation TOperation)      // procedure Is Abstract
	PaintGrid()                                                    // procedure Is Abstract
	ValidateRename(AComponent IComponent, CurName, NewName string) // procedure Is Abstract
	SelectOnlyThisComponent(AComponent IComponent)                 // procedure Is Abstract
	PrepareFreeDesigner(AFreeComponent bool)                       // procedure Is Abstract
}

// TIDesigner Is Abstract Class Parent: TObject
type TIDesigner struct {
	TObject
}

func (m *TIDesigner) LookupRoot() IComponent {
	r1 := designerImportAPI().SysCallN(4, m.Instance())
	return AsComponent(r1)
}

func (m *TIDesigner) DefaultFormBoundsValid() bool {
	r1 := designerImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIDesigner) SetDefaultFormBoundsValid(AValue bool) {
	designerImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TIDesigner) IsDesignMsg(Sender IControl, Message *TLMessage) bool {
	var result1 uintptr
	r1 := designerImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*Message = *(*TLMessage)(getPointer(result1))
	return GoBool(r1)
}

func (m *TIDesigner) GetShiftState() TShiftState {
	r1 := designerImportAPI().SysCallN(2, m.Instance())
	return TShiftState(r1)
}

func (m *TIDesigner) UniqueName(BaseName string) string {
	r1 := designerImportAPI().SysCallN(11, m.Instance(), PascalStr(BaseName))
	return GoStr(r1)
}

func IDesignerClass() TClass {
	ret := designerImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TIDesigner) UTF8KeyPress(UTF8Key *TUTF8Char) {
	var result0 uintptr
	designerImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(&result0)))
	*UTF8Key = *(*TUTF8Char)(getPointer(result0))
}

func (m *TIDesigner) Modified() {
	designerImportAPI().SysCallN(5, m.Instance())
}

func (m *TIDesigner) Notification(AComponent IComponent, Operation TOperation) {
	designerImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(AComponent), uintptr(Operation))
}

func (m *TIDesigner) PaintGrid() {
	designerImportAPI().SysCallN(7, m.Instance())
}

func (m *TIDesigner) ValidateRename(AComponent IComponent, CurName, NewName string) {
	designerImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(AComponent), PascalStr(CurName), PascalStr(NewName))
}

func (m *TIDesigner) SelectOnlyThisComponent(AComponent IComponent) {
	designerImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TIDesigner) PrepareFreeDesigner(AFreeComponent bool) {
	designerImportAPI().SysCallN(8, m.Instance(), PascalBool(AFreeComponent))
}

var (
	designerImport       *imports.Imports = nil
	designerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("IDesigner_Class", 0),
		/*1*/ imports.NewTable("IDesigner_DefaultFormBoundsValid", 0),
		/*2*/ imports.NewTable("IDesigner_GetShiftState", 0),
		/*3*/ imports.NewTable("IDesigner_IsDesignMsg", 0),
		/*4*/ imports.NewTable("IDesigner_LookupRoot", 0),
		/*5*/ imports.NewTable("IDesigner_Modified", 0),
		/*6*/ imports.NewTable("IDesigner_Notification", 0),
		/*7*/ imports.NewTable("IDesigner_PaintGrid", 0),
		/*8*/ imports.NewTable("IDesigner_PrepareFreeDesigner", 0),
		/*9*/ imports.NewTable("IDesigner_SelectOnlyThisComponent", 0),
		/*10*/ imports.NewTable("IDesigner_UTF8KeyPress", 0),
		/*11*/ imports.NewTable("IDesigner_UniqueName", 0),
		/*12*/ imports.NewTable("IDesigner_ValidateRename", 0),
	}
)

func designerImportAPI() *imports.Imports {
	if designerImport == nil {
		designerImport = NewDefaultImports()
		designerImport.SetImportTable(designerImportTables)
		designerImportTables = nil
	}
	return designerImport
}
