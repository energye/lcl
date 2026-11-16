//go:build !windows

package host

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
)

// 非 Windows 平台没有窗口 DC；Presenter 退化为"只记账不提交"。
// 这样 Container 的呈现链路在 CI/headless 里仍可跑通并断言帧序。
func (p *Presenter) submit(*raster.Backing, geom.Rect) bool { return p.hwnd != 0 }
func (p *Presenter) invalidate(geom.Rect)                   {}
func (p *Presenter) release()                               {}
