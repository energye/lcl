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

// ICustomApplication Parent: IComponent
type ICustomApplication interface {
	IComponent
	// FindOptionIndex
	//  Extra methods.
	FindOptionIndex(S string, longopt *bool, startAt int32) int32                                                                       // function
	GetOptionValueWithString(S string) string                                                                                           // function
	GetOptionValueWithCharString(C uint16, S string) string                                                                             // function
	GetOptionValues(C uint16, S string) IStringArrayWrap                                                                                // function
	HasOptionWithString(S string) bool                                                                                                  // function
	HasOptionWithCharString(C uint16, S string) bool                                                                                    // function
	CheckOptionsWithStringStringsX3Bool(shortOptions string, longopts IStrings, opts IStrings, nonOpts IStrings, allErrors bool) string // function
	CheckOptionsWithStringStringsBool(shortOptions string, longopts IStrings, allErrors bool) string                                    // function
	CheckOptionsWithStringX2Bool(shortOptions string, longOpts string, allErrors bool) string                                           // function
	// HandleException
	//  Some Delphi methods.
	HandleException(sender IObject)                                  // procedure
	Initialize()                                                     // procedure
	Run()                                                            // procedure
	Terminate()                                                      // procedure
	TerminateWithInt(exitCode int32)                                 // procedure
	GetEnvironmentListWithStringsBool(list IStrings, namesOnly bool) // procedure
	GetEnvironmentListWithStrings(list IStrings)                     // procedure
	Log(eventType types.TEventType, msg string)                      // procedure
	// ExeName
	//  Delphi properties
	ExeName() string          // property ExeName Getter
	HelpFile() string         // property HelpFile Getter
	SetHelpFile(value string) // property HelpFile Setter
	Terminated() bool         // property Terminated Getter
	Title() string            // property Title Getter
	SetTitle(value string)    // property Title Setter
	// ConsoleApplication
	//  Extra properties
	ConsoleApplication() bool                     // property ConsoleApplication Getter
	Location() string                             // property Location Getter
	Params(index int32) string                    // property Params Getter
	ParamCount() int32                            // property ParamCount Getter
	EnvironmentVariable(envName string) string    // property EnvironmentVariable Getter
	OptionChar() uint16                           // property OptionChar Getter
	SetOptionChar(value uint16)                   // property OptionChar Setter
	CaseSensitiveOptions() bool                   // property CaseSensitiveOptions Getter
	SetCaseSensitiveOptions(value bool)           // property CaseSensitiveOptions Setter
	StopOnException() bool                        // property StopOnException Getter
	SetStopOnException(value bool)                // property StopOnException Setter
	ExceptionExitCode() int32                     // property ExceptionExitCode Getter
	SetExceptionExitCode(value int32)             // property ExceptionExitCode Setter
	EventLogFilter() types.TEventLogTypes         // property EventLogFilter Getter
	SetEventLogFilter(value types.TEventLogTypes) // property EventLogFilter Setter
	SingleInstanceEnabled() bool                  // property SingleInstanceEnabled Getter
	SetSingleInstanceEnabled(value bool)          // property SingleInstanceEnabled Setter
}

type TCustomApplication struct {
	TComponent
}

func (m *TCustomApplication) FindOptionIndex(S string, longopt *bool, startAt int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customApplicationAPI().SysCallN(1, m.Instance(), api.PasStr(S), uintptr(base.UnsafePointer(longopt)), uintptr(startAt))
	return int32(r)
}

func (m *TCustomApplication) GetOptionValueWithString(S string) string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(2, m.Instance(), api.PasStr(S))
	return api.GoStr(r)
}

func (m *TCustomApplication) GetOptionValueWithCharString(C uint16, S string) string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(3, m.Instance(), uintptr(C), api.PasStr(S))
	return api.GoStr(r)
}

func (m *TCustomApplication) GetOptionValues(C uint16, S string) IStringArrayWrap {
	if !m.IsValid() {
		return nil
	}
	r := customApplicationAPI().SysCallN(4, m.Instance(), uintptr(C), api.PasStr(S))
	return AsStringArrayWrap(r)
}

func (m *TCustomApplication) HasOptionWithString(S string) bool {
	if !m.IsValid() {
		return false
	}
	r := customApplicationAPI().SysCallN(5, m.Instance(), api.PasStr(S))
	return api.GoBool(r)
}

