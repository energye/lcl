package host

import "github.com/energye/lcl/wg/input"

// ── 虚拟键 → 语义键码 ───────────────────────────────────────────────
//
// 只做一件事：把平台的虚拟键翻译成 input.KeyCode。
// 上层（快捷键、文本编辑、焦点）只认语义键码，因此换平台只需要换这张表。
//
// 为什么按**物理位置**而不是字符映射：中文输入法激活时 VK_A 对应的字符
// 可能是"a"也可能是任意候选字，但 Ctrl+S 永远是"左下第二排第三个键"。

// VK 常量（只列本表用到的）。
const (
	vkBack     = 0x08
	vkTab      = 0x09
	vkReturn   = 0x0D
	vkShift    = 0x10
	vkControl  = 0x11
	vkMenu     = 0x12 // Alt
	vkPause    = 0x13
	vkCaps     = 0x14
	vkEscape   = 0x1B
	vkSpace    = 0x20
	vkPrior    = 0x21 // PageUp
	vkNext     = 0x22 // PageDown
	vkEnd      = 0x23
	vkHome     = 0x24
	vkLeft     = 0x25
	vkUp       = 0x26
	vkRight    = 0x27
	vkDown     = 0x28
	vkInsert   = 0x2D
	vkDelete   = 0x2E
	vk0        = 0x30
	vk9        = 0x39
	vkA        = 0x41
	vkZ        = 0x5A
	vkLWIN     = 0x5B
	vkRWIN     = 0x5C
	vkNum0     = 0x60
	vkNum9     = 0x69
	vkNumMul   = 0x6A
	vkNumAdd   = 0x6B
	vkNumSub   = 0x6D
	vkNumDec   = 0x6E
	vkNumDiv   = 0x6F
	vkF1       = 0x70
	vkF12      = 0x7B
	vkOEM1     = 0xBA // ;
	vkOEMPlus  = 0xBB // =
	vkOEMCom   = 0xBC // ,
	vkOEMMinus = 0xBD // -
	vkOEMPer   = 0xBE // .
	vkOEMQ     = 0xBF // /
	vkOEMTilde = 0xC0 // `
	vkOEMOpen  = 0xDB // [
	vkOEMPipe  = 0xDC // \
	vkOEMClose = 0xDD // ]
	vkOEMQuote = 0xDE // '
)

// VKToCode 把虚拟键映射成语义键码。
func VKToCode(vk uint16) input.KeyCode {
	switch {
	case vk >= vkA && vk <= vkZ:
		return input.CodeA + input.KeyCode(vk-vkA)
	case vk >= vk0 && vk <= vk9:
		return input.Code0 + input.KeyCode(vk-vk0)
	case vk >= vkF1 && vk <= vkF12:
		return input.CodeF1 + input.KeyCode(vk-vkF1)
	case vk >= vkNum0 && vk <= vkNum9:
		return input.Code0 + input.KeyCode(vk-vkNum0)
	}
	switch vk {
	case vkBack:
		return input.CodeBackspace
	case vkTab:
		return input.CodeTab
	case vkReturn:
		return input.CodeEnter
	case vkEscape:
		return input.CodeEscape
	case vkSpace:
		return input.CodeSpace
	case vkDelete:
		return input.CodeDelete
	case vkInsert:
		return input.CodeInsert
	case vkLeft:
		return input.CodeLeft
	case vkRight:
		return input.CodeRight
	case vkUp:
		return input.CodeUp
	case vkDown:
		return input.CodeDown
	case vkHome:
		return input.CodeHome
	case vkEnd:
		return input.CodeEnd
	case vkPrior:
		return input.CodePageUp
	case vkNext:
		return input.CodePageDown
	case vkShift:
		return input.CodeShift
	case vkControl:
		return input.CodeControl
	case vkMenu:
		return input.CodeAlt
	case vkLWIN, vkRWIN:
		return input.CodeMeta
	case vkCaps:
		return input.CodeCapsLock
	case vkOEMMinus:
		return input.CodeMinus
	case vkOEMPlus:
		return input.CodeEqual
	case vkOEMOpen:
		return input.CodeBracketLeft
	case vkOEMClose:
		return input.CodeBracketRight
	case vkOEMPipe:
		return input.CodeBackslash
	case vkOEM1:
		return input.CodeSemicolon
	case vkOEMQuote:
		return input.CodeQuote
	case vkOEMCom:
		return input.CodeComma
	case vkOEMPer:
		return input.CodePeriod
	case vkOEMQ:
		return input.CodeSlash
	case vkOEMTilde:
		return input.CodeBackquote
	case vkNumAdd:
		return input.CodeNumpadAdd
	case vkNumSub:
		return input.CodeNumpadSub
	case vkNumMul:
		return input.CodeNumpadMul
	case vkNumDiv:
		return input.CodeNumpadDiv
	case vkNumDec:
		return input.CodeNumpadDecimal
	}
	return input.CodeUnknown
}

// ── 修饰键状态 ──────────────────────────────────────────────────────

// VK 修饰键位（与 GetKeyState 的高位约定一致，这里只需要"是否按下"）。
const (
	vkStateDown = 0x8000
)

// modsFrom 由三个"是否按下"合成修饰集。
func modsFrom(shift, ctrl, alt bool) input.Modifiers {
	var m input.Modifiers
	if shift {
		m |= input.ModShift
	}
	if ctrl {
		m |= input.ModControl
	}
	if alt {
		m |= input.ModAlt
	}
	return m
}
