@Echo off
title DOWNLOADING MODULES
go mod init GoDefender
go get github.com/itzngga/GoDefender@v1.0.5
go get github.com/itzngga/GoDefender/AntiVirtualization/TriageDetection
go get github.com/itzngga/GoDefender/AntiVirtualization/MonitorMetrics
go get github.com/itzngga/GoDefender/AntiVirtualization/VirtualboxDetection
go get github.com/itzngga/GoDefender/AntiVirtualization/VMWareDetection
go get github.com/itzngga/GoDefender/AntiVirtualization/KVMCheck
go get github.com/itzngga/GoDefender/AntiVirtualization/UsernameCheck
go get github.com/itzngga/GoDefender/AntiDebug/IsDebuggerPresent
go get github.com/itzngga/GoDefender/AntiDebug/RemoteDebugger
go get github.com/itzngga/GoDefender/AntiDebug/pcuptime
go get github.com/itzngga/GoDefender/AntiDebug/CheckBlacklistedWindowsNames
go get github.com/itzngga/GoDefender/AntiDebug/RunningProcesses
go get github.com/itzngga/GoDefender/AntiDebug/ParentAntiDebug
go get github.com/itzngga/GoDefender/AntiDebug/KillBadProcesses
go get github.com/itzngga/GoDefender/Process/CriticalProcess
go get github.com/itzngga/GoDefender/AntiDebug/UserAntiAntiDebug
pause
