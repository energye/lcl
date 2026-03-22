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

// ISynEditBase Parent: ICustomControl
type ISynEditBase interface {
	ICustomControl
	FindGutterFromGutterPartList(partList IObject) IObject // function
	// CaretXPix
	//  Caret
	CaretXPix() int32 // function
	CaretYPix() int32 // function
	// ScreenXYToTextXY
	//  * ScreenXY:
	//  First visible (scrolled in) screen line is 1
	//  First column is 1 => column does not take scrolling into account
	ScreenXYToTextXY(screenXY types.TRect, limitToLines bool) types.TRect                                                                                                 // function
	TextXYToScreenXY(physTextXY types.TRect) types.TRect                                                                                                                  // function
	GetWordAtRowCol(xY types.TPoint) string                                                                                                                               // function
	FindMatchingBracketWithPointBoolX4(physStartBracket types.TPoint, startIncludeNeighborChars bool, moveCaret bool, selectBrackets bool, onlyVisible bool) types.TPoint // function
	FindMatchingBracketLogical(logicalStartBracket types.TPoint, startIncludeNeighborChars bool, moveCaret bool, selectBrackets bool, onlyVisible bool) types.TPoint      // function
	IsLinkable(Y int32, x1 int32, x2 int32) bool                                                                                                                          // function
	// GetLineState
	//  text / lines
	GetLineState(line int32) types.TSynLineState // function
	// LogicalToPhysicalPos
	//  Byte to Char
	LogicalToPhysicalPos(P types.TPoint) types.TPoint                      // function
	LogicalToPhysicalCol(line string, index int32, logicalPos int32) int32 // function
	// PhysicalToLogicalPos
	//  Char to Byte
	PhysicalToLogicalPos(P types.TPoint) types.TPoint                        // function
	PhysicalToLogicalCol(line string, index int32, physicalPos int32) int32  // function
	PhysicalLineLength(line string, index int32) int32                       // function
	GetWordBoundsAtRowCol(xY types.TPoint, outStartX *int32, outEndX *int32) // procedure
	// UpdateCursorOverride
	//  Cursor
	UpdateCursorOverride() // procedure
	// BeginUndoBlock
	//  Undo Redo
	BeginUndoBlock()                // procedure
	BeginUpdate(withUndoBlock bool) // procedure
	EndUndoBlock()                  // procedure
	EndUpdate()                     // procedure
	ClearUndo()                     // procedure
	Redo()                          // procedure
	Undo()                          // procedure
	// FindMatchingBracket
	//  matching brackets
	FindMatchingBracket() // procedure
	// InvalidateGutter
	//  invalidate lines
	InvalidateGutter()                                     // procedure
	InvalidateLine(line int32)                             // procedure
	InvalidateGutterLines(firstLine int32, lastLine int32) // procedure
	InvalidateLines(firstLine int32, lastLine int32)       // procedure
	CanRedo() bool                                         // property CanRedo Getter
	CanUndo() bool                                         // property CanUndo Getter
	BookMarkOptions() ISynBookMarkOpt                      // property BookMarkOptions Getter
	SetBookMarkOptions(value ISynBookMarkOpt)              // property BookMarkOptions Setter
	ExtraCharSpacing() int32                               // property ExtraCharSpacing Getter
	SetExtraCharSpacing(value int32)                       // property ExtraCharSpacing Setter
	ExtraLineSpacing() int32                               // property ExtraLineSpacing Getter
	SetExtraLineSpacing(value int32)                       // property ExtraLineSpacing Setter
	Lines() IStrings                                       // property Lines Getter
	SetLines(value IStrings)                               // property Lines Setter
	// Options
	//  See SYNEDIT_UNIMPLEMENTED_OPTIONS for deprecated Values
	Options() types.TSynEditorOptions           // property Options Getter
	SetOptions(value types.TSynEditorOptions)   // property Options Setter
	Options2() types.TSynEditorOptions2         // property Options2 Getter
	SetOptions2(value types.TSynEditorOptions2) // property Options2 Setter
	ReadOnly() bool                             // property ReadOnly Getter
	SetReadOnly(value bool)                     // property ReadOnly Setter
	Modified() bool                             // property Modified Getter
	SetModified(value bool)                     // property Modified Setter
	CaretX() int32                              // property CaretX Getter
	SetCaretX(value int32)                      // property CaretX Setter
	CaretY() int32                              // property CaretY Getter
	SetCaretY(value int32)                      // property CaretY Setter
	CaretXY() types.TPoint                      // property CaretXY Getter
	SetCaretXY(value types.TPoint)              // property CaretXY Setter
	LogicalCaretXY() types.TPoint               // property LogicalCaretXY Getter
	SetLogicalCaretXY(value types.TPoint)       // property LogicalCaretXY Setter
	CharsInWindow() int32                       // property CharsInWindow Getter
	CharWidth() int32                           // property CharWidth Getter
	LeftChar() int32                            // property LeftChar Getter
	SetLeftChar(value int32)                    // property LeftChar Setter
	LineHeight() int32                          // property LineHeight Getter
	LinesInWindow() int32                       // property LinesInWindow Getter
	TopLine() int32                             // property TopLine Getter
	SetTopLine(value int32)                     // property TopLine Setter
	BlockBegin() types.TPoint                   // property BlockBegin Getter
	SetBlockBegin(value types.TPoint)           // property BlockBegin Setter
	BlockEnd() types.TPoint                     // property BlockEnd Getter
	SetBlockEnd(value types.TPoint)             // property BlockEnd Setter
	SelStart() int32                            // property SelStart Getter
	SetSelStart(value int32)                    // property SelStart Setter
	SelEnd() int32                              // property SelEnd Getter
	SetSelEnd(value int32)                      // property SelEnd Setter
	IsBackwardSel() bool                        // property IsBackwardSel Getter
	SelText() string                            // property SelText Getter
	SetSelText(value string)                    // property SelText Setter
	MouseActions() ISynEditMouseActions         // property MouseActions Getter
	SetMouseActions(value ISynEditMouseActions) // property MouseActions Setter
	// MouseSelActions
	//  Mouseactions, if mouse is over selection => fallback to normal
	MouseSelActions() ISynEditMouseActions              // property MouseSelActions Getter
	SetMouseSelActions(value ISynEditMouseActions)      // property MouseSelActions Setter
	MouseTextActions() ISynEditMouseActions             // property MouseTextActions Getter
	SetMouseTextActions(value ISynEditMouseActions)     // property MouseTextActions Setter
	MouseOptions() types.TSynEditorMouseOptions         // property MouseOptions Getter
	SetMouseOptions(value types.TSynEditorMouseOptions) // property MouseOptions Setter
	TextViewsManager() ISynTextViewsManager             // property TextViewsManager Getter
	SelectedColor() ISynSelectedColor                   // property SelectedColor Getter
	SetSelectedColor(value ISynSelectedColor)           // property SelectedColor Setter
	SelAvail() bool                                     // property SelAvail Getter
	HideSelection() bool                                // property HideSelection Getter
	SetHideSelection(value bool)                        // property HideSelection Setter
	HighlighterToObject() IObject                       // property Highlighter Getter
	MarksToObject() IObject                             // property Marks Getter
}

