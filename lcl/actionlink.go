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
)

// IActionLink Parent: IBasicActionLink
type IActionLink interface {
	IBasicActionLink
	IsCaptionLinked() bool     // function
	IsCheckedLinked() bool     // function
	IsEnabledLinked() bool     // function
	IsGroupIndexLinked() bool  // function
	IsHelpContextLinked() bool // function
	IsHelpLinked() bool        // function
	IsHintLinked() bool        // function
	IsImageIndexLinked() bool  // function
	IsShortCutLinked() bool    // function
	IsVisibleLinked() bool     // function
}

type TActionLink struct {
	TBasicActionLink
}

func (m *TActionLink) IsCaptionLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsCheckedLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsEnabledLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsGroupIndexLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsHelpContextLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsHelpLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsHintLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsImageIndexLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsShortCutLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TActionLink) IsVisibleLinked() bool {
	if !m.IsValid() {
		return false
	}
	r := actionLinkAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

// NewActionLink class constructor
func NewActionLink(client IObject) IActionLink {
	r := actionLinkAPI().SysCallN(0, base.GetObjectUintptr(client))
	return AsActionLink(r)
}

var (
	actionLinkOnce   base.Once
	actionLinkImport *imports.Imports = nil
)

func actionLinkAPI() *imports.Imports {
	actionLinkOnce.Do(func() {
		actionLinkImport = api.NewDefaultImports()
		actionLinkImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TActionLink_Create", 0), // constructor NewActionLink
			/* 1 */ imports.NewTable("TActionLink_IsCaptionLinked", 0), // function IsCaptionLinked
			/* 2 */ imports.NewTable("TActionLink_IsCheckedLinked", 0), // function IsCheckedLinked
			/* 3 */ imports.NewTable("TActionLink_IsEnabledLinked", 0), // function IsEnabledLinked
			/* 4 */ imports.NewTable("TActionLink_IsGroupIndexLinked", 0), // function IsGroupIndexLinked
			/* 5 */ imports.NewTable("TActionLink_IsHelpContextLinked", 0), // function IsHelpContextLinked
			/* 6 */ imports.NewTable("TActionLink_IsHelpLinked", 0), // function IsHelpLinked
			/* 7 */ imports.NewTable("TActionLink_IsHintLinked", 0), // function IsHintLinked
			/* 8 */ imports.NewTable("TActionLink_IsImageIndexLinked", 0), // function IsImageIndexLinked
			/* 9 */ imports.NewTable("TActionLink_IsShortCutLinked", 0), // function IsShortCutLinked
			/* 10 */ imports.NewTable("TActionLink_IsVisibleLinked", 0), // function IsVisibleLinked
		}
	})
	return actionLinkImport
}
