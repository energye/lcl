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

// ILazAccessibleObject Parent: IObject
type ILazAccessibleObject interface {
	IObject
	HandleAllocated() bool                                            // function
	FindOwnerWinControl() IWinControl                                 // function
	AddChildAccessibleObject(dataObject IObject) ILazAccessibleObject // function
	// GetChildAccessibleObjectWithDataObject
	//  These search only in the child objects added manually
	GetChildAccessibleObjectWithDataObject(dataObject IObject) ILazAccessibleObject // function
	GetChildAccessibleObjectsCount() int32                                          // function
	GetChildAccessibleObject(index int32) ILazAccessibleObject                      // function
	// GetFirstChildAccessibleObject
	//  These search in all subcontrols too
	GetFirstChildAccessibleObject() ILazAccessibleObject // function
	GetNextChildAccessibleObject() ILazAccessibleObject  // function
	// GetSelectedChildAccessibleObject
	//
	GetSelectedChildAccessibleObject() ILazAccessibleObject                     // function
	GetChildAccessibleObjectAtPos(pos types.TPoint) ILazAccessibleObject        // function
	GetEnumerator() ILazAccessibleObjectEnumerator                              // function
	InitializeHandle()                                                          // procedure
	SetAccessibleNameWithString(name string)                                    // procedure
	SetAccessibleDescriptionWithString(description string)                      // procedure
	SetAccessibleValueWithString(value string)                                  // procedure
	SetAccessibleRoleWithLazAccessibilityRole(role types.TLazAccessibilityRole) // procedure
	InsertChildAccessibleObject(object ILazAccessibleObject)                    // procedure
	ClearChildAccessibleObjects()                                               // procedure
	RemoveChildAccessibleObject(object ILazAccessibleObject, freeObject bool)   // procedure
	// AccessibleName
	//  Primary information
	AccessibleName() string                                                    // property AccessibleName Getter
	SetAccessibleNameToString(value string)                                    // property AccessibleName Setter
	AccessibleDescription() string                                             // property AccessibleDescription Getter
	SetAccessibleDescriptionToString(value string)                             // property AccessibleDescription Setter
	AccessibleValue() string                                                   // property AccessibleValue Getter
	SetAccessibleValueToString(value string)                                   // property AccessibleValue Setter
	AccessibleRole() types.TLazAccessibilityRole                               // property AccessibleRole Getter
	SetAccessibleRoleToLazAccessibilityRole(value types.TLazAccessibilityRole) // property AccessibleRole Setter
	Position() types.TPoint                                                    // property Position Getter
	SetPosition(value types.TPoint)                                            // property Position Setter
	Size() types.TSize                                                         // property Size Getter
	SetSize(value types.TSize)                                                 // property Size Setter
	Handle() uintptr                                                           // property Handle Getter
	SetHandle(value uintptr)                                                   // property Handle Setter
}

type TLazAccessibleObject struct {
	TObject
}

func (m *TLazAccessibleObject) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := lazAccessibleObjectAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TLazAccessibleObject) FindOwnerWinControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(2, m.Instance())
	return AsWinControl(r)
}

func (m *TLazAccessibleObject) AddChildAccessibleObject(dataObject IObject) ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(dataObject))
	return AsLazAccessibleObject(r)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectWithDataObject(dataObject IObject) ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(dataObject))
	return AsLazAccessibleObject(r)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectsCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazAccessibleObjectAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TLazAccessibleObject) GetChildAccessibleObject(index int32) ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(6, m.Instance(), uintptr(index))
	return AsLazAccessibleObject(r)
}

func (m *TLazAccessibleObject) GetFirstChildAccessibleObject() ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(7, m.Instance())
	return AsLazAccessibleObject(r)
}

func (m *TLazAccessibleObject) GetNextChildAccessibleObject() ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(8, m.Instance())
	return AsLazAccessibleObject(r)
}

