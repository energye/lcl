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

// IWinControl Parent: IControl
type IWinControl interface {
	IControl
	BorderWidth() TBorderWidth                                                   // property
	SetBorderWidth(AValue TBorderWidth)                                          // property
	BoundsLockCount() int32                                                      // property
	Brush() IBrush                                                               // property
	CachedClientHeight() int32                                                   // property
	CachedClientWidth() int32                                                    // property
	ChildSizing() IControlChildSizing                                            // property
	SetChildSizing(AValue IControlChildSizing)                                   // property
	ControlCount() int32                                                         // property
	Controls(Index int32) IControl                                               // property
	DockClientCount() int32                                                      // property
	DockClients(Index int32) IControl                                            // property
	DockManager() IDockManager                                                   // property
	SetDockManager(AValue IDockManager)                                          // property
	DockSite() bool                                                              // property
	SetDockSite(AValue bool)                                                     // property
	DoubleBuffered() bool                                                        // property
	SetDoubleBuffered(AValue bool)                                               // property
	Handle() HWND                                                                // property
	SetHandle(AValue HWND)                                                       // property
	IsFlipped() bool                                                             // property
	IsResizing() bool                                                            // property
	TabOrder() TTabOrder                                                         // property
	SetTabOrder(AValue TTabOrder)                                                // property
	TabStop() bool                                                               // property
	SetTabStop(AValue bool)                                                      // property
	ParentDoubleBuffered() bool                                                  // property
	SetParentDoubleBuffered(AValue bool)                                         // property
	ParentWindow() HWND                                                          // property
	SetParentWindow(AValue HWND)                                                 // property
	Showing() bool                                                               // property
	UseDockManager() bool                                                        // property
	SetUseDockManager(AValue bool)                                               // property
	SetDesignerDeleting(AValue bool)                                             // property
	IsSpecialSubControl() bool                                                   // property
	VisibleDockClientCount() int32                                               // property
	ControlAtPos(Pos *TPoint, AllowDisabled bool) IControl                       // function
	ControlAtPos1(Pos *TPoint, AllowDisabled, AllowWinControls bool) IControl    // function
	ControlAtPos2(Pos *TPoint, Flags TControlAtPosFlags) IControl                // function
	ContainsControl(Control IControl) bool                                       // function
	ClientRectNeedsInterfaceUpdate() bool                                        // function
	CanFocus() bool                                                              // function
	CanSetFocus() bool                                                           // function
	GetControlIndex(AControl IControl) int32                                     // function
	Focused() bool                                                               // function
	PerformTab(ForwardTab bool) bool                                             // function
	FindChildControl(ControlName string) IControl                                // function
	GetEnumeratorControls() IWinControlEnumerator                                // function
	GetEnumeratorControlsReverse() IWinControlEnumerator                         // function
	GetDockCaption(AControl IControl) string                                     // function
	HandleAllocated() bool                                                       // function
	BrushCreated() bool                                                          // function
	IntfUTF8KeyPress(UTF8Key *TUTF8Char, RepeatCount int32, SystemKey bool) bool // function
	IntfGetDropFilesTarget() IWinControl                                         // function
	BeginUpdateBounds()                                                          // procedure
	EndUpdateBounds()                                                            // procedure
	LockRealizeBounds()                                                          // procedure
	UnlockRealizeBounds()                                                        // procedure
	DoAdjustClientRectChange(InvalidateRect bool)                                // procedure
	InvalidateClientRectCache(WithChildControls bool)                            // procedure
	DisableAlign()                                                               // procedure
	EnableAlign()                                                                // procedure
	ReAlign()                                                                    // procedure
	ScrollByWS(DeltaX, DeltaY int32)                                             // procedure
	ScrollBy(DeltaX, DeltaY int32)                                               // procedure
	FixDesignFontsPPIWithChildren(ADesignTimePPI int32)                          // procedure
	DockDrop(DragDockObject IDragDockObject, X, Y int32)                         // procedure
	SetControlIndex(AControl IControl, NewIndex int32)                           // procedure
	SelectNext(CurControl IWinControl, GoForward, CheckTabStop bool)             // procedure
	NotifyControls(Msg Word)                                                     // procedure
	AddControl()                                                                 // procedure
	InsertControl(AControl IControl)                                             // procedure
	InsertControl1(AControl IControl, Index int32)                               // procedure
	RemoveControl(AControl IControl)                                             // procedure
	SetFocus()                                                                   // procedure
	FlipChildren(AllLevels bool)                                                 // procedure
	ScaleBy(Multiplier, Divider int32)                                           // procedure
	UpdateDockCaption(Exclude IControl)                                          // procedure
	HandleNeeded()                                                               // procedure
	EraseBackground(DC HDC)                                                      // procedure
	PaintTo(DC HDC, X, Y int32)                                                  // procedure
	PaintTo1(ACanvas ICanvas, X, Y int32)                                        // procedure
	SetShape(AShape IBitmap)                                                     // procedure
	SetShape1(AShape IRegion)                                                    // procedure
	SetOnAlignInsertBefore(fn TAlignInsertBeforeEvent)                           // property event
	SetOnAlignPosition(fn TAlignPositionEvent)                                   // property event
	SetOnDockDrop(fn TDockDropEvent)                                             // property event
	SetOnDockOver(fn TDockOverEvent)                                             // property event
	SetOnEnter(fn TNotifyEvent)                                                  // property event
	SetOnExit(fn TNotifyEvent)                                                   // property event
	SetOnKeyDown(fn TKeyEvent)                                                   // property event
	SetOnKeyPress(fn TKeyPressEvent)                                             // property event
	SetOnKeyUp(fn TKeyEvent)                                                     // property event
	SetOnUnDock(fn TUnDockEvent)                                                 // property event
	SetOnUTF8KeyPress(fn TUTF8KeyPressEvent)                                     // property event
}

