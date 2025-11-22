//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import "github.com/energye/lcl/base"

// AsObject Convert a pointer object to an existing class object
func AsObject(obj any) IObject {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TObject)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPersistent Convert a pointer object to an existing class object
func AsPersistent(obj any) IPersistent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPersistent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsComponent Convert a pointer object to an existing class object
func AsComponent(obj any) IComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomApplication Convert a pointer object to an existing class object
func AsCustomApplication(obj any) ICustomApplication {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomApplication)
	base.SetObjectInstance(result, instance)
	return result
}

// AsApplication Convert a pointer object to an existing class object
func AsApplication(obj any) IApplication {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TApplication)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLCLComponent Convert a pointer object to an existing class object
func AsLCLComponent(obj any) ILCLComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLCLComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsControl Convert a pointer object to an existing class object
func AsControl(obj any) IControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWinControl Convert a pointer object to an existing class object
func AsWinControl(obj any) IWinControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWinControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomControl Convert a pointer object to an existing class object
func AsCustomControl(obj any) ICustomControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsScrollingWinControl Convert a pointer object to an existing class object
func AsScrollingWinControl(obj any) IScrollingWinControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TScrollingWinControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDesignControl Convert a pointer object to an existing class object
func AsCustomDesignControl(obj any) ICustomDesignControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDesignControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomForm Convert a pointer object to an existing class object
func AsCustomForm(obj any) ICustomForm {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomForm)
	base.SetObjectInstance(result, instance)
	return result
}

// AsForm Convert a pointer object to an existing class object
func AsForm(obj any) IForm {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TForm)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEngForm Convert a pointer object to an existing class object
func AsEngForm(obj any) IEngForm {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEngForm)
	base.SetObjectInstance(result, instance)
	return result
}

