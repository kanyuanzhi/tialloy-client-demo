package component

import (
	"github.com/shirou/gopsutil/disk"
	"tialloy-client-demo/global"
	"tialloy-client-demo/model"
)

type DiskCollector struct {
	Basic   *model.TerminalDiskBasic
	Running *model.TerminalDiskRunning
	Path    string
}

func NewDiskCollector() *DiskCollector {
	var diskPath string
	if global.OS == "windows" {
		diskPath = global.Object.WindowsPath
	} else {
		diskPath = "/"
	}
	return &DiskCollector{
		Basic:   model.NewTerminalDiskBasic(),
		Running: model.NewTerminalDiskRunning(),
		Path:    diskPath,
	}
}

func (dc *DiskCollector) GetBasic() *model.TerminalDiskBasic {
	diskInfo, _ := disk.Usage(dc.Path)
	dc.Basic.Path = diskInfo.Path
	dc.Basic.Fstype = diskInfo.Fstype
	dc.Basic.Total = diskInfo.Total
	return dc.Basic
}

func (dc *DiskCollector) GetRunning() *model.TerminalDiskRunning {
	diskInfo, _ := disk.Usage(dc.Path)
	dc.Running.Free = diskInfo.Free
	dc.Running.Used = diskInfo.Used
	dc.Running.UsedPercent = diskInfo.UsedPercent
	return dc.Running
}
