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

// ICustomSynEdit Parent: ISynEditBase
type ICustomSynEdit interface {
	ISynEditBase
	HasText(flags types.TSynEditHasTextFlags) bool                                                                                                                                                             // function
	GetBookMarkWithIntX3(bookMark int32, X *int32, Y *int32) bool                                                                                                                                              // function
	GetBookMarkWithIntX5(bookMark int32, X *int32, Y *int32, left *int32, top *int32) bool                                                                                                                     // function
	IsBookmark(bookMark int32) bool                                                                                                                                                                            // function
	GetHighlighterAttriAtRowCol(xY types.TPoint, outToken *string, outAttri *ISynHighlighterAttributes) bool                                                                                                   // function
	GetHighlighterAttriAtRowColExWithPointStrIntX2SHAttributesBool(xY types.TPoint, outToken *string, outTokenType *int32, outStart *int32, outAttri *ISynHighlighterAttributes, continueIfPossible bool) bool // function
	GetHighlighterAttriAtRowColExWithPointIntBool(xY types.TPoint, outTokenType *int32, continueIfPossible bool) bool                                                                                          // function
	NextWordPos() types.TPoint                                                                                                                                                                                 // function
	PrevWordPos() types.TPoint                                                                                                                                                                                 // function
	IdentChars() types.TSynIdentChars                                                                                                                                                                          // function
	IsIdentChar(C string) bool                                                                                                                                                                                 // function
	// ScreenColumnToXValue
	//  End "from SynMemo"
	//  Pixel
	ScreenColumnToXValue(col int32) int32                                                                     // function
	ScreenXYToPixels(rowCol types.TRect) types.TPoint                                                         // function
	RowColumnToPixels(rowCol types.TPoint) types.TPoint                                                       // function
	PixelsToRowColumn(pixels types.TPoint, flags types.TSynCoordinateMappingFlags) types.TPoint               // function
	PixelsToLogicalPos(pixels types.TPoint) types.TPoint                                                      // function
	SearchReplace(search string, replace string, options types.TSynSearchOptions) int32                       // function
	SearchReplaceEx(search string, replace string, options types.TSynSearchOptions, start types.TPoint) int32 // function
	PluginCount() int32                                                                                       // function
	MarkupCount() int32                                                                                       // function
	AfterLoadFromFile()                                                                                       // procedure
	EnsureCursorPosVisible()                                                                                  // procedure
	MoveCaretToVisibleArea()                                                                                  // procedure
	MoveCaretIgnoreEOL(newCaret types.TPoint)                                                                 // procedure
	MoveLogicalCaretIgnoreEOL(newLogCaret types.TPoint)                                                       // procedure
	// ClearSelection
	//  Selection
	ClearSelection()                // procedure
	SelectAll()                     // procedure
	SelectToBrace()                 // procedure
	SelectWord()                    // procedure
	SelectLine(withLeadSpaces bool) // procedure
	SelectParagraph()               // procedure
	// Clear
	//  Text Raw (not undo-able)
	Clear()              // procedure
	Append(value string) // procedure
	// ClearAll
	//  Text (unho-able)
	ClearAll()                                                                                                                                                                                                                                                                   // procedure
	InsertTextAtCaret(text string, caretMode types.TSynCaretAdjustMode)                                                                                                                                                                                                          // procedure
	SetTextBetweenPointsWithPointX2StrSETFlagsSCAModeSMAModeSSMode(startPoint types.TPoint, endPoint types.TPoint, value string, flags types.TSynEditTextFlags, caretMode types.TSynCaretAdjustMode, marksMode types.TSynMarksAdjustMode, selectionMode types.TSynSelectionMode) // procedure
	MarkTextAsSaved()                                                                                                                                                                                                                                                            // procedure
	// ClearBookMark
	//  BoorMark
	ClearBookMark(bookMark int32)                                            // procedure
	GotoBookMark(bookMark int32)                                             // procedure
	SetBookMark(bookMark int32, X int32, Y int32, anLeft int32, anTop int32) // procedure
	// CopyToClipboard
	//  Clipboard
	CopyToClipboard()                                                                                                 // procedure
	CutToClipboard()                                                                                                  // procedure
	PasteFromClipboard(forceColumnMode bool)                                                                          // procedure
	DoCopyToClipboard(sText string, foldInfo string)                                                                  // procedure
	CommandProcessor(command types.TSynEditorCommand, char string, data uintptr, skipHooks types.THookedCommandFlags) // procedure
	ExecuteCommand(command types.TSynEditorCommand, char string, data uintptr)                                        // procedure
	CaretAtIdentOrString(xY types.TPoint, outAtIdent *bool, outNearString *bool)                                      // procedure
	Notification(component IComponent, operation types.TOperation)                                                    // procedure
	SetUseIncrementalColorWithBool(value bool)                                                                        // procedure
	SetDefaultKeystrokes()                                                                                            // procedure
	ResetMouseActions()                                                                                               // procedure
	SetOptionFlag(flag types.TSynEditorOption, value bool)                                                            // procedure
	SetHighlightSearch(search string, options types.TSynSearchOptions)                                                // procedure
	WndProc(msg *types.TMessage)                                                                                      // procedure
	AddKey(command types.TSynEditorCommand, key1 uint16, sS1 types.TShiftState, key2 uint16, sS2 types.TShiftState)   // procedure
	ShareTextBufferFrom(shareEditor ICustomSynEdit)                                                                   // procedure
	UnShareTextBuffer()                                                                                               // procedure
	// SetCaretTypeSize
	//  Caret
	SetCaretTypeSize(type_ types.TSynCaretType, width int32, height int32, xOffs int32, yOffs int32)                          // procedure
	LineText() string                                                                                                         // property LineText Getter
	SetLineText(value string)                                                                                                 // property LineText Setter
	Text() string                                                                                                             // property Text Getter
	SetText(value string)                                                                                                     // property Text Setter
	TextBetweenPoints(startPoint types.TPoint, endPoint types.TPoint) string                                                  // property TextBetweenPoints Getter
	SetTextBetweenPointsWithPointX2ToStr(startPoint types.TPoint, endPoint types.TPoint, value string)                        // property TextBetweenPoints Setter
	SetTextBetweenPointsEx(startPoint types.TPoint, endPoint types.TPoint, caretMode types.TSynCaretAdjustMode, value string) // property TextBetweenPointsEx Setter
	MarksToSynEditMarkList() ISynEditMarkList                                                                                 // property Marks Getter
	CanPaste() bool                                                                                                           // property CanPaste Getter
	FoldState() string                                                                                                        // property FoldState Getter
	SetFoldState(value string)                                                                                                // property FoldState Setter
	MaxLeftChar() int32                                                                                                       // property MaxLeftChar Getter
	SetMaxLeftChar(value int32)                                                                                               // property MaxLeftChar Setter
	ScrollOnEditLeftOptions() ISynScrollOnEditOptions                                                                         // property ScrollOnEditLeftOptions Getter
	SetScrollOnEditLeftOptions(value ISynScrollOnEditOptions)                                                                 // property ScrollOnEditLeftOptions Setter
	ScrollOnEditRightOptions() ISynScrollOnEditOptions                                                                        // property ScrollOnEditRightOptions Getter
	SetScrollOnEditRightOptions(value ISynScrollOnEditOptions)                                                                // property ScrollOnEditRightOptions Setter
	SetUseIncrementalColorToBool(value bool)                                                                                  // property UseIncrementalColor Setter
	PaintLock() int32                                                                                                         // property PaintLock Getter
	UseUTF8() bool                                                                                                            // property UseUTF8 Getter
	ChangeStamp() int64                                                                                                       // property ChangeStamp Getter
	Plugin(index int32) ILazSynEditPlugin                                                                                     // property Plugin Getter
	Markup(index int32) ISynEditMarkup                                                                                        // property Markup Getter
	MarkupByClass(index types.TSynEditMarkupClass) ISynEditMarkup                                                             // property MarkupByClass Getter
	TrimSpaceType() types.TSynEditStringTrimmingType                                                                          // property TrimSpaceType Getter
	SetTrimSpaceType(value types.TSynEditStringTrimmingType)                                                                  // property TrimSpaceType Setter
	InsertCaret() types.TSynEditCaretType                                                                                     // property InsertCaret Getter
	SetInsertCaret(value types.TSynEditCaretType)                                                                             // property InsertCaret Setter
	OverwriteCaret() types.TSynEditCaretType                                                                                  // property OverwriteCaret Getter
	SetOverwriteCaret(value types.TSynEditCaretType)                                                                          // property OverwriteCaret Setter
	// DefaultSelectionMode
	//  Selection
	DefaultSelectionMode() types.TSynSelectionMode         // property DefaultSelectionMode Getter
	SetDefaultSelectionMode(value types.TSynSelectionMode) // property DefaultSelectionMode Setter
	SelectionMode() types.TSynSelectionMode                // property SelectionMode Getter
	SetSelectionMode(value types.TSynSelectionMode)        // property SelectionMode Setter
	IsStickySelecting() bool                               // property IsStickySelecting Getter
	// MarkupManager
	//  Colors
	MarkupManager() ISynEditMarkupManager                               // property MarkupManager Getter
	OffTextCursor() types.TCursor                                       // property OffTextCursor Getter
	SetOffTextCursor(value types.TCursor)                               // property OffTextCursor Setter
	IncrementColor() ISynSelectedColor                                  // property IncrementColor Getter
	SetIncrementColor(value ISynSelectedColor)                          // property IncrementColor Setter
	HighlightAllColor() ISynSelectedColor                               // property HighlightAllColor Getter
	SetHighlightAllColor(value ISynSelectedColor)                       // property HighlightAllColor Setter
	BracketMatchColor() ISynSelectedColor                               // property BracketMatchColor Getter
	SetBracketMatchColor(value ISynSelectedColor)                       // property BracketMatchColor Setter
	MouseLinkColor() ISynSelectedColor                                  // property MouseLinkColor Getter
	SetMouseLinkColor(value ISynSelectedColor)                          // property MouseLinkColor Setter
	LineHighlightColor() ISynSelectedColor                              // property LineHighlightColor Getter
	SetLineHighlightColor(value ISynSelectedColor)                      // property LineHighlightColor Setter
	FoldedCodeColor() ISynSelectedColor                                 // property FoldedCodeColor Getter
	SetFoldedCodeColor(value ISynSelectedColor)                         // property FoldedCodeColor Setter
	FoldedCodeLineColor() ISynSelectedColor                             // property FoldedCodeLineColor Getter
	SetFoldedCodeLineColor(value ISynSelectedColor)                     // property FoldedCodeLineColor Setter
	HiddenCodeLineColor() ISynSelectedColor                             // property HiddenCodeLineColor Getter
	SetHiddenCodeLineColor(value ISynSelectedColor)                     // property HiddenCodeLineColor Setter
	Beautifier() ISynCustomBeautifier                                   // property Beautifier Getter
	SetBeautifier(value ISynCustomBeautifier)                           // property Beautifier Setter
	BlockIndent() int32                                                 // property BlockIndent Getter
	SetBlockIndent(value int32)                                         // property BlockIndent Setter
	BlockTabIndent() int32                                              // property BlockTabIndent Getter
	SetBlockTabIndent(value int32)                                      // property BlockTabIndent Setter
	HighlighterToSynCustomHighlighter() ISynCustomHighlighter           // property Highlighter Getter
	SetHighlighter(value ISynCustomHighlighter)                         // property Highlighter Setter
	Gutter() ISynGutter                                                 // property Gutter Getter
	SetGutter(value ISynGutter)                                         // property Gutter Setter
	RightGutter() ISynGutter                                            // property RightGutter Getter
	SetRightGutter(value ISynGutter)                                    // property RightGutter Setter
	InsertMode() bool                                                   // property InsertMode Getter
	SetInsertMode(value bool)                                           // property InsertMode Setter
	Keystrokes() ISynEditKeyStrokes                                     // property Keystrokes Getter
	SetKeystrokes(value ISynEditKeyStrokes)                             // property Keystrokes Setter
	MaxUndo() int32                                                     // property MaxUndo Getter
	SetMaxUndo(value int32)                                             // property MaxUndo Setter
	ShareOptions() types.TSynEditorShareOptions                         // property ShareOptions Getter
	SetShareOptions(value types.TSynEditorShareOptions)                 // property ShareOptions Setter
	VisibleSpecialChars() types.TSynVisibleSpecialChars                 // property VisibleSpecialChars Getter
	SetVisibleSpecialChars(value types.TSynVisibleSpecialChars)         // property VisibleSpecialChars Setter
	RightEdge() int32                                                   // property RightEdge Getter
	SetRightEdge(value int32)                                           // property RightEdge Setter
	RightEdgeColor() types.TColor                                       // property RightEdgeColor Getter
	SetRightEdgeColor(value types.TColor)                               // property RightEdgeColor Setter
	ScrollBars() types.TScrollStyle                                     // property ScrollBars Getter
	SetScrollBars(value types.TScrollStyle)                             // property ScrollBars Setter
	BracketHighlightStyle() types.TSynEditBracketHighlightStyle         // property BracketHighlightStyle Getter
	SetBracketHighlightStyle(value types.TSynEditBracketHighlightStyle) // property BracketHighlightStyle Setter
	TabWidth() int32                                                    // property TabWidth Getter
	SetTabWidth(value int32)                                            // property TabWidth Setter
	WantTabs() bool                                                     // property WantTabs Getter
	SetWantTabs(value bool)                                             // property WantTabs Setter
	SetTabViewClass(value types.TSynEditStringTabExpanderClass)         // property TabViewClass Setter
	SetOnChange(fn TNotifyEvent)                                        // property event
	SetOnChangeUpdating(fn TChangeUpdatingEvent)                        // property event
	SetOnCutCopy(fn TSynCopyPasteEvent)                                 // property event
	SetOnPaste(fn TSynCopyPasteEvent)                                   // property event
	SetOnDropFiles(fn TSynDropFilesEvent)                               // property event
	SetOnGutterClick(fn TGutterClickEvent)                              // property event
	SetOnPaintToPaintEvent(fn TPaintEvent)                              // property event
	SetOnPlaceBookmark(fn TPlaceMarkEvent)                              // property event
	SetOnClearBookmark(fn TPlaceMarkEvent)                              // property event
	SetOnProcessCommand(fn TProcessCommandEvent)                        // property event
	SetOnProcessUserCommand(fn TProcessCommandEvent)                    // property event
	SetOnCommandProcessed(fn TProcessCommandEvent)                      // property event
	SetOnReplaceText(fn TReplaceTextEvent)                              // property event
	SetOnSpecialLineMarkup(fn TSpecialLineMarkupEvent)                  // property event
}