// AsButtonControl Convert a pointer object to an existing class object
func AsButtonControl(obj any) IButtonControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TButtonControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomButton Convert a pointer object to an existing class object
func AsCustomButton(obj any) ICustomButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsButton Convert a pointer object to an existing class object
func AsButton(obj any) IButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBasicAction Convert a pointer object to an existing class object
func AsBasicAction(obj any) IBasicAction {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBasicAction)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomEdit Convert a pointer object to an existing class object
func AsCustomEdit(obj any) ICustomEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEdit Convert a pointer object to an existing class object
func AsEdit(obj any) IEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomAbstractGroupedEdit Convert a pointer object to an existing class object
func AsCustomAbstractGroupedEdit(obj any) ICustomAbstractGroupedEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomAbstractGroupedEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomEditButton Convert a pointer object to an existing class object
func AsCustomEditButton(obj any) ICustomEditButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomEditButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEditButton Convert a pointer object to an existing class object
func AsEditButton(obj any) IEditButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEditButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomMaskEdit Convert a pointer object to an existing class object
func AsCustomMaskEdit(obj any) ICustomMaskEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomMaskEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGEEdit Convert a pointer object to an existing class object
func AsGEEdit(obj any) IGEEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGEEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsEbEdit Convert a pointer object to an existing class object
func AsEbEdit(obj any) IEbEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TEbEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMenu Convert a pointer object to an existing class object
func AsMenu(obj any) IMenu {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMenu)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMainMenu Convert a pointer object to an existing class object
func AsMainMenu(obj any) IMainMenu {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMainMenu)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPopupMenu Convert a pointer object to an existing class object
func AsPopupMenu(obj any) IPopupMenu {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPopupMenu)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomMemo Convert a pointer object to an existing class object
func AsCustomMemo(obj any) ICustomMemo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomMemo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMemo Convert a pointer object to an existing class object
func AsMemo(obj any) IMemo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMemo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCheckBox Convert a pointer object to an existing class object
func AsCustomCheckBox(obj any) ICustomCheckBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCheckBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCheckBox Convert a pointer object to an existing class object
func AsCheckBox(obj any) ICheckBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCheckBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRadioButton Convert a pointer object to an existing class object
func AsRadioButton(obj any) IRadioButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRadioButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomGroupBox Convert a pointer object to an existing class object
func AsCustomGroupBox(obj any) ICustomGroupBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomGroupBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGroupBox Convert a pointer object to an existing class object
func AsGroupBox(obj any) IGroupBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGroupBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGraphicControl Convert a pointer object to an existing class object
func AsGraphicControl(obj any) IGraphicControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGraphicControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomLabel Convert a pointer object to an existing class object
func AsCustomLabel(obj any) ICustomLabel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomLabel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLabel Convert a pointer object to an existing class object
func AsLabel(obj any) ILabel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLabel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomListBox Convert a pointer object to an existing class object
func AsCustomListBox(obj any) ICustomListBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomListBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListBox Convert a pointer object to an existing class object
func AsListBox(obj any) IListBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomComboBox Convert a pointer object to an existing class object
func AsCustomComboBox(obj any) ICustomComboBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomComboBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsComboBox Convert a pointer object to an existing class object
func AsComboBox(obj any) IComboBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TComboBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPanel Convert a pointer object to an existing class object
func AsCustomPanel(obj any) ICustomPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPanel Convert a pointer object to an existing class object
func AsPanel(obj any) IPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomImage Convert a pointer object to an existing class object
func AsCustomImage(obj any) ICustomImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsImage Convert a pointer object to an existing class object
func AsImage(obj any) IImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCanvasHelper Convert a pointer object to an existing class object
func AsFPCanvasHelper(obj any) IFPCanvasHelper {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCanvasHelper)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomFont Convert a pointer object to an existing class object
func AsFPCustomFont(obj any) IFPCustomFont {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomFont)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFont Convert a pointer object to an existing class object
func AsFont(obj any) IFont {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFont)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGraphicsObject Convert a pointer object to an existing class object
func AsGraphicsObject(obj any) IGraphicsObject {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGraphicsObject)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRegion Convert a pointer object to an existing class object
func AsRegion(obj any) IRegion {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRegion)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLinkLabel Convert a pointer object to an existing class object
func AsLinkLabel(obj any) ILinkLabel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLinkLabel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomSpeedButton Convert a pointer object to an existing class object
func AsCustomSpeedButton(obj any) ICustomSpeedButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomSpeedButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsSpeedButton Convert a pointer object to an existing class object
func AsSpeedButton(obj any) ISpeedButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TSpeedButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomSplitter Convert a pointer object to an existing class object
func AsCustomSplitter(obj any) ICustomSplitter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomSplitter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsSplitter Convert a pointer object to an existing class object
func AsSplitter(obj any) ISplitter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TSplitter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomRadioGroup Convert a pointer object to an existing class object
func AsCustomRadioGroup(obj any) ICustomRadioGroup {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomRadioGroup)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRadioGroup Convert a pointer object to an existing class object
func AsRadioGroup(obj any) IRadioGroup {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRadioGroup)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomStaticText Convert a pointer object to an existing class object
func AsCustomStaticText(obj any) ICustomStaticText {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomStaticText)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStaticText Convert a pointer object to an existing class object
func AsStaticText(obj any) IStaticText {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStaticText)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomColorBox Convert a pointer object to an existing class object
func AsCustomColorBox(obj any) ICustomColorBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomColorBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsColorBox Convert a pointer object to an existing class object
func AsColorBox(obj any) IColorBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TColorBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomColorListBox Convert a pointer object to an existing class object
func AsCustomColorListBox(obj any) ICustomColorListBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomColorListBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsColorListBox Convert a pointer object to an existing class object
func AsColorListBox(obj any) IColorListBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TColorListBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCommonDialog Convert a pointer object to an existing class object
func AsCommonDialog(obj any) ICommonDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCommonDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFileDialog Convert a pointer object to an existing class object
func AsFileDialog(obj any) IFileDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFileDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsOpenDialog Convert a pointer object to an existing class object
func AsOpenDialog(obj any) IOpenDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TOpenDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsSaveDialog Convert a pointer object to an existing class object
func AsSaveDialog(obj any) ISaveDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TSaveDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsColorDialog Convert a pointer object to an existing class object
func AsColorDialog(obj any) IColorDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TColorDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFontDialog Convert a pointer object to an existing class object
func AsFontDialog(obj any) IFontDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFontDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPrintDialog Convert a pointer object to an existing class object
func AsCustomPrintDialog(obj any) ICustomPrintDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPrintDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPrintDialog Convert a pointer object to an existing class object
func AsPrintDialog(obj any) IPrintDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPrintDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPreviewFileDialog Convert a pointer object to an existing class object
func AsPreviewFileDialog(obj any) IPreviewFileDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPreviewFileDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsOpenPictureDialog Convert a pointer object to an existing class object
func AsOpenPictureDialog(obj any) IOpenPictureDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TOpenPictureDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsSavePictureDialog Convert a pointer object to an existing class object
func AsSavePictureDialog(obj any) ISavePictureDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TSavePictureDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsSelectDirectoryDialog Convert a pointer object to an existing class object
func AsSelectDirectoryDialog(obj any) ISelectDirectoryDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TSelectDirectoryDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsExtCommonDialog Convert a pointer object to an existing class object
func AsExtCommonDialog(obj any) IExtCommonDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TExtCommonDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCalendarDialog Convert a pointer object to an existing class object
func AsCalendarDialog(obj any) ICalendarDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCalendarDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomRichMemo Convert a pointer object to an existing class object
func AsCustomRichMemo(obj any) ICustomRichMemo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomRichMemo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRichMemo Convert a pointer object to an existing class object
func AsRichMemo(obj any) IRichMemo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRichMemo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRichEdit Convert a pointer object to an existing class object
func AsRichEdit(obj any) IRichEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRichEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomTrackBar Convert a pointer object to an existing class object
func AsCustomTrackBar(obj any) ICustomTrackBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomTrackBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTrackBar Convert a pointer object to an existing class object
func AsTrackBar(obj any) ITrackBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTrackBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomImageList Convert a pointer object to an existing class object
func AsCustomImageList(obj any) ICustomImageList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomImageList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDragImageList Convert a pointer object to an existing class object
func AsDragImageList(obj any) IDragImageList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDragImageList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsImageList Convert a pointer object to an existing class object
func AsImageList(obj any) IImageList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TImageList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomUpDown Convert a pointer object to an existing class object
func AsCustomUpDown(obj any) ICustomUpDown {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomUpDown)
	base.SetObjectInstance(result, instance)
	return result
}

// AsUpDown Convert a pointer object to an existing class object
func AsUpDown(obj any) IUpDown {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TUpDown)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomProgressBar Convert a pointer object to an existing class object
func AsCustomProgressBar(obj any) ICustomProgressBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomProgressBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsProgressBar Convert a pointer object to an existing class object
func AsProgressBar(obj any) IProgressBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TProgressBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDateTimePicker Convert a pointer object to an existing class object
func AsCustomDateTimePicker(obj any) ICustomDateTimePicker {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDateTimePicker)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDateTimePicker Convert a pointer object to an existing class object
func AsDateTimePicker(obj any) IDateTimePicker {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDateTimePicker)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTimeEdit Convert a pointer object to an existing class object
func AsTimeEdit(obj any) ITimeEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTimeEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFileNameEdit Convert a pointer object to an existing class object
func AsFileNameEdit(obj any) IFileNameEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFileNameEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCalendar Convert a pointer object to an existing class object
func AsCustomCalendar(obj any) ICustomCalendar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCalendar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCalendar Convert a pointer object to an existing class object
func AsCalendar(obj any) ICalendar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCalendar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomListView Convert a pointer object to an existing class object
func AsCustomListView(obj any) ICustomListView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomListView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListView Convert a pointer object to an existing class object
func AsListView(obj any) IListView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomTreeView Convert a pointer object to an existing class object
func AsCustomTreeView(obj any) ICustomTreeView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomTreeView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTreeView Convert a pointer object to an existing class object
func AsTreeView(obj any) ITreeView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTreeView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStatusBar Convert a pointer object to an existing class object
func AsStatusBar(obj any) IStatusBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStatusBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsToolWindow Convert a pointer object to an existing class object
func AsToolWindow(obj any) IToolWindow {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TToolWindow)
	base.SetObjectInstance(result, instance)
	return result
}

