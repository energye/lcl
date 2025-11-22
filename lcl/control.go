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

// IControl Parent: ILCLComponent
type IControl interface {
	ILCLComponent
	ManualDock(newDockSite IWinControl, dropControl IControl, controlSide types.TAlign, keepDockSiteSize bool) bool      // function
	ManualFloat(theScreenRect types.TRect, keepDockSiteSize bool) bool                                                   // function
	ReplaceDockedControl(control IControl, newDockSite IWinControl, dropControl IControl, controlSide types.TAlign) bool // function
	Docked() bool                                                                                                        // function
	Dragging() bool                                                                                                      // function
	// GetAccessibleObject
	//  accessibility
	GetAccessibleObject() ILazAccessibleObject                           // function
	CreateAccessibleObject() ILazAccessibleObject                        // function
	GetSelectedChildAccessibleObject() ILazAccessibleObject              // function
	GetChildAccessibleObjectAtPos(pos types.TPoint) ILazAccessibleObject // function
	// ScaleDesignToForm
	//  scale support
	ScaleDesignToForm(size int32) int32                                    // function
	ScaleFormToDesign(size int32) int32                                    // function
	Scale96ToForm(size int32) int32                                        // function
	ScaleFormTo96(size int32) int32                                        // function
	Scale96ToFont(size int32) int32                                        // function
	ScaleFontTo96(size int32) int32                                        // function
	ScaleScreenToFont(size int32) int32                                    // function
	ScaleFontToScreen(size int32) int32                                    // function
	Scale96ToScreen(size int32) int32                                      // function
	ScaleScreenTo96(size int32) int32                                      // function
	AutoSizePhases() types.TControlAutoSizePhases                          // function
	AutoSizeDelayed() bool                                                 // function
	AutoSizeDelayedReport() string                                         // function
	AutoSizeDelayedHandle() bool                                           // function
	AnchoredControlCount() int32                                           // function
	GetCanvasScaleFactor() float64                                         // function
	GetDefaultWidth() int32                                                // function
	GetDefaultHeight() int32                                               // function
	GetDefaultColor(defaultColorType types.TDefaultColorType) types.TColor // function
	// GetColorResolvingParent
	//  These two are helper routines to help obtain the background color of a control
	GetColorResolvingParent() types.TColor    // function
	GetRGBColorResolvingParent() types.TColor // function
	// GetSidePosition
	//
	GetSidePosition(side types.TAnchorKind) int32                                 // function
	GetAnchorsDependingOnParent(withNormalAnchors bool) types.TAnchors            // function
	IsParentOf(control IControl) bool                                             // function
	GetTopParent() IControl                                                       // function
	FindSubComponent(name string) IComponent                                      // function
	IsVisible() bool                                                              // function
	IsControlVisible() bool                                                       // function
	IsEnabled() bool                                                              // function
	IsParentColor() bool                                                          // function
	IsParentFont() bool                                                           // function
	FormIsUpdating() bool                                                         // function
	IsProcessingPaintMsg() bool                                                   // function
	CheckChildClassAllowed(childClass types.TClass, exceptionOnInvalid bool) bool // function
	GetTextBuf(buffer uintptr, bufSize int32) int32                               // function
	GetTextLen() int32                                                            // function
	Perform(msg uint32, wParam types.WParam, lParam types.LParam) types.LRESULT   // function
	ScreenToClient(point types.TPoint) types.TPoint                               // function
	ClientToScreenWithPoint(point types.TPoint) types.TPoint                      // function
	ClientToScreenWithRect(rect types.TRect) types.TRect                          // function
	ScreenToControl(point types.TPoint) types.TPoint                              // function
	ControlToScreen(point types.TPoint) types.TPoint                              // function
	ClientToParent(point types.TPoint, parent IWinControl) types.TPoint           // function
	ParentToClient(point types.TPoint, parent IWinControl) types.TPoint           // function
	GetChildrenRect(scrolled bool) types.TRect                                    // function
	HandleObjectShouldBeVisible() bool                                            // function
	ParentDestroyingHandle() bool                                                 // function
	ParentHandlesAllocated() bool                                                 // function
	HasHelp() bool                                                                // function
	UseRightToLeftAlignment() bool                                                // function
	UseRightToLeftReading() bool                                                  // function
	UseRightToLeftScrollBar() bool                                                // function
	IsRightToLeft() bool                                                          // function
	// DragDrop
	//  is a hack for speed. It will be replaced by the use of the widgetset
	//  classes.
	//  So, don't use it anymore.
	//  drag and dock
	DragDrop(source IObject, X int32, Y int32)      // procedure
	Dock(newDockSite IWinControl, rect types.TRect) // procedure
	// AdjustSize
	//  size
	AdjustSize()                                                                                                  // procedure
	AnchorToNeighbour(side types.TAnchorKind, space types.TSpacingSize, sibling IControl)                         // procedure
	AnchorParallel(side types.TAnchorKind, space types.TSpacingSize, sibling IControl)                            // procedure
	AnchorHorizontalCenterTo(sibling IControl)                                                                    // procedure
	AnchorVerticalCenterTo(sibling IControl)                                                                      // procedure
	AnchorToCompanion(side types.TAnchorKind, space types.TSpacingSize, sibling IControl, freeCompositeSide bool) // procedure
	AnchorSame(side types.TAnchorKind, sibling IControl)                                                          // procedure
	AnchorAsAlign(theAlign types.TAlign, space types.TSpacingSize)                                                // procedure
	AnchorClient(space types.TSpacingSize)                                                                        // procedure
	SetBounds(left int32, top int32, width int32, height int32)                                                   // procedure
	SetInitialBounds(left int32, top int32, width int32, height int32)                                            // procedure
	SetBoundsKeepBase(left int32, top int32, width int32, height int32)                                           // procedure
	GetPreferredSize(preferredWidth *int32, preferredHeight *int32, raw bool, withThemeSpace bool)                // procedure
	CNPreferredSizeChanged()                                                                                      // procedure
	InvalidatePreferredSize()                                                                                     // procedure
	DisableAutoSizing()                                                                                           // procedure
	EnableAutoSizing()                                                                                            // procedure
	UpdateBaseBounds(storeBounds bool, storeParentClientSize bool, useLoadedValues bool)                          // procedure
	WriteLayoutDebugReport(prefix string)                                                                         // procedure
	// AutoAdjustLayout
	//  LCL Scaling (High-DPI)
	AutoAdjustLayout(mode types.TLayoutAdjustmentPolicy, fromPPI int32, toPPI int32, oldFormWidth int32, newFormWidth int32) // procedure
	ShouldAutoAdjust(width *bool, height *bool)                                                                              // procedure
	FixDesignFontsPPI(designTimePPI int32)                                                                                   // procedure
	ScaleFontsPPI(toPPI int32, proportion float64)                                                                           // procedure
	EditingDone()                                                                                                            // procedure
	ExecuteDefaultAction()                                                                                                   // procedure
	ExecuteCancelAction()                                                                                                    // procedure
	BeginDrag(immediate bool, threshold int32)                                                                               // procedure
	EndDrag(drop bool)                                                                                                       // procedure
	BringToFront()                                                                                                           // procedure
	Hide()                                                                                                                   // procedure
	Refresh()                                                                                                                // procedure
	Repaint()                                                                                                                // procedure
	Invalidate()                                                                                                             // procedure
	CheckNewParent(parent IWinControl)                                                                                       // procedure
	SendToBack()                                                                                                             // procedure
	SetTempCursor(value types.TCursor)                                                                                       // procedure
	UpdateRolesForForm()                                                                                                     // procedure
	ActiveDefaultControlChanged(newControl IControl)                                                                         // procedure
	SetTextBuf(buffer uintptr)                                                                                               // procedure
	Show()                                                                                                                   // procedure
	Update()                                                                                                                 // procedure
	InitiateAction()                                                                                                         // procedure
	ShowHelp()                                                                                                               // procedure
	AnchoredControls(index int32) IControl                                                                                   // property AnchoredControls Getter
	BaseBounds() types.TRect                                                                                                 // property BaseBounds Getter
	ReadBounds() types.TRect                                                                                                 // property ReadBounds Getter
	BaseParentClientSize() types.TSize                                                                                       // property BaseParentClientSize Getter
	// AccessibleName
	//  standard properties, which should be supported by all descendants
	AccessibleName() string                                  // property AccessibleName Getter
	SetAccessibleName(value string)                          // property AccessibleName Setter
	AccessibleDescription() string                           // property AccessibleDescription Getter
	SetAccessibleDescription(value string)                   // property AccessibleDescription Setter
	AccessibleValue() string                                 // property AccessibleValue Getter
	SetAccessibleValue(value string)                         // property AccessibleValue Setter
	AccessibleRole() types.TLazAccessibilityRole             // property AccessibleRole Getter
	SetAccessibleRole(value types.TLazAccessibilityRole)     // property AccessibleRole Setter
	Action() IBasicAction                                    // property Action Getter
	SetAction(value IBasicAction)                            // property Action Setter
	Align() types.TAlign                                     // property Align Getter
	SetAlign(value types.TAlign)                             // property Align Setter
	Anchors() types.TAnchors                                 // property Anchors Getter
	SetAnchors(value types.TAnchors)                         // property Anchors Setter
	AnchorSide(kind types.TAnchorKind) IAnchorSide           // property AnchorSide Getter
	AutoSize() bool                                          // property AutoSize Getter
	SetAutoSize(value bool)                                  // property AutoSize Setter
	BorderSpacing() IControlBorderSpacing                    // property BorderSpacing Getter
	SetBorderSpacing(value IControlBorderSpacing)            // property BorderSpacing Setter
	BoundsRect() types.TRect                                 // property BoundsRect Getter
	SetBoundsRect(value types.TRect)                         // property BoundsRect Setter
	BoundsRectForNewParent() types.TRect                     // property BoundsRectForNewParent Getter
	SetBoundsRectForNewParent(value types.TRect)             // property BoundsRectForNewParent Setter
	Caption() string                                         // property Caption Getter
	SetCaption(value string)                                 // property Caption Setter
	CaptureMouseButtons() types.TCaptureMouseButtons         // property CaptureMouseButtons Getter
	SetCaptureMouseButtons(value types.TCaptureMouseButtons) // property CaptureMouseButtons Setter
	ClientHeight() int32                                     // property ClientHeight Getter
	SetClientHeight(value int32)                             // property ClientHeight Setter
	ClientOrigin() types.TPoint                              // property ClientOrigin Getter
	ClientRect() types.TRect                                 // property ClientRect Getter
	ClientWidth() int32                                      // property ClientWidth Getter
	SetClientWidth(value int32)                              // property ClientWidth Setter
	Color() types.TColor                                     // property Color Getter
	SetColor(value types.TColor)                             // property Color Setter
	Constraints() ISizeConstraints                           // property Constraints Getter
	SetConstraints(value ISizeConstraints)                   // property Constraints Setter
	ControlOrigin() types.TPoint                             // property ControlOrigin Getter
	ControlState() types.TControlState                       // property ControlState Getter
	SetControlState(value types.TControlState)               // property ControlState Setter
	ControlStyle() types.TControlStyle                       // property ControlStyle Getter
	SetControlStyle(value types.TControlStyle)               // property ControlStyle Setter
	Enabled() bool                                           // property Enabled Getter
	SetEnabled(value bool)                                   // property Enabled Setter
	Font() IFont                                             // property Font Getter
	SetFont(value IFont)                                     // property Font Setter
	IsControl() bool                                         // property IsControl Getter
	SetIsControl(value bool)                                 // property IsControl Setter
	MouseInClient() bool                                     // property MouseInClient Getter
	Parent() IWinControl                                     // property Parent Getter
	SetParent(value IWinControl)                             // property Parent Setter
	PopupMenu() IPopupMenu                                   // property PopupMenu Getter
	SetPopupMenu(value IPopupMenu)                           // property PopupMenu Setter
	ShowHint() bool                                          // property ShowHint Getter
	SetShowHint(value bool)                                  // property ShowHint Setter
	Visible() bool                                           // property Visible Getter
	SetVisible(value bool)                                   // property Visible Setter
	// DockOrientation
	//  docking properties
	DockOrientation() types.TDockOrientation               // property DockOrientation Getter
	SetDockOrientation(value types.TDockOrientation)       // property DockOrientation Setter
	Floating() bool                                        // property Floating Getter
	FloatingDockSiteClass() types.TWinControlClass         // property FloatingDockSiteClass Getter
	SetFloatingDockSiteClass(value types.TWinControlClass) // property FloatingDockSiteClass Setter
	HostDockSite() IWinControl                             // property HostDockSite Getter
	SetHostDockSite(value IWinControl)                     // property HostDockSite Setter
	LRDockWidth() int32                                    // property LRDockWidth Getter
	SetLRDockWidth(value int32)                            // property LRDockWidth Setter
	TBDockHeight() int32                                   // property TBDockHeight Getter
	SetTBDockHeight(value int32)                           // property TBDockHeight Setter
	UndockHeight() int32                                   // property UndockHeight Getter
	SetUndockHeight(value int32)                           // property UndockHeight Setter
	UndockWidth() int32                                    // property UndockWidth Getter
	SetUndockWidth(value int32)                            // property UndockWidth Setter
	BiDiMode() types.TBiDiMode                             // property BiDiMode Getter
	SetBiDiMode(value types.TBiDiMode)                     // property BiDiMode Setter
	ParentBiDiMode() bool                                  // property ParentBiDiMode Getter
	SetParentBiDiMode(value bool)                          // property ParentBiDiMode Setter
	AnchorSideLeft() IAnchorSide                           // property AnchorSideLeft Getter
	SetAnchorSideLeft(value IAnchorSide)                   // property AnchorSideLeft Setter
	AnchorSideTop() IAnchorSide                            // property AnchorSideTop Getter
	SetAnchorSideTop(value IAnchorSide)                    // property AnchorSideTop Setter
	AnchorSideRight() IAnchorSide                          // property AnchorSideRight Getter
	SetAnchorSideRight(value IAnchorSide)                  // property AnchorSideRight Setter
	AnchorSideBottom() IAnchorSide                         // property AnchorSideBottom Getter
	SetAnchorSideBottom(value IAnchorSide)                 // property AnchorSideBottom Setter
	Cursor() types.TCursor                                 // property Cursor Getter
	SetCursor(value types.TCursor)                         // property Cursor Setter
	Left() int32                                           // property Left Getter
	SetLeft(value int32)                                   // property Left Setter
	Height() int32                                         // property Height Getter
	SetHeight(value int32)                                 // property Height Setter
	Hint() string                                          // property Hint Getter
	SetHint(value string)                                  // property Hint Setter
	Top() int32                                            // property Top Getter
	SetTop(value int32)                                    // property Top Setter
	Width() int32                                          // property Width Getter
	SetWidth(value int32)                                  // property Width Setter
	HelpType() types.THelpType                             // property HelpType Getter
	SetHelpType(value types.THelpType)                     // property HelpType Setter
	HelpKeyword() string                                   // property HelpKeyword Getter
	SetHelpKeyword(value string)                           // property HelpKeyword Setter
	HelpContext() types.THelpContext                       // property HelpContext Getter
	SetHelpContext(value types.THelpContext)               // property HelpContext Setter
	SetOnChangeBounds(fn TNotifyEvent)                     // property event
	SetOnClick(fn TNotifyEvent)                            // property event
	SetOnResize(fn TNotifyEvent)                           // property event
	SetOnShowHint(fn TControlShowHintEvent)                // property event
	SetWindowProc(fn TWndMethod)                           // property event
}

