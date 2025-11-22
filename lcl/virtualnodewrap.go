//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// IVirtualNodeWrap Parent: IObject
type IVirtualNodeWrap interface {
	IObject
	Data() types.PVirtualNode // function
	// Index
	//  index of node with regard to its parent
	Index() uint32 // function
	// ChildCount
	//  number of child nodes
	ChildCount() uint32 // function
	// NodeHeight
	//  height in pixels
	NodeHeight() uint16 // function
	// States
	//  states describing various properties of the node (expanded, initialized etc.)
	States() types.TVirtualNodeStates // function
	// Align
	//  line/button alignment
	Align() byte // function
	// CheckState
	//  indicates the current check state (e.g. checked, pressed etc.)
	CheckState() types.TCheckState // function
	// CheckType
	//  indicates which check type shall be used for this node
	CheckType() types.TCheckType // function
	// Dummy
	//  dummy value to fill DWORD boundary
	Dummy() byte // function
	// TotalCount
	//  sum of this node, all of its child nodes and their child nodes etc.
	TotalCount() uint32 // function
	// TotalHeight
	//  height in pixels this node covers on screen including the height of all of its children
	//  Note: Some copy routines require that all pointers (as well as the data area) in a node are
	//  located at the end of the node! Hence if you want to add new member fields (except pointers to internal
	//  data) then put them before field Parent.
	TotalHeight() uint32 // function
	// Parent
	//  reference to the node's parent (for the root this contains the treeview)
	Parent() IVirtualNodeWrap // function
	// PrevSibling
	//  link to the node's previous sibling or nil if it is the first node
	PrevSibling() IVirtualNodeWrap // function
	// NextSibling
	//  link to the node's next sibling or nil if it is the last node
	NextSibling() IVirtualNodeWrap // function
	// FirstChild
	//  link to the node's first child...
	FirstChild() IVirtualNodeWrap // function
	// LastChild
	//  link to the node's last child...
	LastChild() IVirtualNodeWrap // function
}

type TVirtualNodeWrap struct {
	TObject
}

func (m *TVirtualNodeWrap) Data() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(1, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TVirtualNodeWrap) Index() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(2, m.Instance())
	return uint32(r)
}

func (m *TVirtualNodeWrap) ChildCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(3, m.Instance())
	return uint32(r)
}

func (m *TVirtualNodeWrap) NodeHeight() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(4, m.Instance())
	return uint16(r)
}

func (m *TVirtualNodeWrap) States() types.TVirtualNodeStates {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(5, m.Instance())
	return types.TVirtualNodeStates(r)
}

func (m *TVirtualNodeWrap) Align() byte {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(6, m.Instance())
	return byte(r)
}

func (m *TVirtualNodeWrap) CheckState() types.TCheckState {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(7, m.Instance())
	return types.TCheckState(r)
}

func (m *TVirtualNodeWrap) CheckType() types.TCheckType {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(8, m.Instance())
	return types.TCheckType(r)
}

func (m *TVirtualNodeWrap) Dummy() byte {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(9, m.Instance())
	return byte(r)
}

func (m *TVirtualNodeWrap) TotalCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(10, m.Instance())
	return uint32(r)
}

func (m *TVirtualNodeWrap) TotalHeight() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualNodeWrapAPI().SysCallN(11, m.Instance())
	return uint32(r)
}

func (m *TVirtualNodeWrap) Parent() IVirtualNodeWrap {
	if !m.IsValid() {
		return nil
	}
	r := virtualNodeWrapAPI().SysCallN(12, m.Instance())
	return AsVirtualNodeWrap(r)
}

func (m *TVirtualNodeWrap) PrevSibling() IVirtualNodeWrap {
	if !m.IsValid() {
		return nil
	}
	r := virtualNodeWrapAPI().SysCallN(13, m.Instance())
	return AsVirtualNodeWrap(r)
}

func (m *TVirtualNodeWrap) NextSibling() IVirtualNodeWrap {
	if !m.IsValid() {
		return nil
	}
	r := virtualNodeWrapAPI().SysCallN(14, m.Instance())
	return AsVirtualNodeWrap(r)
}

func (m *TVirtualNodeWrap) FirstChild() IVirtualNodeWrap {
	if !m.IsValid() {
		return nil
	}
	r := virtualNodeWrapAPI().SysCallN(15, m.Instance())
	return AsVirtualNodeWrap(r)
}

func (m *TVirtualNodeWrap) LastChild() IVirtualNodeWrap {
	if !m.IsValid() {
		return nil
	}
	r := virtualNodeWrapAPI().SysCallN(16, m.Instance())
	return AsVirtualNodeWrap(r)
}

// VirtualNodeWrap  is static instance
var VirtualNodeWrap _VirtualNodeWrapClass

// _VirtualNodeWrapClass is class type defined by TVirtualNodeWrap
type _VirtualNodeWrapClass uintptr

func (_VirtualNodeWrapClass) UnWrap(data types.PVirtualNode) IVirtualNodeWrap {
	r := virtualNodeWrapAPI().SysCallN(17, uintptr(data))
	return AsVirtualNodeWrap(r)
}

// NewVirtualNodeWrap class constructor
func NewVirtualNodeWrap() IVirtualNodeWrap {
	r := virtualNodeWrapAPI().SysCallN(0)
	return AsVirtualNodeWrap(r)
}

var (
	virtualNodeWrapOnce   base.Once
	virtualNodeWrapImport *imports.Imports = nil
)

func virtualNodeWrapAPI() *imports.Imports {
	virtualNodeWrapOnce.Do(func() {
		virtualNodeWrapImport = api.NewDefaultImports()
		virtualNodeWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVirtualNodeWrap_Create", 0), // constructor NewVirtualNodeWrap
			/* 1 */ imports.NewTable("TVirtualNodeWrap_Data", 0), // function Data
			/* 2 */ imports.NewTable("TVirtualNodeWrap_Index", 0), // function Index
			/* 3 */ imports.NewTable("TVirtualNodeWrap_ChildCount", 0), // function ChildCount
			/* 4 */ imports.NewTable("TVirtualNodeWrap_NodeHeight", 0), // function NodeHeight
			/* 5 */ imports.NewTable("TVirtualNodeWrap_States", 0), // function States
			/* 6 */ imports.NewTable("TVirtualNodeWrap_Align", 0), // function Align
			/* 7 */ imports.NewTable("TVirtualNodeWrap_CheckState", 0), // function CheckState
			/* 8 */ imports.NewTable("TVirtualNodeWrap_CheckType", 0), // function CheckType
			/* 9 */ imports.NewTable("TVirtualNodeWrap_Dummy", 0), // function Dummy
			/* 10 */ imports.NewTable("TVirtualNodeWrap_TotalCount", 0), // function TotalCount
			/* 11 */ imports.NewTable("TVirtualNodeWrap_TotalHeight", 0), // function TotalHeight
			/* 12 */ imports.NewTable("TVirtualNodeWrap_Parent", 0), // function Parent
			/* 13 */ imports.NewTable("TVirtualNodeWrap_PrevSibling", 0), // function PrevSibling
			/* 14 */ imports.NewTable("TVirtualNodeWrap_NextSibling", 0), // function NextSibling
			/* 15 */ imports.NewTable("TVirtualNodeWrap_FirstChild", 0), // function FirstChild
			/* 16 */ imports.NewTable("TVirtualNodeWrap_LastChild", 0), // function LastChild
			/* 17 */ imports.NewTable("TVirtualNodeWrap_UnWrap", 0), // static function UnWrap
		}
	})
	return virtualNodeWrapImport
}
