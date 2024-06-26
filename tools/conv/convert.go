//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package conv
package conv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"unsafe"
)

const (
	IntSize   = strconv.IntSize //bit
	IntSize32 = 32              //
	intSize64 = 64              //
)

func StrToInt64(value string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}
func StrToInt32(value string) int32 {
	v, _ := strconv.ParseInt(value, 10, 32)
	return int32(v)
}

func StrToFloat64(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func StrToFloat32(value string) float32 {
	v, _ := strconv.ParseFloat(value, 32)
	return float32(v)
}

// InterfaceToString 接口转 string
func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func IntToBytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	if IntSize == IntSize32 {
		if err := binary.Write(buf, binary.BigEndian, int32(i)); err == nil {
			return buf.Bytes()
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, int64(i)); err == nil {
			return buf.Bytes()
		}
	}
	return nil
}

func UIntToBytes(i uint) []byte {
	buf := bytes.NewBuffer([]byte{})
	if IntSize == IntSize32 {
		if err := binary.Write(buf, binary.BigEndian, uint32(i)); err == nil {
			return buf.Bytes()
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, uint64(i)); err == nil {
			return buf.Bytes()
		}
	}
	return nil
}

func Int8ToBytes(i int8) []byte {
	return []byte{byte(i)}
}

func UInt8ToBytes(i uint8) []byte {
	return []byte{byte(i)}
}

func Int16ToBytes(i int16) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func UInt16ToBytes(i uint16) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func Int32ToBytes(i int32) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func UInt32ToBytes(i uint32) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func Int64ToBytes(i int64) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func UInt64ToBytes(i uint64) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func BytesToInt(b []byte) int {
	var i int64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return int(i)
}

func BytesToUInt(b []byte) uint {
	var i uint64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return uint(i)
}

func ByteToInt8(b byte) int8 {
	return int8(b)
}

func ByteToUInt8(b byte) uint8 {
	return uint8(b)
}

func BytesToInt16(b []byte) int16 {
	var i int16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToUInt16(b []byte) uint16 {
	var i uint16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToInt32(b []byte) int32 {
	var i int32
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToUInt32(b []byte) uint32 {
	var i uint32
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToInt64(b []byte) int64 {
	var i int64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToUInt64(b []byte) uint64 {
	var i uint64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

//func StringToBytes(data string) []byte {
//	return *(*[]byte)(unsafe.Pointer(&data))
//}

// String转换Bytes数组，isDStr转换DString 默认GoString
func StringToBytes(s string, isDStr ...bool) []byte {
	if len(isDStr) > 0 && isDStr[0] {
		temp := []byte(s)
		utf8StrArr := make([]byte, len(temp)+1)
		copy(utf8StrArr, temp)
		return utf8StrArr
	} else {
		return []byte(s)
	}
}

// Float64ToBytes Float64转byte
func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// BytesToFloat64 byte转Float64
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

// Float32ToBytes Float64转byte
func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

// BytesToFloat32 byte转Float64
func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func ByteToBool(b byte) bool {
	if b == 1 {
		return true
	}
	return false
}

const (
	dSecond float64 = 3600
	dDay            = 24 * dSecond
)

// ArrayIndexOf 返回a在array数组的下标, a必须和array元素的类型相同
func ArrayIndexOf(array interface{}, a interface{}) int {
	switch array.(type) {
	case []string:
		arrs := array.([]string)
		if len(arrs) == 0 {
			return -1
		}
		eq := a.(string)
		for i, v := range arrs {
			if v == eq {
				return i
			}
		}
		//case []int: //其它待增加
	}
	return -1
}

// GetInstancePtr 获取指针的指针的地址
func GetInstancePtr(ptr uintptr) unsafe.Pointer {
	ptr = *(*uintptr)(unsafe.Pointer(ptr))
	return unsafe.Pointer(ptr)
}

// GoroutineID 获取当前携程ID
func GoroutineID() (id uint64) {
	var buf [30]byte
	runtime.Stack(buf[:], false)
	for i := 10; buf[i] != ' '; i++ {
		id = id*10 + uint64(buf[i]&15)
	}
	return id
}

type StringHeader struct {
	Data uintptr
	Len  int
}

func GoStr(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	resultString := (*StringHeader)(unsafe.Pointer(ptr))
	if resultString == nil || resultString.Len <= 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(resultString))
}
