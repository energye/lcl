//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !prod
// +build !prod

package config

import (
	"encoding/json"
	"github.com/energye/lcl/tool/exec"
	"os"
	"path/filepath"
)

func init() {
	homeDIR := exec.HomeDir
	if homeDIR == "" {
		println("[Warning] failed to obtain the current user directory.")
	} else {
		filePath := filepath.Join(homeDIR, ".energy", "config.json")
		data, err := os.ReadFile(filePath)
		if err != nil {
			println("[ERROR] Read .energy Error:", err.Error())
			return
		}
		tempConfig := &Config{}
		err = json.Unmarshal(data, tempConfig)
		if err != nil {
			println("[ERROR] Parsing .energy file Error:", err.Error())
			return
		}
		config = tempConfig
	}
}

func (m *Config) FrameworkPath() string {
	return m.Framework
}
