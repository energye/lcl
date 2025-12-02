//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/lcl/api/imports"
	"sync"
	"unsafe"

	"github.com/energye/lcl/types"
)

// WidgetType
//
//	WtWIN32 = 1: Windows Win32
//	WtGTK2 = 2: Linux GTK2
//	WtGTK3 = 3: Linux GTK3
//	WtCOCOA = 4: macOS Cocoa
type WidgetType uint16

var _GlobalWidgetType = WtInvalid

const (
	WtInvalid WidgetType = iota // 无效类型, 默认值
	WtWIN32                     // Windows Win32
	WtGTK2                      // Linux GTK2
	WtGTK3                      // Linux GTK3
	WtCOCOA                     // MacOS Cocoa
)

func (m WidgetType) IsWin32() bool {
	return m == WtWIN32
}

func (m WidgetType) IsGTK2() bool {
	return m == WtGTK2
}

func (m WidgetType) IsGTK3() bool {
	return m == WtGTK3
}

func (m WidgetType) IsCOCOA() bool {
	return m == WtCOCOA
}

// -------------------- TApplication ---------------------------

func Application_Instance() uintptr {
	return libPreDefAPI().SysCallN(_Application_Instance)
}

func Application_CreateForm(app uintptr) uintptr {
	return libPreDefAPI().SysCallN(_Application_CreateForm, app)
}

func Application_Run(app uintptr) {
	defer func() {
		LibRelease()
	}()
	libPreDefAPI().SysCallN(_Application_Run, app)
}

func Application_Initialize(obj uintptr) {
	libPreDefAPI().SysCallN(_Application_Initialize, obj)
}

func Application_SetRunLoopReceived(obj uintptr, proc uintptr) {
	libPreDefAPI().SysCallN(_Application_SetRunLoopReceived, obj, proc)
}

// -------------------- TClipboard ---------------------------

func Clipboard_Instance() uintptr {
	return libPreDefAPI().SysCallN(_Clipboard_Instance)
}

func SetClipboard(obj uintptr) uintptr {
	return libPreDefAPI().SysCallN(_DSetClipboard, obj)
}

func RegisterClipboardFormat(aFormat string) types.TClipboardFormat {
	return types.TClipboardFormat(libPreDefAPI().SysCallN(_DRegisterClipboardFormat, PasStr(aFormat)))
}

func PredefinedClipboardFormat(aFormat types.TPredefinedClipboardFormat) types.TClipboardFormat {
	return types.TClipboardFormat(libPreDefAPI().SysCallN(_DPredefinedClipboardFormat, uintptr(aFormat)))
}

// -------------------- TMouse ---------------------------

func Mouse_Instance() uintptr {
	return libPreDefAPI().SysCallN(_Mouse_Instance)
}

// -------------------- TPrinter ---------------------------

func Printer_Instance() uintptr {
	return libPreDefAPI().SysCallN(_Printer_Instance)
}

// -------------------- TScreen ---------------------------

func Screen_Instance() uintptr {
	return libPreDefAPI().SysCallN(_Screen_Instance)
}

// -------------------- Procedures/Functions ---------------------------

func GetStringArrOf(p uintptr, index int) string {
	return GoStr(libPreDefAPI().SysCallN(_DGetStringArrOf, p, uintptr(index)))
}

func StrLen(p uintptr) int {
	return int(libPreDefAPI().SysCallN(_DStrLen, p))
}

func Move(src, dest uintptr, nLen int) {
	libPreDefAPI().SysCallN(_DMove, src, dest, uintptr(nLen))
}

func ShowMessage(s string) {
	libPreDefAPI().SysCallN(_DShowMessage, PasStr(s))
}

func GetMainInstance() types.HINST {
	r := libPreDefAPI().SysCallN(_DGetMainInstance)
	return types.HINST(r)
}

func MessageDlg(Msg string, DlgType types.TMsgDlgType, Buttons types.TMsgDlgButtons, HelpCtx int32) int32 {
	return int32(libPreDefAPI().SysCallN(_DMessageDlg, PasStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx)))
}

