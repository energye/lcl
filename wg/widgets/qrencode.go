package widgets

// ── QR 标准编码引擎（第一百零五轮）───────────────────────────────
//
// 纯 Go 实现 ISO/IEC 18004 核心编码管线（byte 模式、版本 1-10、
// 四个纠错级别）：
//
//	数据 → [0100 模式指示 + 字符计数 + UTF-8 字节] 比特流
//	     → 终止符 0000 + 对齐到字节 + 补齐 pad 0xEC 0x11
//	     → 按块切分 → Reed-Solomon EC（GF(256)，0x11d）
//	     → 结构交织 → 矩阵布置（finder/align/timing/dark/format）
//	     → 8 掩码 penalty 评分择优
//
// 数值口径（官方表转写，用总码字自洽校验）：
//   - 块结构：thonky error-correction-table v1-10 全行（Total Data
//     Codewords 列的分组乘式即组 1/组 2 块数×块内码字）。
//   - 总码字/版本：26/44/70/100/134/172/196/242/292/346（官方总表，
//     = Σ(块数据码字)+Σ(EC 码字)，单测校验）。
//   - 对齐图案中心：v2-10 标准序列。
//   - format 信息：15 位 BCH(17,15) with gen 0x537，mask XOR 0x5412。
//
// 教训（外部库注入通道保留）：Matrix 显式注入仍优先——宿主可绕过
// 内置引擎（如需 v11+ 或 kanji 模式）。

import "fmt"

// QRErrorLevel 是纠错级别（antd errorLevel）。
type QRErrorLevel uint8

const (
	// QRLow 级别 L（约 7% 可恢复）。
	QRLow QRErrorLevel = iota
	// QRMedium 级别 M（约 15%，最常用——默认）。
	QRMedium
	// QRQuartile 级别 Q（约 25%）。
	QRQuartile
	// QRHigh 级别 H（约 30%）。
	QRHigh
)

// qrBlockSpec 是单个版本+级别的块结构。
type qrBlockSpec struct {
	numBlocks    int // 总块数
	ecPerBlock   int // 每块 EC 码字数
	blocksG1     int // 组 1 块数
	dataG1       int // 组 1 每块数据码字
	blocksG2     int // 组 2 块数（0 = 无组 2）
	dataG2       int // 组 2 每块数据码字
	totalDataCW  int // 总数据码字
	totalCodewds int // 总码字（数据+EC，官方总表）
}

// qrBlocks[version-1][level]：v1-10 四级别块结构（thonky 表转写）。
var qrBlocks = [10][4]qrBlockSpec{
	// v1（总 26）
	{{1, 7, 1, 19, 0, 0, 19, 26}, {1, 10, 1, 16, 0, 0, 16, 26}, {1, 13, 1, 13, 0, 0, 13, 26}, {1, 17, 1, 9, 0, 0, 9, 26}},
	// v2（总 44）
	{{1, 10, 1, 34, 0, 0, 34, 44}, {1, 16, 1, 28, 0, 0, 28, 44}, {1, 22, 1, 22, 0, 0, 22, 44}, {1, 28, 1, 16, 0, 0, 16, 44}},
	// v3（总 70）
	{{1, 15, 1, 55, 0, 0, 55, 70}, {1, 26, 1, 44, 0, 0, 44, 70}, {2, 18, 2, 17, 0, 0, 34, 70}, {2, 22, 2, 13, 0, 0, 26, 70}},
	// v4（总 100）
	{{1, 20, 1, 80, 0, 0, 80, 100}, {2, 18, 2, 32, 0, 0, 64, 100}, {2, 26, 2, 24, 0, 0, 48, 100}, {4, 16, 4, 9, 0, 0, 36, 100}},
	// v5（总 134）
	{{1, 26, 1, 108, 0, 0, 108, 134}, {2, 24, 2, 43, 0, 0, 86, 134}, {4, 18, 2, 15, 2, 16, 62, 134}, {4, 22, 2, 11, 2, 12, 46, 134}},
	// v6（总 172）
	{{2, 18, 2, 68, 0, 0, 136, 172}, {4, 16, 4, 27, 0, 0, 108, 172}, {4, 24, 4, 19, 0, 0, 76, 172}, {4, 28, 4, 15, 0, 0, 60, 172}},
	// v7（总 196）
	{{2, 20, 2, 78, 0, 0, 156, 196}, {4, 18, 4, 31, 0, 0, 124, 196}, {6, 18, 2, 14, 4, 15, 88, 196}, {5, 26, 4, 13, 1, 14, 66, 196}},
	// v8（总 242）
	{{2, 24, 2, 97, 0, 0, 194, 242}, {4, 22, 2, 38, 2, 39, 154, 242}, {6, 22, 4, 18, 2, 19, 110, 242}, {6, 26, 4, 14, 2, 15, 86, 242}},
	// v9（总 292）
	{{2, 30, 2, 116, 0, 0, 232, 292}, {5, 22, 3, 36, 2, 37, 182, 292}, {8, 20, 4, 16, 4, 17, 132, 292}, {8, 24, 4, 12, 4, 13, 100, 292}},
	// v10（总 346）
	{{4, 18, 2, 68, 2, 69, 274, 346}, {5, 26, 4, 43, 1, 44, 216, 346}, {8, 24, 6, 19, 2, 20, 154, 346}, {8, 28, 6, 15, 2, 16, 122, 346}},
}

