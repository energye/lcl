package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// ── 纯展示组件的构建 ──────────────────────────────────────────────

func TestCardComposition(t *testing.T) {
	c := NewCard(Spacer{Width: 100, Height: 20})
	c.Title = Label{Text: "标题"}
	c.Footer = Label{Text: "底部"}
	c.Elevated = false
	n := c.Build(testCtx())
	n.Layout(scene.Unbounded())
	if len(n.Children()) != 1 {
		t.Fatalf("Card 应包含一个纵向 Flex，实际 %d 个子节点", len(n.Children()))
	}
	col := n.Children()[0]
	// 标题 + 分隔线 + 内容 + 底部 = 4
	if got := len(col.Children()); got != 4 {
		t.Fatalf("Card 主体子节点 = %d，应为 4", got)
	}
	if n.Size().W < 100 || n.Size().H < 20 {
		t.Fatalf("Card 固有尺寸未覆盖内容: %+v", n.Size())
	}
}

func TestBadgeTagAvatar(t *testing.T) {
	dot := (&Badge{Dot: true}).Build(testCtx())
	dot.Layout(scene.Unbounded())
	if dot.Size().W != 8 || dot.Size().H != 8 {
		t.Fatalf("圆点角标尺寸 = %+v", dot.Size())
	}
	num := (&Badge{Count: 42}).Build(testCtx())
	num.Layout(scene.Unbounded())
	if num.Size().W < 16 || num.Size().H < 16 {
		t.Fatalf("数字角标尺寸异常: %+v", num.Size())
	}

	tag := (&Tag{Text: "标签"}).Build(testCtx())
	tag.Layout(scene.Unbounded())
	if len(tag.Children()) != 1 || tag.Size().W <= 0 {
		t.Fatalf("Tag 结构/尺寸异常: children=%d size=%+v", len(tag.Children()), tag.Size())
	}

	av := (&Avatar{Text: "中文名字", Size: 40}).Build(testCtx())
	av.Layout(scene.Unbounded())
	if av.Size().W != 40 || av.Size().H != 40 || len(av.Children()) != 1 {
		t.Fatalf("Avatar 异常: size=%+v children=%d", av.Size(), len(av.Children()))
	}
}

func TestEmptySkeletonSpin(t *testing.T) {
	e := (&Empty{Text: "空空如也", Subtext: "试试新建", Width: 260}).Build(testCtx())
	e.Layout(scene.Unbounded())
	if e.Size().W != 260 || e.Size().H != 96 {
		t.Fatalf("Empty 尺寸 = %+v", e.Size())
	}
	if got := len(e.Children()); got != 3 {
		t.Fatalf("Empty 子节点 = %d（图标+主文案+副文案应为 3）", got)
	}

	s := (&Skeleton{Rows: 4, Width: 200, Height: 10}).Build(testCtx())
	s.Layout(scene.Unbounded())
	if got := len(s.Children()); got != 4 {
		t.Fatalf("Skeleton 行数 = %d，应为 4", got)
	}

	sp := (&Spin{Size: 24, Phase: 0.25}).Build(testCtx())
	sp.Layout(scene.Unbounded())
	if sp.Size().W != 24 || sp.Size().H != 24 {
		t.Fatalf("Spin 尺寸 = %+v", sp.Size())
	}
}

func TestSteps(t *testing.T) {
	n := (&Steps{Labels: []string{"第一步", "第二步", "第三步"}, Current: 1}).Build(testCtx())
	n.Layout(scene.Unbounded())
	// 新结构：根节点 = 步骤列与连接线交替（3 步 → 5 个子节点）。
	steps, connectors := 0, 0
	for _, ch := range n.Children() {
		if len(ch.Children()) >= 2 {
			steps++
		} else {
			connectors++
		}
	}
	if steps != 3 || connectors != 2 {
		t.Fatalf("Steps 步数 = %d（应 3），连接线 = %d（应 2）", steps, connectors)
	}
	if got := len(n.Children()[0].Children()); got != 2 {
		t.Fatalf("每一步应有圆形序号+标签两个子节点，实际 %d", got)
	}
}

