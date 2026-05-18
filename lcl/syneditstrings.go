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

// ISynEditStrings Parent: ISynEditStringsBase
type ISynEditStrings interface {
	ISynEditStringsBase
	IsUpdating() bool // function
	// LogicPosAddCharsWithStrIntX2LPFlags
	//  Char bounds // 1 based (1 is the 1st char in the line)
	LogicPosAddCharsWithStrIntX2LPFlags(line string, logicalPos int32, count int32, flags types.LPosFlags) int32 // function
	LogicPosIsAtChar(line string, logicalPos int32, flags types.LPosFlags) bool                                  // function
	LogicPosAdjustToCharWithStrIntLPFlags(line string, logicalPos int32, flags types.LPosFlags) int32            // function
	LogicPosAddCharsWithStrIntX2Bool(line string, logicalPos int32, count int32, allowPastEOL bool) int32        // function
	LogicPosAdjustToCharWithStrIntBoolX2(line string, logicalPos int32, next bool, allowPastEOL bool) int32      // function
	// GetPhysicalCharWidthsWithInt
	//  CharWidths
	GetPhysicalCharWidthsWithInt(index int32) IByteArrayWrap                                   // function
	GetPhysicalCharWidthsWithStrIntX2(line uintptr, lineLen int32, index int32) IByteArrayWrap // function
	// LogicalToPhysicalPos
	//  Byte to Char
	LogicalToPhysicalPos(P types.TPoint) types.TPoint                      // function
	LogicalToPhysicalCol(line string, index int32, logicalPos int32) int32 // function
	// PhysicalToLogicalPos
	//  Char to Byte
	PhysicalToLogicalPos(P types.TPoint) types.TPoint                                      // function
	PhysicalToLogicalCol(line string, index int32, physicalPos int32) int32                // function
	TextToViewIndex(textIndex types.TLineIdx) types.TLineIdx                               // function
	ViewToTextIndex(viewIndex types.TLineIdx) types.TLineIdx                               // function
	AddVisibleOffsetToTextIndex(textIndex types.TLineIdx, lineOffset int32) types.TLineIdx // function
	IsTextIdxVisible(textIndex types.TLineIdx) bool                                        // function
	// ViewXYToTextXY
	//  ViewedToPhysAndLog
	//  * Convert between TextBuffer and ViewedText
	//  X/Y are all 1-based
	ViewXYToTextXY(physViewXY types.TRect) types.TRect                     // function
	TextXYToViewXY(physTextXY types.TRect) types.TRect                     // function
	EditDelete(logX int32, logY int32, byteLen int32) string               // function
	EditReplace(logX int32, logY int32, byteLen int32, text string) string // function
	BeginUpdateWithObject(sender IObject)                                  // procedure
	EndUpdateWithObject(sender IObject)                                    // procedure
	// DeleteLines
	//  Currently Lines are physical
	DeleteLines(index int32, numLines int32)                                                                                                                                  // procedure
	InsertLines(index int32, numLines int32)                                                                                                                                  // procedure
	InsertStrings(index int32, newStrings IStrings)                                                                                                                           // procedure
	SendNotificationWithSENReasonSEStringsIntX2(reason types.TSynEditNotifyReason, sender ISynEditStrings, index int32, count int32)                                          // procedure
	SendNotificationWithSENReasonSEStringsIntX4Str(reason types.TSynEditNotifyReason, sender ISynEditStrings, index int32, count int32, bytePos int32, len int32, txt string) // procedure
	SendNotificationWithSynEditNotifyReasonObject(reason types.TSynEditNotifyReason, sender IObject)                                                                          // procedure
	FlushNotificationCache()                                                                                                                                                  // procedure
	GetPhysicalCharWidthsWithIntX2PCWidths(index int32, result *IByteArrayWrap, outLen *int32)                                                                                // procedure
	GetInfoForViewedXY(viewedXY types.TRect, flags types.TViewedXYInfoFlags, outViewedXYInfo *TViewedXYInfo)                                                                  // procedure
	// EditInsert
	//  Currently Lines are physical
	EditInsert(logX int32, logY int32, text string)         // procedure
	EditLineBreak(logX int32, logY int32)                   // procedure
	EditLineJoin(logY int32, fillText string)               // procedure
	EditLinesInsert(logY int32, count int32, text string)   // procedure
	EditLinesDelete(logY int32, count int32)                // procedure
	EditUndo(item ISynEditUndoItem)                         // procedure
	EditRedo(item ISynEditUndoItem)                         // procedure
	LogPhysConvertor() ISynLogicalPhysicalConvertor         // property LogPhysConvertor Getter
	SetLogPhysConvertor(value ISynLogicalPhysicalConvertor) // property LogPhysConvertor Setter
	// IsInEditAction
	//  * IsInEditAction
	//  While in EditAction a "senrEditAction" event is triggerred, that contains
	//  more detailed line-count change info.
	//  Yet senrLineCount is also sent.
	//  IsInEditAction allows one to skip senrLineCount, that are sent as senrEditAction
	//  Currently used by FoldView
	IsInEditAction() bool            // property IsInEditAction Getter
	UndoList() ISynEditUndoList      // property UndoList Getter
	RedoList() ISynEditUndoList      // property RedoList Getter
	CurUndoList() ISynEditUndoList   // property CurUndoList Getter
	IsUndoing() bool                 // property IsUndoing Getter
	SetIsUndoing(value bool)         // property IsUndoing Setter
	IsRedoing() bool                 // property IsRedoing Getter
	SetIsRedoing(value bool)         // property IsRedoing Setter
	TextChangeStamp() int64          // property TextChangeStamp Getter
	ViewChangeStamp() int64          // property ViewChangeStamp Getter
	LengthOfLongestLine() int32      // property LengthOfLongestLine Getter
	ViewedLines(index int32) string  // property ViewedLines Getter
	ViewedCount() int32              // property ViewedCount Getter
	IsUtf8() bool                    // property IsUtf8 Getter
	SetIsUtf8(value bool)            // property IsUtf8 Setter
	DisplayView() ILazSynDisplayView // property DisplayView Getter
}

