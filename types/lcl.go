//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

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

// TFontStylesBase SET: TFontStyle
type TFontStylesBase = TSet

// TFontStyles SET: TFontStyle
type TFontStyles = TSet

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

// TParaModifyMaske ENUM
type TParaModifyMaske = int32

const (
	Pmm_FirstLine TParaModifyMaske = iota
	Pmm_HeadIndent
	Pmm_TailIndent
	Pmm_SpaceBefore
	Pmm_SpaceAfter
	Pmm_LineSpacing
)

// TParaModifyMask SET: TParaModifyMaske
type TParaModifyMask = TSet

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

// TSearchOption ENUM
type TSearchOption = int32

const (
	SoMatchCase TSearchOption = iota
	SoWholeWord
	SoBackward
)

// TSearchOptions SET: TSearchOption
type TSearchOptions = TSet

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
//
//	the available configuration flags for the Task Dialog
//	- most are standard TDF_* flags used for Vista/Seven native API
//	(see http://msdn.microsoft.com/en-us/library/bb787473(v=vs.85).aspx
//	for TASKDIALOG_FLAGS)
//	- tdfQuery and tdfQueryMasked are custom flags, implemented in pure Delphi
//	code to handle input query
//	- our emulation code will handle only tdfUseCommandLinks,
//	tdfUseCommandLinksNoIcon, and tdfQuery options
type TTaskDialogFlag = int32

const (
	TdfEnableHyperLinks TTaskDialogFlag = iota
	TdfUseHIconMain
	TdfUseHIconFooter
	TdfAllowDialogCancellation
	TdfUseCommandLinks
	TdfUseCommandLinksNoIcon
	TdfExpandFooterArea
	TdfExpandByDefault
	TdfVerificationFlagChecked
	TdfShowProgressBar
	TdfShowMarqueeProgressBar
	TdfCallbackTimer
	TdfPositionRelativeToWindow
	TdfRtlLayout
	TdfNoDefaultRadioButton
	TdfCanBeMinimized
	TdfNoSetForeGround
	TdfSizeToContent
	TdfQuery
	TdfQueryMasked
	TdfQueryFieldFocused
)

// TTaskDialogFlags SET: TTaskDialogFlag
type TTaskDialogFlags = TSet

// TTextModifyMaske ENUM
type TTextModifyMaske = int32

const (
	Tmm_Color TTextModifyMaske = iota
	Tmm_Name
	Tmm_Size
	Tmm_Styles
	Tmm_BackColor
)

// TTextModifyMask SET: TTextModifyMaske
type TTextModifyMask = TSet

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

// CGLContextEnable ENUM
//
//	** Enable names for CGLEnable, CGLDisable, and CGLIsEnabled.
type CGLContextEnable = int32

const (
	KCGLCESwapRectangle           = 201
	KCGLCESwapLimit               = 203
	KCGLCERasterization           = 221
	KCGLCEStateValidation         = 301
	KCGLCESurfaceBackingSize      = 305
	KCGLCEDisplayListOptimization = 307
	KCGLCEMPEngine                = 313
)

// CGLContextParameter ENUM
//
//	** Parameter names for CGLSetParameter and CGLGetParameter.
type CGLContextParameter = int32

const (
	KCGLCPSwapRectangle          = 200
	KCGLCPSwapInterval           = 222
	KCGLCPDispatchTableSize      = 224
	KCGLCPClientStorage          = 226
	KCGLCPSurfaceTexture         = 228
	KCGLCPSurfaceOrder           = 235
	KCGLCPSurfaceOpacity         = 236
	KCGLCPSurfaceBackingSize     = 304
	KCGLCPSurfaceSurfaceVolatile = 306
	KCGLCPReclaimResources       = 308
	KCGLCPCurrentRendererID      = 309
	KCGLCPGPUVertexProcessing    = 310
	KCGLCPGPUFragmentProcessing  = 311
)

// CGLError ENUM
//
//	** Error return values from CGLGetError.
type CGLError = int32

const (
	KCGLNoError         CGLError = iota
	KCGLBadAttribute             = 10000
	KCGLBadProperty              = 10001
	KCGLBadPixelFormat           = 10002
	KCGLBadRendererInfo          = 10003
	KCGLBadContext               = 10004
	KCGLBadDrawable              = 10005
	KCGLBadDisplay               = 10006
	KCGLBadState                 = 10007
	KCGLBadValue                 = 10008
	KCGLBadMatch                 = 10009
	KCGLBadEnumeration           = 10010
	KCGLBadOffScreen             = 10011
	KCGLBadFullScreen            = 10012
	KCGLBadWindow                = 10013
	KCGLBadAddress               = 10014
	KCGLBadCodeModule            = 10015
	KCGLBadAlloc                 = 10016
	KCGLBadConnection            = 10017
)

// CGLGlobalOption ENUM
//
//	** Option names for CGLSetOption and CGLGetOption.
type CGLGlobalOption = int32

const (
	KCGLGOFormatCacheSize  = 501
	KCGLGOClearFormatCache = 502
	KCGLGORetainRenderers  = 503
	KCGLGOResetLibrary     = 504
	KCGLGOUseErrorHandler  = 505
)

// CGLPixelFormatAttribute ENUM
//
//	** Attribute names for CGLChoosePixelFormat and CGLDescribePixelFormat.
type CGLPixelFormatAttribute = int32

const (
	KCGLPFAAllRenderers       = 1
	KCGLPFADoubleBuffer       = 5
	KCGLPFAStereo             = 6
	KCGLPFAAuxBuffers         = 7
	KCGLPFAColorSize          = 8
	KCGLPFAAlphaSize          = 11
	KCGLPFADepthSize          = 12
	KCGLPFAStencilSize        = 13
	KCGLPFAAccumSize          = 14
	KCGLPFAMinimumPolicy      = 51
	KCGLPFAMaximumPolicy      = 52
	KCGLPFAOffScreen          = 53
	KCGLPFAFullScreen         = 54
	KCGLPFASampleBuffers      = 55
	KCGLPFASamples            = 56
	KCGLPFAAuxDepthStencil    = 57
	KCGLPFAColorFloat         = 58
	KCGLPFAMultisample        = 59
	KCGLPFASupersample        = 60
	KCGLPFASampleAlpha        = 61
	KCGLPFARendererID         = 70
	KCGLPFASingleRenderer     = 71
	KCGLPFANoRecovery         = 72
	KCGLPFAAccelerated        = 73
	KCGLPFAClosestPolicy      = 74
	KCGLPFARobust             = 75
	KCGLPFABackingStore       = 76
	KCGLPFAMPSafe             = 78
	KCGLPFAWindow             = 80
	KCGLPFAMultiScreen        = 81
	KCGLPFACompliant          = 83
	KCGLPFADisplayMask        = 84
	KCGLPFAPBuffer            = 90
	KCGLPFARemotePBuffer      = 91
	KCGLPFAVirtualScreenCount = 128
)

// CGLRendererProperty ENUM
//
//	** Property names for CGLDescribeRenderer.
type CGLRendererProperty = int32

const (
	KCGLRPOffScreen          = 53
	KCGLRPFullScreen         = 54
	KCGLRPRendererID         = 70
	KCGLRPAccelerated        = 73
	KCGLRPRobust             = 75
	KCGLRPBackingStore       = 76
	KCGLRPMPSafe             = 78
	KCGLRPWindow             = 80
	KCGLRPMultiScreen        = 81
	KCGLRPCompliant          = 83
	KCGLRPDisplayMask        = 84
	KCGLRPBufferModes        = 100
	KCGLRPColorModes         = 103
	KCGLRPAccumModes         = 104
	KCGLRPDepthModes         = 105
	KCGLRPStencilModes       = 106
	KCGLRPMaxAuxBuffers      = 107
	KCGLRPMaxSampleBuffers   = 108
	KCGLRPMaxSamples         = 109
	KCGLRPSampleModes        = 110
	KCGLRPSampleAlpha        = 111
	KCGLRPVideoMemory        = 120
	KCGLRPTextureMemory      = 121
	KCGLRPGPUVertProcCapable = 122
	KCGLRPGPUFragProcCapable = 123
	KCGLRPRendererCount      = 128
)

// LPosFlag ENUM
type LPosFlag = int32

const (
	LpAllowPastEol LPosFlag = iota
	LpAdjustToNext
	LpStopAtCodePoint
)

