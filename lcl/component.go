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

// IComponent Parent: IPersistent
type IComponent interface {
	IPersistent
	ExecuteAction(action IBasicAction) bool      // function
	FindComponent(name string) IComponent        // function
	GetEnumerator() IComponentEnumerator         // function
	GetParentComponent() IComponent              // function
	HasParent() bool                             // function
	UpdateAction(action IBasicAction) bool       // function
	DestroyComponents()                          // procedure
	Destroying()                                 // procedure
	FreeNotification(component IComponent)       // procedure
	RemoveFreeNotification(component IComponent) // procedure
	FreeOnRelease()                              // procedure
	InsertComponent(component IComponent)        // procedure
	RemoveComponent(component IComponent)        // procedure
	SetSubComponent(subComponent bool)           // procedure
	Components(index int32) IComponent           // property Components Getter
	ComponentCount() int32                       // property ComponentCount Getter
	ComponentIndex() int32                       // property ComponentIndex Getter
	SetComponentIndex(value int32)               // property ComponentIndex Setter
	ComponentState() types.TComponentState       // property ComponentState Getter
	ComponentStyle() types.TComponentStyle       // property ComponentStyle Getter
	DesignInfo() int32                           // property DesignInfo Getter
	SetDesignInfo(value int32)                   // property DesignInfo Setter
	Owner() IComponent                           // property Owner Getter
	VCLComObject() uintptr                       // property VCLComObject Getter
	SetVCLComObject(value uintptr)               // property VCLComObject Setter
	Name() string                                // property Name Getter
	SetName(value string)                        // property Name Setter
	Tag() uintptr                                // property Tag Getter
	SetTag(value uintptr)                        // property Tag Setter
}

type TComponent struct {
	TPersistent
}

func (m *TComponent) ExecuteAction(action IBasicAction) bool {
	if !m.IsValid() {
		return false
	}
	r := componentAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(action))
	return api.GoBool(r)
}

func (m *TComponent) FindComponent(name string) IComponent {
	if !m.IsValid() {
		return nil
	}
	r := componentAPI().SysCallN(2, m.Instance(), api.PasStr(name))
	return AsComponent(r)
}

func (m *TComponent) GetEnumerator() IComponentEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := componentAPI().SysCallN(3, m.Instance())
	return AsComponentEnumerator(r)
}

func (m *TComponent) GetParentComponent() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := componentAPI().SysCallN(4, m.Instance())
	return AsComponent(r)
}

func (m *TComponent) HasParent() bool {
	if !m.IsValid() {
		return false
	}
	r := componentAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TComponent) UpdateAction(action IBasicAction) bool {
	if !m.IsValid() {
		return false
	}
	r := componentAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(action))
	return api.GoBool(r)
}

func (m *TComponent) DestroyComponents() {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(7, m.Instance())
}

func (m *TComponent) Destroying() {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(8, m.Instance())
}

func (m *TComponent) FreeNotification(component IComponent) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(component))
}

func (m *TComponent) RemoveFreeNotification(component IComponent) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(component))
}

func (m *TComponent) FreeOnRelease() {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(11, m.Instance())
}

func (m *TComponent) InsertComponent(component IComponent) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(component))
}

func (m *TComponent) RemoveComponent(component IComponent) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(component))
}

func (m *TComponent) SetSubComponent(subComponent bool) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(14, m.Instance(), api.PasBool(subComponent))
}

func (m *TComponent) Components(index int32) IComponent {
	if !m.IsValid() {
		return nil
	}
	r := componentAPI().SysCallN(15, m.Instance(), uintptr(index))
	return AsComponent(r)
}

func (m *TComponent) ComponentCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := componentAPI().SysCallN(16, m.Instance())
	return int32(r)
}

func (m *TComponent) ComponentIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := componentAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TComponent) SetComponentIndex(value int32) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TComponent) ComponentState() types.TComponentState {
	if !m.IsValid() {
		return 0
	}
	r := componentAPI().SysCallN(18, m.Instance())
	return types.TComponentState(r)
}

