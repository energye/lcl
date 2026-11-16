package examples

import (
	"testing"

	"github.com/energye/lcl/wg"
)

// TestNoSilentDegradation 是渲染能力的门禁：每条示例都必须真的被后端执行，
// 不允许"指令进去了、后端画不出来、只记一个降级数"还能通过。
//
// 这条测试来自一次真实事故：位图绘制在软件光栅里长期未实现（C22 只计
// degraded），而示例里同时有文字，像素断言被文字满足，于是 render-image
// 一直"通过"却什么都画不出来。能力缺失必须让它红灯。
func TestNoSilentDegradation(t *testing.T) {
	for _, ex := range All {
		ex := ex
		t.Run(ex.ID, func(t *testing.T) {
			c := wg.NewHeadless(wg.DefaultOptions())
			c.Resize(1024, 768)
			c.MountWidget(ex.Widget())
			st, err := c.FrameOnce()
			if err != nil {
				t.Fatalf("首帧失败: %v", err)
			}
			if st.Degraded > 0 {
				t.Fatalf("示例有 %d 条指令被后端降级（能力未实现或后端不支持）", st.Degraded)
			}
		})
	}
}