// TestStepsIcons 自定义图标：32×32 容器+Icon 子节点（icon 即内容不叠
// 数字/对勾），未设图标步维持 20×20 默认圆点。
func TestStepsIcons(t *testing.T) {
	s := &Steps{Labels: []string{"账户", "设置", "支付", "完成"}, Current: 2,
		Icons: []string{"user", "setting", "shopping-cart", "check-circle"}}
	n := s.Build(testCtx())
	n.Layout(scene.Unbounded())
	for i, ch := range n.Children() {
		if i%2 == 1 {
			continue // 连接线
		}
		// 步节点首个子节点是图标容器。
		iconNode := ch.Children()[0]
		sz := iconNode.Size()
		if sz.W != 32 || sz.H != 32 {
			t.Fatalf("icon 步 %d 容器应 32×32，got %.0fx%.0f", i/2, sz.W, sz.H)
		}
		if len(iconNode.Children()) != 1 {
			t.Fatalf("icon 步 %d 容器应只含 1 个 Icon 子节点，got %d", i/2, len(iconNode.Children()))
		}
	}
	if s.iconOf(1) != "setting" || s.iconOf(4) != "" {
		t.Fatalf("iconOf 取值错误: %q %q", s.iconOf(1), s.iconOf(4))
	}
	// 未设图标：容器回到 20×20 默认圆点。
	d := (&Steps{Labels: []string{"a", "b"}, Current: 0}).Build(testCtx())
	d.Layout(scene.Unbounded())
	sz := d.Children()[0].Children()[0].Size()
	if sz.W != 20 || sz.H != 20 {
		t.Fatalf("默认圆点应 20×20，got %.0fx%.0f", sz.W, sz.H)
	}
}

// ── 交互组件的点击 ────────────────────────────────────────────────

func TestSegmentedClick(t *testing.T) {
	r := newControlRig()
	seg := NewSegmented("日", "周", "月")
	changed := -1
	seg.OnChange = func(i int) { changed = i }
	r.mount(seg)
	r.click(15, 12)
	if seg.Selected != 0 || changed != -1 {
		t.Fatalf("点中已选段不该回调: selected=%d changed=%d", seg.Selected, changed)
	}
	// 第二段起点在首段右侧（root 是 Box，其 child[0] 才是 Flex 行）
	second := seg.Node().Children()[0].Children()[1].Bounds()
	r.click((second.Min.X+second.Max.X)/2, (second.Min.Y+second.Max.Y)/2)
	if seg.Selected != 1 || changed != 1 {
		t.Fatalf("第二段点击失败: selected=%d changed=%d", seg.Selected, changed)
	}
}

func TestPaginationPagesAndClick(t *testing.T) {
	p := NewPagination(1, 20)
	// ellipsis 语义：首末恒显 + 窗口 ±2 + 间隙省略号（pageCell）。
	got := p.visiblePages()
	if len(got) != 5 || got[0].page != 1 || got[0].ellipsis != 0 ||
		got[3].ellipsis != 5 || got[4].page != 20 {
		t.Fatalf("首页可见序列 = %+v", got)
	}
	p.Page = 10
	got = p.visiblePages()
	if len(got) != 9 || got[1].ellipsis != -5 || got[7].ellipsis != 5 || got[8].page != 20 {
		t.Fatalf("中间页可见序列 = %+v", got)
	}
	p.Page = 20
	got = p.visiblePages()
	if got[len(got)-1].page != 20 || got[1].ellipsis != -5 {
		t.Fatalf("末页可见序列 = %+v", got)
	}
	p.Page = 1

	r := newControlRig()
	changed := -1
	p.OnChange = func(i int) { changed = i }
	r.mount(p)
	// 第二页按钮中心
	second := p.Node().Children()[1].Bounds()
	r.click((second.Min.X+second.Max.X)/2, (second.Min.Y+second.Max.Y)/2)
	if p.Page != 2 || changed != 2 {
		t.Fatalf("翻页失败: page=%d changed=%d", p.Page, changed)
	}
}

func TestBreadcrumbClick(t *testing.T) {
	r := newControlRig()
	bc := &Breadcrumb{Items: []string{"首页", "列表", "详情"}}
	clicked := -1
	bc.OnClick = func(i int) { clicked = i }
	r.mount(bc)
	if len(r.root.Children()) == 0 {
		t.Fatal("Breadcrumb 未挂载")
	}
	// 第一项在 x≈0；最后一项高亮不可点
	r.click(8, 10)
	if clicked != 0 {
		t.Fatalf("面包屑点击回调 = %d，应为 0", clicked)
	}
	clicked = -1
	last := bc.Node().Children()[4].Bounds() // 4 = 项0,›,项1,›,项2
	r.click((last.Min.X+last.Max.X)/2, (last.Min.Y+last.Max.Y)/2)
	if clicked != -1 {
		t.Fatal("最后一项（当前页）不该响应点击")
	}
}

