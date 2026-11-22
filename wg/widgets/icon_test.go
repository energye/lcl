package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/render"
)

// TestIconPaintersNonEmpty 强自证：每个 painter 至少一条非空路径
// （未知图标=橙块兜底——painter 写错名会画橙块，此测在源头拦截）。
func TestIconPaintersNonEmpty(t *testing.T) {
	if len(iconPainters) < 84 {
		t.Fatalf("图标数 = %d，应 ≥84（第一百零九轮扩展）", len(iconPainters))
	}
	for name, def := range iconPainters {
		if def.stroke == nil && def.fill == nil {
			t.Fatalf("图标 %q 无 stroke 也无 fill（会渲染成橙块）", name)
		}
		if def.stroke != nil {
			p := render.NewPath()
			def.stroke(iconBrush{p: p})
			if p.IsEmpty() {
				t.Fatalf("图标 %q stroke 路径为空", name)
			}
		}
		if def.fill != nil {
			p := render.NewPath()
			def.fill(iconBrush{p: p})
			if p.IsEmpty() {
				t.Fatalf("图标 %q fill 路径为空", name)
			}
		}
	}
}

// TestIconNamesSorted IconNames 排序稳定且包含新扩展。
func TestIconNamesSorted(t *testing.T) {
	names := IconNames()
	if len(names) < 84 {
		t.Fatalf("IconNames = %d，应 ≥84", len(names))
	}
	for i := 1; i < len(names); i++ {
		if names[i] < names[i-1] {
			t.Fatalf("IconNames 未排序 at %d: %q > %q", i, names[i-1], names[i])
		}
	}
	// 扩展代表图标在集内。
	for _, want := range []string{"database", "zoom-in", "bold", "play-circle", "global", "team"} {
		found := false
		for _, n := range names {
			if n == want {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("扩展图标 %q 不在 IconNames", want)
		}
	}
}
