//go:build windows

package text

import (
	"sync"
	"syscall"
	"unsafe"

	"github.com/energye/lcl/wg/geom"
)

// Windows 度量与字体句柄：直接走 gdi32，不经过 LCL。
//
// 为什么不用 lcl.ICanvas 度量：text 包必须能在无窗口、无 LCL 初始化的
// 环境下工作（CI、golden 出图、契约自检），而 LCL 的字体对象依赖其运行时。
//
// 线程安全：一个内存 DC 被所有度量复用，配一把互斥锁。度量是纯查询，
// 不涉及绘制状态，串行化的代价可以忽略（有缓存兜底）。

var (
	gdi32                    = syscall.NewLazyDLL("gdi32.dll")
	procCreateCompatibleDC   = gdi32.NewProc("CreateCompatibleDC")
	procCreateFontW          = gdi32.NewProc("CreateFontW")
	procSelectObject         = gdi32.NewProc("SelectObject")
	procDeleteObject         = gdi32.NewProc("DeleteObject")
	procDeleteDC             = gdi32.NewProc("DeleteDC")
	procGetTextExtentPoint32 = gdi32.NewProc("GetTextExtentPoint32W")
	procGetTextMetricsW      = gdi32.NewProc("GetTextMetricsW")
	procGetDeviceCaps        = gdi32.NewProc("GetDeviceCaps")
)

// LOGFONTW 布局（92 字节）。字段顺序必须与 wingdi.h 一致。
type logFontW struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         int32
	Italic         byte
	Underline      byte
	StrikeOut      byte
	CharSet        byte
	OutPrecision   byte
	ClipPrecision  byte
	Quality        byte
	PitchAndFamily byte
	FaceName       [32]uint16
}

// textMetricW 只声明到 tmDescent 为止（后续字段不需要）。
type textMetricW struct {
	Height           int32
	Ascent           int32
	Descent          int32
	InternalLeading  int32
	ExternalLeading  int32
	AveCharWidth     int32
	MaxCharWidth     int32
	Weight           int32
	Overhang         int32
	DigitizedAspectX int32
	DigitizedAspectY int32
	FirstChar        uint16
	LastChar         uint16
	DefaultChar      uint16
	BreakChar        uint16
	Italic           byte
	Underlined       byte
	StruckOut        byte
	PitchAndFamily   byte
	FaceName         [32]uint16
}

const (
	lfWeightNormal     = 400
	cleartypeQuality   = 5 // CLEARTYPE_QUALITY
	defaultCharset     = 1 // DEFAULT_CHARSET
	outTTPrecis        = 4 // OUT_TT_PRECIS
	clipDefaultPrecis  = 0
	pitchAndFamilyVar  = 0
	loGPixelsPerInchY  = 90 // LOGPIXELSY
	antialiasedQuality = 4  // ANTIALIASED_QUALITY
)

var (
	dcOnce   sync.Once
	memDC    uintptr
	dpiOnce  sync.Once
	dpiScale = float32(1)

	fontMu    sync.Mutex
	fontCache = make(map[faceKey]uintptr, 32)
)

// ensureDC 懒加载一个内存 DC（CreateCompatibleDC(0) 不需要窗口）。
func ensureDC() uintptr {
	dcOnce.Do(func() {
		h, _, _ := procCreateCompatibleDC.Call(0)
		memDC = h
	})
	return memDC
}

// scale 返回系统 DPI / 96。
//
// 语义约定：Font.Size 是 DIP。GDI 只认像素，因此度量时按 scale 放大，
// 再把结果除回 scale，得到 DIP 尺寸。这样在 150% 缩放下排版宽度不变，
// 换行位置也不会漂移。
func scale() float32 {
	dpiOnce.Do(func() {
		dc := ensureDC()
		if dc == 0 {
			return
		}
		v, _, _ := procGetDeviceCaps.Call(dc, loGPixelsPerInchY)
		if v > 0 {
			dpiScale = float32(v) / 96.0
		}
	})
	return dpiScale
}

