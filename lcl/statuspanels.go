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
	. "github.com/energye/lcl/types"
)

// IStatusPanels Parent: ICollection
type IStatusPanels interface {
	ICollection
	ItemsForStatusPanel(Index int32) IStatusPanel            // property
	SetItemsForStatusPanel(Index int32, AValue IStatusPanel) // property
	StatusBar() IStatusBar                                   // property
	AddForStatusPanel() IStatusPanel                         // function
}

// TStatusPanels Parent: TCollection
type TStatusPanels struct {
	TCollection
}

func NewStatusPanels(AStatusBar IStatusBar) IStatusPanels {
	r1 := LCL().SysCallN(5197, GetObjectUintptr(AStatusBar))
	return AsStatusPanels(r1)
}

func (m *TStatusPanels) ItemsForStatusPanel(Index int32) IStatusPanel {
	r1 := LCL().SysCallN(5198, 0, m.Instance(), uintptr(Index))
	return AsStatusPanel(r1)
}

func (m *TStatusPanels) SetItemsForStatusPanel(Index int32, AValue IStatusPanel) {
	LCL().SysCallN(5198, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TStatusPanels) StatusBar() IStatusBar {
	r1 := LCL().SysCallN(5199, m.Instance())
	return AsStatusBar(r1)
}

func (m *TStatusPanels) AddForStatusPanel() IStatusPanel {
	r1 := LCL().SysCallN(5195, m.Instance())
	return AsStatusPanel(r1)
}

func StatusPanelsClass() TClass {
	ret := LCL().SysCallN(5196)
	return TClass(ret)
}
