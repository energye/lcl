//go:build !windows

package gdi

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// Painter 是非 Windows 平台的占位文字通道：不绘制，返回 0。
type Painter struct{}

// NewPainter 创建占位文字通道。
func NewPainter() *Painter { return &Painter{} }

// PaintText 不执行绘制。
func (p *Painter) PaintText(dst *raster.Backing, dl *render.DisplayList, clip geom.Rect) int { return 0 }

// Close 无操作。
func (p *Painter) Close() {}