// qrAlignCenters 是对齐图案中心坐标（v2-10 标准序列；v1 无对齐图案）。
var qrAlignCenters = [10][]int{
	nil, {6, 18}, {6, 22}, {6, 26}, {6, 30}, {6, 34}, {6, 22, 38}, {6, 24, 42}, {6, 26, 46}, {6, 28, 52},
}

// qrSize 返回版本对应的矩阵边长（21 + 4*(v-1)）。
func qrSize(version int) int { return 21 + 4*(version-1) }

// ── GF(256) 算术（0x11d = x^8+x^4+x^3+x^2+1）──────────────────────

var qrGFExp [512]byte
var qrGFLog [256]byte

func init() {
	x := 1
	for i := 0; i < 255; i++ {
		qrGFExp[i] = byte(x)
		qrGFLog[x] = byte(i)
		x <<= 1
		if x >= 256 {
			x ^= 0x11d
		}
	}
	for i := 255; i < 512; i++ {
		qrGFExp[i] = qrGFExp[i-255]
	}
}

func qrGFMul(a, b byte) byte {
	if a == 0 || b == 0 {
		return 0
	}
	return qrGFExp[int(qrGFLog[a])+int(qrGFLog[b])]
}

// qrRSEncode 计算一组的 Reed-Solomon EC 码字（多项式长除法）。
// gen = Π(x - α^i)，i=0..n-1；消息多项式乘 x^n 后对 gen 取余。
func qrRSEncode(data []byte, ecCount int) []byte {
	// 生成多项式：系数从最高次到最低次（gen[0] 为 x^(n) 系数=1）。
	gen := make([]byte, ecCount+1)
	gen[0] = 1
	for i := 0; i < ecCount; i++ {
		// gen *= (x + α^i)：新 gen[k] = gen[k]*α^i + gen[k-1]。
		for k := ecCount; k >= 1; k-- {
			gen[k] = qrGFMul(gen[k], qrGFExp[i]) ^ gen[k-1]
		}
		gen[0] = qrGFMul(gen[0], qrGFExp[i])
	}
	// 长除法：消息 ×x^n 逐步 XOR。
	rem := make([]byte, ecCount)
	for _, d := range data {
		factor := d ^ rem[0]
		copy(rem, rem[1:])
		rem[ecCount-1] = 0
		if factor != 0 {
			lf := qrGFLog[factor]
			for k := 0; k < ecCount; k++ {
				g := gen[k+1]
				if g != 0 {
					rem[k] ^= qrGFExp[int(qrGFLog[g])+int(lf)]
				}
			}
		}
	}
	return rem
}

// ── 比特流组装 ────────────────────────────────────────────────────

type qrBitBuf struct {
	bits []bool
}

func (b *qrBitBuf) append(v uint32, n int) {
	for i := n - 1; i >= 0; i-- {
		b.bits = append(b.bits, (v>>uint(i))&1 == 1)
	}
}

func (b *qrBitBuf) appendBytes(data []byte) {
	for _, by := range data {
		b.append(uint32(by), 8)
	}
}

// byteCountBits 返回 byte 模式字符计数指示位数（v1-9=8，v10-26=16）。
func qrByteCountBits(version int) int {
	if version <= 9 {
		return 8
	}
	return 16
}

// qrDataBits 构建 byte 模式数据比特流（0100 + 计数 + 数据）。
func qrDataBits(data []byte, version int) qrBitBuf {
	var b qrBitBuf
	b.append(0b0100, 4) // byte 模式指示
	b.append(uint32(len(data)), qrByteCountBits(version))
	b.appendBytes(data)
	// 终止符（最多 4 个 0）、对齐到字节边界。
	for i := 0; i < 4 && len(b.bits)%8 != 0; i++ {
		b.bits = append(b.bits, false)
	}
	for len(b.bits)%8 != 0 {
		b.bits = append(b.bits, false)
	}
	return b
}

