package component

import (
	"github.com/shirou/gopsutil/disk"
	"tialloy-client-demo/model"
)

type DiskCollector struct {
	DiskBasicInfo   *model.DiskBasicInfo
	DiskRunningInfo *model.DiskRunningInfo
}

func NewDiskCollector() *DiskCollector {
	return &DiskCollector{
		DiskBasicInfo:   model.NewDiskBasicInfo(),
		DiskRunningInfo: model.NewDiskRunningInfo(),
	}
}

func (dc *DiskCollector) GetBasicInfo() *model.DiskBasicInfo {
	diskInfo, _ := disk.Usage(".")
	dc.DiskBasicInfo.Path = diskInfo.Path
	dc.DiskBasicInfo.Fstype = diskInfo.Fstype
	dc.DiskBasicInfo.Total = diskInfo.Total
	return dc.DiskBasicInfo
}

func (dc *DiskCollector) GetRunningInfo() *model.DiskRunningInfo {
	diskInfo, _ := disk.Usage(".")
	dc.DiskRunningInfo.Free = diskInfo.Free
	dc.DiskRunningInfo.Used = diskInfo.Used
	dc.DiskRunningInfo.UsedPercent = diskInfo.UsedPercent
	return dc.DiskRunningInfo
}
