//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package base

// CheckPtr
//
// Checks if the interface is instantiated, and returns an instance pointer if it has been instantiated.
func CheckPtr(value any) uintptr {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case IBase:
		return value.(IBase).Instance()
	}
	return 0
}

// GetInstance
//
// Simplification of As operation.
func GetInstance(value any) UnsafePointer {
	switch value.(type) {
	case uintptr:
		// an object from a pointer to an existing object instance
		return UnsafePointer(value.(uintptr))
	case UnsafePointer:
		// An object from an unsafe address. Note: Using this function may cause some unknown situations. Use it with caution.
		return value.(UnsafePointer)
	case IBase:
		if value == nil {
			return nil
		}
		// An object from an existing object instance.
		return UnsafePointer(value.(IBase).Instance())
	default:
		return UnsafePointer(getUIntPtr(value))
	}
}

// SetObjectInstance Set instance values for external component creation
func SetObjectInstance(object IBase, instance UnsafePointer) {
	if object == nil {
		return
	}
	object.SetInstance(instance)
}

// GetObjectUintptr Get class pointer
func GetObjectUintptr(object IBase) uintptr {
	if object == nil {
		return 0
	}
	switch object.(type) {
	case IBase:
		return object.Instance()
	}
	return 0
}

func getUIntPtr(v any) uintptr {
	switch v.(type) {
	case int:
		return uintptr(v.(int))
	case uint:
		return uintptr(v.(uint))
	case int8:
		return uintptr(v.(int8))
	case uint8:
		return uintptr(v.(uint8))
	case int16:
		return uintptr(v.(int16))
	case uint16:
		return uintptr(v.(uint16))
	case int32:
		return uintptr(v.(int32))
	case uint32:
		return uintptr(v.(uint32))
	case int64:
		return uintptr(v.(int64))
	case uint64:
		return uintptr(v.(uint64))
	}
	return 0
}
