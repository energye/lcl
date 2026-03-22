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

// ISynHighlighterAttributes Parent: ILazSynCustomTextAttributes
type ISynHighlighterAttributes interface {
	ILazSynCustomTextAttributes
	IsEnabled() bool                                                                                 // function
	LoadFromBorlandRegistry(rootKey types.HKEY, attrKey string, attrName string, oldStyle bool) bool // function
	LoadFromRegistry(reg IRegistry) bool                                                             // function
	SaveToRegistry(reg IRegistry) bool                                                               // function
	LoadFromFile(ini IIniFile) bool                                                                  // function
	SaveToFile(ini IIniFile) bool                                                                    // function
	InternalSaveDefaultValues()                                                                      // procedure
	IntegerStyle() int32                                                                             // property IntegerStyle Getter
	SetIntegerStyle(value int32)                                                                     // property IntegerStyle Setter
	IntegerStyleMask() int32                                                                         // property IntegerStyleMask Getter
	SetIntegerStyleMask(value int32)                                                                 // property IntegerStyleMask Setter
	Name() string                                                                                    // property Name Getter
	Caption() string                                                                                 // property Caption Getter
	StoredName() string                                                                              // property StoredName Getter
	SetStoredName(value string)                                                                      // property StoredName Setter
	SetOnChange(fn TNotifyEvent)                                                                     // property event
}

type TSynHighlighterAttributes struct {
	TLazSynCustomTextAttributes
}

func (m *TSynHighlighterAttributes) IsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterAttributesAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TSynHighlighterAttributes) LoadFromBorlandRegistry(rootKey types.HKEY, attrKey string, attrName string, oldStyle bool) bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterAttributesAPI().SysCallN(3, m.Instance(), uintptr(rootKey), api.PasStr(attrKey), api.PasStr(attrName), api.PasBool(oldStyle))
	return api.GoBool(r)
}

func (m *TSynHighlighterAttributes) LoadFromRegistry(reg IRegistry) bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterAttributesAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(reg))
	return api.GoBool(r)
}

func (m *TSynHighlighterAttributes) SaveToRegistry(reg IRegistry) bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterAttributesAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(reg))
	return api.GoBool(r)
}

func (m *TSynHighlighterAttributes) LoadFromFile(ini IIniFile) bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterAttributesAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(ini))
	return api.GoBool(r)
}

func (m *TSynHighlighterAttributes) SaveToFile(ini IIniFile) bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterAttributesAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(ini))
	return api.GoBool(r)
}

func (m *TSynHighlighterAttributes) InternalSaveDefaultValues() {
	if !m.IsValid() {
		return
	}
	synHighlighterAttributesAPI().SysCallN(8, m.Instance())
}

func (m *TSynHighlighterAttributes) IntegerStyle() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterAttributesAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TSynHighlighterAttributes) SetIntegerStyle(value int32) {
	if !m.IsValid() {
		return
	}
	synHighlighterAttributesAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TSynHighlighterAttributes) IntegerStyleMask() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterAttributesAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TSynHighlighterAttributes) SetIntegerStyleMask(value int32) {
	if !m.IsValid() {
		return
	}
	synHighlighterAttributesAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TSynHighlighterAttributes) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := synHighlighterAttributesAPI().SysCallN(11, m.Instance())
	return api.GoStr(r)
}

func (m *TSynHighlighterAttributes) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := synHighlighterAttributesAPI().SysCallN(12, m.Instance())
	return api.GoStr(r)
}

func (m *TSynHighlighterAttributes) StoredName() string {
	if !m.IsValid() {
		return ""
	}
	r := synHighlighterAttributesAPI().SysCallN(13, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TSynHighlighterAttributes) SetStoredName(value string) {
	if !m.IsValid() {
		return
	}
	synHighlighterAttributesAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynHighlighterAttributes) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, synHighlighterAttributesAPI(), api.MakeEventDataPtr(cb))
}

// NewSynHighlighterAttributes class constructor
func NewSynHighlighterAttributes() ISynHighlighterAttributes {
	r := synHighlighterAttributesAPI().SysCallN(0)
	return AsSynHighlighterAttributes(r)
}

// NewSynHighlighterAttributesWithStrX2 class constructor
func NewSynHighlighterAttributesWithStrX2(caption string, storedName string) ISynHighlighterAttributes {
	r := synHighlighterAttributesAPI().SysCallN(1, api.PasStr(caption), api.PasStr(storedName))
	return AsSynHighlighterAttributes(r)
}

var (
	synHighlighterAttributesOnce   base.Once
	synHighlighterAttributesImport *imports.Imports = nil
)

func synHighlighterAttributesAPI() *imports.Imports {
	synHighlighterAttributesOnce.Do(func() {
		synHighlighterAttributesImport = api.NewDefaultImports()
		synHighlighterAttributesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHighlighterAttributes_Create", 0), // constructor NewSynHighlighterAttributes
			/* 1 */ imports.NewTable("TSynHighlighterAttributes_CreateWithStrX2", 0), // constructor NewSynHighlighterAttributesWithStrX2
			/* 2 */ imports.NewTable("TSynHighlighterAttributes_IsEnabled", 0), // function IsEnabled
			/* 3 */ imports.NewTable("TSynHighlighterAttributes_LoadFromBorlandRegistry", 0), // function LoadFromBorlandRegistry
			/* 4 */ imports.NewTable("TSynHighlighterAttributes_LoadFromRegistry", 0), // function LoadFromRegistry
			/* 5 */ imports.NewTable("TSynHighlighterAttributes_SaveToRegistry", 0), // function SaveToRegistry
			/* 6 */ imports.NewTable("TSynHighlighterAttributes_LoadFromFile", 0), // function LoadFromFile
			/* 7 */ imports.NewTable("TSynHighlighterAttributes_SaveToFile", 0), // function SaveToFile
			/* 8 */ imports.NewTable("TSynHighlighterAttributes_InternalSaveDefaultValues", 0), // procedure InternalSaveDefaultValues
			/* 9 */ imports.NewTable("TSynHighlighterAttributes_IntegerStyle", 0), // property IntegerStyle
			/* 10 */ imports.NewTable("TSynHighlighterAttributes_IntegerStyleMask", 0), // property IntegerStyleMask
			/* 11 */ imports.NewTable("TSynHighlighterAttributes_Name", 0), // property Name
			/* 12 */ imports.NewTable("TSynHighlighterAttributes_Caption", 0), // property Caption
			/* 13 */ imports.NewTable("TSynHighlighterAttributes_StoredName", 0), // property StoredName
			/* 14 */ imports.NewTable("TSynHighlighterAttributes_OnChange", 0), // event OnChange
		}
	})
	return synHighlighterAttributesImport
}
