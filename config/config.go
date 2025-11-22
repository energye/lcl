//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package config

import (
	"strconv"
	"strings"
)

var config *Config

type Config struct {
	Framework string `json:"framework"` // 框架目录
	Registry  string `json:"registry"`  // 远程服务配置地址
	Proxy     string `json:"proxy"`     // 代理地址
	IsCEF     bool   `json:"-"`
}

// Get 返回 config 环境
// mode: 构建模式 dev 或 prod
// 当 mode 是 dev 时即开发环境 使用 $HOME/.energy配置, 是 prod 时即生产模式不再使用 $HOME/.energy配置, 使用自定义或当前执行目录
func Get() *Config {
	return config
}

// Version 返回当前使用的 CEF Framework 版本号
// 规则 Framework: CEF-[VER]_[OS]_[ARCH]
// 返回 VER
func (m *Config) Version() int {
	framework := strings.Split(m.Framework, "-")
	// CEF [VER]_[OS]_[ARCH]
	if len(framework) != 2 {
		return 0
	}
	// split: [VER]_[OS]_[ARCH]
	framework = strings.Split(framework[1], "_")
	if len(framework) != 3 {
		return 0
	}
	// [VER] [OS] [ARCH]
	ver, err := strconv.Atoi(framework[0])
	if err != nil {
		return 0
	}
	return ver
}