func (m *TLazAccessibleObject) GetSelectedChildAccessibleObject() ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(9, m.Instance())
	return AsLazAccessibleObject(r)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectAtPos(pos types.TPoint) ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&pos)))
	return AsLazAccessibleObject(r)
}

func (m *TLazAccessibleObject) GetEnumerator() ILazAccessibleObjectEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := lazAccessibleObjectAPI().SysCallN(11, m.Instance())
	return AsLazAccessibleObjectEnumerator(r)
}

func (m *TLazAccessibleObject) InitializeHandle() {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(12, m.Instance())
}

func (m *TLazAccessibleObject) SetAccessibleNameWithString(name string) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(13, m.Instance(), api.PasStr(name))
}

func (m *TLazAccessibleObject) SetAccessibleDescriptionWithString(description string) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(14, m.Instance(), api.PasStr(description))
}

func (m *TLazAccessibleObject) SetAccessibleValueWithString(value string) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(15, m.Instance(), api.PasStr(value))
}

func (m *TLazAccessibleObject) SetAccessibleRoleWithLazAccessibilityRole(role types.TLazAccessibilityRole) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(16, m.Instance(), uintptr(role))
}

func (m *TLazAccessibleObject) InsertChildAccessibleObject(object ILazAccessibleObject) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(object))
}

func (m *TLazAccessibleObject) ClearChildAccessibleObjects() {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(18, m.Instance())
}

func (m *TLazAccessibleObject) RemoveChildAccessibleObject(object ILazAccessibleObject, freeObject bool) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(object), api.PasBool(freeObject))
}

func (m *TLazAccessibleObject) AccessibleName() string {
	if !m.IsValid() {
		return ""
	}
	r := lazAccessibleObjectAPI().SysCallN(20, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TLazAccessibleObject) SetAccessibleNameToString(value string) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(20, 1, m.Instance(), api.PasStr(value))
}

func (m *TLazAccessibleObject) AccessibleDescription() string {
	if !m.IsValid() {
		return ""
	}
	r := lazAccessibleObjectAPI().SysCallN(21, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TLazAccessibleObject) SetAccessibleDescriptionToString(value string) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(21, 1, m.Instance(), api.PasStr(value))
}

func (m *TLazAccessibleObject) AccessibleValue() string {
	if !m.IsValid() {
		return ""
	}
	r := lazAccessibleObjectAPI().SysCallN(22, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TLazAccessibleObject) SetAccessibleValueToString(value string) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(22, 1, m.Instance(), api.PasStr(value))
}

func (m *TLazAccessibleObject) AccessibleRole() types.TLazAccessibilityRole {
	if !m.IsValid() {
		return 0
	}
	r := lazAccessibleObjectAPI().SysCallN(23, 0, m.Instance())
	return types.TLazAccessibilityRole(r)
}

func (m *TLazAccessibleObject) SetAccessibleRoleToLazAccessibilityRole(value types.TLazAccessibilityRole) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TLazAccessibleObject) Position() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(24, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazAccessibleObject) SetPosition(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(24, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TLazAccessibleObject) Size() (result types.TSize) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(25, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazAccessibleObject) SetSize(value types.TSize) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(25, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TLazAccessibleObject) Handle() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := lazAccessibleObjectAPI().SysCallN(26, 0, m.Instance())
	return uintptr(r)
}

