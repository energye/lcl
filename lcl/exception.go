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

// IException Parent: IObject
type IException interface {
	IObject
	HelpContext() int32          // property
	SetHelpContext(AValue int32) // property
	Message() string             // property
	SetMessage(AValue string)    // property
}

// Exception Parent: TObject
type Exception struct {
	TObject
}

func NewException(msg string) IException {
	r1 := LCL().SysCallN(2845, PascalStr(msg))
	return AsException(r1)
}

func NewExceptionHelp(Msg string, AHelpContext int32) IException {
	r1 := LCL().SysCallN(2846, PascalStr(Msg), uintptr(AHelpContext))
	return AsException(r1)
}

func (m *Exception) HelpContext() int32 {
	r1 := LCL().SysCallN(2847, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *Exception) SetHelpContext(AValue int32) {
	LCL().SysCallN(2847, 1, m.Instance(), uintptr(AValue))
}

func (m *Exception) Message() string {
	r1 := LCL().SysCallN(2848, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *Exception) SetMessage(AValue string) {
	LCL().SysCallN(2848, 1, m.Instance(), PascalStr(AValue))
}

func ExceptionClass() TClass {
	ret := LCL().SysCallN(2844)
	return TClass(ret)
}
