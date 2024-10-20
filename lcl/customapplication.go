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

// ICustomApplication Parent: IComponent
type ICustomApplication interface {
	IComponent
	ExeName() string                                                                                    // property
	HelpFile() string                                                                                   // property
	SetHelpFile(AValue string)                                                                          // property
	Terminated() bool                                                                                   // property
	Title() string                                                                                      // property
	SetTitle(AValue string)                                                                             // property
	Location() string                                                                                   // property
	Params(Index int32) string                                                                          // property
	ParamCount() int32                                                                                  // property
	EnvironmentVariable(envName string) string                                                          // property
	OptionChar() Char                                                                                   // property
	SetOptionChar(AValue Char)                                                                          // property
	CaseSensitiveOptions() bool                                                                         // property
	SetCaseSensitiveOptions(AValue bool)                                                                // property
	StopOnException() bool                                                                              // property
	SetStopOnException(AValue bool)                                                                     // property
	ExceptionExitCode() int32                                                                           // property
	SetExceptionExitCode(AValue int32)                                                                  // property
	EventLogFilter() TEventLogTypes                                                                     // property
	SetEventLogFilter(AValue TEventLogTypes)                                                            // property
	FindOptionIndex(S string, Longopt *bool, StartAt int32) int32                                       // function
	GetOptionValue(S string) string                                                                     // function
	GetOptionValue1(C Char, S string) string                                                            // function
	GetOptionValues(C Char, S string) TStringArray                                                      // function
	HasOption(S string) bool                                                                            // function
	HasOption1(C Char, S string) bool                                                                   // function
	CheckOptions(ShortOptions string, Longopts IStrings, Opts, NonOpts IStrings, AllErrors bool) string // function
	CheckOptions1(ShortOptions string, Longopts IStrings, AllErrors bool) string                        // function
	CheckOptions2(ShortOptions string, LongOpts string, AllErrors bool) string                          // function
	HandleException(Sender IObject)                                                                     // procedure
	Initialize()                                                                                        // procedure
	Run()                                                                                               // procedure
	ShowException(E IException)                                                                         // procedure
	Terminate()                                                                                         // procedure
	Terminate1(AExitCode int32)                                                                         // procedure
	GetEnvironmentList(List IStrings, NamesOnly bool)                                                   // procedure
	GetEnvironmentList1(List IStrings)                                                                  // procedure
	Log(EventType TEventType, Msg string)                                                               // procedure
	SetOnException(fn TExceptionEvent)                                                                  // property event
}

// TCustomApplication Parent: TComponent
type TCustomApplication struct {
	TComponent
	exceptionPtr uintptr
}

func NewCustomApplication(AOwner IComponent) ICustomApplication {
	r1 := customApplicationImportAPI().SysCallN(5, GetObjectUintptr(AOwner))
	return AsCustomApplication(r1)
}

func (m *TCustomApplication) ExeName() string {
	r1 := customApplicationImportAPI().SysCallN(9, m.Instance())
	return GoStr(r1)
}

