//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

// LPosFlag ENUM
type LPosFlag = int32

const (
	LpAllowPastEol LPosFlag = iota
	LpAdjustToNext
	LpStopAtCodePoint
)

// LPosFlags SET: LPosFlag
type LPosFlags = TSet

// TAnchorKind ENUM
type TAnchorKind = int32

const (
	AkTop TAnchorKind = iota
	AkLeft
	AkRight
	AkBottom
)

// TAnchors SET: TAnchorKind
type TAnchors = TSet

// TApplicationFlag ENUM
type TApplicationFlag = int32

const (
	AppWaiting TApplicationFlag = iota
	AppIdleEndSent
	AppNoExceptionMessages
	AppActive
	AppDestroying
	AppDoNotCallAsyncQueue
	AppInitialized
)

// TApplicationFlags SET: TApplicationFlag
type TApplicationFlags = TSet

// TApplicationNavigationOption ENUM
type TApplicationNavigationOption = int32

const (
	AnoTabToSelectNext TApplicationNavigationOption = iota
	AnoReturnForDefaultControl
	AnoEscapeForCancelControl
	AnoF1ForHelp
	AnoArrowToSelectNextInParent
)

// TApplicationNavigationOptions SET: TApplicationNavigationOption
type TApplicationNavigationOptions = TSet

// TAutoCompleteOption ENUM
type TAutoCompleteOption = int32

const (
	AcoAutoSuggest TAutoCompleteOption = iota
	AcoAutoAppend
	AcoSearch
	AcoFilterPrefixes
	AcoUseTab
	AcoUpDownKeyDropsList
	AcoRtlReading
)

// TAutoCompleteOptions SET: TAutoCompleteOption
type TAutoCompleteOptions = TSet

// TBandPaintOption ENUM
type TBandPaintOption = int32

const (
	BpoGrabber TBandPaintOption = iota
	BpoFrame
	BpoGradient
	BpoRoundRect
)

// TBandPaintOptions SET: TBandPaintOption
type TBandPaintOptions = TSet

// TBorderIcon ENUM
//
//	TCustomForm
type TBorderIcon = int32

const (
	BiSystemMenu TBorderIcon = iota
	BiMinimize
	BiMaximize
	BiHelp
)

// TBorderIcons SET: TBorderIcon
type TBorderIcons = TSet

// TCanvasStates ENUM
type TCanvasStates = int32

const (
	CsHandleValid TCanvasStates = iota
	CsFontValid
	CsPenvalid
	CsBrushValid
	CsRegionValid
)

// TCanvasState SET: TCanvasStates
type TCanvasState = TSet

// TColorBoxStyles ENUM
type TColorBoxStyles = int32

const (
	CbStandardColors TColorBoxStyles = iota
	CbExtendedColors
	CbSystemColors
	CbIncludeNone
	CbIncludeDefault
	CbCustomColor
	CbPrettyNames
	CbCustomColors
)

// TColorBoxStyle SET: TColorBoxStyles
type TColorBoxStyle = TSet

// TColorDialogOption ENUM
type TColorDialogOption = int32

const (
	CdFullOpen TColorDialogOption = iota
	CdPreventFullOpen
	CdShowHelp
	CdSolidColor
	CdAnyColor
)

// TColorDialogOptions SET: TColorDialogOption
type TColorDialogOptions = TSet

// TComboBoxAutoCompleteTextOption ENUM
//
//	TCustomComboBox
type TComboBoxAutoCompleteTextOption = int32

const (
	CbactEnabled TComboBoxAutoCompleteTextOption = iota
	CbactEndOfLineComplete
	CbactRetainPrefixCase
	CbactSearchCaseSensitive
	CbactSearchAscending
)

// TComboBoxAutoCompleteText SET: TComboBoxAutoCompleteTextOption
type TComboBoxAutoCompleteText = TSet

// TComboBoxExStyleEx ENUM
type TComboBoxExStyleEx = int32

const (
	CsExCaseSensitive TComboBoxExStyleEx = iota
	CsExNoEditImage
	CsExNoEditImageIndent
	CsExNoSizeLimit
	CsExPathWordBreak
)

// TComboBoxExStyles SET: TComboBoxExStyleEx
type TComboBoxExStyles = TSet

// TCommentStyle ENUM
type TCommentStyle = int32

const (
	CsAnsiStyle TCommentStyle = iota
	CsPasStyle
	CsCStyle
	CsAsmStyle
	CsBasStyle
	CsVBStyle
)

// CommentStyles SET: TCommentStyle
type CommentStyles = TSet

// TComponentStatee ENUM
type TComponentStatee = int32

const (
	CsLoading TComponentStatee = iota
	CsReading
	CsWriting
	CsDestroying
	CsDesigning
	CsAncestor
	CsUpdating
	CsFixups
	CsFreeNotification
	CsInline
	CsDesignInstance
)

// TComponentState SET: TComponentStatee
type TComponentState = TSet

// TComponentStylee ENUM
type TComponentStylee = int32

const (
	CsInheritable TComponentStylee = iota
	CsCheckPropAvail
	CsSubComponent
	CsTransient
)

// TComponentStyle SET: TComponentStylee
type TComponentStyle = TSet

// TControlAtPosFlag ENUM
type TControlAtPosFlag = int32

const (
	CapfAllowDisabled TControlAtPosFlag = iota
	CapfAllowWinControls
	CapfOnlyClientAreas
	CapfRecursive
	CapfHasScrollOffset
	CapfOnlyWinControls
)

// TControlAtPosFlags SET: TControlAtPosFlag
type TControlAtPosFlags = TSet

// TControlAutoSizePhase ENUM
type TControlAutoSizePhase = int32

const (
	CaspNone TControlAutoSizePhase = iota
	CaspChangingProperties
	CaspCreatingHandles
	CaspComputingBounds
	CaspRealizingBounds
	CaspShowing
)

// TControlAutoSizePhases SET: TControlAutoSizePhase
type TControlAutoSizePhases = TSet

// TControlRoleForForm ENUM
type TControlRoleForForm = int32

const (
	CrffDefault TControlRoleForForm = iota
	CrffCancel
)

// TControlRolesForForm SET: TControlRoleForForm
type TControlRolesForForm = TSet

// TControlStateType ENUM
type TControlStateType = int32

const (
	CsLButtonDown TControlStateType = iota
	CsClicked
	CsPalette
	CsReadingState
	CsFocusing
	CsCreating
	CsPaintCopy
	CsCustomPaint
	CsDestroyingHandle
	CsDocking
	CsVisibleSetInLoading
)

// TControlState SET: TControlStateType
type TControlState = TSet

// TControlStyleType ENUM
type TControlStyleType = int32

const (
	CsAcceptsControls TControlStyleType = iota
	CsCaptureMouse
	CsDesignInteractive
	CsClickEvents
	CsFramed
	CsSetCaption
	CsOpaque
	CsDoubleClicks
	CsTripleClicks
	CsQuadClicks
	CsFixedWidth
	CsFixedHeight
	CsNoDesignVisible
	CsReplicatable
	CsNoStdEvents
	CsDisplayDragImage
	CsReflector
	CsActionClient
	CsMenuEvents
	CsNoFocus
	CsNeedsBorderPaint
	CsParentBackground
	CsDesignNoSmoothResize
	CsDesignFixedBounds
	CsHasDefaultAction
	CsHasCancelAction
	CsNoDesignSelectable
	CsOwnedChildrenNotSelectable
	CsAutoSize0x0
	CsAutoSizeKeepChildLeft
	CsAutoSizeKeepChildTop
	CsRequiresKeyboardInput
)

// TControlStyle SET: TControlStyleType
type TControlStyle = TSet

// TCTabControlCapability ENUM
type TCTabControlCapability = int32

const (
	NbcShowCloseButtons TCTabControlCapability = iota
	NbcMultiLine
	NbcPageListPopup
	NbcShowAddTabButton
	NbcTabsSizeable
)

// TCTabControlCapabilities SET: TCTabControlCapability
type TCTabControlCapabilities = TSet

// TCTabControlOption ENUM
//
//	These are LCL additions
type TCTabControlOption = int32

const (
	NboShowCloseButtons TCTabControlOption = iota
	NboMultiLine
	NboHidePageListPopup
	NboKeyboardTabSwitch
	NboShowAddTabButton
	NboDoChangeOnSetIndex
)

// TCTabControlOptions SET: TCTabControlOption
type TCTabControlOptions = TSet

// TCustomDrawStateFlag ENUM
type TCustomDrawStateFlag = int32

const (
	CdsSelected TCustomDrawStateFlag = iota
	CdsGrayed
	CdsDisabled
	CdsChecked
	CdsFocused
	CdsDefault
	CdsHot
	CdsMarked
	CdsIndeterminate
)

// TCustomDrawState SET: TCustomDrawStateFlag
type TCustomDrawState = TSet

// TDateTimePickerOption ENUM
type TDateTimePickerOption = int32

const (
	DtpoDoChangeOnSetDateTime TDateTimePickerOption = iota
	DtpoEnabledIfUnchecked
	DtpoAutoCheck
	DtpoFlatButton
	DtpoResetSelection
)

// TDateTimePickerOptions SET: TDateTimePickerOption
type TDateTimePickerOptions = TSet

// TDisplayOption ENUM
type TDisplayOption = int32

const (
	DoColumnTitles TDisplayOption = iota
	DoAutoColResize
	DoKeyColFixed
)

// TDisplayOptions SET: TDisplayOption
type TDisplayOptions = TSet

// TDisplaySetting ENUM
type TDisplaySetting = int32

const (
	DsShowHeadings TDisplaySetting = iota
	DsShowDayNames
	DsNoMonthChange
	DsShowWeekNumbers
)

// TDisplaySettings SET: TDisplaySetting
type TDisplaySettings = TSet

// TDragOperation ENUM
//
//	operations basically allowed during drag'n drop
type TDragOperation = int32

const (
	DoCopy TDragOperation = iota
	DoMove
	DoLink
)

// TDragOperations SET: TDragOperation
type TDragOperations = TSet

// TEdgeBorder ENUM
//
//	TToolWindow
type TEdgeBorder = int32

const (
	EbLeft TEdgeBorder = iota
	EbTop
	EbRight
	EbBottom
)

// TEdgeBorders SET: TEdgeBorder
type TEdgeBorders = TSet

// TEventType ENUM
type TEventType = int32

const (
	EtNotify TEventType = iota
	EtKey
	EtKeyPress
	EtMouseWheel
	EtMouseUpDown
)

// TEventLogTypes SET: TEventType
type TEventLogTypes = TSet

// TFilterStringOption ENUM
type TFilterStringOption = int32

const (
	FsoCaseSensitive TFilterStringOption = iota
	FsoMatchOnlyAtStart
)

// TFilterStringOptions SET: TFilterStringOption
type TFilterStringOptions = TSet

// TFindOption ENUM
//
//	TFindDialog
type TFindOption = int32

const (
	FrDown TFindOption = iota
	FrFindNext
	FrHideMatchCase
	FrHideWholeWord
	FrHideUpDown
	FrMatchCase
	FrDisableMatchCase
	FrDisableUpDown
	FrDisableWholeWord
	FrReplace
	FrReplaceAll
	FrWholeWord
	FrShowHelp
	FrEntireScope
	FrHideEntireScope
	FrPromptOnReplace
	FrHidePromptOnReplace
	FrButtonsAtBottom
)

// TFindOptions SET: TFindOption
type TFindOptions = TSet

// TFontDialogOption ENUM
//
//	TFontDialog
type TFontDialogOption = int32

const (
	FdAnsiOnly TFontDialogOption = iota
	FdTrueTypeOnly
	FdEffects
	FdFixedPitchOnly
	FdForceFontExist
	FdNoFaceSel
	FdNoOEMFonts
	FdNoSimulations
	FdNoSizeSel
	FdNoStyleSel
	FdNoVectorFonts
	FdShowHelp
	FdWysiwyg
	FdLimitSize
	FdScalableOnly
	FdApplyButton
)

// TFontDialogOptions SET: TFontDialogOption
type TFontDialogOptions = TSet

// TFontStyle ENUM
type TFontStyle = int32

const (
	FsBold TFontStyle = iota
	FsItalic
	FsUnderline
	FsStrikeOut
)

// TFontStyles SET: TFontStyle
type TFontStyles = TSet

// TFontStylesBase SET: TFontStyle
type TFontStylesBase = TSet

// TFormStateType ENUM
type TFormStateType = int32

const (
	FsCreating TFormStateType = iota
	FsVisible
	FsShowing
	FsModal
	FsCreatedMDIChild
	FsBorderStyleChanged
	FsFormStyleChanged
	FsFirstShow
	FsDisableAutoSize
)

// TFormState SET: TFormStateType
type TFormState = TSet

// TGridDrawStatee ENUM
type TGridDrawStatee = int32

const (
	GdSelected TGridDrawStatee = iota
	GdFocused
	GdFixed
	GdHot
	GdPushed
	GdRowHighlight
)

// TGridDrawState SET: TGridDrawStatee
type TGridDrawState = TSet

// TGridOption ENUM
type TGridOption = int32

const (
	GoFixedVertLine TGridOption = iota
	GoFixedHorzLine
	GoVertLine
	GoHorzLine
	GoRangeSelect
	GoDrawFocusSelected
	GoRowSizing
	GoColSizing
	GoRowMoving
	GoColMoving
	GoEditing
	GoAutoAddRows
	GoTabs
	GoRowSelect
	GoAlwaysShowEditor
	GoThumbTracking
	GoColSpanning
	GoRelaxedRowSelect
	GoDblClickAutoSize
	GoSmoothScroll
	GoFixedRowNumbering
	GoScrollKeepVisible
	GoHeaderHotTracking
	GoHeaderPushedLook
	GoSelectionActive
	GoFixedColSizing
	GoDontScrollPartCell
	GoCellHints
	GoTruncCellHints
	GoCellEllipsis
	GoAutoAddRowsSkipContentCheck
	GoRowHighlight
)

// TGridOptions SET: TGridOption
type TGridOptions = TSet

// TGridOption2 ENUM
type TGridOption2 = int32

const (
	GoScrollToLastCol TGridOption2 = iota
	GoScrollToLastRow
	GoEditorParentColor
	GoEditorParentFont
	GoCopyWithoutTrailingLinebreak
)

// TGridOptions2 SET: TGridOption2
type TGridOptions2 = TSet

// TGridSaveOptions ENUM
type TGridSaveOptions = int32

const (
	SoDesign TGridSaveOptions = iota
	SoAttributes
	SoContent
	SoPosition
)

// TSaveOptions SET: TGridSaveOptions
type TSaveOptions = TSet

// TGridZone ENUM
type TGridZone = int32

const (
	GzNormal TGridZone = iota
	GzFixedCols
	GzFixedRows
	GzFixedCells
	GzInvalid
)

// TGridZoneSet SET: TGridZone
type TGridZoneSet = TSet

// THeaderPaintElement ENUM
type THeaderPaintElement = int32

const (
	HpeBackground THeaderPaintElement = iota
	HpeDropMark
	HpeHeaderGlyph
	HpeSortGlyph
	HpeText
)

// THeaderPaintElements SET: THeaderPaintElement
type THeaderPaintElements = TSet

// THeaderState ENUM
type THeaderState = int32

const (
	HsAutoSizing THeaderState = iota
	HsDragging
	HsDragPending
	HsLoading
	HsColumnWidthTracking
	HsColumnWidthTrackPending
	HsHeightTracking
	HsHeightTrackPending
	HsResizing
	HsScaling
	HsNeedScaling
)

// THeaderStates SET: THeaderState
type THeaderStates = TSet

// THitPosition ENUM
//
//	These flags are returned by the hit test method.
type THitPosition = int32

const (
	HiAbove THitPosition = iota
	HiBelow
	HiNowhere
	HiOnItem
	HiOnItemButton
	HiOnItemButtonExact
	HiOnItemCheckbox
	HiOnItemIndent
	HiOnItemLabel
	HiOnItemLeft
	HiOnItemRight
	HiOnNormalIcon
	HiOnStateIcon
	HiToLeft
	HiToRight
	HiUpperSplitter
	HiLowerSplitter
)

// THitPositions SET: THitPosition
type THitPositions = TSet

// THitTest ENUM
type THitTest = int32

const (
	HtAbove THitTest = iota
	HtBelow
	HtNowhere
	HtOnItem
	HtOnButton
	HtOnIcon
	HtOnIndent
	HtOnLabel
	HtOnRight
	HtOnStateIcon
	HtToLeft
	HtToRight
)

