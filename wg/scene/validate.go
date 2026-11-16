package scene

import "fmt"

// walkValidate 递归检查结构不变量，返回第一个问题。
//
// 只在测试与 --selftest 里跑：每帧跑会把性能门禁打穿。
// 它抓的是"声明与数据不一致"这一类静默错误——例如 KindText 但 text 为 nil，
// 这类错误在渲染上表现为"什么都不画"，特别难查。
func (n *Node) walkValidate(path []*Node) error {
	if n == nil {
		return nil
	}
	for _, p := range path {
		if p == n {
			return fmt.Errorf("节点成环：id=%d kind=%s", n.id, n.kind)
		}
	}
	path = append(path, n)

	if n.kind >= KindCount {
		return fmt.Errorf("非法 Kind=%d（id=%d）", n.kind, n.id)
	}
	if err := n.validateSelf(); err != nil {
		return fmt.Errorf("%s(id=%d): %w", n.kind, n.id, err)
	}
	for i, c := range n.children {
		if c == nil {
			return fmt.Errorf("%s(id=%d): 第 %d 个子节点为 nil", n.kind, n.id, i)
		}
		if c.parent != n {
			return fmt.Errorf("%s(id=%d): 第 %d 个子节点的 parent 指针不一致", n.kind, n.id, i)
		}
		if err := c.walkValidate(path); err != nil {
			return err
		}
	}
	return nil
}

// validateSelf 检查单个节点的 Kind 与 spec 一致性。
func (n *Node) validateSelf() error {
	// 单子节点容器
	switch n.kind {
	case KindScroll, KindClip, KindTransform, KindOpacity, KindEffect, KindSurface:
		if len(n.children) > 1 {
			return fmt.Errorf("单子节点容器却有 %d 个子节点", len(n.children))
		}
	}
	// Kind 与 spec 一致性
	switch n.kind {
	case KindText:
		if n.text == nil {
			return fmt.Errorf("KindText 但 text spec 为 nil")
		}
		if n.text.Layout != nil && len(n.children) > 0 {
			return fmt.Errorf("文本节点不应有子节点")
		}
	case KindPath:
		if n.path == nil || n.path.Path == nil {
			return fmt.Errorf("KindPath 但 path spec 为 nil")
		}
	case KindImage:
		if n.image == nil || n.image.Image == nil {
			return fmt.Errorf("KindImage 但 image spec 为 nil")
		}
	case KindCustomPaint:
		if n.custom == nil {
			return fmt.Errorf("KindCustomPaint 但回调为 nil")
		}
	case KindBox:
		if n.box == nil {
			return fmt.Errorf("KindBox 但 box spec 为 nil")
		}
	}
	// 尺寸合法性：NaN 会让整个布局链崩掉，必须在源头拦住。
	if n.size.W < 0 || n.size.H < 0 {
		return fmt.Errorf("负尺寸 %v", n.size)
	}
	if n.ratio < 0 {
		return fmt.Errorf("负宽高比 %v", n.ratio)
	}
	if n.opacity < 0 || n.opacity > 1 {
		return fmt.Errorf("opacity 越界 %v", n.opacity)
	}
	if n.kind == KindFlex && n.columns != 1 {
		return fmt.Errorf("flex 节点不应设置 columns")
	}
	return nil
}

// ValidateTree 是 Validate 的别名（便于在遍历处读起来顺）。
func (n *Node) ValidateTree() error { return n.Validate() }

// SetTransformOrigin 设置变换支点（归一化对齐；默认 Center）。
func (n *Node) SetTransformOrigin(a Alignment) *Node {
	n.originAlign = a
	n.MarkNeedsComposite()
	return n
}

// transformOriginAlign 返回支点对齐（零值对齐取 Center）。
func (n *Node) transformOriginAlign() Alignment {
	if n.originAlign == (Alignment{}) {
		return Center
	}
	return n.originAlign
}

// PaintStyle 是 BoxSpec 里 Fill 的便捷判断：是否真的会画出东西。
func (b BoxSpec) Paints() bool {
	return b.Fill.Shader != nil || !b.Fill.Color.IsTransparent()
}

// StrokeWidthOr 返回描边宽度（0 时给 1，避免"设置了描边却看不见"）。
func (b BoxSpec) StrokeWidthOr() float32 {
	if b.Border.StrokeWidth <= 0 {
		return 1
	}
	return b.Border.StrokeWidth
}
