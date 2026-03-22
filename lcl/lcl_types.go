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
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

type TAlignInfo struct {
	AlignList    IFPList      // TFPList
	ControlIndex int32        // Integer
	Align        types.TAlign // TAlign
	Scratch      int32        // Integer
}

type TControlBorderSpacingDefault struct {
	Left   types.TSpacingSize // TSpacingSize
	Top    types.TSpacingSize // TSpacingSize
	Right  types.TSpacingSize // TSpacingSize
	Bottom types.TSpacingSize // TSpacingSize
	Around types.TSpacingSize // TSpacingSize
}

type TFontData struct {
	Handle      types.HFont           // HFont
	Height      int32                 // Integer
	Pitch       types.TFontPitch      // TFontPitch
	Style       types.TFontStylesBase // TFontStylesBase
	CharSet     types.TFontCharSet    // TFontCharSet
	Quality     types.TFontQuality    // TFontQuality
	Name        string                // TFontDataName
	Orientation int32                 // Integer
}

type TFPColor struct {
	Red   uint16 // word
	Green uint16 // word
	Blue  uint16 // word
	Alpha uint16 // word
}

type THeaderPaintInfo struct {
	TargetCanvas    ICanvas               // TCanvas
	Column          IVirtualTreeColumn    // TVirtualTreeColumn
	PaintRectangle  types.TRect           // TRect
	TextRectangle   types.TRect           // TRect
	IsHoverIndex    types.LongBool        // Boolean
	IsDownIndex     types.LongBool        // Boolean
	IsEnabled       types.LongBool        // Boolean
	ShowHeaderGlyph types.LongBool        // Boolean
	ShowSortGlyph   types.LongBool        // Boolean
	ShowRightBorder types.LongBool        // Boolean
	DropMark        types.TVTDropMarkMode // TVTDropMarkMode
	GlyphPos        types.TPoint          // TPoint
	SortGlyphPos    types.TPoint          // TPoint
}

type THintInfo struct {
	HintControl     IControl               // TControl
	HintWindowClass types.TWinControlClass // TWinControlClass
	HintPos         types.TPoint           // TPoint
	HintMaxWidth    int32                  // Integer
	HintColor       types.TColor           // TColor
	CursorRect      types.TRect            // TRect
	CursorPos       types.TPoint           // TPoint
	ReshowTimeout   int32                  // Integer
	HideTimeout     int32                  // Integer
	HintStr         string                 // string
	HintData        uintptr                // Pointer
}

type THitInfo struct {
	HitNode      types.PVirtualNode  // PVirtualNode
	HitPositions types.THitPositions // THitPositions
	HitColumn    int32               // TColumnIndex
	HitPoint     types.TPoint        // TPoint
}

type TLazSynDisplayRtlInfo struct {
	IsRtl     types.LongBool // Boolean
	PhysLeft  int32          // integer
	PhysRight int32          // integer
	LogFirst  int32          // integer
	LogLast   int32          // integer
}

type TLazSynDisplayTokenBound struct {
	Physical int32 // Integer
	Logical  int32 // Integer
	Offset   int32 // Integer
}

type TLazSynDisplayTokenInfo struct {
	TokenStart  uintptr                   // PChar
	TokenLength int32                     // integer
	TokenAttr   ISynHighlighterAttributes // TSynHighlighterAttributes
}

type TLCLTextMetric struct {
	Ascender  int32 // Integer
	Descender int32 // Integer
	Height    int32 // Integer
}

type TLineRange struct {
	Top    types.TLineIdx // TLineIdx
	Bottom types.TLineIdx // TLineIdx
}

type TLogCaretPoint struct {
	X    int32 // Integer
	Y    int32 // Integer
	Offs int32 // Integer
}

type TMethod struct {
	Code uintptr // CodePointer
	Data uintptr // Pointer
}

type TPaperRect struct {
	PhysicalRect types.TRect // TRect
	WorkRect     types.TRect // TRect
}

type TRawImagePosition struct {
	Byte uintptr // PtrUInt
	Bit  uint32  // cardinal
}

type TRegDataInfo struct {
	RegData  types.TRegDataType // TRegDataType
	DataSize int32              // Integer
}

