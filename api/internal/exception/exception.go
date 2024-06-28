//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package exception

type Callback func(funcName, message string)

var HandlerCallback Callback

// CallException 外部主动调用异常回调
func CallException(funcName, message string) {
	if HandlerCallback != nil {
		HandlerCallback(funcName, message)
	}
}
