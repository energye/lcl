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

// ISaveDialog Parent: IOpenDialog
type ISaveDialog interface {
	IOpenDialog
}

// TSaveDialog Parent: TOpenDialog
type TSaveDialog struct {
	TOpenDialog
}

func NewSaveDialog(AOwner IComponent) ISaveDialog {
	r1 := LCL().SysCallN(4889, GetObjectUintptr(AOwner))
	return AsSaveDialog(r1)
}

func SaveDialogClass() TClass {
	ret := LCL().SysCallN(4888)
	return TClass(ret)
}
