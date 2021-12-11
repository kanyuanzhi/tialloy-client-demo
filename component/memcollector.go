package component

import (
	"github.com/shirou/gopsutil/mem"
	"tialloy-client-demo/model"
)

type MemCollector struct {
	Basic   *model.TerminalMemBasic
	Running *model.TerminalMemRunning
}

func NewMemCollector() *MemCollector {
	return &MemCollector{
		Basic:   model.NewTerminalMemBasic(),
		Running: model.NewTerminalMemRunning(),
	}
}

func (mc *MemCollector) GetBasic() *model.TerminalMemBasic {
	memInfo, _ := mem.VirtualMemory()
	mc.Basic.Total = memInfo.Total
	return mc.Basic
}

func (mc *MemCollector) GetRunning() *model.TerminalMemRunning {
	memInfo, _ := mem.VirtualMemory()
	mc.Running.Available = memInfo.Available
	mc.Running.Used = memInfo.Used
	mc.Running.UsedPercent = memInfo.UsedPercent
	return mc.Running
}
