package component

import (
	"github.com/shirou/gopsutil/host"
	"tialloy-client-demo/model"
)

type HostCollector struct {
	HostBasicInfo   *model.HostBasicInfo
	HostRunningInfo *model.HostRunningInfo
}

func NewHostCollector() *HostCollector {
	return &HostCollector{
		HostBasicInfo:   model.NewHostBasicInfo(),
		HostRunningInfo: model.NewHostRunningInfo(),
	}
}

func (hc *HostCollector) GetBasicInfo() *model.HostBasicInfo {
	hostInfo, _ := host.Info()
	hc.HostBasicInfo.Hostname = hostInfo.Hostname
	hc.HostBasicInfo.OS = hostInfo.OS
	hc.HostBasicInfo.Platform = hostInfo.Platform
	hc.HostBasicInfo.PlatformFamily = hostInfo.PlatformFamily
	hc.HostBasicInfo.PlatformVersion = hostInfo.PlatformVersion
	hc.HostBasicInfo.KernelVersion = hostInfo.KernelVersion
	hc.HostBasicInfo.KernelArch = hostInfo.KernelArch
	userInfo, err := host.Users()
	if err != nil {
		hc.HostBasicInfo.User = "None"
	} else {
		hc.HostBasicInfo.User = userInfo[0].User
	}
	return hc.HostBasicInfo
}

func (hc *HostCollector) GetRunningInfo() *model.HostRunningInfo {
	hostInfo, _ := host.Info()
	hc.HostRunningInfo.Uptime = hostInfo.Uptime
	hc.HostRunningInfo.Procs = hostInfo.Procs
	return hc.HostRunningInfo
}
