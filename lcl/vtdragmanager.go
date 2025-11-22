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

// IVTDragManager Parent: IInterfacedObject
type IVTDragManager interface {
	IInterfacedObject
	DragLeave() types.HRESULT                                                // function
	DragOver(keyState uint32, pt types.TPoint, effect *uint32) types.HRESULT // function
	GiveFeedback(effect uint32) types.HRESULT                                // function
	QueryContinueDrag(escapePressed bool, keyState uint32) types.HRESULT     // function
	ForceDragLeave()                                                         // procedure
	AsIntfVTDragManager() uintptr
}

type TVTDragManager struct {
	TInterfacedObject
}

func (m *TVTDragManager) DragLeave() types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := vTDragManagerAPI().SysCallN(1, m.Instance())
	return types.HRESULT(r)
}

func (m *TVTDragManager) DragOver(keyState uint32, pt types.TPoint, effect *uint32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	effectPtr := uintptr(*effect)
	r := vTDragManagerAPI().SysCallN(2, m.Instance(), uintptr(keyState), uintptr(base.UnsafePointer(&pt)), uintptr(base.UnsafePointer(&effectPtr)))
	*effect = uint32(effectPtr)
	return types.HRESULT(r)
}

func (m *TVTDragManager) GiveFeedback(effect uint32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := vTDragManagerAPI().SysCallN(3, m.Instance(), uintptr(effect))
	return types.HRESULT(r)
}

func (m *TVTDragManager) QueryContinueDrag(escapePressed bool, keyState uint32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := vTDragManagerAPI().SysCallN(4, m.Instance(), api.PasBool(escapePressed), uintptr(keyState))
	return types.HRESULT(r)
}

func (m *TVTDragManager) ForceDragLeave() {
	if !m.IsValid() {
		return
	}
	vTDragManagerAPI().SysCallN(5, m.Instance())
}

func (m *TVTDragManager) AsIntfVTDragManager() uintptr {
	return m.GetIntfPointer(0)
}

// NewVTDragManager class constructor
func NewVTDragManager(owner IBaseVirtualTree) IVTDragManager {
	var vTDragManagerPtr uintptr // IVTDragManager
	r := vTDragManagerAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&vTDragManagerPtr)))
	ret := AsVTDragManager(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, vTDragManagerPtr)
	}
	return ret
}

var (
	vTDragManagerOnce   base.Once
	vTDragManagerImport *imports.Imports = nil
)

func vTDragManagerAPI() *imports.Imports {
	vTDragManagerOnce.Do(func() {
		vTDragManagerImport = api.NewDefaultImports()
		vTDragManagerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTDragManager_Create", 0), // constructor NewVTDragManager
			/* 1 */ imports.NewTable("TVTDragManager_DragLeave", 0), // function DragLeave
			/* 2 */ imports.NewTable("TVTDragManager_DragOver", 0), // function DragOver
			/* 3 */ imports.NewTable("TVTDragManager_GiveFeedback", 0), // function GiveFeedback
			/* 4 */ imports.NewTable("TVTDragManager_QueryContinueDrag", 0), // function QueryContinueDrag
			/* 5 */ imports.NewTable("TVTDragManager_ForceDragLeave", 0), // procedure ForceDragLeave
		}
	})
	return vTDragManagerImport
}
