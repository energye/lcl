package widgets

// ── Countdown 倒计时（第一百零六轮，低频项收敛）──────────────────
//
// antd：<Statistic.Countdown value format onFinish onChange>
// （5.25.0 更名 Statistic.Timer，API 同——type=countdown|countup）。
// 独立组件、Statistic 排版规范复用（24 SemiBold 数值+SM 灰标题）。
//
// 驱动：Build 内按 Now 注入点计算剩余，time.AfterFunc 到下一秒边界
// 自排程 MarkDirty（Carousel Autoplay 递归模式）——每秒一帧，静止
// 收敛 0 帧；Now 可注入（Calendar 同款跨天确定性设计）。
//
// format 为 dayjs token 简化子集：D（天，不补零）/HH（补零时，≥100h
// 展开实际数）/mm/ss（补零分秒）/SSS（毫秒三位）。默认 "HH:mm:ss"。

import (
	"fmt"
	"strings"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// Countdown 是倒计时/正计时展示。
type Countdown struct {
	controlBase
	Title  string
	Prefix string
	Suffix string
	// Target countdown=截止时刻 / countup=起始时刻。
	Target time.Time
	// CountUp 正计时（antd Timer type="countup"）。
	CountUp bool
	// Format dayjs token 子集（默认 "HH:mm:ss"）。
	Format string
	// OnFinish 时间到回调（仅 countdown）。
	OnFinish func()
	// OnChange 每 tick 回调（剩余/已过时长）。
	OnChange func(d time.Duration)
	// Now 当前时刻注入点（nil=time.Now；审计/回归跨天确定性）。
	Now func() time.Time

	Width     float32
	timerDone bool // 本次倒计时是否已触发 OnFinish（防重复）
	timer     *time.Timer
}

// NewCountdown 创建倒计时（默认 1 小时后截止、HH:mm:ss）。
func NewCountdown(title string, target time.Time) *Countdown {
	return &Countdown{Title: title, Target: target, Format: "HH:mm:ss"}
}

// nowNow 返回注入时刻或 time.Now。
func (c *Countdown) nowNow() time.Time {
	if c.Now != nil {
		return c.Now()
	}
	return time.Now()
}

// formatDuration 按 dayjs token 子集格式化时长。
// fmt 中 token 逐段替换：D/HH/H/mm/ss/SSS；其余字符原样。
func formatDuration(d time.Duration, f string) string {
	if d < 0 {
		d = 0
	}
	totalMS := int64(d / time.Millisecond)
	millis := totalMS % 1000
	totalSec := totalMS / 1000
	sec := totalSec % 60
	totalMin := totalSec / 60
	min := totalMin % 60
	totalHour := totalMin / 60
	hour := totalHour % 24
	day := totalHour / 24

	rep := func(s string) string {
		s = strings.ReplaceAll(s, "SSS", fmt.Sprintf("%03d", millis))
		if day > 0 {
			s = strings.ReplaceAll(s, "D", fmt.Sprintf("%d", day))
		} else {
			s = strings.ReplaceAll(s, "D", "0")
		}
		if strings.Contains(s, "HH") {
			s = strings.ReplaceAll(s, "HH", fmt.Sprintf("%02d", hour))
		}
		s = strings.ReplaceAll(s, "H", fmt.Sprintf("%d", totalHour))
		s = strings.ReplaceAll(s, "mm", fmt.Sprintf("%02d", min))
		s = strings.ReplaceAll(s, "ss", fmt.Sprintf("%02d", sec))
		return s
	}
	// 先拆分非 token 字面量，避免 "D" 误伤普通字母段落。
	var b strings.Builder
	token := ""
	flush := func() {
		if token != "" {
			b.WriteString(rep(token))
			token = ""
		}
	}
	for _, r := range f {
		if strings.ContainsRune("DHmsS", r) {
			token += string(r)
		} else {
			flush()
			b.WriteRune(r)
		}
	}
	flush()
	return b.String()
}

// remaining 计算当前显示时长（countdown=剩余，countup=已过）。
func (c *Countdown) remaining() time.Duration {
	now := c.nowNow()
	if c.CountUp {
		if c.Target.After(now) {
			return 0
		}
		return now.Sub(c.Target)
	}
	if now.After(c.Target) {
		return 0
	}
	return c.Target.Sub(now)
}

// scheduleNext 安排下一秒边界重绘（剩余>0 时；Carousel 递归模式）。
// c.timer 跨 Build 持有——旧 timer 先停（重建时重排不叠加）。
// AfterFunc 回调直接触发 OnFinish（不依赖重建——MarkDirty 在无帧
// 泵的测试/静窗环境下不产生重 Build）。
func (c *Countdown) scheduleNext() {
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	d := c.remaining()
	if d <= 0 {
		c.fireFinishIfDue()
		return
	}
	if c.OnChange != nil {
		c.OnChange(d)
	}
	// 到下一秒边界（ceil 到整秒）。
	wait := d % time.Second
	if wait == 0 {
		wait = time.Second
	}
	t := time.AfterFunc(wait, func() {
		nd := c.remaining()
		if nd <= 0 && !c.CountUp {
			c.fireFinishIfDue()
		} else {
			// 未到点：递归自排程下一秒（Carousel Autoplay 同款）。
			c.scheduleNext()
		}
		c.MarkDirty()
	})
	c.timer = t
}

// fireFinishIfDue 触发一次 OnFinish（countdown 到点；防重复）。
func (c *Countdown) fireFinishIfDue() {
	if c.CountUp || c.timerDone {
		return
	}
	if c.nowNow().Before(c.Target) {
		return
	}
	c.timerDone = true
	if c.OnFinish != nil {
		c.OnFinish()
	}
}

// Stop 停止定时（宿主销毁/页面卸载时调用）。
func (c *Countdown) Stop() {
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
}

// Build 实现 ui.Widget。
func (c *Countdown) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	th := ctx.ThemeOr()
	c.scheduleNext()
	w := c.Width
	if w <= 0 {
		w = 140
	}
	col := scene.Flex(scene.Vertical, th.Space.XS)
	col.SetWidth(w)
	col.SetHit(scene.HitSpec{Label: "countdown"})

	if c.Title != "" {
		tl := text.LayoutText(c.Title, text.DefaultFont(th.Font.SizeSM), text.Options{})
		title := scene.Text(scene.TextSpec{Src: c.Title, Layout: tl, Color: th.Color.TextMuted})
		title.SetHit(scene.HitSpec{Label: "countdown-title"})
		col.Add(title)
	}
	row := scene.Flex(scene.Horizontal, th.Space.XS)
	if c.Prefix != "" {
		pl := text.LayoutText(c.Prefix, text.DefaultFont(th.Font.SizeLG), text.Options{})
		pv := scene.Text(scene.TextSpec{Src: c.Prefix, Layout: pl, Color: th.Color.TextMuted})
		pv.SetHit(scene.HitSpec{Label: "countdown-prefix"})
		row.Add(pv)
	}
	txt := formatDuration(c.remaining(), c.Format)
	vf := text.DefaultFont(24).WithWeight(text.SemiBold)
	vl := text.LayoutText(txt, vf, text.Options{})
	value := scene.Text(scene.TextSpec{Src: txt, Layout: vl, Color: th.Color.Text})
	value.SetMargin(geom.InLTRB(0, (28-vl.Size.H)/2, 0, 0))
	value.SetHit(scene.HitSpec{Label: "countdown-value"})
	row.Add(value)
	if c.Suffix != "" {
		sl := text.LayoutText(c.Suffix, text.DefaultFont(th.Font.SizeLG), text.Options{})
		sv := scene.Text(scene.TextSpec{Src: c.Suffix, Layout: sl, Color: th.Color.TextMuted})
		sv.SetHit(scene.HitSpec{Label: "countdown-suffix"})
		sv.SetMargin(geom.InLTRB(0, (28-sl.Size.H)/2, 0, 0))
		row.Add(sv)
	}
	col.Add(row)
	return col
}
