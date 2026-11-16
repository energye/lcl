package diag

import (
	"errors"
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/compat"
	"github.com/energye/lcl/wg/migration"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// ── --migrationtest：三示例迁移 + compat 门面门禁（M4）────────────
//
// 单测覆盖了场景构建与交互；这里的额外价值是**公共入口链路**：
// 真实 Container 出帧 + 旧式门面能出节点 + 零 lcl 静态扫描。
// 组件接线断了单测可能仍绿，这里一定红。

// MigrationCase 是单条断言结果。
type MigrationCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// MigrationResult 是汇总。
type MigrationResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []MigrationCase
}

func (r MigrationResult) OK() bool { return r.Failed == 0 && r.Cases >= 7 }

func (r MigrationResult) String() string {
	return fmt.Sprintf("migrationtest: %d 项，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// MigrationTest 运行三示例迁移门禁。
func MigrationTest() (MigrationResult, error) {
	began := time.Now()
	var res MigrationResult
	cases := []struct {
		name string
		run  func() (string, error)
	}{
		{"scenes-render", caseMigrateScenes},
		{"compat-facade", caseMigrateCompat},
		{"zero-lcl-migration", func() (string, error) { return caseMigrateZeroLCL("migration") }},
		{"zero-lcl-chart", func() (string, error) { return caseMigrateZeroLCL("widgets/chart") }},
		{"zero-lcl-compat", func() (string, error) { return caseMigrateZeroLCL("compat") }},
		{"showcase-charts", caseMigrateShowcaseCharts},
		{"imedemo-text", caseMigrateIMEText},
	}
	for _, c := range cases {
		got, err := c.run()
		d := MigrationCase{Name: c.name, OK: err == nil, Got: got}
		if err != nil {
			d.Note = err.Error()
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

func caseMigrateScenes() (string, error) {
	var got []string
	for _, sc := range []struct {
		name string
		vw   int
		make func() ui.Widget
	}{
		{"showcase", 1440, func() ui.Widget { return migration.NewShowcase() }},
		{"canvas_ime", 960, func() ui.Widget { return migration.NewCanvasIME() }},
		{"dbmodeler", 1100, func() ui.Widget { return migration.NewDBModeler() }},
	} {
		c := wg.NewHeadless(wg.DefaultOptions())
		c.Resize(sc.vw, 720)
		c.MountWidget(sc.make())
		stats, err := c.FrameOnce()
		if err != nil {
			return strings.Join(got, " "), fmt.Errorf("%s 出帧失败: %w", sc.name, err)
		}
		if stats.Nodes == 0 {
			return strings.Join(got, " "), fmt.Errorf("%s 无渲染节点", sc.name)
		}
		got = append(got, fmt.Sprintf("%s:%d节点", sc.name, stats.Nodes))
	}
	return strings.Join(got, " "), nil
}

func caseMigrateCompat() (string, error) {
	ctx := &ui.BuildContext{Theme: theme.Default()}
	btn := compat.TButton("确定")
	n := btn.Build(ctx)
	if n == nil || btn.Text != "确定" {
		return "nil", errors.New("compat.TButton 转发失败")
	}
	tbl := compat.TTable(
		[]widgets.TableColumn{{Title: "A", Mode: widgets.ColumnAuto}},
		[][]any{{"x"}},
	)
	if len(tbl.Rows) != 1 {
		return "rows=0", errors.New("compat.TTable 转发失败")
	}
	return "TButton/TTable OK", nil
}

func caseMigrateShowcaseCharts() (string, error) {
	s := migration.NewShowcase()
	ctx := &ui.BuildContext{Theme: theme.Default()}
	n := s.Build(ctx)
	custom := 0
	n.Walk(func(x *scene.Node) bool {
		if x.Kind() == scene.KindCustomPaint {
			custom++
		}
		return true
	})
	if custom < 5 {
		return fmt.Sprintf("自定义绘制=%d", custom), errors.New("showcase 图表节点不足")
	}
	return fmt.Sprintf("图表=%d", custom), nil
}

func caseMigrateIMEText() (string, error) {
	c := migration.NewCanvasIME()
	c.Submit.OnClick()
	got := c.Input.Editor().Value().Text
	if got != "你好" {
		return got, errors.New("IME 组合串上屏失败")
	}
	return got, nil
}

func caseMigrateZeroLCL(pkgPath string) (string, error) {
	pkg, err := build.Import("github.com/energye/lcl/wg/"+pkgPath, "", build.FindOnly)
	if err != nil {
		return err.Error(), err
	}
	fset := token.NewFileSet()
	parsed, err := parser.ParseDir(fset, pkg.Dir, nil, parser.ImportsOnly)
	if err != nil {
		return err.Error(), err
	}
	files, bad := 0, 0
	for _, p := range parsed {
		for _, f := range p.Files {
			files++
			for _, spec := range f.Imports {
				path := strings.Trim(spec.Path.Value, `"`)
				if path == "github.com/energye/lcl" || strings.HasPrefix(path, "github.com/energye/lcl/lcl") {
					bad++
				}
			}
		}
	}
	got := fmt.Sprintf("%s: 文件=%d 违规=%d", filepath.Base(pkgPath), files, bad)
	if bad != 0 {
		return got, errors.New("出现 lcl 绑定依赖")
	}
	return got, nil
}
