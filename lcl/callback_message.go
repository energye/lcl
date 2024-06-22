//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"unsafe"

	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
)

func messageCallbackProc(f uintptr, msg uintptr) uintptr {
	v := PtrToElementValue(f)
	if v != nil {
		v.(TWndProcEvent)(
			(*types.TMessage)(unsafe.Pointer(msg)),
		)
	}
	return 0
}