// THitTests SET: THitTest
type THitTests = TSet

// THookedCommandFlag ENUM
type THookedCommandFlag = int32

const (
	HcfInit THookedCommandFlag = iota
	HcfPreExec
	HcfPostExec
	HcfFinish
)

// THookedCommandFlags SET: THookedCommandFlag
type THookedCommandFlags = TSet

// TIniFileOption ENUM
type TIniFileOption = int32

const (
	IfoStripComments TIniFileOption = iota
	IfoStripInvalid
	IfoEscapeLineFeeds
	IfoCaseSensitive
	IfoStripQuotes
	IfoFormatSettingsActive
	IfoWriteStringBoolean
)

// TIniFileOptions SET: TIniFileOption
type TIniFileOptions = TSet

// TKeyOption ENUM
type TKeyOption = int32

const (
	KeyEdit TKeyOption = iota
	KeyAdd
	KeyDelete
	KeyUnique
)

// TKeyOptions SET: TKeyOption
type TKeyOptions = TSet

// TListBoxOption ENUM
type TListBoxOption = int32

const (
	LboDrawFocusRect TListBoxOption = iota
)

// TListBoxOptions SET: TListBoxOption
type TListBoxOptions = TSet

// TListHotTrackStyle ENUM
type TListHotTrackStyle = int32

const (
	HtHandPoint TListHotTrackStyle = iota
	HtUnderlineCold
	HtUnderlineHot
)

// TListHotTrackStyles SET: TListHotTrackStyle
type TListHotTrackStyles = TSet

// TListItemsFlag ENUM
type TListItemsFlag = int32

const (
	LisfWSItemsCreated TListItemsFlag = iota
)

// TListItemsFlags SET: TListItemsFlag
type TListItemsFlags = TSet

// TListItemState ENUM
type TListItemState = int32

const (
	LisCut TListItemState = iota
	LisDropTarget
	LisFocused
	LisSelected
)

// TListItemStates SET: TListItemState
type TListItemStates = TSet

// TMouseButton ENUM
type TMouseButton = int32

const (
	MbLeft TMouseButton = iota
	MbRight
	MbMiddle
	MbExtra1
	MbExtra2
)

// TCaptureMouseButtons SET: TMouseButton
type TCaptureMouseButtons = TSet

// TMultiSelectStyles ENUM
type TMultiSelectStyles = int32

const (
	MsControlSelect TMultiSelectStyles = iota
	MsShiftSelect
	MsVisibleOnly
	MsSiblingOnly
)

// TMultiSelectStyle SET: TMultiSelectStyles
type TMultiSelectStyle = TSet

// TNodeState ENUM
type TNodeState = int32

const (
	NsCut TNodeState = iota
	NsDropHilited
	NsFocused
	NsSelected
	NsMultiSelected
	NsExpanded
	NsHasChildren
	NsDeleting
	NsVisible
	NsEnabled
	NsBound
	NsValidHasChildren
)

// TNodeStates SET: TNodeState
type TNodeStates = TSet

// TOnBeforeExeucteFlag ENUM
type TOnBeforeExeucteFlag = int32

const (
	BefAbort TOnBeforeExeucteFlag = iota
)

// TOnBeforeExeucteFlags SET: TOnBeforeExeucteFlag
type TOnBeforeExeucteFlags = TSet

// TOpenGLControlOption ENUM
type TOpenGLControlOption = int32

const (
	OcoMacRetinaMode TOpenGLControlOption = iota
	OcoRenderAtDesignTime
)

// TOpenGLControlOptions SET: TOpenGLControlOption
type TOpenGLControlOptions = TSet

// TOpenOption ENUM
//
//	TOpenDialog
type TOpenOption = int32

const (
	OfReadOnly TOpenOption = iota
	OfOverwritePrompt
	OfHideReadOnly
	OfNoChangeDir
	OfShowHelp
	OfNoValidate
	OfAllowMultiSelect
	OfExtensionDifferent
	OfPathMustExist
	OfFileMustExist
	OfCreatePrompt
	OfShareAware
	OfNoReadOnlyReturn
	OfNoTestFileCreate
	OfNoNetworkButton
	OfNoLongNames
	OfOldStyleDialog
	OfNoDereferenceLinks
	OfNoResolveLinks
	OfEnableIncludeNotify
	OfEnableSizing
	OfDontAddToRecent
	OfForceShowHidden
	OfViewDetail
	OfAutoPreview
)

// TOpenOptions SET: TOpenOption
type TOpenOptions = TSet

// TOpenOptionEx ENUM
//
//	WS specific options that cannot be (more or less) mapped to the standard TOpenOptions
//	Currently just Windows Vista+ (IFileDialog) options
type TOpenOptionEx = int32

const (
	OfHidePinnedPlaces TOpenOptionEx = iota
	OfStrictFileTypes
	OfPickFolders
	OfOkButtonNeedsInteraction
	OfForceFileSystem
	OfAllNonStorageItems
)

// TOpenOptionsEx SET: TOpenOptionEx
type TOpenOptionsEx = TSet

// TOwnerDrawStateType ENUM
//
//	ToDo: move this to StdCtrls
type TOwnerDrawStateType = int32

const (
	OdSelected TOwnerDrawStateType = iota
	OdGrayed
	OdDisabled
	OdChecked
	OdFocused
	OdDefault
	OdHotLight
	OdInactive
	OdNoAccel
	OdNoFocusRect
	OdReserved1
	OdReserved2
	OdComboBoxEdit
	OdBackgroundPainted
)

// TOwnerDrawState SET: TOwnerDrawStateType
type TOwnerDrawState = TSet

// TPageSetupDialogOption ENUM
type TPageSetupDialogOption = int32

const (
	PsoDefaultMinMargins TPageSetupDialogOption = iota
	PsoDisableMargins
	PsoDisableOrientation
	PsoDisablePagePainting
	PsoDisablePaper
	PsoDisablePrinter
	PsoMargins
	PsoMinMargins
	PsoShowHelp
	PsoWarning
	PsoNoNetworkButton
)

// TPageSetupDialogOptions SET: TPageSetupDialogOption
type TPageSetupDialogOptions = TSet

// TPanelPart ENUM
type TPanelPart = int32

const (
	PpText TPanelPart = iota
	PpBorder
	PpWidth
)

// TPanelParts SET: TPanelPart
type TPanelParts = TSet

// TPrintDialogOption ENUM
type TPrintDialogOption = int32

const (
	PoPrintToFile TPrintDialogOption = iota
	PoPageNums
	PoSelection
	PoWarning
	PoHelp
	PoDisablePrintToFile
	PoBeforeBeginDoc
)

// TPrintDialogOptions SET: TPrintDialogOption
type TPrintDialogOptions = TSet

// TRawImageQueryFlag ENUM
type TRawImageQueryFlag = int32

const (
	RiqfMono TRawImageQueryFlag = iota
	RiqfGrey
	RiqfRGB
	RiqfAlpha
	RiqfMask
	RiqfPalette
	RiqfUpdate
)

// TRawImageQueryFlags SET: TRawImageQueryFlag
type TRawImageQueryFlags = TSet

// TSectionValuesOption ENUM
type TSectionValuesOption = int32

const (
	SvoIncludeComments TSectionValuesOption = iota
	SvoIncludeInvalid
	SvoIncludeQuotes
)

// TSectionValuesOptions SET: TSectionValuesOption
type TSectionValuesOptions = TSet

// TShiftStateEnum ENUM
//
//	Types used by standard events
type TShiftStateEnum = int32

const (
	SsShift TShiftStateEnum = iota
	SsAlt
	SsCtrl
	SsLeft
	SsRight
	SsMiddle
	SsDouble
	SsMeta
	SsSuper
	SsHyper
	SsAltGr
	SsCaps
	SsNum
	SsScroll
	SsTriple
	SsQuad
	SsExtra1
	SsExtra2
)

// TShiftState SET: TShiftStateEnum
type TShiftState = TSet

// TSizeConstraintsOption ENUM
type TSizeConstraintsOption = int32

const (
	ScoAdviceWidthAsMin TSizeConstraintsOption = iota
	ScoAdviceWidthAsMax
	ScoAdviceHeightAsMin
	ScoAdviceHeightAsMax
)

// TSizeConstraintsOptions SET: TSizeConstraintsOption
type TSizeConstraintsOptions = TSet

// TStringsOption ENUM
type TStringsOption = int32

const (
	SoStrictDelimiter TStringsOption = iota
	SoWriteBOM
	SoTrailingLineBreak
	SoUseLocale
	SoPreserveBOM
)

// TStringsOptions SET: TStringsOption
type TStringsOptions = TSet

// TSynCoordinateMappingFlag ENUM
type TSynCoordinateMappingFlag = int32

const (
	ScmLimitToLines TSynCoordinateMappingFlag = iota
	ScmIncludePartVisible
	ScmForceLeftSidePos
)

// TSynCoordinateMappingFlags SET: TSynCoordinateMappingFlag
type TSynCoordinateMappingFlags = TSet

// TSynCustomFoldConfigMode ENUM
type TSynCustomFoldConfigMode = int32

const (
	FmFold TSynCustomFoldConfigMode = iota
	FmHide
	FmMarkup
	FmOutline
)

// TSynCustomFoldConfigModes SET: TSynCustomFoldConfigMode
type TSynCustomFoldConfigModes = TSet

// TSynEditHasTextFlag ENUM
type TSynEditHasTextFlag = int32

const (
	ShtIncludeVirtual TSynEditHasTextFlag = iota
)

// TSynEditHasTextFlags SET: TSynEditHasTextFlag
type TSynEditHasTextFlags = TSet

// TSynEditorMouseOption ENUM
type TSynEditorMouseOption = int32

const (
	EmUseMouseActions TSynEditorMouseOption = iota
	EmAltSetsColumnMode
	EmDragDropEditing
	EmRightMouseMovesCursor
	EmDoubleClickSelectsLine
	EmShowCtrlMouseLinks
	EmCtrlWheelZoom
)

// TSynEditorMouseOptions SET: TSynEditorMouseOption
type TSynEditorMouseOptions = TSet

// TSynEditorOption ENUM
type TSynEditorOption = int32

const (
	EoAutoIndent TSynEditorOption = iota
	EoBracketHighlight
	EoEnhanceHomeKey
	EoGroupUndo
	EoHalfPageScroll
	EoHideRightMargin
	EoKeepCaretX
	EoNoCaret
	EoNoSelection
	EoPersistentCaret
	EoScrollByOneLess
	EoScrollPastEof
	EoScrollPastEol
	EoScrollHintFollows
	EoShowScrollHint
	EoShowSpecialChars
	EoSmartTabs
	EoTabIndent
	EoTabsToSpaces
	EoTrimTrailingSpaces
	EoAutoSizeMaxScrollWidth
	EoDisableScrollArrows
	EoHideShowScrollbars
	EoDropFiles
	EoSmartTabDelete
	EoSpacesToTabs
	EoAutoIndentOnPaste
	EoAltSetsColumnMode
	EoDragDropEditing
	EoRightMouseMovesCursor
	EoDoubleClickSelectsLine
	EoShowCtrlMouseLinks
)

// TSynEditorOptions SET: TSynEditorOption
type TSynEditorOptions = TSet

// TSynEditorOption2 ENUM
type TSynEditorOption2 = int32

const (
	EoCaretSkipsSelection TSynEditorOption2 = iota
	EoCaretMoveEndsSelection
	EoCaretSkipTab
	EoAlwaysVisibleCaret
	EoEnhanceEndKey
	EoFoldedCopyPaste
	EoPersistentBlock
	EoOverwriteBlock
	EoAutoHideCursor
	EoColorSelectionTillEol
	EoPersistentCaretStopBlink
	EoNoScrollOnSelectRange
	EoAcceptDragDropEditing
	EoScrollPastEolAddPage
	EoScrollPastEolAutoCaret
	EoBookmarkRestoresScroll
)

// TSynEditorOptions2 SET: TSynEditorOption2
type TSynEditorOptions2 = TSet

// TSynEditorShareOption ENUM
//
//	options for textbuffersharing
type TSynEditorShareOption = int32

const (
	EosShareMarks TSynEditorShareOption = iota
)

// TSynEditorShareOptions SET: TSynEditorShareOption
type TSynEditorShareOptions = TSet

// TSynEditTextFlag ENUM
type TSynEditTextFlag = int32

const (
	SetSelect TSynEditTextFlag = iota
	SetPersistentBlock
	SetMoveBlock
	SetExtendBlock
)

// TSynEditTextFlags SET: TSynEditTextFlag
type TSynEditTextFlags = TSet

// TSynFoldAction ENUM
type TSynFoldAction = int32

const (
	SfaOpen TSynFoldAction = iota
	SfaClose
	SfaFold
	SfaFoldFold
	SfaFoldHide
	SfaMultiLine
	SfaSingleLine
	SfaCloseForNextLine
	SfaLastLineClose
	SfaCloseAndOpen
	SfaDefaultCollapsed
	SfaMarkup
	SfaOutline
	SfaOutlineKeepLevel
	SfaOutlineMergeParent
	SfaOutlineForceIndent
	SfaOutlineNoColor
	SfaOutlineNoLine
	SfaInvalid
	SfaOpenFold
	SfaCloseFold
	SfaOneLineOpen
	SfaOneLineClose
)

// TSynFoldActions SET: TSynFoldAction
type TSynFoldActions = TSet

// TSynFoldBlockFilterFlag ENUM
//   - TSynFoldBlockFilter
//     used to specify which folds to include for:
//   - FoldOpenCount, FoldCloseCount, FoldNestCount
//   - maybe in future TLazSynFoldNodeInfoList
//     TLazSynFoldNodeInfoList has additional filters
//     TLazSynFoldNodeInfoList always uses the full set (sfbIncludeDisabled)
//     A Highlighter is not required to implement this, or can choose to implement
//     a subset only. For any field/value a Highlighter may simple assume default.
//   - Highlighter that have only one "FoldGroup" do not require this.
//   - Highlighter that do not store foldblocks that are unavailable (e.g. off by
//     config) always return the same set
//     Using a record, as argument is the virtual methods, allows one to add further
//     fields/values, without breaking inheritance.
//     New fields values are expected to be ignored (handled as default) by existing
//     highlighter.
//     Callers of the method can:
//   - use InitFoldBlockFilter to make sure all fields are set to default
//   - use (none virtual) wrapper methods
type TSynFoldBlockFilterFlag = int32

const (
	SfbIncludeDisabled TSynFoldBlockFilterFlag = iota
)

// TSynFoldBlockFilterFlags SET: TSynFoldBlockFilterFlag
type TSynFoldBlockFilterFlags = TSet

// TSynHighlighterCapability ENUM
type TSynHighlighterCapability = int32

const (
	HcUserSettings TSynHighlighterCapability = iota
	HcRegistry
	HcCodeFolding
)

// TSynHighlighterCapabilities SET: TSynHighlighterCapability
type TSynHighlighterCapabilities = TSet

// TSynLogPhysFlag ENUM
type TSynLogPhysFlag = int32

const (
	LpfAdjustToCharBegin TSynLogPhysFlag = iota
	LpfAdjustToNextChar
)

// TSynLogPhysFlags SET: TSynLogPhysFlag
type TSynLogPhysFlags = TSet

// TSynMAUpRestriction ENUM
type TSynMAUpRestriction = int32

const (
	CrLastDownPos TSynMAUpRestriction = iota
	CrLastDownPosSameLine
	CrLastDownPosSearchAll
	CrLastDownButton
	CrLastDownShift
	CrAllowFallback
)

// TSynMAUpRestrictions SET: TSynMAUpRestriction
type TSynMAUpRestrictions = TSet

// TSynPluginSyncroScanMode ENUM
//   - TSynPluginCustomSyncroEdit implements:
//   - Locking of TrimTrailingSpace
//   - CurrentCell follows Caret / LastCell
//   - DeActivate if Edit outside Cell
//   - DeActivate on undo/redo if needed
//   - various helpers, to set the caret/block
type TSynPluginSyncroScanMode = int32

const (
	SpssNoCase TSynPluginSyncroScanMode = iota
	SpssWithCase
	SpssCtxNoCase
	SpssCtxWithCase
)

// TSynPluginSyncroScanModes SET: TSynPluginSyncroScanMode
type TSynPluginSyncroScanModes = TSet

// TSynSearchOption ENUM
type TSynSearchOption = int32