// qrInterleave 按 QR 标准结构交织数据+EC 码字（组内逐块取第 i 字节，
// 数据在前 EC 在后；尾块较长时组 2 的多余字节接在各组字节之后）。
func qrInterleave(dataBlocks, ecBlocks [][]byte, totalData int) []byte {
	var out []byte
	maxData := 0
	for _, b := range dataBlocks {
		if len(b) > maxData {
			maxData = len(b)
		}
	}
	// 数据列优先：先所有块的第 0 字节、第 1 字节……短块跳过。
	for i := 0; i < maxData; i++ {
		for _, b := range dataBlocks {
			if i < len(b) {
				out = append(out, b[i])
			}
		}
	}
	_ = totalData
	maxEC := 0
	for _, b := range ecBlocks {
		if len(b) > maxEC {
			maxEC = len(b)
		}
	}
	for i := 0; i < maxEC; i++ {
		for _, b := range ecBlocks {
			if i < len(b) {
				out = append(out, b[i])
			}
		}
	}
	return out
}

// ── 矩阵布置 ──────────────────────────────────────────────────────

type qrMatrix struct {
	n        int
	mod      [][]bool // true=黑
	reserved [][]bool // 功能模块占位（布置数据时跳过）
}

func newQRMatrix(n int) *qrMatrix {
	m := &qrMatrix{n: n,
		mod:      make([][]bool, n),
		reserved: make([][]bool, n)}
	for i := range m.mod {
		m.mod[i] = make([]bool, n)
		m.reserved[i] = make([]bool, n)
	}
	return m
}

func (m *qrMatrix) set(x, y int, v bool) {
	if x >= 0 && x < m.n && y >= 0 && y < m.n {
		m.mod[y][x] = v
		m.reserved[y][x] = true
	}
}

// placeFinder 画 7×7 定位角+周围白环。
func (m *qrMatrix) placeFinder(cx, cy int) {
	for dy := -1; dy <= 7; dy++ {
		for dx := -1; dx <= 7; dx++ {
			x, y := cx+dx, cy+dy
			if x < 0 || x >= m.n || y < 0 || y >= m.n {
				continue
			}
			inside := dx >= 0 && dx <= 6 && dy >= 0 && dy <= 6
			edge := dx == 0 || dx == 6 || dy == 0 || dy == 6
			heart := dx >= 2 && dx <= 4 && dy >= 2 && dy <= 4
			m.set(x, y, inside && (edge || heart))
		}
	}
}

// placeAlign 画 5×5 对齐图案（黑框+白环+黑心）。
func (m *qrMatrix) placeAlign(cx, cy int) {
	for dy := -2; dy <= 2; dy++ {
		for dx := -2; dx <= 2; dx++ {
			x, y := cx+dx, cy+dy
			if x < 0 || x >= m.n || y < 0 || y >= m.n {
				continue
			}
			edge := dx == -2 || dx == 2 || dy == -2 || dy == 2
			m.set(x, y, edge || (dx == 0 && dy == 0))
		}
	}
}

// placeFunctionPatterns 画全部功能模块（finder/align/timing/dark）。
func (m *qrMatrix) placeFunctionPatterns(version int) {
	n := m.n
	// 三定位角。
	m.placeFinder(0, 0)
	m.placeFinder(n-7, 0)
	m.placeFinder(0, n-7)
	// 对齐图案（跳过与定位角重叠的左上、其余按中心坐标画）。
	for _, cy := range qrAlignCenters[version-1] {
		for _, cx := range qrAlignCenters[version-1] {
			overlapTopLeft := cx == 6 && cy == 6
			overlapTR := cx == 6 && cy+2 >= n-8 && cy <= 8
			overlapBL := cy == 6 && cx+2 >= n-8 && cx <= 8
			if overlapTopLeft || overlapTR || overlapBL {
				continue
			}
			m.placeAlign(cx, cy)
		}
	}
	// timing：第 6 行/列，黑白交替（避开 finder 区）。
	for i := 8; i < n-8; i++ {
		v := i%2 == 0
		m.set(6, i, v)
		m.set(i, 6, v)
	}
	// dark module（固定黑，format 信息旁）。
	m.set(8, n-8, true)
}

