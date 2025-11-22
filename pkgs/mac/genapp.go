//go:build darwin && !prod
// +build darwin,!prod

package mac

import "github.com/energye/lcl/internal/initialize/macapp/gen"

func GenApp(exeFileDir, cefSubProcess string) {
	gen.GenMacOSApp(exeFileDir, cefSubProcess)
}
