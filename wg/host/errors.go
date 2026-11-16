package host

import "errors"

// Surface 树的错误。全部为哨兵值，便于上层用 errors.Is 判断。
var (
	// ErrNilSurface 表示传入 nil 画布。
	ErrNilSurface = errors.New("wg/host: 画布为 nil")
	// ErrSurfaceCycle 表示挂接会形成环（自己挂自己，或把祖先挂到后代）。
	ErrSurfaceCycle = errors.New("wg/host: 挂接会形成环")
	// ErrSurfaceAttached 表示该画布已经挂在另一个父上，必须先 Detach。
	ErrSurfaceAttached = errors.New("wg/host: 画布已挂在其它父上")
)