// placeFormatInfo 画 format 信息（两份：左上拆 L 形、另两角整条）。
// formatBits 为 15 位（MSB 先），布点按 ISO 标准。
func (m *qrMatrix) placeFormatInfo(bits uint32) {
	n := m.n
	get := func(i int) bool { return (bits>>uint(14-i))&1 == 1 }
	// 左上：位置 0-5 绕 finder 上方，6-8 timing 旁，9-14 左列。
	for i := 0; i <= 5; i++ {
		m.set(8, i, get(i))
	}
	m.set(8, 7, get(6))
	m.set(8, 8, get(7))
	m.set(7, 8, get(8))
	for i := 9; i <= 14; i++ {
		m.set(14-i, 8, get(i))
	}
	// 右上：finder 下方一列（0-7）。
	for i := 0; i <= 7; i++ {
		m.set(n-1-i, 8, get(i))
	}
	// 左下：finder 右侧一行（8-14）。
	for i := 8; i <= 14; i++ {
		m.set(8, n-15+i, get(i))
	}
	// dark module 已在 placeFunctionPatterns 画过（不被 format 覆盖）。
}

// qrFormatBits 计算 15 位 format 信息（2 位级别+3 位掩码 → BCH(15,5)）。
// 生成多项式 0x537（10100110111），异或掩码 0x5412。
func qrFormatBits(level QRErrorLevel, mask int) uint32 {
	var levelBits uint32
	switch level {
	case QRLow:
		levelBits = 0b01
	case QRMedium:
		levelBits = 0b00
	case QRQuartile:
		levelBits = 0b11
	case QRHigh:
		levelBits = 0b10
	}
	data := levelBits<<3 | uint32(mask)
	// BCH：data(5 位) ×x^10 对 0x537 取余，拼在 data 后。
	rem := data
	for i := 0; i < 10; i++ {
		rem <<= 1
		if rem&0x400 != 0 {
			rem ^= 0x537
		}
	}
	fmtBits := (data<<10 | rem&0x3FF) ^ 0x5412
	return fmtBits & 0x7FFF
}

// placeData 沿两列宽蛇形布置数据比特（跳过功能模块）。
func (m *qrMatrix) placeData(codewords []byte) {
	bitAt := func(idx int) bool {
		byteIdx, bitIdx := idx/8, idx%8
		if byteIdx >= len(codewords) {
			return false
		}
		return codewords[byteIdx]>>(7-uint(bitIdx))&1 == 1
	}
	total := len(codewords) * 8
	idx := 0
	col := m.n - 1
	upward := true
	for col > 0 {
		// 跳过 timing 列（col==6 时向左多退一步，两列组是 (col,col-1)）。
		if col == 6 {
			col--
		}
		rows := make([]int, m.n)
		if upward {
			for i := 0; i < m.n; i++ {
				rows[i] = m.n - 1 - i
			}
		} else {
			for i := 0; i < m.n; i++ {
				rows[i] = i
			}
		}
		for _, y := range rows {
			for _, x := range []int{col, col - 1} {
				if m.reserved[y][x] {
					continue
				}
				var v bool
				if idx < total {
					v = bitAt(idx)
					idx++
				}
				m.mod[y][x] = v
			}
		}
		upward = !upward
		col -= 2
	}
}

// applyMask 应用掩码函数（ISO 8 条件，只翻非功能模块）。
func (m *qrMatrix) applyMask(mask int) {
	cond := func(x, y int) bool {
		switch mask {
		case 0:
			return (x+y)%2 == 0
		case 1:
			return y%2 == 0
		case 2:
			return x%3 == 0
		case 3:
			return (x+y)%3 == 0
		case 4:
			return (y/2+x/3)%2 == 0
		case 5:
			return (x*y)%2+(x*y)%3 == 0
		case 6:
			return ((x*y)%2+(x*y)%3)%2 == 0
		default:
			return ((x+y)%2+(x*y)%3)%2 == 0
		}
	}
	for y := 0; y < m.n; y++ {
		for x := 0; x < m.n; x++ {
			if !m.reserved[y][x] && cond(x, y) {
				m.mod[y][x] = !m.mod[y][x]
			}
		}
	}
}