type TCustomSynEdit struct {
	TSynEditBase
}

func (m *TCustomSynEdit) HasText(flags types.TSynEditHasTextFlags) bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(1, m.Instance(), uintptr(flags))
	return api.GoBool(r)
}

func (m *TCustomSynEdit) GetBookMarkWithIntX3(bookMark int32, X *int32, Y *int32) bool {
	if !m.IsValid() {
		return false
	}
	XPtr := uintptr(*X)
	YPtr := uintptr(*Y)
	r := customSynEditAPI().SysCallN(2, m.Instance(), uintptr(bookMark), uintptr(base.UnsafePointer(&XPtr)), uintptr(base.UnsafePointer(&YPtr)))
	*X = int32(XPtr)
	*Y = int32(YPtr)
	return api.GoBool(r)
}

func (m *TCustomSynEdit) GetBookMarkWithIntX5(bookMark int32, X *int32, Y *int32, left *int32, top *int32) bool {
	if !m.IsValid() {
		return false
	}
	XPtr := uintptr(*X)
	YPtr := uintptr(*Y)
	leftPtr := uintptr(*left)
	topPtr := uintptr(*top)
	r := customSynEditAPI().SysCallN(3, m.Instance(), uintptr(bookMark), uintptr(base.UnsafePointer(&XPtr)), uintptr(base.UnsafePointer(&YPtr)), uintptr(base.UnsafePointer(&leftPtr)), uintptr(base.UnsafePointer(&topPtr)))
	*X = int32(XPtr)
	*Y = int32(YPtr)
	*left = int32(leftPtr)
	*top = int32(topPtr)
	return api.GoBool(r)
}

