package diag

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/raster"
)

// ── --resizetest ────────────────────────────────────────────────────
//
// 7 次尺寸切换 + 释放顺序。C10（节流+尾随）与 C21（视图→位图→池）在这里落地。
//
// 为什么这条必须单独测：拖窗口边缘时每秒几十次 WM_SIZE，节流没做好会
// 每一像素都重建一次缓冲（拖一下就分配几十 MB）；释放顺序错了则是
// "偶发花屏 + GDI 句柄缓慢泄漏"，两者都不会在静态检查里暴露。

// ResizeCase 是单条尺寸切换断言的结果。
type ResizeCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// ResizeResult 是 --resizetest 的汇总。
type ResizeResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []ResizeCase
}

// OK 判定是否全绿（门禁 ≥7 次）。
func (r ResizeResult) OK() bool { return r.Failed == 0 && r.Cases >= 7 }

// String 返回可读摘要。
func (r ResizeResult) String() string {
	return fmt.Sprintf("resizetest: %d 次，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// resizeRig 是"调度器 + 池 + 无窗宿主"的确定性环境。
type resizeRig struct {
	clock  *frame.ManualClock
	ftick  *frame.ManualTicker
	rtick  *frame.ManualTicker
	s      *frame.Scheduler
	pool   *host.Pool
	h      *host.Host
	sizes  []string
	frames int
}

// newResizeRig 造环境。
func newResizeRig() *resizeRig {
	r := &resizeRig{
		clock: frame.NewManualClock(0, frame.DefaultStep),
		pool:  host.NewPool(4, 32<<20),
		h:     host.New(nil),
	}
	r.ftick = frame.NewManualTicker(r.clock)
	r.rtick = frame.NewManualTicker(r.clock)
	s, err := frame.NewScheduler(frame.Options{
		Clock:        r.clock,
		Ticker:       r.ftick,
		ResizeTicker: r.rtick,
		OnFrame:      func(now, delta time.Duration) bool { r.frames++; return true },
		OnResize: func(w, h int) {
			r.sizes = append(r.sizes, fmt.Sprintf("%dx%d", w, h))
			r.h.Resize(w, h)
		},
		WantFrame: func(now time.Duration) (bool, time.Duration) { return false, 0 },
	})
	if err != nil {
		panic("wg-diag: 创建调度器失败: " + err.Error())
	}
	r.s = s
	return r
}

// ResizeTest 跑尺寸切换门禁。
func ResizeTest() (ResizeResult, error) {
	began := time.Now()
	var res ResizeResult
	cases := []struct {
		name string
		run  func() (string, error)
	}{
		{"burst-throttled-with-trailing", caseBurstThrottled},
		{"quiet-window-splits", caseQuietSplits},
		{"host-size-follows", caseHostSizeFollows},
		{"grow-invalidates-new-area", caseGrowInvalidates},
		{"pool-hits-same-size", casePoolHit},
		{"release-order-view-first", caseReleaseOrder},
		{"pool-budget-evicts", casePoolBudget},
		{"seven-switches-no-leak", caseSevenSwitches},
	}
	for _, c := range cases {
		got, err := c.run()
		d := ResizeCase{Name: c.name, OK: err == nil, Got: got}
		if err != nil {
			d.Note = err.Error()
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

// caseBurstThrottled：7 次请求间隔 20ms（比 60ms 窗口密）。
// 节流 + 尾随语义下：拖动中每 ~60ms 执行一次（不再是 0 次），停下后
// 尾随把最后一次尺寸落地；最终执行的一定是最后的请求尺寸。
func caseBurstThrottled() (string, error) {
	r := newResizeRig()
	sizes := [][2]int{{800, 600}, {810, 600}, {820, 610}, {830, 620}, {840, 620}, {850, 640}, {860, 640}}
	for _, s := range sizes {
		r.s.RequestResize(s[0], s[1])
		r.rtick.Advance(20 * time.Millisecond)
	}
	r.rtick.Advance(frame.DefaultResizeDebounce)
	got := fmt.Sprintf("%d 次触发 %v", len(r.sizes), r.sizes)
	if len(r.sizes) < 2 || len(r.sizes) > 5 {
		return got, fmt.Errorf("拖动期间应有节流触发的 resize（2–5 次）")
	}
	if r.sizes[0] != "800x600" {
		return got, fmt.Errorf("首次请求应前沿立即执行")
	}
	if r.sizes[len(r.sizes)-1] != "860x640" {
		return got, fmt.Errorf("尾随应取最后一次尺寸")
	}
	return got, nil
}

// caseQuietSplits：中间静默超过节流窗口 → 应该触发两次。
func caseQuietSplits() (string, error) {
	r := newResizeRig()
	r.s.RequestResize(800, 600)
	r.rtick.Advance(frame.DefaultResizeDebounce + 10*time.Millisecond)
	r.s.RequestResize(900, 700)
	r.rtick.Advance(frame.DefaultResizeDebounce + 10*time.Millisecond)
	got := fmt.Sprintf("%d 次 %v", len(r.sizes), r.sizes)
	if len(r.sizes) != 2 {
		return got, fmt.Errorf("静默后应再触发一次")
	}
	return got, nil
}

// caseHostSizeFollows：resize 回调后宿主与根画布尺寸都应更新。
func caseHostSizeFollows() (string, error) {
	r := newResizeRig()
	r.s.RequestResize(640, 480)
	r.rtick.Advance(frame.DefaultResizeDebounce + time.Millisecond)
	cs := r.h.ClientSize()
	rb := r.h.Root().Bounds()
	got := fmt.Sprintf("client=%vx%v root=%vx%v", cs.W, cs.H, rb.W(), rb.H())
	if cs.W != 640 || cs.H != 480 {
		return got, fmt.Errorf("客户区未更新")
	}
	if rb.W() != 640 || rb.H() != 480 {
		return got, fmt.Errorf("根画布未更新")
	}
	return got, nil
}

// caseGrowInvalidates：窗口变大时新增区域必须脏（否则残留旧像素）。
func caseGrowInvalidates() (string, error) {
	r := newResizeRig()
	r.s.RequestResize(400, 300)
	r.rtick.Advance(frame.DefaultResizeDebounce + time.Millisecond)
	r.h.Root().ClearDamageTree()

	r.s.RequestResize(600, 400)
	r.rtick.Advance(frame.DefaultResizeDebounce + time.Millisecond)
	d := r.h.Root().Damage()
	got := fmt.Sprintf("damage=%v（需覆盖 600x400）", d)
	if d.W() < 600 || d.H() < 400 {
		return got, fmt.Errorf("新增区域未标记脏")
	}
	return got, nil
}

// casePoolHit：同尺寸再取应命中，且内容被清零。
func casePoolHit() (string, error) {
	p := host.NewPool(4, 8<<20)
	b1 := p.Acquire(64, 64)
	if b1 == nil {
		return "", fmt.Errorf("取缓冲失败")
	}
	b1.Pix()[0] = 0xFFFF0000
	if !p.ReleaseView(b1) {
		return "", fmt.Errorf("解视图失败")
	}
	if !p.Release(b1) {
		return "", fmt.Errorf("归还失败")
	}
	b2 := p.Acquire(64, 64)
	st := p.Stats()
	got := fmt.Sprintf("hits=%d reuse=%.0f%% first=%08x", st.Hits, st.Reused()*100, b2.Pix()[0])
	if st.Hits != 1 {
		return got, fmt.Errorf("同尺寸应命中池")
	}
	if b2.Pix()[0] != 0 {
		return got, fmt.Errorf("复用缓冲未清零")
	}
	return got, nil
}

// caseReleaseOrder：未解视图就归池必须被拒绝（C21）。
func caseReleaseOrder() (string, error) {
	p := host.NewPool(4, 8<<20)
	b := p.Acquire(32, 32)
	if p.Release(b) {
		return "归池成功", fmt.Errorf("有活视图时不该归池")
	}
	if st := p.Stats(); st.WrongOrder != 1 {
		return fmt.Sprintf("wrongOrder=%d", st.WrongOrder), fmt.Errorf("未记录顺序违规")
	}
	if !p.ReleaseView(b) {
		return "", fmt.Errorf("解视图失败")
	}
	if !p.Release(b) {
		return "", fmt.Errorf("解视图后应可归池")
	}
	return fmt.Sprintf("wrongOrder=%d free=%d", p.Stats().WrongOrder, p.Stats().Free), nil
}

// casePoolBudget：超过字节预算时淘汰缓冲。
func casePoolBudget() (string, error) {
	// 预算 256KiB：一个 128×128 缓冲 64KiB，放得下 4 个，第 5 个必须触发淘汰。
	p := host.NewPool(8, 256<<10)
	var bs []*raster.Backing
	for i := 0; i < 5; i++ {
		b := p.Acquire(128, 128)
		if b == nil {
			return "", fmt.Errorf("第 %d 次取缓冲失败", i)
		}
		bs = append(bs, b)
	}
	for _, b := range bs {
		p.ReleaseView(b)
		p.Release(b)
	}
	st := p.Stats()
	got := fmt.Sprintf("free=%d bytes=%d evict=%d", st.Free, st.Bytes, st.Evictions)
	if st.Bytes > 256<<10 {
		return got, fmt.Errorf("超出字节预算")
	}
	if st.Evictions == 0 {
		return got, fmt.Errorf("应发生淘汰")
	}
	return got, nil
}

// caseSevenSwitches：7 次尺寸切换后不泄漏、池不越界。
func caseSevenSwitches() (string, error) {
	r := newResizeRig()
	dims := [][2]int{{1024, 768}, {1280, 800}, {1024, 768}, {1440, 880}, {800, 600}, {1440, 880}, {1024, 768}}
	for _, d := range dims {
		r.s.RequestResize(d[0], d[1])
		r.rtick.Advance(frame.DefaultResizeDebounce + time.Millisecond)
		// 每次切换都走一遍"借出 → 解视图 → 归池"
		b := r.pool.Acquire(d[0]/4, d[1]/4)
		if b == nil {
			return "", fmt.Errorf("%dx%d 取缓冲失败", d[0], d[1])
		}
		r.pool.ReleaseView(b)
		r.pool.Release(b)
	}
	st := r.pool.Stats()
	got := fmt.Sprintf("触发=%d leaked=%d free=%d wrongOrder=%d reuse=%.0f%%",
		len(r.sizes), r.pool.Leaked(), st.Free, st.WrongOrder, st.Reused()*100)
	if len(r.sizes) != 7 {
		return got, fmt.Errorf("7 次切换应各触发一次")
	}
	if r.pool.Leaked() != 0 {
		return got, fmt.Errorf("存在未解视图（泄漏）")
	}
	if st.WrongOrder != 0 {
		return got, fmt.Errorf("释放顺序违规 %d 次", st.WrongOrder)
	}
	if r.h.ClientSize() != geom.Sz(1024, 768) {
		return got, fmt.Errorf("末次尺寸未生效")
	}
	return got, nil
}
