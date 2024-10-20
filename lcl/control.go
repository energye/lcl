//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IControl Parent: ILCLComponent
type IControl interface {
	ILCLComponent
	AnchoredControls(Index int32) IControl                                                                         // property
	BaseBounds() (resultRect TRect)                                                                                // property
	ReadBounds() (resultRect TRect)                                                                                // property
	BaseParentClientSize() (resultSize TSize)                                                                      // property
	AccessibleName() string                                                                                        // property
	SetAccessibleName(AValue string)                                                                               // property
	AccessibleDescription() string                                                                                 // property
	SetAccessibleDescription(AValue string)                                                                        // property
	AccessibleValue() string                                                                                       // property
	SetAccessibleValue(AValue string)                                                                              // property
	AccessibleRole() TLazAccessibilityRole                                                                         // property
	SetAccessibleRole(AValue TLazAccessibilityRole)                                                                // property
	Action() IBasicAction                                                                                          // property
	SetAction(AValue IBasicAction)                                                                                 // property
	Align() TAlign                                                                                                 // property
	SetAlign(AValue TAlign)                                                                                        // property
	Anchors() TAnchors                                                                                             // property
	SetAnchors(AValue TAnchors)                                                                                    // property
	AnchorSide(Kind TAnchorKind) IAnchorSide                                                                       // property
	AutoSize() bool                                                                                                // property
	SetAutoSize(AValue bool)                                                                                       // property
	BorderSpacing() IControlBorderSpacing                                                                          // property
	SetBorderSpacing(AValue IControlBorderSpacing)                                                                 // property
	BoundsRect() (resultRect TRect)                                                                                // property
	SetBoundsRect(AValue *TRect)                                                                                   // property
	BoundsRectForNewParent() (resultRect TRect)                                                                    // property
	SetBoundsRectForNewParent(AValue *TRect)                                                                       // property
	Caption() string                                                                                               // property
	SetCaption(AValue string)                                                                                      // property
	CaptureMouseButtons() TCaptureMouseButtons                                                                     // property
	SetCaptureMouseButtons(AValue TCaptureMouseButtons)                                                            // property
	ClientHeight() int32                                                                                           // property
	SetClientHeight(AValue int32)                                                                                  // property
	ClientOrigin() (resultPoint TPoint)                                                                            // property
	ClientRect() (resultRect TRect)                                                                                // property
	ClientWidth() int32                                                                                            // property
	SetClientWidth(AValue int32)                                                                                   // property
	Color() TColor                                                                                                 // property
	SetColor(AValue TColor)                                                                                        // property
	Constraints() ISizeConstraints                                                                                 // property
	SetConstraints(AValue ISizeConstraints)                                                                        // property
	ControlOrigin() (resultPoint TPoint)                                                                           // property
	ControlState() TControlState                                                                                   // property
	SetControlState(AValue TControlState)                                                                          // property
	ControlStyle() TControlStyle                                                                                   // property
	SetControlStyle(AValue TControlStyle)                                                                          // property
	Enabled() bool                                                                                                 // property
	SetEnabled(AValue bool)                                                                                        // property
	Font() IFont                                                                                                   // property
	SetFont(AValue IFont)                                                                                          // property
	IsControl() bool                                                                                               // property
	SetIsControl(AValue bool)                                                                                      // property
	MouseInClient() bool                                                                                           // property
	Parent() IWinControl                                                                                           // property
	SetParent(AValue IWinControl)                                                                                  // property
	PopupMenu() IPopupMenu                                                                                         // property
	SetPopupMenu(AValue IPopupMenu)                                                                                // property
	ShowHint() bool                                                                                                // property
	SetShowHint(AValue bool)                                                                                       // property
	Visible() bool                                                                                                 // property
	SetVisible(AValue bool)                                                                                        // property
	DockOrientation() TDockOrientation                                                                             // property
	SetDockOrientation(AValue TDockOrientation)                                                                    // property
	Floating() bool                                                                                                // property
	FloatingDockSiteClass() TWinControlClass                                                                       // property
	SetFloatingDockSiteClass(AValue TWinControlClass)                                                              // property
	HostDockSite() IWinControl                                                                                     // property
	SetHostDockSite(AValue IWinControl)                                                                            // property
	LRDockWidth() int32                                                                                            // property
	SetLRDockWidth(AValue int32)                                                                                   // property
	TBDockHeight() int32                                                                                           // property
	SetTBDockHeight(AValue int32)                                                                                  // property
	UndockHeight() int32                                                                                           // property
	SetUndockHeight(AValue int32)                                                                                  // property
	UndockWidth() int32                                                                                            // property
	SetUndockWidth(AValue int32)                                                                                   // property
	BiDiMode() TBiDiMode                                                                                           // property
	SetBiDiMode(AValue TBiDiMode)                                                                                  // property
	ParentBiDiMode() bool                                                                                          // property
	SetParentBiDiMode(AValue bool)                                                                                 // property
	AnchorSideLeft() IAnchorSide                                                                                   // property
	SetAnchorSideLeft(AValue IAnchorSide)                                                                          // property
	AnchorSideTop() IAnchorSide                                                                                    // property
	SetAnchorSideTop(AValue IAnchorSide)                                                                           // property
	AnchorSideRight() IAnchorSide                                                                                  // property
	SetAnchorSideRight(AValue IAnchorSide)                                                                         // property
	AnchorSideBottom() IAnchorSide                                                                                 // property
	SetAnchorSideBottom(AValue IAnchorSide)                                                                        // property
	Cursor() TCursor                                                                                               // property
	SetCursor(AValue TCursor)                                                                                      // property
	Left() int32                                                                                                   // property
	SetLeft(AValue int32)                                                                                          // property
	Height() int32                                                                                                 // property
	SetHeight(AValue int32)                                                                                        // property
	Hint() string                                                                                                  // property
	SetHint(AValue string)                                                                                         // property
	Top() int32                                                                                                    // property
	SetTop(AValue int32)                                                                                           // property
	Width() int32                                                                                                  // property
	SetWidth(AValue int32)                                                                                         // property
	HelpType() THelpType                                                                                           // property
	SetHelpType(AValue THelpType)                                                                                  // property
	HelpKeyword() string                                                                                           // property
	SetHelpKeyword(AValue string)                                                                                  // property
	HelpContext() THelpContext                                                                                     // property
	SetHelpContext(AValue THelpContext)                                                                            // property
	ManualDock(NewDockSite IWinControl, DropControl IControl, ControlSide TAlign, KeepDockSiteSize bool) bool      // function
	ManualFloat(TheScreenRect *TRect, KeepDockSiteSize bool) bool                                                  // function
	ReplaceDockedControl(Control IControl, NewDockSite IWinControl, DropControl IControl, ControlSide TAlign) bool // function
	Docked() bool                                                                                                  // function
	Dragging() bool                                                                                                // function
	GetAccessibleObject() ILazAccessibleObject                                                                     // function
	CreateAccessibleObject() ILazAccessibleObject                                                                  // function
	GetSelectedChildAccessibleObject() ILazAccessibleObject                                                        // function
	GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject                                               // function
	ScaleDesignToForm(ASize int32) int32                                                                           // function
	ScaleFormToDesign(ASize int32) int32                                                                           // function
	Scale96ToForm(ASize int32) int32                                                                               // function
	ScaleFormTo96(ASize int32) int32                                                                               // function
	Scale96ToFont(ASize int32) int32                                                                               // function
	ScaleFontTo96(ASize int32) int32                                                                               // function
	ScaleScreenToFont(ASize int32) int32                                                                           // function
	ScaleFontToScreen(ASize int32) int32                                                                           // function
	Scale96ToScreen(ASize int32) int32                                                                             // function
	ScaleScreenTo96(ASize int32) int32                                                                             // function
	AutoSizePhases() TControlAutoSizePhases                                                                        // function
	AutoSizeDelayed() bool                                                                                         // function
	AutoSizeDelayedReport() string                                                                                 // function
	AutoSizeDelayedHandle() bool                                                                                   // function
	AnchoredControlCount() int32                                                                                   // function
	GetCanvasScaleFactor() (resultDouble float64)                                                                  // function
	GetDefaultWidth() int32                                                                                        // function
	GetDefaultHeight() int32                                                                                       // function
	GetDefaultColor(DefaultColorType TDefaultColorType) TColor                                                     // function
	GetColorResolvingParent() TColor                                                                               // function
	GetRGBColorResolvingParent() TColor                                                                            // function
	GetSidePosition(Side TAnchorKind) int32                                                                        // function
	GetAnchorsDependingOnParent(WithNormalAnchors bool) TAnchors                                                   // function
	IsParentOf(AControl IControl) bool                                                                             // function
	GetTopParent() IControl                                                                                        // function
	FindSubComponent(AName string) IComponent                                                                      // function
	IsVisible() bool                                                                                               // function
	IsControlVisible() bool                                                                                        // function
	IsEnabled() bool                                                                                               // function
	IsParentColor() bool                                                                                           // function
	IsParentFont() bool                                                                                            // function
	FormIsUpdating() bool                                                                                          // function
	IsProcessingPaintMsg() bool                                                                                    // function
	CheckChildClassAllowed(ChildClass TClass, ExceptionOnInvalid bool) bool                                        // function
	GetTextBuf(Buffer *string, BufSize int32) int32                                                                // function
	GetTextLen() int32                                                                                             // function
	Perform(Msg uint32, WParam WPARAM, LParam LPARAM) LRESULT                                                      // function
	ScreenToClient(APoint *TPoint) (resultPoint TPoint)                                                            // function
	ClientToScreen(APoint *TPoint) (resultPoint TPoint)                                                            // function
	ClientToScreen1(ARect *TRect) (resultRect TRect)                                                               // function
	ScreenToControl(APoint *TPoint) (resultPoint TPoint)                                                           // function
	ControlToScreen(APoint *TPoint) (resultPoint TPoint)                                                           // function
	ClientToParent(Point *TPoint, AParent IWinControl) (resultPoint TPoint)                                        // function
	ParentToClient(Point *TPoint, AParent IWinControl) (resultPoint TPoint)                                        // function
	GetChildrenRect(Scrolled bool) (resultRect TRect)                                                              // function
	HandleObjectShouldBeVisible() bool                                                                             // function
	ParentDestroyingHandle() bool                                                                                  // function
	ParentHandlesAllocated() bool                                                                                  // function
	HasHelp() bool                                                                                                 // function
	UseRightToLeftAlignment() bool                                                                                 // function
	UseRightToLeftReading() bool                                                                                   // function
	UseRightToLeftScrollBar() bool                                                                                 // function
	IsRightToLeft() bool                                                                                           // function
	DragDrop(Source IObject, X, Y int32)                                                                           // procedure
	Dock(NewDockSite IWinControl, ARect *TRect)                                                                    // procedure
	AdjustSize()                                                                                                   // procedure
	AnchorToNeighbour(Side TAnchorKind, Space TSpacingSize, Sibling IControl)                                      // procedure
	AnchorParallel(Side TAnchorKind, Space TSpacingSize, Sibling IControl)                                         // procedure
	AnchorHorizontalCenterTo(Sibling IControl)                                                                     // procedure
	AnchorVerticalCenterTo(Sibling IControl)                                                                       // procedure
	AnchorToCompanion(Side TAnchorKind, Space TSpacingSize, Sibling IControl, FreeCompositeSide bool)              // procedure
	AnchorSame(Side TAnchorKind, Sibling IControl)                                                                 // procedure
	AnchorAsAlign(TheAlign TAlign, Space TSpacingSize)                                                             // procedure
	AnchorClient(Space TSpacingSize)                                                                               // procedure
	SetBounds(aLeft, aTop, aWidth, aHeight int32)                                                                  // procedure
	SetInitialBounds(aLeft, aTop, aWidth, aHeight int32)                                                           // procedure
	SetBoundsKeepBase(aLeft, aTop, aWidth, aHeight int32)                                                          // procedure
	GetPreferredSize(PreferredWidth, PreferredHeight *int32, Raw bool, WithThemeSpace bool)                        // procedure
	CNPreferredSizeChanged()                                                                                       // procedure
	InvalidatePreferredSize()                                                                                      // procedure
	DisableAutoSizing()                                                                                            // procedure
	EnableAutoSizing()                                                                                             // procedure
	UpdateBaseBounds(StoreBounds, StoreParentClientSize, UseLoadedValues bool)                                     // procedure
	WriteLayoutDebugReport(Prefix string)                                                                          // procedure
	AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth int32)          // procedure
	ShouldAutoAdjust(AWidth, AHeight *bool)                                                                        // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)                                                                        // procedure
	ScaleFontsPPI(AToPPI int32, AProportion float64)                                                               // procedure
	EditingDone()                                                                                                  // procedure
	ExecuteDefaultAction()                                                                                         // procedure
	ExecuteCancelAction()                                                                                          // procedure
	BeginDrag(Immediate bool, Threshold int32)                                                                     // procedure
	EndDrag(Drop bool)                                                                                             // procedure
	BringToFront()                                                                                                 // procedure
	Hide()                                                                                                         // procedure
	Refresh()                                                                                                      // procedure
	Repaint()                                                                                                      // procedure
	Invalidate()                                                                                                   // procedure
	CheckNewParent(AParent IWinControl)                                                                            // procedure
	SendToBack()                                                                                                   // procedure
	SetTempCursor(Value TCursor)                                                                                   // procedure
	UpdateRolesForForm()                                                                                           // procedure
	ActiveDefaultControlChanged(NewControl IControl)                                                               // procedure
	SetTextBuf(Buffer string)                                                                                      // procedure
	Show()                                                                                                         // procedure
	Update()                                                                                                       // procedure
	InitiateAction()                                                                                               // procedure
	ShowHelp()                                                                                                     // procedure
	SetOnChangeBounds(fn TNotifyEvent)                                                                             // property event
	SetOnClick(fn TNotifyEvent)                                                                                    // property event
	SetOnResize(fn TNotifyEvent)                                                                                   // property event
	SetOnShowHint(fn TControlShowHintEvent)                                                                        // property event
}

