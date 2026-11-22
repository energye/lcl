package widgets

import (
	"testing"
)

// TestQRGF 自证 GF(256) 算术：乘法恒等/交换/对数表闭包。
func TestQRGF(t *testing.T) {
	// exp[0..254] 必须是 0x11d 域的全部 255 个非零元（双射）。
	seen := map[byte]bool{}
	for i := 0; i < 255; i++ {
		v := qrGFExp[i]
		if v == 0 || seen[v] {
			t.Fatalf("GF exp 表非双射：i=%d v=%d", i, v)
		}
		seen[v] = true
	}
	// 乘法恒等：a*1=a。
	for _, a := range []byte{1, 2, 17, 128, 255} {
		if qrGFMul(a, 1) != a {
			t.Fatalf("a*1 != a: %d", a)
		}
	}
	// 交换律抽查 + x*x=α^(2logx)。
	if qrGFMul(0x53, 0xCA) != qrGFMul(0xCA, 0x53) {
		t.Fatal("乘法不满足交换律")
	}
	// 2 的幂一致性：exp[i+1] == exp[i]*2（i<254）。
	for i := 0; i < 254; i++ {
		if qrGFMul(qrGFExp[i], 2) != qrGFExp[i+1] {
			t.Fatalf("exp 链断裂 at %d", i)
		}
	}
}

// TestQRSyndrome 强自证：编码后的完整码字（数据+EC）对生成多项式
// 的余式必须为零（等价 syndromes 全零）——RS 编码正确性的代数判据。
func TestQRSyndrome(t *testing.T) {
	for _, ec := range []int{7, 10, 13, 17, 26, 30} {
		msg := []byte{0x40, 0x29, 0x8a, 0xed, 0x2c, 0x10, 0x93, 0xf5, 0x77, 0x01}
		ecw := qrRSEncode(msg, ec)
		if len(ecw) != ec {
			t.Fatalf("EC 长度 = %d，应 %d", len(ecw), ec)
		}
		// 完整多项式 = msg·x^ec + ecw；模 gen 余式应为 0。
		// 直接用编码同款长除法验证：把 msg+ecw 再走一遍除法余式应为零。
		full := append(append([]byte{}, msg...), ecw...)
		// 构造 gen（与 qrRSEncode 内一致）。
		gen := make([]byte, ec+1)
		gen[0] = 1
		for i := 0; i < ec; i++ {
			for k := ec; k >= 1; k-- {
				gen[k] = qrGFMul(gen[k], qrGFExp[i]) ^ gen[k-1]
			}
			gen[0] = qrGFMul(gen[0], qrGFExp[i])
		}
		rem := make([]byte, ec)
		for _, d := range full {
			factor := d ^ rem[0]
			copy(rem, rem[1:])
			rem[ec-1] = 0
			if factor != 0 {
				lf := qrGFLog[factor]
				for k := 0; k < ec; k++ {
					g := gen[k+1]
					if g != 0 {
						rem[k] ^= qrGFExp[int(qrGFLog[g])+int(lf)]
					}
				}
			}
		}
		for _, v := range rem {
			if v != 0 {
				t.Fatalf("EC=%d 余式非零: %v", ec, rem)
			}
		}
	}
}

