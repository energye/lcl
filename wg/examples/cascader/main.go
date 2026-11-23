package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有：Build 闭包内新建实例会在每次根重建时状态归零
// （教训⑨变体，TimePicker 轮实证）。
var (
	cdBase   = widgets.NewCascader()
	cdInit   = widgets.NewCascader()
	cdSearch = widgets.NewCascader()
	cdLog    = ""
)

func init() {
	opts := []*widgets.CascaderOption{
		{Value: "zhejiang", Label: "浙江", Children: []*widgets.CascaderOption{
			{Value: "hangzhou", Label: "杭州", Children: []*widgets.CascaderOption{
				{Value: "xihu", Label: "西湖区"},
				{Value: "shangcheng", Label: "上城区"},
			}},
			{Value: "ningbo", Label: "宁波", Children: []*widgets.CascaderOption{
				{Value: "haishu", Label: "海曙区"},
			}},
		}},
		{Value: "jiangsu", Label: "江苏", Children: []*widgets.CascaderOption{
			{Value: "nanjing", Label: "南京", Children: []*widgets.CascaderOption{
				{Value: "xuanwu", Label: "玄武区"},
			}},
			{Value: "suzhou", Label: "苏州", Children: []*widgets.CascaderOption{
				{Value: "gusu", Label: "姑苏区"},
			}},
		}},
	}
	cdBase.Options = opts
	cdInit.Options = opts
	cdInit.Value = []string{"zhejiang", "hangzhou", "xihu"}
	cdSearch.Options = opts
}

func main() {
	launcher.Main(doc.Meta{
		ID: "cascader", Title: "Cascader 级联选择", Group: doc.GroupControl,
		What: []string{
			"触发框（Input 形态+down 图标+路径文本 \" / \" 分隔）→ 点击弹级联面板（多列并排，列 180px/item 28px/视口 240px）",
			"点含子级的选项展开下一列（同帧重开面板——宽度随列数走）；叶子选择关面板（默认只能选叶子）",
			"路径上选项文字主色（antd 级联文字色高亮）；hover 浅底；Scroll 列视口组合复用",
			"**showSearch** 搜索：触发框变输入框，输入按路径 label 过滤，搜索态扁平列表（显示完整路径，点选即选中整链）",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "逐级点选：浙江 → 杭州 → 西湖区；触发框显示完整路径", Build: func(ctx *ui.BuildContext) ui.Widget {
				return cdBase
			}},
			{Title: "带初值", Note: "Value=zhejiang/hangzhou/xihu：触发框显示路径文本", Build: func(ctx *ui.BuildContext) ui.Widget {
				return cdInit
			}},
			{Title: "搜索", Note: "showSearch：输入按路径过滤，搜索态扁平列表显示完整路径", Build: func(ctx *ui.BuildContext) ui.Widget {
				cdSearch.ShowSearch = true
				cdSearch.OnSearch = func(s string) {
					cdLog = "search:" + s
				}
				return cdSearch
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCascader", Type: "func() *Cascader", Desc: "构造（默认占位\"请选择\"）"},
			{Name: "Options", Type: "[]*CascaderOption", Desc: "级联数据（Value/Label/Children）"},
			{Name: "Value", Type: "[]string", Desc: "选中路径（叶子链的 Value 链）"},
			{Name: "OnChange", Type: "func(path []string)", Desc: "叶子选择回调"},
			{Name: "Placeholder", Type: "string", Default: "请选择", Desc: "占位文本"},
			{Name: "Width", Type: "float32", Default: "220", Desc: "触发框宽度"},
			{Name: "ShowSearch", Type: "bool", Default: "false", Desc: "搜索框（触发框变输入框，按路径过滤）"},
			{Name: "Filter / OnSearch", Type: "func(inputValue, path) bool / func(search)", Desc: "自定义过滤 / 搜索回调"},
			{Name: "Show / Close / IsOpen", Type: "func(...)", Desc: "程序化开关面板 / 状态查询"},
		},
	})
}