// AsToolBar Convert a pointer object to an existing class object
func AsToolBar(obj any) IToolBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TToolBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomBitBtn Convert a pointer object to an existing class object
func AsCustomBitBtn(obj any) ICustomBitBtn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomBitBtn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBitBtn Convert a pointer object to an existing class object
func AsBitBtn(obj any) IBitBtn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBitBtn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGraphic Convert a pointer object to an existing class object
func AsGraphic(obj any) IGraphic {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGraphic)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRasterImage Convert a pointer object to an existing class object
func AsRasterImage(obj any) IRasterImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRasterImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomIcon Convert a pointer object to an existing class object
func AsCustomIcon(obj any) ICustomIcon {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomIcon)
	base.SetObjectInstance(result, instance)
	return result
}

// AsIcon Convert a pointer object to an existing class object
func AsIcon(obj any) IIcon {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TIcon)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomBitmap Convert a pointer object to an existing class object
func AsCustomBitmap(obj any) ICustomBitmap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomBitmap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPImageBitmap Convert a pointer object to an existing class object
func AsFPImageBitmap(obj any) IFPImageBitmap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPImageBitmap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBitmap Convert a pointer object to an existing class object
func AsBitmap(obj any) IBitmap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBitmap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStream Convert a pointer object to an existing class object
func AsStream(obj any) IStream {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStream)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomMemoryStream Convert a pointer object to an existing class object
func AsCustomMemoryStream(obj any) ICustomMemoryStream {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomMemoryStream)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMemoryStream Convert a pointer object to an existing class object
func AsMemoryStream(obj any) IMemoryStream {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMemoryStream)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStringWrap Convert a pointer object to an existing class object
func AsStringWrap(obj any) IStringWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStringWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStrings Convert a pointer object to an existing class object
func AsStrings(obj any) IStrings {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStrings)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStringList Convert a pointer object to an existing class object
func AsStringList(obj any) IStringList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStringList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomBrush Convert a pointer object to an existing class object
func AsFPCustomBrush(obj any) IFPCustomBrush {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomBrush)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBrush Convert a pointer object to an existing class object
func AsBrush(obj any) IBrush {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBrush)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomPen Convert a pointer object to an existing class object
func AsFPCustomPen(obj any) IFPCustomPen {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomPen)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPen Convert a pointer object to an existing class object
func AsPen(obj any) IPen {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPen)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMenuItem Convert a pointer object to an existing class object
func AsMenuItem(obj any) IMenuItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMenuItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPicture Convert a pointer object to an existing class object
func AsPicture(obj any) IPicture {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPicture)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCollection Convert a pointer object to an existing class object
func AsCollection(obj any) ICollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListColumns Convert a pointer object to an existing class object
func AsListColumns(obj any) IListColumns {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListColumns)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListItems Convert a pointer object to an existing class object
func AsListItems(obj any) IListItems {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListItems)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTreeNodes Convert a pointer object to an existing class object
func AsTreeNodes(obj any) ITreeNodes {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTreeNodes)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListItem Convert a pointer object to an existing class object
func AsListItem(obj any) IListItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTreeNode Convert a pointer object to an existing class object
func AsTreeNode(obj any) ITreeNode {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTreeNode)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomTabControl Convert a pointer object to an existing class object
func AsCustomTabControl(obj any) ICustomTabControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomTabControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPageControl Convert a pointer object to an existing class object
func AsPageControl(obj any) IPageControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPageControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPage Convert a pointer object to an existing class object
func AsCustomPage(obj any) ICustomPage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTabSheet Convert a pointer object to an existing class object
func AsTabSheet(obj any) ITabSheet {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTabSheet)
	base.SetObjectInstance(result, instance)
	return result
}

// AsScreen Convert a pointer object to an existing class object
func AsScreen(obj any) IScreen {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TScreen)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMouse Convert a pointer object to an existing class object
func AsMouse(obj any) IMouse {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMouse)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCollectionItem Convert a pointer object to an existing class object
func AsCollectionItem(obj any) ICollectionItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCollectionItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListColumn Convert a pointer object to an existing class object
func AsListColumn(obj any) IListColumn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListColumn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStatusPanels Convert a pointer object to an existing class object
func AsStatusPanels(obj any) IStatusPanels {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStatusPanels)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStatusPanel Convert a pointer object to an existing class object
func AsStatusPanel(obj any) IStatusPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStatusPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomFloatSpinEdit Convert a pointer object to an existing class object
func AsCustomFloatSpinEdit(obj any) ICustomFloatSpinEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomFloatSpinEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFloatSpinEdit Convert a pointer object to an existing class object
func AsFloatSpinEdit(obj any) IFloatSpinEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFloatSpinEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDirectoryEdit Convert a pointer object to an existing class object
func AsDirectoryEdit(obj any) IDirectoryEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDirectoryEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFlowPanelControl Convert a pointer object to an existing class object
func AsFlowPanelControl(obj any) IFlowPanelControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFlowPanelControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsOwnedCollection Convert a pointer object to an existing class object
func AsOwnedCollection(obj any) IOwnedCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TOwnedCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFlowPanelControlList Convert a pointer object to an existing class object
func AsFlowPanelControlList(obj any) IFlowPanelControlList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFlowPanelControlList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsColorButton Convert a pointer object to an existing class object
func AsColorButton(obj any) IColorButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TColorButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomSpinEdit Convert a pointer object to an existing class object
func AsCustomSpinEdit(obj any) ICustomSpinEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomSpinEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsSpinEdit Convert a pointer object to an existing class object
func AsSpinEdit(obj any) ISpinEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TSpinEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomCanvas Convert a pointer object to an existing class object
func AsFPCustomCanvas(obj any) IFPCustomCanvas {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomCanvas)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCanvas Convert a pointer object to an existing class object
func AsCanvas(obj any) ICanvas {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCanvas)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPortableNetworkGraphic Convert a pointer object to an existing class object
func AsPortableNetworkGraphic(obj any) IPortableNetworkGraphic {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPortableNetworkGraphic)
	base.SetObjectInstance(result, instance)
	return result
}

