// runningprocesses.go
package runningprocesses

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
	pep         = kernel32DLL.NewProc("K32EnumProcesses")
)

// GetRunningProcessesCount returns the number of currently running processes.
func GetRunningProcessesCount() int {
	var ids [1024]uint32
	var needed uint32
	pep.Call(uintptr(unsafe.Pointer(&ids)), uintptr(len(ids)), uintptr(unsafe.Pointer(&needed)))
	return int(needed / 4)
}

// CheckRunningProcessesCount checks if the number of currently running processes is less than a specified count.
func CheckRunningProcessesCount(count int) {
	processesCount := GetRunningProcessesCount()
	//fmt.Printf("Number of running processes: %d\n", processesCount)
	if processesCount < count {
		os.Exit(-1)
	}
}
