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

// IWinControl Parent: IControl
type IWinControl interface {
	IControl
	ControlAtPosWithPointBool(pos types.TPoint, allowDisabled bool) IControl                          // function
	ControlAtPosWithPointBoolX2(pos types.TPoint, allowDisabled bool, allowWinControls bool) IControl // function
	ControlAtPosWithPointControlAtPosFlags(pos types.TPoint, flags types.TControlAtPosFlags) IControl // function
	ContainsControl(control IControl) bool                                                            // function
	ClientRectNeedsInterfaceUpdate() bool                                                             // function
	CanFocus() bool                                                                                   // function
	CanSetFocus() bool                                                                                // function
	GetControlIndex(control IControl) int32                                                           // function
	Focused() bool                                                                                    // function
	PerformTab(forwardTab bool) bool                                                                  // function
	FindChildControl(controlName string) IControl                                                     // function
	// GetEnumeratorControls
	//  enumerators
	GetEnumeratorControls() IWinControlEnumerator                             // function
	GetEnumeratorControlsReverse() IWinControlEnumerator                      // function
	GetDockCaption(control IControl) string                                   // function
	HandleAllocated() bool                                                    // function
	BrushCreated() bool                                                       // function
	IntfUTF8KeyPress(uTF8Key *string, repeatCount int32, systemKey bool) bool // function
	IntfGetDropFilesTarget() IWinControl                                      // function
	BeginUpdateBounds()                                                       // procedure
	EndUpdateBounds()                                                         // procedure
	LockRealizeBounds()                                                       // procedure
	UnlockRealizeBounds()                                                     // procedure
	DoAdjustClientRectChange(invalidateRect bool)                             // procedure
	InvalidateClientRectCache(withChildControls bool)                         // procedure
	DisableAlign()                                                            // procedure
	EnableAlign()                                                             // procedure
	ReAlign()                                                                 // procedure
	ScrollByWS(deltaX int32, deltaY int32)                                    // procedure
	ScrollBy(deltaX int32, deltaY int32)                                      // procedure
	FixDesignFontsPPIWithChildren(designTimePPI int32)                        // procedure
	DockDrop(dragDockObject IDragDockObject, X int32, Y int32)                // procedure
	SetControlIndex(control IControl, newIndex int32)                         // procedure
	SelectNext(curControl IWinControl, goForward bool, checkTabStop bool)     // procedure
	BroadCast(toAllMessage *uintptr)                                          // procedure
	NotifyControls(msg uint16)                                                // procedure
	AddControl()                                                              // procedure
	InsertControlWithControl(control IControl)                                // procedure
	InsertControlWithControlInt(control IControl, index int32)                // procedure
	RemoveControl(control IControl)                                           // procedure
	SetFocus()                                                                // procedure
	FlipChildren(allLevels bool)                                              // procedure
	ScaleBy(multiplier int32, divider int32)                                  // procedure
	UpdateDockCaption(exclude IControl)                                       // procedure
	GetTabOrderList(list IFPList)                                             // procedure
	HandleNeeded()                                                            // procedure
	EraseBackground(dC types.HDC)                                             // procedure
	PaintToWithHDCIntX2(dC types.HDC, X int32, Y int32)                       // procedure
	PaintToWithCanvasIntX2(canvas ICanvas, X int32, Y int32)                  // procedure
	SetShapeWithBitmap(shape IBitmap)                                         // procedure
	SetShapeWithRegion(shape IRegion)                                         // procedure
	// BorderWidth
	//  properties which are supported by all descendents
	BorderWidth() types.TBorderWidth                   // property BorderWidth Getter
	SetBorderWidth(value types.TBorderWidth)           // property BorderWidth Setter
	BoundsLockCount() int32                            // property BoundsLockCount Getter
	Brush() IBrush                                     // property Brush Getter
	CachedClientHeight() int32                         // property CachedClientHeight Getter
	CachedClientWidth() int32                          // property CachedClientWidth Getter
	ChildSizing() IControlChildSizing                  // property ChildSizing Getter
	SetChildSizing(value IControlChildSizing)          // property ChildSizing Setter
	ControlCount() int32                               // property ControlCount Getter
	Controls(index int32) IControl                     // property Controls Getter
	DefWndProc() uintptr                               // property DefWndProc Getter
	SetDefWndProc(value uintptr)                       // property DefWndProc Setter
	DockClientCount() int32                            // property DockClientCount Getter
	DockClients(index int32) IControl                  // property DockClients Getter
	DockManager() IDockManager                         // property DockManager Getter
	SetDockManager(value IDockManager)                 // property DockManager Setter
	DockSite() bool                                    // property DockSite Getter
	SetDockSite(value bool)                            // property DockSite Setter
	DoubleBuffered() bool                              // property DoubleBuffered Getter
	SetDoubleBuffered(value bool)                      // property DoubleBuffered Setter
	Handle() types.HWND                                // property Handle Getter
	SetHandle(value types.HWND)                        // property Handle Setter
	IsFlipped() bool                                   // property IsFlipped Getter
	IsResizing() bool                                  // property IsResizing Getter
	TabOrder() types.TTabOrder                         // property TabOrder Getter
	SetTabOrder(value types.TTabOrder)                 // property TabOrder Setter
	TabStop() bool                                     // property TabStop Getter
	SetTabStop(value bool)                             // property TabStop Setter
	ParentDoubleBuffered() bool                        // property ParentDoubleBuffered Getter
	SetParentDoubleBuffered(value bool)                // property ParentDoubleBuffered Setter
	ParentWindow() types.HWND                          // property ParentWindow Getter
	SetParentWindow(value types.HWND)                  // property ParentWindow Setter
	Showing() bool                                     // property Showing Getter
	UseDockManager() bool                              // property UseDockManager Getter
	SetUseDockManager(value bool)                      // property UseDockManager Setter
	SetDesignerDeleting(value bool)                    // property DesignerDeleting Setter
	IsSpecialSubControl() bool                         // property IsSpecialSubControl Getter
	VisibleDockClientCount() int32                     // property VisibleDockClientCount Getter
	SetOnAlignInsertBefore(fn TAlignInsertBeforeEvent) // property event
	SetOnAlignPosition(fn TAlignPositionEvent)         // property event
	SetOnDockDrop(fn TDockDropEvent)                   // property event
	SetOnDockOver(fn TDockOverEvent)                   // property event
	SetOnEnter(fn TNotifyEvent)                        // property event
	SetOnExit(fn TNotifyEvent)                         // property event
	SetOnKeyDown(fn TKeyEvent)                         // property event
	SetOnKeyPress(fn TKeyPressEvent)                   // property event
	SetOnKeyUp(fn TKeyEvent)                           // property event
	SetOnUnDock(fn TUnDockEvent)                       // property event
	SetOnUTF8KeyPress(fn TUTF8KeyPressEvent)           // property event
}

