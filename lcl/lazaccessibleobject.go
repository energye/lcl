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

// ILazAccessibleObject Parent: IObject
type ILazAccessibleObject interface {
	IObject
	AccessibleName() string                                                          // property
	SetAccessibleName(AValue string)                                                 // property
	AccessibleDescription() string                                                   // property
	SetAccessibleDescription(AValue string)                                          // property
	AccessibleValue() string                                                         // property
	SetAccessibleValue(AValue string)                                                // property
	AccessibleRole() TLazAccessibilityRole                                           // property
	SetAccessibleRole(AValue TLazAccessibilityRole)                                  // property
	Position() (resultPoint TPoint)                                                  // property
	SetPosition(AValue *TPoint)                                                      // property
	Size() (resultSize TSize)                                                        // property
	SetSize(AValue *TSize)                                                           // property
	Handle() uint32                                                                  // property
	SetHandle(AValue uint32)                                                         // property
	HandleAllocated() bool                                                           // function
	FindOwnerWinControl() IWinControl                                                // function
	AddChildAccessibleObject(ADataObject IObject) ILazAccessibleObject               // function
	GetChildAccessibleObjectWithDataObject(ADataObject IObject) ILazAccessibleObject // function
	GetChildAccessibleObjectsCount() int32                                           // function
	GetChildAccessibleObject(AIndex int32) ILazAccessibleObject                      // function
	GetFirstChildAccessibleObject() ILazAccessibleObject                             // function
	GetNextChildAccessibleObject() ILazAccessibleObject                              // function
	GetSelectedChildAccessibleObject() ILazAccessibleObject                          // function
	GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject                 // function
	GetEnumerator() ILazAccessibleObjectEnumerator                                   // function
	InitializeHandle()                                                               // procedure
	SetAccessibleName1(AName string)                                                 // procedure
	SetAccessibleDescription1(ADescription string)                                   // procedure
	SetAccessibleValue1(AValue string)                                               // procedure
	SetAccessibleRole1(ARole TLazAccessibilityRole)                                  // procedure
	InsertChildAccessibleObject(AObject ILazAccessibleObject)                        // procedure
	ClearChildAccessibleObjects()                                                    // procedure
	RemoveChildAccessibleObject(AObject ILazAccessibleObject, AFreeObject bool)      // procedure
}

// TLazAccessibleObject Parent: TObject
type TLazAccessibleObject struct {
	TObject
}