type TControl struct {
	TLCLComponent
}

func (m *TControl) ManualDock(newDockSite IWinControl, dropControl IControl, controlSide types.TAlign, keepDockSiteSize bool) bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(newDockSite), base.GetObjectUintptr(dropControl), uintptr(controlSide), api.PasBool(keepDockSiteSize))
	return api.GoBool(r)
}

func (m *TControl) ManualFloat(theScreenRect types.TRect, keepDockSiteSize bool) bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&theScreenRect)), api.PasBool(keepDockSiteSize))
	return api.GoBool(r)
}

func (m *TControl) ReplaceDockedControl(control IControl, newDockSite IWinControl, dropControl IControl, controlSide types.TAlign) bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(control), base.GetObjectUintptr(newDockSite), base.GetObjectUintptr(dropControl), uintptr(controlSide))
	return api.GoBool(r)
}

func (m *TControl) Docked() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) Dragging() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) GetAccessibleObject() ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(6, m.Instance())
	return AsLazAccessibleObject(r)
}

func (m *TControl) CreateAccessibleObject() ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(7, m.Instance())
	return AsLazAccessibleObject(r)
}

func (m *TControl) GetSelectedChildAccessibleObject() ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(8, m.Instance())
	return AsLazAccessibleObject(r)
}

func (m *TControl) GetChildAccessibleObjectAtPos(pos types.TPoint) ILazAccessibleObject {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&pos)))
	return AsLazAccessibleObject(r)
}

