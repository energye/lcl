//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"container/list"
	"sync"
	"unsafe"
)

var (
	// 防止GC的表
	events = list.New()

	// ThreadSync
	threadSyncProc func()

	// sync
	idRefs, threadSyncRef sync.Mutex
)

func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

func PasBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

// typedef struct { void *type; void *value; } GoInterface;
type interfacePtr struct {
	tpy uintptr
	val uintptr
}

func interfaceNotNil(v interface{}) bool {
	ptr := (*interfacePtr)(unsafe.Pointer(&v))
	return ptr != nil && ptr.tpy != 0 && ptr.val != 0
}

func PtrToElementPtr(v uintptr) *list.Element {
	if v == 0 {
		return nil
	}
	return (*list.Element)(unsafe.Pointer(v))
}

func PtrToElementValue(v uintptr) interface{} {
	element := PtrToElementPtr(v)
	if element != nil {
		return element.Value
	}
	return nil
}

func removeEventElement(v *list.Element) bool {
	if v != nil {
		idRefs.Lock()
		defer idRefs.Unlock()
		events.Remove(v)
		return true
	}
	return false
}

func RemoveEventElement(v uintptr) bool {
	return removeEventElement(PtrToElementPtr(v))
}

func MakeEventDataPtr(fn interface{}) uintptr {
	idRefs.Lock()
	defer idRefs.Unlock()
	if fn == nil {
		return 0
	}
	if interfaceNotNil(fn) {
		return uintptr(unsafe.Pointer(events.PushBack(fn)))
	}
	return 0
}

func ThreadSyncCallbackFn() func() {
	return threadSyncProc
}