type TRegKeyInfo struct {
	NumSubKeys   int32           // Integer
	MaxSubKeyLen int32           // Integer
	NumValues    int32           // Integer
	MaxValueLen  int32           // Integer
	MaxDataLen   int32           // Integer
	FileTime     types.TDateTime // TDateTime
}

type TRGBAQuad struct {
	Blue  byte // Byte
	Green byte // Byte
	Red   byte // Byte
	Alpha byte // Byte
}

type TStatStg struct {
	PwcsName          string      // string
	DwType            types.DWORD // DWORD
	CbSize            uint64      // uint64
	Mtime             string      // string
	Ctime             string      // string
	Atime             string      // string
	GrfMode           types.DWORD // DWORD
	GrfLocksSupported types.DWORD // DWORD
	Clsid             string      // string
	GrfStateBits      types.DWORD // DWORD
	Reserved          types.DWORD // DWORD
}

type TSynDividerDrawConfigSetting struct {
	Color types.TColor // TColor
}

type TSynEditGutterLineInfo struct {
	ShowDot   types.LongBool // boolean
	ViewLine  int32          // Integer
	LineRange TLineRange     // TLineRange
}

type TSynEditMouseActionInfo struct {
	NewCaret      ISynEditCaret             // TSynEditCaret
	Button        types.TSynMouseButton     // TSynMouseButton
	Shift         types.TShiftState         // TShiftState
	MouseX        int32                     // Integer
	MouseY        int32                     // Integer
	WheelDelta    int32                     // Integer
	CCount        types.TSynMAClickCount    // TSynMAClickCount
	Dir           types.TSynMAClickDir      // TSynMAClickDir
	CaretDone     types.LongBool            // Boolean
	IgnoreUpClick types.LongBool            // Boolean
	ActionResult  TSynEditMouseActionResult // TSynEditMouseActionResult
}

type TSynEditMouseActionResult struct {
	DoPopUpEvent types.LongBool // Boolean
	PopUpEventX  int32          // Integer
	PopUpEventY  int32          // Integer
	PopUpMenu    IPopupMenu     // TPopupMenu
}

type TSynFoldBlockFilter struct {
	FoldGroup int32                          // integer
	Flags     types.TSynFoldBlockFilterFlags // TSynFoldBlockFilterFlags
}

type TSynFoldNodeInfo struct {
	LineIndex          int32                 // Integer
	NodeIndex          int32                 // Integer
	AllNodeIndex       int32                 // Integer
	LogXStart          int32                 // Integer
	LogXEnd            int32                 // Integer
	FoldLvlStart       int32                 // Integer
	FoldLvlEnd         int32                 // Integer
	NestLvlStart       int32                 // Integer
	NestLvlEnd         int32                 // Integer
	FoldAction         types.TSynFoldActions // TSynFoldActions
	FoldType           uintptr               // Pointer
	FoldTypeCompatible uintptr               // Pointer
	FoldGroup          int32                 // Integer
}

type TSynHLightMultiVirtualSection struct {
	StartPos      types.TPoint // TPoint
	EndPos        types.TPoint // TPoint
	TokenStartPos int32        // Integer
	TokenEndPos   int32        // Integer
	VirtualLine   int32        // Integer
}

type TTextStyle struct {
	Alignment   types.TAlignment  // TAlignment
	Layout      types.TTextLayout // TTextLayout
	SingleLine  types.LongBool    // boolean
	Clipping    types.LongBool    // boolean
	ExpandTabs  types.LongBool    // boolean
	ShowPrefix  types.LongBool    // boolean
	Wordbreak   types.LongBool    // boolean
	Opaque      types.LongBool    // boolean
	SystemFont  types.LongBool    // Boolean
	RightToLeft types.LongBool    // Boolean
	EndEllipsis types.LongBool    // Boolean
}

type TThemedElementDetails struct {
	Element types.TThemedElement // TThemedElement
	Part    int32                // Integer
	State   int32                // Integer
}

type TViewedXYInfo struct {
	CorrectedViewedXY types.TRect    // TPhysPoint
	PhysXY            types.TRect    // TPhysPoint
	LogicalXY         TLogCaretPoint // TLogCaretPoint
	PhysBoundOffset   int32          // Integer
	LogEOLPos         int32          // Integer
	FirstViewedX      types.IntPos   // IntPos
}

