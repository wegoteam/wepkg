package file

import (
	"runtime"
)

// IsOnMacOS
// @Description: 指示应用程序是否在macOS上运行。
// @return bool
func IsOnMacOS() bool {
	return runtime.GOOS == `darwin`
}

// IsOnWindows
// @Description: 指示应用程序是否在Windows上运行。
// @return bool
func IsOnWindows() bool {
	return runtime.GOOS == `windows`
}

// IsOnLinux
// @Description: 指示应用程序是否在Linux上运行。
// @return bool
func IsOnLinux() bool {
	return runtime.GOOS == `linux`
}

// IsOn32bitArch
// @Description: 指示应用程序是否在32位架构上运行。
// @return bool
func IsOn32bitArch() bool {
	return (^uintptr(0) >> 31) == 1
}

// IsOn64bitArch
// @Description: 指示应用程序是否在64位架构上运行。
// @return bool
func IsOn64bitArch() bool {
	return (^uintptr(0) >> 63) == 1
}