func (m *TLazAccessibleObject) SetHandle(value uintptr) {
	if !m.IsValid() {
		return
	}
	lazAccessibleObjectAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

// NewLazAccessibleObject class constructor
func NewLazAccessibleObject(owner IControl) ILazAccessibleObject {
	r := lazAccessibleObjectAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazAccessibleObject(r)
}

var (
	lazAccessibleObjectOnce   base.Once
	lazAccessibleObjectImport *imports.Imports = nil
)

func lazAccessibleObjectAPI() *imports.Imports {
	lazAccessibleObjectOnce.Do(func() {
		lazAccessibleObjectImport = api.NewDefaultImports()
		lazAccessibleObjectImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazAccessibleObject_Create", 0), // constructor NewLazAccessibleObject
			/* 1 */ imports.NewTable("TLazAccessibleObject_HandleAllocated", 0), // function HandleAllocated
			/* 2 */ imports.NewTable("TLazAccessibleObject_FindOwnerWinControl", 0), // function FindOwnerWinControl
			/* 3 */ imports.NewTable("TLazAccessibleObject_AddChildAccessibleObject", 0), // function AddChildAccessibleObject
			/* 4 */ imports.NewTable("TLazAccessibleObject_GetChildAccessibleObjectWithDataObject", 0), // function GetChildAccessibleObjectWithDataObject
			/* 5 */ imports.NewTable("TLazAccessibleObject_GetChildAccessibleObjectsCount", 0), // function GetChildAccessibleObjectsCount
			/* 6 */ imports.NewTable("TLazAccessibleObject_GetChildAccessibleObject", 0), // function GetChildAccessibleObject
			/* 7 */ imports.NewTable("TLazAccessibleObject_GetFirstChildAccessibleObject", 0), // function GetFirstChildAccessibleObject
			/* 8 */ imports.NewTable("TLazAccessibleObject_GetNextChildAccessibleObject", 0), // function GetNextChildAccessibleObject
			/* 9 */ imports.NewTable("TLazAccessibleObject_GetSelectedChildAccessibleObject", 0), // function GetSelectedChildAccessibleObject
			/* 10 */ imports.NewTable("TLazAccessibleObject_GetChildAccessibleObjectAtPos", 0), // function GetChildAccessibleObjectAtPos
			/* 11 */ imports.NewTable("TLazAccessibleObject_GetEnumerator", 0), // function GetEnumerator
			/* 12 */ imports.NewTable("TLazAccessibleObject_InitializeHandle", 0), // procedure InitializeHandle
			/* 13 */ imports.NewTable("TLazAccessibleObject_SetAccessibleNameWithString", 0), // procedure SetAccessibleNameWithString
			/* 14 */ imports.NewTable("TLazAccessibleObject_SetAccessibleDescriptionWithString", 0), // procedure SetAccessibleDescriptionWithString
			/* 15 */ imports.NewTable("TLazAccessibleObject_SetAccessibleValueWithString", 0), // procedure SetAccessibleValueWithString
			/* 16 */ imports.NewTable("TLazAccessibleObject_SetAccessibleRoleWithLazAccessibilityRole", 0), // procedure SetAccessibleRoleWithLazAccessibilityRole
			/* 17 */ imports.NewTable("TLazAccessibleObject_InsertChildAccessibleObject", 0), // procedure InsertChildAccessibleObject
			/* 18 */ imports.NewTable("TLazAccessibleObject_ClearChildAccessibleObjects", 0), // procedure ClearChildAccessibleObjects
			/* 19 */ imports.NewTable("TLazAccessibleObject_RemoveChildAccessibleObject", 0), // procedure RemoveChildAccessibleObject
			/* 20 */ imports.NewTable("TLazAccessibleObject_AccessibleName", 0), // property AccessibleName
			/* 21 */ imports.NewTable("TLazAccessibleObject_AccessibleDescription", 0), // property AccessibleDescription
			/* 22 */ imports.NewTable("TLazAccessibleObject_AccessibleValue", 0), // property AccessibleValue
			/* 23 */ imports.NewTable("TLazAccessibleObject_AccessibleRole", 0), // property AccessibleRole
			/* 24 */ imports.NewTable("TLazAccessibleObject_Position", 0), // property Position
			/* 25 */ imports.NewTable("TLazAccessibleObject_Size", 0), // property Size
			/* 26 */ imports.NewTable("TLazAccessibleObject_Handle", 0), // property Handle
		}
	})
	return lazAccessibleObjectImport
}
