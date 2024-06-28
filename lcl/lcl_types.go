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
	"github.com/energye/lcl/types"
)

// TRawImagePosition
//
//	Record describing a position in Raw Image data.
//	Byte is the byte offset, Bit the bit number/offset in that byte.
type TRawImagePosition struct {
	Byte types.PtrUInt
	Bit  types.Cardinal
}

// TColumnIndex int32
type TColumnIndex = int32

// TNodeArray array of PVirtualNode
// TODO no impl, 当前写法不正确
type TNodeArray = uintptr

// TColumnPosition Cardinal
type TColumnPosition = types.Cardinal

// TVTPaintInfo Record
type TVTPaintInfo struct {
	Canvas           ICanvas                       // the canvas to paint on
	PaintOptions     types.TVTInternalPaintOptions // a copy of the paint options passed to PaintTree
	Node             IVirtualNode                  // the node to paint
	Column           TColumnIndex                  // the node's column index to paint
	Position         TColumnPosition               // the column position of the node
	CellRect         types.TRect                   // the node cell
	ContentRect      types.TRect                   // the area of the cell used for the node's content
	NodeWidth        types.Integer                 // the actual node width
	Alignment        types.TAlignment              // how to align within the node rectangle
	CaptionAlignment types.TAlignment              // how to align text within the caption rectangle
	BidiMode         types.TBiDiMode               // directionality to be used for painting
	BrushOrigin      types.TPoint                  // the alignment for the brush used to draw dotted lines
	imageInfoPtr     unsafePointer                 // [4]TVTImageInfo // info about each possible node image, array[TVTImageInfoIndex] of TVTImageInfo
}

// TVTImageInfo Record
//
//	For painting a node and its columns/cells a lot of information must be passed frequently around.
type TVTImageInfo struct {
	Index   types.Integer    // Index in the associated image list.
	XPos    types.Integer    // Horizontal position in the current target canvas.
	YPos    types.Integer    // Vertical position in the current target canvas.
	Ghosted types.Boolean    // Flag to indicate that the image must be drawn slightly lighter.
	Images  TCustomImageList // The image list to be used for painting.
}

// ImageInfo array[TVTImageInfoIndex] of TVTImageInfo
func (m *TVTPaintInfo) ImageInfo(index types.TVTImageInfoIndex) *TVTImageInfo {
	if index >= 0 && index <= 3 {
		var result uintptr
		api.LCLPreDef().SysCallN(api.VTImageInfoGet(), uintptr(m.imageInfoPtr), uintptr(index), uintptr(unsafePointer(&result)))
	}
	return nil
}

// TVTHeaderHitInfo Record
//
//	Structure used when info about a certain position in the header is needed.
type TVTHeaderHitInfo struct {
	X           types.Integer
	Y           types.Integer
	Button      types.TMouseButton
	Shift       types.TShiftState
	Column      TColumnIndex
	HitPosition types.TVTHeaderHitPositions
}

// THeaderPaintInfo Record
//
//	This structure carries all important information about header painting and is used in the advanced header painting.
type THeaderPaintInfo struct {
	instance        *tHeaderPaintInfo
	TargetCanvas    ICanvas
	Column          IVirtualTreeColumn
	PaintRectangle  *types.TRect
	TextRectangle   *types.TRect
	IsHoverIndex    types.Boolean
	IsDownIndex     types.Boolean
	IsEnabled       types.Boolean
	ShowHeaderGlyph types.Boolean
	ShowSortGlyph   types.Boolean
	ShowRightBorder types.Boolean
	DropMark        types.TVTDropMarkMode
	GlyphPos        *types.TPoint
	SortGlyphPos    *types.TPoint
}

