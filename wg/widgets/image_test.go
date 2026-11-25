package widgets

import "testing"

// TestImagePreviewTransform 预览变换：缩放步进/钳制、旋转循环。
func TestImagePreviewTransform(t *testing.T) {
	im := &Image{}
	im.previewScale = 1

	// 放大 1.5×。
	im.applyZoom(1.5)
	if im.previewScale != 1.5 {
		t.Fatalf("放大后 scale=%v，want 1.5", im.previewScale)
	}
	// 缩小回 1 以下钳制到 1。
	im.applyZoom(0.1)
	if im.previewScale != 1 {
		t.Fatalf("缩小钳制后 scale=%v，want 1", im.previewScale)
	}
	// 超大钳制到 50。
	im.applyZoom(1000)
	if im.previewScale != 50 {
		t.Fatalf("超大钳制后 scale=%v，want 50", im.previewScale)
	}

	// 旋转 90° 步进 + 循环。
	im.previewRotate = 0
	im.applyRotate(1)
	if im.previewRotate != 1 {
		t.Fatalf("右转后 rotate=%d，want 1", im.previewRotate)
	}
	im.applyRotate(-1)
	im.applyRotate(-1)
	if im.previewRotate != 3 {
		t.Fatalf("左转两次后 rotate=%d，want 3（0-1+4 循环）", im.previewRotate)
	}
	im.applyRotate(1)
	if im.previewRotate != 0 {
		t.Fatalf("再右转后 rotate=%d，want 0", im.previewRotate)
	}
}
