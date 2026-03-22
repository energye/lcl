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

// ISynBaseHighlighter Parent: ISynCustomHighlighter
type ISynBaseHighlighter interface {
	ISynCustomHighlighter
	SetOnGetInstanceLanguageName(fn TSynBaseHighlighterRStrEvent)         // property event
	SetOnGetAttribCount(fn TSynBaseHighlighterRIntEvent)                  // property event
	SetOnGetAttribute(fn TSynBaseHighlighterIIntRHAttrEvent)              // property event
	SetOnGetDefaultAttribute(fn TSynBaseHighlighterIIntRHAttrEvent)       // property event
	SetOnGetDefaultFilter(fn TSynBaseHighlighterRStrEvent)                // property event
	SetOnGetIdentChars(fn TSynBaseHighlighterRIdeCharEvent)               // property event
	SetOnSetWordBreakChars(fn TSynBaseHighlighterIIdeCharEvent)           // property event
	SetOnGetSampleSource(fn TSynBaseHighlighterRStrEvent)                 // property event
	SetOnIsFilterStored(fn TSynBaseHighlighterRBoolEvent)                 // property event
	SetOnSetDefaultFilter(fn TSynBaseHighlighterIStrEvent)                // property event
	SetOnSetSampleSource(fn TSynBaseHighlighterIStrEvent)                 // property event
	SetOnGetRangeIdentifier(fn TSynBaseHighlighterRPtrEvent)              // property event
	SetOnCreateRangeList(fn TSynBaseHighlighterIEsbRHrlCharEvent)         // property event
	SetOnAfterAttachedToRangeList(fn TSynBaseHighlighterIHrlCharEvent)    // property event
	SetOnBeforeDetachedFromRangeList(fn TSynBaseHighlighterIHrlCharEvent) // property event
	SetOnUpdateRangeInfoAtLine(fn TSynBaseHighlighterIIntRBoolEvent)      // property event
	SetOnDoCurrentLinesChanged(fn TSynBaseHighlighterEvent)               // property event
	SetOnGetDrawDivider(fn TSynBaseHighlighterIIntRDdcsEvent)             // property event
	SetOnGetDividerDrawConfig(fn TSynBaseHighlighterIIntRDdcEvent)        // property event
	SetOnGetDividerDrawConfigCount(fn TSynBaseHighlighterRIntEvent)       // property event
	SetOnPerformScan(fn TSynBaseHighlighterIIntBoolRIntEvent)             // property event
	SetOnGetEol(fn TSynBaseHighlighterRBoolEvent)                         // property event
	SetOnGetRange(fn TSynBaseHighlighterRPtrEvent)                        // property event
	SetOnGetToken(fn TSynBaseHighlighterRStrEvent)                        // property event
	SetOnGetTokenEx(fn TSynBaseHighlighterOStrIntEvent)                   // property event
	SetOnGetEndOfLineAttribute(fn TSynBaseHighlighterRHAttrEvent)         // property event
	SetOnGetTokenAttribute(fn TSynBaseHighlighterRHAttrEvent)             // property event
	SetOnGetTokenKind(fn TSynBaseHighlighterRIntEvent)                    // property event
	SetOnGetTokenPos(fn TSynBaseHighlighterRIntEvent)                     // property event
	SetOnGetTokenLen(fn TSynBaseHighlighterRIntEvent)                     // property event
	SetOnIsKeyword(fn TSynBaseHighlighterIStrRBoolEvent)                  // property event
	SetOnNext(fn TSynBaseHighlighterEvent)                                // property event
	SetOnStartAtLineIndex(fn TSynBaseHighlighterIIntEvent)                // property event
	SetOnSetRange(fn TSynBaseHighlighterIPtrEvent)                        // property event
	SetOnResetRange(fn TSynBaseHighlighterEvent)                          // property event
	SetOnSetLine(fn TSynBaseHighlighterIStrIntEvent)                      // property event
	SetOnUseUserSettings(fn TSynBaseHighlighterIIntRBoolEvent)            // property event
	SetOnEnumUserSettings(fn TSynBaseHighlighterIStrsEvent)               // property event
}

