//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package font

type CharSet = byte

const (
	ANSI_CHARSET        CharSet = 0
	DEFAULT_CHARSET     CharSet = 1
	SYMBOL_CHARSET      CharSet = 2
	FCS_ISO_10646_1     CharSet = 4  // Unicode;
	FCS_ISO_8859_1      CharSet = 5  //  ISO Latin-1 (Western Europe);
	FCS_ISO_8859_2      CharSet = 6  //  ISO Latin-2 (Eastern Europe);
	FCS_ISO_8859_3      CharSet = 7  //  ISO Latin-3 (Southern Europe);
	FCS_ISO_8859_4      CharSet = 8  //  ISO Latin-4 (Northern Europe);
	FCS_ISO_8859_5      CharSet = 9  //  ISO Cyrillic;
	FCS_ISO_8859_6      CharSet = 10 //  ISO Arabic;
	FCS_ISO_8859_7      CharSet = 11 //  ISO Greek;
	FCS_ISO_8859_8      CharSet = 12 //  ISO Hebrew;
	FCS_ISO_8859_9      CharSet = 13 //  ISO Latin-5 (Turkish);
	FCS_ISO_8859_10     CharSet = 14 //  ISO Latin-6 (Nordic);
	FCS_ISO_8859_15     CharSet = 15 //  ISO Latin-9, or Latin-0 (Revised Western-European);
	MAC_CHARSET         CharSet = 77
	SHIFTJIS_CHARSET    CharSet = 128
	HANGEUL_CHARSET     CharSet = 129
	JOHAB_CHARSET       CharSet = 130
	GB2312_CHARSET      CharSet = 134
	CHINESEBIG5_CHARSET CharSet = 136
	GREEK_CHARSET       CharSet = 161
	TURKISH_CHARSET     CharSet = 162
	VIETNAMESE_CHARSET  CharSet = 163
	HEBREW_CHARSET      CharSet = 177
	ARABIC_CHARSET      CharSet = 178
	BALTIC_CHARSET      CharSet = 186
	RUSSIAN_CHARSET     CharSet = 204
	THAI_CHARSET        CharSet = 222
	EASTEUROPE_CHARSET  CharSet = 238
	OEM_CHARSET         CharSet = 255
)

func CharSetToString(charset CharSet) string {
	switch charset {
	case ANSI_CHARSET:
		return "ANSI_CHARSET"
	case DEFAULT_CHARSET:
		return "DEFAULT_CHARSET"
	case SYMBOL_CHARSET:
		return "SYMBOL_CHARSET"
	case MAC_CHARSET:
		return "MAC_CHARSET"
	case SHIFTJIS_CHARSET:
		return "SHIFTJIS_CHARSET"
	case HANGEUL_CHARSET:
		return "HANGEUL_CHARSET"
	case JOHAB_CHARSET:
		return "JOHAB_CHARSET"
	case GB2312_CHARSET:
		return "GB2312_CHARSET"
	case CHINESEBIG5_CHARSET:
		return "CHINESEBIG5_CHARSET"
	case GREEK_CHARSET:
		return "GREEK_CHARSET"
	case TURKISH_CHARSET:
		return "TURKISH_CHARSET"
	case VIETNAMESE_CHARSET:
		return "VIETNAMESE_CHARSET"
	case HEBREW_CHARSET:
		return "HEBREW_CHARSET"
	case ARABIC_CHARSET:
		return "ARABIC_CHARSET"
	case BALTIC_CHARSET:
		return "BALTIC_CHARSET"
	case RUSSIAN_CHARSET:
		return "RUSSIAN_CHARSET"
	case THAI_CHARSET:
		return "THAI_CHARSET"
	case EASTEUROPE_CHARSET:
		return "EASTEUROPE_CHARSET"
	case OEM_CHARSET:
		return "OEM_CHARSET"
	case FCS_ISO_10646_1:
		return "UNICODE"
	case FCS_ISO_8859_1:
		return "FCS_ISO_8859_1"
	case FCS_ISO_8859_2:
		return "FCS_ISO_8859_2"
	case FCS_ISO_8859_3:
		return "FCS_ISO_8859_3"
	case FCS_ISO_8859_4:
		return "FCS_ISO_8859_4"
	case FCS_ISO_8859_5:
		return "FCS_ISO_8859_5"
	case FCS_ISO_8859_6:
		return "FCS_ISO_8859_6"
	case FCS_ISO_8859_7:
		return "FCS_ISO_8859_7"
	case FCS_ISO_8859_8:
		return "FCS_ISO_8859_8"
	case FCS_ISO_8859_9:
		return "FCS_ISO_8859_9"
	case FCS_ISO_8859_10:
		return "FCS_ISO_8859_10"
	case FCS_ISO_8859_15:
		return "FCS_ISO_8859_15"
	default:
		return ""
	}
}
