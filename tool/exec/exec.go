package exec

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/energye/lcl/tool/homedir"
)

var (
	CurrentDir string // 当前执行目录
	Dir        string // 执行文件所在目录
	Path       string // 执行文件所在完整目录
	Name       string // 执行文件名称
	HomeDir    string // 当前系统用户目录
)

func init() {
	var err error
	CurrentDir, err = os.Getwd()
	Path, err = os.Executable()
	Dir, Name = filepath.Split(Path)
	HomeDir, err = os.UserHomeDir()
	if err != nil {
		HomeDir, err = homedir.Dir()
	}
}

// AppDir 返回应用所在目录
func AppDir() string {
	tempDir := Dir
	if runtime.GOOS != "darwin" {
		return tempDir
	}
	if strings.Contains(tempDir, ".app/Contents/MacOS") {
		return filepath.Join(tempDir, "../../../")
	}
	return tempDir
}
