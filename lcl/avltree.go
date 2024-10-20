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

// IAVLTree Parent: IObject
type IAVLTree interface {
	IObject
	Root() IAVLTreeNode                                                                              // property
	Count() SizeInt                                                                                  // property
	NewNode() IAVLTreeNode                                                                           // function
	Add(Data uintptr) IAVLTreeNode                                                                   // function
	AddAscendingSequence(Data uintptr, LastAdded IAVLTreeNode, Successor *IAVLTreeNode) IAVLTreeNode // function
	Remove(Data uintptr) bool                                                                        // function
	RemovePointer(Data uintptr) bool                                                                 // function
	IsEqual(aTree IAVLTree, CheckDataPointer bool) bool                                              // function
	Compare(Data1, Data2 uintptr) int32                                                              // function
	Find(Data uintptr) IAVLTreeNode                                                                  // function
	FindSuccessor(ANode IAVLTreeNode) IAVLTreeNode                                                   // function
	FindPrecessor(ANode IAVLTreeNode) IAVLTreeNode                                                   // function
	FindLowest() IAVLTreeNode                                                                        // function
	FindHighest() IAVLTreeNode                                                                       // function
	FindNearest(Data uintptr) IAVLTreeNode                                                           // function
	FindPointer(Data uintptr) IAVLTreeNode                                                           // function
	FindLeftMost(Data uintptr) IAVLTreeNode                                                          // function
	FindRightMost(Data uintptr) IAVLTreeNode                                                         // function
	FindLeftMostSameKey(ANode IAVLTreeNode) IAVLTreeNode                                             // function
	FindRightMostSameKey(ANode IAVLTreeNode) IAVLTreeNode                                            // function
	GetEnumerator() IAVLTreeNodeEnumerator                                                           // function
	GetEnumeratorHighToLow() IAVLTreeNodeEnumerator                                                  // function
	NodeToReportStr(aNode IAVLTreeNode) string                                                       // function
	ReportAsString() string                                                                          // function
	DisposeNode(ANode IAVLTreeNode)                                                                  // procedure
	Add1(ANode IAVLTreeNode)                                                                         // procedure
	Delete(ANode IAVLTreeNode)                                                                       // procedure
	MoveDataLeftMost(ANode *IAVLTreeNode)                                                            // procedure
	MoveDataRightMost(ANode *IAVLTreeNode)                                                           // procedure
	Clear()                                                                                          // procedure
	FreeAndClear()                                                                                   // procedure
	FreeAndDelete(ANode IAVLTreeNode)                                                                // procedure
	Assign(aTree IAVLTree)                                                                           // procedure
	ConsistencyCheck()                                                                               // procedure
	WriteReportToStream(s IStream)                                                                   // procedure
}

// TAVLTree Parent: TObject
type TAVLTree struct {
	TObject
}

func NewAVLTree() IAVLTree {
	r1 := aVLTreeImportAPI().SysCallN(9)
	return AsAVLTree(r1)
}

func (m *TAVLTree) Root() IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(35, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) Count() SizeInt {
	r1 := aVLTreeImportAPI().SysCallN(8, m.Instance())
	return SizeInt(r1)
}

func (m *TAVLTree) NewNode() IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(30, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) Add(Data uintptr) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(0, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) AddAscendingSequence(Data uintptr, LastAdded IAVLTreeNode, Successor *IAVLTreeNode) IAVLTreeNode {
	var result2 uintptr
	r1 := aVLTreeImportAPI().SysCallN(2, m.Instance(), uintptr(Data), GetObjectUintptr(LastAdded), uintptr(unsafePointer(&result2)))
	*Successor = AsAVLTreeNode(result2)
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) Remove(Data uintptr) bool {
	r1 := aVLTreeImportAPI().SysCallN(32, m.Instance(), uintptr(Data))
	return GoBool(r1)
}

func (m *TAVLTree) RemovePointer(Data uintptr) bool {
	r1 := aVLTreeImportAPI().SysCallN(33, m.Instance(), uintptr(Data))
	return GoBool(r1)
}

func (m *TAVLTree) IsEqual(aTree IAVLTree, CheckDataPointer bool) bool {
	r1 := aVLTreeImportAPI().SysCallN(27, m.Instance(), GetObjectUintptr(aTree), PascalBool(CheckDataPointer))
	return GoBool(r1)
}

func (m *TAVLTree) Compare(Data1, Data2 uintptr) int32 {
	r1 := aVLTreeImportAPI().SysCallN(6, m.Instance(), uintptr(Data1), uintptr(Data2))
	return int32(r1)
}