func (m *TCustomSynEdit) IsBookmark(bookMark int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(4, m.Instance(), uintptr(bookMark))
	return api.GoBool(r)
}

func (m *TCustomSynEdit) GetHighlighterAttriAtRowCol(xY types.TPoint, outToken *string, outAttri *ISynHighlighterAttributes) bool {
	if !m.IsValid() {
		return false
	}
	var tokenPtr uintptr
	var attriPtr uintptr
	r := customSynEditAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&xY)), uintptr(base.UnsafePointer(&tokenPtr)), uintptr(base.UnsafePointer(&attriPtr)))
	*outToken = api.GoStr(tokenPtr)
	*outAttri = AsSynHighlighterAttributes(attriPtr)
	return api.GoBool(r)
}

func (m *TCustomSynEdit) GetHighlighterAttriAtRowColExWithPointStrIntX2SHAttributesBool(xY types.TPoint, outToken *string, outTokenType *int32, outStart *int32, outAttri *ISynHighlighterAttributes, continueIfPossible bool) bool {
	if !m.IsValid() {
		return false
	}
	var tokenPtr uintptr
	var tokenTypePtr uintptr
	var startPtr uintptr
	var attriPtr uintptr
	r := customSynEditAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&xY)), uintptr(base.UnsafePointer(&tokenPtr)), uintptr(base.UnsafePointer(&tokenTypePtr)), uintptr(base.UnsafePointer(&startPtr)), uintptr(base.UnsafePointer(&attriPtr)), api.PasBool(continueIfPossible))
	*outToken = api.GoStr(tokenPtr)
	*outTokenType = int32(tokenTypePtr)
	*outStart = int32(startPtr)
	*outAttri = AsSynHighlighterAttributes(attriPtr)
	return api.GoBool(r)
}

func (m *TCustomSynEdit) GetHighlighterAttriAtRowColExWithPointIntBool(xY types.TPoint, outTokenType *int32, continueIfPossible bool) bool {
	if !m.IsValid() {
		return false
	}
	var tokenTypePtr uintptr
	r := customSynEditAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&xY)), uintptr(base.UnsafePointer(&tokenTypePtr)), api.PasBool(continueIfPossible))
	*outTokenType = int32(tokenTypePtr)
	return api.GoBool(r)
}

func (m *TCustomSynEdit) NextWordPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomSynEdit) PrevWordPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomSynEdit) IdentChars() types.TSynIdentChars {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(10, m.Instance())
	return types.TSynIdentChars(r)
}

func (m *TCustomSynEdit) IsIdentChar(C string) bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(11, m.Instance(), api.PasStr(C))
	return api.GoBool(r)
}

func (m *TCustomSynEdit) ScreenColumnToXValue(col int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(12, m.Instance(), uintptr(col))
	return int32(r)
}

func (m *TCustomSynEdit) ScreenXYToPixels(rowCol types.TRect) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&rowCol)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomSynEdit) RowColumnToPixels(rowCol types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&rowCol)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomSynEdit) PixelsToRowColumn(pixels types.TPoint, flags types.TSynCoordinateMappingFlags) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&pixels)), uintptr(flags), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomSynEdit) PixelsToLogicalPos(pixels types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&pixels)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomSynEdit) SearchReplace(search string, replace string, options types.TSynSearchOptions) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(17, m.Instance(), api.PasStr(search), api.PasStr(replace), uintptr(options))
	return int32(r)
}