type TWinControl struct {
	TControl
}

func (m *TWinControl) ControlAtPosWithPointBool(pos types.TPoint, allowDisabled bool) IControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&pos)), api.PasBool(allowDisabled))
	return AsControl(r)
}

func (m *TWinControl) ControlAtPosWithPointBoolX2(pos types.TPoint, allowDisabled bool, allowWinControls bool) IControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&pos)), api.PasBool(allowDisabled), api.PasBool(allowWinControls))
	return AsControl(r)
}

func (m *TWinControl) ControlAtPosWithPointControlAtPosFlags(pos types.TPoint, flags types.TControlAtPosFlags) IControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&pos)), uintptr(flags))
	return AsControl(r)
}

func (m *TWinControl) ContainsControl(control IControl) bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(control))
	return api.GoBool(r)
}

func (m *TWinControl) ClientRectNeedsInterfaceUpdate() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) CanFocus() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) CanSetFocus() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) GetControlIndex(control IControl) int32 {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(control))
	return int32(r)
}

func (m *TWinControl) Focused() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(11, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) PerformTab(forwardTab bool) bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(12, m.Instance(), api.PasBool(forwardTab))
	return api.GoBool(r)
}

func (m *TWinControl) FindChildControl(controlName string) IControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(13, m.Instance(), api.PasStr(controlName))
	return AsControl(r)
}

func (m *TWinControl) GetEnumeratorControls() IWinControlEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(14, m.Instance())
	return AsWinControlEnumerator(r)
}

func (m *TWinControl) GetEnumeratorControlsReverse() IWinControlEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(15, m.Instance())
	return AsWinControlEnumerator(r)
}

