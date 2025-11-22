//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package imports

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

// CallImport 调用导入表接口
type CallImport interface {
	Proc(index int) ProcAddr
	SysCallN(index int, args ...uintptr) (result uintptr)
}

// 全局 API 表所属类对象类型
var classType int32

// Imports 导入表实例
type Imports struct {
	Dll   DLL
	Table []*Table
	Type  int32
}

// LoadLib 加载动态库，适用自定义导入动态链接库，支持Windows, MacOS, Linux
func LoadLib(libName string) (*Imports, error) {
	if libName == "" {
		return &Imports{}, errors.New("the dynamic link library is empty")
	}
	dll, err := NewDLL(libName)
	if dll == 0 && err != nil {
		return nil, err
	}
	return &Imports{Dll: dll}, nil
}

// NextType 设置API表类型, 用于区分不同所属类对象
func (m *Imports) NextType() {
	// m.Type >= math.AddInt32 ???
	m.Type = atomic.AddInt32(&classType, 1)
}

// Proc 返回api实例
//
//	index: 导入表索引
func (m *Imports) Proc(index int) ProcAddr {
	return m.internalGetImportFunc(index)
}

// SysCallN 调用api
//
//	 参数
//		index: 导入表索引
//		args: 调用api参数, 指针数组
func (m *Imports) SysCallN(index int, args ...uintptr) uintptr {
	proc := m.internalGetImportFunc(index)
	if proc > 0 {
		defer func() {
			//if err := recover(); err != nil {
			//	println("[ERROR] SysCallN Name:", m.Table[index].Name(), "Message:", err.(error).Error())
			//	// TODO 增加用户回调
			//	//exception.GlobalException(m.table[index].Name(), err.(error).Error())
			//}
		}()
		r1, _, _ := proc.Call(args...)
		//if err != 0 && err != 1400 {
		//exception.CallException(m.table[index].Name(), err.Error())
		//	}
		return r1
	}
	return 0
}

func (m *Imports) internalGetImportFunc(index int) ProcAddr {
	item := m.Table[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.addr))) == nil {
		var err error
		item.addr, err = m.Dll.GetProcAddr(item.name)
		if err != nil {
			println("[ERROR] GetImport Name:", item.name, "Message:", err.Error())
			return 0
		}
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.Table[index].addr)), unsafe.Pointer(item.addr))
	}
	return item.addr
}
