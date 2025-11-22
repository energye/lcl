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

type callback struct {
	name string
	cb   func(getVal func(i int) uintptr)
}

func getPtr(val uintptr) base.UnsafePointer {
	return base.UnsafePointer(val)
}

func makeTAcceptDateEvent(cb TAcceptDateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TAcceptDateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender : TObject; var ADate : TDateTime; var AcceptDate: Boolean);
			sender := AsObject(getVal(0))
			date := (*types.TDateTime)(getPtr(getVal(1)))
			acceptDate := (*bool)(getPtr(getVal(2)))
			cb(sender, date, acceptDate)
		},
	}
}

func makeTAcceptFileNameEvent(cb TAcceptFileNameEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TAcceptFileNameEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender : TObject; Var Value : String);
			sender := AsObject(getVal(0))
			value := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &value)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(value)
		},
	}
}

func makeTAcceptTimeEvent(cb TAcceptTimeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TAcceptTimeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender : TObject; var ATime : TDateTime; var AcceptTime: Boolean);
			sender := AsObject(getVal(0))
			time := (*types.TDateTime)(getPtr(getVal(1)))
			acceptTime := (*bool)(getPtr(getVal(2)))
			cb(sender, time, acceptTime)
		},
	}
}

func makeTActionEvent(cb TActionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TActionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (AAction: TBasicAction; var Handled: Boolean);
			action := AsBasicAction(getVal(0))
			handled := (*bool)(getPtr(getVal(1)))
			cb(action, handled)
		},
	}
}

func makeTAddHeaderPopupItemEvent(cb TAddHeaderPopupItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TAddHeaderPopupItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TBaseVirtualTree; const Column: TColumnIndex; var Cmd: TAddPopupItemType);
			sender := AsBaseVirtualTree(getVal(0))
			column := int32(getVal(1))
			cmd := (*types.TAddPopupItemType)(getPtr(getVal(2)))
			cb(sender, column, cmd)
		},
	}
}

func makeTAlignInsertBeforeEvent(cb TAlignInsertBeforeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TAlignInsertBeforeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function (Sender: TWinControl; Control1, Control2: TControl): Boolean;
			sender := AsWinControl(getVal(0))
			control1 := AsControl(getVal(1))
			control2 := AsControl(getVal(2))
			ret := cb(sender, control1, control2)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTAlignPositionEvent(cb TAlignPositionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TAlignPositionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 8 : procedure (Sender: TWinControl; Control: TControl; var NewLeft, NewTop, NewWidth, NewHeight: Integer; var AlignRect: TRect; AlignInfo: TAlignInfo);
			sender := AsWinControl(getVal(0))
			control := AsControl(getVal(1))
			newLeft := (*int32)(getPtr(getVal(2)))
			newTop := (*int32)(getPtr(getVal(3)))
			newWidth := (*int32)(getPtr(getVal(4)))
			newHeight := (*int32)(getPtr(getVal(5)))
			alignRect := (*types.TRect)(getPtr(getVal(6)))
			alignInfoPtr := (*tAlignInfo)(getPtr(getVal(7)))
			alignInfo := alignInfoPtr.ToGo()
			cb(sender, control, newLeft, newTop, newWidth, newHeight, alignRect, alignInfo)
		},
	}
}

func makeTBandDragEvent(cb TBandDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TBandDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender: TObject; Control: TControl; var Drag: Boolean);
			sender := AsObject(getVal(0))
			control := AsControl(getVal(1))
			drag := (*bool)(getPtr(getVal(2)))
			cb(sender, control, drag)
		},
	}
}

func makeTBandInfoEvent(cb TBandInfoEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TBandInfoEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure (Sender: TObject; Control: TControl; var Insets: TRect; var PreferredSize, RowCount: Integer);
			sender := AsObject(getVal(0))
			control := AsControl(getVal(1))
			insets := (*types.TRect)(getPtr(getVal(2)))
			preferredSize := (*int32)(getPtr(getVal(3)))
			rowCount := (*int32)(getPtr(getVal(4)))
			cb(sender, control, insets, preferredSize, rowCount)
		},
	}
}

func makeTBandMoveEvent(cb TBandMoveEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TBandMoveEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender: TObject; Control: TControl; var ARect: TRect);
			sender := AsObject(getVal(0))
			control := AsControl(getVal(1))
			rect := (*types.TRect)(getPtr(getVal(2)))
			cb(sender, control, rect)
		},
	}
}

func makeTBandPaintEvent(cb TBandPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TBandPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure (Sender: TObject; Control: TControl; ACanvas: TCanvas; var ARect: TRect; var Options: TBandPaintOptions);
			sender := AsObject(getVal(0))
			control := AsControl(getVal(1))
			canvas := AsCanvas(getVal(2))
			rect := (*types.TRect)(getPtr(getVal(3)))
			options := (*types.TBandPaintOptions)(getPtr(getVal(4)))
			cb(sender, control, canvas, rect, options)
		},
	}
}

func makeTCanOffsetEvent(cb TCanOffsetEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCanOffsetEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; var NewOffset: Integer; var Accept: Boolean);
			sender := AsObject(getVal(0))
			newOffset := (*int32)(getPtr(getVal(1)))
			accept := (*bool)(getPtr(getVal(2)))
			cb(sender, newOffset, accept)
		},
	}
}

func makeTCanResizeEvent(cb TCanResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCanResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; var NewSize: Integer; var Accept: Boolean);
			sender := AsObject(getVal(0))
			newSize := (*int32)(getPtr(getVal(1)))
			accept := (*bool)(getPtr(getVal(2)))
			cb(sender, newSize, accept)
		},
	}
}

func makeTCellProcessEvent(cb TCellProcessEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCellProcessEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; aCol, aRow: Integer; processType: TCellProcessType; var aValue: string);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			processType := types.TCellProcessType(getVal(3))
			value := api.GoStr(*(*uintptr)(getPtr(getVal(4))))
			cb(sender, col, row, processType, &value)
			*(*uintptr)(getPtr(getVal(4))) = api.PasStr(value)
		},
	}
}

func makeTCheckGroupClicked(cb TCheckGroupClicked) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCheckGroupClicked",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Index: integer);
			sender := AsObject(getVal(0))
			index := int32(getVal(1))
			cb(sender, index)
		},
	}
}

func makeTCheckItemChange(cb TCheckItemChange) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCheckItemChange",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; AIndex: Integer);
			sender := AsObject(getVal(0))
			index := int32(getVal(1))
			cb(sender, index)
		},
	}
}

func makeTCheckItemEvent(cb TCheckItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCheckItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : function (Item: TObject): Boolean;
			item := AsObject(getVal(0))
			ret := cb(item)
			*(*bool)(getPtr(getVal(1))) = ret
		},
	}
}

func makeTClipboardRequestEvent(cb TClipboardRequestEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TClipboardRequestEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(const RequestedFormatID: TClipboardFormat; Data: TStream);
			requestedFormatID := types.TClipboardFormat(getVal(0))
			data := AsStream(getVal(1))
			cb(requestedFormatID, data)
		},
	}
}

func makeTCloseEvent(cb TCloseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCloseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var CloseAction: TCloseAction);
			sender := AsObject(getVal(0))
			closeAction := (*types.TCloseAction)(getPtr(getVal(1)))
			cb(sender, closeAction)
		},
	}
}

func makeTCloseQueryEvent(cb TCloseQueryEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCloseQueryEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender : TObject; var CanClose: Boolean);
			sender := AsObject(getVal(0))
			canClose := (*bool)(getPtr(getVal(1)))
			cb(sender, canClose)
		},
	}
}

func makeTColumnChangeEvent(cb TColumnChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TColumnChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(const Sender: TBaseVirtualTree; const Column: TColumnIndex; Visible: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			column := int32(getVal(1))
			visible := api.GoBool(getVal(2))
			cb(sender, column, visible)
		},
	}
}

func makeTConstrainedResizeEvent(cb TConstrainedResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TConstrainedResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; var MinWidth, MinHeight, MaxWidth, MaxHeight: TConstraintSize);
			sender := AsObject(getVal(0))
			minWidth := (*types.TConstraintSize)(getPtr(getVal(1)))
			minHeight := (*types.TConstraintSize)(getPtr(getVal(2)))
			maxWidth := (*types.TConstraintSize)(getPtr(getVal(3)))
			maxHeight := (*types.TConstraintSize)(getPtr(getVal(4)))
			cb(sender, minWidth, minHeight, maxWidth, maxHeight)
		},
	}
}

func makeTContextPopupEvent(cb TContextPopupEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TContextPopupEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; MousePos: TPoint; var Handled: Boolean);
			sender := AsObject(getVal(0))
			mousePos := *(*types.TPoint)(getPtr(getVal(1)))
			handled := (*bool)(getPtr(getVal(2)))
			cb(sender, mousePos, handled)
		},
	}
}

func makeTControlShowHintEvent(cb TControlShowHintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TControlShowHintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; HintInfo: PHintInfo);
			sender := AsObject(getVal(0))
			hintInfoPtr := (*tHintInfo)(getPtr(getVal(1)))
			hintInfo := hintInfoPtr.ToGo()
			cb(sender, hintInfo)
		},
	}
}

func makeTCustomDateEvent(cb TCustomDateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCustomDateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender : TObject; var ADate : string);
			sender := AsObject(getVal(0))
			date := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &date)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(date)
		},
	}
}

func makeTCustomGraphicNotify(cb TCustomGraphicNotify) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCustomGraphicNotify",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTCustomHCCreateSectionClassEvent(cb TCustomHCCreateSectionClassEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCustomHCCreateSectionClassEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TCustomHeaderControl; var SectionClass: THeaderSectionClass);
			sender := AsCustomHeaderControl(getVal(0))
			sectionClass := (*types.THeaderSectionClass)(getPtr(getVal(1)))
			cb(sender, sectionClass)
		},
	}
}

func makeTCustomImageListGetWidthForPPI(cb TCustomImageListGetWidthForPPI) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCustomImageListGetWidthForPPI",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TCustomImageList; AImageWidth, APPI: Integer; var AResultWidth: Integer);
			sender := AsCustomImageList(getVal(0))
			imageWidth := int32(getVal(1))
			pPI := int32(getVal(2))
			resultWidth := (*int32)(getPtr(getVal(3)))
			cb(sender, imageWidth, pPI, resultWidth)
		},
	}
}

func makeTCustomSectionNotifyEvent(cb TCustomSectionNotifyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCustomSectionNotifyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(HeaderControl: TCustomHeaderControl; Section: THeaderSection);
			headerControl := AsCustomHeaderControl(getVal(0))
			section := AsHeaderSection(getVal(1))
			cb(headerControl, section)
		},
	}
}

func makeTCustomSectionTrackEvent(cb TCustomSectionTrackEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCustomSectionTrackEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(HeaderControl: TCustomHeaderControl; Section: THeaderSection; Width: Integer; State: TSectionTrackState);
			headerControl := AsCustomHeaderControl(getVal(0))
			section := AsHeaderSection(getVal(1))
			width := int32(getVal(2))
			state := types.TSectionTrackState(getVal(3))
			cb(headerControl, section, width, state)
		},
	}
}

func makeTCustomTimeEvent(cb TCustomTimeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TCustomTimeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender : TObject; var ATime : TDateTime);
			sender := AsObject(getVal(0))
			time := (*types.TDateTime)(getPtr(getVal(1)))
			cb(sender, time)
		},
	}
}

func makeTDateRangeCheckEvent(cb TDateRangeCheckEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDateRangeCheckEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender : TObject; var ADate : TDateTime);
			sender := AsObject(getVal(0))
			date := (*types.TDateTime)(getPtr(getVal(1)))
			cb(sender, date)
		},
	}
}

func makeTDesignerGetShiftStateEvent(cb TDesignerGetShiftStateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerGetShiftStateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): TShiftState;
			ret := cb()
			*(*types.TShiftState)(getPtr(getVal(0))) = ret
		},
	}
}

