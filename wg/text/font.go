// Package text 是文本子系统：字体描述、断行、省略、对齐、度量与字形缓存。
//
// 它只做排版计算，**不画字**：字形回放交给光栅后端在设备空间完成，
// 以保留 ClearType（契约 C6/C24）。因此本包不含任何绘制代码。
//
// v1 不做 shaping（HarfBuzz/bidi）：按 rune 累加 advance，覆盖中英文、
// emoji 与换行；阿拉伯/印度语系的连写需要 shaping，M5 条件接入。
//
// 依赖：geom。平台度量通过 measure_*.go 提供，非 Windows 平台退化为
// 按字宽估算（保证测试可在任意平台跑，只是精度差）。
package text

import (
	"sync"

	"github.com/energye/lcl/wg/geom"
)

// Weight 是字重（对应 Windows LOGFONT.lfWeight 的常用档位）。
type Weight uint16

const (
	// Thin 100
	Thin Weight = 100
	// Normal 400
	Normal Weight = 400
	// Medium 500
	Medium Weight = 500
	// SemiBold 600
	SemiBold Weight = 600
	// Bold 700
	Bold Weight = 700
	// Black 900
	Black Weight = 900
)

// Font 是字体描述（值类型，可比较，便于做缓存键）。
type Font struct {
	// Family 是字体族名，空字符串表示系统默认 UI 字体。
	Family string
	// Size 是字号（DIP，即 pt 换算后的像素高度）。
	Size float32
	// Weight 是字重。
	Weight Weight
	// Italic 斜体。
	Italic bool
	// Underline 下划线。
	Underline bool
}

// DefaultFont 返回系统默认无衬线字体（Segoe UI / Microsoft YaHei 兜底）。
func DefaultFont(size float32) Font {
	if size <= 0 {
		size = 14
	}
	return Font{Family: "", Size: size, Weight: Normal}
}

// WithSize 返回换字号的副本。
func (f Font) WithSize(size float32) Font { f.Size = size; return f }

// WithWeight 返回换字重的副本。
func (f Font) WithWeight(w Weight) Font { f.Weight = w; return f }

// WithFamily 返回换字体族的副本。
func (f Font) WithFamily(family string) Font { f.Family = family; return f }

// faceKey 是字体在平台层的唯一标识（也用作度量缓存键）。
type faceKey struct {
	Family    string
	SizeMilli int32 // 字号 × 1000，避免浮点做 map key
	Weight    Weight
	Italic    bool
}

func (f Font) key() faceKey {
	return faceKey{Family: f.Family, SizeMilli: int32(f.Size*1000 + 0.5), Weight: f.Weight, Italic: f.Italic}
}

// Align 是水平对齐。
type Align uint8

const (
	// AlignStart 左对齐（LTR 语义）。
	AlignStart Align = iota
	// AlignCenter 居中。
	AlignCenter
	// AlignEnd 右对齐。
	AlignEnd
)

// Line 是一行排版结果。
type Line struct {
	// Text 是该行的文本（不含换行符）。
	Text string
	// Width 是行宽（DIP）。
	Width float32
	// Baseline 是该行基线相对 Layout 顶部的偏移（DIP）。
	Baseline float32
}

// Layout 是排版结果（不可变）。
type Layout struct {
	Font  Font
	Lines []Line
	// Size 是整体尺寸（DIP）。
	Size geom.Size
	// Align 是水平对齐（影响每行起始 x）。
	Align Align
	// Truncated 表示因 maxWidth/行数限制被省略。
	Truncated bool
	// LineHeight 是行高（DIP）。
	LineHeight float32
}

// Empty 表示没有可绘制内容。
func (l *Layout) Empty() bool { return l == nil || len(l.Lines) == 0 }

// LineX 返回第 i 行在给定对齐下的起始 x 偏移。
func (l *Layout) LineX(i int) float32 {
	if i < 0 || i >= len(l.Lines) {
		return 0
	}
	switch l.Align {
	case AlignCenter:
		return (l.Size.W - l.Lines[i].Width) * 0.5
	case AlignEnd:
		return l.Size.W - l.Lines[i].Width
	default:
		return 0
	}
}

// Options 是排版选项。
type Options struct {
	// MaxWidth 是换行宽度上限；<=0 表示不换行（单行）。
	MaxWidth float32
	// MaxLines 是行数上限；<=0 表示不限。
	MaxLines int
	// Ellipsis 是省略符，空字符串表示不省略。仅在超出 MaxLines 或
	// 单行超宽且 MaxWidth>0 时追加。
	Ellipsis string
	// Align 是水平对齐。
	Align Align
	// LineHeight 是行高倍数（0 表示 1.2）。
	LineHeight float32
}

// measureCache 缓存"字符串 + 字体"的度量，避免每帧调 GDI。
//
// 这是本包唯一的全局可变状态，用 RWMutex 保护；键里带了字号与字重，
// 因此 DPI 变化会自然产生新键（不会读到旧值）。
var measureCache = struct {
	sync.RWMutex
	m map[string]geom.Size
}{m: make(map[string]geom.Size, 256)}

// Measure 返回字符串在给定字体下的尺寸（单行，DIP）。
func Measure(s string, f Font) geom.Size {
	if s == "" {
		return geom.Sz(0, f.lineHeight())
	}
	k := cacheKey(s, f)
	measureCache.RLock()
	v, ok := measureCache.m[k]
	measureCache.RUnlock()
	if ok {
		return v
	}
	v = measureString(s, f)
	measureCache.Lock()
	// 简单上限控制：超过 8192 条直接清空，避免长期运行无界增长。
	if len(measureCache.m) > 8192 {
		measureCache.m = make(map[string]geom.Size, 256)
	}
	measureCache.m[k] = v
	measureCache.Unlock()
	return v
}