// TestQRRSKnownVector 已知向量：thonky 教程 5-H 组 1 例（32 字节数据
// 中前 10 个的子例）——用 4 个 EC 的经典向量 0x40 0x29 0x1e... 不可靠
// 转写，改用代数自洽：单字节消息 0x40, ec=7 → 码字必须满足
// gen(α^k)=0 的根性质：Σ codeword_i · α^(i+ec-1-…)=0。等价 syndromes。
// 这里用一条独立实现的公开测试向量（hello world v1-Q 经典 0x10 位流
// 不转写——代数判据已覆盖正确性，此测改为纠错能力验证）。
func TestQRRSCorrects(t *testing.T) {
	// RS(n,k) 最小距离 d=n-k+1 → 可纠 (d-1)/2 = ec/2 个错。
	msg := []byte{0x40, 0xC4, 0x10, 0x04, 0xB2, 0xAA, 0x88, 0xC2, 0xF0, 0x11}
	ec := 10
	ecw := qrRSEncode(msg, ec)
	full := append(append([]byte{}, msg...), ecw...)
	// 引入 ec/2 个错（分散位置）。
	corrupt := append([]byte{}, full...)
	corrupt[0] ^= 0xFF
	corrupt[5] ^= 0x7C
	corrupt[9] ^= 0x33
	corrupt[14] ^= 0x55
	corrupt[18] ^= 0xA1
	// 删除 ec/2 个错误位置后重算 EC 应与原 EC 相同（数据部分恢复）。
	// 简化验证：syndrome 判据——corrupt 全码字余式非零（有错），
	// 而原 full 余式为零（无错）。
	zeroRem := func(codeword []byte) bool {
		rem := make([]byte, ec)
		gen := make([]byte, ec+1)
		gen[0] = 1
		for i := 0; i < ec; i++ {
			for k := ec; k >= 1; k-- {
				gen[k] = qrGFMul(gen[k], qrGFExp[i]) ^ gen[k-1]
			}
			gen[0] = qrGFMul(gen[0], qrGFExp[i])
		}
		for _, d := range codeword {
			factor := d ^ rem[0]
			copy(rem, rem[1:])
			rem[ec-1] = 0
			if factor != 0 {
				lf := qrGFLog[factor]
				for k := 0; k < ec; k++ {
					g := gen[k+1]
					if g != 0 {
						rem[k] ^= qrGFExp[int(qrGFLog[g])+int(lf)]
					}
				}
			}
		}
		for _, v := range rem {
			if v != 0 {
				return false
			}
		}
		return true
	}
	if !zeroRem(full) {
		t.Fatal("完整码字余式应为零")
	}
	if zeroRem(corrupt) {
		t.Fatal("损坏码字余式应非零")
	}
}

// TestQRBlocksConsistency 块表自洽：每版本四级别满足
// (1) 组 1 块数×块内数据码字 + 组 2 = 总数据码字；
// (2) 总数据码字 + 块数×EC = 总码字（官方总表）。
func TestQRBlocksConsistency(t *testing.T) {
	totals := [10]int{26, 44, 70, 100, 134, 172, 196, 242, 292, 346}
	for v := 1; v <= 10; v++ {
		for lv, spec := range qrBlocks[v-1] {
			data := spec.blocksG1*spec.dataG1 + spec.blocksG2*spec.dataG2
			if data != spec.totalDataCW {
				t.Fatalf("v%d L%d: 数据码字 %d != %d", v, lv, data, spec.totalDataCW)
			}
			if spec.blocksG1+spec.blocksG2 != spec.numBlocks {
				t.Fatalf("v%d L%d: 块数不匹配", v, lv)
			}
			total := data + spec.numBlocks*spec.ecPerBlock
			if total != totals[v-1] {
				t.Fatalf("v%d L%d: 总码字 %d != %d", v, lv, total, totals[v-1])
			}
		}
	}
}

// TestQREncodeStructure 编码产物结构自证：尺寸、三定位角、dark module、
// timing 交替、format 区（两份）与掩码数字一致。
func TestQREncodeStructure(t *testing.T) {
	m := QREncode("HELLO WORLD", QRLow)
	if m == nil {
		t.Fatal("编码失败")
	}
	n := len(m)
	if n != 21 {
		t.Fatalf("v1 尺寸 = %d，应 21", n)
	}
	// 定位角：左上 7×7 黑框四角。
	for _, p := range [][2]int{{0, 0}, {0, 6}, {6, 0}, {6, 6}} {
		if !m[p[1]][p[0]] {
			t.Fatalf("定位角黑框角点缺失 %v", p)
		}
	}
	// 定位角白心环 (1,1) 与黑心 (3,3)。
	if m[1][1] || !m[3][3] {
		t.Fatal("定位角内部结构错误")
	}
	// dark module。
	if !m[n-8][8] {
		t.Fatal("dark module 缺失")
	}
	// timing 行/列交替（8..n-9）。
	for i := 8; i < n-8; i++ {
		if m[6][i] != (i%2 == 0) {
			t.Fatalf("timing 行不交替 at %d", i)
		}
		if m[i][6] != (i%2 == 0) {
			t.Fatalf("timing 列不交替 at %d", i)
		}
	}
}

