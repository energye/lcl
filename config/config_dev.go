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
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/tool/exec"
	"os"
	"path/filepath"
)

func init() {
	homeDIR := exec.HomeDir
	if homeDIR == "" {
		println("[Warning] failed to obtain the current user directory.")
	} else {
		configPath := filepath.Join(homeDIR, ".energy", "config.json")
		tempConfig := &Config{}

		if !tool.IsExist(configPath) {
			// 不存在创建 config.json
			tempConfig.Framework = filepath.Join(exec.AppDir(), "frameworks")
			data, err := json.MarshalIndent(tempConfig, "", "\t")
			if err != nil {
				println("[ERROR] Create .energy config MarshalIndent Error:", err.Error())
				return
			}
			err = os.WriteFile(configPath, data, 0644)
			if err != nil {
				println("[Warning] Create .energy config WriteFile Error:", err.Error())
			}
			config = tempConfig
			return
		}
		// 存在读取 config.json
		data, err := os.ReadFile(configPath)
		if err != nil {
			println("[ERROR] ReadFile .energy Error:", err.Error())
			return
		}
		err = json.Unmarshal(data, tempConfig)
		if err != nil {
			println("[ERROR] Unmarshal .energy file Error:", err.Error())
			return
		}
		config = tempConfig
	}
}

func (m *Config) FrameworkPath() string {
	return m.Framework
}