const (
	SsoMatchCase TSynSearchOption = iota
	SsoWholeWord
	SsoBackwards
	SsoEntireScope
	SsoSelectedOnly
	SsoReplace
	SsoReplaceAll
	SsoPrompt
	SsoSearchInReplacement
	SsoRegExpr
	SsoRegExprMultiLine
	SsoFindContinue
)

// TSynSearchOptions SET: TSynSearchOption
type TSynSearchOptions = TSet

// TSynStatusChange ENUM
type TSynStatusChange = int32

const (
	ScCaretX TSynStatusChange = iota
	ScCaretY
	ScLeftChar
	ScTopLine
	ScLinesInWindow
	ScCharsInWindow
	ScInsertMode
	ScModified
	ScSelection
	ScReadOnly
	ScFocus
	ScOptions
)

// TSynStatusChanges SET: TSynStatusChange
type TSynStatusChanges = TSet

// TSynVisibleSpecialChar ENUM
type TSynVisibleSpecialChar = int32

const (
	VscSpace TSynVisibleSpecialChar = iota
	VscTabAtFirst
	VscTabAtLast
)

// TSynVisibleSpecialChars SET: TSynVisibleSpecialChar
type TSynVisibleSpecialChars = TSet

// TTaskDialogCommonButton ENUM
type TTaskDialogCommonButton = int32

const (
	TcbOk TTaskDialogCommonButton = iota
	TcbYes
	TcbNo
	TcbCancel
	TcbRetry
	TcbClose
)

// TTaskDialogCommonButtons SET: TTaskDialogCommonButton
type TTaskDialogCommonButtons = TSet

// TTaskDialogFlag ENUM
type TTaskDialogFlag = int32

const (
	TfEnableHyperlinks TTaskDialogFlag = iota
	TfUseHiconMain
	TfUseHiconFooter
	TfAllowDialogCancellation
	TfUseCommandLinks
	TfUseCommandLinksNoIcon
	TfExpandFooterArea
	TfExpandedByDefault
	TfVerificationFlagChecked
	TfShowProgressBar
	TfShowMarqueeProgressBar
	TfCallbackTimer
	TfPositionRelativeToWindow
	TfRtlLayout
	TfNoDefaultRadioButton
	TfCanBeMinimized
	TfNoSetForeGround
	TfSizeToContent
	TfForceNonNative
	TfEmulateClassicStyle
	TfQuery
	TfSimpleQuery
	TfQueryFixedChoices
	TfQueryFocused
)

// TTaskDialogFlags SET: TTaskDialogFlag
type TTaskDialogFlags = TSet

// TTreeViewOption ENUM
type TTreeViewOption = int32

const (
	TvoAllowMultiselect TTreeViewOption = iota
	TvoAutoExpand
	TvoAutoInsertMark
	TvoAutoItemHeight
	TvoHideSelection
	TvoHotTrack
	TvoKeepCollapsedNodes
	TvoReadOnly
	TvoRightClickSelect
	TvoRowSelect
	TvoShowButtons
	TvoShowLines
	TvoShowRoot
	TvoShowSeparators
	TvoToolTips
	TvoNoDoubleClickExpand
	TvoThemedDraw
	TvoEmptySpaceUnselect
)

// TTreeViewOptions SET: TTreeViewOption
type TTreeViewOptions = TSet

// TViewedXYInfoFlag ENUM
type TViewedXYInfoFlag = int32

const (
	VifAdjustLogXYToNextChar TViewedXYInfoFlag = iota
	VifReturnPhysXY
	VifReturnLogXY
	VifReturnLogEOL
	VifReturnPhysOffset
)

// TViewedXYInfoFlags SET: TViewedXYInfoFlag
type TViewedXYInfoFlags = TSet

// TVirtualNodeInitState ENUM
//
//	States used in InitNode to indicate states a node shall initially have.
type TVirtualNodeInitState = int32

const (
	IvsDisabled TVirtualNodeInitState = iota
	IvsExpanded
	IvsHasChildren
	IvsMultiline
	IvsSelected
	IvsFiltered
	IvsReInit
)

// TVirtualNodeInitStates SET: TVirtualNodeInitState
type TVirtualNodeInitStates = TSet

// TVirtualNodeState ENUM
//
//	Be careful when adding new states as this might change the size of the type which in turn
//	changes the alignment in the node record as well as the stream chunks.
//	Do not reorder the states and always add new states at the end of this enumeration in order to avoid
//	breaking existing code.
type TVirtualNodeState = int32

const (
	VsInitialized TVirtualNodeState = iota
	VsChecking
	VsCutOrCopy
	VsDisabled
	VsDeleting
	VsExpanded
	VsHasChildren
	VsVisible
	VsSelected
	VsOnFreeNodeCallRequired
	VsAllChildrenHidden
	VsClearing
	VsMultiline
	VsHeightMeasured
	VsToggling
	VsFiltered
)

// TVirtualNodeStates SET: TVirtualNodeState
type TVirtualNodeStates = TSet

// TVirtualTreeState ENUM
type TVirtualTreeState = int32

const (
	TsCancelHintAnimation TVirtualTreeState = iota
	TsChangePending
	TsCheckPropagation
	TsCollapsing
	TsToggleFocusedSelection
	TsClearPending
	TsClipboardFlushing
	TsCopyPending
	TsCutPending
	TsDrawSelPending
	TsDrawSelecting
	TsEditing
	TsEditPending
	TsExpanding
	TsNodeHeightTracking
	TsNodeHeightTrackPending
	TsHint
	TsInAnimation
	TsIncrementalSearching
	TsIncrementalSearchPending
	TsIterating
	TsKeyCheckPending
	TsLeftButtonDown
	TsLeftDblClick
	TsMouseCheckPending
	TsMiddleButtonDown
	TsMiddleDblClick
	TsNeedRootCountUpdate
	TsOLEDragging
	TsOLEDragPending
	TsPainting
	TsRightButtonDown
	TsRightDblClick
	TsPopupMenuShown
	TsScrolling
	TsScrollPending
	TsSizing
	TsStopValidation
	TsStructureChangePending
	TsSynchMode
	TsThumbTracking
	TsToggling
	TsUpdateHiddenChildrenNeeded
	TsUpdating
	TsUseCache
	TsUserDragObject
	TsUseThemes
	TsValidating
	TsPreviouslySelectedLocked
	TsValidationNeeded
	TsVCLDragging
	TsVCLDragPending
	TsVCLDragFinished
	TsWheelPanning
	TsWheelScrolling
	TsWindowCreating
	TsUseExplorerTheme
)

// TVirtualTreeStates SET: TVirtualTreeState
type TVirtualTreeStates = TSet

// TVTAnimationOption ENUM
//
//	Options to toggle animation support:
type TVTAnimationOption = int32

const (
	ToAnimatedToggle TVTAnimationOption = iota
	ToAdvancedAnimatedToggle
)

// TVTAnimationOptions SET: TVTAnimationOption
type TVTAnimationOptions = TSet

// TVTAutoOption ENUM
//
//	Options which toggle automatic handling of certain situations:
type TVTAutoOption = int32

const (
	ToAutoDropExpand TVTAutoOption = iota
	ToAutoExpand
	ToAutoScroll
	ToAutoScrollOnExpand
	ToAutoSort
	ToAutoSpanColumns
	ToAutoTristateTracking
	ToAutoHideButtons
	ToAutoDeleteMovedNodes
	ToDisableAutoscrollOnFocus
	ToAutoChangeScale
	ToAutoFreeOnCollapse
	ToDisableAutoscrollOnEdit
	ToAutoBidiColumnOrdering
)

// TVTAutoOptions SET: TVTAutoOption
type TVTAutoOptions = TSet

// TVTColumnOption ENUM
//
//	Options per column.
type TVTColumnOption = int32

const (
	CoAllowClick TVTColumnOption = iota
	CoDraggable
	CoEnabled
	CoParentBidiMode
	CoParentColor
	CoResizable
	CoShowDropMark
	CoVisible
	CoAutoSpring
	CoFixed
	CoSmartResize
	CoAllowFocus
	CoDisableAnimatedResize
	CoWrapCaption
	CoUseCaptionAlignment
	CoEditable
)

// TVTColumnOptions SET: TVTColumnOption
type TVTColumnOptions = TSet

// TVTHeaderHitPosition ENUM
//
//	These flags are used to indicate where a click in the header happened.
type TVTHeaderHitPosition = int32

const (
	HhiNoWhere TVTHeaderHitPosition = iota
	HhiOnColumn
	HhiOnIcon
	HhiOnCheckbox
)

// TVTHeaderHitPositions SET: TVTHeaderHitPosition
type TVTHeaderHitPositions = TSet

// TVTHeaderOption ENUM
type TVTHeaderOption = int32

const (
	HoAutoResize TVTHeaderOption = iota
	HoColumnResize
	HoDblClickResize
	HoDrag
	HoHotTrack
	HoOwnerDraw
	HoRestrictDrag
	HoShowHint
	HoShowImages
	HoShowSortGlyphs
	HoVisible
	HoAutoSpring
	HoFullRepaintOnResize
	HoDisableAnimatedResize
	HoHeightResize
	HoHeightDblClickResize
	HoHeaderClickAutoSort
)

// TVTHeaderOptions SET: TVTHeaderOption
type TVTHeaderOptions = TSet

// TVTHeaderPopupOption ENUM
type TVTHeaderPopupOption = int32

const (
	PoOriginalOrder TVTHeaderPopupOption = iota
	PoAllowHideAll
	PoResizeToFitItem
)

// TVTHeaderPopupOptions SET: TVTHeaderPopupOption
type TVTHeaderPopupOptions = TSet

// TVTInternalPaintOption ENUM
//
//	options which determine what to draw in PaintTree
type TVTInternalPaintOption = int32

const (
	PoBackground TVTInternalPaintOption = iota
	PoColumnColor
	PoDrawFocusRect
	PoDrawSelection
	PoDrawDropMark
	PoGridLines
	PoMainOnly
	PoSelectedOnly
	PoUnbuffered
)

// TVTInternalPaintOptions SET: TVTInternalPaintOption
type TVTInternalPaintOptions = TSet

// TVTMiscOption ENUM
//
//	Options which do not fit into any of the other groups:
type TVTMiscOption = int32

const (
	ToAcceptOLEDrop TVTMiscOption = iota
	ToCheckSupport
	ToEditable
	ToFullRepaintOnResize
	ToGridExtensions
	ToInitOnSave
	ToReportMode
	ToToggleOnDblClick
	ToWheelPanning
	ToReadOnly
	ToVariableNodeHeight
	ToFullRowDrag
	ToNodeHeightResize
	ToNodeHeightDblClickResize
	ToEditOnClick
	ToEditOnDblClick
	ToReverseFullExpandHotKey
)

// TVTMiscOptions SET: TVTMiscOption
type TVTMiscOptions = TSet

// TVTPaintOption ENUM
//
//	There is a heap of switchable behavior in the tree. Since published properties may never exceed 4 bytes,
//	which limits sets to at most 32 members, and because for better overview tree options are splitted
//	in various sub-options and are held in a commom options class.
//
//	Options to customize tree appearance:
type TVTPaintOption = int32

const (
	ToHideFocusRect TVTPaintOption = iota
	ToHideSelection
	ToHotTrack
	ToPopupMode
	ToShowBackground
	ToShowButtons
	ToShowDropmark
	ToShowHorzGridLines
	ToShowRoot
	ToShowTreeLines
	ToShowVertGridLines
	ToThemeAware
	ToUseBlendedImages
	ToGhostedIfUnfocused
	ToFullVertGridLines
	ToAlwaysHideSelection
	ToUseBlendedSelection
	ToStaticBackground
	ToChildrenAbove
	ToFixedIndent
	ToUseExplorerTheme
	ToHideTreeLinesIfThemed
	ToShowFilteredNodes
)

// TVTPaintOptions SET: TVTPaintOption
type TVTPaintOptions = TSet

// TVTSelectionOption ENUM
//
//	Options which determine the tree's behavior when selecting nodes:
type TVTSelectionOption = int32

const (
	ToDisableDrawSelection TVTSelectionOption = iota
	ToExtendedFocus
	ToFullRowSelect
	ToLevelSelectConstraint
	ToMiddleClickSelect
	ToMultiSelect
	ToRightClickSelect
	ToSiblingSelectConstraint
	ToCenterScrollIntoView
	ToSimpleDrawSelection
	ToAlwaysSelectNode
	ToRestoreSelection
)

// TVTSelectionOptions SET: TVTSelectionOption
type TVTSelectionOptions = TSet

// TVTStringOption ENUM
//
//	--------- TCustomVirtualStringTree
//	Options regarding strings (useful only for the string tree and descendants):
type TVTStringOption = int32

const (
	ToSaveCaptions TVTStringOption = iota
	ToShowStaticText
	ToAutoAcceptEditChange
)

// TVTStringOptions SET: TVTStringOption
type TVTStringOptions = TSet

// TSortDirection ENUM Rename: ComCtrls.TSortDirection
type CCtrlsTSortDirection = int32

const (
	SdAscending CCtrlsTSortDirection = iota
	SdDescending
)

// TTaskDialogIcon ENUM Rename: Dialogs.TTaskDialogIcon
type DialogsTTaskDialogIcon = int32

const (
	TdiNone DialogsTTaskDialogIcon = iota
	TdiWarning
	TdiError
	TdiInformation
	TdiShield
	TdiQuestion
)

// TSortDirection ENUM Rename: laz.VirtualTrees.TSortDirection
type LVTreesTSortDirection = int32

const (
	LVTreesSdAscending LVTreesTSortDirection = iota
	LVTreesSdDescending
)

// TtkTokenKind ENUM Rename: SynHighlighterAny.TtkTokenKind
type SHAnyTtkTokenKind = int32

const (
	TkComment SHAnyTtkTokenKind = iota
	TkIdentifier
	TkKey
	TkNull
	TkNumber
	TkPreprocessor
	TkSpace
	TkString
	TkSymbol
	TkUnknown
	TkConstant
	TkObject
	TkEntity
	TkDollarVariable
	TkDot
)

// TtkTokenKind ENUM Rename: SynHighlighterCpp.TtkTokenKind
type SHCppTtkTokenKind = int32

const (
	TkAsm SHCppTtkTokenKind = iota
	SHCppTkComment
	TkDirective
	SHCppTkIdentifier
	SHCppTkKey
	SHCppTkNull
	SHCppTkNumber
	SHCppTkSpace
	SHCppTkString
	SHCppTkSymbol
	SHCppTkUnknown
)

// TxtkTokenKind ENUM Rename: SynHighlighterCpp.TxtkTokenKind
type SHCppTxtkTokenKind = int32

const (
	XtkAdd SHCppTxtkTokenKind = iota
	XtkAddAssign
	XtkAnd
	XtkAndAssign
	XtkArrow
	XtkAssign
	XtkBitComplement
	XtkBraceClose
	XtkBraceOpen
	XtkColon
	XtkComma
	XtkDecrement
	XtkDivide
	XtkDivideAssign
	XtkEllipse
	XtkGreaterThan
	XtkGreaterThanEqual
	XtkIncOr
	XtkIncOrAssign
	XtkIncrement
	XtkLessThan
	XtkLessThanEqual
	XtkLogAnd
	XtkLogComplement
	XtkLogEqual
	XtkLogOr
	XtkMod
	XtkModAssign
	XtkMultiplyAssign
	XtkNotEqual
	XtkPoint
	XtkQuestion
	XtkRoundClose
	XtkRoundOpen
	XtkScopeResolution
	XtkSemiColon
	XtkShiftLeft
	XtkShiftLeftAssign
	XtkShiftRight
	XtkShiftRightAssign
	XtkSquareClose
	XtkSquareOpen
	XtkStar
	XtkSubtract
	XtkSubtractAssign
	XtkXor
	XtkXorAssign
)

// TtkTokenKind ENUM Rename: SynHighlighterCss.TtkTokenKind
type SHCssTtkTokenKind = int32

const (
	SHCssTkComment SHCssTtkTokenKind = iota
	SHCssTkIdentifier
	SHCssTkKey
	SHCssTkNull
	SHCssTkNumber
	SHCssTkSpace
	SHCssTkString
	SHCssTkSymbol
	TkMeasurementUnit
	TkSelector
	SHCssTkUnknown
)

// TtkTokenKind ENUM Rename: SynHighlighterHTML.TtkTokenKind
type SHHTMLTtkTokenKind = int32