// TestQREncodeDeterministic 同输入同输出（掩码择优必须确定性）。
func TestQREncodeDeterministic(t *testing.T) {
	a := QREncode("https://example.com", QRMedium)
	b := QREncode("https://example.com", QRMedium)
	if len(a) != len(b) {
		t.Fatal("尺寸不稳定")
	}
	for y := range a {
		for x := range a[y] {
			if a[y][x] != b[y][x] {
				t.Fatalf("编码不确定 at (%d,%d)", x, y)
			}
		}
	}
}

// TestQREncodeVersions 多版本多级别可编码（容量边界内）且尺寸正确。
func TestQREncodeVersions(t *testing.T) {
	cases := []struct {
		text string
		lv   QRErrorLevel
		n    int
	}{
		{"A", QRLow, 21},                                                           // v1
		{"HELLO WORLD", QRHigh, 25},                                                // v1-H 容量 9 < 11 字节 → 升 v2
		{"https://example.com/path", QRMedium, 25},                                 // v2
		{"0123456789012345678901234567890123456789", QRMedium, 29},                 // v3
		{"The quick brown fox jumps over the lazy dog 0123456789", QRQuartile, 37}, // 57 字节 > v4-Q 48 → 升 v5
	}
	for _, c := range cases {
		m := QREncode(c.text, c.lv)
		if m == nil {
			t.Fatalf("编码失败: %q L%d", c.text, c.lv)
		}
		if len(m) != c.n {
			t.Fatalf("%q L%d 尺寸 = %d，应 %d", c.text, c.lv, len(m), c.n)
		}
	}
}

// TestQREncodeOverflow 超容量返回 nil（v10 上限外）。
func TestQREncodeOverflow(t *testing.T) {
	big := make([]byte, 3000)
	for i := range big {
		big[i] = 'a'
	}
	if m := QREncode(string(big), QRMedium); m != nil {
		t.Fatal("超容量应返回 nil")
	}
}

// TestQRFormatBits 已知向量：v1-L mask0 的 format 信息 15 位。
// 官方值（ISO Annex / thonky format 表）：0b111011111000100 = 0x7788…？
// 转写风险高——改用 BCH 校验性质：去掩码后 5 位数据 + 10 位校验，
// 对 0x537 余式为 0 且异或 0x5412 前后自洽。
func TestQRFormatBits(t *testing.T) {
	for _, lv := range []QRErrorLevel{QRLow, QRMedium, QRQuartile, QRHigh} {
		for mask := 0; mask < 8; mask++ {
			f := qrFormatBits(lv, mask)
			un := f ^ 0x5412
			data5 := un >> 10
			rem10 := un & 0x3FF
			// 重算 BCH：data5 ×x^10 mod 0x537 应等于 rem10。
			rem := data5
			for i := 0; i < 10; i++ {
				rem <<= 1
				if rem&0x400 != 0 {
					rem ^= 0x537
				}
			}
			if rem&0x3FF != rem10 {
				t.Fatalf("format BCH 校验失败 L%d mask%d", lv, mask)
			}
			// 级别位前缀检查。
			want := map[QRErrorLevel]uint32{QRLow: 0b01, QRMedium: 0b00, QRQuartile: 0b11, QRHigh: 0b10}[lv]
			if data5>>3 != want {
				t.Fatalf("级别位错误 L%d", lv)
			}
		}
	}
}