type TSynEditBase struct {
	TCustomControl
}

func (m *TSynEditBase) FindGutterFromGutterPartList(partList IObject) IObject {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(0, m.Instance(), base.GetObjectUintptr(partList))
	return AsObject(r)
}

func (m *TSynEditBase) CaretXPix() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) CaretYPix() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) ScreenXYToTextXY(screenXY types.TRect, limitToLines bool) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&screenXY)), api.PasBool(limitToLines), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) TextXYToScreenXY(physTextXY types.TRect) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&physTextXY)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) GetWordAtRowCol(xY types.TPoint) string {
	if !m.IsValid() {
		return ""
	}
	r := synEditBaseAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&xY)))
	return api.GoStr(r)
}

func (m *TSynEditBase) FindMatchingBracketWithPointBoolX4(physStartBracket types.TPoint, startIncludeNeighborChars bool, moveCaret bool, selectBrackets bool, onlyVisible bool) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&physStartBracket)), api.PasBool(startIncludeNeighborChars), api.PasBool(moveCaret), api.PasBool(selectBrackets), api.PasBool(onlyVisible), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) FindMatchingBracketLogical(logicalStartBracket types.TPoint, startIncludeNeighborChars bool, moveCaret bool, selectBrackets bool, onlyVisible bool) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&logicalStartBracket)), api.PasBool(startIncludeNeighborChars), api.PasBool(moveCaret), api.PasBool(selectBrackets), api.PasBool(onlyVisible), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) IsLinkable(Y int32, x1 int32, x2 int32) bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(8, m.Instance(), uintptr(Y), uintptr(x1), uintptr(x2))
	return api.GoBool(r)
}

func (m *TSynEditBase) GetLineState(line int32) types.TSynLineState {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(9, m.Instance(), uintptr(line))
	return types.TSynLineState(r)
}