// TControl Parent: TLCLComponent
type TControl struct {
	TLCLComponent
	changeBoundsPtr uintptr
	clickPtr        uintptr
	resizePtr       uintptr
	showHintPtr     uintptr
}

func NewControl(TheOwner IComponent) IControl {
	r1 := controlImportAPI().SysCallN(57, GetObjectUintptr(TheOwner))
	return AsControl(r1)
}

func (m *TControl) AnchoredControls(Index int32) IControl {
	r1 := controlImportAPI().SysCallN(22, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TControl) BaseBounds() (resultRect TRect) {
	controlImportAPI().SysCallN(30, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) ReadBounds() (resultRect TRect) {
	controlImportAPI().SysCallN(127, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) BaseParentClientSize() (resultSize TSize) {
	controlImportAPI().SysCallN(31, m.Instance(), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TControl) AccessibleName() string {
	r1 := controlImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetAccessibleName(AValue string) {
	controlImportAPI().SysCallN(1, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) AccessibleDescription() string {
	r1 := controlImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetAccessibleDescription(AValue string) {
	controlImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) AccessibleValue() string {
	r1 := controlImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetAccessibleValue(AValue string) {
	controlImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) AccessibleRole() TLazAccessibilityRole {
	r1 := controlImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TLazAccessibilityRole(r1)
}

func (m *TControl) SetAccessibleRole(AValue TLazAccessibilityRole) {
	controlImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Action() IBasicAction {
	r1 := controlImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return AsBasicAction(r1)
}

func (m *TControl) SetAction(AValue IBasicAction) {
	controlImportAPI().SysCallN(4, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) Align() TAlign {
	r1 := controlImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TAlign(r1)
}

func (m *TControl) SetAlign(AValue TAlign) {
	controlImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Anchors() TAnchors {
	r1 := controlImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return TAnchors(r1)
}

func (m *TControl) SetAnchors(AValue TAnchors) {
	controlImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) AnchorSide(Kind TAnchorKind) IAnchorSide {
	r1 := controlImportAPI().SysCallN(13, m.Instance(), uintptr(Kind))
	return AsAnchorSide(r1)
}

func (m *TControl) AutoSize() bool {
	r1 := controlImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetAutoSize(AValue bool) {
	controlImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) BorderSpacing() IControlBorderSpacing {
	r1 := controlImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return AsControlBorderSpacing(r1)
}

func (m *TControl) SetBorderSpacing(AValue IControlBorderSpacing) {
	controlImportAPI().SysCallN(34, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) BoundsRect() (resultRect TRect) {
	controlImportAPI().SysCallN(35, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) SetBoundsRect(AValue *TRect) {
	controlImportAPI().SysCallN(35, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TControl) BoundsRectForNewParent() (resultRect TRect) {
	controlImportAPI().SysCallN(36, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) SetBoundsRectForNewParent(AValue *TRect) {
	controlImportAPI().SysCallN(36, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TControl) Caption() string {
	r1 := controlImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetCaption(AValue string) {
	controlImportAPI().SysCallN(39, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) CaptureMouseButtons() TCaptureMouseButtons {
	r1 := controlImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return TCaptureMouseButtons(r1)
}

func (m *TControl) SetCaptureMouseButtons(AValue TCaptureMouseButtons) {
	controlImportAPI().SysCallN(40, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ClientHeight() int32 {
	r1 := controlImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetClientHeight(AValue int32) {
	controlImportAPI().SysCallN(44, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ClientOrigin() (resultPoint TPoint) {
	controlImportAPI().SysCallN(45, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientRect() (resultRect TRect) {
	controlImportAPI().SysCallN(46, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) ClientWidth() int32 {
	r1 := controlImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetClientWidth(AValue int32) {
	controlImportAPI().SysCallN(50, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Color() TColor {
	r1 := controlImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TControl) SetColor(AValue TColor) {
	controlImportAPI().SysCallN(51, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Constraints() ISizeConstraints {
	r1 := controlImportAPI().SysCallN(52, 0, m.Instance(), 0)
	return AsSizeConstraints(r1)
}

func (m *TControl) SetConstraints(AValue ISizeConstraints) {
	controlImportAPI().SysCallN(52, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) ControlOrigin() (resultPoint TPoint) {
	controlImportAPI().SysCallN(53, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ControlState() TControlState {
	r1 := controlImportAPI().SysCallN(54, 0, m.Instance(), 0)
	return TControlState(r1)
}

func (m *TControl) SetControlState(AValue TControlState) {
	controlImportAPI().SysCallN(54, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ControlStyle() TControlStyle {
	r1 := controlImportAPI().SysCallN(55, 0, m.Instance(), 0)
	return TControlStyle(r1)
}

func (m *TControl) SetControlStyle(AValue TControlStyle) {
	controlImportAPI().SysCallN(55, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Enabled() bool {
	r1 := controlImportAPI().SysCallN(68, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetEnabled(AValue bool) {
	controlImportAPI().SysCallN(68, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) Font() IFont {
	r1 := controlImportAPI().SysCallN(76, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TControl) SetFont(AValue IFont) {
	controlImportAPI().SysCallN(76, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) IsControl() bool {
	r1 := controlImportAPI().SysCallN(106, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetIsControl(AValue bool) {
	controlImportAPI().SysCallN(106, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) MouseInClient() bool {
	r1 := controlImportAPI().SysCallN(119, m.Instance())
	return GoBool(r1)
}

func (m *TControl) Parent() IWinControl {
	r1 := controlImportAPI().SysCallN(120, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TControl) SetParent(AValue IWinControl) {
	controlImportAPI().SysCallN(120, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) PopupMenu() IPopupMenu {
	r1 := controlImportAPI().SysCallN(126, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TControl) SetPopupMenu(AValue IPopupMenu) {
	controlImportAPI().SysCallN(126, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) ShowHint() bool {
	r1 := controlImportAPI().SysCallN(157, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetShowHint(AValue bool) {
	controlImportAPI().SysCallN(157, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) Visible() bool {
	r1 := controlImportAPI().SysCallN(168, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetVisible(AValue bool) {
	controlImportAPI().SysCallN(168, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) DockOrientation() TDockOrientation {
	r1 := controlImportAPI().SysCallN(62, 0, m.Instance(), 0)
	return TDockOrientation(r1)
}

func (m *TControl) SetDockOrientation(AValue TDockOrientation) {
	controlImportAPI().SysCallN(62, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Floating() bool {
	r1 := controlImportAPI().SysCallN(74, m.Instance())
	return GoBool(r1)
}

func (m *TControl) FloatingDockSiteClass() TWinControlClass {
	r1 := controlImportAPI().SysCallN(75, 0, m.Instance(), 0)
	return TWinControlClass(r1)
}

func (m *TControl) SetFloatingDockSiteClass(AValue TWinControlClass) {
	controlImportAPI().SysCallN(75, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) HostDockSite() IWinControl {
	r1 := controlImportAPI().SysCallN(102, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TControl) SetHostDockSite(AValue IWinControl) {
	controlImportAPI().SysCallN(102, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) LRDockWidth() int32 {
	r1 := controlImportAPI().SysCallN(115, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetLRDockWidth(AValue int32) {
	controlImportAPI().SysCallN(115, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) TBDockHeight() int32 {
	r1 := controlImportAPI().SysCallN(158, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetTBDockHeight(AValue int32) {
	controlImportAPI().SysCallN(158, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) UndockHeight() int32 {
	r1 := controlImportAPI().SysCallN(160, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetUndockHeight(AValue int32) {
	controlImportAPI().SysCallN(160, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) UndockWidth() int32 {
	r1 := controlImportAPI().SysCallN(161, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetUndockWidth(AValue int32) {
	controlImportAPI().SysCallN(161, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) BiDiMode() TBiDiMode {
	r1 := controlImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TControl) SetBiDiMode(AValue TBiDiMode) {
	controlImportAPI().SysCallN(33, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ParentBiDiMode() bool {
	r1 := controlImportAPI().SysCallN(121, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetParentBiDiMode(AValue bool) {
	controlImportAPI().SysCallN(121, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) AnchorSideLeft() IAnchorSide {
	r1 := controlImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideLeft(AValue IAnchorSide) {
	controlImportAPI().SysCallN(15, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) AnchorSideTop() IAnchorSide {
	r1 := controlImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideTop(AValue IAnchorSide) {
	controlImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) AnchorSideRight() IAnchorSide {
	r1 := controlImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideRight(AValue IAnchorSide) {
	controlImportAPI().SysCallN(16, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) AnchorSideBottom() IAnchorSide {
	r1 := controlImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideBottom(AValue IAnchorSide) {
	controlImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) Cursor() TCursor {
	r1 := controlImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TControl) SetCursor(AValue TCursor) {
	controlImportAPI().SysCallN(59, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Left() int32 {
	r1 := controlImportAPI().SysCallN(116, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetLeft(AValue int32) {
	controlImportAPI().SysCallN(116, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Height() int32 {
	r1 := controlImportAPI().SysCallN(96, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetHeight(AValue int32) {
	controlImportAPI().SysCallN(96, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Hint() string {
	r1 := controlImportAPI().SysCallN(101, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetHint(AValue string) {
	controlImportAPI().SysCallN(101, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) Top() int32 {
	r1 := controlImportAPI().SysCallN(159, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetTop(AValue int32) {
	controlImportAPI().SysCallN(159, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Width() int32 {
	r1 := controlImportAPI().SysCallN(169, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetWidth(AValue int32) {
	controlImportAPI().SysCallN(169, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) HelpType() THelpType {
	r1 := controlImportAPI().SysCallN(99, 0, m.Instance(), 0)
	return THelpType(r1)
}

func (m *TControl) SetHelpType(AValue THelpType) {
	controlImportAPI().SysCallN(99, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) HelpKeyword() string {
	r1 := controlImportAPI().SysCallN(98, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetHelpKeyword(AValue string) {
	controlImportAPI().SysCallN(98, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) HelpContext() THelpContext {
	r1 := controlImportAPI().SysCallN(97, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TControl) SetHelpContext(AValue THelpContext) {
	controlImportAPI().SysCallN(97, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ManualDock(NewDockSite IWinControl, DropControl IControl, ControlSide TAlign, KeepDockSiteSize bool) bool {
	r1 := controlImportAPI().SysCallN(117, m.Instance(), GetObjectUintptr(NewDockSite), GetObjectUintptr(DropControl), uintptr(ControlSide), PascalBool(KeepDockSiteSize))
	return GoBool(r1)
}

func (m *TControl) ManualFloat(TheScreenRect *TRect, KeepDockSiteSize bool) bool {
	r1 := controlImportAPI().SysCallN(118, m.Instance(), uintptr(unsafePointer(TheScreenRect)), PascalBool(KeepDockSiteSize))
	return GoBool(r1)
}

func (m *TControl) ReplaceDockedControl(Control IControl, NewDockSite IWinControl, DropControl IControl, ControlSide TAlign) bool {
	r1 := controlImportAPI().SysCallN(130, m.Instance(), GetObjectUintptr(Control), GetObjectUintptr(NewDockSite), GetObjectUintptr(DropControl), uintptr(ControlSide))
	return GoBool(r1)
}

func (m *TControl) Docked() bool {
	r1 := controlImportAPI().SysCallN(63, m.Instance())
	return GoBool(r1)
}

func (m *TControl) Dragging() bool {
	r1 := controlImportAPI().SysCallN(65, m.Instance())
	return GoBool(r1)
}

func (m *TControl) GetAccessibleObject() ILazAccessibleObject {
	r1 := controlImportAPI().SysCallN(78, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TControl) CreateAccessibleObject() ILazAccessibleObject {
	r1 := controlImportAPI().SysCallN(58, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TControl) GetSelectedChildAccessibleObject() ILazAccessibleObject {
	r1 := controlImportAPI().SysCallN(89, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TControl) GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject {
	r1 := controlImportAPI().SysCallN(81, m.Instance(), uintptr(unsafePointer(APos)))
	return AsLazAccessibleObject(r1)
}

func (m *TControl) ScaleDesignToForm(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(134, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFormToDesign(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(139, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) Scale96ToForm(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(132, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFormTo96(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(138, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) Scale96ToFont(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(131, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFontTo96(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(135, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleScreenToFont(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(141, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFontToScreen(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(136, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) Scale96ToScreen(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(133, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleScreenTo96(ASize int32) int32 {
	r1 := controlImportAPI().SysCallN(140, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) AutoSizePhases() TControlAutoSizePhases {
	r1 := controlImportAPI().SysCallN(29, m.Instance())
	return TControlAutoSizePhases(r1)
}

func (m *TControl) AutoSizeDelayed() bool {
	r1 := controlImportAPI().SysCallN(26, m.Instance())
	return GoBool(r1)
}

func (m *TControl) AutoSizeDelayedReport() string {
	r1 := controlImportAPI().SysCallN(28, m.Instance())
	return GoStr(r1)
}

func (m *TControl) AutoSizeDelayedHandle() bool {
	r1 := controlImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TControl) AnchoredControlCount() int32 {
	r1 := controlImportAPI().SysCallN(21, m.Instance())
	return int32(r1)
}

func (m *TControl) GetCanvasScaleFactor() (resultDouble float64) {
	controlImportAPI().SysCallN(80, m.Instance(), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TControl) GetDefaultWidth() int32 {
	r1 := controlImportAPI().SysCallN(86, m.Instance())
	return int32(r1)
}

func (m *TControl) GetDefaultHeight() int32 {
	r1 := controlImportAPI().SysCallN(85, m.Instance())
	return int32(r1)
}

func (m *TControl) GetDefaultColor(DefaultColorType TDefaultColorType) TColor {
	r1 := controlImportAPI().SysCallN(84, m.Instance(), uintptr(DefaultColorType))
	return TColor(r1)
}

func (m *TControl) GetColorResolvingParent() TColor {
	r1 := controlImportAPI().SysCallN(83, m.Instance())
	return TColor(r1)
}

func (m *TControl) GetRGBColorResolvingParent() TColor {
	r1 := controlImportAPI().SysCallN(88, m.Instance())
	return TColor(r1)
}

func (m *TControl) GetSidePosition(Side TAnchorKind) int32 {
	r1 := controlImportAPI().SysCallN(90, m.Instance(), uintptr(Side))
	return int32(r1)
}

func (m *TControl) GetAnchorsDependingOnParent(WithNormalAnchors bool) TAnchors {
	r1 := controlImportAPI().SysCallN(79, m.Instance(), PascalBool(WithNormalAnchors))
	return TAnchors(r1)
}

func (m *TControl) IsParentOf(AControl IControl) bool {
	r1 := controlImportAPI().SysCallN(111, m.Instance(), GetObjectUintptr(AControl))
	return GoBool(r1)
}

func (m *TControl) GetTopParent() IControl {
	r1 := controlImportAPI().SysCallN(93, m.Instance())
	return AsControl(r1)
}

func (m *TControl) FindSubComponent(AName string) IComponent {
	r1 := controlImportAPI().SysCallN(72, m.Instance(), PascalStr(AName))
	return AsComponent(r1)
}

func (m *TControl) IsVisible() bool {
	r1 := controlImportAPI().SysCallN(114, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsControlVisible() bool {
	r1 := controlImportAPI().SysCallN(107, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsEnabled() bool {
	r1 := controlImportAPI().SysCallN(108, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsParentColor() bool {
	r1 := controlImportAPI().SysCallN(109, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsParentFont() bool {
	r1 := controlImportAPI().SysCallN(110, m.Instance())
	return GoBool(r1)
}

func (m *TControl) FormIsUpdating() bool {
	r1 := controlImportAPI().SysCallN(77, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsProcessingPaintMsg() bool {
	r1 := controlImportAPI().SysCallN(112, m.Instance())
	return GoBool(r1)
}

func (m *TControl) CheckChildClassAllowed(ChildClass TClass, ExceptionOnInvalid bool) bool {
	r1 := controlImportAPI().SysCallN(41, m.Instance(), uintptr(ChildClass), PascalBool(ExceptionOnInvalid))
	return GoBool(r1)
}

func (m *TControl) GetTextBuf(Buffer *string, BufSize int32) int32 {
	r1 := sysCallGetTextBuf(controlImportAPI(), 91, m.Instance(), Buffer, BufSize)
	return int32(r1)
}

func (m *TControl) GetTextLen() int32 {
	r1 := controlImportAPI().SysCallN(92, m.Instance())
	return int32(r1)
}

func (m *TControl) Perform(Msg uint32, WParam WPARAM, LParam LPARAM) LRESULT {
	r1 := controlImportAPI().SysCallN(125, m.Instance(), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return LRESULT(r1)
}

func (m *TControl) ScreenToClient(APoint *TPoint) (resultPoint TPoint) {
	controlImportAPI().SysCallN(142, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientToScreen(APoint *TPoint) (resultPoint TPoint) {
	controlImportAPI().SysCallN(48, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientToScreen1(ARect *TRect) (resultRect TRect) {
	controlImportAPI().SysCallN(49, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) ScreenToControl(APoint *TPoint) (resultPoint TPoint) {
	controlImportAPI().SysCallN(143, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ControlToScreen(APoint *TPoint) (resultPoint TPoint) {
	controlImportAPI().SysCallN(56, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientToParent(Point *TPoint, AParent IWinControl) (resultPoint TPoint) {
	controlImportAPI().SysCallN(47, m.Instance(), uintptr(unsafePointer(Point)), GetObjectUintptr(AParent), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ParentToClient(Point *TPoint, AParent IWinControl) (resultPoint TPoint) {
	controlImportAPI().SysCallN(124, m.Instance(), uintptr(unsafePointer(Point)), GetObjectUintptr(AParent), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) GetChildrenRect(Scrolled bool) (resultRect TRect) {
	controlImportAPI().SysCallN(82, m.Instance(), PascalBool(Scrolled), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) HandleObjectShouldBeVisible() bool {
	r1 := controlImportAPI().SysCallN(94, m.Instance())
	return GoBool(r1)
}

func (m *TControl) ParentDestroyingHandle() bool {
	r1 := controlImportAPI().SysCallN(122, m.Instance())
	return GoBool(r1)
}

func (m *TControl) ParentHandlesAllocated() bool {
	r1 := controlImportAPI().SysCallN(123, m.Instance())
	return GoBool(r1)
}

func (m *TControl) HasHelp() bool {
	r1 := controlImportAPI().SysCallN(95, m.Instance())
	return GoBool(r1)
}

func (m *TControl) UseRightToLeftAlignment() bool {
	r1 := controlImportAPI().SysCallN(165, m.Instance())
	return GoBool(r1)
}

func (m *TControl) UseRightToLeftReading() bool {
	r1 := controlImportAPI().SysCallN(166, m.Instance())
	return GoBool(r1)
}

func (m *TControl) UseRightToLeftScrollBar() bool {
	r1 := controlImportAPI().SysCallN(167, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsRightToLeft() bool {
	r1 := controlImportAPI().SysCallN(113, m.Instance())
	return GoBool(r1)
}

func ControlClass() TClass {
	ret := controlImportAPI().SysCallN(43)
	return TClass(ret)
}

func (m *TControl) DragDrop(Source IObject, X, Y int32) {
	controlImportAPI().SysCallN(64, m.Instance(), GetObjectUintptr(Source), uintptr(X), uintptr(Y))
}

func (m *TControl) Dock(NewDockSite IWinControl, ARect *TRect) {
	controlImportAPI().SysCallN(61, m.Instance(), GetObjectUintptr(NewDockSite), uintptr(unsafePointer(ARect)))
}

func (m *TControl) AdjustSize() {
	controlImportAPI().SysCallN(6, m.Instance())
}

func (m *TControl) AnchorToNeighbour(Side TAnchorKind, Space TSpacingSize, Sibling IControl) {
	controlImportAPI().SysCallN(19, m.Instance(), uintptr(Side), uintptr(Space), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorParallel(Side TAnchorKind, Space TSpacingSize, Sibling IControl) {
	controlImportAPI().SysCallN(11, m.Instance(), uintptr(Side), uintptr(Space), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorHorizontalCenterTo(Sibling IControl) {
	controlImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorVerticalCenterTo(Sibling IControl) {
	controlImportAPI().SysCallN(20, m.Instance(), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorToCompanion(Side TAnchorKind, Space TSpacingSize, Sibling IControl, FreeCompositeSide bool) {
	controlImportAPI().SysCallN(18, m.Instance(), uintptr(Side), uintptr(Space), GetObjectUintptr(Sibling), PascalBool(FreeCompositeSide))
}

func (m *TControl) AnchorSame(Side TAnchorKind, Sibling IControl) {
	controlImportAPI().SysCallN(12, m.Instance(), uintptr(Side), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorAsAlign(TheAlign TAlign, Space TSpacingSize) {
	controlImportAPI().SysCallN(8, m.Instance(), uintptr(TheAlign), uintptr(Space))
}

func (m *TControl) AnchorClient(Space TSpacingSize) {
	controlImportAPI().SysCallN(9, m.Instance(), uintptr(Space))
}

func (m *TControl) SetBounds(aLeft, aTop, aWidth, aHeight int32) {
	controlImportAPI().SysCallN(145, m.Instance(), uintptr(aLeft), uintptr(aTop), uintptr(aWidth), uintptr(aHeight))
}

func (m *TControl) SetInitialBounds(aLeft, aTop, aWidth, aHeight int32) {
	controlImportAPI().SysCallN(147, m.Instance(), uintptr(aLeft), uintptr(aTop), uintptr(aWidth), uintptr(aHeight))
}

func (m *TControl) SetBoundsKeepBase(aLeft, aTop, aWidth, aHeight int32) {
	controlImportAPI().SysCallN(146, m.Instance(), uintptr(aLeft), uintptr(aTop), uintptr(aWidth), uintptr(aHeight))
}

func (m *TControl) GetPreferredSize(PreferredWidth, PreferredHeight *int32, Raw bool, WithThemeSpace bool) {
	var result0 uintptr
	var result1 uintptr
	controlImportAPI().SysCallN(87, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)), PascalBool(Raw), PascalBool(WithThemeSpace))
	*PreferredWidth = int32(result0)
	*PreferredHeight = int32(result1)
}

func (m *TControl) CNPreferredSizeChanged() {
	controlImportAPI().SysCallN(38, m.Instance())
}

func (m *TControl) InvalidatePreferredSize() {
	controlImportAPI().SysCallN(105, m.Instance())
}

func (m *TControl) DisableAutoSizing() {
	controlImportAPI().SysCallN(60, m.Instance())
}

func (m *TControl) EnableAutoSizing() {
	controlImportAPI().SysCallN(67, m.Instance())
}

func (m *TControl) UpdateBaseBounds(StoreBounds, StoreParentClientSize, UseLoadedValues bool) {
	controlImportAPI().SysCallN(163, m.Instance(), PascalBool(StoreBounds), PascalBool(StoreParentClientSize), PascalBool(UseLoadedValues))
}

func (m *TControl) WriteLayoutDebugReport(Prefix string) {
	controlImportAPI().SysCallN(170, m.Instance(), PascalStr(Prefix))
}

func (m *TControl) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth int32) {
	controlImportAPI().SysCallN(24, m.Instance(), uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func (m *TControl) ShouldAutoAdjust(AWidth, AHeight *bool) {
	var result0 uintptr
	var result1 uintptr
	controlImportAPI().SysCallN(154, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*AWidth = GoBool(result0)
	*AHeight = GoBool(result1)
}

func (m *TControl) FixDesignFontsPPI(ADesignTimePPI int32) {
	controlImportAPI().SysCallN(73, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TControl) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	controlImportAPI().SysCallN(137, m.Instance(), uintptr(AToPPI), uintptr(unsafePointer(&AProportion)))
}

func (m *TControl) EditingDone() {
	controlImportAPI().SysCallN(66, m.Instance())
}

func (m *TControl) ExecuteDefaultAction() {
	controlImportAPI().SysCallN(71, m.Instance())
}

func (m *TControl) ExecuteCancelAction() {
	controlImportAPI().SysCallN(70, m.Instance())
}

func (m *TControl) BeginDrag(Immediate bool, Threshold int32) {
	controlImportAPI().SysCallN(32, m.Instance(), PascalBool(Immediate), uintptr(Threshold))
}

func (m *TControl) EndDrag(Drop bool) {
	controlImportAPI().SysCallN(69, m.Instance(), PascalBool(Drop))
}

func (m *TControl) BringToFront() {
	controlImportAPI().SysCallN(37, m.Instance())
}

func (m *TControl) Hide() {
	controlImportAPI().SysCallN(100, m.Instance())
}

func (m *TControl) Refresh() {
	controlImportAPI().SysCallN(128, m.Instance())
}

func (m *TControl) Repaint() {
	controlImportAPI().SysCallN(129, m.Instance())
}

func (m *TControl) Invalidate() {
	controlImportAPI().SysCallN(104, m.Instance())
}

func (m *TControl) CheckNewParent(AParent IWinControl) {
	controlImportAPI().SysCallN(42, m.Instance(), GetObjectUintptr(AParent))
}

func (m *TControl) SendToBack() {
	controlImportAPI().SysCallN(144, m.Instance())
}

func (m *TControl) SetTempCursor(Value TCursor) {
	controlImportAPI().SysCallN(152, m.Instance(), uintptr(Value))
}

func (m *TControl) UpdateRolesForForm() {
	controlImportAPI().SysCallN(164, m.Instance())
}

func (m *TControl) ActiveDefaultControlChanged(NewControl IControl) {
	controlImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(NewControl))
}

func (m *TControl) SetTextBuf(Buffer string) {
	controlImportAPI().SysCallN(153, m.Instance(), PascalStr(Buffer))
}

func (m *TControl) Show() {
	controlImportAPI().SysCallN(155, m.Instance())
}

func (m *TControl) Update() {
	controlImportAPI().SysCallN(162, m.Instance())
}

func (m *TControl) InitiateAction() {
	controlImportAPI().SysCallN(103, m.Instance())
}

func (m *TControl) ShowHelp() {
	controlImportAPI().SysCallN(156, m.Instance())
}

func (m *TControl) SetOnChangeBounds(fn TNotifyEvent) {
	if m.changeBoundsPtr != 0 {
		RemoveEventElement(m.changeBoundsPtr)
	}
	m.changeBoundsPtr = MakeEventDataPtr(fn)
	controlImportAPI().SysCallN(148, m.Instance(), m.changeBoundsPtr)
}

func (m *TControl) SetOnClick(fn TNotifyEvent) {
	if m.clickPtr != 0 {
		RemoveEventElement(m.clickPtr)
	}
	m.clickPtr = MakeEventDataPtr(fn)
	controlImportAPI().SysCallN(149, m.Instance(), m.clickPtr)
}

func (m *TControl) SetOnResize(fn TNotifyEvent) {
	if m.resizePtr != 0 {
		RemoveEventElement(m.resizePtr)
	}
	m.resizePtr = MakeEventDataPtr(fn)
	controlImportAPI().SysCallN(150, m.Instance(), m.resizePtr)
}

func (m *TControl) SetOnShowHint(fn TControlShowHintEvent) {
	if m.showHintPtr != 0 {
		RemoveEventElement(m.showHintPtr)
	}
	m.showHintPtr = MakeEventDataPtr(fn)
	controlImportAPI().SysCallN(151, m.Instance(), m.showHintPtr)
}

var (
	controlImport       *imports.Imports = nil
	controlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Control_AccessibleDescription", 0),
		/*1*/ imports.NewTable("Control_AccessibleName", 0),
		/*2*/ imports.NewTable("Control_AccessibleRole", 0),
		/*3*/ imports.NewTable("Control_AccessibleValue", 0),
		/*4*/ imports.NewTable("Control_Action", 0),
		/*5*/ imports.NewTable("Control_ActiveDefaultControlChanged", 0),
		/*6*/ imports.NewTable("Control_AdjustSize", 0),
		/*7*/ imports.NewTable("Control_Align", 0),
		/*8*/ imports.NewTable("Control_AnchorAsAlign", 0),
		/*9*/ imports.NewTable("Control_AnchorClient", 0),
		/*10*/ imports.NewTable("Control_AnchorHorizontalCenterTo", 0),
		/*11*/ imports.NewTable("Control_AnchorParallel", 0),
		/*12*/ imports.NewTable("Control_AnchorSame", 0),
		/*13*/ imports.NewTable("Control_AnchorSide", 0),
		/*14*/ imports.NewTable("Control_AnchorSideBottom", 0),
		/*15*/ imports.NewTable("Control_AnchorSideLeft", 0),
		/*16*/ imports.NewTable("Control_AnchorSideRight", 0),
		/*17*/ imports.NewTable("Control_AnchorSideTop", 0),
		/*18*/ imports.NewTable("Control_AnchorToCompanion", 0),
		/*19*/ imports.NewTable("Control_AnchorToNeighbour", 0),
		/*20*/ imports.NewTable("Control_AnchorVerticalCenterTo", 0),
		/*21*/ imports.NewTable("Control_AnchoredControlCount", 0),
		/*22*/ imports.NewTable("Control_AnchoredControls", 0),
		/*23*/ imports.NewTable("Control_Anchors", 0),
		/*24*/ imports.NewTable("Control_AutoAdjustLayout", 0),
		/*25*/ imports.NewTable("Control_AutoSize", 0),
		/*26*/ imports.NewTable("Control_AutoSizeDelayed", 0),
		/*27*/ imports.NewTable("Control_AutoSizeDelayedHandle", 0),
		/*28*/ imports.NewTable("Control_AutoSizeDelayedReport", 0),
		/*29*/ imports.NewTable("Control_AutoSizePhases", 0),
		/*30*/ imports.NewTable("Control_BaseBounds", 0),
		/*31*/ imports.NewTable("Control_BaseParentClientSize", 0),
		/*32*/ imports.NewTable("Control_BeginDrag", 0),
		/*33*/ imports.NewTable("Control_BiDiMode", 0),
		/*34*/ imports.NewTable("Control_BorderSpacing", 0),
		/*35*/ imports.NewTable("Control_BoundsRect", 0),
		/*36*/ imports.NewTable("Control_BoundsRectForNewParent", 0),
		/*37*/ imports.NewTable("Control_BringToFront", 0),
		/*38*/ imports.NewTable("Control_CNPreferredSizeChanged", 0),
		/*39*/ imports.NewTable("Control_Caption", 0),
		/*40*/ imports.NewTable("Control_CaptureMouseButtons", 0),
		/*41*/ imports.NewTable("Control_CheckChildClassAllowed", 0),
		/*42*/ imports.NewTable("Control_CheckNewParent", 0),
		/*43*/ imports.NewTable("Control_Class", 0),
		/*44*/ imports.NewTable("Control_ClientHeight", 0),
		/*45*/ imports.NewTable("Control_ClientOrigin", 0),
		/*46*/ imports.NewTable("Control_ClientRect", 0),
		/*47*/ imports.NewTable("Control_ClientToParent", 0),
		/*48*/ imports.NewTable("Control_ClientToScreen", 0),
		/*49*/ imports.NewTable("Control_ClientToScreen1", 0),
		/*50*/ imports.NewTable("Control_ClientWidth", 0),
		/*51*/ imports.NewTable("Control_Color", 0),
		/*52*/ imports.NewTable("Control_Constraints", 0),
		/*53*/ imports.NewTable("Control_ControlOrigin", 0),
		/*54*/ imports.NewTable("Control_ControlState", 0),
		/*55*/ imports.NewTable("Control_ControlStyle", 0),
		/*56*/ imports.NewTable("Control_ControlToScreen", 0),
		/*57*/ imports.NewTable("Control_Create", 0),
		/*58*/ imports.NewTable("Control_CreateAccessibleObject", 0),
		/*59*/ imports.NewTable("Control_Cursor", 0),
		/*60*/ imports.NewTable("Control_DisableAutoSizing", 0),
		/*61*/ imports.NewTable("Control_Dock", 0),
		/*62*/ imports.NewTable("Control_DockOrientation", 0),
		/*63*/ imports.NewTable("Control_Docked", 0),
		/*64*/ imports.NewTable("Control_DragDrop", 0),
		/*65*/ imports.NewTable("Control_Dragging", 0),
		/*66*/ imports.NewTable("Control_EditingDone", 0),
		/*67*/ imports.NewTable("Control_EnableAutoSizing", 0),
		/*68*/ imports.NewTable("Control_Enabled", 0),
		/*69*/ imports.NewTable("Control_EndDrag", 0),
		/*70*/ imports.NewTable("Control_ExecuteCancelAction", 0),
		/*71*/ imports.NewTable("Control_ExecuteDefaultAction", 0),
		/*72*/ imports.NewTable("Control_FindSubComponent", 0),
		/*73*/ imports.NewTable("Control_FixDesignFontsPPI", 0),
		/*74*/ imports.NewTable("Control_Floating", 0),
		/*75*/ imports.NewTable("Control_FloatingDockSiteClass", 0),
		/*76*/ imports.NewTable("Control_Font", 0),
		/*77*/ imports.NewTable("Control_FormIsUpdating", 0),
		/*78*/ imports.NewTable("Control_GetAccessibleObject", 0),
		/*79*/ imports.NewTable("Control_GetAnchorsDependingOnParent", 0),
		/*80*/ imports.NewTable("Control_GetCanvasScaleFactor", 0),
		/*81*/ imports.NewTable("Control_GetChildAccessibleObjectAtPos", 0),
		/*82*/ imports.NewTable("Control_GetChildrenRect", 0),
		/*83*/ imports.NewTable("Control_GetColorResolvingParent", 0),
		/*84*/ imports.NewTable("Control_GetDefaultColor", 0),
		/*85*/ imports.NewTable("Control_GetDefaultHeight", 0),
		/*86*/ imports.NewTable("Control_GetDefaultWidth", 0),
		/*87*/ imports.NewTable("Control_GetPreferredSize", 0),
		/*88*/ imports.NewTable("Control_GetRGBColorResolvingParent", 0),
		/*89*/ imports.NewTable("Control_GetSelectedChildAccessibleObject", 0),
		/*90*/ imports.NewTable("Control_GetSidePosition", 0),
		/*91*/ imports.NewTable("Control_GetTextBuf", 0),
		/*92*/ imports.NewTable("Control_GetTextLen", 0),
		/*93*/ imports.NewTable("Control_GetTopParent", 0),
		/*94*/ imports.NewTable("Control_HandleObjectShouldBeVisible", 0),
		/*95*/ imports.NewTable("Control_HasHelp", 0),
		/*96*/ imports.NewTable("Control_Height", 0),
		/*97*/ imports.NewTable("Control_HelpContext", 0),
		/*98*/ imports.NewTable("Control_HelpKeyword", 0),
		/*99*/ imports.NewTable("Control_HelpType", 0),
		/*100*/ imports.NewTable("Control_Hide", 0),
		/*101*/ imports.NewTable("Control_Hint", 0),
		/*102*/ imports.NewTable("Control_HostDockSite", 0),
		/*103*/ imports.NewTable("Control_InitiateAction", 0),
		/*104*/ imports.NewTable("Control_Invalidate", 0),
		/*105*/ imports.NewTable("Control_InvalidatePreferredSize", 0),
		/*106*/ imports.NewTable("Control_IsControl", 0),
		/*107*/ imports.NewTable("Control_IsControlVisible", 0),
		/*108*/ imports.NewTable("Control_IsEnabled", 0),
		/*109*/ imports.NewTable("Control_IsParentColor", 0),
		/*110*/ imports.NewTable("Control_IsParentFont", 0),
		/*111*/ imports.NewTable("Control_IsParentOf", 0),
		/*112*/ imports.NewTable("Control_IsProcessingPaintMsg", 0),
		/*113*/ imports.NewTable("Control_IsRightToLeft", 0),
		/*114*/ imports.NewTable("Control_IsVisible", 0),
		/*115*/ imports.NewTable("Control_LRDockWidth", 0),
		/*116*/ imports.NewTable("Control_Left", 0),
		/*117*/ imports.NewTable("Control_ManualDock", 0),
		/*118*/ imports.NewTable("Control_ManualFloat", 0),
		/*119*/ imports.NewTable("Control_MouseInClient", 0),
		/*120*/ imports.NewTable("Control_Parent", 0),
		/*121*/ imports.NewTable("Control_ParentBiDiMode", 0),
		/*122*/ imports.NewTable("Control_ParentDestroyingHandle", 0),
		/*123*/ imports.NewTable("Control_ParentHandlesAllocated", 0),
		/*124*/ imports.NewTable("Control_ParentToClient", 0),
		/*125*/ imports.NewTable("Control_Perform", 0),
		/*126*/ imports.NewTable("Control_PopupMenu", 0),
		/*127*/ imports.NewTable("Control_ReadBounds", 0),
		/*128*/ imports.NewTable("Control_Refresh", 0),
		/*129*/ imports.NewTable("Control_Repaint", 0),
		/*130*/ imports.NewTable("Control_ReplaceDockedControl", 0),
		/*131*/ imports.NewTable("Control_Scale96ToFont", 0),
		/*132*/ imports.NewTable("Control_Scale96ToForm", 0),
		/*133*/ imports.NewTable("Control_Scale96ToScreen", 0),
		/*134*/ imports.NewTable("Control_ScaleDesignToForm", 0),
		/*135*/ imports.NewTable("Control_ScaleFontTo96", 0),
		/*136*/ imports.NewTable("Control_ScaleFontToScreen", 0),
		/*137*/ imports.NewTable("Control_ScaleFontsPPI", 0),
		/*138*/ imports.NewTable("Control_ScaleFormTo96", 0),
		/*139*/ imports.NewTable("Control_ScaleFormToDesign", 0),
		/*140*/ imports.NewTable("Control_ScaleScreenTo96", 0),
		/*141*/ imports.NewTable("Control_ScaleScreenToFont", 0),
		/*142*/ imports.NewTable("Control_ScreenToClient", 0),
		/*143*/ imports.NewTable("Control_ScreenToControl", 0),
		/*144*/ imports.NewTable("Control_SendToBack", 0),
		/*145*/ imports.NewTable("Control_SetBounds", 0),
		/*146*/ imports.NewTable("Control_SetBoundsKeepBase", 0),
		/*147*/ imports.NewTable("Control_SetInitialBounds", 0),
		/*148*/ imports.NewTable("Control_SetOnChangeBounds", 0),
		/*149*/ imports.NewTable("Control_SetOnClick", 0),
		/*150*/ imports.NewTable("Control_SetOnResize", 0),
		/*151*/ imports.NewTable("Control_SetOnShowHint", 0),
		/*152*/ imports.NewTable("Control_SetTempCursor", 0),
		/*153*/ imports.NewTable("Control_SetTextBuf", 0),
		/*154*/ imports.NewTable("Control_ShouldAutoAdjust", 0),
		/*155*/ imports.NewTable("Control_Show", 0),
		/*156*/ imports.NewTable("Control_ShowHelp", 0),
		/*157*/ imports.NewTable("Control_ShowHint", 0),
		/*158*/ imports.NewTable("Control_TBDockHeight", 0),
		/*159*/ imports.NewTable("Control_Top", 0),
		/*160*/ imports.NewTable("Control_UndockHeight", 0),
		/*161*/ imports.NewTable("Control_UndockWidth", 0),
		/*162*/ imports.NewTable("Control_Update", 0),
		/*163*/ imports.NewTable("Control_UpdateBaseBounds", 0),
		/*164*/ imports.NewTable("Control_UpdateRolesForForm", 0),
		/*165*/ imports.NewTable("Control_UseRightToLeftAlignment", 0),
		/*166*/ imports.NewTable("Control_UseRightToLeftReading", 0),
		/*167*/ imports.NewTable("Control_UseRightToLeftScrollBar", 0),
		/*168*/ imports.NewTable("Control_Visible", 0),
		/*169*/ imports.NewTable("Control_Width", 0),
		/*170*/ imports.NewTable("Control_WriteLayoutDebugReport", 0),
	}
)

func controlImportAPI() *imports.Imports {
	if controlImport == nil {
		controlImport = NewDefaultImports()
		controlImport.SetImportTable(controlImportTables)
		controlImportTables = nil
	}
	return controlImport
}
