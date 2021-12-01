package service

import (
	"ergate/face"
	"ergate/model"
	"github.com/shirou/gopsutil/mem"
)

type MemCollector struct {
	MemBasicInfo   *model.MemBasicInfo
	MemRunningInfo *model.MemRunningInfo
}

func NewMemCollector() face.ICollector {
	return &MemCollector{
		MemBasicInfo:   model.NewMemBasicInfo(),
		MemRunningInfo: model.NewMemRunningInfo(),
	}
}

func (mc *MemCollector) GetBasicInfo() interface{} {
	memInfo, _ := mem.VirtualMemory()
	mc.MemBasicInfo.Total = memInfo.Total
	return mc.MemBasicInfo
}

func (mc *MemCollector) GetRunningInfo() interface{} {
	memInfo, _ := mem.VirtualMemory()
	mc.MemRunningInfo.Available = memInfo.Available
	mc.MemRunningInfo.Used = memInfo.Used
	mc.MemRunningInfo.UsedPercent = memInfo.UsedPercent
	return mc.MemRunningInfo
}