func (m *TSynEditBase) LogicalToPhysicalPos(P types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&P)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) LogicalToPhysicalCol(line string, index int32, logicalPos int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(11, m.Instance(), api.PasStr(line), uintptr(index), uintptr(logicalPos))
	return int32(r)
}

func (m *TSynEditBase) PhysicalToLogicalPos(P types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&P)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) PhysicalToLogicalCol(line string, index int32, physicalPos int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(13, m.Instance(), api.PasStr(line), uintptr(index), uintptr(physicalPos))
	return int32(r)
}

func (m *TSynEditBase) PhysicalLineLength(line string, index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(14, m.Instance(), api.PasStr(line), uintptr(index))
	return int32(r)
}

func (m *TSynEditBase) GetWordBoundsAtRowCol(xY types.TPoint, outStartX *int32, outEndX *int32) {
	if !m.IsValid() {
		return
	}
	var startXPtr uintptr
	var endXPtr uintptr
	synEditBaseAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&xY)), uintptr(base.UnsafePointer(&startXPtr)), uintptr(base.UnsafePointer(&endXPtr)))
	*outStartX = int32(startXPtr)
	*outEndX = int32(endXPtr)
}

func (m *TSynEditBase) UpdateCursorOverride() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(16, m.Instance())
}

func (m *TSynEditBase) BeginUndoBlock() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(17, m.Instance())
}

func (m *TSynEditBase) BeginUpdate(withUndoBlock bool) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(18, m.Instance(), api.PasBool(withUndoBlock))
}

func (m *TSynEditBase) EndUndoBlock() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(19, m.Instance())
}

func (m *TSynEditBase) EndUpdate() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(20, m.Instance())
}

func (m *TSynEditBase) ClearUndo() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(21, m.Instance())
}

func (m *TSynEditBase) Redo() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(22, m.Instance())
}

func (m *TSynEditBase) Undo() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(23, m.Instance())
}

func (m *TSynEditBase) FindMatchingBracket() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(24, m.Instance())
}

func (m *TSynEditBase) InvalidateGutter() {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(25, m.Instance())
}

func (m *TSynEditBase) InvalidateLine(line int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(26, m.Instance(), uintptr(line))
}

func (m *TSynEditBase) InvalidateGutterLines(firstLine int32, lastLine int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(27, m.Instance(), uintptr(firstLine), uintptr(lastLine))
}

func (m *TSynEditBase) InvalidateLines(firstLine int32, lastLine int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(28, m.Instance(), uintptr(firstLine), uintptr(lastLine))
}

func (m *TSynEditBase) CanRedo() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditBase) CanUndo() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(30, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditBase) BookMarkOptions() ISynBookMarkOpt {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(31, 0, m.Instance())
	return AsSynBookMarkOpt(r)
}

func (m *TSynEditBase) SetBookMarkOptions(value ISynBookMarkOpt) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(31, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditBase) ExtraCharSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(32, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetExtraCharSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(32, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) ExtraLineSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(33, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetExtraLineSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(33, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) Lines() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(34, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynEditBase) SetLines(value IStrings) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(34, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditBase) Options() types.TSynEditorOptions {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(35, 0, m.Instance())
	return types.TSynEditorOptions(r)
}

func (m *TSynEditBase) SetOptions(value types.TSynEditorOptions) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) Options2() types.TSynEditorOptions2 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(36, 0, m.Instance())
	return types.TSynEditorOptions2(r)
}