const (
	TkAmpersand SHHTMLTtkTokenKind = iota
	TkASP
	TkCDATA
	SHHTMLTkComment
	SHHTMLTkIdentifier
	SHHTMLTkKey
	SHHTMLTkNull
	SHHTMLTkSpace
	SHHTMLTkString
	SHHTMLTkSymbol
	TkText
	TkUndefKey
	TkValue
	TkDOCTYPE
)

// TtkTokenKind ENUM Rename: SynHighlighterJava.TtkTokenKind
type SHJavaTtkTokenKind = int32

const (
	SHJavaTkComment SHJavaTtkTokenKind = iota
	TkDocument
	SHJavaTkIdentifier
	TkInvalid
	SHJavaTkKey
	SHJavaTkNull
	SHJavaTkNumber
	SHJavaTkSpace
	SHJavaTkString
	SHJavaTkSymbol
	SHJavaTkUnknown
	TkAnnotation
)

// TxtkTokenKind ENUM Rename: SynHighlighterJava.TxtkTokenKind
type SHJavaTxtkTokenKind = int32

const (
	SHJavaXtkAdd SHJavaTxtkTokenKind = iota
	SHJavaXtkAddAssign
	SHJavaXtkAnd
	SHJavaXtkAndAssign
	SHJavaXtkAssign
	SHJavaXtkBitComplement
	SHJavaXtkBraceClose
	SHJavaXtkBraceOpen
	SHJavaXtkColon
	XtkCondAnd
	XtkCondOr
	SHJavaXtkDecrement
	SHJavaXtkDivide
	SHJavaXtkDivideAssign
	SHJavaXtkGreaterThan
	SHJavaXtkGreaterThanEqual
	SHJavaXtkIncOr
	SHJavaXtkIncOrAssign
	SHJavaXtkIncrement
	SHJavaXtkLessThan
	SHJavaXtkLessThanEqual
	SHJavaXtkLogComplement
	SHJavaXtkLogEqual
	XtkMultiply
	SHJavaXtkMultiplyAssign
	SHJavaXtkNotEqual
	SHJavaXtkPoint
	SHJavaXtkQuestion
	XtkRemainder
	XtkRemainderAssign
	SHJavaXtkRoundClose
	SHJavaXtkRoundOpen
	SHJavaXtkSemiColon
	SHJavaXtkShiftLeft
	SHJavaXtkShiftLeftAssign
	SHJavaXtkShiftRight
	SHJavaXtkShiftRightAssign
	SHJavaXtkSquareClose
	SHJavaXtkSquareOpen
	SHJavaXtkSubtract
	SHJavaXtkSubtractAssign
	XtkUnsignShiftRight
	XtkUnsignShiftRightAssign
	SHJavaXtkXor
	SHJavaXtkXorAssign
	SHJavaXtkComma
	XtkNonSymbol
)

// TtkTokenKind ENUM Rename: SynHighlighterJScript.TtkTokenKind
type SHJScriptTtkTokenKind = int32

const (
	SHJScriptTkComment SHJScriptTtkTokenKind = iota
	SHJScriptTkIdentifier
	SHJScriptTkKey
	SHJScriptTkNull
	SHJScriptTkNumber
	SHJScriptTkSpace
	SHJScriptTkString
	SHJScriptTkSymbol
	TkBracket
	SHJScriptTkUnknown
	TkNonReservedKey
	TkEvent
)

// TtkTokenKind ENUM Rename: SynHighlighterPython.TtkTokenKind
type SHPythonTtkTokenKind = int32

const (
	SHPythonTkComment SHPythonTtkTokenKind = iota
	SHPythonTkIdentifier
	SHPythonTkKey
	SHPythonTkNull
	SHPythonTkNumber
	SHPythonTkSpace
	SHPythonTkString
	SHPythonTkSymbol
	TkNonKeyword
	TkTripleQuotedString
	TkSystemDefined
	TkHex
	TkOct
	TkFloat
	SHPythonTkUnknown
)

// TtkTokenKind ENUM Rename: SynHighlighterSQL.TtkTokenKind
type SHSQLTtkTokenKind = int32

const (
	SHSQLTkComment SHSQLTtkTokenKind = iota
	TkDatatype
	TkDefaultPackage
	TkException
	TkFunction
	SHSQLTkIdentifier
	SHSQLTkKey
	SHSQLTkNull
	SHSQLTkNumber
	SHSQLTkSpace
	TkPLSQL
	TkSQLPlus
	SHSQLTkString
	SHSQLTkSymbol
	TkTableName
	SHSQLTkUnknown
	TkVariable
)

// TtkTokenKind ENUM Rename: SynHighlighterXML.TtkTokenKind
type SHXMLTtkTokenKind = int32

const (
	TkAposAttrValue SHXMLTtkTokenKind = iota
	TkAposEntityRef
	TkAttribute
	SHXMLTkCDATA
	SHXMLTkComment
	TkElement
	TkEntityRef
	TkEqual
	SHXMLTkNull
	TkProcessingInstruction
	TkQuoteAttrValue
	TkQuoteEntityRef
	SHXMLTkSpace
	SHXMLTkSymbol
	SHXMLTkText
	TknsAposAttrValue
	TknsAposEntityRef
	TknsAttribute
	TknsEqual
	TknsQuoteAttrValue
	TknsQuoteEntityRef
	SHXMLTkDocType
)

// TActionListState ENUM
type TActionListState = int32

const (
	AsNormal TActionListState = iota
	AsSuspended
	AsSuspendedEnabled
)

// TAddPopupItemType ENUM
type TAddPopupItemType = int32

const (
	ApNormal TAddPopupItemType = iota
	ApDisabled
	ApHidden
)

// TAlign ENUM
type TAlign = int32

const (
	AlNone TAlign = iota
	AlTop
	AlBottom
	AlLeft
	AlRight
	AlClient
	AlCustom
)

// TAlignment ENUM
//
//	Text alignment types
type TAlignment = int32

const (
	TaLeftJustify TAlignment = iota
	TaRightJustify
	TaCenter
)

// TAnchorSideReference ENUM
type TAnchorSideReference = int32

const (
	AsrTop TAnchorSideReference = iota
	AsrBottom
	AsrCenter
)

// TAntialiasingMode ENUM
type TAntialiasingMode = int32

const (
	AmDontCare TAntialiasingMode = iota
	AmOn
	AmOff
)

// TApplicationDoubleBuffered ENUM
type TApplicationDoubleBuffered = int32

const (
	AdbDefault TApplicationDoubleBuffered = iota
	AdbFalse
	AdbTrue
)

// TApplicationExceptionDlg ENUM
type TApplicationExceptionDlg = int32

const (
	AedOkCancelDialog TApplicationExceptionDlg = iota
	AedOkMessageBox
)

// TApplicationShowGlyphs ENUM
type TApplicationShowGlyphs = int32

const (
	SbgAlways TApplicationShowGlyphs = iota
	SbgNever
	SbgSystem
)

// TApplicationType ENUM
//
//	This identifies the kind of device where the application currently runs on
//	Note that the same application can run in all kinds of devices if it has a
//	user interface flexible enough
type TApplicationType = int32

const (
	AtDefault TApplicationType = iota
	AtDesktop
	AtPDA
	AtKeyPadDevice
	AtTablet
	AtTV
	AtMobileEmulator
)

// TArrowShape ENUM
type TArrowShape = int32

const (
	AsClassicSmaller TArrowShape = iota
	AsClassicLarger
	AsModernSmaller
	AsModernLarger
	AsYetAnotherShape
	AsTheme
)

// TAutoAdvance ENUM
type TAutoAdvance = int32

const (
	AaNone TAutoAdvance = iota
	AaDown
	AaRight
	AaLeft
	AaRightDown
	AaLeftDown
	AaRightUp
	AaLeftUp
)

// TBalloonFlags ENUM
//
//	TCustomTrayIcon
type TBalloonFlags = int32

const (
	BfNone TBalloonFlags = iota
	BfInfo
	BfWarning
	BfError
)

// TBandDrawingStyle ENUM
//
//	TControlBar
type TBandDrawingStyle = int32

const (
	DsNormal TBandDrawingStyle = iota
	DsGradient
)

// TBevelShape ENUM
type TBevelShape = int32

const (
	BsBox TBevelShape = iota
	BsFrame
	BsTopLine
	BsBottomLine
	BsLeftLine
	BsRightLine
	BsSpacer
)

// TBevelStyle ENUM
//
//	TBevel
type TBevelStyle = int32

const (
	BsLowered TBevelStyle = iota
	BsRaised
)

// TBiDiMode ENUM
type TBiDiMode = int32

const (
	BdLeftToRight TBiDiMode = iota
	BdRightToLeft
	BdRightToLeftNoAlign
	BdRightToLeftReadingOnly
)

// TBitBtnKind ENUM
//
//	TCustomBitBtn
//	when adding items here, also update TBitBtn.GetCaptionOfKind
type TBitBtnKind = int32

const (
	BkCustom TBitBtnKind = iota
	BkOK
	BkCancel
	BkHelp
	BkYes
	BkNo
	BkClose
	BkAbort
	BkRetry
	BkIgnore
	BkAll
	BkNoToAll
	BkYesToAll
)

// TBitmapHandleType ENUM
//
//	TCustomBitmapImage
//	Descendent of TSharedImage for TCustomBitmap. If a TCustomBitmap is assigned to another
//	TCustomBitmap, only the reference count will be increased and both will share the
//	same TCustomBitmapImage
type TBitmapHandleType = int32

const (
	BmDIB TBitmapHandleType = iota
	BmDDB
)

// TButtonLayout ENUM
type TButtonLayout = int32

const (
	BlGlyphLeft TButtonLayout = iota
	BlGlyphRight
	BlGlyphTop
	BlGlyphBottom
)

// TButtonOrder ENUM
type TButtonOrder = int32

const (
	BoDefault TButtonOrder = iota
	BoCloseCancelOK
	BoCloseOKCancel
)

// TCalculatorLayout ENUM
type TCalculatorLayout = int32

const (
	ClNormal TCalculatorLayout = iota
	ClSimple
)

// TCalDayOfWeek ENUM
type TCalDayOfWeek = int32

const (
	DowMonday TCalDayOfWeek = iota
	DowTuesday
	DowWednesday
	DowThursday
	DowFriday
	DowSaturday
	DowSunday
	DowDefault
)

// TCalendarPart ENUM
type TCalendarPart = int32

const (
	CpNoWhere TCalendarPart = iota
	CpDate
	CpWeekNumber
	CpTitle
	CpTitleBtn
	CpTitleMonth
	CpTitleYear
)

// TCalendarView ENUM
//
//	In Windows since Vista native calendar control has four possible views.
//	In other widgetsets, as well as in older windows, calendar can only have
//	standard "month view" - grid with days representing a month.
type TCalendarView = int32

const (
	CvMonth TCalendarView = iota
	CvYear
	CvDecade
	CvCentury
)

// TCellHintPriority ENUM
type TCellHintPriority = int32

const (
	ChpAll TCellHintPriority = iota
	ChpAllNoDefault
	ChpTruncOnly
)

// TCellProcessType ENUM
//
//	The grid can display three types of hint: the default hint (Hint property),
//	individual cell hints (OnCellHint event), and hints for truncated cells.
//	TCellHintPriority determines how the overall hint is combined when more
//	multiple hint texts are to be displayed.
type TCellProcessType = int32

const (
	CpCopy TCellProcessType = iota
	CpPaste
)

// TChangeReason ENUM
type TChangeReason = int32

const (
	CrIgnore TChangeReason = iota
	CrAccumulated
	CrChildAdded
	CrChildDeleted
	CrNodeAdded
	CrNodeCopied
	CrNodeMoved
)

// TCheckBoxState ENUM
//
//	TCustomCheckBox
type TCheckBoxState = int32

const (
	CbUnchecked TCheckBoxState = iota
	CbChecked
	CbGrayed
)

// TCheckImageKind ENUM
type TCheckImageKind = int32

const (
	CkLightCheck TCheckImageKind = iota
	CkDarkCheck
	CkLightTick
	CkDarkTick
	CkFlat
	CkXP
	CkCustom
	CkSystemFlat
	CkSystemDefault
)

// TCheckState ENUM
//
//	The check states include both, transient and fluent (temporary) states. The only temporary state defined so
//	far is the pressed state.
type TCheckState = int32

const (
	CsUncheckedNormal TCheckState = iota
	CsUncheckedPressed
	CsCheckedNormal
	CsCheckedPressed
	CsMixedNormal
	CsMixedPressed
)

// TCheckType ENUM
type TCheckType = int32

const (
	CtNone TCheckType = iota
	CtTriStateCheckBox
	CtCheckBox
	CtRadioButton
	CtButton
)

// TChildControlResizeStyle ENUM
//
//	TControlChildSizing
//	LeftRightSpacing, TopBottomSpacing: Integer;
//	minimum space between left client border and left most children.
//	For example: ClientLeftRight=5 means children Left position is at least 5.
//	HorizontalSpacing, VerticalSpacing: Integer;
//	minimum space between each child horizontally
//	Defines how child controls are resized/aligned.
//	crsAnchorAligning
//	Anchors and Align work like Delphi. For example if Anchors property of
//	the control is [akLeft], it means fixed distance between left border of
//	parent's client area. [akRight] means fixed distance between right
//	border of the control and the right border of the parent's client area.
//	When the parent is resized the child is moved to keep the distance.
//	[akLeft,akRight] means fixed distance to left border and fixed distance
//	to right border. When the parent is resized, the controls width is
//	changed (resized) to keep the left and right distance.
//	Same for akTop,akBottom.
//	Align=alLeft for a control means set Left leftmost, Top topmost and
//	maximize Height. The width is kept, if akRight is not set. If akRight
//	is set in the Anchors property, then the right distance is kept and
//	the control's width is resized.
//	If there several controls with Align=alLeft, they will not overlapp and
//	be put side by side.
//	Same for alRight, alTop, alBottom. (Always expand 3 sides).
//	Align=alClient. The control will fill the whole remaining space.
//	Setting two children to Align=alClient does only make sense, if you set
//	maximum Constraints.
//	Order: First all alTop children are resized, then alBottom, then alLeft,
//	then alRight and finally alClient.
//	crsScaleChilds
//	Scale children, keep space between them fixed.
//	Children are resized to their normal/adviced size. If there is some space
//	left in the client area of the parent, then the children are scaled to
//	fill the space. You can set maximum Constraints. Then the other children
//	are scaled more.
//	For example: 3 child controls A, B, C with A.Width=10, B.Width=20 and
//	C.Width=30 (total=60). If the Parent's client area has a ClientWidth of
//	120, then the children are scaled with Factor 2.
//	If B has a maximum constraint width of 30, then first the children will be
//	scaled with 1.5 (A.Width=15, B.Width=30, C.Width=45). Then A and C
//	(15+45=60 and 30 pixel space left) will be scaled by 1.5 again, to a
//	final result of: A.Width=23, B.Width=30, C.Width=67 (23+30+67=120).
//	crsHomogenousChildResize
//	Enlarge children equally.
//	Children are resized to their normal/adviced size. If there is some space
//	left in the client area of the parent, then the remaining space is
//	distributed equally to each child.
//	For example: 3 child controls A, B, C with A.Width=10, B.Width=20 and
//	C.Width=30 (total=60). If the Parent's client area has a ClientWidth of
//	120, then 60/3=20 is added to each Child.
//	If B has a maximum constraint width of 30, then first 10 is added to
//	all children (A.Width=20, B.Width=30, C.Width=40). Then A and C
//	(20+40=60 and 30 pixel space left) will get 30/2=15 additional,
//	resulting in: A.Width=35, B.Width=30, C.Width=55 (35+30+55=120).
//	crsHomogenousSpaceResize
//	Enlarge space between children equally.
//	Children are resized to their normal/adviced size. If there is some space
//	left in the client area of the parent, then the space between the children
//	is expanded.
//	For example: 3 child controls A, B, C with A.Width=10, B.Width=20 and
//	C.Width=30 (total=60). If the Parent's client area has a ClientWidth of
//	120, then there will be 60/2=30 space between A and B and between
//	B and C.
//	crsSameSize - not implemented yet
//	Set each child to the same size (maybe one pixel difference).
//	The client area is divided by the number of controls and each control
//	gets the same size. The remainder is distributed to the first children.
type TChildControlResizeStyle = int32

const (
	CrsAnchorAligning TChildControlResizeStyle = iota
	CrsScaleChilds
	CrsHomogenousChildResize
	CrsHomogenousSpaceResize
	CrsSameSize
)