func (m *TWinControl) GetDockCaption(control IControl) string {
	if !m.IsValid() {
		return ""
	}
	r := winControlAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(control))
	return api.GoStr(r)
}

func (m *TWinControl) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(17, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) BrushCreated() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(18, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) IntfUTF8KeyPress(uTF8Key *string, repeatCount int32, systemKey bool) bool {
	if !m.IsValid() {
		return false
	}
	uTF8KeyPtr := api.PasStr(*uTF8Key)
	r := winControlAPI().SysCallN(19, m.Instance(), uintptr(base.UnsafePointer(&uTF8KeyPtr)), uintptr(repeatCount), api.PasBool(systemKey))
	*uTF8Key = api.GoStr(uTF8KeyPtr)
	return api.GoBool(r)
}

func (m *TWinControl) IntfGetDropFilesTarget() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(20, m.Instance())
	return AsWinControl(r)
}

func (m *TWinControl) BeginUpdateBounds() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(21, m.Instance())
}

func (m *TWinControl) EndUpdateBounds() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(22, m.Instance())
}

func (m *TWinControl) LockRealizeBounds() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(23, m.Instance())
}

func (m *TWinControl) UnlockRealizeBounds() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(24, m.Instance())
}

func (m *TWinControl) DoAdjustClientRectChange(invalidateRect bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(25, m.Instance(), api.PasBool(invalidateRect))
}

func (m *TWinControl) InvalidateClientRectCache(withChildControls bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(26, m.Instance(), api.PasBool(withChildControls))
}

func (m *TWinControl) DisableAlign() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(27, m.Instance())
}

func (m *TWinControl) EnableAlign() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(28, m.Instance())
}

func (m *TWinControl) ReAlign() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(29, m.Instance())
}

func (m *TWinControl) ScrollByWS(deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(30, m.Instance(), uintptr(deltaX), uintptr(deltaY))
}

func (m *TWinControl) ScrollBy(deltaX int32, deltaY int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(31, m.Instance(), uintptr(deltaX), uintptr(deltaY))
}

func (m *TWinControl) FixDesignFontsPPIWithChildren(designTimePPI int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(32, m.Instance(), uintptr(designTimePPI))
}

func (m *TWinControl) DockDrop(dragDockObject IDragDockObject, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(33, m.Instance(), base.GetObjectUintptr(dragDockObject), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetControlIndex(control IControl, newIndex int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(34, m.Instance(), base.GetObjectUintptr(control), uintptr(newIndex))
}

func (m *TWinControl) SelectNext(curControl IWinControl, goForward bool, checkTabStop bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(35, m.Instance(), base.GetObjectUintptr(curControl), api.PasBool(goForward), api.PasBool(checkTabStop))
}

func (m *TWinControl) BroadCast(toAllMessage *uintptr) {
	if !m.IsValid() {
		return
	}
	toAllMessagePtr := uintptr(*toAllMessage)
	winControlAPI().SysCallN(36, m.Instance(), uintptr(base.UnsafePointer(&toAllMessagePtr)))
	*toAllMessage = uintptr(toAllMessagePtr)
}

func (m *TWinControl) NotifyControls(msg uint16) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(37, m.Instance(), uintptr(msg))
}

func (m *TWinControl) AddControl() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(38, m.Instance())
}

func (m *TWinControl) InsertControlWithControl(control IControl) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(39, m.Instance(), base.GetObjectUintptr(control))
}

func (m *TWinControl) InsertControlWithControlInt(control IControl, index int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(40, m.Instance(), base.GetObjectUintptr(control), uintptr(index))
}

func (m *TWinControl) RemoveControl(control IControl) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(41, m.Instance(), base.GetObjectUintptr(control))
}

func (m *TWinControl) SetFocus() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(42, m.Instance())
}

func (m *TWinControl) FlipChildren(allLevels bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(43, m.Instance(), api.PasBool(allLevels))
}

func (m *TWinControl) ScaleBy(multiplier int32, divider int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(44, m.Instance(), uintptr(multiplier), uintptr(divider))
}

func (m *TWinControl) UpdateDockCaption(exclude IControl) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(45, m.Instance(), base.GetObjectUintptr(exclude))
}

func (m *TWinControl) GetTabOrderList(list IFPList) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(46, m.Instance(), base.GetObjectUintptr(list))
}

