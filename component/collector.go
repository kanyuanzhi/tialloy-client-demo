package component

import (
	"tialloy-client-demo/model"
)

type Collector struct {
	BasicInfo   *model.BasicInfo
	RunningInfo *model.RunningInfo

	Localhost string

	HostCollector *HostCollector
	CpuCollector  *CpuCollector
	MemCollector  *MemCollector
	NetCollector  *NetCollector
	DiskCollector *DiskCollector
}

func NewCollector() *Collector {
	return &Collector{
		BasicInfo:   &model.BasicInfo{},
		RunningInfo: &model.RunningInfo{},

		HostCollector: NewHostCollector(),
		CpuCollector:  NewCpuCollector(),
		MemCollector:  NewMemCollector(),
		NetCollector:  NewNetCollector(),
		DiskCollector: NewDiskCollector(),
	}
}

func (c *Collector) GetBasicInfo() *model.BasicInfo {
	c.BasicInfo.Data = &model.BasicInfoData{
		HostBasicInfo: c.HostCollector.GetBasicInfo(),
		CpuBasicInfo:  c.CpuCollector.GetBasicInfo(),
		MemBasicInfo:  c.MemCollector.GetBasicInfo(),
		NetBasicInfo:  c.NetCollector.GetBasicInfo(),
		DiskBasicInfo: c.DiskCollector.GetBasicInfo(),
	}
	return c.BasicInfo
}

func (c *Collector) GetRunningInfo() *model.RunningInfo {
	c.RunningInfo.Data = &model.RunningInfoData{
		HostRunningInfo: c.HostCollector.GetRunningInfo(),
		CpuRunningInfo:  c.CpuCollector.GetRunningInfo(),
		MemRunningInfo:  c.MemCollector.GetRunningInfo(),
		NetRunningInfo:  c.NetCollector.GetRunningInfo(),
		DiskRunningInfo: c.DiskCollector.GetRunningInfo(),
	}
	return c.RunningInfo
}
