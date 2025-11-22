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

// ICustomFrame Parent: ICustomDesignControl
type ICustomFrame interface {
	ICustomDesignControl
	ParentBackground() bool         // property ParentBackground Getter
	SetParentBackground(value bool) // property ParentBackground Setter
}

type TCustomFrame struct {
	TCustomDesignControl
}

func (m *TCustomFrame) ParentBackground() bool {
	if !m.IsValid() {
		return false
	}
	r := customFrameAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomFrame) SetParentBackground(value bool) {
	if !m.IsValid() {
		return
	}
	customFrameAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

// NewCustomFrame class constructor
func NewCustomFrame(owner IComponent) ICustomFrame {
	r := customFrameAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomFrame(r)
}

func TCustomFrameClass() types.TClass {
	r := customFrameAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	customFrameOnce   base.Once
	customFrameImport *imports.Imports = nil
)

func customFrameAPI() *imports.Imports {
	customFrameOnce.Do(func() {
		customFrameImport = api.NewDefaultImports()
		customFrameImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomFrame_Create", 0), // constructor NewCustomFrame
			/* 1 */ imports.NewTable("TCustomFrame_ParentBackground", 0), // property ParentBackground
			/* 2 */ imports.NewTable("TCustomFrame_TClass", 0), // function TCustomFrameClass
		}
	})
	return customFrameImport
}
