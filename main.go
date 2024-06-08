package main

import (
	"fmt"
	"github.com/itzngga/GoDefender/AntiDebug/InternetCheck"
	"github.com/itzngga/GoDefender/AntiDebug/IsDebuggerPresent"
	processkiller "github.com/itzngga/GoDefender/AntiDebug/KillBadProcesses"
	runningprocesses "github.com/itzngga/GoDefender/AntiDebug/RunningProcesses"
	"github.com/itzngga/GoDefender/AntiDebug/pcuptime"
	kvmcheck "github.com/itzngga/GoDefender/AntiVirtualization/KVMCheck"
	triagecheck "github.com/itzngga/GoDefender/AntiVirtualization/TriageDetection"
	"github.com/itzngga/GoDefender/AntiVirtualization/USBCheck"
	usernamecheck "github.com/itzngga/GoDefender/AntiVirtualization/UsernameCheck"
	"os"

	"github.com/itzngga/GoDefender/AntiVirtualization/MonitorMetrics"
	"github.com/itzngga/GoDefender/AntiVirtualization/RecentFileActivity"
	"github.com/itzngga/GoDefender/AntiVirtualization/VMWareDetection"
	"github.com/itzngga/GoDefender/AntiVirtualization/VirtualboxDetection"

	// Anti-Debug
	"github.com/itzngga/GoDefender/AntiDebug/CheckBlacklistedWindowsNames"
	"github.com/itzngga/GoDefender/AntiDebug/RemoteDebugger"

	"github.com/itzngga/GoDefender/AntiDebug/UserAntiAntiDebug"
	// Process Related
	//"github.com/itzngga/GoDefender/Process/CriticalProcess"
)

func main() {
	/*
		ANTIDEBUG
		-----------
		- IsDebuggerPresent
		- RemoteDebugger
		- PC Uptime Check
		- Running Proccesses Count
		- Check blacklisted windows
		- KillBlacklisted Proceseses
		- Parent AntiDebug
	*/
	RecentFileActivity.RecentFileActivityCheck()
	USBCheck.PluggedIn()
	userantiantidebug.AntiAntiDebug()
	IsDebuggerPresent.IsDebuggerPresent()
	remotedebuggercheck.RemoteDebugger()
	pcuptime.CheckUptime(1200)
	runningprocesses.CheckRunningProcessesCount(50)
	blacklistcheck.CheckBlacklistedWindows()
	//parentantidebug.ParentAntiDebug()
	processkiller.KillProcesses()

	/*
		AntiVirulization
		----------------
		- Triage Check
		- VMWare Check
		- Anti KVM
		- Username Check
		-
	*/

	InternetCheck.CheckConnection()
	triagecheck.TriageCheckDebug()
	MonitorMetrics.IsScreenSmall()
	VirtualboxDetection.GraphicsCardCheck()
	fmt.Println("Debug Check: VirtualBox isnt present")
	VMWare.GraphicsCardCheck()
	fmt.Println("Debug Check: VMWare isnt present")
	if kvmcheck.CheckForKVM() {
		os.Exit(-1)
	}
	usernamecheck.CheckForBlacklistedNames()
	fmt.Println("IF YOURE HERE YOU PASSED LOL")
	/*
		EXTRA THINGS NOW:
	*/
	//programutils.SetDebugPrivilege() this is for devs who plan on continuing
	//programutils.SetProcessCritical() // this automatically gets the SeDebugPrivillige
	fmt.Scanln()
}
