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

// ICustomShape Parent: IGraphicControl
type ICustomShape interface {
	IGraphicControl
	PtInShape(P types.TPoint) bool         // function
	Paint()                                // procedure
	StyleChanged(sender IObject)           // procedure
	Brush() IBrush                         // property Brush Getter
	SetBrush(value IBrush)                 // property Brush Setter
	Pen() IPen                             // property Pen Getter
	SetPen(value IPen)                     // property Pen Setter
	Shape() types.TShapeType               // property Shape Getter
	SetShape(value types.TShapeType)       // property Shape Setter
	SetOnShapeClick(fn TNotifyEvent)       // property event
	SetOnShapePoints(fn TShapePointsEvent) // property event
}

type TCustomShape struct {
	TGraphicControl
}

func (m *TCustomShape) PtInShape(P types.TPoint) bool {
	if !m.IsValid() {
		return false
	}
	r := customShapeAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return api.GoBool(r)
}

func (m *TCustomShape) Paint() {
	if !m.IsValid() {
		return
	}
	customShapeAPI().SysCallN(2, m.Instance())
}

func (m *TCustomShape) StyleChanged(sender IObject) {
	if !m.IsValid() {
		return
	}
	customShapeAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TCustomShape) Brush() IBrush {
	if !m.IsValid() {
		return nil
	}
	r := customShapeAPI().SysCallN(4, 0, m.Instance())
	return AsBrush(r)
}

func (m *TCustomShape) SetBrush(value IBrush) {
	if !m.IsValid() {
		return
	}
	customShapeAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomShape) Pen() IPen {
	if !m.IsValid() {
		return nil
	}
	r := customShapeAPI().SysCallN(5, 0, m.Instance())
	return AsPen(r)
}

func (m *TCustomShape) SetPen(value IPen) {
	if !m.IsValid() {
		return
	}
	customShapeAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomShape) Shape() types.TShapeType {
	if !m.IsValid() {
		return 0
	}
	r := customShapeAPI().SysCallN(6, 0, m.Instance())
	return types.TShapeType(r)
}

func (m *TCustomShape) SetShape(value types.TShapeType) {
	if !m.IsValid() {
		return
	}
	customShapeAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomShape) SetOnShapeClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, customShapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomShape) SetOnShapePoints(fn TShapePointsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTShapePointsEvent(fn)
	base.SetEvent(m, 8, customShapeAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomShape class constructor
func NewCustomShape(theOwner IComponent) ICustomShape {
	r := customShapeAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomShape(r)
}

func TCustomShapeClass() types.TClass {
	r := customShapeAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	customShapeOnce   base.Once
	customShapeImport *imports.Imports = nil
)

func customShapeAPI() *imports.Imports {
	customShapeOnce.Do(func() {
		customShapeImport = api.NewDefaultImports()
		customShapeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomShape_Create", 0), // constructor NewCustomShape
			/* 1 */ imports.NewTable("TCustomShape_PtInShape", 0), // function PtInShape
			/* 2 */ imports.NewTable("TCustomShape_Paint", 0), // procedure Paint
			/* 3 */ imports.NewTable("TCustomShape_StyleChanged", 0), // procedure StyleChanged
			/* 4 */ imports.NewTable("TCustomShape_Brush", 0), // property Brush
			/* 5 */ imports.NewTable("TCustomShape_Pen", 0), // property Pen
			/* 6 */ imports.NewTable("TCustomShape_Shape", 0), // property Shape
			/* 7 */ imports.NewTable("TCustomShape_OnShapeClick", 0), // event OnShapeClick
			/* 8 */ imports.NewTable("TCustomShape_OnShapePoints", 0), // event OnShapePoints
			/* 9 */ imports.NewTable("TCustomShape_TClass", 0), // function TCustomShapeClass
		}
	})
	return customShapeImport
}
