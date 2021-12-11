package component

import (
	"tialloy-client-demo/model"
)

type Collector struct {
	Basic   *model.TerminalBasic
	Running *model.TerminalRunning

	Localhost string

	HostCollector *HostCollector
	CpuCollector  *CpuCollector
	MemCollector  *MemCollector
	NetCollector  *NetCollector
	DiskCollector *DiskCollector
}

func NewCollector() *Collector {
	return &Collector{
		Basic:   model.NewTerminalBasic(),
		Running: model.NewTerminalRunning(),

		HostCollector: NewHostCollector(),
		CpuCollector:  NewCpuCollector(),
		MemCollector:  NewMemCollector(),
		NetCollector:  NewNetCollector(),
		DiskCollector: NewDiskCollector(),
	}
}

func (c *Collector) GetBasic() *model.TerminalBasic {
	c.Basic = &model.TerminalBasic{
		HostBasic: c.HostCollector.GetBasic(),
		CpuBasic:  c.CpuCollector.GetBasic(),
		MemBasic:  c.MemCollector.GetBasic(),
		NetBasic:  c.NetCollector.GetBasic(),
		DiskBasic: c.DiskCollector.GetBasic(),
	}
	return c.Basic
}

func (c *Collector) GetRunning() *model.TerminalRunning {
	c.Running = &model.TerminalRunning{
		HostRunning: c.HostCollector.GetRunning(),
		CpuRunning:  c.CpuCollector.GetRunning(),
		MemRunning:  c.MemCollector.GetRunning(),
		NetRunning:  c.NetCollector.GetRunning(),
		DiskRunning: c.DiskCollector.GetRunning(),
	}
	return c.Running
}
