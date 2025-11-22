//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package base

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"sync"
)

// Once 别名
type Once = sync.Once

type IEvents interface {
	ClearEvents()
}

type IIntfs interface {
	SetIntfPointer(indexType int32, ptr uintptr)
	GetIntfPointer(indexType int32) uintptr
	Create(size int)
}

// IBase 所有结构的基接口
type IBase interface {
	Instance() uintptr                  // 返回实例
	SetInstance(instance UnsafePointer) // 设置实例
	UnsafeAddr() UnsafePointer          // 返回实例指针
	IsValid() bool                      // 返回当前对象是否有效
	//FreeAndNil()                        // 将当前实现释放并设置nil
	//Nil()  // 将当前实例设置为nil
	// Free 直接释放对象资源, 适用于 Class 类型指针
	Free()
	// 设置事件
	setEvent(apiIndex int, apiTable *imports.Imports, eventPtr uintptr)
}

// 事件列表
type eventList struct {
	apiTable *imports.Imports
	ptrs     map[int]uintptr
}

type TBase struct {
	instance UnsafePointer
	// events 事件所属对象分类集合
	// key: apiIndex + apiTable.Type，也是当前对象事件的唯一类型
	// value: 事件列表
	events map[int32]*eventList
	// 实例接口指针列表
	// key: 当前实例的接口类型 value: 接口指针地址
	intfs []uintptr
}

// SetEvent 给带有事件的组件设置事件
// base: 基接口
// apiIndex: 事件API表索引，也是当前对象事件的唯一类型
// apiTable: API表
// eventPtr: 事件指针
func SetEvent(base IBase, apiIndex int, apiTable *imports.Imports, eventPtr uintptr) {
	base.setEvent(apiIndex, apiTable, eventPtr)
}

func (m *TBase) Create(size int) {
	m.intfs = make([]uintptr, size)
}

func (m *TBase) SetIntfPointer(indexType int32, ptr uintptr) {
	m.intfs[indexType] = ptr
}

func (m *TBase) GetIntfPointer(indexType int32) uintptr {
	return m.intfs[indexType]
}

func (m *TBase) ClearEvents() {
	if !m.IsValid() {
		return
	}
	// 移除所有事件
	for _, list := range m.events {
		for apiIndex, eventPtr := range list.ptrs {
			if eventPtr != 0 {
				api.RemoveEventElement(eventPtr)
				if list.apiTable != nil {
					list.apiTable.SysCallN(apiIndex, m.Instance(), 0)
				}
			}
		}
		list.apiTable = nil
		list.ptrs = nil
	}
	m.events = nil
}

func (m *TBase) setEvent(apiIndex int, apiTable *imports.Imports, eventPtr uintptr) {
	if !m.IsValid() || apiTable == nil {
		return
	}
	if m.events == nil {
		m.events = make(map[int32]*eventList)
	}
	var (
		classType = apiTable.Type
		list      *eventList
		ok        bool
	)
	if list, ok = m.events[classType]; !ok {
		list = &eventList{ptrs: make(map[int]uintptr)}
		m.events[classType] = list
	}
	if oldEventPtr, ok := list.ptrs[apiIndex]; ok && oldEventPtr != 0 {
		api.RemoveEventElement(oldEventPtr)
		delete(list.ptrs, apiIndex)
	}
	if eventPtr != 0 {
		list.ptrs[apiIndex] = eventPtr
	}
	apiTable.SysCallN(apiIndex, m.Instance(), eventPtr)
}

func (m *TBase) Instance() uintptr {
	if m == nil || m.instance == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TBase) UnsafeAddr() UnsafePointer {
	return m.instance
}

func (m *TBase) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *TBase) SetInstance(instance UnsafePointer) {
	if uintptr(m.instance) != uintptr(instance) {
		m.instance = instance
	}
}

// FreeAndNil Free and set nil, auto release memory and set pointer to nil
//func (m *TBase) FreeAndNil() {
//	if m.instance != nil {
//		api.FreeAndNil(m.Instance())
//		m.instance = nil
//	}
//}

// Nil Set the current object instance to nil, similar to obj = nil
//func (m *TBase) Nil() {
//	if m.instance != nil {
//		api.SetNil(m.Instance())
//		m.instance = nil
//	}
//}

// Free empty
func (m *TBase) Free() {
	if m.instance != nil {
		m.ClearEvents()
		m.intfs = nil
		m.instance = nil
	}
}

// TODO 直接添加函数名实现 "ClassName", "InstanceSize", "InheritsFrom", "ClassType", "ClassParent"
