//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

import "github.com/energye/lcl/types/colors"

// PByte = ^Byte = *uintptr
type PByte = uintptr
type HDWP = uintptr
type HMODULE = uintptr
type HINST = uintptr
type LPCWSTR = uintptr
type HRGN = uintptr
type LRESULT = uintptr
type HRSRC = uintptr
type HGLOBAL = uintptr
type TFNWndEnumProc = uintptr
type TXId = uint64
type ATOM = uint16
type TAtom = uint16
type SIZE_T = uintptr
type DWORD_PTR = uintptr
type TModalResult = int32
type THelpEventData = uintptr
type TTabOrder = int16
type PFNLVCOMPARE = uintptr
type PFNTVCOMPARE = uintptr
type Byte = uint8
type TSpacingSize = int32
type TClass = uintptr
type TThreadID = uintptr

// TClipboardFormat = TPredefinedClipboardFormat
type TClipboardFormat = int32
type Single = float32
type Char = uint16   // Char Unicode 主要用于keymap, 参见types/keys包
type PChar = uintptr // ^Char
type PPChar = PChar  // ^PChar
type PPAnsiChar = PPChar
type PWideChar = uintptr    // ^WideChar
type PPWideChar = PWideChar // ^PWideChar
type AnsiChar = Char
type HBitmap = uintptr
type TPenPattern = LongWord
type Integer = int32
type LongInt = int32
type LongPtr = uintptr
type LongWord = uint32
type NativeUInt = uint32
type Cardinal = uint32
type HWND = uintptr
type WParam = uintptr
type LParam = uintptr
type HDC = uintptr
type Pointer = uintptr
type QWord = uintptr
type Word = uint16
type HGDIOBJ = uintptr
type HPALETTE = uintptr
type LResult = uintptr
type HMENU = uintptr
type PtrInt = uintptr
type HBITMAP = uintptr
type HICON = uintptr
type HMONITOR = uintptr
type HBRUSH = uintptr
type HPEN = uintptr
type HKEY = uintptr
type Variant = uintptr
type TGraphicsColor = int32
type SmallInt = int16
type HFONT = uintptr
type HRESULT = uintptr
type SizeInt = int
type DWORD = Cardinal
type PDWORD = uintptr // PDWORD = ^DWORD
type ACCESS_MASK = DWORD
type REGSAM = ACCESS_MASK
type Boolean = bool
type UINT = LongWord
type ULONG_PTR = QWord
type THandle = QWord
type HANDLE = THandle
type COLORREF = Cardinal
type TColorRef = COLORREF
type PtrUInt = QWord
type HCURSOR = HICON
type TFarProc = Pointer
type LongBool int32
type BOOL = LongBool
type LANGID = Word

// type TCefStringList = Pointer
type TCriticalSection = PtrUInt
type HOOK = QWord
type TCopyMode = LongInt
type TImageIndex = Integer
type TFontDataName = string
type HFont = HANDLE
type TFPResourceHandle = PtrUInt
type TFontCharSet = byte
type HIMAGELIST = HANDLE
type TLCLHandle = PtrUInt

// Currency -922337203685477.5808到922337203685477.5807
// 10000倍
type Currency = int64

// TFPJPEGCompressionQuality = 1..100;
type TFPJPEGCompressionQuality = uint8
type TJPEGQualityRange = TFPJPEGCompressionQuality

// TDateTime Double
type TDateTime = float64

// TDate TDateTime
type TDate = TDateTime

// TTime TDateTime
type TTime = TDateTime

type TDockZoneClass = TClass
type TCollectionItemClass = TClass
type TWinControlClass = TClass
type TFPCustomImageClass = TClass
type TGraphicClass = TClass
type TStatusPanelClass = TClass
type THeaderSectionClass = TClass
type TListItemClass = TClass
type TTreeNodeClass = TClass
type TAVLTreeNodeClass = TClass
type TFPImageBitmapClass = TClass
type TSynEditMarkupClass = TClass
type TSynEditStringTabExpanderClass = TClass
type TSynEditStringsLinkedClass = TClass
type TSynGutterPartBaseClass = TClass

// TSynEdit
type TSynEditorMouseCommand = Word
type TSynEditorMouseCommandOpt = Word
type TLinePos = Integer
type TLineIdx = Integer
type IntPos = Integer
type IntIdx = Integer

func (m LongBool) Bool() bool {
	return m != 0
}

// TSet Pascal集合类型 set of xxx
type TSet uint32

type TColor = colors.TColor

// TUTF8Char
//
//	UTF-8 character is at most 6 bytes plus a #0
type TUTF8Char struct {
	Len     byte
	Content [7]byte
}

// TBrushPattern
//
//	32 = PatternBitCount = SizeOf(uint32) * 8;
//	[32]uint32 = array[0..PatternBitCount-1] of TPenPattern
type TBrushPattern = uintptr

// TOverlay = 0..14; // windows limitation
type TOverlay uint8

type TPoint struct {
	X int32
	Y int32
}

func Point(x, y int32) TPoint {
	return TPoint{X: x, Y: y}
}

type TRect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type TSize struct {
	Cx int32
	Cy int32
}

func (u *TUTF8Char) ToString() string {
	if u.Len > 0 && u.Len < 7 {
		return string(u.Content[0:u.Len])
	}
	return ""
}

