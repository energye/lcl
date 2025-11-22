//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package messages

const (
	LM_NULL              = 0x0000
	LM_CREATE            = 0x0001
	LM_DESTROY           = 0x0002
	LM_MOVE              = 0x0003
	LM_SIZE              = 0x0005
	LM_ACTIVATE          = 0x0006
	LM_SETFOCUS          = 0x0007
	LM_KILLFOCUS         = 0x0008
	LM_ENABLE            = 0x000A
	LM_GETTEXTLENGTH     = 0x000E
	LM_PAINT             = 0x000F
	LM_ERASEBKGND        = 0x0014
	LM_SHOWWINDOW        = 0x0018
	LM_CANCELMODE        = 0x001F
	LM_SETCURSOR         = 0x0020
	LM_DRAWITEM          = 0x002B
	LM_MEASUREITEM       = 0x002C
	LM_DELETEITEM        = 0x002D
	LM_VKEYTOITEM        = 0x002E
	LM_CHARTOITEM        = 0x002F
	LM_SETFONT           = 0x0030
	LM_COMPAREITEM       = 0x0039
	LM_WINDOWPOSCHANGING = 0x0046
	LM_WINDOWPOSCHANGED  = 0x0047
	LM_NOTIFY            = 0x004E
	LM_HELP              = 0x0053
	LM_NOTIFYFORMAT      = 0x0055
	LM_CONTEXTMENU       = 0x007B
	LM_NCCALCSIZE        = 0x0083
	LM_NCHITTEST         = 0x0084
	LM_NCPAINT           = 0x0085
	LM_NCACTIVATE        = 0x0086
	LM_GETDLGCODE        = 0x0087
	LM_NCMOUSEMOVE       = 0x00A0
	LM_NCLBUTTONDOWN     = 0x00A1
	LM_NCLBUTTONUP       = 0x00A2
	LM_NCLBUTTONDBLCLK   = 0x00A3
	LM_KEYFIRST          = 0x0100
	LM_KEYDOWN           = 0x0100
	LM_KEYUP             = 0x0101
	LM_CHAR              = 0x0102
	LM_SYSKEYDOWN        = 0x0104
	LM_SYSKEYUP          = 0x0105
	LM_SYSCHAR           = 0x0106
	LM_KEYLAST           = 0x0108
	LM_COMMAND           = 0x0111
	LM_SYSCOMMAND        = 0x0112
	LM_TIMER             = 0x0113
	LM_HSCROLL           = 0x0114
	LM_VSCROLL           = 0x0115
	LM_CTLCOLORMSGBOX    = 0x0132
	LM_CTLCOLOREDIT      = 0x0133
	LM_CTLCOLORLISTBOX   = 0x0134
	LM_CTLCOLORBTN       = 0x0135
	LM_CTLCOLORDLG       = 0x0136
	LM_CTLCOLORSCROLLBAR = 0x0137
	LM_CTLCOLORSTATIC    = 0x0138
	LM_MOUSEFIRST        = 0x0200
	LM_MOUSEMOVE         = 0x0200
	LM_LBUTTONDOWN       = 0x0201
	LM_LBUTTONUP         = 0x0202
	LM_LBUTTONDBLCLK     = 0x0203
	LM_RBUTTONDOWN       = 0x0204
	LM_RBUTTONUP         = 0x0205
	LM_RBUTTONDBLCLK     = 0x0206
	LM_MBUTTONDOWN       = 0x0207
	LM_MBUTTONUP         = 0x0208
	LM_MBUTTONDBLCLK     = 0x0209
	LM_MOUSEWHEEL        = 0x020A
	LM_XBUTTONDOWN       = 0x020B
	LM_XBUTTONUP         = 0x020C
	LM_XBUTTONDBLCLK     = 0x020D
	LM_MOUSEHWHEEL       = 0x020E
	LM_MOUSELAST         = 0x020E
	LM_PARENTNOTIFY      = 0x0210
	LM_CAPTURECHANGED    = 0x0215
	LM_DROPFILES         = 0x0233
	LM_SELCHANGE         = 0x0234
	LM_DPICHANGED        = 0x02E0
	LM_CUT               = 0x0300
	LM_COPY              = 0x0301
	LM_PASTE             = 0x0302
	LM_CLEAR             = 0x0303
)