// TClipboardType ENUM
type TClipboardType = int32

const (
	CtPrimarySelection TClipboardType = iota
	CtSecondarySelection
	CtClipboard
)

// TCloseAction ENUM
type TCloseAction = int32

const (
	CaNone TCloseAction = iota
	CaHide
	CaFree
	CaMinimize
)

// TColumnButtonStyle ENUM
type TColumnButtonStyle = int32

const (
	CbsAuto TColumnButtonStyle = iota
	CbsEllipsis
	CbsNone
	CbsPickList
	CbsCheckboxColumn
	CbsButton
	CbsButtonColumn
)

// TColumnLayout ENUM
//
//	TCustomRadioGroup
type TColumnLayout = int32

const (
	ClHorizontalThenVertical TColumnLayout = iota
	ClVerticalThenHorizontal
)

// TComboBoxExStyle ENUM
type TComboBoxExStyle = int32

const (
	CsExDropDown TComboBoxExStyle = iota
	CsExSimple
	CsExDropDownList
)

// TComboBoxStyle ENUM
type TComboBoxStyle = int32

const (
	CsDropDown TComboBoxStyle = iota
	CsSimple
	CsDropDownList
	CsOwnerDrawFixed
	CsOwnerDrawVariable
	CsOwnerDrawEditableFixed
	CsOwnerDrawEditableVariable
)

// TControlCellAlign ENUM
type TControlCellAlign = int32

const (
	CcaFill TControlCellAlign = iota
	CcaLeftTop
	CcaRightBottom
	CcaCenter
)

// TControlChildrenLayout ENUM
type TControlChildrenLayout = int32

const (
	CclNone TControlChildrenLayout = iota
	CclLeftToRightThenTopToBottom
	CclTopToBottomThenLeftToRight
)

// TCoolBandMaximize ENUM
//
//	BandMaximize is not used now but is needed for Delphi compatibility.
//	It is not used in Delphi's TCoolBar either.
type TCoolBandMaximize = int32

const (
	BmNone TCoolBandMaximize = iota
	BmClick
	BmDblClick
)

// TCustomDrawStage ENUM
type TCustomDrawStage = int32

const (
	CdPrePaint TCustomDrawStage = iota
	CdPostPaint
	CdPreErase
	CdPostErase
)

// TDateDisplayOrder ENUM
//
//	Used by DateDisplayOrder property to determine the order to display date
//	parts -- d-m-y, m-d-y or y-m-d.
//	When ddoTryDefault is set, the actual order is determined from
//	ShortDateFormat global variable -- see comments above
//	AdjustEffectiveDateDisplayOrder procedure
type TDateDisplayOrder = int32

const (
	DdoDMY TDateDisplayOrder = iota
	DdoMDY
	DdoYMD
	DdoTryDefault
)

// TDateOrder ENUM
type TDateOrder = int32

const (
	DoNone TDateOrder = iota
	DoMDY
	DoDMY
	DoYMd
)

// TDateTimeKind ENUM
//
//	TDateTimeKind determines if we should display date, time or both:
type TDateTimeKind = int32

const (
	DtkDate TDateTimeKind = iota
	DtkTime
	DtkDateTime
)

// TDefaultColorType ENUM
type TDefaultColorType = int32

const (
	DctBrush TDefaultColorType = iota
	DctFont
)

// TDefaultMonitor ENUM
type TDefaultMonitor = int32

const (
	DmDesktop TDefaultMonitor = iota
	DmPrimary
	DmMainForm
	DmActiveForm
)

// TDialogKind ENUM
type TDialogKind = int32

const (
	DkOpen TDialogKind = iota
	DkSave
	DkPictureOpen
	DkPictureSave
)

// TDirection ENUM
type TDirection = int32

const (
	FromBeginning TDirection = iota
	FromEnd
)

// TDisplayCode ENUM
type TDisplayCode = int32

const (
	DrBounds TDisplayCode = iota
	DrIcon
	DrLabel
	DrSelectBounds
)

// TDockOrientation ENUM
type TDockOrientation = int32

const (
	DoNoOrient TDockOrientation = iota
	DoHorizontal
	DoVertical
	DoPages
)

// TDragKind ENUM
type TDragKind = int32

const (
	DkDrag TDragKind = iota
	DkDock
)

// TDragMode ENUM
type TDragMode = int32

const (
	DmManual TDragMode = iota
	DmAutomatic
)

// TDragState ENUM
type TDragState = int32

const (
	DsDragEnter TDragState = iota
	DsDragLeave
	DsDragMove
)

// TDrawImageMode ENUM
//
//	TXButton
type TDrawImageMode = int32

const (
	DimNormal TDrawImageMode = iota
	DimCenter
	DimStretch
)

// TDrawingStyle ENUM
//
//	TCustomImageList
//	@abstract(Contains a list of images)
//	Introduced by Marc Weustink <marc@dommelstein.net>
//	Delphis TCustomImageList is based on the Win32 imagelists which has
//	internally only one bitmap to hold all images. This reduces handle
//	allocation.
//	The original TCustomImageList implementation was LCL only based, so for
//	other platforms the single bitmap implementation had some speed drawbacks.
//	Therefore it was implemented as list of bitmaps, however it doesnt reduce
//	handle allocation.
//	In its current form, the imagelist is again based on a 32bit RGBA raw
//	imagedata and the widgetset is notified when images are added or removed,
//	so the widgetset can create its own optimal storage. The LCL keeps only the
//	data, so all transparency info will be stored cross platform. (not all
//	platforms have a 8bit alpha channel).
//	NOTE: due to its implementation, the TCustomImageList is not a TBitmap
//	collection. If a fast storage of bitmaps is needed, create your own list!
//	Some temp rework defines, for old functionality both need so be set
type TDrawingStyle = int32

const (
	DsFocus TDrawingStyle = iota
	DsSelected
	TDStyleDsNormal
	DsTransparent
)

// TDropMode ENUM
//
//	modes to determine drop position further
type TDropMode = int32

const (
	DmNowhere TDropMode = iota
	DmAbove
	DmOnNode
	DmBelow
)

// TDTCalAlignment ENUM
//
//	calendar alignment - left or right,
//	dtaDefault means it is determined by BiDiMode
type TDTCalAlignment = int32

const (
	DtaLeft TDTCalAlignment = iota
	DtaRight
	DtaDefault
)

// TDTDateMode ENUM
type TDTDateMode = int32

const (
	DmComboBox TDTDateMode = iota
	DmUpDown
	DmNone
)

// TDuplicates ENUM
type TDuplicates = int32

const (
	DupIgnore TDuplicates = iota
	DupAccept
	DupError
)

// TEchoMode ENUM
type TEchoMode = int32

const (
	EmNormal TEchoMode = iota
	EmNone
	EmPassword
)

// TEdgeStyle ENUM
type TEdgeStyle = int32

const (
	EsNone TEdgeStyle = iota
	EsRaised
	EsLowered
)

// TEditCharCase ENUM
//
//	TCustomEdit Options
type TEditCharCase = int32

const (
	EcNormal TEditCharCase = iota
	EcUppercase
	EcLowerCase
)

// TEditStyle ENUM
type TEditStyle = int32

const (
	EsSimple TEditStyle = iota
	EsEllipsis
	EsPickList
)

// TEmulatedTextHintStatus ENUM
//
//	Used for TCustomComboBox and TCustomEdit.
type TEmulatedTextHintStatus = int32

const (
	ThsHidden TEmulatedTextHintStatus = iota
	ThsShowing
	ThsChanging
)

// TFindItemKind ENUM
//
//	TMenu
type TFindItemKind = int32

const (
	FkCommand TFindItemKind = iota
	FkHandle
	FkShortCut
)

// TFlowStyle ENUM
type TFlowStyle = int32

const (
	FsLeftRightTopBottom TFlowStyle = iota
	FsRightLeftTopBottom
	FsLeftRightBottomTop
	FsRightLeftBottomTop
	FsTopBottomLeftRight
	FsBottomTopLeftRight
	FsTopBottomRightLeft
	FsBottomTopRightLeft
)

// TFontPitch ENUM
type TFontPitch = int32

const (
	FpDefault TFontPitch = iota
	FpVariable
	FpFixed
)

// TFontQuality ENUM
type TFontQuality = int32

const (
	FqDefault TFontQuality = iota
	FqDraft
	FqProof
	FqNonAntialiased
	FqAntialiased
	FqCleartype
	FqCleartypeNatural
)

// TFormBorderStyle ENUM
type TFormBorderStyle = int32

const (
	BsNone TFormBorderStyle = iota
	BsSingle
	BsSizeable
	BsDialog
	BsToolWindow
	BsSizeToolWin
)

// TFormStyle ENUM
type TFormStyle = int32

const (
	FsNormal TFormStyle = iota
	FsMDIChild
	FsMDIForm
	FsStayOnTop
	FsSplash
	FsSystemStayOnTop
)

// TFPBrushStyle ENUM
type TFPBrushStyle = int32

const (
	BsSolid TFPBrushStyle = iota
	BsClear
	BsHorizontal
	BsVertical
	BsFDiagonal
	BsBDiagonal
	BsCross
	BsDiagCross
	BsImage
	BsPattern
)

// TFPDrawingMode ENUM
type TFPDrawingMode = int32

const (
	DmOpaque TFPDrawingMode = iota
	DmAlphaBlend
	DmCustom
)

// TFPImgProgressStage ENUM
type TFPImgProgressStage = int32

const (
	PsStarting TFPImgProgressStage = iota
	PsRunning
	PsEnding
)

// TFPObservedOperation ENUM
//
//	Notification operations :
//	Observer has changed, is freed, item added to/deleted from list, custom event.
type TFPObservedOperation = int32

const (
	OoChange TFPObservedOperation = iota
	OoFree
	OoAddItem
	OoDeleteItem
	OoCustom
)

// TFPPenEndCap ENUM
type TFPPenEndCap = int32

const (
	PecRound TFPPenEndCap = iota
	PecSquare
	PecFlat
)

// TFPPenJoinStyle ENUM
type TFPPenJoinStyle = int32

const (
	PjsRound TFPPenJoinStyle = iota
	PjsBevel
	PjsMiter
)

// TFPPenMode ENUM
type TFPPenMode = int32

const (
	PmBlack TFPPenMode = iota
	PmWhite
	PmNop
	PmNot
	PmCopy
	PmNotCopy
	PmMergePenNot
	PmMaskPenNot
	PmMergeNotPen
	PmMaskNotPen
	PmMerge
	PmNotMerge
	PmMask
	PmNotMask
	PmXor
	PmNotXor
)

// TFPPenStyle ENUM
type TFPPenStyle = int32

const (
	PsSolid TFPPenStyle = iota
	PsDash
	PsDot
	PsDashDot
	PsDashDotDot
	PsinsideFrame
	PsPattern
	PsClear
)

// TGlyphShowMode ENUM
type TGlyphShowMode = int32

const (
	GsmAlways TGlyphShowMode = iota
	GsmNever
	GsmApplication
	GsmSystem
)

// TGrabStyle ENUM
//
//	TCoolBar
type TGrabStyle = int32

const (
	GsSimple TGrabStyle = iota
	GsDouble
	GsHorLines
	GsVerLines
	GsGripper
	GsButton
)

// TGradientDirection ENUM
type TGradientDirection = int32

const (
	GdVertical TGradientDirection = iota
	GdHorizontal
)

// TGraphicsBevelCut ENUM
type TGraphicsBevelCut = int32

const (
	BvNone TGraphicsBevelCut = iota
	BvLowered
	BvRaised
	BvSpace
)

// TGraphicsDrawEffect ENUM
type TGraphicsDrawEffect = int32

const (
	GdeNormal TGraphicsDrawEffect = iota
	GdeDisabled
	GdeHighlighted
	GdeShadowed
	Gde1Bit
)

// TGridCursorState ENUM
type TGridCursorState = int32

const (
	GcsDefault TGridCursorState = iota
	GcsColWidthChanging
	GcsRowHeightChanging
	GcsDragging
)

// THeaderSectionState ENUM
//
//	THeaderSection
type THeaderSectionState = int32

const (
	HsNormal THeaderSectionState = iota
	HsHot
	HsPressed
)

// THelpType ENUM
type THelpType = int32

const (
	HtKeyword THelpType = iota
	HtContext
)

// TIconArrangement ENUM
//
//	TIconOptions
type TIconArrangement = int32

const (
	IaTop TIconArrangement = iota
	IaLeft
)

// TImageOrientation ENUM
//
//	TImageButton
type TImageOrientation = int32

const (
	IoHorizontal TImageOrientation = iota
	IoVertical
)

// TImageType ENUM
type TImageType = int32

const (
	ItImage TImageType = iota
	ItMask
)

// TItemChange ENUM
//
//	TCustomListView
type TItemChange = int32

const (
	CtText TItemChange = iota
	CtImage
	CtState
)

// TItemEraseAction ENUM
//
//	Used to describe the action to do when using the OnBeforeItemErase event.
type TItemEraseAction = int32

const (
	EaColor TItemEraseAction = iota
	EaDefault
	EaNone
)

// TItemFind ENUM
type TItemFind = int32

const (
	IfData TItemFind = iota
	IfPartialString
	IfExactString
	IfNearest
)

// TLabelPosition ENUM
//
//	TCustomLabeledEdit
type TLabelPosition = int32

const (
	LpAbove TLabelPosition = iota
	LpBelow
	LpLeft
	LpRight
)

// TLayoutAdjustmentPolicy ENUM
type TLayoutAdjustmentPolicy = int32

const (
	LapDefault TLayoutAdjustmentPolicy = iota
	LapFixedLayout
	LapAutoAdjustWithoutHorizontalScrolling
	LapAutoAdjustForDPI
)

// TLazAccessibilityRole ENUM
type TLazAccessibilityRole = int32

const (
	LarIgnore TLazAccessibilityRole = iota
	LarAnimation
	LarButton
	LarCell
	LarChart
	LarCheckBox
	LarClock
	LarColorPicker
	LarColumn
	LarComboBox
	LarDateField
	LarGrid
	LarGroup
	LarImage
	LarLabel
	LarListBox
	LarListItem
	LarMenuBar
	LarMenuItem
	LarProgressIndicator
	LarRadioButton
	LarResizeGrip
	LarRow
	LarScrollBar
	LarSpinner
	LarTabControl
	LarText
	LarTextEditorMultiline
	LarTextEditorSingleline
	LarToolBar
	LarToolBarButton
	LarTrackBar
	LarTreeView
	LarTreeItem
	LarUnknown
	LarWindow
)

// TLazDockHeaderPart ENUM
type TLazDockHeaderPart = int32

const (
	LdhpAll TLazDockHeaderPart = iota
	LdhpCaption
	LdhpRestoreButton
	LdhpCloseButton
)

// TLazSynBorderSide ENUM
type TLazSynBorderSide = int32

const (
	BsLeft TLazSynBorderSide = iota
	BsTop
	BsRight
	BsBottom
)

// TListAssignOp ENUM
type TListAssignOp = int32

const (
	LaCopy TListAssignOp = iota
	LaAnd
	LaOr
	LaXor
	LaSrcUnique
	LaDestUnique
)

// TListBoxStyle ENUM
//
//	TCustomListBox
type TListBoxStyle = int32

const (
	LbStandard TListBoxStyle = iota
	LbOwnerDrawFixed
	LbOwnerDrawVariable
	LbVirtual
)

// TMaskEditValidationErrorMode ENUM
type TMaskEditValidationErrorMode = int32

const (
	MvemException TMaskEditValidationErrorMode = iota
	MvemEvent
)

// TMenuItemAutoFlag ENUM
type TMenuItemAutoFlag = int32

const (
	MaAutomatic TMenuItemAutoFlag = iota
	MaManual
	MaParent
)

// TMenuItemHandlerType ENUM
//
//	TMenuItem
type TMenuItemHandlerType = int32

const (
	MihtDestroy TMenuItemHandlerType = iota
)

// TMissingNameValueSeparatorAction ENUM
type TMissingNameValueSeparatorAction = int32

const (
	MnvaValue TMissingNameValueSeparatorAction = iota
	MnvaName
	MnvaEmpty
	MnvaError
)

// TMonitorDefaultTo ENUM
type TMonitorDefaultTo = int32

const (
	MdNearest TMonitorDefaultTo = iota
	MdNull
	MdPrimary
)

// TMonthDisplay ENUM
type TMonthDisplay = int32

const (
	MdShort TMonthDisplay = iota
	MdLong
	MdCustom
)