// MONITOR_DPI_TYPE ENUM
type MONITOR_DPI_TYPE = int32

const (
	MDT_EFFECTIVE_DPI MONITOR_DPI_TYPE = iota
	MDT_ANGULAR_DPI                    = 1
	MDT_RAW_DPI                        = 2
	MDT_DEFAULT                        = MDT_EFFECTIVE_DPI
)

// taComponents ENUM
type taComponents = int32

const (
	TacTarrow taComponents = iota
	TacTbitbtn
	TacTbutton
	TacTbuttonpanel
	TacTcalcedit
	TacTcalendar
	TacTcheckbox
	TacTcheckcombobox
	TacTcheckgroup
	TacTchecklistbox
	TacTcolorbox
	TacTcolorbutton
	TacTcolorlistbox
	TacTcombobox
	TacTcomboboxex
	TacTcontrolbar
	TacTcoolbar
	TacTdateedit
	TacTdirectoryedit
	TacTedit
	TacTeditbutton
	TacTfilelistbox
	TacTfilenameedit
	TacTfiltercombobox
	TacTfloatspinedit
	TacTgroupbox
	TacTheadercontrol
	TacTimage
	TacTlabel
	TacTlabelededit
	TacTlistbox
	TacTlistview
	TacTmaskedit
	TacTmemo
	TacTnotebook
	TacTpagecontrol
	TacTpaintbox
	TacTpanel
	TacTprogressbar
	TacTradiobutton
	TacTradiogroup
	TacTshape
	TacTshelllistview
	TacTshelltreeview
	TacTspeedbutton
	TacTspinedit
	TacTsplitter
	TacTstacTictext
	TacTstatusbar
	TacTstringgrid
	TacTtabcontrol
	TacTtimeedit
	TacTtimer
	TacTtogglebox
	TacTtoolbar
	TacTtrackbar
	TacTtreeview
	TacTupdown
	TacTvaluelisteditor
)

// TActionListState ENUM
type TActionListState = int32

const (
	AsNormal TActionListState = iota
	AsSuspended
	AsSuspendedEnabled
)

// TActiveXRegType ENUM
//
//	Component registration handlers
type TActiveXRegType = int32

const (
	AxrComponentOnly TActiveXRegType = iota
	AxrIncludeDescendants
)

// TAddMode ENUM
type TAddMode = int32

const (
	TaAddFirst TAddMode = iota
	TaAdd
	TaInsert
)

// TAddPopupItemType ENUM
type TAddPopupItemType = int32

const (
	ApNormal TAddPopupItemType = iota
	ApDisabled
	ApHidden
)

// taDialogs ENUM
type taDialogs = int32

