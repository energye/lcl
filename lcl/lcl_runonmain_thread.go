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
//	Execute in main thread, async queue.
func RunOnMainThreadAsync(callback TMainThreadAsyncProc) int {
	if Application.Initialed() {
		id := queueAsyncCall.next(callback)
		api.RunOnMainAsyncCall(id)
		return id
	} else {
		callback(0)
	}
	return 0
}

// RunOnMainThreadSync
//
//	Execute in main thread, synchronous.
func RunOnMainThreadSync(callback TMainThreadSyncProc) {
	if Application.Initialed() {
		api.RunOnMainSyncCall(callback)
	} else {
		callback()
	}
}
