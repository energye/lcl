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

// ICustomLabeledEdit Parent: ICustomEdit
type ICustomLabeledEdit interface {
	ICustomEdit
	EditLabel() IBoundLabel                      // property EditLabel Getter
	LabelPosition() types.TLabelPosition         // property LabelPosition Getter
	SetLabelPosition(value types.TLabelPosition) // property LabelPosition Setter
	LabelSpacing() int32                         // property LabelSpacing Getter
	SetLabelSpacing(value int32)                 // property LabelSpacing Setter
}

type TCustomLabeledEdit struct {
	TCustomEdit
}

func (m *TCustomLabeledEdit) EditLabel() IBoundLabel {
	if !m.IsValid() {
		return nil
	}
	r := customLabeledEditAPI().SysCallN(1, m.Instance())
	return AsBoundLabel(r)
}

func (m *TCustomLabeledEdit) LabelPosition() types.TLabelPosition {
	if !m.IsValid() {
		return 0
	}
	r := customLabeledEditAPI().SysCallN(2, 0, m.Instance())
	return types.TLabelPosition(r)
}

func (m *TCustomLabeledEdit) SetLabelPosition(value types.TLabelPosition) {
	if !m.IsValid() {
		return
	}
	customLabeledEditAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomLabeledEdit) LabelSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customLabeledEditAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCustomLabeledEdit) SetLabelSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	customLabeledEditAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// NewCustomLabeledEdit class constructor
func NewCustomLabeledEdit(theOwner IComponent) ICustomLabeledEdit {
	r := customLabeledEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomLabeledEdit(r)
}

func TCustomLabeledEditClass() types.TClass {
	r := customLabeledEditAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	customLabeledEditOnce   base.Once
	customLabeledEditImport *imports.Imports = nil
)

func customLabeledEditAPI() *imports.Imports {
	customLabeledEditOnce.Do(func() {
		customLabeledEditImport = api.NewDefaultImports()
		customLabeledEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomLabeledEdit_Create", 0), // constructor NewCustomLabeledEdit
			/* 1 */ imports.NewTable("TCustomLabeledEdit_EditLabel", 0), // property EditLabel
			/* 2 */ imports.NewTable("TCustomLabeledEdit_LabelPosition", 0), // property LabelPosition
			/* 3 */ imports.NewTable("TCustomLabeledEdit_LabelSpacing", 0), // property LabelSpacing
			/* 4 */ imports.NewTable("TCustomLabeledEdit_TClass", 0), // function TCustomLabeledEditClass
		}
	})
	return customLabeledEditImport
}