type TVTHeaderHitInfo struct {
	X           int32                       // Integer
	Y           int32                       // Integer
	Button      types.TMouseButton          // TMouseButton
	Shift       types.TShiftState           // TShiftState
	Column      int32                       // TColumnIndex
	HitPosition types.TVTHeaderHitPositions // TVTHeaderHitPositions
}

type TVTImageInfo struct {
	Index   int32            // Integer
	XPos    int32            // Integer
	YPos    int32            // Integer
	Ghosted types.LongBool   // Boolean
	Images  ICustomImageList // TCustomImageList
}

type TVTPaintInfo struct {
	Canvas           ICanvas                       // TCanvas
	PaintOptions     types.TVTInternalPaintOptions // TVTInternalPaintOptions
	Node             types.PVirtualNode            // PVirtualNode
	Column           int32                         // TColumnIndex
	Position         uint32                        // TColumnPosition
	CellRect         types.TRect                   // TRect
	ContentRect      types.TRect                   // TRect
	NodeWidth        int32                         // Integer
	Alignment        types.TAlignment              // TAlignment
	CaptionAlignment types.TAlignment              // TAlignment
	BidiMode         types.TBiDiMode               // TBidiMode
	BrushOrigin      types.TPoint                  // TPoint
	ImageInfo        uintptr                       // array[TVTImageInfoIndex] of TVTImageInfo
	ImageInfoSize    int32                         // SizeOf: array[TVTImageInfoIndex] of TVTImageInfo
}

func (m *TAlignInfo) ToPas() *tAlignInfo {
	if m == nil {
		return nil
	}
	return &tAlignInfo{
		AlignList:    base.GetObjectUintptr(m.AlignList),
		ControlIndex: m.ControlIndex,
		Align:        m.Align,
		Scratch:      m.Scratch,
	}
}

type tAlignInfo struct {
	AlignList    uintptr      // TFPList
	ControlIndex int32        // Integer
	Align        types.TAlign // TAlign
	Scratch      int32        // Integer
}

func (m *tAlignInfo) ToGo() TAlignInfo {
	if m == nil {
		return TAlignInfo{}
	}
	return TAlignInfo{
		AlignList:    AsFPList(m.AlignList),
		ControlIndex: m.ControlIndex,
		Align:        m.Align,
		Scratch:      m.Scratch,
	}
}
func (m *TFontData) ToPas() *tFontData {
	if m == nil {
		return nil
	}
	return &tFontData{
		Handle:      m.Handle,
		Height:      m.Height,
		Pitch:       m.Pitch,
		Style:       m.Style,
		CharSet:     m.CharSet,
		Quality:     m.Quality,
		Name:        api.PasStr(m.Name),
		Orientation: m.Orientation,
	}
}

type tFontData struct {
	Handle      types.HFont           // HFont
	Height      int32                 // Integer
	Pitch       types.TFontPitch      // TFontPitch
	Style       types.TFontStylesBase // TFontStylesBase
	CharSet     types.TFontCharSet    // TFontCharSet
	Quality     types.TFontQuality    // TFontQuality
	Name        uintptr               // TFontDataName
	Orientation int32                 // Integer
}

func (m *tFontData) ToGo() TFontData {
	if m == nil {
		return TFontData{}
	}
	return TFontData{
		Handle:      m.Handle,
		Height:      m.Height,
		Pitch:       m.Pitch,
		Style:       m.Style,
		CharSet:     m.CharSet,
		Quality:     m.Quality,
		Name:        api.GoStr(m.Name),
		Orientation: m.Orientation,
	}
}
func (m *THeaderPaintInfo) ToPas() *tHeaderPaintInfo {
	if m == nil {
		return nil
	}
	return &tHeaderPaintInfo{
		TargetCanvas:    base.GetObjectUintptr(m.TargetCanvas),
		Column:          base.GetObjectUintptr(m.Column),
		PaintRectangle:  m.PaintRectangle,
		TextRectangle:   m.TextRectangle,
		IsHoverIndex:    m.IsHoverIndex,
		IsDownIndex:     m.IsDownIndex,
		IsEnabled:       m.IsEnabled,
		ShowHeaderGlyph: m.ShowHeaderGlyph,
		ShowSortGlyph:   m.ShowSortGlyph,
		ShowRightBorder: m.ShowRightBorder,
		DropMark:        m.DropMark,
		GlyphPos:        m.GlyphPos,
		SortGlyphPos:    m.SortGlyphPos,
	}
}

