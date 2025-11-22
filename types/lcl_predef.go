//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

// TSelectDirOpts SET: TSelectDirOpt
type TSelectDirOpts = TSet

// TShortCut Low(Word)..High(Word)
type TShortCut = uint16

// TMsgDlgBtn ENUM
type TMsgDlgBtn = int32

const (
	MbYes TMsgDlgBtn = iota
	MbNo
	MbOK
	MbCancel
	MbAbort
	MbRetry
	MbIgnore
	MbAll
	MbNoToAll
	MbYesToAll
	MbHelp
	MbClose
)

// TMsgDlgButtons TMsgDlgBtn SET
type TMsgDlgButtons = TSet

// TMsgDlgType ENUM
type TMsgDlgType = int32

const (
	MtWarning TMsgDlgType = iota
	MtError
	MtInformation
	MtConfirmation
	MtCustom
)

// THelpContext ENUM
type THelpContext = int32

// TFillStyle ENUM
type TFillStyle = TGraphicsFillStyle

// TMenuAutoFlag maAutomatic..maManual
type TMenuAutoFlag = TMenuItemAutoFlag

// TEventLogTypes SET TEventType
type TEventLogTypes = TSet

// TSearchType ENUM
type TSearchType = int32

const (
	StWholeWord TSearchType = iota
	StMatchCase
)

// TSearchTypes SET: TSearchType
type TSearchTypes = TSet

// TIniFileOptions SET TIniFileOption
type TIniFileOptions = TSet

// TAutoScrollInterval ENUM
//
//	Limits the speed interval which can be used for auto scrolling (milliseconds).
//	1..1000
type TAutoScrollInterval = int32

// TVTScrollIncrement ENUM
//
//	1..10000
type TVTScrollIncrement = int32

// TLeftRight = TAlignment.taLeftJustify..TAlignment.taRightJustify;
type TLeftRight = int32

// TDateTimeParts SET: TDateTimePart
type TDateTimeParts = TSet

// TScrollBarInc = 1..32767;
type TScrollBarInc int32

// TBorderStyle ENUM
type TBorderStyle = TFormBorderStyle

// TVTBias ENUM -128..127
type TVTBias = int8

// TVTTransparency ENUM  0..255
// Drag image support for the tree.
type TVTTransparency = uint8

// TJPEGScale ENUM
type TJPEGScale = int32

const (
	JsFullSize TJPEGScale = iota
	JsHalf
	JsQuarter
	JsEighth
)

// IdButton ENUM Stock Pixmap Types
type IdButton = int32

const (
	IdButtonBase     IdButton = 0
	IdButtonOk                = IdButtonBase + 1
	IdButtonCancel            = IdButtonBase + 2
	IdButtonHelp              = IdButtonBase + 3
	IdButtonYes               = IdButtonBase + 4
	IdButtonNo                = IdButtonBase + 5
	IdButtonClose             = IdButtonBase + 6
	IdButtonAbort             = IdButtonBase + 7
	IdButtonRetry             = IdButtonBase + 8
	IdButtonIgnore            = IdButtonBase + 9
	IdButtonAll               = IdButtonBase + 10
	IdButtonYesToAll          = IdButtonBase + 11
	IdButtonNoToAll           = IdButtonBase + 12
	IdButtonOpen              = IdButtonBase + 13
	IdButtonSave              = IdButtonBase + 14
	IdButtonShield            = IdButtonBase + 15
	IdButtonContinue          = IdButtonBase + 16
	IdButtonTryAgain          = IdButtonBase + 17
)

// TButtonImage ENUM IdButtonOk.. IdButtonNoToAll
type TButtonImage = IdButton

// TListItemsSortType ENUM
type TListItemsSortType = TSortType

// TPenStyle ENUM
type TPenStyle = TFPPenStyle

// TVTConstraintPercent ENUM 0..100
type TVTConstraintPercent = uint8

// TBevelCut ENUM
type TBevelCut = TGraphicsBevelCut

// TPanelBevel ENUM
type TPanelBevel = TBevelCut

// TBevelWidth ENUM
type TBevelWidth = int32

// TJPEGPerformance ENUM
type TJPEGPerformance = int32

const (
	JpBestQuality TJPEGPerformance = iota
	JpBestSpeed
)

// TPanelButton ENUM = PbOK .. PbHelp
type TPanelButton = TPanelButtonEx

// TPanelButtons SET: TPanelButton
type TPanelButtons = TSet

// TRowSize 1..MaxInt
type TRowSize = int32

// TSynIdentChars = set of char
type TSynIdentChars = TSet

type TSynEditorCommand = Word

type TSynEditCaretType = TSynCaretType