// TWinControl Parent: TControl
type TWinControl struct {
	TControl
	alignInsertBeforePtr uintptr
	alignPositionPtr     uintptr
	dockDropPtr          uintptr
	dockOverPtr          uintptr
	enterPtr             uintptr
	exitPtr              uintptr
	keyDownPtr           uintptr
	keyPressPtr          uintptr
	keyUpPtr             uintptr
	unDockPtr            uintptr
	uTF8KeyPressPtr      uintptr
}

func NewWinControl(TheOwner IComponent) IWinControl {
	r1 := winControlImportAPI().SysCallN(19, GetObjectUintptr(TheOwner))
	return AsWinControl(r1)
}

func NewWinControlParented(AParentWindow HWND) IWinControl {
	r1 := winControlImportAPI().SysCallN(20, uintptr(AParentWindow))
	return AsWinControl(r1)
}

func (m *TWinControl) BorderWidth() TBorderWidth {
	r1 := winControlImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TBorderWidth(r1)
}

func (m *TWinControl) SetBorderWidth(AValue TBorderWidth) {
	winControlImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) BoundsLockCount() int32 {
	r1 := winControlImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func (m *TWinControl) Brush() IBrush {
	r1 := winControlImportAPI().SysCallN(4, m.Instance())
	return AsBrush(r1)
}

func (m *TWinControl) CachedClientHeight() int32 {
	r1 := winControlImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TWinControl) CachedClientWidth() int32 {
	r1 := winControlImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TWinControl) ChildSizing() IControlChildSizing {
	r1 := winControlImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsControlChildSizing(r1)
}

func (m *TWinControl) SetChildSizing(AValue IControlChildSizing) {
	winControlImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TWinControl) ControlCount() int32 {
	r1 := winControlImportAPI().SysCallN(17, m.Instance())
	return int32(r1)
}

func (m *TWinControl) Controls(Index int32) IControl {
	r1 := winControlImportAPI().SysCallN(18, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TWinControl) DockClientCount() int32 {
	r1 := winControlImportAPI().SysCallN(24, m.Instance())
	return int32(r1)
}

func (m *TWinControl) DockClients(Index int32) IControl {
	r1 := winControlImportAPI().SysCallN(25, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TWinControl) DockManager() IDockManager {
	r1 := winControlImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return AsDockManager(r1)
}

func (m *TWinControl) SetDockManager(AValue IDockManager) {
	winControlImportAPI().SysCallN(27, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TWinControl) DockSite() bool {
	r1 := winControlImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetDockSite(AValue bool) {
	winControlImportAPI().SysCallN(28, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) DoubleBuffered() bool {
	r1 := winControlImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetDoubleBuffered(AValue bool) {
	winControlImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) Handle() HWND {
	r1 := winControlImportAPI().SysCallN(41, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TWinControl) SetHandle(AValue HWND) {
	winControlImportAPI().SysCallN(41, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) IsFlipped() bool {
	r1 := winControlImportAPI().SysCallN(49, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) IsResizing() bool {
	r1 := winControlImportAPI().SysCallN(50, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) TabOrder() TTabOrder {
	r1 := winControlImportAPI().SysCallN(81, 0, m.Instance(), 0)
	return TTabOrder(r1)
}

func (m *TWinControl) SetTabOrder(AValue TTabOrder) {
	winControlImportAPI().SysCallN(81, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) TabStop() bool {
	r1 := winControlImportAPI().SysCallN(82, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetTabStop(AValue bool) {
	winControlImportAPI().SysCallN(82, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) ParentDoubleBuffered() bool {
	r1 := winControlImportAPI().SysCallN(56, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetParentDoubleBuffered(AValue bool) {
	winControlImportAPI().SysCallN(56, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) ParentWindow() HWND {
	r1 := winControlImportAPI().SysCallN(57, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TWinControl) SetParentWindow(AValue HWND) {
	winControlImportAPI().SysCallN(57, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) Showing() bool {
	r1 := winControlImportAPI().SysCallN(80, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) UseDockManager() bool {
	r1 := winControlImportAPI().SysCallN(85, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetUseDockManager(AValue bool) {
	winControlImportAPI().SysCallN(85, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) SetDesignerDeleting(AValue bool) {
	winControlImportAPI().SysCallN(21, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) IsSpecialSubControl() bool {
	r1 := winControlImportAPI().SysCallN(51, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) VisibleDockClientCount() int32 {
	r1 := winControlImportAPI().SysCallN(86, m.Instance())
	return int32(r1)
}

func (m *TWinControl) ControlAtPos(Pos *TPoint, AllowDisabled bool) IControl {
	r1 := winControlImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(Pos)), PascalBool(AllowDisabled))
	return AsControl(r1)
}

func (m *TWinControl) ControlAtPos1(Pos *TPoint, AllowDisabled, AllowWinControls bool) IControl {
	r1 := winControlImportAPI().SysCallN(15, m.Instance(), uintptr(unsafePointer(Pos)), PascalBool(AllowDisabled), PascalBool(AllowWinControls))
	return AsControl(r1)
}

func (m *TWinControl) ControlAtPos2(Pos *TPoint, Flags TControlAtPosFlags) IControl {
	r1 := winControlImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(Pos)), uintptr(Flags))
	return AsControl(r1)
}

func (m *TWinControl) ContainsControl(Control IControl) bool {
	r1 := winControlImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TWinControl) ClientRectNeedsInterfaceUpdate() bool {
	r1 := winControlImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) CanFocus() bool {
	r1 := winControlImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) CanSetFocus() bool {
	r1 := winControlImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) GetControlIndex(AControl IControl) int32 {
	r1 := winControlImportAPI().SysCallN(37, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func (m *TWinControl) Focused() bool {
	r1 := winControlImportAPI().SysCallN(36, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) PerformTab(ForwardTab bool) bool {
	r1 := winControlImportAPI().SysCallN(58, m.Instance(), PascalBool(ForwardTab))
	return GoBool(r1)
}

func (m *TWinControl) FindChildControl(ControlName string) IControl {
	r1 := winControlImportAPI().SysCallN(33, m.Instance(), PascalStr(ControlName))
	return AsControl(r1)
}

func (m *TWinControl) GetEnumeratorControls() IWinControlEnumerator {
	r1 := winControlImportAPI().SysCallN(39, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControl) GetEnumeratorControlsReverse() IWinControlEnumerator {
	r1 := winControlImportAPI().SysCallN(40, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControl) GetDockCaption(AControl IControl) string {
	r1 := winControlImportAPI().SysCallN(38, m.Instance(), GetObjectUintptr(AControl))
	return GoStr(r1)
}

func (m *TWinControl) HandleAllocated() bool {
	r1 := winControlImportAPI().SysCallN(42, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) BrushCreated() bool {
	r1 := winControlImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) IntfUTF8KeyPress(UTF8Key *TUTF8Char, RepeatCount int32, SystemKey bool) bool {
	var result0 uintptr
	r1 := winControlImportAPI().SysCallN(47, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(RepeatCount), PascalBool(SystemKey))
	*UTF8Key = *(*TUTF8Char)(getPointer(result0))
	return GoBool(r1)
}

func (m *TWinControl) IntfGetDropFilesTarget() IWinControl {
	r1 := winControlImportAPI().SysCallN(46, m.Instance())
	return AsWinControl(r1)
}

func WinControlClass() TClass {
	ret := winControlImportAPI().SysCallN(11)
	return TClass(ret)
}

func (m *TWinControl) BeginUpdateBounds() {
	winControlImportAPI().SysCallN(1, m.Instance())
}

func (m *TWinControl) EndUpdateBounds() {
	winControlImportAPI().SysCallN(31, m.Instance())
}

func (m *TWinControl) LockRealizeBounds() {
	winControlImportAPI().SysCallN(52, m.Instance())
}

func (m *TWinControl) UnlockRealizeBounds() {
	winControlImportAPI().SysCallN(83, m.Instance())
}

func (m *TWinControl) DoAdjustClientRectChange(InvalidateRect bool) {
	winControlImportAPI().SysCallN(23, m.Instance(), PascalBool(InvalidateRect))
}

func (m *TWinControl) InvalidateClientRectCache(WithChildControls bool) {
	winControlImportAPI().SysCallN(48, m.Instance(), PascalBool(WithChildControls))
}

func (m *TWinControl) DisableAlign() {
	winControlImportAPI().SysCallN(22, m.Instance())
}

func (m *TWinControl) EnableAlign() {
	winControlImportAPI().SysCallN(30, m.Instance())
}

func (m *TWinControl) ReAlign() {
	winControlImportAPI().SysCallN(59, m.Instance())
}

func (m *TWinControl) ScrollByWS(DeltaX, DeltaY int32) {
	winControlImportAPI().SysCallN(63, m.Instance(), uintptr(DeltaX), uintptr(DeltaY))
}

func (m *TWinControl) ScrollBy(DeltaX, DeltaY int32) {
	winControlImportAPI().SysCallN(62, m.Instance(), uintptr(DeltaX), uintptr(DeltaY))
}

func (m *TWinControl) FixDesignFontsPPIWithChildren(ADesignTimePPI int32) {
	winControlImportAPI().SysCallN(34, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TWinControl) DockDrop(DragDockObject IDragDockObject, X, Y int32) {
	winControlImportAPI().SysCallN(26, m.Instance(), GetObjectUintptr(DragDockObject), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetControlIndex(AControl IControl, NewIndex int32) {
	winControlImportAPI().SysCallN(65, m.Instance(), GetObjectUintptr(AControl), uintptr(NewIndex))
}

func (m *TWinControl) SelectNext(CurControl IWinControl, GoForward, CheckTabStop bool) {
	winControlImportAPI().SysCallN(64, m.Instance(), GetObjectUintptr(CurControl), PascalBool(GoForward), PascalBool(CheckTabStop))
}

func (m *TWinControl) NotifyControls(Msg Word) {
	winControlImportAPI().SysCallN(53, m.Instance(), uintptr(Msg))
}

func (m *TWinControl) AddControl() {
	winControlImportAPI().SysCallN(0, m.Instance())
}

func (m *TWinControl) InsertControl(AControl IControl) {
	winControlImportAPI().SysCallN(44, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TWinControl) InsertControl1(AControl IControl, Index int32) {
	winControlImportAPI().SysCallN(45, m.Instance(), GetObjectUintptr(AControl), uintptr(Index))
}

func (m *TWinControl) RemoveControl(AControl IControl) {
	winControlImportAPI().SysCallN(60, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TWinControl) SetFocus() {
	winControlImportAPI().SysCallN(66, m.Instance())
}

func (m *TWinControl) FlipChildren(AllLevels bool) {
	winControlImportAPI().SysCallN(35, m.Instance(), PascalBool(AllLevels))
}

func (m *TWinControl) ScaleBy(Multiplier, Divider int32) {
	winControlImportAPI().SysCallN(61, m.Instance(), uintptr(Multiplier), uintptr(Divider))
}

func (m *TWinControl) UpdateDockCaption(Exclude IControl) {
	winControlImportAPI().SysCallN(84, m.Instance(), GetObjectUintptr(Exclude))
}

func (m *TWinControl) HandleNeeded() {
	winControlImportAPI().SysCallN(43, m.Instance())
}

func (m *TWinControl) EraseBackground(DC HDC) {
	winControlImportAPI().SysCallN(32, m.Instance(), uintptr(DC))
}

func (m *TWinControl) PaintTo(DC HDC, X, Y int32) {
	winControlImportAPI().SysCallN(54, m.Instance(), uintptr(DC), uintptr(X), uintptr(Y))
}

func (m *TWinControl) PaintTo1(ACanvas ICanvas, X, Y int32) {
	winControlImportAPI().SysCallN(55, m.Instance(), GetObjectUintptr(ACanvas), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetShape(AShape IBitmap) {
	winControlImportAPI().SysCallN(78, m.Instance(), GetObjectUintptr(AShape))
}

func (m *TWinControl) SetShape1(AShape IRegion) {
	winControlImportAPI().SysCallN(79, m.Instance(), GetObjectUintptr(AShape))
}

func (m *TWinControl) SetOnAlignInsertBefore(fn TAlignInsertBeforeEvent) {
	if m.alignInsertBeforePtr != 0 {
		RemoveEventElement(m.alignInsertBeforePtr)
	}
	m.alignInsertBeforePtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(67, m.Instance(), m.alignInsertBeforePtr)
}

func (m *TWinControl) SetOnAlignPosition(fn TAlignPositionEvent) {
	if m.alignPositionPtr != 0 {
		RemoveEventElement(m.alignPositionPtr)
	}
	m.alignPositionPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(68, m.Instance(), m.alignPositionPtr)
}

func (m *TWinControl) SetOnDockDrop(fn TDockDropEvent) {
	if m.dockDropPtr != 0 {
		RemoveEventElement(m.dockDropPtr)
	}
	m.dockDropPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(69, m.Instance(), m.dockDropPtr)
}

func (m *TWinControl) SetOnDockOver(fn TDockOverEvent) {
	if m.dockOverPtr != 0 {
		RemoveEventElement(m.dockOverPtr)
	}
	m.dockOverPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(70, m.Instance(), m.dockOverPtr)
}

func (m *TWinControl) SetOnEnter(fn TNotifyEvent) {
	if m.enterPtr != 0 {
		RemoveEventElement(m.enterPtr)
	}
	m.enterPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(71, m.Instance(), m.enterPtr)
}

func (m *TWinControl) SetOnExit(fn TNotifyEvent) {
	if m.exitPtr != 0 {
		RemoveEventElement(m.exitPtr)
	}
	m.exitPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(72, m.Instance(), m.exitPtr)
}

func (m *TWinControl) SetOnKeyDown(fn TKeyEvent) {
	if m.keyDownPtr != 0 {
		RemoveEventElement(m.keyDownPtr)
	}
	m.keyDownPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(73, m.Instance(), m.keyDownPtr)
}

func (m *TWinControl) SetOnKeyPress(fn TKeyPressEvent) {
	if m.keyPressPtr != 0 {
		RemoveEventElement(m.keyPressPtr)
	}
	m.keyPressPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(74, m.Instance(), m.keyPressPtr)
}

func (m *TWinControl) SetOnKeyUp(fn TKeyEvent) {
	if m.keyUpPtr != 0 {
		RemoveEventElement(m.keyUpPtr)
	}
	m.keyUpPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(75, m.Instance(), m.keyUpPtr)
}

func (m *TWinControl) SetOnUnDock(fn TUnDockEvent) {
	if m.unDockPtr != 0 {
		RemoveEventElement(m.unDockPtr)
	}
	m.unDockPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(77, m.Instance(), m.unDockPtr)
}

func (m *TWinControl) SetOnUTF8KeyPress(fn TUTF8KeyPressEvent) {
	if m.uTF8KeyPressPtr != 0 {
		RemoveEventElement(m.uTF8KeyPressPtr)
	}
	m.uTF8KeyPressPtr = MakeEventDataPtr(fn)
	winControlImportAPI().SysCallN(76, m.Instance(), m.uTF8KeyPressPtr)
}

var (
	winControlImport       *imports.Imports = nil
	winControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WinControl_AddControl", 0),
		/*1*/ imports.NewTable("WinControl_BeginUpdateBounds", 0),
		/*2*/ imports.NewTable("WinControl_BorderWidth", 0),
		/*3*/ imports.NewTable("WinControl_BoundsLockCount", 0),
		/*4*/ imports.NewTable("WinControl_Brush", 0),
		/*5*/ imports.NewTable("WinControl_BrushCreated", 0),
		/*6*/ imports.NewTable("WinControl_CachedClientHeight", 0),
		/*7*/ imports.NewTable("WinControl_CachedClientWidth", 0),
		/*8*/ imports.NewTable("WinControl_CanFocus", 0),
		/*9*/ imports.NewTable("WinControl_CanSetFocus", 0),
		/*10*/ imports.NewTable("WinControl_ChildSizing", 0),
		/*11*/ imports.NewTable("WinControl_Class", 0),
		/*12*/ imports.NewTable("WinControl_ClientRectNeedsInterfaceUpdate", 0),
		/*13*/ imports.NewTable("WinControl_ContainsControl", 0),
		/*14*/ imports.NewTable("WinControl_ControlAtPos", 0),
		/*15*/ imports.NewTable("WinControl_ControlAtPos1", 0),
		/*16*/ imports.NewTable("WinControl_ControlAtPos2", 0),
		/*17*/ imports.NewTable("WinControl_ControlCount", 0),
		/*18*/ imports.NewTable("WinControl_Controls", 0),
		/*19*/ imports.NewTable("WinControl_Create", 0),
		/*20*/ imports.NewTable("WinControl_CreateParented", 0),
		/*21*/ imports.NewTable("WinControl_DesignerDeleting", 0),
		/*22*/ imports.NewTable("WinControl_DisableAlign", 0),
		/*23*/ imports.NewTable("WinControl_DoAdjustClientRectChange", 0),
		/*24*/ imports.NewTable("WinControl_DockClientCount", 0),
		/*25*/ imports.NewTable("WinControl_DockClients", 0),
		/*26*/ imports.NewTable("WinControl_DockDrop", 0),
		/*27*/ imports.NewTable("WinControl_DockManager", 0),
		/*28*/ imports.NewTable("WinControl_DockSite", 0),
		/*29*/ imports.NewTable("WinControl_DoubleBuffered", 0),
		/*30*/ imports.NewTable("WinControl_EnableAlign", 0),
		/*31*/ imports.NewTable("WinControl_EndUpdateBounds", 0),
		/*32*/ imports.NewTable("WinControl_EraseBackground", 0),
		/*33*/ imports.NewTable("WinControl_FindChildControl", 0),
		/*34*/ imports.NewTable("WinControl_FixDesignFontsPPIWithChildren", 0),
		/*35*/ imports.NewTable("WinControl_FlipChildren", 0),
		/*36*/ imports.NewTable("WinControl_Focused", 0),
		/*37*/ imports.NewTable("WinControl_GetControlIndex", 0),
		/*38*/ imports.NewTable("WinControl_GetDockCaption", 0),
		/*39*/ imports.NewTable("WinControl_GetEnumeratorControls", 0),
		/*40*/ imports.NewTable("WinControl_GetEnumeratorControlsReverse", 0),
		/*41*/ imports.NewTable("WinControl_Handle", 0),
		/*42*/ imports.NewTable("WinControl_HandleAllocated", 0),
		/*43*/ imports.NewTable("WinControl_HandleNeeded", 0),
		/*44*/ imports.NewTable("WinControl_InsertControl", 0),
		/*45*/ imports.NewTable("WinControl_InsertControl1", 0),
		/*46*/ imports.NewTable("WinControl_IntfGetDropFilesTarget", 0),
		/*47*/ imports.NewTable("WinControl_IntfUTF8KeyPress", 0),
		/*48*/ imports.NewTable("WinControl_InvalidateClientRectCache", 0),
		/*49*/ imports.NewTable("WinControl_IsFlipped", 0),
		/*50*/ imports.NewTable("WinControl_IsResizing", 0),
		/*51*/ imports.NewTable("WinControl_IsSpecialSubControl", 0),
		/*52*/ imports.NewTable("WinControl_LockRealizeBounds", 0),
		/*53*/ imports.NewTable("WinControl_NotifyControls", 0),
		/*54*/ imports.NewTable("WinControl_PaintTo", 0),
		/*55*/ imports.NewTable("WinControl_PaintTo1", 0),
		/*56*/ imports.NewTable("WinControl_ParentDoubleBuffered", 0),
		/*57*/ imports.NewTable("WinControl_ParentWindow", 0),
		/*58*/ imports.NewTable("WinControl_PerformTab", 0),
		/*59*/ imports.NewTable("WinControl_ReAlign", 0),
		/*60*/ imports.NewTable("WinControl_RemoveControl", 0),
		/*61*/ imports.NewTable("WinControl_ScaleBy", 0),
		/*62*/ imports.NewTable("WinControl_ScrollBy", 0),
		/*63*/ imports.NewTable("WinControl_ScrollByWS", 0),
		/*64*/ imports.NewTable("WinControl_SelectNext", 0),
		/*65*/ imports.NewTable("WinControl_SetControlIndex", 0),
		/*66*/ imports.NewTable("WinControl_SetFocus", 0),
		/*67*/ imports.NewTable("WinControl_SetOnAlignInsertBefore", 0),
		/*68*/ imports.NewTable("WinControl_SetOnAlignPosition", 0),
		/*69*/ imports.NewTable("WinControl_SetOnDockDrop", 0),
		/*70*/ imports.NewTable("WinControl_SetOnDockOver", 0),
		/*71*/ imports.NewTable("WinControl_SetOnEnter", 0),
		/*72*/ imports.NewTable("WinControl_SetOnExit", 0),
		/*73*/ imports.NewTable("WinControl_SetOnKeyDown", 0),
		/*74*/ imports.NewTable("WinControl_SetOnKeyPress", 0),
		/*75*/ imports.NewTable("WinControl_SetOnKeyUp", 0),
		/*76*/ imports.NewTable("WinControl_SetOnUTF8KeyPress", 0),
		/*77*/ imports.NewTable("WinControl_SetOnUnDock", 0),
		/*78*/ imports.NewTable("WinControl_SetShape", 0),
		/*79*/ imports.NewTable("WinControl_SetShape1", 0),
		/*80*/ imports.NewTable("WinControl_Showing", 0),
		/*81*/ imports.NewTable("WinControl_TabOrder", 0),
		/*82*/ imports.NewTable("WinControl_TabStop", 0),
		/*83*/ imports.NewTable("WinControl_UnlockRealizeBounds", 0),
		/*84*/ imports.NewTable("WinControl_UpdateDockCaption", 0),
		/*85*/ imports.NewTable("WinControl_UseDockManager", 0),
		/*86*/ imports.NewTable("WinControl_VisibleDockClientCount", 0),
	}
)

func winControlImportAPI() *imports.Imports {
	if winControlImport == nil {
		winControlImport = NewDefaultImports()
		winControlImport.SetImportTable(winControlImportTables)
		winControlImportTables = nil
	}
	return winControlImport
}
