package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

func dataCtx() *ui.BuildContext { return &ui.BuildContext{Theme: theme.Default()} }

func sampleRows() [][]any {
	rows := make([][]any, 100)
	for i := range rows {
		rows[i] = []any{i, []string{"张三", "李四", "王五"}[i%3], float64(i % 7)}
	}
	return rows
}

func tableCols() []TableColumn {
	return []TableColumn{
		{Title: "ID", Mode: ColumnFixed, Width: 60},
		{Title: "姓名", Mode: ColumnAuto},
		{Title: "值", Mode: ColumnFlex, Flex: 1, Sortable: true},
	}
}

func TestTableColumnWidthsQ9(t *testing.T) {
	tbl := NewTable(tableCols(), sampleRows())
	tbl.Width = 600
	w := tbl.columnWidths(theme.Default().Font.SizeSM)
	if len(w) != 3 {
		t.Fatalf("列宽数量 = %d", len(w))
	}
	if w[0] != 60 {
		t.Fatalf("fixed 列宽 = %v，应为 60", w[0])
	}
	if w[1] <= 12 || w[1] > 600 {
		t.Fatalf("auto 列宽 = %v，应包含文本宽+padding", w[1])
	}
	sum := w[0] + w[1] + w[2]
	if sum > 600+1 && w[2] > 0 {
		t.Fatalf("列宽总和超出: %v", sum)
	}
	if w[2] <= 0 {
		t.Fatalf("flex 列宽 = %v，应分到剩余空间", w[2])
	}
}

func TestTableColumnFlexWeight(t *testing.T) {
	cols := []TableColumn{
		{Title: "A", Mode: ColumnFlex, Flex: 1},
		{Title: "B", Mode: ColumnFlex, Flex: 3},
	}
	tbl := NewTable(cols, nil)
	tbl.Width = 400
	w := tbl.columnWidths(theme.Default().Font.SizeSM)
	if w[0] <= 0 || w[1] <= 0 || w[1] < w[0]*2 {
		t.Fatalf("flex 权重分配错误: %v", w)
	}
}

func TestTableSortStable(t *testing.T) {
	tbl := NewTable(tableCols(), sampleRows())
	tbl.Sort(2, true)
	idx := tbl.sortedIndex()
	if len(idx) != len(tbl.Rows) {
		t.Fatal("排序后行数变了")
	}
	// 索引必须是对全部行的一个排列
	seen := make(map[int]bool)
	for _, i := range idx {
		if seen[i] {
			t.Fatalf("排序索引重复: %d", i)
		}
		seen[i] = true
	}
	// 降序翻转
	tbl.Sort(2, false)
	asc, desc := tbl.sortedIndex(), idx
	if cellString(tbl.Rows[asc[0]][2]) < cellString(tbl.Rows[desc[0]][2]) {
		t.Fatal("降序首行应不小于升序首行")
	}
}

func TestTableVirtualization(t *testing.T) {
	tbl := NewTable(tableCols(), sampleRows())
	tbl.Width, tbl.Height, tbl.RowHeight = 600, 120, 30
	n := tbl.Build(dataCtx())
	// 1 表头 + 可见行（远小于 100）
	if got := len(n.Children()); got >= len(tbl.Rows) {
		t.Fatalf("虚拟化失效：构建了 %d 个节点（总行 %d）", got, len(tbl.Rows))
	}
	visible := tbl.VisibleRows()
	if len(visible) == 0 || len(visible) >= len(tbl.Rows) {
		t.Fatalf("可见行数量异常: %d", len(visible))
	}
	// 滚到底部
	tbl.SetScrollOffset(tbl.maxOffset())
	n2 := tbl.Build(dataCtx())
	if len(n2.Children()) >= len(tbl.Rows) {
		t.Fatal("滚动后仍构建了全部行")
	}
	last := tbl.VisibleRows()
	if last[len(last)-1] != len(tbl.Rows)-1 {
		t.Fatalf("滚到底部应包含末行，实际末可见 = %d", last[len(last)-1])
	}
}

func TestTableScrollClamp(t *testing.T) {
	tbl := NewTable(tableCols(), sampleRows())
	tbl.Height, tbl.RowHeight = 120, 30
	tbl.SetScrollOffset(-100)
	if tbl.ScrollOffset() != 0 {
		t.Fatalf("负偏移未钳制: %v", tbl.ScrollOffset())
	}
	tbl.SetScrollOffset(1e9)
	if tbl.ScrollOffset() > tbl.maxOffset() {
		t.Fatalf("越界偏移未钳制: %v", tbl.ScrollOffset())
	}
}