type TSynEditStrings struct {
	TSynEditStringsBase
}

func (m *TSynEditStrings) IsUpdating() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditStringsAPI().SysCallN(0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditStrings) LogicPosAddCharsWithStrIntX2LPFlags(line string, logicalPos int32, count int32, flags types.LPosFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(1, m.Instance(), api.PasStr(line), uintptr(logicalPos), uintptr(count), uintptr(flags))
	return int32(r)
}

func (m *TSynEditStrings) LogicPosIsAtChar(line string, logicalPos int32, flags types.LPosFlags) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditStringsAPI().SysCallN(2, m.Instance(), api.PasStr(line), uintptr(logicalPos), uintptr(flags))
	return api.GoBool(r)
}

func (m *TSynEditStrings) LogicPosAdjustToCharWithStrIntLPFlags(line string, logicalPos int32, flags types.LPosFlags) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(3, m.Instance(), api.PasStr(line), uintptr(logicalPos), uintptr(flags))
	return int32(r)
}

func (m *TSynEditStrings) LogicPosAddCharsWithStrIntX2Bool(line string, logicalPos int32, count int32, allowPastEOL bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(4, m.Instance(), api.PasStr(line), uintptr(logicalPos), uintptr(count), api.PasBool(allowPastEOL))
	return int32(r)
}

func (m *TSynEditStrings) LogicPosAdjustToCharWithStrIntBoolX2(line string, logicalPos int32, next bool, allowPastEOL bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(5, m.Instance(), api.PasStr(line), uintptr(logicalPos), api.PasBool(next), api.PasBool(allowPastEOL))
	return int32(r)
}

