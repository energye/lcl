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

var (
	config *Config
)

type Config struct {
	Framework string `json:"framework"` // 框架目录
}

// Get 返回 config 环境
// mode: 构建模式 dev 或 prod
// 当 mode 是 dev 时即开发环境 使用 $HOME/.energy配置, 是 prod 时即生产模式不再使用 $HOME/.energy配置, 使用自定义或当前执行目录
func Get() *Config {
	return config
}