func (m *TControl) ScaleDesignToForm(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(10, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) ScaleFormToDesign(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(11, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) Scale96ToForm(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(12, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) ScaleFormTo96(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(13, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) Scale96ToFont(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(14, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) ScaleFontTo96(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(15, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) ScaleScreenToFont(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(16, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) ScaleFontToScreen(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(17, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) Scale96ToScreen(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(18, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) ScaleScreenTo96(size int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(19, m.Instance(), uintptr(size))
	return int32(r)
}

func (m *TControl) AutoSizePhases() types.TControlAutoSizePhases {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(20, m.Instance())
	return types.TControlAutoSizePhases(r)
}

func (m *TControl) AutoSizeDelayed() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(21, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) AutoSizeDelayedReport() string {
	if !m.IsValid() {
		return ""
	}
	r := controlAPI().SysCallN(22, m.Instance())
	return api.GoStr(r)
}

func (m *TControl) AutoSizeDelayedHandle() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(23, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) AnchoredControlCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(24, m.Instance())
	return int32(r)
}

func (m *TControl) GetCanvasScaleFactor() (result float64) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) GetDefaultWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(26, m.Instance())
	return int32(r)
}

func (m *TControl) GetDefaultHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(27, m.Instance())
	return int32(r)
}

func (m *TControl) GetDefaultColor(defaultColorType types.TDefaultColorType) types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(28, m.Instance(), uintptr(defaultColorType))
	return types.TColor(r)
}

func (m *TControl) GetColorResolvingParent() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(29, m.Instance())
	return types.TColor(r)
}

func (m *TControl) GetRGBColorResolvingParent() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(30, m.Instance())
	return types.TColor(r)
}

func (m *TControl) GetSidePosition(side types.TAnchorKind) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(31, m.Instance(), uintptr(side))
	return int32(r)
}

func (m *TControl) GetAnchorsDependingOnParent(withNormalAnchors bool) types.TAnchors {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(32, m.Instance(), api.PasBool(withNormalAnchors))
	return types.TAnchors(r)
}

func (m *TControl) IsParentOf(control IControl) bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(33, m.Instance(), base.GetObjectUintptr(control))
	return api.GoBool(r)
}

func (m *TControl) GetTopParent() IControl {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(34, m.Instance())
	return AsControl(r)
}

func (m *TControl) FindSubComponent(name string) IComponent {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(35, m.Instance(), api.PasStr(name))
	return AsComponent(r)
}

func (m *TControl) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(36, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) IsControlVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(37, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) IsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(38, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) IsParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(39, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) IsParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(40, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) FormIsUpdating() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(41, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) IsProcessingPaintMsg() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(42, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) CheckChildClassAllowed(childClass types.TClass, exceptionOnInvalid bool) bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(43, m.Instance(), uintptr(childClass), api.PasBool(exceptionOnInvalid))
	return api.GoBool(r)
}

func (m *TControl) GetTextBuf(buffer uintptr, bufSize int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(44, m.Instance(), uintptr(buffer), uintptr(bufSize))
	return int32(r)
}

func (m *TControl) GetTextLen() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(45, m.Instance())
	return int32(r)
}

