//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"fmt"
	"github.com/energye/lcl/pkgs/win"
)

func showError(err interface{}) {
	win.MessageBox(0, fmt.Sprint(err), "Error", win.MB_ICONERROR)
}