func NewLazAccessibleObject(AOwner IControl) ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(7, GetObjectUintptr(AOwner))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) AccessibleName() string {
	r1 := lazAccessibleObjectImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazAccessibleObject) SetAccessibleName(AValue string) {
	lazAccessibleObjectImportAPI().SysCallN(1, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) AccessibleDescription() string {
	r1 := lazAccessibleObjectImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazAccessibleObject) SetAccessibleDescription(AValue string) {
	lazAccessibleObjectImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) AccessibleValue() string {
	r1 := lazAccessibleObjectImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazAccessibleObject) SetAccessibleValue(AValue string) {
	lazAccessibleObjectImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) AccessibleRole() TLazAccessibilityRole {
	r1 := lazAccessibleObjectImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TLazAccessibilityRole(r1)
}

func (m *TLazAccessibleObject) SetAccessibleRole(AValue TLazAccessibilityRole) {
	lazAccessibleObjectImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazAccessibleObject) Position() (resultPoint TPoint) {
	lazAccessibleObjectImportAPI().SysCallN(21, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TLazAccessibleObject) SetPosition(AValue *TPoint) {
	lazAccessibleObjectImportAPI().SysCallN(21, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TLazAccessibleObject) Size() (resultSize TSize) {
	lazAccessibleObjectImportAPI().SysCallN(27, 0, m.Instance(), uintptr(unsafePointer(&resultSize)), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TLazAccessibleObject) SetSize(AValue *TSize) {
	lazAccessibleObjectImportAPI().SysCallN(27, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TLazAccessibleObject) Handle() uint32 {
	r1 := lazAccessibleObjectImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazAccessibleObject) SetHandle(AValue uint32) {
	lazAccessibleObjectImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazAccessibleObject) HandleAllocated() bool {
	r1 := lazAccessibleObjectImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TLazAccessibleObject) FindOwnerWinControl() IWinControl {
	r1 := lazAccessibleObjectImportAPI().SysCallN(8, m.Instance())
	return AsWinControl(r1)
}

func (m *TLazAccessibleObject) AddChildAccessibleObject(ADataObject IObject) ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(ADataObject))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectWithDataObject(ADataObject IObject) ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(ADataObject))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectsCount() int32 {
	r1 := lazAccessibleObjectImportAPI().SysCallN(12, m.Instance())
	return int32(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObject(AIndex int32) ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(9, m.Instance(), uintptr(AIndex))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetFirstChildAccessibleObject() ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(14, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetNextChildAccessibleObject() ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(15, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetSelectedChildAccessibleObject() ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(16, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject {
	r1 := lazAccessibleObjectImportAPI().SysCallN(10, m.Instance(), uintptr(unsafePointer(APos)))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetEnumerator() ILazAccessibleObjectEnumerator {
	r1 := lazAccessibleObjectImportAPI().SysCallN(13, m.Instance())
	return AsLazAccessibleObjectEnumerator(r1)
}

func LazAccessibleObjectClass() TClass {
	ret := lazAccessibleObjectImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TLazAccessibleObject) InitializeHandle() {
	lazAccessibleObjectImportAPI().SysCallN(19, m.Instance())
}

func (m *TLazAccessibleObject) SetAccessibleName1(AName string) {
	lazAccessibleObjectImportAPI().SysCallN(24, m.Instance(), PascalStr(AName))
}

func (m *TLazAccessibleObject) SetAccessibleDescription1(ADescription string) {
	lazAccessibleObjectImportAPI().SysCallN(23, m.Instance(), PascalStr(ADescription))
}

func (m *TLazAccessibleObject) SetAccessibleValue1(AValue string) {
	lazAccessibleObjectImportAPI().SysCallN(26, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) SetAccessibleRole1(ARole TLazAccessibilityRole) {
	lazAccessibleObjectImportAPI().SysCallN(25, m.Instance(), uintptr(ARole))
}

func (m *TLazAccessibleObject) InsertChildAccessibleObject(AObject ILazAccessibleObject) {
	lazAccessibleObjectImportAPI().SysCallN(20, m.Instance(), GetObjectUintptr(AObject))
}

func (m *TLazAccessibleObject) ClearChildAccessibleObjects() {
	lazAccessibleObjectImportAPI().SysCallN(6, m.Instance())
}

func (m *TLazAccessibleObject) RemoveChildAccessibleObject(AObject ILazAccessibleObject, AFreeObject bool) {
	lazAccessibleObjectImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(AObject), PascalBool(AFreeObject))
}

var (
	lazAccessibleObjectImport       *imports.Imports = nil
	lazAccessibleObjectImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazAccessibleObject_AccessibleDescription", 0),
		/*1*/ imports.NewTable("LazAccessibleObject_AccessibleName", 0),
		/*2*/ imports.NewTable("LazAccessibleObject_AccessibleRole", 0),
		/*3*/ imports.NewTable("LazAccessibleObject_AccessibleValue", 0),
		/*4*/ imports.NewTable("LazAccessibleObject_AddChildAccessibleObject", 0),
		/*5*/ imports.NewTable("LazAccessibleObject_Class", 0),
		/*6*/ imports.NewTable("LazAccessibleObject_ClearChildAccessibleObjects", 0),
		/*7*/ imports.NewTable("LazAccessibleObject_Create", 0),
		/*8*/ imports.NewTable("LazAccessibleObject_FindOwnerWinControl", 0),
		/*9*/ imports.NewTable("LazAccessibleObject_GetChildAccessibleObject", 0),
		/*10*/ imports.NewTable("LazAccessibleObject_GetChildAccessibleObjectAtPos", 0),
		/*11*/ imports.NewTable("LazAccessibleObject_GetChildAccessibleObjectWithDataObject", 0),
		/*12*/ imports.NewTable("LazAccessibleObject_GetChildAccessibleObjectsCount", 0),
		/*13*/ imports.NewTable("LazAccessibleObject_GetEnumerator", 0),
		/*14*/ imports.NewTable("LazAccessibleObject_GetFirstChildAccessibleObject", 0),
		/*15*/ imports.NewTable("LazAccessibleObject_GetNextChildAccessibleObject", 0),
		/*16*/ imports.NewTable("LazAccessibleObject_GetSelectedChildAccessibleObject", 0),
		/*17*/ imports.NewTable("LazAccessibleObject_Handle", 0),
		/*18*/ imports.NewTable("LazAccessibleObject_HandleAllocated", 0),
		/*19*/ imports.NewTable("LazAccessibleObject_InitializeHandle", 0),
		/*20*/ imports.NewTable("LazAccessibleObject_InsertChildAccessibleObject", 0),
		/*21*/ imports.NewTable("LazAccessibleObject_Position", 0),
		/*22*/ imports.NewTable("LazAccessibleObject_RemoveChildAccessibleObject", 0),
		/*23*/ imports.NewTable("LazAccessibleObject_SetAccessibleDescription1", 0),
		/*24*/ imports.NewTable("LazAccessibleObject_SetAccessibleName1", 0),
		/*25*/ imports.NewTable("LazAccessibleObject_SetAccessibleRole1", 0),
		/*26*/ imports.NewTable("LazAccessibleObject_SetAccessibleValue1", 0),
		/*27*/ imports.NewTable("LazAccessibleObject_Size", 0),
	}
)

func lazAccessibleObjectImportAPI() *imports.Imports {
	if lazAccessibleObjectImport == nil {
		lazAccessibleObjectImport = NewDefaultImports()
		lazAccessibleObjectImport.SetImportTable(lazAccessibleObjectImportTables)
		lazAccessibleObjectImportTables = nil
	}
	return lazAccessibleObjectImport
}