// TMouseWheelOption ENUM
type TMouseWheelOption = int32

const (
	MwCursor TMouseWheelOption = iota
	MwGrid
)

// TNodeAttachMode ENUM
type TNodeAttachMode = int32

const (
	NaAdd TNodeAttachMode = iota
	NaAddFirst
	NaAddChild
	NaAddChildFirst
	NaInsert
	NaInsertBehind
)

// TOperation ENUM
//
//	TComponent class
type TOperation = int32

const (
	OpInsert TOperation = iota
	OpRemove
)

// TPageMeasureUnits ENUM
type TPageMeasureUnits = int32

const (
	PmDefault TPageMeasureUnits = iota
	PmMillimeters
	PmInches
)

// TPairSplitterType ENUM
//
//	TCustomPairSplitter
type TPairSplitterType = int32

const (
	PstHorizontal TPairSplitterType = iota
	PstVertical
)

// TPanelButtonEx ENUM
type TPanelButtonEx = int32

const (
	PbOK TPanelButtonEx = iota
	PbCancel
	PbClose
	PbHelp
	PbNone
)

// TPenEndCap ENUM
type TPenEndCap = int32

const (
	TPECapPecRound TPenEndCap = iota
	TPECapPecSquare
	TPECapPecFlat
)

// TPenJoinStyle ENUM
type TPenJoinStyle = int32

const (
	TPJStylePjsRound TPenJoinStyle = iota
	TPJStylePjsBevel
	TPJStylePjsMiter
)

// TPixelFormat ENUM
//
//	For Delphi compatibility
type TPixelFormat = int32

const (
	PfDevice TPixelFormat = iota
	Pf1bit
	Pf4bit
	Pf8bit
	Pf15bit
	Pf16bit
	Pf24bit
	Pf32bit
	PfCustom
)

// TPopupAlignment ENUM
//
//	TPopupMenu
type TPopupAlignment = int32

const (
	PaLeft TPopupAlignment = iota
	PaRight
	PaCenter
)

// TPopupMode ENUM
type TPopupMode = int32

const (
	PmNone TPopupMode = iota
	PmAuto
	PmExplicit
)

// TPosition ENUM
//
//	form position policies:
type TPosition = int32

const (
	PoDesigned TPosition = iota
	PoDefault
	PoDefaultPosOnly
	PoDefaultSizeOnly
	PoScreenCenter
	PoDesktopCenter
	PoMainFormCenter
	PoOwnerFormCenter
	PoWorkAreaCenter
)

// TPrefixOption ENUM
type TPrefixOption = int32

const (
	PoNone TPrefixOption = iota
	PoHeaderClick
)

// TPrinterOrientation ENUM
type TPrinterOrientation = int32

const (
	PoPortrait TPrinterOrientation = iota
	PoLandscape
	PoReverseLandscape
	PoReversePortrait
)

// TPrinterState ENUM
type TPrinterState = int32

const (
	PsNoDefine TPrinterState = iota
	PsReady
	PsPrinting
	PsStopped
)

// TPrinterType ENUM
type TPrinterType = int32

const (
	PtLocal TPrinterType = iota
	PtNetWork
)

// TPrintRange ENUM
//
//	TPrintDialog
type TPrintRange = int32

const (
	PrAllPages TPrintRange = iota
	PrSelection
	PrPageNums
	PrCurrentPage
)

// TProgressBarOrientation ENUM
type TProgressBarOrientation = int32

const (
	PbHorizontal TProgressBarOrientation = iota
	PbVertical
	PbRightToLeft
	PbTopDown
)

// TProgressBarState ENUM
//
//	used by TTaskDialogProgressBar, but Delphi defines it in ComCtrls unit
type TProgressBarState = int32

const (
	PbsNormal TProgressBarState = iota
	PbsError
	PbsPaused
)

// TProgressBarStyle ENUM
type TProgressBarStyle = int32

const (
	PbstNormal TProgressBarStyle = iota
	PbstMarquee
)

// TRangeSelectMode ENUM
//
//	Option goRangeSelect: --> select a single range only, or multiple ranges
type TRangeSelectMode = int32

const (
	RsmSingle TRangeSelectMode = iota
	RsmMulti
)

// TRawImageBitOrder ENUM
type TRawImageBitOrder = int32

const (
	RiboBitsInOrder TRawImageBitOrder = iota
	RiboReversedBits
)

// TRawImageByteOrder ENUM
type TRawImageByteOrder = int32

const (
	RiboLSBFirst TRawImageByteOrder = iota
	RiboMSBFirst
)

// TRawImageColorFormat ENUM
//
//	------------------------------------------------------------------------------
//	raw image data
//	Colorformat: Higher values means higher intensity.
//	For example: Red=0 means no red, Alpha=0 means transparent
type TRawImageColorFormat = int32

const (
	RicfNone TRawImageColorFormat = iota
	RicfRGBA
	RicfGray
)

// TRawImageLineEnd ENUM
type TRawImageLineEnd = int32

const (
	RileTight TRawImageLineEnd = iota
	RileByteBoundary
	RileWordBoundary
	RileDWordBoundary
	RileQWordBoundary
	RileDQWordBoundary
)

// TRawImageLineOrder ENUM
type TRawImageLineOrder = int32

const (
	RiloTopToBottom TRawImageLineOrder = iota
	RiloBottomToTop
)

// TRegDataType ENUM
type TRegDataType = int32

const (
	RdUnknown TRegDataType = iota
	RdString
	RdExpandString
	RdBinary
	RdInteger
	RdIntegerBigEndian
	RdLink
	RdMultiString
	RdResourceList
	RdFullResourceDescriptor
	RdResourceRequirementList
	RdInt64
)

// TReplacedChildSite ENUM
type TReplacedChildSite = int32

const (
	RplcLeft TReplacedChildSite = iota
	RplcRight
)

// TResizeStyle ENUM
//
//	TCustomSplitter
type TResizeStyle = int32

const (
	RsLine TResizeStyle = iota
	RsNone
	RsPattern
	RsUpdate
)

// TScrollBarKind ENUM
//
//	TControlScrollBar
type TScrollBarKind = int32

const (
	SbHorizontal TScrollBarKind = iota
	SbVertical
)

// TScrollCode ENUM
type TScrollCode = int32

const (
	ScLineUp TScrollCode = iota
	ScLineDown
	ScPageUp
	ScPageDown
	ScPosition
	ScTrack
	ScTop
	ScBottom
	ScEndScroll
)

// TScrollStyle ENUM
//
//	TScrollBar
type TScrollStyle = int32

const (
	SsNone TScrollStyle = iota
	SsHorizontal
	SsVertical
	SsBoth
	SsAutoHorizontal
	SsAutoVertical
	SsAutoBoth
)

// TSearchDirection ENUM
type TSearchDirection = int32

const (
	SdLeft TSearchDirection = iota
	SdRight
	SdAbove
	SdBelow
	SdAll
)

// TSectionTrackState ENUM
type TSectionTrackState = int32

const (
	TsTrackBegin TSectionTrackState = iota
	TsTrackMove
	TsTrackEnd
)

// TSeekOrigin ENUM
type TSeekOrigin = int32

const (
	SoBeginning TSeekOrigin = iota
	SoCurrent
	SoEnd
)

// TShapeType ENUM
//
//	TShape
type TShapeType = int32

const (
	StRectangle TShapeType = iota
	StSquare
	StRoundRect
	StRoundSquare
	StEllipse
	StCircle
	StSquaredDiamond
	StDiamond
	StTriangle
	StTriangleLeft
	StTriangleRight
	StTriangleDown
	StStar
	StStarDown
	StPolygon
)

// TShowInTaskBar ENUM
type TShowInTaskBar = int32

const (
	StDefault TShowInTaskBar = iota
	StAlways
	StNever
)

// TSmartAutoFitType ENUM
type TSmartAutoFitType = int32

const (
	SmaAllColumns TSmartAutoFitType = iota
	SmaNoColumn
	SmaUseColumnOption
)

// TSortIndicator ENUM
type TSortIndicator = int32

const (
	SiNone TSortIndicator = iota
	SiAscending
	SiDescending
)

// TSortOrder ENUM
type TSortOrder = int32

const (
	SoAscending TSortOrder = iota
	SoDescending
)

// TSortType ENUM
type TSortType = int32

const (
	StNone TSortType = iota
	StData
	StText
	StBoth
)

// TSQLDialect ENUM
type TSQLDialect = int32

const (
	SqlStandard TSQLDialect = iota
	SqlInterbase6
	SqlMSSQL7
	SqlMySQL
	SqlOracle
	SqlSybase
	SqlIngres
	SqlMSSQL2K
	SqlPostgres
	SqlSQLite
	SqlFirebird25
	SqlFirebird30
	SqlFirebird40
)

// TStaticBorderStyle ENUM
//
//	TCustomStaticText
type TStaticBorderStyle = int32

const (
	SbsNone TStaticBorderStyle = iota
	SbsSingle
	SbsSunken
)

// TStatusPanelBevel ENUM
type TStatusPanelBevel = int32

const (
	TSPBevelPbNone TStatusPanelBevel = iota
	PbLowered
	PbRaised
)

// TStatusPanelStyle ENUM
type TStatusPanelStyle = int32

const (
	PsText TStatusPanelStyle = iota
	PsOwnerDraw
)

// TStreamOwnership ENUM
//
//	TStreamAdapter
type TStreamOwnership = int32

const (
	SoReference TStreamOwnership = iota
	SoOwned
)

// TStringDelim ENUM
type TStringDelim = int32

const (
	SdSingleQuote TStringDelim = iota
	SdDoubleQuote
)

// TStringsSortStyle ENUM
type TStringsSortStyle = int32

const (
	SslNone TStringsSortStyle = iota
	SslUser
	SslAuto
)

// TSynCaretAdjustMode ENUM
type TSynCaretAdjustMode = int32

const (
	ScamIgnore TSynCaretAdjustMode = iota
	ScamAdjust
	ScamForceAdjust
	ScamEnd
	ScamBegin
)

// TSynCaretType ENUM
type TSynCaretType = int32

const (
	CtVerticalLine TSynCaretType = iota
	CtHorizontalLine
	CtHalfBlock
	CtBlock
	CtCostum
)

// TSynCompletionLongHintType ENUM
type TSynCompletionLongHintType = int32

const (
	SclpNone TSynCompletionLongHintType = iota
	SclpExtendRightOnly
	SclpExtendHalfLeft
	SclpExtendUnlimitedLeft
)

// TSynCopyPasteAction ENUM
type TSynCopyPasteAction = int32

const (
	ScaContinue TSynCopyPasteAction = iota
	ScaPlainText
	ScaAbort
)

// TSynDefaultPopupMenu ENUM
type TSynDefaultPopupMenu = int32

const (
	DpmDisabled TSynDefaultPopupMenu = iota
	DpmBefore
	DpmAfter
)

// TSynEditBracketHighlightStyle ENUM
type TSynEditBracketHighlightStyle = int32

const (
	SbhsLeftOfCursor TSynEditBracketHighlightStyle = iota
	SbhsRightOfCursor
	SbhsBoth
)

// TSynEditMarkSortOrder ENUM
type TSynEditMarkSortOrder = int32

const (
	SmsoUnsorted TSynEditMarkSortOrder = iota
	SmsoColumn
	SmsoPriority
	SmsoBookmarkFirst
	SmsoBookMarkLast
)

// TSynEditNotifyReason ENUM
type TSynEditNotifyReason = int32

const (
	SenrLineCount TSynEditNotifyReason = iota
	SenrLineChange
	SenrLinesModified
	SenrHighlightChanged
	SenrLineMappingChanged
	SenrEditAction
	SenrCleared
	SenrUndoRedoAdded
	SenrModifiedChanged
	SenrIncOwnedPaintLock
	SenrDecOwnedPaintLock
	SenrIncPaintLock
	SenrDecPaintLock
	SenrBeforeIncPaintLock
	SenrAfterIncPaintLock
	SenrBeforeDecPaintLock
	SenrAfterDecPaintLock
	SenrTextBufferChanging
	SenrTextBufferChanged
	SenrBeginUndoRedo
	SenrEndUndoRedo
)

// TSynEditStringTrimmingType ENUM
type TSynEditStringTrimmingType = int32

const (
	SettLeaveLine TSynEditStringTrimmingType = iota
	SettEditLine
	SettMoveCaret
	SettIgnoreAll
)

// TSynEventParamType ENUM
type TSynEventParamType = int32

const (
	PtString TSynEventParamType = iota
	PtInteger
)

// TSynFrameEdges ENUM
type TSynFrameEdges = int32

const (
	SfeNone TSynFrameEdges = iota
	SfeAround
	SfeBottom
	SfeLeft
)

// TSynGutterSide ENUM
//
//	TSynGutterBase
type TSynGutterSide = int32

const (
	GsLeft TSynGutterSide = iota
	GsRight
)

// TSynHTMLSynMode ENUM
type TSynHTMLSynMode = int32

const (
	ShmHtml TSynHTMLSynMode = iota
	ShmXHtml
)

// TSynLineState ENUM
type TSynLineState = int32

const (
	SlsNone TSynLineState = iota
	SlsSaved
	SlsUnsaved
)

// TSynLineStyle ENUM
type TSynLineStyle = int32

const (
	SlsSolid TSynLineStyle = iota
	SlsDashed
	SlsDotted
	SlsWaved
)

// TSynLogCharSide ENUM
type TSynLogCharSide = int32

const (
	CslBefore TSynLogCharSide = iota
	CslAfter
	CslFollowLtr
	CslFollowRtl
)

// TSynMAClickCount ENUM
type TSynMAClickCount = int32

const (
	CcSingle TSynMAClickCount = iota
	CcDouble
	CcTriple
	CcQuad
	CcAny
)

// TSynMAClickDir ENUM
type TSynMAClickDir = int32

const (
	CdUp TSynMAClickDir = iota
	CdDown
)

// TSynMacroState ENUM
type TSynMacroState = int32

const (
	MsStopped TSynMacroState = iota
	MsRecording
	MsPlaying
	MsPaused
)

// TSynMarksAdjustMode ENUM
//   - This is used, if text is *replaced*.
//     What to do with marks in text that is deleted/replaced
type TSynMarksAdjustMode = int32

const (
	SmaMoveUp TSynMarksAdjustMode = iota
	SmaKeep
)

// TSynMouseButton ENUM
//   - For streaming compatibility the enum members of TSynMouseButton must have the
//     same names as Controls.TMouseButton
//     To avoid conflicts the definiton will be hidden here and aliases be defind for
//     common usage
type TSynMouseButton = int32

const (
	TSMButtonMbLeft TSynMouseButton = iota
	TSMButtonMbRight
	TSMButtonMbMiddle
	TSMButtonMbExtra1
	TSMButtonMbExtra2
	MbWheelUp
	MbWheelDown
	MbWheelLeft
	MbWheelRight
)

// TSynPhysCharSide ENUM
type TSynPhysCharSide = int32

const (
	CspLeft TSynPhysCharSide = iota
	CspRight
	CspFollowLtr
	CspFollowRtl
)

// TSynReplaceAction ENUM
type TSynReplaceAction = int32

const (
	RaCancel TSynReplaceAction = iota
	RaSkip
	RaReplace
	RaReplaceAll
)

// TSynSelectionMode ENUM
type TSynSelectionMode = int32

const (
	SmNormal TSynSelectionMode = iota
	SmLine
	SmColumn
	SmCurrent
)

// TSynSizedDiffAVLFindMode ENUM
type TSynSizedDiffAVLFindMode = int32

const (
	AfmNil TSynSizedDiffAVLFindMode = iota
	AfmCreate
	AfmPrev
	AfmNext
)

// TSysLinkType ENUM
type TSysLinkType = int32

const (
	SltURL TSysLinkType = iota
	SltID
)

// TTabPosition ENUM
type TTabPosition = int32

const (
	TpTop TTabPosition = iota
	TpBottom
	TpLeft
	TpRight
)

// TTabStyle ENUM
type TTabStyle = int32

const (
	TsTabs TTabStyle = iota
	TsButtons
	TsFlatButtons
)

// TTaskBarBehavior ENUM
type TTaskBarBehavior = int32

const (
	TbDefault TTaskBarBehavior = iota
	TbMultiButton
	TbSingleButton
)

// TTextLayout ENUM
//
//	Reflects text style when drawn in a rectangle
type TTextLayout = int32

const (
	TlTop TTextLayout = iota
	TlCenter
	TlBottom
)

