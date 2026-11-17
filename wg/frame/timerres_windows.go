//go:build windows

package frame

import (
	"sync"
	"syscall"
)

var (
	winmm               = syscall.NewLazyDLL("winmm.dll")
	procTimeBeginPeriod = winmm.NewProc("timeBeginPeriod")
	timerResOnce        sync.Once
)

// ensureTimerResolution 把系统定时器精度提到 1ms（整进程一次）。
//
// 为什么必须：Go 的 time.AfterFunc 走系统定时器粒度（Windows 默认
// 15.625ms），16ms 的帧闹钟实际要到 ~31ms 才触发——动画被钉在 ~32fps，
// 实测 34.6fps、91% 的帧间隔落在 20–40ms（showcase 的实测结论一致：
// setTimerResolution(1) 是 60Hz 的前提）。帧闹钟精度不够时，帧调度器
// 再怎么算 16ms 也只是把期望值写进一个 15.6ms 量化的格子。
func ensureTimerResolution() {
	timerResOnce.Do(func() {
		procTimeBeginPeriod.Call(1)
	})
}