func (m *TCustomApplication) HelpFile() string {
	r1 := customApplicationImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomApplication) SetHelpFile(AValue string) {
	customApplicationImportAPI().SysCallN(19, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomApplication) Terminated() bool {
	r1 := customApplicationImportAPI().SysCallN(32, m.Instance())
	return GoBool(r1)
}

func (m *TCustomApplication) Title() string {
	r1 := customApplicationImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomApplication) SetTitle(AValue string) {
	customApplicationImportAPI().SysCallN(33, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomApplication) Location() string {
	r1 := customApplicationImportAPI().SysCallN(21, m.Instance())
	return GoStr(r1)
}

func (m *TCustomApplication) Params(Index int32) string {
	r1 := customApplicationImportAPI().SysCallN(25, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TCustomApplication) ParamCount() int32 {
	r1 := customApplicationImportAPI().SysCallN(24, m.Instance())
	return int32(r1)
}

func (m *TCustomApplication) EnvironmentVariable(envName string) string {
	r1 := customApplicationImportAPI().SysCallN(6, m.Instance(), PascalStr(envName))
	return GoStr(r1)
}

func (m *TCustomApplication) OptionChar() Char {
	r1 := customApplicationImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomApplication) SetOptionChar(AValue Char) {
	customApplicationImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomApplication) CaseSensitiveOptions() bool {
	r1 := customApplicationImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomApplication) SetCaseSensitiveOptions(AValue bool) {
	customApplicationImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomApplication) StopOnException() bool {
	r1 := customApplicationImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomApplication) SetStopOnException(AValue bool) {
	customApplicationImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomApplication) ExceptionExitCode() int32 {
	r1 := customApplicationImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomApplication) SetExceptionExitCode(AValue int32) {
	customApplicationImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomApplication) EventLogFilter() TEventLogTypes {
	r1 := customApplicationImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TEventLogTypes(r1)
}

func (m *TCustomApplication) SetEventLogFilter(AValue TEventLogTypes) {
	customApplicationImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomApplication) FindOptionIndex(S string, Longopt *bool, StartAt int32) int32 {
	var result1 uintptr
	r1 := customApplicationImportAPI().SysCallN(10, m.Instance(), PascalStr(S), uintptr(unsafePointer(&result1)), uintptr(StartAt))
	*Longopt = GoBool(result1)
	return int32(r1)
}

func (m *TCustomApplication) GetOptionValue(S string) string {
	r1 := customApplicationImportAPI().SysCallN(13, m.Instance(), PascalStr(S))
	return GoStr(r1)
}

func (m *TCustomApplication) GetOptionValue1(C Char, S string) string {
	r1 := customApplicationImportAPI().SysCallN(14, m.Instance(), uintptr(C), PascalStr(S))
	return GoStr(r1)
}

func (m *TCustomApplication) GetOptionValues(C Char, S string) TStringArray {
	r1 := customApplicationImportAPI().SysCallN(15, m.Instance(), uintptr(C), PascalStr(S))
	return TStringArray(r1)
}

func (m *TCustomApplication) HasOption(S string) bool {
	r1 := customApplicationImportAPI().SysCallN(17, m.Instance(), PascalStr(S))
	return GoBool(r1)
}

func (m *TCustomApplication) HasOption1(C Char, S string) bool {
	r1 := customApplicationImportAPI().SysCallN(18, m.Instance(), uintptr(C), PascalStr(S))
	return GoBool(r1)
}

func (m *TCustomApplication) CheckOptions(ShortOptions string, Longopts IStrings, Opts, NonOpts IStrings, AllErrors bool) string {
	r1 := customApplicationImportAPI().SysCallN(1, m.Instance(), PascalStr(ShortOptions), GetObjectUintptr(Longopts), GetObjectUintptr(Opts), GetObjectUintptr(NonOpts), PascalBool(AllErrors))
	return GoStr(r1)
}

func (m *TCustomApplication) CheckOptions1(ShortOptions string, Longopts IStrings, AllErrors bool) string {
	r1 := customApplicationImportAPI().SysCallN(2, m.Instance(), PascalStr(ShortOptions), GetObjectUintptr(Longopts), PascalBool(AllErrors))
	return GoStr(r1)
}

func (m *TCustomApplication) CheckOptions2(ShortOptions string, LongOpts string, AllErrors bool) string {
	r1 := customApplicationImportAPI().SysCallN(3, m.Instance(), PascalStr(ShortOptions), PascalStr(LongOpts), PascalBool(AllErrors))
	return GoStr(r1)
}

func CustomApplicationClass() TClass {
	ret := customApplicationImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TCustomApplication) HandleException(Sender IObject) {
	customApplicationImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(Sender))
}

func (m *TCustomApplication) Initialize() {
	customApplicationImportAPI().SysCallN(20, m.Instance())
}

func (m *TCustomApplication) Run() {
	customApplicationImportAPI().SysCallN(26, m.Instance())
}

func (m *TCustomApplication) ShowException(E IException) {
	customApplicationImportAPI().SysCallN(28, m.Instance(), GetObjectUintptr(E))
}

func (m *TCustomApplication) Terminate() {
	customApplicationImportAPI().SysCallN(30, m.Instance())
}

func (m *TCustomApplication) Terminate1(AExitCode int32) {
	customApplicationImportAPI().SysCallN(31, m.Instance(), uintptr(AExitCode))
}

func (m *TCustomApplication) GetEnvironmentList(List IStrings, NamesOnly bool) {
	customApplicationImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(List), PascalBool(NamesOnly))
}

func (m *TCustomApplication) GetEnvironmentList1(List IStrings) {
	customApplicationImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(List))
}

