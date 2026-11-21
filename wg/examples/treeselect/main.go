package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨：Build 闭包内新建实例每次根重建状态归零）。
var (
	tsBase = widgets.NewTreeSelect([]widgets.TreeNode{
		{Key: "comp", Label: "组件", Children: []widgets.TreeNode{
			{Key: "form", Label: "表单", Children: []widgets.TreeNode{
				{Key: "input", Label: "输入框"},
				{Key: "select", Label: "选择器"},
			}},
			{Key: "display", Label: "展示", Children: []widgets.TreeNode{
				{Key: "table", Label: "表格"},
				{Key: "list", Label: "列表"},
			}},
		}},
		{Key: "nav", Label: "导航", Children: []widgets.TreeNode{
			{Key: "top", Label: "顶部导航"},
			{Key: "side", Label: "侧边导航"},
		}},
	})
	tsInit = widgets.NewTreeSelect([]widgets.TreeNode{
		{Key: "comp", Label: "组件", Children: []widgets.TreeNode{
			{Key: "form", Label: "表单", Children: []widgets.TreeNode{
				{Key: "input", Label: "输入框"},
				{Key: "select", Label: "选择器"},
			}},
			{Key: "display", Label: "展示", Children: []widgets.TreeNode{
				{Key: "table", Label: "表格"},
				{Key: "list", Label: "列表"},
			}},
		}},
		{Key: "nav", Label: "导航", Children: []widgets.TreeNode{
			{Key: "top", Label: "顶部导航"},
			{Key: "side", Label: "侧边导航"},
		}},
	})
	tsInited bool
)

func main() {
	launcher.Main(doc.Meta{
		ID: "treeselect", Title: "TreeSelect 树选择", Group: doc.GroupControl,
		What: []string{
			"触发框（Input 形态+down 图标）→ 点击弹树面板（**Tree 组件原样内嵌**——展开/缩进/选中/hover 全部组合复用，零重复实现）",
			"选中叶子节点 → 触发框显示 Label + 面板关闭（antd TreeSelect 默认选后关）",
			"Tree 实例跨面板开合持有：展开状态保留；树面板高度按行数自适应",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "展开「组件 → 表单」后点「输入框」；触发框显示节点 Label", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tsBase
			}},
			{Title: "带初值", Note: "Value=comp/form/input：触发框显示「输入框」，面板选中高亮", Build: func(ctx *ui.BuildContext) ui.Widget {
				if !tsInited {
					tsInit.Value = "input" // Value 语义=选中节点 Key（非路径）
					tsInited = true
				}
				return tsInit
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTreeSelect", Type: "func(nodes []TreeNode) *TreeSelect", Desc: "构造（复用 Tree 的 TreeNode）"},
			{Name: "Nodes", Type: "[]TreeNode", Desc: "树数据（Key/Label/Children——与 Tree 组件同构）"},
			{Name: "Value", Type: "string", Desc: "选中节点 Key"},
			{Name: "OnChange", Type: "func(key string)", Desc: "叶子选择回调"},
			{Name: "Placeholder", Type: "string", Default: "请选择", Desc: "占位文本"},
			{Name: "Width", Type: "float32", Default: "220", Desc: "触发框宽度"},
			{Name: "Show / Close / IsOpen", Type: "func(...)", Desc: "程序化开关面板 / 状态查询"},
		},
	})
}
