package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/input"
)

// newNavMenu 造三层菜单：根 a/b(父)/c，b 下 b-1/b-2(父)/b-3，b-2 下 b-2-1。
func newNavMenu() *Menu {
	return NewMenu(
		&MenuItem{Key: "a", Label: "甲"},
		&MenuItem{Key: "b", Label: "乙", Children: []*MenuItem{
			{Key: "b-1", Label: "乙一"},
			{Key: "b-2", Label: "乙二", Children: []*MenuItem{
				{Key: "b-2-1", Label: "乙二一"},
			}},
			{Key: "b-3", Label: "乙三"},
		}},
		&MenuItem{Key: "c", Label: "丙"},
	)
}

func keyDown(code input.KeyCode) *input.KeyEvent {
	return &input.KeyEvent{Kind: input.KeyDown, Code: code}
}

// TestMenuNavItemsVisibility navItems 只收可见项：收起时子项不出现、
// 展开后出现、分组标题跳过。
func TestMenuNavItemsVisibility(t *testing.T) {
	m := NewMenu(
		&MenuItem{Key: "grp", Label: "分组", Group: true, Children: []*MenuItem{
			{Key: "g1", Label: "组员"},
		}},
		&MenuItem{Key: "b", Label: "乙", Children: []*MenuItem{
			{Key: "b-1", Label: "乙一"},
		}},
		&MenuItem{Key: "c", Label: "丙"},
	)
	nav := m.navItems()
	if len(nav) != 2 || nav[0].key != "b" || nav[1].key != "c" {
		t.Fatalf("收起态应只有 b/c（分组跳过），got %v", nav)
	}
	m.OpenKeys = map[string]bool{"b": true}
	nav = m.navItems()
	if len(nav) != 3 || nav[1].key != "b-1" {
		t.Fatalf("展开后应含 b-1，got %v", nav)
	}
}

// TestMenuNavUpDown ↓↑ 在可见项间移动焦点（含展开态越界钳制）。
func TestMenuNavUpDown(t *testing.T) {
	m := newNavMenu()
	m.OpenKeys = map[string]bool{"b": true}
	// 可见序：a, b, b-1, b-2, b-3, c
	m.handleNavKey(keyDown(input.CodeDown), nil)
	if m.FocusedKey != "a" {
		t.Fatalf("无焦点时 ↓ 应落首项，got %q", m.FocusedKey)
	}
	m.handleNavKey(keyDown(input.CodeDown), nil)
	if m.FocusedKey != "b" {
		t.Fatalf("↓ 应到 b，got %q", m.FocusedKey)
	}
	m.handleNavKey(keyDown(input.CodeUp), nil)
	if m.FocusedKey != "a" {
		t.Fatalf("↑ 应回 a，got %q", m.FocusedKey)
	}
	m.handleNavKey(keyDown(input.CodeUp), nil)
	if m.FocusedKey != "a" {
		t.Fatalf("首项 ↑ 应钳制不动，got %q", m.FocusedKey)
	}
}

// TestMenuNavRightLeft → 展开父项并进首子项、← 收起焦点回父项。
func TestMenuNavRightLeft(t *testing.T) {
	m := newNavMenu()
	// → 在未聚焦时返回 false（不消费）。
	if m.handleNavKey(keyDown(input.CodeRight), nil) {
		t.Fatal("未聚焦时 → 不该消费")
	}
	m.FocusedKey = "b"
	if !m.handleNavKey(keyDown(input.CodeRight), nil) {
		t.Fatal("→ 应消费")
	}
	if !m.OpenKeys["b"] {
		t.Fatal("→ 应展开 b")
	}
	if m.FocusedKey != "b" {
		t.Fatalf("→ 展开后焦点留在父项，got %q", m.FocusedKey)
	}
	// 再 → 进首子项 b-1。
	m.handleNavKey(keyDown(input.CodeRight), nil)
	if m.FocusedKey != "b-1" {
		t.Fatalf("→ 应进首子项 b-1，got %q", m.FocusedKey)
	}
	// ← 收起 b，焦点回父项。
	if !m.handleNavKey(keyDown(input.CodeLeft), nil) {
		t.Fatal("← 应消费")
	}
	if m.OpenKeys["b"] {
		t.Fatal("← 应收起 b")
	}
	if m.FocusedKey != "b" {
		t.Fatalf("← 后焦点应回父项 b，got %q", m.FocusedKey)
	}
}

// TestMenuNavEnter Enter 选中叶子（父项=切换展开，不触发 OnClick）。
func TestMenuNavEnter(t *testing.T) {
	m := newNavMenu()
	var clicked []string
	m.OnClick = func(k string) { clicked = append(clicked, k) }
	m.FocusedKey = "c"
	if !m.handleNavKey(keyDown(input.CodeEnter), nil) {
		t.Fatal("Enter 应消费")
	}
	if m.Selected != "c" || len(clicked) != 1 || clicked[0] != "c" {
		t.Fatalf("Enter 应选中 c 并回调，got selected=%q clicked=%v", m.Selected, clicked)
	}
	// 父项 Enter=切换展开。
	m.FocusedKey = "b"
	m.handleNavKey(keyDown(input.CodeEnter), nil)
	if !m.OpenKeys["b"] || len(clicked) != 1 {
		t.Fatalf("父项 Enter 应只切展开，OpenKeys[b]=%v clicked=%v", m.OpenKeys["b"], clicked)
	}
	m.handleNavKey(keyDown(input.CodeEnter), nil)
	if m.OpenKeys["b"] {
		t.Fatal("父项再次 Enter 应收起")
	}
}
