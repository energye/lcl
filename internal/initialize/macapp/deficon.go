// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

//go:build darwin && !prod
// +build darwin,!prod

package macapp

import _ "embed"

//go:embed icon.icns
var macOSAppIcon []byte
