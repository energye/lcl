package widgets

import "testing"

// TestFormListAddRemove 动态行：初始 1 行、Add/Remove、字段键命名。
func TestFormListAddRemove(t *testing.T) {
	l := NewFormList("contacts", "name", "phone")
	if l.Len() != 1 {
		t.Fatalf("初始行数=%d，want 1", l.Len())
	}

	// Add → 2 行。
	l.Add()
	if l.Len() != 2 {
		t.Fatalf("Add 后行数=%d，want 2", l.Len())
	}
	l.Add()
	if l.Len() != 3 {
		t.Fatalf("再 Add 后行数=%d，want 3", l.Len())
	}

	// Remove 中间行 → 2 行。
	l.Remove(1)
	if l.Len() != 2 {
		t.Fatalf("Remove(1) 后行数=%d，want 2", l.Len())
	}
	// Remove 越界/负索引无副作用。
	l.Remove(99)
	l.Remove(-1)
	if l.Len() != 2 {
		t.Fatalf("越界 Remove 后行数=%d，want 2", l.Len())
	}

	// 字段键命名 contacts[i].field。
	if got := l.fieldKey(0, "name"); got != "contacts[0].name" {
		t.Fatalf("fieldKey=%q，want contacts[0].name", got)
	}
	if got := l.fieldKey(2, "phone"); got != "contacts[2].phone" {
		t.Fatalf("fieldKey=%q，want contacts[2].phone", got)
	}

	// Values 汇总所有行字段。
	l.Rows[0].Values["name"] = "张三"
	l.Rows[0].Values["phone"] = "123"
	l.Rows[1].Values["name"] = "李四"
	vals := l.Values()
	if vals["contacts[0].name"] != "张三" || vals["contacts[0].phone"] != "123" {
		t.Fatalf("行 0 值=%v", vals)
	}
	if vals["contacts[1].name"] != "李四" {
		t.Fatalf("行 1 值=%v", vals)
	}

	// 构建不 panic + 锚存在。
	if n := l.Build(dataCtx()); n == nil {
		t.Fatal("FormList 构建失败")
	}
}
