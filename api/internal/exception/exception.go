// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package exception

type Callback func(exceptionID, message string)

var (
	exceptionCallback Callback
	Debug             bool
)

func SetOnException(fn Callback) {
	exceptionCallback = fn
}

func CallOnException(exceptionID string, message string) bool {
	if exceptionCallback != nil {
		exceptionCallback(exceptionID, message)
		return true
	}
	return false
}
