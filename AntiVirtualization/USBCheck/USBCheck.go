package USBCheck

import (
	"os"
	"os/exec"
	"strings"
)

func PluggedIn() {
	usbcheckcmd := exec.Command("reg", "query", "HKLM\\SYSTEM\\ControlSet001\\Enum\\USBSTOR")
	outputusb, err := usbcheckcmd.CombinedOutput()
	if err != nil {
		return
	}
	usblines := strings.Split(string(outputusb), "\n")
	pluggedusb := 0
	for _, line := range usblines {
		if strings.TrimSpace(line) != "" {
			pluggedusb++
		}
	}
	if pluggedusb < 1 {
		if pluggedusb < 0 {
			os.Exit(-1)
		}
		return
	}
}