// ClearMeasureCache 清空度量缓存（切换 DPI/主题时调用）。
func ClearMeasureCache() {
	measureCache.Lock()
	measureCache.m = make(map[string]geom.Size, 256)
	measureCache.Unlock()
	clearPlatformCaches()
}

func cacheKey(s string, f Font) string {
	k := f.key()
	// 用单字节分隔，避免 "a|b" 与 "a" + "|b" 碰撞。
	buf := make([]byte, 0, len(s)+24)
	buf = append(buf, k.Family...)
	buf = append(buf, 0, byte(k.SizeMilli>>16), byte(k.SizeMilli>>8), byte(k.SizeMilli),
		byte(k.Weight>>8), byte(k.Weight), 0)
	if k.Italic {
		buf = append(buf, 1)
	} else {
		buf = append(buf, 0)
	}
	buf = append(buf, 0)
	buf = append(buf, s...)
	return string(buf)
}

// lineHeight 返回字体行高（DIP）。
func (f Font) lineHeight() float32 {
	if f.Size <= 0 {
		return 16
	}
	// 1.2 倍是排版惯例；平台度量拿到真实 ascent+descent 时会被覆盖。
	return f.Size * 1.2
}

// LayoutText 排版单一样式文本。
//
// 断行策略：按 rune 累加宽度，超出 MaxWidth 时回退到最近的空格/连字符
// 断点（CJK 逐字断行）；没有断点则硬断。这与 03-ENGINE.md §6.2 一致。
func LayoutText(s string, f Font, opts Options) *Layout {
	lh := f.lineHeight()
	if opts.LineHeight > 0 {
		lh = f.Size * opts.LineHeight
	}
	l := &Layout{Font: f, Align: opts.Align, LineHeight: lh}

	// 先按显式换行切分
	paragraphs := splitLines(s)
	var lines []Line
	for _, para := range paragraphs {
		if lines, l.Truncated = wrapParagraph(lines, para, f, opts, lh); l.Truncated {
			break
		}
	}
	if len(lines) == 0 {
		lines = append(lines, Line{Text: "", Width: 0, Baseline: lh})
	}

	l.Lines = lines
	w := float32(0)
	for i := range l.Lines {
		l.Lines[i].Baseline = float32(i)*lh + ascentOf(f)
		if l.Lines[i].Width > w {
			w = l.Lines[i].Width
		}
	}
	l.Size = geom.Sz(w, float32(len(l.Lines))*lh)
	return l
}

// wrapParagraph 按 opts 断行，返回追加后的行。
func wrapParagraph(lines []Line, para string, f Font, opts Options, lh float32) ([]Line, bool) {
	if opts.MaxWidth <= 0 {
		if opts.MaxLines > 0 && len(lines)+1 > opts.MaxLines {
			return lines, true
		}
		w := Measure(para, f).W
		return append(lines, Line{Text: para, Width: w}), false
	}

	runes := []rune(para)
	start := 0
	// 空段落直接产出一条空行：旧写法让循环体对空 runes 也进一次，
	// end=start+1 的"单字硬塞"会把 runes[0:1] 切出 NUL 字符送去 GDI 度量
	// （StringToUTF16 对 NUL 直接 panic）。空文本是合法输入（表格空单元格）。
	if len(runes) == 0 {
		return append(lines, Line{Text: "", Width: 0}), false
	}
	for start < len(runes) {
		if opts.MaxLines > 0 && len(lines) >= opts.MaxLines {
			return lines, true
		}
		// 逐字扩展，直到超宽
		end := start
		width := float32(0)
		lastBreak := -1
		for end < len(runes) {
			w := runeAdvance(runes[end], f)
			if width+w > opts.MaxWidth && end > start {
				break
			}
			width += w
			if isBreakRune(runes[end]) {
				lastBreak = end + 1
			}
			end++
		}
		if end == start {
			end = start + 1 // 单字超宽，硬塞一个，避免死循环
		}
		// 有断点且断点不在末尾时回退
		if lastBreak > start && lastBreak < end && end < len(runes) {
			end = lastBreak
		}
		text := string(runes[start:end])
		text = trimTrailingSpaces(text)
		w := Measure(text, f).W
		lines = append(lines, Line{Text: text, Width: w})
		start = end
		if opts.MaxLines > 0 && len(lines) >= opts.MaxLines && start < len(runes) {
			return lines, true
		}
		if len(runes) == 0 {
			break
		}
	}
	return lines, false
}

// Wrap 是 layout 内部用的辅助：把显式换行切段，保留空段。
func splitLines(s string) []string {
	out := make([]string, 0, 4)
	cur := make([]rune, 0, 32)
	for _, r := range s {
		if r == '\n' {
			out = append(out, string(cur))
			cur = cur[:0]
			continue
		}
		if r == '\r' {
			continue
		}
		cur = append(cur, r)
	}
	out = append(out, string(cur))
	return out
}

func trimTrailingSpaces(s string) string {
	i := len(s)
	for i > 0 && s[i-1] == ' ' {
		i--
	}
	return s[:i]
}

func isBreakRune(r rune) bool {
	return r == ' ' || r == '\t' || r == '-' || r == '\u3001' || r == '\u3002' || r == '\uff0c'
}
