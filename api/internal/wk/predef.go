//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wk

import "github.com/energye/lcl/api/imports"

var wkImportDefs = []*imports.Table{
	imports.NewTable("WVLoader_SetOnEnvironmentCreated", 0),
}

// InitWKPreDefsImport 初始化预定义api
func InitWKPreDefsImport(imp *imports.Imports) {
	imp.SetImportTable(wkImportDefs)
	imp.SetOk(true)
}