func (m *TCustomSynEdit) SearchReplaceEx(search string, replace string, options types.TSynSearchOptions, start types.TPoint) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(18, m.Instance(), api.PasStr(search), api.PasStr(replace), uintptr(options), uintptr(base.UnsafePointer(&start)))
	return int32(r)
}

func (m *TCustomSynEdit) PluginCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(19, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) MarkupCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(20, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) AfterLoadFromFile() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(21, m.Instance())
}

func (m *TCustomSynEdit) EnsureCursorPosVisible() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(22, m.Instance())
}

func (m *TCustomSynEdit) MoveCaretToVisibleArea() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(23, m.Instance())
}

func (m *TCustomSynEdit) MoveCaretIgnoreEOL(newCaret types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&newCaret)))
}

func (m *TCustomSynEdit) MoveLogicalCaretIgnoreEOL(newLogCaret types.TPoint) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&newLogCaret)))
}

func (m *TCustomSynEdit) ClearSelection() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(26, m.Instance())
}

func (m *TCustomSynEdit) SelectAll() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(27, m.Instance())
}

func (m *TCustomSynEdit) SelectToBrace() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(28, m.Instance())
}

func (m *TCustomSynEdit) SelectWord() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(29, m.Instance())
}

func (m *TCustomSynEdit) SelectLine(withLeadSpaces bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(30, m.Instance(), api.PasBool(withLeadSpaces))
}

func (m *TCustomSynEdit) SelectParagraph() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(31, m.Instance())
}

func (m *TCustomSynEdit) Clear() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(32, m.Instance())
}

func (m *TCustomSynEdit) Append(value string) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(33, m.Instance(), api.PasStr(value))
}

func (m *TCustomSynEdit) ClearAll() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(34, m.Instance())
}

func (m *TCustomSynEdit) InsertTextAtCaret(text string, caretMode types.TSynCaretAdjustMode) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(35, m.Instance(), api.PasStr(text), uintptr(caretMode))
}

func (m *TCustomSynEdit) SetTextBetweenPointsWithPointX2StrSETFlagsSCAModeSMAModeSSMode(startPoint types.TPoint, endPoint types.TPoint, value string, flags types.TSynEditTextFlags, caretMode types.TSynCaretAdjustMode, marksMode types.TSynMarksAdjustMode, selectionMode types.TSynSelectionMode) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(36, m.Instance(), uintptr(base.UnsafePointer(&startPoint)), uintptr(base.UnsafePointer(&endPoint)), api.PasStr(value), uintptr(flags), uintptr(caretMode), uintptr(marksMode), uintptr(selectionMode))
}

func (m *TCustomSynEdit) MarkTextAsSaved() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(37, m.Instance())
}

func (m *TCustomSynEdit) ClearBookMark(bookMark int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(38, m.Instance(), uintptr(bookMark))
}

func (m *TCustomSynEdit) GotoBookMark(bookMark int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(39, m.Instance(), uintptr(bookMark))
}

func (m *TCustomSynEdit) SetBookMark(bookMark int32, X int32, Y int32, anLeft int32, anTop int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(40, m.Instance(), uintptr(bookMark), uintptr(X), uintptr(Y), uintptr(anLeft), uintptr(anTop))
}

func (m *TCustomSynEdit) CopyToClipboard() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(41, m.Instance())
}

func (m *TCustomSynEdit) CutToClipboard() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(42, m.Instance())
}

func (m *TCustomSynEdit) PasteFromClipboard(forceColumnMode bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(43, m.Instance(), api.PasBool(forceColumnMode))
}

func (m *TCustomSynEdit) DoCopyToClipboard(sText string, foldInfo string) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(44, m.Instance(), api.PasStr(sText), api.PasStr(foldInfo))
}

func (m *TCustomSynEdit) CommandProcessor(command types.TSynEditorCommand, char string, data uintptr, skipHooks types.THookedCommandFlags) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(45, m.Instance(), uintptr(command), api.PasStr(char), uintptr(data), uintptr(skipHooks))
}

func (m *TCustomSynEdit) ExecuteCommand(command types.TSynEditorCommand, char string, data uintptr) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(46, m.Instance(), uintptr(command), api.PasStr(char), uintptr(data))
}

func (m *TCustomSynEdit) CaretAtIdentOrString(xY types.TPoint, outAtIdent *bool, outNearString *bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(47, m.Instance(), uintptr(base.UnsafePointer(&xY)), uintptr(base.UnsafePointer(outAtIdent)), uintptr(base.UnsafePointer(outNearString)))
}

func (m *TCustomSynEdit) Notification(component IComponent, operation types.TOperation) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(48, m.Instance(), base.GetObjectUintptr(component), uintptr(operation))
}

func (m *TCustomSynEdit) SetUseIncrementalColorWithBool(value bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(49, m.Instance(), api.PasBool(value))
}

func (m *TCustomSynEdit) SetDefaultKeystrokes() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(50, m.Instance())
}

func (m *TCustomSynEdit) ResetMouseActions() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(51, m.Instance())
}

func (m *TCustomSynEdit) SetOptionFlag(flag types.TSynEditorOption, value bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(52, m.Instance(), uintptr(flag), api.PasBool(value))
}

func (m *TCustomSynEdit) SetHighlightSearch(search string, options types.TSynSearchOptions) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(53, m.Instance(), api.PasStr(search), uintptr(options))
}

func (m *TCustomSynEdit) WndProc(msg *types.TMessage) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(54, m.Instance(), uintptr(base.UnsafePointer(msg)))
}

func (m *TCustomSynEdit) AddKey(command types.TSynEditorCommand, key1 uint16, sS1 types.TShiftState, key2 uint16, sS2 types.TShiftState) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(55, m.Instance(), uintptr(command), uintptr(key1), uintptr(sS1), uintptr(key2), uintptr(sS2))
}

func (m *TCustomSynEdit) ShareTextBufferFrom(shareEditor ICustomSynEdit) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(56, m.Instance(), base.GetObjectUintptr(shareEditor))
}

func (m *TCustomSynEdit) UnShareTextBuffer() {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(57, m.Instance())
}

func (m *TCustomSynEdit) SetCaretTypeSize(type_ types.TSynCaretType, width int32, height int32, xOffs int32, yOffs int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(58, m.Instance(), uintptr(type_), uintptr(width), uintptr(height), uintptr(xOffs), uintptr(yOffs))
}

