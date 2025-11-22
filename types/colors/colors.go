//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package colors

import (
	"fmt"
	"strconv"
)

type TColor uint32

func (c TColor) R() byte {
	return byte(c)
}

func (c TColor) G() byte {
	return byte(c >> 8)
}

func (c TColor) B() byte {
	return byte(c >> 16)
}

func (c TColor) RGB(r, g, b byte) TColor {
	return TColor(uint32(r) | (uint32(g) << 8) | (uint32(b) << 16))
}

func (c TColor) BGR() TColor {
	return TColor(byte(c>>16)) | (TColor(byte(c>>8)) << 8) | (TColor(byte(c)) << 16)
}

// Lazarus中的颜色表，Lazarus中的TColor与一般HTML的RGB有点不一样，反过来的以BGR表示.

// 原Delphi中的定义
const (
	ClClSysNone            TColor = 0x1FFFFFFF
	ClSysDefault           TColor = 0x20000000
	ClAliceblue            TColor = 0xFFF8F0
	ClAntiquewhite         TColor = 0xD7EBFA
	ClAqua                 TColor = 0xFFFF00
	ClAquamarine           TColor = 0xD4FF7F
	ClAzure                TColor = 0xFFFFF0
	ClBeige                TColor = 0xDCF5F5
	ClBisque               TColor = 0xC4E4FF
	ClBlack                TColor = 0x000000
	ClBlanchedalmond       TColor = 0xCDEBFF
	ClBlue                 TColor = 0xFF0000
	ClBlueviolet           TColor = 0xE22B8A
	ClBrown                TColor = 0x2A2AA5
	ClBurlywood            TColor = 0x87B8DE
	ClCadetblue            TColor = 0xA09E5F
	ClChartreuse           TColor = 0x00FF7F
	ClChocolate            TColor = 0x1E69D2
	ClCoral                TColor = 0x507FFF
	ClCornflowerblue       TColor = 0xED9564
	ClCornsilk             TColor = 0xDCF8FF
	ClCrimson              TColor = 0x3C14DC
	ClCyan                 TColor = 0xFFFF00
	ClDarkblue             TColor = 0x8B0000
	ClDarkcyan             TColor = 0x8B8B00
	ClDarkgoldenrod        TColor = 0x0B86B8
	ClDarkgray             TColor = 0xA9A9A9
	ClDarkgreen            TColor = 0x006400
	ClDarkgrey             TColor = 0xA9A9A9
	ClDarkkhaki            TColor = 0x6BB7BD
	ClDarkmagenta          TColor = 0x8B008B
	ClDarkolivegreen       TColor = 0x2F6B55
	ClDarkorange           TColor = 0x008CFF
	ClDarkorchid           TColor = 0xCC3299
	ClDarkred              TColor = 0x00008B
	ClDarksalmon           TColor = 0x7A96E9
	ClDarkseagreen         TColor = 0x8FBC8F
	ClDarkslateblue        TColor = 0x8B3D48
	ClDarkslategray        TColor = 0x4F4F2F
	ClDarkslategrey        TColor = 0x4F4F2F
	ClDarkturquoise        TColor = 0xD1CE00
	ClDarkviolet           TColor = 0xD30094
	ClDeeppink             TColor = 0x9314FF
	ClDeepskyblue          TColor = 0xFFBF00
	ClDimgray              TColor = 0x696969
	ClDimgrey              TColor = 0x696969
	ClDodgerblue           TColor = 0xFF901E
	ClFirebrick            TColor = 0x2222B2
	ClFloralwhite          TColor = 0xF0FAFF
	ClForestgreen          TColor = 0x228B22
	ClFuchsia              TColor = 0xFF00FF
	ClGainsboro            TColor = 0xDCDCDC
	ClGhostwhite           TColor = 0xFFF8F8
	ClGold                 TColor = 0x00D7FF
	ClGoldenrod            TColor = 0x20A5DA
	ClGray                 TColor = 0x808080
	ClGreen                TColor = 0x008000
	ClGreenyellow          TColor = 0x2FFFAD
	ClGrey                 TColor = 0x808080
	ClHoneydew             TColor = 0xF0FFF0
	ClHotpink              TColor = 0xB469FF
	ClIndianred            TColor = 0x5C5CCD
	ClIndigo               TColor = 0x82004B
	ClIvory                TColor = 0xF0FFFF
	ClKhaki                TColor = 0x8CE6F0
	ClLavender             TColor = 0xFAE6E6
	ClLavenderblush        TColor = 0xF5F0FF
	ClLawngreen            TColor = 0x00FC7C
	ClLemonchiffon         TColor = 0xCDFAFF
	ClLightblue            TColor = 0xE6D8AD
	ClLightcoral           TColor = 0x8080F0
	ClLightcyan            TColor = 0xFFFFE0
	ClLightgoldenrodyellow TColor = 0xD2FAFA
	ClLightgray            TColor = 0xD3D3D3
	ClLightgreen           TColor = 0x90EE90
	ClLightgrey            TColor = 0xD3D3D3
	ClLightpink            TColor = 0xC1B6FF
	ClLightsalmon          TColor = 0x7AA0FF
	ClLightseagreen        TColor = 0xAAB220
	ClLightskyblue         TColor = 0xFACE87
	ClLightslategray       TColor = 0x998877
	ClLightslategrey       TColor = 0x998877
	ClLightsteelblue       TColor = 0xDEC4B0
	ClLightyellow          TColor = 0xE0FFFF
	ClLtGray               TColor = 0xC0C0C0
	ClMedGray              TColor = 0xA4A0A0
	ClDkGray               TColor = 0x808080
	ClMoneyGreen           TColor = 0xC0DCC0
	ClLegacySkyBlue        TColor = 0xF0CAA6
	ClCream                TColor = 0xF0FBFF
	ClLime                 TColor = 0x00FF00
	ClLimegreen            TColor = 0x32CD32
	ClLinen                TColor = 0xE6F0FA
	ClMagenta              TColor = 0xFF00FF
	ClMaroon               TColor = 0x000080
	ClMediumaquamarine     TColor = 0xAACD66
	ClMediumblue           TColor = 0xCD0000
	ClMediumorchid         TColor = 0xD355BA
	ClMediumpurple         TColor = 0xDB7093
	ClMediumseagreen       TColor = 0x71B33C
	ClMediumslateblue      TColor = 0xEE687B
	ClMediumspringgreen    TColor = 0x9AFA00
	ClMediumturquoise      TColor = 0xCCD148
	ClMediumvioletred      TColor = 0x8515C7
	ClMidnightblue         TColor = 0x701919
	ClMintcream            TColor = 0xFAFFF5
	ClMistyrose            TColor = 0xE1E4FF
	ClMoccasin             TColor = 0xB5E4FF
	ClNavajowhite          TColor = 0xADDEFF
	ClNavy                 TColor = 0x800000
	ClOldlace              TColor = 0xE6F5FD
	ClOlive                TColor = 0x008080
	ClOlivedrab            TColor = 0x238E6B
	ClOrange               TColor = 0x00A5FF
	ClOrangered            TColor = 0x0045FF
	ClOrchid               TColor = 0xD670DA
	ClPalegoldenrod        TColor = 0xAAE8EE
	ClPalegreen            TColor = 0x98FB98
	ClPaleturquoise        TColor = 0xEEEEAF
	ClPalevioletred        TColor = 0x9370DB
	ClPapayawhip           TColor = 0xD5EFFF
	ClPeachpuff            TColor = 0xB9DAFF
	ClPeru                 TColor = 0x3F85CD
	ClPink                 TColor = 0xCBC0FF
	ClPlum                 TColor = 0xDDA0DD
	ClPowderblue           TColor = 0xE6E0B0
	ClPurple               TColor = 0x800080
	ClRed                  TColor = 0x0000FF
	ClRosybrown            TColor = 0x8F8FBC
	ClRoyalblue            TColor = 0xE16941
	ClSaddlebrown          TColor = 0x13458B
	ClSalmon               TColor = 0x7280FA
	ClSandybrown           TColor = 0x60A4F4
	ClSeagreen             TColor = 0x578B2E
	ClSeashell             TColor = 0xEEF5FF
	ClSienna               TColor = 0x2D52A0
	ClSilver               TColor = 0xC0C0C0
	ClSkyblue              TColor = 0xEBCE87
	ClSlateblue            TColor = 0xCD5A6A
	ClSlategray            TColor = 0x908070
	ClSlategrey            TColor = 0x908070
	ClSnow                 TColor = 0xFAFAFF
	ClSpringgreen          TColor = 0x7FFF00
	ClSteelblue            TColor = 0xB48246
	ClTan                  TColor = 0x8CB4D2
	ClTeal                 TColor = 0x808000
	ClThistle              TColor = 0xD8BFD8
	ClTomato               TColor = 0x4763FF
	ClTurquoise            TColor = 0xD0E040
	ClViolet               TColor = 0xEE82EE
	ClWheat                TColor = 0xB3DEF5
	ClWhite                TColor = 0xFFFFFF
	ClWhitesmoke           TColor = 0xF5F5F5
	ClYellow               TColor = 0x00FFFF
	ClYellowgreen          TColor = 0x32CD9A
	ClBtnFace              TColor = 0xFF00000F
	ClNull                 TColor = 0x00000000
)

