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

// IItemProp Parent: IPersistent
type IItemProp interface {
	IPersistent
	// EditMask
	//  function HasPickList: Boolean;
	EditMask() string                    // property EditMask Getter
	SetEditMask(value string)            // property EditMask Setter
	EditStyle() types.TEditStyle         // property EditStyle Getter
	SetEditStyle(value types.TEditStyle) // property EditStyle Setter
	KeyDesc() string                     // property KeyDesc Getter
	SetKeyDesc(value string)             // property KeyDesc Setter
	PickList() IStrings                  // property PickList Getter
	SetPickList(value IStrings)          // property PickList Setter
	MaxLength() int32                    // property MaxLength Getter
	SetMaxLength(value int32)            // property MaxLength Setter
	ReadOnly() bool                      // property ReadOnly Getter
	SetReadOnly(value bool)              // property ReadOnly Setter
}

type TItemProp struct {
	TPersistent
}

func (m *TItemProp) EditMask() string {
	if !m.IsValid() {
		return ""
	}
	r := itemPropAPI().SysCallN(1, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TItemProp) SetEditMask(value string) {
	if !m.IsValid() {
		return
	}
	itemPropAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

func (m *TItemProp) EditStyle() types.TEditStyle {
	if !m.IsValid() {
		return 0
	}
	r := itemPropAPI().SysCallN(2, 0, m.Instance())
	return types.TEditStyle(r)
}

func (m *TItemProp) SetEditStyle(value types.TEditStyle) {
	if !m.IsValid() {
		return
	}
	itemPropAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TItemProp) KeyDesc() string {
	if !m.IsValid() {
		return ""
	}
	r := itemPropAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TItemProp) SetKeyDesc(value string) {
	if !m.IsValid() {
		return
	}
	itemPropAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TItemProp) PickList() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := itemPropAPI().SysCallN(4, 0, m.Instance())
	return AsStrings(r)
}

func (m *TItemProp) SetPickList(value IStrings) {
	if !m.IsValid() {
		return
	}
	itemPropAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TItemProp) MaxLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := itemPropAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TItemProp) SetMaxLength(value int32) {
	if !m.IsValid() {
		return
	}
	itemPropAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TItemProp) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := itemPropAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TItemProp) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	itemPropAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

// NewItemProp class constructor
func NewItemProp(owner IValueListEditor) IItemProp {
	r := itemPropAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsItemProp(r)
}

var (
	itemPropOnce   base.Once
	itemPropImport *imports.Imports = nil
)

func itemPropAPI() *imports.Imports {
	itemPropOnce.Do(func() {
		itemPropImport = api.NewDefaultImports()
		itemPropImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TItemProp_Create", 0), // constructor NewItemProp
			/* 1 */ imports.NewTable("TItemProp_EditMask", 0), // property EditMask
			/* 2 */ imports.NewTable("TItemProp_EditStyle", 0), // property EditStyle
			/* 3 */ imports.NewTable("TItemProp_KeyDesc", 0), // property KeyDesc
			/* 4 */ imports.NewTable("TItemProp_PickList", 0), // property PickList
			/* 5 */ imports.NewTable("TItemProp_MaxLength", 0), // property MaxLength
			/* 6 */ imports.NewTable("TItemProp_ReadOnly", 0), // property ReadOnly
		}
	})
	return itemPropImport
}
