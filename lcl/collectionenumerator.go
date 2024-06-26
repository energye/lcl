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

// ICollectionEnumerator Parent: IObject
type ICollectionEnumerator interface {
	IObject
	Current() ICollectionItem    // property
	GetCurrent() ICollectionItem // function
	MoveNext() bool              // function
}

// TCollectionEnumerator Parent: TObject
type TCollectionEnumerator struct {
	TObject
}

func NewCollectionEnumerator(ACollection ICollection) ICollectionEnumerator {
	r1 := LCL().SysCallN(688, GetObjectUintptr(ACollection))
	return AsCollectionEnumerator(r1)
}

func (m *TCollectionEnumerator) Current() ICollectionItem {
	r1 := LCL().SysCallN(689, m.Instance())
	return AsCollectionItem(r1)
}

func (m *TCollectionEnumerator) GetCurrent() ICollectionItem {
	r1 := LCL().SysCallN(690, m.Instance())
	return AsCollectionItem(r1)
}

func (m *TCollectionEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(691, m.Instance())
	return GoBool(r1)
}

func CollectionEnumeratorClass() TClass {
	ret := LCL().SysCallN(687)
	return TClass(ret)
}