func makeTDesignerIsDesignMsgEvent(cb TDesignerIsDesignMsgEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerIsDesignMsgEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(Sender: TControl; var Message: TLMessage): boolean;
			sender := AsControl(getVal(0))
			message := (*types.TLMessage)(getPtr(getVal(1)))
			ret := cb(sender, message)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTDesignerModifiedEvent(cb TDesignerModifiedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerModifiedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTDesignerNotificationEvent(cb TDesignerNotificationEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerNotificationEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(AComponent: TComponent; Operation: TOperation);
			component := AsComponent(getVal(0))
			operation := types.TOperation(getVal(1))
			cb(component, operation)
		},
	}
}

func makeTDesignerPaintGridEvent(cb TDesignerPaintGridEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerPaintGridEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : procedure();
			cb()
		},
	}
}

func makeTDesignerPrepareFreeDesignerEvent(cb TDesignerPrepareFreeDesignerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerPrepareFreeDesignerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(AFreeComponent: boolean);
			freeComponent := api.GoBool(getVal(0))
			cb(freeComponent)
		},
	}
}

func makeTDesignerSelectOnlyThisComponentEvent(cb TDesignerSelectOnlyThisComponentEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerSelectOnlyThisComponentEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(AComponent: TComponent);
			component := AsComponent(getVal(0))
			cb(component)
		},
	}
}

func makeTDesignerUTF8KeyPressEvent(cb TDesignerUTF8KeyPressEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerUTF8KeyPressEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var UTF8Key: TUTF8Char);
			uTF8Key := api.GoStr(*(*uintptr)(getPtr(getVal(0))))
			cb(&uTF8Key)
			*(*uintptr)(getPtr(getVal(0))) = api.PasStr(uTF8Key)
		},
	}
}

func makeTDesignerUniqueNameEvent(cb TDesignerUniqueNameEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerUniqueNameEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : function(const BaseName: string): string;
			baseName := api.GoStr(getVal(0))
			ret := cb(baseName)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(ret)
		},
	}
}

func makeTDesignerValidateRenameEvent(cb TDesignerValidateRenameEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDesignerValidateRenameEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(AComponent: TComponent; const CurName, NewName: string);
			component := AsComponent(getVal(0))
			curName := api.GoStr(getVal(1))
			newName := api.GoStr(getVal(2))
			cb(component, curName, newName)
		},
	}
}

func makeTDestroyResolutionHandleEvent(cb TDestroyResolutionHandleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDestroyResolutionHandleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TCustomImageList; AWidth: Integer; AReferenceHandle: TLCLHandle);
			sender := AsCustomImageList(getVal(0))
			width := int32(getVal(1))
			referenceHandle := types.TLCLHandle(getVal(2))
			cb(sender, width, referenceHandle)
		},
	}
}

func makeTDialogResultEvent(cb TDialogResultEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDialogResultEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Success: Boolean);
			sender := AsObject(getVal(0))
			success := api.GoBool(getVal(1))
			cb(sender, success)
		},
	}
}

func makeTDockDropEvent(cb TDockDropEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDockDropEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; Source: TDragDockObject; X, Y: Integer);
			sender := AsObject(getVal(0))
			source := AsDragDockObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, source, X, Y)
		},
	}
}

func makeTDockOverEvent(cb TDockOverEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDockOverEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; Source: TDragDockObject; X, Y: Integer; State: TDragState; var Accept: Boolean);
			sender := AsObject(getVal(0))
			source := AsDragDockObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			state := types.TDragState(getVal(4))
			accept := (*bool)(getPtr(getVal(5)))
			cb(sender, source, X, Y, state, accept)
		},
	}
}

func makeTDragDropEvent(cb TDragDropEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDragDropEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender, Source: TObject; X,Y: Integer);
			sender := AsObject(getVal(0))
			source := AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, source, X, Y)
		},
	}
}

func makeTDragOverEvent(cb TDragOverEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDragOverEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender, Source: TObject; X,Y: Integer; State: TDragState; var Accept: Boolean);
			sender := AsObject(getVal(0))
			source := AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			state := types.TDragState(getVal(4))
			accept := (*bool)(getPtr(getVal(5)))
			cb(sender, source, X, Y, state, accept)
		},
	}
}

func makeTDrawItemEvent(cb TDrawItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDrawItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Control: TWinControl; Index: Integer; ARect: TRect; State: TOwnerDrawState);
			control := AsWinControl(getVal(0))
			index := int32(getVal(1))
			rect := *(*types.TRect)(getPtr(getVal(2)))
			state := types.TOwnerDrawState(getVal(3))
			cb(control, index, rect, state)
		},
	}
}

func makeTDrawPanelEvent(cb TDrawPanelEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDrawPanelEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(StatusBar: TStatusBar; Panel: TStatusPanel; const Rect: TRect);
			statusBar := AsStatusBar(getVal(0))
			panel := AsStatusPanel(getVal(1))
			rect := *(*types.TRect)(getPtr(getVal(2)))
			cb(statusBar, panel, rect)
		},
	}
}

func makeTDropFilesEvent(cb TDropFilesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDropFilesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender: TObject; const FileNames: array of string);
			sender := AsObject(getVal(0))
			fileNamesPtr := getVal(1)
			fileNamesLen := int32(getVal(2))
			fileNames := NewStringArray(int(fileNamesLen), fileNamesPtr)
			cb(sender, fileNames)
		},
	}
}

func makeTEndDragEvent(cb TEndDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TEndDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender, Target: TObject; X,Y: Integer);
			sender := AsObject(getVal(0))
			target := AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, target, X, Y)
		},
	}
}

func makeTFPCanvasCombineColors(cb TFPCanvasCombineColors) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TFPCanvasCombineColors",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(const color1, color2: TFPColor): TFPColor;
			color1 := *(*TFPColor)(getPtr(getVal(0)))
			color2 := *(*TFPColor)(getPtr(getVal(1)))
			ret := cb(color1, color2)
			*(*TFPColor)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTFPImgProgressEvent(cb TFPImgProgressEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TFPImgProgressEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure (Sender: TObject; Stage: TFPImgProgressStage; PercentDone: Byte; RedrawNow: Boolean; const R: TRect; const Msg: AnsiString; var Continue : Boolean);
			sender := AsObject(getVal(0))
			stage := types.TFPImgProgressStage(getVal(1))
			percentDone := byte(getVal(2))
			redrawNow := api.GoBool(getVal(3))
			R := *(*types.TRect)(getPtr(getVal(4)))
			msg := api.GoStr(getVal(5))
			continue_ := (*bool)(getPtr(getVal(6)))
			cb(sender, stage, percentDone, redrawNow, R, msg, continue_)
		},
	}
}

func makeTFilterItemEvent(cb TFilterItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TFilterItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function (ItemData: Pointer; out Done: Boolean): Boolean;
			itemData := uintptr(getVal(0))
			done := (*bool)(getPtr(getVal(1)))
			ret := cb(itemData, done)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTFilterItemExEvent(cb TFilterItemExEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TFilterItemExEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function (const ACaption: string; ItemData: Pointer; out Done: Boolean): Boolean;
			caption := api.GoStr(getVal(0))
			itemData := uintptr(getVal(1))
			done := (*bool)(getPtr(getVal(2)))
			ret := cb(caption, itemData, done)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTFilterNodeEvent(cb TFilterNodeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TFilterNodeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function (ItemNode: TTreeNode; out Done: Boolean): Boolean;
			itemNode := AsTreeNode(getVal(0))
			done := (*bool)(getPtr(getVal(1)))
			ret := cb(itemNode, done)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTGetCellHintEvent(cb TGetCellHintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetCellHintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; ACol, ARow: Integer; var HintText: String);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			hintText := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
			cb(sender, col, row, &hintText)
			*(*uintptr)(getPtr(getVal(3))) = api.PasStr(hintText)
		},
	}
}

func makeTGetCheckboxStateEvent(cb TGetCheckboxStateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetCheckboxStateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; ACol, ARow: Integer; var Value: TCheckboxState);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			value := (*types.TCheckBoxState)(getPtr(getVal(3)))
			cb(sender, col, row, value)
		},
	}
}

func makeTGetColorsEvent(cb TGetColorsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetColorsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TCustomColorBox; Items: TStrings);
			sender := AsCustomColorBox(getVal(0))
			items := AsStrings(getVal(1))
			cb(sender, items)
		},
	}
}

func makeTGetDockCaptionEvent(cb TGetDockCaptionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetDockCaptionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; AControl: TControl; var ACaption: string);
			sender := AsObject(getVal(0))
			control := AsControl(getVal(1))
			caption := api.GoStr(*(*uintptr)(getPtr(getVal(2))))
			cb(sender, control, &caption)
			*(*uintptr)(getPtr(getVal(2))) = api.PasStr(caption)
		},
	}
}

func makeTGetEditEvent(cb TGetEditEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetEditEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; ACol, ARow: Integer; var Value: string);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			value := api.GoStr(*(*uintptr)(getPtr(getVal(3))))
			cb(sender, col, row, &value)
			*(*uintptr)(getPtr(getVal(3))) = api.PasStr(value)
		},
	}
}

func makeTGetHandleEvent(cb TGetHandleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetHandleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var Handle: HWND);
			handle := (*types.HWND)(getPtr(getVal(0)))
			cb(handle)
		},
	}
}

func makeTGetPickListEvent(cb TGetPickListEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetPickListEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const KeyName: string; Values: TStrings);
			sender := AsObject(getVal(0))
			keyName := api.GoStr(getVal(1))
			values := AsStrings(getVal(2))
			cb(sender, keyName, values)
		},
	}
}

func makeTGetSiteInfoEvent(cb TGetSiteInfoEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TGetSiteInfoEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; DockClient: TControl; var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean);
			sender := AsObject(getVal(0))
			dockClient := AsControl(getVal(1))
			influenceRect := (*types.TRect)(getPtr(getVal(2)))
			mousePos := *(*types.TPoint)(getPtr(getVal(3)))
			canDock := (*bool)(getPtr(getVal(4)))
			cb(sender, dockClient, influenceRect, mousePos, canDock)
		},
	}
}

func makeTHdrEvent(cb THdrEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "THdrEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; IsColumn: Boolean; Index: Integer);
			sender := AsObject(getVal(0))
			isColumn := api.GoBool(getVal(1))
			index := int32(getVal(2))
			cb(sender, isColumn, index)
		},
	}
}

func makeTHeaderSizingEvent(cb THeaderSizingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "THeaderSizingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const IsColumn: boolean; const aIndex, aSize: Integer);
			sender := AsObject(getVal(0))
			isColumn := api.GoBool(getVal(1))
			index := int32(getVal(2))
			size := int32(getVal(3))
			cb(sender, isColumn, index, size)
		},
	}
}

func makeTHelpEvent(cb THelpEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "THelpEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(Command: Word; Data: PtrInt; var CallHelp: Boolean): Boolean;
			command := uint16(getVal(0))
			data := uintptr(getVal(1))
			callHelp := (*bool)(getPtr(getVal(2)))
			ret := cb(command, data, callHelp)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTHintEvent(cb THintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "THintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (var HintStr: string; var CanShow: Boolean);
			hintStr := api.GoStr(*(*uintptr)(getPtr(getVal(0))))
			canShow := (*bool)(getPtr(getVal(1)))
			cb(&hintStr, canShow)
			*(*uintptr)(getPtr(getVal(0))) = api.PasStr(hintStr)
		},
	}
}

func makeTIdleEvent(cb TIdleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TIdleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender: TObject; var Done: Boolean);
			sender := AsObject(getVal(0))
			done := (*bool)(getPtr(getVal(1)))
			cb(sender, done)
		},
	}
}

func makeTImageIndexEvent(cb TImageIndexEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TImageIndexEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function (Str: String; Data: TObject; var AIsEnabled: Boolean): Integer;
			str := api.GoStr(getVal(0))
			data := AsObject(getVal(1))
			isEnabled := (*bool)(getPtr(getVal(2)))
			ret := cb(str, data, isEnabled)
			*(*int32)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTImagePaintBackgroundEvent(cb TImagePaintBackgroundEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TImagePaintBackgroundEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (ASender: TObject; ACanvas: TCanvas; ARect: TRect);
			sender := AsObject(getVal(0))
			canvas := AsCanvas(getVal(1))
			rect := *(*types.TRect)(getPtr(getVal(2)))
			cb(sender, canvas, rect)
		},
	}
}