// penaltyScore 计算掩码罚分（ISO 四规则之和，越高越差）。
func penaltyScore(m *qrMatrix) int {
	n, score := m.n, 0
	// 规则 1：同色连续 ≥5 每段 +3+(len-5)。
	for y := 0; y < n; y++ {
		run := 1
		for x := 1; x < n; x++ {
			if m.mod[y][x] == m.mod[y][x-1] {
				run++
			} else {
				if run >= 5 {
					score += 3 + run - 5
				}
				run = 1
			}
		}
		if run >= 5 {
			score += 3 + run - 5
		}
	}
	for x := 0; x < n; x++ {
		run := 1
		for y := 1; y < n; y++ {
			if m.mod[y][x] == m.mod[y-1][x] {
				run++
			} else {
				if run >= 5 {
					score += 3 + run - 5
				}
				run = 1
			}
		}
		if run >= 5 {
			score += 3 + run - 5
		}
	}
	// 规则 2：2×2 同色块每处 +3。
	for y := 0; y < n-1; y++ {
		for x := 0; x < n-1; x++ {
			v := m.mod[y][x]
			if v == m.mod[y][x+1] && v == m.mod[y+1][x] && v == m.mod[y+1][x+1] {
				score += 3
			}
		}
	}
	// 规则 3：1011101 带四白边模式（简化：行/列扫描 11 位窗口）。
	pattern := [11]bool{true, false, true, true, true, false, true, false, false, false, false}
	match := func(get func(int, int) bool) int {
		cnt := 0
		for y := 0; y < n; y++ {
			for x := 0; x+11 <= n; x++ {
				ok, inv := true, true
				for i := 0; i < 11; i++ {
					v := get(x+i, y)
					if v != pattern[i] {
						ok = false
					}
					if v != pattern[10-i] {
						inv = false
					}
				}
				if ok || inv {
					cnt++
				}
			}
		}
		return cnt
	}
	row := func(x, y int) bool { return m.mod[y][x] }
	col := func(x, y int) bool { return m.mod[y][x] }
	score += (match(row) + match(col)) * 40
	// 规则 4：黑白比例偏离 50%（每 5% +10）。
	dark := 0
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if m.mod[y][x] {
				dark++
			}
		}
	}
	pct := dark * 100 / (n * n)
	if d := pct - 50; d < 0 {
		d = -d
		score += (d / 5) * 10
	}
	return score
}

// QREncode 生成标准 QR 矩阵（byte 模式，v1-10 四级别）。
// 返回 nil 表示数据超出 v10 该级别容量。
func QREncode(value string, level QRErrorLevel) [][]bool {
	data := []byte(value)
	// 选最小可容纳版本。
	version := 0
	for v := 1; v <= 10; v++ {
		spec := qrBlocks[v-1][level]
		// 容量 = 总数据码字 - 模式 4 位 - 计数位（向上取整到字节）。
		overhead := (4 + qrByteCountBits(v) + 7) / 8
		if len(data)+overhead <= spec.totalDataCW {
			version = v
			break
		}
	}
	if version == 0 {
		return nil
	}
	spec := qrBlocks[version-1][level]

	// 1) 数据比特流 + pad 补齐。
	bits := qrDataBits(data, version)
	for len(bits.bits)/8 < spec.totalDataCW {
		bits.append(0xEC, 8)
		if len(bits.bits)/8 >= spec.totalDataCW {
			break
		}
		bits.append(0x11, 8)
	}
	dataBytes := make([]byte, 0, spec.totalDataCW)
	for i := 0; i+8 <= len(bits.bits); i += 8 {
		var v byte
		for k := 0; k < 8; k++ {
			if bits.bits[i+k] {
				v |= 1 << uint(7-k)
			}
		}
		dataBytes = append(dataBytes, v)
	}

	// 2) 切块 + RS EC。
	var dataBlocks, ecBlocks [][]byte
	g1, d1, g2, d2 := spec.blocksG1, spec.dataG1, spec.blocksG2, spec.dataG2
	pos := 0
	for i := 0; i < g1; i++ {
		blk := dataBytes[pos : pos+d1]
		pos += d1
		dataBlocks = append(dataBlocks, blk)
		ecBlocks = append(ecBlocks, qrRSEncode(blk, spec.ecPerBlock))
	}
	for i := 0; i < g2; i++ {
		blk := dataBytes[pos : pos+d2]
		pos += d2
		dataBlocks = append(dataBlocks, blk)
		ecBlocks = append(ecBlocks, qrRSEncode(blk, spec.ecPerBlock))
	}

	// 3) 交织。
	_ = qrInterleave(dataBlocks, ecBlocks, spec.totalDataCW)
	codewords := qrInterleave(dataBlocks, ecBlocks, spec.totalDataCW)

	// 4-6) 矩阵 + 8 掩码择优。
	n := qrSize(version)
	best, bestScore := (*qrMatrix)(nil), -1
	for mask := 0; mask < 8; mask++ {
		m := newQRMatrix(n)
		m.placeFunctionPatterns(version)
		m.placeData(codewords)
		m.applyMask(mask)
		m.placeFormatInfo(qrFormatBits(level, mask))
		s := penaltyScore(m)
		if best == nil || s < bestScore {
			best, bestScore = m, s
		}
	}
	return best.mod
}

var _ = fmt.Sprintf // 占位引用（fmt 在其他构建路径使用）
