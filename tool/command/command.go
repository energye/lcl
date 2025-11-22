//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package command

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Level 消息类型等级
type Level int32

const (
	LInfo Level = iota
	LError
)

func (m Level) String() string {
	switch m {
	case LInfo:
		return "INFO"
	case LError:
		return "ERROR"
	}
	return "UNKNOWN"
}

type CMD struct {
	HideWindow bool
	IsPrint    bool
	Dir        string
	Console    func(data string, level Level)
	Cmd        *exec.Cmd
}

func NewCMD() *CMD {
	return &CMD{IsPrint: true}
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func (m *CMD) Command(name string, args ...string) {
	if m.IsPrint {
		fmt.Println("run command:", name, strings.Join(args, " "))
	}
	cmd := exec.Command(name, args...)
	m.Cmd = cmd
	if m.Dir != "" {
		cmd.Dir = m.Dir
	}
	//隐藏调用外部命令窗口
	if m.HideWindow && IsWindows() {
		cmd.SysProcAttr = HideWindow(true)
	}

	console := func(data string, level Level) {
		if m.Console != nil {
			m.Console(data, level)
		}
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		console(err.Error(), LError)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		console(err.Error(), LError)
		return
	}

	err = cmd.Start()
	if err != nil {
		console(err.Error(), LError)
		return
	}

	// 处理标准输出
	go func() {
		scanner := bufio.NewScanner(stdout)
		// 增加缓冲区大小，避免 token too long 错误
		buf := make([]byte, 0, 64*1024) // 64KB 缓冲区
		scanner.Buffer(buf, 1024*1024)  // 最大 1MB

		for scanner.Scan() {
			data := string(scanner.Bytes())
			if m.IsPrint {
				fmt.Println(data)
			}
			console(data, LInfo)
		}
		if err := scanner.Err(); err != nil {
			console(err.Error(), LError)
		}
	}()

	// 处理标准错误
	go func() {
		scanner := bufio.NewScanner(stderr)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)
		for scanner.Scan() {
			data := string(scanner.Bytes())
			if m.IsPrint {
				fmt.Println(data)
			}
			console(data, LInfo)
		}
		if err := scanner.Err(); err != nil {
			console(err.Error(), LError)
		}
	}()

	err = cmd.Wait()
	if err != nil {
		console(err.Error(), LError)
	}
}
