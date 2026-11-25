package widgets

// ── Statistic 统计数值（M12 长尾）────────────────────────────────
//
// antd：<Statistic title value precision separator prefix suffix>
// 纯展示排版：标题灰字 + 大号粗体数值（千分位分隔/小数精度）+
// 前后缀小字。Countdown 倒计时形态与 valueStyle 深度定制未实现
// （如实标注）。

import (
	"fmt"
	"strings"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// Statistic 是统计数值展示。
type Statistic struct {
	controlBase
	Title     string
	Value     float64
	Precision int    // 小数位（默认 0）
	Separator string // 千分位分隔符（默认 ","；"" 不分隔）
	Prefix    string // 前缀小字（如 ¥）
	Suffix    string // 后缀小字（如 元）
	Width     float32
}

// NewStatistic 创建统计数值。
func NewStatistic(title string, value float64) *Statistic {
	return &Statistic{Title: title, Value: value, Separator: ","}
}

// formatValue 千分位+精度格式化（antd 默认 separator=","）。
func (s *Statistic) formatValue() string {
	prec := s.Precision
	if prec < 0 {
		prec = 0
	}
	body := fmt.Sprintf("%.*f", prec, s.Value)
	sep := s.Separator
	if sep == "" || (len(body) <= 3 && !strings.Contains(body, ".")) {
		return body
	}
	dot := strings.Index(body, ".")
	intPart, frac := body, ""
	if dot >= 0 {
		intPart, frac = body[:dot], body[dot:]
	}
	n := len(intPart)
	if n > 3 {
		var b strings.Builder
		for i, ch := range intPart {
			if i > 0 && (n-i)%3 == 0 {
				b.WriteString(sep)
			}
			b.WriteRune(ch)
		}
		intPart = b.String()
	}
	return intPart + frac
}

// Build 实现 ui.Widget。
func (s *Statistic) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(s)
	th := ctx.ThemeOr()
	w := s.Width
	if w <= 0 {
		w = 140
	}
	col := scene.Flex(scene.Vertical, th.Space.XS)
	col.SetWidth(w)
	col.SetHit(scene.HitSpec{Label: "statistic"})

	if s.Title != "" {
		tl := text.LayoutText(s.Title, text.DefaultFont(th.Font.SizeBase), text.Options{})
		title := scene.Text(scene.TextSpec{Src: s.Title, Layout: tl, Color: th.Color.TextMuted})
		title.SetHit(scene.HitSpec{Label: "statistic-title"})
		col.Add(title)
	}

	// 数值行：前后缀继承 contentFontSize（antd 前后缀与数值同为 fontSizeHeading3=24），数值 SemiBold 600。
	row := scene.Flex(scene.Horizontal, th.Space.XS)
	if s.Prefix != "" {
		pl := text.LayoutText(s.Prefix, text.DefaultFont(th.Font.SizeTitle), text.Options{})
		pv := scene.Text(scene.TextSpec{Src: s.Prefix, Layout: pl, Color: th.Color.TextMuted})
		pv.SetHit(scene.HitSpec{Label: "statistic-prefix"})
		row.Add(pv)
	}
	vf := text.DefaultFont(th.Font.SizeTitle).WithWeight(text.SemiBold)
	vl := text.LayoutText(s.formatValue(), vf, text.Options{})
	value := scene.Text(scene.TextSpec{Src: s.formatValue(), Layout: vl, Color: th.Color.Text})
	value.SetMargin(geom.InLTRB(0, (28-vl.Size.H)/2, 0, 0))
	value.SetHit(scene.HitSpec{Label: "statistic-value"})
	row.Add(value)
	if s.Suffix != "" {
		sl := text.LayoutText(s.Suffix, text.DefaultFont(th.Font.SizeTitle), text.Options{})
		sv := scene.Text(scene.TextSpec{Src: s.Suffix, Layout: sl, Color: th.Color.TextMuted})
		sv.SetHit(scene.HitSpec{Label: "statistic-suffix"})
		sv.SetMargin(geom.InLTRB(0, (28-sl.Size.H)/2, 0, 0))
		row.Add(sv)
	}
	col.Add(row)
	return col
}
