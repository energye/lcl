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
	r1 := LCL().SysCallN(3363, m.Instance())
	return AsComponent(r1)
}

func (m *TIDesigner) DefaultFormBoundsValid() bool {
	r1 := LCL().SysCallN(3360, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIDesigner) SetDefaultFormBoundsValid(AValue bool) {
	LCL().SysCallN(3360, 1, m.Instance(), PascalBool(AValue))
}

func (m *TIDesigner) IsDesignMsg(Sender IControl, Message *TLMessage) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(3362, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*Message = *(*TLMessage)(getPointer(result1))
	return GoBool(r1)
}

func (m *TIDesigner) GetShiftState() TShiftState {
	r1 := LCL().SysCallN(3361, m.Instance())
	return TShiftState(r1)
}

func (m *TIDesigner) UniqueName(BaseName string) string {
	r1 := LCL().SysCallN(3370, m.Instance(), PascalStr(BaseName))
	return GoStr(r1)
}

func IDesignerClass() TClass {
	ret := LCL().SysCallN(3359)
	return TClass(ret)
}

func (m *TIDesigner) UTF8KeyPress(UTF8Key *TUTF8Char) {
	var result0 uintptr
	LCL().SysCallN(3369, m.Instance(), uintptr(unsafePointer(&result0)))
	*UTF8Key = *(*TUTF8Char)(getPointer(result0))
}

func (m *TIDesigner) Modified() {
	LCL().SysCallN(3364, m.Instance())
}

func (m *TIDesigner) Notification(AComponent IComponent, Operation TOperation) {
	LCL().SysCallN(3365, m.Instance(), GetObjectUintptr(AComponent), uintptr(Operation))
}

func (m *TIDesigner) PaintGrid() {
	LCL().SysCallN(3366, m.Instance())
}

func (m *TIDesigner) ValidateRename(AComponent IComponent, CurName, NewName string) {
	LCL().SysCallN(3371, m.Instance(), GetObjectUintptr(AComponent), PascalStr(CurName), PascalStr(NewName))
}

func (m *TIDesigner) SelectOnlyThisComponent(AComponent IComponent) {
	LCL().SysCallN(3368, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TIDesigner) PrepareFreeDesigner(AFreeComponent bool) {
	LCL().SysCallN(3367, m.Instance(), PascalBool(AFreeComponent))
}
