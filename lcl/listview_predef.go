//----------------------------------------
// The code is automatically generated by the generate-tool.
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

func (m *TListView) DeleteSelected() {
	selected := m.Selected()
	if selected != nil && selected.IsValid() {
		selected.Delete()
	}
}
