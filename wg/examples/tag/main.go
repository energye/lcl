package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// closableDemo 是可关闭标签演示组件：持有存活标签集合，× 点击真实移除。
// 嵌入 ui.Base：OnClose 移除元素后 SetState 触发整窗重建，否则点击后
// 界面表现为"点了没反应"（静止窗 C14 无帧请求）。
type closableDemo struct {
	ui.Base
	alive []*widgets.Tag
}

func newClosableDemo() *closableDemo {
	tags := []*widgets.Tag{
		{Text: "golang", Preset: "processing"},
		{Text: "React", Preset: "success"},
		{Text: "Vue", Preset: "error"},
	}
	for _, t := range tags {
		t.Closable = true
	}
	return &closableDemo{alive: tags}
}

// closableDemoInst 包级持有：闭包内新建会让存活状态/重建随演示重建丢失。
var closableDemoInst = newClosableDemo()

func (d *closableDemo) Build(ctx *ui.BuildContext) *scene.Node {
	if ctx.InvalidateRoot != nil {
		d.Attach(ctx.InvalidateRoot)
	}
	row := scene.Flex(scene.Horizontal, 8)
	for _, t := range d.alive {
		tt := t
		tt.OnClose = func() {
			d.SetState(func() {
				for i, x := range d.alive {
					if x == tt {
						d.alive = append(d.alive[:i], d.alive[i+1:]...)
						break
					}
				}
			})
		}
		row.Add(tt.Build(ctx))
	}
	note := (&widgets.Label{Text: fmt.Sprintf("存活 %d 个", len(d.alive))}).Build(ctx)
	note.SetMargin(geom.InLTRB(0, 8, 0, 0))
	row.Add(note)
	return row
}

// checkableDemo 可选中标签演示（包级持有实例：点击切换选中态）。
var checkableTags = func() []*widgets.Tag {
	tags := []*widgets.Tag{
		{Text: "后端", Checkable: true},
		{Text: "前端", Checkable: true},
		{Text: "测试", Checkable: true},
	}
	return tags
}()

func main() {
	launcher.Main(doc.Meta{
		ID: "tag", Title: "Tag 标签", Group: doc.GroupDisplay,
		What: []string{
			"默认（无 Color）= 灰底 #FAFAFA + colorBorder 描边 + 深色文字（antd 默认 Tag）",
			"显式 Color = 浅底（基色 40% 透明）+ 同色文字 + 基色淡描边",
			"preset 预设色：success/error/warning/processing 语义色（浅底+同色文字）",
			"closable 可关闭：× 图标点击真实移除；checkable 可选中：主色实底+白字切换",
		},
		Demos: []doc.Demo{
			{Title: "标签组", Note: "宽度随文本自适应", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 8)
				for _, t := range []string{"Go", "LCL", "无句柄画布", "渲染引擎"} {
					row.Add((&widgets.Tag{Text: t}).Build(ctx))
				}
				row.Add((&widgets.Tag{Text: "可配置色", Color: geom.RGBA(22, 119, 255)}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "preset 预设色", Note: "success/error/warning/processing 语义色", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 8)
					for _, p := range []string{"success", "error", "warning", "processing"} {
						row.Add((&widgets.Tag{Text: p, Preset: p}).Build(c))
					}
					return row
				})
			}},
			{Title: "closable 可关闭", Note: "点 × 真实移除（存活数联动）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return closableDemoInst
			}},
			{Title: "checkable 可选中", Note: "点击切换主色实底+白字", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 8)
					for _, t := range checkableTags {
						row.Add(t.Build(c))
					}
					return row
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "Text", Type: "string", Desc: "标签文字"},
			{Name: "Color", Type: "geom.Color", Default: "零值=默认灰底", Desc: "显式时按预设色渲染（浅底+同色文字）"},
			{Name: "Preset", Type: "string", Desc: "预设色名 success/error/warning/processing（优先于 Color）"},
			{Name: "Closable/OnClose", Type: "bool / func()", Desc: "× 可关闭与移除回调"},
			{Name: "Checkable/Checked/OnCheck", Type: "bool / bool / func(bool)", Desc: "可选中态与切换回调"},
		},
	})
}
