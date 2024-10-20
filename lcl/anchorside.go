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

// IAnchorSide Parent: IPersistent
type IAnchorSide interface {
	IPersistent
	Owner() IControl                                                                                                                                                     // property
	Kind() TAnchorKind                                                                                                                                                   // property
	Control() IControl                                                                                                                                                   // property
	SetControl(AValue IControl)                                                                                                                                          // property
	Side() TAnchorSideReference                                                                                                                                          // property
	SetSide(AValue TAnchorSideReference)                                                                                                                                 // property
	CheckSidePosition(NewControl IControl, NewSide TAnchorSideReference, OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32) bool // function
	IsAnchoredToParent(ParentSide TAnchorKind) bool                                                                                                                      // function
	GetSidePosition(OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32)                                                           // procedure
	FixCenterAnchoring()                                                                                                                                                 // procedure
}

// TAnchorSide Parent: TPersistent
type TAnchorSide struct {
	TPersistent
}

func NewAnchorSide(TheOwner IControl, TheKind TAnchorKind) IAnchorSide {
	r1 := anchorSideImportAPI().SysCallN(3, GetObjectUintptr(TheOwner), uintptr(TheKind))
	return AsAnchorSide(r1)
}

func (m *TAnchorSide) Owner() IControl {
	r1 := anchorSideImportAPI().SysCallN(8, m.Instance())
	return AsControl(r1)
}

func (m *TAnchorSide) Kind() TAnchorKind {
	r1 := anchorSideImportAPI().SysCallN(7, m.Instance())
	return TAnchorKind(r1)
}

func (m *TAnchorSide) Control() IControl {
	r1 := anchorSideImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TAnchorSide) SetControl(AValue IControl) {
	anchorSideImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TAnchorSide) Side() TAnchorSideReference {
	r1 := anchorSideImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TAnchorSideReference(r1)
}

func (m *TAnchorSide) SetSide(AValue TAnchorSideReference) {
	anchorSideImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TAnchorSide) CheckSidePosition(NewControl IControl, NewSide TAnchorSideReference, OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32) bool {
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := anchorSideImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(NewControl), uintptr(NewSide), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*OutReferenceControl = AsControl(result2)
	*OutReferenceSide = TAnchorSideReference(result3)
	*OutPosition = int32(result4)
	return GoBool(r1)
}

func (m *TAnchorSide) IsAnchoredToParent(ParentSide TAnchorKind) bool {
	r1 := anchorSideImportAPI().SysCallN(6, m.Instance(), uintptr(ParentSide))
	return GoBool(r1)
}

func AnchorSideClass() TClass {
	ret := anchorSideImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TAnchorSide) GetSidePosition(OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32) {
	var result0 uintptr
	var result1 uintptr
	var result2 uintptr
	anchorSideImportAPI().SysCallN(5, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutReferenceControl = AsControl(result0)
	*OutReferenceSide = TAnchorSideReference(result1)
	*OutPosition = int32(result2)
}

func (m *TAnchorSide) FixCenterAnchoring() {
	anchorSideImportAPI().SysCallN(4, m.Instance())
}

var (
	anchorSideImport       *imports.Imports = nil
	anchorSideImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("AnchorSide_CheckSidePosition", 0),
		/*1*/ imports.NewTable("AnchorSide_Class", 0),
		/*2*/ imports.NewTable("AnchorSide_Control", 0),
		/*3*/ imports.NewTable("AnchorSide_Create", 0),
		/*4*/ imports.NewTable("AnchorSide_FixCenterAnchoring", 0),
		/*5*/ imports.NewTable("AnchorSide_GetSidePosition", 0),
		/*6*/ imports.NewTable("AnchorSide_IsAnchoredToParent", 0),
		/*7*/ imports.NewTable("AnchorSide_Kind", 0),
		/*8*/ imports.NewTable("AnchorSide_Owner", 0),
		/*9*/ imports.NewTable("AnchorSide_Side", 0),
	}
)

func anchorSideImportAPI() *imports.Imports {
	if anchorSideImport == nil {
		anchorSideImport = NewDefaultImports()
		anchorSideImport.SetImportTable(anchorSideImportTables)
		anchorSideImportTables = nil
	}
	return anchorSideImport
}
