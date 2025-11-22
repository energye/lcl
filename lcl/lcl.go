//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import (
	"encoding/json"
	"math"

	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
)

func MathRound(AValue float64) int64 {
	if AValue >= 0 {
		return int64(math.Trunc(AValue + 0.5))
	} else {
		return int64(math.Trunc(AValue - 0.5))
	}
}

func MulDiv(nNumber, nNumerator, nDenominator int32) int32 {
	if nDenominator == 0 {
		return -1
	} else if nNumerator == nDenominator {
		return nNumber
	} else {
		return int32(MathRound(float64(nNumber) * float64(nNumerator) / float64(nDenominator)))
	}
}

func EqualsObject(obj1, obj2 IObject) bool {
	if obj1 == nil && obj2 == nil {
		return true
	} else if obj1 == nil || obj2 == nil {
		return false
	}
	return obj1.Instance() == obj2.Instance()
}

func FindControl(handle types.HWND) IWinControl {
	return AsWinControl(api.FindControl(handle))
}

func FindLCLControl(screenPos types.TPoint) IControl {
	return AsControl(api.FindLCLControl(screenPos))
}

func FindOwnerControl(handle types.HWND) IWinControl {
	return AsWinControl(api.FindOwnerControl(handle))
}

func FindControlAtPosition(position types.TPoint, allowDisabled bool) IControl {
	return AsControl(api.FindControlAtPosition(position, allowDisabled))
}

func FindLCLWindow(screenPos types.TPoint, allowDisabled bool) IWinControl {
	return AsWinControl(api.FindLCLWindow(screenPos, allowDisabled))
}

func FindDragTarget(position types.TPoint, allowDisabled bool) IControl {
	return AsControl(api.FindDragTarget(position, allowDisabled))
}

type ComponentProperties struct {
	Name    string `json:"name"`    // 属性名
	Value   string `json:"value"`   // 属性值, 或默认值
	Kind    string `json:"kind"`    // 属性类型
	Type    string `json:"type"`    // 属性类型名
	Options string `json:"options"` // 属性复合类型(枚举、集合)选项
}

func (m *ComponentProperties) ToJSON() string {
	str, _ := json.Marshal(m)
	return string(str)
}

var gSetDesigningComponent *_TSetDesigningComponent

type _TSetDesigningComponent struct {
	designerFormEventData uintptr
}

// 设计组件实例 在设计阶段时使用
func DesigningComponent() *_TSetDesigningComponent {
	if gSetDesigningComponent == nil {
		gSetDesigningComponent = new(_TSetDesigningComponent)
	}
	return gSetDesigningComponent
}

func (m *_TSetDesigningComponent) SetComponentDesignMode(component IComponent, value bool) {
	if component == nil || !component.IsValid() {
		return
	}
	api.SetComponentDesignMode(component.Instance(), value)
}

func (m *_TSetDesigningComponent) SetComponentDesignInstanceMode(component IComponent, value bool) {
	if component == nil || !component.IsValid() {
		return
	}
	api.SetComponentDesignInstanceMode(component.Instance(), value)
}

func (m *_TSetDesigningComponent) SetComponentInlineMode(component IComponent, value bool) {
	if component == nil || !component.IsValid() {
		return
	}
	api.SetComponentInlineMode(component.Instance(), value)
}

func (m *_TSetDesigningComponent) SetComponentAncestorMode(component IComponent, value bool) {
	if component == nil || !component.IsValid() {
		return
	}
	api.SetComponentAncestorMode(component.Instance(), value)
}

func (m *_TSetDesigningComponent) SetWidgetSetDesigning(component IComponent) {
	if component == nil || !component.IsValid() {
		return
	}
	api.SetWidgetSetDesigning(component.Instance())
}

// GetComponentProperties 获取组件属性
func (m *_TSetDesigningComponent) GetComponentProperties(component IObject) []ComponentProperties {
	if component == nil || !component.IsValid() {
		return nil
	}
	propList := NewStringList()
	defer propList.Free()
	api.GetComponentProperties(component.Instance(), propList.Instance())
	count := int(propList.Count())
	if count > 0 {
		props := make([]ComponentProperties, count)
		for i := 0; i < count; i++ {
			jsonStr := propList.Strings(int32(i))
			var propItem ComponentProperties
			if err := json.Unmarshal([]byte(jsonStr), &propItem); err == nil {
				props[i] = propItem
			}
		}
		return props
	}
	return nil
}

type TGetDesignerFormEvent func(persistent IPersistent) ICustomForm

func (m *_TSetDesigningComponent) SetOnGetDesignerFormEvent(fn TGetDesignerFormEvent) {
	makeTGetDesignerFormEvent := func(cb TGetDesignerFormEvent) *callback {
		if cb == nil {
			return nil
		}
		return &callback{
			name: "TGetDesignerFormEvent",
			cb: func(getVal func(i int) uintptr) {
				persistent := AsPersistent(getVal(0))
				customForm := (*uintptr)(getPtr(getVal(1)))
				result := cb(persistent)
				if result != nil {
					*customForm = result.Instance()
				}
			},
		}
	}
	cb := makeTGetDesignerFormEvent(fn)
	if cb == nil && m.designerFormEventData != 0 {
		api.RemoveEventElement(m.designerFormEventData)
		api.SetOnGetDesignerFormEvent(0)
		m.designerFormEventData = 0
	} else {
		m.designerFormEventData = api.MakeEventDataPtr(cb)
		api.SetOnGetDesignerFormEvent(m.designerFormEventData)
	}
}
