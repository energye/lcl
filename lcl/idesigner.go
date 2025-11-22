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
)

// IIDesigner Parent: IObject
type IIDesigner interface {
	IObject
	LookupRoot() IComponent               // property LookupRoot Getter
	DefaultFormBoundsValid() bool         // property DefaultFormBoundsValid Getter
	SetDefaultFormBoundsValid(value bool) // property DefaultFormBoundsValid Setter
}

type TIDesigner struct {
	TObject
}

func (m *TIDesigner) LookupRoot() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := iDesignerAPI().SysCallN(0, m.Instance())
	return AsComponent(r)
}

func (m *TIDesigner) DefaultFormBoundsValid() bool {
	if !m.IsValid() {
		return false
	}
	r := iDesignerAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TIDesigner) SetDefaultFormBoundsValid(value bool) {
	if !m.IsValid() {
		return
	}
	iDesignerAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

var (
	iDesignerOnce   base.Once
	iDesignerImport *imports.Imports = nil
)

func iDesignerAPI() *imports.Imports {
	iDesignerOnce.Do(func() {
		iDesignerImport = api.NewDefaultImports()
		iDesignerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TIDesigner_LookupRoot", 0), // property LookupRoot
			/* 1 */ imports.NewTable("TIDesigner_DefaultFormBoundsValid", 0), // property DefaultFormBoundsValid
		}
	})
	return iDesignerImport
}