func (m *TCustomSynEdit) LineText() string {
	if !m.IsValid() {
		return ""
	}
	r := customSynEditAPI().SysCallN(59, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomSynEdit) SetLineText(value string) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(59, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomSynEdit) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := customSynEditAPI().SysCallN(60, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomSynEdit) SetText(value string) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(60, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomSynEdit) TextBetweenPoints(startPoint types.TPoint, endPoint types.TPoint) string {
	if !m.IsValid() {
		return ""
	}
	r := customSynEditAPI().SysCallN(61, 0, m.Instance(), uintptr(base.UnsafePointer(&startPoint)), uintptr(base.UnsafePointer(&endPoint)))
	return api.GoStr(r)
}

func (m *TCustomSynEdit) SetTextBetweenPointsWithPointX2ToStr(startPoint types.TPoint, endPoint types.TPoint, value string) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(61, 1, m.Instance(), uintptr(base.UnsafePointer(&startPoint)), uintptr(base.UnsafePointer(&endPoint)), api.PasStr(value))
}

func (m *TCustomSynEdit) SetTextBetweenPointsEx(startPoint types.TPoint, endPoint types.TPoint, caretMode types.TSynCaretAdjustMode, value string) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(62, m.Instance(), uintptr(base.UnsafePointer(&startPoint)), uintptr(base.UnsafePointer(&endPoint)), uintptr(caretMode), api.PasStr(value))
}

func (m *TCustomSynEdit) MarksToSynEditMarkList() ISynEditMarkList {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(63, m.Instance())
	return AsSynEditMarkList(r)
}

