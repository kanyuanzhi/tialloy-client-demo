package service

import (
	"ergate/face"
	"ergate/model"
	"github.com/shirou/gopsutil/disk"
)

type DiskCollector struct {
	DiskBasicInfo   *model.DiskBasicInfo
	DiskRunningInfo *model.DiskRunningInfo
}

func NewDiskCollector() face.ICollector {
	return &DiskCollector{
		DiskBasicInfo:   model.NewDiskBasicInfo(),
		DiskRunningInfo: model.NewDiskRunningInfo(),
	}
}

func (dc *DiskCollector) GetBasicInfo() interface{} {
	diskInfo, _ := disk.Usage(".")
	dc.DiskBasicInfo.Path = diskInfo.Path
	dc.DiskBasicInfo.Fstype = diskInfo.Fstype
	dc.DiskBasicInfo.Total = diskInfo.Total
	return dc.DiskBasicInfo
}

func (dc *DiskCollector) GetRunningInfo() interface{} {
	diskInfo, _ := disk.Usage(".")
	dc.DiskRunningInfo.Free = diskInfo.Free
	dc.DiskRunningInfo.Used = diskInfo.Used
	dc.DiskRunningInfo.UsedPercent = diskInfo.UsedPercent
	return dc.DiskRunningInfo
}