func makeTKeyEvent(cb TKeyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TKeyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; var Key: Word; Shift: TShiftState);
			sender := AsObject(getVal(0))
			key := (*uint16)(getPtr(getVal(1)))
			shift := types.TShiftState(getVal(2))
			cb(sender, key, shift)
		},
	}
}

func makeTKeyPressEvent(cb TKeyPressEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TKeyPressEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var Key: char);
			sender := AsObject(getVal(0))
			key := (*uint16)(getPtr(getVal(1)))
			cb(sender, key)
		},
	}
}

func makeTLBGetColorsEvent(cb TLBGetColorsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLBGetColorsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TCustomColorListBox; Items: TStrings);
			sender := AsCustomColorListBox(getVal(0))
			items := AsStrings(getVal(1))
			cb(sender, items)
		},
	}
}

func makeTLVAdvancedCustomDrawEvent(cb TLVAdvancedCustomDrawEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVAdvancedCustomDrawEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TCustomListView; const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
			sender := AsCustomListView(getVal(0))
			rect := *(*types.TRect)(getPtr(getVal(1)))
			stage := types.TCustomDrawStage(getVal(2))
			defaultDraw := (*bool)(getPtr(getVal(3)))
			cb(sender, rect, stage, defaultDraw)
		},
	}
}

func makeTLVAdvancedCustomDrawItemEvent(cb TLVAdvancedCustomDrawItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVAdvancedCustomDrawItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TCustomListView; Item: TListItem; State: TCustomDrawState; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
			sender := AsCustomListView(getVal(0))
			item := AsListItem(getVal(1))
			state := types.TCustomDrawState(getVal(2))
			stage := types.TCustomDrawStage(getVal(3))
			defaultDraw := (*bool)(getPtr(getVal(4)))
			cb(sender, item, state, stage, defaultDraw)
		},
	}
}

func makeTLVAdvancedCustomDrawSubItemEvent(cb TLVAdvancedCustomDrawSubItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVAdvancedCustomDrawSubItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TCustomListView; Item: TListItem; SubItem: Integer; State: TCustomDrawState; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
			sender := AsCustomListView(getVal(0))
			item := AsListItem(getVal(1))
			subItem := int32(getVal(2))
			state := types.TCustomDrawState(getVal(3))
			stage := types.TCustomDrawStage(getVal(4))
			defaultDraw := (*bool)(getPtr(getVal(5)))
			cb(sender, item, subItem, state, stage, defaultDraw)
		},
	}
}

func makeTLVChangeEvent(cb TLVChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Item: TListItem; Change: TItemChange);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			change := types.TItemChange(getVal(2))
			cb(sender, item, change)
		},
	}
}

func makeTLVChangingEvent(cb TLVChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; Item: TListItem; Change: TItemChange; var AllowChange: Boolean);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			change := types.TItemChange(getVal(2))
			allowChange := (*bool)(getPtr(getVal(3)))
			cb(sender, item, change, allowChange)
		},
	}
}

func makeTLVCheckedItemEvent(cb TLVCheckedItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVCheckedItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender: TObject; Item: TListItem);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			cb(sender, item)
		},
	}
}

func makeTLVColumnClickEvent(cb TLVColumnClickEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVColumnClickEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Column: TListColumn);
			sender := AsObject(getVal(0))
			column := AsListColumn(getVal(1))
			cb(sender, column)
		},
	}
}

func makeTLVCompareEvent(cb TLVCompareEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVCompareEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; Item1, Item2: TListItem; Data: Integer; var Compare: Integer);
			sender := AsObject(getVal(0))
			item1 := AsListItem(getVal(1))
			item2 := AsListItem(getVal(2))
			data := int32(getVal(3))
			compare := (*int32)(getPtr(getVal(4)))
			cb(sender, item1, item2, data, compare)
		},
	}
}

func makeTLVCreateItemClassEvent(cb TLVCreateItemClassEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVCreateItemClassEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TCustomListView; var ItemClass: TListItemClass);
			sender := AsCustomListView(getVal(0))
			itemClass := (*types.TListItemClass)(getPtr(getVal(1)))
			cb(sender, itemClass)
		},
	}
}

func makeTLVCustomDrawEvent(cb TLVCustomDrawEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVCustomDrawEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TCustomListView; const ARect: TRect; var DefaultDraw: Boolean);
			sender := AsCustomListView(getVal(0))
			rect := *(*types.TRect)(getPtr(getVal(1)))
			defaultDraw := (*bool)(getPtr(getVal(2)))
			cb(sender, rect, defaultDraw)
		},
	}
}

func makeTLVCustomDrawItemEvent(cb TLVCustomDrawItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVCustomDrawItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TCustomListView; Item: TListItem; State: TCustomDrawState; var DefaultDraw: Boolean);
			sender := AsCustomListView(getVal(0))
			item := AsListItem(getVal(1))
			state := types.TCustomDrawState(getVal(2))
			defaultDraw := (*bool)(getPtr(getVal(3)))
			cb(sender, item, state, defaultDraw)
		},
	}
}

func makeTLVCustomDrawSubItemEvent(cb TLVCustomDrawSubItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVCustomDrawSubItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TCustomListView; Item: TListItem; SubItem: Integer; State: TCustomDrawState; var DefaultDraw: Boolean);
			sender := AsCustomListView(getVal(0))
			item := AsListItem(getVal(1))
			subItem := int32(getVal(2))
			state := types.TCustomDrawState(getVal(3))
			defaultDraw := (*bool)(getPtr(getVal(4)))
			cb(sender, item, subItem, state, defaultDraw)
		},
	}
}

func makeTLVDataEvent(cb TLVDataEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVDataEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Item: TListItem);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			cb(sender, item)
		},
	}
}

func makeTLVDataFindEvent(cb TLVDataFindEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVDataFindEvent",
		cb: func(getVal func(i int) uintptr) {
			// 9 : procedure(Sender: TObject; AFind: TItemFind; const AFindString: string; const AFindPosition: TPoint; AFindData: Pointer; AStartIndex: Integer; ADirection: TSearchDirection; AWrap: Boolean; var AIndex: Integer);
			sender := AsObject(getVal(0))
			find := types.TItemFind(getVal(1))
			findString := api.GoStr(getVal(2))
			findPosition := *(*types.TPoint)(getPtr(getVal(3)))
			findData := uintptr(getVal(4))
			startIndex := int32(getVal(5))
			direction := types.TSearchDirection(getVal(6))
			wrap := api.GoBool(getVal(7))
			index := (*int32)(getPtr(getVal(8)))
			cb(sender, find, findString, findPosition, findData, startIndex, direction, wrap, index)
		},
	}
}

func makeTLVDataHintEvent(cb TLVDataHintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVDataHintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; StartIndex, EndIndex: Integer);
			sender := AsObject(getVal(0))
			startIndex := int32(getVal(1))
			endIndex := int32(getVal(2))
			cb(sender, startIndex, endIndex)
		},
	}
}

func makeTLVDataStateChangeEvent(cb TLVDataStateChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVDataStateChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; StartIndex, EndIndex: Integer; OldState, NewState: TListItemStates);
			sender := AsObject(getVal(0))
			startIndex := int32(getVal(1))
			endIndex := int32(getVal(2))
			oldState := types.TListItemStates(getVal(3))
			newState := types.TListItemStates(getVal(4))
			cb(sender, startIndex, endIndex, oldState, newState)
		},
	}
}

func makeTLVDeletedEvent(cb TLVDeletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVDeletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Item: TListItem);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			cb(sender, item)
		},
	}
}

func makeTLVDrawItemEvent(cb TLVDrawItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVDrawItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TCustomListView; AItem: TListItem; ARect: TRect; AState: TOwnerDrawState);
			sender := AsCustomListView(getVal(0))
			item := AsListItem(getVal(1))
			rect := *(*types.TRect)(getPtr(getVal(2)))
			state := types.TOwnerDrawState(getVal(3))
			cb(sender, item, rect, state)
		},
	}
}

func makeTLVEditedEvent(cb TLVEditedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVEditedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Item: TListItem; var AValue: string);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			value := api.GoStr(*(*uintptr)(getPtr(getVal(2))))
			cb(sender, item, &value)
			*(*uintptr)(getPtr(getVal(2))) = api.PasStr(value)
		},
	}
}

func makeTLVEditingEvent(cb TLVEditingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVEditingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Item: TListItem; var AllowEdit: Boolean);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			allowEdit := (*bool)(getPtr(getVal(2)))
			cb(sender, item, allowEdit)
		},
	}
}

func makeTLVInsertEvent(cb TLVInsertEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVInsertEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Item: TListItem);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			cb(sender, item)
		},
	}
}

func makeTLVSelectItemEvent(cb TLVSelectItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLVSelectItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Item: TListItem; Selected: Boolean);
			sender := AsObject(getVal(0))
			item := AsListItem(getVal(1))
			selected := api.GoBool(getVal(2))
			cb(sender, item, selected)
		},
	}
}

func makeTLinkActionEvent(cb TLinkActionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLinkActionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure (Sender: TObject; ALinkAction: TLinkAction; const info: TLinkMouseInfo; LinkStart, LinkLen: Integer);
			sender := AsObject(getVal(0))
			linkAction := types.TLinkAction(getVal(1))
			infoPtr := (*tLinkMouseInfo)(getPtr(getVal(2)))
			info := infoPtr.ToGo()
			linkStart := int32(getVal(3))
			linkLen := int32(getVal(4))
			cb(sender, linkAction, info, linkStart, linkLen)
		},
	}
}

func makeTListCompareEvent(cb TListCompareEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TListCompareEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(AList: TListControlItems; AItem1, AItem2: TListControlItem): Integer;
			list := AsListControlItems(getVal(0))
			item1 := AsListControlItem(getVal(1))
			item2 := AsListControlItem(getVal(2))
			ret := cb(list, item1, item2)
			*(*int32)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTMeasureItemEvent(cb TMeasureItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMeasureItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Control: TWinControl; Index: Integer; var AHeight: Integer);
			control := AsWinControl(getVal(0))
			index := int32(getVal(1))
			height := (*int32)(getPtr(getVal(2)))
			cb(control, index, height)
		},
	}
}

func makeTMenuChangeEvent(cb TMenuChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMenuChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender: TObject; Source: TMenuItem; Rebuild: Boolean);
			sender := AsObject(getVal(0))
			source := AsMenuItem(getVal(1))
			rebuild := api.GoBool(getVal(2))
			cb(sender, source, rebuild)
		},
	}
}

func makeTMenuDrawItemEvent(cb TMenuDrawItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMenuDrawItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; ACanvas: TCanvas; ARect: TRect; AState: TOwnerDrawState);
			sender := AsObject(getVal(0))
			canvas := AsCanvas(getVal(1))
			rect := *(*types.TRect)(getPtr(getVal(2)))
			state := types.TOwnerDrawState(getVal(3))
			cb(sender, canvas, rect, state)
		},
	}
}

func makeTMenuMeasureItemEvent(cb TMenuMeasureItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMenuMeasureItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; ACanvas: TCanvas; var AWidth, AHeight: Integer);
			sender := AsObject(getVal(0))
			canvas := AsCanvas(getVal(1))
			width := (*int32)(getPtr(getVal(2)))
			height := (*int32)(getPtr(getVal(3)))
			cb(sender, canvas, width, height)
		},
	}
}

func makeTModalDialogFinished(cb TModalDialogFinished) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TModalDialogFinished",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender: TObject; AResult: Integer);
			sender := AsObject(getVal(0))
			result := int32(getVal(1))
			cb(sender, result)
		},
	}
}

