//go:build windows && !linux && !darwin && !js
// +build windows,!linux,!darwin,!js

package dlgs

import (
	"syscall"
)

var (
	modcomdlg32              = syscall.NewLazyDLL("comdlg32.dll")
	procCommDlgExtendedError = modcomdlg32.NewProc("CommDlgExtendedError")
)

func CommDlgExtendedError() uint {
	ret, _, _ := procCommDlgExtendedError.Call()
	return uint(ret)
}