func (m *TWinControl) HandleNeeded() {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(47, m.Instance())
}

func (m *TWinControl) EraseBackground(dC types.HDC) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(48, m.Instance(), uintptr(dC))
}

func (m *TWinControl) PaintToWithHDCIntX2(dC types.HDC, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(49, m.Instance(), uintptr(dC), uintptr(X), uintptr(Y))
}

func (m *TWinControl) PaintToWithCanvasIntX2(canvas ICanvas, X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(50, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetShapeWithBitmap(shape IBitmap) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(51, m.Instance(), base.GetObjectUintptr(shape))
}

func (m *TWinControl) SetShapeWithRegion(shape IRegion) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(52, m.Instance(), base.GetObjectUintptr(shape))
}

func (m *TWinControl) BorderWidth() types.TBorderWidth {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(53, 0, m.Instance())
	return types.TBorderWidth(r)
}

func (m *TWinControl) SetBorderWidth(value types.TBorderWidth) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(53, 1, m.Instance(), uintptr(value))
}

func (m *TWinControl) BoundsLockCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(54, m.Instance())
	return int32(r)
}

func (m *TWinControl) Brush() IBrush {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(55, m.Instance())
	return AsBrush(r)
}

func (m *TWinControl) CachedClientHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(56, m.Instance())
	return int32(r)
}

func (m *TWinControl) CachedClientWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(57, m.Instance())
	return int32(r)
}

func (m *TWinControl) ChildSizing() IControlChildSizing {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(58, 0, m.Instance())
	return AsControlChildSizing(r)
}

func (m *TWinControl) SetChildSizing(value IControlChildSizing) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(58, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TWinControl) ControlCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(59, m.Instance())
	return int32(r)
}

func (m *TWinControl) Controls(index int32) IControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(60, m.Instance(), uintptr(index))
	return AsControl(r)
}

func (m *TWinControl) DefWndProc() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(61, 0, m.Instance())
	return uintptr(r)
}

func (m *TWinControl) SetDefWndProc(value uintptr) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(61, 1, m.Instance(), uintptr(value))
}

func (m *TWinControl) DockClientCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(62, m.Instance())
	return int32(r)
}

func (m *TWinControl) DockClients(index int32) IControl {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(63, m.Instance(), uintptr(index))
	return AsControl(r)
}

func (m *TWinControl) DockManager() IDockManager {
	if !m.IsValid() {
		return nil
	}
	r := winControlAPI().SysCallN(64, 0, m.Instance())
	return AsDockManager(r)
}

func (m *TWinControl) SetDockManager(value IDockManager) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(64, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TWinControl) DockSite() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(65, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) SetDockSite(value bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(65, 1, m.Instance(), api.PasBool(value))
}

func (m *TWinControl) DoubleBuffered() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(66, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) SetDoubleBuffered(value bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(66, 1, m.Instance(), api.PasBool(value))
}

func (m *TWinControl) Handle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(67, 0, m.Instance())
	return types.HWND(r)
}

func (m *TWinControl) SetHandle(value types.HWND) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(67, 1, m.Instance(), uintptr(value))
}

func (m *TWinControl) IsFlipped() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(68, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) IsResizing() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(69, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) TabOrder() types.TTabOrder {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(70, 0, m.Instance())
	return types.TTabOrder(r)
}

func (m *TWinControl) SetTabOrder(value types.TTabOrder) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(70, 1, m.Instance(), uintptr(value))
}

func (m *TWinControl) TabStop() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(71, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) SetTabStop(value bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(71, 1, m.Instance(), api.PasBool(value))
}

func (m *TWinControl) ParentDoubleBuffered() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(72, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) SetParentDoubleBuffered(value bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(72, 1, m.Instance(), api.PasBool(value))
}

func (m *TWinControl) ParentWindow() types.HWND {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(73, 0, m.Instance())
	return types.HWND(r)
}

func (m *TWinControl) SetParentWindow(value types.HWND) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(73, 1, m.Instance(), uintptr(value))
}

func (m *TWinControl) Showing() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(74, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) UseDockManager() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(75, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) SetUseDockManager(value bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(75, 1, m.Instance(), api.PasBool(value))
}

func (m *TWinControl) SetDesignerDeleting(value bool) {
	if !m.IsValid() {
		return
	}
	winControlAPI().SysCallN(76, m.Instance(), api.PasBool(value))
}

