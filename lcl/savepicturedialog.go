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

// ISavePictureDialog Parent: IOpenPictureDialog
type ISavePictureDialog interface {
	IOpenPictureDialog
}

// TSavePictureDialog Parent: TOpenPictureDialog
type TSavePictureDialog struct {
	TOpenPictureDialog
}

func NewSavePictureDialog(TheOwner IComponent) ISavePictureDialog {
	r1 := LCL().SysCallN(4891, GetObjectUintptr(TheOwner))
	return AsSavePictureDialog(r1)
}

func SavePictureDialogClass() TClass {
	ret := LCL().SysCallN(4890)
	return TClass(ret)
}
