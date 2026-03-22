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

// ISynEditFriend Parent: IComponent
type ISynEditFriend interface {
	IComponent
}

type TSynEditFriend struct {
	TComponent
}

// NewSynEditFriend class constructor
func NewSynEditFriend(owner IComponent) ISynEditFriend {
	r := synEditFriendAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynEditFriend(r)
}

func TSynEditFriendClass() types.TClass {
	r := synEditFriendAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	synEditFriendOnce   base.Once
	synEditFriendImport *imports.Imports = nil
)

func synEditFriendAPI() *imports.Imports {
	synEditFriendOnce.Do(func() {
		synEditFriendImport = api.NewDefaultImports()
		synEditFriendImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditFriend_Create", 0), // constructor NewSynEditFriend
			/* 1 */ imports.NewTable("TSynEditFriend_TClass", 0), // function TSynEditFriendClass
		}
	})
	return synEditFriendImport
}