func (m *TControl) Perform(msg uint32, wParam types.WParam, lParam types.LParam) types.LRESULT {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(46, m.Instance(), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return types.LRESULT(r)
}

func (m *TControl) ScreenToClient(point types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(47, m.Instance(), uintptr(base.UnsafePointer(&point)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ClientToScreenWithPoint(point types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(48, m.Instance(), uintptr(base.UnsafePointer(&point)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ClientToScreenWithRect(rect types.TRect) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(49, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ScreenToControl(point types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(50, m.Instance(), uintptr(base.UnsafePointer(&point)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ControlToScreen(point types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(51, m.Instance(), uintptr(base.UnsafePointer(&point)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ClientToParent(point types.TPoint, parent IWinControl) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(52, m.Instance(), uintptr(base.UnsafePointer(&point)), base.GetObjectUintptr(parent), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ParentToClient(point types.TPoint, parent IWinControl) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(53, m.Instance(), uintptr(base.UnsafePointer(&point)), base.GetObjectUintptr(parent), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) GetChildrenRect(scrolled bool) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(54, m.Instance(), api.PasBool(scrolled), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) HandleObjectShouldBeVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(55, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) ParentDestroyingHandle() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(56, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) ParentHandlesAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(57, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) HasHelp() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(58, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) UseRightToLeftAlignment() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(59, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) UseRightToLeftReading() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(60, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) UseRightToLeftScrollBar() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(61, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) IsRightToLeft() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(62, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) DragDrop(source IObject, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(63, m.Instance(), base.GetObjectUintptr(source), uintptr(X), uintptr(Y))
}

func (m *TControl) Dock(newDockSite IWinControl, rect types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(64, m.Instance(), base.GetObjectUintptr(newDockSite), uintptr(base.UnsafePointer(&rect)))
}

func (m *TControl) AdjustSize() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(65, m.Instance())
}

func (m *TControl) AnchorToNeighbour(side types.TAnchorKind, space types.TSpacingSize, sibling IControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(66, m.Instance(), uintptr(side), uintptr(space), base.GetObjectUintptr(sibling))
}

func (m *TControl) AnchorParallel(side types.TAnchorKind, space types.TSpacingSize, sibling IControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(67, m.Instance(), uintptr(side), uintptr(space), base.GetObjectUintptr(sibling))
}

func (m *TControl) AnchorHorizontalCenterTo(sibling IControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(68, m.Instance(), base.GetObjectUintptr(sibling))
}

func (m *TControl) AnchorVerticalCenterTo(sibling IControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(69, m.Instance(), base.GetObjectUintptr(sibling))
}

func (m *TControl) AnchorToCompanion(side types.TAnchorKind, space types.TSpacingSize, sibling IControl, freeCompositeSide bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(70, m.Instance(), uintptr(side), uintptr(space), base.GetObjectUintptr(sibling), api.PasBool(freeCompositeSide))
}

func (m *TControl) AnchorSame(side types.TAnchorKind, sibling IControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(71, m.Instance(), uintptr(side), base.GetObjectUintptr(sibling))
}

func (m *TControl) AnchorAsAlign(theAlign types.TAlign, space types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(72, m.Instance(), uintptr(theAlign), uintptr(space))
}

func (m *TControl) AnchorClient(space types.TSpacingSize) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(73, m.Instance(), uintptr(space))
}

func (m *TControl) SetBounds(left int32, top int32, width int32, height int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(74, m.Instance(), uintptr(left), uintptr(top), uintptr(width), uintptr(height))
}

func (m *TControl) SetInitialBounds(left int32, top int32, width int32, height int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(75, m.Instance(), uintptr(left), uintptr(top), uintptr(width), uintptr(height))
}

func (m *TControl) SetBoundsKeepBase(left int32, top int32, width int32, height int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(76, m.Instance(), uintptr(left), uintptr(top), uintptr(width), uintptr(height))
}

func (m *TControl) GetPreferredSize(preferredWidth *int32, preferredHeight *int32, raw bool, withThemeSpace bool) {
	if !m.IsValid() {
		return
	}
	preferredWidthPtr := uintptr(*preferredWidth)
	preferredHeightPtr := uintptr(*preferredHeight)
	controlAPI().SysCallN(77, m.Instance(), uintptr(base.UnsafePointer(&preferredWidthPtr)), uintptr(base.UnsafePointer(&preferredHeightPtr)), api.PasBool(raw), api.PasBool(withThemeSpace))
	*preferredWidth = int32(preferredWidthPtr)
	*preferredHeight = int32(preferredHeightPtr)
}

func (m *TControl) CNPreferredSizeChanged() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(78, m.Instance())
}

func (m *TControl) InvalidatePreferredSize() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(79, m.Instance())
}

func (m *TControl) DisableAutoSizing() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(80, m.Instance())
}

func (m *TControl) EnableAutoSizing() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(81, m.Instance())
}

func (m *TControl) UpdateBaseBounds(storeBounds bool, storeParentClientSize bool, useLoadedValues bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(82, m.Instance(), api.PasBool(storeBounds), api.PasBool(storeParentClientSize), api.PasBool(useLoadedValues))
}

func (m *TControl) WriteLayoutDebugReport(prefix string) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(83, m.Instance(), api.PasStr(prefix))
}

func (m *TControl) AutoAdjustLayout(mode types.TLayoutAdjustmentPolicy, fromPPI int32, toPPI int32, oldFormWidth int32, newFormWidth int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(84, m.Instance(), uintptr(mode), uintptr(fromPPI), uintptr(toPPI), uintptr(oldFormWidth), uintptr(newFormWidth))
}

func (m *TControl) ShouldAutoAdjust(width *bool, height *bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(85, m.Instance(), uintptr(base.UnsafePointer(width)), uintptr(base.UnsafePointer(height)))
}

func (m *TControl) FixDesignFontsPPI(designTimePPI int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(86, m.Instance(), uintptr(designTimePPI))
}

func (m *TControl) ScaleFontsPPI(toPPI int32, proportion float64) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(87, m.Instance(), uintptr(toPPI), uintptr(base.UnsafePointer(&proportion)))
}

func (m *TControl) EditingDone() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(88, m.Instance())
}

func (m *TControl) ExecuteDefaultAction() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(89, m.Instance())
}

func (m *TControl) ExecuteCancelAction() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(90, m.Instance())
}

func (m *TControl) BeginDrag(immediate bool, threshold int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(91, m.Instance(), api.PasBool(immediate), uintptr(threshold))
}

func (m *TControl) EndDrag(drop bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(92, m.Instance(), api.PasBool(drop))
}

func (m *TControl) BringToFront() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(93, m.Instance())
}

func (m *TControl) Hide() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(94, m.Instance())
}

func (m *TControl) Refresh() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(95, m.Instance())
}

func (m *TControl) Repaint() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(96, m.Instance())
}

func (m *TControl) Invalidate() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(97, m.Instance())
}

func (m *TControl) CheckNewParent(parent IWinControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(98, m.Instance(), base.GetObjectUintptr(parent))
}

func (m *TControl) SendToBack() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(99, m.Instance())
}

func (m *TControl) SetTempCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(100, m.Instance(), uintptr(value))
}

func (m *TControl) UpdateRolesForForm() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(101, m.Instance())
}

func (m *TControl) ActiveDefaultControlChanged(newControl IControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(102, m.Instance(), base.GetObjectUintptr(newControl))
}

func (m *TControl) SetTextBuf(buffer uintptr) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(103, m.Instance(), uintptr(buffer))
}

func (m *TControl) Show() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(104, m.Instance())
}

func (m *TControl) Update() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(105, m.Instance())
}

