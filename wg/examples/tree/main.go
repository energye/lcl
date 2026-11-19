package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	tree = widgets.NewTree([]widgets.TreeNode{
		{Key: "src", Label: "src", Children: []widgets.TreeNode{
			{Key: "render", Label: "render", Children: []widgets.TreeNode{
				{Key: "canvas", Label: "canvas.go"},
			}},
			{Key: "scene", Label: "scene"},
		}},
		{Key: "widgets", Label: "widgets"},
	})

	treeCheck = widgets.NewTree([]widgets.TreeNode{
		{Key: "doc", Label: "文档", Children: []widgets.TreeNode{
			{Key: "doc-a", Label: "文档 A"},
			{Key: "doc-b", Label: "文档 B"},
		}},
		{Key: "media", Label: "媒体", Children: []widgets.TreeNode{
			{Key: "media-img", Label: "图片"},
		}},
		{Key: "misc", Label: "其他"},
	})
)

func init() {
	treeCheck.Checkable = true
	treeCheck.Checked["doc"] = true
	treeCheck.Checked["doc-a"] = true
}

func main() {
	launcher.Main(doc.Meta{
		ID: "tree", Title: "Tree 树", Group: doc.GroupData,
		What: []string{
			"层级数据展示：点击箭头展开/收起，点击标签选中",
			"Expanded 状态按 Key 记录；缩进固定 18px、行高 28px",
			"Checkable 复选模式（每行复选框独立勾选，checkStrictly 语义）",
			"计划（M9）：级联勾选/半选、节点图标、异步加载",
		},
		Demos: []doc.Demo{
			{Title: "目录树", Note: "展开 src → render，点选 canvas.go", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tree
			}},
			{Title: "checkable 复选", Note: "Checkable=true：每行复选框独立勾选（checkStrictly）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return treeCheck
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTree", Type: "func(nodes []TreeNode) *Tree", Desc: "构造；TreeNode{Key,Label,Children}"},
			{Name: "Expanded", Type: "map[string]bool", Desc: "展开状态（按 Key）"},
			{Name: "Selected", Type: "string", Desc: "选中节点 Key"},
			{Name: "Indent", Type: "float32", Default: "18", Desc: "层级缩进"},
			{Name: "RowH", Type: "float32", Default: "28", Desc: "行高"},
			{Name: "OnSelect", Type: "func(key string)", Desc: "选中回调"},
			{Name: "OnToggle", Type: "func(key string)", Desc: "展开/收起回调"},
			{Name: "Toggle", Type: "func(key string)", Desc: "程序化切换展开态"},
			{Name: "Checkable/Checked", Type: "bool / map[string]bool", Desc: "复选模式与勾选集合"},
			{Name: "OnCheck", Type: "func(checkedKeys []string)", Desc: "勾选变化回调"},
		},
	})
}
