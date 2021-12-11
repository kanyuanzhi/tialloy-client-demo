package component

import (
	"github.com/shirou/gopsutil/cpu"
	"tialloy-client-demo/model"
)

type CpuCollector struct {
	Basic   *model.TerminalCpuBasic
	Running *model.TerminalCpuRunning
}

func NewCpuCollector() *CpuCollector {
	return &CpuCollector{
		Basic:   model.NewTerminalCpuBasic(),
		Running: model.NewTerminalCpuRunning(),
	}
}

func (cc *CpuCollector) GetBasic() *model.TerminalCpuBasic {
	cpuInfo, _ := cpu.Info()
	cc.Basic.ModelName = cpuInfo[0].ModelName
	cc.Basic.PhysicalCores, _ = cpu.Counts(false)
	cc.Basic.LogicalCores, _ = cpu.Counts(true)
	return cc.Basic
}

func (cc *CpuCollector) GetRunning() *model.TerminalCpuRunning {
	//cc.cpuRunningInfo.TotalPercent, _ = cpu.Percent(time.Duration(cpuCollector.duration)*time.Second, false)
	//cc.cpuRunningInfo.PerPercent, _ = cpu.Percent(time.Duration(cpuCollector.duration)*time.Second, true)
	cc.Running.TotalPercent, _ = cpu.Percent(0, false)
	cc.Running.PerPercent, _ = cpu.Percent(0, true)
	return cc.Running
}
