package component

import (
	"github.com/shirou/gopsutil/mem"
	"tialloy-client-demo/model"
)

type MemCollector struct {
	MemBasicInfo   *model.MemBasicInfo
	MemRunningInfo *model.MemRunningInfo
}

func NewMemCollector() *MemCollector {
	return &MemCollector{
		MemBasicInfo:   model.NewMemBasicInfo(),
		MemRunningInfo: model.NewMemRunningInfo(),
	}
}

func (mc *MemCollector) GetBasicInfo() *model.MemBasicInfo {
	memInfo, _ := mem.VirtualMemory()
	mc.MemBasicInfo.Total = memInfo.Total
	return mc.MemBasicInfo
}

func (mc *MemCollector) GetRunningInfo() *model.MemRunningInfo {
	memInfo, _ := mem.VirtualMemory()
	mc.MemRunningInfo.Available = memInfo.Available
	mc.MemRunningInfo.Used = memInfo.Used
	mc.MemRunningInfo.UsedPercent = memInfo.UsedPercent
	return mc.MemRunningInfo
}