func TextToShortCut(val string) types.TShortCut {
	return types.TShortCut(libPreDefAPI().SysCallN(_DTextToShortCut, PasStr(val)))
}

func ShortCutToText(val types.TShortCut) string {
	return GoStr(libPreDefAPI().SysCallN(_DShortCutToText, uintptr(val)))
}

func SysOpen(filename string) {
	libPreDefAPI().SysCallN(_DSysOpen, PasStr(filename))
}

func ExtractFilePath(filename string) string {
	return GoStr(libPreDefAPI().SysCallN(_DExtractFilePath, PasStr(filename)))
}

func FileExists(filename string) bool {
	return GoBool(libPreDefAPI().SysCallN(_DFileExists, PasStr(filename)))
}

func SelectDirectory1(options types.TSelectDirOpts) (bool, string) {
	var ptr uintptr
	v := GoBool(libPreDefAPI().SysCallN(_DSelectDirectory1, uintptr(unsafe.Pointer(&ptr)), uintptr(options), 0))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func SelectDirectory2(caption, root string, showHidden bool) (bool, string) {
	var ptr uintptr
	v := GoBool(libPreDefAPI().SysCallN(_DSelectDirectory2, PasStr(caption), PasStr(root), PasBool(showHidden), uintptr(unsafe.Pointer(&ptr))))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func RunOnMainAsyncCall(id int) {
	libPreDefAPI().SysCallN(_DRunMainAsyncCall, uintptr(unsafe.Pointer(&id)))
}

func RunOnMainSyncCall(fn func()) {
	threadSyncRef.Lock()
	defer threadSyncRef.Unlock()
	threadSyncProc = fn
	libPreDefAPI().SysCallN(_DRunMainSyncCall)
	threadSyncProc = nil
}

func InputBox(aCaption, aPrompt, aDefault string) string {
	return GoStr(libPreDefAPI().SysCallN(_DInputBox, PasStr(aCaption), PasStr(aPrompt), PasStr(aDefault)))
}

func InputQuery(aCaption, aPrompt string, value *string) bool {
	if value == nil {
		return false
	}
	var strPtr uintptr
	r := libPreDefAPI().SysCallN(_DInputQuery, PasStr(aCaption), PasStr(aPrompt), PasStr(*value), uintptr(unsafe.Pointer(&strPtr)))
	if strPtr != 0 {
		*value = GoStr(strPtr)
	}
	return GoBool(r)
}

func PasswordBox(aCaption, aPrompt string) string {
	return GoStr(libPreDefAPI().SysCallN(_DPasswordBox, PasStr(aCaption), PasStr(aPrompt)))
}

// InputCombo stringList = IStringList
func InputCombo(aCaption, aPrompt string, stringList uintptr) int32 {
	return int32(libPreDefAPI().SysCallN(_DInputCombo, PasStr(aCaption), PasStr(aPrompt), stringList))
}

// InputComboEx stringList = IStringList
func InputComboEx(aCaption, aPrompt string, stringList uintptr, allowCustomText bool) string {
	return GoStr(libPreDefAPI().SysCallN(_DInputComboEx, PasStr(aCaption), PasStr(aPrompt), stringList, PasBool(allowCustomText)))
}

func SysLocale(aInfo *types.TSysLocale) {
	libPreDefAPI().SysCallN(_DSysLocale, uintptr(unsafe.Pointer(aInfo)))
}

func SetPropertyValue(instance uintptr, propName, value string) {
	libPreDefAPI().SysCallN(_DSetPropertyValue, instance, PasStr(propName), PasStr(value))
}

func SetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	libPreDefAPI().SysCallN(_DSetPropertySecValue, instance, PasStr(propName), PasStr(secPropName), PasStr(value))
}

func GUIDToString(guid types.TGUID) string {
	return GoStr(libPreDefAPI().SysCallN(_DGUIDToString, uintptr(unsafe.Pointer(&guid))))
}

func StringToGUID(str string) (guid types.TGUID) {
	libPreDefAPI().SysCallN(_DStringToGUID, PasStr(str), uintptr(unsafe.Pointer(&guid)))
	return
}

func CreateGUID() (guid types.TGUID) {
	libPreDefAPI().SysCallN(_DCreateGUID, uintptr(unsafe.Pointer(&guid)))
	return
}

func GetLibResourceCount() int32 {
	return int32(libPreDefAPI().SysCallN(_DGetLibResourceCount))
}

func GetLibResourceItem(aIndex int32) (ret types.TLibResource) {
	item := struct {
		Name     uintptr
		ValuePtr uintptr
	}{}
	libPreDefAPI().SysCallN(_DGetLibResourceItem, uintptr(aIndex), uintptr(unsafe.Pointer(&item)))
	ret.Name = GoStr(item.Name)
	ret.Ptr = item.ValuePtr
	return
}

func ModifyLibResource(aPtr uintptr, aValue string) {
	libPreDefAPI().SysCallN(_DModifyLibResource, aPtr, PasStr(aValue))
}

func LCLVersion() (major, minor, release, patch uint16, fullVersion uint32, version string) {
	var versionPtr uintptr
	libPreDefAPI().SysCallN(_GetLCLVersion, uintptr(unsafe.Pointer(&major)), uintptr(unsafe.Pointer(&minor)), uintptr(unsafe.Pointer(&release)),
		uintptr(unsafe.Pointer(&patch)), uintptr(unsafe.Pointer(&fullVersion)), uintptr(unsafe.Pointer(&versionPtr)))
	version = GoStr(versionPtr)
	return
}

// Widget 获当前组件库类型
func Widget() WidgetType {
	if _GlobalWidgetType == WtInvalid {
		r := libPreDefAPI().SysCallN(_GetLCLWidget)
		_GlobalWidgetType = WidgetType(r)
	}
	return _GlobalWidgetType
}

func LibAbout() string {
	return GoStr(libPreDefAPI().SysCallN(_DLibAbout))
}

func MainThreadId() uintptr {
	return libPreDefAPI().SysCallN(_DMainThreadId)
}

func CurrentThreadId() uintptr {
	return libPreDefAPI().SysCallN(_DCurrentThreadId)
}

func InitGoDll(aMainThreadId uintptr) {
	libPreDefAPI().SysCallN(_DInitGoDll, aMainThreadId)
}

func FindControl(handle types.HWND) uintptr {
	return libPreDefAPI().SysCallN(_DFindControl, handle)
}

func FindLCLControl(screenPos types.TPoint) uintptr {
	return libPreDefAPI().SysCallN(_DFindLCLControl, uintptr(unsafe.Pointer(&screenPos)))
}

func FindOwnerControl(handle types.HWND) uintptr {
	return libPreDefAPI().SysCallN(_DFindOwnerControl, handle)
}

func FindControlAtPosition(position types.TPoint, allowDisabled bool) uintptr {
	return libPreDefAPI().SysCallN(_DFindControlAtPosition, uintptr(unsafe.Pointer(&position)), PasBool(allowDisabled))
}

func FindLCLWindow(screenPos types.TPoint, allowDisabled bool) uintptr {
	return libPreDefAPI().SysCallN(_DFindLCLWindow, uintptr(unsafe.Pointer(&screenPos)), PasBool(allowDisabled))
}

func FindDragTarget(position types.TPoint, allowDisabled bool) uintptr {
	return libPreDefAPI().SysCallN(_DFindDragTarget, uintptr(unsafe.Pointer(&position)), PasBool(allowDisabled))
}

// FreeAndNil 用于释放以类形式返回的指针
// 该函数不适用于 IInterface 类型
func FreeAndNil(obj uintptr) {
	libPreDefAPI().SysCallN(_DFreeAndNil, uintptr(unsafe.Pointer(&obj)))
}

func SetNil(obj uintptr) {
	libPreDefAPI().SysCallN(_DSetNil, uintptr(unsafe.Pointer(&obj)))
}

func FreePointer(data uintptr) {
	libPreDefAPI().SysCallN(_DFreePointer, data)
}

func DateTimeToUnix(dateTime float64) (result int64) {
	libPreDefAPI().SysCallN(_DToUnixTime, uintptr(unsafe.Pointer(&dateTime)), uintptr(unsafe.Pointer(&result)))
	return
}

func UnixToDateTime(dateTime int64) (result float64) {
	libPreDefAPI().SysCallN(_DUnixToTime, uintptr(unsafe.Pointer(&dateTime)), uintptr(unsafe.Pointer(&result)))
	return
}

func SendMessage(hWd types.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return libPreDefAPI().SysCallN(_DSendMessage, hWd, uintptr(msg), wParam, lParam)
}

func PostMessage(hWd types.HWND, msg uint32, wParam, lParam uintptr) bool {
	return GoBool(libPreDefAPI().SysCallN(_DPostMessage, hWd, uintptr(msg), wParam, lParam))
}

func IsIconic(hWnd types.HWND) bool {
	return GoBool(libPreDefAPI().SysCallN(_DIsIconic, hWnd))
}

func IsWindow(hWnd types.HWND) bool {
	return GoBool(libPreDefAPI().SysCallN(_DIsWindow, hWnd))
}

func IsZoomed(hWnd types.HWND) bool {
	return GoBool(libPreDefAPI().SysCallN(_DIsZoomed, hWnd))
}

func IsWindowVisible(hWnd types.HWND) bool {
	return GoBool(libPreDefAPI().SysCallN(_DIsWindowVisible, hWnd))
}

func GetDC(hWnd types.HWND) types.HDC {
	return libPreDefAPI().SysCallN(_DGetDC, hWnd)
}

func ReleaseDC(hWnd types.HWND, dc types.HDC) int {
	return int(libPreDefAPI().SysCallN(_DReleaseDC, hWnd, dc))
}

func SetForegroundWindow(hWnd types.HWND) bool {
	return GoBool(libPreDefAPI().SysCallN(_DSetForegroundWindow, hWnd))
}

func WindowFromPoint(point types.TPoint) types.HWND {
	return libPreDefAPI().SysCallN(_DWindowFromPoint, uintptr(unsafe.Pointer(&point)))
}

// SetEventCallback 设置所有不同类型的回调函数
func SetEventCallback(ptr uintptr, eventType EventCallbackType) {
	libPreDefAPI().SysCallN(_SetEventCallback, ptr, uintptr(eventType))
}

func SetComponentDesignMode(componentInstance uintptr, value bool) {
	if componentInstance == 0 {
		return
	}
	libPreDefAPI().SysCallN(_SetComponentDesignMode, componentInstance, PasBool(value))
}

func SetComponentDesignInstanceMode(componentInstance uintptr, value bool) {
	if componentInstance == 0 {
		return
	}
	libPreDefAPI().SysCallN(_SetComponentDesignInstanceMode, componentInstance, PasBool(value))
}

func SetComponentInlineMode(componentInstance uintptr, value bool) {
	if componentInstance == 0 {
		return
	}
	libPreDefAPI().SysCallN(_SetComponentInlineMode, componentInstance, PasBool(value))
}

func SetComponentAncestorMode(componentInstance uintptr, value bool) {
	if componentInstance == 0 {
		return
	}
	libPreDefAPI().SysCallN(_SetComponentAncestorMode, componentInstance, PasBool(value))
}

func SetWidgetSetDesigning(componentInstance uintptr) {
	if componentInstance == 0 {
		return
	}
	libPreDefAPI().SysCallN(_SetWidgetSetDesigning, componentInstance)
}

func GetComponentProperties(componentInstance, stringListInstance uintptr) {
	if componentInstance == 0 || stringListInstance == 0 {
		return
	}
	libPreDefAPI().SysCallN(_GetComponentProperties, componentInstance, stringListInstance)
}

func SetOnGetDesignerFormEvent(fn uintptr) {
	libPreDefAPI().SysCallN(_SetOnGetDesignerFormEvent, fn)
}

func ControlWindowProc(controlInstance uintptr, message *types.TLMessage) {
	if controlInstance == 0 || message == nil {
		return
	}
	libPreDefAPI().SysCallN(_Control_WindowProc, controlInstance, uintptr(unsafe.Pointer(message)))
}

// LocalesLCLResourceStrings
// 使用指定语言的本地化文件，翻译 LCLStrConsts.pas 中的资源字符串。
// TranslateLCLResourceStrings 是一个字符串类型函数，用于将 LCLStrConsts.pas 单元中的资源字符串翻译为 Lang 参数指定的语言标识符。Lang 的值需符合 ISO 639 标准（语言名称表示代码），详见：ISO 639 - 语言名称表示代码。
// 例如：'de'（德语）或 'ru'（俄语）。传入空字符串（”）将使用系统默认语言。
// Dir 为可选路径，用于指定存储字符串常量翻译值的 .po 或 .mo 文件所在位置。传入空字符串（”）将使用预定义目录（如 languages 或 locale），该目录用于存放 LCLStrConsts.pas 单元对应的本地化文件（文件名格式：lclstrconsts.<语言标识符>.po 或 lclstrconsts.<语言标识符>.mo）。
// TranslateLCLResourceStrings 会调用 TranslateUnitResourceStringsEx 例程，传入指定的参数值，并以 lclstrconsts 作为翻译的基准单元名。TranslateUnitResourceStringsEx 会解析平台 / 操作系统所需的路径、语言标识符和文件扩展名，并应用翻译后的值。
// 返回值
// 返回翻译时实际使用的语言标识符或区域设置代码。若未找到指定语言的本地化文件，或发生错误，返回值可能为空字符串（”）。
func LocalesLCLResourceStrings(lang, dir string) string {
	if lang == "" || dir == "" {
		return ""
	}
	r := libPreDefAPI().SysCallN(_LocalesLCLResourceStrings, PasStr(lang), PasStr(dir))
	return GoStr(r)
}

// LocalesUnitResourceStringsEx
// 使用指定语言和路径对应的本地化文件，翻译某个单元中的字符串常量。
// TranslateUnitResourceStringsEx 是一个便捷例程，用于将指定源文件中的字符串常量翻译为目标语言。它提供参数以指定源文件、本地化文件路径以及所需的语言。
// Lang 参数存储用于翻译字符串常量的可选语言标识符，其值示例为 'en'（英语）、'ru'（俄语）或 'de'（德语）。传入空字符串（”）将使用平台或操作系统的默认语言标识符。
// Dir 参数存储本地化文件（.po、.mo）的所在路径，可为相对于应用程序可执行文件所在目录的相对路径，也可为指向本地化文件目录的完全限定路径。示例：'languages' 或 /usr/share/locale。传入空字符串（”）表示使用 'languages' 或 'locale' 等预定义目录查找本地化文件。
// LocaleFileName 参数存储用于翻译指定单元中值的本地化文件基准名称，其值示例为 'lazaruside' 或 'lclstrconts'。若该参数为空字符串（”），例程不执行任何操作。
// LocaleUnitName 参数存储该例程中待翻译字符串常量所在源代码文件的基准名称。传入空字符串（”）时，将使用 LocaleFileName 的值作为源代码文件的基准名称。
// TranslateUnitResourceStringsEx 会调用 FindLocaleFileName 实现例程，解析所需的语言标识符、本地化文件的路径及文件扩展名。该例程优先使用 .po 文件，若未找到对应的 .po 文件，则使用 .mo 文件。
// TranslateUnitResourceStringsEx 会调用 translations.pp 单元中的 TranslateUnitResourceStrings 例程，通过本地化文件路径翻译指定文件。
// 若未找到 .po 文件路径，TranslateUnitResourceStringsEx 会调用 FCL（Free Pascal 类库）的 gettext.pp 单元中的 TranslateResourceStrings 例程，使用 .mo 文件进行翻译。
// 返回值
// 返回值存储翻译时实际使用的语言代码或区域标识。若返回空字符串（”），表示未找到与指定参数匹配的本地化文件，或翻译过程中发生错误。
func LocalesUnitResourceStringsEx(lang, dir, localeFileName string, localeUnitName string) string {
	if lang == "" || dir == "" || localeFileName == "" {
		return ""
	}
	r := libPreDefAPI().SysCallN(_LocalesUnitResourceStringsEx, PasStr(lang), PasStr(dir), PasStr(localeFileName), PasStr(localeUnitName))
	return GoStr(r)
}

// LocalesSetDefaultLang
// 设置字符串翻译所使用的默认语言。
// SetDefaultLang 是一个用于设置字符串翻译默认语言的过程。
// Lang 参数存储用于翻译字符串的语言标识符。其值需符合 ISO 639 标准（语言名称表示代码），详见：ISO 639 - 语言名称表示代码[http://www.loc.gov/standards/iso639-2/php/code_list.php]
// 该参数的默认值为空字符串（”），表示使用系统的默认语言标识符。此值可通过命令行参数传入（例如 '--lang ru' 或 '--lang=es'），也可通过 LANG 环境变量覆盖。降级备选值为平台的语言标识符：在 Windows 系统中，该值是通过 GetLocaleInfo 函数获取的语言代码和国家 / 地区代码，格式如 'en_US' 或 'zh_CN'；在其他平台中，降级备选值取自 LC_ALL 或 LC_MESSAGE 环境变量。
// Dir 参数存储 .po 或 .mo 本地化文件所在的目录名（例如 'mylng'）。默认值为空字符串（”），表示使用 'languages' 或 'locale' 等预定义目录查找翻译所需的本地化文件。
// LocaleFileName 参数存储用于翻译字符串值的本地化文件基准名称（例如 'lazaruside' 或 'debuggerstrconst'）。该参数的默认值为空字符串（”），此时将使用应用程序可执行文件的基准名称（例如 'project1.exe' 对应的基准名称为 'project1'）。
// 该基准名称与语言标识符、.po 或 .mo 文件扩展名组合后，形成包含本地化字符串值的完整文件名。例如：'lazaruside.ru.po'、'debuggerstrconst.zh_CN.po' 或 'project1.es.po'。Dir 参数的值会作为前缀添加，构成本地化文件的完整路径。
// ForceUpdate 参数指示加载翻译时是否立即更新用户界面。若从某个单元的初始化部分调用 SetDefaultLang，需将 ForceUpdate 设置为 False。该参数的默认值为 True。
// SetDefaultLang 会先验证语言标识符 / 区域标识是否有效，以及指定的 .po 或 .mo 文件是否存在，再应用这些设置。该过程会调用 FindLocaleFileName 函数获取翻译资源对应的文件名：优先应用 .po 文件，若不存在则使用 .mo 文件获取翻译后的字符串值。若找到符合指定条件的本地化文件，还会通过 TranslateLCLResourceStrings 翻译 LCL 资源字符串常量。
// 当 ForceUpdate 为 True 时，表单（Forms）、数据模块（Data Modules）、控件（Controls）、组件（Components）及持久化对象（Persistent objects）中使用的字符串属性会被翻译。此操作仅对包含设置器（写访问）过程且具备运行时类型信息（RTTI，Run Time Type Information）的成员执行。用户界面元素的更新通过 TUpdateTranslator 完成。
func LocalesSetDefaultLang(lang, dir, localeFileName string, forceUpdate ...bool) string {
	if lang == "" || dir == "" || localeFileName == "" {
		return ""
	}
	fu := len(forceUpdate) > 0
	if fu {
		fu = forceUpdate[0]
	}
	r := libPreDefAPI().SysCallN(_LocalesSetDefaultLang, PasStr(lang), PasStr(dir), PasStr(localeFileName), PasBool(fu))
	return GoStr(r)
}

// LocalesUnitResourceStringsFormStream
// 翻译单元字符串, 例如 LCLStrConsts
// 参数: AStream 是 lang xx.po 文件字节流
// 参数: ALocaleUnitName 待翻译的单元资源字符串的单元名, 例如: lclstrconsts 单元名
func LocalesUnitResourceStringsFormStream(stream uintptr, localeUnitName string) bool {
	if stream == 0 || localeUnitName == "" {
		return false
	}
	r := libPreDefAPI().SysCallN(_LocalesUnitResourceStringsFormStream, stream, PasStr(localeUnitName))
	return GoBool(r)
}

// LocalesSetDefaultLangFormStream
// 设置默认语言
// 参数: AStream 是 lang xx.po 文件字节流
// 参数: ALocaleUnitName 待翻译的单元资源字符串的单元名, 例如: lclstrconsts 单元名
// 参数: ForceUpdate 是否强制翻译更新组件
func LocalesSetDefaultLangFormStream(stream uintptr, localeUnitName string, forceUpdate ...bool) bool {
	if stream == 0 {
		return false
	}
	fu := len(forceUpdate) > 0
	if fu {
		fu = forceUpdate[0]
	}
	r := libPreDefAPI().SysCallN(_LocalesSetDefaultLangFormStream, stream, PasStr(localeUnitName), PasBool(fu))
	return GoBool(r)
}

// NewInstanceByComponentClass 通过组件类型创建新的实例
// 该函数调用底层系统API来实例化指定类型的组件对象
//
//	typeClass - 组件类型的指针地址，用于标识要创建的组件类型
//
//	uintptr - 新创建实例的指针地址，如果创建失败则返回0
func NewInstanceByComponentClass(class types.TClass) uintptr {
	r := libPreDefAPI().SysCallN(_NewInstanceByComponentClass, class)
	return r
}

var (
	libPreDefOnce   sync.Once
	libPreDefImport *imports.Imports
)

func libPreDefAPI() *imports.Imports {
	libPreDefOnce.Do(func() {
		libPreDefImport = NewDefaultImports()
		libPreDefImport.Table = []*imports.Table{
			/* iota */ imports.NewTable("DSysLocale", 0),
			/* iota */ imports.NewTable("AddLazarusResources", 0),
			/* iota */ imports.NewTable("AddLazarusResources2", 0),
			/* iota */ imports.NewTable("GetLCLVersion", 0),
			/* iota */ imports.NewTable("GetLCLWidget", 0),
			/* iota */ imports.NewTable("Application_Instance", 0),
			/* iota */ imports.NewTable("Application_CreateForm", 0),
			/* iota */ imports.NewTable("Application_Run", 0),
			/* iota */ imports.NewTable("Application_Initialize", 0),
			/* iota */ imports.NewTable("Application_SetRunLoopReceived", 0),
			/* iota */ imports.NewTable("Mouse_Instance", 0),
			/* iota */ imports.NewTable("Screen_Instance", 0),
			/* iota */ imports.NewTable("DSendMessage", 0),
			/* iota */ imports.NewTable("DPostMessage", 0),
			/* iota */ imports.NewTable("DIsIconic", 0),
			/* iota */ imports.NewTable("DIsWindow", 0),
			/* iota */ imports.NewTable("DIsZoomed", 0),
			/* iota */ imports.NewTable("DIsWindowVisible", 0),
			/* iota */ imports.NewTable("DGetDC", 0),
			/* iota */ imports.NewTable("DReleaseDC", 0),
			/* iota */ imports.NewTable("DSetForegroundWindow", 0),
			/* iota */ imports.NewTable("DRegisterClipboardFormat", 0),
			/* iota */ imports.NewTable("DWindowFromPoint", 0),
			/* iota */ imports.NewTable("SetEventCallback", 0),
			/* iota */ imports.NewTable("DRunMainAsyncCall", 0),
			/* iota */ imports.NewTable("DRunMainSyncCall", 0),
			/* iota */ imports.NewTable("DGetStringArrOf", 0),
			/* iota */ imports.NewTable("DStrLen", 0),
			/* iota */ imports.NewTable("DMove", 0),
			/* iota */ imports.NewTable("DShowMessage", 0),
			/* iota */ imports.NewTable("DGetMainInstance", 0),
			/* iota */ imports.NewTable("DSysOpen", 0),
			/* iota */ imports.NewTable("DTextToShortCut", 0),
			/* iota */ imports.NewTable("DShortCutToText", 0),
			/* iota */ imports.NewTable("DMessageDlg", 0),
			/* iota */ imports.NewTable("DExtractFilePath", 0),
			/* iota */ imports.NewTable("DFileExists", 0),
			/* iota */ imports.NewTable("DMainThreadId", 0),
			/* iota */ imports.NewTable("DCurrentThreadId", 0),
			/* iota */ imports.NewTable("DSelectDirectory1", 0),
			/* iota */ imports.NewTable("DSelectDirectory2", 0),
			/* iota */ imports.NewTable("DInputBox", 0),
			/* iota */ imports.NewTable("DInputQuery", 0),
			/* iota */ imports.NewTable("DPasswordBox", 0),
			/* iota */ imports.NewTable("DInputCombo", 0),
			/* iota */ imports.NewTable("DInputComboEx", 0),
			/* iota */ imports.NewTable("DSetPropertyValue", 0),
			/* iota */ imports.NewTable("DSetPropertySecValue", 0),
			/* iota */ imports.NewTable("Clipboard_Instance", 0),
			/* iota */ imports.NewTable("DPredefinedClipboardFormat", 0),
			/* iota */ imports.NewTable("DSetClipboard", 0),
			/* iota */ imports.NewTable("DGUIDToString", 0),
			/* iota */ imports.NewTable("DStringToGUID", 0),
			/* iota */ imports.NewTable("DCreateGUID", 0),
			/* iota */ imports.NewTable("Printer_Instance", 0),
			/* iota */ imports.NewTable("DGetLibResourceCount", 0),
			/* iota */ imports.NewTable("DGetLibResourceItem", 0),
			/* iota */ imports.NewTable("DModifyLibResource", 0),
			/* iota */ imports.NewTable("DLibAbout", 0),
			/* iota */ imports.NewTable("DInitGoDll", 0),
			/* iota */ imports.NewTable("DFindControl", 0),
			/* iota */ imports.NewTable("DFindLCLControl", 0),
			/* iota */ imports.NewTable("DFindOwnerControl", 0),
			/* iota */ imports.NewTable("DFindControlAtPosition", 0),
			/* iota */ imports.NewTable("DFindLCLWindow", 0),
			/* iota */ imports.NewTable("DFindDragTarget", 0),
			/* iota */ imports.NewTable("DFreeAndNil", 0),
			/* iota */ imports.NewTable("DSetNil", 0),
			/* iota */ imports.NewTable("DFreePointer", 0),
			/* iota */ imports.NewTable("DToUnixTime", 0),
			/* iota */ imports.NewTable("DUnixToTime", 0),
			/* iota */ imports.NewTable("SetComponentDesignMode", 0),
			/* iota */ imports.NewTable("SetComponentDesignInstanceMode", 0),
			/* iota */ imports.NewTable("SetComponentInlineMode", 0),
			/* iota */ imports.NewTable("SetComponentAncestorMode", 0),
			/* iota */ imports.NewTable("SetWidgetSetDesigning", 0),
			/* iota */ imports.NewTable("GetComponentProperties", 0),
			/* iota */ imports.NewTable("SetOnGetDesignerFormEvent", 0),
			/* iota */ imports.NewTable("Control_WindowProc", 0),
			/* iota */ imports.NewTable("LocalesLCLResourceStrings", 0),
			/* iota */ imports.NewTable("LocalesUnitResourceStringsEx", 0),
			/* iota */ imports.NewTable("LocalesSetDefaultLang", 0),
			/* iota */ imports.NewTable("LocalesUnitResourceStringsFormStream", 0),
			/* iota */ imports.NewTable("LocalesSetDefaultLangFormStream", 0),
			/* iota */ imports.NewTable("NewInstanceByComponentClass", 0),
		}
	})
	return libPreDefImport
}
