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
	"github.com/energye/lcl/types"
)

// IOnCreate TForm OnCreate
type IOnCreate interface {
	FormCreate(sender IObject)
}

// IOnAfterCreate TForm OnAfterCreate
type IOnAfterCreate interface {
	FormAfterCreate(sender IObject)
}

// IOnCreateParams TForm CreateParams
type IOnCreateParams interface {
	CreateParams(params *types.TCreateParams)
}

// IOnCloseQuery TForm OnCloseQuery
type IOnCloseQuery interface {
	OnCloseQuery(sender IObject, canClose *bool)
}

// IOnClose TForm OnClose
type IOnClose interface {
	OnClose(sender IObject, closeAction *types.TCloseAction)
}