func (m *TSynEditBase) SetOptions2(value types.TSynEditorOptions2) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(37, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditBase) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(37, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditBase) Modified() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(38, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditBase) SetModified(value bool) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(38, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditBase) CaretX() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(39, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetCaretX(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(39, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) CaretY() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(40, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetCaretY(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) CaretXY() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(41, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) SetCaretXY(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(41, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBase) LogicalCaretXY() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(42, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) SetLogicalCaretXY(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(42, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBase) CharsInWindow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(43, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) CharWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(44, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) LeftChar() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(45, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetLeftChar(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(45, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) LineHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(46, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) LinesInWindow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(47, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) TopLine() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(48, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetTopLine(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(48, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) BlockBegin() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(49, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) SetBlockBegin(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(49, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBase) BlockEnd() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(50, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynEditBase) SetBlockEnd(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(50, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynEditBase) SelStart() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(51, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetSelStart(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(51, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) SelEnd() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(52, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditBase) SetSelEnd(value int32) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(52, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) IsBackwardSel() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(53, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditBase) SelText() string {
	if !m.IsValid() {
		return ""
	}
	r := synEditBaseAPI().SysCallN(54, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TSynEditBase) SetSelText(value string) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(54, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynEditBase) MouseActions() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(55, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynEditBase) SetMouseActions(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(55, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditBase) MouseSelActions() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(56, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynEditBase) SetMouseSelActions(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(56, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditBase) MouseTextActions() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(57, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynEditBase) SetMouseTextActions(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(57, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditBase) MouseOptions() types.TSynEditorMouseOptions {
	if !m.IsValid() {
		return 0
	}
	r := synEditBaseAPI().SysCallN(58, 0, m.Instance())
	return types.TSynEditorMouseOptions(r)
}

func (m *TSynEditBase) SetMouseOptions(value types.TSynEditorMouseOptions) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(58, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditBase) TextViewsManager() ISynTextViewsManager {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(59, m.Instance())
	return AsSynTextViewsManager(r)
}

func (m *TSynEditBase) SelectedColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(60, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynEditBase) SetSelectedColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(60, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditBase) SelAvail() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(61, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditBase) HideSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditBaseAPI().SysCallN(62, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditBase) SetHideSelection(value bool) {
	if !m.IsValid() {
		return
	}
	synEditBaseAPI().SysCallN(62, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditBase) HighlighterToObject() IObject {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(63, m.Instance())
	return AsObject(r)
}

func (m *TSynEditBase) MarksToObject() IObject {
	if !m.IsValid() {
		return nil
	}
	r := synEditBaseAPI().SysCallN(64, m.Instance())
	return AsObject(r)
}

var (
	synEditBaseOnce   base.Once
	synEditBaseImport *imports.Imports = nil
)

func synEditBaseAPI() *imports.Imports {
	synEditBaseOnce.Do(func() {
		synEditBaseImport = api.NewDefaultImports()
		synEditBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditBase_FindGutterFromGutterPartList", 0), // function FindGutterFromGutterPartList
			/* 1 */ imports.NewTable("TSynEditBase_CaretXPix", 0), // function CaretXPix
			/* 2 */ imports.NewTable("TSynEditBase_CaretYPix", 0), // function CaretYPix
			/* 3 */ imports.NewTable("TSynEditBase_ScreenXYToTextXY", 0), // function ScreenXYToTextXY
			/* 4 */ imports.NewTable("TSynEditBase_TextXYToScreenXY", 0), // function TextXYToScreenXY
			/* 5 */ imports.NewTable("TSynEditBase_GetWordAtRowCol", 0), // function GetWordAtRowCol
			/* 6 */ imports.NewTable("TSynEditBase_FindMatchingBracketWithPointBoolX4", 0), // function FindMatchingBracketWithPointBoolX4
			/* 7 */ imports.NewTable("TSynEditBase_FindMatchingBracketLogical", 0), // function FindMatchingBracketLogical
			/* 8 */ imports.NewTable("TSynEditBase_IsLinkable", 0), // function IsLinkable
			/* 9 */ imports.NewTable("TSynEditBase_GetLineState", 0), // function GetLineState
			/* 10 */ imports.NewTable("TSynEditBase_LogicalToPhysicalPos", 0), // function LogicalToPhysicalPos
			/* 11 */ imports.NewTable("TSynEditBase_LogicalToPhysicalCol", 0), // function LogicalToPhysicalCol
			/* 12 */ imports.NewTable("TSynEditBase_PhysicalToLogicalPos", 0), // function PhysicalToLogicalPos
			/* 13 */ imports.NewTable("TSynEditBase_PhysicalToLogicalCol", 0), // function PhysicalToLogicalCol
			/* 14 */ imports.NewTable("TSynEditBase_PhysicalLineLength", 0), // function PhysicalLineLength
			/* 15 */ imports.NewTable("TSynEditBase_GetWordBoundsAtRowCol", 0), // procedure GetWordBoundsAtRowCol
			/* 16 */ imports.NewTable("TSynEditBase_UpdateCursorOverride", 0), // procedure UpdateCursorOverride
			/* 17 */ imports.NewTable("TSynEditBase_BeginUndoBlock", 0), // procedure BeginUndoBlock
			/* 18 */ imports.NewTable("TSynEditBase_BeginUpdate", 0), // procedure BeginUpdate
			/* 19 */ imports.NewTable("TSynEditBase_EndUndoBlock", 0), // procedure EndUndoBlock
			/* 20 */ imports.NewTable("TSynEditBase_EndUpdate", 0), // procedure EndUpdate
			/* 21 */ imports.NewTable("TSynEditBase_ClearUndo", 0), // procedure ClearUndo
			/* 22 */ imports.NewTable("TSynEditBase_Redo", 0), // procedure Redo
			/* 23 */ imports.NewTable("TSynEditBase_Undo", 0), // procedure Undo
			/* 24 */ imports.NewTable("TSynEditBase_FindMatchingBracket", 0), // procedure FindMatchingBracket
			/* 25 */ imports.NewTable("TSynEditBase_InvalidateGutter", 0), // procedure InvalidateGutter
			/* 26 */ imports.NewTable("TSynEditBase_InvalidateLine", 0), // procedure InvalidateLine
			/* 27 */ imports.NewTable("TSynEditBase_InvalidateGutterLines", 0), // procedure InvalidateGutterLines
			/* 28 */ imports.NewTable("TSynEditBase_InvalidateLines", 0), // procedure InvalidateLines
			/* 29 */ imports.NewTable("TSynEditBase_CanRedo", 0), // property CanRedo
			/* 30 */ imports.NewTable("TSynEditBase_CanUndo", 0), // property CanUndo
			/* 31 */ imports.NewTable("TSynEditBase_BookMarkOptions", 0), // property BookMarkOptions
			/* 32 */ imports.NewTable("TSynEditBase_ExtraCharSpacing", 0), // property ExtraCharSpacing
			/* 33 */ imports.NewTable("TSynEditBase_ExtraLineSpacing", 0), // property ExtraLineSpacing
			/* 34 */ imports.NewTable("TSynEditBase_Lines", 0), // property Lines
			/* 35 */ imports.NewTable("TSynEditBase_Options", 0), // property Options
			/* 36 */ imports.NewTable("TSynEditBase_Options2", 0), // property Options2
			/* 37 */ imports.NewTable("TSynEditBase_ReadOnly", 0), // property ReadOnly
			/* 38 */ imports.NewTable("TSynEditBase_Modified", 0), // property Modified
			/* 39 */ imports.NewTable("TSynEditBase_CaretX", 0), // property CaretX
			/* 40 */ imports.NewTable("TSynEditBase_CaretY", 0), // property CaretY
			/* 41 */ imports.NewTable("TSynEditBase_CaretXY", 0), // property CaretXY
			/* 42 */ imports.NewTable("TSynEditBase_LogicalCaretXY", 0), // property LogicalCaretXY
			/* 43 */ imports.NewTable("TSynEditBase_CharsInWindow", 0), // property CharsInWindow
			/* 44 */ imports.NewTable("TSynEditBase_CharWidth", 0), // property CharWidth
			/* 45 */ imports.NewTable("TSynEditBase_LeftChar", 0), // property LeftChar
			/* 46 */ imports.NewTable("TSynEditBase_LineHeight", 0), // property LineHeight
			/* 47 */ imports.NewTable("TSynEditBase_LinesInWindow", 0), // property LinesInWindow
			/* 48 */ imports.NewTable("TSynEditBase_TopLine", 0), // property TopLine
			/* 49 */ imports.NewTable("TSynEditBase_BlockBegin", 0), // property BlockBegin
			/* 50 */ imports.NewTable("TSynEditBase_BlockEnd", 0), // property BlockEnd
			/* 51 */ imports.NewTable("TSynEditBase_SelStart", 0), // property SelStart
			/* 52 */ imports.NewTable("TSynEditBase_SelEnd", 0), // property SelEnd
			/* 53 */ imports.NewTable("TSynEditBase_IsBackwardSel", 0), // property IsBackwardSel
			/* 54 */ imports.NewTable("TSynEditBase_SelText", 0), // property SelText
			/* 55 */ imports.NewTable("TSynEditBase_MouseActions", 0), // property MouseActions
			/* 56 */ imports.NewTable("TSynEditBase_MouseSelActions", 0), // property MouseSelActions
			/* 57 */ imports.NewTable("TSynEditBase_MouseTextActions", 0), // property MouseTextActions
			/* 58 */ imports.NewTable("TSynEditBase_MouseOptions", 0), // property MouseOptions
			/* 59 */ imports.NewTable("TSynEditBase_TextViewsManager", 0), // property TextViewsManager
			/* 60 */ imports.NewTable("TSynEditBase_SelectedColor", 0), // property SelectedColor
			/* 61 */ imports.NewTable("TSynEditBase_SelAvail", 0), // property SelAvail
			/* 62 */ imports.NewTable("TSynEditBase_HideSelection", 0), // property HideSelection
			/* 63 */ imports.NewTable("TSynEditBase_HighlighterToObject", 0), // property HighlighterToObject
			/* 64 */ imports.NewTable("TSynEditBase_MarksToObject", 0), // property MarksToObject
		}
	})
	return synEditBaseImport
}
