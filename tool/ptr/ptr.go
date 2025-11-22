//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package ptr

import "unsafe"

const ptrSize = unsafe.Sizeof(uintptr(0))

// GetParamOf 获取参数指针
func GetParamOf(index int, ptr uintptr) uintptr {
	//return *(*uintptr)(unsafe.Pointer(ptr + uintptr(index)*unsafe.Sizeof(ptr)))
	return *(*uintptr)(unsafe.Pointer(ptr + uintptr(index)*ptrSize))
}

// GetParamPtr 根据指定指针位置开始 偏移获取指针
func GetParamPtr(ptr uintptr, offset int) unsafe.Pointer {
	return unsafe.Pointer(ptr + uintptr(offset))
}

// GetPointerOffset 指针偏移地址
func GetPointerOffset(ptr uintptr, offset uintptr) unsafe.Pointer {
	return unsafe.Pointer(ptr + offset)
}
