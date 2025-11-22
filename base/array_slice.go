//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package base

import (
	"unsafe"

	"github.com/energye/lcl/tool/ptr"
)

type IInterface interface {
	IBase
	Release()
}

// IArraySlice 内部和外部数组维护
type IArraySlice[T any] interface {
	Instance() uintptr // 返回数组指针
	Count() int        // 返回数组个数
	Get(index int) T   // 获取指定元素
	Add(value T)       // 内部数组添加元素
	Free()             // 销毁
}

type TStructArraySlice[T any] struct {
	instance UnsafePointer // 返回外部数组地址
	sizeOf   uintptr       // 数组元素大小
	count    int           // 外部数组数量
	elements []T           // 内部数组集合
}

// NewStructArraySlice 初始化数组结构, 结构类型
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewStructArraySlice[T any](count int, instance uintptr) IArraySlice[T] {
	if instance == 0 {
		count = 0
	}
	var t T
	return &TStructArraySlice[T]{
		instance: UnsafePointer(instance),
		sizeOf:   unsafe.Sizeof(t),
		count:    count,
		elements: make([]T, count),
	}
}

func (m *TStructArraySlice[T]) Instance() uintptr {
	if m.instance != nil {
		// 返回外部数组地址
		return uintptr(m.instance)
	} else if m.instance == nil && len(m.elements) > 0 {
		// 返回内部数组地址
		return uintptr(UnsafePointer(&m.elements[0]))
	}
	return 0
}

func (m *TStructArraySlice[T]) Count() int {
	if m.instance != nil {
		// 返回外部数组个数
		return m.count
	} else if m.instance == nil {
		// 返回内部数组个数
		return len(m.elements)
	}
	return 0
}

func (m *TStructArraySlice[T]) Get(index int) T {
	if m.instance == nil && index >= 0 && index < len(m.elements) {
		// 内部数组
		return m.elements[index]
	} else if m.instance != nil && index >= 0 && index < m.count {
		// 外部数组
		return *(*T)(ptr.GetPointerOffset(uintptr(m.instance), uintptr(index)*m.sizeOf))
	}
	var t T
	return t
}

func (m *TStructArraySlice[T]) Add(value T) {
	if m.instance == nil {
		// 内部数组
		m.elements = append(m.elements, value)
	}
}

func (m *TStructArraySlice[T]) Free() {
	m.elements = nil
	m.instance = nil
	m.count = 0
}

type TClassArraySlice[T IBase] struct {
	instance UnsafePointer // 返回外部数组地址
	count    int           // 外部数组数量
	elements []T           // 内部数组集合
	address  []uintptr     // 内部数组地址集合
	asObject func(any) T
}

type TInterfaceArraySlice[T IInterface] struct {
	instance UnsafePointer // 返回外部数组地址
	count    int           // 外部数组数量
	elements []T           // 内部数组集合
	address  []uintptr     // 内部数组地址集合
	asObject func(any) T
}

// NewClassArraySlice 初始化数组结构, 实例对象类型
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewClassArraySlice[T IBase](count int, instance uintptr, asObject func(any) T) IArraySlice[T] {
	if instance == 0 {
		count = 0
	}
	return &TClassArraySlice[T]{
		instance: UnsafePointer(instance),
		count:    count,
		elements: make([]T, count),
		asObject: asObject,
	}
}

func (m *TClassArraySlice[T]) Instance() uintptr {
	if m.instance != nil {
		// 返回外部数组地址
		return uintptr(m.instance)
	}
	if len(m.elements) > 0 {
		// 返回内部数组地址
		return uintptr(UnsafePointer(&m.address[0]))
	}
	return 0
}

func (m *TClassArraySlice[T]) Count() int {
	if m.instance != nil {
		// 返回外部数组个数
		return m.count
	}
	// 返回内部数组个数
	return len(m.elements)
}

func (m *TClassArraySlice[T]) Get(index int) T {
	var r T
	if m.instance == nil && index >= 0 && index < len(m.elements) {
		// 内部数组
		r = m.elements[index]
	} else if m.instance != nil && index >= 0 && index < m.count {
		// 外部数组
		if r = m.elements[index]; any(r) != nil {
			return r
		}
		r = m.asObject(ptr.GetParamOf(index, m.Instance()))
		m.elements[index] = r
	}
	return r
}

func (m *TClassArraySlice[T]) Add(value T) {
	if m.instance == nil {
		// 内部数组
		m.elements = append(m.elements, value)
		m.address = append(m.address, value.Instance())
	}
}

func (m *TClassArraySlice[T]) Free() {
	if m.elements != nil {
		for _, v := range m.elements {
			if any(v) != nil && v.Instance() != 0 {
				v.Free()
			}
		}
		m.elements = nil
	}
	m.count = 0
	m.instance = nil
	m.address = nil
	m.count = 0
}

// NewInterfaceArraySlice 初始化数组结构, 实例对象类型
//
//	count: 外部数组元素个数
//	instance: 数组首地址, 值0(nil)表示内部维护数组, 否则为外部数组首地址
func NewInterfaceArraySlice[T IInterface](count int, instance uintptr, asObject func(any) T) IArraySlice[T] {
	if instance == 0 {
		count = 0
	}
	return &TInterfaceArraySlice[T]{
		instance: UnsafePointer(instance),
		count:    count,
		elements: make([]T, count),
		asObject: asObject,
	}
}

func (m *TInterfaceArraySlice[T]) Instance() uintptr {
	if m.instance != nil {
		// 返回外部数组地址
		return uintptr(m.instance)
	}
	if len(m.elements) > 0 {
		// 返回内部数组地址
		return uintptr(UnsafePointer(&m.address[0]))
	}
	return 0
}

func (m *TInterfaceArraySlice[T]) Count() int {
	if m.instance != nil {
		// 返回外部数组个数
		return m.count
	}
	// 返回内部数组个数
	return len(m.elements)
}

func (m *TInterfaceArraySlice[T]) Get(index int) T {
	var r T
	if m.instance == nil && index >= 0 && index < len(m.elements) {
		// 内部数组
		r = m.elements[index]
	} else if m.instance != nil && index >= 0 && index < m.count {
		// 外部数组
		if r = m.elements[index]; any(r) != nil {
			return r
		}
		r = m.asObject(ptr.GetParamOf(index, m.Instance()))
		m.elements[index] = r
	}
	return r
}

func (m *TInterfaceArraySlice[T]) Add(value T) {
	if m.instance == nil {
		// 内部数组
		m.elements = append(m.elements, value)
		m.address = append(m.address, value.Instance())
	}
}

func (m *TInterfaceArraySlice[T]) Free() {
	if m.elements != nil {
		for _, v := range m.elements {
			if any(v) != nil && v.Instance() != 0 {
				v.Release()
			}
		}
		m.elements = nil
	}
	m.count = 0
	m.instance = nil
	m.address = nil
	m.count = 0
}
