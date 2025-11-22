//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import (
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

type IDataObject uintptr

type TFormatEtc struct {
	CfFormat types.Word // {TCLIPFORMAT};
	Ptd      uintptr    // ^tagDVTARGETDEVICE
	dwAspect types.DWORD
	lindex   int32
	tymed    types.DWORD
}

type IUnknown = base.IBase

// IVTEditLink
// It is an abstract interface that needs to inherit the ICustomVTEditLink implementation when used,
type IVTEditLink interface {
	IInterfacedObject
}

func AsVTEditLink(instance uintptr) IVTEditLink {
	if instance == 0 {
		return nil
	}
	m := new(TInterfacedObject)
	base.SetObjectInstance(m, base.UnsafePointer(instance))
	return m
}
