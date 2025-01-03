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

import (
	"bytes"
	"fmt"
	"github.com/energye/lcl/api/libname"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func init() {
	MacApp.execName = os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	MacApp.macContentsDir = os.Args[0] + ".app/Contents"
	MacApp.macOSDir = MacApp.macContentsDir + "/MacOS"
	MacApp.macResources = MacApp.macContentsDir + "/Resources"
	MacApp.execFile = MacApp.macOSDir + "/" + MacApp.execName
	MacApp.plistFileName = MacApp.macContentsDir + "/Info.plist"
	MacApp.pkgInfoFileName = MacApp.macContentsDir + "/PkgInfo"
	MacApp.macAppFrameworksDir = MacApp.macContentsDir + "/Frameworks"
	MacApp.lclLibFileName = MacApp.macContentsDir + "/Frameworks/liblcl.dylib" //liblcl to frameworks
	MacApp.lsUIElement = "false"
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 生成目录并拷贝文件
func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	destSplitPathDirs := strings.Split(dest, "/")
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b := fileExists(destSplitPath)
			if b == false {
				if err := os.Mkdir(destSplitPath, 0755); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func Init() {
	if strings.Contains(os.Args[0], ".app/Contents/MacOS") {
		return
	}
	MacApp.isMain = true
	// 创建 xxx.app
	if MacApp.createMacOSApp(MacApp) {
		MacApp.copyDylib()
		MacApp.runMacOSApp()
	}
}

func (m *macApp) runMacOSApp() {
	cmd := exec.Command(m.execFile, os.Args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func (m *macApp) copyDylib() {
	var libPath = libname.GetLibPath(libname.GetDLLName())
	// 文件不存在，复制
	if !fileExists(m.lclLibFileName) {
		if fileExists(libPath) {
			copyFile(libPath, m.lclLibFileName)
		}
	} else {
		// 文件存在，对比后更新
		if fileExists(libPath) {
			f1, _ := os.Stat(libPath)
			f2, _ := os.Stat(m.lclLibFileName)
			if f1.Size() != f2.Size() {
				copyFile(libPath, m.lclLibFileName)
			}
		}
	}
}

// 以一个Mac下的app形式运行
// 调试下使用，正式发布的时候虽然可 以不用去掉，但也不咋好
func (*macApp) createMacOSApp(m *macApp) bool {
	if !fileExists(m.macOSDir) {
		if err := os.MkdirAll(m.macOSDir, 0755); err != nil {
			return false
		}
	}
	if !fileExists(m.macResources) {
		os.MkdirAll(m.macResources, 0755)
	}
	if !fileExists(m.macAppFrameworksDir) {
		os.MkdirAll(m.macAppFrameworksDir, 0755)
	}
	resName := fmt.Sprintf("%s/%s.icns", m.macResources, m.execName)
	if !fileExists(resName) {
		ioutil.WriteFile(resName, macOSAppIcon, 0666)
	}
	if !fileExists(m.plistFileName) {
		datas := map[string]string{
			"execName":    m.execName,
			"devRegion":   "China", // China English
			"locale":      "zh_CN", //os.Getenv("LANG"),
			"copyright":   "copyright xxxx",
			"LSUIElement": m.lsUIElement,
		}
		buff := bytes.NewBuffer([]byte{})
		tmp := template.New("file")
		tmp.Parse(infoplist)
		tmp.Execute(buff, datas)
		ioutil.WriteFile(m.plistFileName, buff.Bytes(), 0666)
	}
	if !fileExists(m.pkgInfoFileName) {
		ioutil.WriteFile(m.pkgInfoFileName, pkgInfo, 0666)
	}
	// 主进程执行文件
	if _, err := copyFile(os.Args[0], m.execFile); err == nil {
		os.Chmod(m.execFile, 0755)
		return true
	}
	return false
}
