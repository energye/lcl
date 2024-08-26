<p align="center">
   <span style="color: #2ba9f1;font-size: 24px;font-weight: bold;">Go LCL</span>
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        是Go基于 LCL(Lazarus Component Library) 构建桌面应用的框架
    </strong>
</p>
<p align="center">
    用于开发Windows、Mac OS和Linux平台的跨平台桌面应用程序
</p>

---

This project is under development in version 3.0 of [energy](https://github.com/energye/energy)

---
![go-version](https://img.shields.io/github/go-mod/go-version/energye/lcl?logo=git&logoColor=green)
[![github](https://img.shields.io/github/last-commit/energye/lcl/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/lcl)
[![release](https://img.shields.io/github/v/release/energye/lcl?logo=git&logoColor=green)](https://github.com/energye/lcl/releases)
![repo](https://img.shields.io/github/repo-size/energye/lcl.svg?logo=github&logoColor=green&label=repo-size)
[![Go Report](https://goreportcard.com/badge/github.com/energye/lcl)](https://goreportcard.com/report/github.com/energye/lcl)
[![Go Reference](https://pkg.go.dev/badge/github.com/energye/lcl)](https://pkg.go.dev/github.com/energye/lcl)
[![license](https://img.shields.io/github/license/energye/lcl.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
---

### 项目简介

> [Go LCL](https://github.com/energye/lcl)
> 是 Go 基于
> [LCL](https://www.lazarus-ide.org/)
> 开发的框架
>
>> LCL - 基础库, 图形用户界面(GUI)组件库, 提供了非常丰富的系统原生控件
>
> 构建&使用
>
>> LCL 开发原生图形用户界面(GUI)应用. 轻量级, 丰富的系统原生控件



### 特点

> - 仅需 `Go` 和 `liblcl` 动态链接库
> - 丰富的系统原生控件, 跨平台-支持 Windows、Mac OS、Linux

### 内置依赖&集成

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/lcl)

#### 基本需求

> - Golang >= 1.20
> - 动态链接库 `liblcl.dll` 当前仅提供了Windows测试版本

#### 开发环境

1. 安装 [Golang](https://golang.google.cn/dl/), Windows版本, 仅支持intel架构 [https://golang.google.cn/dl](https://golang.google.cn/dl)
2. 下载  `LCL` 控件库动态链接库, [下载地址](https://github.com/energye/lcl/tree/main/bins/webview2/windows)
3. 将动态链接库配置到环境变量 `ENERGY_HOME` 目录下, 或放置到和执行文件 `exe` 同一目录
4. 创建Go项目开始使用 `LCL` 构建桌面应用, 参考 `Go` 示例 [LCL examples](https://github.com/energye/examples/tree/main/lcl)

### 相关项目
* [Go LCL](https://github.com/energye/lcl)
* [Go Webview2](https://github.com/energye/wv)
* [Go CEF](https://github.com/energye/cef)
* [Go Webkit](https://github.com/energye/wk)
* [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi)
* [CEF](https://github.com/chromiumembedded/cef)
* [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
* [CefSharp](https://github.com/cefsharp/CefSharp)
* [Java-CEF](https://bitbucket.org/chromiumembedded/java-cef)
* [cefpython](https://github.com/cztomczak/cefpython)
* [Chromium](https://chromium.googlesource.com/chromium/src/)

---

### 欢迎加入
energy底层由多个项目模块组成, 因过于复杂扔处于建设的过程中，有很多的事情无法独自完成，如果有感兴趣的同学想参与energy的实现或学习，可通过微信或QQ联系我。

如果你觉得此项目对你有帮助，请点亮 Star

---

### ENERGY QQ交流群 & 微信

<p align="center">
    <img src="https://assets.yanghy.cn/qq-group.jpg" width="250" title="QQ交流群: 541258627" alt="QQ交流群: 541258627">
    <img src="https://assets.yanghy.cn/we-chat.jpg" width="250" title="微信: sniawmdf" alt="微信: sniawmdf" style="margin-left: 30px;">
</p>

---

### 鸣谢 Jetbrains

<p align="center">
    <a href="https://www.jetbrains.com?from=energy">
        <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo.">
    </a>
</p>

---

### 项目截图
##### Windows-10
<img src="https://assets.yanghy.cn/lcl-simple1.png">
<img src="https://assets.yanghy.cn/lcl-simple2.png">
----

### 开源协议

[![license](https://img.shields.io/github/license/energye/wv.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)