func makeTMouseEvent(cb TMouseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMouseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; Button: TMouseButton; Shift: TShiftState; X, Y: Integer);
			sender := AsObject(getVal(0))
			button := types.TMouseButton(getVal(1))
			shift := types.TShiftState(getVal(2))
			X := int32(getVal(3))
			Y := int32(getVal(4))
			cb(sender, button, shift, X, Y)
		},
	}
}

func makeTMouseMoveEvent(cb TMouseMoveEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMouseMoveEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; Shift: TShiftState; X, Y: Integer);
			sender := AsObject(getVal(0))
			shift := types.TShiftState(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, shift, X, Y)
		},
	}
}

func makeTMouseWheelEvent(cb TMouseWheelEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMouseWheelEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; Shift: TShiftState; WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean);
			sender := AsObject(getVal(0))
			shift := types.TShiftState(getVal(1))
			wheelDelta := int32(getVal(2))
			mousePos := *(*types.TPoint)(getPtr(getVal(3)))
			handled := (*bool)(getPtr(getVal(4)))
			cb(sender, shift, wheelDelta, mousePos, handled)
		},
	}
}

func makeTMouseWheelUpDownEvent(cb TMouseWheelUpDownEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TMouseWheelUpDownEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean);
			sender := AsObject(getVal(0))
			shift := types.TShiftState(getVal(1))
			mousePos := *(*types.TPoint)(getPtr(getVal(2)))
			handled := (*bool)(getPtr(getVal(3)))
			cb(sender, shift, mousePos, handled)
		},
	}
}

func makeTNotebookTabDragDropEvent(cb TNotebookTabDragDropEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TNotebookTabDragDropEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender, Source: TObject; OldIndex, NewIndex: Integer; CopyDrag: Boolean; var Done: Boolean);
			sender := AsObject(getVal(0))
			source := AsObject(getVal(1))
			oldIndex := int32(getVal(2))
			newIndex := int32(getVal(3))
			copyDrag := api.GoBool(getVal(4))
			done := (*bool)(getPtr(getVal(5)))
			cb(sender, source, oldIndex, newIndex, copyDrag, done)
		},
	}
}

func makeTNotebookTabDragOverEvent(cb TNotebookTabDragOverEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TNotebookTabDragOverEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender, Source: TObject; OldIndex, NewIndex: Integer; CopyDrag: Boolean; var Accept: Boolean);
			sender := AsObject(getVal(0))
			source := AsObject(getVal(1))
			oldIndex := int32(getVal(2))
			newIndex := int32(getVal(3))
			copyDrag := api.GoBool(getVal(4))
			accept := (*bool)(getPtr(getVal(5)))
			cb(sender, source, oldIndex, newIndex, copyDrag, accept)
		},
	}
}

func makeTNotifyEvent(cb TNotifyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TNotifyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(Sender: TObject);
			sender := AsObject(getVal(0))
			cb(sender)
		},
	}
}

func makeTOnCompareCells(cb TOnCompareCells) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCompareCells",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure (Sender: TObject; ACol, ARow, BCol,BRow: Integer; var Result: integer);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			bCol := int32(getVal(3))
			bRow := int32(getVal(4))
			result := (*int32)(getPtr(getVal(5)))
			cb(sender, col, row, bCol, bRow, result)
		},
	}
}

func makeTOnDrawCell(cb TOnDrawCell) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDrawCell",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; aCol, aRow: Integer; aRect: TRect; aState:TGridDrawState);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			rect := *(*types.TRect)(getPtr(getVal(3)))
			state := types.TGridDrawState(getVal(4))
			cb(sender, col, row, rect, state)
		},
	}
}

func makeTOnPrepareCanvasEvent(cb TOnPrepareCanvasEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrepareCanvasEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; aCol, aRow: Integer; aState: TGridDrawState);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			state := types.TGridDrawState(getVal(3))
			cb(sender, col, row, state)
		},
	}
}

func makeTOnSelectCellEvent(cb TOnSelectCellEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSelectCellEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; aCol, aRow: Integer; var CanSelect: Boolean);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			canSelect := (*bool)(getPtr(getVal(3)))
			cb(sender, col, row, canSelect)
		},
	}
}

func makeTOnSelectEvent(cb TOnSelectEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSelectEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aCol, aRow: Integer);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			cb(sender, col, row)
		},
	}
}

func makeTOnUserInputEvent(cb TOnUserInputEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnUserInputEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Msg: Cardinal);
			sender := AsObject(getVal(0))
			msg := uint32(getVal(1))
			cb(sender, msg)
		},
	}
}

func makeTOnValidateEvent(cb TOnValidateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnValidateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; ACol, ARow: Longint; const KeyName, KeyValue: string);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			keyName := api.GoStr(getVal(3))
			keyValue := api.GoStr(getVal(4))
			cb(sender, col, row, keyName, keyValue)
		},
	}
}

func makeTOpenGlCtrlMakeCurrentEvent(cb TOpenGlCtrlMakeCurrentEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOpenGlCtrlMakeCurrentEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var Allow: boolean);
			sender := AsObject(getVal(0))
			allow := (*bool)(getPtr(getVal(1)))
			cb(sender, allow)
		},
	}
}

func makeTPrintActionEvent(cb TPrintActionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TPrintActionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure (Sender: TObject; APrintAction: TPrintAction; PrintCanvas: TCanvas; CurrentPage: Integer; var AbortPrint: Boolean);
			sender := AsObject(getVal(0))
			printAction := types.TPrintAction(getVal(1))
			printCanvas := AsCanvas(getVal(2))
			currentPage := int32(getVal(3))
			abortPrint := (*bool)(getPtr(getVal(4)))
			cb(sender, printAction, printCanvas, currentPage, abortPrint)
		},
	}
}

func makeTProgressEvent(cb TProgressEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TProgressEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure (Sender: TObject; Stage: TFPImgProgressStage; PercentDone: Byte; RedrawNow: Boolean; const R: TRect; const Msg: AnsiString; var Continue : Boolean);
			sender := AsObject(getVal(0))
			stage := types.TFPImgProgressStage(getVal(1))
			percentDone := byte(getVal(2))
			redrawNow := api.GoBool(getVal(3))
			R := *(*types.TRect)(getPtr(getVal(4)))
			msg := api.GoStr(getVal(5))
			continue_ := (*bool)(getPtr(getVal(6)))
			cb(sender, stage, percentDone, redrawNow, R, msg, continue_)
		},
	}
}

func makeTQueryEndSessionEvent(cb TQueryEndSessionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TQueryEndSessionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure (var Cancel: Boolean);
			cancel := (*bool)(getPtr(getVal(0)))
			cb(cancel)
		},
	}
}

func makeTSBCreatePanelClassEvent(cb TSBCreatePanelClassEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TSBCreatePanelClassEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TStatusBar; var PanelClass: TStatusPanelClass);
			sender := AsStatusBar(getVal(0))
			panelClass := (*types.TStatusPanelClass)(getPtr(getVal(1)))
			cb(sender, panelClass)
		},
	}
}

func makeTScrollEvent(cb TScrollEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TScrollEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; ScrollCode: TScrollCode; var ScrollPos: Integer);
			sender := AsObject(getVal(0))
			scrollCode := types.TScrollCode(getVal(1))
			scrollPos := (*int32)(getPtr(getVal(2)))
			cb(sender, scrollCode, scrollPos)
		},
	}
}

func makeTSectionDragEvent(cb TSectionDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TSectionDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; FromSection, ToSection: THeaderSection; var AllowDrag: Boolean);
			sender := AsObject(getVal(0))
			fromSection := AsHeaderSection(getVal(1))
			toSection := AsHeaderSection(getVal(2))
			allowDrag := (*bool)(getPtr(getVal(3)))
			cb(sender, fromSection, toSection, allowDrag)
		},
	}
}

func makeTSelectEditorEvent(cb TSelectEditorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TSelectEditorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; aCol, aRow: Integer; var Editor: TWinControl);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			var editor IWinControl
			editor = AsWinControl(*(*uintptr)(getPtr(getVal(3))))
			cb(sender, col, row, &editor)
			if editor != nil {
				*(*uintptr)(getPtr(getVal(3))) = editor.Instance()
			}
		},
	}
}

func makeTSelectionChangeEvent(cb TSelectionChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TSelectionChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; User: boolean);
			sender := AsObject(getVal(0))
			user := api.GoBool(getVal(1))
			cb(sender, user)
		},
	}
}

func makeTSetCheckboxStateEvent(cb TSetCheckboxStateEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TSetCheckboxStateEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; ACol, ARow: Integer; const Value: TCheckboxState);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			value := types.TCheckBoxState(getVal(3))
			cb(sender, col, row, value)
		},
	}
}

func makeTSetEditEvent(cb TSetEditEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TSetEditEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; ACol, ARow: Integer; const Value: string);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			value := api.GoStr(getVal(3))
			cb(sender, col, row, value)
		},
	}
}

func makeTShapePointsEvent(cb TShapePointsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TShapePointsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender: TObject; var Points: TPointArray; var Winding: Boolean);
			sender := AsObject(getVal(0))
			var points IPointArrayWrap
			points = AsPointArrayWrap(*(*uintptr)(getPtr(getVal(1))))
			winding := (*bool)(getPtr(getVal(2)))
			cb(sender, &points, winding)
			if points != nil {
				*(*uintptr)(getPtr(getVal(1))) = points.Instance()
			}
		},
	}
}

func makeTShortcutEvent(cb TShortcutEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TShortcutEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (var Msg: TLMKey; var Handled: Boolean);
			msg := (*types.TLMKey)(getPtr(getVal(0)))
			handled := (*bool)(getPtr(getVal(1)))
			cb(msg, handled)
		},
	}
}

func makeTShowHintEvent(cb TShowHintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TShowHintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (var HintStr: string; var CanShow: Boolean; var HintInfo: THintInfo);
			hintStr := api.GoStr(*(*uintptr)(getPtr(getVal(0))))
			canShow := (*bool)(getPtr(getVal(1)))
			hintInfoPtr := *(*tHintInfo)(getPtr(getVal(2)))
			hintInfo := hintInfoPtr.ToGo()
			cb(&hintStr, canShow, &hintInfo)
			*(*uintptr)(getPtr(getVal(0))) = api.PasStr(hintStr)
			if r := hintInfo.ToPas(); r != nil {
				*(*tHintInfo)(getPtr(getVal(2))) = *r
			}
		},
	}
}

func makeTStartDockEvent(cb TStartDockEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TStartDockEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var DragObject: TDragDockObject);
			sender := AsObject(getVal(0))
			var dragObject IDragDockObject
			dragObject = AsDragDockObject(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &dragObject)
			if dragObject != nil {
				*(*uintptr)(getPtr(getVal(1))) = dragObject.Instance()
			}
		},
	}
}

func makeTStartDragEvent(cb TStartDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TStartDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var DragObject: TDragObject);
			sender := AsObject(getVal(0))
			var dragObject IDragObject
			dragObject = AsDragObject(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &dragObject)
			if dragObject != nil {
				*(*uintptr)(getPtr(getVal(1))) = dragObject.Instance()
			}
		},
	}
}

func makeTSysLinkEvent(cb TSysLinkEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TSysLinkEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const Link: string; LinkType: TSysLinkType);
			sender := AsObject(getVal(0))
			link := api.GoStr(getVal(1))
			linkType := types.TSysLinkType(getVal(2))
			cb(sender, link, linkType)
		},
	}
}

func makeTTVAdvancedCustomDrawEvent(cb TTVAdvancedCustomDrawEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVAdvancedCustomDrawEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TCustomTreeView; const ARect: TRect; Stage: TCustomDrawStage; var DefaultDraw: Boolean);
			sender := AsCustomTreeView(getVal(0))
			rect := *(*types.TRect)(getPtr(getVal(1)))
			stage := types.TCustomDrawStage(getVal(2))
			defaultDraw := (*bool)(getPtr(getVal(3)))
			cb(sender, rect, stage, defaultDraw)
		},
	}
}