// Lazarus中的定义
const (
	//CLR_NONE    = 0xFFFFFFFF
	//CLR_DEFAULT = 0xFF000000
	//CLR_INVALID = 0xFFFFFFFF

	cOLOR_SCROLLBAR           = 0
	cOLOR_BACKGROUND          = 1
	cOLOR_ACTIVECAPTION       = 2
	cOLOR_INACTIVECAPTION     = 3
	cOLOR_MENU                = 4
	cOLOR_WINDOW              = 5
	cOLOR_WINDOWFRAME         = 6
	cOLOR_MENUTEXT            = 7
	cOLOR_WINDOWTEXT          = 8
	cOLOR_CAPTIONTEXT         = 9
	cOLOR_ACTIVEBORDER        = 10
	cOLOR_INACTIVEBORDER      = 11
	cOLOR_APPWORKSPACE        = 12
	cOLOR_HIGHLIGHT           = 13
	cOLOR_HIGHLIGHTTEXT       = 14
	cOLOR_BTNFACE             = 15
	cOLOR_BTNSHADOW           = 16
	cOLOR_GRAYTEXT            = 17
	cOLOR_BTNTEXT             = 18
	cOLOR_INACTIVECAPTIONTEXT = 19
	cOLOR_BTNHIGHLIGHT        = 20
	cOLOR_3DDKSHADOW          = 21
	cOLOR_3DLIGHT             = 22
	cOLOR_INFOTEXT            = 23
	cOLOR_INFOBK              = 24
	// PBD: 25 is unassigned in all the docs I can find
	//      if someone finds what this is supposed to be then fill it in
	//      note defaults below, and cl[ColorConst] in graphics
	cOLOR_HOTLIGHT                = 26
	cOLOR_GRADIENTACTIVECAPTION   = 27
	cOLOR_GRADIENTINACTIVECAPTION = 28
	cOLOR_MENUHILIGHT             = 29
	cOLOR_MENUBAR                 = 30
	cOLOR_FORM                    = 31
	cOLOR_ENDCOLORS               = cOLOR_FORM
	cOLOR_DESKTOP                 = cOLOR_BACKGROUND
	cOLOR_3DFACE                  = cOLOR_BTNFACE
	cOLOR_3DSHADOW                = cOLOR_BTNSHADOW
	cOLOR_3DHIGHLIGHT             = cOLOR_BTNHIGHLIGHT
	cOLOR_3DHILIGHT               = cOLOR_BTNHIGHLIGHT
	cOLOR_BTNHILIGHT              = cOLOR_BTNHIGHLIGHT
	mAX_SYS_COLORS                = cOLOR_ENDCOLORS
	sYS_COLOR_BASE                = 0x80000000

	// The following colors match the predefined Delphi Colors
	// standard colors
	//ClBlack   = 0x000000
	//ClMaroon  = 0x000080
	//ClGreen   = 0x008000
	//ClOlive   = 0x008080
	//ClNavy    = 0x800000
	//ClPurple  = 0x800080
	//ClTeal    = 0x808000
	//ClGray    = 0x808080
	//ClSilver  = 0xC0C0C0
	//ClRed     = 0x0000FF
	//ClLime    = 0x00FF00
	//ClYellow  = 0x00FFFF
	//ClBlue    = 0xFF0000
	//ClFuchsia = 0xFF00FF
	//ClAqua    = 0xFFFF00
	//ClLtGray  = 0xC0C0C0 // ClSilver alias
	//ClDkGray  = 0x808080 // ClGray alias
	//ClWhite   = 0xFFFFFF

	// extended colors
	//ClMoneyGreen = 0xC0DCC0
	//ClSkyBlue    = 0xF0CAA6
	//ClCream      = 0xF0FBFF
	//ClMedGray    = 0xA4A0A0

	// special colors
	ClNone    = 0x1FFFFFFF
	ClDefault = 0x20000000

	// system colors
	ClScrollBar       = sYS_COLOR_BASE | cOLOR_SCROLLBAR
	ClBackground      = sYS_COLOR_BASE | cOLOR_BACKGROUND
	ClActiveCaption   = sYS_COLOR_BASE | cOLOR_ACTIVECAPTION
	ClInactiveCaption = sYS_COLOR_BASE | cOLOR_INACTIVECAPTION
	ClMenu            = sYS_COLOR_BASE | cOLOR_MENU
	ClWindow          = sYS_COLOR_BASE | cOLOR_WINDOW
	ClWindowFrame     = sYS_COLOR_BASE | cOLOR_WINDOWFRAME
	ClMenuText        = sYS_COLOR_BASE | cOLOR_MENUTEXT
	ClWindowText      = sYS_COLOR_BASE | cOLOR_WINDOWTEXT
	ClCaptionText     = sYS_COLOR_BASE | cOLOR_CAPTIONTEXT
	ClActiveBorder    = sYS_COLOR_BASE | cOLOR_ACTIVEBORDER
	ClInactiveBorder  = sYS_COLOR_BASE | cOLOR_INACTIVEBORDER
	ClAppWorkspace    = sYS_COLOR_BASE | cOLOR_APPWORKSPACE
	ClHighlight       = sYS_COLOR_BASE | cOLOR_HIGHLIGHT
	ClHighlightText   = sYS_COLOR_BASE | cOLOR_HIGHLIGHTTEXT
	//ClBtnFace             = sYS_COLOR_BASE | cOLOR_BTNFACE
	ClBtnShadow           = sYS_COLOR_BASE | cOLOR_BTNSHADOW
	ClGrayText            = sYS_COLOR_BASE | cOLOR_GRAYTEXT
	ClBtnText             = sYS_COLOR_BASE | cOLOR_BTNTEXT
	ClInactiveCaptionText = sYS_COLOR_BASE | cOLOR_INACTIVECAPTIONTEXT
	ClBtnHighlight        = sYS_COLOR_BASE | cOLOR_BTNHIGHLIGHT
	Cl3DDkShadow          = sYS_COLOR_BASE | cOLOR_3DDKSHADOW
	Cl3DLight             = sYS_COLOR_BASE | cOLOR_3DLIGHT
	ClInfoText            = sYS_COLOR_BASE | cOLOR_INFOTEXT
	ClInfoBk              = sYS_COLOR_BASE | cOLOR_INFOBK

	ClHotLight                = sYS_COLOR_BASE | cOLOR_HOTLIGHT
	ClGradientActiveCaption   = sYS_COLOR_BASE | cOLOR_GRADIENTACTIVECAPTION
	ClGradientInactiveCaption = sYS_COLOR_BASE | cOLOR_GRADIENTINACTIVECAPTION
	ClMenuHighlight           = sYS_COLOR_BASE | cOLOR_MENUHILIGHT
	ClMenuBar                 = sYS_COLOR_BASE | cOLOR_MENUBAR
	ClForm                    = sYS_COLOR_BASE | cOLOR_FORM

	// synonyms: do not show them in color lists
	ClColorDesktop = sYS_COLOR_BASE | cOLOR_DESKTOP
	Cl3DFace       = sYS_COLOR_BASE | cOLOR_3DFACE
	Cl3DShadow     = sYS_COLOR_BASE | cOLOR_3DSHADOW
	Cl3DHiLight    = sYS_COLOR_BASE | cOLOR_3DHIGHLIGHT
	ClBtnHiLight   = sYS_COLOR_BASE | cOLOR_BTNHILIGHT

	ClFirstSpecialColor = ClBtnHiLight
	ClMask              = ClWhite
	ClDontMask          = ClBlack
)

// RGB
func RGB(r, g, b byte) uint32 {
	return uint32(r) | (uint32(g) << 8) | (uint32(b) << 16)
}

// RGBToBGR
func RGBToBGR(rgb uint32) uint32 {
	return uint32(byte(rgb>>16)) | (uint32(byte(rgb>>8)) << 8) | (uint32(byte(rgb)) << 16)
}

type ARGB = uint32

type TARGB struct {
	A uint32
	R uint32
	G uint32
	B uint32
}

func NewARGB(a, r, g, b uint32) *TARGB {
	return &TARGB{
		A: a,
		R: r,
		G: g,
		B: b,
	}
}

func (m *TARGB) ARGB() ARGB {
	ca, _ := strconv.ParseUint(fmt.Sprintf("%02X%02X%02X%02X", m.A, m.R, m.G, m.B), 16, 32)
	return ARGB(ca)
}