// -------------
// lcl messages
//
// This should be a list of LCL specific messages
// RECEIVED from the interface, here are no defines
// of messages send to the interface
// -------------
const (
	LM_USER = 0x400 // standard WM_USER value
	WM_USER = LM_USER

	// reserve some space for LM_USER messages
	LM_LCL              = LM_USER + 0x10000
	LM_ACTIVATEITEM     = LM_LCL + 4 // GTK internal. Should be removed later
	LM_CHANGED          = LM_LCL + 5
	LM_FOCUS            = LM_LCL + 6
	LM_CLICKED          = LM_LCL + 7
	LM_RELEASED         = LM_LCL + 9
	LM_ENTER            = LM_LCL + 11
	LM_LEAVE            = LM_LCL + 12
	LM_CHECKRESIZE      = LM_LCL + 14
	LM_SETEDITABLE      = LM_LCL + 18
	LM_MOVEWORD         = LM_LCL + 19
	LM_MOVEPAGE         = LM_LCL + 20
	LM_MOVETOROW        = LM_LCL + 21
	LM_MOVETOCOLUMN     = LM_LCL + 22
	LM_KILLCHAR         = LM_LCL + 23
	LM_KILLWORD         = LM_LCL + 24
	LM_KILLLINE         = LM_LCL + 25
	LM_CONFIGUREEVENT   = LM_LCL + 31 // GTK internal. Should be removed later.
	LM_EXIT             = LM_LCL + 60
	LM_CLOSEQUERY       = LM_LCL + 62
	LM_DRAGSTART        = LM_LCL + 63
	LM_QUIT             = LM_LCL + 65
	LM_MONTHCHANGED     = LM_LCL + 66
	LM_YEARCHANGED      = LM_LCL + 67
	LM_DAYCHANGED       = LM_LCL + 68
	LM_MOUSEFIRST2      = LM_LCL + 70
	LM_LBUTTONTRIPLECLK = LM_MOUSEFIRST2 + 0
	LM_LBUTTONQUADCLK   = LM_MOUSEFIRST2 + 1
	LM_MBUTTONTRIPLECLK = LM_MOUSEFIRST2 + 2
	LM_MBUTTONQUADCLK   = LM_MOUSEFIRST2 + 3
	LM_RBUTTONTRIPLECLK = LM_MOUSEFIRST2 + 4
	LM_RBUTTONQUADCLK   = LM_MOUSEFIRST2 + 5
	LM_MOUSEENTER       = LM_MOUSEFIRST2 + 6
	LM_MOUSELEAVE       = LM_MOUSEFIRST2 + 7
	LM_XBUTTONTRIPLECLK = LM_MOUSEFIRST2 + 8
	LM_XBUTTONQUADCLK   = LM_MOUSEFIRST2 + 9
	LM_MOUSELAST2       = LM_XBUTTONQUADCLK
	// for triple and quad clicks see below

	LM_GRABFOCUS    = LM_LCL + 80
	LM_DRAWLISTITEM = LM_LCL + 81
	LM_DEFERREDEDIT = LM_LCL + 82 // used in customdbcombobox

	// these IDs are reserved for internal messages in the interfaces
	LM_INTERFACEFIRST = LM_LCL + 99
	LM_INTERFACELAST  = LM_LCL + 199
	LM_UNKNOWN        = LM_INTERFACELAST + 1
	LM_IM_COMPOSITION = LM_USER + 0xFFF0 // gtk IM

	// GTK IM Flags
	GTK_IM_FLAG_START   = 1
	GTK_IM_FLAG_PREEDIT = 2
	GTK_IM_FLAG_END     = 4
	GTK_IM_FLAG_COMMIT  = 8
	GTK_IM_FLAG_REPLACE = 16
)