func TestTreeExpand(t *testing.T) {
	nodes := []TreeNode{
		{Key: "a", Label: "A", Children: []TreeNode{{Key: "a1", Label: "A1"}}},
		{Key: "b", Label: "B"},
	}
	tr := NewTree(nodes)
	tr.OnToggle = func(string, bool) {}
	n := tr.Build(dataCtx())
	if got := len(n.Children()); got != 2 {
		t.Fatalf("收起态应只有 2 个顶层节点，实际 %d", got)
	}
	tr.Toggle("a")
	n2 := tr.Build(dataCtx())
	if got := len(n2.Children()); got != 3 {
		t.Fatalf("展开 A 后应有 3 行，实际 %d", got)
	}
}

func TestFormValidation(t *testing.T) {
	f := NewForm()
	f.Rule("姓名", RequiredRule("请输入姓名"))
	f.Rule("城市", RequiredRule("请输入城市"))
	if f.Validate() {
		t.Fatal("空表单应校验失败")
	}
	if f.Errors["姓名"] != "请输入姓名" {
		t.Fatalf("错误提示 = %q", f.Errors["姓名"])
	}
	f.Values["姓名"] = "张三"
	f.Values["城市"] = "深圳"
	if !f.Validate() {
		t.Fatalf("填写完整应通过: %v", f.Errors)
	}
}

func TestFormRuleSet(t *testing.T) {
	// 声明式规则集：required/min/email/url/pattern/validator 全路径。
	f := NewForm()
	f.Rules("用户名", FormRule{Required: true, Min: 2}, FormRule{Whitespace: true, Message: "此项必填"})
	f.Rules("邮箱", FormRule{Required: true, Email: true})
	f.Rules("网址", FormRule{URL: true})
	f.Rules("验证码", FormRule{Pattern: `^\d{4}$`, Message: "验证码为 4 位数字"})
	f.Rules("密码", FormRule{Validator: func(v string) string {
		if len([]rune(v)) < 6 {
			return "密码至少 6 位"
		}
		return ""
	}})
	if f.Validate() {
		t.Fatal("空表单应校验失败")
	}
	if f.Errors["用户名"] != "此项必填" || f.Errors["邮箱"] != "此项必填" {
		t.Fatalf("required 报错 = %v", f.Errors)
	}
	f.Values["用户名"] = "w"  // min 2 触发
	f.Values["邮箱"] = "bad" // email 触发
	f.Values["网址"] = "notaurl"
	f.Values["验证码"] = "12ab" // pattern 触发
	f.Values["密码"] = "123"   // validator 触发
	if f.Validate() {
		t.Fatal("非法值应校验失败")
	}
	if f.Errors["用户名"] != "长度不能少于 2" {
		t.Fatalf("min 报错 = %q", f.Errors["用户名"])
	}
	if f.Errors["邮箱"] != "邮箱格式不正确" {
		t.Fatalf("email 报错 = %q", f.Errors["邮箱"])
	}
	if f.Errors["网址"] != "网址格式不正确" {
		t.Fatalf("url 报错 = %q", f.Errors["网址"])
	}
	if f.Errors["验证码"] != "验证码为 4 位数字" {
		t.Fatalf("pattern 报错 = %q", f.Errors["验证码"])
	}
	if f.Errors["密码"] != "密码至少 6 位" {
		t.Fatalf("validator 报错 = %q", f.Errors["密码"])
	}
	// 合法值全通过（含 whitespace/len 边界）。
	f.Values["用户名"] = "wg用户"
	f.Values["邮箱"] = "name@example.com"
	f.Values["网址"] = "https://example.com"
	f.Values["验证码"] = "1234"
	f.Values["密码"] = "secret"
	if !f.Validate() {
		t.Fatalf("合法值应通过: %v", f.Errors)
	}
	// whitespace：显式开启的规则里纯空白算空（antd whitespace opt-in）。
	f.Values["用户名"] = "   "
	if f.Validate() {
		t.Fatal("纯空白应校验失败")
	}
	if f.Errors["用户名"] != "此项必填" {
		t.Fatalf("whitespace 报错 = %q", f.Errors["用户名"])
	}
	// 声明式规则与过程式校验器并存：首错即停。
	f2 := NewForm()
	f2.Rules("字段", FormRule{Required: true, Message: "规则报错"})
	f2.Rule("字段", func(v string) string { return "校验器报错" })
	f2.Validate()
	if f2.Errors["字段"] != "规则报错" {
		t.Fatalf("首错即停 = %q", f2.Errors["字段"])
	}
}

func TestSelectOptionsAndValue(t *testing.T) {
	s := NewSelect([]string{"北京", "上海", "广州"})
	if s.Selected != -1 {
		t.Fatal("初始未选")
	}
	changed := -1
	s.OnChange = func(i int) { changed = i }
	s.SetSelected(1)
	if s.Selected != 1 || changed != 1 {
		t.Fatalf("选中失败: %d / %d", s.Selected, changed)
	}
	s.SetSelected(1)
	if changed != 1 {
		t.Fatal("重复选同一项不应再回调")
	}
	// 无 ctx.Overlay 时 show 是安全 no-op
	data := s.Build(dataCtx())
	if data == nil {
		t.Fatal("Select 构建失败")
	}
}
