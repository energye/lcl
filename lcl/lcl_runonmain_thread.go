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
	"github.com/energye/lcl/api"
)

// RunOnMainThreadAsync
//
//	主线程中执行, 异步
func RunOnMainThreadAsync(callback TMainThreadAsyncProc) int {
	id := queueAsyncCall.next(callback)
	api.RunOnMainAsyncCall(id)
	return id
}

// RunOnMainThreadSync
//
//	主线程中执行, 同步
func RunOnMainThreadSync(callback TMainThreadSyncProc) {
	api.RunOnMainSyncCall(callback)
}
