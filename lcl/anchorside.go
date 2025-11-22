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

// IAnchorSide Parent: IPersistent
type IAnchorSide interface {
	IPersistent
	CheckSidePosition(newControl IControl, newSide types.TAnchorSideReference, outReferenceControl *IControl, outReferenceSide *types.TAnchorSideReference, outPosition *int32) bool // function
	IsAnchoredToParent(parentSide types.TAnchorKind) bool                                                                                                                            // function
	GetSidePosition(outReferenceControl *IControl, outReferenceSide *types.TAnchorSideReference, outPosition *int32)                                                                 // procedure
	FixCenterAnchoring()                                                                                                                                                             // procedure
	Owner() IControl                                                                                                                                                                 // property Owner Getter
	Kind() types.TAnchorKind                                                                                                                                                         // property Kind Getter
	Control() IControl                                                                                                                                                               // property Control Getter
	SetControl(value IControl)                                                                                                                                                       // property Control Setter
	Side() types.TAnchorSideReference                                                                                                                                                // property Side Getter
	SetSide(value types.TAnchorSideReference)                                                                                                                                        // property Side Setter
}

type TAnchorSide struct {
	TPersistent
}

func (m *TAnchorSide) CheckSidePosition(newControl IControl, newSide types.TAnchorSideReference, outReferenceControl *IControl, outReferenceSide *types.TAnchorSideReference, outPosition *int32) bool {
	if !m.IsValid() {
		return false
	}
	var referenceControlPtr uintptr
	var referenceSidePtr uintptr
	var positionPtr uintptr
	r := anchorSideAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(newControl), uintptr(newSide), uintptr(base.UnsafePointer(&referenceControlPtr)), uintptr(base.UnsafePointer(&referenceSidePtr)), uintptr(base.UnsafePointer(&positionPtr)))
	*outReferenceControl = AsControl(referenceControlPtr)
	*outReferenceSide = types.TAnchorSideReference(referenceSidePtr)
	*outPosition = int32(positionPtr)
	return api.GoBool(r)
}

func (m *TAnchorSide) IsAnchoredToParent(parentSide types.TAnchorKind) bool {
	if !m.IsValid() {
		return false
	}
	r := anchorSideAPI().SysCallN(2, m.Instance(), uintptr(parentSide))
	return api.GoBool(r)
}

func (m *TAnchorSide) GetSidePosition(outReferenceControl *IControl, outReferenceSide *types.TAnchorSideReference, outPosition *int32) {
	if !m.IsValid() {
		return
	}
	var referenceControlPtr uintptr
	var referenceSidePtr uintptr
	var positionPtr uintptr
	anchorSideAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&referenceControlPtr)), uintptr(base.UnsafePointer(&referenceSidePtr)), uintptr(base.UnsafePointer(&positionPtr)))
	*outReferenceControl = AsControl(referenceControlPtr)
	*outReferenceSide = types.TAnchorSideReference(referenceSidePtr)
	*outPosition = int32(positionPtr)
}

func (m *TAnchorSide) FixCenterAnchoring() {
	if !m.IsValid() {
		return
	}
	anchorSideAPI().SysCallN(4, m.Instance())
}

func (m *TAnchorSide) Owner() IControl {
	if !m.IsValid() {
		return nil
	}
	r := anchorSideAPI().SysCallN(5, m.Instance())
	return AsControl(r)
}

func (m *TAnchorSide) Kind() types.TAnchorKind {
	if !m.IsValid() {
		return 0
	}
	r := anchorSideAPI().SysCallN(6, m.Instance())
	return types.TAnchorKind(r)
}

func (m *TAnchorSide) Control() IControl {
	if !m.IsValid() {
		return nil
	}
	r := anchorSideAPI().SysCallN(7, 0, m.Instance())
	return AsControl(r)
}

func (m *TAnchorSide) SetControl(value IControl) {
	if !m.IsValid() {
		return
	}
	anchorSideAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TAnchorSide) Side() types.TAnchorSideReference {
	if !m.IsValid() {
		return 0
	}
	r := anchorSideAPI().SysCallN(8, 0, m.Instance())
	return types.TAnchorSideReference(r)
}

func (m *TAnchorSide) SetSide(value types.TAnchorSideReference) {
	if !m.IsValid() {
		return
	}
	anchorSideAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

// NewAnchorSide class constructor
func NewAnchorSide(theOwner IControl, theKind types.TAnchorKind) IAnchorSide {
	r := anchorSideAPI().SysCallN(0, base.GetObjectUintptr(theOwner), uintptr(theKind))
	return AsAnchorSide(r)
}

var (
	anchorSideOnce   base.Once
	anchorSideImport *imports.Imports = nil
)

func anchorSideAPI() *imports.Imports {
	anchorSideOnce.Do(func() {
		anchorSideImport = api.NewDefaultImports()
		anchorSideImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TAnchorSide_Create", 0), // constructor NewAnchorSide
			/* 1 */ imports.NewTable("TAnchorSide_CheckSidePosition", 0), // function CheckSidePosition
			/* 2 */ imports.NewTable("TAnchorSide_IsAnchoredToParent", 0), // function IsAnchoredToParent
			/* 3 */ imports.NewTable("TAnchorSide_GetSidePosition", 0), // procedure GetSidePosition
			/* 4 */ imports.NewTable("TAnchorSide_FixCenterAnchoring", 0), // procedure FixCenterAnchoring
			/* 5 */ imports.NewTable("TAnchorSide_Owner", 0), // property Owner
			/* 6 */ imports.NewTable("TAnchorSide_Kind", 0), // property Kind
			/* 7 */ imports.NewTable("TAnchorSide_Control", 0), // property Control
			/* 8 */ imports.NewTable("TAnchorSide_Side", 0), // property Side
		}
	})
	return anchorSideImport
}