// TTextLineBreakStyle ENUM
type TTextLineBreakStyle = int32

const (
	TlbsLF TTextLineBreakStyle = iota
	TlbsCRLF
	TlbsCR
)

// TThemedButton ENUM
//
//	'Button' theme data
type TThemedButton = int32

const (
	TbButtonDontCare TThemedButton = iota
	TbButtonRoot
	TbPushButtonNormal
	TbPushButtonHot
	TbPushButtonPressed
	TbPushButtonDisabled
	TbPushButtonDefaulted
	TbRadioButtonUncheckedNormal
	TbRadioButtonUncheckedHot
	TbRadioButtonUncheckedPressed
	TbRadioButtonUncheckedDisabled
	TbRadioButtonCheckedNormal
	TbRadioButtonCheckedHot
	TbRadioButtonCheckedPressed
	TbRadioButtonCheckedDisabled
	TbCheckBoxUncheckedNormal
	TbCheckBoxUncheckedHot
	TbCheckBoxUncheckedPressed
	TbCheckBoxUncheckedDisabled
	TbCheckBoxCheckedNormal
	TbCheckBoxCheckedHot
	TbCheckBoxCheckedPressed
	TbCheckBoxCheckedDisabled
	TbCheckBoxMixedNormal
	TbCheckBoxMixedHot
	TbCheckBoxMixedPressed
	TbCheckBoxMixedDisabled
	TbGroupBoxNormal
	TbGroupBoxDisabled
	TbUserButton
)

// TThemedClock ENUM
//
//	'Clock' theme data
type TThemedClock = int32

const (
	TcClockDontCare TThemedClock = iota
	TcClockRoot
	TcTimeNormal
)

// TThemedComboBox ENUM
//
//	'ComboBox' theme data
type TThemedComboBox = int32

const (
	TcComboBoxDontCare TThemedComboBox = iota
	TcComboBoxRoot
	TcDropDownButtonNormal
	TcDropDownButtonHot
	TcDropDownButtonPressed
	TcDropDownButtonDisabled
)

// TThemedEdit ENUM
//
//	'Edit' theme data
type TThemedEdit = int32

const (
	TeEditDontCare TThemedEdit = iota
	TeEditRoot
	TeEditTextNormal
	TeEditTextHot
	TeEditTextSelected
	TeEditTextDisabled
	TeEditTextFocused
	TeEditTextReadOnly
	TeEditTextAssist
	TeEditCaret
)

// TThemedElement ENUM
//
//	These are all elements which can be themed.
type TThemedElement = int32

const (
	TeButton TThemedElement = iota
	TeClock
	TeComboBox
	TeEdit
	TeExplorerBar
	TeHeader
	TeListView
	TeMenu
	TePage
	TeProgress
	TeRebar
	TeScrollBar
	TeSpin
	TeStartPanel
	TeStatus
	TeTab
	TeTaskBand
	TeTaskBar
	TeToolBar
	TeToolTip
	TeTrackBar
	TeTrayNotify
	TeTreeview
	TeWindow
)

// TThemedExplorerBar ENUM
//
//	'ExplorerBar' theme data
type TThemedExplorerBar = int32

const (
	TebExplorerBarDontCare TThemedExplorerBar = iota
	TebExplorerBarRoot
	TebHeaderBackgroundNormal
	TebHeaderBackgroundHot
	TebHeaderBackgroundPressed
	TebHeaderCloseNormal
	TebHeaderCloseHot
	TebHeaderClosePressed
	TebHeaderPinNormal
	TebHeaderPinHot
	TebHeaderPinPressed
	TebHeaderPinSelectedNormal
	TebHeaderPinSelectedHot
	TebHeaderPinSelectedPressed
	TebIEBarMenuNormal
	TebIEBarMenuHot
	TebIEBarMenuPressed
	TebNormalGroupBackground
	TebNormalGroupCollapseNormal
	TebNormalGroupCollapseHot
	TebNormalGroupCollapsePressed
	TebNormalGroupExpandNormal
	TebNormalGroupExpandHot
	TebNormalGroupExpandPressed
	TebNormalGroupHead
	TebSpecialGroupBackground
	TebSpecialGroupCollapseSpecial
	TebSpecialGroupCollapseHot
	TebSpecialGroupCollapsePressed
	TebSpecialGroupExpandSpecial
	TebSpecialGroupExpandHot
	TebSpecialGroupExpandPressed
	TebSpecialGroupHead
)

// TThemedHeader ENUM
//
//	'Header' theme data
type TThemedHeader = int32

const (
	ThHeaderDontCare TThemedHeader = iota
	ThHeaderRoot
	ThHeaderItemNormal
	ThHeaderItemHot
	ThHeaderItemPressed
	ThHeaderItemLeftNormal
	ThHeaderItemLeftHot
	ThHeaderItemLeftPressed
	ThHeaderItemRightNormal
	ThHeaderItemRightHot
	ThHeaderItemRightPressed
	ThHeaderSortArrowSortedUp
	ThHeaderSortArrowSortedDown
)

// TThemedListView ENUM
//
//	'ListView' theme data
type TThemedListView = int32

const (
	TlListviewDontCare TThemedListView = iota
	TlListviewRoot
	TlListItemNormal
	TlListItemHot
	TlListItemSelected
	TlListItemDisabled
	TlListItemSelectedNotFocus
	TlListGroup
	TlListDetail
	TlListSortDetail
	TlEmptyText
)

// TThemedMenu ENUM
//
//	'Menu' theme data
type TThemedMenu = int32

const (
	TmMenuDontCare TThemedMenu = iota
	TmMenuRoot
	TmMenuItemNormal
	TmMenuItemSelected
	TmMenuItemDemoted
	TmMenuDropDown
	TmMenuBarItem
	TmMenuBarDropDown
	TmChevron
	TmSeparator
	TmBarBackgroundActive
	TmBarBackgroundInactive
	TmBarItemNormal
	TmBarItemHot
	TmBarItemPushed
	TmBarItemDisabled
	TmBarItemDisabledHot
	TmBarItemDisabledPushed
	TmPopupBackground
	TmPopupBorders
	TmPopupCheckMarkNormal
	TmPopupCheckMarkDisabled
	TmPopupBulletNormal
	TmPopupBulletDisabled
	TmPopupCheckBackgroundDisabled
	TmPopupCheckBackgroundNormal
	TmPopupCheckBackgroundBitmap
	TmPopupGutter
	TmPopupItemNormal
	TmPopupItemHot
	TmPopupItemDisabled
	TmPopupItemDisabledHot
	TmPopupSeparator
	TmPopupSubmenuNormal
	TmPopupSubmenuDisabled
	TmSystemCloseNormal
	TmSystemCloseDisabled
	TmSystemMaximizeNormal
	TmSystemMaximizeDisabled
	TmSystemMinimizeNormal
	TmSystemMinimizeDisabled
	TmSystemRestoreNormal
	TmSystemRestoreDisabled
)

// TThemedPage ENUM
//
//	'Page' theme data
type TThemedPage = int32

const (
	TpPageDontCare TThemedPage = iota
	TpPageRoot
	TpUpNormal
	TpUpHot
	TpUpPressed
	TpUpDisabled
	TpDownNormal
	TpDownHot
	TpDownPressed
	TpDownDisabled
	TpUpHorzNormal
	TpUpHorzHot
	TpUpHorzPressed
	TpUpHorzDisabled
	TpDownHorzNormal
	TpDownHorzHot
	TpDownHorzPressed
	TpDownHorzDisabled
)

// TThemedProgress ENUM
//
//	'Progress' theme data
type TThemedProgress = int32

const (
	TpProgressDontCare TThemedProgress = iota
	TpProgressRoot
	TpBar
	TpBarVert
	TpChunk
	TpChunkVert
)

// TThemedRebar ENUM
//
//	'Rebar' theme data
type TThemedRebar = int32

const (
	TrRebarDontCare TThemedRebar = iota
	TrRebarRoot
	TrGripper
	TrGripperVert
	TrBandNormal
	TrBandHot
	TrBandPressed
	TrBandDisabled
	TrBandChecked
	TrBandHotChecked
	TrChevronNormal
	TrChevronHot
	TrChevronPressed
	TrChevronDisabled
	TrChevronVertNormal
	TrChevronVertHot
	TrChevronVertPressed
	TrChevronVertDisabled
)

// TThemedScrollBar ENUM
//
//	'ScrollBar' theme data
type TThemedScrollBar = int32

const (
	TsScrollBarDontCare TThemedScrollBar = iota
	TsScrollBarRoot
	TsArrowBtnUpNormal
	TsArrowBtnUpHot
	TsArrowBtnUpPressed
	TsArrowBtnUpDisabled
	TsArrowBtnDownNormal
	TsArrowBtnDownHot
	TsArrowBtnDownPressed
	TsArrowBtnDownDisabled
	TsArrowBtnLeftNormal
	TsArrowBtnLeftHot
	TsArrowBtnLeftPressed
	TsArrowBtnLeftDisabled
	TsArrowBtnRightNormal
	TsArrowBtnRightHot
	TsArrowBtnRightPressed
	TsArrowBtnRightDisabled
	TsThumbBtnHorzNormal
	TsThumbBtnHorzHot
	TsThumbBtnHorzPressed
	TsThumbBtnHorzDisabled
	TsThumbBtnVertNormal
	TsThumbBtnVertHot
	TsThumbBtnVertPressed
	TsThumbBtnVertDisabled
	TsLowerTrackHorzNormal
	TsLowerTrackHorzHot
	TsLowerTrackHorzPressed
	TsLowerTrackHorzDisabled
	TsUpperTrackHorzNormal
	TsUpperTrackHorzHot
	TsUpperTrackHorzPressed
	TsUpperTrackHorzDisabled
	TsLowerTrackVertNormal
	TsLowerTrackVertHot
	TsLowerTrackVertPressed
	TsLowerTrackVertDisabled
	TsUpperTrackVertNormal
	TsUpperTrackVertHot
	TsUpperTrackVertPressed
	TsUpperTrackVertDisabled
	TsGripperHorzNormal
	TsGripperHorzHot
	TsGripperHorzPressed
	TsGripperHorzDisabled
	TsGripperVertNormal
	TsGripperVertHot
	TsGripperVertPressed
	TsGripperVertDisabled
	TsSizeBoxRightAlign
	TsSizeBoxLeftAlign
)

// TThemedSpin ENUM
//
//	'Spin' theme data
type TThemedSpin = int32

const (
	TsSpinDontCare TThemedSpin = iota
	TsSpinRoot
	TsUpNormal
	TsUpHot
	TsUpPressed
	TsUpDisabled
	TsDownNormal
	TsDownHot
	TsDownPressed
	TsDownDisabled
	TsUpHorzNormal
	TsUpHorzHot
	TsUpHorzPressed
	TsUpHorzDisabled
	TsDownHorzNormal
	TsDownHorzHot
	TsDownHorzPressed
	TsDownHorzDisabled
)

// TThemedStartPanel ENUM
//
//	'StartPanel' theme data
type TThemedStartPanel = int32

const (
	TspStartPanelDontCare TThemedStartPanel = iota
	TspStartPanelRoot
	TspUserPane
	TspMorePrograms
	TspMoreProgramsArrowNormal
	TspMoreProgramsArrowHot
	TspMoreProgramsArrowPressed
	TspProgList
	TspProgListSeparator
	TspPlacesList
	TspPlacesListSeparator
	TspLogOff
	TspLogOffButtonsNormal
	TspLogOffButtonsHot
	TspLogOffButtonsPressed
	TspUserPicture
	TspPreview
)

// TThemedStatus ENUM
//
//	'Status' theme data
type TThemedStatus = int32

const (
	TsStatusDontCare TThemedStatus = iota
	TsStatusRoot
	TsPane
	TsGripperPane
	TsGripper
)

// TThemedTab ENUM
//
//	'Tab' theme data
type TThemedTab = int32

const (
	TtTabDontCare TThemedTab = iota
	TtTabRoot
	TtTabItemNormal
	TtTabItemHot
	TtTabItemSelected
	TtTabItemDisabled
	TtTabItemFocused
	TtTabItemLeftEdgeNormal
	TtTabItemLeftEdgeHot
	TtTabItemLeftEdgeSelected
	TtTabItemLeftEdgeDisabled
	TtTabItemLeftEdgeFocused
	TtTabItemRightEdgeNormal
	TtTabItemRightEdgeHot
	TtTabItemRightEdgeSelected
	TtTabItemRightEdgeDisabled
	TtTabItemRightEdgeFocused
	TtTabItemBothEdgeNormal
	TtTabItemBothEdgeHot
	TtTabItemBothEdgeSelected
	TtTabItemBothEdgeDisabled
	TtTabItemBothEdgeFocused
	TtTopTabItemNormal
	TtTopTabItemHot
	TtTopTabItemSelected
	TtTopTabItemDisabled
	TtTopTabItemFocused
	TtTopTabItemLeftEdgeNormal
	TtTopTabItemLeftEdgeHot
	TtTopTabItemLeftEdgeSelected
	TtTopTabItemLeftEdgeDisabled
	TtTopTabItemLeftEdgeFocused
	TtTopTabItemRightEdgeNormal
	TtTopTabItemRightEdgeHot
	TtTopTabItemRightEdgeSelected
	TtTopTabItemRightEdgeDisabled
	TtTopTabItemRightEdgeFocused
	TtTopTabItemBothEdgeNormal
	TtTopTabItemBothEdgeHot
	TtTopTabItemBothEdgeSelected
	TtTopTabItemBothEdgeDisabled
	TtTopTabItemBothEdgeFocused
	TtPane
	TtBody
)

// TThemedTaskBand ENUM
//
//	'TaskBand' theme data
type TThemedTaskBand = int32

const (
	TtbTaskBandDontCare TThemedTaskBand = iota
	TtbTaskBandRoot
	TtbGroupCount
	TtbFlashButton
	TtpFlashButtonGroupMenu
)

// TThemedTaskBar ENUM
//
//	'TaskBar' theme data
type TThemedTaskBar = int32

const (
	TtTaskBarDontCare TThemedTaskBar = iota
	TtTaskBarRoot
	TtbTimeNormal
)

// TThemedToolBar ENUM
//
//	'ToolBar' theme data
type TThemedToolBar = int32

const (
	TtbToolBarDontCare TThemedToolBar = iota
	TtbToolBarRoot
	TtbButtonNormal
	TtbButtonHot
	TtbButtonPressed
	TtbButtonDisabled
	TtbButtonChecked
	TtbButtonCheckedHot
	TtbDropDownButtonNormal
	TtbDropDownButtonHot
	TtbDropDownButtonPressed
	TtbDropDownButtonDisabled
	TtbDropDownButtonChecked
	TtbDropDownButtonCheckedHot
	TtbSplitButtonNormal
	TtbSplitButtonHot
	TtbSplitButtonPressed
	TtbSplitButtonDisabled
	TtbSplitButtonChecked
	TtbSplitButtonCheckedHot
	TtbSplitButtonDropDownNormal
	TtbSplitButtonDropDownHot
	TtbSplitButtonDropDownPressed
	TtbSplitButtonDropDownDisabled
	TtbSplitButtonDropDownChecked
	TtbSplitButtonDropDownCheckedHot
	TtbSeparatorNormal
	TtbSeparatorHot
	TtbSeparatorPressed
	TtbSeparatorDisabled
	TtbSeparatorChecked
	TtbSeparatorCheckedHot
	TtbSeparatorVertNormal
	TtbSeparatorVertHot
	TtbSeparatorVertPressed
	TtbSeparatorVertDisabled
	TtbSeparatorVertChecked
	TtbSeparatorVertCheckedHot
)

// TThemedToolTip ENUM
//
//	'ToolTip' theme data
type TThemedToolTip = int32

const (
	TttToolTipDontCare TThemedToolTip = iota
	TttToolTipRoot
	TttStandardNormal
	TttStandardLink
	TttStandardTitleNormal
	TttStandardTitleLink
	TttBaloonNormal
	TttBaloonLink
	TttBaloonTitleNormal
	TttBaloonTitleLink
	TttCloseNormal
	TttCloseHot
	TttClosePressed
)

// TThemedTrackBar ENUM
//
//	'TrackBar' theme data
type TThemedTrackBar = int32