type tHeaderPaintInfo struct {
	TargetCanvas    uintptr               // TCanvas
	Column          uintptr               // TVirtualTreeColumn
	PaintRectangle  types.TRect           // TRect
	TextRectangle   types.TRect           // TRect
	IsHoverIndex    types.LongBool        // Boolean
	IsDownIndex     types.LongBool        // Boolean
	IsEnabled       types.LongBool        // Boolean
	ShowHeaderGlyph types.LongBool        // Boolean
	ShowSortGlyph   types.LongBool        // Boolean
	ShowRightBorder types.LongBool        // Boolean
	DropMark        types.TVTDropMarkMode // TVTDropMarkMode
	GlyphPos        types.TPoint          // TPoint
	SortGlyphPos    types.TPoint          // TPoint
}

func (m *tHeaderPaintInfo) ToGo() THeaderPaintInfo {
	if m == nil {
		return THeaderPaintInfo{}
	}
	return THeaderPaintInfo{
		TargetCanvas:    AsCanvas(m.TargetCanvas),
		Column:          AsVirtualTreeColumn(m.Column),
		PaintRectangle:  m.PaintRectangle,
		TextRectangle:   m.TextRectangle,
		IsHoverIndex:    m.IsHoverIndex,
		IsDownIndex:     m.IsDownIndex,
		IsEnabled:       m.IsEnabled,
		ShowHeaderGlyph: m.ShowHeaderGlyph,
		ShowSortGlyph:   m.ShowSortGlyph,
		ShowRightBorder: m.ShowRightBorder,
		DropMark:        m.DropMark,
		GlyphPos:        m.GlyphPos,
		SortGlyphPos:    m.SortGlyphPos,
	}
}
func (m *THintInfo) ToPas() *tHintInfo {
	if m == nil {
		return nil
	}
	return &tHintInfo{
		HintControl:     base.GetObjectUintptr(m.HintControl),
		HintWindowClass: m.HintWindowClass,
		HintPos:         m.HintPos,
		HintMaxWidth:    m.HintMaxWidth,
		HintColor:       m.HintColor,
		CursorRect:      m.CursorRect,
		CursorPos:       m.CursorPos,
		ReshowTimeout:   m.ReshowTimeout,
		HideTimeout:     m.HideTimeout,
		HintStr:         api.PasStr(m.HintStr),
		HintData:        m.HintData,
	}
}

type tHintInfo struct {
	HintControl     uintptr                // TControl
	HintWindowClass types.TWinControlClass // TWinControlClass
	HintPos         types.TPoint           // TPoint
	HintMaxWidth    int32                  // Integer
	HintColor       types.TColor           // TColor
	CursorRect      types.TRect            // TRect
	CursorPos       types.TPoint           // TPoint
	ReshowTimeout   int32                  // Integer
	HideTimeout     int32                  // Integer
	HintStr         uintptr                // string
	HintData        uintptr                // Pointer
}

func (m *tHintInfo) ToGo() THintInfo {
	if m == nil {
		return THintInfo{}
	}
	return THintInfo{
		HintControl:     AsControl(m.HintControl),
		HintWindowClass: m.HintWindowClass,
		HintPos:         m.HintPos,
		HintMaxWidth:    m.HintMaxWidth,
		HintColor:       m.HintColor,
		CursorRect:      m.CursorRect,
		CursorPos:       m.CursorPos,
		ReshowTimeout:   m.ReshowTimeout,
		HideTimeout:     m.HideTimeout,
		HintStr:         api.GoStr(m.HintStr),
		HintData:        m.HintData,
	}
}
func (m *TLazSynDisplayTokenInfo) ToPas() *tLazSynDisplayTokenInfo {
	if m == nil {
		return nil
	}
	return &tLazSynDisplayTokenInfo{
		TokenStart:  m.TokenStart,
		TokenLength: m.TokenLength,
		TokenAttr:   base.GetObjectUintptr(m.TokenAttr),
	}
}

type tLazSynDisplayTokenInfo struct {
	TokenStart  uintptr // PChar
	TokenLength int32   // integer
	TokenAttr   uintptr // TSynHighlighterAttributes
}