func (m *TAVLTree) Find(Data uintptr) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(12, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindSuccessor(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindPrecessor(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(19, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindLowest() IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(16, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindHighest() IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(13, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindNearest(Data uintptr) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(17, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindPointer(Data uintptr) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(18, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindLeftMost(Data uintptr) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(14, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindRightMost(Data uintptr) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(20, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindLeftMostSameKey(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindRightMostSameKey(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := aVLTreeImportAPI().SysCallN(21, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) GetEnumerator() IAVLTreeNodeEnumerator {
	r1 := aVLTreeImportAPI().SysCallN(25, m.Instance())
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTree) GetEnumeratorHighToLow() IAVLTreeNodeEnumerator {
	r1 := aVLTreeImportAPI().SysCallN(26, m.Instance())
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTree) NodeToReportStr(aNode IAVLTreeNode) string {
	r1 := aVLTreeImportAPI().SysCallN(31, m.Instance(), GetObjectUintptr(aNode))
	return GoStr(r1)
}

func (m *TAVLTree) ReportAsString() string {
	r1 := aVLTreeImportAPI().SysCallN(34, m.Instance())
	return GoStr(r1)
}

func AVLTreeClass() TClass {
	ret := aVLTreeImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TAVLTree) DisposeNode(ANode IAVLTreeNode) {
	aVLTreeImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) Add1(ANode IAVLTreeNode) {
	aVLTreeImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) Delete(ANode IAVLTreeNode) {
	aVLTreeImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) MoveDataLeftMost(ANode *IAVLTreeNode) {
	var result0 uintptr
	aVLTreeImportAPI().SysCallN(28, m.Instance(), uintptr(unsafePointer(&result0)))
	*ANode = AsAVLTreeNode(result0)
}

func (m *TAVLTree) MoveDataRightMost(ANode *IAVLTreeNode) {
	var result0 uintptr
	aVLTreeImportAPI().SysCallN(29, m.Instance(), uintptr(unsafePointer(&result0)))
	*ANode = AsAVLTreeNode(result0)
}

func (m *TAVLTree) Clear() {
	aVLTreeImportAPI().SysCallN(5, m.Instance())
}

func (m *TAVLTree) FreeAndClear() {
	aVLTreeImportAPI().SysCallN(23, m.Instance())
}

func (m *TAVLTree) FreeAndDelete(ANode IAVLTreeNode) {
	aVLTreeImportAPI().SysCallN(24, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) Assign(aTree IAVLTree) {
	aVLTreeImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(aTree))
}

func (m *TAVLTree) ConsistencyCheck() {
	aVLTreeImportAPI().SysCallN(7, m.Instance())
}

func (m *TAVLTree) WriteReportToStream(s IStream) {
	aVLTreeImportAPI().SysCallN(36, m.Instance(), GetObjectUintptr(s))
}

var (
	aVLTreeImport       *imports.Imports = nil
	aVLTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("AVLTree_Add", 0),
		/*1*/ imports.NewTable("AVLTree_Add1", 0),
		/*2*/ imports.NewTable("AVLTree_AddAscendingSequence", 0),
		/*3*/ imports.NewTable("AVLTree_Assign", 0),
		/*4*/ imports.NewTable("AVLTree_Class", 0),
		/*5*/ imports.NewTable("AVLTree_Clear", 0),
		/*6*/ imports.NewTable("AVLTree_Compare", 0),
		/*7*/ imports.NewTable("AVLTree_ConsistencyCheck", 0),
		/*8*/ imports.NewTable("AVLTree_Count", 0),
		/*9*/ imports.NewTable("AVLTree_Create", 0),
		/*10*/ imports.NewTable("AVLTree_Delete", 0),
		/*11*/ imports.NewTable("AVLTree_DisposeNode", 0),
		/*12*/ imports.NewTable("AVLTree_Find", 0),
		/*13*/ imports.NewTable("AVLTree_FindHighest", 0),
		/*14*/ imports.NewTable("AVLTree_FindLeftMost", 0),
		/*15*/ imports.NewTable("AVLTree_FindLeftMostSameKey", 0),
		/*16*/ imports.NewTable("AVLTree_FindLowest", 0),
		/*17*/ imports.NewTable("AVLTree_FindNearest", 0),
		/*18*/ imports.NewTable("AVLTree_FindPointer", 0),
		/*19*/ imports.NewTable("AVLTree_FindPrecessor", 0),
		/*20*/ imports.NewTable("AVLTree_FindRightMost", 0),
		/*21*/ imports.NewTable("AVLTree_FindRightMostSameKey", 0),
		/*22*/ imports.NewTable("AVLTree_FindSuccessor", 0),
		/*23*/ imports.NewTable("AVLTree_FreeAndClear", 0),
		/*24*/ imports.NewTable("AVLTree_FreeAndDelete", 0),
		/*25*/ imports.NewTable("AVLTree_GetEnumerator", 0),
		/*26*/ imports.NewTable("AVLTree_GetEnumeratorHighToLow", 0),
		/*27*/ imports.NewTable("AVLTree_IsEqual", 0),
		/*28*/ imports.NewTable("AVLTree_MoveDataLeftMost", 0),
		/*29*/ imports.NewTable("AVLTree_MoveDataRightMost", 0),
		/*30*/ imports.NewTable("AVLTree_NewNode", 0),
		/*31*/ imports.NewTable("AVLTree_NodeToReportStr", 0),
		/*32*/ imports.NewTable("AVLTree_Remove", 0),
		/*33*/ imports.NewTable("AVLTree_RemovePointer", 0),
		/*34*/ imports.NewTable("AVLTree_ReportAsString", 0),
		/*35*/ imports.NewTable("AVLTree_Root", 0),
		/*36*/ imports.NewTable("AVLTree_WriteReportToStream", 0),
	}
)

func aVLTreeImportAPI() *imports.Imports {
	if aVLTreeImport == nil {
		aVLTreeImport = NewDefaultImports()
		aVLTreeImport.SetImportTable(aVLTreeImportTables)
		aVLTreeImportTables = nil
	}
	return aVLTreeImport
}