func makeTTVAdvancedCustomDrawItemEvent(cb TTVAdvancedCustomDrawItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVAdvancedCustomDrawItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState; Stage: TCustomDrawStage; var PaintImages, DefaultDraw: Boolean);
			sender := AsCustomTreeView(getVal(0))
			node := AsTreeNode(getVal(1))
			state := types.TCustomDrawState(getVal(2))
			stage := types.TCustomDrawStage(getVal(3))
			paintImages := (*bool)(getPtr(getVal(4)))
			defaultDraw := (*bool)(getPtr(getVal(5)))
			cb(sender, node, state, stage, paintImages, defaultDraw)
		},
	}
}

func makeTTVChangedEvent(cb TTVChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Node: TTreeNode);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTTVChangingEvent(cb TTVChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Node: TTreeNode; var AllowChange: Boolean);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			allowChange := (*bool)(getPtr(getVal(2)))
			cb(sender, node, allowChange)
		},
	}
}

func makeTTVCollapsingEvent(cb TTVCollapsingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVCollapsingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Node: TTreeNode; var AllowCollapse: Boolean);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			allowCollapse := (*bool)(getPtr(getVal(2)))
			cb(sender, node, allowCollapse)
		},
	}
}

func makeTTVCompareEvent(cb TTVCompareEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVCompareEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; Node1, Node2: TTreeNode; var Compare: Integer);
			sender := AsObject(getVal(0))
			node1 := AsTreeNode(getVal(1))
			node2 := AsTreeNode(getVal(2))
			compare := (*int32)(getPtr(getVal(3)))
			cb(sender, node1, node2, compare)
		},
	}
}

func makeTTVCreateNodeClassEvent(cb TTVCreateNodeClassEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVCreateNodeClassEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TCustomTreeView; var NodeClass: TTreeNodeClass);
			sender := AsCustomTreeView(getVal(0))
			nodeClass := (*types.TTreeNodeClass)(getPtr(getVal(1)))
			cb(sender, nodeClass)
		},
	}
}

func makeTTVCustomCreateNodeEvent(cb TTVCustomCreateNodeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVCustomCreateNodeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TCustomTreeView; var ATreeNode: TTreenode);
			sender := AsCustomTreeView(getVal(0))
			var treeNode ITreeNode
			treeNode = AsTreeNode(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &treeNode)
			if treeNode != nil {
				*(*uintptr)(getPtr(getVal(1))) = treeNode.Instance()
			}
		},
	}
}

func makeTTVCustomDrawArrowEvent(cb TTVCustomDrawArrowEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVCustomDrawArrowEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TCustomTreeView; const ARect: TRect; ACollapsed: Boolean);
			sender := AsCustomTreeView(getVal(0))
			rect := *(*types.TRect)(getPtr(getVal(1)))
			collapsed := api.GoBool(getVal(2))
			cb(sender, rect, collapsed)
		},
	}
}

func makeTTVCustomDrawEvent(cb TTVCustomDrawEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVCustomDrawEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TCustomTreeView; const ARect: TRect; var DefaultDraw: Boolean);
			sender := AsCustomTreeView(getVal(0))
			rect := *(*types.TRect)(getPtr(getVal(1)))
			defaultDraw := (*bool)(getPtr(getVal(2)))
			cb(sender, rect, defaultDraw)
		},
	}
}

func makeTTVCustomDrawItemEvent(cb TTVCustomDrawItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVCustomDrawItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState; var DefaultDraw: Boolean);
			sender := AsCustomTreeView(getVal(0))
			node := AsTreeNode(getVal(1))
			state := types.TCustomDrawState(getVal(2))
			defaultDraw := (*bool)(getPtr(getVal(3)))
			cb(sender, node, state, defaultDraw)
		},
	}
}

func makeTTVEditedEvent(cb TTVEditedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVEditedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Node: TTreeNode; var S: string);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			S := api.GoStr(*(*uintptr)(getPtr(getVal(2))))
			cb(sender, node, &S)
			*(*uintptr)(getPtr(getVal(2))) = api.PasStr(S)
		},
	}
}

func makeTTVEditingEndEvent(cb TTVEditingEndEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVEditingEndEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Node: TTreeNode; Cancel: Boolean);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			cancel := api.GoBool(getVal(2))
			cb(sender, node, cancel)
		},
	}
}

func makeTTVEditingEvent(cb TTVEditingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVEditingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Node: TTreeNode; var AllowEdit: Boolean);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			allowEdit := (*bool)(getPtr(getVal(2)))
			cb(sender, node, allowEdit)
		},
	}
}

func makeTTVExpandedEvent(cb TTVExpandedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVExpandedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; Node: TTreeNode);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTTVExpandingEvent(cb TTVExpandingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVExpandingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Node: TTreeNode; var AllowExpansion: Boolean);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			allowExpansion := (*bool)(getPtr(getVal(2)))
			cb(sender, node, allowExpansion)
		},
	}
}

func makeTTVHasChildrenEvent(cb TTVHasChildrenEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVHasChildrenEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : function(Sender: TCustomTreeView; ANode: TTreeNode): Boolean;
			sender := AsCustomTreeView(getVal(0))
			node := AsTreeNode(getVal(1))
			ret := cb(sender, node)
			*(*bool)(getPtr(getVal(2))) = ret
		},
	}
}

func makeTTVNodeChangedEvent(cb TTVNodeChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTVNodeChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; Node: TTreeNode; ChangeReason: TTreeNodeChangeReason);
			sender := AsObject(getVal(0))
			node := AsTreeNode(getVal(1))
			changeReason := types.TTreeNodeChangeReason(getVal(2))
			cb(sender, node, changeReason)
		},
	}
}

func makeTTabChangingEvent(cb TTabChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTabChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var AllowChange: Boolean);
			sender := AsObject(getVal(0))
			allowChange := (*bool)(getPtr(getVal(1)))
			cb(sender, allowChange)
		},
	}
}

func makeTTabGetImageEvent(cb TTabGetImageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTabGetImageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer);
			sender := AsObject(getVal(0))
			tabIndex := int32(getVal(1))
			imageIndex := (*int32)(getPtr(getVal(2)))
			cb(sender, tabIndex, imageIndex)
		},
	}
}

func makeTTaskDlgClickEvent(cb TTaskDlgClickEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTaskDlgClickEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; AModalResult: TModalResult; var ACanClose: Boolean);
			sender := AsObject(getVal(0))
			modalResult := types.TModalResult(getVal(1))
			canClose := (*bool)(getPtr(getVal(2)))
			cb(sender, modalResult, canClose)
		},
	}
}

func makeTTaskDlgTimerEvent(cb TTaskDlgTimerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TTaskDlgTimerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; TickCount: Cardinal; var Reset: Boolean);
			sender := AsObject(getVal(0))
			tickCount := uint32(getVal(1))
			reset := (*bool)(getPtr(getVal(2)))
			cb(sender, tickCount, reset)
		},
	}
}

func makeTToggledCheckboxEvent(cb TToggledCheckboxEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TToggledCheckboxEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; aCol, aRow: Integer; aState: TCheckboxState);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			state := types.TCheckBoxState(getVal(3))
			cb(sender, col, row, state)
		},
	}
}

func makeTToolBarOnPaintButton(cb TToolBarOnPaintButton) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TToolBarOnPaintButton",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TToolButton; State: integer);
			sender := AsToolButton(getVal(0))
			state := int32(getVal(1))
			cb(sender, state)
		},
	}
}

func makeTUDChangingEvent(cb TUDChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TUDChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender: TObject; var AllowChange: Boolean);
			sender := AsObject(getVal(0))
			allowChange := (*bool)(getPtr(getVal(1)))
			cb(sender, allowChange)
		},
	}
}

func makeTUDChangingEventEx(cb TUDChangingEventEx) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TUDChangingEventEx",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure (Sender: TObject; var AllowChange: Boolean; NewValue: SmallInt; Direction: TUpDownDirection);
			sender := AsObject(getVal(0))
			allowChange := (*bool)(getPtr(getVal(1)))
			newValue := types.SmallInt(getVal(2))
			direction := types.TUpDownDirection(getVal(3))
			cb(sender, allowChange, newValue, direction)
		},
	}
}

func makeTUDClickEvent(cb TUDClickEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TUDClickEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure (Sender: TObject; Button: TUDBtnType);
			sender := AsObject(getVal(0))
			button := types.TUDBtnType(getVal(1))
			cb(sender, button)
		},
	}
}

func makeTUTF8KeyPressEvent(cb TUTF8KeyPressEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TUTF8KeyPressEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var UTF8Key: TUTF8Char);
			sender := AsObject(getVal(0))
			uTF8Key := api.GoStr(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &uTF8Key)
			*(*uintptr)(getPtr(getVal(1))) = api.PasStr(uTF8Key)
		},
	}
}

func makeTUnDockEvent(cb TUnDockEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TUnDockEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean);
			sender := AsObject(getVal(0))
			client := AsControl(getVal(1))
			newTarget := AsWinControl(getVal(2))
			allow := (*bool)(getPtr(getVal(3)))
			cb(sender, client, newTarget, allow)
		},
	}
}

func makeTUserCheckBoxImageEvent(cb TUserCheckBoxImageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TUserCheckBoxImageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TObject; const aCol, aRow: Integer; const CheckedState: TCheckBoxState; var ImageList: TCustomImageList; var ImageIndex: TImageIndex);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			checkedState := types.TCheckBoxState(getVal(3))
			var imageList ICustomImageList
			imageList = AsCustomImageList(*(*uintptr)(getPtr(getVal(4))))
			imageIndex := (*int32)(getPtr(getVal(5)))
			cb(sender, col, row, checkedState, &imageList, imageIndex)
			if imageList != nil {
				*(*uintptr)(getPtr(getVal(4))) = imageList.Instance()
			}
		},
	}
}

func makeTUserCheckboxBitmapEvent(cb TUserCheckboxBitmapEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TUserCheckboxBitmapEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const aCol, aRow: Integer; const CheckedState: TCheckboxState; var ABitmap: TBitmap);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			checkedState := types.TCheckBoxState(getVal(3))
			var bitmap IBitmap
			bitmap = AsBitmap(*(*uintptr)(getPtr(getVal(4))))
			cb(sender, col, row, checkedState, &bitmap)
			if bitmap != nil {
				*(*uintptr)(getPtr(getVal(4))) = bitmap.Instance()
			}
		},
	}
}

func makeTVSTGetTextEvent(cb TVSTGetTextEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVSTGetTextEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; TextType: TVSTTextType; var CellText: String);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			textType := types.TVSTTextType(getVal(3))
			cellText := api.GoStr(*(*uintptr)(getPtr(getVal(4))))
			cb(sender, node, column, textType, &cellText)
			*(*uintptr)(getPtr(getVal(4))) = api.PasStr(cellText)
		},
	}
}

func makeTVSTNewTextEvent(cb TVSTNewTextEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVSTNewTextEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; const NewText: String);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			newText := api.GoStr(getVal(3))
			cb(sender, node, column, newText)
		},
	}
}

func makeTVSTShortenStringEvent(cb TVSTShortenStringEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVSTShortenStringEvent",
		cb: func(getVal func(i int) uintptr) {
			// 8 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; Column: TColumnIndex; const S: String; TextSpace: Integer; var Result: String; var Done: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			S := api.GoStr(getVal(4))
			textSpace := int32(getVal(5))
			result := api.GoStr(*(*uintptr)(getPtr(getVal(6))))
			done := (*bool)(getPtr(getVal(7)))
			cb(sender, targetCanvas, node, column, S, textSpace, &result, done)
			*(*uintptr)(getPtr(getVal(6))) = api.PasStr(result)
		},
	}
}

func makeTVTAddToSelectionEvent(cb TVTAddToSelectionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAddToSelectionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTVTAdvancedHeaderPaintEvent(cb TVTAdvancedHeaderPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAdvancedHeaderPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; var PaintInfo: THeaderPaintInfo; const Elements: THeaderPaintElements);
			sender := AsVTHeader(getVal(0))
			paintInfoPtr := *(*tHeaderPaintInfo)(getPtr(getVal(1)))
			paintInfo := paintInfoPtr.ToGo()
			elements := types.THeaderPaintElements(getVal(2))
			cb(sender, &paintInfo, elements)
			if r := paintInfo.ToPas(); r != nil {
				*(*tHeaderPaintInfo)(getPtr(getVal(1))) = *r
			}
		},
	}
}

