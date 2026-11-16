package contracts

import (
	"strings"
	"testing"
)

func TestTableComplete(t *testing.T) {
	list := All()
	if len(list) != 28 {
		t.Fatalf("契约数 = %d，应为 28", len(list))
	}
	// 按编号排序：C2 必须在 C10 之前
	idx := map[string]int{}
	for i, c := range list {
		idx[c.ID] = i
	}
	if idx["C2"] > idx["C10"] {
		t.Fatal("编号未按数值排序（C2 应在 C10 前）")
	}
	for i := 1; i <= 28; i++ {
		id := "C" + itoa(i)
		if _, ok := idx[id]; !ok {
			t.Fatalf("缺少 %s", id)
		}
	}
}

func TestValidate(t *testing.T) {
	if err := Validate(); err != nil {
		t.Fatalf("契约表自检失败: %v", err)
	}
}

func TestGet(t *testing.T) {
	c, ok := Get("C17")
	if !ok || c.ID != "C17" {
		t.Fatalf("Get(C17) = %v %v", c, ok)
	}
	if c, ok := Get("c17"); !ok || c.ID != "C17" {
		t.Fatal("Get 应大小写不敏感")
	}
	if _, ok := Get("C99"); ok {
		t.Fatal("不存在的编号应返回 false")
	}
}

func TestCheckAllAtM0(t *testing.T) {
	rs := CheckAll()
	if len(rs) != 28 {
		t.Fatalf("结果数 = %d", len(rs))
	}
	s := Summarize(rs)
	if !s.OK() {
		for _, r := range rs {
			if r.Status == Fail || r.Status == Missing {
				t.Errorf("%s %s: %s", r.ID, r.Status, r.Detail)
			}
		}
		t.Fatalf("M0 门禁应为绿: %s", s)
	}
	// C17 在 M0 必须真的跑过断言（不是 Pending）
	for _, r := range rs {
		if r.ID == "C17" {
			if r.Status != Pass {
				t.Fatalf("C17 状态 = %s: %s", r.Status, r.Detail)
			}
			return
		}
	}
	t.Fatal("未找到 C17")
}

// TestM1GateGreen 是 M1 的门禁：所有 Since<=M1 的契约都必须有断言且通过。
func TestM1GateGreen(t *testing.T) {
	rs := Run(M1)
	s := Summarize(rs)
	if !s.OK() {
		for _, r := range rs {
			if r.Status == Fail || r.Status == Missing {
				t.Errorf("%s %s: %s", r.ID, r.Status, r.Detail)
			}
		}
		t.Fatalf("M1 门禁应为绿: %s", s)
	}
	// M1 到期 16 条（C1/C2/C5/C6/C8/C9/C13/C14/C16/C17/C18/C20/C22/C23/C24/C25），
	// 其余 12 条按能力分期排到 M2/M3。
	if s.Pass < 16 {
		t.Fatalf("M1 应有至少 16 条契约在跑，实际 %d", s.Pass)
	}
}

// TestMissingGateIsRed 保护契约机制本身：把门禁提到一条"没有断言"的
// 未来契约时，必须报 MISSING 而不是静默通过——否则"契约与断言同生共死"
// 就形同虚设。
//
// 不复用某个真实契约当样本：C1–C28 在 M3 已全部到期，若把其中一条
// 故意留空来测门禁，就等于给引擎留了一个永远不写断言的死角。
// 这里在**测试内**临时插入一条 Since=M5 的契约，跑完立刻还原。
func TestMissingGateIsRed(t *testing.T) {
	saved := table
	defer func() { table = saved }()
	table = append(append([]Contract{}, saved...), Contract{
		ID: "C99", Title: "测试用未来契约", Owner: "③", Gate: "--future", Since: M5,
	})
	s := Summarize(Run(M5))
	if s.Missing != 1 {
		t.Fatalf("M5 门禁下应有 1 个 MISSING，实际 %s", s)
	}
	if s.OK() {
		t.Fatal("含无断言契约的门禁不应为绿")
	}
}

// TestM2GateGreen 确认 M2 的 11 条到期契约已全部落地。
func TestM2GateGreen(t *testing.T) {
	s := Summarize(Run(M2))
	if s.Missing != 0 || s.Fail != 0 {
		t.Fatalf("M2 门禁应为绿，实际 %s", s)
	}
}

// TestM3GateGreen 确认全部 28 条契约在 M3 门禁下为绿。
func TestM3GateGreen(t *testing.T) {
	s := Summarize(Run(M3))
	if s.Missing != 0 || s.Fail != 0 || s.Pass != 28 {
		t.Fatalf("M3 门禁应为 28 条全绿，实际 %s", s)
	}
}

func TestRunDoesNotMutateCurrent(t *testing.T) {
	before := Current
	_ = Run(M5)
	if Current != before {
		t.Fatalf("Run 修改了全局 Current: %v → %v", before, Current)
	}
}

func TestWGRootResolves(t *testing.T) {
	root := WGRoot()
	if root == "" || !strings.HasSuffix(strings.ReplaceAll(root, "\\", "/"), "/wg") {
		t.Fatalf("WGRoot = %q，应以 /wg 结尾", root)
	}
}

func TestIDNum(t *testing.T) {
	if idNum("C17") != 17 || idNum("C7") != 7 || idNum("C") != 0 {
		t.Fatal("idNum 解析错误")
	}
	if !lessID("C2", "C10") || lessID("C10", "C2") {
		t.Fatal("lessID 排序错误")
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
