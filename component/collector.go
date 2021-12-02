package component

import (
	"ergate/model"
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
		BasicInfo:   model.NewBasicInfo(),
		RunningInfo: model.NewRunningInfo(),

		HostCollector: NewHostCollector(),
		CpuCollector:  NewCpuCollector(),
		MemCollector:  NewMemCollector(),
		NetCollector:  NewNetCollector(),
		DiskCollector: NewDiskCollector(),
	}
}

func (c *Collector) GetBasicInfo() *model.BasicInfo {
	c.BasicInfo.HostBasicInfo = c.HostCollector.GetBasicInfo()
	c.BasicInfo.CpuBasicInfo = c.CpuCollector.GetBasicInfo()
	c.BasicInfo.MemBasicInfo = c.MemCollector.GetBasicInfo()
	c.BasicInfo.NetBasicInfo = c.NetCollector.GetBasicInfo()
	c.BasicInfo.DiskBasicInfo = c.DiskCollector.GetBasicInfo()
	return c.BasicInfo
}

func (c *Collector) GetRunningInfo() *model.RunningInfo {
	c.RunningInfo.HostRunningInfo = c.HostCollector.GetRunningInfo()
	c.RunningInfo.CpuRunningInfo = c.CpuCollector.GetRunningInfo()
	c.RunningInfo.MemRunningInfo = c.MemCollector.GetRunningInfo()
	c.RunningInfo.NetRunningInfo = c.NetCollector.GetRunningInfo()
	c.RunningInfo.DiskRunningInfo = c.DiskCollector.GetRunningInfo()
	return c.RunningInfo
}
