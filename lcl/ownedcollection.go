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

// IOwnedCollection Parent: ICollection
type IOwnedCollection interface {
	ICollection
}

// TOwnedCollection Parent: TCollection
type TOwnedCollection struct {
	TCollection
}

func NewOwnedCollection(AOwner IPersistent, AItemClass TCollectionItemClass) IOwnedCollection {
	r1 := LCL().SysCallN(4425, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsOwnedCollection(r1)
}

func OwnedCollectionClass() TClass {
	ret := LCL().SysCallN(4424)
	return TClass(ret)
}
