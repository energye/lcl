package compat

import (
	"testing"

	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func ctx() *ui.BuildContext { return &ui.BuildContext{Theme: theme.Default()} }

// 门面必须转发到新组件：结构与新式构造器完全一致。
func TestFacadesForward(t *testing.T) {
	// 按钮
	if TButton("确定").Text != "确定" {
		t.Fatal("TButton 未转发 Text")
	}
	// 布局
	row := TRow(TSpacer(10, 10), TSpacer(20, 20))
	if len(row.Children) != 2 {
		t.Fatal("TRow 未转发子节点")
	}
	// 列表
	l := TList(widgets.ListItem{Key: "a", Widget: TLabel("A")})
	if len(l.Items) != 1 {
		t.Fatal("TList 未转发 items")
	}
	// 表格
	tbl := TTable(
		[]widgets.TableColumn{{Title: "A", Mode: widgets.ColumnAuto}},
		[][]any{{"x"}, {"y"}},
	)
	if len(tbl.Rows) != 2 {
		t.Fatal("TTable 未转发 rows")
	}
	// 浮层
	if TModal("t", nil).Title != "t" {
		t.Fatal("TModal 未转发 title")
	}
	if TDrawer("d", nil).Title != "d" {
		t.Fatal("TDrawer 未转发 title")
	}
	if TSelect([]string{"a", "b"}).Options[1] != "b" {
		t.Fatal("TSelect 未转发 options")
	}
	if !TCheckbox("x", true).Checked || !TSwitch(true).On {
		t.Fatal("旧式布尔门面未转发")
	}
	if TProgress(0.5).Value != 0.5 {
		t.Fatal("TProgress 未转发 value")
	}
}

// 门面构建出的节点必须能进入新引擎出帧（不是空壳）。
func TestFacadeNodesBuild(t *testing.T) {
	c := ctx()
	nodes := []ui.Widget{
		TButton("B"), TInput("P"), TLabel("L"), TCard(TLabel("C")),
		TRow(TSpacer(1, 1)), TColumn(TSpacer(1, 1)),
		TPadding(4, 4, TLabel("X")), TSelect([]string{"a"}),
		TSegmented("1", "2"), TPagination(1, 3),
		TCollapse(widgets.CollapsePanel{Title: "p"}),
	}
	for i, w := range nodes {
		if n := AsNode(c, w); n == nil {
			t.Fatalf("门面[%d] 构建出空节点", i)
		}
	}
}