func (u *TUTF8Char) SetString(str string) {
	if str != "" {
		bs := []byte(str)
		u.Len = byte(len(bs))
		if u.Len > 6 {
			u.Len = 6
		}
		copy(u.Content[:], bs[:u.Len])
	}
}

//----------------------------------------------------------------------------------------------------------------------
// -- TRect

func Rect(left, top, right, bottom int32) TRect {
	return TRect{Left: left, Top: top, Right: right, Bottom: bottom}
}

func (r TRect) PtInRect(P TPoint) bool {
	return P.X >= r.Left && P.X < r.Right && P.Y >= r.Top && P.Y < r.Bottom
}

func (r TRect) Width() int32 {
	return r.Right - r.Left
}

func (r *TRect) SetWidth(val int32) {
	r.Right = r.Left + val
}

func (r TRect) Height() int32 {
	return r.Bottom - r.Top
}

func (r *TRect) SetHeight(val int32) {
	r.Bottom = r.Top + val
}

func (r TRect) IsEmpty() bool {
	return r.Right <= r.Left || r.Bottom <= r.Top
}

func (r *TRect) Empty() {
	r.Left = 0
	r.Top = 0
	r.Right = 0
	r.Bottom = 0
}

func (r TRect) Size() TSize {
	return TSize{r.Width(), r.Height()}
}

func (r *TRect) SetSize(w, h int32) {
	r.SetWidth(w)
	r.SetHeight(h)
}

func (r *TRect) Inflate(dx, dy int32) {
	r.Left += -dx
	r.Top += -dy
	r.Right += dx
	r.Bottom += dy
}

func (r TRect) Contains(aR TRect) bool {
	return r.Left <= aR.Left && r.Right >= aR.Right && r.Top <= aR.Top && r.Bottom >= aR.Bottom
}

func (r TRect) IntersectsWith(aR TRect) bool {
	return r.Left < aR.Right && r.Right > aR.Left && r.Top < aR.Bottom && r.Bottom > aR.Top
}

func (r TRect) CenterPoint() (ret TPoint) {
	ret.X = (r.Right-r.Left)/2 + r.Left
	ret.Y = (r.Bottom-r.Top)/2 + r.Top
	return
}

func (r *TRect) Scale(val float64) {
	r.Left = int32(float64(r.Left) * val)
	r.Top = int32(float64(r.Top) * val)
	r.Right = int32(float64(r.Right) * val)
	r.Bottom = int32(float64(r.Bottom) * val)
}

func (r *TRect) Scale2(val int) {
	r.Scale(float64(val))
}

// -- TPoint

func NewPoint(x, y int32) TPoint {
	return TPoint{X: x, Y: y}
}

func (p TPoint) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

func (p *TPoint) Offset(dx, dy int32) {
	p.X += dx
	p.Y += dy
}

func (p *TPoint) Scale(val float64) {
	p.X = int32(float64(p.X) * val)
	p.Y = int32(float64(p.Y) * val)
}

func (p *TPoint) Scale2(val int) {
	p.Scale(float64(val))
}

// TMsg
//
// Only Windows, tagMSG
type TMsg struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      TPoint
}

// TCursorInfo
type TCursorInfo struct {
	CbSize      uint32
	Flags       uint32
	HCursor     HCURSOR
	PtScreenPos TPoint
}

// TWndClass
type TWndClass struct {
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     uintptr
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  LPCWSTR
	LpszClassName LPCWSTR
}

// NewSet
//
// 新建TSet，初始值为0，然后添加元素
//
// Create a new TSet, the initial value is 0, and then add elements.
func NewSet(opts ...int32) TSet {
	var s TSet
	return s.Include(opts...)
}

// Include
//
// 集合加法，val...中存储为位的索引，下标为0
//
// Set addition, stored as bit index in val..., subscript 0.
func (s TSet) Include(val ...int32) TSet {
	r := uint32(s)
	for _, v := range val {
		r |= 1 << uint8(v)
	}
	return TSet(r)
}

// Exclude
//
// 集合减法，val...中存储为位的索引，下标为0
//
// Set subtraction, stored as bit index in val..., subscript 0.
func (s TSet) Exclude(val ...int32) TSet {
	r := uint32(s)
	for _, v := range val {
		r &= ^(1 << uint8(v))
	}
	return TSet(r)
}

// In
//
// 集合类型的判断，val表示位数，下标为0
//
// Judgment of the Set type, val represents the number of digits, and the subscript is 0.
func (s TSet) In(val int32) bool {
	if s&(1<<uint8(val)) != 0 {
		return true
	}
	return false
}

type TBitmapInfoHeader struct { // use packed, this is the .bmp file format
	BiSize          DWORD
	BiWidth         LongInt
	BiHeight        LongInt
	BiPlanes        Word
	BiBitCount      Word
	BiCompression   DWORD
	BiSizeImage     DWORD
	BiXPelsPerMeter LongInt
	BiYPelsPerMeter LongInt
	BiClrUsed       DWORD
	BiClrImportant  DWORD
}
type TRGBQuad struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

type TBitmapInfo struct {
	BmiHeader TBitmapInfoHeader
	BmiColors [1]TRGBQuad
}

type TMonitorInfo struct {
	CbSize    DWORD
	RcMonitor TRect
	RcWork    TRect
	DwFlags   DWORD
}