func (m *TSynEditStrings) GetPhysicalCharWidthsWithInt(index int32) IByteArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsAPI().SysCallN(6, m.Instance(), uintptr(index))
	return AsByteArrayWrap(r)
}

func (m *TSynEditStrings) GetPhysicalCharWidthsWithStrIntX2(line uintptr, lineLen int32, index int32) IByteArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsAPI().SysCallN(7, m.Instance(), uintptr(line), uintptr(lineLen), uintptr(index))
	return AsByteArrayWrap(r)
}

func (m *TSynEditStrings) LogicalToPhysicalPos(P types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&P)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditStrings) LogicalToPhysicalCol(line string, index int32, logicalPos int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(9, m.Instance(), api.PasStr(line), uintptr(index), uintptr(logicalPos))
	return int32(r)
}

func (m *TSynEditStrings) PhysicalToLogicalPos(P types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&P)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditStrings) PhysicalToLogicalCol(line string, index int32, physicalPos int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(11, m.Instance(), api.PasStr(line), uintptr(index), uintptr(physicalPos))
	return int32(r)
}

func (m *TSynEditStrings) TextToViewIndex(textIndex types.TLineIdx) types.TLineIdx {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(12, m.Instance(), uintptr(textIndex))
	return types.TLineIdx(r)
}

func (m *TSynEditStrings) ViewToTextIndex(viewIndex types.TLineIdx) types.TLineIdx {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(13, m.Instance(), uintptr(viewIndex))
	return types.TLineIdx(r)
}

func (m *TSynEditStrings) AddVisibleOffsetToTextIndex(textIndex types.TLineIdx, lineOffset int32) types.TLineIdx {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(14, m.Instance(), uintptr(textIndex), uintptr(lineOffset))
	return types.TLineIdx(r)
}

func (m *TSynEditStrings) IsTextIdxVisible(textIndex types.TLineIdx) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditStringsAPI().SysCallN(15, m.Instance(), uintptr(textIndex))
	return api.GoBool(r)
}

func (m *TSynEditStrings) ViewXYToTextXY(physViewXY types.TRect) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&physViewXY)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditStrings) TextXYToViewXY(physTextXY types.TRect) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&physTextXY)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditStrings) EditDelete(logX int32, logY int32, byteLen int32) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	synEditStringsAPI().SysCallN(18, m.Instance(), uintptr(logX), uintptr(logY), uintptr(byteLen), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynEditStrings) EditReplace(logX int32, logY int32, byteLen int32, text string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	synEditStringsAPI().SysCallN(19, m.Instance(), uintptr(logX), uintptr(logY), uintptr(byteLen), api.PasStr(text), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynEditStrings) BeginUpdateWithObject(sender IObject) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TSynEditStrings) EndUpdateWithObject(sender IObject) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TSynEditStrings) DeleteLines(index int32, numLines int32) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(22, m.Instance(), uintptr(index), uintptr(numLines))
}

func (m *TSynEditStrings) InsertLines(index int32, numLines int32) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(23, m.Instance(), uintptr(index), uintptr(numLines))
}

func (m *TSynEditStrings) InsertStrings(index int32, newStrings IStrings) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(24, m.Instance(), uintptr(index), base.GetObjectUintptr(newStrings))
}

func (m *TSynEditStrings) SendNotificationWithSENReasonSEStringsIntX2(reason types.TSynEditNotifyReason, sender ISynEditStrings, index int32, count int32) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(25, m.Instance(), uintptr(reason), base.GetObjectUintptr(sender), uintptr(index), uintptr(count))
}

func (m *TSynEditStrings) SendNotificationWithSENReasonSEStringsIntX4Str(reason types.TSynEditNotifyReason, sender ISynEditStrings, index int32, count int32, bytePos int32, len int32, txt string) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(26, m.Instance(), uintptr(reason), base.GetObjectUintptr(sender), uintptr(index), uintptr(count), uintptr(bytePos), uintptr(len), api.PasStr(txt))
}

