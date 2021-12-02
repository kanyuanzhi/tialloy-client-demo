package component

import (
	"ergate/model"
	"github.com/shirou/gopsutil/net"
	"strings"
)

type NetCollector struct {
	NetBasicInfo   *model.NetBasicInfo
	NetRunningInfo *model.NetRunningInfo

	IP string
}

func NewNetCollector() *NetCollector {
	return &NetCollector{
		NetBasicInfo:   model.NewNetBasicInfo(),
		NetRunningInfo: model.NewNetRunningInfo(),
	}
}

func (nc *NetCollector) GetBasicInfo() *model.NetBasicInfo {
	nc.NetBasicInfo.IP = nc.IP
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		for _, addr := range inter.Addrs {
			ip_ := strings.Split(addr.Addr, "/")[0]
			if ip_ == nc.IP {
				nc.NetBasicInfo.Name = inter.Name
				nc.NetBasicInfo.Mac = inter.HardwareAddr
			}
		}
	}
	return nc.NetBasicInfo
}

func (nc *NetCollector) GetRunningInfo() *model.NetRunningInfo {
	IOCounters, _ := net.IOCounters(true) // true统计每一个网卡的信息，false统计总信息
	for _, IOCounter := range IOCounters {
		if IOCounter.Name == nc.NetBasicInfo.Name {
			nc.NetRunningInfo.BytesSent = IOCounter.BytesSent
			nc.NetRunningInfo.BytesRecv = IOCounter.BytesRecv
			nc.NetRunningInfo.PacketsSent = IOCounter.PacketsSent
			nc.NetRunningInfo.PacketsRecv = IOCounter.PacketsRecv
			nc.NetRunningInfo.Errin = IOCounter.Errin
			nc.NetRunningInfo.Errout = IOCounter.Errout
			nc.NetRunningInfo.Dropin = IOCounter.Dropin
			nc.NetRunningInfo.Dropout = IOCounter.Dropout
		}
	}
	return nc.NetRunningInfo
}

func (nc *NetCollector) SetIP(ip string) {
	nc.IP = ip
}