func (m *tLazSynDisplayTokenInfo) ToGo() TLazSynDisplayTokenInfo {
	if m == nil {
		return TLazSynDisplayTokenInfo{}
	}
	return TLazSynDisplayTokenInfo{
		TokenStart:  m.TokenStart,
		TokenLength: m.TokenLength,
		TokenAttr:   AsSynHighlighterAttributes(m.TokenAttr),
	}
}
func (m *TStatStg) ToPas() *tStatStg {
	if m == nil {
		return nil
	}
	return &tStatStg{
		PwcsName:          api.PasStr(m.PwcsName),
		DwType:            m.DwType,
		CbSize:            m.CbSize,
		Mtime:             api.PasStr(m.Mtime),
		Ctime:             api.PasStr(m.Ctime),
		Atime:             api.PasStr(m.Atime),
		GrfMode:           m.GrfMode,
		GrfLocksSupported: m.GrfLocksSupported,
		Clsid:             api.PasStr(m.Clsid),
		GrfStateBits:      m.GrfStateBits,
		Reserved:          m.Reserved,
	}
}

type tStatStg struct {
	PwcsName          uintptr     // string
	DwType            types.DWORD // DWORD
	CbSize            uint64      // uint64
	Mtime             uintptr     // string
	Ctime             uintptr     // string
	Atime             uintptr     // string
	GrfMode           types.DWORD // DWORD
	GrfLocksSupported types.DWORD // DWORD
	Clsid             uintptr     // string
	GrfStateBits      types.DWORD // DWORD
	Reserved          types.DWORD // DWORD
}

func (m *tStatStg) ToGo() TStatStg {
	if m == nil {
		return TStatStg{}
	}
	return TStatStg{
		PwcsName:          api.GoStr(m.PwcsName),
		DwType:            m.DwType,
		CbSize:            m.CbSize,
		Mtime:             api.GoStr(m.Mtime),
		Ctime:             api.GoStr(m.Ctime),
		Atime:             api.GoStr(m.Atime),
		GrfMode:           m.GrfMode,
		GrfLocksSupported: m.GrfLocksSupported,
		Clsid:             api.GoStr(m.Clsid),
		GrfStateBits:      m.GrfStateBits,
		Reserved:          m.Reserved,
	}
}
func (m *TSynEditMouseActionInfo) ToPas() *tSynEditMouseActionInfo {
	if m == nil {
		return nil
	}
	return &tSynEditMouseActionInfo{
		NewCaret:      base.GetObjectUintptr(m.NewCaret),
		Button:        m.Button,
		Shift:         m.Shift,
		MouseX:        m.MouseX,
		MouseY:        m.MouseY,
		WheelDelta:    m.WheelDelta,
		CCount:        m.CCount,
		Dir:           m.Dir,
		CaretDone:     m.CaretDone,
		IgnoreUpClick: m.IgnoreUpClick,
		ActionResult:  m.ActionResult,
	}
}

type tSynEditMouseActionInfo struct {
	NewCaret      uintptr                   // TSynEditCaret
	Button        types.TSynMouseButton     // TSynMouseButton
	Shift         types.TShiftState         // TShiftState
	MouseX        int32                     // Integer
	MouseY        int32                     // Integer
	WheelDelta    int32                     // Integer
	CCount        types.TSynMAClickCount    // TSynMAClickCount
	Dir           types.TSynMAClickDir      // TSynMAClickDir
	CaretDone     types.LongBool            // Boolean
	IgnoreUpClick types.LongBool            // Boolean
	ActionResult  TSynEditMouseActionResult // TSynEditMouseActionResult
}

func (m *tSynEditMouseActionInfo) ToGo() TSynEditMouseActionInfo {
	if m == nil {
		return TSynEditMouseActionInfo{}
	}
	return TSynEditMouseActionInfo{
		NewCaret:      AsSynEditCaret(m.NewCaret),
		Button:        m.Button,
		Shift:         m.Shift,
		MouseX:        m.MouseX,
		MouseY:        m.MouseY,
		WheelDelta:    m.WheelDelta,
		CCount:        m.CCount,
		Dir:           m.Dir,
		CaretDone:     m.CaretDone,
		IgnoreUpClick: m.IgnoreUpClick,
		ActionResult:  m.ActionResult,
	}
}
func (m *TSynEditMouseActionResult) ToPas() *tSynEditMouseActionResult {
	if m == nil {
		return nil
	}
	return &tSynEditMouseActionResult{
		DoPopUpEvent: m.DoPopUpEvent,
		PopUpEventX:  m.PopUpEventX,
		PopUpEventY:  m.PopUpEventY,
		PopUpMenu:    base.GetObjectUintptr(m.PopUpMenu),
	}
}