func (m *TSynEditStrings) SendNotificationWithSynEditNotifyReasonObject(reason types.TSynEditNotifyReason, sender IObject) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(27, m.Instance(), uintptr(reason), base.GetObjectUintptr(sender))
}

func (m *TSynEditStrings) FlushNotificationCache() {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(28, m.Instance())
}

func (m *TSynEditStrings) GetPhysicalCharWidthsWithIntX2PCWidths(index int32, result *IByteArrayWrap, outLen *int32) {
	if !m.IsValid() {
		return
	}
	resultPtr := base.GetObjectUintptr(*result)
	var lenPtr uintptr
	synEditStringsAPI().SysCallN(29, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&resultPtr)), uintptr(base.UnsafePointer(&lenPtr)))
	*result = AsByteArrayWrap(resultPtr)
	*outLen = int32(lenPtr)
}

func (m *TSynEditStrings) GetInfoForViewedXY(viewedXY types.TRect, flags types.TViewedXYInfoFlags, outViewedXYInfo *TViewedXYInfo) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(30, m.Instance(), uintptr(base.UnsafePointer(&viewedXY)), uintptr(flags), uintptr(base.UnsafePointer(outViewedXYInfo)))
}

func (m *TSynEditStrings) EditInsert(logX int32, logY int32, text string) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(31, m.Instance(), uintptr(logX), uintptr(logY), api.PasStr(text))
}

func (m *TSynEditStrings) EditLineBreak(logX int32, logY int32) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(32, m.Instance(), uintptr(logX), uintptr(logY))
}

func (m *TSynEditStrings) EditLineJoin(logY int32, fillText string) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(33, m.Instance(), uintptr(logY), api.PasStr(fillText))
}

func (m *TSynEditStrings) EditLinesInsert(logY int32, count int32, text string) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(34, m.Instance(), uintptr(logY), uintptr(count), api.PasStr(text))
}

func (m *TSynEditStrings) EditLinesDelete(logY int32, count int32) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(35, m.Instance(), uintptr(logY), uintptr(count))
}

func (m *TSynEditStrings) EditUndo(item ISynEditUndoItem) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(36, m.Instance(), base.GetObjectUintptr(item))
}

func (m *TSynEditStrings) EditRedo(item ISynEditUndoItem) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(37, m.Instance(), base.GetObjectUintptr(item))
}

func (m *TSynEditStrings) LogPhysConvertor() ISynLogicalPhysicalConvertor {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsAPI().SysCallN(38, 0, m.Instance())
	return AsSynLogicalPhysicalConvertor(r)
}

func (m *TSynEditStrings) SetLogPhysConvertor(value ISynLogicalPhysicalConvertor) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(38, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditStrings) IsInEditAction() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditStringsAPI().SysCallN(39, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditStrings) UndoList() ISynEditUndoList {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsAPI().SysCallN(40, m.Instance())
	return AsSynEditUndoList(r)
}

func (m *TSynEditStrings) RedoList() ISynEditUndoList {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsAPI().SysCallN(41, m.Instance())
	return AsSynEditUndoList(r)
}

func (m *TSynEditStrings) CurUndoList() ISynEditUndoList {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsAPI().SysCallN(42, m.Instance())
	return AsSynEditUndoList(r)
}

func (m *TSynEditStrings) IsUndoing() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditStringsAPI().SysCallN(43, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditStrings) SetIsUndoing(value bool) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(43, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditStrings) IsRedoing() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditStringsAPI().SysCallN(44, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditStrings) SetIsRedoing(value bool) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(44, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditStrings) TextChangeStamp() (result int64) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(45, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditStrings) ViewChangeStamp() (result int64) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(46, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditStrings) LengthOfLongestLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(47, m.Instance())
	return int32(r)
}

func (m *TSynEditStrings) ViewedLines(index int32) (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synEditStringsAPI().SysCallN(48, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynEditStrings) ViewedCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsAPI().SysCallN(49, m.Instance())
	return int32(r)
}