func makeTVTAfterAutoFitColumnEvent(cb TVTAfterAutoFitColumnEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterAutoFitColumnEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TVTHeader; Column: TColumnIndex);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			cb(sender, column)
		},
	}
}

func makeTVTAfterAutoFitColumnsEvent(cb TVTAfterAutoFitColumnsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterAutoFitColumnsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(Sender: TVTHeader);
			sender := AsVTHeader(getVal(0))
			cb(sender)
		},
	}
}

func makeTVTAfterCellPaintEvent(cb TVTAfterCellPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterCellPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; Column: TColumnIndex; const CellRect: TRect);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			cellRect := *(*types.TRect)(getPtr(getVal(4)))
			cb(sender, targetCanvas, node, column, cellRect)
		},
	}
}

func makeTVTAfterColumnWidthTrackingEvent(cb TVTAfterColumnWidthTrackingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterColumnWidthTrackingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TVTHeader; Column: TColumnIndex);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			cb(sender, column)
		},
	}
}

func makeTVTAfterGetMaxColumnWidthEvent(cb TVTAfterGetMaxColumnWidthEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterGetMaxColumnWidthEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; Column: TColumnIndex; var MaxWidth: Integer);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			maxWidth := (*int32)(getPtr(getVal(2)))
			cb(sender, column, maxWidth)
		},
	}
}

func makeTVTAfterHeaderHeightTrackingEvent(cb TVTAfterHeaderHeightTrackingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterHeaderHeightTrackingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(Sender: TVTHeader);
			sender := AsVTHeader(getVal(0))
			cb(sender)
		},
	}
}

func makeTVTAfterItemEraseEvent(cb TVTAfterItemEraseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterItemEraseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; const ItemRect: TRect);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			itemRect := *(*types.TRect)(getPtr(getVal(3)))
			cb(sender, targetCanvas, node, itemRect)
		},
	}
}

func makeTVTAfterItemPaintEvent(cb TVTAfterItemPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTAfterItemPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; const ItemRect: TRect);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			itemRect := *(*types.TRect)(getPtr(getVal(3)))
			cb(sender, targetCanvas, node, itemRect)
		},
	}
}

func makeTVTBackgroundPaintEvent(cb TVTBackgroundPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBackgroundPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; const R: TRect; var Handled: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			R := *(*types.TRect)(getPtr(getVal(2)))
			handled := (*bool)(getPtr(getVal(3)))
			cb(sender, targetCanvas, R, handled)
		},
	}
}

func makeTVTBeforeAutoFitColumnEvent(cb TVTBeforeAutoFitColumnEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeAutoFitColumnEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TVTHeader; Column: TColumnIndex; var SmartAutoFitType: TSmartAutoFitType; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			smartAutoFitType := (*types.TSmartAutoFitType)(getPtr(getVal(2)))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, column, smartAutoFitType, allowed)
		},
	}
}

func makeTVTBeforeAutoFitColumnsEvent(cb TVTBeforeAutoFitColumnsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeAutoFitColumnsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TVTHeader; var SmartAutoFitType: TSmartAutoFitType);
			sender := AsVTHeader(getVal(0))
			smartAutoFitType := (*types.TSmartAutoFitType)(getPtr(getVal(1)))
			cb(sender, smartAutoFitType)
		},
	}
}

func makeTVTBeforeCellPaintEvent(cb TVTBeforeCellPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeCellPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; Column: TColumnIndex; CellPaintMode: TVTCellPaintMode; CellRect: TRect; var ContentRect: TRect);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			cellPaintMode := types.TVTCellPaintMode(getVal(4))
			cellRect := *(*types.TRect)(getPtr(getVal(5)))
			contentRect := (*types.TRect)(getPtr(getVal(6)))
			cb(sender, targetCanvas, node, column, cellPaintMode, cellRect, contentRect)
		},
	}
}

func makeTVTBeforeColumnWidthTrackingEvent(cb TVTBeforeColumnWidthTrackingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeColumnWidthTrackingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; Column: TColumnIndex; Shift: TShiftState);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			shift := types.TShiftState(getVal(2))
			cb(sender, column, shift)
		},
	}
}

func makeTVTBeforeDrawLineImageEvent(cb TVTBeforeDrawLineImageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeDrawLineImageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Level: Integer; var PosX: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			level := int32(getVal(2))
			posX := (*int32)(getPtr(getVal(3)))
			cb(sender, node, level, posX)
		},
	}
}

func makeTVTBeforeGetMaxColumnWidthEvent(cb TVTBeforeGetMaxColumnWidthEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeGetMaxColumnWidthEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; Column: TColumnIndex; var UseSmartColumnWidth: Boolean);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			useSmartColumnWidth := (*bool)(getPtr(getVal(2)))
			cb(sender, column, useSmartColumnWidth)
		},
	}
}

func makeTVTBeforeHeaderHeightTrackingEvent(cb TVTBeforeHeaderHeightTrackingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeHeaderHeightTrackingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TVTHeader; Shift: TShiftState);
			sender := AsVTHeader(getVal(0))
			shift := types.TShiftState(getVal(1))
			cb(sender, shift)
		},
	}
}

func makeTVTBeforeItemEraseEvent(cb TVTBeforeItemEraseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeItemEraseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; const ItemRect: TRect; var ItemColor: TColor; var EraseAction: TItemEraseAction);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			itemRect := *(*types.TRect)(getPtr(getVal(3)))
			itemColor := (*types.TColor)(getPtr(getVal(4)))
			eraseAction := (*types.TItemEraseAction)(getPtr(getVal(5)))
			cb(sender, targetCanvas, node, itemRect, itemColor, eraseAction)
		},
	}
}

func makeTVTBeforeItemPaintEvent(cb TVTBeforeItemPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeforeItemPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; const ItemRect: TRect; var CustomDraw: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			itemRect := *(*types.TRect)(getPtr(getVal(3)))
			customDraw := (*bool)(getPtr(getVal(4)))
			cb(sender, targetCanvas, node, itemRect, customDraw)
		},
	}
}

func makeTVTBeginCancelEndEditEvent(cb TVTBeginCancelEndEditEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTBeginCancelEndEditEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): Boolean;
			ret := cb()
			*(*bool)(getPtr(getVal(0))) = ret
		},
	}
}

func makeTVTCanSplitterResizeColumnEvent(cb TVTCanSplitterResizeColumnEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCanSplitterResizeColumnEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TVTHeader; P: TPoint; Column: TColumnIndex; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			P := *(*types.TPoint)(getPtr(getVal(1)))
			column := int32(getVal(2))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, P, column, allowed)
		},
	}
}

func makeTVTCanSplitterResizeHeaderEvent(cb TVTCanSplitterResizeHeaderEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCanSplitterResizeHeaderEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; P: TPoint; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			P := *(*types.TPoint)(getPtr(getVal(1)))
			allowed := (*bool)(getPtr(getVal(2)))
			cb(sender, P, allowed)
		},
	}
}

func makeTVTCanSplitterResizeNodeEvent(cb TVTCanSplitterResizeNodeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCanSplitterResizeNodeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; P: TPoint; Node: PVirtualNode; Column: TColumnIndex; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			P := *(*types.TPoint)(getPtr(getVal(1)))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			allowed := (*bool)(getPtr(getVal(4)))
			cb(sender, P, node, column, allowed)
		},
	}
}

func makeTVTChangeEvent(cb TVTChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTVTChangingEvent(cb TVTChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			allowed := (*bool)(getPtr(getVal(2)))
			cb(sender, node, allowed)
		},
	}
}

func makeTVTCheckChangingEvent(cb TVTCheckChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCheckChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; var NewState: TCheckState; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			newState := (*types.TCheckState)(getPtr(getVal(2)))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, node, newState, allowed)
		},
	}
}

func makeTVTColumnClickEvent(cb TVTColumnClickEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTColumnClickEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender: TBaseVirtualTree; Column: TColumnIndex; Shift: TShiftState);
			sender := AsBaseVirtualTree(getVal(0))
			column := int32(getVal(1))
			shift := types.TShiftState(getVal(2))
			cb(sender, column, shift)
		},
	}
}

func makeTVTColumnDblClickEvent(cb TVTColumnDblClickEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTColumnDblClickEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender: TBaseVirtualTree; Column: TColumnIndex; Shift: TShiftState);
			sender := AsBaseVirtualTree(getVal(0))
			column := int32(getVal(1))
			shift := types.TShiftState(getVal(2))
			cb(sender, column, shift)
		},
	}
}

func makeTVTColumnExportEvent(cb TVTColumnExportEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTColumnExportEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure (Sender: TBaseVirtualTree; aExportType: TVTExportType; Column: TVirtualTreeColumn);
			sender := AsBaseVirtualTree(getVal(0))
			exportType := types.TVTExportType(getVal(1))
			column := AsVirtualTreeColumn(getVal(2))
			cb(sender, exportType, column)
		},
	}
}

func makeTVTColumnWidthDblClickResizeEvent(cb TVTColumnWidthDblClickResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTColumnWidthDblClickResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TVTHeader; Column: TColumnIndex; Shift: TShiftState; P: TPoint; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			shift := types.TShiftState(getVal(2))
			P := *(*types.TPoint)(getPtr(getVal(3)))
			allowed := (*bool)(getPtr(getVal(4)))
			cb(sender, column, shift, P, allowed)
		},
	}
}

func makeTVTColumnWidthTrackingEvent(cb TVTColumnWidthTrackingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTColumnWidthTrackingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TVTHeader; Column: TColumnIndex; Shift: TShiftState; var TrackPoint: TPoint; P: TPoint; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			shift := types.TShiftState(getVal(2))
			trackPoint := (*types.TPoint)(getPtr(getVal(3)))
			P := *(*types.TPoint)(getPtr(getVal(4)))
			allowed := (*bool)(getPtr(getVal(5)))
			cb(sender, column, shift, trackPoint, P, allowed)
		},
	}
}

func makeTVTCompareEvent(cb TVTCompareEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCompareEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; Node1, Node2: PVirtualNode; Column: TColumnIndex; var Result: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			node1 := types.PVirtualNode(getVal(1))
			node2 := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			result := (*int32)(getPtr(getVal(4)))
			cb(sender, node1, node2, column, result)
		},
	}
}

func makeTVTCreateDataObjectEvent(cb TVTCreateDataObjectEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCreateDataObjectEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; out IDataObject: IDataObject);
			sender := AsBaseVirtualTree(getVal(0))
			iDataObject := (*IDataObject)(getPtr(getVal(1)))
			cb(sender, iDataObject)
		},
	}
}

func makeTVTCreateDragManagerEvent(cb TVTCreateDragManagerEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCreateDragManagerEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; out DragManager: IVTDragManager);
			sender := AsBaseVirtualTree(getVal(0))
			var dragManager IVTDragManager
			cb(sender, &dragManager)
			if dragManager != nil {
				*(*uintptr)(getPtr(getVal(1))) = dragManager.Instance()
			}
		},
	}
}

func makeTVTCreateEditorEvent(cb TVTCreateEditorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTCreateEditorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; out EditLink: IVTEditLink);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			var editLink IVTEditLink
			cb(sender, node, column, &editLink)
			if editLink != nil {
				*(*uintptr)(getPtr(getVal(3))) = editLink.Instance()
			}
		},
	}
}

func makeTVTDragAllowedEvent(cb TVTDragAllowedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTDragAllowedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, node, column, allowed)
		},
	}
}

