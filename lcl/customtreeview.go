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

// ICustomTreeView Parent: ICustomControl
type ICustomTreeView interface {
	ICustomControl
	AccessibilityOn() bool                                                                                  // property
	SetAccessibilityOn(AValue bool)                                                                         // property
	BackgroundColor() TColor                                                                                // property
	SetBackgroundColor(AValue TColor)                                                                       // property
	BottomItem() ITreeNode                                                                                  // property
	SetBottomItem(AValue ITreeNode)                                                                         // property
	DefaultItemHeight() int32                                                                               // property
	SetDefaultItemHeight(AValue int32)                                                                      // property
	DropTarget() ITreeNode                                                                                  // property
	SetDropTarget(AValue ITreeNode)                                                                         // property
	ExpandSignColor() TColor                                                                                // property
	SetExpandSignColor(AValue TColor)                                                                       // property
	ExpandSignSize() int32                                                                                  // property
	SetExpandSignSize(AValue int32)                                                                         // property
	ExpandSignWidth() int32                                                                                 // property
	SetExpandSignWidth(AValue int32)                                                                        // property
	ExpandSignType() TTreeViewExpandSignType                                                                // property
	SetExpandSignType(AValue TTreeViewExpandSignType)                                                       // property
	Images() ICustomImageList                                                                               // property
	SetImages(AValue ICustomImageList)                                                                      // property
	ImagesWidth() int32                                                                                     // property
	SetImagesWidth(AValue int32)                                                                            // property
	InsertMarkNode() ITreeNode                                                                              // property
	SetInsertMarkNode(AValue ITreeNode)                                                                     // property
	InsertMarkType() TTreeViewInsertMarkType                                                                // property
	SetInsertMarkType(AValue TTreeViewInsertMarkType)                                                       // property
	Items() ITreeNodes                                                                                      // property
	SetItems(AValue ITreeNodes)                                                                             // property
	KeepCollapsedNodes() bool                                                                               // property
	SetKeepCollapsedNodes(AValue bool)                                                                      // property
	MultiSelectStyle() TMultiSelectStyle                                                                    // property
	SetMultiSelectStyle(AValue TMultiSelectStyle)                                                           // property
	Options() TTreeViewOptions                                                                              // property
	SetOptions(AValue TTreeViewOptions)                                                                     // property
	ScrollBars() TScrollStyle                                                                               // property
	SetScrollBars(AValue TScrollStyle)                                                                      // property
	Selected() ITreeNode                                                                                    // property
	SetSelected(AValue ITreeNode)                                                                           // property
	SelectionColor() TColor                                                                                 // property
	SetSelectionColor(AValue TColor)                                                                        // property
	SelectionCount() uint32                                                                                 // property
	SelectionFontColor() TColor                                                                             // property
	SetSelectionFontColor(AValue TColor)                                                                    // property
	SelectionFontColorUsed() bool                                                                           // property
	SetSelectionFontColorUsed(AValue bool)                                                                  // property
	Selections(AIndex int32) ITreeNode                                                                      // property
	SeparatorColor() TColor                                                                                 // property
	SetSeparatorColor(AValue TColor)                                                                        // property
	StateImages() ICustomImageList                                                                          // property
	SetStateImages(AValue ICustomImageList)                                                                 // property
	StateImagesWidth() int32                                                                                // property
	SetStateImagesWidth(AValue int32)                                                                       // property
	TopItem() ITreeNode                                                                                     // property
	SetTopItem(AValue ITreeNode)                                                                            // property
	TreeLineColor() TColor                                                                                  // property
	SetTreeLineColor(AValue TColor)                                                                         // property
	TreeLinePenStyle() TPenStyle                                                                            // property
	SetTreeLinePenStyle(AValue TPenStyle)                                                                   // property
	AlphaSort() bool                                                                                        // function
	CustomSort(fn TTreeNodeCompare) bool                                                                    // function
	DefaultTreeViewSort(Node1, Node2 ITreeNode) int32                                                       // function
	GetHitTestInfoAt(X, Y int32) THitTests                                                                  // function
	GetNodeAt(X, Y int32) ITreeNode                                                                         // function
	GetNodeWithExpandSignAt(X, Y int32) ITreeNode                                                           // function
	IsEditing() bool                                                                                        // function
	GetFirstMultiSelected() ITreeNode                                                                       // function
	GetLastMultiSelected() ITreeNode                                                                        // function
	SelectionVisible() bool                                                                                 // function
	StoreCurrentSelection() IStringList                                                                     // function
	ClearSelection(KeepPrimary bool)                                                                        // procedure
	ConsistencyCheck()                                                                                      // procedure
	GetInsertMarkAt(X, Y int32, OutNInsertMarkNode *ITreeNode, OutNInsertMarkType *TTreeViewInsertMarkType) // procedure
	SetInsertMark(AnInsertMarkNode ITreeNode, AnInsertMarkType TTreeViewInsertMarkType)                     // procedure
	SetInsertMarkAt(X, Y int32)                                                                             // procedure
	BeginUpdate()                                                                                           // procedure
	EndUpdate()                                                                                             // procedure
	FullCollapse()                                                                                          // procedure
	FullExpand()                                                                                            // procedure
	LoadFromFile(FileName string)                                                                           // procedure
	LoadFromStream(Stream IStream)                                                                          // procedure
	SaveToFile(FileName string)                                                                             // procedure
	SaveToStream(Stream IStream)                                                                            // procedure
	WriteDebugReport(Prefix string, AllNodes bool)                                                          // procedure
	LockSelectionChangeEvent()                                                                              // procedure
	UnlockSelectionChangeEvent()                                                                            // procedure
	Select(Node ITreeNode, ShiftState TShiftState)                                                          // procedure
	Select1(Nodes IList)                                                                                    // procedure
	MakeSelectionVisible()                                                                                  // procedure
	ClearInvisibleSelection()                                                                               // procedure
	ApplyStoredSelection(ASelection IStringList, FreeList bool)                                             // procedure
	MoveToNextNode(ASelect bool)                                                                            // procedure
	MoveToPrevNode(ASelect bool)                                                                            // procedure
	MovePageDown(ASelect bool)                                                                              // procedure
	MovePageUp(ASelect bool)                                                                                // procedure
	MoveLeft(ASelect bool)                                                                                  // procedure
	MoveRight(ASelect bool)                                                                                 // procedure
	MoveExpand(ASelect bool)                                                                                // procedure
	MoveCollapse(ASelect bool)                                                                              // procedure
	MoveHome(ASelect bool)                                                                                  // procedure
	MoveEnd(ASelect bool)                                                                                   // procedure
}

