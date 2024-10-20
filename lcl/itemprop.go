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

// IItemProp Parent: IPersistent
type IItemProp interface {
	IPersistent
	EditMask() string               // property
	SetEditMask(AValue string)      // property
	EditStyle() TEditStyle          // property
	SetEditStyle(AValue TEditStyle) // property
	KeyDesc() string                // property
	SetKeyDesc(AValue string)       // property
	PickList() IStrings             // property
	SetPickList(AValue IStrings)    // property
	MaxLength() int32               // property
	SetMaxLength(AValue int32)      // property
	ReadOnly() bool                 // property
	SetReadOnly(AValue bool)        // property
}

// TItemProp Parent: TPersistent
type TItemProp struct {
	TPersistent
}

func NewItemProp(AOwner IValueListEditor) IItemProp {
	r1 := temPropImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsItemProp(r1)
}

func (m *TItemProp) EditMask() string {
	r1 := temPropImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TItemProp) SetEditMask(AValue string) {
	temPropImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TItemProp) EditStyle() TEditStyle {
	r1 := temPropImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TEditStyle(r1)
}

func (m *TItemProp) SetEditStyle(AValue TEditStyle) {
	temPropImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TItemProp) KeyDesc() string {
	r1 := temPropImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TItemProp) SetKeyDesc(AValue string) {
	temPropImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TItemProp) PickList() IStrings {
	r1 := temPropImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TItemProp) SetPickList(AValue IStrings) {
	temPropImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TItemProp) MaxLength() int32 {
	r1 := temPropImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TItemProp) SetMaxLength(AValue int32) {
	temPropImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TItemProp) ReadOnly() bool {
	r1 := temPropImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TItemProp) SetReadOnly(AValue bool) {
	temPropImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func ItemPropClass() TClass {
	ret := temPropImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	temPropImport       *imports.Imports = nil
	temPropImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ItemProp_Class", 0),
		/*1*/ imports.NewTable("ItemProp_Create", 0),
		/*2*/ imports.NewTable("ItemProp_EditMask", 0),
		/*3*/ imports.NewTable("ItemProp_EditStyle", 0),
		/*4*/ imports.NewTable("ItemProp_KeyDesc", 0),
		/*5*/ imports.NewTable("ItemProp_MaxLength", 0),
		/*6*/ imports.NewTable("ItemProp_PickList", 0),
		/*7*/ imports.NewTable("ItemProp_ReadOnly", 0),
	}
)

func temPropImportAPI() *imports.Imports {
	if temPropImport == nil {
		temPropImport = NewDefaultImports()
		temPropImport.SetImportTable(temPropImportTables)
		temPropImportTables = nil
	}
	return temPropImport
}