func (m *TCustomSynEdit) CanPaste() bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(64, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSynEdit) FoldState() string {
	if !m.IsValid() {
		return ""
	}
	r := customSynEditAPI().SysCallN(65, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomSynEdit) SetFoldState(value string) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(65, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomSynEdit) MaxLeftChar() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(66, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) SetMaxLeftChar(value int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(66, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) ScrollOnEditLeftOptions() ISynScrollOnEditOptions {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(67, 0, m.Instance())
	return AsSynScrollOnEditOptions(r)
}

func (m *TCustomSynEdit) SetScrollOnEditLeftOptions(value ISynScrollOnEditOptions) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(67, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) ScrollOnEditRightOptions() ISynScrollOnEditOptions {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(68, 0, m.Instance())
	return AsSynScrollOnEditOptions(r)
}

func (m *TCustomSynEdit) SetScrollOnEditRightOptions(value ISynScrollOnEditOptions) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(68, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) SetUseIncrementalColorToBool(value bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(69, m.Instance(), api.PasBool(value))
}

func (m *TCustomSynEdit) PaintLock() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(70, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) UseUTF8() bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(71, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSynEdit) ChangeStamp() (result int64) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(72, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomSynEdit) Plugin(index int32) ILazSynEditPlugin {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(73, m.Instance(), uintptr(index))
	return AsLazSynEditPlugin(r)
}

func (m *TCustomSynEdit) Markup(index int32) ISynEditMarkup {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(74, m.Instance(), uintptr(index))
	return AsSynEditMarkup(r)
}

func (m *TCustomSynEdit) MarkupByClass(index types.TSynEditMarkupClass) ISynEditMarkup {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(75, m.Instance(), uintptr(index))
	return AsSynEditMarkup(r)
}

func (m *TCustomSynEdit) TrimSpaceType() types.TSynEditStringTrimmingType {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(76, 0, m.Instance())
	return types.TSynEditStringTrimmingType(r)
}

func (m *TCustomSynEdit) SetTrimSpaceType(value types.TSynEditStringTrimmingType) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(76, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) InsertCaret() types.TSynEditCaretType {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(77, 0, m.Instance())
	return types.TSynEditCaretType(r)
}

func (m *TCustomSynEdit) SetInsertCaret(value types.TSynEditCaretType) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(77, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) OverwriteCaret() types.TSynEditCaretType {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(78, 0, m.Instance())
	return types.TSynEditCaretType(r)
}

func (m *TCustomSynEdit) SetOverwriteCaret(value types.TSynEditCaretType) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(78, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) DefaultSelectionMode() types.TSynSelectionMode {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(79, 0, m.Instance())
	return types.TSynSelectionMode(r)
}

func (m *TCustomSynEdit) SetDefaultSelectionMode(value types.TSynSelectionMode) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(79, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) SelectionMode() types.TSynSelectionMode {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(80, 0, m.Instance())
	return types.TSynSelectionMode(r)
}

func (m *TCustomSynEdit) SetSelectionMode(value types.TSynSelectionMode) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(80, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) IsStickySelecting() bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(81, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSynEdit) MarkupManager() ISynEditMarkupManager {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(82, m.Instance())
	return AsSynEditMarkupManager(r)
}

func (m *TCustomSynEdit) OffTextCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(83, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCustomSynEdit) SetOffTextCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(83, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) IncrementColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(84, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetIncrementColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(84, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) HighlightAllColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(85, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetHighlightAllColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(85, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) BracketMatchColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(86, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetBracketMatchColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(86, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) MouseLinkColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(87, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetMouseLinkColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(87, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) LineHighlightColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(88, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetLineHighlightColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(88, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) FoldedCodeColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(89, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetFoldedCodeColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(89, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) FoldedCodeLineColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(90, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetFoldedCodeLineColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(90, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) HiddenCodeLineColor() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(91, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TCustomSynEdit) SetHiddenCodeLineColor(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(91, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) Beautifier() ISynCustomBeautifier {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(92, 0, m.Instance())
	return AsSynCustomBeautifier(r)
}

func (m *TCustomSynEdit) SetBeautifier(value ISynCustomBeautifier) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(92, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) BlockIndent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(93, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) SetBlockIndent(value int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(93, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) BlockTabIndent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(94, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) SetBlockTabIndent(value int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(94, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) HighlighterToSynCustomHighlighter() ISynCustomHighlighter {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(95, 0, m.Instance())
	return AsSynCustomHighlighter(r)
}

func (m *TCustomSynEdit) SetHighlighter(value ISynCustomHighlighter) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(95, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) Gutter() ISynGutter {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(96, 0, m.Instance())
	return AsSynGutter(r)
}

func (m *TCustomSynEdit) SetGutter(value ISynGutter) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(96, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) RightGutter() ISynGutter {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(97, 0, m.Instance())
	return AsSynGutter(r)
}

func (m *TCustomSynEdit) SetRightGutter(value ISynGutter) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(97, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) InsertMode() bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(98, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSynEdit) SetInsertMode(value bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(98, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSynEdit) Keystrokes() ISynEditKeyStrokes {
	if !m.IsValid() {
		return nil
	}
	r := customSynEditAPI().SysCallN(99, 0, m.Instance())
	return AsSynEditKeyStrokes(r)
}

func (m *TCustomSynEdit) SetKeystrokes(value ISynEditKeyStrokes) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(99, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSynEdit) MaxUndo() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(100, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) SetMaxUndo(value int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(100, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) ShareOptions() types.TSynEditorShareOptions {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(101, 0, m.Instance())
	return types.TSynEditorShareOptions(r)
}

func (m *TCustomSynEdit) SetShareOptions(value types.TSynEditorShareOptions) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(101, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) VisibleSpecialChars() types.TSynVisibleSpecialChars {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(102, 0, m.Instance())
	return types.TSynVisibleSpecialChars(r)
}

func (m *TCustomSynEdit) SetVisibleSpecialChars(value types.TSynVisibleSpecialChars) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(102, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) RightEdge() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(103, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) SetRightEdge(value int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(103, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) RightEdgeColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(104, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomSynEdit) SetRightEdgeColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(104, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) ScrollBars() types.TScrollStyle {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(105, 0, m.Instance())
	return types.TScrollStyle(r)
}

func (m *TCustomSynEdit) SetScrollBars(value types.TScrollStyle) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(105, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) BracketHighlightStyle() types.TSynEditBracketHighlightStyle {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(106, 0, m.Instance())
	return types.TSynEditBracketHighlightStyle(r)
}

func (m *TCustomSynEdit) SetBracketHighlightStyle(value types.TSynEditBracketHighlightStyle) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(106, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) TabWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSynEditAPI().SysCallN(107, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSynEdit) SetTabWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(107, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) WantTabs() bool {
	if !m.IsValid() {
		return false
	}
	r := customSynEditAPI().SysCallN(108, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSynEdit) SetWantTabs(value bool) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(108, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSynEdit) SetTabViewClass(value types.TSynEditStringTabExpanderClass) {
	if !m.IsValid() {
		return
	}
	customSynEditAPI().SysCallN(109, m.Instance(), uintptr(value))
}

func (m *TCustomSynEdit) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 110, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnChangeUpdating(fn TChangeUpdatingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTChangeUpdatingEvent(fn)
	base.SetEvent(m, 111, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnCutCopy(fn TSynCopyPasteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynCopyPasteEvent(fn)
	base.SetEvent(m, 112, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnPaste(fn TSynCopyPasteEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynCopyPasteEvent(fn)
	base.SetEvent(m, 113, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnDropFiles(fn TSynDropFilesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynDropFilesEvent(fn)
	base.SetEvent(m, 114, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnGutterClick(fn TGutterClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGutterClickEvent(fn)
	base.SetEvent(m, 115, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnPaintToPaintEvent(fn TPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTPaintEvent(fn)
	base.SetEvent(m, 116, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnPlaceBookmark(fn TPlaceMarkEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTPlaceMarkEvent(fn)
	base.SetEvent(m, 117, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnClearBookmark(fn TPlaceMarkEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTPlaceMarkEvent(fn)
	base.SetEvent(m, 118, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnProcessCommand(fn TProcessCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTProcessCommandEvent(fn)
	base.SetEvent(m, 119, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnProcessUserCommand(fn TProcessCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTProcessCommandEvent(fn)
	base.SetEvent(m, 120, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnCommandProcessed(fn TProcessCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTProcessCommandEvent(fn)
	base.SetEvent(m, 121, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnReplaceText(fn TReplaceTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTReplaceTextEvent(fn)
	base.SetEvent(m, 122, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynEdit) SetOnSpecialLineMarkup(fn TSpecialLineMarkupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSpecialLineMarkupEvent(fn)
	base.SetEvent(m, 123, customSynEditAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomSynEdit class constructor
func NewCustomSynEdit(owner IComponent) ICustomSynEdit {
	r := customSynEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomSynEdit(r)
}

func TCustomSynEditClass() types.TClass {
	r := customSynEditAPI().SysCallN(124)
	return types.TClass(r)
}

var (
	customSynEditOnce   base.Once
	customSynEditImport *imports.Imports = nil
)

func customSynEditAPI() *imports.Imports {
	customSynEditOnce.Do(func() {
		customSynEditImport = api.NewDefaultImports()
		customSynEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomSynEdit_Create", 0), // constructor NewCustomSynEdit
			/* 1 */ imports.NewTable("TCustomSynEdit_HasText", 0), // function HasText
			/* 2 */ imports.NewTable("TCustomSynEdit_GetBookMarkWithIntX3", 0), // function GetBookMarkWithIntX3
			/* 3 */ imports.NewTable("TCustomSynEdit_GetBookMarkWithIntX5", 0), // function GetBookMarkWithIntX5
			/* 4 */ imports.NewTable("TCustomSynEdit_IsBookmark", 0), // function IsBookmark
			/* 5 */ imports.NewTable("TCustomSynEdit_GetHighlighterAttriAtRowCol", 0), // function GetHighlighterAttriAtRowCol
			/* 6 */ imports.NewTable("TCustomSynEdit_GetHighlighterAttriAtRowColExWithPointStrIntX2SHAttributesBool", 0), // function GetHighlighterAttriAtRowColExWithPointStrIntX2SHAttributesBool
			/* 7 */ imports.NewTable("TCustomSynEdit_GetHighlighterAttriAtRowColExWithPointIntBool", 0), // function GetHighlighterAttriAtRowColExWithPointIntBool
			/* 8 */ imports.NewTable("TCustomSynEdit_NextWordPos", 0), // function NextWordPos
			/* 9 */ imports.NewTable("TCustomSynEdit_PrevWordPos", 0), // function PrevWordPos
			/* 10 */ imports.NewTable("TCustomSynEdit_IdentChars", 0), // function IdentChars
			/* 11 */ imports.NewTable("TCustomSynEdit_IsIdentChar", 0), // function IsIdentChar
			/* 12 */ imports.NewTable("TCustomSynEdit_ScreenColumnToXValue", 0), // function ScreenColumnToXValue
			/* 13 */ imports.NewTable("TCustomSynEdit_ScreenXYToPixels", 0), // function ScreenXYToPixels
			/* 14 */ imports.NewTable("TCustomSynEdit_RowColumnToPixels", 0), // function RowColumnToPixels
			/* 15 */ imports.NewTable("TCustomSynEdit_PixelsToRowColumn", 0), // function PixelsToRowColumn
			/* 16 */ imports.NewTable("TCustomSynEdit_PixelsToLogicalPos", 0), // function PixelsToLogicalPos
			/* 17 */ imports.NewTable("TCustomSynEdit_SearchReplace", 0), // function SearchReplace
			/* 18 */ imports.NewTable("TCustomSynEdit_SearchReplaceEx", 0), // function SearchReplaceEx
			/* 19 */ imports.NewTable("TCustomSynEdit_PluginCount", 0), // function PluginCount
			/* 20 */ imports.NewTable("TCustomSynEdit_MarkupCount", 0), // function MarkupCount
			/* 21 */ imports.NewTable("TCustomSynEdit_AfterLoadFromFile", 0), // procedure AfterLoadFromFile
			/* 22 */ imports.NewTable("TCustomSynEdit_EnsureCursorPosVisible", 0), // procedure EnsureCursorPosVisible
			/* 23 */ imports.NewTable("TCustomSynEdit_MoveCaretToVisibleArea", 0), // procedure MoveCaretToVisibleArea
			/* 24 */ imports.NewTable("TCustomSynEdit_MoveCaretIgnoreEOL", 0), // procedure MoveCaretIgnoreEOL
			/* 25 */ imports.NewTable("TCustomSynEdit_MoveLogicalCaretIgnoreEOL", 0), // procedure MoveLogicalCaretIgnoreEOL
			/* 26 */ imports.NewTable("TCustomSynEdit_ClearSelection", 0), // procedure ClearSelection
			/* 27 */ imports.NewTable("TCustomSynEdit_SelectAll", 0), // procedure SelectAll
			/* 28 */ imports.NewTable("TCustomSynEdit_SelectToBrace", 0), // procedure SelectToBrace
			/* 29 */ imports.NewTable("TCustomSynEdit_SelectWord", 0), // procedure SelectWord
			/* 30 */ imports.NewTable("TCustomSynEdit_SelectLine", 0), // procedure SelectLine
			/* 31 */ imports.NewTable("TCustomSynEdit_SelectParagraph", 0), // procedure SelectParagraph
			/* 32 */ imports.NewTable("TCustomSynEdit_Clear", 0), // procedure Clear
			/* 33 */ imports.NewTable("TCustomSynEdit_Append", 0), // procedure Append
			/* 34 */ imports.NewTable("TCustomSynEdit_ClearAll", 0), // procedure ClearAll
			/* 35 */ imports.NewTable("TCustomSynEdit_InsertTextAtCaret", 0), // procedure InsertTextAtCaret
			/* 36 */ imports.NewTable("TCustomSynEdit_SetTextBetweenPointsWithPointX2StrSETFlagsSCAModeSMAModeSSMode", 0), // procedure SetTextBetweenPointsWithPointX2StrSETFlagsSCAModeSMAModeSSMode
			/* 37 */ imports.NewTable("TCustomSynEdit_MarkTextAsSaved", 0), // procedure MarkTextAsSaved
			/* 38 */ imports.NewTable("TCustomSynEdit_ClearBookMark", 0), // procedure ClearBookMark
			/* 39 */ imports.NewTable("TCustomSynEdit_GotoBookMark", 0), // procedure GotoBookMark
			/* 40 */ imports.NewTable("TCustomSynEdit_SetBookMark", 0), // procedure SetBookMark
			/* 41 */ imports.NewTable("TCustomSynEdit_CopyToClipboard", 0), // procedure CopyToClipboard
			/* 42 */ imports.NewTable("TCustomSynEdit_CutToClipboard", 0), // procedure CutToClipboard
			/* 43 */ imports.NewTable("TCustomSynEdit_PasteFromClipboard", 0), // procedure PasteFromClipboard
			/* 44 */ imports.NewTable("TCustomSynEdit_DoCopyToClipboard", 0), // procedure DoCopyToClipboard
			/* 45 */ imports.NewTable("TCustomSynEdit_CommandProcessor", 0), // procedure CommandProcessor
			/* 46 */ imports.NewTable("TCustomSynEdit_ExecuteCommand", 0), // procedure ExecuteCommand
			/* 47 */ imports.NewTable("TCustomSynEdit_CaretAtIdentOrString", 0), // procedure CaretAtIdentOrString
			/* 48 */ imports.NewTable("TCustomSynEdit_Notification", 0), // procedure Notification
			/* 49 */ imports.NewTable("TCustomSynEdit_SetUseIncrementalColorWithBool", 0), // procedure SetUseIncrementalColorWithBool
			/* 50 */ imports.NewTable("TCustomSynEdit_SetDefaultKeystrokes", 0), // procedure SetDefaultKeystrokes
			/* 51 */ imports.NewTable("TCustomSynEdit_ResetMouseActions", 0), // procedure ResetMouseActions
			/* 52 */ imports.NewTable("TCustomSynEdit_SetOptionFlag", 0), // procedure SetOptionFlag
			/* 53 */ imports.NewTable("TCustomSynEdit_SetHighlightSearch", 0), // procedure SetHighlightSearch
			/* 54 */ imports.NewTable("TCustomSynEdit_WndProc", 0), // procedure WndProc
			/* 55 */ imports.NewTable("TCustomSynEdit_AddKey", 0), // procedure AddKey
			/* 56 */ imports.NewTable("TCustomSynEdit_ShareTextBufferFrom", 0), // procedure ShareTextBufferFrom
			/* 57 */ imports.NewTable("TCustomSynEdit_UnShareTextBuffer", 0), // procedure UnShareTextBuffer
			/* 58 */ imports.NewTable("TCustomSynEdit_SetCaretTypeSize", 0), // procedure SetCaretTypeSize
			/* 59 */ imports.NewTable("TCustomSynEdit_LineText", 0), // property LineText
			/* 60 */ imports.NewTable("TCustomSynEdit_Text", 0), // property Text
			/* 61 */ imports.NewTable("TCustomSynEdit_TextBetweenPoints", 0), // property TextBetweenPoints
			/* 62 */ imports.NewTable("TCustomSynEdit_TextBetweenPointsEx", 0), // property TextBetweenPointsEx
			/* 63 */ imports.NewTable("TCustomSynEdit_MarksToSynEditMarkList", 0), // property MarksToSynEditMarkList
			/* 64 */ imports.NewTable("TCustomSynEdit_CanPaste", 0), // property CanPaste
			/* 65 */ imports.NewTable("TCustomSynEdit_FoldState", 0), // property FoldState
			/* 66 */ imports.NewTable("TCustomSynEdit_MaxLeftChar", 0), // property MaxLeftChar
			/* 67 */ imports.NewTable("TCustomSynEdit_ScrollOnEditLeftOptions", 0), // property ScrollOnEditLeftOptions
			/* 68 */ imports.NewTable("TCustomSynEdit_ScrollOnEditRightOptions", 0), // property ScrollOnEditRightOptions
			/* 69 */ imports.NewTable("TCustomSynEdit_UseIncrementalColor", 0), // property UseIncrementalColor
			/* 70 */ imports.NewTable("TCustomSynEdit_PaintLock", 0), // property PaintLock
			/* 71 */ imports.NewTable("TCustomSynEdit_UseUTF8", 0), // property UseUTF8
			/* 72 */ imports.NewTable("TCustomSynEdit_ChangeStamp", 0), // property ChangeStamp
			/* 73 */ imports.NewTable("TCustomSynEdit_Plugin", 0), // property Plugin
			/* 74 */ imports.NewTable("TCustomSynEdit_Markup", 0), // property Markup
			/* 75 */ imports.NewTable("TCustomSynEdit_MarkupByClass", 0), // property MarkupByClass
			/* 76 */ imports.NewTable("TCustomSynEdit_TrimSpaceType", 0), // property TrimSpaceType
			/* 77 */ imports.NewTable("TCustomSynEdit_InsertCaret", 0), // property InsertCaret
			/* 78 */ imports.NewTable("TCustomSynEdit_OverwriteCaret", 0), // property OverwriteCaret
			/* 79 */ imports.NewTable("TCustomSynEdit_DefaultSelectionMode", 0), // property DefaultSelectionMode
			/* 80 */ imports.NewTable("TCustomSynEdit_SelectionMode", 0), // property SelectionMode
			/* 81 */ imports.NewTable("TCustomSynEdit_IsStickySelecting", 0), // property IsStickySelecting
			/* 82 */ imports.NewTable("TCustomSynEdit_MarkupManager", 0), // property MarkupManager
			/* 83 */ imports.NewTable("TCustomSynEdit_OffTextCursor", 0), // property OffTextCursor
			/* 84 */ imports.NewTable("TCustomSynEdit_IncrementColor", 0), // property IncrementColor
			/* 85 */ imports.NewTable("TCustomSynEdit_HighlightAllColor", 0), // property HighlightAllColor
			/* 86 */ imports.NewTable("TCustomSynEdit_BracketMatchColor", 0), // property BracketMatchColor
			/* 87 */ imports.NewTable("TCustomSynEdit_MouseLinkColor", 0), // property MouseLinkColor
			/* 88 */ imports.NewTable("TCustomSynEdit_LineHighlightColor", 0), // property LineHighlightColor
			/* 89 */ imports.NewTable("TCustomSynEdit_FoldedCodeColor", 0), // property FoldedCodeColor
			/* 90 */ imports.NewTable("TCustomSynEdit_FoldedCodeLineColor", 0), // property FoldedCodeLineColor
			/* 91 */ imports.NewTable("TCustomSynEdit_HiddenCodeLineColor", 0), // property HiddenCodeLineColor
			/* 92 */ imports.NewTable("TCustomSynEdit_Beautifier", 0), // property Beautifier
			/* 93 */ imports.NewTable("TCustomSynEdit_BlockIndent", 0), // property BlockIndent
			/* 94 */ imports.NewTable("TCustomSynEdit_BlockTabIndent", 0), // property BlockTabIndent
			/* 95 */ imports.NewTable("TCustomSynEdit_HighlighterToSynCustomHighlighter", 0), // property HighlighterToSynCustomHighlighter
			/* 96 */ imports.NewTable("TCustomSynEdit_Gutter", 0), // property Gutter
			/* 97 */ imports.NewTable("TCustomSynEdit_RightGutter", 0), // property RightGutter
			/* 98 */ imports.NewTable("TCustomSynEdit_InsertMode", 0), // property InsertMode
			/* 99 */ imports.NewTable("TCustomSynEdit_Keystrokes", 0), // property Keystrokes
			/* 100 */ imports.NewTable("TCustomSynEdit_MaxUndo", 0), // property MaxUndo
			/* 101 */ imports.NewTable("TCustomSynEdit_ShareOptions", 0), // property ShareOptions
			/* 102 */ imports.NewTable("TCustomSynEdit_VisibleSpecialChars", 0), // property VisibleSpecialChars
			/* 103 */ imports.NewTable("TCustomSynEdit_RightEdge", 0), // property RightEdge
			/* 104 */ imports.NewTable("TCustomSynEdit_RightEdgeColor", 0), // property RightEdgeColor
			/* 105 */ imports.NewTable("TCustomSynEdit_ScrollBars", 0), // property ScrollBars
			/* 106 */ imports.NewTable("TCustomSynEdit_BracketHighlightStyle", 0), // property BracketHighlightStyle
			/* 107 */ imports.NewTable("TCustomSynEdit_TabWidth", 0), // property TabWidth
			/* 108 */ imports.NewTable("TCustomSynEdit_WantTabs", 0), // property WantTabs
			/* 109 */ imports.NewTable("TCustomSynEdit_TabViewClass", 0), // property TabViewClass
			/* 110 */ imports.NewTable("TCustomSynEdit_OnChange", 0), // event OnChange
			/* 111 */ imports.NewTable("TCustomSynEdit_OnChangeUpdating", 0), // event OnChangeUpdating
			/* 112 */ imports.NewTable("TCustomSynEdit_OnCutCopy", 0), // event OnCutCopy
			/* 113 */ imports.NewTable("TCustomSynEdit_OnPaste", 0), // event OnPaste
			/* 114 */ imports.NewTable("TCustomSynEdit_OnDropFiles", 0), // event OnDropFiles
			/* 115 */ imports.NewTable("TCustomSynEdit_OnGutterClick", 0), // event OnGutterClick
			/* 116 */ imports.NewTable("TCustomSynEdit_OnPaintToPaintEvent", 0), // event OnPaint
			/* 117 */ imports.NewTable("TCustomSynEdit_OnPlaceBookmark", 0), // event OnPlaceBookmark
			/* 118 */ imports.NewTable("TCustomSynEdit_OnClearBookmark", 0), // event OnClearBookmark
			/* 119 */ imports.NewTable("TCustomSynEdit_OnProcessCommand", 0), // event OnProcessCommand
			/* 120 */ imports.NewTable("TCustomSynEdit_OnProcessUserCommand", 0), // event OnProcessUserCommand
			/* 121 */ imports.NewTable("TCustomSynEdit_OnCommandProcessed", 0), // event OnCommandProcessed
			/* 122 */ imports.NewTable("TCustomSynEdit_OnReplaceText", 0), // event OnReplaceText
			/* 123 */ imports.NewTable("TCustomSynEdit_OnSpecialLineMarkup", 0), // event OnSpecialLineMarkup
			/* 124 */ imports.NewTable("TCustomSynEdit_TClass", 0), // function TCustomSynEditClass
		}
	})
	return customSynEditImport
}