func (m *TCustomApplication) HasOptionWithCharString(C uint16, S string) bool {
	if !m.IsValid() {
		return false
	}
	r := customApplicationAPI().SysCallN(6, m.Instance(), uintptr(C), api.PasStr(S))
	return api.GoBool(r)
}

func (m *TCustomApplication) CheckOptionsWithStringStringsX3Bool(shortOptions string, longopts IStrings, opts IStrings, nonOpts IStrings, allErrors bool) string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(7, m.Instance(), api.PasStr(shortOptions), base.GetObjectUintptr(longopts), base.GetObjectUintptr(opts), base.GetObjectUintptr(nonOpts), api.PasBool(allErrors))
	return api.GoStr(r)
}

func (m *TCustomApplication) CheckOptionsWithStringStringsBool(shortOptions string, longopts IStrings, allErrors bool) string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(8, m.Instance(), api.PasStr(shortOptions), base.GetObjectUintptr(longopts), api.PasBool(allErrors))
	return api.GoStr(r)
}

func (m *TCustomApplication) CheckOptionsWithStringX2Bool(shortOptions string, longOpts string, allErrors bool) string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(9, m.Instance(), api.PasStr(shortOptions), api.PasStr(longOpts), api.PasBool(allErrors))
	return api.GoStr(r)
}

func (m *TCustomApplication) HandleException(sender IObject) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TCustomApplication) Initialize() {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(11, m.Instance())
}

func (m *TCustomApplication) Run() {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(12, m.Instance())
}

func (m *TCustomApplication) Terminate() {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(13, m.Instance())
}

func (m *TCustomApplication) TerminateWithInt(exitCode int32) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(14, m.Instance(), uintptr(exitCode))
}

func (m *TCustomApplication) GetEnvironmentListWithStringsBool(list IStrings, namesOnly bool) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(list), api.PasBool(namesOnly))
}

func (m *TCustomApplication) GetEnvironmentListWithStrings(list IStrings) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(list))
}

func (m *TCustomApplication) Log(eventType types.TEventType, msg string) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(17, m.Instance(), uintptr(eventType), api.PasStr(msg))
}

func (m *TCustomApplication) ExeName() string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(18, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomApplication) HelpFile() string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(19, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomApplication) SetHelpFile(value string) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(19, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomApplication) Terminated() bool {
	if !m.IsValid() {
		return false
	}
	r := customApplicationAPI().SysCallN(20, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomApplication) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(21, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomApplication) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(21, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomApplication) ConsoleApplication() bool {
	if !m.IsValid() {
		return false
	}
	r := customApplicationAPI().SysCallN(22, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomApplication) Location() string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(23, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomApplication) Params(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(24, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TCustomApplication) ParamCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customApplicationAPI().SysCallN(25, m.Instance())
	return int32(r)
}

func (m *TCustomApplication) EnvironmentVariable(envName string) string {
	if !m.IsValid() {
		return ""
	}
	r := customApplicationAPI().SysCallN(26, m.Instance(), api.PasStr(envName))
	return api.GoStr(r)
}

func (m *TCustomApplication) OptionChar() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := customApplicationAPI().SysCallN(27, 0, m.Instance())
	return uint16(r)
}

func (m *TCustomApplication) SetOptionChar(value uint16) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TCustomApplication) CaseSensitiveOptions() bool {
	if !m.IsValid() {
		return false
	}
	r := customApplicationAPI().SysCallN(28, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomApplication) SetCaseSensitiveOptions(value bool) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(28, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomApplication) StopOnException() bool {
	if !m.IsValid() {
		return false
	}
	r := customApplicationAPI().SysCallN(29, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomApplication) SetStopOnException(value bool) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(29, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomApplication) ExceptionExitCode() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customApplicationAPI().SysCallN(30, 0, m.Instance())
	return int32(r)
}

func (m *TCustomApplication) SetExceptionExitCode(value int32) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TCustomApplication) EventLogFilter() types.TEventLogTypes {
	if !m.IsValid() {
		return 0
	}
	r := customApplicationAPI().SysCallN(31, 0, m.Instance())
	return types.TEventLogTypes(r)
}

func (m *TCustomApplication) SetEventLogFilter(value types.TEventLogTypes) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TCustomApplication) SingleInstanceEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := customApplicationAPI().SysCallN(32, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomApplication) SetSingleInstanceEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	customApplicationAPI().SysCallN(32, 1, m.Instance(), api.PasBool(value))
}

// NewCustomApplication class constructor
func NewCustomApplication(owner IComponent) ICustomApplication {
	r := customApplicationAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomApplication(r)
}

