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

// ITreeFilterEdit Parent: ICustomControlFilterEdit
type ITreeFilterEdit interface {
	ICustomControlFilterEdit
	GetExistingBranch(rootNode ITreeNode) ITreeFilterBranch    // function
	GetCleanBranch(rootNode ITreeNode) ITreeFilterBranch       // function
	DeleteBranch(rootNode ITreeNode) bool                      // function
	SetTreeFilterSilently(tree ICustomTreeView, filter string) // procedure
	ImageIndexDirectory() int32                                // property ImageIndexDirectory Getter
	SetImageIndexDirectory(value int32)                        // property ImageIndexDirectory Setter
	ScrolledPos() types.TPoint                                 // property ScrolledPos Getter
	SelectionList() IStringList                                // property SelectionList Getter
	ShowDirHierarchy() bool                                    // property ShowDirHierarchy Getter
	SetShowDirHierarchy(value bool)                            // property ShowDirHierarchy Setter
	FilteredTreeview() ICustomTreeView                         // property FilteredTreeview Getter
	SetFilteredTreeview(value ICustomTreeView)                 // property FilteredTreeview Setter
	ExpandAllInitially() bool                                  // property ExpandAllInitially Getter
	SetExpandAllInitially(value bool)                          // property ExpandAllInitially Setter
	SetOnGetImageIndex(fn TImageIndexEvent)                    // property event
	SetOnFilterNode(fn TFilterNodeEvent)                       // property event
}

type TTreeFilterEdit struct {
	TCustomControlFilterEdit
}

func (m *TTreeFilterEdit) GetExistingBranch(rootNode ITreeNode) ITreeFilterBranch {
	if !m.IsValid() {
		return nil
	}
	r := treeFilterEditAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(rootNode))
	return AsTreeFilterBranch(r)
}

func (m *TTreeFilterEdit) GetCleanBranch(rootNode ITreeNode) ITreeFilterBranch {
	if !m.IsValid() {
		return nil
	}
	r := treeFilterEditAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(rootNode))
	return AsTreeFilterBranch(r)
}

func (m *TTreeFilterEdit) DeleteBranch(rootNode ITreeNode) bool {
	if !m.IsValid() {
		return false
	}
	r := treeFilterEditAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(rootNode))
	return api.GoBool(r)
}

func (m *TTreeFilterEdit) SetTreeFilterSilently(tree ICustomTreeView, filter string) {
	if !m.IsValid() {
		return
	}
	treeFilterEditAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(tree), api.PasStr(filter))
}

func (m *TTreeFilterEdit) ImageIndexDirectory() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeFilterEditAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TTreeFilterEdit) SetImageIndexDirectory(value int32) {
	if !m.IsValid() {
		return
	}
	treeFilterEditAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TTreeFilterEdit) ScrolledPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	treeFilterEditAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TTreeFilterEdit) SelectionList() IStringList {
	if !m.IsValid() {
		return nil
	}
	r := treeFilterEditAPI().SysCallN(7, m.Instance())
	return AsStringList(r)
}

func (m *TTreeFilterEdit) ShowDirHierarchy() bool {
	if !m.IsValid() {
		return false
	}
	r := treeFilterEditAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeFilterEdit) SetShowDirHierarchy(value bool) {
	if !m.IsValid() {
		return
	}
	treeFilterEditAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeFilterEdit) FilteredTreeview() ICustomTreeView {
	if !m.IsValid() {
		return nil
	}
	r := treeFilterEditAPI().SysCallN(9, 0, m.Instance())
	return AsCustomTreeView(r)
}

func (m *TTreeFilterEdit) SetFilteredTreeview(value ICustomTreeView) {
	if !m.IsValid() {
		return
	}
	treeFilterEditAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TTreeFilterEdit) ExpandAllInitially() bool {
	if !m.IsValid() {
		return false
	}
	r := treeFilterEditAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeFilterEdit) SetExpandAllInitially(value bool) {
	if !m.IsValid() {
		return
	}
	treeFilterEditAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeFilterEdit) SetOnGetImageIndex(fn TImageIndexEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTImageIndexEvent(fn)
	base.SetEvent(m, 11, treeFilterEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeFilterEdit) SetOnFilterNode(fn TFilterNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTFilterNodeEvent(fn)
	base.SetEvent(m, 12, treeFilterEditAPI(), api.MakeEventDataPtr(cb))
}

// NewTreeFilterEdit class constructor
func NewTreeFilterEdit(owner IComponent) ITreeFilterEdit {
	r := treeFilterEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsTreeFilterEdit(r)
}

func TTreeFilterEditClass() types.TClass {
	r := treeFilterEditAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	treeFilterEditOnce   base.Once
	treeFilterEditImport *imports.Imports = nil
)

func treeFilterEditAPI() *imports.Imports {
	treeFilterEditOnce.Do(func() {
		treeFilterEditImport = api.NewDefaultImports()
		treeFilterEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTreeFilterEdit_Create", 0), // constructor NewTreeFilterEdit
			/* 1 */ imports.NewTable("TTreeFilterEdit_GetExistingBranch", 0), // function GetExistingBranch
			/* 2 */ imports.NewTable("TTreeFilterEdit_GetCleanBranch", 0), // function GetCleanBranch
			/* 3 */ imports.NewTable("TTreeFilterEdit_DeleteBranch", 0), // function DeleteBranch
			/* 4 */ imports.NewTable("TTreeFilterEdit_SetTreeFilterSilently", 0), // procedure SetTreeFilterSilently
			/* 5 */ imports.NewTable("TTreeFilterEdit_ImageIndexDirectory", 0), // property ImageIndexDirectory
			/* 6 */ imports.NewTable("TTreeFilterEdit_ScrolledPos", 0), // property ScrolledPos
			/* 7 */ imports.NewTable("TTreeFilterEdit_SelectionList", 0), // property SelectionList
			/* 8 */ imports.NewTable("TTreeFilterEdit_ShowDirHierarchy", 0), // property ShowDirHierarchy
			/* 9 */ imports.NewTable("TTreeFilterEdit_FilteredTreeview", 0), // property FilteredTreeview
			/* 10 */ imports.NewTable("TTreeFilterEdit_ExpandAllInitially", 0), // property ExpandAllInitially
			/* 11 */ imports.NewTable("TTreeFilterEdit_OnGetImageIndex", 0), // event OnGetImageIndex
			/* 12 */ imports.NewTable("TTreeFilterEdit_OnFilterNode", 0), // event OnFilterNode
			/* 13 */ imports.NewTable("TTreeFilterEdit_TClass", 0), // function TTreeFilterEditClass
		}
	})
	return treeFilterEditImport
}