func (m *TCustomApplication) Log(EventType TEventType, Msg string) {
	customApplicationImportAPI().SysCallN(22, m.Instance(), uintptr(EventType), PascalStr(Msg))
}

func (m *TCustomApplication) SetOnException(fn TExceptionEvent) {
	if m.exceptionPtr != 0 {
		RemoveEventElement(m.exceptionPtr)
	}
	m.exceptionPtr = MakeEventDataPtr(fn)
	customApplicationImportAPI().SysCallN(27, m.Instance(), m.exceptionPtr)
}

var (
	customApplicationImport       *imports.Imports = nil
	customApplicationImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomApplication_CaseSensitiveOptions", 0),
		/*1*/ imports.NewTable("CustomApplication_CheckOptions", 0),
		/*2*/ imports.NewTable("CustomApplication_CheckOptions1", 0),
		/*3*/ imports.NewTable("CustomApplication_CheckOptions2", 0),
		/*4*/ imports.NewTable("CustomApplication_Class", 0),
		/*5*/ imports.NewTable("CustomApplication_Create", 0),
		/*6*/ imports.NewTable("CustomApplication_EnvironmentVariable", 0),
		/*7*/ imports.NewTable("CustomApplication_EventLogFilter", 0),
		/*8*/ imports.NewTable("CustomApplication_ExceptionExitCode", 0),
		/*9*/ imports.NewTable("CustomApplication_ExeName", 0),
		/*10*/ imports.NewTable("CustomApplication_FindOptionIndex", 0),
		/*11*/ imports.NewTable("CustomApplication_GetEnvironmentList", 0),
		/*12*/ imports.NewTable("CustomApplication_GetEnvironmentList1", 0),
		/*13*/ imports.NewTable("CustomApplication_GetOptionValue", 0),
		/*14*/ imports.NewTable("CustomApplication_GetOptionValue1", 0),
		/*15*/ imports.NewTable("CustomApplication_GetOptionValues", 0),
		/*16*/ imports.NewTable("CustomApplication_HandleException", 0),
		/*17*/ imports.NewTable("CustomApplication_HasOption", 0),
		/*18*/ imports.NewTable("CustomApplication_HasOption1", 0),
		/*19*/ imports.NewTable("CustomApplication_HelpFile", 0),
		/*20*/ imports.NewTable("CustomApplication_Initialize", 0),
		/*21*/ imports.NewTable("CustomApplication_Location", 0),
		/*22*/ imports.NewTable("CustomApplication_Log", 0),
		/*23*/ imports.NewTable("CustomApplication_OptionChar", 0),
		/*24*/ imports.NewTable("CustomApplication_ParamCount", 0),
		/*25*/ imports.NewTable("CustomApplication_Params", 0),
		/*26*/ imports.NewTable("CustomApplication_Run", 0),
		/*27*/ imports.NewTable("CustomApplication_SetOnException", 0),
		/*28*/ imports.NewTable("CustomApplication_ShowException", 0),
		/*29*/ imports.NewTable("CustomApplication_StopOnException", 0),
		/*30*/ imports.NewTable("CustomApplication_Terminate", 0),
		/*31*/ imports.NewTable("CustomApplication_Terminate1", 0),
		/*32*/ imports.NewTable("CustomApplication_Terminated", 0),
		/*33*/ imports.NewTable("CustomApplication_Title", 0),
	}
)

func customApplicationImportAPI() *imports.Imports {
	if customApplicationImport == nil {
		customApplicationImport = NewDefaultImports()
		customApplicationImport.SetImportTable(customApplicationImportTables)
		customApplicationImportTables = nil
	}
	return customApplicationImport
}
