package macapp

// MacApp
//
//  1. macos调式时临时创建一个符合macapp的程序包
//
//  2. 如果基于cef，需要指定cef frameworks 根目录【/homt/xxx/cef_binary_xxxxxxx_macosx64/Release】
var MacApp = &macApp{}

const (
	envDev = "dev" //MacOS ENERGY的开发环境常量配置
)

type macApp struct {
	execName             string
	execFile             string
	macContentsDir       string
	macOSDir             string
	macResources         string
	lclLibFileName       string
	plistFileName        string
	pkgInfoFileName      string
	macAppFrameworksDir  string
	isCEF                bool
	isMain               bool
	cefFrameworksDir     string
	browseSubprocessPath string
	lsUIElement          string
}
