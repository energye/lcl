//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IDockZone Parent: IObject
type IDockZone interface {
	IObject
	ChildControl() IControl                        // property
	ChildCount() int32                             // property
	FirstChild() IDockZone                         // property
	Height() int32                                 // property
	SetHeight(AValue int32)                        // property
	Left() int32                                   // property
	SetLeft(AValue int32)                          // property
	LimitBegin() int32                             // property
	SetLimitBegin(AValue int32)                    // property
	LimitSize() int32                              // property
	SetLimitSize(AValue int32)                     // property
	Orientation() TDockOrientation                 // property
	SetOrientation(AValue TDockOrientation)        // property
	Parent() IDockZone                             // property
	Top() int32                                    // property
	SetTop(AValue int32)                           // property
	Tree() IDockTree                               // property
	Visible() bool                                 // property
	VisibleChildCount() int32                      // property
	Width() int32                                  // property
	SetWidth(AValue int32)                         // property
	NextSibling() IDockZone                        // property
	PrevSibling() IDockZone                        // property
	FindZone(AControl IControl) IDockZone          // function
	FirstVisibleChild() IDockZone                  // function
	GetNextVisibleZone() IDockZone                 // function
	NextVisible() IDockZone                        // function
	PrevVisible() IDockZone                        // function
	GetLastChild() IDockZone                       // function
	GetIndex() int32                               // function
	AddSibling(NewZone IDockZone, InsertAt TAlign) // procedure
	AddAsFirstChild(NewChildZone IDockZone)        // procedure
	AddAsLastChild(NewChildZone IDockZone)         // procedure
	ReplaceChild(OldChild, NewChild IDockZone)     // procedure
	Remove(ChildZone IDockZone)                    // procedure
}

// TDockZone Parent: TObject
type TDockZone struct {
	TObject
}

func NewDockZone(TheTree IDockTree, TheChildControl IControl) IDockZone {
	r1 := dockZoneImportAPI().SysCallN(6, GetObjectUintptr(TheTree), GetObjectUintptr(TheChildControl))
	return AsDockZone(r1)
}

func (m *TDockZone) ChildControl() IControl {
	r1 := dockZoneImportAPI().SysCallN(3, m.Instance())
	return AsControl(r1)
}

