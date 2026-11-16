package examples

import (
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	Register(Example{ID: "card", Name: "Card 卡片", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		card := widgets.NewCard(&widgets.Label{Text: "卡片正文 · 内容由 ui.Widget 组合"})
		card.Title = &widgets.Label{Text: "卡片标题"}
		card.Footer = &widgets.Label{Text: "底部区域"}
		return Section(ctx, "Card", "标题/内容/底部三段式，Elevated 控制阴影", card.Build(ctx))
	}})

	Register(Example{ID: "avatar", Name: "Avatar 头像", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		row := scene.Flex(scene.Horizontal, 12)
		for _, s := range []struct {
			text string
			size float32
		}{{"张", 28}, {"李", 36}, {"WG", 48}} {
			row.Add((&widgets.Avatar{Text: s.text, Size: s.size}).Build(ctx))
		}
		return Section(ctx, "Avatar", "纯画布圆形头像，无位图依赖", row)
	}})

	Register(Example{ID: "badge", Name: "Badge 徽标", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		row := scene.Flex(scene.Horizontal, 18)
		row.Add((&widgets.Badge{Count: 3}).Build(ctx))
		row.Add((&widgets.Badge{Text: "NEW"}).Build(ctx))
		row.Add((&widgets.Badge{Dot: true}).Build(ctx))
		return Section(ctx, "Badge", "计数/文本/圆点三态，可与任意节点组合使用", row)
	}})

	Register(Example{ID: "tag", Name: "Tag 标签", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		row := scene.Flex(scene.Horizontal, 8)
		for _, t := range []string{"Go", "LCL", "无句柄画布", "渲染引擎"} {
			row.Add((&widgets.Tag{Text: t}).Build(ctx))
		}
		return Section(ctx, "Tag", "按文本 intrinsic 宽度自适应", row)
	}})

	Register(Example{ID: "empty", Name: "Empty 空状态", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		e := widgets.Empty{Text: "暂无数据", Subtext: "点击左侧其他示例看看", Width: 320}
		return Section(ctx, "Empty", "空状态插图与两级文案，全部用 scene 图元画出", e.Build(ctx))
	}})

	Register(Example{ID: "skeleton", Name: "Skeleton 骨架屏", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		s := widgets.Skeleton{Rows: 4, Width: 360, Height: 14}
		return Section(ctx, "Skeleton", "加载占位；Phase 由调用方驱动，组件本身无定时器", s.Build(ctx))
	}})

	Register(Example{ID: "spin", Name: "Spin 加载指示", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		row := scene.Flex(scene.Horizontal, 24)
		for _, p := range []float32{0, 0.25, 0.5, 0.75} {
			row.Add((&widgets.Spin{Size: 28, Phase: p}).Build(ctx))
		}
		return Section(ctx, "Spin", "CustomPaintNode + render.Path 弧线；Phase 明确交调用方控制", row)
	}})

	Register(Example{ID: "steps", Name: "Steps 步骤条", Group: GroupDisplay, Build: func(ctx *ui.BuildContext) *scene.Node {
		s := widgets.Steps{Labels: []string{"选择", "配置", "确认", "完成"}, Current: 2}
		return Section(ctx, "Steps", "当前步及之前高亮，后续置灰", s.Build(ctx))
	}})

	{
		seg := widgets.NewSegmented("日", "周", "月", "年")
		Register(Example{ID: "segmented", Name: "Segmented 分段选择", Group: GroupDisplay,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Segmented", "点击切换选中段；手势检测器跨 Build 存活", seg.Build(ctx))
			}})
	}

	{
		page := widgets.NewPagination(3, 8)
		Register(Example{ID: "pagination", Name: "Pagination 分页", Group: GroupDisplay,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Pagination", "只显示当前页附近页码，首末页可跳转", page.Build(ctx))
			}})
	}

	{
		status := widgets.NewButton("未点击")
		crumb := &widgets.Breadcrumb{Items: []string{"首页", "资源管理", "渲染引擎"}}
		crumb.OnClick = func(i int) {
			status.Text = "点击了：" + crumb.Items[i]
			status.MarkDirty()
		}
		Register(Example{ID: "breadcrumb", Name: "Breadcrumb 面包屑", Group: GroupDisplay,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(crumb.Build(ctx))
				col.Add(status.Build(ctx))
				return Section(ctx, "Breadcrumb", "最后一项为当前页不可点；其余项可点击回调", col)
			}})
	}

	{
		collapse := widgets.NewCollapse(
			widgets.CollapsePanel{Title: "第一节 · 渲染管线", Content: &widgets.Label{Text: "RenderNode → Canvas → DisplayList → Layer → damage"}},
			widgets.CollapsePanel{Title: "第二节 · 交互链路", Content: &widgets.Label{Text: "命中路径 → Dispatcher → GestureArena → 回调"}},
			widgets.CollapsePanel{Title: "第三节 · 输入与 IME", Content: &widgets.Label{Text: "C28：TextKeyHandler 与 IME 结果分离"}},
		)
		Register(Example{ID: "collapse", Name: "Collapse 折叠面板", Group: GroupDisplay,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Collapse", "点击表头展开/收起；表头宽度走 CrossStretch", collapse.Build(ctx))
			}})
	}
}
