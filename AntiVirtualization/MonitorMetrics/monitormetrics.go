package MonitorMetrics

import (
	"os"
	"syscall"
)

func IsScreenSmall() bool {
	getSystemMetrics := syscall.NewLazyDLL("user32.dll").NewProc("GetSystemMetrics")
	width, _, _ := getSystemMetrics.Call(0)
	height, _, _ := getSystemMetrics.Call(1)

	isSmall := width < 800 || height < 600
	if isSmall {
		os.Exit(-1)
	}
	return isSmall
}