// AsJPEGImage Convert a pointer object to an existing class object
func AsJPEGImage(obj any) IJPEGImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TJPEGImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGIFImage Convert a pointer object to an existing class object
func AsGIFImage(obj any) IGIFImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGIFImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomActionList Convert a pointer object to an existing class object
func AsCustomActionList(obj any) ICustomActionList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomActionList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsActionList Convert a pointer object to an existing class object
func AsActionList(obj any) IActionList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TActionList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsContainedAction Convert a pointer object to an existing class object
func AsContainedAction(obj any) IContainedAction {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TContainedAction)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomAction Convert a pointer object to an existing class object
func AsCustomAction(obj any) ICustomAction {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomAction)
	base.SetObjectInstance(result, instance)
	return result
}

// AsAction Convert a pointer object to an existing class object
func AsAction(obj any) IAction {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TAction)
	base.SetObjectInstance(result, instance)
	return result
}

// AsToolButton Convert a pointer object to an existing class object
func AsToolButton(obj any) IToolButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TToolButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomIniFile Convert a pointer object to an existing class object
func AsCustomIniFile(obj any) ICustomIniFile {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomIniFile)
	base.SetObjectInstance(result, instance)
	return result
}

// AsIniFile Convert a pointer object to an existing class object
func AsIniFile(obj any) IIniFile {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TIniFile)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRegistry Convert a pointer object to an existing class object
func AsRegistry(obj any) IRegistry {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRegistry)
	base.SetObjectInstance(result, instance)
	return result
}

// AsClipboard Convert a pointer object to an existing class object
func AsClipboard(obj any) IClipboard {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TClipboard)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMonitor Convert a pointer object to an existing class object
func AsMonitor(obj any) IMonitor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMonitor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPaintBox Convert a pointer object to an existing class object
func AsPaintBox(obj any) IPaintBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPaintBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomTimer Convert a pointer object to an existing class object
func AsCustomTimer(obj any) ICustomTimer {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomTimer)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTimer Convert a pointer object to an existing class object
func AsTimer(obj any) ITimer {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTimer)
	base.SetObjectInstance(result, instance)
	return result
}

// AsList Convert a pointer object to an existing class object
func AsList(obj any) IList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsParaAttributes Convert a pointer object to an existing class object
func AsParaAttributes(obj any) IParaAttributes {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TParaAttributes)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTextAttributes Convert a pointer object to an existing class object
func AsTextAttributes(obj any) ITextAttributes {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTextAttributes)
	base.SetObjectInstance(result, instance)
	return result
}

// AsIconOptions Convert a pointer object to an existing class object
func AsIconOptions(obj any) IIconOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TIconOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomScrollBar Convert a pointer object to an existing class object
func AsCustomScrollBar(obj any) ICustomScrollBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomScrollBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsScrollBar Convert a pointer object to an existing class object
func AsScrollBar(obj any) IScrollBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TScrollBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMaskEdit Convert a pointer object to an existing class object
func AsMaskEdit(obj any) IMaskEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMaskEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomShape Convert a pointer object to an existing class object
func AsCustomShape(obj any) ICustomShape {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomShape)
	base.SetObjectInstance(result, instance)
	return result
}

// AsShape Convert a pointer object to an existing class object
func AsShape(obj any) IShape {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TShape)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBevel Convert a pointer object to an existing class object
func AsBevel(obj any) IBevel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBevel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsScrollBox Convert a pointer object to an existing class object
func AsScrollBox(obj any) IScrollBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TScrollBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCheckListBox Convert a pointer object to an existing class object
func AsCustomCheckListBox(obj any) ICustomCheckListBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCheckListBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCheckListBox Convert a pointer object to an existing class object
func AsCheckListBox(obj any) ICheckListBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCheckListBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsATGauge Convert a pointer object to an existing class object
func AsATGauge(obj any) IATGauge {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TATGauge)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGauge Convert a pointer object to an existing class object
func AsGauge(obj any) IGauge {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGauge)
	base.SetObjectInstance(result, instance)
	return result
}

// AsImageButton Convert a pointer object to an existing class object
func AsImageButton(obj any) IImageButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TImageButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFindDialog Convert a pointer object to an existing class object
func AsFindDialog(obj any) IFindDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFindDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsReplaceDialog Convert a pointer object to an existing class object
func AsReplaceDialog(obj any) IReplaceDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TReplaceDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPrinterSetupDialog Convert a pointer object to an existing class object
func AsCustomPrinterSetupDialog(obj any) ICustomPrinterSetupDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPrinterSetupDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPrinterSetupDialog Convert a pointer object to an existing class object
func AsPrinterSetupDialog(obj any) IPrinterSetupDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPrinterSetupDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPageSetupDialog Convert a pointer object to an existing class object
func AsPageSetupDialog(obj any) IPageSetupDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPageSetupDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDragObject Convert a pointer object to an existing class object
func AsDragObject(obj any) IDragObject {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDragObject)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDragDockObject Convert a pointer object to an existing class object
func AsDragDockObject(obj any) IDragDockObject {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDragDockObject)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomGrid Convert a pointer object to an existing class object
func AsCustomGrid(obj any) ICustomGrid {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomGrid)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomDrawGrid Convert a pointer object to an existing class object
func AsCustomDrawGrid(obj any) ICustomDrawGrid {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomDrawGrid)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomStringGrid Convert a pointer object to an existing class object
func AsCustomStringGrid(obj any) ICustomStringGrid {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomStringGrid)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStringGrid Convert a pointer object to an existing class object
func AsStringGrid(obj any) IStringGrid {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStringGrid)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDrawGrid Convert a pointer object to an existing class object
func AsDrawGrid(obj any) IDrawGrid {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDrawGrid)
	base.SetObjectInstance(result, instance)
	return result
}

