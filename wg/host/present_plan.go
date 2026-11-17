package host

import "github.com/energye/lcl/wg/geom"

// presentPlan 描述一帧过渡呈现（帧尺寸 ≠ 客户区）的几何方案。
//
// 语义与成熟框架一致：**内容永不缩放**。只把帧与客户区的交集按 1:1
// 画到客户区左上角（窗口缩小时自然成为裁剪），客户区多出的条带用背景色
// 填住，等下一帧（尺寸吻合）覆盖。整窗缩放是显式选项，不走这里。
type presentPlan struct {
	// Copy 是 1:1 拷贝区（左上对齐，尺寸 = min(帧, 客户区)）。
	Copy geom.Rect
	// Fill 是新暴露的条带（右/下），需要填背景色。
	Fill []geom.Rect
}

// exactPresentPlan 计算精确模式的过渡几何。
func exactPresentPlan(fw, fh, cw, ch int) presentPlan {
	var p presentPlan
	iw, ih := fw, fh
	if cw < iw {
		iw = cw
	}
	if ch < ih {
		ih = ch
	}
	if iw <= 0 || ih <= 0 {
		return p
	}
	p.Copy = geom.RectXYWH(0, 0, float32(iw), float32(ih))
	if cw > fw {
		p.Fill = append(p.Fill, geom.RectXYWH(float32(fw), 0, float32(cw-fw), float32(ch)))
	}
	if ch > fh {
		hx := fw
		if cw < hx {
			hx = cw
		}
		if hx > 0 {
			p.Fill = append(p.Fill, geom.RectXYWH(0, float32(fh), float32(hx), float32(ch-fh)))
		}
	}
	return p
}