// FontHandle 返回缓存中的 GDI 字体句柄（HFONT）。调用方不得删除它。
func FontHandle(f Font) uintptr {
	if f.Size <= 0 {
		f.Size = 14
	}
	k := f.key()
	fontMu.Lock()
	defer fontMu.Unlock()
	if h, ok := fontCache[k]; ok {
		return h
	}
	dc := ensureDC()
	if dc == 0 {
		return 0
	}
	lf := logFontW{
		Height:         -int32(f.Size*scale() + 0.5),
		Weight:         int32(f.Weight),
		CharSet:        defaultCharset,
		OutPrecision:   outTTPrecis,
		ClipPrecision:  clipDefaultPrecis,
		Quality:        cleartypeQuality,
		PitchAndFamily: pitchAndFamilyVar,
		Italic:         boolByte(f.Italic),
		Underline:      boolByte(f.Underline),
	}
	family := f.Family
	if family == "" {
		family = "Segoe UI"
	}
	copyUTF16(lf.FaceName[:], family)

	h, _, _ := procCreateFontW.Call(
		uintptr(lf.Height), uintptr(lf.Width),
		uintptr(lf.Escapement), uintptr(lf.Orientation), uintptr(lf.Weight),
		uintptr(lf.Italic), uintptr(lf.Underline), uintptr(lf.StrikeOut),
		uintptr(lf.CharSet), uintptr(lf.OutPrecision), uintptr(lf.ClipPrecision),
		uintptr(lf.Quality), uintptr(lf.PitchAndFamily),
		uintptr(unsafe.Pointer(&lf.FaceName[0])),
	)
	if h == 0 {
		return 0
	}
	fontCache[k] = h
	return h
}

// measureString 用 GDI 度量单行文本，结果换算为 DIP。
func measureString(s string, f Font) geom.Size {
	dc := ensureDC()
	hf := FontHandle(f)
	if dc == 0 || hf == 0 {
		return geom.Sz(0, f.lineHeight())
	}
	sc := scale()

	fontMu.Lock()
	prev, _, _ := procSelectObject.Call(dc, hf)
	var sz struct{ CX, CY int32 }
	u := syscall.StringToUTF16Ptr(s)
	procGetTextExtentPoint32.Call(dc, uintptr(unsafe.Pointer(u)), uintptr(len(syscall.StringToUTF16(s))-1), uintptr(unsafe.Pointer(&sz)))
	if prev != 0 {
		procSelectObject.Call(dc, prev)
	}
	var tm textMetricW
	procSelectObject.Call(dc, hf)
	procGetTextMetricsW.Call(dc, uintptr(unsafe.Pointer(&tm)))
	if prev != 0 {
		procSelectObject.Call(dc, prev)
	}
	fontMu.Unlock()

	h := float32(sz.CY) / sc
	if h < f.Size {
		h = f.lineHeight()
	}
	return geom.Sz(float32(sz.CX)/sc, h)
}

// runeAdvance 返回单个字符的推进宽度（DIP）。
func runeAdvance(r rune, f Font) float32 {
	return measureString(string(r), f).W
}

// ascentOf 返回基线到顶部的距离（DIP）。
func ascentOf(f Font) float32 {
	dc := ensureDC()
	hf := FontHandle(f)
	if dc == 0 || hf == 0 {
		return f.lineHeight() * 0.8
	}
	fontMu.Lock()
	prev, _, _ := procSelectObject.Call(dc, hf)
	var tm textMetricW
	procGetTextMetricsW.Call(dc, uintptr(unsafe.Pointer(&tm)))
	if prev != 0 {
		procSelectObject.Call(dc, prev)
	}
	fontMu.Unlock()
	a := float32(tm.Ascent) / scale()
	if a <= 0 {
		return f.lineHeight() * 0.8
	}
	return a
}

// ReleaseFonts 释放所有缓存的字体句柄（进程退出或 DPI 变化时调用）。
func ReleaseFonts() {
	fontMu.Lock()
	for k, h := range fontCache {
		procDeleteObject.Call(h)
		delete(fontCache, k)
	}
	fontMu.Unlock()
	ClearMeasureCache()
}

func boolByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

// copyUTF16 把字符串写入定长 uint16 数组（截断，保证 0 结尾）。
func copyUTF16(dst []uint16, s string) {
	u := syscall.StringToUTF16(s)
	n := len(u)
	if n > len(dst) {
		n = len(dst)
	}
	copy(dst[:n], u[:n])
	dst[len(dst)-1] = 0
}