// AsValueListEditor Convert a pointer object to an existing class object
func AsValueListEditor(obj any) IValueListEditor {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TValueListEditor)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomHeaderControl Convert a pointer object to an existing class object
func AsCustomHeaderControl(obj any) ICustomHeaderControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomHeaderControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsHeaderControl Convert a pointer object to an existing class object
func AsHeaderControl(obj any) IHeaderControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(THeaderControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsHeaderSection Convert a pointer object to an existing class object
func AsHeaderSection(obj any) IHeaderSection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(THeaderSection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsHeaderSections Convert a pointer object to an existing class object
func AsHeaderSections(obj any) IHeaderSections {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(THeaderSections)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomLabeledEdit Convert a pointer object to an existing class object
func AsCustomLabeledEdit(obj any) ICustomLabeledEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomLabeledEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLabeledEdit Convert a pointer object to an existing class object
func AsLabeledEdit(obj any) ILabeledEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLabeledEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBoundLabel Convert a pointer object to an existing class object
func AsBoundLabel(obj any) IBoundLabel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBoundLabel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomFlowPanel Convert a pointer object to an existing class object
func AsCustomFlowPanel(obj any) ICustomFlowPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomFlowPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFlowPanel Convert a pointer object to an existing class object
func AsFlowPanel(obj any) IFlowPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFlowPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCoolBar Convert a pointer object to an existing class object
func AsCustomCoolBar(obj any) ICustomCoolBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCoolBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoolBar Convert a pointer object to an existing class object
func AsCoolBar(obj any) ICoolBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoolBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoolBands Convert a pointer object to an existing class object
func AsCoolBands(obj any) ICoolBands {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoolBands)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoolBand Convert a pointer object to an existing class object
func AsCoolBand(obj any) ICoolBand {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoolBand)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPrinter Convert a pointer object to an existing class object
func AsPrinter(obj any) IPrinter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPrinter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomTaskDialog Convert a pointer object to an existing class object
func AsCustomTaskDialog(obj any) ICustomTaskDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomTaskDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTaskDialog Convert a pointer object to an existing class object
func AsTaskDialog(obj any) ITaskDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTaskDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTaskDialogButtons Convert a pointer object to an existing class object
func AsTaskDialogButtons(obj any) ITaskDialogButtons {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTaskDialogButtons)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTaskDialogBaseButtonItem Convert a pointer object to an existing class object
func AsTaskDialogBaseButtonItem(obj any) ITaskDialogBaseButtonItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTaskDialogBaseButtonItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTaskDialogButtonItem Convert a pointer object to an existing class object
func AsTaskDialogButtonItem(obj any) ITaskDialogButtonItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTaskDialogButtonItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTaskDialogRadioButtonItem Convert a pointer object to an existing class object
func AsTaskDialogRadioButtonItem(obj any) ITaskDialogRadioButtonItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTaskDialogRadioButtonItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomComboBoxEx Convert a pointer object to an existing class object
func AsCustomComboBoxEx(obj any) ICustomComboBoxEx {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomComboBoxEx)
	base.SetObjectInstance(result, instance)
	return result
}

// AsComboBoxEx Convert a pointer object to an existing class object
func AsComboBoxEx(obj any) IComboBoxEx {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TComboBoxEx)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListControlItems Convert a pointer object to an existing class object
func AsListControlItems(obj any) IListControlItems {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListControlItems)
	base.SetObjectInstance(result, instance)
	return result
}

// AsComboExItems Convert a pointer object to an existing class object
func AsComboExItems(obj any) IComboExItems {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TComboExItems)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListControlItem Convert a pointer object to an existing class object
func AsListControlItem(obj any) IListControlItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListControlItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsComboExItem Convert a pointer object to an existing class object
func AsComboExItem(obj any) IComboExItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TComboExItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomFrame Convert a pointer object to an existing class object
func AsCustomFrame(obj any) ICustomFrame {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomFrame)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFrame Convert a pointer object to an existing class object
func AsFrame(obj any) IFrame {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFrame)
	base.SetObjectInstance(result, instance)
	return result
}

// AsControlScrollBar Convert a pointer object to an existing class object
func AsControlScrollBar(obj any) IControlScrollBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TControlScrollBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsSizeConstraints Convert a pointer object to an existing class object
func AsSizeConstraints(obj any) ISizeConstraints {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TSizeConstraints)
	base.SetObjectInstance(result, instance)
	return result
}

// AsXButton Convert a pointer object to an existing class object
func AsXButton(obj any) IXButton {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TXButton)
	base.SetObjectInstance(result, instance)
	return result
}

// AsAnchorSide Convert a pointer object to an existing class object
func AsAnchorSide(obj any) IAnchorSide {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TAnchorSide)
	base.SetObjectInstance(result, instance)
	return result
}

// AsControlBorderSpacing Convert a pointer object to an existing class object
func AsControlBorderSpacing(obj any) IControlBorderSpacing {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TControlBorderSpacing)
	base.SetObjectInstance(result, instance)
	return result
}

// AsControlChildSizing Convert a pointer object to an existing class object
func AsControlChildSizing(obj any) IControlChildSizing {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TControlChildSizing)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCheckGroup Convert a pointer object to an existing class object
func AsCustomCheckGroup(obj any) ICustomCheckGroup {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCheckGroup)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCheckGroup Convert a pointer object to an existing class object
func AsCheckGroup(obj any) ICheckGroup {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCheckGroup)
	base.SetObjectInstance(result, instance)
	return result
}

