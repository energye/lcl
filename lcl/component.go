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

// IComponent Parent: IPersistent
type IComponent interface {
	IPersistent
	Components(Index int32) IComponent            // property
	ComponentCount() int32                        // property
	ComponentIndex() int32                        // property
	SetComponentIndex(AValue int32)               // property
	ComponentState() TComponentStates             // property
	ComponentStyle() TComponentStyles             // property
	DesignInfo() int32                            // property
	SetDesignInfo(AValue int32)                   // property
	Owner() IComponent                            // property
	VCLComObject() uintptr                        // property
	SetVCLComObject(AValue uintptr)               // property
	Name() string                                 // property
	SetName(AValue string)                        // property
	Tag() uint32                                  // property
	SetTag(AValue uint32)                         // property
	ExecuteAction(Action IBasicAction) bool       // function
	FindComponent(AName string) IComponent        // function
	GetEnumerator() IComponentEnumerator          // function
	GetParentComponent() IComponent               // function
	HasParent() bool                              // function
	UpdateAction(Action IBasicAction) bool        // function
	DestroyComponents()                           // procedure
	Destroying()                                  // procedure
	FreeNotification(AComponent IComponent)       // procedure
	RemoveFreeNotification(AComponent IComponent) // procedure
	FreeOnRelease()                               // procedure
	InsertComponent(AComponent IComponent)        // procedure
	RemoveComponent(AComponent IComponent)        // procedure
	SetSubComponent(ASubComponent bool)           // procedure
}

// TComponent Parent: TPersistent
type TComponent struct {
	TPersistent
}

func NewComponent(AOwner IComponent) IComponent {
	r1 := componentImportAPI().SysCallN(6, GetObjectUintptr(AOwner))
	return AsComponent(r1)
}

func (m *TComponent) Components(Index int32) IComponent {
	r1 := componentImportAPI().SysCallN(5, m.Instance(), uintptr(Index))
	return AsComponent(r1)
}

func (m *TComponent) ComponentCount() int32 {
	r1 := componentImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TComponent) ComponentIndex() int32 {
	r1 := componentImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComponent) SetComponentIndex(AValue int32) {
	componentImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) ComponentState() TComponentStates {
	r1 := componentImportAPI().SysCallN(3, m.Instance())
	return TComponentStates(r1)
}

func (m *TComponent) ComponentStyle() TComponentStyles {
	r1 := componentImportAPI().SysCallN(4, m.Instance())
	return TComponentStyles(r1)
}

func (m *TComponent) DesignInfo() int32 {
	r1 := componentImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComponent) SetDesignInfo(AValue int32) {
	componentImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) Owner() IComponent {
	r1 := componentImportAPI().SysCallN(19, m.Instance())
	return AsComponent(r1)
}

func (m *TComponent) VCLComObject() uintptr {
	r1 := componentImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TComponent) SetVCLComObject(AValue uintptr) {
	componentImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) Name() string {
	r1 := componentImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TComponent) SetName(AValue string) {
	componentImportAPI().SysCallN(18, 1, m.Instance(), PascalStr(AValue))
}

func (m *TComponent) Tag() uint32 {
	r1 := componentImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TComponent) SetTag(AValue uint32) {
	componentImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) ExecuteAction(Action IBasicAction) bool {
	r1 := componentImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(Action))
	return GoBool(r1)
}

func (m *TComponent) FindComponent(AName string) IComponent {
	r1 := componentImportAPI().SysCallN(11, m.Instance(), PascalStr(AName))
	return AsComponent(r1)
}

func (m *TComponent) GetEnumerator() IComponentEnumerator {
	r1 := componentImportAPI().SysCallN(14, m.Instance())
	return AsComponentEnumerator(r1)
}

func (m *TComponent) GetParentComponent() IComponent {
	r1 := componentImportAPI().SysCallN(15, m.Instance())
	return AsComponent(r1)
}

func (m *TComponent) HasParent() bool {
	r1 := componentImportAPI().SysCallN(16, m.Instance())
	return GoBool(r1)
}

func (m *TComponent) UpdateAction(Action IBasicAction) bool {
	r1 := componentImportAPI().SysCallN(24, m.Instance(), GetObjectUintptr(Action))
	return GoBool(r1)
}

func ComponentClass() TClass {
	ret := componentImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TComponent) DestroyComponents() {
	componentImportAPI().SysCallN(8, m.Instance())
}

func (m *TComponent) Destroying() {
	componentImportAPI().SysCallN(9, m.Instance())
}

func (m *TComponent) FreeNotification(AComponent IComponent) {
	componentImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) RemoveFreeNotification(AComponent IComponent) {
	componentImportAPI().SysCallN(21, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) FreeOnRelease() {
	componentImportAPI().SysCallN(13, m.Instance())
}

func (m *TComponent) InsertComponent(AComponent IComponent) {
	componentImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) RemoveComponent(AComponent IComponent) {
	componentImportAPI().SysCallN(20, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) SetSubComponent(ASubComponent bool) {
	componentImportAPI().SysCallN(22, m.Instance(), PascalBool(ASubComponent))
}

var (
	componentImport       *imports.Imports = nil
	componentImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Component_Class", 0),
		/*1*/ imports.NewTable("Component_ComponentCount", 0),
		/*2*/ imports.NewTable("Component_ComponentIndex", 0),
		/*3*/ imports.NewTable("Component_ComponentState", 0),
		/*4*/ imports.NewTable("Component_ComponentStyle", 0),
		/*5*/ imports.NewTable("Component_Components", 0),
		/*6*/ imports.NewTable("Component_Create", 0),
		/*7*/ imports.NewTable("Component_DesignInfo", 0),
		/*8*/ imports.NewTable("Component_DestroyComponents", 0),
		/*9*/ imports.NewTable("Component_Destroying", 0),
		/*10*/ imports.NewTable("Component_ExecuteAction", 0),
		/*11*/ imports.NewTable("Component_FindComponent", 0),
		/*12*/ imports.NewTable("Component_FreeNotification", 0),
		/*13*/ imports.NewTable("Component_FreeOnRelease", 0),
		/*14*/ imports.NewTable("Component_GetEnumerator", 0),
		/*15*/ imports.NewTable("Component_GetParentComponent", 0),
		/*16*/ imports.NewTable("Component_HasParent", 0),
		/*17*/ imports.NewTable("Component_InsertComponent", 0),
		/*18*/ imports.NewTable("Component_Name", 0),
		/*19*/ imports.NewTable("Component_Owner", 0),
		/*20*/ imports.NewTable("Component_RemoveComponent", 0),
		/*21*/ imports.NewTable("Component_RemoveFreeNotification", 0),
		/*22*/ imports.NewTable("Component_SetSubComponent", 0),
		/*23*/ imports.NewTable("Component_Tag", 0),
		/*24*/ imports.NewTable("Component_UpdateAction", 0),
		/*25*/ imports.NewTable("Component_VCLComObject", 0),
	}
)

func componentImportAPI() *imports.Imports {
	if componentImport == nil {
		componentImport = NewDefaultImports()
		componentImport.SetImportTable(componentImportTables)
		componentImportTables = nil
	}
	return componentImport
}
