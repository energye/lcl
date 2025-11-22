//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import "github.com/energye/lcl/base"

// AsUnknown Convert a pointer object to an existing class object
func AsUnknown(obj any) IUnknown {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(base.TBase)
	base.SetObjectInstance(result, instance)
	return result
}
