package widgets

// ── Layout 布局容器（M12 收尾）───────────────────────────────────
//
// antd：<Layout><Sider/><Header/><Content/><Footer/></Layout>
// 语义化布局包装：Flex Vertical [Header, Flex Horizontal [Sider,
// Content], Footer]——各子区语义锚（layout-sider/header/content/
// footer）。antd 默认 Sider 深底 #001529（dark token 缺——用
// SurfaceSunken 近似并如实标注）；Content 白底。
// 未实现（如实标注）：Sider 折叠动画/collapsedWidth/breakpoint、
// trigger 收起钮、theme=dark 深色 token、Layout 嵌套语义样式。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// Layout 是语义化布局容器。
type Layout struct {
	// Header 顶栏内容（高 64，antd 默认）。
	Header ui.Widget
	// Sider 侧边栏内容（宽 SiderWidth，antd 默认 200）。
	Sider ui.Widget
	// Content 主内容。
	Content ui.Widget
	// Footer 底栏内容（高 48，antd 默认）。
	Footer ui.Widget
	// SiderWidth 侧边栏宽（默认 200）。
	SiderWidth float32
	// HeaderHeight 顶栏高（默认 64）。
	HeaderHeight float32
	// FooterHeight 底栏高（默认 48）。
	FooterHeight float32
	// Width/Height 总尺寸（Height<=0 自适应内容高）。
	Width, Height float32
}

// NewLayout 创建布局容器。
func NewLayout(content ui.Widget) *Layout {
	return &Layout{Content: content, SiderWidth: 200, HeaderHeight: 64, FooterHeight: 48}
}

// Build 实现 ui.Widget。
func (l *Layout) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	root := scene.Flex(scene.Vertical, 0)
	if l.Width > 0 {
		root.SetWidth(l.Width)
	}
	root.SetHit(scene.HitSpec{Label: "layout"})

	hh := l.HeaderHeight
	if hh <= 0 {
		hh = 64
	}
	fh := l.FooterHeight
	if fh <= 0 {
		fh = 48
	}
	sw := l.SiderWidth
	if sw <= 0 {
		sw = 200
	}

	if l.Header != nil {
		header := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)}).SetWidth(l.Width).SetHeight(hh)
		if inner := l.Header.Build(ctx); inner != nil {
			inner.SetMargin(geom.InLTRB(16, 0, 16, 0))
			header.Add(inner)
		}
		header.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "layout-header"})
		root.Add(header)
	}

	mid := scene.Flex(scene.Horizontal, 0)
	if l.Height > 0 {
		mid.SetHeight(l.Height - hh - fh)
	}
	if l.Sider != nil {
		sider := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)}).SetWidth(sw)
		if l.Height > 0 {
			sider.SetHeight(l.Height - hh - fh)
		}
		if inner := l.Sider.Build(ctx); inner != nil {
			inner.SetMargin(geom.InLTRB(8, 8, 8, 8))
			sider.Add(inner)
		}
		sider.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "layout-sider"})
		mid.Add(sider)
	}
	if l.Content != nil {
		content := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)})
		if l.Width > 0 {
			content.SetWidth(l.Width - sw)
		}
		if l.Height > 0 {
			content.SetHeight(l.Height - hh - fh)
		}
		if inner := l.Content.Build(ctx); inner != nil {
			inner.SetMargin(geom.InLTRB(16, 16, 16, 16))
			content.Add(inner)
		}
		content.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "layout-content"})
		mid.Add(content)
	}
	root.Add(mid)

	if l.Footer != nil {
		footer := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)}).SetWidth(l.Width).SetHeight(fh)
		if inner := l.Footer.Build(ctx); inner != nil {
			inner.SetMargin(geom.InLTRB(16, 0, 16, 0))
			footer.Add(inner)
		}
		footer.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "layout-footer"})
		root.Add(footer)
	}
	return root
}