// AsToggleBox Convert a pointer object to an existing class object
func AsToggleBox(obj any) IToggleBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TToggleBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCheckCombo Convert a pointer object to an existing class object
func AsCustomCheckCombo(obj any) ICustomCheckCombo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCheckCombo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCheckComboBox Convert a pointer object to an existing class object
func AsCheckComboBox(obj any) ICheckComboBox {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCheckComboBox)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGridColumnTitle Convert a pointer object to an existing class object
func AsGridColumnTitle(obj any) IGridColumnTitle {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGridColumnTitle)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGridColumn Convert a pointer object to an existing class object
func AsGridColumn(obj any) IGridColumn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGridColumn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsGridColumns Convert a pointer object to an existing class object
func AsGridColumns(obj any) IGridColumns {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TGridColumns)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDockManager Convert a pointer object to an existing class object
func AsDockManager(obj any) IDockManager {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDockManager)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDockZone Convert a pointer object to an existing class object
func AsDockZone(obj any) IDockZone {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDockZone)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDockTree Convert a pointer object to an existing class object
func AsDockTree(obj any) IDockTree {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDockTree)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazDockZone Convert a pointer object to an existing class object
func AsLazDockZone(obj any) ILazDockZone {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazDockZone)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazDockTree Convert a pointer object to an existing class object
func AsLazDockTree(obj any) ILazDockTree {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazDockTree)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazDockForm Convert a pointer object to an existing class object
func AsLazDockForm(obj any) ILazDockForm {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazDockForm)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazDockPage Convert a pointer object to an existing class object
func AsLazDockPage(obj any) ILazDockPage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazDockPage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazDockPages Convert a pointer object to an existing class object
func AsLazDockPages(obj any) ILazDockPages {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazDockPages)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazDockSplitter Convert a pointer object to an existing class object
func AsLazDockSplitter(obj any) ILazDockSplitter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazDockSplitter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomTrayIcon Convert a pointer object to an existing class object
func AsCustomTrayIcon(obj any) ICustomTrayIcon {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomTrayIcon)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTrayIcon Convert a pointer object to an existing class object
func AsTrayIcon(obj any) ITrayIcon {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTrayIcon)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomRegion Convert a pointer object to an existing class object
func AsFPCustomRegion(obj any) IFPCustomRegion {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomRegion)
	base.SetObjectInstance(result, instance)
	return result
}

// AsIDesigner Convert a pointer object to an existing class object
func AsIDesigner(obj any) IIDesigner {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TIDesigner)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWindowMagnetOptions Convert a pointer object to an existing class object
func AsWindowMagnetOptions(obj any) IWindowMagnetOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWindowMagnetOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLCLReferenceComponent Convert a pointer object to an existing class object
func AsLCLReferenceComponent(obj any) ILCLReferenceComponent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLCLReferenceComponent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomImageListResolution Convert a pointer object to an existing class object
func AsCustomImageListResolution(obj any) ICustomImageListResolution {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomImageListResolution)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMemoScrollbar Convert a pointer object to an existing class object
func AsMemoScrollbar(obj any) IMemoScrollbar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMemoScrollbar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDragImageListResolution Convert a pointer object to an existing class object
func AsDragImageListResolution(obj any) IDragImageListResolution {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDragImageListResolution)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomImage Convert a pointer object to an existing class object
func AsFPCustomImage(obj any) IFPCustomImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPMemoryImage Convert a pointer object to an existing class object
func AsFPMemoryImage(obj any) IFPMemoryImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPMemoryImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazIntfImage Convert a pointer object to an existing class object
func AsLazIntfImage(obj any) ILazIntfImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazIntfImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPPalette Convert a pointer object to an existing class object
func AsFPPalette(obj any) IFPPalette {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPPalette)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMergedMenuItems Convert a pointer object to an existing class object
func AsMergedMenuItems(obj any) IMergedMenuItems {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMergedMenuItems)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPixmap Convert a pointer object to an existing class object
func AsPixmap(obj any) IPixmap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPixmap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPortableAnyMapGraphic Convert a pointer object to an existing class object
func AsPortableAnyMapGraphic(obj any) IPortableAnyMapGraphic {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPortableAnyMapGraphic)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPreviewFileControl Convert a pointer object to an existing class object
func AsPreviewFileControl(obj any) IPreviewFileControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPreviewFileControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPaperSize Convert a pointer object to an existing class object
func AsPaperSize(obj any) IPaperSize {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPaperSize)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPrinterCanvas Convert a pointer object to an existing class object
func AsPrinterCanvas(obj any) IPrinterCanvas {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPrinterCanvas)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDataModule Convert a pointer object to an existing class object
func AsDataModule(obj any) IDataModule {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDataModule)
	base.SetObjectInstance(result, instance)
	return result
}

// AsItemProp Convert a pointer object to an existing class object
func AsItemProp(obj any) IItemProp {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TItemProp)
	base.SetObjectInstance(result, instance)
	return result
}

// AsValueListStrings Convert a pointer object to an existing class object
func AsValueListStrings(obj any) IValueListStrings {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TValueListStrings)
	base.SetObjectInstance(result, instance)
	return result
}

// AsShortCutList Convert a pointer object to an existing class object
func AsShortCutList(obj any) IShortCutList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TShortCutList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsComponentEnumerator Convert a pointer object to an existing class object
func AsComponentEnumerator(obj any) IComponentEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TComponentEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomImageListResolutionEnumerator Convert a pointer object to an existing class object
func AsCustomImageListResolutionEnumerator(obj any) ICustomImageListResolutionEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomImageListResolutionEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsActionListEnumerator Convert a pointer object to an existing class object
func AsActionListEnumerator(obj any) IActionListEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TActionListEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStringsEnumerator Convert a pointer object to an existing class object
func AsStringsEnumerator(obj any) IStringsEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStringsEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTaskDialogButtonsEnumerator Convert a pointer object to an existing class object
func AsTaskDialogButtonsEnumerator(obj any) ITaskDialogButtonsEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTaskDialogButtonsEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCollectionEnumerator Convert a pointer object to an existing class object
func AsCollectionEnumerator(obj any) ICollectionEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCollectionEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsToolBarEnumerator Convert a pointer object to an existing class object
func AsToolBarEnumerator(obj any) IToolBarEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TToolBarEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTreeNodesEnumerator Convert a pointer object to an existing class object
func AsTreeNodesEnumerator(obj any) ITreeNodesEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTreeNodesEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWinControlEnumerator Convert a pointer object to an existing class object
func AsWinControlEnumerator(obj any) IWinControlEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWinControlEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListItemsEnumerator Convert a pointer object to an existing class object
func AsListItemsEnumerator(obj any) IListItemsEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListItemsEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBasicActionLink Convert a pointer object to an existing class object
func AsBasicActionLink(obj any) IBasicActionLink {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBasicActionLink)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomImageHandler Convert a pointer object to an existing class object
func AsFPCustomImageHandler(obj any) IFPCustomImageHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomImageHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomImageReader Convert a pointer object to an existing class object
func AsFPCustomImageReader(obj any) IFPCustomImageReader {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomImageReader)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomImageWriter Convert a pointer object to an existing class object
func AsFPCustomImageWriter(obj any) IFPCustomImageWriter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomImageWriter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsMenuItemEnumerator Convert a pointer object to an existing class object
func AsMenuItemEnumerator(obj any) IMenuItemEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TMenuItemEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazAccessibleObjectEnumerator Convert a pointer object to an existing class object
func AsLazAccessibleObjectEnumerator(obj any) ILazAccessibleObjectEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazAccessibleObjectEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsChangeLink Convert a pointer object to an existing class object
func AsChangeLink(obj any) IChangeLink {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TChangeLink)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazAccessibleObject Convert a pointer object to an existing class object
func AsLazAccessibleObject(obj any) ILazAccessibleObject {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazAccessibleObject)
	base.SetObjectInstance(result, instance)
	return result
}

