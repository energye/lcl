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

// ISynTextViewsManager Parent: IObject
type ISynTextViewsManager interface {
	IObject
	IndexOf(textView ISynEditStringsLinked) int32                                                        // function
	Count() int32                                                                                        // function
	ReconnectViews()                                                                                     // procedure
	AddTextViewWithSynEditStringsLinkedBool(textView ISynEditStringsLinked, asFirst bool)                // procedure
	AddTextViewWithSynEditStringsLinkedInt(textView ISynEditStringsLinked, index int32)                  // procedure
	AddTextViewWithSynEditStringsLinkedX2(textView ISynEditStringsLinked, anAfter ISynEditStringsLinked) // procedure
	RemoveSynTextView(textView ISynEditStringsLinked, destroy bool)                                      // procedure
	SynTextView(index int32) ISynEditStringsLinked                                                       // property SynTextView Getter
	SynTextViewByClass(index types.TSynEditStringsLinkedClass) ISynEditStringsLinked                     // property SynTextViewByClass Getter
}

type TSynTextViewsManager struct {
	TObject
}

func (m *TSynTextViewsManager) IndexOf(textView ISynEditStringsLinked) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synTextViewsManagerAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(textView))
	return int32(r)
}

func (m *TSynTextViewsManager) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synTextViewsManagerAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TSynTextViewsManager) ReconnectViews() {
	if !m.IsValid() {
		return
	}
	synTextViewsManagerAPI().SysCallN(3, m.Instance())
}

func (m *TSynTextViewsManager) AddTextViewWithSynEditStringsLinkedBool(textView ISynEditStringsLinked, asFirst bool) {
	if !m.IsValid() {
		return
	}
	synTextViewsManagerAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(textView), api.PasBool(asFirst))
}

func (m *TSynTextViewsManager) AddTextViewWithSynEditStringsLinkedInt(textView ISynEditStringsLinked, index int32) {
	if !m.IsValid() {
		return
	}
	synTextViewsManagerAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(textView), uintptr(index))
}

func (m *TSynTextViewsManager) AddTextViewWithSynEditStringsLinkedX2(textView ISynEditStringsLinked, anAfter ISynEditStringsLinked) {
	if !m.IsValid() {
		return
	}
	synTextViewsManagerAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(textView), base.GetObjectUintptr(anAfter))
}

func (m *TSynTextViewsManager) RemoveSynTextView(textView ISynEditStringsLinked, destroy bool) {
	if !m.IsValid() {
		return
	}
	synTextViewsManagerAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(textView), api.PasBool(destroy))
}

func (m *TSynTextViewsManager) SynTextView(index int32) ISynEditStringsLinked {
	if !m.IsValid() {
		return nil
	}
	r := synTextViewsManagerAPI().SysCallN(8, m.Instance(), uintptr(index))
	return AsSynEditStringsLinked(r)
}

func (m *TSynTextViewsManager) SynTextViewByClass(index types.TSynEditStringsLinkedClass) ISynEditStringsLinked {
	if !m.IsValid() {
		return nil
	}
	r := synTextViewsManagerAPI().SysCallN(9, m.Instance(), uintptr(index))
	return AsSynEditStringsLinked(r)
}

// NewSynTextViewsManager class constructor
func NewSynTextViewsManager(textBuffer ISynEditStringListBase, topViewChangedCallback INotifyEventDelegate) ISynTextViewsManager {
	r := synTextViewsManagerAPI().SysCallN(0, base.GetObjectUintptr(textBuffer), base.GetObjectUintptr(topViewChangedCallback))
	return AsSynTextViewsManager(r)
}

var (
	synTextViewsManagerOnce   base.Once
	synTextViewsManagerImport *imports.Imports = nil
)

func synTextViewsManagerAPI() *imports.Imports {
	synTextViewsManagerOnce.Do(func() {
		synTextViewsManagerImport = api.NewDefaultImports()
		synTextViewsManagerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynTextViewsManager_Create", 0), // constructor NewSynTextViewsManager
			/* 1 */ imports.NewTable("TSynTextViewsManager_IndexOf", 0), // function IndexOf
			/* 2 */ imports.NewTable("TSynTextViewsManager_Count", 0), // function Count
			/* 3 */ imports.NewTable("TSynTextViewsManager_ReconnectViews", 0), // procedure ReconnectViews
			/* 4 */ imports.NewTable("TSynTextViewsManager_AddTextViewWithSynEditStringsLinkedBool", 0), // procedure AddTextViewWithSynEditStringsLinkedBool
			/* 5 */ imports.NewTable("TSynTextViewsManager_AddTextViewWithSynEditStringsLinkedInt", 0), // procedure AddTextViewWithSynEditStringsLinkedInt
			/* 6 */ imports.NewTable("TSynTextViewsManager_AddTextViewWithSynEditStringsLinkedX2", 0), // procedure AddTextViewWithSynEditStringsLinkedX2
			/* 7 */ imports.NewTable("TSynTextViewsManager_RemoveSynTextView", 0), // procedure RemoveSynTextView
			/* 8 */ imports.NewTable("TSynTextViewsManager_SynTextView", 0), // property SynTextView
			/* 9 */ imports.NewTable("TSynTextViewsManager_SynTextViewByClass", 0), // property SynTextViewByClass
		}
	})
	return synTextViewsManagerImport
}