type tSynEditMouseActionResult struct {
	DoPopUpEvent types.LongBool // Boolean
	PopUpEventX  int32          // Integer
	PopUpEventY  int32          // Integer
	PopUpMenu    uintptr        // TPopupMenu
}

func (m *tSynEditMouseActionResult) ToGo() TSynEditMouseActionResult {
	if m == nil {
		return TSynEditMouseActionResult{}
	}
	return TSynEditMouseActionResult{
		DoPopUpEvent: m.DoPopUpEvent,
		PopUpEventX:  m.PopUpEventX,
		PopUpEventY:  m.PopUpEventY,
		PopUpMenu:    AsPopupMenu(m.PopUpMenu),
	}
}
func (m *TVTImageInfo) ToPas() *tVTImageInfo {
	if m == nil {
		return nil
	}
	return &tVTImageInfo{
		Index:   m.Index,
		XPos:    m.XPos,
		YPos:    m.YPos,
		Ghosted: m.Ghosted,
		Images:  base.GetObjectUintptr(m.Images),
	}
}

type tVTImageInfo struct {
	Index   int32          // Integer
	XPos    int32          // Integer
	YPos    int32          // Integer
	Ghosted types.LongBool // Boolean
	Images  uintptr        // TCustomImageList
}

func (m *tVTImageInfo) ToGo() TVTImageInfo {
	if m == nil {
		return TVTImageInfo{}
	}
	return TVTImageInfo{
		Index:   m.Index,
		XPos:    m.XPos,
		YPos:    m.YPos,
		Ghosted: m.Ghosted,
		Images:  AsCustomImageList(m.Images),
	}
}
func (m *TVTPaintInfo) ToPas() *tVTPaintInfo {
	if m == nil {
		return nil
	}
	return &tVTPaintInfo{
		Canvas:           base.GetObjectUintptr(m.Canvas),
		PaintOptions:     m.PaintOptions,
		Node:             m.Node,
		Column:           m.Column,
		Position:         m.Position,
		CellRect:         m.CellRect,
		ContentRect:      m.ContentRect,
		NodeWidth:        m.NodeWidth,
		Alignment:        m.Alignment,
		CaptionAlignment: m.CaptionAlignment,
		BidiMode:         m.BidiMode,
		BrushOrigin:      m.BrushOrigin,
		ImageInfo:        m.ImageInfo,
		ImageInfoSize:    m.ImageInfoSize,
	}
}

type tVTPaintInfo struct {
	Canvas           uintptr                       // TCanvas
	PaintOptions     types.TVTInternalPaintOptions // TVTInternalPaintOptions
	Node             types.PVirtualNode            // PVirtualNode
	Column           int32                         // TColumnIndex
	Position         uint32                        // TColumnPosition
	CellRect         types.TRect                   // TRect
	ContentRect      types.TRect                   // TRect
	NodeWidth        int32                         // Integer
	Alignment        types.TAlignment              // TAlignment
	CaptionAlignment types.TAlignment              // TAlignment
	BidiMode         types.TBiDiMode               // TBidiMode
	BrushOrigin      types.TPoint                  // TPoint
	ImageInfo        uintptr                       // array[TVTImageInfoIndex] of TVTImageInfo
	ImageInfoSize    int32                         // SizeOf: array[TVTImageInfoIndex] of TVTImageInfo
}

func (m *tVTPaintInfo) ToGo() TVTPaintInfo {
	if m == nil {
		return TVTPaintInfo{}
	}
	return TVTPaintInfo{
		Canvas:           AsCanvas(m.Canvas),
		PaintOptions:     m.PaintOptions,
		Node:             m.Node,
		Column:           m.Column,
		Position:         m.Position,
		CellRect:         m.CellRect,
		ContentRect:      m.ContentRect,
		NodeWidth:        m.NodeWidth,
		Alignment:        m.Alignment,
		CaptionAlignment: m.CaptionAlignment,
		BidiMode:         m.BidiMode,
		BrushOrigin:      m.BrushOrigin,
		ImageInfo:        m.ImageInfo,
		ImageInfoSize:    m.ImageInfoSize,
	}
}
