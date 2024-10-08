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

// IOpenPictureDialog Parent: IPreviewFileDialog
type IOpenPictureDialog interface {
	IPreviewFileDialog
	DefaultFilter() string // property
	GetFilterExt() string  // function
}

// TOpenPictureDialog Parent: TPreviewFileDialog
type TOpenPictureDialog struct {
	TPreviewFileDialog
}

func NewOpenPictureDialog(TheOwner IComponent) IOpenPictureDialog {
	r1 := LCL().SysCallN(4421, GetObjectUintptr(TheOwner))
	return AsOpenPictureDialog(r1)
}

func (m *TOpenPictureDialog) DefaultFilter() string {
	r1 := LCL().SysCallN(4422, m.Instance())
	return GoStr(r1)
}

func (m *TOpenPictureDialog) GetFilterExt() string {
	r1 := LCL().SysCallN(4423, m.Instance())
	return GoStr(r1)
}

func OpenPictureDialogClass() TClass {
	ret := LCL().SysCallN(4420)
	return TClass(ret)
}