func (m *TWinControl) IsSpecialSubControl() bool {
	if !m.IsValid() {
		return false
	}
	r := winControlAPI().SysCallN(77, m.Instance())
	return api.GoBool(r)
}

func (m *TWinControl) VisibleDockClientCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := winControlAPI().SysCallN(78, m.Instance())
	return int32(r)
}

func (m *TWinControl) SetOnAlignInsertBefore(fn TAlignInsertBeforeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTAlignInsertBeforeEvent(fn)
	base.SetEvent(m, 79, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnAlignPosition(fn TAlignPositionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTAlignPositionEvent(fn)
	base.SetEvent(m, 80, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnDockDrop(fn TDockDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDockDropEvent(fn)
	base.SetEvent(m, 81, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnDockOver(fn TDockOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDockOverEvent(fn)
	base.SetEvent(m, 82, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 83, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnExit(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 84, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnKeyDown(fn TKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTKeyEvent(fn)
	base.SetEvent(m, 85, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnKeyPress(fn TKeyPressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTKeyPressEvent(fn)
	base.SetEvent(m, 86, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnKeyUp(fn TKeyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTKeyEvent(fn)
	base.SetEvent(m, 87, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnUnDock(fn TUnDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUnDockEvent(fn)
	base.SetEvent(m, 88, winControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TWinControl) SetOnUTF8KeyPress(fn TUTF8KeyPressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUTF8KeyPressEvent(fn)
	base.SetEvent(m, 89, winControlAPI(), api.MakeEventDataPtr(cb))
}

// WinControl  is static instance
var WinControl _WinControlClass

// _WinControlClass is class type defined by TWinControl
type _WinControlClass uintptr

func (_WinControlClass) CreateParentedControl(parentWindow types.HWND) IWinControl {
	r := winControlAPI().SysCallN(7, uintptr(parentWindow))
	return AsWinControl(r)
}

// NewWinControl class constructor
func NewWinControl(theOwner IComponent) IWinControl {
	r := winControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsWinControl(r)
}

// NewWinControlParented class constructor
func NewWinControlParented(parentWindow types.HWND) IWinControl {
	r := winControlAPI().SysCallN(1, uintptr(parentWindow))
	return AsWinControl(r)
}

func TWinControlClass() types.TClass {
	r := winControlAPI().SysCallN(90)
	return types.TClass(r)
}

var (
	winControlOnce   base.Once
	winControlImport *imports.Imports = nil
)

func winControlAPI() *imports.Imports {
	winControlOnce.Do(func() {
		winControlImport = api.NewDefaultImports()
		winControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TWinControl_Create", 0), // constructor NewWinControl
			/* 1 */ imports.NewTable("TWinControl_CreateParented", 0), // constructor NewWinControlParented
			/* 2 */ imports.NewTable("TWinControl_ControlAtPosWithPointBool", 0), // function ControlAtPosWithPointBool
			/* 3 */ imports.NewTable("TWinControl_ControlAtPosWithPointBoolX2", 0), // function ControlAtPosWithPointBoolX2
			/* 4 */ imports.NewTable("TWinControl_ControlAtPosWithPointControlAtPosFlags", 0), // function ControlAtPosWithPointControlAtPosFlags
			/* 5 */ imports.NewTable("TWinControl_ContainsControl", 0), // function ContainsControl
			/* 6 */ imports.NewTable("TWinControl_ClientRectNeedsInterfaceUpdate", 0), // function ClientRectNeedsInterfaceUpdate
			/* 7 */ imports.NewTable("TWinControl_CreateParentedControl", 0), // static function CreateParentedControl
			/* 8 */ imports.NewTable("TWinControl_CanFocus", 0), // function CanFocus
			/* 9 */ imports.NewTable("TWinControl_CanSetFocus", 0), // function CanSetFocus
			/* 10 */ imports.NewTable("TWinControl_GetControlIndex", 0), // function GetControlIndex
			/* 11 */ imports.NewTable("TWinControl_Focused", 0), // function Focused
			/* 12 */ imports.NewTable("TWinControl_PerformTab", 0), // function PerformTab
			/* 13 */ imports.NewTable("TWinControl_FindChildControl", 0), // function FindChildControl
			/* 14 */ imports.NewTable("TWinControl_GetEnumeratorControls", 0), // function GetEnumeratorControls
			/* 15 */ imports.NewTable("TWinControl_GetEnumeratorControlsReverse", 0), // function GetEnumeratorControlsReverse
			/* 16 */ imports.NewTable("TWinControl_GetDockCaption", 0), // function GetDockCaption
			/* 17 */ imports.NewTable("TWinControl_HandleAllocated", 0), // function HandleAllocated
			/* 18 */ imports.NewTable("TWinControl_BrushCreated", 0), // function BrushCreated
			/* 19 */ imports.NewTable("TWinControl_IntfUTF8KeyPress", 0), // function IntfUTF8KeyPress
			/* 20 */ imports.NewTable("TWinControl_IntfGetDropFilesTarget", 0), // function IntfGetDropFilesTarget
			/* 21 */ imports.NewTable("TWinControl_BeginUpdateBounds", 0), // procedure BeginUpdateBounds
			/* 22 */ imports.NewTable("TWinControl_EndUpdateBounds", 0), // procedure EndUpdateBounds
			/* 23 */ imports.NewTable("TWinControl_LockRealizeBounds", 0), // procedure LockRealizeBounds
			/* 24 */ imports.NewTable("TWinControl_UnlockRealizeBounds", 0), // procedure UnlockRealizeBounds
			/* 25 */ imports.NewTable("TWinControl_DoAdjustClientRectChange", 0), // procedure DoAdjustClientRectChange
			/* 26 */ imports.NewTable("TWinControl_InvalidateClientRectCache", 0), // procedure InvalidateClientRectCache
			/* 27 */ imports.NewTable("TWinControl_DisableAlign", 0), // procedure DisableAlign
			/* 28 */ imports.NewTable("TWinControl_EnableAlign", 0), // procedure EnableAlign
			/* 29 */ imports.NewTable("TWinControl_ReAlign", 0), // procedure ReAlign
			/* 30 */ imports.NewTable("TWinControl_ScrollBy_WS", 0), // procedure ScrollByWS
			/* 31 */ imports.NewTable("TWinControl_ScrollBy", 0), // procedure ScrollBy
			/* 32 */ imports.NewTable("TWinControl_FixDesignFontsPPIWithChildren", 0), // procedure FixDesignFontsPPIWithChildren
			/* 33 */ imports.NewTable("TWinControl_DockDrop", 0), // procedure DockDrop
			/* 34 */ imports.NewTable("TWinControl_SetControlIndex", 0), // procedure SetControlIndex
			/* 35 */ imports.NewTable("TWinControl_SelectNext", 0), // procedure SelectNext
			/* 36 */ imports.NewTable("TWinControl_BroadCast", 0), // procedure BroadCast
			/* 37 */ imports.NewTable("TWinControl_NotifyControls", 0), // procedure NotifyControls
			/* 38 */ imports.NewTable("TWinControl_AddControl", 0), // procedure AddControl
			/* 39 */ imports.NewTable("TWinControl_InsertControlWithControl", 0), // procedure InsertControlWithControl
			/* 40 */ imports.NewTable("TWinControl_InsertControlWithControlInt", 0), // procedure InsertControlWithControlInt
			/* 41 */ imports.NewTable("TWinControl_RemoveControl", 0), // procedure RemoveControl
			/* 42 */ imports.NewTable("TWinControl_SetFocus", 0), // procedure SetFocus
			/* 43 */ imports.NewTable("TWinControl_FlipChildren", 0), // procedure FlipChildren
			/* 44 */ imports.NewTable("TWinControl_ScaleBy", 0), // procedure ScaleBy
			/* 45 */ imports.NewTable("TWinControl_UpdateDockCaption", 0), // procedure UpdateDockCaption
			/* 46 */ imports.NewTable("TWinControl_GetTabOrderList", 0), // procedure GetTabOrderList
			/* 47 */ imports.NewTable("TWinControl_HandleNeeded", 0), // procedure HandleNeeded
			/* 48 */ imports.NewTable("TWinControl_EraseBackground", 0), // procedure EraseBackground
			/* 49 */ imports.NewTable("TWinControl_PaintToWithHDCIntX2", 0), // procedure PaintToWithHDCIntX2
			/* 50 */ imports.NewTable("TWinControl_PaintToWithCanvasIntX2", 0), // procedure PaintToWithCanvasIntX2
			/* 51 */ imports.NewTable("TWinControl_SetShapeWithBitmap", 0), // procedure SetShapeWithBitmap
			/* 52 */ imports.NewTable("TWinControl_SetShapeWithRegion", 0), // procedure SetShapeWithRegion
			/* 53 */ imports.NewTable("TWinControl_BorderWidth", 0), // property BorderWidth
			/* 54 */ imports.NewTable("TWinControl_BoundsLockCount", 0), // property BoundsLockCount
			/* 55 */ imports.NewTable("TWinControl_Brush", 0), // property Brush
			/* 56 */ imports.NewTable("TWinControl_CachedClientHeight", 0), // property CachedClientHeight
			/* 57 */ imports.NewTable("TWinControl_CachedClientWidth", 0), // property CachedClientWidth
			/* 58 */ imports.NewTable("TWinControl_ChildSizing", 0), // property ChildSizing
			/* 59 */ imports.NewTable("TWinControl_ControlCount", 0), // property ControlCount
			/* 60 */ imports.NewTable("TWinControl_Controls", 0), // property Controls
			/* 61 */ imports.NewTable("TWinControl_DefWndProc", 0), // property DefWndProc
			/* 62 */ imports.NewTable("TWinControl_DockClientCount", 0), // property DockClientCount
			/* 63 */ imports.NewTable("TWinControl_DockClients", 0), // property DockClients
			/* 64 */ imports.NewTable("TWinControl_DockManager", 0), // property DockManager
			/* 65 */ imports.NewTable("TWinControl_DockSite", 0), // property DockSite
			/* 66 */ imports.NewTable("TWinControl_DoubleBuffered", 0), // property DoubleBuffered
			/* 67 */ imports.NewTable("TWinControl_Handle", 0), // property Handle
			/* 68 */ imports.NewTable("TWinControl_IsFlipped", 0), // property IsFlipped
			/* 69 */ imports.NewTable("TWinControl_IsResizing", 0), // property IsResizing
			/* 70 */ imports.NewTable("TWinControl_TabOrder", 0), // property TabOrder
			/* 71 */ imports.NewTable("TWinControl_TabStop", 0), // property TabStop
			/* 72 */ imports.NewTable("TWinControl_ParentDoubleBuffered", 0), // property ParentDoubleBuffered
			/* 73 */ imports.NewTable("TWinControl_ParentWindow", 0), // property ParentWindow
			/* 74 */ imports.NewTable("TWinControl_Showing", 0), // property Showing
			/* 75 */ imports.NewTable("TWinControl_UseDockManager", 0), // property UseDockManager
			/* 76 */ imports.NewTable("TWinControl_DesignerDeleting", 0), // property DesignerDeleting
			/* 77 */ imports.NewTable("TWinControl_IsSpecialSubControl", 0), // property IsSpecialSubControl
			/* 78 */ imports.NewTable("TWinControl_VisibleDockClientCount", 0), // property VisibleDockClientCount
			/* 79 */ imports.NewTable("TWinControl_OnAlignInsertBefore", 0), // event OnAlignInsertBefore
			/* 80 */ imports.NewTable("TWinControl_OnAlignPosition", 0), // event OnAlignPosition
			/* 81 */ imports.NewTable("TWinControl_OnDockDrop", 0), // event OnDockDrop
			/* 82 */ imports.NewTable("TWinControl_OnDockOver", 0), // event OnDockOver
			/* 83 */ imports.NewTable("TWinControl_OnEnter", 0), // event OnEnter
			/* 84 */ imports.NewTable("TWinControl_OnExit", 0), // event OnExit
			/* 85 */ imports.NewTable("TWinControl_OnKeyDown", 0), // event OnKeyDown
			/* 86 */ imports.NewTable("TWinControl_OnKeyPress", 0), // event OnKeyPress
			/* 87 */ imports.NewTable("TWinControl_OnKeyUp", 0), // event OnKeyUp
			/* 88 */ imports.NewTable("TWinControl_OnUnDock", 0), // event OnUnDock
			/* 89 */ imports.NewTable("TWinControl_OnUTF8KeyPress", 0), // event OnUTF8KeyPress
			/* 90 */ imports.NewTable("TWinControl_TClass", 0), // function TWinControlClass
		}
	})
	return winControlImport
}
