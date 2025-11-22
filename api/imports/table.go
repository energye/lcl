//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// 导入动态链接库函数表

package imports

// Table lib API 表
type Table struct {
	name string   // API 名称
	addr ProcAddr // API 函数地址
}

// NewTable 创建一个新的 API  table
func NewTable(name string, addr ProcAddr) *Table {
	r := &Table{}
	r.name = name
	r.addr = addr
	return r
}

// Name 返回这个 API 名称
func (m *Table) Name() string {
	return m.name
}

// Addr 返回这个 API 函数地址
func (m *Table) Addr() ProcAddr {
	return m.addr
}
