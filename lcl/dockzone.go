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

// IDockZone Parent: IObject
type IDockZone interface {
	IObject
	FindZone(control IControl) IDockZone                 // function
	FirstVisibleChild() IDockZone                        // function
	GetNextVisibleZone() IDockZone                       // function
	NextVisible() IDockZone                              // function
	PrevVisible() IDockZone                              // function
	GetLastChild() IDockZone                             // function
	GetIndex() int32                                     // function
	AddSibling(newZone IDockZone, insertAt types.TAlign) // procedure
	AddAsFirstChild(newChildZone IDockZone)              // procedure
	AddAsLastChild(newChildZone IDockZone)               // procedure
	ReplaceChild(oldChild IDockZone, newChild IDockZone) // procedure
	Remove(childZone IDockZone)                          // procedure
	ChildControl() IControl                              // property ChildControl Getter
	ChildCount() int32                                   // property ChildCount Getter
	FirstChild() IDockZone                               // property FirstChild Getter
	Height() int32                                       // property Height Getter
	SetHeight(value int32)                               // property Height Setter
	Left() int32                                         // property Left Getter
	SetLeft(value int32)                                 // property Left Setter
	LimitBegin() int32                                   // property LimitBegin Getter
	SetLimitBegin(value int32)                           // property LimitBegin Setter
	LimitSize() int32                                    // property LimitSize Getter
	SetLimitSize(value int32)                            // property LimitSize Setter
	Orientation() types.TDockOrientation                 // property Orientation Getter
	SetOrientation(value types.TDockOrientation)         // property Orientation Setter
	Parent() IDockZone                                   // property Parent Getter
	Top() int32                                          // property Top Getter
	SetTop(value int32)                                  // property Top Setter
	Tree() IDockTree                                     // property Tree Getter
	Visible() bool                                       // property Visible Getter
	VisibleChildCount() int32                            // property VisibleChildCount Getter
	Width() int32                                        // property Width Getter
	SetWidth(value int32)                                // property Width Setter
	NextSibling() IDockZone                              // property NextSibling Getter
	PrevSibling() IDockZone                              // property PrevSibling Getter
}

type TDockZone struct {
	TObject
}

func (m *TDockZone) FindZone(control IControl) IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(control))
	return AsDockZone(r)
}

func (m *TDockZone) FirstVisibleChild() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(2, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) GetNextVisibleZone() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(3, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) NextVisible() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(4, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) PrevVisible() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(5, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) GetLastChild() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(6, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) GetIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TDockZone) AddSibling(newZone IDockZone, insertAt types.TAlign) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(newZone), uintptr(insertAt))
}

func (m *TDockZone) AddAsFirstChild(newChildZone IDockZone) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(newChildZone))
}

func (m *TDockZone) AddAsLastChild(newChildZone IDockZone) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(newChildZone))
}

func (m *TDockZone) ReplaceChild(oldChild IDockZone, newChild IDockZone) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(oldChild), base.GetObjectUintptr(newChild))
}

func (m *TDockZone) Remove(childZone IDockZone) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(childZone))
}

func (m *TDockZone) ChildControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(13, m.Instance())
	return AsControl(r)
}

func (m *TDockZone) ChildCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(14, m.Instance())
	return int32(r)
}

func (m *TDockZone) FirstChild() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(15, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TDockZone) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TDockZone) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TDockZone) SetLeft(value int32) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TDockZone) LimitBegin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TDockZone) SetLimitBegin(value int32) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TDockZone) LimitSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TDockZone) SetLimitSize(value int32) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TDockZone) Orientation() types.TDockOrientation {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(20, 0, m.Instance())
	return types.TDockOrientation(r)
}

func (m *TDockZone) SetOrientation(value types.TDockOrientation) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TDockZone) Parent() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(21, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TDockZone) SetTop(value int32) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TDockZone) Tree() IDockTree {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(23, m.Instance())
	return AsDockTree(r)
}

func (m *TDockZone) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := dockZoneAPI().SysCallN(24, m.Instance())
	return api.GoBool(r)
}