func makeTVTDragDropEvent(cb TVTDragDropEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTDragDropEvent",
		cb: func(getVal func(i int) uintptr) {
			// 8 : procedure(Sender: TBaseVirtualTree; Source: TObject; DataObject: IDataObject; Formats: TFormatArray; Shift: TShiftState; const Pt: TPoint; var Effect: LongWord; Mode: TDropMode);
			sender := AsBaseVirtualTree(getVal(0))
			source := AsObject(getVal(1))
			dataObject := IDataObject(getVal(2))
			formatsPtr := getVal(3)
			formatsLen := int32(getVal(4))
			formats := NewFormatArray(int(formatsLen), formatsPtr)
			shift := types.TShiftState(getVal(5))
			pt := *(*types.TPoint)(getPtr(getVal(6)))
			effect := (*uint32)(getPtr(getVal(7)))
			mode := types.TDropMode(getVal(8))
			cb(sender, source, dataObject, formats, shift, pt, effect, mode)
		},
	}
}

func makeTVTDragOverEvent(cb TVTDragOverEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTDragOverEvent",
		cb: func(getVal func(i int) uintptr) {
			// 8 : procedure(Sender: TBaseVirtualTree; Source: TObject; Shift: TShiftState; State: TDragState; const Pt: TPoint; Mode: TDropMode; var Effect: LongWord; var Accept: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			source := AsObject(getVal(1))
			shift := types.TShiftState(getVal(2))
			state := types.TDragState(getVal(3))
			pt := *(*types.TPoint)(getPtr(getVal(4)))
			mode := types.TDropMode(getVal(5))
			effect := (*uint32)(getPtr(getVal(6)))
			accept := (*bool)(getPtr(getVal(7)))
			cb(sender, source, shift, state, pt, mode, effect, accept)
		},
	}
}

func makeTVTDrawHintEvent(cb TVTDrawHintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTDrawHintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; HintCanvas: TCanvas; Node: PVirtualNode; R: TRect; Column: TColumnIndex);
			sender := AsBaseVirtualTree(getVal(0))
			hintCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			R := *(*types.TRect)(getPtr(getVal(3)))
			column := int32(getVal(4))
			cb(sender, hintCanvas, node, R, column)
		},
	}
}

func makeTVTDrawNodeEvent(cb TVTDrawNodeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTDrawNodeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; const PaintInfo: TVTPaintInfo);
			sender := AsBaseVirtualTree(getVal(0))
			paintInfoPtr := (*tVTPaintInfo)(getPtr(getVal(1)))
			paintInfo := paintInfoPtr.ToGo()
			cb(sender, paintInfo)
		},
	}
}

func makeTVTDrawTextEvent(cb TVTDrawTextEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTDrawTextEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; Column: TColumnIndex; const CellText: String; const CellRect: TRect; var DefaultDraw: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			cellText := api.GoStr(getVal(4))
			cellRect := *(*types.TRect)(getPtr(getVal(5)))
			defaultDraw := (*bool)(getPtr(getVal(6)))
			cb(sender, targetCanvas, node, column, cellText, cellRect, defaultDraw)
		},
	}
}

func makeTVTEditCancelEvent(cb TVTEditCancelEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTEditCancelEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Column: TColumnIndex);
			sender := AsBaseVirtualTree(getVal(0))
			column := int32(getVal(1))
			cb(sender, column)
		},
	}
}

func makeTVTEditChangeEvent(cb TVTEditChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTEditChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			cb(sender, node, column)
		},
	}
}

func makeTVTEditChangingEvent(cb TVTEditChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTEditChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, node, column, allowed)
		},
	}
}

func makeTVTFocusChangeEvent(cb TVTFocusChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTFocusChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			cb(sender, node, column)
		},
	}
}

func makeTVTFocusChangingEvent(cb TVTFocusChangingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTFocusChangingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TBaseVirtualTree; OldNode, NewNode: PVirtualNode; OldColumn, NewColumn: TColumnIndex; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			oldNode := types.PVirtualNode(getVal(1))
			newNode := types.PVirtualNode(getVal(2))
			oldColumn := int32(getVal(3))
			newColumn := int32(getVal(4))
			allowed := (*bool)(getPtr(getVal(5)))
			cb(sender, oldNode, newNode, oldColumn, newColumn, allowed)
		},
	}
}

func makeTVTFreeNodeEvent(cb TVTFreeNodeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTFreeNodeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTVTGetBoundsEvent(cb TVTGetBoundsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetBoundsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 0 : function(): TRect;
			ret := cb()
			*(*types.TRect)(getPtr(getVal(0))) = ret
		},
	}
}

func makeTVTGetCellIsEmptyEvent(cb TVTGetCellIsEmptyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetCellIsEmptyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; var IsEmpty: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			isEmpty := (*bool)(getPtr(getVal(3)))
			cb(sender, node, column, isEmpty)
		},
	}
}

func makeTVTGetCursorEvent(cb TVTGetCursorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetCursorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; var Cursor: TCursor);
			sender := AsBaseVirtualTree(getVal(0))
			cursor := (*types.TCursor)(getPtr(getVal(1)))
			cb(sender, cursor)
		},
	}
}

func makeTVTGetHeaderCursorEvent(cb TVTGetHeaderCursorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetHeaderCursorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TVTHeader; var Cursor: HCURSOR);
			sender := AsVTHeader(getVal(0))
			cursor := (*types.HCURSOR)(getPtr(getVal(1)))
			cb(sender, cursor)
		},
	}
}

func makeTVTGetHintEvent(cb TVTGetHintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetHintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; var LineBreakStyle: TVTTooltipLineBreakStyle; var HintText: String);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			lineBreakStyle := (*types.TVTTooltipLineBreakStyle)(getPtr(getVal(3)))
			hintText := api.GoStr(*(*uintptr)(getPtr(getVal(4))))
			cb(sender, node, column, lineBreakStyle, &hintText)
			*(*uintptr)(getPtr(getVal(4))) = api.PasStr(hintText)
		},
	}
}

func makeTVTGetHintSizeEvent(cb TVTGetHintSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetHintSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; var R: TRect);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			R := (*types.TRect)(getPtr(getVal(3)))
			cb(sender, node, column, R)
		},
	}
}

func makeTVTGetImageEvent(cb TVTGetImageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetImageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Kind: TVTImageKind; Column: TColumnIndex; var Ghosted: Boolean; var ImageIndex: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			kind := types.TVTImageKind(getVal(2))
			column := int32(getVal(3))
			ghosted := (*bool)(getPtr(getVal(4)))
			imageIndex := (*int32)(getPtr(getVal(5)))
			cb(sender, node, kind, column, ghosted, imageIndex)
		},
	}
}

func makeTVTGetImageExEvent(cb TVTGetImageExEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetImageExEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Kind: TVTImageKind; Column: TColumnIndex; var Ghosted: Boolean; var ImageIndex: Integer; var ImageList: TCustomImageList);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			kind := types.TVTImageKind(getVal(2))
			column := int32(getVal(3))
			ghosted := (*bool)(getPtr(getVal(4)))
			imageIndex := (*int32)(getPtr(getVal(5)))
			var imageList ICustomImageList
			imageList = AsCustomImageList(*(*uintptr)(getPtr(getVal(6))))
			cb(sender, node, kind, column, ghosted, imageIndex, &imageList)
			if imageList != nil {
				*(*uintptr)(getPtr(getVal(6))) = imageList.Instance()
			}
		},
	}
}

func makeTVTGetImageTextEvent(cb TVTGetImageTextEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetImageTextEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Kind: TVTImageKind; Column: TColumnIndex; var ImageText: String);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			kind := types.TVTImageKind(getVal(2))
			column := int32(getVal(3))
			imageText := api.GoStr(*(*uintptr)(getPtr(getVal(4))))
			cb(sender, node, kind, column, &imageText)
			*(*uintptr)(getPtr(getVal(4))) = api.PasStr(imageText)
		},
	}
}

func makeTVTGetLineStyleEvent(cb TVTGetLineStyleEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetLineStyleEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; var Bits: Pointer);
			sender := AsBaseVirtualTree(getVal(0))
			bits := (*uintptr)(getPtr(getVal(1)))
			cb(sender, bits)
		},
	}
}

func makeTVTGetNodeDataSizeEvent(cb TVTGetNodeDataSizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetNodeDataSizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; var NodeDataSize: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			nodeDataSize := (*int32)(getPtr(getVal(1)))
			cb(sender, nodeDataSize)
		},
	}
}

func makeTVTGetNodeWidthEvent(cb TVTGetNodeWidthEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetNodeWidthEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; HintCanvas: TCanvas; Node: PVirtualNode; Column: TColumnIndex; var NodeWidth: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			hintCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			nodeWidth := (*int32)(getPtr(getVal(4)))
			cb(sender, hintCanvas, node, column, nodeWidth)
		},
	}
}

func makeTVTGetUserClipboardFormatsEvent(cb TVTGetUserClipboardFormatsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTGetUserClipboardFormatsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; var Formats: TFormatEtcArray);
			sender := AsBaseVirtualTree(getVal(0))
			var formats IFormatEtcArrayWrap
			formats = AsFormatEtcArrayWrap(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &formats)
			if formats != nil {
				*(*uintptr)(getPtr(getVal(1))) = formats.Instance()
			}
		},
	}
}

func makeTVTHeaderClickEvent(cb TVTHeaderClickEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderClickEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TVTHeader; HitInfo: TVTHeaderHitInfo);
			sender := AsVTHeader(getVal(0))
			hitInfo := *(*TVTHeaderHitInfo)(getPtr(getVal(1)))
			cb(sender, hitInfo)
		},
	}
}

func makeTVTHeaderDraggedEvent(cb TVTHeaderDraggedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderDraggedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; Column: TColumnIndex; OldPosition: Integer);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			oldPosition := int32(getVal(2))
			cb(sender, column, oldPosition)
		},
	}
}

func makeTVTHeaderDraggedOutEvent(cb TVTHeaderDraggedOutEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderDraggedOutEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; Column: TColumnIndex; const DropPosition: TPoint);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			dropPosition := *(*types.TPoint)(getPtr(getVal(2)))
			cb(sender, column, dropPosition)
		},
	}
}

func makeTVTHeaderDraggingEvent(cb TVTHeaderDraggingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderDraggingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; Column: TColumnIndex; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			allowed := (*bool)(getPtr(getVal(2)))
			cb(sender, column, allowed)
		},
	}
}

func makeTVTHeaderHeightDblClickResizeEvent(cb TVTHeaderHeightDblClickResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderHeightDblClickResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TVTHeader; var P: TPoint; Shift: TShiftState; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			P := (*types.TPoint)(getPtr(getVal(1)))
			shift := types.TShiftState(getVal(2))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, P, shift, allowed)
		},
	}
}

func makeTVTHeaderHeightTrackingEvent(cb TVTHeaderHeightTrackingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderHeightTrackingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TVTHeader; var P: TPoint; Shift: TShiftState; var Allowed: Boolean);
			sender := AsVTHeader(getVal(0))
			P := (*types.TPoint)(getPtr(getVal(1)))
			shift := types.TShiftState(getVal(2))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, P, shift, allowed)
		},
	}
}

func makeTVTHeaderMouseEvent(cb TVTHeaderMouseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderMouseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TVTHeader; Button: TMouseButton; Shift: TShiftState; X, Y: Integer);
			sender := AsVTHeader(getVal(0))
			button := types.TMouseButton(getVal(1))
			shift := types.TShiftState(getVal(2))
			X := int32(getVal(3))
			Y := int32(getVal(4))
			cb(sender, button, shift, X, Y)
		},
	}
}

func makeTVTHeaderMouseMoveEvent(cb TVTHeaderMouseMoveEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderMouseMoveEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TVTHeader; Shift: TShiftState; X, Y: Integer);
			sender := AsVTHeader(getVal(0))
			shift := types.TShiftState(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, shift, X, Y)
		},
	}
}

func makeTVTHeaderNotifyEvent(cb TVTHeaderNotifyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderNotifyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TVTHeader; Column: TColumnIndex);
			sender := AsVTHeader(getVal(0))
			column := int32(getVal(1))
			cb(sender, column)
		},
	}
}

