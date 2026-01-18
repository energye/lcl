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

// IOnFormCreate TForm OnCreate
type IOnFormCreate interface {
	FormCreate(sender IObject)
}

// IOnFormAfterCreate TForm OnAfterCreate
type IOnFormAfterCreate interface {
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

// IOnShow TForm OnShow
type IOnShow interface {
	OnShow(sender IObject)
}