// TCustomTreeView Parent: TCustomControl
type TCustomTreeView struct {
	TCustomControl
	customSortPtr uintptr
}

func NewCustomTreeView(AnOwner IComponent) ICustomTreeView {
	r1 := customTreeViewImportAPI().SysCallN(10, GetObjectUintptr(AnOwner))
	return AsCustomTreeView(r1)
}

func (m *TCustomTreeView) AccessibilityOn() bool {
	r1 := customTreeViewImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTreeView) SetAccessibilityOn(AValue bool) {
	customTreeViewImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTreeView) BackgroundColor() TColor {
	r1 := customTreeViewImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetBackgroundColor(AValue TColor) {
	customTreeViewImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) BottomItem() ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetBottomItem(AValue ITreeNode) {
	customTreeViewImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) DefaultItemHeight() int32 {
	r1 := customTreeViewImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetDefaultItemHeight(AValue int32) {
	customTreeViewImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) DropTarget() ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetDropTarget(AValue ITreeNode) {
	customTreeViewImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignColor() TColor {
	r1 := customTreeViewImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetExpandSignColor(AValue TColor) {
	customTreeViewImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignSize() int32 {
	r1 := customTreeViewImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetExpandSignSize(AValue int32) {
	customTreeViewImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignWidth() int32 {
	r1 := customTreeViewImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetExpandSignWidth(AValue int32) {
	customTreeViewImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignType() TTreeViewExpandSignType {
	r1 := customTreeViewImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TTreeViewExpandSignType(r1)
}

func (m *TCustomTreeView) SetExpandSignType(AValue TTreeViewExpandSignType) {
	customTreeViewImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Images() ICustomImageList {
	r1 := customTreeViewImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTreeView) SetImages(AValue ICustomImageList) {
	customTreeViewImportAPI().SysCallN(28, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) ImagesWidth() int32 {
	r1 := customTreeViewImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetImagesWidth(AValue int32) {
	customTreeViewImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) InsertMarkNode() ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetInsertMarkNode(AValue ITreeNode) {
	customTreeViewImportAPI().SysCallN(30, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) InsertMarkType() TTreeViewInsertMarkType {
	r1 := customTreeViewImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return TTreeViewInsertMarkType(r1)
}

func (m *TCustomTreeView) SetInsertMarkType(AValue TTreeViewInsertMarkType) {
	customTreeViewImportAPI().SysCallN(31, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Items() ITreeNodes {
	r1 := customTreeViewImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return AsTreeNodes(r1)
}

func (m *TCustomTreeView) SetItems(AValue ITreeNodes) {
	customTreeViewImportAPI().SysCallN(33, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) KeepCollapsedNodes() bool {
	r1 := customTreeViewImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTreeView) SetKeepCollapsedNodes(AValue bool) {
	customTreeViewImportAPI().SysCallN(34, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTreeView) MultiSelectStyle() TMultiSelectStyle {
	r1 := customTreeViewImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return TMultiSelectStyle(r1)
}

func (m *TCustomTreeView) SetMultiSelectStyle(AValue TMultiSelectStyle) {
	customTreeViewImportAPI().SysCallN(49, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Options() TTreeViewOptions {
	r1 := customTreeViewImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return TTreeViewOptions(r1)
}

func (m *TCustomTreeView) SetOptions(AValue TTreeViewOptions) {
	customTreeViewImportAPI().SysCallN(50, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ScrollBars() TScrollStyle {
	r1 := customTreeViewImportAPI().SysCallN(53, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TCustomTreeView) SetScrollBars(AValue TScrollStyle) {
	customTreeViewImportAPI().SysCallN(53, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Selected() ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(56, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetSelected(AValue ITreeNode) {
	customTreeViewImportAPI().SysCallN(56, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) SelectionColor() TColor {
	r1 := customTreeViewImportAPI().SysCallN(57, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetSelectionColor(AValue TColor) {
	customTreeViewImportAPI().SysCallN(57, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) SelectionCount() uint32 {
	r1 := customTreeViewImportAPI().SysCallN(58, m.Instance())
	return uint32(r1)
}

func (m *TCustomTreeView) SelectionFontColor() TColor {
	r1 := customTreeViewImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetSelectionFontColor(AValue TColor) {
	customTreeViewImportAPI().SysCallN(59, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) SelectionFontColorUsed() bool {
	r1 := customTreeViewImportAPI().SysCallN(60, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTreeView) SetSelectionFontColorUsed(AValue bool) {
	customTreeViewImportAPI().SysCallN(60, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTreeView) Selections(AIndex int32) ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(62, m.Instance(), uintptr(AIndex))
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SeparatorColor() TColor {
	r1 := customTreeViewImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetSeparatorColor(AValue TColor) {
	customTreeViewImportAPI().SysCallN(63, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) StateImages() ICustomImageList {
	r1 := customTreeViewImportAPI().SysCallN(66, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTreeView) SetStateImages(AValue ICustomImageList) {
	customTreeViewImportAPI().SysCallN(66, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) StateImagesWidth() int32 {
	r1 := customTreeViewImportAPI().SysCallN(67, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetStateImagesWidth(AValue int32) {
	customTreeViewImportAPI().SysCallN(67, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) TopItem() ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(69, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetTopItem(AValue ITreeNode) {
	customTreeViewImportAPI().SysCallN(69, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) TreeLineColor() TColor {
	r1 := customTreeViewImportAPI().SysCallN(70, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetTreeLineColor(AValue TColor) {
	customTreeViewImportAPI().SysCallN(70, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) TreeLinePenStyle() TPenStyle {
	r1 := customTreeViewImportAPI().SysCallN(71, 0, m.Instance(), 0)
	return TPenStyle(r1)
}

func (m *TCustomTreeView) SetTreeLinePenStyle(AValue TPenStyle) {
	customTreeViewImportAPI().SysCallN(71, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) AlphaSort() bool {
	r1 := customTreeViewImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTreeView) CustomSort(fn TTreeNodeCompare) bool {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	r1 := customTreeViewImportAPI().SysCallN(11, m.Instance(), m.customSortPtr)
	return GoBool(r1)
}

func (m *TCustomTreeView) DefaultTreeViewSort(Node1, Node2 ITreeNode) int32 {
	r1 := customTreeViewImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(Node1), GetObjectUintptr(Node2))
	return int32(r1)
}

func (m *TCustomTreeView) GetHitTestInfoAt(X, Y int32) THitTests {
	r1 := customTreeViewImportAPI().SysCallN(23, m.Instance(), uintptr(X), uintptr(Y))
	return THitTests(r1)
}

func (m *TCustomTreeView) GetNodeAt(X, Y int32) ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(26, m.Instance(), uintptr(X), uintptr(Y))
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) GetNodeWithExpandSignAt(X, Y int32) ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(27, m.Instance(), uintptr(X), uintptr(Y))
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) IsEditing() bool {
	r1 := customTreeViewImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTreeView) GetFirstMultiSelected() ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(22, m.Instance())
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) GetLastMultiSelected() ITreeNode {
	r1 := customTreeViewImportAPI().SysCallN(25, m.Instance())
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SelectionVisible() bool {
	r1 := customTreeViewImportAPI().SysCallN(61, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTreeView) StoreCurrentSelection() IStringList {
	r1 := customTreeViewImportAPI().SysCallN(68, m.Instance())
	return AsStringList(r1)
}

func CustomTreeViewClass() TClass {
	ret := customTreeViewImportAPI().SysCallN(6)
	return TClass(ret)
}

func (m *TCustomTreeView) ClearSelection(KeepPrimary bool) {
	customTreeViewImportAPI().SysCallN(8, m.Instance(), PascalBool(KeepPrimary))
}

func (m *TCustomTreeView) ConsistencyCheck() {
	customTreeViewImportAPI().SysCallN(9, m.Instance())
}

func (m *TCustomTreeView) GetInsertMarkAt(X, Y int32, OutNInsertMarkNode *ITreeNode, OutNInsertMarkType *TTreeViewInsertMarkType) {
	var result1 uintptr
	var result2 uintptr
	customTreeViewImportAPI().SysCallN(24, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutNInsertMarkNode = AsTreeNode(result1)
	*OutNInsertMarkType = TTreeViewInsertMarkType(result2)
}

func (m *TCustomTreeView) SetInsertMark(AnInsertMarkNode ITreeNode, AnInsertMarkType TTreeViewInsertMarkType) {
	customTreeViewImportAPI().SysCallN(64, m.Instance(), GetObjectUintptr(AnInsertMarkNode), uintptr(AnInsertMarkType))
}

func (m *TCustomTreeView) SetInsertMarkAt(X, Y int32) {
	customTreeViewImportAPI().SysCallN(65, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TCustomTreeView) BeginUpdate() {
	customTreeViewImportAPI().SysCallN(4, m.Instance())
}

func (m *TCustomTreeView) EndUpdate() {
	customTreeViewImportAPI().SysCallN(15, m.Instance())
}

func (m *TCustomTreeView) FullCollapse() {
	customTreeViewImportAPI().SysCallN(20, m.Instance())
}

func (m *TCustomTreeView) FullExpand() {
	customTreeViewImportAPI().SysCallN(21, m.Instance())
}

func (m *TCustomTreeView) LoadFromFile(FileName string) {
	customTreeViewImportAPI().SysCallN(35, m.Instance(), PascalStr(FileName))
}

func (m *TCustomTreeView) LoadFromStream(Stream IStream) {
	customTreeViewImportAPI().SysCallN(36, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TCustomTreeView) SaveToFile(FileName string) {
	customTreeViewImportAPI().SysCallN(51, m.Instance(), PascalStr(FileName))
}

func (m *TCustomTreeView) SaveToStream(Stream IStream) {
	customTreeViewImportAPI().SysCallN(52, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TCustomTreeView) WriteDebugReport(Prefix string, AllNodes bool) {
	customTreeViewImportAPI().SysCallN(73, m.Instance(), PascalStr(Prefix), PascalBool(AllNodes))
}

func (m *TCustomTreeView) LockSelectionChangeEvent() {
	customTreeViewImportAPI().SysCallN(37, m.Instance())
}

func (m *TCustomTreeView) UnlockSelectionChangeEvent() {
	customTreeViewImportAPI().SysCallN(72, m.Instance())
}

func (m *TCustomTreeView) Select(Node ITreeNode, ShiftState TShiftState) {
	customTreeViewImportAPI().SysCallN(54, m.Instance(), GetObjectUintptr(Node), uintptr(ShiftState))
}

func (m *TCustomTreeView) Select1(Nodes IList) {
	customTreeViewImportAPI().SysCallN(55, m.Instance(), GetObjectUintptr(Nodes))
}

func (m *TCustomTreeView) MakeSelectionVisible() {
	customTreeViewImportAPI().SysCallN(38, m.Instance())
}

func (m *TCustomTreeView) ClearInvisibleSelection() {
	customTreeViewImportAPI().SysCallN(7, m.Instance())
}

func (m *TCustomTreeView) ApplyStoredSelection(ASelection IStringList, FreeList bool) {
	customTreeViewImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(ASelection), PascalBool(FreeList))
}

func (m *TCustomTreeView) MoveToNextNode(ASelect bool) {
	customTreeViewImportAPI().SysCallN(47, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveToPrevNode(ASelect bool) {
	customTreeViewImportAPI().SysCallN(48, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MovePageDown(ASelect bool) {
	customTreeViewImportAPI().SysCallN(44, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MovePageUp(ASelect bool) {
	customTreeViewImportAPI().SysCallN(45, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveLeft(ASelect bool) {
	customTreeViewImportAPI().SysCallN(43, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveRight(ASelect bool) {
	customTreeViewImportAPI().SysCallN(46, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveExpand(ASelect bool) {
	customTreeViewImportAPI().SysCallN(41, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveCollapse(ASelect bool) {
	customTreeViewImportAPI().SysCallN(39, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveHome(ASelect bool) {
	customTreeViewImportAPI().SysCallN(42, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveEnd(ASelect bool) {
	customTreeViewImportAPI().SysCallN(40, m.Instance(), PascalBool(ASelect))
}

var (
	customTreeViewImport       *imports.Imports = nil
	customTreeViewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomTreeView_AccessibilityOn", 0),
		/*1*/ imports.NewTable("CustomTreeView_AlphaSort", 0),
		/*2*/ imports.NewTable("CustomTreeView_ApplyStoredSelection", 0),
		/*3*/ imports.NewTable("CustomTreeView_BackgroundColor", 0),
		/*4*/ imports.NewTable("CustomTreeView_BeginUpdate", 0),
		/*5*/ imports.NewTable("CustomTreeView_BottomItem", 0),
		/*6*/ imports.NewTable("CustomTreeView_Class", 0),
		/*7*/ imports.NewTable("CustomTreeView_ClearInvisibleSelection", 0),
		/*8*/ imports.NewTable("CustomTreeView_ClearSelection", 0),
		/*9*/ imports.NewTable("CustomTreeView_ConsistencyCheck", 0),
		/*10*/ imports.NewTable("CustomTreeView_Create", 0),
		/*11*/ imports.NewTable("CustomTreeView_CustomSort", 0),
		/*12*/ imports.NewTable("CustomTreeView_DefaultItemHeight", 0),
		/*13*/ imports.NewTable("CustomTreeView_DefaultTreeViewSort", 0),
		/*14*/ imports.NewTable("CustomTreeView_DropTarget", 0),
		/*15*/ imports.NewTable("CustomTreeView_EndUpdate", 0),
		/*16*/ imports.NewTable("CustomTreeView_ExpandSignColor", 0),
		/*17*/ imports.NewTable("CustomTreeView_ExpandSignSize", 0),
		/*18*/ imports.NewTable("CustomTreeView_ExpandSignType", 0),
		/*19*/ imports.NewTable("CustomTreeView_ExpandSignWidth", 0),
		/*20*/ imports.NewTable("CustomTreeView_FullCollapse", 0),
		/*21*/ imports.NewTable("CustomTreeView_FullExpand", 0),
		/*22*/ imports.NewTable("CustomTreeView_GetFirstMultiSelected", 0),
		/*23*/ imports.NewTable("CustomTreeView_GetHitTestInfoAt", 0),
		/*24*/ imports.NewTable("CustomTreeView_GetInsertMarkAt", 0),
		/*25*/ imports.NewTable("CustomTreeView_GetLastMultiSelected", 0),
		/*26*/ imports.NewTable("CustomTreeView_GetNodeAt", 0),
		/*27*/ imports.NewTable("CustomTreeView_GetNodeWithExpandSignAt", 0),
		/*28*/ imports.NewTable("CustomTreeView_Images", 0),
		/*29*/ imports.NewTable("CustomTreeView_ImagesWidth", 0),
		/*30*/ imports.NewTable("CustomTreeView_InsertMarkNode", 0),
		/*31*/ imports.NewTable("CustomTreeView_InsertMarkType", 0),
		/*32*/ imports.NewTable("CustomTreeView_IsEditing", 0),
		/*33*/ imports.NewTable("CustomTreeView_Items", 0),
		/*34*/ imports.NewTable("CustomTreeView_KeepCollapsedNodes", 0),
		/*35*/ imports.NewTable("CustomTreeView_LoadFromFile", 0),
		/*36*/ imports.NewTable("CustomTreeView_LoadFromStream", 0),
		/*37*/ imports.NewTable("CustomTreeView_LockSelectionChangeEvent", 0),
		/*38*/ imports.NewTable("CustomTreeView_MakeSelectionVisible", 0),
		/*39*/ imports.NewTable("CustomTreeView_MoveCollapse", 0),
		/*40*/ imports.NewTable("CustomTreeView_MoveEnd", 0),
		/*41*/ imports.NewTable("CustomTreeView_MoveExpand", 0),
		/*42*/ imports.NewTable("CustomTreeView_MoveHome", 0),
		/*43*/ imports.NewTable("CustomTreeView_MoveLeft", 0),
		/*44*/ imports.NewTable("CustomTreeView_MovePageDown", 0),
		/*45*/ imports.NewTable("CustomTreeView_MovePageUp", 0),
		/*46*/ imports.NewTable("CustomTreeView_MoveRight", 0),
		/*47*/ imports.NewTable("CustomTreeView_MoveToNextNode", 0),
		/*48*/ imports.NewTable("CustomTreeView_MoveToPrevNode", 0),
		/*49*/ imports.NewTable("CustomTreeView_MultiSelectStyle", 0),
		/*50*/ imports.NewTable("CustomTreeView_Options", 0),
		/*51*/ imports.NewTable("CustomTreeView_SaveToFile", 0),
		/*52*/ imports.NewTable("CustomTreeView_SaveToStream", 0),
		/*53*/ imports.NewTable("CustomTreeView_ScrollBars", 0),
		/*54*/ imports.NewTable("CustomTreeView_Select", 0),
		/*55*/ imports.NewTable("CustomTreeView_Select1", 0),
		/*56*/ imports.NewTable("CustomTreeView_Selected", 0),
		/*57*/ imports.NewTable("CustomTreeView_SelectionColor", 0),
		/*58*/ imports.NewTable("CustomTreeView_SelectionCount", 0),
		/*59*/ imports.NewTable("CustomTreeView_SelectionFontColor", 0),
		/*60*/ imports.NewTable("CustomTreeView_SelectionFontColorUsed", 0),
		/*61*/ imports.NewTable("CustomTreeView_SelectionVisible", 0),
		/*62*/ imports.NewTable("CustomTreeView_Selections", 0),
		/*63*/ imports.NewTable("CustomTreeView_SeparatorColor", 0),
		/*64*/ imports.NewTable("CustomTreeView_SetInsertMark", 0),
		/*65*/ imports.NewTable("CustomTreeView_SetInsertMarkAt", 0),
		/*66*/ imports.NewTable("CustomTreeView_StateImages", 0),
		/*67*/ imports.NewTable("CustomTreeView_StateImagesWidth", 0),
		/*68*/ imports.NewTable("CustomTreeView_StoreCurrentSelection", 0),
		/*69*/ imports.NewTable("CustomTreeView_TopItem", 0),
		/*70*/ imports.NewTable("CustomTreeView_TreeLineColor", 0),
		/*71*/ imports.NewTable("CustomTreeView_TreeLinePenStyle", 0),
		/*72*/ imports.NewTable("CustomTreeView_UnlockSelectionChangeEvent", 0),
		/*73*/ imports.NewTable("CustomTreeView_WriteDebugReport", 0),
	}
)

func customTreeViewImportAPI() *imports.Imports {
	if customTreeViewImport == nil {
		customTreeViewImport = NewDefaultImports()
		customTreeViewImport.SetImportTable(customTreeViewImportTables)
		customTreeViewImportTables = nil
	}
	return customTreeViewImport
}
