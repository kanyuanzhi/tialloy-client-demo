package model

type Info struct {
	BasicInfo   *BasicInfo   `json:"basic_info,omitempty"`
	RunningInfo *RunningInfo `json:"running_info,omitempty"`
}

type BasicInfo struct {
	HostBasicInfo *HostBasicInfo `json:"host_basic_info"`
	CpuBasicInfo  *CpuBasicInfo  `json:"cpu_basic_info"`
	MemBasicInfo  *MemBasicInfo  `json:"mem_basic_info"`
	NetBasicInfo  *NetBasicInfo  `json:"net_basic_info"`
	DiskBasicInfo *DiskBasicInfo `json:"disk_basic_info"`
}

func NewBasicInfo() *BasicInfo {
	return &BasicInfo{
		HostBasicInfo: NewHostBasicInfo(),
		CpuBasicInfo:  NewCpuBasicInfo(),
		MemBasicInfo:  NewMemBasicInfo(),
		NetBasicInfo:  NewNetBasicInfo(),
		DiskBasicInfo: NewDiskBasicInfo(),
	}
}

type RunningInfo struct {
	HostRunningInfo *HostRunningInfo `json:"host_running_info"`
	CpuRunningInfo  *CpuRunningInfo  `json:"cpu_running_info"`
	MemRunningInfo  *MemRunningInfo  `json:"mem_running_info"`
	NetRunningInfo  *NetRunningInfo  `json:"net_running_info"`
	DiskRunningInfo *DiskRunningInfo `json:"disk_running_info"`
}

func NewRunningInfo() *RunningInfo {
	return &RunningInfo{
		HostRunningInfo: NewHostRunningInfo(),
		CpuRunningInfo:  NewCpuRunningInfo(),
		MemRunningInfo:  NewMemRunningInfo(),
		NetRunningInfo:  NewNetRunningInfo(),
		DiskRunningInfo: NewDiskRunningInfo(),
	}
}