func (m *TControl) InitiateAction() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(106, m.Instance())
}

func (m *TControl) ShowHelp() {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(107, m.Instance())
}

func (m *TControl) AnchoredControls(index int32) IControl {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(108, m.Instance(), uintptr(index))
	return AsControl(r)
}

func (m *TControl) BaseBounds() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(109, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ReadBounds() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(110, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) BaseParentClientSize() (result types.TSize) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(111, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) AccessibleName() string {
	if !m.IsValid() {
		return ""
	}
	r := controlAPI().SysCallN(112, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TControl) SetAccessibleName(value string) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(112, 1, m.Instance(), api.PasStr(value))
}

func (m *TControl) AccessibleDescription() string {
	if !m.IsValid() {
		return ""
	}
	r := controlAPI().SysCallN(113, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TControl) SetAccessibleDescription(value string) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(113, 1, m.Instance(), api.PasStr(value))
}

func (m *TControl) AccessibleValue() string {
	if !m.IsValid() {
		return ""
	}
	r := controlAPI().SysCallN(114, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TControl) SetAccessibleValue(value string) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(114, 1, m.Instance(), api.PasStr(value))
}

func (m *TControl) AccessibleRole() types.TLazAccessibilityRole {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(115, 0, m.Instance())
	return types.TLazAccessibilityRole(r)
}

func (m *TControl) SetAccessibleRole(value types.TLazAccessibilityRole) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(115, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Action() IBasicAction {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(116, 0, m.Instance())
	return AsBasicAction(r)
}

func (m *TControl) SetAction(value IBasicAction) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(116, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) Align() types.TAlign {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(117, 0, m.Instance())
	return types.TAlign(r)
}

func (m *TControl) SetAlign(value types.TAlign) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(117, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Anchors() types.TAnchors {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(118, 0, m.Instance())
	return types.TAnchors(r)
}

func (m *TControl) SetAnchors(value types.TAnchors) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(118, 1, m.Instance(), uintptr(value))
}

func (m *TControl) AnchorSide(kind types.TAnchorKind) IAnchorSide {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(119, m.Instance(), uintptr(kind))
	return AsAnchorSide(r)
}

func (m *TControl) AutoSize() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(120, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) SetAutoSize(value bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(120, 1, m.Instance(), api.PasBool(value))
}

func (m *TControl) BorderSpacing() IControlBorderSpacing {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(121, 0, m.Instance())
	return AsControlBorderSpacing(r)
}

func (m *TControl) SetBorderSpacing(value IControlBorderSpacing) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(121, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) BoundsRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(122, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) SetBoundsRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(122, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TControl) BoundsRectForNewParent() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(123, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) SetBoundsRectForNewParent(value types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(123, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TControl) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := controlAPI().SysCallN(124, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TControl) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(124, 1, m.Instance(), api.PasStr(value))
}

func (m *TControl) CaptureMouseButtons() types.TCaptureMouseButtons {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(125, 0, m.Instance())
	return types.TCaptureMouseButtons(r)
}

func (m *TControl) SetCaptureMouseButtons(value types.TCaptureMouseButtons) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(125, 1, m.Instance(), uintptr(value))
}

func (m *TControl) ClientHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(126, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetClientHeight(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(126, 1, m.Instance(), uintptr(value))
}

func (m *TControl) ClientOrigin() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(127, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ClientRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(128, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ClientWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(129, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetClientWidth(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(129, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(130, 0, m.Instance())
	return types.TColor(r)
}

func (m *TControl) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(130, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Constraints() ISizeConstraints {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(131, 0, m.Instance())
	return AsSizeConstraints(r)
}

func (m *TControl) SetConstraints(value ISizeConstraints) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(131, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) ControlOrigin() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(132, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TControl) ControlState() types.TControlState {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(133, 0, m.Instance())
	return types.TControlState(r)
}

func (m *TControl) SetControlState(value types.TControlState) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(133, 1, m.Instance(), uintptr(value))
}

func (m *TControl) ControlStyle() types.TControlStyle {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(134, 0, m.Instance())
	return types.TControlStyle(r)
}

func (m *TControl) SetControlStyle(value types.TControlStyle) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(134, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(135, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(135, 1, m.Instance(), api.PasBool(value))
}

func (m *TControl) Font() IFont {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(136, 0, m.Instance())
	return AsFont(r)
}

func (m *TControl) SetFont(value IFont) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(136, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) IsControl() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(137, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) SetIsControl(value bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(137, 1, m.Instance(), api.PasBool(value))
}

func (m *TControl) MouseInClient() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(138, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) Parent() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(139, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TControl) SetParent(value IWinControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(139, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) PopupMenu() IPopupMenu {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(140, 0, m.Instance())
	return AsPopupMenu(r)
}

func (m *TControl) SetPopupMenu(value IPopupMenu) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(140, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) ShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(141, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) SetShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(141, 1, m.Instance(), api.PasBool(value))
}

func (m *TControl) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(142, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(142, 1, m.Instance(), api.PasBool(value))
}

func (m *TControl) DockOrientation() types.TDockOrientation {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(143, 0, m.Instance())
	return types.TDockOrientation(r)
}

func (m *TControl) SetDockOrientation(value types.TDockOrientation) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(143, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Floating() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(144, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) FloatingDockSiteClass() types.TWinControlClass {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(145, 0, m.Instance())
	return types.TWinControlClass(r)
}

func (m *TControl) SetFloatingDockSiteClass(value types.TWinControlClass) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(145, 1, m.Instance(), uintptr(value))
}

func (m *TControl) HostDockSite() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(146, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TControl) SetHostDockSite(value IWinControl) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(146, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) LRDockWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(147, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetLRDockWidth(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(147, 1, m.Instance(), uintptr(value))
}

func (m *TControl) TBDockHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(148, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetTBDockHeight(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(148, 1, m.Instance(), uintptr(value))
}

func (m *TControl) UndockHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(149, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetUndockHeight(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(149, 1, m.Instance(), uintptr(value))
}

func (m *TControl) UndockWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(150, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetUndockWidth(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(150, 1, m.Instance(), uintptr(value))
}

func (m *TControl) BiDiMode() types.TBiDiMode {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(151, 0, m.Instance())
	return types.TBiDiMode(r)
}

func (m *TControl) SetBiDiMode(value types.TBiDiMode) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(151, 1, m.Instance(), uintptr(value))
}

func (m *TControl) ParentBiDiMode() bool {
	if !m.IsValid() {
		return false
	}
	r := controlAPI().SysCallN(152, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControl) SetParentBiDiMode(value bool) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(152, 1, m.Instance(), api.PasBool(value))
}

func (m *TControl) AnchorSideLeft() IAnchorSide {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(153, 0, m.Instance())
	return AsAnchorSide(r)
}

func (m *TControl) SetAnchorSideLeft(value IAnchorSide) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(153, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) AnchorSideTop() IAnchorSide {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(154, 0, m.Instance())
	return AsAnchorSide(r)
}

func (m *TControl) SetAnchorSideTop(value IAnchorSide) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(154, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) AnchorSideRight() IAnchorSide {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(155, 0, m.Instance())
	return AsAnchorSide(r)
}

func (m *TControl) SetAnchorSideRight(value IAnchorSide) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(155, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) AnchorSideBottom() IAnchorSide {
	if !m.IsValid() {
		return nil
	}
	r := controlAPI().SysCallN(156, 0, m.Instance())
	return AsAnchorSide(r)
}

func (m *TControl) SetAnchorSideBottom(value IAnchorSide) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(156, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TControl) Cursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(157, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TControl) SetCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(157, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(158, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetLeft(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(158, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(159, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(159, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Hint() string {
	if !m.IsValid() {
		return ""
	}
	r := controlAPI().SysCallN(160, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TControl) SetHint(value string) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(160, 1, m.Instance(), api.PasStr(value))
}

func (m *TControl) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(161, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetTop(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(161, 1, m.Instance(), uintptr(value))
}

func (m *TControl) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(162, 0, m.Instance())
	return int32(r)
}

func (m *TControl) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(162, 1, m.Instance(), uintptr(value))
}

func (m *TControl) HelpType() types.THelpType {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(163, 0, m.Instance())
	return types.THelpType(r)
}

func (m *TControl) SetHelpType(value types.THelpType) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(163, 1, m.Instance(), uintptr(value))
}

func (m *TControl) HelpKeyword() string {
	if !m.IsValid() {
		return ""
	}
	r := controlAPI().SysCallN(164, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TControl) SetHelpKeyword(value string) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(164, 1, m.Instance(), api.PasStr(value))
}

func (m *TControl) HelpContext() types.THelpContext {
	if !m.IsValid() {
		return 0
	}
	r := controlAPI().SysCallN(165, 0, m.Instance())
	return types.THelpContext(r)
}

func (m *TControl) SetHelpContext(value types.THelpContext) {
	if !m.IsValid() {
		return
	}
	controlAPI().SysCallN(165, 1, m.Instance(), uintptr(value))
}

func (m *TControl) SetOnChangeBounds(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 166, controlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControl) SetOnClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 167, controlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControl) SetOnResize(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 168, controlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControl) SetOnShowHint(fn TControlShowHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTControlShowHintEvent(fn)
	base.SetEvent(m, 169, controlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControl) SetWindowProc(fn TWndMethod) {
	if !m.IsValid() {
		return
	}
	cb := makeTWndMethod(fn)
	base.SetEvent(m, 170, controlAPI(), api.MakeEventDataPtr(cb))
}

// NewControl class constructor
func NewControl(theOwner IComponent) IControl {
	r := controlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsControl(r)
}

func TControlClass() types.TClass {
	r := controlAPI().SysCallN(171)
	return types.TClass(r)
}

var (
	controlOnce   base.Once
	controlImport *imports.Imports = nil
)

func controlAPI() *imports.Imports {
	controlOnce.Do(func() {
		controlImport = api.NewDefaultImports()
		controlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TControl_Create", 0), // constructor NewControl
			/* 1 */ imports.NewTable("TControl_ManualDock", 0), // function ManualDock
			/* 2 */ imports.NewTable("TControl_ManualFloat", 0), // function ManualFloat
			/* 3 */ imports.NewTable("TControl_ReplaceDockedControl", 0), // function ReplaceDockedControl
			/* 4 */ imports.NewTable("TControl_Docked", 0), // function Docked
			/* 5 */ imports.NewTable("TControl_Dragging", 0), // function Dragging
			/* 6 */ imports.NewTable("TControl_GetAccessibleObject", 0), // function GetAccessibleObject
			/* 7 */ imports.NewTable("TControl_CreateAccessibleObject", 0), // function CreateAccessibleObject
			/* 8 */ imports.NewTable("TControl_GetSelectedChildAccessibleObject", 0), // function GetSelectedChildAccessibleObject
			/* 9 */ imports.NewTable("TControl_GetChildAccessibleObjectAtPos", 0), // function GetChildAccessibleObjectAtPos
			/* 10 */ imports.NewTable("TControl_ScaleDesignToForm", 0), // function ScaleDesignToForm
			/* 11 */ imports.NewTable("TControl_ScaleFormToDesign", 0), // function ScaleFormToDesign
			/* 12 */ imports.NewTable("TControl_Scale96ToForm", 0), // function Scale96ToForm
			/* 13 */ imports.NewTable("TControl_ScaleFormTo96", 0), // function ScaleFormTo96
			/* 14 */ imports.NewTable("TControl_Scale96ToFont", 0), // function Scale96ToFont
			/* 15 */ imports.NewTable("TControl_ScaleFontTo96", 0), // function ScaleFontTo96
			/* 16 */ imports.NewTable("TControl_ScaleScreenToFont", 0), // function ScaleScreenToFont
			/* 17 */ imports.NewTable("TControl_ScaleFontToScreen", 0), // function ScaleFontToScreen
			/* 18 */ imports.NewTable("TControl_Scale96ToScreen", 0), // function Scale96ToScreen
			/* 19 */ imports.NewTable("TControl_ScaleScreenTo96", 0), // function ScaleScreenTo96
			/* 20 */ imports.NewTable("TControl_AutoSizePhases", 0), // function AutoSizePhases
			/* 21 */ imports.NewTable("TControl_AutoSizeDelayed", 0), // function AutoSizeDelayed
			/* 22 */ imports.NewTable("TControl_AutoSizeDelayedReport", 0), // function AutoSizeDelayedReport
			/* 23 */ imports.NewTable("TControl_AutoSizeDelayedHandle", 0), // function AutoSizeDelayedHandle
			/* 24 */ imports.NewTable("TControl_AnchoredControlCount", 0), // function AnchoredControlCount
			/* 25 */ imports.NewTable("TControl_GetCanvasScaleFactor", 0), // function GetCanvasScaleFactor
			/* 26 */ imports.NewTable("TControl_GetDefaultWidth", 0), // function GetDefaultWidth
			/* 27 */ imports.NewTable("TControl_GetDefaultHeight", 0), // function GetDefaultHeight
			/* 28 */ imports.NewTable("TControl_GetDefaultColor", 0), // function GetDefaultColor
			/* 29 */ imports.NewTable("TControl_GetColorResolvingParent", 0), // function GetColorResolvingParent
			/* 30 */ imports.NewTable("TControl_GetRGBColorResolvingParent", 0), // function GetRGBColorResolvingParent
			/* 31 */ imports.NewTable("TControl_GetSidePosition", 0), // function GetSidePosition
			/* 32 */ imports.NewTable("TControl_GetAnchorsDependingOnParent", 0), // function GetAnchorsDependingOnParent
			/* 33 */ imports.NewTable("TControl_IsParentOf", 0), // function IsParentOf
			/* 34 */ imports.NewTable("TControl_GetTopParent", 0), // function GetTopParent
			/* 35 */ imports.NewTable("TControl_FindSubComponent", 0), // function FindSubComponent
			/* 36 */ imports.NewTable("TControl_IsVisible", 0), // function IsVisible
			/* 37 */ imports.NewTable("TControl_IsControlVisible", 0), // function IsControlVisible
			/* 38 */ imports.NewTable("TControl_IsEnabled", 0), // function IsEnabled
			/* 39 */ imports.NewTable("TControl_IsParentColor", 0), // function IsParentColor
			/* 40 */ imports.NewTable("TControl_IsParentFont", 0), // function IsParentFont
			/* 41 */ imports.NewTable("TControl_FormIsUpdating", 0), // function FormIsUpdating
			/* 42 */ imports.NewTable("TControl_IsProcessingPaintMsg", 0), // function IsProcessingPaintMsg
			/* 43 */ imports.NewTable("TControl_CheckChildClassAllowed", 0), // function CheckChildClassAllowed
			/* 44 */ imports.NewTable("TControl_GetTextBuf", 0), // function GetTextBuf
			/* 45 */ imports.NewTable("TControl_GetTextLen", 0), // function GetTextLen
			/* 46 */ imports.NewTable("TControl_Perform", 0), // function Perform
			/* 47 */ imports.NewTable("TControl_ScreenToClient", 0), // function ScreenToClient
			/* 48 */ imports.NewTable("TControl_ClientToScreenWithPoint", 0), // function ClientToScreenWithPoint
			/* 49 */ imports.NewTable("TControl_ClientToScreenWithRect", 0), // function ClientToScreenWithRect
			/* 50 */ imports.NewTable("TControl_ScreenToControl", 0), // function ScreenToControl
			/* 51 */ imports.NewTable("TControl_ControlToScreen", 0), // function ControlToScreen
			/* 52 */ imports.NewTable("TControl_ClientToParent", 0), // function ClientToParent
			/* 53 */ imports.NewTable("TControl_ParentToClient", 0), // function ParentToClient
			/* 54 */ imports.NewTable("TControl_GetChildrenRect", 0), // function GetChildrenRect
			/* 55 */ imports.NewTable("TControl_HandleObjectShouldBeVisible", 0), // function HandleObjectShouldBeVisible
			/* 56 */ imports.NewTable("TControl_ParentDestroyingHandle", 0), // function ParentDestroyingHandle
			/* 57 */ imports.NewTable("TControl_ParentHandlesAllocated", 0), // function ParentHandlesAllocated
			/* 58 */ imports.NewTable("TControl_HasHelp", 0), // function HasHelp
			/* 59 */ imports.NewTable("TControl_UseRightToLeftAlignment", 0), // function UseRightToLeftAlignment
			/* 60 */ imports.NewTable("TControl_UseRightToLeftReading", 0), // function UseRightToLeftReading
			/* 61 */ imports.NewTable("TControl_UseRightToLeftScrollBar", 0), // function UseRightToLeftScrollBar
			/* 62 */ imports.NewTable("TControl_IsRightToLeft", 0), // function IsRightToLeft
			/* 63 */ imports.NewTable("TControl_DragDrop", 0), // procedure DragDrop
			/* 64 */ imports.NewTable("TControl_Dock", 0), // procedure Dock
			/* 65 */ imports.NewTable("TControl_AdjustSize", 0), // procedure AdjustSize
			/* 66 */ imports.NewTable("TControl_AnchorToNeighbour", 0), // procedure AnchorToNeighbour
			/* 67 */ imports.NewTable("TControl_AnchorParallel", 0), // procedure AnchorParallel
			/* 68 */ imports.NewTable("TControl_AnchorHorizontalCenterTo", 0), // procedure AnchorHorizontalCenterTo
			/* 69 */ imports.NewTable("TControl_AnchorVerticalCenterTo", 0), // procedure AnchorVerticalCenterTo
			/* 70 */ imports.NewTable("TControl_AnchorToCompanion", 0), // procedure AnchorToCompanion
			/* 71 */ imports.NewTable("TControl_AnchorSame", 0), // procedure AnchorSame
			/* 72 */ imports.NewTable("TControl_AnchorAsAlign", 0), // procedure AnchorAsAlign
			/* 73 */ imports.NewTable("TControl_AnchorClient", 0), // procedure AnchorClient
			/* 74 */ imports.NewTable("TControl_SetBounds", 0), // procedure SetBounds
			/* 75 */ imports.NewTable("TControl_SetInitialBounds", 0), // procedure SetInitialBounds
			/* 76 */ imports.NewTable("TControl_SetBoundsKeepBase", 0), // procedure SetBoundsKeepBase
			/* 77 */ imports.NewTable("TControl_GetPreferredSize", 0), // procedure GetPreferredSize
			/* 78 */ imports.NewTable("TControl_CNPreferredSizeChanged", 0), // procedure CNPreferredSizeChanged
			/* 79 */ imports.NewTable("TControl_InvalidatePreferredSize", 0), // procedure InvalidatePreferredSize
			/* 80 */ imports.NewTable("TControl_DisableAutoSizing", 0), // procedure DisableAutoSizing
			/* 81 */ imports.NewTable("TControl_EnableAutoSizing", 0), // procedure EnableAutoSizing
			/* 82 */ imports.NewTable("TControl_UpdateBaseBounds", 0), // procedure UpdateBaseBounds
			/* 83 */ imports.NewTable("TControl_WriteLayoutDebugReport", 0), // procedure WriteLayoutDebugReport
			/* 84 */ imports.NewTable("TControl_AutoAdjustLayout", 0), // procedure AutoAdjustLayout
			/* 85 */ imports.NewTable("TControl_ShouldAutoAdjust", 0), // procedure ShouldAutoAdjust
			/* 86 */ imports.NewTable("TControl_FixDesignFontsPPI", 0), // procedure FixDesignFontsPPI
			/* 87 */ imports.NewTable("TControl_ScaleFontsPPI", 0), // procedure ScaleFontsPPI
			/* 88 */ imports.NewTable("TControl_EditingDone", 0), // procedure EditingDone
			/* 89 */ imports.NewTable("TControl_ExecuteDefaultAction", 0), // procedure ExecuteDefaultAction
			/* 90 */ imports.NewTable("TControl_ExecuteCancelAction", 0), // procedure ExecuteCancelAction
			/* 91 */ imports.NewTable("TControl_BeginDrag", 0), // procedure BeginDrag
			/* 92 */ imports.NewTable("TControl_EndDrag", 0), // procedure EndDrag
			/* 93 */ imports.NewTable("TControl_BringToFront", 0), // procedure BringToFront
			/* 94 */ imports.NewTable("TControl_Hide", 0), // procedure Hide
			/* 95 */ imports.NewTable("TControl_Refresh", 0), // procedure Refresh
			/* 96 */ imports.NewTable("TControl_Repaint", 0), // procedure Repaint
			/* 97 */ imports.NewTable("TControl_Invalidate", 0), // procedure Invalidate
			/* 98 */ imports.NewTable("TControl_CheckNewParent", 0), // procedure CheckNewParent
			/* 99 */ imports.NewTable("TControl_SendToBack", 0), // procedure SendToBack
			/* 100 */ imports.NewTable("TControl_SetTempCursor", 0), // procedure SetTempCursor
			/* 101 */ imports.NewTable("TControl_UpdateRolesForForm", 0), // procedure UpdateRolesForForm
			/* 102 */ imports.NewTable("TControl_ActiveDefaultControlChanged", 0), // procedure ActiveDefaultControlChanged
			/* 103 */ imports.NewTable("TControl_SetTextBuf", 0), // procedure SetTextBuf
			/* 104 */ imports.NewTable("TControl_Show", 0), // procedure Show
			/* 105 */ imports.NewTable("TControl_Update", 0), // procedure Update
			/* 106 */ imports.NewTable("TControl_InitiateAction", 0), // procedure InitiateAction
			/* 107 */ imports.NewTable("TControl_ShowHelp", 0), // procedure ShowHelp
			/* 108 */ imports.NewTable("TControl_AnchoredControls", 0), // property AnchoredControls
			/* 109 */ imports.NewTable("TControl_BaseBounds", 0), // property BaseBounds
			/* 110 */ imports.NewTable("TControl_ReadBounds", 0), // property ReadBounds
			/* 111 */ imports.NewTable("TControl_BaseParentClientSize", 0), // property BaseParentClientSize
			/* 112 */ imports.NewTable("TControl_AccessibleName", 0), // property AccessibleName
			/* 113 */ imports.NewTable("TControl_AccessibleDescription", 0), // property AccessibleDescription
			/* 114 */ imports.NewTable("TControl_AccessibleValue", 0), // property AccessibleValue
			/* 115 */ imports.NewTable("TControl_AccessibleRole", 0), // property AccessibleRole
			/* 116 */ imports.NewTable("TControl_Action", 0), // property Action
			/* 117 */ imports.NewTable("TControl_Align", 0), // property Align
			/* 118 */ imports.NewTable("TControl_Anchors", 0), // property Anchors
			/* 119 */ imports.NewTable("TControl_AnchorSide", 0), // property AnchorSide
			/* 120 */ imports.NewTable("TControl_AutoSize", 0), // property AutoSize
			/* 121 */ imports.NewTable("TControl_BorderSpacing", 0), // property BorderSpacing
			/* 122 */ imports.NewTable("TControl_BoundsRect", 0), // property BoundsRect
			/* 123 */ imports.NewTable("TControl_BoundsRectForNewParent", 0), // property BoundsRectForNewParent
			/* 124 */ imports.NewTable("TControl_Caption", 0), // property Caption
			/* 125 */ imports.NewTable("TControl_CaptureMouseButtons", 0), // property CaptureMouseButtons
			/* 126 */ imports.NewTable("TControl_ClientHeight", 0), // property ClientHeight
			/* 127 */ imports.NewTable("TControl_ClientOrigin", 0), // property ClientOrigin
			/* 128 */ imports.NewTable("TControl_ClientRect", 0), // property ClientRect
			/* 129 */ imports.NewTable("TControl_ClientWidth", 0), // property ClientWidth
			/* 130 */ imports.NewTable("TControl_Color", 0), // property Color
			/* 131 */ imports.NewTable("TControl_Constraints", 0), // property Constraints
			/* 132 */ imports.NewTable("TControl_ControlOrigin", 0), // property ControlOrigin
			/* 133 */ imports.NewTable("TControl_ControlState", 0), // property ControlState
			/* 134 */ imports.NewTable("TControl_ControlStyle", 0), // property ControlStyle
			/* 135 */ imports.NewTable("TControl_Enabled", 0), // property Enabled
			/* 136 */ imports.NewTable("TControl_Font", 0), // property Font
			/* 137 */ imports.NewTable("TControl_IsControl", 0), // property IsControl
			/* 138 */ imports.NewTable("TControl_MouseInClient", 0), // property MouseInClient
			/* 139 */ imports.NewTable("TControl_Parent", 0), // property Parent
			/* 140 */ imports.NewTable("TControl_PopupMenu", 0), // property PopupMenu
			/* 141 */ imports.NewTable("TControl_ShowHint", 0), // property ShowHint
			/* 142 */ imports.NewTable("TControl_Visible", 0), // property Visible
			/* 143 */ imports.NewTable("TControl_DockOrientation", 0), // property DockOrientation
			/* 144 */ imports.NewTable("TControl_Floating", 0), // property Floating
			/* 145 */ imports.NewTable("TControl_FloatingDockSiteClass", 0), // property FloatingDockSiteClass
			/* 146 */ imports.NewTable("TControl_HostDockSite", 0), // property HostDockSite
			/* 147 */ imports.NewTable("TControl_LRDockWidth", 0), // property LRDockWidth
			/* 148 */ imports.NewTable("TControl_TBDockHeight", 0), // property TBDockHeight
			/* 149 */ imports.NewTable("TControl_UndockHeight", 0), // property UndockHeight
			/* 150 */ imports.NewTable("TControl_UndockWidth", 0), // property UndockWidth
			/* 151 */ imports.NewTable("TControl_BiDiMode", 0), // property BiDiMode
			/* 152 */ imports.NewTable("TControl_ParentBiDiMode", 0), // property ParentBiDiMode
			/* 153 */ imports.NewTable("TControl_AnchorSideLeft", 0), // property AnchorSideLeft
			/* 154 */ imports.NewTable("TControl_AnchorSideTop", 0), // property AnchorSideTop
			/* 155 */ imports.NewTable("TControl_AnchorSideRight", 0), // property AnchorSideRight
			/* 156 */ imports.NewTable("TControl_AnchorSideBottom", 0), // property AnchorSideBottom
			/* 157 */ imports.NewTable("TControl_Cursor", 0), // property Cursor
			/* 158 */ imports.NewTable("TControl_Left", 0), // property Left
			/* 159 */ imports.NewTable("TControl_Height", 0), // property Height
			/* 160 */ imports.NewTable("TControl_Hint", 0), // property Hint
			/* 161 */ imports.NewTable("TControl_Top", 0), // property Top
			/* 162 */ imports.NewTable("TControl_Width", 0), // property Width
			/* 163 */ imports.NewTable("TControl_HelpType", 0), // property HelpType
			/* 164 */ imports.NewTable("TControl_HelpKeyword", 0), // property HelpKeyword
			/* 165 */ imports.NewTable("TControl_HelpContext", 0), // property HelpContext
			/* 166 */ imports.NewTable("TControl_OnChangeBounds", 0), // event OnChangeBounds
			/* 167 */ imports.NewTable("TControl_OnClick", 0), // event OnClick
			/* 168 */ imports.NewTable("TControl_OnResize", 0), // event OnResize
			/* 169 */ imports.NewTable("TControl_OnShowHint", 0), // event OnShowHint
			/* 170 */ imports.NewTable("TControl_WindowProc", 0), // event WindowProc
			/* 171 */ imports.NewTable("TControl_TClass", 0), // function TControlClass
		}
	})
	return controlImport
}
