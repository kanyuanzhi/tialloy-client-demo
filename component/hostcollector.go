package component

import (
	"github.com/shirou/gopsutil/host"
	"tialloy-client-demo/model"
)

type HostCollector struct {
	Basic   *model.TerminalHostBasic
	Running *model.TerminalHostRunning
}

func NewHostCollector() *HostCollector {
	return &HostCollector{
		Basic:   model.NewTerminalHostBasic(),
		Running: model.NewTerminalHostRunning(),
	}
}

func (hc *HostCollector) GetBasic() *model.TerminalHostBasic {
	hostInfo, _ := host.Info()
	hc.Basic.Hostname = hostInfo.Hostname
	hc.Basic.BootTime = hostInfo.BootTime
	hc.Basic.OS = hostInfo.OS
	hc.Basic.Platform = hostInfo.Platform
	hc.Basic.PlatformFamily = hostInfo.PlatformFamily
	hc.Basic.PlatformVersion = hostInfo.PlatformVersion
	hc.Basic.KernelVersion = hostInfo.KernelVersion
	hc.Basic.KernelArch = hostInfo.KernelArch
	userInfo, err := host.Users()
	if err != nil {
		hc.Basic.User = "None"
	} else {
		hc.Basic.User = userInfo[0].User
	}
	return hc.Basic
}

func (hc *HostCollector) GetRunning() *model.TerminalHostRunning {
	hostInfo, _ := host.Info()
	hc.Running.Uptime = hostInfo.Uptime
	hc.Running.Procs = hostInfo.Procs
	return hc.Running
}
