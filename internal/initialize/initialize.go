// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

// Package Application environment initialization loads dynamic libraries

package initialize

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/rtl/version"
)

// Initialize
// runtime libenergy
func Initialize() {
	loadLibENERGY()
	api.Init()
	// rtl
	rtl.Init()
	// version
	version.Init()
	// system platform api
	APIInit()
}
