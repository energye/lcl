package exec

import (
	"os"
	"path/filepath"

	"github.com/energye/lcl/tool/homedir"
)

var (
	CurrentDir string          // 当前执行目录
	Dir        string          // 执行文件所在目录
	Path       string          // 执行文件所在完整目录
	Name       string          // 执行文件名称
	HomeDir, _ = homedir.Dir() // 当前系统用户目录
)

func init() {
	CurrentDir, _ = os.Getwd()
	Path, _ = os.Executable()
	Dir, Name = filepath.Split(Path)
	HomeDir, _ = homedir.Dir()
}