func (m *TDockZone) ChildCount() int32 {
	r1 := dockZoneImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TDockZone) FirstChild() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(8, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) Height() int32 {
	r1 := dockZoneImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetHeight(AValue int32) {
	dockZoneImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Left() int32 {
	r1 := dockZoneImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLeft(AValue int32) {
	dockZoneImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) LimitBegin() int32 {
	r1 := dockZoneImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLimitBegin(AValue int32) {
	dockZoneImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) LimitSize() int32 {
	r1 := dockZoneImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetLimitSize(AValue int32) {
	dockZoneImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Orientation() TDockOrientation {
	r1 := dockZoneImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TDockOrientation(r1)
}

func (m *TDockZone) SetOrientation(AValue TDockOrientation) {
	dockZoneImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Parent() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(20, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) Top() int32 {
	r1 := dockZoneImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetTop(AValue int32) {
	dockZoneImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) Tree() IDockTree {
	r1 := dockZoneImportAPI().SysCallN(26, m.Instance())
	return AsDockTree(r1)
}

func (m *TDockZone) Visible() bool {
	r1 := dockZoneImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TDockZone) VisibleChildCount() int32 {
	r1 := dockZoneImportAPI().SysCallN(28, m.Instance())
	return int32(r1)
}

func (m *TDockZone) Width() int32 {
	r1 := dockZoneImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDockZone) SetWidth(AValue int32) {
	dockZoneImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TDockZone) NextSibling() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(17, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) PrevSibling() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(21, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) FindZone(AControl IControl) IDockZone {
	r1 := dockZoneImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(AControl))
	return AsDockZone(r1)
}

func (m *TDockZone) FirstVisibleChild() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(9, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetNextVisibleZone() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(12, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) NextVisible() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(18, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) PrevVisible() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(22, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetLastChild() IDockZone {
	r1 := dockZoneImportAPI().SysCallN(11, m.Instance())
	return AsDockZone(r1)
}

func (m *TDockZone) GetIndex() int32 {
	r1 := dockZoneImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func DockZoneClass() TClass {
	ret := dockZoneImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TDockZone) AddSibling(NewZone IDockZone, InsertAt TAlign) {
	dockZoneImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(NewZone), uintptr(InsertAt))
}

func (m *TDockZone) AddAsFirstChild(NewChildZone IDockZone) {
	dockZoneImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(NewChildZone))
}

func (m *TDockZone) AddAsLastChild(NewChildZone IDockZone) {
	dockZoneImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(NewChildZone))
}

func (m *TDockZone) ReplaceChild(OldChild, NewChild IDockZone) {
	dockZoneImportAPI().SysCallN(24, m.Instance(), GetObjectUintptr(OldChild), GetObjectUintptr(NewChild))
}

func (m *TDockZone) Remove(ChildZone IDockZone) {
	dockZoneImportAPI().SysCallN(23, m.Instance(), GetObjectUintptr(ChildZone))
}

var (
	dockZoneImport       *imports.Imports = nil
	dockZoneImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DockZone_AddAsFirstChild", 0),
		/*1*/ imports.NewTable("DockZone_AddAsLastChild", 0),
		/*2*/ imports.NewTable("DockZone_AddSibling", 0),
		/*3*/ imports.NewTable("DockZone_ChildControl", 0),
		/*4*/ imports.NewTable("DockZone_ChildCount", 0),
		/*5*/ imports.NewTable("DockZone_Class", 0),
		/*6*/ imports.NewTable("DockZone_Create", 0),
		/*7*/ imports.NewTable("DockZone_FindZone", 0),
		/*8*/ imports.NewTable("DockZone_FirstChild", 0),
		/*9*/ imports.NewTable("DockZone_FirstVisibleChild", 0),
		/*10*/ imports.NewTable("DockZone_GetIndex", 0),
		/*11*/ imports.NewTable("DockZone_GetLastChild", 0),
		/*12*/ imports.NewTable("DockZone_GetNextVisibleZone", 0),
		/*13*/ imports.NewTable("DockZone_Height", 0),
		/*14*/ imports.NewTable("DockZone_Left", 0),
		/*15*/ imports.NewTable("DockZone_LimitBegin", 0),
		/*16*/ imports.NewTable("DockZone_LimitSize", 0),
		/*17*/ imports.NewTable("DockZone_NextSibling", 0),
		/*18*/ imports.NewTable("DockZone_NextVisible", 0),
		/*19*/ imports.NewTable("DockZone_Orientation", 0),
		/*20*/ imports.NewTable("DockZone_Parent", 0),
		/*21*/ imports.NewTable("DockZone_PrevSibling", 0),
		/*22*/ imports.NewTable("DockZone_PrevVisible", 0),
		/*23*/ imports.NewTable("DockZone_Remove", 0),
		/*24*/ imports.NewTable("DockZone_ReplaceChild", 0),
		/*25*/ imports.NewTable("DockZone_Top", 0),
		/*26*/ imports.NewTable("DockZone_Tree", 0),
		/*27*/ imports.NewTable("DockZone_Visible", 0),
		/*28*/ imports.NewTable("DockZone_VisibleChildCount", 0),
		/*29*/ imports.NewTable("DockZone_Width", 0),
	}
)

func dockZoneImportAPI() *imports.Imports {
	if dockZoneImport == nil {
		dockZoneImport = NewDefaultImports()
		dockZoneImport.SetImportTable(dockZoneImportTables)
		dockZoneImportTables = nil
	}
	return dockZoneImport
}
