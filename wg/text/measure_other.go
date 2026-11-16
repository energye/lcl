//go:build !windows

package text

import "github.com/energye/lcl/wg/geom"

// 非 Windows 平台没有 GDI，用经验字宽估算：
//   - ASCII 取 0.55em（接近 Segoe UI 的平均字宽）；
//   - CJK 取 1.0em（方块字）；
//   - 其它按 0.75em。
//
// 这套估算只保证"能排版、能跑测试"，不保证与 Windows 出图一致。
// 真正的度量与字形回放只在 Windows 上可用（见 measure_windows.go）。

func measureString(s string, f Font) geom.Size {
	em := f.Size
	if em <= 0 {
		em = 14
	}
	w := float32(0)
	for _, r := range s {
		w += runeAdvance(r, f)
	}
	return geom.Sz(w, f.lineHeight())
}

func runeAdvance(r rune, f Font) float32 {
	em := f.Size
	if em <= 0 {
		em = 14
	}
	switch {
	case r < 0x80:
		if r == ' ' {
			return em * 0.28
		}
		if r >= 'A' && r <= 'Z' {
			return em * 0.62
		}
		return em * 0.55
	case r >= 0x2E80 && r <= 0x9FFF, r >= 0xF900 && r <= 0xFAFF,
		r >= 0x3000 && r <= 0x303F, r >= 0xFF00 && r <= 0xFFEF:
		return em // CJK 与全角标点
	default:
		return em * 0.75
	}
}

func ascentOf(f Font) float32 {
	em := f.Size
	if em <= 0 {
		em = 14
	}
	return em * 0.8
}

// FontHandle 在非 Windows 平台无意义，返回 0。
func FontHandle(f Font) uintptr { return 0 }
