package component

import (
	"ergate/model"
	"github.com/shirou/gopsutil/cpu"
)

type CpuCollector struct {
	CpuBasicInfo   *model.CpuBasicInfo
	CpuRunningInfo *model.CpuRunningInfo
}

func NewCpuCollector() *CpuCollector {
	return &CpuCollector{
		CpuBasicInfo:   model.NewCpuBasicInfo(),
		CpuRunningInfo: model.NewCpuRunningInfo(),
	}
}

func (cc *CpuCollector) GetBasicInfo() *model.CpuBasicInfo {
	cpuInfo, _ := cpu.Info()
	cc.CpuBasicInfo.ModelName = cpuInfo[0].ModelName
	cc.CpuBasicInfo.PhysicalCores, _ = cpu.Counts(false)
	cc.CpuBasicInfo.LogicalCores, _ = cpu.Counts(true)
	return cc.CpuBasicInfo
}

func (cc *CpuCollector) GetRunningInfo() *model.CpuRunningInfo {
	//cc.cpuRunningInfo.TotalPercent, _ = cpu.Percent(time.Duration(cpuCollector.duration)*time.Second, false)
	//cc.cpuRunningInfo.PerPercent, _ = cpu.Percent(time.Duration(cpuCollector.duration)*time.Second, true)
	cc.CpuRunningInfo.TotalPercent, _ = cpu.Percent(0, false)
	cc.CpuRunningInfo.PerPercent, _ = cpu.Percent(0, true)
	return cc.CpuRunningInfo
}
