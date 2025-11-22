//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package colors

func Blue(rgb TColor) byte {
	return byte((rgb >> 16) & 0xff)
}

func Green(rgb TColor) byte {
	return byte((rgb >> 8) & 0xff)
}

func Red(rgb TColor) byte {
	return byte(rgb & 0xff)
}

func RGBToColor(R, G, B byte) TColor {
	return TColor(B)<<16 | TColor(G)<<8 | TColor(R)
}