// TestBreadcrumbMenu 下拉菜单：menuOf 取值/越界、OnMenuClick 回调。
func TestBreadcrumbMenu(t *testing.T) {
	r := newControlRig()
	bc := &Breadcrumb{
		Items: []string{"控制台", "应用中心"},
		Menus: [][]string{{"General", "Layout"}, nil},
	}
	var got [2]int
	var labels []string
	bc.OnMenuClick = func(itemIndex int, label string) {
		got[itemIndex]++
		labels = append(labels, label)
	}
	r.mount(bc)
	if bc.menuOf(0) == nil || len(bc.menuOf(0)) != 2 {
		t.Fatal("menuOf(0) 应返回菜单列表")
	}
	if bc.menuOf(1) != nil {
		t.Fatal("menuOf(1) 空列表应返回 nil")
	}
	if bc.menuOf(5) != nil {
		t.Fatal("menuOf 越界应返回 nil")
	}
	// 菜单项浮层构建（headless 无 Overlay，仅验证状态机）：直接调
	// OnMenuClick 路径已由 showMenu 内闭包覆盖，这里锁回调签名。
	bc.OnMenuClick(0, "Layout")
	if got[0] != 1 || len(labels) != 1 || labels[0] != "Layout" {
		t.Fatalf("OnMenuClick 回调错误: %v %v", got, labels)
	}
}

func TestCollapseToggle(t *testing.T) {
	// 展开：初始全收起，点第一项表头应展开并带出内容
	r := newControlRig()
	c := NewCollapse(
		CollapsePanel{Title: "面板一", Content: Spacer{Width: 30, Height: 10}},
		CollapsePanel{Title: "面板二", Content: Spacer{Width: 30, Height: 10}},
	)
	changed := -2
	c.OnChange = func(i int) { changed = i }
	r.mount(c)
	if got := len(c.Node().Children()); got != 2 {
		t.Fatalf("折叠组节点数 = %d，应含两个纵向项", got)
	}
	if got := len(c.Node().Children()[0].Children()); got != 1 {
		t.Fatalf("收起态应只含表头，实际 %d 个子节点", got)
	}
	r.click(20, 12)
	if c.Open != 0 || changed != 0 {
		t.Fatalf("展开失败: open=%d changed=%d", c.Open, changed)
	}
	// headless 下 MarkDirty 不自动重建，手动重建验证展开态
	// 新结构：collapse 根 → 面板容器（Box）→ 内容根（表头+内容包装）。
	rebuilt := c.Build(r.ctx)
	panel := rebuilt.Children()[0].Children()[0]
	if got := len(panel.Children()); got != 2 {
		t.Fatalf("展开后应含表头+内容两个子节点，实际 %d", got)
	}

	// 收起：初始就展开第一项，再点应收回
	r2 := newControlRig()
	c2 := NewCollapse(
		CollapsePanel{Title: "面板一", Content: Spacer{Width: 30, Height: 10}},
	)
	c2.Open = 0
	changed2 := -2
	c2.OnChange = func(i int) { changed2 = i }
	r2.mount(c2)
	r2.click(20, 12)
	if c2.Open != -1 || changed2 != -1 {
		t.Fatalf("收起失败: open=%d changed=%d", c2.Open, changed2)
	}
}

// ── 类型约束：确保组件满足 Widget 接口 ────────────────────────────

var (
	_ ui.Widget = (*Card)(nil)
	_ ui.Widget = (*Badge)(nil)
	_ ui.Widget = (*Tag)(nil)
	_ ui.Widget = (*Avatar)(nil)
	_ ui.Widget = (*Empty)(nil)
	_ ui.Widget = (*Skeleton)(nil)
	_ ui.Widget = (*Spin)(nil)
	_ ui.Widget = (*Steps)(nil)
	_ ui.Widget = (*Segmented)(nil)
	_ ui.Widget = (*Pagination)(nil)
	_ ui.Widget = (*Breadcrumb)(nil)
	_ ui.Widget = (*Collapse)(nil)
)