// AsThemeServices Convert a pointer object to an existing class object
func AsThemeServices(obj any) IThemeServices {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TThemeServices)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRichMemoInline Convert a pointer object to an existing class object
func AsRichMemoInline(obj any) IRichMemoInline {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRichMemoInline)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListEnumerator Convert a pointer object to an existing class object
func AsListEnumerator(obj any) IListEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPListEnumerator Convert a pointer object to an existing class object
func AsFPListEnumerator(obj any) IFPListEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPListEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPList Convert a pointer object to an existing class object
func AsFPList(obj any) IFPList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBaseVirtualTree Convert a pointer object to an existing class object
func AsBaseVirtualTree(obj any) IBaseVirtualTree {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBaseVirtualTree)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomVirtualDrawTree Convert a pointer object to an existing class object
func AsCustomVirtualDrawTree(obj any) ICustomVirtualDrawTree {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomVirtualDrawTree)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazVirtualDrawTree Convert a pointer object to an existing class object
func AsLazVirtualDrawTree(obj any) ILazVirtualDrawTree {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazVirtualDrawTree)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomVirtualStringTree Convert a pointer object to an existing class object
func AsCustomVirtualStringTree(obj any) ICustomVirtualStringTree {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomVirtualStringTree)
	base.SetObjectInstance(result, instance)
	return result
}

// AsLazVirtualStringTree Convert a pointer object to an existing class object
func AsLazVirtualStringTree(obj any) ILazVirtualStringTree {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TLazVirtualStringTree)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDividerBevel Convert a pointer object to an existing class object
func AsDividerBevel(obj any) IDividerBevel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDividerBevel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsClipboardFormats Convert a pointer object to an existing class object
func AsClipboardFormats(obj any) IClipboardFormats {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TClipboardFormats)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTColors Convert a pointer object to an existing class object
func AsVTColors(obj any) IVTColors {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTColors)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTHeader Convert a pointer object to an existing class object
func AsVTHeader(obj any) IVTHeader {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTHeader)
	base.SetObjectInstance(result, instance)
	return result
}

// AsScrollBarOptions Convert a pointer object to an existing class object
func AsScrollBarOptions(obj any) IScrollBarOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TScrollBarOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomVirtualTreeOptions Convert a pointer object to an existing class object
func AsCustomVirtualTreeOptions(obj any) ICustomVirtualTreeOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomVirtualTreeOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVirtualTreeOptions Convert a pointer object to an existing class object
func AsVirtualTreeOptions(obj any) IVirtualTreeOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVirtualTreeOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTDragImage Convert a pointer object to an existing class object
func AsVTDragImage(obj any) IVTDragImage {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTDragImage)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTVirtualNodeEnumerator Convert a pointer object to an existing class object
func AsVTVirtualNodeEnumerator(obj any) IVTVirtualNodeEnumerator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTVirtualNodeEnumerator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomStringTreeOptions Convert a pointer object to an existing class object
func AsCustomStringTreeOptions(obj any) ICustomStringTreeOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomStringTreeOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStringTreeOptions Convert a pointer object to an existing class object
func AsStringTreeOptions(obj any) IStringTreeOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStringTreeOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVirtualTreeColumns Convert a pointer object to an existing class object
func AsVirtualTreeColumns(obj any) IVirtualTreeColumns {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVirtualTreeColumns)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTFixedAreaConstraints Convert a pointer object to an existing class object
func AsVTFixedAreaConstraints(obj any) IVTFixedAreaConstraints {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTFixedAreaConstraints)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVirtualTreeColumn Convert a pointer object to an existing class object
func AsVirtualTreeColumn(obj any) IVirtualTreeColumn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVirtualTreeColumn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomOpenGLControl Convert a pointer object to an existing class object
func AsCustomOpenGLControl(obj any) ICustomOpenGLControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomOpenGLControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsOpenGLControl Convert a pointer object to an existing class object
func AsOpenGLControl(obj any) IOpenGLControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TOpenGLControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRawImageDescriptionWrap Convert a pointer object to an existing class object
func AsRawImageDescriptionWrap(obj any) IRawImageDescriptionWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRawImageDescriptionWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsRawImageWrap Convert a pointer object to an existing class object
func AsRawImageWrap(obj any) IRawImageWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TRawImageWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVirtualNodeWrap Convert a pointer object to an existing class object
func AsVirtualNodeWrap(obj any) IVirtualNodeWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVirtualNodeWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTVirtualNodeEnumerationWrap Convert a pointer object to an existing class object
func AsVTVirtualNodeEnumerationWrap(obj any) IVTVirtualNodeEnumerationWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTVirtualNodeEnumerationWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPCustomInterpolation Convert a pointer object to an existing class object
func AsFPCustomInterpolation(obj any) IFPCustomInterpolation {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPCustomInterpolation)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFPBoxInterpolation Convert a pointer object to an existing class object
func AsFPBoxInterpolation(obj any) IFPBoxInterpolation {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFPBoxInterpolation)
	base.SetObjectInstance(result, instance)
	return result
}