const (
	TadTcalculatordialog taDialogs = iota
	TadTcalendardialog
	TadTcolordialog
	TadTfinddialog
	TadTfontdialog
	TadTMessageDialog
	TadTopendialog
	TadTopenpicturedialog
	TadTreplacedialog
	TadTsavedialog
	TadTsavepicturedialog
	TadTselectdirectorydialog
	TadTQuestionDialog
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

// TAnchorSideChangeOperation ENUM
//
//	TAnchorSide
//	Class holding the reference sides of the anchors of a TControl.
//	Every TControl has four AnchorSides:
//	AnchorSide[akLeft], AnchorSide[akRight], AnchorSide[akTop] and
//	AnchorSide[akBottom].
//	Normally if Anchors contain akLeft, and the Parent is resized, the LCL
//	tries to keep the distance between the left side of the control and the
//	right side of its parent client area.
//	With AnchorSide[akLeft] you can define a different reference side. The
//	kept distance is defined by the BorderSpacing and Parent.ChildSizing.
//	Example1:
//	+-----+ +-----+
//	| B | | C |
//	| | +-----+
//	+-----+
//	If you want to have the top of B the same as the top of C use
//	B.AnchorSide[akTop].Side:=asrTop;
//	B.AnchorSide[akTop].Control:=C;
//	If you want to keep a distance of 10 pixels between B and C use
//	B.BorderSpacing.Right:=10;
//	B.AnchorSide[akRight].Side:=asrLeft;
//	B.AnchorSide[akRight].Control:=C;
//	Do not setup in both directions, because this will create a circle, and
//	circles are not allowed.
//	Example2:
//	+-------+
//	+---+ | |
//	| A | | B |
//	+---+ | |
//	+-------+
//	Centering A relative to B:
//	A.AnchorSide[akTop].Side:=arsCenter;
//	A.AnchorSide[akTop].Control:=B;
//	Or use this. It's equivalent:
//	A.AnchorSide[akBottom].Side:=arsCenter;
//	A.AnchorSide[akBottom].Control:=B;
type TAnchorSideChangeOperation = int32

const (
	AscoAdd TAnchorSideChangeOperation = iota
	AscoRemove
	AscoChangeSide
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

// TAppHintTimerType ENUM
type TAppHintTimerType = int32

const (
	AhttNone TAppHintTimerType = iota
	AhttShowHint
	AhttHideHint
	AhttReshowHint
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

// TApplicationHandlerType ENUM
type TApplicationHandlerType = int32

const (
	AhtIdle TApplicationHandlerType = iota
	AhtIdleEnd
	AhtKeyDownBefore
	AhtKeyDownAfter
	AhtActivate
	AhtDeactivate
	AhtUserInput
	AhtException
	AhtEndSession
	AhtQueryEndSession
	AhtMinimize
	AhtModalBegin
	AhtModalEnd
	AhtRestore
	AhtDropFiles
	AhtHelp
	AhtHint
	AhtShowHint
	AhtGetMainFormHandle
	AhtActionExecute
	AhtActionUpdate
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

// TArrowType ENUM
type TArrowType = int32

const (
	AtUp TArrowType = iota
	AtDown
	AtLeft
	AtRight
)

// TATButtonOverlayPosition ENUM
type TATButtonOverlayPosition = int32

const (
	BopLeftTop TATButtonOverlayPosition = iota
	BopRightTop
	BopLeftBottom
	BopRightBottom
)

// TATGaugeKind ENUM
type TATGaugeKind = int32

const (
	GkText1 TATGaugeKind = iota
	GkHorizontalBar1
	GkVerticalBar1
	GkPie1
	GkNeedle1
	GkHalfPie1
)

// TAttributeType ENUM
type TAttributeType = int32

const (
	AtSelected TAttributeType = iota
	AtDefaultText
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

// TAutoSizeBoxOrientation ENUM
type TAutoSizeBoxOrientation = int32

const (
	AsboHorizontal TAutoSizeBoxOrientation = iota
	AsboVertical
)

// TAutoSizeSideDataState ENUM
type TAutoSizeSideDataState = int32

const (
	AssdfInvalid TAutoSizeSideDataState = iota
	AssdfComputing
	AssdfUncomputable
	AssdfValid
)

// TAutoSizeSideDistDirection ENUM
type TAutoSizeSideDistDirection = int32

const (
	AssddLeftTop TAutoSizeSideDistDirection = iota
	AssddRightBottom
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

// TBandMove ENUM
type TBandMove = int32

const (
	BmNone1 TBandMove = iota
	BmReady1
	BmMoving1
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

// TBom ENUM
type TBom = int32

const (
	BomUtf8 TBom = iota
	BomUtf16BE
	BomUtf16LE
	BomUndefined
)

// TBookmarkedRecordEnumeratorOptions ENUM
type TBookmarkedRecordEnumeratorOptions = int32

const (
	BreDisableDataset TBookmarkedRecordEnumeratorOptions = iota
	BreStopOnInvalidBookmark
	BreRestoreCurrent
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

// TButtonState ENUM
type TButtonState = int32

const (
	BsUp TButtonState = iota
	BsDisabled
	BsDown
	BsExclusive
	BsHot
)

// TCalculatorLayout ENUM
type TCalculatorLayout = int32

const (
	ClNormal TCalculatorLayout = iota
	ClSimple
)

// TCalculatorState ENUM
type TCalculatorState = int32

const (
	CsFirst TCalculatorState = iota
	CsValid
	CsError
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

// TCanvasOrientation ENUM
type TCanvasOrientation = int32

const (
	CsLefttoRight TCanvasOrientation = iota
	CoRighttoLeft
)

// TCarbonBitmapAlignment ENUM
//
//	TCarbonBitmap
type TCarbonBitmapAlignment = int32

const (
	CbaByte TCarbonBitmapAlignment = iota
	CbaWord
	CbaDWord
	CbaQWord
	CbaDQWord
)

// TCarbonBitmapType ENUM
type TCarbonBitmapType = int32

const (
	CbtMono TCarbonBitmapType = iota
	CbtGray
	CbtRGB
	CbtARGB
	CbtRGBA
	CbtBGR
	CbtBGRA
)

// TCarbonControlEvent ENUM
type TCarbonControlEvent = int32

const (
	CceValueChanged TCarbonControlEvent = iota
	CceIndicatorMoved
	CceDoAction
	CceDraw
	CceHit
)

// TCarbonCursorType ENUM
type TCarbonCursorType = int32

const (
	CctUnknown TCarbonCursorType = iota
	CctQDHardware
	CctQDColor
	CctTheme
	CctAnimated
	CctWait
)

// TCarbonWidgetFlag ENUM
type TCarbonWidgetFlag = int32

const (
	CwfNone TCarbonWidgetFlag = iota
	CwdTToolBar
	CwdTTabControl
)

// TCDWSEventCapability ENUM
//
//	TCommonDialog
type TCDWSEventCapability = int32

const (
	CdecWSPerformsDoShow TCDWSEventCapability = iota
	CdecWSPerformsDoCanClose
	CdecWSPerformsDoClose
	CdecWSNOCanCloseSupport
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

// TChildExitReason ENUM
type TChildExitReason = int32

const (
	CerExit TChildExitReason = iota
	CerSignal
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

// TCocoaBitmapAlignment ENUM
type TCocoaBitmapAlignment = int32

const (
	CbaByte1 TCocoaBitmapAlignment = iota
	CbaWord1
	CbaDWord1
	CbaQWord1
	CbaDQWord1
)

// TCocoaBitmapType ENUM
type TCocoaBitmapType = int32

const (
	CbtMono1 TCocoaBitmapType = iota
	CbtGray1
	CbtRGB1
	CbtARGB1
	CbtRGBA1
	CbtABGR1
	CbtBGRA1
)

// TCocoaClipboardDataType ENUM
type TCocoaClipboardDataType = int32

const (
	CcdtText TCocoaClipboardDataType = iota
	CcdtCocoaStandard
	CcdtBitmap
	CcdtNonStandard
)

// TCocoaCombine ENUM
type TCocoaCombine = int32

const (
	Cc_And TCocoaCombine = iota
	Cc_Xor
	Cc_Or
	Cc_Diff
	Cc_Copy
)

// TCocoaPatternColorMode ENUM
type TCocoaPatternColorMode = int32

const (
	CpmBitmap TCocoaPatternColorMode = iota
	CpmBrushColor
	CpmContextColor
)

// TCocoaRegionType ENUM
type TCocoaRegionType = int32

const (
	Crt_Error TCocoaRegionType = iota
	Crt_Empty
	Crt_Rectangle
	Crt_Complex
)

// TCollectionNotification ENUM
type TCollectionNotification = int32

const (
	CnAdded TCollectionNotification = iota
	CnExtracting
	CnDeleting
)

// TColorFormat ENUM
type TColorFormat = int32

const (
	CfMono TColorFormat = iota
	CfGray2
	CfGray4
	CfGray8
	CfGray16
	CfGray24
	CfGrayA8
	CfGrayA16
	CfGrayA32
	CfRGB15
	CfRGB16
	CfRGB24
	CfRGB32
	CfRGB48
	CfRGBA8
	CfRGBA16
	CfRGBA32
	CfRGBA64
	CfBGR15
	CfBGR16
	CfBGR24
	CfBGR32
	CfBGR48
	CfABGR8
	CfABGR16
	CfABGR32
	CfABGR64
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

// TColumnOrder ENUM
type TColumnOrder = int32

const (
	CoDesignOrder TColumnOrder = iota
	CoFieldIndexOrder
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

// TCommonButton ENUM
//
//	the standard kind of common buttons handled by the Task Dialog
type TCommonButton = int32

const (
	CbOK TCommonButton = iota
	CbYes
	CbNo
	CbCancel
	CbRetry
	CbClose
)

// TComputeResult ENUM
type TComputeResult = int32

const (
	CrSuccess TComputeResult = iota
	CrCircle
	CrFixedCircled
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

// TControlFlag ENUM
type TControlFlag = int32

const (
	CfLoading TControlFlag = iota
	CfAutoSizeNeeded
	CfLeftLoaded
	CfTopLoaded
	CfWidthLoaded
	CfHeightLoaded
	CfClientWidthLoaded
	CfClientHeightLoaded
	CfBoundsRectForNewParentValid
	CfBaseBoundsValid
	CfPreferredSizeValid
	CfPreferredMinSizeValid
	CfOnChangeBoundsNeeded
	CfProcessingWMPaint
	CfKillChangeBounds
	CfKillInvalidatePreferredSize
	CfKillAdjustSize
)

// TControlHandlerType ENUM
type TControlHandlerType = int32

const (
	ChtOnResize TControlHandlerType = iota
	ChtOnChangeBounds
	ChtOnVisibleChanging
	ChtOnVisibleChanged
	ChtOnEnabledChanging
	ChtOnEnabledChanged
	ChtOnKeyDown
	ChtOnBeforeDestruction
	ChtOnMouseWheel
	ChtOnMouseWheelHorz
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

// TCursorDesign ENUM
type TCursorDesign = int32

const (
	CdDefault TCursorDesign = iota
	CdGrabber
	CdRestricted
)

// TCustomDrawResultFlag ENUM
type TCustomDrawResultFlag = int32

const (
	CdrSkipDefault TCustomDrawResultFlag = iota
	CdrNotifyPostpaint
	CdrNotifyItemdraw
	CdrNotifySubitemdraw
	CdrNotifyPosterase
	CdrNotifyItemerase
)

// TCustomDrawStage ENUM
type TCustomDrawStage = int32

const (
	CdPrePaint TCustomDrawStage = iota
	CdPostPaint
	CdPreErase
	CdPostErase
)

// TCustomDrawTarget ENUM
//
//	Custom draw
type TCustomDrawTarget = int32

const (
	DtControl TCustomDrawTarget = iota
	DtItem
	DtSubItem
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

// TDateTimePart ENUM
type TDateTimePart = int32

const (
	DtpDay TDateTimePart = iota
	DtpMonth
	DtpYear
	DtpHour
	DtpMinute
	DtpSecond
	DtpMiliSec
	DtpAMPM
)

// TDbGridExtraOption ENUM
type TDbGridExtraOption = int32

const (
	DgeAutoColumns TDbGridExtraOption = iota
	DgeCheckboxColumn
)

// TDBGridOption ENUM
type TDBGridOption = int32

const (
	DgEditing TDBGridOption = iota
	DgTitles
	DgIndicator
	DgColumnResize
	DgColumnMove
	DgColLines
	DgRowLines
	DgTabs
	DgAlwaysShowEditor
	DgRowSelect
	DgAlwaysShowSelection
	DgConfirmDelete
	DgCancelOnExit
	DgMultiselect
	DgHeaderHotTracking
	DgHeaderPushedLook
	DgPersistentMultiSelect
	DgAutoSizeColumns
	DgAnyButtonCanSelect
	DgDisableDelete
	DgDisableInsert
	DgCellHints
	DgTruncCellHints
	DgCellEllipsis
	DgRowHighlight
	DgThumbTracking
	DgDblClickAutoSize
	DgDisplayMemoText
)

// TDbGridStatusItem ENUM
type TDbGridStatusItem = int32

const (
	GsUpdatingData TDbGridStatusItem = iota
	GsAddingAutoColumns
	GsRemovingAutoColumns
	GsAutoSized
	GsStartEditing
	GsLoadingGrid
)

// TDBNavButtonDirection ENUM
type TDBNavButtonDirection = int32

const (
	NbdHorizontal TDBNavButtonDirection = iota
	NbdVertical
)

// TDBNavButtonType ENUM
type TDBNavButtonType = int32

const (
	NbFirst TDBNavButtonType = iota
	NbPrior
	NbNext
	NbLast
	NbInsert
	NbDelete
	NbEdit
	NbPost
	NbCancel
	NbRefresh
)

// TDBNavGlyph ENUM
type TDBNavGlyph = int32

const (
	NgEnabled TDBNavGlyph = iota
	NgDisabled
)

// TDBNavigatorOption ENUM
type TDBNavigatorOption = int32

const (
	NavFocusableButtons TDBNavigatorOption = iota
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

// TDockHeaderImageKind ENUM
type TDockHeaderImageKind = int32

const (
	DhiRestore TDockHeaderImageKind = iota
	DhiClose
)

// TDockImageOperation ENUM
type TDockImageOperation = int32

const (
	DisShow TDockImageOperation = iota
	DisMove
	DisHide
)

// TDockOrientation ENUM
type TDockOrientation = int32

const (
	DoNoOrient TDockOrientation = iota
	DoHorizontal
	DoVertical
	DoPages
)

// TDockTreeFlag ENUM
type TDockTreeFlag = int32

const (
	DtfUpdateAllNeeded TDockTreeFlag = iota
)

// TDragBand ENUM
type TDragBand = int32

const (
	DbNone TDragBand = iota
	DbMove
	DbResize
)

// TDragKind ENUM
type TDragKind = int32

const (
	DkDrag TDragKind = iota
	DkDock
)

// TDragMessage ENUM
type TDragMessage = int32

const (
	DmDragEnter TDragMessage = iota
	DmDragLeave
	DmDragMove
	DmDragDrop
	DmDragCancel
	DmFindTarget
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
	DsFocus1 TDrawingStyle = iota
	DsSelected1
	DsNormal1
	DsTransparent1
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

// TEraseBkgndCommand ENUM
type TEraseBkgndCommand = int32

const (
	EcDefault TEraseBkgndCommand = iota
	EcDiscard
	EcDiscardNoRemove
	EcDoubleBufferNoRemove
)

// TEventType ENUM
type TEventType = int32

const (
	EtNotify TEventType = iota
	EtKey
	EtKeyPress
	EtMouseWheel
	EtMouseUpDown
)

// TExpandCollapseMode ENUM
type TExpandCollapseMode = int32

const (
	EcmRefreshedExpanding TExpandCollapseMode = iota
	EcmKeepChildren
	EcmCollapseAndClear
)

// TFileAttr ENUM
//
//	TCustomFileListBox
type TFileAttr = int32

const (
	FtReadOnly TFileAttr = iota
	FtHidden
	FtSystem
	FtVolumeID
	FtDirectory
	FtArchive
	FtNormal
)

// TFilerFlag ENUM
type TFilerFlag = int32

const (
	FfInherited TFilerFlag = iota
	FfChildPos
	FfInline
)

// TFileSortType ENUM
type TFileSortType = int32

const (
	FstNone TFileSortType = iota
	FstAlphabet
	FstFoldersFirst
	FstCustom
)

// TFillMode ENUM
type TFillMode = int32

const (
	FmAlternate TFillMode = iota
	FmWinding
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

// TFormHandlerType ENUM
type TFormHandlerType = int32

const (
	FhtFirstShow TFormHandlerType = iota
	FhtClose
	FhtCreate
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
	PecRound1 TFPPenEndCap = iota
	PecSquare1
	PecFlat1
)

// TFPPenJoinStyle ENUM
type TFPPenJoinStyle = int32

const (
	PjsRound1 TFPPenJoinStyle = iota
	PjsBevel1
	PjsMiter1
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

// TGaugeKind ENUM
//
//	重新定义，目的为了兼容ATGauge的
type TGaugeKind = int32

const (
	GkText TGaugeKind = iota
	GkHorizontalBar
	GkVerticalBar
	GkPie
	GkNeedle
	GkHalfPie
)

// TGlyphShowMode ENUM
type TGlyphShowMode = int32

const (
	GsmAlways TGlyphShowMode = iota
	GsmNever
	GsmApplication
	GsmSystem
)

// TGlyphTransparencyMode ENUM
//
//	TButtonGlyph
type TGlyphTransparencyMode = int32

const (
	GtmGlyph TGlyphTransparencyMode = iota
	GtmOpaque
	GtmTransparent
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

// TGraphicsFillStyle ENUM
type TGraphicsFillStyle = int32

const (
	FsSurface TGraphicsFillStyle = iota
	FsBorder
)

// TGridCursorState ENUM
type TGridCursorState = int32

const (
	GcsDefault TGridCursorState = iota
	GcsColWidthChanging
	GcsRowHeightChanging
	GcsDragging
)

// TGridFlagsOption ENUM
type TGridFlagsOption = int32

const (
	GfEditorUpdateLock TGridFlagsOption = iota
	GfNeedsSelectActive
	GfEditorTab
	GfRevEditorTab
	GfVisualChange
	GfColumnsLocked
	GfEditingDone
	GfSizingStarted
	GfPainting
	GfUpdatingSize
	GfClientRectChange
	GfAutoEditPending
	GfUpdatingScrollbar
)

// TGridState ENUM
type TGridState = int32

const (
	GsNormal TGridState = iota
	GsSelecting
	GsRowSizing
	GsColSizing
	GsRowMoving
	GsColMoving
	GsHeaderClicking
	GsButtonColumnClicking
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

// THelpNodeType ENUM
//
//	THelpNode
//	A help node is a position/place in a help database.
//	For example it points to a Help file or to a Link on a HTML file.
type THelpNodeType = int32

const (
	HntURLIDContext THelpNodeType = iota
	HntURL
	HntURLID
	HntID
	HntContext
	HntURLContext
)

// THelpType ENUM
type THelpType = int32

const (
	HtKeyword THelpType = iota
	HtContext
)

// THookedCommandFlag ENUM
type THookedCommandFlag = int32

const (
	HcfInit THookedCommandFlag = iota
	HcfPreExec
	HcfPostExec
	HcfFinish
)

// TicnsIconType ENUM
//
//	from lower to higher
type TicnsIconType = int32

const (
	IitNone TicnsIconType = iota
	IitMini4BitData
	IitMini8BitData
	IitSmall4BitData
	IitSmall8BitData
	IitSmall32BitData
	IitLarge4BitData
	IitLarge8BitData
	IitLarge32BitData
	IitHuge4BitData
	IitHuge8BitData
	IitHuge32BitData
	IitThumbnail32BitData
	IitMini1BitMask
	IitSmall1BitMask
	IitSmall8BitMask
	IitLarge1BitMask
	IitLarge8BitMask
	IitHuge1BitMask
	IitHuge8BitMask
	IitThumbnail8BitMask
	Iit256PixelDataARGB
	Iit512PixelDataARGB
)

// TIconArrangement ENUM
//
//	TIconOptions
type TIconArrangement = int32

const (
	IaTop TIconArrangement = iota
	IaLeft
)

// TIdleTimerAutoEvent ENUM
//
//	TIdleTimer
//	For example:
//	Do something after 2 seconds after user input and idle.
//	AutoEnabled:=true;
//	AutoStartEvent:=itaOnIdle; // start the timer on first idle
//	AutoEndEvent:=itaOnUserInput; // end on any user input
//	If the OnTimer event works in several chunks, set FireOnIdle:=true.
//	The OnTimer event will then be called on idle until FireOnIdle is false.
//	FireOnIdle is set to false on any user input.
type TIdleTimerAutoEvent = int32

const (
	ItaOnIdle TIdleTimerAutoEvent = iota
	ItaOnIdleEnd
	ItaOnUserInput
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

// tinterfaceentrytype ENUM
//
//	This enumerate is found both in the rtl and compiler. Do not change the order of the fields.
type tinterfaceentrytype = int32

const (
	EtStandard tinterfaceentrytype = iota
	EtVirtualMethodResult
	EtStaticMethodResult
	EtFieldValue
	EtVirtualMethodClass
	EtStaticMethodClass
	EtFieldValueClass
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

// TItemType ENUM
type TItemType = int32

const (
	ItNormal TItemType = iota
	ItCell
	ItColumn
	ItRow
	ItFixed
	ItFixedColumn
	ItFixedRow
	ItSelected
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

// TLazCanvasImageFormat ENUM
type TLazCanvasImageFormat = int32

const (
	ClfOther TLazCanvasImageFormat = iota
	ClfRGB16_R5G6B5
	ClfRGB24
	ClfRGB24UpsideDown
	ClfBGR24
	ClfBGRA32
	ClfRGBA32
	ClfARGB32
)

// TLazDeviceMessageKind ENUM
//
//	TLazMessaging
type TLazDeviceMessageKind = int32

const (
	DmkSMS TLazDeviceMessageKind = iota
	DmkMMS
	DmkEMail
)

// TLazDockHeaderPart ENUM
type TLazDockHeaderPart = int32

const (
	LdhpAll TLazDockHeaderPart = iota
	LdhpCaption
	LdhpRestoreButton
	LdhpCloseButton
)

// TLazFileDialogKind ENUM
type TLazFileDialogKind = int32

const (
	LdkOpenDesktop TLazFileDialogKind = iota
	LdkSaveDesktop
	LdkOpenPDA
	LdkSavePDA
	LdkSelectDirectory
)

// TLazMessagingStatus ENUM
type TLazMessagingStatus = int32

const (
	MssSentSuccessfully TLazMessagingStatus = iota
	MssSendingGeneralError
	MssRadioOff
	MssNoService
	MssReceivedSuccessfully
	MssReceivingGeneralError
)

// TLazPositionMethod ENUM
//
//	TLazPositionInfo
type TLazPositionMethod = int32

const (
	PmGPS TLazPositionMethod = iota
	PmNetwork
)

// TLazReaderDIBEncoding ENUM
//
//	(*) Note: when reading images with an alpha channel and the alpha channel
//	has no influence on the mask (unless the maskcolor is transparent)
type TLazReaderDIBEncoding = int32

const (
	LrdeRGB TLazReaderDIBEncoding = iota
	LrdeRLE
	LrdeBitfield
	LrdeJpeg
	LrdePng
	LrdeHuffman
)

// TLazReaderMaskMode ENUM
//
//	TLazReaderDIB
//	This is an imroved FPImage reader for dib images.
type TLazReaderMaskMode = int32

const (
	LrmmNone TLazReaderMaskMode = iota
	LrmmAuto
	LrmmColor
)

// TLazRegionFillMode ENUM
type TLazRegionFillMode = int32

const (
	RfmOddEven TLazRegionFillMode = iota
	RfmWinding
)

// TLazSynBorderSide ENUM
type TLazSynBorderSide = int32

const (
	BsLeft TLazSynBorderSide = iota
	BsTop
	BsRight
	BsBottom
)

// TLazSynWordBoundary ENUM
type TLazSynWordBoundary = int32

const (
	SwbWordBegin TLazSynWordBoundary = iota
	SwbWordEnd
	SwbTokenBegin
	SwbTokenEnd
	SwbCaseChange
	SwbWordSmart
)

// TLCLCapability ENUM
type TLCLCapability = int32

const (
	LcAsyncProcess TLCLCapability = iota
	LcCanDrawOutsideOnPaint
	LcNeedMininimizeAppWithMainForm
	LcApplicationTitle
	LcApplicationWindow
	LcFormIcon
	LcModalWindow
	LcDragDockStartOnTitleClick
	LcAntialiasingEnabledByDefault
	LcLMHelpSupport
	LcReceivesLMClearCutCopyPasteReliably
	LcSendsUTF8KeyPress
	LcAllowChildControlsInNativeControls
	LcEmulatedMDI
	LcAccessibilitySupport
	LcRadialGradientBrush
	LcTransparentWindow
	LcTextHint
	LcNativeTaskDialog
	LcCanDrawHidden
	LcAccelleratorKeys
)

// TLCLGlyphsMissingResources ENUM
type TLCLGlyphsMissingResources = int32

const (
	GmrAllMustExist TLCLGlyphsMissingResources = iota
	GmrOneMustExist
	GmrIgnoreAll
)

// TLCLPlatform ENUM
type TLCLPlatform = int32

const (
	LpGtk TLCLPlatform = iota
	LpGtk2
	LpGtk3
	LpWin32
	LpWinCE
	LpCarbon
	LpQT
	LpQt5
	LpQt6
	LpfpGUI
	LpNoGUI
	LpCocoa
	LpCustomDrawn
	LpMUI
)

// TLCLTaskDialogFooterIcon ENUM
type TLCLTaskDialogFooterIcon = int32

const (
	TfiBlank1 TLCLTaskDialogFooterIcon = iota
	TfiWarning1
	TfiQuestion1
	TfiError1
	TfiInformation1
	TfiShield1
)

// TLCLTaskDialogIcon ENUM
type TLCLTaskDialogIcon = int32

const (
	TiBlank1 TLCLTaskDialogIcon = iota
	TiWarning1
	TiQuestion1
	TiError1
	TiInformation1
	TiNotUsed1
	TiShield1
)

// TLinkAction ENUM
type TLinkAction = int32

const (
	LaClick TLinkAction = iota
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

// TListItemFlag ENUM
type TListItemFlag = int32

const (
	LifDestroying TListItemFlag = iota
	LifCreated
)

// TListNotification ENUM
//
//	TList class
type TListNotification = int32

const (
	LnAdded TListNotification = iota
	LnExtracted
	LnDeleted
)

// TListViewFlag ENUM
type TListViewFlag = int32

const (
	LffSelectedValid TListViewFlag = iota
	LffItemsMoving
	LffItemsSorting
	LffPreparingSorting
)

// TListViewImageList ENUM
type TListViewImageList = int32

const (
	LvilSmall TListViewImageList = iota
	LvilLarge
	LvilState
)

// TListViewProperty ENUM
type TListViewProperty = int32

const (
	LvpAutoArrange TListViewProperty = iota
	LvpCheckboxes
	LvpColumnClick
	LvpFlatScrollBars
	LvpFullDrag
	LvpGridLines
	LvpHideSelection
	LvpHotTrack
	LvpMultiSelect
	LvpOwnerDraw
	LvpReadOnly
	LvpRowSelect
	LvpShowColumnHeaders
	LvpShowWorkAreas
	LvpWrapText
	LvpToolTips
)

// TLRSItemType ENUM
type TLRSItemType = int32

const (
	LrsitCollection TLRSItemType = iota
	LrsitComponent
	LrsitList
	LrsitProperty
)

// TLRSOWStackItemState ENUM
//
//	TLRSOWStackItem
//	The TLRSObjectWriter can find empty entries and omit writing them to stream.
//	For example:
//	inline ConditionalOptionsFrame: TCompOptsConditionalsFrame
//	inherited COCTreeView: TTreeView
//	end
//	inherited COCPopupMenu: TPopupMenu
//	end
//	end
//	The empty inherited child components will not be written if
//	WriteEmptyInheritedChilds = false (default).
//	Reason:
//	This allows one to delete/rename controls in ancestors without the need
//	to update all descendants.
type TLRSOWStackItemState = int32

const (
	LrsowsisStarted TLRSOWStackItemState = iota
	LrsowsisHeaderWritten
	LrsowsisDataWritten
)

// TLRSStreamOriginalFormat ENUM
type TLRSStreamOriginalFormat = int32

const (
	SofUnknown TLRSStreamOriginalFormat = iota
	SofBinary
	SofText
)

// TLVStyleType ENUM
type TLVStyleType = int32

const (
	LsStyle TLVStyleType = iota
	LsInvert
	LsExStyle
	LsNone
)

// TMaskCaseSensitivity ENUM
type TMaskCaseSensitivity = int32

const (
	McsPlatformDefault TMaskCaseSensitivity = iota
	McsCaseInsensitive
	McsCaseSensitive
)

// TMaskeditTrimType ENUM
type TMaskeditTrimType = int32

const (
	MetTrimLeft TMaskeditTrimType = iota
	MetTrimRight
)

// TMaskEditValidationErrorMode ENUM
type TMaskEditValidationErrorMode = int32

const (
	MvemException TMaskEditValidationErrorMode = iota
	MvemEvent
)

// tMaskedType ENUM
//
//	Type for mask (internal)
//	When adding more: make sure to add them in procedure InitcMaskToMaskedTypeArray if appropriate
type tMaskedType = int32

const (
	Char_Invalid tMaskedType = iota
	Char_IsLiteral
	Char_Number
	Char_NumberFixed
	Char_NumberPlusMin
	Char_Letter
	Char_LetterFixed
	Char_LetterUpCase
	Char_LetterDownCase
	Char_LetterFixedUpCase
	Char_LetterFixedDownCase
	Char_AlphaNum
	Char_AlphaNumFixed
	Char_AlphaNumUpCase
	Char_AlphaNumDownCase
	Char_AlphaNumFixedUpCase
	Char_AlphaNumFixedDownCase
	Char_All
	Char_AllFixed
	Char_AllUpCase
	Char_AllDownCase
	Char_AllFixedUpCase
	Char_AllFixedDownCase
	Char_HourSeparator
	Char_DateSeparator
	Char_Hex
	Char_HexFixed
	Char_HexUpCase
	Char_HexDownCase
	Char_HexFixedUpCase
	Char_HexFixedDownCase
	Char_Binary
	Char_BinaryFixed
	Char_Set
	Char_SetFixed
	Char_SetNegateFixed
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

// TNativeCanvasType ENUM
type TNativeCanvasType = int32

const (
	NctWindowsDC TNativeCanvasType = iota
	NctLazCanvas
)

// TNativeHandleType ENUM
//
//	Types for native Handle integration
type TNativeHandleType = int32

const (
	NhtWindowsHWND TNativeHandleType = iota
	NhtX11TWindow
	NhtCocoaNSWindow
	NhtQtQWidget
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

// TNotifierXButtonButtonState ENUM
//
//	TNotifierXButton
//	To avoid dependency on Buttons
type TNotifierXButtonButtonState = int32

const (
	NbsUp TNotifierXButtonButtonState = iota
	NbsDown
	NbsHot
)

// TNumberingStyle ENUM
type TNumberingStyle = int32

const (
	NsNone TNumberingStyle = iota
	NsBullet
)

// TObjectTextEncoding ENUM
type TObjectTextEncoding = int32

const (
	OteDFM TObjectTextEncoding = iota
	OteLFM
)

// TObjectType ENUM
type TObjectType = int32

const (
	OtFolders TObjectType = iota
	OtNonFolders
	OtHidden
)

// TOnBeforeExeucteFlag ENUM
type TOnBeforeExeucteFlag = int32

const (
	BefAbort TOnBeforeExeucteFlag = iota
)

// TOperation ENUM
//
//	TComponent class
type TOperation = int32

const (
	OpInsert TOperation = iota
	OpRemove
)

// TPageFlag ENUM
//
//	TCustomPage
type TPageFlag = int32

const (
	PfAdded TPageFlag = iota
	PfAdding
	PfRemoving
	PfInserting
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

// TParaAlignment ENUM
type TParaAlignment = int32

const (
	PaLeft1 TParaAlignment = iota
	PaRight1
	PaCenter1
	PaJustify1
)

// TParaNumStyle ENUM
type TParaNumStyle = int32

const (
	PnNone TParaNumStyle = iota
	PnBullet
	PnNumber
	PnLowLetter
	PnLowRoman
	PnUpLetter
	PnUpRoman
	PnCustomChar
)

// TPenEndCap ENUM
type TPenEndCap = int32

const (
	PecRound TPenEndCap = iota
	PecSquare
	PecFlat
)

// TPenJoinStyle ENUM
type TPenJoinStyle = int32

const (
	PjsRound TPenJoinStyle = iota
	PjsBevel
	PjsMiter
)

// TPipeReason ENUM
type TPipeReason = int32

const (
	PrDataAvailable TPipeReason = iota
	PrBroken
	PrCanWrite
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

// TPlacementOperation ENUM
type TPlacementOperation = int32

const (
	PoSave TPlacementOperation = iota
	PoRestore
)

// TPluginState ENUM
type TPluginState = int32

const (
	PsNone TPluginState = iota
	PsExecuting
	PsAccepting
	PsCancelling
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

// TPredefinedClipboardFormat ENUM
type TPredefinedClipboardFormat = int32

const (
	PcfText TPredefinedClipboardFormat = iota
	PcfBitmap
	PcfPixmap
	PcfIcon
	PcfPicture
	PcfMetaFilePict
	PcfObject
	PcfComponent
	PcfCustomData
)

// TPrefixOption ENUM
type TPrefixOption = int32

const (
	PoNone TPrefixOption = iota
	PoHeaderClick
)

// TPrintAction ENUM
type TPrintAction = int32

const (
	PaDocStart TPrintAction = iota
	PaPageStart
	PaPageEnd
	PaDocEnd
)

// TPrinterCapability ENUM
type TPrinterCapability = int32

const (
	PcCopies TPrinterCapability = iota
	PcOrientation
	PcCollation
)

// TPrinterFlags ENUM
type TPrinterFlags = int32

const (
	PfPrinting TPrinterFlags = iota
	PfAborted
	PfDestroying
	PfPrintersValid
	PfRawMode
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

// TPsCanvasState ENUM
type TPsCanvasState = int32

const (
	PcsPosValid TPsCanvasState = iota
	PcsClipping
	PcsClipSaved
)

// TPSPaintType ENUM
type TPSPaintType = int32

const (
	PtColored TPSPaintType = iota
	PtUncolored
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

// TRegionCombineMode ENUM
type TRegionCombineMode = int32

const (
	RgnAnd TRegionCombineMode = iota
	RgnCopy
	RgnDiff
	RgnOr
	RgnXOR
)

// TRegionOperationType ENUM
type TRegionOperationType = int32

const (
	RgnNewRect TRegionOperationType = iota
	RgnCombine
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

// TRubberBandShape ENUM
type TRubberBandShape = int32

const (
	RbsLine TRubberBandShape = iota
	RbsRectangle
)

// TRuntimeError ENUM
type TRuntimeError = int32

const (
	ReIntOverflow TRuntimeError = iota
	ReInvalidOp
	ReZeroDivide
	ReOverflow
	ReUnderflow
	ReInvalidCast
	ReAccessViolation
	RePrivInstruction
	ReControlBreak
	ReStackOverflow
	ReVarTypeCast
	ReVarInvalidOp
	ReVarDispatch
	ReVarArrayCreate
	ReVarNotArray
	ReVarArrayBounds
	ReAssertionFailed
	ReExternalException
	ReIntfCastError
	ReSafeCallError
	ReQuit
	ReCodesetConversion
	ReNoDynLibsSupport
	ReThreadError
)

// TScreenNotification ENUM
type TScreenNotification = int32

const (
	SnNewFormCreated TScreenNotification = iota
	SnFormAdded
	SnRemoveForm
	SnActiveControlChanged
	SnActiveFormChanged
	SnFormVisibleChanged
)

// TScreenRotation ENUM
//
//	TLazDevice
type TScreenRotation = int32

const (
	SrRotation_0 TScreenRotation = iota
	SrRotation_90
	SrRotation_180
	SrRotation_270
)

// TScrollBarKind ENUM
//
//	TControlScrollBar
type TScrollBarKind = int32

const (
	SbHorizontal TScrollBarKind = iota
	SbVertical
)

// TScrollBarStyle ENUM
type TScrollBarStyle = int32

const (
	SsRegular TScrollBarStyle = iota
	SsFlat
	SsHotTrack
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

// TSelectDirOpt ENUM
type TSelectDirOpt = int32

const (
	SdAllowCreate TSelectDirOpt = iota
	SdPerformCreate
	SdPrompt
)

// TShadowType ENUM
type TShadowType = int32

const (
	StNone TShadowType = iota
	StIn
	StOut
	StEtchedIn
	StEtchedOut
	StFilled
)

// TShapeDirection ENUM
type TShapeDirection = int32

const (
	AtUp1 TShapeDirection = iota
	AtDown1
	AtLeft1
	AtRight1
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

// TShowHelpResult ENUM
type TShowHelpResult = int32

const (
	ShrNone TShowHelpResult = iota
	ShrSuccess
	ShrCancel
	ShrDatabaseNotFound
	ShrContextNotFound
	ShrViewerNotFound
	ShrHelpNotFound
	ShrViewerError
	ShrSelectorError
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

// TSortDirection ENUM
type TSortDirection = int32

const (
	SdAscending TSortDirection = iota
	SdDescending
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
	StNone1 TSortType = iota
	StData1
	StText1
	StBoth1
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
	PbNone1 TStatusPanelBevel = iota
	PbLowered1
	PbRaised1
)

// TStatusPanelStyle ENUM
type TStatusPanelStyle = int32

const (
	PsText TStatusPanelStyle = iota
	PsOwnerDraw
)

// TStockFont ENUM
//
//	enum to use with InitStockFont
type TStockFont = int32

const (
	SfSystem TStockFont = iota
	SfHint
	SfIcon
	SfMenu
)

// TStreamOwnership ENUM
//
//	TStreamAdapter
type TStreamOwnership = int32

const (
	SoReference TStreamOwnership = iota
	SoOwned
)

// TStringsSortStyle ENUM
type TStringsSortStyle = int32

const (
	SslNone TStringsSortStyle = iota
	SslUser
	SslAuto
)

// TSubItemUpdate ENUM
type TSubItemUpdate = int32

const (
	SiuText TSubItemUpdate = iota
	SiuImage
)

// TSynBeautifierIndentType ENUM
type TSynBeautifierIndentType = int32

const (
	SbitSpace TSynBeautifierIndentType = iota
	SbitCopySpaceTab
	SbitPositionCaret
	SbitConvertToTabSpace
	SbitConvertToTabOnly
)

// TSynBlockPersistMode ENUM
type TSynBlockPersistMode = int32

const (
	SbpDefault TSynBlockPersistMode = iota
	SbpWeak
	SbpStrong
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

// TSynCommentContineMode ENUM
type TSynCommentContineMode = int32

const (
	SccNoPrefix TSynCommentContineMode = iota
	SccPrefixAlways
	SccPrefixMatch
)

// TSynCommentExtendMode ENUM
type TSynCommentExtendMode = int32

const (
	SceNever TSynCommentExtendMode = iota
	SceAlways
	SceSplitLine
	SceMatching
	SceMatchingSplitLine
)

// TSynCommentIndentFlag ENUM
type TSynCommentIndentFlag = int32

const (
	SciNone TSynCommentIndentFlag = iota
	SciAlignOpen
	SciAlignOpenOnce
	SciAlignOpenSkipBOL
	SciAddTokenLen
	SciAddPastTokenIndent
	SciMatchOnlyTokenLen
	SciMatchOnlyPastTokenIndent
	SciAlignOnlyTokenLen
	SciAlignOnlyPastTokenIndent
	SciApplyIndentForNoMatch
)

// TSynCommentMatchLine ENUM
type TSynCommentMatchLine = int32

const (
	SclMatchFirst TSynCommentMatchLine = iota
	SclMatchPrev
)

// TSynCommentMatchMode ENUM
type TSynCommentMatchMode = int32

const (
	ScmMatchAfterOpening TSynCommentMatchMode = iota
	ScmMatchOpening
	ScmMatchWholeLine
	ScmMatchAtAsterisk
)

// TSynCommentType ENUM
type TSynCommentType = int32

const (
	SctAnsi TSynCommentType = iota
	SctBor
	SctSlash
)

// TSynCompletionLongHintType ENUM
type TSynCompletionLongHintType = int32

const (
	SclpNone TSynCompletionLongHintType = iota
	SclpExtendRightOnly
	SclpExtendHalfLeft
	SclpExtendUnlimitedLeft
)

// TSynCoordinateMappingFlag ENUM
type TSynCoordinateMappingFlag = int32

const (
	ScmLimitToLines TSynCoordinateMappingFlag = iota
	ScmIncludePartVisible
	ScmForceLeftSidePos
)

// TSynCopyPasteAction ENUM
type TSynCopyPasteAction = int32

const (
	ScaContinue TSynCopyPasteAction = iota
	ScaPlainText
	ScaAbort
)

// TSynCustomCaretSizeFlag ENUM
//
//	relative dimensions in percent from 0 to 1024 (=100%)
type TSynCustomCaretSizeFlag = int32

const (
	CcsRelativeLeft TSynCustomCaretSizeFlag = iota
	CcsRelativeTop
	CcsRelativeWidth
	CcsRelativeHeight
)

// TSynEditBracketHighlightStyle ENUM
type TSynEditBracketHighlightStyle = int32

const (
	SbhsLeftOfCursor TSynEditBracketHighlightStyle = iota
	SbhsRightOfCursor
	SbhsBoth
)

// TSynEditCaretFlag ENUM
//
//	TSynEditCaret
type TSynEditCaretFlag = int32

const (
	ScCharPosValid TSynEditCaretFlag = iota
	ScBytePosValid
	ScViewedPosValid
	ScHasLineMapHandler
	ScfUpdateLastCaretX
)

// TSynEditCaretUpdateFlag ENUM
type TSynEditCaretUpdateFlag = int32

const (
	ScuForceSet TSynEditCaretUpdateFlag = iota
	ScuChangedX
	ScuChangedY
	ScuNoInvalidate
)

// TSynEditHasTextFlag ENUM
type TSynEditHasTextFlag = int32

const (
	ShtIncludeVirtual TSynEditHasTextFlag = iota
)

// TSynEditMarkChangeReason ENUM
type TSynEditMarkChangeReason = int32

const (
	SmcrLine TSynEditMarkChangeReason = iota
	SmcrColumn
	SmcrVisible
	SmcrChanged
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

// TSynEditorShareOption ENUM
//
//	options for textbuffersharing
type TSynEditorShareOption = int32

const (
	EosShareMarks TSynEditorShareOption = iota
)

// TSynEditStringTrimmingType ENUM
type TSynEditStringTrimmingType = int32

const (
	SettLeaveLine TSynEditStringTrimmingType = iota
	SettEditLine
	SettMoveCaret
	SettIgnoreAll
)

// TSynEditTextFlag ENUM
type TSynEditTextFlag = int32

const (
	SetSelect TSynEditTextFlag = iota
	SetPersistentBlock
	SetMoveBlock
	SetExtendBlock
)

// TSynFrameEdges ENUM
type TSynFrameEdges = int32

const (
	SfeNone TSynFrameEdges = iota
	SfeAround
	SfeBottom
	SfeLeft
)

// TSynHighlighterCapability ENUM
type TSynHighlighterCapability = int32

const (
	HcUserSettings TSynHighlighterCapability = iota
	HcRegistry
	HcCodeFolding
)

// TSynHomeMode ENUM
type TSynHomeMode = int32

const (
	SynhmDefault TSynHomeMode = iota
	SynhmFirstWord
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

// TSynLogPhysFlag ENUM
type TSynLogPhysFlag = int32

const (
	LpfAdjustToCharBegin TSynLogPhysFlag = iota
	LpfAdjustToNextChar
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

// TSynMarksAdjustMode ENUM
//   - This is used, if text is *replaced*.
//     What to do with marks in text that is deleted/replaced
type TSynMarksAdjustMode = int32

const (
	SmaMoveUp TSynMarksAdjustMode = iota
	SmaKeep
)

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

// TSynPaintEvent ENUM
type TSynPaintEvent = int32

const (
	PeBeforePaint TSynPaintEvent = iota
	PeAfterPaint
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

// TSynScrollEvent ENUM
type TSynScrollEvent = int32

const (
	PeBeforeScroll TSynScrollEvent = iota
	PeAfterScroll
	PeAfterScrollFailed
)

// TSynSearchOption ENUM
type TSynSearchOption = int32

const (
	SsoBackwards TSynSearchOption = iota
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

// TSynSelectedColorEnum ENUM
type TSynSelectedColorEnum = int32

const (
	SscBack TSynSelectedColorEnum = iota
	SscFore
	SscFrameLeft
	SscFrameRight
	SscFrameTop
	SscFrameBottom
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

// TSynStateFlag ENUM
type TSynStateFlag = int32

const (
	SfCaretChanged TSynStateFlag = iota
	SfHideCursor
	SfEnsureCursorPos
	SfEnsureCursorPosAtResize
	SfEnsureCursorPosForEditRight
	SfEnsureCursorPosForEditLeft
	SfExplicitTopLine
	SfExplicitLeftChar
	SfPreventScrollAfterSelect
	SfIgnoreNextChar
	SfPainting
	SfHasPainted
	SfHasScrolled
	SfScrollbarChanged
	SfHorizScrollbarVisible
	SfVertScrollbarVisible
	SfGutterResized
	SfAfterLoadFromFileNeeded
	SfAfterHandleCreatedNeeded
	SfLeftGutterClick
	SfRightGutterClick
	SfInClick
	SfDblClicked
	SfTripleClicked
	SfQuadClicked
	SfWaitForDragging
	SfWaitForDraggingNoCaret
	SfIsDragging
	SfDraggingOver
	SfWaitForMouseSelecting
	SfMouseSelecting
	SfMouseDoneSelecting
	SfIgnoreUpClick
	SfSelChanged
)

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

// TSynVisibleSpecialChar ENUM
type TSynVisibleSpecialChar = int32

const (
	VscSpace TSynVisibleSpecialChar = iota
	VscTabAtFirst
	VscTabAtLast
)

// TSysLinkType ENUM
type TSysLinkType = int32

const (
	SltURL TSysLinkType = iota
	SltID
)

// TTabAlignment ENUM
type TTabAlignment = int32

const (
	TabLeft TTabAlignment = iota
	TabCenter
	TabRight
	TabDecimal
	TabWordBar
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

// TTaskDialogElement ENUM
type TTaskDialogElement = int32

const (
	TdeContent TTaskDialogElement = iota
	TdeExpandedInfo
	TdeFooter
	TdeMainInstruction
	TdeEdit
	TdeVerif
)

// TTaskDialogFooterIcon ENUM
//
//	the available footer icons for the Task Dialog
type TTaskDialogFooterIcon = int32

const (
	TfiBlank TTaskDialogFooterIcon = iota
	TfiWarning
	TfiQuestion
	TfiError
	TfiInformation
	TfiShield
)

// TTaskDialogIcon ENUM
//
//	the available main icons for the Task Dialog
type TTaskDialogIcon = int32

const (
	TiBlank TTaskDialogIcon = iota
	TiWarning
	TiQuestion
	TiError
	TiInformation
	TiNotUsed
	TiShield
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

// TTextUIFeature ENUM
type TTextUIFeature = int32

const (
	UiLink TTextUIFeature = iota
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

// TThreadPriority ENUM
type TThreadPriority = int32

const (
	TpIdle TThreadPriority = iota
	TpLowest
	TpLower
	TpNormal
	TpHigher
	TpHighest
	TpTimeCritical
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

// TTiffUnit ENUM
//
//	TTiffImage
type TTiffUnit = int32

const (
	TuUnknown TTiffUnit = iota
	TuNone
	TuInch
	TuCentimeter
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

// TToolBarFlag ENUM
type TToolBarFlag = int32

const (
	TbfUpdateVisibleBarNeeded TToolBarFlag = iota
	TbfPlacingControls
)

// TToolButtonFlag ENUM
type TToolButtonFlag = int32

const (
	TbfPressed TToolButtonFlag = iota
	TbfArrowPressed
	TbfMouseInArrow
	TbfDropDownMenuShown
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

// TTreeViewState ENUM
//
//	TCustomTreeView
type TTreeViewState = int32

const (
	TvsScrollbarChanged TTreeViewState = iota
	TvsMaxRightNeedsUpdate
	TvsTopsNeedsUpdate
	TvsMaxLvlNeedsUpdate
	TvsTopItemNeedsUpdate
	TvsBottomItemNeedsUpdate
	TvsCanvasChanged
	TvsDragged
	TvsIsEditing
	TvsStateChanging
	TvsManualNotify
	TvsUpdating
	TvsPainting
	TvoFocusedPainting
	TvsDblClicked
	TvsTripleClicked
	TvsQuadClicked
	TvsSelectionChanged
	TvsEditOnMouseUp
	TvsSingleSelectOnMouseUp
)

// TTriPts ENUM
type TTriPts = int32

const (
	PtA TTriPts = iota
	PtB
	PtC
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

// TValueType ENUM
//
//	TFiler
type TValueType = int32

const (
	VaNull TValueType = iota
	VaList
	VaInt8
	VaInt16
	VaInt32
	VaExtended
	VaString
	VaIdent
	VaFalse
	VaTrue
	VaBinary
	VaSet
	VaLString
	VaNil
	VaCollection
	VaSingle
	VaCurrency
	VaDate
	VaWString
	VaInt64
	VaUTF8String
	VaUString
	VaQWord
)

// TVerticalAlignment ENUM
type TVerticalAlignment = int32

const (
	TaAlignTop TVerticalAlignment = iota
	TaAlignBottom
	TaVerticalCenter
)

// TViewedXYInfoFlag ENUM
type TViewedXYInfoFlag = int32

const (
	VifAdjustLogXYToNextChar TViewedXYInfoFlag = iota
	VifReturnPhysXY
	VifReturnLogXY
	VifReturnLogEOL
	VifReturnPhysOffset
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

// TVScriptPos ENUM
type TVScriptPos = int32

const (
	VpNormal TVScriptPos = iota
	VpSubScript
	VpSuperScript
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

// TVTCellContentMarginType ENUM
//
//	Determines which sides of the cell content margin should be considered.
type TVTCellContentMarginType = int32

const (
	CcmtAllSides TVTCellContentMarginType = iota
	CcmtTopLeftOnly
	CcmtBottomRightOnly
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
	BlGlyphLeft1 TVTHeaderColumnLayout = iota
	BlGlyphRight1
	BlGlyphTop1
	BlGlyphBottom1
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

// TVTImageInfoIndex ENUM
type TVTImageInfoIndex = int32

const (
	IiNormal TVTImageInfoIndex = iota
	IiState
	IiCheck
	IiOverlay
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

// TVTLineType ENUM
//
//	TVTLineType is used during painting a tree
type TVTLineType = int32

const (
	LtNone TVTLineType = iota
	LtBottomRight
	LtTopDown
	LtTopDownRight
	LtRight
	LtTopRight
	LtLeft
	LtLeftBottom
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

// TVZVirtualNodeEnumerationMode ENUM
type TVZVirtualNodeEnumerationMode = int32

const (
	VneAll TVZVirtualNodeEnumerationMode = iota
	VneChecked
	VneChild
	VneCutCopy
	VneInitialized
	VneLeaf
	VneLevel
	VneNoInit
	VneSelected
	VneVisible
	VneVisibleChild
	VneVisibleNoInitChild
	VneVisibleNoInit
)

// TWinCETitlePolicy ENUM
//
//	Policy for using the "OK" close button in the title instead of
//	the default "X" minimize button
type TWinCETitlePolicy = int32

const (
	TpAlwaysUseOKButton TWinCETitlePolicy = iota
	TpOKButtonOnlyOnDialogs
	TpControlWithBorderIcons
)

// TWinCEVersion ENUM
type TWinCEVersion = int32

const (
	Wince_1 TWinCEVersion = iota
	Wince_2
	Wince_3
	Wince_4
	Wince_5
	Wince_6
	Wince_6_1
	Wince_6_5
	Wince_7
	Wince_other
)

// TWinControlFlag ENUM
//
//	TWinControl
type TWinControlFlag = int32

const (
	WcfClientRectNeedsUpdate TWinControlFlag = iota
	WcfColorChanged
	WcfFontChanged
	WcfAllAutoSizing
	WcfAligningControls
	WcfEraseBackground
	WcfCreatingHandle
	WcfInitializing
	WcfCreatingChildHandles
	WcfRealizingBounds
	WcfBoundsRealized
	WcfUpdateShowing
	WcfHandleVisible
	WcfAdjustedLogicalClientRectValid
	WcfKillIntfSetBounds
	WcfDesignerDeleting
	WcfSpecialSubControl
)

// TWindowState ENUM
type TWindowState = int32

const (
	WsNormal TWindowState = iota
	WsMinimized
	WsMaximized
	WsFullScreen
)

// TWindowsVersion ENUM
type TWindowsVersion = int32

const (
	WvUnknown TWindowsVersion = iota
	Wv95
	WvNT4
	Wv98
	WvMe
	Wv2000
	WvXP
	WvServer2003
	WvVista
	Wv7
	Wv8
	Wv8_1
	Wv10
	Wv11
	WvLater
)

// TWrapAfter ENUM
type TWrapAfter = int32

const (
	WaAuto TWrapAfter = iota
	WaForce
	WaAvoid
	WaForbid
)

// TWSListViewItemChange ENUM
//
//	TWSCustomListView
type TWSListViewItemChange = int32

const (
	LvicText TWSListViewItemChange = iota
	LvicImage
)

// TWSZPosition ENUM
//
//	TWSWinControl
type TWSZPosition = int32

const (
	WszpBack TWSZPosition = iota
	WszpFront
)

// TXButtonState ENUM
type TXButtonState = int32

const (
	XbsNone TXButtonState = iota
	XbsHot
	XbsDown
	XbsDisabled
)

// TXPMRange ENUM
type TXPMRange = int32

const (
	XrCode TXPMRange = iota
	XrStaticKeyWord
	XrCharKeyWord
)
