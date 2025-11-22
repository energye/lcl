//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// IStreamAdapter Parent: IInterfacedObject
type IStreamAdapter interface {
	IInterfacedObject
	Read(pv uintptr, cb types.DWORD, pcbRead types.PDWORD) types.HRESULT                      // function
	Write(pv uintptr, cb types.DWORD, pcbWritten types.PDWORD) types.HRESULT                  // function
	Seek(dlibMove int64, dwOrigin types.DWORD, outLibNewPosition *int64) types.HRESULT        // function
	SetSize(libNewSize int64) types.HRESULT                                                   // function
	CopyTo(stm IStreamAdapter, cb int64, outCbRead *int64, outCbWritten *int64) types.HRESULT // function
	Commit(grfCommitFlags types.DWORD) types.HRESULT                                          // function
	Revert() types.HRESULT                                                                    // function
	LockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT               // function
	UnlockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT             // function
	Stat(outStatstg *TStatStg, grfStatFlag types.DWORD) types.HRESULT                         // function
	Clone(outStm *IStreamAdapter) types.HRESULT                                               // function
	Stream() IStream                                                                          // property Stream Getter
	StreamOwnership() types.TStreamOwnership                                                  // property StreamOwnership Getter
	SetStreamOwnership(value types.TStreamOwnership)                                          // property StreamOwnership Setter
	AsIntfStream() uintptr
}

type TStreamAdapter struct {
	TInterfacedObject
}

func (m *TStreamAdapter) Read(pv uintptr, cb types.DWORD, pcbRead types.PDWORD) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(1, m.Instance(), uintptr(pv), uintptr(cb), uintptr(pcbRead))
	return types.HRESULT(r)
}

func (m *TStreamAdapter) Write(pv uintptr, cb types.DWORD, pcbWritten types.PDWORD) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(2, m.Instance(), uintptr(pv), uintptr(cb), uintptr(pcbWritten))
	return types.HRESULT(r)
}

func (m *TStreamAdapter) Seek(dlibMove int64, dwOrigin types.DWORD, outLibNewPosition *int64) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var libNewPositionPtr uintptr
	r := streamAdapterAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&dlibMove)), uintptr(dwOrigin), uintptr(base.UnsafePointer(&libNewPositionPtr)))
	*outLibNewPosition = int64(libNewPositionPtr)
	return types.HRESULT(r)
}

func (m *TStreamAdapter) SetSize(libNewSize int64) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&libNewSize)))
	return types.HRESULT(r)
}

func (m *TStreamAdapter) CopyTo(stm IStreamAdapter, cb int64, outCbRead *int64, outCbWritten *int64) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var cbReadPtr uintptr
	var cbWrittenPtr uintptr
	r := streamAdapterAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(stm), uintptr(base.UnsafePointer(&cb)), uintptr(base.UnsafePointer(&cbReadPtr)), uintptr(base.UnsafePointer(&cbWrittenPtr)))
	*outCbRead = int64(cbReadPtr)
	*outCbWritten = int64(cbWrittenPtr)
	return types.HRESULT(r)
}

func (m *TStreamAdapter) Commit(grfCommitFlags types.DWORD) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(6, m.Instance(), uintptr(grfCommitFlags))
	return types.HRESULT(r)
}

func (m *TStreamAdapter) Revert() types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(7, m.Instance())
	return types.HRESULT(r)
}

func (m *TStreamAdapter) LockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&libOffset)), uintptr(base.UnsafePointer(&cb)), uintptr(dwLockType))
	return types.HRESULT(r)
}

func (m *TStreamAdapter) UnlockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&libOffset)), uintptr(base.UnsafePointer(&cb)), uintptr(dwLockType))
	return types.HRESULT(r)
}

func (m *TStreamAdapter) Stat(outStatstg *TStatStg, grfStatFlag types.DWORD) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	statstgPtr := &tStatStg{}
	r := streamAdapterAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(statstgPtr)), uintptr(grfStatFlag))
	*outStatstg = statstgPtr.ToGo()
	return types.HRESULT(r)
}

func (m *TStreamAdapter) Clone(outStm *IStreamAdapter) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var stmPtr uintptr
	r := streamAdapterAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&stmPtr)))
	*outStm = AsStreamAdapter(stmPtr)
	return types.HRESULT(r)
}

func (m *TStreamAdapter) Stream() IStream {
	if !m.IsValid() {
		return nil
	}
	r := streamAdapterAPI().SysCallN(12, m.Instance())
	return AsStream(r)
}

func (m *TStreamAdapter) StreamOwnership() types.TStreamOwnership {
	if !m.IsValid() {
		return 0
	}
	r := streamAdapterAPI().SysCallN(13, 0, m.Instance())
	return types.TStreamOwnership(r)
}

func (m *TStreamAdapter) SetStreamOwnership(value types.TStreamOwnership) {
	if !m.IsValid() {
		return
	}
	streamAdapterAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TStreamAdapter) AsIntfStream() uintptr {
	return m.GetIntfPointer(0)
}

// NewStreamAdapter class constructor
func NewStreamAdapter(stream IStream, ownership types.TStreamOwnership) IStreamAdapter {
	var streamPtr uintptr // IStream
	r := streamAdapterAPI().SysCallN(0, base.GetObjectUintptr(stream), uintptr(ownership), uintptr(base.UnsafePointer(&streamPtr)))
	ret := AsStreamAdapter(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, streamPtr)
	}
	return ret
}

var (
	streamAdapterOnce   base.Once
	streamAdapterImport *imports.Imports = nil
)

func streamAdapterAPI() *imports.Imports {
	streamAdapterOnce.Do(func() {
		streamAdapterImport = api.NewDefaultImports()
		streamAdapterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStreamAdapter_Create", 0), // constructor NewStreamAdapter
			/* 1 */ imports.NewTable("TStreamAdapter_Read", 0), // function Read
			/* 2 */ imports.NewTable("TStreamAdapter_Write", 0), // function Write
			/* 3 */ imports.NewTable("TStreamAdapter_Seek", 0), // function Seek
			/* 4 */ imports.NewTable("TStreamAdapter_SetSize", 0), // function SetSize
			/* 5 */ imports.NewTable("TStreamAdapter_CopyTo", 0), // function CopyTo
			/* 6 */ imports.NewTable("TStreamAdapter_Commit", 0), // function Commit
			/* 7 */ imports.NewTable("TStreamAdapter_Revert", 0), // function Revert
			/* 8 */ imports.NewTable("TStreamAdapter_LockRegion", 0), // function LockRegion
			/* 9 */ imports.NewTable("TStreamAdapter_UnlockRegion", 0), // function UnlockRegion
			/* 10 */ imports.NewTable("TStreamAdapter_Stat", 0), // function Stat
			/* 11 */ imports.NewTable("TStreamAdapter_Clone", 0), // function Clone
			/* 12 */ imports.NewTable("TStreamAdapter_Stream", 0), // property Stream
			/* 13 */ imports.NewTable("TStreamAdapter_StreamOwnership", 0), // property StreamOwnership
		}
	})
	return streamAdapterImport
}
