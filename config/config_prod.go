//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build prod

package config

import (
	"github.com/energye/lcl/tool/exec"
)

func init() {
	config = &Config{Framework: exec.AppDir()}
}

func (m *Config) FrameworkPath() string {
	return m.Framework
}
