//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import (
	"sync"

	"github.com/energye/lcl/api"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/tool/ptr"
	"github.com/energye/lcl/types"
)

var (
	createParamsMap sync.Map
	newFormCreate   IForm
	queueAsyncCall  = _QueueAsyncCall{id: 0, calls: sync.Map{}}
)

// TMainThreadSyncProc 主线程运行回调函数
type TMainThreadSyncProc func()

// TMainThreadAsyncProc 主线程运行回调函数
type TMainThreadAsyncProc func(id uint32)

// Inserts an asynchronous call into the queue. This is thread safe.
// Errors Exceptions on various inconsistencies of the queue.
type _QueueAsyncCall struct {
	id    int
	calls sync.Map
}

func (m *_QueueAsyncCall) call(id int) {
	if call, ok := m.calls.Load(id); ok {
		m.calls.Delete(id)
		call.(TMainThreadAsyncProc)(uint32(id))
	}
}
func (m *_QueueAsyncCall) next(fn TMainThreadAsyncProc) int {
	// TODO 这里需要验证一下个 id 是否已被处理, 在累加到最大 int 数值后会被重置
	m.id++
	m.calls.Store(m.id, fn)
	return m.id
}

func addNewFormCreate(form IForm) {
	newFormCreate = form
}

func addRequestCreateParamsMap(ptr uintptr, proc IForm) {
	createParamsMap.Store(ptr, proc)
}

func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	fn := api.PtrToElementValue(f)
	if fn != nil {
		if cb, ok := fn.(*callback); ok {
			defer func() {
				//if err := recover(); err != nil {
				//	// TODO 增加回调到用户
				//	switch err.(type) {
				//	case error:
				//		println("[ERROR] CallbackEvent:", cb.name, err.(error).Error())
				//	case string:
				//		println("[ERROR] CallbackEvent:", cb.name, err.(string))
				//	default:
				//		println("[ERROR] CallbackEvent:", cb.name, err)
				//	}
				//}
			}()
			getVal := func(i int) uintptr {
				return ptr.GetParamOf(i, args)
			}
			cb.cb(getVal)
		}
	}
	return 0
}

// 移除事件，释放相关的引用
func removeEventCallbackProc(f uintptr) uintptr {
	api.RemoveEventElement(f)
	return 0
}

// 消息回调处理
func messageCallbackProc(f uintptr, msg uintptr) uintptr {
	fn := api.PtrToElementValue(f)
	if fn != nil {
		if cb, ok := fn.(func(msg uintptr)); ok {
			cb(msg)
		}
	}
	return 0
}

// UI 线程执行回调 同步
func threadSyncCallbackProc() uintptr {
	fn := api.ThreadSyncCallbackFn()
	if fn != nil {
		fn()
	}
	return 0
}

// UI 线程执行回调 异步
func threadAsyncCallbackProc(data uintptr) uintptr {
	id := *(*int)(base.UnsafePointer(data))
	queueAsyncCall.call(id)
	return 0
}

// 控件创建时 CreateParams 函数回调
func createParamsCallbackProc(ptr uintptr, sender, params uintptr) uintptr {
	if val, ok := createParamsMap.Load(ptr); ok {
		createParamsMap.Delete(ptr)
		if form, ok := val.(IForm); ok {
			form.SetInstance(base.UnsafePointer(sender))
			if impl, ok := form.(IOnCreateParams); ok {
				impl.CreateParams((*types.TCreateParams)(base.UnsafePointer(params)))
			}
		}
	}
	return 0
}

// 窗口创建时 FormCreate 函数回调
func formCreateCallbackProc(ptr uintptr, sender uintptr) uintptr {
	if newFormCreate != nil {
		currentForm := newFormCreate
		newFormCreate = nil
		currentForm.SetInstance(base.UnsafePointer(sender))
		if impl, ok := currentForm.(IOnCreate); ok {
			impl.FormCreate(currentForm)
		}
	}
	return 0
}