func (m *TComponent) ComponentStyle() types.TComponentStyle {
	if !m.IsValid() {
		return 0
	}
	r := componentAPI().SysCallN(19, m.Instance())
	return types.TComponentStyle(r)
}

func (m *TComponent) DesignInfo() int32 {
	if !m.IsValid() {
		return 0
	}
	r := componentAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TComponent) SetDesignInfo(value int32) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TComponent) Owner() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := componentAPI().SysCallN(21, m.Instance())
	return AsComponent(r)
}

func (m *TComponent) VCLComObject() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := componentAPI().SysCallN(22, 0, m.Instance())
	return uintptr(r)
}

func (m *TComponent) SetVCLComObject(value uintptr) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TComponent) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := componentAPI().SysCallN(23, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TComponent) SetName(value string) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(23, 1, m.Instance(), api.PasStr(value))
}

func (m *TComponent) Tag() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := componentAPI().SysCallN(24, 0, m.Instance())
	return uintptr(r)
}

func (m *TComponent) SetTag(value uintptr) {
	if !m.IsValid() {
		return
	}
	componentAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

// NewComponent class constructor
func NewComponent(owner IComponent) IComponent {
	r := componentAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsComponent(r)
}

var (
	componentOnce   base.Once
	componentImport *imports.Imports = nil
)

func componentAPI() *imports.Imports {
	componentOnce.Do(func() {
		componentImport = api.NewDefaultImports()
		componentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TComponent_Create", 0), // constructor NewComponent
			/* 1 */ imports.NewTable("TComponent_ExecuteAction", 0), // function ExecuteAction
			/* 2 */ imports.NewTable("TComponent_FindComponent", 0), // function FindComponent
			/* 3 */ imports.NewTable("TComponent_GetEnumerator", 0), // function GetEnumerator
			/* 4 */ imports.NewTable("TComponent_GetParentComponent", 0), // function GetParentComponent
			/* 5 */ imports.NewTable("TComponent_HasParent", 0), // function HasParent
			/* 6 */ imports.NewTable("TComponent_UpdateAction", 0), // function UpdateAction
			/* 7 */ imports.NewTable("TComponent_DestroyComponents", 0), // procedure DestroyComponents
			/* 8 */ imports.NewTable("TComponent_Destroying", 0), // procedure Destroying
			/* 9 */ imports.NewTable("TComponent_FreeNotification", 0), // procedure FreeNotification
			/* 10 */ imports.NewTable("TComponent_RemoveFreeNotification", 0), // procedure RemoveFreeNotification
			/* 11 */ imports.NewTable("TComponent_FreeOnRelease", 0), // procedure FreeOnRelease
			/* 12 */ imports.NewTable("TComponent_InsertComponent", 0), // procedure InsertComponent
			/* 13 */ imports.NewTable("TComponent_RemoveComponent", 0), // procedure RemoveComponent
			/* 14 */ imports.NewTable("TComponent_SetSubComponent", 0), // procedure SetSubComponent
			/* 15 */ imports.NewTable("TComponent_Components", 0), // property Components
			/* 16 */ imports.NewTable("TComponent_ComponentCount", 0), // property ComponentCount
			/* 17 */ imports.NewTable("TComponent_ComponentIndex", 0), // property ComponentIndex
			/* 18 */ imports.NewTable("TComponent_ComponentState", 0), // property ComponentState
			/* 19 */ imports.NewTable("TComponent_ComponentStyle", 0), // property ComponentStyle
			/* 20 */ imports.NewTable("TComponent_DesignInfo", 0), // property DesignInfo
			/* 21 */ imports.NewTable("TComponent_Owner", 0), // property Owner
			/* 22 */ imports.NewTable("TComponent_VCLComObject", 0), // property VCLComObject
			/* 23 */ imports.NewTable("TComponent_Name", 0), // property Name
			/* 24 */ imports.NewTable("TComponent_Tag", 0), // property Tag
		}
	})
	return componentImport
}
