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
	"unsafe"

	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"

	"github.com/energye/lcl/base"
)

// StreamHelper IStream helper class
var StreamHelper _StreamHelper

// IStream helper class
type _StreamHelper uintptr

// ReadBuffer is IStream.ReadBuffer helper func
func (_StreamHelper) ReadBuffer(stream IStream, count int) []byte {
	if stream == nil || count <= 0 {
		return nil
	}
	buffer := make([]byte, count)
	bufferPtr := uintptr(base.UnsafePointer(&buffer[0]))
	stream.ReadBuffer(&bufferPtr, int32(count))
	return buffer
}

// Read is IStream.Read helper func
func (_StreamHelper) Read(stream IStream, count int) []byte {
	if stream == nil || count <= 0 {
		return nil
	}
	buffer := make([]byte, count)
	bufferPtr := uintptr(base.UnsafePointer(&buffer[0]))
	stream.Read(&bufferPtr, int32(count))
	return buffer
}

// WriteBuffer is IStream.WriteBuffer helper func
func (_StreamHelper) WriteBuffer(stream IStream, buffer []byte) {
	if stream == nil || len(buffer) == 0 {
		return
	}
	stream.WriteBuffer(uintptr(unsafe.Pointer(&buffer[0])), int32(len(buffer)))
}

// Write is IStream.Write helper func
func (_StreamHelper) Write(stream IStream, buffer []byte) {
	if stream == nil || len(buffer) == 0 {
		return
	}
	stream.Write(uintptr(unsafe.Pointer(&buffer[0])), int32(len(buffer)))
}

// ClipboardHelper IClipboard helper class
var ClipboardHelper _ClipboardHelper

// IClipboard helper class
type _ClipboardHelper uintptr

// GetTextBuf is IClipboard.GetTextBuf  helper func
func (_ClipboardHelper) GetTextBuf(clipboard IClipboard, buffer *string, bufSize int32) int32 {
	if clipboard == nil || buffer == nil || bufSize == 0 {
		return 0
	}
	bufferPtr := make([]uint8, bufSize+1)
	sLen := clipboard.GetTextBuf(uintptr(unsafe.Pointer(&bufferPtr[0])), bufSize)
	if sLen > 0 {
		*buffer = string(bufferPtr[:sLen])
	}
	return 0
}

// ControlHelper IClipboard helper class
var ControlHelper _ControlHelper

// IControl helper class
type _ControlHelper uintptr

// GetTextBuf is IControl.GetTextBuf  helper func
func (_ControlHelper) GetTextBuf(clipboard IControl, buffer *string, bufSize int32) int32 {
	if clipboard == nil || buffer == nil || bufSize == 0 {
		return 0
	}
	bufferPtr := make([]uint8, bufSize+1)
	sLen := clipboard.GetTextBuf(uintptr(unsafe.Pointer(&bufferPtr[0])), bufSize)
	if sLen > 0 {
		*buffer = string(bufferPtr[:sLen])
	}
	return 0
}

// WindowProc  The handler for all messages.
// WindowProc is set to TControl.WndProc by default.
// When you ever change WindowProc, then remember the previous setting and delegate all unhandled messages to that method.
func (_ControlHelper) WindowProc(control IControl, message *types.TLMessage) {
	api.ControlWindowProc(control.Instance(), message)
}