func (m *TSynEditStrings) IsUtf8() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditStringsAPI().SysCallN(50, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditStrings) SetIsUtf8(value bool) {
	if !m.IsValid() {
		return
	}
	synEditStringsAPI().SysCallN(50, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditStrings) DisplayView() ILazSynDisplayView {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsAPI().SysCallN(51, m.Instance())
	return AsLazSynDisplayView(r)
}

var (
	synEditStringsOnce   base.Once
	synEditStringsImport *imports.Imports = nil
)

func synEditStringsAPI() *imports.Imports {
	synEditStringsOnce.Do(func() {
		synEditStringsImport = api.NewDefaultImports()
		synEditStringsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditStrings_IsUpdating", 0), // function IsUpdating
			/* 1 */ imports.NewTable("TSynEditStrings_LogicPosAddCharsWithStrIntX2LPFlags", 0), // function LogicPosAddCharsWithStrIntX2LPFlags
			/* 2 */ imports.NewTable("TSynEditStrings_LogicPosIsAtChar", 0), // function LogicPosIsAtChar
			/* 3 */ imports.NewTable("TSynEditStrings_LogicPosAdjustToCharWithStrIntLPFlags", 0), // function LogicPosAdjustToCharWithStrIntLPFlags
			/* 4 */ imports.NewTable("TSynEditStrings_LogicPosAddCharsWithStrIntX2Bool", 0), // function LogicPosAddCharsWithStrIntX2Bool
			/* 5 */ imports.NewTable("TSynEditStrings_LogicPosAdjustToCharWithStrIntBoolX2", 0), // function LogicPosAdjustToCharWithStrIntBoolX2
			/* 6 */ imports.NewTable("TSynEditStrings_GetPhysicalCharWidthsWithInt", 0), // function GetPhysicalCharWidthsWithInt
			/* 7 */ imports.NewTable("TSynEditStrings_GetPhysicalCharWidthsWithStrIntX2", 0), // function GetPhysicalCharWidthsWithStrIntX2
			/* 8 */ imports.NewTable("TSynEditStrings_LogicalToPhysicalPos", 0), // function LogicalToPhysicalPos
			/* 9 */ imports.NewTable("TSynEditStrings_LogicalToPhysicalCol", 0), // function LogicalToPhysicalCol
			/* 10 */ imports.NewTable("TSynEditStrings_PhysicalToLogicalPos", 0), // function PhysicalToLogicalPos
			/* 11 */ imports.NewTable("TSynEditStrings_PhysicalToLogicalCol", 0), // function PhysicalToLogicalCol
			/* 12 */ imports.NewTable("TSynEditStrings_TextToViewIndex", 0), // function TextToViewIndex
			/* 13 */ imports.NewTable("TSynEditStrings_ViewToTextIndex", 0), // function ViewToTextIndex
			/* 14 */ imports.NewTable("TSynEditStrings_AddVisibleOffsetToTextIndex", 0), // function AddVisibleOffsetToTextIndex
			/* 15 */ imports.NewTable("TSynEditStrings_IsTextIdxVisible", 0), // function IsTextIdxVisible
			/* 16 */ imports.NewTable("TSynEditStrings_ViewXYToTextXY", 0), // function ViewXYToTextXY
			/* 17 */ imports.NewTable("TSynEditStrings_TextXYToViewXY", 0), // function TextXYToViewXY
			/* 18 */ imports.NewTable("TSynEditStrings_EditDelete", 0), // function EditDelete
			/* 19 */ imports.NewTable("TSynEditStrings_EditReplace", 0), // function EditReplace
			/* 20 */ imports.NewTable("TSynEditStrings_BeginUpdateWithObject", 0), // procedure BeginUpdateWithObject
			/* 21 */ imports.NewTable("TSynEditStrings_EndUpdateWithObject", 0), // procedure EndUpdateWithObject
			/* 22 */ imports.NewTable("TSynEditStrings_DeleteLines", 0), // procedure DeleteLines
			/* 23 */ imports.NewTable("TSynEditStrings_InsertLines", 0), // procedure InsertLines
			/* 24 */ imports.NewTable("TSynEditStrings_InsertStrings", 0), // procedure InsertStrings
			/* 25 */ imports.NewTable("TSynEditStrings_SendNotificationWithSENReasonSEStringsIntX2", 0), // procedure SendNotificationWithSENReasonSEStringsIntX2
			/* 26 */ imports.NewTable("TSynEditStrings_SendNotificationWithSENReasonSEStringsIntX4Str", 0), // procedure SendNotificationWithSENReasonSEStringsIntX4Str
			/* 27 */ imports.NewTable("TSynEditStrings_SendNotificationWithSynEditNotifyReasonObject", 0), // procedure SendNotificationWithSynEditNotifyReasonObject
			/* 28 */ imports.NewTable("TSynEditStrings_FlushNotificationCache", 0), // procedure FlushNotificationCache
			/* 29 */ imports.NewTable("TSynEditStrings_GetPhysicalCharWidthsWithIntX2PCWidths", 0), // procedure GetPhysicalCharWidthsWithIntX2PCWidths
			/* 30 */ imports.NewTable("TSynEditStrings_GetInfoForViewedXY", 0), // procedure GetInfoForViewedXY
			/* 31 */ imports.NewTable("TSynEditStrings_EditInsert", 0), // procedure EditInsert
			/* 32 */ imports.NewTable("TSynEditStrings_EditLineBreak", 0), // procedure EditLineBreak
			/* 33 */ imports.NewTable("TSynEditStrings_EditLineJoin", 0), // procedure EditLineJoin
			/* 34 */ imports.NewTable("TSynEditStrings_EditLinesInsert", 0), // procedure EditLinesInsert
			/* 35 */ imports.NewTable("TSynEditStrings_EditLinesDelete", 0), // procedure EditLinesDelete
			/* 36 */ imports.NewTable("TSynEditStrings_EditUndo", 0), // procedure EditUndo
			/* 37 */ imports.NewTable("TSynEditStrings_EditRedo", 0), // procedure EditRedo
			/* 38 */ imports.NewTable("TSynEditStrings_LogPhysConvertor", 0), // property LogPhysConvertor
			/* 39 */ imports.NewTable("TSynEditStrings_IsInEditAction", 0), // property IsInEditAction
			/* 40 */ imports.NewTable("TSynEditStrings_UndoList", 0), // property UndoList
			/* 41 */ imports.NewTable("TSynEditStrings_RedoList", 0), // property RedoList
			/* 42 */ imports.NewTable("TSynEditStrings_CurUndoList", 0), // property CurUndoList
			/* 43 */ imports.NewTable("TSynEditStrings_IsUndoing", 0), // property IsUndoing
			/* 44 */ imports.NewTable("TSynEditStrings_IsRedoing", 0), // property IsRedoing
			/* 45 */ imports.NewTable("TSynEditStrings_TextChangeStamp", 0), // property TextChangeStamp
			/* 46 */ imports.NewTable("TSynEditStrings_ViewChangeStamp", 0), // property ViewChangeStamp
			/* 47 */ imports.NewTable("TSynEditStrings_LengthOfLongestLine", 0), // property LengthOfLongestLine
			/* 48 */ imports.NewTable("TSynEditStrings_ViewedLines", 0), // property ViewedLines
			/* 49 */ imports.NewTable("TSynEditStrings_ViewedCount", 0), // property ViewedCount
			/* 50 */ imports.NewTable("TSynEditStrings_IsUtf8", 0), // property IsUtf8
			/* 51 */ imports.NewTable("TSynEditStrings_DisplayView", 0), // property DisplayView
		}
	})
	return synEditStringsImport
}
