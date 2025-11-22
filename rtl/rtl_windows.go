//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import "github.com/energye/lcl/api"

// CreateURLShortCut
//
// 创建一个url的快捷方式
//
// Create a shortcut to a URL
//
//	rtl.CreateURLShortCut("C:\\aaa\\bbb\\", "MyDir", "https://github.com/energye/lcl")
func CreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	api.CreateURLShortCut(aDestPath, aShortCutName, aURL)
}

// CreateShortCut
//
// 创建一个快捷方式
//
// Create a shortcut
//  1. rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "MyDir", os.Args[0], "", "", "")
//  2. rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "MyDir", os.Args[0], "", "Description text", "-a -b")
func CreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return api.CreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs)
}