func TCustomApplicationClass() types.TClass {
	r := customApplicationAPI().SysCallN(33)
	return types.TClass(r)
}

var (
	customApplicationOnce   base.Once
	customApplicationImport *imports.Imports = nil
)

func customApplicationAPI() *imports.Imports {
	customApplicationOnce.Do(func() {
		customApplicationImport = api.NewDefaultImports()
		customApplicationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomApplication_Create", 0), // constructor NewCustomApplication
			/* 1 */ imports.NewTable("TCustomApplication_FindOptionIndex", 0), // function FindOptionIndex
			/* 2 */ imports.NewTable("TCustomApplication_GetOptionValueWithString", 0), // function GetOptionValueWithString
			/* 3 */ imports.NewTable("TCustomApplication_GetOptionValueWithCharString", 0), // function GetOptionValueWithCharString
			/* 4 */ imports.NewTable("TCustomApplication_GetOptionValues", 0), // function GetOptionValues
			/* 5 */ imports.NewTable("TCustomApplication_HasOptionWithString", 0), // function HasOptionWithString
			/* 6 */ imports.NewTable("TCustomApplication_HasOptionWithCharString", 0), // function HasOptionWithCharString
			/* 7 */ imports.NewTable("TCustomApplication_CheckOptionsWithStringStringsX3Bool", 0), // function CheckOptionsWithStringStringsX3Bool
			/* 8 */ imports.NewTable("TCustomApplication_CheckOptionsWithStringStringsBool", 0), // function CheckOptionsWithStringStringsBool
			/* 9 */ imports.NewTable("TCustomApplication_CheckOptionsWithStringX2Bool", 0), // function CheckOptionsWithStringX2Bool
			/* 10 */ imports.NewTable("TCustomApplication_HandleException", 0), // procedure HandleException
			/* 11 */ imports.NewTable("TCustomApplication_Initialize", 0), // procedure Initialize
			/* 12 */ imports.NewTable("TCustomApplication_Run", 0), // procedure Run
			/* 13 */ imports.NewTable("TCustomApplication_Terminate", 0), // procedure Terminate
			/* 14 */ imports.NewTable("TCustomApplication_TerminateWithInt", 0), // procedure TerminateWithInt
			/* 15 */ imports.NewTable("TCustomApplication_GetEnvironmentListWithStringsBool", 0), // procedure GetEnvironmentListWithStringsBool
			/* 16 */ imports.NewTable("TCustomApplication_GetEnvironmentListWithStrings", 0), // procedure GetEnvironmentListWithStrings
			/* 17 */ imports.NewTable("TCustomApplication_Log", 0), // procedure Log
			/* 18 */ imports.NewTable("TCustomApplication_ExeName", 0), // property ExeName
			/* 19 */ imports.NewTable("TCustomApplication_HelpFile", 0), // property HelpFile
			/* 20 */ imports.NewTable("TCustomApplication_Terminated", 0), // property Terminated
			/* 21 */ imports.NewTable("TCustomApplication_Title", 0), // property Title
			/* 22 */ imports.NewTable("TCustomApplication_ConsoleApplication", 0), // property ConsoleApplication
			/* 23 */ imports.NewTable("TCustomApplication_Location", 0), // property Location
			/* 24 */ imports.NewTable("TCustomApplication_Params", 0), // property Params
			/* 25 */ imports.NewTable("TCustomApplication_ParamCount", 0), // property ParamCount
			/* 26 */ imports.NewTable("TCustomApplication_EnvironmentVariable", 0), // property EnvironmentVariable
			/* 27 */ imports.NewTable("TCustomApplication_OptionChar", 0), // property OptionChar
			/* 28 */ imports.NewTable("TCustomApplication_CaseSensitiveOptions", 0), // property CaseSensitiveOptions
			/* 29 */ imports.NewTable("TCustomApplication_StopOnException", 0), // property StopOnException
			/* 30 */ imports.NewTable("TCustomApplication_ExceptionExitCode", 0), // property ExceptionExitCode
			/* 31 */ imports.NewTable("TCustomApplication_EventLogFilter", 0), // property EventLogFilter
			/* 32 */ imports.NewTable("TCustomApplication_SingleInstanceEnabled", 0), // property SingleInstanceEnabled
			/* 33 */ imports.NewTable("TCustomApplication_TClass", 0), // function TCustomApplicationClass
		}
	})
	return customApplicationImport
}