func (m *THeaderPaintInfo) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	var setRectPtrVal = func(src, target *types.TRect) {
		target.Left = src.Left
		target.Top = src.Top
		target.Right = src.Right
		target.Bottom = src.Bottom
	}
	var setPointPtrVal = func(src, target *types.TPoint) {
		target.X = src.X
		target.Y = src.Y
	}
	*(*uintptr)(unsafePointer(m.instance.TargetCanvas)) = m.TargetCanvas.Instance()
	*(*uintptr)(unsafePointer(m.instance.Column)) = m.Column.Instance()
	paintRectangle := (*types.TRect)(unsafePointer(m.instance.PaintRectangle))
	textRectangleL := (*types.TRect)(unsafePointer(m.instance.TextRectangle))
	setRectPtrVal(m.PaintRectangle, paintRectangle)
	setRectPtrVal(m.TextRectangle, textRectangleL)
	*(*types.Boolean)(unsafePointer(m.instance.IsHoverIndex)) = m.IsHoverIndex
	*(*types.Boolean)(unsafePointer(m.instance.IsDownIndex)) = m.IsDownIndex
	*(*types.Boolean)(unsafePointer(m.instance.IsEnabled)) = m.IsEnabled
	*(*types.Boolean)(unsafePointer(m.instance.ShowHeaderGlyph)) = m.ShowHeaderGlyph
	*(*types.Boolean)(unsafePointer(m.instance.ShowSortGlyph)) = m.ShowSortGlyph
	*(*types.Boolean)(unsafePointer(m.instance.ShowRightBorder)) = m.ShowRightBorder
	*(*types.TVTDropMarkMode)(unsafePointer(m.instance.DropMark)) = m.DropMark
	glyphPos := (*types.TPoint)(unsafePointer(m.instance.GlyphPos))
	sortGlyphPos := (*types.TPoint)(unsafePointer(m.instance.SortGlyphPos))
	setPointPtrVal(m.GlyphPos, glyphPos)
	setPointPtrVal(m.SortGlyphPos, sortGlyphPos)
}

// THitInfo Record
//
//	Structure used when info about a certain position in the tree is needed.
type THitInfo struct {
	HitNode      IVirtualNode
	HitPositions types.THitPositions
	HitColumn    TColumnIndex
	HitPoint     types.TPoint
}

// THintInfo record
type THintInfo struct {
	instance        *tHintInfo
	HintControl     IControl
	HintWindowClass types.TWinControlClass
	HintPos         *types.TPoint // screen coordinates
	HintMaxWidth    int32
	HintColor       types.TColor
	CursorRect      *types.TRect
	CursorPos       *types.TPoint
	ReshowTimeout   int32
	HideTimeout     int32
	HintStr         string
	HintData        types.Pointer
}

func (m *THintInfo) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	var setRectPtrVal = func(src, target *types.TRect) {
		target.Left = src.Left
		target.Top = src.Top
		target.Right = src.Right
		target.Bottom = src.Bottom
	}
	var setPointPtrVal = func(src, target *types.TPoint) {
		target.X = src.X
		target.Y = src.Y
	}
	*(*uintptr)(unsafePointer(m.instance.HintControl)) = m.HintControl.Instance()
	*(*types.TWinControlClass)(unsafePointer(m.instance.HintWindowClass)) = m.HintWindowClass
	hintPos := (*types.TPoint)(unsafePointer(m.instance.HintPos))
	setPointPtrVal(m.HintPos, hintPos)
	*(*int32)(unsafePointer(m.instance.HintMaxWidth)) = m.HintMaxWidth
	*(*types.TColor)(unsafePointer(m.instance.HintColor)) = m.HintColor
	cursorRect := (*types.TRect)(unsafePointer(m.instance.CursorRect))
	setRectPtrVal(m.CursorRect, cursorRect)
	cursorPos := (*types.TPoint)(unsafePointer(m.instance.CursorPos))
	setPointPtrVal(m.CursorPos, cursorPos)
	*(*int32)(unsafePointer(m.instance.ReshowTimeout)) = m.ReshowTimeout
	*(*int32)(unsafePointer(m.instance.HideTimeout)) = m.HideTimeout
	m.instance.HintStr = api.PascalStr(m.HintStr)
	m.instance.HintData = m.HintData
}