// AsScaledImageListResolutionWrap Convert a pointer object to an existing class object
func AsScaledImageListResolutionWrap(obj any) IScaledImageListResolutionWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TScaledImageListResolutionWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsInterfacedObject Convert a pointer object to an existing class object
func AsInterfacedObject(obj any) IInterfacedObject {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TInterfacedObject)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTDragManager Convert a pointer object to an existing class object
func AsVTDragManager(obj any) IVTDragManager {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTDragManager)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStreamAdapter Convert a pointer object to an existing class object
func AsStreamAdapter(obj any) IStreamAdapter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStreamAdapter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsUint32ArrayWrap Convert a pointer object to an existing class object
func AsUint32ArrayWrap(obj any) IUint32ArrayWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TUint32ArrayWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsUint64ArrayWrap Convert a pointer object to an existing class object
func AsUint64ArrayWrap(obj any) IUint64ArrayWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TUint64ArrayWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsInt64ArrayWrap Convert a pointer object to an existing class object
func AsInt64ArrayWrap(obj any) IInt64ArrayWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TInt64ArrayWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsStringArrayWrap Convert a pointer object to an existing class object
func AsStringArrayWrap(obj any) IStringArrayWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TStringArrayWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsBrushPatternWrap Convert a pointer object to an existing class object
func AsBrushPatternWrap(obj any) IBrushPatternWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TBrushPatternWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPointArrayWrap Convert a pointer object to an existing class object
func AsPointArrayWrap(obj any) IPointArrayWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPointArrayWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsFormatEtcArrayWrap Convert a pointer object to an existing class object
func AsFormatEtcArrayWrap(obj any) IFormatEtcArrayWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TFormatEtcArrayWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTaskDialogProgressBar Convert a pointer object to an existing class object
func AsTaskDialogProgressBar(obj any) ITaskDialogProgressBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTaskDialogProgressBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomGraphicControl Convert a pointer object to an existing class object
func AsCustomGraphicControl(obj any) ICustomGraphicControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomGraphicControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsActionLink Convert a pointer object to an existing class object
func AsActionLink(obj any) IActionLink {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TActionLink)
	base.SetObjectInstance(result, instance)
	return result
}

// AsControlActionLink Convert a pointer object to an existing class object
func AsControlActionLink(obj any) IControlActionLink {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TControlActionLink)
	base.SetObjectInstance(result, instance)
	return result
}

// AsApplicationProperties Convert a pointer object to an existing class object
func AsApplicationProperties(obj any) IApplicationProperties {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TApplicationProperties)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPanelBitBtn Convert a pointer object to an existing class object
func AsPanelBitBtn(obj any) IPanelBitBtn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPanelBitBtn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomButtonPanel Convert a pointer object to an existing class object
func AsCustomButtonPanel(obj any) ICustomButtonPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomButtonPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsButtonPanel Convert a pointer object to an existing class object
func AsButtonPanel(obj any) IButtonPanel {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TButtonPanel)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomControlBar Convert a pointer object to an existing class object
func AsCustomControlBar(obj any) ICustomControlBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomControlBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsControlBar Convert a pointer object to an existing class object
func AsControlBar(obj any) IControlBar {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TControlBar)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCtrlBand Convert a pointer object to an existing class object
func AsCtrlBand(obj any) ICtrlBand {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCtrlBand)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomPairSplitter Convert a pointer object to an existing class object
func AsCustomPairSplitter(obj any) ICustomPairSplitter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomPairSplitter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPairSplitter Convert a pointer object to an existing class object
func AsPairSplitter(obj any) IPairSplitter {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPairSplitter)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPairSplitterSide Convert a pointer object to an existing class object
func AsPairSplitterSide(obj any) IPairSplitterSide {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPairSplitterSide)
	base.SetObjectInstance(result, instance)
	return result
}

// AsPopupNotifier Convert a pointer object to an existing class object
func AsPopupNotifier(obj any) IPopupNotifier {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TPopupNotifier)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCalculatorDialog Convert a pointer object to an existing class object
func AsCalculatorDialog(obj any) ICalculatorDialog {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCalculatorDialog)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDateEdit Convert a pointer object to an existing class object
func AsDateEdit(obj any) IDateEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDateEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomCheckBoxThemed Convert a pointer object to an existing class object
func AsCustomCheckBoxThemed(obj any) ICustomCheckBoxThemed {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomCheckBoxThemed)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCheckBoxThemed Convert a pointer object to an existing class object
func AsCheckBoxThemed(obj any) ICheckBoxThemed {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCheckBoxThemed)
	base.SetObjectInstance(result, instance)
	return result
}

// AsExtendedNotebook Convert a pointer object to an existing class object
func AsExtendedNotebook(obj any) IExtendedNotebook {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TExtendedNotebook)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomControlFilterEdit Convert a pointer object to an existing class object
func AsCustomControlFilterEdit(obj any) ICustomControlFilterEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomControlFilterEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsListFilterEdit Convert a pointer object to an existing class object
func AsListFilterEdit(obj any) IListFilterEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TListFilterEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTreeFilterEdit Convert a pointer object to an existing class object
func AsTreeFilterEdit(obj any) ITreeFilterEdit {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTreeFilterEdit)
	base.SetObjectInstance(result, instance)
	return result
}

// AsTreeFilterBranch Convert a pointer object to an existing class object
func AsTreeFilterBranch(obj any) ITreeFilterBranch {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TTreeFilterBranch)
	base.SetObjectInstance(result, instance)
	return result
}

// AsVTHeaderPopupMenu Convert a pointer object to an existing class object
func AsVTHeaderPopupMenu(obj any) IVTHeaderPopupMenu {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TVTHeaderPopupMenu)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCustomVTEditLink Convert a pointer object to an existing class object
func AsCustomVTEditLink(obj any) ICustomVTEditLink {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCustomVTEditLink)
	base.SetObjectInstance(result, instance)
	return result
}

// AsHintWindow Convert a pointer object to an existing class object
func AsHintWindow(obj any) IHintWindow {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(THintWindow)
	base.SetObjectInstance(result, instance)
	return result
}

// AsDesigner Convert a pointer object to an existing class object
func AsDesigner(obj any) IDesigner {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TDesigner)
	base.SetObjectInstance(result, instance)
	return result
}
