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

// ICustomLabeledEdit Parent: ICustomEdit
type ICustomLabeledEdit interface {
	ICustomEdit
	EditLabel() IBoundLabel                 // property
	LabelPosition() TLabelPosition          // property
	SetLabelPosition(AValue TLabelPosition) // property
	LabelSpacing() int32                    // property
	SetLabelSpacing(AValue int32)           // property
}

// TCustomLabeledEdit Parent: TCustomEdit
type TCustomLabeledEdit struct {
	TCustomEdit
}

func NewCustomLabeledEdit(TheOwner IComponent) ICustomLabeledEdit {
	r1 := customLabeledEditImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsCustomLabeledEdit(r1)
}

func (m *TCustomLabeledEdit) EditLabel() IBoundLabel {
	r1 := customLabeledEditImportAPI().SysCallN(2, m.Instance())
	return AsBoundLabel(r1)
}

func (m *TCustomLabeledEdit) LabelPosition() TLabelPosition {
	r1 := customLabeledEditImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TLabelPosition(r1)
}

func (m *TCustomLabeledEdit) SetLabelPosition(AValue TLabelPosition) {
	customLabeledEditImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomLabeledEdit) LabelSpacing() int32 {
	r1 := customLabeledEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomLabeledEdit) SetLabelSpacing(AValue int32) {
	customLabeledEditImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func CustomLabeledEditClass() TClass {
	ret := customLabeledEditImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customLabeledEditImport       *imports.Imports = nil
	customLabeledEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomLabeledEdit_Class", 0),
		/*1*/ imports.NewTable("CustomLabeledEdit_Create", 0),
		/*2*/ imports.NewTable("CustomLabeledEdit_EditLabel", 0),
		/*3*/ imports.NewTable("CustomLabeledEdit_LabelPosition", 0),
		/*4*/ imports.NewTable("CustomLabeledEdit_LabelSpacing", 0),
	}
)

func customLabeledEditImportAPI() *imports.Imports {
	if customLabeledEditImport == nil {
		customLabeledEditImport = NewDefaultImports()
		customLabeledEditImport.SetImportTable(customLabeledEditImportTables)
		customLabeledEditImportTables = nil
	}
	return customLabeledEditImport
}
