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

// IDockManager Is Abstract Class Parent: IPersistent
type IDockManager interface {
	IPersistent
	GetDockEdge(ADockObject IDragDockObject) bool                                  // function
	AutoFreeByControl() bool                                                       // function
	IsEnabledControl(Control IControl) bool                                        // function
	CanBeDoubleDocked() bool                                                       // function
	BeginUpdate()                                                                  // procedure
	EndUpdate()                                                                    // procedure
	GetControlBounds(Control IControl, OutControlBounds *TRect)                    // procedure Is Abstract
	InsertControl(ADockObject IDragDockObject)                                     // procedure
	InsertControl1(Control IControl, InsertAt TAlign, DropCtl IControl)            // procedure Is Abstract
	LoadFromStream(Stream IStream)                                                 // procedure Is Abstract
	PaintSite(DC HDC)                                                              // procedure
	MessageHandler(Sender IControl, Message *TLMessage)                            // procedure
	PositionDockRect(ADockObject IDragDockObject)                                  // procedure
	PositionDockRect1(Client, DropCtl IControl, DropAlign TAlign, DockRect *TRect) // procedure Is Abstract
	RemoveControl(Control IControl)                                                // procedure Is Abstract
	ResetBounds(Force bool)                                                        // procedure Is Abstract
	SaveToStream(Stream IStream)                                                   // procedure Is Abstract
	SetReplacingControl(Control IControl)                                          // procedure
}

// TDockManager Is Abstract Class Parent: TPersistent
type TDockManager struct {
	TPersistent
}

func (m *TDockManager) GetDockEdge(ADockObject IDragDockObject) bool {
	r1 := dockManagerImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(ADockObject))
	return GoBool(r1)
}

func (m *TDockManager) AutoFreeByControl() bool {
	r1 := dockManagerImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TDockManager) IsEnabledControl(Control IControl) bool {
	r1 := dockManagerImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TDockManager) CanBeDoubleDocked() bool {
	r1 := dockManagerImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func DockManagerClass() TClass {
	ret := dockManagerImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TDockManager) BeginUpdate() {
	dockManagerImportAPI().SysCallN(1, m.Instance())
}

func (m *TDockManager) EndUpdate() {
	dockManagerImportAPI().SysCallN(4, m.Instance())
}

func (m *TDockManager) GetControlBounds(Control IControl, OutControlBounds *TRect) {
	var result1 uintptr
	dockManagerImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Control), uintptr(unsafePointer(&result1)))
	*OutControlBounds = *(*TRect)(getPointer(result1))
}

func (m *TDockManager) InsertControl(ADockObject IDragDockObject) {
	dockManagerImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(ADockObject))
}

func (m *TDockManager) InsertControl1(Control IControl, InsertAt TAlign, DropCtl IControl) {
	dockManagerImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(Control), uintptr(InsertAt), GetObjectUintptr(DropCtl))
}

func (m *TDockManager) LoadFromStream(Stream IStream) {
	dockManagerImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TDockManager) PaintSite(DC HDC) {
	dockManagerImportAPI().SysCallN(12, m.Instance(), uintptr(DC))
}

func (m *TDockManager) MessageHandler(Sender IControl, Message *TLMessage) {
	var result1 uintptr
	dockManagerImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*Message = *(*TLMessage)(getPointer(result1))
}

func (m *TDockManager) PositionDockRect(ADockObject IDragDockObject) {
	dockManagerImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(ADockObject))
}

func (m *TDockManager) PositionDockRect1(Client, DropCtl IControl, DropAlign TAlign, DockRect *TRect) {
	var result2 uintptr
	dockManagerImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(Client), GetObjectUintptr(DropCtl), uintptr(DropAlign), uintptr(unsafePointer(&result2)))
	*DockRect = *(*TRect)(getPointer(result2))
}

func (m *TDockManager) RemoveControl(Control IControl) {
	dockManagerImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(Control))
}

func (m *TDockManager) ResetBounds(Force bool) {
	dockManagerImportAPI().SysCallN(16, m.Instance(), PascalBool(Force))
}

func (m *TDockManager) SaveToStream(Stream IStream) {
	dockManagerImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TDockManager) SetReplacingControl(Control IControl) {
	dockManagerImportAPI().SysCallN(18, m.Instance(), GetObjectUintptr(Control))
}

var (
	dockManagerImport       *imports.Imports = nil
	dockManagerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DockManager_AutoFreeByControl", 0),
		/*1*/ imports.NewTable("DockManager_BeginUpdate", 0),
		/*2*/ imports.NewTable("DockManager_CanBeDoubleDocked", 0),
		/*3*/ imports.NewTable("DockManager_Class", 0),
		/*4*/ imports.NewTable("DockManager_EndUpdate", 0),
		/*5*/ imports.NewTable("DockManager_GetControlBounds", 0),
		/*6*/ imports.NewTable("DockManager_GetDockEdge", 0),
		/*7*/ imports.NewTable("DockManager_InsertControl", 0),
		/*8*/ imports.NewTable("DockManager_InsertControl1", 0),
		/*9*/ imports.NewTable("DockManager_IsEnabledControl", 0),
		/*10*/ imports.NewTable("DockManager_LoadFromStream", 0),
		/*11*/ imports.NewTable("DockManager_MessageHandler", 0),
		/*12*/ imports.NewTable("DockManager_PaintSite", 0),
		/*13*/ imports.NewTable("DockManager_PositionDockRect", 0),
		/*14*/ imports.NewTable("DockManager_PositionDockRect1", 0),
		/*15*/ imports.NewTable("DockManager_RemoveControl", 0),
		/*16*/ imports.NewTable("DockManager_ResetBounds", 0),
		/*17*/ imports.NewTable("DockManager_SaveToStream", 0),
		/*18*/ imports.NewTable("DockManager_SetReplacingControl", 0),
	}
)

func dockManagerImportAPI() *imports.Imports {
	if dockManagerImport == nil {
		dockManagerImport = NewDefaultImports()
		dockManagerImport.SetImportTable(dockManagerImportTables)
		dockManagerImportTables = nil
	}
	return dockManagerImport
}