const (
	TtbTrackBarDontCare TThemedTrackBar = iota
	TtbTrackBarRoot
	TtbTrack
	TtbTrackVert
	TtbThumbNormal
	TtbThumbHot
	TtbThumbPressed
	TtbThumbFocused
	TtbThumbDisabled
	TtbThumbBottomNormal
	TtbThumbBottomHot
	TtbThumbBottomPressed
	TtbThumbBottomFocused
	TtbThumbBottomDisabled
	TtbThumbTopNormal
	TtbThumbTopHot
	TtbThumbTopPressed
	TtbThumbTopFocused
	TtbThumbTopDisabled
	TtbThumbVertNormal
	TtbThumbVertHot
	TtbThumbVertPressed
	TtbThumbVertFocused
	TtbThumbVertDisabled
	TtbThumbLeftNormal
	TtbThumbLeftHot
	TtbThumbLeftPressed
	TtbThumbLeftFocused
	TtbThumbLeftDisabled
	TtbThumbRightNormal
	TtbThumbRightHot
	TtbThumbRightPressed
	TtbThumbRightFocused
	TtbThumbRightDisabled
	TtbThumbTics
	TtbThumbTicsVert
)

// TThemedTrayNotify ENUM
//
//	'TrayNotify' theme data
type TThemedTrayNotify = int32

const (
	TtnTrayNotifyDontCare TThemedTrayNotify = iota
	TtnTrayNotifyRoot
	TtnBackground
	TtnAnimBackground
)

// TThemedTreeview ENUM
//
//	'Treeview' theme data
type TThemedTreeview = int32

const (
	TtTreeviewDontCare TThemedTreeview = iota
	TtTreeviewRoot
	TtItemNormal
	TtItemHot
	TtItemSelected
	TtItemDisabled
	TtItemSelectedNotFocus
	TtGlyphClosed
	TtGlyphOpened
	TtBranch
	TtHotGlyphClosed
	TtHotGlyphOpened
)

// TThemedWindow ENUM
//
//	'Window' theme data
type TThemedWindow = int32

const (
	TwWindowDontCare TThemedWindow = iota
	TwWindowRoot
	TwCaptionActive
	TwCaptionInactive
	TwCaptionDisabled
	TwSmallCaptionActive
	TwSmallCaptionInactive
	TwSmallCaptionDisabled
	TwMinCaptionActive
	TwMinCaptionInactive
	TwMinCaptionDisabled
	TwSmallMinCaptionActive
	TwSmallMinCaptionInactive
	TwSmallMinCaptionDisabled
	TwMaxCaptionActive
	TwMaxCaptionInactive
	TwMaxCaptionDisabled
	TwSmallMaxCaptionActive
	TwSmallMaxCaptionInactive
	TwSmallMaxCaptionDisabled
	TwFrameLeftActive
	TwFrameLeftInactive
	TwFrameRightActive
	TwFrameRightInactive
	TwFrameBottomActive
	TwFrameBottomInactive
	TwSmallFrameLeftActive
	TwSmallFrameLeftInactive
	TwSmallFrameRightActive
	TwSmallFrameRightInactive
	TwSmallFrameBottomActive
	TwSmallFrameBottomInactive
	TwSysButtonNormal
	TwSysButtonHot
	TwSysButtonPushed
	TwSysButtonDisabled
	TwSysButtonInactive
	TwMDISysButtonNormal
	TwMDISysButtonHot
	TwMDISysButtonPushed
	TwMDISysButtonDisabled
	TwMDISysButtonInactive
	TwMinButtonNormal
	TwMinButtonHot
	TwMinButtonPushed
	TwMinButtonDisabled
	TwMinButtonInactive
	TwMDIMinButtonNormal
	TwMDIMinButtonHot
	TwMDIMinButtonPushed
	TwMDIMinButtonDisabled
	TwMDIMinButtonInactive
	TwMaxButtonNormal
	TwMaxButtonHot
	TwMaxButtonPushed
	TwMaxButtonDisabled
	TwMaxButtonInactive
	TwCloseButtonNormal
	TwCloseButtonHot
	TwCloseButtonPushed
	TwCloseButtonDisabled
	TwCloseButtonInactive
	TwSmallCloseButtonNormal
	TwSmallCloseButtonHot
	TwSmallCloseButtonPushed
	TwSmallCloseButtonDisabled
	TwSmallCloseButtonInactive
	TwMDICloseButtonNormal
	TwMDICloseButtonHot
	TwMDICloseButtonPushed
	TwMDICloseButtonDisabled
	TwMDICloseButtonInactive
	TwRestoreButtonNormal
	TwRestoreButtonHot
	TwRestoreButtonPushed
	TwRestoreButtonDisabled
	TwRestoreButtonInactive
	TwMDIRestoreButtonNormal
	TwMDIRestoreButtonHot
	TwMDIRestoreButtonPushed
	TwMDIRestoreButtonDisabled
	TwMDIRestoreButtonInactive
	TwHelpButtonNormal
	TwHelpButtonHot
	TwHelpButtonPushed
	TwHelpButtonDisabled
	TwHelpButtonInactive
	TwMDIHelpButtonNormal
	TwMDIHelpButtonHot
	TwMDIHelpButtonPushed
	TwMDIHelpButtonDisabled
	TwMDIHelpButtonInactive
	TwHorzScrollNormal
	TwHorzScrollHot
	TwHorzScrollPushed
	TwHorzScrollDisabled
	TwHorzThumbNormal
	TwHorzThumbHot
	TwHorzThumbPushed
	TwHorzThumbDisabled
	TwVertScrollNormal
	TwVertScrollHot
	TwVertScrollPushed
	TwVertScrollDisabled
	TwVertThumbNormal
	TwVertThumbHot
	TwVertThumbPushed
	TwVertThumbDisabled
	TwDialog
	TwCaptionSizingTemplate
	TwSmallCaptionSizingTemplate
	TwFrameLeftSizingTemplate
	TwSmallFrameLeftSizingTemplate
	TwFrameRightSizingTemplate
	TwSmallFrameRightSizingTemplate
	TwFrameBottomSizingTemplate
	TwSmallFrameBottomSizingTemplate
)

// TThemeOption ENUM
type TThemeOption = int32

const (
	ToShowButtonImages TThemeOption = iota
	ToShowMenuImages
	ToUseGlyphEffects
)

// TTickMark ENUM
type TTickMark = int32

const (
	TmBottomRight TTickMark = iota
	TmTopLeft
	TmBoth
)

// TTickStyle ENUM
type TTickStyle = int32

const (
	TsNone TTickStyle = iota
	TsAuto
	TsManual
)

// TTimeDisplay ENUM
type TTimeDisplay = int32

const (
	TdHM TTimeDisplay = iota
	TdHMS
	TdHMSMs
)

// TTimeFormat ENUM
type TTimeFormat = int32

const (
	Tf12 TTimeFormat = iota
	Tf24
)

// TTitleStyle ENUM
type TTitleStyle = int32

const (
	TsLazarus TTitleStyle = iota
	TsStandard
	TsNative
)

// TToolButtonStyle ENUM
type TToolButtonStyle = int32

const (
	TbsButton TToolButtonStyle = iota
	TbsCheck
	TbsDropDown
	TbsSeparator
	TbsDivider
	TbsButtonDrop
)

// TTrackBarOrientation ENUM
//
//	TCustomTrackBar
type TTrackBarOrientation = int32

const (
	TrHorizontal TTrackBarOrientation = iota
	TrVertical
)

// TTrackBarScalePos ENUM
type TTrackBarScalePos = int32

const (
	TrLeft TTrackBarScalePos = iota
	TrRight
	TrTop
	TrBottom
)

// TTrackButton ENUM
type TTrackButton = int32

const (
	TbRightButton TTrackButton = iota
	TbLeftButton
)

// TTransparentMode ENUM
type TTransparentMode = int32

const (
	TmAuto TTransparentMode = iota
	TmFixed
)

// TTreeNodeChangeReason ENUM
type TTreeNodeChangeReason = int32

const (
	NcTextChanged TTreeNodeChangeReason = iota
	NcDataChanged
	NcHeightChanged
	NcImageEffect
	NcImageIndex
	NcParentChanged
	NcVisibility
	NcEnablement
	NcOverlayIndex
	NcStateIndex
	NcSelectedIndex
)

// TTreeViewExpandSignType ENUM
type TTreeViewExpandSignType = int32

const (
	TvestTheme TTreeViewExpandSignType = iota
	TvestPlusMinus
	TvestArrow
	TvestArrowFill
	TvestAngleBracket
)

// TTreeViewInsertMarkType ENUM
type TTreeViewInsertMarkType = int32

const (
	TvimNone TTreeViewInsertMarkType = iota
	TvimAsFirstChild
	TvimAsNextSibling
	TvimAsPrevSibling
)

// TUDAlignButton ENUM
type TUDAlignButton = int32

const (
	UdLeft TUDAlignButton = iota
	UdRight
	UdTop
	UdBottom
)

// TUDBtnType ENUM
type TUDBtnType = int32

const (
	BtNext TUDBtnType = iota
	BtPrev
)

// TUDOrientation ENUM
type TUDOrientation = int32

const (
	UdHorizontal TUDOrientation = iota
	UdVertical
)

// TUpDownDirection ENUM
type TUpDownDirection = int32

const (
	UpdNone TUpDownDirection = iota
	UpdUp
	UpdDown
)

// TVerticalAlignment ENUM
type TVerticalAlignment = int32

const (
	TaAlignTop TVerticalAlignment = iota
	TaAlignBottom
	TaVerticalCenter
)

// TViewStyle ENUM
type TViewStyle = int32

const (
	VsIcon TViewStyle = iota
	VsSmallIcon
	VsList
	VsReport
)

// TVirtualTreeColumnStyle ENUM
type TVirtualTreeColumnStyle = int32

const (
	VsText TVirtualTreeColumnStyle = iota
	VsOwnerDraw
)

// TVleSortCol ENUM
type TVleSortCol = int32

const (
	ColKey TVleSortCol = iota
	ColValue
)

// TVSTTextSourceType ENUM
//
//	Describes the source to use when converting a string tree into a string for clipboard etc.
type TVSTTextSourceType = int32

const (
	TstAll TVSTTextSourceType = iota
	TstInitialized
	TstSelected
	TstCutCopySet
	TstVisible
	TstChecked
)

// TVSTTextType ENUM
//
//	Describes the type of text to return in the text and draw info retrival events.
type TVSTTextType = int32

const (
	TtNormal TVSTTextType = iota
	TtStatic
)

// TVTButtonFillMode ENUM
//
//	TButtonFillMode is only used when the button style is bsRectangle and determines how to fill the interior.
type TVTButtonFillMode = int32

const (
	FmTreeColor TVTButtonFillMode = iota
	FmWindowColor
	FmShaded
	FmTransparent
)

// TVTButtonStyle ENUM
//
//	Determines the look of a tree's buttons.
type TVTButtonStyle = int32

const (
	BsRectangle TVTButtonStyle = iota
	BsTriangle
)

// TVTCellPaintMode ENUM
//
//	Determines for which purpose the cell paint event is called.
type TVTCellPaintMode = int32

const (
	CpmPaint TVTCellPaintMode = iota
	CpmGetContentMargin
)

// TVTDragImageKind ENUM
//
//	determines whether and how the drag image is to show
type TVTDragImageKind = int32

const (
	DiComplete TVTDragImageKind = iota
	DiMainColumnOnly
	DiNoImage
)

// TVTDragMoveRestriction ENUM
//
//	Simple move limitation for the drag image.
type TVTDragMoveRestriction = int32

const (
	DmrNone TVTDragMoveRestriction = iota
	DmrHorizontalOnly
	DmrVerticalOnly
)

// TVTDragType ENUM
//
//	Switch for OLE and VCL drag'n drop. Because it is not possible to have both simultanously.
type TVTDragType = int32

const (
	DtOLE TVTDragType = iota
	DtVCL
)

// TVTDrawSelectionMode ENUM
//
//	Determines how to draw the selection rectangle used for draw selection.
type TVTDrawSelectionMode = int32

const (
	SmDottedRectangle TVTDrawSelectionMode = iota
	SmBlendedRectangle
)

// TVTDropMarkMode ENUM
//
//	Used during owner draw of the header to indicate which drop mark for the column must be drawn.
type TVTDropMarkMode = int32

const (
	DmmNone TVTDropMarkMode = iota
	DmmLeft
	DmmRight
)

// TVTExportMode ENUM
//
//	Options to control data export
type TVTExportMode = int32

const (
	EmAll TVTExportMode = iota
	EmChecked
	EmUnchecked
	EmVisibleDueToExpansion
	EmSelected
)

// TVTExportType ENUM
//
//	Export type
type TVTExportType = int32

const (
	EtRTF TVTExportType = iota
	EtHTML
	EtText
	EtExcel
	EtWord
	EtCustom
)

// TVTHeaderColumnLayout ENUM
type TVTHeaderColumnLayout = int32

const (
	TVTHCLayoutBlGlyphLeft TVTHeaderColumnLayout = iota
	TVTHCLayoutBlGlyphRight
	TVTHCLayoutBlGlyphTop
	TVTHCLayoutBlGlyphBottom
)

// TVTHeaderStyle ENUM
type TVTHeaderStyle = int32

const (
	HsThickButtons TVTHeaderStyle = iota
	HsFlatButtons
	HsPlates
)

// TVTHintKind ENUM
type TVTHintKind = int32

const (
	VhkText TVTHintKind = iota
	VhkOwnerDraw
)

// TVTHintMode ENUM
type TVTHintMode = int32

const (
	HmDefault TVTHintMode = iota
	HmHint
	HmHintAndDefault
	HmTooltip
)

// TVTImageKind ENUM
type TVTImageKind = int32

const (
	IkNormal TVTImageKind = iota
	IkSelected
	IkState
	IkOverlay
)

// TVTIncrementalSearch ENUM
type TVTIncrementalSearch = int32

const (
	IsAll TVTIncrementalSearch = iota
	IsNone
	IsInitializedOnly
	IsVisibleOnly
)

// TVTLineMode ENUM
//
//	Determines how to draw tree lines.
type TVTLineMode = int32

const (
	LmNormal TVTLineMode = iota
	LmBands
)

// TVTLineStyle ENUM
//
//	Determines the look of a tree's lines.
type TVTLineStyle = int32

const (
	LsCustomStyle TVTLineStyle = iota
	LsDotted
	LsSolid
)

// TVTNodeAlignment ENUM
//
//	Determines how to use the align member of a node.
type TVTNodeAlignment = int32

const (
	NaFromBottom TVTNodeAlignment = iota
	NaFromTop
	NaProportional
)

// TVTNodeAttachMode ENUM
//
//	mode to describe a move action
type TVTNodeAttachMode = int32

const (
	AmNoWhere TVTNodeAttachMode = iota
	AmInsertBefore
	AmInsertAfter
	AmAddChildFirst
	AmAddChildLast
)

// TVTOperationKind ENUM
//
//	Kinds of operations
type TVTOperationKind = int32

const (
	OkAutoFitColumns TVTOperationKind = iota
	OkGetMaxColumnWidth
	OkSortNode
	OkSortTree
)

// TVTScrollBarStyle ENUM
type TVTScrollBarStyle = int32

const (
	SbmRegular TVTScrollBarStyle = iota
	Sbm3D
)

// TVTSearchDirection ENUM
//
//	Determines which direction to use when advancing nodes during an incremental search.
type TVTSearchDirection = int32

const (
	SdForward TVTSearchDirection = iota
	SdBackward
)

// TVTSearchStart ENUM
//
//	Determines where to start incremental searching for each key press.
type TVTSearchStart = int32

const (
	SsAlwaysStartOver TVTSearchStart = iota
	SsLastHit
	SsFocusedNode
)

// TVTTooltipLineBreakStyle ENUM
//
//	Indicates how to format a tooltip.
type TVTTooltipLineBreakStyle = int32

const (
	HlbDefault TVTTooltipLineBreakStyle = iota
	HlbForceSingleLine
	HlbForceMultiLine
)

// TVTUpdateState ENUM
//
//	Indicates in the OnUpdating event what state the tree is currently in.
type TVTUpdateState = int32

const (
	UsBegin TVTUpdateState = iota
	UsBeginSynch
	UsSynch
	UsUpdate
	UsEnd
	UsEndSynch
)

// TWindowState ENUM
type TWindowState = int32

const (
	WsNormal TWindowState = iota
	WsMinimized
	WsMaximized
	WsFullScreen
)

// TWrapAfter ENUM
type TWrapAfter = int32

const (
	WaAuto TWrapAfter = iota
	WaForce
	WaAvoid
	WaForbid
)