func (m *TDockZone) VisibleChildCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(25, m.Instance())
	return int32(r)
}

func (m *TDockZone) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dockZoneAPI().SysCallN(26, 0, m.Instance())
	return int32(r)
}

func (m *TDockZone) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	dockZoneAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TDockZone) NextSibling() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(27, m.Instance())
	return AsDockZone(r)
}

func (m *TDockZone) PrevSibling() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockZoneAPI().SysCallN(28, m.Instance())
	return AsDockZone(r)
}

// NewDockZone class constructor
func NewDockZone(theTree IDockTree, theChildControl IControl) IDockZone {
	r := dockZoneAPI().SysCallN(0, base.GetObjectUintptr(theTree), base.GetObjectUintptr(theChildControl))
	return AsDockZone(r)
}

var (
	dockZoneOnce   base.Once
	dockZoneImport *imports.Imports = nil
)

func dockZoneAPI() *imports.Imports {
	dockZoneOnce.Do(func() {
		dockZoneImport = api.NewDefaultImports()
		dockZoneImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDockZone_Create", 0), // constructor NewDockZone
			/* 1 */ imports.NewTable("TDockZone_FindZone", 0), // function FindZone
			/* 2 */ imports.NewTable("TDockZone_FirstVisibleChild", 0), // function FirstVisibleChild
			/* 3 */ imports.NewTable("TDockZone_GetNextVisibleZone", 0), // function GetNextVisibleZone
			/* 4 */ imports.NewTable("TDockZone_NextVisible", 0), // function NextVisible
			/* 5 */ imports.NewTable("TDockZone_PrevVisible", 0), // function PrevVisible
			/* 6 */ imports.NewTable("TDockZone_GetLastChild", 0), // function GetLastChild
			/* 7 */ imports.NewTable("TDockZone_GetIndex", 0), // function GetIndex
			/* 8 */ imports.NewTable("TDockZone_AddSibling", 0), // procedure AddSibling
			/* 9 */ imports.NewTable("TDockZone_AddAsFirstChild", 0), // procedure AddAsFirstChild
			/* 10 */ imports.NewTable("TDockZone_AddAsLastChild", 0), // procedure AddAsLastChild
			/* 11 */ imports.NewTable("TDockZone_ReplaceChild", 0), // procedure ReplaceChild
			/* 12 */ imports.NewTable("TDockZone_Remove", 0), // procedure Remove
			/* 13 */ imports.NewTable("TDockZone_ChildControl", 0), // property ChildControl
			/* 14 */ imports.NewTable("TDockZone_ChildCount", 0), // property ChildCount
			/* 15 */ imports.NewTable("TDockZone_FirstChild", 0), // property FirstChild
			/* 16 */ imports.NewTable("TDockZone_Height", 0), // property Height
			/* 17 */ imports.NewTable("TDockZone_Left", 0), // property Left
			/* 18 */ imports.NewTable("TDockZone_LimitBegin", 0), // property LimitBegin
			/* 19 */ imports.NewTable("TDockZone_LimitSize", 0), // property LimitSize
			/* 20 */ imports.NewTable("TDockZone_Orientation", 0), // property Orientation
			/* 21 */ imports.NewTable("TDockZone_Parent", 0), // property Parent
			/* 22 */ imports.NewTable("TDockZone_Top", 0), // property Top
			/* 23 */ imports.NewTable("TDockZone_Tree", 0), // property Tree
			/* 24 */ imports.NewTable("TDockZone_Visible", 0), // property Visible
			/* 25 */ imports.NewTable("TDockZone_VisibleChildCount", 0), // property VisibleChildCount
			/* 26 */ imports.NewTable("TDockZone_Width", 0), // property Width
			/* 27 */ imports.NewTable("TDockZone_NextSibling", 0), // property NextSibling
			/* 28 */ imports.NewTable("TDockZone_PrevSibling", 0), // property PrevSibling
		}
	})
	return dockZoneImport
}
