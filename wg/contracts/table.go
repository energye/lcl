package contracts

// 本表是 C1–C28 的唯一事实来源。
//
// 编号与标题来自：
//   - C1–C17：ARCHITECTURE.md §10.1（既有实证，全部继承）
//   - C18–C21：02-CANVAS.md（画布体系）
//   - C22–C25：03-ENGINE.md（渲染内核）
//   - C26–C28：04-UI-INTERACTION.md（交互）
//
// 新增契约必须同时：① 在三份文档里出现；② 在本表登记；③ 提供断言或明确 Since。
// 只改文档不改表，Validate 不会发现；只改表不改文档，评审会发现——两边都要动。
//
// Since 的语义是"这条约束从哪个里程碑起有**可执行断言**"，不是"计划哪天做"。
// 能力没做完的一律往后排：宁可 PENDING，也不写一个假断言——那会让门禁说谎。
var table = []Contract{
	// ── C1–C17 继承自既有实现（examples/wg 的实证约束） ──────────────────
	{
		// 描边必须只落在边缘。"先填充后描边"在本引擎里的等价物是
		// "stroke-to-fill 之后内部仍为空"——若实现把描边当填充，内部会被涂满。
		ID: "C1", Title: "画刷顺序：先填充后描边，描边不污染内部", Owner: "①",
		Gate: "--render-matrix", Since: M1,
		Assert: assertBrushOrder,
	},
	{
		ID: "C2", Title: "呈现格式固定 Pf24bit，不做整窗 alpha 合成", Owner: "⓪",
		Gate: "--screenshot", Since: M1,
		Assert: assertPresentFormat,
	},
	{
		// 依赖 WndProcPolicy（吃掉 WM_ERASEBKGND），M2 随宿主消息接管落地。
		ID: "C3", Title: "ERASEBKGND 必须吃掉，否则闪烁", Owner: "⓪",
		Gate: "--schedtest", Since: M2,
		Assert: assertEraseEaten,
	},
	{
		// 静态烘焙依赖 RasterCache，M2 落地。
		ID: "C4", Title: "静态内容只烘焙一次，帧内不重画", Owner: "①",
		Gate: "--bench", Since: M2,
		Assert: assertBakeOnce,
	},
	{
		ID: "C5", Title: "AA 就地混色，不额外分配整帧 alpha 缓冲", Owner: "①",
		Gate: "--render-matrix", Since: M1,
		Assert: assertInPlaceAA,
	},
	{
		ID: "C6", Title: "文字在像素写回后回放，保留 ClearType", Owner: "①",
		Gate: "--screenshot", Since: M1,
		Assert: assertTextReplay,
	},
	{
		ID: "C7", Title: "光栅缓存归属唯一（RasterCache 只由 ① 持有）", Owner: "①",
		Gate: "--layers", Since: M2,
		Assert: assertCacheOwnership,
	},
	{
		ID: "C8", Title: "资源释放顺序固定，避免悬垂视图", Owner: "⓪①",
		Gate: "--resizetest", Since: M1,
		Assert: assertReleaseOrder,
	},
	{
		ID: "C9", Title: "S() 同缩放：DIP→像素只有一处换算", Owner: "⓪",
		Gate: "--render-matrix", Since: M1,
		Assert: assertSingleScale,
	},
	{
		ID: "C10", Title: "resize 去抖，连续拖动只重排一次", Owner: "⓪",
		Gate: "--resizetest", Since: M2,
		Assert: assertResizeDebounce,
	},
	{
		ID: "C11", Title: "单 outstanding 帧槽，槽未回填不申请下一帧", Owner: "②",
		Gate: "--schedtest", Since: M2,
		Assert: assertSingleOutstanding,
	},
	{
		ID: "C12", Title: "capture 自愈：丢失捕获必须补 Cancel", Owner: "③",
		Gate: "--inputtest", Since: M2,
		Assert: assertCaptureHeal,
	},
	{
		ID: "C13", Title: "帧=帧槽：一帧一次不透明提交", Owner: "②",
		Gate: "--schedtest", Since: M1,
		Assert: assertSingleCommit,
	},
	{
		ID: "C14", Title: "静止零帧：无动画无输入不出帧", Owner: "②",
		Gate: "--schedtest", Since: M1,
		Assert: assertIdleFrame,
	},
	{
		// 墙钟相位依赖 ticker 与动画系统，M2 落地。
		ID: "C15", Title: "墙钟相位：动画按墙钟积分，不按帧计数", Owner: "②",
		Gate: "--schedtest", Since: M2,
		Assert: assertWallClockPhase,
	},
	{
		ID: "C16", Title: "帧间隔用中位数统计，不用平均值", Owner: "②",
		Gate: "--bench", Since: M1,
		Assert: assertMedian,
	},
	{
		ID: "C17", Title: "单向依赖：lcl/unsafe 不得泄漏到上层", Owner: "全层",
		Gate: "checkdeps", Since: M0,
		Assert: checkDeps,
	},

	// ── C18–C21 画布体系（02-CANVAS.md） ────────────────────────────────
	{
		ID: "C18", Title: "帧内禁止资源拓扑变更，变更进 pending 队列", Owner: "⓪",
		Gate: "--selftest", Since: M1,
		Assert: assertNoTopologyChange,
	},
	{
		// Layer 树与 Reason 校验在 M2 随合成层落地。
		ID: "C19", Title: "Layer 必须声明 Reason，且稳态命中缓存", Owner: "①",
		Gate: "--layers", Since: M2,
		Assert: assertLayerReason,
	},
	{
		ID: "C20", Title: "子 Surface damage 必须上卷，并做遮挡剔除", Owner: "①",
		Gate: "--damage", Since: M1,
		Assert: assertDamageBubbles,
	},
	{
		// 池在 M2 引入（离屏层复用），届时才谈得上"归还池"的顺序。
		ID: "C21", Title: "后备存储释放顺序：视图 → 位图 → 池", Owner: "⓪①",
		Gate: "--resizetest", Since: M2,
		Assert: assertPoolReleaseOrder,
	},

	// ── C22–C25 渲染内核（03-ENGINE.md） ───────────────────────────────
	{
		ID: "C22", Title: "指令必须执行或显式降级，不得静默丢弃", Owner: "①",
		Gate: "--render-matrix", Since: M1,
		Assert: assertExplicitDegrade,
	},
	{
		ID: "C23", Title: "DisplayList 不可变，重录换 version", Owner: "①",
		Gate: "--selftest", Since: M1,
		Assert: assertDisplayListImmutable,
	},
	{
		ID: "C24", Title: "文字在像素写回后设备空间回放并裁剪", Owner: "①",
		Gate: "--screenshot", Since: M1,
		Assert: assertTextClipped,
	},
	{
		ID: "C25", Title: "两后端几何位置误差 ≤ 1px", Owner: "①",
		Gate: "--render-matrix", Since: M1,
		Assert: assertTwoBackendsAgree,
	},

	// ── C26–C28 交互（04-UI-INTERACTION.md） ───────────────────────────
	{
		ID: "C26", Title: "捕获指针必有 Up 或 Cancel，Cancel 回滚状态", Owner: "③",
		Gate: "--inputtest", Since: M2,
		Assert: assertCaptureTerminated,
	},
	{
		ID: "C27", Title: "单指针单胜者，败者 Reject 且不改状态", Owner: "③",
		Gate: "--gesturetest", Since: M2,
		Assert: assertSingleWinner,
	},
	{
		ID: "C28", Title: "文本与按键分离，控制字符不进 Text", Owner: "③⓪",
		Gate: "--keyselftest", Since: M3,
	},
}
