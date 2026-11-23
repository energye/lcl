package main

import (
	"time"

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

	// treeCascade 级联勾选（CheckStrictly=false 默认）：勾父全选子孙、
	// 子孙部分勾选父半选。预置 doc 半选态（doc-a 勾 doc-b 不勾）。
	// Label 唯一化（资料/图库/手册 A…）：多卡页面 text 锚撞车是定位
	// 混乱的根源（Select/Tag 轮同款教训）。
	treeCascade = widgets.NewTree([]widgets.TreeNode{
		{Key: "all", Label: "全部资料", Children: []widgets.TreeNode{
			{Key: "grp-doc", Label: "手册", Children: []widgets.TreeNode{
				{Key: "doc-a", Label: "手册 A"},
				{Key: "doc-b", Label: "手册 B"},
			}},
			{Key: "grp-media", Label: "图库"}, // 无子节点：压缩卡高（页尾卡滚动可达性）
		}},
	})

	// treeAsync 异步加载卡（antd loadData）：LoadData 注入模拟取数，
	// 异步形态——AfterFunc 300ms 延迟 SetChildren（loading 态可见）。
	treeAsync = widgets.NewTree([]widgets.TreeNode{
		{Key: "srv", Label: "服务列表"},
		{Key: "db", Label: "数据库"},
	})
)

func init() {
	treeAsync.LoadData = func(key string) []widgets.TreeNode {
		// 异步形态：返回 nil + 宿主延迟挂载（模拟网络往返）。
		go func() {
			time.Sleep(300 * time.Millisecond)
			switch key {
			case "srv":
				treeAsync.SetChildren(key, []widgets.TreeNode{
					{Key: "srv-1", Label: "用户服务"},
					{Key: "srv-2", Label: "订单服务"},
				})
			case "db":
				treeAsync.SetChildren(key, []widgets.TreeNode{
					{Key: "db-1", Label: "主库"},
				})
			}
		}()
		return nil
	}
}

func init() {
	treeCheck.Checkable = true
	treeCheck.CheckStrictly = true // 旧行为保持：独立勾选
	treeCheck.Checked["doc"] = true
	treeCheck.Checked["doc-a"] = true
	treeCascade.Checkable = true
	treeCascade.Expanded["all"] = true
	treeCascade.Expanded["grp-doc"] = true  // 「文档」展开：子孙复选框入树

	treeCascade.Checked["doc-a"] = true // 部分勾选 → 「手册」半选
	treeCascade.SyncChecked()           // 外部预置后重算父链半选态
}

func main() {
	launcher.Main(doc.Meta{
		ID: "tree", Title: "Tree 树", Group: doc.GroupData,
		What: []string{
			"层级数据展示：点击箭头展开/收起，点击标签选中",
			"Expanded 状态按 Key 记录；缩进固定 18px、行高 28px",
			"Checkable 复选：CheckStrictly=true 独立勾选 / false 级联（勾父全选子孙、部分选父半选）",
		},
		Demos: []doc.Demo{
			{Title: "目录树", Note: "展开 src → render，点选 canvas.go", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tree
			}},
			{Title: "checkable 复选", Note: "CheckStrictly=true：每行复选框独立勾选", Build: func(ctx *ui.BuildContext) ui.Widget {
				return treeCheck
			}},
			{Title: "级联勾选", Note: "CheckStrictly=false：预置「手册 A」勾选 →「手册」半选；勾「全部资料」级联全选", Build: func(ctx *ui.BuildContext) ui.Widget {
				return treeCascade
			}},
			{Title: "异步加载", Note: "LoadData 注入+宿主延迟 SetChildren：首次展开触发加载（loading 态），300ms 后子节点挂载", Build: func(ctx *ui.BuildContext) ui.Widget {
				return treeAsync
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
			{Name: "CheckStrictly", Type: "bool", Default: "false", Desc: "true=独立勾选 / false=级联（父全选子孙+半选联动）"},
			{Name: "OnCheck", Type: "func(checkedKeys []string)", Desc: "勾选变化回调"},
			{Name: "LoadData", Type: "func(key string) []TreeNode", Desc: "异步加载子节点注入（首次展开触发；同步返回数据/异步 nil+SetChildren）"},
			{Name: "SetChildren", Type: "func(key string, children []TreeNode)", Desc: "宿主挂载异步加载结果（MarkDirty 刷新）"},
		},
	})
}