type TSynBaseHighlighter struct {
	TSynCustomHighlighter
}

func (m *TSynBaseHighlighter) SetOnGetInstanceLanguageName(fn TSynBaseHighlighterRStrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRStrEvent(fn)
	base.SetEvent(m, 1, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetAttribCount(fn TSynBaseHighlighterRIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRIntEvent(fn)
	base.SetEvent(m, 2, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetAttribute(fn TSynBaseHighlighterIIntRHAttrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntRHAttrEvent(fn)
	base.SetEvent(m, 3, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetDefaultAttribute(fn TSynBaseHighlighterIIntRHAttrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntRHAttrEvent(fn)
	base.SetEvent(m, 4, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetDefaultFilter(fn TSynBaseHighlighterRStrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRStrEvent(fn)
	base.SetEvent(m, 5, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetIdentChars(fn TSynBaseHighlighterRIdeCharEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRIdeCharEvent(fn)
	base.SetEvent(m, 6, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnSetWordBreakChars(fn TSynBaseHighlighterIIdeCharEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIdeCharEvent(fn)
	base.SetEvent(m, 7, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetSampleSource(fn TSynBaseHighlighterRStrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRStrEvent(fn)
	base.SetEvent(m, 8, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnIsFilterStored(fn TSynBaseHighlighterRBoolEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRBoolEvent(fn)
	base.SetEvent(m, 9, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnSetDefaultFilter(fn TSynBaseHighlighterIStrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIStrEvent(fn)
	base.SetEvent(m, 10, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnSetSampleSource(fn TSynBaseHighlighterIStrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIStrEvent(fn)
	base.SetEvent(m, 11, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetRangeIdentifier(fn TSynBaseHighlighterRPtrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRPtrEvent(fn)
	base.SetEvent(m, 12, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnCreateRangeList(fn TSynBaseHighlighterIEsbRHrlCharEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIEsbRHrlCharEvent(fn)
	base.SetEvent(m, 13, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnAfterAttachedToRangeList(fn TSynBaseHighlighterIHrlCharEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIHrlCharEvent(fn)
	base.SetEvent(m, 14, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnBeforeDetachedFromRangeList(fn TSynBaseHighlighterIHrlCharEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIHrlCharEvent(fn)
	base.SetEvent(m, 15, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnUpdateRangeInfoAtLine(fn TSynBaseHighlighterIIntRBoolEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntRBoolEvent(fn)
	base.SetEvent(m, 16, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnDoCurrentLinesChanged(fn TSynBaseHighlighterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterEvent(fn)
	base.SetEvent(m, 17, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetDrawDivider(fn TSynBaseHighlighterIIntRDdcsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntRDdcsEvent(fn)
	base.SetEvent(m, 18, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetDividerDrawConfig(fn TSynBaseHighlighterIIntRDdcEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntRDdcEvent(fn)
	base.SetEvent(m, 19, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetDividerDrawConfigCount(fn TSynBaseHighlighterRIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRIntEvent(fn)
	base.SetEvent(m, 20, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnPerformScan(fn TSynBaseHighlighterIIntBoolRIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntBoolRIntEvent(fn)
	base.SetEvent(m, 21, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetEol(fn TSynBaseHighlighterRBoolEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRBoolEvent(fn)
	base.SetEvent(m, 22, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetRange(fn TSynBaseHighlighterRPtrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRPtrEvent(fn)
	base.SetEvent(m, 23, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetToken(fn TSynBaseHighlighterRStrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRStrEvent(fn)
	base.SetEvent(m, 24, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetTokenEx(fn TSynBaseHighlighterOStrIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterOStrIntEvent(fn)
	base.SetEvent(m, 25, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetEndOfLineAttribute(fn TSynBaseHighlighterRHAttrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRHAttrEvent(fn)
	base.SetEvent(m, 26, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetTokenAttribute(fn TSynBaseHighlighterRHAttrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRHAttrEvent(fn)
	base.SetEvent(m, 27, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetTokenKind(fn TSynBaseHighlighterRIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRIntEvent(fn)
	base.SetEvent(m, 28, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetTokenPos(fn TSynBaseHighlighterRIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRIntEvent(fn)
	base.SetEvent(m, 29, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnGetTokenLen(fn TSynBaseHighlighterRIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterRIntEvent(fn)
	base.SetEvent(m, 30, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnIsKeyword(fn TSynBaseHighlighterIStrRBoolEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIStrRBoolEvent(fn)
	base.SetEvent(m, 31, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnNext(fn TSynBaseHighlighterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterEvent(fn)
	base.SetEvent(m, 32, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnStartAtLineIndex(fn TSynBaseHighlighterIIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntEvent(fn)
	base.SetEvent(m, 33, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnSetRange(fn TSynBaseHighlighterIPtrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIPtrEvent(fn)
	base.SetEvent(m, 34, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnResetRange(fn TSynBaseHighlighterEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterEvent(fn)
	base.SetEvent(m, 35, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnSetLine(fn TSynBaseHighlighterIStrIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIStrIntEvent(fn)
	base.SetEvent(m, 36, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnUseUserSettings(fn TSynBaseHighlighterIIntRBoolEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIIntRBoolEvent(fn)
	base.SetEvent(m, 37, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynBaseHighlighter) SetOnEnumUserSettings(fn TSynBaseHighlighterIStrsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBaseHighlighterIStrsEvent(fn)
	base.SetEvent(m, 38, synBaseHighlighterAPI(), api.MakeEventDataPtr(cb))
}

// NewSynBaseHighlighter class constructor
func NewSynBaseHighlighter(owner IComponent) ISynBaseHighlighter {
	r := synBaseHighlighterAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynBaseHighlighter(r)
}

func TSynBaseHighlighterClass() types.TClass {
	r := synBaseHighlighterAPI().SysCallN(39)
	return types.TClass(r)
}

var (
	synBaseHighlighterOnce   base.Once
	synBaseHighlighterImport *imports.Imports = nil
)

func synBaseHighlighterAPI() *imports.Imports {
	synBaseHighlighterOnce.Do(func() {
		synBaseHighlighterImport = api.NewDefaultImports()
		synBaseHighlighterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynBaseHighlighter_Create", 0), // constructor NewSynBaseHighlighter
			/* 1 */ imports.NewTable("TSynBaseHighlighter_OnGetInstanceLanguageName", 0), // event OnGetInstanceLanguageName
			/* 2 */ imports.NewTable("TSynBaseHighlighter_OnGetAttribCount", 0), // event OnGetAttribCount
			/* 3 */ imports.NewTable("TSynBaseHighlighter_OnGetAttribute", 0), // event OnGetAttribute
			/* 4 */ imports.NewTable("TSynBaseHighlighter_OnGetDefaultAttribute", 0), // event OnGetDefaultAttribute
			/* 5 */ imports.NewTable("TSynBaseHighlighter_OnGetDefaultFilter", 0), // event OnGetDefaultFilter
			/* 6 */ imports.NewTable("TSynBaseHighlighter_OnGetIdentChars", 0), // event OnGetIdentChars
			/* 7 */ imports.NewTable("TSynBaseHighlighter_OnSetWordBreakChars", 0), // event OnSetWordBreakChars
			/* 8 */ imports.NewTable("TSynBaseHighlighter_OnGetSampleSource", 0), // event OnGetSampleSource
			/* 9 */ imports.NewTable("TSynBaseHighlighter_OnIsFilterStored", 0), // event OnIsFilterStored
			/* 10 */ imports.NewTable("TSynBaseHighlighter_OnSetDefaultFilter", 0), // event OnSetDefaultFilter
			/* 11 */ imports.NewTable("TSynBaseHighlighter_OnSetSampleSource", 0), // event OnSetSampleSource
			/* 12 */ imports.NewTable("TSynBaseHighlighter_OnGetRangeIdentifier", 0), // event OnGetRangeIdentifier
			/* 13 */ imports.NewTable("TSynBaseHighlighter_OnCreateRangeList", 0), // event OnCreateRangeList
			/* 14 */ imports.NewTable("TSynBaseHighlighter_OnAfterAttachedToRangeList", 0), // event OnAfterAttachedToRangeList
			/* 15 */ imports.NewTable("TSynBaseHighlighter_OnBeforeDetachedFromRangeList", 0), // event OnBeforeDetachedFromRangeList
			/* 16 */ imports.NewTable("TSynBaseHighlighter_OnUpdateRangeInfoAtLine", 0), // event OnUpdateRangeInfoAtLine
			/* 17 */ imports.NewTable("TSynBaseHighlighter_OnDoCurrentLinesChanged", 0), // event OnDoCurrentLinesChanged
			/* 18 */ imports.NewTable("TSynBaseHighlighter_OnGetDrawDivider", 0), // event OnGetDrawDivider
			/* 19 */ imports.NewTable("TSynBaseHighlighter_OnGetDividerDrawConfig", 0), // event OnGetDividerDrawConfig
			/* 20 */ imports.NewTable("TSynBaseHighlighter_OnGetDividerDrawConfigCount", 0), // event OnGetDividerDrawConfigCount
			/* 21 */ imports.NewTable("TSynBaseHighlighter_OnPerformScan", 0), // event OnPerformScan
			/* 22 */ imports.NewTable("TSynBaseHighlighter_OnGetEol", 0), // event OnGetEol
			/* 23 */ imports.NewTable("TSynBaseHighlighter_OnGetRange", 0), // event OnGetRange
			/* 24 */ imports.NewTable("TSynBaseHighlighter_OnGetToken", 0), // event OnGetToken
			/* 25 */ imports.NewTable("TSynBaseHighlighter_OnGetTokenEx", 0), // event OnGetTokenEx
			/* 26 */ imports.NewTable("TSynBaseHighlighter_OnGetEndOfLineAttribute", 0), // event OnGetEndOfLineAttribute
			/* 27 */ imports.NewTable("TSynBaseHighlighter_OnGetTokenAttribute", 0), // event OnGetTokenAttribute
			/* 28 */ imports.NewTable("TSynBaseHighlighter_OnGetTokenKind", 0), // event OnGetTokenKind
			/* 29 */ imports.NewTable("TSynBaseHighlighter_OnGetTokenPos", 0), // event OnGetTokenPos
			/* 30 */ imports.NewTable("TSynBaseHighlighter_OnGetTokenLen", 0), // event OnGetTokenLen
			/* 31 */ imports.NewTable("TSynBaseHighlighter_OnIsKeyword", 0), // event OnIsKeyword
			/* 32 */ imports.NewTable("TSynBaseHighlighter_OnNext", 0), // event OnNext
			/* 33 */ imports.NewTable("TSynBaseHighlighter_OnStartAtLineIndex", 0), // event OnStartAtLineIndex
			/* 34 */ imports.NewTable("TSynBaseHighlighter_OnSetRange", 0), // event OnSetRange
			/* 35 */ imports.NewTable("TSynBaseHighlighter_OnResetRange", 0), // event OnResetRange
			/* 36 */ imports.NewTable("TSynBaseHighlighter_OnSetLine", 0), // event OnSetLine
			/* 37 */ imports.NewTable("TSynBaseHighlighter_OnUseUserSettings", 0), // event OnUseUserSettings
			/* 38 */ imports.NewTable("TSynBaseHighlighter_OnEnumUserSettings", 0), // event OnEnumUserSettings
			/* 39 */ imports.NewTable("TSynBaseHighlighter_TClass", 0), // function TSynBaseHighlighterClass
		}
	})
	return synBaseHighlighterImport
}
