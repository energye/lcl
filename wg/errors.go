package wg

import "errors"

// 门面层错误。全部为哨兵值，便于用 errors.Is 判断。
var (
	// ErrClosed 表示容器已关闭，不再接受任何操作。
	ErrClosed = errors.New("wg: 容器已关闭")
	// ErrNilForm 表示传入的窗体为 nil（无窗口模式下应使用 NewHeadless）。
	ErrNilForm = errors.New("wg: 窗体为 nil")
	// ErrNotReady 表示该能力在当前里程碑尚未接入。
	//
	// 这是**故意**的错误而不是空实现：空实现会让调用方以为已经生效，
	// 例如 Run() 静默返回会让窗口一帧都不画、却看不出问题。
	ErrNotReady = errors.New("wg: 该能力尚未接入（见 docs/wg-engine/05-API-ROADMAP.md §5）")
)