func makeTVTHeaderPaintEvent(cb TVTHeaderPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TVTHeader; HeaderCanvas: TCanvas; Column: TVirtualTreeColumn; const R: TRect; Hover, Pressed: Boolean; DropMark: TVTDropMarkMode);
			sender := AsVTHeader(getVal(0))
			headerCanvas := AsCanvas(getVal(1))
			column := AsVirtualTreeColumn(getVal(2))
			R := *(*types.TRect)(getPtr(getVal(3)))
			hover := api.GoBool(getVal(4))
			pressed := api.GoBool(getVal(5))
			dropMark := types.TVTDropMarkMode(getVal(6))
			cb(sender, headerCanvas, column, R, hover, pressed, dropMark)
		},
	}
}

func makeTVTHeaderPaintQueryElementsEvent(cb TVTHeaderPaintQueryElementsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHeaderPaintQueryElementsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TVTHeader; var PaintInfo: THeaderPaintInfo; var Elements: THeaderPaintElements);
			sender := AsVTHeader(getVal(0))
			paintInfoPtr := *(*tHeaderPaintInfo)(getPtr(getVal(1)))
			paintInfo := paintInfoPtr.ToGo()
			elements := (*types.THeaderPaintElements)(getPtr(getVal(2)))
			cb(sender, &paintInfo, elements)
			if r := paintInfo.ToPas(); r != nil {
				*(*tHeaderPaintInfo)(getPtr(getVal(1))) = *r
			}
		},
	}
}

func makeTVTHelpContextEvent(cb TVTHelpContextEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHelpContextEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; var HelpContext: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			helpContext := (*int32)(getPtr(getVal(3)))
			cb(sender, node, column, helpContext)
		},
	}
}

func makeTVTHintKindEvent(cb TVTHintKindEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHintKindEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; var Kind: TVTHintKind);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			kind := (*types.TVTHintKind)(getPtr(getVal(3)))
			cb(sender, node, column, kind)
		},
	}
}

func makeTVTHotNodeChangeEvent(cb TVTHotNodeChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTHotNodeChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; OldNode, NewNode: PVirtualNode);
			sender := AsBaseVirtualTree(getVal(0))
			oldNode := types.PVirtualNode(getVal(1))
			newNode := types.PVirtualNode(getVal(2))
			cb(sender, oldNode, newNode)
		},
	}
}

func makeTVTIncrementalSearchEvent(cb TVTIncrementalSearchEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTIncrementalSearchEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; const SearchText: String; var Result: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			searchText := api.GoStr(getVal(2))
			result := (*int32)(getPtr(getVal(3)))
			cb(sender, node, searchText, result)
		},
	}
}

func makeTVTInitChildrenEvent(cb TVTInitChildrenEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTInitChildrenEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; var ChildCount: Cardinal);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			childCount := (*uint32)(getPtr(getVal(2)))
			cb(sender, node, childCount)
		},
	}
}

func makeTVTInitNodeEvent(cb TVTInitNodeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTInitNodeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; ParentNode, Node: PVirtualNode; var InitialStates: TVirtualNodeInitStates);
			sender := AsBaseVirtualTree(getVal(0))
			parentNode := types.PVirtualNode(getVal(1))
			node := types.PVirtualNode(getVal(2))
			initialStates := (*types.TVirtualNodeInitStates)(getPtr(getVal(3)))
			cb(sender, parentNode, node, initialStates)
		},
	}
}

func makeTVTKeyActionEvent(cb TVTKeyActionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTKeyActionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; var CharCode: Word; var Shift: TShiftState; var DoDefault: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			charCode := (*uint16)(getPtr(getVal(1)))
			shift := (*types.TShiftState)(getPtr(getVal(2)))
			doDefault := (*bool)(getPtr(getVal(3)))
			cb(sender, charCode, shift, doDefault)
		},
	}
}

func makeTVTMeasureItemEvent(cb TVTMeasureItemEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTMeasureItemEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; var NodeHeight: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			nodeHeight := (*int32)(getPtr(getVal(3)))
			cb(sender, targetCanvas, node, nodeHeight)
		},
	}
}

func makeTVTMeasureTextEvent(cb TVTMeasureTextEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTMeasureTextEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas; Node: PVirtualNode; Column: TColumnIndex; const CellText: String; var Extent: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			cellText := api.GoStr(getVal(4))
			extent := (*int32)(getPtr(getVal(5)))
			cb(sender, targetCanvas, node, column, cellText, extent)
		},
	}
}

func makeTVTNodeClickEvent(cb TVTNodeClickEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeClickEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; const HitInfo: THitInfo);
			sender := AsBaseVirtualTree(getVal(0))
			hitInfo := *(*THitInfo)(getPtr(getVal(1)))
			cb(sender, hitInfo)
		},
	}
}

func makeTVTNodeCopiedEvent(cb TVTNodeCopiedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeCopiedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTVTNodeCopyingEvent(cb TVTNodeCopyingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeCopyingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node, Target: PVirtualNode; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			target := types.PVirtualNode(getVal(2))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, node, target, allowed)
		},
	}
}

func makeTVTNodeExportEvent(cb TVTNodeExportEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeExportEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function (Sender: TBaseVirtualTree; aExportType: TVTExportType; Node: PVirtualNode): Boolean;
			sender := AsBaseVirtualTree(getVal(0))
			exportType := types.TVTExportType(getVal(1))
			node := types.PVirtualNode(getVal(2))
			ret := cb(sender, exportType, node)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTVTNodeHeightDblClickResizeEvent(cb TVTNodeHeightDblClickResizeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeHeightDblClickResizeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; Shift: TShiftState; P: TPoint; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			shift := types.TShiftState(getVal(3))
			P := *(*types.TPoint)(getPtr(getVal(4)))
			allowed := (*bool)(getPtr(getVal(5)))
			cb(sender, node, column, shift, P, allowed)
		},
	}
}

func makeTVTNodeHeightTrackingEvent(cb TVTNodeHeightTrackingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeHeightTrackingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 7 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; Shift: TShiftState; var TrackPoint: TPoint; P: TPoint; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			shift := types.TShiftState(getVal(3))
			trackPoint := (*types.TPoint)(getPtr(getVal(4)))
			P := *(*types.TPoint)(getPtr(getVal(5)))
			allowed := (*bool)(getPtr(getVal(6)))
			cb(sender, node, column, shift, trackPoint, P, allowed)
		},
	}
}

func makeTVTNodeMovedEvent(cb TVTNodeMovedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeMovedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTVTNodeMovingEvent(cb TVTNodeMovingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTNodeMovingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TBaseVirtualTree; Node, Target: PVirtualNode; var Allowed: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			target := types.PVirtualNode(getVal(2))
			allowed := (*bool)(getPtr(getVal(3)))
			cb(sender, node, target, allowed)
		},
	}
}

func makeTVTOperationEvent(cb TVTOperationEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTOperationEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; OperationKind: TVTOperationKind);
			sender := AsBaseVirtualTree(getVal(0))
			operationKind := types.TVTOperationKind(getVal(1))
			cb(sender, operationKind)
		},
	}
}

func makeTVTPaintEvent(cb TVTPaintEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTPaintEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; TargetCanvas: TCanvas);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			cb(sender, targetCanvas)
		},
	}
}

func makeTVTPaintText(cb TVTPaintText) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTPaintText",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TBaseVirtualTree; const TargetCanvas: TCanvas; Node: PVirtualNode; Column: TColumnIndex; TextType: TVSTTextType);
			sender := AsBaseVirtualTree(getVal(0))
			targetCanvas := AsCanvas(getVal(1))
			node := types.PVirtualNode(getVal(2))
			column := int32(getVal(3))
			textType := types.TVSTTextType(getVal(4))
			cb(sender, targetCanvas, node, column, textType)
		},
	}
}

func makeTVTPopupEvent(cb TVTPopupEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTPopupEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex; const P: TPoint; var AskParent: Boolean; var PopupMenu: TPopupMenu);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			P := *(*types.TPoint)(getPtr(getVal(3)))
			askParent := (*bool)(getPtr(getVal(4)))
			var popupMenu IPopupMenu
			popupMenu = AsPopupMenu(*(*uintptr)(getPtr(getVal(5))))
			cb(sender, node, column, P, askParent, &popupMenu)
			if popupMenu != nil {
				*(*uintptr)(getPtr(getVal(5))) = popupMenu.Instance()
			}
		},
	}
}

func makeTVTPrepareEditEvent(cb TVTPrepareEditEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTPrepareEditEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(Tree: TBaseVirtualTree; Node: PVirtualNode; Column: TColumnIndex): Boolean;
			tree := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			column := int32(getVal(2))
			ret := cb(tree, node, column)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTVTProcessMessageEvent(cb TVTProcessMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTProcessMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(var Msg: TLMessage);
			msg := (*types.TLMessage)(getPtr(getVal(0)))
			cb(msg)
		},
	}
}

func makeTVTRemoveFromSelectionEvent(cb TVTRemoveFromSelectionEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTRemoveFromSelectionEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			cb(sender, node)
		},
	}
}

func makeTVTSaveNodeEvent(cb TVTSaveNodeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTSaveNodeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Stream: TStream);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			stream := AsStream(getVal(2))
			cb(sender, node, stream)
		},
	}
}

func makeTVTSaveTreeEvent(cb TVTSaveTreeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTSaveTreeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; Stream: TStream);
			sender := AsBaseVirtualTree(getVal(0))
			stream := AsStream(getVal(1))
			cb(sender, stream)
		},
	}
}

func makeTVTScrollBarShowEvent(cb TVTScrollBarShowEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTScrollBarShowEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Bar: Integer; Show: Boolean);
			sender := AsBaseVirtualTree(getVal(0))
			bar := int32(getVal(1))
			show := api.GoBool(getVal(2))
			cb(sender, bar, show)
		},
	}
}

func makeTVTScrollEvent(cb TVTScrollEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTScrollEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; DeltaX, DeltaY: Integer);
			sender := AsBaseVirtualTree(getVal(0))
			deltaX := int32(getVal(1))
			deltaY := int32(getVal(2))
			cb(sender, deltaX, deltaY)
		},
	}
}

func makeTVTSetBoundsEvent(cb TVTSetBoundsEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTSetBoundsEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(R: TRect);
			R := *(*types.TRect)(getPtr(getVal(0)))
			cb(R)
		},
	}
}

func makeTVTStateChangeEvent(cb TVTStateChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTStateChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Enter, Leave: TVirtualTreeStates);
			sender := AsBaseVirtualTree(getVal(0))
			enter := types.TVirtualTreeStates(getVal(1))
			leave := types.TVirtualTreeStates(getVal(2))
			cb(sender, enter, leave)
		},
	}
}

func makeTVTStructureChangeEvent(cb TVTStructureChangeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTStructureChangeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TBaseVirtualTree; Node: PVirtualNode; Reason: TChangeReason);
			sender := AsBaseVirtualTree(getVal(0))
			node := types.PVirtualNode(getVal(1))
			reason := types.TChangeReason(getVal(2))
			cb(sender, node, reason)
		},
	}
}

func makeTVTTreeExportEvent(cb TVTTreeExportEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTTreeExportEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; aExportType: TVTExportType);
			sender := AsBaseVirtualTree(getVal(0))
			exportType := types.TVTExportType(getVal(1))
			cb(sender, exportType)
		},
	}
}

func makeTVTUpdatingEvent(cb TVTUpdatingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TVTUpdatingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TBaseVirtualTree; State: TVTUpdateState);
			sender := AsBaseVirtualTree(getVal(0))
			state := types.TVTUpdateState(getVal(1))
			cb(sender, state)
		},
	}
}

func makeTValidateEntryEvent(cb TValidateEntryEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TValidateEntryEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; aCol, aRow: Integer; const OldValue: string; var NewValue: String);
			sender := AsObject(getVal(0))
			col := int32(getVal(1))
			row := int32(getVal(2))
			oldValue := api.GoStr(getVal(3))
			newValue := api.GoStr(*(*uintptr)(getPtr(getVal(4))))
			cb(sender, col, row, oldValue, &newValue)
			*(*uintptr)(getPtr(getVal(4))) = api.PasStr(newValue)
		},
	}
}

func makeTWndMethod(cb TWndMethod) func(msg uintptr) {
	if cb == nil {
		return nil
	}
	return func(msg uintptr) {
		message := (*types.TLMessage)(getPtr(msg))
		cb(message)
	}
}
