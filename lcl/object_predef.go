//----------------------------------------
// The code is automatically generated by the generate-tool.
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"unsafe"
)

// TIs : Is操作符
type TIs uintptr

func (m *TObject) Instance() uintptr {
	if m.instance == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TObject) UnsafeAddr() unsafe.Pointer {
	return m.instance
}

func (m *TObject) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *TObject) Is() TIs {
	return TIs(m.Instance())
}

func (m *TObject) SetInstance(instance unsafe.Pointer) {
	if uintptr(m.instance) != uintptr(instance) {
		m.instance = instance
	}
}

func (m *TObject) free(index int) {
	if m.instance != nil {
		api.LCL().SysCallN(index, m.Instance())
		m.instance = nil
	}
}