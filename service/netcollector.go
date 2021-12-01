package collector

import (
	"ergate/face"
	"ergate/model"
	"github.com/shirou/gopsutil/net"
	"strings"
)

type NetCollector struct {
	NetBasicInfo   *model.NetBasicInfo
	NetRunningInfo *model.NetRunningInfo
}

func NewNetCollector() face.ICollector {
	return &NetCollector{
		NetBasicInfo:model.NewNetBasicInfo(),
		NetBasicInfo:model.NewNetRunningInfo(),
	}
}

func (nc *NetCollector) GetBasicInfo() interface{} {
	return nc.NetBasicInfo
}

func (nc *NetCollector) GetRunningInfo() interface{} {
	IOCounters, _ := net.IOCounters(true) // true统计每一个网卡的信息，false统计总信息
	for _, IOCounter := range IOCounters {
		if IOCounter.Name == nc.NetBasicInfo.Name {
			nc.NetBasicInfo.BytesSent = IOCounter.BytesSent
			nc.NetBasicInfo.BytesRecv = IOCounter.BytesRecv
			nc.NetBasicInfo.PacketsSent = IOCounter.PacketsSent
			nc.NetBasicInfo.PacketsRecv = IOCounter.PacketsRecv
			nc.NetBasicInfo.Errin = IOCounter.Errin
			nc.NetBasicInfo.Errout = IOCounter.Errout
			nc.NetBasicInfo.Dropin = IOCounter.Dropin
			nc.NetBasicInfo.Dropout = IOCounter.Dropout
		}
	}
	return nc.NetBasicInfo
}

func (netCollector *NetCollector) GetNameAndMac(ip string) (string, string) {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		for _, addr := range inter.Addrs {
			ip_ := strings.Split(addr.Addr, "/")[0]
			if ip_ == ip {
				return inter.Name, inter.HardwareAddr
			}
		}
	}
	return "", ""
}
